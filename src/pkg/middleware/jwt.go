package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/idprm/go-alesse/src/pkg/config"
	"github.com/idprm/go-alesse/src/pkg/model"
)

func GenerateJWTToken(user model.User) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = exp

	t, err := token.SignedString([]byte(config.ViperEnv("JWT_SECRET_AUTH")))

	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}

func GenerateJWTTokenDoctor(doctor model.Doctor) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["doctor_id"] = doctor.ID
	claims["exp"] = exp

	t, err := token.SignedString([]byte(config.ViperEnv("JWT_SECRET_AUTH")))

	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}

func GenerateJWTTokenOfficer(officer model.Officer) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["officer_id"] = officer.ID
	claims["exp"] = exp

	t, err := token.SignedString([]byte(config.ViperEnv("JWT_SECRET_AUTH")))

	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}

func GenerateJWTTokenApothecary(apothecary model.Apothecary) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["apothecary_id"] = apothecary.ID
	claims["exp"] = exp

	t, err := token.SignedString([]byte(config.ViperEnv("JWT_SECRET_AUTH")))

	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}

func GenerateJWTTokenCourier(courier model.Courier) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["courier_id"] = courier.ID
	claims["exp"] = exp

	t, err := token.SignedString([]byte(config.ViperEnv("JWT_SECRET_AUTH")))

	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}

func GenerateJWTTokenSpecialist(specialist model.Specialist) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["specialist_id"] = specialist.ID
	claims["exp"] = exp

	t, err := token.SignedString([]byte(config.ViperEnv("JWT_SECRET_AUTH")))

	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}

func RefreshJWTToken() error {
	return nil
}
