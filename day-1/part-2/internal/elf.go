package internal

type Elf struct {
	CarriedCalories int
}

func NewElf() *Elf {
	return &Elf{}
}

func (e *Elf) AddCalories(caloriesToAdd int) {
	e.CarriedCalories += caloriesToAdd
}
