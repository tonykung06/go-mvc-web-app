package viewmodels

type Home struct {
	Title   string
	Message string
	Path    string
	Fruits  [3]string
}

func GetHome() Home {
	return Home{
		Title:   "Home Page",
		Message: "the message",
		Path:    "/Home",
		Fruits:  [3]string{"Apple", "Orange", "Banana"},
	}
}
