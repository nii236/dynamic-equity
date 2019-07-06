package pie

// Provider of pies
type Provider interface {
	Add(int)
	Remove(int)
	Total(int)
}
