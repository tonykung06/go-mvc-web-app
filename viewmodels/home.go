package viewmodels

type Home struct {
	Title   string
	Message string
	Path    string
	Fruits  []string
}

type Login struct {
	Title string
	Path  string
}

var fruits = []string{"Apple", "Orange", "Banana"}

func GetLogin() Login {
	return Login{
		Title: "Login Now",
		Path:  "/login",
	}
}

func GetHome() Home {
	return Home{
		Title:   "Home Page",
		Message: "the message",
		Path:    "/home",
		Fruits:  fruits,
	}
}

func GetHomeById(fruitId int) Home {
	fruits, _ := GetFruits(fruitId)
	return Home{
		Title:   "Home Page",
		Message: "the message",
		Path:    "/home",
		Fruits:  fruits,
	}
}

func GetFruits(fruitId int) ([]string, error) {
	if fruitId < len(fruits) {
		return []string{fruits[fruitId]}, nil
	}
	return nil, &fruitNotFound{"We don't have this fruit"}
}

type fruitNotFound struct {
	s string
}

func (e *fruitNotFound) Error() string {
	return e.s
}
