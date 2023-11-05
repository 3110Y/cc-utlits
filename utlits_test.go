package utlits

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	a := []string{"1", "2", "3"}
	b := Map(a, func(f string) int {
		atoi, err := strconv.Atoi(f)
		if err != nil {
			return 0
		}
		return atoi
	})
	c := make([]int, len(a), len(a))
	for i, value := range a {
		atoi, err := strconv.Atoi(value)
		if err != nil {
			c[i] = atoi
			continue
		}
		c[i] = atoi
	}
	assert.Equal(t, c, b)
}

func TestPointer(t *testing.T) {
	a := "b"
	b := Pointer(a)
	assert.Equal(t, a, *b)
}

func TestNullSafeFunction(t *testing.T) {
	a := "1"
	b := NullSafeFunction(&a, func(t *string) *int {
		c := 31
		return &c
	})
	assert.Equal(t, 31, *b)
}

func TestValueOrDefault(t *testing.T) {
	a := 1
	b := ValueOrDefault(&a)
	assert.Equal(t, a, b)
	var d *int
	d = nil
	e := ValueOrDefault(&d)
	assert.Equal(t, d, e)
}
