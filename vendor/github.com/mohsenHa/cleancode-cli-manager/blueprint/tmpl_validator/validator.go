package tmpl_validator

func GetSampleTmpl() string {
	return `package {{.}}validator

import (
	"clean-code-structure/param/{{.}}param"
	"clean-code-structure/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Validator layer MUST return validator.Error

func (v Validator) ValidateSampleRequest(req {{.}}param.SampleRequest) error {
	const op = "{{.}}validator.ValidateSampleRequest"
	
	_, span := otela.TraceBuilder(
		"{{.}}validator",
		"ValidateSampleRequest",
		otela.WithContext(req.Ctx),
	)

	if err := validation.ValidateStruct(&req); err != nil {
		fieldErrors := make(map[string]string)

		var errV validation.Errors
		if errors.As(err, &errV) {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		span.AddEvent("request not valid")
		err = validator.Error{
			Fields: fieldErrors,
			Err: richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
				WithKind(richerror.KindInvalid).
				WithMeta(map[string]interface{}{"req": req}).WithErr(err),
		}
		span.RecordError(err)
		return err
	}

	return nil
}
`
}

func GetValidatorTmpl() string {
	return `package {{.}}validator

type Validator struct {
}

func New() Validator {
	return Validator{}
}
`
}
