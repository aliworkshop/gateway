package gateway

type Paginator interface {
	GetPageSize() int32
	GetPage() int32
	Total() uint64
}

type IPaginator interface {
	Paginator
	SetPage(int32)
	SetPageSize(int32)
	SetTotal(uint64)
}

type paginator struct {
	limit int32
	page  int32
	total uint64
}

func NewPaginator() IPaginator {
	return &paginator{}
}

func (p *paginator) SetPageSize(limit int32) {
	p.limit = limit
}

func (p *paginator) SetPage(page int32) {
	p.page = page
}

func (p *paginator) GetPageSize() int32 {
	if p.limit == 0 {
		p.limit = 10
	}
	return p.limit
}

func (p *paginator) GetPage() int32 {
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