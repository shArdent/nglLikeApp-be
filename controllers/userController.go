package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nglLike/models"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	rows, err := uc.DB.Query("SELECT id, username, name, email, password, bio, created_at, updated_at FROM users;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Password, &user.Bio, &user.CreatedAt, &user.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (uc *UserController) AddNewUser(c *gin.Context) {
	var newDataPayload models.NewUserPayload

	if err := c.BindJSON(&newDataPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	pass := []byte(newDataPayload.Password)

	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	_, err = uc.DB.Exec("INSERT INTO users (username, email, password) VALUES (?,?,?)", newDataPayload.Username, newDataPayload.Email, hashed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Akun Berhasil Dibuat",
	})
}
