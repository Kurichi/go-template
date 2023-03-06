package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

// f is the path to the service account file
func FirebaseAuthMiddleware(f string) echo.MiddlewareFunc {
	opt := option.WithCredentialsFile(f)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "authorization header is required"})
			}
			if !strings.Contains(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "authorization header must be Bearer"})
			}
			idToken := strings.Split(authHeader, "Bearer ")[1]

			auth, err := client.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization header"})
			}
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "uid", auth.UID)))
			// auth.UID

			return next(c)
		}
	}
}
