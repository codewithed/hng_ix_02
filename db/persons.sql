-- name: CreatePerson :one
INSERT INTO persons (
    name, bio
) VALUES ($1, $2)
RETURNING *;

-- name: GetPerson :one
SELECT * FROM persons
WHERE name = $1 LIMIT 1;

-- name: UpdatePerson :one
UPDATE persons SET bio = $1
WHERE name = $2
RETURNING *;

-- name: DeletePerson :exec
DELETE FROM persons WHERE name = $1;