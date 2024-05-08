package main

func helloHighOrder() func() string {
	return func() string {
		return "Hello"
	}
}

func helloHighOrderWithParam(name string) func() string {
	return func() string {
		return "hello " + name
	}
}

func helloHighOrderWithParam2(name string) func(greeting string) string {
	return func(greeting string) string {
		return greeting + " " + name
	}
}
