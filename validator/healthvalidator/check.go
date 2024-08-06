package healthvalidator

import (
	"clean-code-structure/param/healthparam"
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	"clean-code-structure/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateCheckRequest(req healthparam.CheckRequest) error {
	const op = "messagevalidator.ValidateSendRequest"

	if err := validation.ValidateStruct(&req); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
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
