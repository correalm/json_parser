package parser

import (
	"fmt"
)

func Parser(raw string) map[string]any {
  fmt.Println("len :: ", len(raw), raw[27])

  result := map[string]any{
    "foo": "bar",
    "bar":"foo",
  }

  return result
}
