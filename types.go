package main

type Phones map[string]string
type EMails map[string]string

type Person struct {
        FirstName string
        LastName  string
        Phones    map[string]string
        Emails    map[string]string
        Note      string
        Photo     string
}

type Reader interface {
	Read() (Person, error)
	Close()
}

type Writer interface {
	Write(p Person)
	Close()
}
