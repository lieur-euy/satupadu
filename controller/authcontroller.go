package controller

import (
	"fmt"
	"regexp"
	"strings"
	"web/database"
	"web/models"

	"github.com/gofiber/fiber/v2"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.Users
	if err := c.BodyParser(&data); err != nil {
		fmt.Print("unable to parshe body")
	}

	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "pasword harus lebih 6 character",
		})
	}
	if !ValidateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid email addrees",
		})
	}
	database.DB.Where("email=?", strings.TrimSpace((data["email"]).(string))).First(*&userData)
	if userData.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "email sudah ada",
		})
	}
	user := models.Users{
		Username:     data["username"].(string),
		Email:        data["email"].(string),
		Hp:           data["hp"].(string),
		Image:        data["image"].(string),
		PasswordHash: data["password"].(string),
	}
}
