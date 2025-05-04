package sample

// Type embedding is a technique where you include one type into another
// but as a parameter without a name.

import "fmt"

// Structs are used to create complex data types
// Structs are used to group related data together

type Circle struct {
	length  float64
	breadth float64
	color   string
}

type Thermostat struct {
	ID    string
	value float64
}

//Value return the current temperature in Celsius
func (t *Thermostat) Value() float64 {
	return t.value
}

//Set tells the thermostat to set the temperature
func (t *Thermostat) Set(value float64) {
	t.value = value
}

//Kind returns the device kind
func (t *Thermostat) Kind() string {
	return "thermostat"
}

func main2() {

	fmt.Println("Wihtout instance: ", Circle{10, 20, "red"})

	circle3 := new(Circle) /* circle3 is a pointer to an instance of circle struct */
	circle3.length = 10
	circle3.breadth = 20
	circle3.color = "red"
	fmt.Println("With instance: ", circle3)

	var circle6 = &Circle{}
	circle6.length = 10
	circle6.color = "Red"
	fmt.Println(circle6) //value for breadth is skipped

	var circle7 = &Circle{}
	(*circle7).breadth = 10
	(*circle7).color = "Blue"
	fmt.Println(circle7) //value for length is skipped

	// func(reciver_name Type) method_name(parameter_list) (return_type){}

	t := Thermostat{"Living Room", 16.2}
	fmt.Printf("%s Before: %.2f\n", t.ID, t.Value())
	t.Set(20.0)
	fmt.Printf("%s After: %.2f\n", t.ID, t.Value())

	// Interfaces are used to define a contract for a type
	// Interfaces are used to define a set of methods that a type must implement
	// type interface_name interface{}

	// Accessing Elements of the Shape Interface
	var s Shape
	s = Rectangle{10, 14}
	fmt.Println("Area of Shape :", s.Area())
	fmt.Println("Perimeter of Shape:", s.Perimeter())

	c := Camera{"Baby Room"}
	sensors := []Sensor{&t, &c}
	printAll(sensors)

}

// Interfaces

// Creating an Interface
type Shape interface {
	// Method Signatures
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

// Implementing Methods of the Shape Interface
func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

// Camera

//Camera is a security camera
type Camera struct {
	id string
}

type Sensor interface {
	ID() string
	Kind() string
}

//ID return the camera ID
func (c *Camera) ID() string {
	return c.id
}

func (*Camera) Kind() string {
	return "camera"
}

func printAll(sensors []Sensor) {
	for _, s := range sensors {
		fmt.Printf("%s <%s>\n", s.ID(), s.Kind())
	}
}
