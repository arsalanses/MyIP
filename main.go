package main

import (
	"log"
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app := fiber.New(fiber.Config{
		Prefork:      true,
		Network:      fiber.NetworkTCP6,
		AppName:      "MyIP",
		ServerHeader: "Go-Fiber",
	})
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(clientIP(c))
	})
	log.Fatal(app.Listen("[::]:3000"))
	//	log.Fatal(app.ListenTLS("[::]:443", "/fullchain.pem", "/privkey.pem"))
}
