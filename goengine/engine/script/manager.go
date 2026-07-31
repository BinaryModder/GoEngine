package script

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"goengine/engine/logger"
	"goengine/scene"
)

type Instance struct {
	interp     *interp.Interpreter
	object     *scene.SceneObject
	scriptName string
	updateFn   reflect.Value
}
type Manager struct {
	instances   []*Instance
	basePath    string
	initialized bool
	hasScripts  bool
	currentDT   float64
}

var State = &Manager{}

func (m *Manager) Init(basePath string) {
	m.basePath = basePath
	m.initialized = true
}

func (m *Manager) IsInitialized() bool {
	return m.initialized
}

func (m *Manager) HasScripts() bool {
	return m.hasScripts
}

func (m *Manager) LoadAndStartAll(sceneData *scene.Scene) {
	if !m.initialized || sceneData == nil {
		return
	}

	m.Destroy()

	for i := range sceneData.Objects {
		obj := &sceneData.Objects[i]
		if obj.Script == "" || obj.Script == "No script" {
			continue
		}
		m.startScript(obj)
	}

	m.hasScripts = len(m.instances) > 0
	logger.Info(fmt.Sprintf("Script manager: loaded %d script instances", len(m.instances)))
}

func (m *Manager) startScript(obj *scene.SceneObject) {
	scriptPath := filepath.Join(m.basePath, "Assets", "Scripts", obj.Script+".go")
	src, err := os.ReadFile(scriptPath)
	if err != nil {
		logger.Error(fmt.Sprintf("Script [%s]: failed to read file: %v", obj.Script, err))
		return
	}

	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(interp.Exports{
		"goengine/bridge/goengine/bridge": {
			"Self": reflect.ValueOf(obj),
			"SetPos": reflect.ValueOf(func(x, y, z float64) {
				obj.Transform.Position = [3]float32{float32(x), float32(y), float32(z)}
			}),
			"Move": reflect.ValueOf(func(dx, dy, dz float64) {
				obj.Transform.Position[0] += float32(dx)
				obj.Transform.Position[1] += float32(dy)
				obj.Transform.Position[2] += float32(dz)
			}),
			"SetRot": reflect.ValueOf(func(x, y, z float64) {
				obj.Transform.Rotation = [3]float32{float32(x), float32(y), float32(z)}
			}),
			"Rotate": reflect.ValueOf(func(dx, dy, dz float64) {
				obj.Transform.Rotation[0] += float32(dx)
				obj.Transform.Rotation[1] += float32(dy)
				obj.Transform.Rotation[2] += float32(dz)
			}),
			"SetScale": reflect.ValueOf(func(x, y, z float64) {
				obj.Transform.Scale = [3]float32{float32(x), float32(y), float32(z)}
			}),
			"GetPosX":   reflect.ValueOf(func() float64 { return float64(obj.Transform.Position[0]) }),
			"GetPosY":   reflect.ValueOf(func() float64 { return float64(obj.Transform.Position[1]) }),
			"GetPosZ":   reflect.ValueOf(func() float64 { return float64(obj.Transform.Position[2]) }),
			"GetRotX":   reflect.ValueOf(func() float64 { return float64(obj.Transform.Rotation[0]) }),
			"GetRotY":   reflect.ValueOf(func() float64 { return float64(obj.Transform.Rotation[1]) }),
			"GetRotZ":   reflect.ValueOf(func() float64 { return float64(obj.Transform.Rotation[2]) }),
			"DeltaTime": reflect.ValueOf(func() float64 { return m.currentDT }),
			"Write":     reflect.ValueOf(func(msg string) { logger.Message(msg) }),
		},
	})

	_, err = i.Eval(string(src))
	if err != nil {
		logger.Error(fmt.Sprintf("Script [%s]: compile error: %v", obj.Script, err))
		return
	}

	updateFn, err := i.Eval("Update")
	if err != nil {
		logger.Error(fmt.Sprintf("Script [%s]: cannot find Update function: %v", obj.Script, err))
		return
	}

	inst := &Instance{
		interp:     i,
		object:     obj,
		scriptName: obj.Script,
		updateFn:   updateFn,
	}
	m.instances = append(m.instances, inst)

	_, err = i.Eval("Start()")
	if err != nil {
		logger.Error(fmt.Sprintf("Script [%s]: Start() error: %v", obj.Script, err))
		return
	}

	logger.Info(fmt.Sprintf("Script [%s] started for object [%s]", obj.Script, obj.Name))
}

func (m *Manager) Update(dt float32) {
	m.currentDT = float64(dt)

	if !m.hasScripts {
		return
	}

	for _, inst := range m.instances {
		inst.updateFn.Call(nil)
	}
}

func (m *Manager) Destroy() {
	for _, inst := range m.instances {
		_, _ = inst.interp.Eval("Destroy()")
	}
	m.instances = nil
	m.hasScripts = false
}
