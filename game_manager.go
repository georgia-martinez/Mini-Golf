package main

type GameManager struct {
	level int
}

func (gm *GameManager) GetGameObjects(level int) []GameObject {

	var gameObjects []GameObject

	switch level {
	case 1:
		gameObjects = append(gameObjects, NewWall(100, 100, 100, 100))
	}

	return gameObjects
}