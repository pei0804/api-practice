package model

type Foo struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

func (f *Foo) Get() Sample {
	return Sample{
		Title: "title",
		Name:  "name",
	}
}
