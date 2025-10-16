package config

import (
	// "os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"http://localhost:3000",
}

var headers = []string{
	"Origin",
	"Content-Type",
	"Accept",
	"Authorization",
	"Access-Control-Request-Headers",
	"Token",
	"Login",
	"Access-Control-Allow-Origin",
	"Bearer",
	"X-Requested-With",
}

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins, ","),
	AllowHeaders:     strings.Join(headers, ","),
	AllowCredentials: true,
	ExposeHeaders:    "Content-Length",
	MaxAge:           3600,
}
