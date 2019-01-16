package components

// CameraTarget contains the data to move the 2D camera around.
type CameraTarget struct {
	X        float32
	Y        float32
	OffsetX  float32
	OffsetY  float32
	Rotation float32
	Zoom     float32
}

func (c *CameraTarget) Name() string {
	return "cameraTarget"
}
