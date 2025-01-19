package lexer

import (
	"errors"
	"regexp"
	"strconv"
)

const (
  QUOTE = '"'
  LEFT_BRACE = '{'
  RIGTH_BRACE = '}'
  LEFT_BRACKET = '['
  RIGTH_BRACKET = ']'
  COLON = ':'
  COMMA = ','
  SPACE = ' '
)

func Lex(raw string) []string {
  raw_len := len(raw)

  result := make([]string, raw_len)

  // percorrer cada um dos caracteres do input
  // identificar o que eles representam
  // adicionar ao resultado a unidade de valor
  
  return result

func IsQuote(token byte) bool {
  return token == QUOTE
}

func IsDelim(token byte) bool {
  return (
    token == COLON ||
    token == COMMA ||
    token == LEFT_BRACE ||
    token == RIGTH_BRACE ||
    token == LEFT_BRACKET ||
    token == RIGTH_BRACKET ||
    IsQuote(token)) && token != SPACE
}

func LexString(raw string) (string, int, error) {
  result := []byte{}
  index := 0

  for index := 0; index < len(raw); index++ {
    result = append(result, raw[index])

    if IsQuote(raw[index]) {
      break
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
