package tools_test

import (
	"testing"

	"github.com/shompys/go-2025/tools"
	"github.com/stretchr/testify/assert"
)

func TestElementAtIndex(t *testing.T) {

	//arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 2
	expected := 3
	//act
	result, _ := tools.ElementAtIndex(slice, index)
	//assert
	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
		return
	}

	t.Log("success")

}

func TestElementAtIndex_NonExistingIndex(t *testing.T) {

	slice := []int{1, 2, 3, 4, 5}
	idx := len(slice)

	result, err := tools.ElementAtIndex(slice, idx)

	// Verifica que err NO sea nil (que sí haya error)
	assert.NotNil(t, err)

	// O también puedes verificar que sea un error específico
	assert.Error(t, err) // Verifica que err != nil

	// Y si quieres verificar que result sea el zero value
	assert.Equal(t, 0, result, "result should be 0")

}
