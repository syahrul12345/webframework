package api

import (
	"log"
	"net/http"
	"scratchuniversity/apps/db"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func createAccountHandler(c *gin.Context) {
	/***
		Incoming payload looks like
		{
			email: syahrul@example.com
			password: somepassword
		}
	***/
	account := &db.Account{}
	err := c.ShouldBindBodyWith(account, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Crate the account
	err = account.Create()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func loginAccountHandler(c *gin.Context) {
	/***
		Incoming payload looks like
		{
			email: syahrul@example.com
			password: somepassword
		}
	***/
	account := &db.Account{}
	err := c.ShouldBindBodyWith(account, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Login the account
	err = account.Login()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Return the account details with token
	c.JSON(http.StatusOK, gin.H{
		"account": account,
	})
}

func changePasswordHandler(c *gin.Context) {
	/***
		Incoming payload looks like
		{
			email: syahrul@example.com
			password: somepassword
			newPassword: newPassword
		}
	***/
	newAccount := &db.NewAccount{}
	err := c.ShouldBindBodyWith(newAccount, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	account := &db.Account{}
	account.Email = newAccount.Email
	account.Password = newAccount.Password

	if newAccount.NewPassword == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "New Password cannot be empty",
		})
		return
	}
	// Change the account password
	err = account.ChangePassword(newAccount.NewPassword)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully changed password, you can login now",
	})
	// Change the password

}

func forgetPasswordHandler(c *gin.Context) {
	/***
		Incoming payload looks like
		{
			email: syahrul@example.com
		}
	***/
	account := &db.Account{}
	c.ShouldBindWith(account, binding.JSON)

	err := account.Exists()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "An email has been sent to your account to reset your password",
	})
}

func getAccountDetailHandler(c *gin.Context) {
	cookieXToken, err := c.Cookie("x-token")
	if err != nil {
		// No cookie found, unauthorized route.
		log.Println("No cookie found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
	}
	acc, err := getAccountFromCookie(cookieXToken)
	if err != nil {
		// No cookie found, unauthorized route.
		log.Println("invalid cookie")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": acc,
	})

}
