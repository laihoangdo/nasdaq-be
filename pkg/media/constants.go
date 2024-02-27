package media

// FileType defines file types allowed for media
type FileType string

// Media file types
const (
	FileTypePNG FileType = "png"
)

const (
	DefaultImgWidth   = 1920
	DefaultImgHeight  = 0
	DefaultImgQuality = 100
)

const (
	defaultWkhtmltoimageBin   = "/usr/local/bin/wkhtmltoimage"
	bucketClassScheduleImages = "class-schedules-images"
	pathClassSchedule         = "/class-schedules"
)
