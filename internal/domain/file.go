package domain

type FileLogicFace interface {
	UploadFile(fileName string, fileSize int64, fileUrl string, directory_id uint64, uploader uint64) error
	DeleteFile(fileId uint64) error
}

type FileRepositoryFace interface {
	UploadFile(fileName string, fileSize int64, fileUrl string, directory_id uint64, uploader uint64) error
	DeleteFile(fileId uint64) error
}

// File Domain And Interface
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
