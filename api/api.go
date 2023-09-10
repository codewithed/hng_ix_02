package api

import (
	db "github.com/codewithed/hng_ix_02/db/sqlc"
	"github.com/gofiber/fiber"
)

type server struct {
	db     db.Querier
	router *fiber.Router
}
