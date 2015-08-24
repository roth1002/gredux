package main

type config struct {
	name   string
	author string
}

func (this *config) SetName(Name string) {
	this.name = Name
}

func (this *config) GetName() string {
	return this.name
}

func (this *config) SetAuthorName(authorName string) {
	this.author = authorName
}

func (this *config) GetAuthorName() string {
	return this.author
}
