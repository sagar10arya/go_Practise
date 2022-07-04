package main

import (
	"fmt"
)

func main() {

	// Formatting verbs for printf function
	/*	Verb	Description
		%v		Prints the value in the default format
		%#v		Prints the value in Go-syntax format
		%T		Prints the type of the value
		%%		Prints the % sign */

	var j = 10.78
	var name = "TaylorSwift"

	fmt.Printf("%v\n", j)
	fmt.Printf("%#v\n", j)
	fmt.Printf("%T\n", j)
	fmt.Printf("%v%%\n", j)

	fmt.Printf("%v\n", name)
	fmt.Printf("%#v\n", name)
	fmt.Printf("%T\n", name)
	fmt.Printf("\n")
	/* Integer Formatting Verbs
	The following verbs can be used with the integer data type:

	Verb	Description
	%b		Base 2
	%d		Base 10
	%+d		Base 10 and always show sign
	%o		Base 8
	%O		Base 8, with leading 0o
	%x		Base 16, lowercase
	%X		Base 16, uppercase
	%#x		Base 16, with leading 0x
	%4d		Pad with spaces (width 4, right justified)
	%-4d	Pad with spaces (width 4, left justified)
	%04d	Pad with zeroes (width 4) */

	var i = 15

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)

	/* String Formatting Verbs
	   The following verbs can be used with the string data type:

	   Verb	Description
	   %s		Prints the value as plain string
	   %q		Prints the value as a double-quoted string
	   %8s		Prints the value as plain string (width 8, right justified)
	   %-8s		Prints the value as plain string (width 8, left justified)
	   %x		Prints the value as hex dump of byte values
	   % x		Prints the value as hex dump with spaces*/
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	/* Float Formatting Verbs
	The following verbs can be used with the float data type:

	Verb	Description
	%e		Scientific notation with 'e' as exponent
	%f		Decimal point, no exponent
	%.2f	Default width, precision 2
	%6.2f	Width 6, precision 2
	%g		Exponent as needed, only necessary digits*/
	var k = 3.141
	fmt.Printf("\n")
	fmt.Printf("%e\n", k)
	fmt.Printf("%f\n", k)
	fmt.Printf("%.2f\n", k)
	fmt.Printf("%6.2f\n", k)
	fmt.Printf("%g\n", k)

}
