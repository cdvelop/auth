package auth

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

func (a *Auth) buildObject() {

	if len(a.fields) == 0 {

		fields := []model.Field{
			{Name: "mail", Legend: "Mail", Input: input.Mail()},
			{Name: "password", Legend: "Contraseña", Input: input.Password()},
		}

		a.fields = append(a.fields, fields...)

	}

	a.object = model.Object{
		ObjectName:          "auth",
		PrincipalFieldsName: []string{},
		Fields:              a.fields,
		BackHandler:         model.BackendHandler{},
		FrontHandler:        model.FrontendHandler{},
	}

}
