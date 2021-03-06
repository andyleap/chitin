// This file is automatically generated, DO NOT EDIT

package person

import (
	"encoding/binary"
	"io"
	"reflect"
	"unsafe"

	"chitin.io/chitin"
	"github.com/dchest/varuint"
)

// use all packages to avoid errors
var (
	_ = io.ErrUnexpectedEOF
	_ reflect.StringHeader
	_ unsafe.Pointer
	_ = varuint.Uint64
)

func chitinParseLengthPrefixed(data []byte) (msg []byte, next []byte, err error) {
loop:
	l, n := varuint.Uint64(data)
	if n < 0 {
		return nil, nil, io.ErrUnexpectedEOF
	}
	if l == 0 {
		// padding
		data = data[n:]
		goto loop
	}
	l--

	const maxInt = int(^uint(0) >> 1)
	if l > uint64(maxInt) {
		// technically, it has to be truncated because it wouldn't fit
		// in memory ;)
		return nil, nil, io.ErrUnexpectedEOF
	}
	li := int(l)

	// TODO prevent overflow here
	end := n + li
	if end > len(data) {
		return nil, nil, io.ErrUnexpectedEOF
	}

	low := n
	high := low + li
	return data[low:high], data[high:], nil
}

func NewPersonV2View(data []byte) (*PersonV2View, error) {
	if len(data) < minLenPersonV2View {
		return nil, chitin.ErrWrongSize
	}
	view := &PersonV2View{
		data: data,
	}
	return view, nil
}

const (
	slotsLenPersonV2View  = 4
	numFieldsPersonV2View = 2
	minLenPersonV2View    = slotsLenPersonV2View + 1*numFieldsPersonV2View
)

type PersonV2View struct {
	data []byte
}

func (v *PersonV2View) Age() uint16 {
	data := v.data[0:2]
	return binary.BigEndian.Uint16(data)
}

func (v *PersonV2View) Siblings() uint16 {
	data := v.data[2:4]
	return binary.BigEndian.Uint16(data)
}

func (v *PersonV2View) Fields() (*PersonV2ViewFields, error) {
	f := &PersonV2ViewFields{}
	data := v.data[slotsLenPersonV2View:]

	{
		msg, next, err := chitinParseLengthPrefixed(data)
		if err != nil {
			return nil, err
		}
		p := (*reflect.StringHeader)(unsafe.Pointer(&f.fieldName))
		p.Data = uintptr(unsafe.Pointer(&msg[0]))
		p.Len = len(msg)
		data = next
	}

	{
		msg, next, err := chitinParseLengthPrefixed(data)
		if err != nil {
			return nil, err
		}
		p := (*reflect.StringHeader)(unsafe.Pointer(&f.fieldPhone))
		p.Data = uintptr(unsafe.Pointer(&msg[0]))
		p.Len = len(msg)
		data = next
	}

	return f, nil
}

type PersonV2ViewFields struct {
	fieldName string

	fieldPhone string
}

func (f *PersonV2ViewFields) Name() string {
	return f.fieldName
}

func (f *PersonV2ViewFields) Phone() string {
	return f.fieldPhone
}

func NewPersonV2Maker() *PersonV2Maker {
	maker := &PersonV2Maker{}
	return maker
}

type PersonV2Maker struct {
	slots [slotsLenPersonV2View]byte

	fieldName string

	fieldPhone string
}

func (m *PersonV2Maker) Bytes() []byte {
	// TODO what do we guarantee about immutability of return value?

	// TODO do this in just one allocation
	data := m.slots[:]

	{
		var lb [varuint.MaxUint64Len]byte
		var ll int

		ll = varuint.PutUint64(lb[:], uint64(len(m.fieldName))+1)
		data = append(data, lb[:ll]...)
		data = append(data, m.fieldName...)
	}

	{
		var lb [varuint.MaxUint64Len]byte
		var ll int

		ll = varuint.PutUint64(lb[:], uint64(len(m.fieldPhone))+1)
		data = append(data, lb[:ll]...)
		data = append(data, m.fieldPhone...)
	}

	return data
}

func (m *PersonV2Maker) SetAge(v uint16) {
	data := m.slots[0:2]
	binary.BigEndian.PutUint16(data, v)
}

func (m *PersonV2Maker) SetSiblings(v uint16) {
	data := m.slots[2:4]
	binary.BigEndian.PutUint16(data, v)
}

func (m *PersonV2Maker) SetName(v string) {
	m.fieldName = v
}

func (m *PersonV2Maker) SetPhone(v string) {
	m.fieldPhone = v
}
