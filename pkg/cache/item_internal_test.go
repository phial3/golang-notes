package cache

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItem_Bool(t *testing.T) {
	decoder := stringDecoder{}
	tests := []struct {
		item   Item
		ok     bool
		expect bool
	}{
		{Item{decoder: decoder, value: []byte("1")}, true, true},
		{Item{decoder: decoder, err: errors.New("")}, false, false},
	}

	for i, tt := range tests {
		got, err := tt.item.Bool()
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				assert.FailNow(t, "test %d, unexpected failure: %v", i, err)
			} else {
				assert.FailNow(t, "test %d, unexpected success", i)
			}
		}

		assert.Equal(t, tt.expect, got)
	}
}

func TestItem_Bytes(t *testing.T) {
	tests := []struct {
		item   Item
		ok     bool
		expect []byte
	}{
		{Item{value: []byte{0x01}}, true, []byte{0x01}},
		{Item{err: errors.New("")}, false, nil},
	}

	for i, tt := range tests {
		got, err := tt.item.Bytes()
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				assert.FailNow(t, "test %d, unexpected failure: %v", i, err)
			} else {
				assert.FailNow(t, "test %d, unexpected success", i)
			}
		}

		assert.Equal(t, tt.expect, got)
	}
}

func TestItem_String(t *testing.T) {
	tests := []struct {
		item   Item
		ok     bool
		expect string
	}{
		{Item{value: []byte("hello")}, true, "hello"},
		{Item{err: errors.New("")}, false, ""},
	}

	for i, tt := range tests {
		got, err := tt.item.String()
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				assert.FailNow(t, "test %d, unexpected failure: %v", i, err)
			} else {
				assert.FailNow(t, "test %d, unexpected success", i)
			}
		}

		assert.Equal(t, tt.expect, got)
	}
}

func TestItem_Int64(t *testing.T) {
	decoder := stringDecoder{}
	tests := []struct {
		item   Item
		ok     bool
		expect int64
	}{
		{Item{decoder: decoder, value: []byte("1")}, true, 1},
		{Item{decoder: decoder, value: []byte("a")}, false, 0},
		{Item{decoder: decoder, err: errors.New("")}, false, 0},
	}

	for i, tt := range tests {
		got, err := tt.item.Int64()
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				assert.FailNow(t, "test %d, unexpected failure: %v", i, err)
			} else {
				assert.FailNow(t, "test %d, unexpected success", i)
			}
		}

		assert.Equal(t, tt.expect, got)
	}
}

func TestItem_Uint64(t *testing.T) {
	decoder := stringDecoder{}
	tests := []struct {
		item   Item
		ok     bool
		expect uint64
	}{
		{Item{decoder: decoder, value: []byte("1")}, true, 1},
		{Item{decoder: decoder, value: []byte("a")}, false, 0},
		{Item{decoder: decoder, err: errors.New("")}, false, 0},
	}

	for i, tt := range tests {
		got, err := tt.item.Uint64()
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				assert.FailNow(t, "test %d, unexpected failure: %v", i, err)
			} else {
				assert.FailNow(t, "test %d, unexpected success", i)
			}
		}

		assert.Equal(t, tt.expect, got)
	}
}

func TestItem_Float64(t *testing.T) {
	decoder := stringDecoder{}
	tests := []struct {
		item   Item
		ok     bool
		expect float64
	}{
		{Item{decoder: decoder, value: []byte("1.2")}, true, 1.2},
		{Item{decoder: decoder, value: []byte("a")}, false, 0},
		{Item{decoder: decoder, err: errors.New("")}, false, 0},
	}

	for i, tt := range tests {
		got, err := tt.item.Float64()
		if ok := (err == nil); ok != tt.ok {
			if err != nil {
				assert.FailNow(t, "test %d, unexpected failure: %v", i, err)
			} else {
				assert.FailNow(t, "test %d, unexpected success", i)
			}
		}

		assert.Equal(t, tt.expect, got)
	}
}

func TestItem_Err(t *testing.T) {
	expect := errors.New("err")
	tests := []struct {
		item   Item
		expect error
	}{
		{Item{}, nil},
		{Item{err: expect}, expect},
	}

	for _, tt := range tests {
		err := tt.item.Err()

		assert.Equal(t, tt.expect, err)
	}
}
