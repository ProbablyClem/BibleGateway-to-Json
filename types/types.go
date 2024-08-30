package types

type Bible struct {
	Books map[string]Book
}

type Book struct {
	Chapters map[int]Chapter
}

type Chapter struct {
	Verses map[int]Verse
}

type Verse struct {
	Text string
	Uri  string
	Line int
}
