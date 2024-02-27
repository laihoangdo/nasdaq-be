package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"nasdaqvfs/internal/models"
	pkgErr "nasdaqvfs/pkg/errors"

	"github.com/gin-gonic/gin"
)

const (
	defaultPageNumber       = 1
	defaultPageOffset       = 0
	defaultDelimiter        = ","
	defaultOrderByDelimiter = "|"
	pageSizeGetAll          = 0

	OrderByAsc  = "ASC"
	OrderByDesc = "DESC"
)

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
// Allow size = 0 or "" (get all)
func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		return nil
	}

	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return errors.New("invalid page size")
	}

	if n < 0 {
		return errors.New("invalid page size")
	}

	if n > 0 {
		q.Size = n
	}

	return nil
}

// Set page number
// Allow page = 0 or "" (get all)
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
		q.OrderBy = []OrderBy{
			{
				Column: "created_at",
				Order:  OrderByAsc,
			},
		}
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

	if err := q.SetPage(c.Query("page")); err != nil {
		errs.Add(pkgErr.NewValidationError("page", err.Error()))
	}

	if err := q.SetSize(c.Query("size")); err != nil {
		errs.Add(pkgErr.NewValidationError("size", err.Error()))
	}

	if err := q.SetOrderBy(c.Query("order_by")); err != nil {
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

// CheckValidAttributes checks if pagination fields are existed in database table
func (q PaginationQuery) CheckValidAttributes(entityAttributes []string) error {
	return q.CheckValidOrderByAttributes(entityAttributes)
}

// CheckValidOrderByAttributes checks if pagination order by existed fields
func (q PaginationQuery) CheckValidOrderByAttributes(entityAttributes []string) error {
	errs := pkgErr.NewValidationErrorsCollector()

	for _, orderBy := range q.OrderBy {
		if !Contains(entityAttributes, orderBy.Column) {
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

func NewPaginationResponse(pq PaginationQuery, count int64) *models.Pagination {
	return &models.Pagination{
		TotalCount: count,
		TotalPages: GetTotalPages(count, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    GetHasMore(pq.GetPage(), count, pq.GetSize()),
	}
}
