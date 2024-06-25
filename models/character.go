package models

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
