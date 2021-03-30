package generator

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func (r *Runner) P(gf *protogen.GeneratedFile, v ...string) {
	indentation := strings.Repeat(" ", r.indentLevel)
	for _, vv := range v {
		lines := strings.Split(vv, "\n")
		for _, line := range lines {
			gf.P(fmt.Sprintf("%s%s", indentation, line))
		}
	}
}
