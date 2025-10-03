package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam 1", 71, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF_Empty(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade for empty calculator to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

// UPDATED: Now checks the length of the single 'grades' slice.
func TestAddGrade_InvalidType(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	const Invalid GradeType = 99

	initial_length := len(gradeCalculator.grades) // Check single slice length

	gradeCalculator.AddGrade("bad grade", 100, Invalid)

	final_length := len(gradeCalculator.grades) // Check single slice length

	if final_length != initial_length {
		t.Errorf("AddGrade with invalid GradeType should not modify grade list. Expected length %d, got %d", initial_length, final_length)
	}
}

func TestGetGradeB_OneCategoryEmpty(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("A1", 100, Assignment)

	gradeCalculator.AddGrade("E1", 100, Exam)

	// Note: No Essay grade added. The logic relies on filterGrades returning an empty slice for Essay.
	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetFinalGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected Assignment.String() to be 'assignment', got %s", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected Exam.String() to be 'exam', got %s", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected Essay.String() to be 'essay', got %s", Essay.String())
	}

	const Unmapped GradeType = 99
	if Unmapped.String() != "" {
		t.Errorf("Expected unmapped GradeType.String() to be '', got %s", Unmapped.String())
	}
}
