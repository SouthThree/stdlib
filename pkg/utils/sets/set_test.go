package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(t *testing.T)
	}{
		{
			name: "new set",
			testFunc: func(t *testing.T) {
				s := New[int](1, 2)
				assert.Equal(t, Set[int]{2: {}, 1: {}}, s)
			},
		},
		{
			name: "insert",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}}
				assert.Equal(t, Set[int]{1: {}, 2: {}, 3: {}}, s.Insert(3))
			},
		},
		{
			name: "delete",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}}
				assert.Equal(t, Set[int]{2: {}}, s.Delete(1))
			},
		},
		{
			name: "clear",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}}
				assert.Equal(t, Set[int]{}, s.Clear())
			},
		},
		{
			name: "has",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}}
				assert.True(t, s.Has(1))
			},
		},
		{
			name: "not has",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}}
				assert.False(t, s.Has(3))
			},
		},
		{
			name: "has all",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}, 3: {}}
				assert.True(t, s.HasAll(1, 3))
			},
		},
		{
			name: "not has all",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}, 3: {}}
				assert.False(t, s.HasAll(1, 3, 4))
			},
		},
		{
			name: "has any",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}, 3: {}}
				assert.True(t, s.HasAny(1, 3, 4))
			},
		},
		{
			name: "not has any",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}, 3: {}}
				assert.False(t, s.HasAll(4, 5))
			},
		},
		{
			name: "clone",
			testFunc: func(t *testing.T) {
				s := Set[int]{1: {}, 2: {}}
				assert.Equal(t, Set[int]{1: {}, 2: {}}, s.Clone())
			},
		},
		{
			name: "difference",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}}
				s2 := Set[int]{2: {}, 3: {}, 4: {}}
				assert.Equal(t, Set[int]{1: {}}, s1.Difference(s2))
			},
		},
		{
			name: "symmetic difference",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}}
				s2 := Set[int]{2: {}, 3: {}, 4: {}}
				assert.Equal(t, Set[int]{1: {}, 3: {}, 4: {}}, s1.SymmeticDifference(s2))
			},
		},
		{
			name: "union",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}}
				s2 := Set[int]{2: {}, 3: {}, 4: {}}
				assert.Equal(t, Set[int]{1: {}, 2: {}, 3: {}, 4: {}}, s1.Union(s2))
			},
		},
		{
			name: "is superset",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}, 3: {}}
				s2 := Set[int]{2: {}, 3: {}}
				assert.True(t, s1.IsSuperset(s2))
			},
		},
		{
			name: "is not superset",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}, 3: {}}
				s2 := Set[int]{2: {}, 3: {}, 4: {}}
				assert.False(t, s1.IsSuperset(s2))
			},
		},
		{
			name: "equal",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}, 3: {}}
				s2 := Set[int]{2: {}, 3: {}, 1: {}}
				assert.True(t, s1.Equal(s2))
			},
		},
		{
			name: "not equal",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}, 3: {}}
				s2 := Set[int]{2: {}, 3: {}, 4: {}}
				assert.False(t, s1.Equal(s2))
			},
		},
		{
			name: "pop any ok",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}}
				key, ok := s1.PopAny()
				assert.True(t, ok)
				assert.Equal(t, 1, key)
			},
		},
		{
			name: "pop any not",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{}
				key, ok := s1.PopAny()
				assert.False(t, ok)
				assert.Equal(t, 0, key)
			},
		},
		{
			name: "unsorted list",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}}
				keys := s1.UnsortedList()
				assert.Equal(t, []int{1}, keys)
			},
		},
		{
			name: "len",
			testFunc: func(t *testing.T) {
				s1 := Set[int]{1: {}, 2: {}, 3: {}}
				assert.Equal(t, 3, s1.Len())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testFunc)
	}
}
