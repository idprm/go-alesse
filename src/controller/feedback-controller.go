package controller

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type FeedbackRequest struct {
	ChatID     uint64  `query:"chat_id" validate:"required" json:"chat_id"`
	Slug       string  `query:"slug" validate:"required" json:"slug"`
	Rating     float32 `query:"rating" validate:"required" json:"rating"`
	Suggestion string  `query:"suggestion" json:"suggestion"`
}

func ValidateFeedback(req FeedbackRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func GetFeedback(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var feedback model.Feedback
	database.Datasource.DB().Where("slug", slug).First(&feedback)
	return c.Status(fiber.StatusOK).JSON(feedback)
}

func SaveFeedback(c *fiber.Ctx) error {
	c.Accepts("application/json")

	req := new(FeedbackRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := ValidateFeedback(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}

	var feedback model.Feedback
	existFeedback := database.Datasource.DB().Where("chat_id", req.ChatID).First(&feedback)
	if existFeedback.RowsAffected == 0 {
		database.Datasource.DB().Create(
			&model.Feedback{
				ChatID:     req.ChatID,
				Slug:       req.Slug,
				Rating:     req.Rating,
				Suggestion: req.Suggestion,
				SubmitedAt: time.Now(),
				IsSubmited: true,
			},
		)
	} else {
		feedback.Rating = req.Rating
		feedback.Suggestion = req.Suggestion
		feedback.SubmitedAt = time.Now()
		feedback.IsSubmited = true
		database.Datasource.DB().Save(&feedback)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "Submited",
		"data":    feedback,
	})
}
