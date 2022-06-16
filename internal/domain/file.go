package domain

import "context"

type FileLogicFace interface {
	UploadFile(ctx context.Context, fileName string, fileSize int64, fileType string, fileUrl string, directory_id uint64, uploader uint64) error
	DeleteFile(ctx context.Context, fileId uint64) error
	RetrieveFiles(ctx context.Context) ([]*FileEntityVo, error)
}

type FileRepositoryFace interface {
	UploadFile(ctx context.Context, fileName string, fileSize int64, fileType string, fileUrl string, directory_id uint64, uploader uint64) error
	DeleteFile(ctx context.Context, fileId uint64) error
	RetrieveFiles(ctx context.Context) ([]*FileEntityDTO, error)
}

// File Domain And Interface
type File struct {
	Id        uint64 `db:"id"`
	FileId    string `db:"file_id"`
	Name      string `db:"name"`
	Size      uint64 `db:"size"`
	FileType  string `db:"file_type"`
	Uploader  uint64 `db:"uploader"`
	Directory uint64 `db:"directory_id"`
	Url       string `db:"url"`
	UpdateAt  string `db:"update_at"`
	CreateAt  string `db:"create_at"`
	DeleteAt  string `db:"delete_at"`
}

type FileEntityDTO struct {
	FileId    string `db:"file_id"`
	FileName  string `db:"file_name"`
	FileSize  uint64 `db:"file_size"`
	FileType  string `db:"file_type"`
	Uploader  uint64 `db:"uploader"`
	Directory uint64 `db:"directory_id"`
	FileUrl   string `db:"file_url"`
	UpdateAt  string `db:"update_at"`
	CreateAt  string `db:"create_at"`
}

type FileEntityVo struct {
	FileId    string `json:"file_id"`
	FileName  string `json:"file_name"`
	FileSize  uint64 `json:"file_size"`
	FileType  string `json:"file_type"`
	Uploader  uint64 `json:"uploader"`
	Directory uint64 `json:"directory_id"`
	FileUrl   string `json:"file_url"`
	UpdateAt  string `json:"update_at"`
	CreateAt  string `json:"create_at"`
}
