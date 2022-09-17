package stl4go

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuiltinSet_Interface(t *testing.T) {
	s := make(BuiltinSet[int])
	_ = Set[int](&s)
}

func TestMakeBuiltinSet(t *testing.T) {
	s := make(BuiltinSet[string])
	expectEq(t, s.Len(), 0)
	expectEq(t, s.IsEmpty(), true)
}

func TestMakeBuiltinSet2(t *testing.T) {
	s := BuiltinSet[string]{}
	expectEq(t, s.Len(), 0)
	expectEq(t, s.IsEmpty(), true)
}

func TestSetOf(t *testing.T) {
	s := SetOf("hello", "world")
	expectEq(t, s.Len(), 2)
}

func TestBuiltinSet_IsEmpty(t *testing.T) {
	s := make(BuiltinSet[string])
	expectEq(t, s.IsEmpty(), true)
	s.Insert("hello")
	expectEq(t, s.IsEmpty(), false)
}

func TestBuiltinSet_Clear(t *testing.T) {
	s := SetOf("hello", "world")
	s.Clear()
	expectTrue(t, s.IsEmpty())
}

func TestBuiltinSet_String(t *testing.T) {
	s := SetOf("hello", "world")
	expectTrue(t, strings.HasPrefix(fmt.Sprintf("%v", s), "BuiltinSet[string]"))
}

func TestBuiltinSet_Has(t *testing.T) {
	s := SetOf("hello", "world")
	expectTrue(t, s.Has("hello"))
	expectTrue(t, s.Has("world"))
	expectFalse(t, s.Has("!"))
}

func TestBuiltinSet_Insert(t *testing.T) {
	s := make(BuiltinSet[string])
	s.Insert("hello")
	s.Insert("hello")
	expectEq(t, s.Has("world"), false)
	s.Insert("world")
	expectEq(t, s.Has("hello"), true)
	expectEq(t, s.Len(), 2)
}

func TestBuiltinSet_InsertN(t *testing.T) {
	s := make(BuiltinSet[string])
	s.InsertN("hello", "world")
	expectEq(t, s.Len(), 2)
}

func TestBuiltinSet_Remove(t *testing.T) {
	s := SetOf("hello", "world")
	expectTrue(t, s.Remove("hello"))
	expectEq(t, s.Len(), 1)
	expectFalse(t, s.Remove("hello"))
	expectEq(t, s.Len(), 1)
	expectTrue(t, s.Remove("world"))
	expectEq(t, s.Len(), 0)
}

func TestBuiltinSet_Delete(t *testing.T) {
	s := SetOf("hello", "world")
	s.Delete("hello")
	expectEq(t, s.Len(), 1)
	s.Delete("hello")
	expectEq(t, s.Len(), 1)
	s.Delete("world")
	expectEq(t, s.Len(), 0)
}

func TestBuiltinSet_RemoveN(t *testing.T) {
	s := SetOf("hello", "world")
	s.RemoveN("hello", "world")
	s.Remove("world")
	expectTrue(t, s.IsEmpty())
}

func TestBuiltinSet_Keys(t *testing.T) {
	s := SetOf("hello", "world")
	ks := s.Keys()
	expectEq(t, 2, len(ks))
}

func TestBuiltinSet_For(t *testing.T) {
	s := SetOf("hello", "world")
	for v := range s {
		expectTrue(t, v == "hello" || v == "world")
	}
}

func TestBuiltinSet_ForEach(t *testing.T) {
	s := SetOf("hello", "world")
	c := 0
	s.ForEach(func(string) {
		c++
	})
	expectEq(t, c, 2)
}

func TestBuiltinSet_ForEachIf(t *testing.T) {
	s := SetOf("hello", "world")
	c := 0
	s.ForEachIf(func(string) bool {
		c++
		return false
	})
	expectLt(t, c, 2)
}

func TestBuiltinSet_Update(t *testing.T) {
	s := SetOf(1, 2, 3)
	s.Update(SetOf(3, 4))
	expectEq(t, s.Len(), 4)
	expectTrue(t, s.Has(4))
}

func TestBuiltinSet_Union(t *testing.T) {
	s := SetOf(1, 2, 3)
	s2 := s.Union(SetOf(3, 4))
	expectEq(t, s2.Len(), 4)
	expectTrue(t, s2.Has(4))
}

func TestBuiltinSet_Intersection(t *testing.T) {
	s := SetOf(1, 2, 3).Intersection(SetOf(3, 4))
	expectEq(t, s.Len(), 1)
	expectTrue(t, s.Has(3))
	s = SetOf(3, 4).Intersection(SetOf(1, 2, 3))
	expectEq(t, s.Len(), 1)
	expectTrue(t, s.Has(3))
}

func TestBuiltinSet_Difference(t *testing.T) {
	s := SetOf(1, 2, 3).Difference(SetOf(3, 4))
	expectEq(t, s.Len(), 2)
	expectTrue(t, s.Has(1))
	expectTrue(t, s.Has(2))
	s = SetOf(1, 2).Difference(SetOf(3, 4))
	expectEq(t, s.Len(), 2)
	expectTrue(t, s.Has(1))
	expectTrue(t, s.Has(2))
}

func TestBuiltinSet_IsDisjointOf(t *testing.T) {
	s1 := SetOf(1, 2, 3)
	s2 := SetOf(3, 4)
	expectFalse(t, s1.IsDisjointOf(s2))
	expectTrue(t, s1.IsDisjointOf(SetOf(4, 5)))
}

func TestBuiltinSet_IsSubsetOf(t *testing.T) {
	expectTrue(t, SetOf[int]().IsSubsetOf(SetOf[int]()))
	expectTrue(t, SetOf[int]().IsSubsetOf(SetOf(1)))
	expectTrue(t, SetOf(1, 2, 3).IsSubsetOf(SetOf(1, 2, 3)))
	expectTrue(t, SetOf(1, 2).IsSubsetOf(SetOf(1, 2, 3)))
	expectFalse(t, SetOf(1, 2, 3).IsSubsetOf(SetOf(1, 2)))
	expectFalse(t, SetOf(1, 2).IsSubsetOf(SetOf(2, 3)))
}

func TestBuiltinSet_IsSupersetOf(t *testing.T) {
	expectTrue(t, SetOf[int]().IsSupersetOf(SetOf[int]()))
	expectTrue(t, SetOf(1).IsSupersetOf(SetOf[int]()))
	expectTrue(t, SetOf(1, 2, 3).IsSupersetOf(SetOf(1, 2, 3)))
	expectTrue(t, SetOf(1, 2, 3).IsSupersetOf(SetOf(1, 2)))
	expectFalse(t, SetOf(1, 2).IsSupersetOf(SetOf(1, 2, 3)))
	expectFalse(t, SetOf(1, 2).IsSupersetOf(SetOf(2, 3)))
}
