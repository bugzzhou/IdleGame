package everything

import (
	"encoding/json"
	"fight/gotool/file"
	"fight/models"
	"fmt"
	"math/rand"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

const saveDir = "./saves"

var savePath = filepath.Join(saveDir, "save.json")

// TODO jszhou 背包机制
func Combat(adventurer, monster models.Character) (bool, int, models.Booty) {
	rand.Seed(time.Now().UnixNano())

	var res bool
	for {
		OneCombat(&adventurer, &monster)

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
	// time.Sleep(1 * time.Second)

	exp := monster.Level*3 + 50

	return res, exp, models.Booty{}
}

// TODO jszhou 攻击 可以设置上范围内上下浮动的机制
// TODO jszhou 防御可以修改成 减少 f(defence) % 的形式，而不是减少绝对值
func OneCombat(adventurer, monster *models.Character) {
	// time.Sleep(200 * time.Millisecond)
	// fmt.Printf("玩家血量为: %d 敌人血量为: %d\n", adventurer.Health, monster.Health)

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

func InitAdventurer() models.Character {
	return models.Character{
		Name:        "Adventurer",
		Level:       1,
		Exp:         0,
		Health:      100,
		BaseAttack:  10,
		BaseDefence: 5,
	}
}

// TODO jszhou monster根据level生成的数值上下浮动
func Encounter(level int) models.Character {
	return models.Character{
		Name:        "Monster",
		Health:      95 + rand.Intn(10),
		BaseAttack:  8 + rand.Intn(4),
		BaseDefence: 3 + rand.Intn(4),
	}
}

func Display(adventurer, monster models.Character) {
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

func Develop(adventurer *models.Character, exp int) {
	exp, level, attack, defence := levelUp(adventurer.Exp+exp, adventurer.Level, adventurer.BaseAttack, adventurer.BaseDefence)
	adventurer.Exp = exp
	adventurer.Level = level
	adventurer.BaseAttack = attack
	adventurer.BaseDefence = defence
}

// TODO jszhou 将经验机制完善
func levelUp(exp, level, attack, defence int) (int, int, int, int) {
	if exp-(level*50+50) < 0 {
		return exp, level, attack, defence
	}

	for exp-(level*50+50) >= 0 {
		level += 1
		attack += (rand.Intn(2) + 1)
		defence += rand.Intn(2)
	}

	return exp - level*50, level, attack, defence

}

func HandleUserInput(adventurer *models.Character, monster models.Character) bool {
	var decide string

	fmt.Println("是否应对敌怪？")
	fmt.Scan(&decide)

	isNumber, _ := regexp.MatchString(`^\d+$`, decide)
	if isNumber {
		n, _ := strconv.Atoi(decide)
		return CombatMultipleTimes(adventurer, monster, n)
	}

	switch decide {
	case "y", "Y":
		return CombatMultipleTimes(adventurer, monster, 1)
	case "n", "N":
		return true
	case "x", "X":
		file.Create(savePath, *adventurer)
		return false
	default:
		fmt.Println("选择无效，重新选择，自动为你跳过该敌怪")
		return true
	}
}

// CombatMultipleTimes 函数执行多次战斗并处理结果
func CombatMultipleTimes(adventurer *models.Character, monster models.Character, times int) bool {
	for i := 0; i < times; i++ {
		success, exp, _ := Combat(*adventurer, monster)
		if success {
			Develop(adventurer, exp)
		} else {
			fmt.Println("输了")
			return false
		}
	}
	return true
}

func GetAdventurer() (models.Character, error) {
	adventurer := models.Character{}

	if file.IsExist(savePath) {
		a := file.GetBytesByPath(savePath)
		err := json.Unmarshal(a, &adventurer)
		if err != nil {
			fmt.Printf("failed to unmarshal, and err is: %s\n", err.Error())
			return adventurer, err
		}
	} else {
		adventurer = InitAdventurer()
		file.Create(savePath, adventurer)
	}

	return adventurer, nil
}
