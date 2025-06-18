-- name: CreateEntry :one
INSERT INTO entries (
 account_id,amount
) VALUES (
  $1, $2 
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 Limit 1;

-- name: UpdateEntry :exec
UPDATE entries
SET account_id = $2, amount = $3
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;