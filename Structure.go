package main

import "fmt"

type nspevent struct {
	date    string
	name    string
	contact int
	members organizingteam
}

type organizingteam struct {
	names       []string
	fun_contact int
}

func main() {

	nsp1 := nspevent{
		date:    "15jan2024",
		name:    "Sankranthi",
		contact: 7001,
		members: organizingteam{
			names:       []string{"NOKIA", "ALU"},
			fun_contact: 9999,
		},
	}

	nsp2 := nspevent{
		date:    "26jan2024",
		name:    "RepublicDay",
		contact: 7002,
		members: organizingteam{
			names:       []string{"INDIA", "BHARATH"},
			fun_contact: 8888,
		},
	}
	//fmt.Println(nsp1.date , nsp1.name ,nsp1.members.fun_contact, ,nsp1.members.names)
	fmt.Println(nsp1.date, nsp1.name, nsp1.members.fun_contact)
	fmt.Println(nsp1.date, nsp1.name, nsp1.contact, nsp1.members.names, nsp1.members.fun_contact)
	fmt.Println("Second Organization Members")
	fmt.Println(nsp2.members.fun_contact)
	fmt.Println("Structures")
	fmt.Println(nsp2.date, nsp2.name, nsp2.contact, nsp2.members.names, nsp2.members.fun_contact)

}
