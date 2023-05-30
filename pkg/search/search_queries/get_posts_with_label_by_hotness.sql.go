// Code generated by sqlc. DO NOT EDIT.
// source: get_posts_with_label_by_hotness.sql

package search_queries

import (
	"context"
	"database/sql"
	"time"
)

const getPostsPageWithPostLabelSortedByHotness = `-- name: GetPostsPageWithPostLabelSortedByHotness :many
SELECT h.id, h.text, h.parent_post_id, h.root_post_id, h.author_did, h.created_at, 
       h.has_embedded_media, h.parent_relationship, h.sentiment, h.sentiment_confidence, h.hotness::float as hotness
FROM post_hotness h
WHERE h.label = $1 AND 
      (CASE WHEN $2::float = -1 THEN TRUE ELSE hotness < $2::float END)
ORDER BY h.hotness DESC, h.id DESC
LIMIT $3
`

type GetPostsPageWithPostLabelSortedByHotnessParams struct {
	Label  string  `json:"label"`
	Cursor float64 `json:"cursor"`
	Limit  int32   `json:"limit"`
}

type GetPostsPageWithPostLabelSortedByHotnessRow struct {
	ID                  string          `json:"id"`
	Text                string          `json:"text"`
	ParentPostID        sql.NullString  `json:"parent_post_id"`
	RootPostID          sql.NullString  `json:"root_post_id"`
	AuthorDid           string          `json:"author_did"`
	CreatedAt           time.Time       `json:"created_at"`
	HasEmbeddedMedia    bool            `json:"has_embedded_media"`
	ParentRelationship  sql.NullString  `json:"parent_relationship"`
	Sentiment           sql.NullString  `json:"sentiment"`
	SentimentConfidence sql.NullFloat64 `json:"sentiment_confidence"`
	Hotness             float64         `json:"hotness"`
}

func (q *Queries) GetPostsPageWithPostLabelSortedByHotness(ctx context.Context, arg GetPostsPageWithPostLabelSortedByHotnessParams) ([]GetPostsPageWithPostLabelSortedByHotnessRow, error) {
	rows, err := q.query(ctx, q.getPostsPageWithPostLabelSortedByHotnessStmt, getPostsPageWithPostLabelSortedByHotness, arg.Label, arg.Cursor, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostsPageWithPostLabelSortedByHotnessRow
	for rows.Next() {
		var i GetPostsPageWithPostLabelSortedByHotnessRow
		if err := rows.Scan(
			&i.ID,
			&i.Text,
			&i.ParentPostID,
			&i.RootPostID,
			&i.AuthorDid,
			&i.CreatedAt,
			&i.HasEmbeddedMedia,
			&i.ParentRelationship,
			&i.Sentiment,
			&i.SentimentConfidence,
			&i.Hotness,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
