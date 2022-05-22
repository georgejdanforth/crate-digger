package server

func Init() {
	r := NewRouter()
	// TODO: make port configurable
	r.Run()
}
