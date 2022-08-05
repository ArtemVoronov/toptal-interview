package task1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEqualChange(t *testing.T, expected []int, actual []int) {
	assert.Equal(t, len(expected), len(actual))
	for i := range expected {
		assert.Equal(t, expected[i], actual[i])
	}
}

func TestBasicCase1(t *testing.T) {
	input := 5.0
	price := 0.99
	expected := []int{1, 0, 0, 0, 0, 4}

	actual, err := getChange(input, price)

	assert.Nil(t, err)
	assertEqualChange(t, expected, actual)
}

func TestBasicCase2(t *testing.T) {
	input := 3.14
	price := 1.99
	expected := []int{0, 1, 1, 0, 0, 1}

	actual, err := getChange(input, price)

	assert.Nil(t, err)
	assertEqualChange(t, expected, actual)
}

func TestBasicCase3(t *testing.T) {
	input := 3.0
	price := 0.01
	expected := []int{4, 0, 2, 1, 1, 2}

	actual, err := getChange(input, price)

	assert.Nil(t, err)
	assertEqualChange(t, expected, actual)
}

func TestBasicCase4(t *testing.T) {
	input := 4.0
	price := 3.14
	expected := []int{1, 0, 1, 1, 1, 0}

	actual, err := getChange(input, price)

	assert.Nil(t, err)
	assertEqualChange(t, expected, actual)
}

func TestBasicCase5(t *testing.T) {
	input := 0.45
	price := 0.34
	expected := []int{1, 0, 1, 0, 0, 0}

	actual, err := getChange(input, price)

	assert.Nil(t, err)
	assertEqualChange(t, expected, actual)
}

func TestEmptyInput(t *testing.T) {
	input := 0.0
	price := 0.1
	expectedErrText := "wrong input"

	_, err := getChange(input, price)

	assert.NotNil(t, err)
	assert.Equal(t, expectedErrText, err.Error())
}

func TestEmptyPrice(t *testing.T) {
	input := 1.0
	price := 0.0
	expectedErrText := "wrong price"

	_, err := getChange(input, price)

	assert.NotNil(t, err)
	assert.Equal(t, expectedErrText, err.Error())
}

func TestNotEnoughtMoney(t *testing.T) {
	input := 1.0
	price := 2.0
	expectedErrText := "not enough money"

	_, err := getChange(input, price)

	assert.NotNil(t, err)
	assert.Equal(t, expectedErrText, err.Error())
}
