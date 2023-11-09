-- name: CreateEntrie :one
INSERT INTO entries (
  account_id, amnount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntrie :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;