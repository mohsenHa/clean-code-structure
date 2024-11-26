package healthvalidator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"clean-code-structure/param/healthparam"
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	"clean-code-structure/validator"
)

// Validator layer MUST return validator.Error

func (v Validator) ValidateCheckRequest(req healthparam.CheckRequest) error {
	const op = "messagevalidator.ValidateCheckRequest"

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

		return validator.Error{
			Fields: fieldErrors,
			Err: richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
				WithKind(richerror.KindInvalid).
				WithMeta(map[string]interface{}{"req": req}).WithErr(err),
		}
	}

	return nil
}
