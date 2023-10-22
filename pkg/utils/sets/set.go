package sets

type Set[T comparable] map[T]Empty

func New[T comparable](items ...T) Set[T] {
	ss := make(Set[T], len(items))
	for _, item := range items {
		ss[item] = Empty{}
	}
	return ss
}

func (s Set[T]) UnsortedList() []T {
	ret := make([]T, 0)
	for key := range s {
		ret = append(ret, key)
	}
	return ret
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Insert(items ...T) Set[T] {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

func (s Set[T]) Delete(items ...T) Set[T] {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

func (s Set[T]) Clear() Set[T] {
	for key := range s {
		delete(s, key)
	}
	return s
}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) HasAll(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) HasAny(items ...T) bool {
	for _, item := range items {
		if _, ok := s[item]; ok {
			return true
		}
	}
	return false
}

func (s Set[T]) Clone() Set[T] {
	ret := make(Set[T], len(s))
	for key := range s {
		ret[key] = Empty{}
	}
	return ret
}

func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	diff := New[T]()

	for key := range s {
		if !s2.Has(key) {
			diff.Insert(key)
		}
	}
	return diff
}

func (s Set[T]) SymmeticDifference(s2 Set[T]) Set[T] {
	return s.Difference(s2).Union(s2.Difference(s))
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	ret := s.Clone()
	for key := range s2 {
		ret.Insert(key)
	}
	return ret
}

func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	ret := New[T]()
	for key := range s {
		if s2.Has(key) {
			ret.Insert(key)
		}
	}
	return ret
}

func (s Set[T]) IsSuperset(s2 Set[T]) bool {
	return s.HasAll(s2.UnsortedList()...)
}

func (s Set[T]) Equal(s2 Set[T]) bool {
	return s.IsSuperset(s2) && s.IsSuperset(s)
}

func (s Set[T]) PopAny() (T, bool) {
	for key := range s {
		return key, true
	}
	var zeroValue T
	return zeroValue, false
}
