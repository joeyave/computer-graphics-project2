package other

import (
	"github.com/joeyave/computer-graphics-project2/utils"
	"math"
	"time"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
	"github.com/joeyave/computer-graphics-project2/app"
)

func init() {
	app.DemoMap["other.pitch"] = &Pitch{}
}

const otherPitchHelp = `
W - вперед
S - назад
A - влево
D - вправо
I - нос вверх
K - нос вниз
J - наклон влево
L - наклон вправо
R - исходная позиция
`

type Pitch struct {
	base *graphic.Mesh
}

// Start is called once at the start of the demo.
func (p *Pitch) Start(a *app.App) {

	// Subscribe to key events
	a.Subscribe(window.OnKeyRepeat, p.onKey)
	a.Subscribe(window.OnKeyDown, p.onKey)

	// Add help label
	label1 := gui.NewLabel(otherPitchHelp)
	label1.SetFontSize(16)
	label1.SetPosition(10, 10)
	a.DemoPanel().Add(label1)

	// Top directional light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 0.5)
	l1.SetPosition(0, 1, 0)
	a.Scene().Add(l1)

	// Creates plane base mesh
	base_geom := geometry.NewDisk(1, 3)
	base_mat := material.NewStandard(&math32.Color{0, 1, 0})
	base_mat.SetWireframe(false)
	base_mat.SetSide(material.SideDouble)
	p.base = graphic.NewMesh(base_geom, base_mat)

	vert_geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	normals := math32.NewArrayF32(0, 0)
	indices := math32.NewArrayU32(0, 0)
	positions.Append(0, 0, 0, 1, 0, 0, 0, 1, 0)
	normals.Append(0, 0, 1, 0, 0, 1, 0, 0, 1)
	indices.Append(0, 1, 2)

	vert_geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	vert_geom.AddVBO(gls.NewVBO(normals).AddAttrib(gls.VertexNormal))
	vert_geom.SetIndices(indices)

	vert_mat := material.NewStandard(&math32.Color{0, 0, 1})
	vert_mat.SetSide(material.SideDouble)
	vert := graphic.NewMesh(vert_geom, vert_mat)
	vert.SetScale(1.5, 1, 1)
	vert.SetPosition(-0.5, 0, 0)
	vert.SetRotationX(math.Pi / 2)
	p.base.Add(vert)

	p.base.SetScale(3, 1, 1)
	p.base.SetRotationX(-math.Pi / 2)
	p.base.SetPosition(0, 0, 0)

	a.Scene().Add(p.base)

	// Add animation controls
	g1 := a.ControlFolder().AddGroup("Animation")
	cb1 := g1.AddCheckBox("Paused").SetValue(true)
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		if paused == false {
			paused = true
		} else {
			paused = false
		}
	})

	cam := a.Camera()
	cam.SetPosition(-3, 3, 3)
	pos := cam.Position()
	cam.UpdateSize(pos.Length())
	spos := a.Scene().Position()
	cam.LookAt(&spos, &math32.Vector3{0, 1, 0})

	// Create axes helper
	axes := helper.NewAxes(3)
	a.Scene().Add(axes)
}

func (p *Pitch) onKey(evname string, ev interface{}) { // TODO use keystate?

	kev := ev.(*window.KeyEvent)

	var q math32.Quaternion
	xaxis := math32.Vector3{1, 0, 0}
	yaxis := math32.Vector3{0, 1, 0}
	zaxis := math32.Vector3{0, 0, 1}
	step := float32(0.05)

	// Pitch forward
	if kev.Key == window.KeyW {

		// Get world direction
		var quat math32.Quaternion
		p.base.Node.WorldQuaternion(&quat)
		direction := math32.Vector3{1, 0, 0}
		utils.ApplyQuaternionToVector(&quat, &direction)
		utils.NormalizeVector(&direction)
		utils.MultiplyScalarVector(step, &direction)

		// Get world position
		var position math32.Vector3
		p.base.Node.WorldPosition(&position)
		position.Add(&direction)
		p.base.Node.SetPositionVec(&position)
	}

	// Pitch backward
	if kev.Key == window.KeyS {

		var quat math32.Quaternion
		p.base.Node.WorldQuaternion(&quat)
		direction := math32.Vector3{1, 0, 0}
		utils.ApplyQuaternionToVector(&quat, &direction)
		utils.NormalizeVector(&direction)
		utils.MultiplyScalarVector(step, &direction)

		utils.NegateVector(&direction)

		var position math32.Vector3
		p.base.Node.WorldPosition(&position)
		position.Add(&direction)
		p.base.Node.SetPositionVec(&position)
	}

	// Pitch up
	if kev.Key == window.KeyI {
		utils.SetAxisAngle(&q, &yaxis, -step)
		baseQ := p.base.Quaternion()
		rotQ := utils.MultiplyQuaternions(&baseQ, &q)
		p.base.SetRotationQuat(rotQ)
	}
	// Pitch down
	if kev.Key == window.KeyK {
		utils.SetAxisAngle(&q, &yaxis, step)
		baseQ := p.base.Quaternion()
		rotQ := utils.MultiplyQuaternions(&baseQ, &q)
		p.base.SetRotationQuat(rotQ)
	}

	// Heading left
	if kev.Key == window.KeyA {
		utils.SetAxisAngle(&q, &zaxis, step)
		baseQ := p.base.Quaternion()
		rotQ := utils.MultiplyQuaternions(&baseQ, &q)
		p.base.SetRotationQuat(rotQ)
	}

	// Heading right
	if kev.Key == window.KeyD {
		utils.SetAxisAngle(&q, &zaxis, -step)
		baseQ := p.base.Quaternion()
		rotQ := utils.MultiplyQuaternions(&baseQ, &q)
		p.base.SetRotationQuat(rotQ)
	}

	// Bank left
	if kev.Key == window.KeyJ {
		utils.SetAxisAngle(&q, &xaxis, -step)
		baseQ := p.base.Quaternion()
		rotQ := utils.MultiplyQuaternions(&baseQ, &q)
		p.base.SetRotationQuat(rotQ)
	}

	// Bank right
	if kev.Key == window.KeyL {
		utils.SetAxisAngle(&q, &xaxis, step)
		baseQ := p.base.Quaternion()
		rotQ := utils.MultiplyQuaternions(&baseQ, &q)
		p.base.SetRotationQuat(rotQ)
	}

	// Reset
	if kev.Key == window.KeyR {
		p.base.SetRotation(-math.Pi/2, 0, 0)
		p.base.SetPosition(0, 0, 0)
	}

	if kev.Key == window.KeyKPAdd {
		scaleVec := p.base.Scale()
		utils.MultiplyScalarVector(1.05, &scaleVec)
		p.base.SetScaleVec(&scaleVec)
	}

	if kev.Key == window.KeyKPSubtract {
		scaleVec := p.base.Scale()
		utils.DivideScalarVector(1.05, &scaleVec)
		p.base.SetScaleVec(&scaleVec)
	}
}

var timeSince float64
var paused = true

// Update is called every frame.
func (p *Pitch) Update(a *app.App, deltaTime time.Duration) {
	if paused == false {
		timeSince += deltaTime.Seconds()

		x := timeSince
		y := math.Cos(timeSince)
		z := math.Sin(timeSince)

		p.base.SetPosition(float32(x), float32(y), float32(z))
	}
}

// Cleanup is called once at the end of the demo.
func (p *Pitch) Cleanup(a *app.App) {}
