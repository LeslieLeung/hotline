package workflow

import (
	"testing"
)

func TestGetValueByPath_ReturnsValueForValidKey(t *testing.T) {
	data := map[string]interface{}{
		"step1": map[string]interface{}{
			"outputs": map[string]interface{}{
				"result": "success",
			},
		},
	}

	value, err := getValueByPath(data, "step1.outputs.result")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if value != "success" {
		t.Errorf("Expected 'success', got %v", value)
	}
}

func TestGetValueByPath_ReturnsErrorForInvalidKey(t *testing.T) {
	data := map[string]interface{}{
		"step1": map[string]interface{}{
			"outputs": map[string]interface{}{
				"result": "success",
			},
		},
	}

	_, err := getValueByPath(data, "step1.outputs.missing")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestGetValueByPath_ReturnsErrorForPartialKeyMatch(t *testing.T) {
	data := map[string]interface{}{
		"step1": map[string]interface{}{
			"outputs": map[string]interface{}{},
		},
	}

	_, err := getValueByPath(data, "step1.outputs.result")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestGetValueByPath_ReturnsErrorForNonMapIntermediate(t *testing.T) {
	data := map[string]interface{}{
		"step1": "not a map",
	}

	_, err := getValueByPath(data, "step1.outputs.result")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
