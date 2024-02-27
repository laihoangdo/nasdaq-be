package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaginationQuery_SetOrderBy(t *testing.T) {
	tests := map[string]struct {
		input      string
		expOrderBy []OrderBy
		expErr     error
	}{
		"success": {
			input: "id|asc,name|asc,date|desc",
			expOrderBy: []OrderBy{
				{
					Column: "id",
					Order:  "ASC",
				},
				{
					Column: "name",
					Order:  "ASC",
				},
				{
					Column: "date",
					Order:  "DESC",
				},
			},
			expErr: nil,
		},
		"success when order is case-insensitive": {
			input: "id|aSc,name|AsC,date|dESc",
			expOrderBy: []OrderBy{
				{
					Column: "id",
					Order:  "ASC",
				},
				{
					Column: "name",
					Order:  "ASC",
				},
				{
					Column: "date",
					Order:  "DESC",
				},
			},
			expErr: nil,
		},
		"success when order by is empty": {
			input: "",
			expOrderBy: []OrderBy{
				{
					Column: "created_at",
					Order:  "ASC",
				},
			},
			expErr: nil,
		},
		"error when attributes are not splitted by comma": {
			input:      "id|aSc name|AsC date|dESc",
			expOrderBy: []OrderBy{},
			expErr:     errors.New("invalid order_by format"),
		},
	}
	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			// Given:
			q := &PaginationQuery{}

			// When:j
			err := q.SetOrderBy(tc.input)

			// Then:
			if tc.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.expOrderBy, q.OrderBy)
		})
	}
}

func TestPaginationQuery_SetSize(t *testing.T) {
	tests := map[string]struct {
		input           string
		paginationQuery *PaginationQuery
		expErr          error
	}{
		"success when size is omitted": {
			input:           "",
			paginationQuery: &PaginationQuery{Size: 0},
			expErr:          nil,
		},
		"success when size is gte 0": {
			input:           "3",
			paginationQuery: &PaginationQuery{Size: 3},
			expErr:          nil,
		},
		"errors when size is negative": {
			input:           "-3",
			paginationQuery: &PaginationQuery{},
			expErr:          errors.New("invalid page size"),
		},
		"error when size is not int": {
			input:           "abc",
			paginationQuery: &PaginationQuery{},
			expErr:          errors.New("invalid page size"),
		},
	}

	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			// Given:
			q := &PaginationQuery{}

			// When:
			err := q.SetSize(tc.input)

			// Then:
			if tc.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.paginationQuery, q)
		})
	}
}

func TestPaginationQuery_SetPage(t *testing.T) {
	tests := map[string]struct {
		input           string
		paginationQuery *PaginationQuery
		expErr          error
	}{
		"success": {
			input:           "5",
			paginationQuery: &PaginationQuery{Page: 5},
			expErr:          nil,
		},
		"success when page is omitted": {
			input:           "",
			paginationQuery: &PaginationQuery{Page: 1},
			expErr:          nil,
		},
		"error when page is not int": {
			input:           "orders",
			paginationQuery: &PaginationQuery{},
			expErr:          errors.New("invalid page number"),
		},
		"error when page is negative": {
			input:           "-10",
			paginationQuery: &PaginationQuery{},
			expErr:          errors.New("invalid page number"),
		},
	}

	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			// Given:
			q := &PaginationQuery{}

			// When:
			err := q.SetPage(tc.input)

			// Then:
			if tc.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tc.paginationQuery, q)
		})
	}
}
