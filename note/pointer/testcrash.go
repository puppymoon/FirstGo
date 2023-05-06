package main

func main() {
	// 指針指向nil
	var p *int = nil
	// 不能將空指針賦值
	*p = 0
}

// in Windows: stops only with: <exit code="-1073741819" msg="process crashed"/>
// runtime error: invalid memory address or nil pointer dereference
