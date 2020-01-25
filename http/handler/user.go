package handler

import (
	"context"
	"fmt"
	pb "github.com/ArtGooner/test-microservice/config"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	cl pb.userServiceClient
}

func New(cl pb.userServiceClient) *UserHandler {
	return &UserHandler{cl: cl}
}

func (u UserHandler) Login(c echo.Context) error {
	email := c.QueryParam("email")
	password := c.QueryParam("password")

	acc := &pb.Account{Email: email, Password: password}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Authenticate(ctx, acc)

	if err != nil {
		log.Fatalf("auth error: %v", err)
	}

	log.Printf("Authenticated: %v", r)

	return c.String(http.StatusOK, fmt.Sprintf("Authenticated: %v", r))
}

//// Create the user
//func (u UserHandler) Create(c echo.Context) error {
//	email := c.QueryParam("email")
//	password := c.QueryParam("password")
//	h := sha256.New()
//	h.Write([]byte(password))
//	usr := &user.User{Email: email, PasswordHash: h.Sum(nil)}
//	err := u.rep.Create(usr)
//	if err != nil {
//		return err
//	}
//	return c.String(http.StatusOK, "Ok")
//}
//
////Get the user
//func (u UserHandler) Get(c echo.Context) error {
//	email := c.QueryParam("email")
//	password := c.QueryParam("password")
//	h := sha256.New()
//	h.Write([]byte(password))
//	usr := &user.User{Email: email, PasswordHash: h.Sum(nil)}
//	res, err := u.rep.Get(usr)
//	if err != nil {
//		return err
//	}
//	if res == nil {
//		return fmt.Errorf("not found")
//	}
//	return c.String(http.StatusOK, res.Email)
//}
