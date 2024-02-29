// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// LoginRequest defines model for LoginRequest.
type LoginRequest struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

// LoginResponse defines model for LoginResponse.
type LoginResponse struct {
	Id       *int    `json:"id,omitempty"`
	JwtToken *string `json:"jwt_token,omitempty"`
}

// RegistrationRequest defines model for RegistrationRequest.
type RegistrationRequest struct {
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

// RegistrationResponse defines model for RegistrationResponse.
type RegistrationResponse struct {
	Id *int `json:"id,omitempty"`
}

// UserInfoRequest defines model for UserInfoRequest.
type UserInfoRequest struct {
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

// UserInfoResponse defines model for UserInfoResponse.
type UserInfoResponse struct {
	FullName    *string `json:"full_name,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// GetUserIdParams defines parameters for GetUserId.
type GetUserIdParams struct {
	Authorization string `json:"Authorization"`
}

// PatchUserIdParams defines parameters for PatchUserId.
type PatchUserIdParams struct {
	Authorization string `json:"Authorization"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = LoginRequest

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = RegistrationRequest

// PatchUserIdJSONRequestBody defines body for PatchUserId for application/json ContentType.
type PatchUserIdJSONRequestBody = UserInfoRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// User login.
	// (POST /login)
	PostLogin(ctx echo.Context) error
	// Register user.
	// (POST /register)
	PostRegister(ctx echo.Context) error
	// Get user info.
	// (GET /user/{id})
	GetUserId(ctx echo.Context, id string, params GetUserIdParams) error
	// Update user profile.
	// (PATCH /user/{id})
	PatchUserId(ctx echo.Context, id string, params PatchUserIdParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostLogin(ctx)
	return err
}

// PostRegister converts echo context to params.
func (w *ServerInterfaceWrapper) PostRegister(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostRegister(ctx)
	return err
}

// GetUserId converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserIdParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for Authorization, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Authorization: %s", err))
		}

		params.Authorization = Authorization
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter Authorization is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserId(ctx, id, params)
	return err
}

// PatchUserId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PatchUserIdParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "Authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for Authorization, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, valueList[0], &Authorization)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Authorization: %s", err))
		}

		params.Authorization = Authorization
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter Authorization is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUserId(ctx, id, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/login", wrapper.PostLogin)
	router.POST(baseURL+"/register", wrapper.PostRegister)
	router.GET(baseURL+"/user/:id", wrapper.GetUserId)
	router.PATCH(baseURL+"/user/:id", wrapper.PatchUserId)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RVTW+bQBD9K2jaIzJumxO3Vq0iq59y01MURRsYYFPY2cwOiVKL/17tYmo7xnUj2ZJ7",
	"g2V23sx7M48FZNRYMmjEQboAl1XYqPD4gZl4js6ScegPLJNFFo3hc4POqTJ8kEeLkIIT1qaErouB8a7V",
	"jDmkl38Cr+IhkG5uMRPoYvhEpTZzvGvRyTaEVc49EOcjGDHYigxem7a5Qd5fxEZ0vMr8l6J2Na7X69FG",
	"sET2124f5FroJ5rxarZg5lhqJ6xE024Kiraur41qcJyDoxG0wt1D1mYXz+NsjJUfDnlmCvp/GVl1sIuN",
	"PS3srfIJpD/SpiAfXOsMl5g9AHyeXfisoqX2r7666Dvyvc58J/fITpOBFF5NppOpjySLRlkNKbwJR75d",
	"qULhSe1XIzREvTa+raD9LIcUvpGTsD3QU4lO3lH+6AMzMoIm3FHW1joLt5JbR2blO/7pJWMBKbxIVsaU",
	"LF0p2bCLrusV61kO9b2eTg+NtdQwYOXoMtZWesK+fvRsnU3PDga56bgjkF9IooJak4cxcG3TKH4cRA3a",
	"TMKXhMNW9vOzW6r5EHUctcYM7siijbrRaWs3iBC1Dnkpn39MFjrvPGqJI/KdowSfyeGIbG452UhbF+gk",
	"YpSWzclQeo4S2Iy8K0763wKrBgXZQXq5eJJm9j6iIlwAb6SQBsODeHBQncP6n0G4xXithS1/XvRJKlR5",
	"SLlM87aVilj/ClQ8K+NVsOCsGtlif7w2CIdf4qf/4yMv8L+M3CkZr82VYD9slqnQNU76+w75fhi3lms/",
	"DyI2TZKaMlVX3pO7q+53AAAA//9PM7eWfwsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
