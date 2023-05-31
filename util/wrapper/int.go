package wrapper

import "fmt"

var _ Wrapper[int, Int] = Int(0)
var _ Wrapper[int8, Int8] = Int8(0)
var _ Wrapper[int16, Int16] = Int16(0)
var _ Wrapper[int32, Int32] = Int32(0)
var _ Wrapper[int64, Int64] = Int64(0)
var _ Wrapper[rune, Rune] = Rune('a')

// Int is a wrapper type for int.
type Int int

// Equal returns true if i and o are equals.
func (i Int) Equal(o Int) bool {

	return i == o

}

// Compare returns -1 if i is less than o,
// 1 if i is greater than o,
// 0 if i and o are equals.
func (i Int) Compare(o Int) int {

	if i < o {

		return -1

	}
	if i == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of i.
func (i Int) Hash() string {

	return fmt.Sprintf("%v", i)

}

// ToValue returns the wrapped value by i.
func (i Int) ToValue() int {

	return int(i)

}

// Int8 is a wrapper type for int8.
type Int8 int8

// Equal returns true if i and o are equals.
func (i Int8) Equal(o Int8) bool {

	return i == o

}

// Compare returns -1 if i is less than o,
// 1 if i is greater than o,
// 0 if i and o are equals.
func (i Int8) Compare(o Int8) int {

	if i < o {

		return -1

	}
	if i == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of i.
func (i Int8) Hash() string {

	return fmt.Sprintf("%v", i)

}

// ToValue returns the wrapped value by i.
func (i Int8) ToValue() int8 {

	return int8(i)

}

// Int16 is a wrapper type for int16.
type Int16 int16

// Equal returns true if i and o are equals.
func (i Int16) Equal(o Int16) bool {

	return i == o

}

// Compare returns -1 if i is less than o,
// 1 if i is greater than o,
// 0 if i and o are equals.
func (i Int16) Compare(o Int16) int {

	if i < o {

		return -1

	}
	if i == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of i.
func (i Int16) Hash() string {

	return fmt.Sprintf("%v", i)

}

// ToValue returns the wrapped value by i.
func (i Int16) ToValue() int16 {

	return int16(i)

}

// Int32 is a wrapper type for int32.
type Int32 int32

// Equal returns true if i and o are equals.
func (i Int32) Equal(o Int32) bool {

	return i == o

}

// Compare returns -1 if i is less than o,
// 1 if i is greater than o,
// 0 if i and o are equals.
func (i Int32) Compare(o Int32) int {

	if i < o {

		return -1

	}
	if i == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of i.
func (i Int32) Hash() string {

	return fmt.Sprintf("%v", i)

}

// ToValue returns the wrapped value by i.
func (i Int32) ToValue() int32 {

	return int32(i)

}

// Int64 is a wrapper type for int64.
type Int64 int64

// Equal returns true if i and o are equals.
func (i Int64) Equal(o Int64) bool {

	return i == o

}

// Compare returns -1 if i is less than o,
// 1 if i is greater than o,
// 0 if i and o are equals.
func (i Int64) Compare(o Int64) int {

	if i < o {

		return -1

	}
	if i == o {

		return 0

	}
	return 1

}

// Hash returns the hash code of i.
func (i Int64) Hash() string {

	return fmt.Sprintf("%v", i)

}

// ToValue returns the wrapped value by i.
func (i Int64) ToValue() int64 {

	return int64(i)

}

// Rune is a wrapper type for rune.
type Rune = Int32
