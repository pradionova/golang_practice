package main 

	type Employee struct {
		name string
		position string
		salary int
		bonus int
	}

	func (r Employee) CalculateTotalSalary() int {
		sum := r.salary + r.bonus
		return sum
	}
