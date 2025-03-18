package task01

type User interface {
	GetAge() int
}

type Employee struct {
	Age int
}

type Customer struct {
	Age int 
}

func (e *Employee) GetAge() int {
	return e.Age
}

func (c *Customer) GetAge() int {
	return c.Age
}

func GetOlder(u ...User) int {
	age := u[0].GetAge()
	
	for _, user := range u {
		if user.GetAge() > age {
			age = user.GetAge()
		}
	}
	
	return age
}
