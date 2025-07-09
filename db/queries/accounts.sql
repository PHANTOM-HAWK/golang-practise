-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: CreateAccount :exec
INSERT INTO accounts (owner,balance) VALUES ($1,$2);

-- name: UpdateAccount :exec
UPDATE accounts SET owner = $2,balance = $3,currency = $4
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;