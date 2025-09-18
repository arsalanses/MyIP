package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func clientIP(c *fiber.Ctx) string {
	hdrs := []string{
		"CF-Connecting-IP",
		"True-Client-IP",
		"X-Forwarded-For",
		"X-Real-IP",
		"X-Client-IP",
		"Forwarded",
	}
	for _, h := range hdrs {
		if v := c.Get(h); v != "" {
			if h == "X-Forwarded-For" {
				return strings.TrimSpace(strings.Split(v, ",")[0])
			}
			if h == "Forwarded" {
				for _, part := range strings.Split(v, ";") {
					part = strings.TrimSpace(part)
					if strings.HasPrefix(strings.ToLower(part), "for=") {
						val := strings.TrimPrefix(part, "for=")
						val = strings.Trim(val, `[]"`)
						if idx := strings.Index(val, ":"); idx != -1 && strings.Count(val, ":") == 1 {
							val = val[:idx]
						}
						return val
					}
				}
			}
			return strings.TrimSpace(v)
		}
	}
	host, _, err := net.SplitHostPort(c.Context().RemoteAddr().String())
	if err == nil {
		return host
	}
	return c.Context().RemoteAddr().String()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	app := fiber.New(fiber.Config{
		Prefork:      true,
		Network:      fiber.NetworkTCP6,
		AppName:      "MyIP",
		ServerHeader: "Go-Fiber",
	})

	if os.Getenv("LOG") == "true" {
		app.Use(logger.New())
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(clientIP(c))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	listenAddr := fmt.Sprintf("[::]:%s", port)

	useTLS := os.Getenv("USE_TLS")
	if useTLS == "true" {
		certFile := os.Getenv("CERT_FILE")
		keyFile := os.Getenv("KEY_FILE")
		if certFile == "" || keyFile == "" {
			log.Fatal("USE_TLS is true, but CERT_FILE or KEY_FILE is not set.")
		}
		log.Fatal(app.ListenTLS(listenAddr, certFile, keyFile))
	} else {
		log.Fatal(app.Listen(listenAddr))
	}
}
