package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	// todos.Add("Learn Go")
	// todos.Add("Learn Python")
	// todos.Add("Building cli todo")
	// todos.Toggle(1)
	// todos.Toggle(2)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	// todos.Print()
	storage.Save(todos)
}
