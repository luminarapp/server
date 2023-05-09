package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/auth"
	"github.com/luminarapp/server/config"
	"github.com/luminarapp/server/models"
	"github.com/luminarapp/server/utils"
)

// GET /users/me
func CurrentUser(c *gin.Context) {
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/me
func UpdateCurrentUser(c *gin.Context) {
	var payload models.UpdateUserRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Make sure user exists
	var user models.User

	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user
	if err := models.DB.Model(&user).Updates(models.User{
		Username: payload.Username,
	}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /users/auth
// TODO: Refactor this function to be more readable and move evtl. to new file or auth package
// Also consider adding new route for the submission - /users/auth/submit or verify
func UserAuthChallenge(c *gin.Context) {
	var payload models.UserAuthChallengeRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if token provided in payload
	if payload.Token == "" || payload.Token == "null" {
		// Check if user exists with email
		var user models.User

		// Generate magic token
		magicToken := utils.GenerateMagicToken()

		if err := models.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
			// Create user
			user = models.User{
				ID: shortuuid.New(),
				Username: strings.ToLower(strings.Split(payload.Email, "@")[0]),
				Email: payload.Email,
				AuthToken: models.AuthToken{
					Token: magicToken,
					ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Config().ChallengeTokenLifetime)).Unix(),
				},
			}

			if err := models.DB.Create(&user).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
					Success: false,
					AuthStep: "tokenRequest",
					ErrorMsg: err.Error(),
				}})

				return
			}

			c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
				Success: true,
				AuthStep: "tokenRequest",
			}})

			return
		} else {
			// Update user token
			user.AuthToken = models.AuthToken{
				Token: magicToken,
				ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.Config().ChallengeTokenLifetime)).Unix(),
			}

			if err := models.DB.Save(&user).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
					Success: false,
					AuthStep: "tokenRequest",
					ErrorMsg: err.Error(),
				}})
				return
			}

			fmt.Println("Auth token (Request): ", user.AuthToken.Token)

			c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
				Success: true,
				AuthStep: "tokenRequest",
			}})

			return
		}
	}

	// Check if token provided in payload is valid
	var user models.User

	// Get user by email
	if err := models.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
			Success: false,
			AuthStep: "tokenSubmittion",
			ErrorMsg: "user not found",
		}})
	
		return
	}

	fmt.Println("Auth token (Submittion): ", payload.Token)
	fmt.Println("Current Auth token: ", user.AuthToken.Token)

	// Check if token matches
	if payload.Token != user.AuthToken.Token {
		c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
			Success: false,
			AuthStep: "tokenSubmittion",
			ErrorMsg: "invalid token",
		}})
		return
	}

	// Check if token has expired
	if user.AuthToken.ExpiresAt < time.Now().Unix() {
		c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
			Success: false,
			AuthStep: "tokenSubmittion",
			ErrorMsg: "token expired",
		}})
		return
	}

	// Delete token
	user.AuthToken = models.AuthToken{}

	if err := models.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
			Success: false,
			AuthStep: "tokenSubmittion",
			ErrorMsg: err.Error(),
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": models.UserAuthChallengeStatus{
		Success: true,
		AuthStep: "tokenSubmittion",
		Session: auth.GenerateToken(user.ID),
	}})
}