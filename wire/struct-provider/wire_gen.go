// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitEndingA(name string) EndingA {
	player := NewPlayer(name)
	monster := NewMonster()
	endingA := EndingA{
		Player:  player,
		Monster: monster,
	}
	return endingA
}

func InitEndingB(name string) EndingB {
	player := NewPlayer(name)
	monster := NewMonster()
	endingB := EndingB{
		Player:  player,
		Monster: monster,
	}
	return endingB
}

// wire.go:

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)

var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "Player", "Monster"))

var endingBSet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingB), "Player", "Monster"))
