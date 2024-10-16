package list

import (
	"reflect"
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewArrayList(t *testing.T) {

	var list structures.Structure[int] = NewArrayList[int]()

	if list == nil {
		t.Log("list is nil")
		t.Fail()
	}
	if list.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewFromSliceArrayList(t *testing.T) {

	var list *ArrayList[int] = NewArrayListFromSlice([]int{1, 2, 3, -4})

	if list == nil {
		t.Log("list is nil")
		t.Fail()
	}
	if list.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
	if !reflect.DeepEqual(list.objects, []int{1, 2, 3, -4}) {
		t.Log("list is", list)
		t.Fail()
	}
}
func TestContainsArrayList(t *testing.T) {

	var list List[int] = NewArrayListFromSlice([]int{1, 2, 3, -4})

	if list.Contains(-1) {
		t.Log("found -1 in list")
		t.Fail()
	}
	if !list.Contains(2) {
		t.Log("not found 2 in list")
		t.Fail()
	}
}
func TestIndexOfArrayList(t *testing.T) {

	var list *ArrayList[float32] = NewArrayListFromSlice([]float32{1.4, 3.45, 2.5, 3.45, 1.4, -5.9})

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
func TestToSliceArrayList(t *testing.T) {

	var list *ArrayList[int] = NewArrayList(1, 2, 3, -4)

	if slice := list.ToSlice(); !reflect.DeepEqual(slice, []int{1, 2, 3, -4}) {
		t.Log("slice is", slice)
		t.Fail()
	}
}
func TestGetArrayList(t *testing.T) {

	var list List[int] = NewArrayList(1, 2, 3, 5)

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
func TestGetDefaultArrayList(t *testing.T) {

	var list List[int] = NewArrayList(1, 2, 3, 5)

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
func TestSetArrayList(t *testing.T) {
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
func TestAddArrayList(t *testing.T) {

	var list *ArrayList[int] = NewArrayList(1, 2, 3, 5)

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
func TestAddSliceArrayList(t *testing.T) {

	var list *ArrayList[int] = NewArrayList(1, 2, 3, 5)

	list.AddSlice([]int{12, -13})
	if !reflect.DeepEqual(list.objects, []int{1, 2, 3, 5, 12, -13}) {
		t.Log("list is", list)
		t.Fail()
	}
	list.Add(1, 2)
	if !reflect.DeepEqual(list.objects, []int{1, 2, 3, 5, 12, -13, 1, 2}) {
		t.Log("list is", list)
		t.Fail()
	}
	list = NewArrayList(1, 2, 3, 5)
	if err := list.AddSliceAtIndex(0, []int{12, -13}); err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if !reflect.DeepEqual(list.objects, []int{12, -13, 1, 2, 3, 5}) {
		t.Log("list is", list)
		t.Fail()
	}
	if err := list.AddAtIndex(2, 1, 2); err != nil {
		t.Log("error is", err)
		t.Fail()
	}
	if !reflect.DeepEqual(list.objects, []int{12, -13, 1, 2, 1, 2, 3, 5}) {
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
func TestRemoveArrayList(t *testing.T) {

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
func TestSortArrayList(t *testing.T) {

	var list List[wrapper.Int] = NewArrayList[wrapper.Int](1, -2, 5, -3)
	var arrayList *ArrayList[int] = NewArrayList[int](1, -2, 5, -3)

	list.Sort()
	if !list.Equal(NewArrayList[wrapper.Int](-3, -2, 1, 5)) {
		t.Log("list is", list)
		t.Fail()
	}
	arrayList.SortFunc(func(i, j int) bool {
		return i < j
	})
	if !arrayList.Equal(NewArrayList[int](-3, -2, 1, 5)) {
		t.Log("list is", list)
		t.Fail()
	}
}
func TestIterArrayList(t *testing.T) {

	var list *ArrayList[int] = NewArrayList(1, -2, 3, 5)
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
func TestEqualArrayList(t *testing.T) {

	var list List[int] = NewArrayList(1, 2, 3, 5)
	var listTest List[test] = NewArrayList[test](test{n1: 1, n2: 2}, test{n1: -2, n2: -4})

	if !list.Equal(NewArrayListFromSlice([]int{1, 2, 3, 5})) {
		t.Log("lists are not equals")
		t.Fail()
	}
	if list.Equal(NewArrayListFromSlice([]int{-1, 2, 3, 5})) {
		t.Log("lists are equals")
		t.Fail()
	}
	if !list.Equal(NewLinkedListFromSlice([]int{1, 2, 3, 5})) {
		t.Log("lists are not equals")
		t.Fail()
	}
	if list.Equal(NewLinkedListFromSlice([]int{-1, 2, 3, 5})) {
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
func TestCompareArrayList(t *testing.T) {

	var list List[List[int]] = NewArrayList[List[int]](
		NewLinkedList(1, 2, 3),
		NewArrayList(4, 2),
		NewArrayList(5, 6, 8, 8),
	)
	var listWrapper List[List[wrapper.Int]] = NewArrayList[List[wrapper.Int]](
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
		NewArrayList(1, 2, 3),
		NewArrayList(5, 6, 8, 8),
	)) != 1 {
		t.Log("compare is not 1")
		t.Fail()
	}
	if listWrapper.Compare(NewArrayList[List[wrapper.Int]](
		NewArrayList[wrapper.Int](1, 2, 3),
		NewArrayList[wrapper.Int](5, 2),
		NewArrayList[wrapper.Int](5, 6, 8, 8),
	)) != -1 {
		t.Log("compare is not -1")
		t.Fail()
	}
	if listWrapper.Compare(NewArrayList[List[wrapper.Int]](
		NewArrayList[wrapper.Int](1, 2, 3),
		NewArrayList[wrapper.Int](5, 6, 8, 8),
	)) != 1 {
		t.Log("compare is not 1")
		t.Fail()
	}
}
func TestCopyArrayList(t *testing.T) {

	var list List[int] = NewArrayList(1, -2, 5, -3)
	var arrayList *ArrayList[int] = NewArrayList(1, -2, 5, -3)

	if !reflect.DeepEqual(list.Copy(), NewArrayList(1, -2, 5, -3)) {
		t.Log("list is", list.Copy())
		t.Fail()
	}
	if !reflect.DeepEqual(arrayList.Copy(), NewArrayList(1, -2, 5, -3)) {
		t.Log("list is", arrayList.Copy())
		t.Fail()
	}
}

type test struct {
	n1, n2 int
}

func (t test) Equal(o any) bool {
	test, ok := o.(test)
	return ok && t.n2 == test.n2
}

func (t test) Compare(o any) int {
	test, ok := o.(test)
	if !ok {
		return -2
	}
	return wrapper.Int(t.n1).Compare(wrapper.Int(test.n1))
}
