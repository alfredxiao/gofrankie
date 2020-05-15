package models

// PuppyObject represents a success response
type PuppyObject struct {
	Puppy bool `json:"puppy" binding:"required"`
}
