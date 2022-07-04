/* Create a new type vehicle.
	- the underlying type is a struct
	- the fields are doors and color
Create two new types truck and sedan.
	- the underlying type of each type is a struct
	-Embed the vehicle type in both truck and sedan
	-Give truck the field "fourwheel" which will set to bool
	- Give sedan the field luxury which will be set to bool
--> Using the vehicle,truck,sedan structs:
    -using a composite literal, create a value of type truck and assign values to the fields
    -using a composite literal, create a value of type sedan and assign values to the fields
Print each of these values
Printout a single field from each of these fields */

package main

import "fmt"

type vehicle struct {
	doors  int
	colors string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{vehicle: vehicle{doors: 2, colors: "black"},
		fourwheel: true}
	s1 := sedan{vehicle: vehicle{doors: 4, colors: "red"},
		luxury: true}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.fourwheel)
	fmt.Println(s1.luxury)
}
