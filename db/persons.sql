-- name: CreatePerson :one
INSERT INTO persons (
    name
) VALUES ($1)
RETURNING *;

-- name: GetPerson :one
SELECT * FROM persons
WHERE name = $1 LIMIT 1;

-- name: UpdatePerson :one
UPDATE persons SET name = $1
WHERE name = $2
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM persons WHERE name = $1;