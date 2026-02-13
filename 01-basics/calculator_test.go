package basics

import "testing"

func TestAdd(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3) = %d; want %d", result, expected)
	}
}

func TestAddNegativeNumbers(t *testing.T) {
	calc := Calculator{}
	result := calc.Add(-5, 3)
	expected := -2

	if result != expected {
		t.Errorf("Add(-5, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	calc := Calculator{}
	result := calc.Substract(10, 3)
	expected := 7

	if result != expected {
		t.Errorf("Subtract(10, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	calc := Calculator{}
	result := calc.Multiply(4, 5)
	expected := 20

	if result != expected {
		t.Errorf("Multiply(4, 5) = %d; want %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	calc := Calculator{}
	result, err := calc.Divide(10, 2)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 5
	if result != expected {
		t.Errorf("Divide(10, 2) = %d; want %d", result, expected)
	}
}

// TestDivideByZero tests error handling
func TestDivideByZero(t *testing.T) {
	calc := Calculator{}
	_, err := calc.Divide(10, 0)

	if err == nil {
		t.Fatal("expected error for division by zero, got nil")
	}

	expectedMsg := "division by zero"
	if err.Error() != expectedMsg {
		t.Errorf("error message = %q; want %q", err.Error(), expectedMsg)
	}
}

func TestPower(t *testing.T) {
	calc := Calculator{}
	result := calc.Power(2, 3)
	expected := 8

	if result != expected {
		t.Errorf("Power(2, 3) = %d; want %d", result, expected)
	}

	result = calc.Power(5, 0)
	expected = 1
	if result != expected {
		t.Errorf("Power(5, 0) = %d; want %d", result, expected)
	}

	result = calc.Power(7, 1)
	expected = 7
	if result != expected {
		t.Errorf("Power(7, 1) = %d; want %d", result, expected)
	}
}


func TestCalculatorMultipleOperations(t *testing.T){
	calc := Calculator{}

	sum := calc.Add(10,5)
	result := calc.Multiply(sum, 2)
	expected := 30

	if result != expected {
		t.Errorf("(10 + 5) * 2 = %d; want %d", result, expected)
	}
}