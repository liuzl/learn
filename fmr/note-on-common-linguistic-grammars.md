### Copied from http://libzx.so/note/2018/08/10/note-on-common-linguistic-grammars.html

## Formal Language Aspect of Grammar

From the aspect of formal languages, we know there are mainly four types of grammar.

* Type-0: Unrestricted
* Type-1: Context-Sensitive Grammar
* Type-2: Context-Free Grammar
* Type-3: Regular Grammar

The larger the number, the stricter and less expressive the grammar. By less expressive we means the less language could be generated from that grammar.

It is common to augment a CFG with some extensions, making it more powerful and expressive. Some examples are as follows,

* Generalized CFG
* Linear Context-Free Rewriting Grammar, which as its name suggested is a subclass of GCFG
* Head Grammar is less powerful LCFRG
* Indexed Grammar is another kind of CFG
* Linear Indexed Grammar is a specialized Indexed Grammar

## Linguistic View of Grammar

On the other linguistic side, people are also talking about grammars, which try to explain how human languages are generated and understood. Though they could be classified as some of the four types above, it is usually not the key point to proclaim or prove that. There are some philosophical debate and different approaches for grammars in natural language.

There are 2 main approaches to describe grammar or word relations in human language, namely the **consituency relation** and the **dependency relation**. There are also other important approaches like cognitive grammar, functional grammar and stochastic grammar, which will not be covered in this post.

The ***Consituency Relation*** side mainly focus on syntactic structures like phrases, e.g. noun phrases, verb phrases, extracted from sentences.

* Phrase Structure Grammar
* General Phrase Structure Grammar
* Head-driven Phrase Structure Grammar(HPSG)
* Lexical Functional Grammar(LFG)
* Minimalist Programm, which favors bare phrase structure and abandoned X-bar theory
* Nanosyntax
* Arcpair Grammar
* Categorial Grammar, and the famous Compositional Categorical Grammar(CCG)

Note the last two are not directly originated from Chomsky’s theory.

And the *English Resources Grammar*(ERG) is a hand-crafted, linguistic-motivated HPSG for English only, and produces English Resource Semantics(ERS). But this lies on the semantic topic, which will not be covered by the post.

And here is the ***Dependency Relation*** side. Rather than composing words into phrases, this approach emphasize relations between any two words in a sentence.

* Recursive Categorical Syntax(algebraic syntax)
* Functional Generative Description
* Lexicase
* Link Grammar
* Meaning-text Theory
* Operator Grammar
* Word Grammar

There’re common grammars not belonging to any of the two approaches. For example,

* Tree Adjoining Grammar(TAG)
* Lexicalized Tree Adjoining Grammar(LTAG), a TAG armed with lexical surface forms

Although it is possible to classify these grammar to Type-0/1/2/3, that’s not critical questions thus in literature they are usually discussed in parallel to the grammar of formal language.

## Equivalence between Grammars

Since there’re so much grammars so far, it’s natural to compare any two of them. For example, how powerful is some one compared to another, i.e. the equivalence between two grammars.

>Two grammar are weakly equivalent, iff they can generate the same language string. 
Two grammar are strongly equivalent, iff they can generate not only the same language string but also the derived tree.

It has been proven that these grammar are weakly equivalent: *Head Grammar*, *Linear Indexed Grammar*, *Compositional Categorical Grammar*(CCG), and *Tree Adjoining Grammar*(TAG). And *Lexicalized Tree Adjoining Grammar*(LTAG) is strongly equivalent to the *Head-driver Phrase Structure Grammar*(HPSG).

Then we can conclude the expressive power of them, from less expressive to stronger:

* Regular Grammar
* CFG
* Head Grammar, Linear Indexed Grammar, CCG, TAG, LTAG, HPSG
* LCFRS
* GCFG
* CSG
* Unrestricted
