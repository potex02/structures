package list

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewLinkedList(t *testing.T) {

	var list structures.Structure[int] = NewLinkedList[int]()

	if list == nil {
		t.Log("list is nil")
		t.Fail()
	}
	if list.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewFromSliceLinkedList(t *testing.T) {

	var list *LinkedList[int] = NewLinkedListFromSlice([]int{1, 2, 3, -4})

	if list == nil {
		t.Log("list is nil")
		t.Fail()
	}
	if list.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
	if !reflect.DeepEqual(list.ToSlice(), []int{1, 2, 3, -4}) {
		t.Log("list is", list)
		t.Fail()
	}
}
func TestContainsLinkedList(t *testing.T) {

	var list List[int] = NewLinkedListFromSlice([]int{1, 2, 3, -4})

	if list.Contains(-1) {
		t.Log("found -1 in list")
		t.Fail()
	}
	if !list.Contains(2) {
		t.Log("not found 2 in list")
		t.Fail()
	}
}
func TestIndexOfLinkedList(t *testing.T) {

	var list *LinkedList[float32] = NewLinkedListFromSlice([]float32{1.4, 3.45, 2.5, 3.45, 1.4, -5.9})

	if list.IndexOf(3.45) != 1 {
		t.Log("3.45 is not in 1 position")
		t.Fail()
	}
	if list.IndexOf(2) != -1 {
		t.Log("2 is present")
		t.Fail()
	}
	if list.LastIndexOf(1.4) != 4 {
		t.Log("last 1.4 is not in 4 position")
		t.Fail()
	}
	if list.LastIndexOf(2) != -1 {
		t.Log("2 is present (last)")
		t.Fail()
	}
}
func TestToSliceLinkedList(t *testing.T) {

	var list *LinkedList[int] = NewLinkedList(1, 2, 3, -4)

	if slice := list.ToSlice(); !reflect.DeepEqual(slice, []int{1, 2, 3, -4}) {
		t.Log("slice is", slice)
		t.Fail()
	}

}
func TestGetLinkedList(t *testing.T) {

	var list List[int] = NewLinkedList(1, 2, 3, 5)

	e, err := list.Get(1)
	if err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if e != 2 {
		t.Log("e is", e)
		t.Fail()
	}
	e, err = list.Get(-1)
	if err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if e != 5 {
		t.Log("e is", e)
		t.Fail()
	}
	if _, err = list.Get(4); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
	if _, err = list.Get(-5); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
}
func TestGetDefaultLinkedList(t *testing.T) {

	var list List[int] = NewLinkedList(1, 2, 3, 5)

	if e := list.GetDefault(1); e != 2 {
		t.Log("e is", e)
		t.Fail()
	}
	if e := list.GetDefaultValue(-1, 4); e != 5 {
		t.Log("e is", e)
		t.Fail()
	}
	if e := list.GetDefaultValue(4, 10); e != 10 {
		t.Log("e is", e)
		t.Fail()
	}
	if e := list.GetDefault(-5); e != 0 {
		t.Log("e is", e)
		t.Fail()
	}
}
func TestSetLinkedList(t *testing.T) {

	var list *ArrayList[int] = NewArrayList(1, 2, 3, 5)

	e, err := list.Set(1, 4)
	if err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if e != 2 {
		t.Log("e is", e)
		t.Fail()
	}
	e, err = list.Set(4, -1)
	if err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if e != 0 {
		t.Log("e is", e)
		t.Fail()
	}
	e, err = list.Set(-2, 10)
	if err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if e != 5 {
		t.Log("e is", e)
		t.Fail()
	}
	if _, err = list.Set(6, -1); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
	if _, err = list.Set(-10, 3); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
	if !reflect.DeepEqual(list.objects, []int{1, 4, 3, 10, -1}) {
		t.Log("list is", list)
		t.Fail()
	}
}
func TestAddLinkedList(t *testing.T) {

	var list *LinkedList[int] = NewLinkedList(1, 2, 3, 5)

	list.Add(2)
	if e, _ := list.Get(4); e != 2 {
		t.Log("list[4] is", e)
		t.Fail()
	}
	if err := list.AddAtIndex(2, -3); err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if e, _ := list.Get(2); e != -3 {
		t.Log("list[2] is", e)
		t.Fail()
	}
	if err := list.AddAtIndex(5, -8); err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if e, _ := list.Get(5); e != -8 {
		t.Log("list[5] is", e)
		t.Fail()
	}
	if err := list.AddAtIndex(-1, 3); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
	if err := list.AddAtIndex(10, 8); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
}
func TestAddSliceLinkedList(t *testing.T) {

	var list *LinkedList[int] = NewLinkedList(1, 2, 3, 5)

	list.AddSlice([]int{12, -13})
	if !reflect.DeepEqual(list.ToSlice(), []int{1, 2, 3, 5, 12, -13}) {
		t.Log("list is", list)
		t.Fail()
	}
	list.Add(1, 2)
	if !reflect.DeepEqual(list.ToSlice(), []int{1, 2, 3, 5, 12, -13, 1, 2}) {
		t.Log("list is", list)
		t.Fail()
	}
	list = NewLinkedList(1, 2, 3, 5)
	if err := list.AddSliceAtIndex(0, []int{12, -13}); err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if !reflect.DeepEqual(list.ToSlice(), []int{12, -13, 1, 2, 3, 5}) {
		t.Log("list is", list)
		t.Fail()
	}
	if err := list.AddAtIndex(2, 1, 2); err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if !reflect.DeepEqual(list.ToSlice(), []int{12, -13, 1, 2, 1, 2, 3, 5}) {
		t.Log("list is", list)
		t.Fail()
	}
	if err := list.AddSliceAtIndex(-1, []int{12, -13}); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
	if err := list.AddAtIndex(list.Len()+1, 12, -13); err == nil {
		t.Log("error is", err)
		t.Fail()
	}
}
func TestRemoveLinkedList(t *testing.T) {

	var list *LinkedList[int] = NewLinkedList(1, 2, 3, 5)

	if !list.RemoveElement(2) {
		t.Fail()
	}
	if list.Len() != 3 {
		t.Fail()
	}
	if !reflect.DeepEqual(list, NewLinkedListFromSlice([]int{1, 3, 5})) {
		t.Log("lists not equals")
		t.Fail()
	}
	if _, err := list.Remove(0); err != nil {
		t.Fail()
	}
	if _, err := list.Remove(-2); err != nil {
		t.Fail()
	}
	if _, err := list.Remove(1); err == nil {
		t.Fail()
	}
	if _, err := list.Remove(-2); err == nil {
		t.Fail()
	}
	if !reflect.DeepEqual(list, NewLinkedList(5)) {
		t.Log("lists not equals")
		t.Fail()
	}
}
func TestSortLinkedList(t *testing.T) {

	var list List[wrapper.Int] = NewLinkedList[wrapper.Int](1, -2, 5, -3)
	var linkedList *LinkedList[int] = NewLinkedList[int](1, -2, 5, -3)

	list.Sort()
	if !list.Equal(NewArrayList[wrapper.Int](-3, -2, 1, 5)) {
		t.Log("list is", list)
		t.Fail()
	}
	linkedList.SortFunc(func(i, j int) int {
		if i < j {
			return -1
		}
		if i == j {
			return 0
		}
		return 1
	})
	if !linkedList.Equal(NewArrayList[int](-3, -2, 1, 5)) {
		t.Log("list is", linkedList)
		t.Fail()
	}
}
func TestIterLinkedList(t *testing.T) {

	var list *LinkedList[int] = NewLinkedList(1, -2, 3, 5)
	var j int = 0

	for i := list.Iter(); !i.End(); i = i.Next() {
		value, err := list.Get(j)
		if err != nil {
			t.Log("error is", err)
			t.Fail()
		}
		if j != i.Index() {
			t.Log("index is", i.Index())
			t.Fail()
		}
		if value != i.Element() {
			t.Log("element is", i.Element())
			t.Fail()
		}
		j++
	}
	j = list.Len() - 1
	for i := list.IterReverse(); !i.End(); i = i.Prev() {
		value, err := list.Get(j)
		if err != nil {
			t.Log("error is", err)
			t.Fail()
		}
		if j != i.Index() {
			t.Log("index is", i.Index())
			t.Fail()
		}
		if value != i.Element() {
			t.Log("element is", i.Element())
			t.Fail()
		}
		j--
	}
}
func TestEqualLinkedList(t *testing.T) {

	var list List[int] = NewLinkedList(1, 2, 3, 5)
	var listTest List[test] = NewLinkedList[test](test{n1: 1, n2: 2}, test{n1: -2, n2: -4})

	if !list.Equal(NewLinkedListFromSlice([]int{1, 2, 3, 5})) {
		t.Log("lists are not equals")
		t.Fail()
	}
	if list.Equal(NewLinkedListFromSlice([]int{-1, 2, 3, 5})) {
		t.Log("lists are equals")
		t.Fail()
	}
	if !list.Equal(NewArrayListFromSlice([]int{1, 2, 3, 5})) {
		t.Log("lists are not equals")
		t.Fail()
	}
	if list.Equal(NewArrayListFromSlice([]int{-1, 2, 3, 5})) {
		t.Log("lists are equals")
		t.Fail()
	}
	if !listTest.Equal(NewArrayList[test](test{n1: 2, n2: 2}, test{n1: 0, n2: -4})) {
		t.Log("lists are not equals")
		t.Fail()
	}
	if listTest.Equal(NewLinkedList[test](test{n1: 1, n2: 1}, test{n1: -2, n2: -4})) {
		t.Log("lists are equals")
		t.Fail()
	}
}
func TestCompareLinkedList(t *testing.T) {

	var list List[List[int]] = NewLinkedList[List[int]](
		NewLinkedList(1, 2, 3),
		NewArrayList(4, 2),
		NewArrayList(5, 6, 8, 8),
	)
	var listWrapper List[List[wrapper.Int]] = NewLinkedList[List[wrapper.Int]](
		NewLinkedList[wrapper.Int](1, 2, 3),
		NewArrayList[wrapper.Int](4, 2),
		NewArrayList[wrapper.Int](5, 6, 8, 8),
	)

	if list.Compare(NewArrayList[List[int]](
		NewArrayList(1, 2, 3),
		NewLinkedList(2, 2),
		NewArrayList(5, 6, 8, 8),
	)) != 0 {
		t.Log("compare is not 0")
		t.Fail()
	}
	if list.Compare(NewArrayList[List[int]](
		NewLinkedList(1, 2, 3),
		NewArrayList(5, 6, 8, 8),
	)) != 1 {
		t.Log("compare is not 1")
		t.Fail()
	}
	if listWrapper.Compare(NewArrayList[List[wrapper.Int]](
		NewLinkedList[wrapper.Int](1, 2, 3),
		NewLinkedList[wrapper.Int](5, 2),
		NewArrayList[wrapper.Int](5, 6, 8, 8),
	)) != -1 {
		t.Log("compare is not -1")
		t.Fail()
	}
	if listWrapper.Compare(NewArrayList[List[wrapper.Int]](
		NewLinkedList[wrapper.Int](1, 2, 3),
		NewLinkedList[wrapper.Int](5, 6, 8, 8),
	)) != 1 {
		t.Log("compare is not 1")
		t.Fail()
	}
}
func TestCopyLinkedList(t *testing.T) {

	var list List[int] = NewLinkedList(1, -2, 5, -3)
	var linkedList *LinkedList[int] = NewLinkedList(1, -2, 5, -3)

	if !reflect.DeepEqual(list.Copy(), NewLinkedList(1, -2, 5, -3)) {
		t.Log("list is", list.Copy())
		t.Fail()
	}
	if !reflect.DeepEqual(linkedList.Copy(), NewLinkedList(1, -2, 5, -3)) {
		t.Log("list is", linkedList.Copy())
		t.Fail()
	}
}
