package main

import (
	"flag"
	"fmt"
	"github.com/pseudomuto/protokit/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"os"
)

const version = "1.1.0"

var requireUnimplemented *bool

func f(gen *protogen.Plugin) error {
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		generateFile(gen, f)
	}
	return nil
}

func main() {

	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-grpc %v\n", version)
		return
	}

	var flags flag.FlagSet
	requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")

	fds, err := utils.LoadDescriptorSet("/Users/zhangjinglei/myapp/hello.pb")
	if err != nil {
		println(err.Error())
	}
	req := utils.CreateGenRequest(fds, "hello.proto")
	opts := protogen.Options{
		ParamFunc: flags.Set,
	}

	gen, err := opts.New(req)
	if err != nil {
		return
	}
	if err := f(gen); err != nil {
		// Errors from the plugin function are reported by setting the
		// error field in the CodeGeneratorResponse.
		//
		// In contrast, errors that indicate a problem in protoc
		// itself (unparsable input, I/O errors, etc.) are reported
		// to stderr.
		gen.Error(err)
	}
	resp := gen.Response()
	out, err := proto.Marshal(resp)
	if err != nil {
		return
	}
	if _, err := os.Stdout.Write(out); err != nil {
		return
	}

	//
	//showVersion := flag.Bool("version", false, "print the version and exit")
	//flag.Parse()
	//if *showVersion {
	//	fmt.Printf("protoc-gen-go-grpc %v\n", version)
	//	return
	//}
	//
	//var flags flag.FlagSet
	//requireUnimplemented = flags.Bool("require_unimplemented_servers", true, "set to false to match legacy behavior")
	//
	//protogen.Options{
	//	ParamFunc: flags.Set,
	//}.Run(func(gen *protogen.Plugin) error {
	//	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	//	for _, f := range gen.Files {
	//		if !f.Generate {
	//			continue
	//		}
	//		generateFile(gen, f)
	//	}
	//	return nil
	//})
}

//package main
//
//import (
//	"eric/internal/thirdpart/embedswagger"
//	"github.com/labstack/echo/v4"
//)
//
//func main() {
//	//codeup.ListRepositoryTree(602072,"dev","public")
//	e := echo.New()
//
//	e.GET("/swagger/*", embedswagger.WrapHandler)
//	e.GET("/swagger/doc",test)
//
//
//	e.Logger.Fatal(e.Start(":1323"))
//}
//
//func test(c echo.Context) error {
//
//	a:=`basePath: /v2
//definitions:
//  web.APIError:
//    properties:
//      createdAt:
//        type: string
//      errorCode:
//        type: integer
//      errorMessage:
//        type: string
//    type: object
//  web.Pet:
//    properties:
//      category:
//        properties:
//          id:
//            type: integer
//          name:
//            type: string
//        type: object
//      id:
//        type: integer
//      name:
//        type: string
//      photoUrls:
//        items:
//          type: string
//        type: array
//      status:
//        type: string
//      tags:
//        items:
//          $ref: '#/definitions/web.Tag'
//        type: array
//    type: object
//  web.RevValue:
//    properties:
//      Data:
//        type: integer
//      Err:
//        type: integer
//      Status:
//        type: boolean
//    type: object
//  web.Tag:
//    properties:
//      id:
//        type: integer
//      name:
//        type: string
//    type: object
//host: petstore.swagger.io
//info:
//  contact:
//    email: support@swagger.io
//    name: API Support
//    url: http://www.swagger.io/support
//  description: This is a sample server Petstore server.
//  license:
//    name: Apache 2.0
//    url: http://www.apache.org/licenses/LICENSE-2.0.html
//  termsOfService: http://swagger.io/terms/
//  title: Swagger Example API
//  version: "1.0"
//paths:
//  /file/upload:
//    post:
//      consumes:
//      - multipart/form-data
//      description: Upload file
//      operationId: file.upload
//      parameters:
//      - description: this is a test file
//        in: formData
//        name: file
//        required: true
//        type: file
//      produces:
//      - application/json
//      responses:
//        "200":
//          description: ok
//          schema:
//            type: string
//        "400":
//          description: We need ID!!
//          schema:
//            $ref: '#/definitions/web.APIError'
//        "404":
//          description: Can not find ID
//          schema:
//            $ref: '#/definitions/web.APIError'
//      summary: Upload file
//  /testapi/get-string-by-int/{some_id}:
//    get:
//      consumes:
//      - application/json
//      description: get string by ID
//      operationId: get-string-by-int
//      parameters:
//      - description: Some ID
//        in: path
//        name: some_id
//        required: true
//        type: integer
//      - description: Some ID
//        in: body
//        name: some_id
//        required: true
//        schema:
//          $ref: '#/definitions/web.Pet'
//          type: object
//      produces:
//      - application/json
//      responses:
//        "200":
//          description: ok
//          schema:
//            type: string
//        "400":
//          description: We need ID!!
//          schema:
//            $ref: '#/definitions/web.APIError'
//        "404":
//          description: Can not find ID
//          schema:
//            $ref: '#/definitions/web.APIError'
//      summary: Add a new pet to the store
//  /testapi/get-struct-array-by-string/{some_id}:
//    get:
//      consumes:
//      - application/json
//      description: get struct array by ID
//      operationId: get-struct-array-by-string
//      parameters:
//      - description: Some ID
//        in: path
//        name: some_id
//        required: true
//        type: string
//      - description: Offset
//        in: query
//        name: offset
//        required: true
//        type: integer
//      - description: Offset
//        in: query
//        name: limit
//        required: true
//        type: integer
//      produces:
//      - application/json
//      responses:
//        "200":
//          description: ok
//          schema:
//            type: string
//        "400":
//          description: We need ID!!
//          schema:
//            $ref: '#/definitions/web.APIError'
//        "404":
//          description: Can not find ID
//          schema:
//            $ref: '#/definitions/web.APIError'
//swagger: "2.0"`
//	return c.String(200,a)
//}
