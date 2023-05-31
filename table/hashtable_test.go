package table

import (
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewHashTable(t *testing.T) {

	var table structures.Structure[int] = NewHashTable[wrapper.String, int]()

	if table == nil {

		t.Log("table is nil")
		t.Fail()

	}
	if table.Len() != 0 {

		t.Log("length is not 0")
		t.Fail()

	}

}
func TestNewHashTableFromSlice(t *testing.T) {

	var table *HashTable[wrapper.String, float32] = NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.2, 5.6})

	if table == nil {

		t.Log("table is nil")
		t.Fail()

	}
	if table.Len() != 2 {

		t.Log("length is not 2")
		t.Fail()

	}
	if e, _ := table.Get("Hello"); e != 1.2 {

		t.Log("element is", e)
		t.Fail()

	}
	if e, _ := table.Get("Ciao"); e != 5.6 {

		t.Log("element is", e)
		t.Fail()

	}

}
func TestContainsKeyHashTable(t *testing.T) {

	var table *HashTable[wrapper.String, float32] = NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.2, 5.6})

	if ok := table.ContainsKey("hello"); ok {

		t.Log("found \"hello\" in table")
		t.Fail()

	}
	if ok := table.ContainsKey("Hello"); !ok {

		t.Log("not found \"Hello\" in table")
		t.Fail()

	}

}
func TestContainsElementHashTable(t *testing.T) {

	var table *HashTable[wrapper.String, float32] = NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.2, 5.6})

	if ok := table.ContainsElement(-1); ok {

		t.Log("found -1 in table")
		t.Fail()

	}
	if ok := table.ContainsElement(1.2); !ok {

		t.Log("not found 1.2 in table")
		t.Fail()

	}

}
func TestGetHashTable(t *testing.T) {

	var table *HashTable[wrapper.Byte, int] = NewHashTableFromSlice([]wrapper.Byte{'a', 'b'}, []int{1, -1})

	e, ok := table.Get('a')
	if !ok {

		t.Log("not found 'a' in table")
		t.Fail()

	}
	if e != 1 {

		t.Log("e is", e)
		t.Fail()

	}
	if _, ok := table.Get('c'); ok {

		t.Log("found 'c' in table")
		t.Fail()

	}

}
func TestPutHashTable(t *testing.T) {

	var table *HashTable[wrapper.String, float32] = NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.2, 5.6})

	if _, ok := table.Put("a", -7.8); ok {

		t.Log("found \"a\" in table")
		t.Fail()

	}
	if e, _ := table.Get("a"); e != -7.8 {

		t.Log("table[\"a\"] is", e)
		t.Fail()

	}
	if _, ok := table.Put("Hello", -7.85); !ok {

		t.Log("not found \"Hello\" in table")
		t.Fail()

	}
	if e, _ := table.Get("Hello"); e != -7.85 {

		t.Log("table[\"Hello\"] is", e)
		t.Fail()

	}

}
func TestRemove(t *testing.T) {

	var table *HashTable[wrapper.String, float32] = NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.2, 5.6})

	e, ok := table.Remove("Ciao")
	if !ok {

		t.Log("not found \"Ciao\" in table")
		t.Fail()

	}
	if e != 5.6 {

		t.Log("e is", e)
		t.Fail()

	}
	_, ok = table.Remove("a")
	if ok {

		t.Log("found \"a\" in table")
		t.Fail()

	}

}
func TestEqualsHashTable(t *testing.T) {

	var table *HashTable[wrapper.String, float32] = NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.2, 5.6})
	var tableTest *HashTable[wrapper.Int, test] = NewHashTableFromSlice[wrapper.Int, test]([]wrapper.Int{1, 2}, []test{{n1: 1, n2: 2}, {n1: -2, n2: -4}})

	if !table.Equal(NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.2, 5.6})) {

		t.Log("tables are not equals")
		t.Fail()

	}
	if table.Equal(NewHashTableFromSlice([]wrapper.String{"Hello", "Ciao"}, []float32{1.5, 5.6})) {

		t.Log("tables are equals")
		t.Fail()

	}
	if table.Equal(NewHashTableFromSlice([]wrapper.String{"Hello", "ciao"}, []float32{1.2, 5.6})) {

		t.Log("tables not equals")
		t.Fail()

	}
	if !tableTest.Equal(NewHashTableFromSlice[wrapper.Int, test]([]wrapper.Int{1, 2}, []test{{n1: 2, n2: 2}, {n1: 0, n2: -4}})) {

		t.Log("tables are not equals")
		t.Fail()

	}
	if tableTest.Equal(NewHashTableFromSlice[wrapper.Int, test]([]wrapper.Int{1, 2}, []test{{n1: 1, n2: 1}, {n1: -2, n2: -4}})) {

		t.Log("tables are equals")
		t.Fail()

	}
	if tableTest.Equal(NewHashTableFromSlice[wrapper.Int, test]([]wrapper.Int{-1, 2}, []test{{n1: 1, n2: 2}, {n1: -2, n2: -4}})) {

		t.Log("tables are equals")
		t.Fail()

	}

}

type test struct {
	n1 int
	n2 int
}

func (t test) Equal(o test) bool {

	return t.n2 == o.n2

}
