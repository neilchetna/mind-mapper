package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/repository"
	"github.com/neilchetna/mind-mapper/internal/service"
	"github.com/neilchetna/mind-mapper/pkg/utils"
	"gorm.io/gorm"
)

func SyncUser(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			clerkUserId, ok := c.Get(utils.ClerkUserId).(string)
			if !ok {
				return echo.ErrUnauthorized
			}
			userRepo := repository.UserRepositoryBuilder(db)
			userService := service.UserServiceBuilder(userRepo)

			user, err := userService.SyncClerkUserToDatabase(ctx, clerkUserId)

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			c.Set(utils.User, user)
			return next(c)
		}
	}
}
