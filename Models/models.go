package model

type ShortUrl struct {
	Url      string `json:"url,omitempty" bson:"url,omitempty"`
	PID      string `json:"pid,omitempty" bson:"pid,omitempty"`
	ShortUrl string `json:"-"`
}
