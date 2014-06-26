# Clean Wiki

A case study to try [clean
architecture](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)
on this [golang wiki tutorial](http://golang.org/doc/articles/wiki/).

# Next

-   ~~Make `FindByTitle` more idiomatic go. Let it return an error.~~

    A disadvantage of functions that return more than one value is that
    these cannot be used "in place". Code becomes more verbose.

    > One of the *idiomatic discusions* is whether or not to use
    > pointers in `domain.Store()` and `domain.FindByTitle()`.

    > For now, I think it is best *not* to use pointers unless there is
    > a good reason. Possible good reasons are

    > -   ability to modify the passed value (pass-by-reference)

    > -   too heavy too pass (expensive pass-by-value)

    > -   implement an interface

    > This is confirmed by the [go style
    > guide](https://code.google.com/p/go-wiki/wiki/CodeReviewComments#Pass_Values)

    > [Effective
    > go](http://golang.org/doc/effective_go.html#pointers_vs_values)
    > even suggests to implicitly implements some basic go interfaces
    > such as `Reader` and `Writer`.

-   Basic *interface* and / or *infrastructure* implementation for
    reading and writing notes.

    It turns out that implementing the page repository "leaks"
    domain.Page into the infrastructure package.

    This contrasts a little bit my previous attempts to prevent leaking
    from `usecase.Add` and `usecase.Read`.

    On the other hand, we may have succeeded to prevent leaking in the
    repository test, which partly mimics the use of this code in
    infrastructure or main.

-   Proceed with the
    [tutorial](http://golang.org/doc/articles/wiki/#tmp_3).

# Notes

Some notes from the clean architecture text (as well as some other
remarks).

-   Possible customer go in "domain", users go in "use cases"

-   MVC Controllers will end up in "interfaces"

-   Distinction between `interfaces` and `infrastructure` is a bit
    vague. According to the clean architecture text:

    > "The definition that makes sense for me is that both contain code
    > that interacts with the outer world, like code that talks to a
    > database, but while the code in interfaces is specific to your
    > program at hand, infrastructure code is not and can be used in
    > completely different applications."

    In the example, `interfaces` defines a `Row` interface and states
    the required sql commands, whereas `infrastructure` configures and
    uses a sqlite database and implements `Row`

    This reminds me of another observation. If we're able to define
    interfaces in terms of `Reader` and `Writer` in `interfaces`, we can
    use these in `infrastructure` where we chain these into a link wit
    zippers, encrypters and so forth.

    > The clean architecture `interface` package defines a struct that
    > on the one hand implements the repo interfaces from `domain` (and
    > customer from `usecases`) and on the other and uses the
    > `interface.DbHandler` and `interface.Row` interfaces which are
    > implemented in the `infrastructure` package.

    > I think in a similar fashion we can define a struct that
    > implements the `domain.PageRepository` interface and uses a
    > `ReadWriter` interface that is implemented or set (determined) in
    > the `infrastructure` package

-   I am still a little bit confused about whether or not to use the
    term `wiki` in `usecases`.

    Because we in fact have use cases such as "add a page to the wiki"
    and "read a page from the wiki" the use of wiki is justified.

    I think due to "requirement drift" (a wiki is *one way* to reach the
    notes) I think the term wiki has been exchanged for "notes".

    In light of the objective to re-implement the golang wiki, the term
    "wiki" is fair enough.
