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
		Name:                "auth",
		NamePrincipalFields: []string{},
		Fields:              a.fields,
		BackendHandler:      model.BackendHandler{},
		FrontendHandler:     model.FrontendHandler{},
	}

}
