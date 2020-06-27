package serializer

type Response struct {
	Status int
	Data   interface{}
	Msg    string
	Error  string
}
