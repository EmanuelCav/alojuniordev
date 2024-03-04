package middleware

import (
	"context"
	"time"

	"github.com/EmanuelCav/alojuniordev/helper"
	"github.com/EmanuelCav/alojuniordev/models"
	"go.mongodb.org/mongo-driver/bson"
)

func RoleValid(role models.RoleCreateModel) string {

	var roleValid models.RoleModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := helper.ConnectionRole().FindOne(ctx, bson.M{"role": role.Role}).Decode(&roleValid)

	if err == nil {
		if roleValid.Role == role.Role {
			return "Role has been already registered"
		}
	}

	if !helper.ValidateString(role.Role) {
		return "Role only accepts letters"
	}

	return ""

}

func RegisterValid(user models.UserRegisterModel) string {

	var userValid models.UserModel

	if len(user.Password) <= 6 {
		return "La conntraseña debe contener más de 6 caracteres"
	}

	if user.Password != user.Confirm {
		return "Las contraseñas no coinciden"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := helper.ConnectionUser().FindOne(ctx, bson.M{"email": user.Email}).Decode(&userValid)

	if err == nil {
		if userValid.Email == user.Email {
			return "El email ya ha sido registrado"
		}

		if userValid.Username == user.Username {
			return "El nombre de usuario ya ha sido registrado"
		}
	}

	if len(user.Username) <= 2 {
		return "El nombre de usuario debe contener mas de dos caracteres"
	}

	if !helper.ValidateStringAndNumber(user.Username) {
		return "El nombre de usuario permite unicamente letras y números"
	}

	if !helper.ValidateEmail(user.Email) {
		return "Email no válido"
	}

	return ""

}

func LoginValid(user models.UserLoginModel) string {

	var userValid models.UserModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := helper.ConnectionUser().FindOne(ctx, bson.M{"email": user.Email}).Decode(&userValid)

	if err != nil {
		return "Los campos no coinciden"
	}

	if !helper.CompareHash(userValid.Password, user.Password) {
		return "Los campos no coinciden"
	}

	return ""
}
