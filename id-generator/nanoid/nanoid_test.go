package nanoid_test

import (
	"fmt"
	"testing"
)

import (
	"github.com/phial3/id-generator/nanoid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStandard(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.Standard(21)
		assert.NoError(t, err, "should be no error")
		id := f()
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})

	t.Run("negative ID length", func(t *testing.T) {
		_, err := nanoid.Standard(-1)
		assert.Error(t, err, "should error if passed ID length is negative")
	})

	t.Run("invalid length (256)", func(t *testing.T) {
		_, err := nanoid.Standard(256)
		assert.Error(t, err, "should error if length > 255")
	})

	t.Run("invalid length (1)", func(t *testing.T) {
		_, err := nanoid.Standard(1)
		assert.Error(t, err, "should error if length < 2")
	})
}

func TestCustom(t *testing.T) {
	t.Run("general", func(t *testing.T) {
		f, err := nanoid.CustomASCII("abcdef", 21)
		assert.NoError(t, err, "should be no error")
		id := f()
		assert.Len(t, id, 21, "should return the same length as the ID specified length")
		t.Log(id)
	})
}

func TestFlatDistribution(t *testing.T) {
	tries := 500_000

	set := "0123456789" // 10.
	length := len(set)
	hits := make(map[rune]int)

	f, err := nanoid.CustomASCII(set, length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < tries; i++ {
		id := f()
		for _, r := range id {
			hits[r]++
		}
	}

	for _, count := range hits {
		require.InEpsilon(t, length*tries/len(set), count, 0.01, "should have flat-distribution")
	}
}

func TestCollisions(t *testing.T) {
	tries := 500_000

	used := make(map[string]bool)
	f, err := nanoid.Standard(8)
	if err != nil {
		panic(err)
	}

	for i := 0; i < tries; i++ {
		id := f()
		require.False(t, used[id], "shouldn't be any colliding IDs")
		used[id] = true
	}
}

func Benchmark8NanoID(b *testing.B) {
	f, err := nanoid.Standard(8)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func Benchmark21NanoID(b *testing.B) {
	f, err := nanoid.Standard(21)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func Benchmark36NanoID(b *testing.B) {
	f, err := nanoid.Standard(36)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func Benchmark255NanoID(b *testing.B) {
	f, err := nanoid.Standard(255)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func BenchmarkCustomUnicodeNanoID(b *testing.B) {
	f, err := nanoid.CustomASCII("°Ô‘š±?¿⾃ⶃⵏ⟎⸸ⵌ", 8)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}

func BenchmarkCustomASCIINanoID(b *testing.B) {
	f, err := nanoid.CustomASCII("0123456789", 10)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		fmt.Println(f())
	}
}

func BenchmarkASCIINanoID(b *testing.B) {
	f, err := nanoid.ASCII(21)
	if err != nil {
		panic(err)
	}

	for n := 0; n < b.N; n++ {
		f()
	}
}
