package handlers

import (
	"API_GO/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByID(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}
	users := []schemas.User{}
	if err := db.First(&users, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "User not found")
		return
	}
	sendSucces(ctx, "Listing-users", users)

}

func GetUsers(ctx *gin.Context) {
	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error listing users")
		return
	}
	sendSucces(ctx, "Listing-users", users)
}

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := schemas.User{
		Name:          request.Name,
		LastName:      request.LastName,
		CI:            request.CI,
		F_inscripcion: request.F_inscripicon,
		Phone:         request.Phone,
		Email:         request.Email,
		Titular:       request.Titular,
	}
	if err := db.Create(&user).Error; err != nil {
		logger.ErrorF("error creating USER: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user on databse")
		return
	}

	sendSucces(ctx, "create-user", user)
}

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}
	user := schemas.User{}
	//Encuentra usuario
	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("User with id:%s not found", id))
		return
	}
	//elimina usuario
	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting user with id %s", id))
		return
	}
	sendSucces(ctx, "deleting-user", user)
}

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.ErrorF("validation erro: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}
	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	//actualizar usuario
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.LastName != "" {
		user.LastName = request.LastName
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.CI != "" {
		user.CI = request.CI
	}
	if request.F_inscripicon != "" {
		user.F_inscripcion = request.F_inscripicon
	}
	if request.Titular != nil {
		user.Titular = request.Titular
	}
	if err := db.Save(&user).Error; err != nil {
		logger.ErrorF("error updating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}
	sendSucces(ctx, "Updatin-user", user)
}

func GetAllOpenningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"menssage": "GET ALL mundo",
	})
}
