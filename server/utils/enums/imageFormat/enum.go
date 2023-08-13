package enums

type Format int64

const (
	UNDEFINED Format = iota
	PNG
	JPEG
	BMP
)

func (f Format) String() string {
	switch f {
	case PNG:
		return "png"
	case JPEG:
		return "jpeg"
	case BMP:
		return "bmp"
	}
	return "jpeg"
}
