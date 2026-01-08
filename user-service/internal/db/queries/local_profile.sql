-- name: GetLocalProfileByID :one
SELECT *
FROM local_profile
WHERE id = $1;

-- name: ListLocalProfiles :many
SELECT *
FROM local_profile
ORDER BY id;

-- name: CreateLocalProfile :one
INSERT INTO local_profile (
    id,
    name,
    gender,
    email,
    status,
    role_id
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;