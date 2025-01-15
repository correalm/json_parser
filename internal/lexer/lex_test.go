package lexer_test

import (
	"json_parser/internal/lexer"
	"reflect"
	"testing"
	"strings"
)

func TestLexer(t *testing.T) {
  t.Run("returns the correct tokens", func(t *testing.T) {
    result := lexer.Lex("\"{\"foo\":\"bar\",\"bar\":\"foo\", \"number\": 4}\"")
    expected := []string{"{", "foo", ":", "bar", ",", "bar", ":", "foo", "number", "4", "}"}

    if !reflect.DeepEqual(result, expected) {
      t.Fatalf("Expect %s, got %s", expected, result)
    }
  })
}

func TestLexString(t *testing.T) {
	t.Run("returns the correct result", func(t *testing.T) {
		result, index, _ := lexer.LexString("\"foo_bar\"")

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
		result, index, err := lexer.LexString("\"foo_bar")

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

func TestLexInteger(t *testing.T) {
  t.Run("returns the correct value", func(t *testing.T) {
    result, index := lexer.LexInteger("250,")

    if result != 250 {
      t.Fatalf("Expected 250, got %d", result)
    }

    if index != 2 {
      t.Fatalf("Expected index 2, got %d", index)
    }
  })

  t.Run("returns the correct value for negative values", func(t *testing.T) {
    result, index := lexer.LexInteger("-45 ")

    if result != -45 {
      t.Fatalf("Expected -45, got %d", result)
    }

    if index != 2 {
      t.Fatalf("Expected index 2, got %d", index)
    }
  })
}

func TestLexFloat(t *testing.T) {
  t.Run("returns the correct value", func(t *testing.T) {
    result, index := lexer.LexFloat("22.2 ")
    expected := 22.2

    if result != expected {
      t.Fatalf("Expected %f, got %f", expected, result)
    }

    if index != 3 {
      t.Fatalf("Expected index 3, got %d", index)
    }
  })

  t.Run("returns the correct value for negative values", func(t *testing.T) {
    result, index := lexer.LexFloat("-789.08,")
    expected := -789.08

    if result != expected {
      t.Fatalf("Expected %f, got %f", expected, result)
    }

    if index != 6 {
      t.Fatalf("Expected index 6, got %d", index)
    }
  })
}
