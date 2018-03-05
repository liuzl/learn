### Example CFG file
```
arithmetic <: super {
  E = "one" {1} | "two" {2} | "three" {3} | "four" {4}
  E = "minus" E {nf.math.neg($2)}
  E = E "plus" E {nf.math.sum($1, $3)}
  E = E "minus" E {nf.math.sum($1, nf.math.neg($3))}
  E = E "times" E {nf.math.mul($1, $3)}
}
```
### BuiltInRules
```
BuitInRules {
  alnum = letter | digit
  digit = "0".."9"
  letter = "a".."z" | "A".."Z"
  ListOf<elem, sep> = elem (sep elem)+
}
```
### Grammar for CFG
```
CFG {
  Grammar = ident SuperGrammar? "{" Rule* "}"
  SuperGrammar = "<:" ident
  ident = identFirst identRest*
  identFirst = "_" | letter
  identRest = "_" | alnum

  Rule
    = ident Formals? ruleDescr? "="  RuleBody -- define
    | ident Formals?            ":=" RuleBody -- override
    | ident Formals?            "+=" RuleBody -- extend

  Formals = "<" ListOf<ident, ","> ">"

  RuleBody = ListOf<TopLevelTerm, "|">

  TopLevelTerm = Seq caseName -- inline | Seq
  Seq = Iter*
  Iter
    = Base "*" -- star
    | Base "+" -- plus
    | Base "?" -- opt
    | Base

  Base
    = ident Params
}
```
