# Go course

## Advantages of `go`

1. Code runs fast (why it runs fast)
2. Garbage collector, some other languages that run fast as Go does they don't have garbage collection
3. Simpler object, object orientation is a little simplified as compared with other languages
4. efficient concurrency implementation built into language

## Code runs fast

Three categories of languages:
1. machine language - low level language and it's directly executed on the CPU (simple, small test. e.g. add, move, read from registry etc..) 
2. assembly language - equivalent to machine language but easier to read. Map one-to-one the machine language (not completely but very close) equivalent, but you can read it because it's english. When you want to something very fast and really efficiently then you'll write it directly in assembly language.
3. high level language - everything else, they are much easier to use. Go is high level language!

> All software must be translated into the machine language of processor

If you've got a processor, i7 or whatever processor you're working with, that processor does not know C or Java or Python or Go or C++ or any of those, right?  
All it knows is its own machine language, say x86 machine language. So C, Python, Java or whatever it is, has to be translated. There is software translation step that has to go on:
1. **Compiled** - high level language to machine code happened one time before you execute the code. Translation occurs once, and when you run the code you are just running the machine code directly.
2. **interpreted** - instructions are translated while the code is excuted. So the translations occurs every time you run the Python code and so that slows you down.

Compiled code is generally faster to execute.
Interpreters make coding easier: manage memory automatically and infer variable types.

Go is a good compromise between compiled and interpreted languages, it is compiled language but it has some good features of interpreted languages.

## Garbage collection

Automatic memory management
- where should memory be allocated?
- when can memory be deallocated?

Manual memory management is hard
- Deallocated too early, false memory access
- Deallocated too late, wasted memory

Go has garbage collector, downside is that it slows  down exection a bit, but it is an efficient garbage collector.

## Objects

Go language is object-oriented -> weakly object-oriented. It implements objects but they have fewer features than you would see in another object-oriented language.

### Object-oriented

- Object-oriented programming is for code organization (encapsulation).
- Group code together data and function which are related.
- User-defined type which is specific to an application.
    - generic type -> `ints` have data (the number) and functions (*, +, - etc.)
    - e.g. 3d geomtric applcication -> `point(x, y, z) { DistToOrigin(), Quadrant() }`  
      difference between object and class
      
### Object in Go

- does not use `class` but use `structs`
- simplified implementation of classes
    - no inheritance
    - no constructors
    - no generics
    
This makes it easier to code and it makes it efficient to run. But if you like these features this is a disadvantage.

## Concurrency

There are built-in constructs in the language that make it easy to use concurrency.

A lot of motivation for concurrency comes from the need for speed, not all.  
How concurrency get around these performance limitations.

### Performance limitation of computers

- Moore's Law used to help performance  
  number of transistors on a chip doubles every 18 months.
- more transistors used to lead to higher clock frequencies  
  Programmers would get lazy, because they knew that pretty soon hardware people would double the number of transistors and fixed all their problems for them, but that si not happening anymore, because Moore's Law had to slow down.
- Power/temperature constraints limit clock frequency now  
  So you can't keep increasing the clock rates.

How do you get performance improvement even though you can't just crank up the clock?

### Parallelism

- Increase number of cores
- you can perform multiple tasks at the same time
    - it does not necessarily improve your latency but your throughput will improve (potentially) 🤞
- Difficult with parallelism
    - when do tasks start/stop?
    - what if one task needs data from another task? How does this data transfer occur?
    - Do tasks conflicts in memory?  
    You do not want one task to write to its variable A and that overrides vairable B of another task.
    Code that can be executed in parallel can be difficult.

### Concurrency programming

- Concurrency is management of multiple tasks at the same time  
  They might not actually be executed at the same time (e.g. single core processor) but they are alive at the same time. So they could be executed at the same time if you had the resource.
- Key requirement for large systems  
  You want to be able to consider 20 things at one time.
- Concurrent programming enables parallelism
    - Management of task execution
    - Communication between tasks
    - Synchronization between tasks  
    So You can not just take a piece of regular code and say oaky I'm going to run it on 5 cores, that won't work. The programmer has to decide hot to partition this code (I want this data here and this data there etc.).  
  The program is making these decisions that allow these things to run in parallel.

### Concurrency in go

- Go includes concurrency primitives
- **Goroutines** represent concurrent tasks
- **Channels** are used to communicate between tasks
- **Select** enables task synchronization
- Concurrency primitives are efficient and easy to use

## Install Go

```brew install go```

## Workspaces and packages

### Workspace

Workspace is a directory where your Go stuff will go.

- hierarchy of directories
- Common organization is good for sharing  
  You never programming alone, you can work with people all over the place, and they have to be able to work with your code (merge, link, review, edit). So it is nice have a common standardized organization of your files.
- Three subdirectories
    1. src - Contains source code files
    2. pkg - Contains packages (libraries)
      Packages that you are going to link in that you need
    3. bin - Contains executables
      your compiled executables
- Directory hierarchy is recommended, not enforced.
- Workspace directory defined by **GOPATH** environment variable.

### Packages

- Group of related files
- Each package can be imported by other packages
- The first line of file names the package  
  `package mypackage`
  ```go
    import (
        mypackage
    )
  ```
- There must be one package called `main`  
  That's where the execution starts.
- Building the main package generate an executable program  
  when you build a non main package does not make it executable
- Main package needs have a function called `main()`
- `main()` is where code execution starts

## Go tool

### import

- `import` keyword is used to access other packages
- Go standard library includes many packages  
  i.e. "fmt" - printf statement
- Searches directories spefified by `GOROOT` and `GOPATH`
  if you import packages from some other place you have to change `GOROOT` and `GOPATH` so that it can find them
  
### Go Tool (CLI)

- Tool to manage Go source code
- Several commands
- `go build` - compiles the program  
  arguments can be a list of packages ora a list of `.go` files
  creates an executable for the main package, same name as the first `.go` file
- `go doc` - prints documentation for a package
- `go fmt` - format source code files
- `go get` - download packages and install them
- `go list` - list all the install packages
- `go run` - compiles go files and runs the executable
- `go test` - runs tests using files ending in `_test.go`
 
## Variables

- names must start with a letter
- any number of letters, digits or underscores
- case sensitive
- don't use keywords (`if`, `case`, `package` etc..)
- data stored in memory
- must have a **name** and a **type**.
  `var x int`
- can declare many on the same line
  `var x, y int`
- type defined the values a variable may take and operation that can be performed on it
- How much space in memory you are going to need to allocate for this variable 
  - integer
  - floating point
  integer division is different from floating point division (depends on hardware) 
  - Strings

- You can define an alias, an alternative name for a type, some time is useful for improve clarity.  
e.g. `type Celsius float64` or `type IDnum int`
- you can declare variables using the type alias  
e.g. `var temp Celsius` or `var pid IDnum`

### Variable initialization

- initialize in declaration `var x int = 100` or `var x = 100` - infer the type
- initialize after declaration `var x int` and then `x = 100`
- uninitialized variables have a zero value `var x int // x = 0` or `var x string // x = ""`
- Short variable declaration

`x := 100`

variable is declared as the type of expression on the right hand side  
Can only do this inside a function.


