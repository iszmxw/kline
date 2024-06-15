package page

const (
	DESC = "DESC"
	ASC  = "ASC"
)

type PageApple struct {
	PageSize   *int
	PageNumber *int
	Design     string
	Index      *int
	Ends       *int
}

func NewPageApple(pageSize, pageNumber int) *PageApple {
	p := &PageApple{}
	p.SetPageSize(pageSize)
	p.SetPageNumber(pageNumber)
	p.SetEnds(pageSize)
	p.SetIndex(pageSize, pageNumber)
	return p
}

func (p *PageApple) SetIndex(pageSize, pageNumber int) {
	if *p.PageNumber == 1 {
		index := 0
		p.Index = &index
	} else {
		index := *p.PageSize * (*p.PageNumber - 1)
		p.Index = &index
	}
}

func (p *PageApple) SetEnds(pageSize int) {
	p.Ends = p.PageSize
}

func (p *PageApple) GetPageSize() *int {
	return p.PageSize
}

func (p *PageApple) SetPageSize(pageSize int) {
	p.PageSize = &pageSize
}

func (p *PageApple) GetPageNumber() *int {
	return p.PageNumber
}

func (p *PageApple) SetPageNumber(pageNumber int) {
	p.PageNumber = &pageNumber
}

func (p *PageApple) GetDesign() string {
	return p.Design
}

func (p *PageApple) SetDesign(design string) {
	p.Design = design
}

func GetPage(total, pageSize int) int {
	page := 0
	if total < pageSize {
		return 1
	}
	flag := total%pageSize > 0
	if flag {
		page = total/pageSize + 1
	} else {
		page = total / pageSize
	}
	return page
}
