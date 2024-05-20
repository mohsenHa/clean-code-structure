package healthvalidator

import (
	"clean-code-structure/param/healthparam"
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateCheckRequest(req healthparam.CheckRequest) (map[string]string, error) {
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

		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}

	return nil, nil
}
