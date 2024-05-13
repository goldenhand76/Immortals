-- name: CreateNode :one
INSERT INTO nodes (
    name, agent_id, client_id, is_online)
    VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetNode :one
SELECT * FROM nodes
WHERE id = $1 LIMIT 1;

-- name: ListNodes :many
SELECT * FROM nodes
ORDER BY name;

-- name: UpdateNode :one
UPDATE nodes SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteNode :exec
DELETE FROM nodes WHERE id = $1;