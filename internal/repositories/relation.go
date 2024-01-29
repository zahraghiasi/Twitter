package repositories

import "github.com/jinzhu/gorm"

type Relation struct {
	gorm.Model
	UserID        int
	SubjectUserID int
	Type          RelationType
}

type RelationType int

const (
	Follow RelationType = iota
	Block
)
