package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"rakamin.com/final-task/database"
	"rakamin.com/final-task/models"
)

// post data
func Register(c *gin.Context) {
	// get req body data
	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		fmt.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//  post data
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash), CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := database.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": user,
	})
}

// login
func Login(c *gin.Context) {
	// get the email body
	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		fmt.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// find user
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	// compare hash password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "succesful login",
	})
}

// get user by token
func GetUserLogin(c *gin.Context) {
	user, _ := c.Get("user")

	// Cek tipe data pengguna
	if userObj, ok := user.(models.User); ok {
		// Memuat data foto-foto terkait dengan pengguna
		database.DB.Model(&userObj).Association("Photos").Find(&userObj.Photos)

		c.JSON(http.StatusOK, gin.H{
			"user": userObj,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

// update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// find user
	var user models.User
	result := database.DB.First(&user, id)

	// check user
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("User with ID %s not found", id),
		})
		return
	}

	database.DB.Model(&user).Updates(models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Account with ID %s has been updated", id),
	})

}

// delete user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	result := database.DB.First(&user, id)

	// check user
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("User with ID %s not found", id),
		})
		return
	}

	database.DB.Delete(&user, id)

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Account with ID %s has been deleted", id),
	})
}

// get data all data account
func PostsIndex(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)

	for i := range user {
		database.DB.Model(&user[i]).Association("Photos").Find(&user[i].Photos)
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetPhoto(c *gin.Context) {
	var photo []models.Photo
	database.DB.Find(&photo)

	c.JSON(200, gin.H{
		"data": photo,
	})
}
