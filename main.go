package main

import (
	"fmt"
)


func main(){
	var a int;
	b:=2;
	var c,d,e = 1, 2.0,"hello";
	f:=010; /*octal no always starts with 0*/
	g:=0xa; /*hexadecimal*/
	fmt.Println("Hello-Go");
	fmt.Printf("a=%d type(a)=%T\n",a,a);
	fmt.Printf("b=%d type(b)=%T\n",b,b);
	fmt.Printf("c=%d type(c)=%T\n",c,c);
	fmt.Printf("d=%f type(d)=%T\n",d,d);
	fmt.Printf("e=%s type(e)=%T\n",e,e);
	fmt.Printf("f=%d type(f)=%T\n",f,f);
	fmt.Printf("g=%d type(g)=%T\n",g,g);
}
