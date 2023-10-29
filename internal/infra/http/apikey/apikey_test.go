package apikey

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

// go test -run Test_ApiKey_Next
func Test_ApiKey_Next(t *testing.T) {
	t.Parallel()

	app := fiber.New()
	app.Use(New(Config{
		Next: func(_ *fiber.Ctx) bool {
			return true
		},
	}))

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode)
}

func Test_Middleware_ApiKey(t *testing.T) {
	t.Parallel()
	app := fiber.New()

	app.Use(New(Config{
		Authorizer: func(token string) bool {
			return token == "john"
		},
	}))

	app.Get("/testauth", func(c *fiber.Ctx) error {
		token := c.Locals("token").(string)

		return c.SendString(token)
	})

	tests := []struct {
		url        string
		statusCode int
		token      string
	}{
		{
			url:        "/testauth",
			statusCode: 200,
			token:      "Bearer john",
		},
		{
			url:        "/testauth",
			statusCode: 401,
			token:      "Bearer ee",
		},
		{
			url:        "/testauth",
			statusCode: 401,
			token:      "Bearer ",
		},
		{
			url:        "/testauth",
			statusCode: 401,
			token:      "Bearer",
		},
		{
			url:        "/testauth",
			statusCode: 401,
			token:      "",
		},
	}

	for _, tt := range tests {
		// Base64 encode credentials for http auth header

		req := httptest.NewRequest("GET", "/testauth", nil)
		req.Header.Add("Authorization", tt.token)
		resp, err := app.Test(req)
		utils.AssertEqual(t, nil, err)

		body, err := io.ReadAll(resp.Body)

		utils.AssertEqual(t, nil, err)
		utils.AssertEqual(t, tt.statusCode, resp.StatusCode)

		if tt.statusCode == 200 {
			utils.AssertEqual(t, tt.token, "Bearer "+string(body))
		}
	}
}

// go test -v -run=^$ -bench=Benchmark_Middleware_ApiKey -benchmem -count=4
func Benchmark_Middleware_ApiKey(b *testing.B) {
	app := fiber.New()

	app.Use(New(Config{
		Authorizer: func(token string) bool {
			return token == "john"
		},
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusTeapot)
	})

	h := app.Handler()

	fctx := &fasthttp.RequestCtx{}
	fctx.Request.Header.SetMethod("GET")
	fctx.Request.SetRequestURI("/")
	fctx.Request.Header.Set(fiber.HeaderAuthorization, "bearer john")

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		h(fctx)
	}

	utils.AssertEqual(b, fiber.StatusTeapot, fctx.Response.Header.StatusCode())
}
