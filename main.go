package main

import (
	"fmt"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

// this works like stringer
func (i *MyError) Error() string {
	return fmt.Sprintf("at %v: %v\n", i.when, i.what)
}

// go has a buit in interface error same thing as fmt.Stringer to handle errors
// here we are returning error interface type so it will look for Error() method that returns string by default and use it's value 
func check(i int) error {
	if i < 10 {
		return &MyError{
			when: time.Now(),
			what: "Bro it's less than 10 not gonna work",
		}
	} else {
		return nil
	}
}

func big (i int){
	if err := check(i); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pass aai dee")
	}
}

func main() {
  big(5)
  big(10)
}
