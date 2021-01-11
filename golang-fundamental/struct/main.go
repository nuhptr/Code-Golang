package main

import "fmt"

func main() {
	var user = User{
		ID: 1,
		email: "nugrahaadi733@gmail.com",
		FirstName: "Adi",
		LastName: "Nugraha Putra",
		isActive: true,
	}
	var user2 = User{
		ID: 2,
		email: "makanmalam@gmail.com",
		FirstName: "Fiki",
		LastName: "Naki",
		isActive: false,
	}
	//user.email = "nugrahaadi733@gmail.com"
	//user.FirstName = "Adi"
	//user.LastName = "Nugraha Putra"
	//user.ID = 1
	//user.isActive = true
	fmt.Println(user)


	displayUser := displayUser(user)
	fmt.Println(displayUser)
	// Name : Adi Nugraha Putra, Email: nugrahaadi733@gmail.com

	listUsers:= []User{user, user2}

	group := Group{
		name: "Malika",
		admin: user,
		users: listUsers,
		isAvailable: true,
	}
	group.DisplayGroup()
}
// method
func (group Group) DisplayGroup() {
	fmt.Printf("Name: %s", group.name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.users))
	fmt. Println("")
	fmt.Println("Users name: ")
	for _, user := range group.users {
		fmt.Println(user.FirstName)
	}
}

type User struct {
	ID int // default 0
	FirstName string
	LastName string
	email string
	isActive bool // default false
}

// embedded struct -> menjadi kan struct lain sebagai member
type Group struct {
	name string
	admin User
	users []User
	isAvailable bool
}

// struct as parameter
func displayUser(user User) string {
	result := fmt.Sprintf("Name: %s %s, Email: %s",
		user.FirstName, user.LastName,user.email)
	return result
}