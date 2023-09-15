package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	db "github.com/codewithed/hng_ix_02/db/sqlc"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type UpdatePerson struct {
	Age int32 `json:"age,omitempty"`
}

func main() {

	conn := os.Getenv("DB_SOURCE")
	dbDriver := os.Getenv("DB_DRIVER")
	sqlDB, err := sql.Open(dbDriver, conn)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(sqlDB)

	router := fiber.New()

	router.Post("/api", func(c *fiber.Ctx) error {
		params := db.CreatePersonParams{}
		err := c.BodyParser(&params)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		user, err := queries.CreatePerson(context.Background(), params)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(user)
	})

	router.Get("/api/:name", func(c *fiber.Ctx) error {
		name_param := c.Params("name")
		name := strings.ReplaceAll(name_param, "%20", " ")
		user, err := queries.GetPerson(context.Background(), name)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(user)
	})

	router.Put("/api/:name", func(c *fiber.Ctx) error {
		name_param := c.Params("name")
		req := UpdatePerson{}
		err := c.BodyParser(&req)
		name := strings.ReplaceAll(name_param, "%20", " ")
		params := db.UpdatePersonParams{
			Name: name,
			Age:  req.Age,
		}

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		user, err := queries.UpdatePerson(context.Background(), params)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(user)
	})

	router.Patch("/api/:name", func(c *fiber.Ctx) error {
		name_param := c.Params("name")
		req := UpdatePerson{}
		err := c.BodyParser(&req)
		name := strings.ReplaceAll(name_param, "%20", " ")
		params := db.UpdatePersonParams{
			Name: name,
			Age:  req.Age,
		}

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		user, err := queries.UpdatePerson(context.Background(), params)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(http.StatusOK).JSON(user)
	})

	router.Delete("/api/:name", func(c *fiber.Ctx) error {
		name_param := c.Params("name")
		name := strings.ReplaceAll(name_param, "%20", " ")
		err := queries.DeletePerson(context.Background(), name)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		msg := fmt.Sprintf("deleted person: %v", name)
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": msg,
		})
	})

	router.Listen(":" + os.Getenv("SERVER_ADDRESS"))
}
