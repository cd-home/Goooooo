package domain

import "context"

type FileLogicFace interface {
	UploadFile(ctx context.Context, fileName string, fileSize int64, fileUrl string, directory_id uint64, uploader uint64) error
	DeleteFile(ctx context.Context, fileId uint64) error
}

type FileRepositoryFace interface {
	UploadFile(ctx context.Context, fileName string, fileSize int64, fileUrl string, directory_id uint64, uploader uint64) error
	DeleteFile(ctx context.Context, fileId uint64) error
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
