package main

type John struct {
	On    bool
	Ammo  int
	Power int
}

func (j *John) Shoot() bool {
	if j.On == false {
		return false
	}
	if j.Ammo < 1 {
		return false
	}
	j.Ammo = j.Ammo - 1
	return true
}

func (j *John) RideBike() bool {
	if j.On == false {
		return false
	}
	if j.Power < 1 {
		return false
	}
	j.Power = j.Power - 1
	return true
}

func main() {
	testStruct := &John{}
}
