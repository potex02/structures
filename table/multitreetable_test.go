package table

import (
	"testing"

	"github.com/potex02/structures"
	"github.com/potex02/structures/util/wrapper"
)

func TestNewMultiTreeTable(t *testing.T) {

	var table structures.Structure[int] = NewMultiTreeTable[wrapper.Int, int]()

	if table == nil {
		t.Log("table is nil")
		t.Fail()
	}
	if table.Len() != 0 {
		t.Log("length is not 0")
		t.Fail()
	}
}
func TestNewMultiTreeTableFromSlice(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if table == nil {
		t.Log("table is nil")
		t.Fail()
	}
	if table.Len() != 4 {
		t.Log("length is not 4")
		t.Fail()
	}
}
func TestContainsMultiTreeTable(t *testing.T) {

	var table MultiTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if !table.Contains("Ciao", 0) {
		t.Log("not found \"Ciao\"-0 in table")
		t.Fail()
	}
	if !table.Contains("Ciao", 5.6) {
		t.Log("not found \"Ciao\"-5.6 in table")
		t.Fail()
	}
	if table.Contains("Ciao", -1) {
		t.Log("found \"Ciao\"-1 in table")
		t.Fail()
	}
	if table.Contains("ciao", 1.2) {
		t.Log("found \"Ciao\"-1.2 in table")
		t.Fail()
	}
}
func TestContainsKeyMultiTreeTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if !table.ContainsKey("Ciao") {
		t.Log("not found \"Ciao\" in table")
		t.Fail()
	}
	if !table.ContainsKey("Hello") {
		t.Log("not found \"Hello\" in table")
		t.Fail()
	}
	if table.ContainsKey("b") {
		t.Log("found \"b\" in table")
		t.Fail()
	}
}
func TestContainsElementMultiTreeTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if !table.ContainsElement(0) {
		t.Log("not found 0 in table")
		t.Fail()
	}
	if !table.ContainsElement(5.6) {
		t.Log("not found 5.6 in table")
		t.Fail()
	}
	if !table.ContainsElement(-1) {
		t.Log("not found -1 in table")
		t.Fail()
	}
	if table.ContainsElement(10) {
		t.Log("found 10 in table")
		t.Fail()
	}
}
func TestGetMultiTreehTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if len(table.Get("Ciao")) != 2 {
		t.Log("not found 2 elements")
		t.Fail()
	}
	if len(table.Get("Hello")) != 1 {
		t.Log("not found 1 element")
		t.Fail()
	}
	if len(table.Get("A")) != 0 {
		t.Log("not found 0 elements")
		t.Fail()
	}
}
func TestPutMultiTreeTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	table.Put("b", -9.3, 2, 5.6)
	if len(table.Get("b")) != 3 {
		t.Log("not found 3 elements")
		t.Fail()
	}
	table.Put("Ciao", -9.3)
	if len(table.Get("Ciao")) != 3 {
		t.Log("not found 3 elements")
		t.Fail()
	}
}
func TestReplaceMultiTreeTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if len(table.Replace("B", 3.4, 2)) != 0 {
		t.Log("not found 0 elements")
		t.Fail()
	}
	if len(table.Get("B")) != 2 {
		t.Log("not found 2 elements")
		t.Fail()
	}
	if len(table.Replace("Ciao", 2)) != 2 {
		t.Log("not found 2 elements")
		t.Fail()
	}
	if len(table.Get("Ciao")) != 1 {
		t.Log("not found 1 element")
		t.Fail()
	}
}
func TestRemoveMultiTreeTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if !table.Remove("Ciao", 0) {
		t.Log("not found \"Ciao\"-0 in table")
		t.Fail()
	}
	if len(table.Get("Ciao")) != 1 {
		t.Log("not found 1 element")
		t.Fail()
	}
	if table.Remove("Ciao", 0) {
		t.Log("found \"Ciao\"-0 in table")
		t.Fail()
	}
	if table.Remove("Ciao", 12) {
		t.Log("found \"Ciao\"-12 elements")
		t.Fail()
	}
	if table.Remove("A", 12) {
		t.Log("found \"A\"-12 elements")
		t.Fail()
	}
}
func TestRemoveKeyMultiTreeTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)

	if len(table.RemoveKey("Ciao")) != 2 {
		t.Log("not found 2 elements")
		t.Fail()
	}
	if len(table.Get("Ciao")) != 0 {
		t.Log("not found 0 element")
		t.Fail()
	}
	if len(table.RemoveKey("A")) != 0 {
		t.Log("not found 0 elements")
		t.Fail()
	}
}
func TestEqualMultiTreeTable(t *testing.T) {

	var table BaseTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)
	var tableTest *MultiTreeTable[wrapper.Int, test] = NewMultiTreeTableFromSlice[wrapper.Int, test](
		[]wrapper.Int{1, -2, -2, 3},
		[]test{{n1: 1, n2: 2}, {n1: -2, n2: -4}, {n1: 1, n2: 2}, {n1: 12, n2: -3}},
	)

	if !table.Equal(NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)) {
		t.Log("tables are not equals")
		t.Fail()
	}
	if table.Equal(NewMultiHashTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.5, 5.6, -1, 0},
	)) {
		t.Log("tables are equals")
		t.Fail()
	}
	if !table.Equal(NewMultiHashTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)) {
		t.Log("tables are not equals")
		t.Fail()
	}
	if !tableTest.Equal(NewMultiTreeTableFromSlice[wrapper.Int, test](
		[]wrapper.Int{1, -2, -2, 3},
		[]test{{n1: -1, n2: 2}, {n1: 2, n2: -4}, {n1: 1, n2: 2}, {n1: 12, n2: -3}},
	)) {
		t.Log("tables are not equals")
		t.Fail()
	}
	if tableTest.Equal(NewMultiTreeTableFromSlice[wrapper.Int, test](
		[]wrapper.Int{1, -2, -2, 3},
		[]test{{n1: 1, n2: -1}, {n1: -2, n2: -4}, {n1: 1, n2: 2}, {n1: 12, n2: -3}},
	)) {
		t.Log("tables are equals")
		t.Fail()
	}
	if !tableTest.Equal(NewMultiHashTableFromSlice[wrapper.Int, test](
		[]wrapper.Int{1, -2, -2, 3},
		[]test{{n1: 1, n2: 2}, {n1: -2, n2: -4}, {n1: 1, n2: 2}, {n1: 12, n2: -3}},
	)) {
		t.Log("tables are not equals")
		t.Fail()
	}
}
func TestCompareMultiTreeTable(t *testing.T) {

	var table *MultiTreeTable[wrapper.String, float32] = NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)
	var tableTest *MultiTreeTable[wrapper.Int, test] = NewMultiTreeTableFromSlice[wrapper.Int, test](
		[]wrapper.Int{1, -2},
		[]test{{n1: 1, n2: 2}, {n1: 12, n2: -3}},
	)

	if table.Compare(NewMultiTreeTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "Ciao", "a", "Ciao"},
		[]float32{1.2, 5.6, -1, 0},
	)) != 0 {
		t.Log("compare is not 0")
		t.Fail()
	}
	if table.Compare(NewMultiHashTableFromSlice[wrapper.String, float32](
		[]wrapper.String{"Hello", "a", "Ciao"},
		[]float32{1.2, -1, 0},
	)) != 1 {
		t.Log("compare is not 1")
		t.Fail()
	}
	if tableTest.Compare(NewMultiTreeTableFromSlice[wrapper.Int, test](
		[]wrapper.Int{1, 2, 4, 6},
		[]test{{n1: -1, n2: 2}, {n1: 2, n2: -4}, {n1: 1, n2: 2}, {n1: 12, n2: -3}},
	)) != -1 {
		t.Log("compare is not -1")
		t.Fail()
	}
	if table.Compare(nil) != -2 {
		t.Log("compare is not -2")
		t.Fail()
	}
}
