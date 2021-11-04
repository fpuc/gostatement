package a

func foo() {
	go func() { // want "go statement found"

	}()

	go bar() // want "go statement found"
}

func bar() {
}
