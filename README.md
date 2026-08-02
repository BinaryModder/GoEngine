<img width="512" height="512" alt="GoEngineIcon" src="https://github.com/user-attachments/assets/f942986e-55d6-422b-af5d-0a0fb55d978e" />

# GoEngine

**GoEngine** is a cross-platform 3D game engine and editor written **100% in Go**.

Unlike many engines that rely on C++, C#, or scripting languages for their core, GoEngine is built entirely in Go. The goal of this project is to explore how far a modern game engine can be taken using only the Go ecosystem.

> **100% Golang** — engine, editor, runtime, asset management, project system, and scripting pipeline are all implemented in Go.

---

## Features

* 3D scene editor
* OpenGL renderer
* Scene hierarchy
* Inspector panel
* Project Hub
* Runtime mode
* Project serialization
* Primitive mesh generation

  * Cube
  * Pyramid
  * Sphere (planned)
* Camera controls
* Script attachment system
* Cross-platform architecture
* Fully written in Go

---


## Script Example

```go
package main

import (
	"goengine/bridge/goengine"
)

var speed float64 = 0.2

func Start() {}

func Update() {
	bridge.Rotate(RotateSpeed*bridge.DeltaTime(), RotateSpeed*bridge.DeltaTime(), RotateSpeed*bridge.DeltaTime())
    bridge.Move(0, MoveSpeed*bridge.DeltaTime(), 0)
}

func Destroy() {}
```

## Allowed bridge methods


 * **Self** - self object

 * **Move** (func(dx, dy , dz float64)) - moving object using increment
 * **Rotate** (func(dx, dy, dz float64)) - rotating object using increment

 * **SetPos** (func(dx , dy, dz float64)) - set object position
 * **SetRot** (func(x, y, z float64)) - set object rotation
 * **SetScale** (func(dx, dy, dz)) - set object scale

 * **GetPosX** - float64
 * **GetPosY** - float64
 * **GetPosZ** - float64
 * **GetRotX** - float64
 * **GetRotY** - float64
 * **GetRotZ** - float64

 * **DeltaTime** - float64
 * **Write** - writing to the console


## Technology Stack

* **Language:** [Go (Golang)](https://github.com/golang/go)
* **GUI:** [giu (Dear ImGui)](https://github.com/AllenDang/giu)
* **Rendering:** [OpenGL 4.1](https://www.opengl.org/)
* **Math:** [mathgl](https://github.com/go-gl/mathgl)
* **Windowing:** [GLFW](https://github.com/go-gl/glfw)
* **Logger:** [Zap](https://github.com/uber-go/zap)
* **Interpreter** [Yaegi](https://github.com/traefik/yaegi.git)

## Project Structure

```text
goengine/
├── app/           Launch engine binary in editor/hub/runtime mode as a subprocess
├── cmd/           CLI entry point with flags: -editor, -runtime, -project
├── core/          Core abstractions: State interface (project path + scene access),
│                 filesystem helpers for folder traversal and absolute paths
├── editor/        Editor mode: global EditorState (materials, scene, project files),
│   ├── functions/   business logic for inspector, hierarchy, project browser
│   └── ui/          ImGui panels: viewport, inspector, hierarchy, menu bar,
│                    material create/load popups, project browser, error messages
├── engine/        Engine subsystems:
│   ├── renderer/    OpenGL 4.1: framebuffers, camera, shaders, primitives, textures
│   ├── platform/    OS detection (Windows/macOS/Linux)
│   └── console/     in-engine log console
    └── logger/      zap logger
├── hub/           Project hub/launcher: browse recent projects, create new ones,
│   └── ui/          paginated project cards, settings navigation
├── io/            I/O layer:
│   ├── loader/      JSON readers: scene, material, project config, settings
│   ├── saver/       JSON writers: scene, material, project config
│   └── dialog/      Native file/folder picker wrappers (sqweek/dialog)
├── project/       Domain types: Project, ProjectConfig, ProjectFile;
│                 scaffolding a new project + name/path validation
├── runtime/       Play-in-editor mode: loads current scene + project config,
│   └── ui/          live-rendering viewport
├── scene/         Scene graph: SceneObject, Scene, Transform, Material;
│                 factory primitives (cube, pyramid, sphere, plane); validation
├── script/        Script lifecycle interface (Start, Update, Destroy);
│                 BaseScript for user-defined behaviour
├── settings/      Global engine settings (login, theme); json reading/writing
└── ui/            Reusable utilities: HiDPI-aware scaling, PNG texture decoding
```

## Philosophy

GoEngine is designed around a simple idea:

* Keep the engine lightweight.
* Keep everything written in Go.
* Avoid mixing multiple programming languages.
* Make the source code easy to read, extend, and maintain.
* Exploring Game Engine work

The project focuses on clean architecture rather than unnecessary complexity.

---

## Current Status

GoEngine is under active development.

---

GoEngine demonstrates that a modern 3D engine can be built using **only Go**.

---

Run a project:

```bash
go run ./cmd/main.go 
```

---

Run the editor:

```bash
go run ./cmd/main.go -editor -project <ProjectPath>
```

---

## License

MIT-License

---

## Author

Developed by **BinaryModder**.


