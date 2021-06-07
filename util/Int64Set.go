package util

type Int64Set map[int64]struct{}

func (s Int64Set) Has(key int64) bool {
	_, ok := s[key]
	return ok
}

func (s Int64Set) Add(key int64) {
	s[key] = struct{}{}
}

func (s Int64Set) Delete(key int64) {
	delete(s, key)
}
