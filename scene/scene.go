package scene

import (
	"io"

	"github.com/andresvie/gorillatracer/camera"
	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/light"
	"github.com/andresvie/gorillatracer/utils"
)

type Scene struct {
	Camera  *camera.Camera
	Lights  []light.Light
	Objects []geometry.Geometry
}

func (s *Scene) Render(reader io.Reader) [][]utils.REAL {

	return nil
}
