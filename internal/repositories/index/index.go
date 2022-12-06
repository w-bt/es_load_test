package index

type Request struct {
	IndexName string
	Refresh   bool
	Flush     bool
	BatchSize int
}
