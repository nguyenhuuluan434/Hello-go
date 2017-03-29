package luannh

//Animal interface
type Animal interface {
	Speak() string
}

//Dog struct
type Dog struct {
}

//Speak string
func (d Dog) Speak() string {
	return "Woof!"
}

//Cat struct
type Cat struct {
}

//Speak string
func (c Cat) Speak() string {
	return "Meow!"
}

//Llama struct
type Llama struct {
}

//Speak ????
func (l Llama) Speak() string {
	return "?????"
}

//JavaProgrammer struct
type JavaProgrammer struct {
}

//Speak string
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

type luanga struct {
}

func (a luanga) Speak() string {
	return "13579"
}

//Stack is a struct have data with a generic type,
type Stack struct {
	data []interface{}
}

//NewStack func
func NewStack() *Stack {
	s := Stack{data: make([]interface{}, 0)}
	return &s
}

//Push func
func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)
}

//Pop func
func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return data
}
