package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nglLike/models"
	"golang.org/x/crypto/bcrypt"
)

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
	}

	var user models.User
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	err := ac.DB.Get(&user, query, userData.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "akun tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau password anda salah"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "data": user})
}
