package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kadsin/sms-gateway/database"
	"github.com/kadsin/sms-gateway/database/models"
	"github.com/kadsin/sms-gateway/internal/server/requests"
)

func ChangeUserBalance(c *fiber.Ctx) error {
	data, err := requests.Prepare[requests.UserBalanceRequest](c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	tx := database.Instance().Where("email", data.Email).
		Assign("balance", data.Balance).
		FirstOrCreate(&models.User{})

	return tx.Error
}
