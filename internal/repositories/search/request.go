package search

type Request struct {
	Query     string
	IndexName string
	Page      int
	NextPage  int
	Size      int
}

func (r Request) FromIndex() int {
	return r.Page * r.Size
}

func (r Request) NextPageFromIndex() int {
	page := r.NextPage

	if page == 0 {
		page = r.Page + 1
	}

	return page * r.Size
}

func (r Request) GetSize(availableLength int) int {
	if r.Size == 0 || availableLength < r.Size {
		return availableLength
	}

	return r.Size
}
