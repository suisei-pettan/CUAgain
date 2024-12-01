package dao

import (
	"CUAgain/models"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
	"time"
)

func GetIp(c *gin.Context, getIpMethod string) string {
	if getIpMethod == "0" {
		ip := c.Request.RemoteAddr
		// 去除端口
		if index := strings.LastIndex(ip, ":"); index != -1 {
			ip = ip[:index]
		}
		return ip
	} else {
		return c.Request.Header.Get(getIpMethod)
	}
}

func DecodePassword(password string, privateKeyPath string) (string, error) {
	// Decode the base64 encoded password
	cipherText, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return "", err
	}

	privateKey, err := os.ReadFile(privateKeyPath)
	// Decode the PEM formatted private key
	block, _ := pem.Decode(privateKey)
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("failed to decode PEM block containing the key")
	}

	// Parse the private key
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Decrypt the password using the private key
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, parsedKey.(*rsa.PrivateKey), cipherText)
	if err != nil {
		return "", err
	}

	// Convert decrypted bytes to string
	decryptedPassword := string(decryptedBytes)
	return decryptedPassword, nil
}

func CheckTheSignInStatus(ip string) bool {
	if Db.QueryRow("SELECT ip FROM users WHERE ip = $1", ip).Scan(&ip) != nil {
		return false
	} else {
		var dbTime string
		err := Db.QueryRow("SELECT expiration_time FROM users WHERE ip = $1", ip).Scan(&dbTime)
		formatTime, err := time.Parse(time.DateTime, dbTime)
		if err != nil {
			log.Fatal(err)
			return false
		}

		currentTime := time.Now().UTC()
		if currentTime.Before(formatTime) {
			return true
		} else {
			DeleteUserIp(ip)
			return false
		}
	}
}

func AllowUser(ip string) {
	var expirationTime time.Time
	expirationTime = time.Now().UTC().Add(time.Minute * time.Duration(models.GetConfig().CUAgain.LoginTimeout))
	dbTime := expirationTime.Format(time.DateTime)
	if CheckTheSignInStatus(ip) {
		_, err := Db.Exec("UPDATE users SET expiration_time = $1 WHERE ip = $2", dbTime, ip)
		if err != nil {
			return
		}
	} else {
		_, err := Db.Exec("INSERT INTO users (ip,expiration_time) VALUES ($1,$2)", ip, dbTime)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func DeleteUserIp(ip string) {
	_, err := Db.Exec("DELETE FROM users WHERE ip = $1", ip)
	if err != nil {
		log.Fatal("Database error,please check the database config")
		return
	}
}
