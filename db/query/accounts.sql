-- name: CreateAccount :one
INSERT INTO accounts (
  owner,balance,currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListAccount :many
SELECT * FROM accounts
WHERE owner = $1
OFFSET $2 
LIMIT  $3;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;


-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateAccountBalace :one 
UPDATE "accounts"
set balance = $1
where id = $2
RETURNING *;

-- name: GetAccountByOwner :one
SELECT * FROM accounts
WHERE id = $1
AND owner = $2
LIMIT 1;