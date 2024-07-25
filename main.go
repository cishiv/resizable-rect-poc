package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 600
	sidebarWidth = 200
)

type Shape struct {
	rect     rl.Rectangle
	color    rl.Color
	dragging bool
	resizing bool
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Drag and Resize Shapes")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	shapes := []Shape{}
	sidebarShapes := []Shape{
		{rect: rl.Rectangle{X: 50, Y: 50, Width: 100, Height: 100}, color: rl.Red},
		{rect: rl.Rectangle{X: 50, Y: 200, Width: 100, Height: 50}, color: rl.Blue},
		{rect: rl.Rectangle{X: 50, Y: 350, Width: 50, Height: 100}, color: rl.Green},
	}

	resizeHandle := rl.Rectangle{Width: 10, Height: 10}

	for !rl.WindowShouldClose() {
		mousePos := rl.GetMousePosition()

		// Handle shape dragging and resizing
		for i := range shapes {
			handleShapeInteraction(&shapes[i], mousePos, resizeHandle)
		}

		// Handle dragging shapes from sidebar
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			for _, sidebarShape := range sidebarShapes {
				if rl.CheckCollisionPointRec(mousePos, sidebarShape.rect) {
					newShape := sidebarShape
					newShape.rect.X = sidebarWidth + 10
					newShape.rect.Y = 10
					newShape.dragging = true
					shapes = append(shapes, newShape)
					break
				}
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Draw sidebar
		rl.DrawRectangle(0, 0, sidebarWidth, screenHeight, rl.LightGray)

		// Draw sidebar shapes
		for _, shape := range sidebarShapes {
			rl.DrawRectangleRec(shape.rect, shape.color)
		}

		// Draw shapes on main screen
		for _, shape := range shapes {
			rl.DrawRectangleRec(shape.rect, shape.color)
			drawResizeHandle(shape, resizeHandle)
		}

		rl.EndDrawing()
	}
}

func handleShapeInteraction(shape *Shape, mousePos rl.Vector2, resizeHandle rl.Rectangle) {
	resizeHandle.X = shape.rect.X + shape.rect.Width - resizeHandle.Width/2
	resizeHandle.Y = shape.rect.Y + shape.rect.Height - resizeHandle.Height/2

	if rl.CheckCollisionPointRec(mousePos, resizeHandle) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			shape.resizing = true
		}
	} else if rl.CheckCollisionPointRec(mousePos, shape.rect) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			shape.dragging = true
		}
	}

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		shape.dragging = false
		shape.resizing = false
	}

	if shape.dragging {
		shape.rect.X = mousePos.X - shape.rect.Width/2
		shape.rect.Y = mousePos.Y - shape.rect.Height/2
	}

	if shape.resizing {
		shape.rect.Width = mousePos.X - shape.rect.X
		shape.rect.Height = mousePos.Y - shape.rect.Y

		if shape.rect.Width < 20 {
			shape.rect.Width = 20
		}
		if shape.rect.Height < 20 {
			shape.rect.Height = 20
		}
	}
}

func drawResizeHandle(shape Shape, resizeHandle rl.Rectangle) {
	resizeHandle.X = shape.rect.X + shape.rect.Width - resizeHandle.Width/2
	resizeHandle.Y = shape.rect.Y + shape.rect.Height - resizeHandle.Height/2
	rl.DrawRectangleRec(resizeHandle, rl.Black)
}
