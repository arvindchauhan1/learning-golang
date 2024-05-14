### 🚀 Welcome to My Go Learning Repository! 🎉

Hey there, fellow Gopher! Welcome to my journey of learning Go. This repository is my personal space where I dive into the depths of Go, exploring its features, quirks, and best practices.

### 📚 What's Inside?

This repository is all about my learning process. Here's what you can expect to find:

- **Code Snippets**: I'll be experimenting with various Go features, writing code snippets to understand how they work.
- **Projects**: As I progress, I'll embark on small projects to apply what I've learned and showcase my skills.
- **Learning Notes**: I'll jot down my thoughts, discoveries, and insights as I navigate through the Go ecosystem.

### 🌟 Why Go?

Go is an awesome programming language known for its simplicity, efficiency, and concurrency support. I'm excited to explore its clean syntax, powerful standard library, and vibrant community.

## Go Variables

<details open>
<summary><strong>Variables in G(click here to expand)</strong></summary>

#### Syntax:

To declare a variable in Go, you use the `var` keyword followed by the variable name and its data type.

```go
var age int // Declaration with int type
```

To initialize a variable with a value, you can do it in the same line using the assignment operator `=`.

```go
var age int = 30 // Declaration and initialization
```

Alternatively, you can use type inference and let Go determine the type based on the value provided:

```go
var age = 30 // Type inference
```

In Go, To use shorthand notation, you use the `:=` operator to declare and initialize a variable in one step.

```go
name := "Alice" // Shorthand declaration and initialization
```

Just like regular variable declarations, Go will infer the data type based on the value provided.

```go
age := 25 // Shorthand declaration and initialization
```

#### Data Types:

Go has several basic data types, including:

- **int**: Integer type (e.g., 1, 42, -10)
- **float64**: Floating-point type (e.g., 3.14, -0.001)
- **string**: Textual data type (e.g., "hello", "world")
- **bool**: Boolean type (e.g., true, false)

#### Public and Private:

In Go, variable names starting with a capital letter are exported and can be accessed from other packages. These are considered public variables.

```go
var Name string = "John" // Public variable
```

Variables starting with a lowercase letter are scoped to the package they are defined in and cannot be accessed from outside. These are considered private variables.

```go
var age int = 30 // Private variable
```

That's all for now! Have fun coding with Go! 🚀🔥

</details>


### 📞 Let's Connect!

If you're also learning Go or just want to chat about programming, hit me up on [LinkedIn](https://www.linkedin.com/in/arvindchauhan1/). Let's learn and grow together!

Happy coding! 🐹🎉
