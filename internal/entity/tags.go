package entity

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateTagDTO struct {
	Name string `json:"name"`
}

func NewTag(id, name string) *Tag {
	return &Tag{
		ID:   id,
		Name: name,
	}
}
