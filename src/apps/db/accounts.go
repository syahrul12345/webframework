package db

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token is a JWT token that will returned to the frontend
type Token struct {
	UserID   uint
	UserName string
	jwt.StandardClaims
}

// Account represents a user to be saved in the DB
type Account struct {
	gorm.Model
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string
}

// NewAccount struct
type NewAccount struct {
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	NewPassword string `json:"NewPassword"`
}

// Validate account that have yet to be created. This is called when a new account object has to be created
func (acc *Account) Validate() error {
	if !strings.Contains(acc.Email, "@") {
		return errors.New("Email Address required")
	}
	if len(acc.Password) < 6 {
		return errors.New("Password is has to be more than 6 characters")
	}

	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("Connection Error. Please retry")
	}
	//Email must be unique
	if temp.Email != "" {
		return errors.New("Email already in use")
	}
	return nil
}

//ValidateLogin is used to validate accounts that are attempting to login
func (acc *Account) ValidateLogin() error {
	// Get the account object from the DB
	// Allocate a temp account object
	// dbAccount := &Account{}
	providedPassword := acc.Password
	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(acc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return errors.New("Account does not exist")
	}
	// let's validate the old password
	// temp is the record that exists in the database. The password is hashed using bcrypt earlier during creation.
	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(providedPassword))
	if err != nil {
		return errors.New("Wrong password")
	}
	return nil
}

//Create account
func (acc *Account) Create() error {
	err := acc.Validate()
	if err != nil {
		return err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost)
	acc.Password = string(hashedPassword)
	//stores the account into the database
	GetDB().Create(acc)

	if acc.ID <= 0 {
		return errors.New("Failed to create new account dew to database error of ID less than 0")
	}
	//create a new JWT token
	tk := &Token{acc.ID, acc.Email, jwt.StandardClaims{
		ExpiresAt: getExpiryDate(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "scatchuniversity",
		Subject:   "Authtoken",
	}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return err
	}
	acc.Token = tokenString
	acc.Password = ""

	return nil
}

// ChangePassword will change the password of the account
func (acc *Account) ChangePassword(newPassword string) error {
	err := acc.ValidateLogin()
	if err != nil {
		return err
	}
	// Get new hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	// Update the password field only
	GetDB().Model(acc).Update("Password", string(hashedPassword))
	//create JWT TOKEN
	tk := &Token{acc.ID, acc.Email, jwt.StandardClaims{
		ExpiresAt: getExpiryDate(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "scatchuniversity",
		Subject:   "Authtoken",
	}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	claims := token.Claims.(*Token)
	claims.StandardClaims.ExpiresAt = getExpiryDate()

	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return err
	}
	acc.Token = tokenString
	fmt.Println("User: " + acc.Email + " has succesfully changed password")
	return nil
}

// Login : Will attemp a login
func (acc *Account) Login() error {
	err := acc.ValidateLogin()
	if err != nil {
		return err
	}
	tk := &Token{acc.ID, acc.Email, jwt.StandardClaims{
		ExpiresAt: getExpiryDate(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "scatchuniversity",
		Subject:   "Authtoken",
	}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// Our standard claims is stored in the db.Token struct
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return err
	}
	acc.Password = ""
	acc.Token = tokenString
	fmt.Println("User: " + acc.Email + " has succesfully logged in.")
	return nil
}

// Exists checks if an account exists
func (acc *Account) Exists() error {
	err := GetDB().Table("accounts").Where("email = ?", acc.Email).First(acc).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errors.New("Connection Error. Please retry")
	}
	if err == gorm.ErrRecordNotFound {
		return errors.New("Account does not exist")
	}
	return nil
}

func getExpiryDate() int64 {
	start := time.Now()
	end := start.Add(time.Hour * 1)
	return end.Unix()
}
