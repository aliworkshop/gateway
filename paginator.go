package gateway

type Paginator interface {
	GetPageSize() int
	GetPage() int
	Total() uint64
}

type IPaginator interface {
	Paginator
	SetPage(int)
	SetPageSize(int)
	SetTotal(uint64)
}

type paginator struct {
	limit int
	page  int
	total uint64
}

func NewPaginator() IPaginator {
	return &paginator{}
}

func (p *paginator) SetPageSize(limit int) {
	p.limit = limit
}

func (p *paginator) SetPage(page int) {
	p.page = page
}

func (p *paginator) GetPageSize() int {
	if p.limit == 0 {
		p.limit = 10
	}
	return p.limit
}

func (p *paginator) GetPage() int {
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

func (p *paginator) Total() uint64 {
	return p.total
}

func (p *paginator) SetTotal(total uint64) {
	p.total = total
}
