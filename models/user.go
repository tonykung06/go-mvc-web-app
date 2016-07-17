package models

type User struct {
	id        int
	email     string
	firstName string
	lastName  string
}

func (this *User) Id() int {
	return this.id
}
func (this *User) Email() string {
	return this.email
}
func (this *User) FirstName() string {
	return this.firstName
}
func (this *User) LastName() string {
	return this.lastName
}
func (this *User) SetId(val int) {
	this.id = val
}
func (this *User) SetEmail(val string) {
	this.email = val
}
func (this *User) SetFirstName(val string) {
	this.firstName = val
}
func (this *User) SetLastName(val string) {
	this.lastName = val
}

func GetUsers() []User {
	return []User{
		{
			id:        1,
			email:     "user1@gmail.com",
			firstName: "my first name 1",
			lastName:  "my last name 1",
		},
		{
			id:        2,
			email:     "user2@gmail.com",
			firstName: "my first name 2",
			lastName:  "my last name 2",
		},
		{
			id:        3,
			email:     "user3@gmail.com",
			firstName: "my first name 3",
			lastName:  "my last name 3",
		},
		{
			id:        4,
			email:     "user4@gmail.com",
			firstName: "my first name 4",
			lastName:  "my last name 4",
		},
	}
}
