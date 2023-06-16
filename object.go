package auth

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

func (a *auth) Object() model.Object {

	if len(a.fields) == 0 {

		fields := []model.Field{
			{Name: "mail", Legend: "Mail", Input: input.Mail()},
			{Name: "password", Unique: true, Legend: "Contraseña", Input: input.Password()},
		}

		a.fields = append(a.fields, fields...)

	}

	o := model.Object{
		Name:           "auth",
		TextFieldNames: []string{},
		Fields:         a.fields,
	}

	return o

}
