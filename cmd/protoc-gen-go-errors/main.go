package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var version = flag.Bool("version", false, "show the version")

func main() {
	// show versions
	flag.Parse()
	if *version {
		fmt.Println("protoc-gen-go-errors " + release)
		return
	}

	// gen errors
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(p *protogen.Plugin) error {
		p.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range p.Files {
			if !f.Generate {
				continue
			}
			generateFile(p, f)
		}
		return nil
	})
}

func generateFile(p *protogen.Plugin, f *protogen.File) {
	panic("unimplemented")
}
