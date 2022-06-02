package domain

type File struct {
	Id        uint64
	UUID      string
	Name      string
	Size      uint64
	Type      string
	Uploader  uint64
	Directory uint64
	Url       string
	UpdateAt  string
	CreateAt  string
	DeleteAt  string
}
