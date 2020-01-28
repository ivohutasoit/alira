package util

import (
	"math"

	"github.com/jinzhu/gorm"
)

type PaginatorParameter struct {
	Database *gorm.DB
	Count    interface{}
	Result   interface{}
	Page     int
	Limit    int
	OrderBy  []string
	ShowSQL  bool
}

type Paginator struct {
	TotalRecord  int         `json:"total_record"`
	TotalPage    int         `json:"total_page"`
	Records      interface{} `json:"records"`
	Offset       int         `json:"offset"`
	Limit        int         `json:"limit"`
	CurrentPage  int         `json:"current_page"`
	PreviousPage int         `json:"previous_page"`
	NextPage     int         `json:"next_page"`
}

func NewPaginator(p *PaginatorParameter) *Paginator {
	database := p.Database

	if p.ShowSQL {
		database = database.Debug()
	}

	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 10
	}

	if len(p.OrderBy) > 0 {
		for _, v := range p.OrderBy {
			database = database.Order(v)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count, offset int

	go countRecords(database, p.Count, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	database.Limit(p.Limit).Offset(offset).Scan(p.Result)
	<-done

	paginator.TotalRecord = count
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))
	paginator.CurrentPage = p.Page
	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.Records = p.Result

	if p.Page > 1 {
		paginator.PreviousPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}

	return &paginator
}

func countRecords(database *gorm.DB, out interface{}, done chan bool, count *int) {
	database.Scan(out).Count(count)
	done <- true
}
