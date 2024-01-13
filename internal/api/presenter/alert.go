package presenter

import (
	"github.com/aldisypu/go-huberi/internal/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Alert struct {
	ID              uint   `json:"id"`
	BackgroundColor string `json:"background_color"`
	HighlightColor  string `json:"highlight_color"`
	TextColor       string `json:"text_color"`
	Duration        int16  `json:"duration"`
	IsUnit          bool   `json:"is_unit"`
	IsBorder        bool   `json:"is_border"`
}

func AlertSuccessResponse(data *entities.Alert) *fiber.Map {
	alert := Alert{
		ID:              data.ID,
		BackgroundColor: data.BackgroundColor,
		HighlightColor:  data.HighlightColor,
		TextColor:       data.TextColor,
		Duration:        data.Duration,
		IsUnit:          data.IsUnit,
		IsBorder:        data.IsBorder,
	}
	return &fiber.Map{
		"status": true,
		"data":   alert,
		"error":  nil,
	}
}

func AlertsSuccessResponse(data *[]Alert) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func AlertErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
