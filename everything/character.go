package everything

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name        string
	Level       int
	Exp         int
	Health      int
	BaseAttack  int
	BaseDefence int
	Weapon      WeaponS
	Armor       ArmorS
	Props       PropsS
}

// TODO jszhou 完善武器、护甲、道具
type WeaponS struct {
}
type ArmorS struct {
}
type PropsS struct {
}
type Booty struct {
}

// TODO jszhou 背包机制
func Combat(adventurer, monster Character) (bool, int, Booty) {
	rand.Seed(time.Now().UnixNano())

	var res bool
	for {
		OneCombat(&adventurer, &monster)
		time.Sleep(200 * time.Millisecond)
		if monster.Health <= 0 {
			fmt.Println("win")
			fmt.Println("player's heath is: ", adventurer.Health)
			fmt.Println("darkness's heath is: ", monster.Health)
			res = true
			break
		}
		if adventurer.Health <= 0 {
			fmt.Println("waste")
			fmt.Println("player's heath is: ", adventurer.Health)
			fmt.Println("darkness's heath is: ", monster.Health)
			res = false
			break
		}
	}
	time.Sleep(1 * time.Second)

	exp := monster.Level*3 + 50

	return res, exp, Booty{}
}

// TODO jszhou 攻击 可以设置上范围内上下浮动的机制
// TODO jszhou 防御可以修改成 减少 f(defence) % 的形式，而不是减少绝对值
func OneCombat(adventurer, monster *Character) {
	fmt.Printf("玩家血量为: %d 敌人血量为: %d\n", adventurer.Health, monster.Health)

	if adventurer.BaseAttack > monster.BaseDefence {
		monster.Health = monster.Health - (adventurer.BaseAttack - monster.BaseDefence)
	} else {
		monster.Health = monster.Health - 1
	}

	if monster.BaseAttack > adventurer.BaseDefence {
		adventurer.Health = adventurer.Health - (monster.BaseAttack - adventurer.BaseDefence)
	} else {
		adventurer.Health = adventurer.Health - 1
	}
}

func InitAdventurer() Character {
	return Character{
		Name:        "Adventurer",
		Level:       1,
		Exp:         0,
		Health:      100,
		BaseAttack:  10,
		BaseDefence: 5,
	}
}

// TODO jszhou monster根据level生成的数值上下浮动
func Encounter(level int) Character {
	return Character{
		Name:        "Monster",
		Health:      95 + rand.Intn(10),
		BaseAttack:  8 + rand.Intn(4),
		BaseDefence: 3 + rand.Intn(4),
	}
}

func Display(adventurer, monster Character) {
	fmt.Printf("冒险者的属性如下：\n")
	fmt.Printf("health: %d\n", adventurer.Health)
	fmt.Printf("level: %d\n", adventurer.Level)
	fmt.Printf("exp: %d\n", adventurer.Exp)
	fmt.Printf("Attack: %d\n", adventurer.BaseAttack)
	fmt.Printf("Defence: %d\n\n", adventurer.BaseDefence)
	fmt.Printf("敌怪的属性如下：\n")
	fmt.Printf("health: %d\n", monster.Health)
	fmt.Printf("Attack: %d\n", monster.BaseAttack)
	fmt.Printf("Defence: %d\n\n", monster.BaseDefence)
}

func Develop(adventurer *Character, exp int) {
	exp, level := levelUp(adventurer.Exp+exp, adventurer.Level)
	adventurer.Exp = exp
	adventurer.Level = level
}

// TODO jszhou 将经验机制完善
func levelUp(exp, level int) (int, int) {
	if exp-(level*50+50) < 0 {
		return exp, level
	}

	for exp-(level*50+50) >= 0 {
		level += 1
	}

	return exp - level*50, level

}
