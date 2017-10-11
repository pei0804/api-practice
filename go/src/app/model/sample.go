package model

type Sample struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

func (s *Sample) Get() Sample {
	return Sample{
		Title: "title",
		Name:  "name",
	}
}
