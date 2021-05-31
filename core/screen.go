package core

import (
	//"fmt"
	gc "github.com/rthornton128/goncurses"
)

const GAME_AREA_RATIO = .75

type Screen struct {
	mainWindow *gc.Window
	gameWindow *gc.Window
	width int
	height int
	GameWindowX int
	GameWindowY int
	GameAreaWidth int
	GameAreaHeight int
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

	center := &Point{
		X: maxWidth / 2,
		Y: maxHeight / 2,
	}

	gameWindow, gameWindowX, gameWindowY, gameAreaWidth, gameAreaHeight := newGameWindow(mainWindow, center)
	mainWindow.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)

	return &Screen{
		mainWindow: mainWindow,
		gameWindow: gameWindow,
		width: maxWidth,
		height: maxHeight,
		GameWindowX: gameWindowX,
		GameWindowY: gameWindowY,
		GameAreaWidth: gameAreaWidth,
		GameAreaHeight: gameAreaHeight,
	}
}

func (s *Screen) Center() Point {
	p := Point{
		X: s.width / 2,
		Y: s.height / 2,
	}

	return p
}

func (s *Screen) GameAreaCenter() Point {
	p := Point{
		X: s.GameAreaWidth / 2,
		Y: s.GameAreaHeight / 2,
	}

	return p
}

func (s *Screen) HideMessage() {
	charsToClear := s.width - 2
	for i := 0; i < charsToClear; i++ {
		s.mainWindow.MovePrint(s.GameWindowY - 2, i + 1, " ")
	}
}

func (s *Screen) ShowMessage(msg string) {
	y := s.GameWindowY - 2
	x := s.Center().X - len(msg) / 2
	s.mainWindow.MovePrint(y, x, msg)
}

func (s *Screen) ShowScore(score uint) {
	y := s.GameWindowY + s.GameAreaHeight + 1
	x := s.GameWindowX
	s.mainWindow.MovePrint(y, x, "Score: ", score)
}

func (s *Screen) DrawFood(coords Point) {
	s.gameWindow.MovePrint(coords.Y, coords.X, "*");
}

func (s *Screen) DrawSnakeSegment(coords Point) {
	s.gameWindow.MovePrint(coords.Y, coords.X, "*");
}

func (s *Screen) EraseGameArea() {
	s.gameWindow.Erase()
}

func (s *Screen) Refresh() {
	s.mainWindow.Refresh()
	s.gameWindow.Refresh()
}

func (s *Screen) GetInput() Key {
	return Key(s.mainWindow.GetChar())
}

func (s *Screen) Cleanup() {
	gc.End()
}

func newGameWindow(parent *gc.Window, center *Point) (*gc.Window, int, int, int, int) {
	maxHeight, maxWidth := parent.MaxYX()
	centerX := int(center.X)
	centerY := int(center.Y)

	gameAreaWidth := int(float32(maxWidth) * GAME_AREA_RATIO)
	gameAreaHeight := int(float32(maxHeight) * GAME_AREA_RATIO)
	x := centerX - (gameAreaWidth / 2)
	y := centerY - (gameAreaHeight / 2)
	container := parent.Derived(gameAreaHeight, gameAreaWidth, y, x)
	container.Border(gc.ACS_VLINE, gc.ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE, gc.ACS_ULCORNER, gc.ACS_URCORNER, gc.ACS_LLCORNER, gc.ACS_LRCORNER)

	window := container.Derived(gameAreaHeight - 2, gameAreaWidth - 2, 1, 1)
	return window, x + 1, y + 1, gameAreaWidth - 2, gameAreaHeight - 2
}
