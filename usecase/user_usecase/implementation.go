package user_usecase

import (
	"majoo/helpers"
	"majoo/middleware"
	"majoo/model/dto"
	"majoo/model/entity"
	"net/http"
)

func (user *userUsecase) Login(request dto.LoginRequest) dto.Response {
	loginUser := entity.User{
		Username: request.Username,
		Password: request.Password,
	}
	userData, err := user.userRepo.Login(loginUser)

	if err != nil {
		return helpers.SendResponse(http.StatusInternalServerError, "Internal server error", err.Error(), nil)   
   	}

	token, errToken := middleware.GenerateToken(userData.Id)
	if errToken != nil {
		return helpers.SendResponse(http.StatusInternalServerError, "Error on generating token", err.Error(), nil)  
	}
	res := make(map[string]interface{})
	res["name"] = userData.Name
	res["user_name"] = userData.Username
	res["token"] = token

   	return helpers.SendResponse(http.StatusOK, "Login success", nil, res)
}

func (user *userUsecase) Register(request dto.RegisterRequest) dto.Response {
	hashed_pwd, _ := helpers.HashPassword(request.Password)
	newUser := entity.User{
		Name: request.Name,
		Username: request.Username,
		Password: hashed_pwd,
	}
	userData, err := user.userRepo.Register(newUser)

	
	if err != nil {
		return helpers.SendResponse(http.StatusInternalServerError, "Internal server error", err.Error(), nil)   
   	}

	token, errToken := middleware.GenerateToken(userData.Id)
	if errToken != nil {
		return helpers.SendResponse(http.StatusInternalServerError, "Error on generating token", err.Error(), nil)  
	}
	res := make(map[string]interface{})
	res["name"] = userData.Name
	res["user_name"] = userData.Username
	res["token"] = token

   	return helpers.SendResponse(http.StatusOK, "Register success", nil, res)
}