## TODO

Read/Revise more about the bufio package in Go. My knowledge is so Rusty it's controversial (see, I have a sense of humour)

- Let's parse Additive expressions:

5 + 6

- One thing to note is that to take order of precedence into account, you just push it down the recursive tree, so a multiplicative term is parsed inside, and thus before, an additive term.
    - Yay, this is done. Now, look at it this way:
    
    5 + 6 * 4

    - If I want the "6 * 4" part to be parsed before the 5 + 6 part, I need to convince the compiler that the expression is actually a sum of 2 exprs -> 5 and (6 * 4). Thus, 6 * 4 must be treated as an expression that gets parsed as the left or right side branch of the expr trees while parsing an additive binary expr.
    - Since parenthesis, or brackets as we call them back home, have a really high precedence, they have to be parsed on the same level as primary exprs. SO, they must be parsed inside multiplicative exprs. In other words, you start parsing an expr, which parses additive, which need a left and right. The left and right nodes in turn parse multiplicative exprs, which in turn parse primary ones including parenthesis.

## Parsing "let" statements

As it currently stands, let statements in Rosso are of the form: 

```
let y = 10;
let foobar = add(5, 6);
```

A VarDecl struct really only needs four things - 
1. The symbol of the identifier that is being declared
  - A question might arise - why not store an `*ast.Ident` type? 
  - okay so what is happening is, when parsing the tokens into an AST, the parser encounters this statement, and so 
    it start parsing this into a node. How will this node looks like?
  - VarDecl node -> Kind: VarDeclNode
                 -> IsConstant: bool
                 ->   