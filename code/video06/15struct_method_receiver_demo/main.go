package main

import (
	"fmt"
)

type Player struct {
	Name        string
	HealthPoint int
	Level       int
	NowPosition []int
	Prop        []string
}

func NewPlayer(name string, hp, level int, np []int, prop []string) *Player {
	p := Player{
		Name:        name,
		HealthPoint: hp,
		Level:       level,
		NowPosition: np,
		Prop:        prop,
	}
	return &p
}

func (p *Player) Attack() {
	fmt.Printf("%s发起攻击\n", p.Name)
}

func (p *Player) Attacked() { // 这里需要标识为指针接收器，否则无法修改结构体的属性
	fmt.Printf("%s被攻击\n", p.Name)
	p.HealthPoint -= 10
}

func (p *Player) buyProp(propName string) {
	fmt.Printf("%s购买了%s\n", p.Name, propName)
	p.Prop = append(p.Prop, propName)
}

func main() {
	var player1 = NewPlayer("rain", 100, 3, []int{10, 100}, []string{"血瓶", "屠龙宝刀"})
	fmt.Println(player1) // &{rain 100 3 [10 100] [血瓶 屠龙宝刀]}
	player1.Attack()
	player1.Attacked()
	fmt.Println(player1.HealthPoint) // 变为了90

	player1.buyProp("魔法书")
	fmt.Println(player1.Prop) // [血瓶 屠龙宝刀 魔法书]
}
