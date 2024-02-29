package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// (POST /register)
func (s *Server) PostRegister(ctx echo.Context) error {
	req := new(generated.RegistrationRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	id, err := s.Repository.RegisterUser(ctx.Request().Context(), repository.RegisterUserInput{
		PhoneNumber: req.PhoneNumber,
		FullName:    req.FullName,
		Password:    string(hashedPassword),
	})
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	var resp generated.RegistrationResponse
	resp.Id = &id
	return ctx.JSON(http.StatusOK, resp)
}

// (POST /login)
func (s *Server) PostLogin(ctx echo.Context) error {
	req := new(generated.LoginRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "failed to bind")
	}

	userData, err := s.Repository.GetUserDataByPhoneNumber(ctx.Request().Context(), req.PhoneNumber)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "error get user data")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(req.Password))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "unauthorized")
	}

	token, err := util.GenerateJWT(userData)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "error generate token")
	}

	err = s.Repository.UpdateLoginActivity(ctx.Request().Context(), userData.Id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "error log activity")
	}

	var resp generated.LoginResponse
	resp.Id = &userData.Id
	resp.JwtToken = &token
	return ctx.JSON(http.StatusOK, resp)
}

// (GET /user/{id})
func (s *Server) GetUserId(ctx echo.Context, id string, header generated.GetUserIdParams) error {
	userId := util.ValidateUserByToken(header.Authorization)
	if userId == 0 {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	userData, err := s.Repository.GetUserDataByID(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "error get user data")
	}

	var resp generated.UserInfoResponse
	resp.PhoneNumber = &userData.PhoneNumber
	resp.FullName = &userData.FullName
	return ctx.JSON(http.StatusOK, resp)
}

// (PATCH /user/{id})
func (s *Server) PatchUserId(ctx echo.Context, id string, header generated.PatchUserIdParams) error {
	userId := util.ValidateUserByToken(header.Authorization)
	if userId == 0 {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	req := new(generated.UserInfoRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	userData, err := s.Repository.GetUserDataByID(ctx.Request().Context(), userId)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "error get user data")
	}

	if req.FullName != "" {
		userData.FullName = req.FullName
	}

	if req.PhoneNumber != "" {
		_, err := s.Repository.GetUserDataByPhoneNumber(ctx.Request().Context(), req.PhoneNumber)
		if err == nil {
			return ctx.String(http.StatusBadRequest, "phone number exist")
		}

		userData.PhoneNumber = req.PhoneNumber
	}

	err = s.Repository.UpdateUserDataByID(ctx.Request().Context(), userData)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "error update user data")
	}

	var resp generated.UserInfoResponse
	resp.PhoneNumber = &userData.PhoneNumber
	resp.FullName = &userData.FullName
	return ctx.JSON(http.StatusOK, resp)
}
