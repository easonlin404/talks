
func main() {
	var a Abser	
	v := Vertex{3, 4}
	
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs()) // will output: 5
}