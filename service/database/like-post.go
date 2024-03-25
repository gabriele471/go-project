package database

import (
	"myproject/service/types"
)

func (db *appdbimpl) InsertLike(likeid string, postId string, user types.User) error {

	query := `
        INSERT INTO likes (like_id, user_id, post_id)
        SELECT ?, ?, ? 
        WHERE EXISTS (SELECT post_id FROM posts WHERE post_id = ?)
    `

	_, err := db.c.Exec(query, likeid, user.Id, postId, postId)
	return err
}
