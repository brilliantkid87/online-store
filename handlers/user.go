package handlers

import (
	"log"
	"net/http"
	"time"

	"synapsis/models"
	"synapsis/pkg/bcrypt"
	jwtToken "synapsis/pkg/jwt"
	"synapsis/pkg/logger"
	"synapsis/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

type UserHandler struct {
	Repo      *repositories.UserRepository
	Validator *validator.Validate
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		Repo:      repo,
		Validator: validator.New(),
	}
}

var Logger = logger.New()

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		Logger.Errorln("Invalid input")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := h.Validator.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			errorMessages[fieldName] = fieldName + " is required"
		}
		return c.JSON(http.StatusBadRequest, errorMessages)
	}

	password, err := bcrypt.HashingPassword(user.Password)
	if err != nil {
		Logger.Errorln("hashing pass failed")
		return c.JSON(http.StatusBadRequest, "hashing pass failed")
	}

	userParams := map[string]interface{}{
		"email": user.Email,
	}
	users, err := h.Repo.GetAllUsers(c.Request().Context(), userParams)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to check user")
	}

	if len(users) > 0 {
		Logger.Errorln("Email is already registered")
		return c.String(http.StatusInternalServerError, "Email is already registered")
	}

	userID, err := h.Repo.RegisterUser(c.Request().Context(), map[string]interface{}{
		"email":     user.Email,
		"user_name": user.UserName,
		"password":  password,
	})
	if err != nil {
		Logger.Errorln("Register user failed")
		return c.String(http.StatusInternalServerError, "Register user failed")
	}

	return c.JSON(http.StatusOK, map[string]string{"user_id": userID})
}

func (h *UserHandler) Login(c echo.Context) error {
	var params map[string]interface{}
	if err := c.Bind(&params); err != nil {
		Logger.Errorln("Invalid input")
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	// Check if user exists
	userParams := map[string]interface{}{
		"email": params["email"],
	}
	users, err := h.Repo.GetAllUsers(c.Request().Context(), userParams)
	if err != nil {
		Logger.Errorln("failed to check user")
		return c.JSON(http.StatusInternalServerError, "Failed to check user")
	}

	if len(users) == 0 {
		Logger.Errorln("Email not registered. Please register first")
		return c.JSON(http.StatusUnauthorized, "Email not registered. Please register first.")
	}

	user := users[0]

	passwordHash, ok := user["password"].(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Password hash is invalid")
	}

	// Verify password
	err = bcrypt.CheckPasswordHash(params["password"].(string), passwordHash)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	userID, ok := user["user_id"].(string)
	if !ok {
		Logger.Errorln("userid invalid")
		return c.JSON(http.StatusInternalServerError, "User ID is invalid")
	}

	// Generate JWT token
	claims := jwt.MapClaims{}
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
