package clgt

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
