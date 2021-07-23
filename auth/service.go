package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SCREET_KEY = []byte("B15mill4h_K3y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {

	Expired := time.Now().Add(time.Hour * 24 * 7).Unix()
	fmt.Println(Expired)

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = Expired

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SCREET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) { // validasi token
	// masukan token , parametetrnya adalah func lalu mengembalikan interface dan err
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		//cek tokennya
		if !ok {
			return nil, errors.New("Iinvalid token")
		}

		return []byte(SCREET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil

}
