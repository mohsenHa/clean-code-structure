package tmpl_param

func GetSampleTmpl() string {
	return `package {{.}}param

import (
	"clean-code-structure/param"
)

type SampleRequest struct {
	param.BaseRequest
}

type SampleResponse struct {
	param.BaseResponse
	Message string ` + "`json:\"message\"`" + `
}
`
}
