// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AgAA//8pLjfFslIAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
