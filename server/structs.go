package main

import "github.com/lib/pq"

type User struct {
	UserID   int           `json:"UserID" db:"user_id"`
	Segments pq.Int32Array `json:"Segments" db:"segments"`
}

type Segmentlist struct {
	SegmentID int    `json:"SegmentID" db:"segment_id"`
	Slug      string `json:"Slug" db:"slug"`
}

type Updater struct {
	SlugsToAdd    pq.Int32Array `json:"SlugsToAdd" db:"slugsToAdd"`
	SlugsToDelete pq.Int32Array `json:"SlugsToDelete" db:"slugsToDelete"`
	UserID        int           `json:"UserID" db:"user_id"`
}
