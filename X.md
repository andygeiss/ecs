In the first step we have to be clear about what our game engine should do.
The main task is to make sure that we have all the essential components that are necessary for the technical and logical aspects 
are responsible, are combined with each other.

An Entity Component System (ECS) helps us to do just that, as the logical components (data) such as entities and their components
can be separated from the actual logic. 
One of the advantages of this is that we can implement and test the game mechanics independently of the rest.
So let's start...

We decide to use 2D and define the **three most important components**:
* Position
* Size
* Velocity

We store these as <code>components.go</code> (Example: [here](https://github.com/andygeiss/ecs/examples/engine/components.go)).

In the next step, the **three most important systems** implement
* Collision
* Movement
* Rendering

We store these as <code>systems.go</code> (Example: [here](https://github.com/andygeiss/ecs/examples/engine/systems.go)

The collision and movement system contains the actual game mechanics:

```go
func (m *Collision) Process(em *ecs.EntityManager) (state int) {
	for _, entity := range em.FilterByMask(MaskPosition | MaskVelocity) {
		position := entity.Get(MaskPosition).(*Position)
		velocity := entity.Get(MaskVelocity).(*Velocity)
		if position.X >= m.width || position.X <= 0 {
			velocity.X = -velocity.X
		}
		if position.Y >= m.height || position.Y <= 0 {
			velocity.Y = -velocity.Y
		}
	}
	return ecs.StateEngineContinue
}
```
The rendering system must be adapted to a specific game library.
In our example we have used [SDL](https://github.com/veandco/go-sdl2)
In the example of [Pong](https://github.com/andygeiss/ecs-pong)) we used [Raylib](https://github.com/gen2brain/raylib-go).

Finally we create a <code>main.go</code> file (Example: [here](https://github.com/andygeiss/ecs/examples/main.go)
and link the systems together:

```go
func run() {
	em := ecs.NewEntityManager()
	em.Add(generateEntities(1000)...)
	sm := ecs.NewSystemManager()
	sm.Add(
		engine.NewMovement(),
		engine.NewCollision(Width, Height),
		engine.NewRendering(Width, Height, "ECS with SDL Demo"),
	)
	ecs.Run(em, sm)
}

func main() {
	ecs.Main(func() {
		run()
	})
}
```