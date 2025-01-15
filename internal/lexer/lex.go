package lexer

import (
	"errors"
	"regexp"
	"strconv"
)

const (
  JSON_QUOTE = '"'
)

func Lex(raw string) []string {
  raw_len := len(raw)

  result := make([]string, raw_len)

  // percorrer cada um dos caracteres do input
  // identificar o que eles representam
  // adicionar ao resultado a unidade de valor
  
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

func isValidNumberChar(c byte) bool {
  matched, err := regexp.Match(`[\d\.\-\+]`, []byte{c})

  if err != nil {
    return false
  }

  return matched
}

func LexInteger(raw string) (int64, int) {
  result := []byte{}
  index := 0

  for ; index < len(raw); index++ {
    current := raw[index]

    if !isValidNumberChar(current) { break }

    result = append(result, current)
  }

  int_value, err := strconv.ParseInt(string(result), 0, 64)

  if err != nil {
    return 0, 0
  }

  return int_value, index - 1
}

// TODO: validate a better way to treat errors
func LexFloat(raw string) (float64, int) {
  result := []byte{}
  index := 0

  for ; index < len(raw); index++ {
    current := raw[index]

    if !isValidNumberChar(current) { break }

    result = append(result, current)
  }

  float_value, err := strconv.ParseFloat(string(result), 64)

  if err != nil {
    return 0, 0
  }

  return float_value, index - 1
}

func LexNull(raw string) string {
  return ""
}
