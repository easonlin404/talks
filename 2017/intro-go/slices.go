package main


func main() {
	
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:] 
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}
}
