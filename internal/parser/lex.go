package parser

import "errors"

const (
  JSON_QUOTE = '"'
)

func Lex(raw string) []string {
  raw_len := len(raw)

  result := make([]string, raw_len)
  
  return result
}

func LexString(raw string) (string, int, error) {
  result := []byte{}
  index := 0

  if raw[0] == JSON_QUOTE {
    for i := 0; i < len(raw); i++ {
      result = append(result, raw[i])
      index = i

      if raw[i] == JSON_QUOTE && i > 0 {
        break
      }
    }
  }

  if result[len(result) - 1] != '"' {
    return "", 0, errors.New("Unterminated string")
  }

  return string(result), index, nil
}

func LexBool(raw string) string {
  return ""
}

func LexNumber(raw string) string {
  return ""
}

func LexNull(raw string) string {
  return ""
}
