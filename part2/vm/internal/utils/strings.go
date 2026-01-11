package utils

import (
	"fmt"
	"strings"
)

func WriteSBf(sb *strings.Builder, str string, args ...any) *strings.Builder {
	sb.WriteString(fmt.Sprintf(str+"\n", args...))
	return sb
}
