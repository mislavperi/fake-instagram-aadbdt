package interfaces

type DailyUploadService interface {
	InsertLog(userID int64, pictureID int64, uploadSizeKb uint64) error
}
