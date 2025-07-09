-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: CreateTransfer :exec
INSERT INTO transfers (from_account_id,to_account_id,amount) VALUES ($1,$2,$3);

-- name: UpdateTransfer :exec
UPDATE transfers SET from_account_id = $2,to_account_id = $3,amount = $4
WHERE id = $1;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;