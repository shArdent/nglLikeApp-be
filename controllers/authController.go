package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"github.com/nglLike/models"
	"golang.org/x/crypto/bcrypt"
)

const secretKey string = "apacb"

type AuthController struct {
	DB *sqlx.DB
}

func NewAuthController(db *sqlx.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) RegisterHandler(c *gin.Context) {
	var newDataPayload models.NewUserPayload
	if err := c.BindJSON(&newDataPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newDataPayload.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	_, err = ac.DB.Exec(query, newDataPayload.Email, hashed)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email sudah terpakai"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Akun Berhasil Dibuat",
	})
}

func (ac *AuthController) LoginHandler(c *gin.Context) {
	var userData models.LoginPayload
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	err := ac.DB.Get(&user, query, userData.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "akun tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email atau password anda salah"})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expiredAt := time.Now().Add(time.Hour * 24).Unix()

	c.SetCookie("jwt", token, int(expiredAt), "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": true, "data": []string{}, "message": "Berhasil login"})
}
