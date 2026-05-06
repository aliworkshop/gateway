package gateway

type Sorter interface {
	GetSortBy() string
}

type ISorter interface {
	Sorter
	SetSort(string)
}

type sort struct {
	sortBy string
}

func NewSorter() ISorter {
	return &sort{}
}

func (s *sort) GetSortBy() string {
	if s.sortBy == "" {
		s.sortBy = "-id"
	}
	return s.sortBy
}

func (s *sort) SetSort(sort string) {
	s.sortBy = sort
}