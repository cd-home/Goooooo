package types

type DeleteFileParam struct {
	FileId uint64 `form:"file_id" binding:"required"`
}
