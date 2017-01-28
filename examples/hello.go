package main

import . "github.com/fogleman/fauxgl"

const (
	width  = 2000
	height = 1000
	fovy   = 20
	near   = 1
	far    = 50
)

var (
	eye    = V(-1, -2, 2)
	center = V(-0.07, 0, 0)
	up     = V(0, 0, 1)
)

func main() {
	mesh, err := LoadSTL("examples/hello.stl")
	if err != nil {
		panic(err)
	}
	mesh.BiUnitCube()
	mesh.SmoothNormalsThreshold(Radians(30))

	context := NewContext(width, height)
	context.ClearColorBuffer(V(0, 0, 0))

	aspect := float64(width) / float64(height)
	matrix := LookAt(eye, center, up).Perspective(fovy, aspect, near, far)
	light := V(-2, 0, 1).Normalize()
	color := V(0.5, 1, 0.65)

	shader := NewDefaultShader(matrix, light, eye, color)
	context.DrawMesh(mesh, shader)

	SavePNG("out.png", context.Image())
}