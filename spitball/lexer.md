# Getting started with the lexer

The first iteration of the lexer will be designed to lex the following subset of the Rosso programming language:

```
let five = 5;
let ten = 10;
let add = fn(x, y) {
    x + y;
};
let result = add(five, ten);
```