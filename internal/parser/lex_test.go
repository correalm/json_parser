package parser_test

import (
	"json_parser/internal/parser"
	"reflect"
	"testing"
	"strings"
)

func TestLex(t *testing.T) {
  t.Run("returns the correct tokens", func(t *testing.T) {
    result := parser.Lex("\"{\"foo\":\"bar\",\"bar\":\"foo\"}\"")
    expected := []string{"{", "foo", ":", "bar", ",", "bar", ":", "foo", "}"}

    if !reflect.DeepEqual(result, expected) {
      t.Fatalf("Expect %s, got %s", expected, result)
    }
  })
}

func TestLexString(t *testing.T) {
	t.Run("returns the correct result", func(t *testing.T) {
		result, index, _ := parser.LexString("\"foo_bar\"")

		expected := "\"foo_bar\""
    expected_idex := 8

		if strings.Compare(result, expected) != 0 {
			t.Fatalf("Expected %s, got %s", expected, result)
		}

    if index != expected_idex {
      t.Fatalf("Expected index %d, got %d", expected_idex, index)
    }
	})

	t.Run("returns an error when the passed string is no terminated", func(t *testing.T) {
		result, index, err := parser.LexString("\"foo_bar")

		if result != "" {
			t.Fatalf("Expected empty string, got %s", result)
		}

    if index != 0 {
      t.Fatalf("Expected index 0, got %d", index)
    }

    if err == nil {
      t.Fatalf("Expected an error but got nil")
    }

    expected_error_message := "Unterminated string"
    error_message := err.Error()

    if strings.Compare(error_message, expected_error_message) != 0 {
      t.Fatalf("Expected %s, got %s error message", expected_error_message, error_message)
    }
	})
}
