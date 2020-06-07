package query

type Query struct {
	ID           uint64
	Keywords     []string
	Location     string
	Active       bool
	CreationDate string
}
