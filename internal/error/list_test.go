package error_test

import (
	"errors"
	"testing"

	error_list "github.com/instabledesign/go-skeleton/internal/error"
	"github.com/stretchr/testify/assert"
)

func TestList_Empty(t *testing.T) {
	list := &error_list.List{}

	assert.True(t, list.Empty())
	assert.Equal(t, "", list.Error())
}

func TestList_Add(t *testing.T) {
	list := &error_list.List{}
	list.Add(nil)
	list.Add(errors.New("my first error"))
	list.Add(errors.New("my second error"))

	assert.False(t, list.Empty())
	assert.Equal(t, "my first error\nmy second error\n", list.Error())
}
