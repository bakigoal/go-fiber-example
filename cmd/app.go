package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

type App struct {
	Server *fiber.App
}

func NewApp() *App {
	return &App{
		Server: createServer(),
	}
}

func createServer() *fiber.App {
	app := fiber.New()

	setupMiddlewares(app)
	setupRoutes(app)
	return app
}

func setupMiddlewares(app *fiber.App) {

	// Middleware that matches any route
	app.Use(func(c fiber.Ctx) error {
		fmt.Println("🥇 First handler", c.Request().URI())
		return c.Next()
	})

	// Middleware that matches all routes starting with /api
	app.Use("/static", func(c fiber.Ctx) error {
		fmt.Println("🥈 Second handler", c.Request().URI())
		return c.Next()
	})
}

func setupRoutes(app *fiber.App) {
	// Static files
	app.Get("/static/*", static.New("./public"))

	// GET /api/register
	app.Get("/api/*", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /john
	app.Get("/:name", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	// Serve a single file for any unmatched routes
	app.Get("*", static.New("./public/index.html"))
}
