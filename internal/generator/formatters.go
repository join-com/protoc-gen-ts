package generator

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func (r *Runner) P(gf *protogen.GeneratedFile, v ...interface{}) {
	indentation := strings.Repeat(" ", r.indentLevel)
	for _, vv := range v {
		gf.P(fmt.Sprintf("%s%s", indentation, vv))
	}
}
