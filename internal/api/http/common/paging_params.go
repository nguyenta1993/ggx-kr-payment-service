package httpcommon

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paging struct {
	Offset     int64 `json:"offset"`
	Limit      int64 `json:"limit"`
	TotalCount int64 `json:"totalCount"`
}

const pageSize = "pageSize"

const currentPage = "currentPage"

func ParseParams(c *gin.Context) *Paging {
	var p Paging

	if pageSize, err := strconv.Atoi(c.Query(pageSize)); err == nil && pageSize > 0 {
		p.Limit = int64(pageSize)
	}

	if currentPage, err := strconv.Atoi(c.Query(currentPage)); err == nil && currentPage > 0 {
		p.Offset = int64(currentPage-1) * p.Limit
	}

	return &p
}

type Filters map[string]string

func ParseFilters(c *gin.Context) Filters {
	filters := Filters{}
	for key, value := range c.Request.URL.Query() {
		if key != pageSize && key != currentPage {
			filters[key] = value[0]
		}
	}
	return filters
}

type SortDirection string

const (
	Ascend  SortDirection = "ascend"
	Descend SortDirection = "descend"
)

const (
	SortDirectionAscend  = "asc"
	SortDirectionDescend = "desc"
)
