package main

import (
	fiber "github.com/gofiber/fiber/v2"
)

type ipLogs struct {
	IpSoucrce  string
	IpForwards []string
}

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("homepage")
	})

	app.Get("/ipSource", func(c *fiber.Ctx) error {
		ip := ipLogs{
			IpSoucrce:  c.IP(),
			IpForwards: c.IPs(),
		}

		return c.JSON(ip)
	})

	app.Listen(":8083")
}
