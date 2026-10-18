-- name: GetFeedFollowsForUser :many
SELECT
    *,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
JOIN users ON feed_follows.user_id = users.id
JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;
