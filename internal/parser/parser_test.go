package parser_test

import (
	"json_parser/internal/parser"
	"reflect"
	"testing"
)

func TestParser (t *testing.T) {
  result := parser.Parser("\"{\"foo\":\"bar\",\"bar\":\"foo:\"}\"")

  expect := map[string]any {
    "foo": "bar",
    "bar": "foo",
  }

  if !reflect.DeepEqual(result, expect) {
    t.Fatal("Oh men")
  }
}
