package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldAddTwoNumbersGivenOption1(t *testing.T) {
	expected := float64(7)
	got, _ := Calculate(3, 4, 1)
	assert.Equal(t, expected, got)
}

func TestShouldSubtractTwoNumbersGivenOption2(t *testing.T) {
	expected := float64(-1)
	got, _ := Calculate(3, 4, 2)
	assert.Equal(t, expected, got)
}

func TestShouldMultiplyTwoNumbersGivenOption3(t *testing.T) {
	expected := float64(15)
	got, _ := Calculate(3, 5, 3)
	assert.Equal(t, expected, got)
}

func TestShouldDivideTwoNumbersGivenOption4AndNonZeroSecondOperand(t *testing.T) {
	expected := 2.5
	got, _ := Calculate(5, 2, 4)
	assert.Equal(t, expected, got)
}

func TestShouldReturnErrorGivenOption4AndZeroSecondOperand(t *testing.T) {
	expected := "division by zero error"
	_, got := Calculate(5, 0, 4)
	assert.Equal(t, expected, got.Error())
}

func TestShouldFindPowerOfTwoNumbersGivenOption5(t *testing.T) {
	expected := float64(25)
	got, _ := Calculate(5, 2, 5)
	assert.Equal(t, expected, got)
}
func TestShouldReturnErrorGivenInvalidOption(t *testing.T) {
	expected := "invalid option"
	_, got := Calculate(5, 0, 6)
	assert.Equal(t, expected, got.Error())
}
