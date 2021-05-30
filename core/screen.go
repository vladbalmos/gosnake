package core

import (
	gc "github.com/rthornton128/goncurses"
)

const GAME_AREA_RATIO = .75

type Point struct {
	X uint
	Y uint
}

type Screen struct {
	mainWindow *gc.Window
	gameWindow *gc.Window
	width uint
	height uint
}

type Key gc.Key

func NewScreen() *Screen {
	mainWindow, err := gc.Init()
	if err != nil {
		panic(err)
	}

	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)

	mainWindow.Timeout(0)
	mainWindow.Keypad(true)

	maxHeight, maxWidth := mainWindow.MaxYX()

	gameWindow := newGameWindow(mainWindow)
	mainWindow.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)

	return &Screen{
		mainWindow: mainWindow,
		gameWindow: gameWindow,
		width: uint(maxWidth),
		height: uint(maxHeight),
	}
}

func (s *Screen) ShowMessage(msg string) {
// CLEATOEOL
}

func (s *Screen) Refresh() {
	gc.Update()
}

func (s *Screen) GetInput() Key {
	return Key(s.mainWindow.GetChar())
}

func (s *Screen) Cleanup() {
	gc.End()
}

func newGameWindow(parent *gc.Window) *gc.Window {
	maxHeight, maxWidth := parent.MaxYX()
	centerX := maxWidth / 2
	centerY := maxHeight / 2

	gameAreaWidth := int(float32(maxWidth) * GAME_AREA_RATIO)
	gameAreaHeight := int(float32(maxHeight) * GAME_AREA_RATIO)
	x := centerX - (gameAreaWidth / 2)
	y := centerY - (gameAreaHeight / 2)
	window := parent.Derived(gameAreaHeight, gameAreaWidth, y, x)
	window.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)
	return window
}
