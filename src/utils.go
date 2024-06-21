package forum

import (
	"strconv"
	"strings"
)

func (E *Engine) StrToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func (E *Engine) filterMsg(msg string) string {
	return strings.ReplaceAll(msg, "'", "[[apostroph]]")
}

func (E *Engine) reversefilterMsg(msg string) string {
	return strings.ReplaceAll(msg, "[[apostroph]]", "'")
}

