/*

                    mmmmmmm   mm   mmmmmm mmmmmmm     m
                       #      ##   #      #      "m m"
                       #     #  #  #mmmmm #mmmmm  "#"
                       #     #mm#  #      #        #
                       #    #    # #      #        #

 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) James Tarver, 2019

project: taffy - compile taffy code to C.  C code will be ugly but efficient. Taffy code will be ugly but terse.
taffy = Terse AF, freakin' yeah!
keywords:
?  if
|  else
|? else if
@  loop init; cond; post
?? switch
fall
brk
cont
ret
use = import
pkg

types are inferred or specified with:
$(a, u) string
(+,-)#(8,16,32,64,u) integer //default per arch
(+,-)#. (32,64,e)float // default 64
user defined types are defined with
type name
	field1,f2,f3 = v1, v2, v3
	method(p, ...) ret  //not exported
	Method() ret,ret,.. //exported
	_method or _field //private to type.
	add %r %w %rw to have getters and setters auto generated or defined by user.
	To use getter/setter use x = Field or Field = x.  Will be translated to x = Get_Field() or
	Set_Field(x).  User defined get/set may have the same characteristics as any other function.

identifiers with () are function definitions if name is not defined previously, otherwise call
f(p)$
	locals
	block
	ret val

ids with [const int] are fixed arrays
ids with [] are slices
ids with (v...v) are ranges
ids with {} are maps
ids with  <T> are parameterized types/functions. Emitted code will create a new name as Id_T_Type
interfaces are referred to inside {prints}


packages: scan, state, parse, ast, gen
states of lexer:
in name - func scans till end of name.
if not defined then
	statefn = inDecl<Type>
else state = in_Expr

in func decl
*/
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)
	
func lex() {
	var src = `// test
				pkg test;
				use testpkg;
				i = 2;
				f(p1$, p2#, p3#., ...)
					? i == 2
						@ 0..9
							print(i)
							pronto()
							xfile=pronto1(i)
				exit(0)
				`
	var s scanner.Scanner

	s.Init(strings.NewReader(src))
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Println(s.TokenText())
		// tok = s.Scan()
		// fmt.Println(s.TokenText())
		// tok = s.Scan()
		// fmt.Println(s.TokenText())
	}
}
