package palisade

type Expression struct {
	// This order is extremely important
	Monadic   *Monadic  `( @@`
	Operator  *Operator `| @@`
	Dyadic    *Dyadic   `| @@`
	Morphemes *Morpheme `| @@ )`
}

type Monadic struct {
	Subexpr    *Expression `(("(" @@ ")")`
	Verb       *Ident      `| @@)`
	Expression *Expression `@@?`
}

type Dyadic struct {
	Monadic    *Monadic    `( @@`
	Morphemes  *Morpheme   `| @@ )`
	Subexpr    *Expression `(("(" @@ ")")`
	Verb       *Ident      `| @@)`
	Expression *Expression `@@?`
}

type Operator struct {
	Verb       *Ident      `@@`
	Operator   *string     `@Operator`
	Expression *Expression `@@`
}

type Morpheme struct {
	Char   *[]string  `@Char+`
	Alpha  *[]string  `| @Alpha+`
	Omega  *[]string  `| @Omega+`
	Real   *[]float64 `| @('-'? Float)+`
	Int    *[]int64   `| @('-'? Int)+`
	String *[]string  `| @String+`
}
