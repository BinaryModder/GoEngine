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

## Technology Stack

* **Language:** Go (Golang)
* **GUI:** giu (Dear ImGui)
* **Rendering:** OpenGL 4.1
* **Math:** mathgl
* **Windowing:** GLFW

---

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
go run main.go 
```

---

Run the editor:

```bash
go run main.go -editor -project <ProjectPath>
```

---

## License

MIT-License

---

## Author

Developed by **BinaryModder**.


