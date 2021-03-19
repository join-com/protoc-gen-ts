package generator

import "google.golang.org/protobuf/compiler/protogen"

func (r *Runner) collectData(protoFile *protogen.File) {
	r.packagesByFile[protoFile.Desc.Path()] = protoFile.Proto.GetPackage()
}
