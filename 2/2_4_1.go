package main

type Game struct {
	On    bool
	Ammo  int
	Power int
}

func (s *Game) Shoot() bool {
	if s.On && s.Ammo > 0 {
		s.Ammo--
		return true
	}
	return false
}

func (s *Game) RideBike() bool {
	if s.On && s.Power > 0 {
		s.Power--
		return true
	}
	return false
}

func main() {

	//testStruct := new(Game)
	/*
	* Экземпляр созданной вами структуры необходимо передать в качестве
	* аргумента функции testStruct, которая выполнит проверку соблюдения
	* всех условий задания/
	 */

}
