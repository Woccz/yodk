# NOLOL-stdlib

This is the documentation for the NOLOL standard library. All macros and definitions listed here can simply be used by just inlcuding the right "std/*" file inside your own nolol file.

For example:

[stdlib_demo.yolol](generated/code/nolol/stdlib_demo.nolol ':include')

**As every part of NOLOL, the standard-library is still subject to change and may be changed in backwards-incompatible ways any time!**



# std/logic

This file contains basic definitions and macros for logic-operations  
Import using ' include "std/logic" '  



## Macros
 
### **logic_continue_line**
```
logic_continue_line(condition)<ignore> line
```
If condition is 0, produces a runtime-error that will skip the remaining line  
Usage: "logic_continue_line(var); do=1; stuff=2 $"  
The $ is important to not skip too much  


 
### **logic_ternary**
```
logic_ternary(condition, a, b) expr
```
Returns a if condition is true, otherwise b  
condition, a and b must be numbers. condition must be 0 or 1  


 
### **logic_wait**
```
logic_wait(condition) line
```
Blocks as long as condition is true  


 
### **logic_xor**
```
logic_xor(a, b) expr
```
Returns 1 if a or b is true, but not both  
a and b must be 0 or 1  





# std/math

This file contains basic definitions and macros for math  
Import using ' include "std/math" '  


## Definitions
 
### **math_e**
The mathematical constant e  

 
### **math_pi**
The mathematical constant pi  



## Macros
 
### **math_abs_basic**
```
math_abs_basic(x) expr
```
Returns the absolute value of x.  
Works on a basic-chip  


 
### **math_floor**
```
math_floor(x) expr
```
Returns the next lower integer to x  
Requires at least an advanced chip  


 
### **math_floor_basic**
```
math_floor_basic(x) expr
```
Returns the next lower integer to x  


 
### **math_sign**
```
math_sign(x) expr
```
If x is >0, returns 1 if <0 return -1, otherwise returns 0  


 
### **math_xor**
```
math_xor(a, b, out) block
```
Returns the bitwise xor of a and b in out  
Requires at least an advanced chip  





# std/string

This file contains basic definitions and macros for string-manipulation  
Import using ' include "std/string" '  



## Macros
 
### **string_contains**
```
string_contains(str, x) expr
```
Returns 1 if str contains x  


 
### **string_len**
```
string_len(str, out) line
```
Adds the lenght of str to out  
str is set to "" in the process  


 
### **string_pop**
```
string_pop(str, out)<ignore> line
```
Removes the last character from str and places it into out  


 
### **string_reverse**
```
string_reverse(str, out)<ignore> line
```
Appends the reverse of str to out.  
str is set to "" in the process  





