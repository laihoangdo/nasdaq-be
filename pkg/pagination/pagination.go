package pagination

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	pkgErr "nasdaqvfs/pkg/errors"

	"github.com/gin-gonic/gin"
)

const (
	defaultPageSize         = 10
	defaultPageNumber       = 1
	defaultPageOffset       = 0
	defaultDelimiter        = ","
	defaultOrderByDelimiter = "|"
	pageSizeGetAll          = 0

	OrderByAsc  = "ASC"
	OrderByDesc = "DESC"
)

type PaginationResponse struct {
	TotalCount *int64 `json:"total_count,omitempty"`
	TotalPages *int   `json:"total_pages,omitempty"`
	Page       *int   `json:"page,omitempty"`
	Size       *int   `json:"size,omitempty"`
	HasMore    *bool  `json:"has_more,omitempty"`
}

// Pagination
type Pagination struct {
	Total int64
	Page  int
	Size  int
}

// IsAll checks if the pagination is for getting all
func (p Pagination) IsAll() bool {
	return p.Size == pageSizeGetAll
}

// ToResponse transforms to response struct
func (p Pagination) ToResponse() PaginationResponse {
	if p.IsAll() {
		return PaginationResponse{}
	}

	totalPages := p.GetTotalPages()
	hasMore := p.HasMore()

	return PaginationResponse{
		TotalCount: &p.Total,
		TotalPages: &totalPages,
		Page:       &p.Page,
		Size:       &p.Size,
		HasMore:    &hasMore,
	}
}

// GetTotalPages get total pages
func (p Pagination) GetTotalPages() int {
	if p.Size == 0 {
		return 0
	}

	return int(math.Ceil(float64(p.Total) / float64(p.Size)))
}

// HasMore check if has more page
func (p Pagination) HasMore() bool {
	return p.Page < p.GetTotalPages()
}

// Pagination query params
type PaginationQuery struct {
	Size       int         `json:"size,omitempty"`
	Page       int         `json:"page,omitempty"`
	OrderBy    []OrderBy   `json:"-"`
	Filter     interface{} `json:"-"`
	TotalCount int64       `json:"-"`
}

type OrderBy struct {
	Column string `json:"-"`
	Order  string `json:"-"`
}

// Set page size
// Allow size="0" (get all)
// size="" (default page size = 10)
func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultPageSize
		return nil
	}

	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return errors.New("invalid page size")
	}

	if n == 0 {
		q.Size = pageSizeGetAll
		return nil
	}

	if n < 0 {
		return errors.New("invalid page size")
	}

	if n > 0 {
		q.Size = n
	}

	return nil
}

// SetPage Set page number
// If page = 0 or "", setting page = defaultPageNumber
func (q *PaginationQuery) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Page = defaultPageNumber
		return nil
	}

	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return errors.New("invalid page number")
	}

	if n < 0 {
		return errors.New("invalid page number")
	}

	if n == 0 {
		q.Page = defaultPageNumber
	} else {
		q.Page = n
	}

	return nil
}

// Set order by
func (q *PaginationQuery) SetOrderBy(orderByQuery string) error {
	if strings.TrimSpace(orderByQuery) == "" {
		q.OrderBy = []OrderBy{}
		return nil
	}

	listOrderBy := strings.Split(orderByQuery, defaultDelimiter)

	if q.OrderBy == nil {
		q.OrderBy = []OrderBy{}
	}

	for _, orderBy := range listOrderBy {
		ele := strings.Split(orderBy, defaultOrderByDelimiter)
		if len(ele) != 2 {
			return errors.New("invalid order_by format")
		}

		ele[1] = strings.ToUpper(strings.TrimSpace(ele[1]))
		if ele[1] != OrderByAsc && ele[1] != OrderByDesc {
			return errors.New("invalid order_by param")
		}

		q.OrderBy = append(q.OrderBy, OrderBy{
			Column: strings.ToLower(strings.TrimSpace(ele[0])),
			Order:  ele[1],
		})
	}

	return nil
}

// Set page filter
func (q *PaginationQuery) SetFilter(val interface{}) {
	q.Filter = val
}

// Set total count
func (q *PaginationQuery) SetTotalCount(totalCount int64) {
	q.TotalCount = totalCount
}

// Get offset
func (q *PaginationQuery) GetOffset() int {
	return (q.Page - 1) * q.Size
}

// Get limit
func (q *PaginationQuery) GetLimit() int {
	return q.Size
}

// Get OrderBy
func (q *PaginationQuery) GetOrderBy() []OrderBy {
	return q.OrderBy
}

// Get OrderBy
func (q *PaginationQuery) GetPage() int {
	return q.Page
}

// Get OrderBy
func (q *PaginationQuery) GetSize() int {
	return q.Size
}

func (q *PaginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", q.GetPage(), q.GetSize(), q.GetOrderBy())
}

// Get pagination query struct from
func GetPaginationFromCtx(c *gin.Context) (PaginationQuery, error) {
	q := PaginationQuery{}
	errs := pkgErr.NewValidationErrorsCollector()

	if err := q.SetPage(strings.TrimSpace(c.Query("page"))); err != nil {
		errs.Add(pkgErr.NewValidationError("page", err.Error()))
	}

	if err := q.SetSize(strings.TrimSpace(c.Query("size"))); err != nil {
		errs.Add(pkgErr.NewValidationError("size", err.Error()))
	}

	if err := q.SetOrderBy(strings.TrimSpace(c.Query("order_by"))); err != nil {
		errs.Add(pkgErr.NewValidationError("order_by", err.Error()))
	}

	if errs.HasError() {
		return PaginationQuery{}, errs
	}

	return q, nil
}

// Get total pages int
func (q PaginationQuery) GetTotalPages() int {
	if q.TotalCount == 0 {
		return 0
	}
	if q.Size == 0 {
		return 1
	}
	d := float64(q.TotalCount) / float64(q.Size)
	return int(math.Ceil(d))
}

// Get has more
func (q PaginationQuery) GetHasMore() bool {
	if !q.IsPaginate() {
		return false
	}
	return q.TotalCount-int64(q.Page*q.Size) > 0
}

// ValidateAttributes checks if pagination fields are existed in database table
func (q PaginationQuery) ValidateAttributes(entityAttributes []string) error {
	return q.ValidateOrderByAttributes(entityAttributes)
}

func contains(arrStr []string, s string) bool {
	for idx := range arrStr {
		if arrStr[idx] == s {
			return true
		}
	}
	return false
}

// ValidateOrderByAttributes checks if pagination order by existed fields
func (q PaginationQuery) ValidateOrderByAttributes(entityAttributes []string) error {
	errs := pkgErr.NewValidationErrorsCollector()

	for _, orderBy := range q.OrderBy {
		if !contains(entityAttributes, orderBy.Column) {
			errs.Add(pkgErr.NewValidationError(orderBy.Column, fmt.Sprintf("invalid attribute %s", orderBy.Column)))
		}
	}

	if errs.HasError() {
		return errs
	}

	return nil
}

// IsGetAll checks if needing get all items
func (q PaginationQuery) IsPaginate() bool {
	return q.Size != pageSizeGetAll
}

// Get total pages int
func GetTotalPages(totalCount int64, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)
	return int(math.Ceil(d))
}

// Get has more
func GetHasMore(currentPage int, totalCount int64, pageSize int) bool {
	return currentPage < int(totalCount)/pageSize
}
