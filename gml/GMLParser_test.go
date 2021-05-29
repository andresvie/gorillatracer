package gml

import (
	"strings"
	"testing"

	"github.com/andresvie/gorillatracer/geometry"
	"github.com/andresvie/gorillatracer/light"
	"github.com/andresvie/gorillatracer/vector"
)

func TestParseSphere(t *testing.T) {
	gmlText := `
	sphere{
		center = (0,0,-2)
		color = (1.0,0,0)
		radius = 0.5
		specular = 1000.2
	}`
	reader := strings.NewReader(gmlText)
	scene, err := ParseGML(reader)
	if err != nil {
		t.Fatalf("expected parse valid scene %s %v", gmlText, err)
	}
	if len(scene.Objects) != 1 {
		t.Fatalf("expected one sphere and got %d", len(scene.Objects))
	}
	var geo geometry.Geometry = scene.Objects[0]
	sphere, ok := geo.(*geometry.Sphere)
	if !ok {
		t.Fatalf("expected parse object to be sphere")
	}
	assertVector("center", sphere.Center, &vector.Vector{X: 0, Y: 0, Z: -2}, t)
	assertVector("color", sphere.Color, &vector.Vector{X: 1.0, Y: 0, Z: 0}, t)
	if sphere.Radius != 0.5 {
		t.Fatalf("expected %s %v to be equal %v", "radius", sphere.Radius, 0.5)
	}
	if sphere.SpecularFactor != 1000.2 {
		t.Fatalf("expected %s %v to be equal %v", "specular", sphere.SpecularFactor, 1000.2)
	}
}

func TestParsePointOfLight(t *testing.T) {
	gmlText := `
	pointOfLight{
		point = (0,5,0)		
		intensity = 0.9
	}`
	reader := strings.NewReader(gmlText)
	scene, err := ParseGML(reader)
	if err != nil {
		t.Fatalf("expected parse valid scene %s %v", gmlText, err)
	}
	if len(scene.Lights) != 1 {
		t.Fatalf("expected one light and got %d", len(scene.Lights))
	}
	var l light.Light = scene.Lights[0]
	pointOfLight, ok := l.(*light.PointOfLight)
	if !ok {
		t.Fatalf("expected parse object to be pointOfLight")
	}
	assertVector("point", pointOfLight.Point, &vector.Vector{X: 0, Y: 5, Z: 0, W: 0}, t)

	if pointOfLight.Intensity != 0.9 {
		t.Fatalf("expected %s %v to be equal %v", "intensity", pointOfLight.Intensity, 0.9)
	}

}

func assertVector(fieldName string, v, expectedVector *vector.Vector, t *testing.T) {
	if v.X != expectedVector.X || v.Y != expectedVector.Y || v.Z != expectedVector.Z {
		t.Fatalf("expected %s %v to be equal %v", fieldName, v, expectedVector)
	}
}
