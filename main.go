package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/codewithed/hng_ix_02/db/sqlc"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	conn := fmt.Sprintf("%s,%s", os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	sqlDB, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(sqlDB)

	router := fiber.New()

	router.Post("/", func(c *fiber.Ctx) error {
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

	router.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
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

	router.Put("/:name", func(c *fiber.Ctx) error {
		params := db.UpdatePersonParams{}
		err := c.BodyParser(&params)
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

	router.Patch("/:name", func(c *fiber.Ctx) error {
		params := db.UpdatePersonParams{}
		err := c.BodyParser(&params)
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

	router.Delete("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		err := queries.DeletePerson(context.Background(), name)
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

		msg := fmt.Sprintf("deleted person: %v", name)
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": msg,
		})
	})

	router.Listen(":" + os.Getenv("SERVER_ADDRESS"))
}
