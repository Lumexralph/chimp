# Chimp

![chimp](chi.jpeg)

Chimp is a programming language created solely for the main reason of learning how to build an interpreter using Go.

I love the web and also love low-level stuff too, I want to know what's happening under the hood.

I owe so much of this inspiration to [Thorsten Ball](https://twitter.com/thorstenball). He gave me his book at almost give away and I went straight to learning.

Features will get added to chimp with time.

## Features of Chimp

Chimp looks like this

    // Bind values to names with let-statements
    let version = 1;
    let name = "Monkey programming language";
    let myArray = [1, 2, 3, 4, 5];
    let coolBooleanLiteral = true;

    // Use expressions to produce values
    let awesomeValue = (10 / 2) * 5 + 30;
    let arrayWithValues = [1 + 1, 2 * 2, 3];

Function literals and we can use them to bind a function to a name:

    // Define a `fibonacci` function
    let fibonacci = fn(x) {
    if (x == 0) {
        0                // Monkey supports implicit returning of values
    } else {
        if (x == 1) {
        return 1;      // ... and explicit return statements
        } else {
        fibonacci(x - 1) + fibonacci(x - 2); // Recursion! Yay!
        }
    }
    };

Data types are booleans, strings, hashes, integers and arrays. They can be combined.

    // Here is an array containing two hashes, that use strings as keys and integers
    // and strings as values
    let people = [{"name": "Anna", "age": 24}, {"name": "Bob", "age": 99}];

    // Getting elements out of the data types is also supported.
    // Here is how we can access array elements by using index expressions:
    fibonacci(myArray[4]);
    // => 5

    // We can also access hash elements with index expressions:
    let getName = fn(person) { person["name"]; };

    // And here we access array elements and call a function with the element as
    // argument:
    getName(people[0]); // => "Anna"
    getName(people[1]); // => "Bob"

functions are first-class citizens, they are treated like any other value. Thus we can use higher-order functions and pass functions around as values.

    // Define the higher-order function `map`, that calls the given function `f`
    // on each element in `arr` and returns an array of the produced values.
    let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
        accumulated
        } else {
        iter(rest(arr), push(accumulated, f(first(arr))));
        }
    };

    iter(arr, []);
    };

    // Now let's take the `people` array and the `getName` function from above and
    // use them with `map`.
    map(people, getName); // => ["Anna", "Bob"]

Closures: This is a cool feature I love from learning in the book

    // newGreeter returns a new function, that greets a `name` with the given
    // `greeting`.
    let newGreeter = fn(greeting) {
    // `puts` is a built-in function we add to the interpreter
    return fn(name) { puts(greeting + " " + name); }
    };

    // `hello` is a greeter function that says "Hello"
    let hello = newGreeter("Hello");

    // Calling it outputs the greeting:
    hello("dear, future Reader!"); // => Hello dear, future Reader!

## Major parts of the interpreter

* the lexer
* the parser
* the Abstract Syntax Tree (AST)
* the internal object system
* the evaluator
