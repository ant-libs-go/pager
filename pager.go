/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2021-01-26 17:25:29
# File Name: pager.go
# Description:
####################################################################### */

package pager

const (
	DefaultPageSize = 20
)

type Pager struct {
	count    int32
	page     int32
	pageSize int32
}

type RenderData struct {
	CurPage   int32 `json:"cur_page"`
	PageSize  int32 `json:"page_size"`
	ItemCount int32 `json:"item_count"`
	PageCount int32 `json:"page_count"`
}

func New(count, page, pageSize int32) (r *Pager) {
	o := &Pager{count: count, page: page, pageSize: pageSize}
	if o.count < 0 {
		o.count = 0
	}
	if o.page <= 1 {
		o.page = 1
	}
	if o.pageSize <= 1 {
		o.pageSize = DefaultPageSize
	}
	return o
}

func (this *Pager) ItemCount() (r int32) {
	return this.count
}

func (this *Pager) PageCount() (r int32) {
	return (this.count + this.pageSize - 1) / this.pageSize
}

func (this *Pager) Offset() (r int32) {
	return (this.page - 1) * this.pageSize
}

func (this *Pager) CurPage() (r int32) {
	return this.page
}

func (this *Pager) PageSize() (r int32) {
	return this.pageSize
}

func (this *Pager) Render() (r *RenderData) {
	return &RenderData{
		CurPage:   this.CurPage(),
		PageSize:  this.pageSize(),
		ItemCount: this.ItemCount(),
		PageCount: this.PageCount()}
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
