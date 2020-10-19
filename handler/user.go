package handler

import (
	"fmt"
	"gomar/auth"
	"gomar/helper"
	"gomar/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap inputan user
	// map input dari user ke struct RegisrterUserInput
	// struct diatas akan di parsing sbg param service
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err) //get metode helper error
		errorMessage := gin.H{"errors": errors}     // error tsb di map oleh gin.H
		response := helper.APIResponse("failed register", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("failed register", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("failed register", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user input email dan pass
	// input ditangkap handler
	// maping inputan ke input struck
	// input strcuk parsing ke service
	// service mencari dg bantuan repository user dengan email
	// if ketemu cocokan password

	var input user.LoginInput // tangkap inputan login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err) //get metode helper error
		errorMessage := gin.H{"errors": errors}     // error tsb di map oleh gin.H

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Loggin failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CekEmailDB(c *gin.Context) {
	// if email ada di db atau belum ?
	// input di maping ke struck
	// struck input di parsing ke service
	// service akan memanggil repo - amil ada or belum ?
	// repository melakukan query di db

	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err) //get metode helper error
		errorMessage := gin.H{"errors": errors}     // error tsb di map oleh gin.H
		response := helper.APIResponse("Email Check failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email Check failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	pesan := "Email has been registered"

	if IsEmailAvailable {
		pesan = "Email is available"
	}

	response := helper.APIResponse(pesan, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	return

}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	// tangkap inputan user
	// save gambar ke folder "image"
	// di service panggil repository
	// JWT (HARCODE) ,user login ID == 1
	// Repository get data where user ID == 1
	// repo update data user, simpan lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image ", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 1
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image ", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image ", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar success upload ", http.StatusOK, "susccess", data)

	c.JSON(http.StatusOK, response)
	return
}
