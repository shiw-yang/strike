package main

import (
	"flag"
	"fmt"

	"github.com/shiw-yang/strike/cmd/protoc-gen-gin/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "0.0.1"

func main() {
	getVersion := flag.Bool("version", false, "show the version")
	flag.Parse()
	if *getVersion {
		fmt.Printf("protoc-gen-gin %v\n", version)
		return
	}
	var flags flag.FlagSet

	var paramErrCode string
	flags.StringVar(&paramErrCode, "code", "400", "the code of response data when param parse error")

	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			generator.GenerateFile(gen, f)
		}

		return nil
	})
}
