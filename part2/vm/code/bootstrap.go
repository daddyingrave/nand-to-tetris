package code

import (
	"fmt"
	"strings"
	"vm/internal/utils"
)

func BootstrapCode(fileName string) (string, error) {
	sb := &strings.Builder{}

	utils.WriteSBf(sb, "// bootstrap start")
	utils.WriteSBf(sb, "")

	utils.WriteSBf(sb, "  @256")
	utils.WriteSBf(sb, "  D=A")
	utils.WriteSBf(sb, "  @SP")
	utils.WriteSBf(sb, "  M=D")

	callCode, err := Call(fileName, "bootstrap", "Sys.init", 0, 1)
	if err != nil {
		return "", fmt.Errorf("fail to generate call code for bootstrap %w", err)
	}

	utils.WriteSBf(sb, callCode)
	utils.WriteSBf(sb, "")
	utils.WriteSBf(sb, "// bootstrap ends")
	utils.WriteSBf(sb, "")

	return sb.String(), nil
}
