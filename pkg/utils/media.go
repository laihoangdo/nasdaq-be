package utils

const (
	byteUnit = 1
	kilobyte = 1024 * byteUnit
	megabyte = 1024 * kilobyte
)

const (
	maxImgSize = 2 * megabyte
)

func IsValidImageSize(size int64) bool {
	return size <= maxImgSize
}
