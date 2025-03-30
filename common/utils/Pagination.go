package utils

// Pagination 增强版分页结构体
type Pagination struct {
	Page      int   `form:"page" json:"page"`           // 当前页码
	PageSize  int   `form:"page_size" json:"page_size"` // 每页数量
	Total     int64 `json:"total,omitempty"`            // 总记录数
	TotalPage int   `json:"total_page,omitempty"`       // 总页数
}

// 计算偏移量
func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// 获取每页数量
func (p *Pagination) Limit() int {
	return p.PageSize
}

// 默认分页参数
func DefaultPagination() Pagination {
	return Pagination{
		Page:     1,
		PageSize: 10,
	}
}

// 验证分页参数
func (p *Pagination) Validate() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	// 设置最大每页数量限制
	if p.PageSize > 100 {
		p.PageSize = 100
	}
}

// 计算总页数
func (p *Pagination) SetTotal(total int64) {
	p.Total = total
	if p.PageSize > 0 {
		p.TotalPage = int(total) / p.PageSize
		if int(total)%p.PageSize > 0 {
			p.TotalPage++
		}
	}
}

// 是否是第一页
func (p *Pagination) IsFirstPage() bool {
	return p.Page <= 1
}

// 是否是最后一页
func (p *Pagination) IsLastPage() bool {
	if p.TotalPage == 0 {
		return true
	}
	return p.Page >= p.TotalPage
}
