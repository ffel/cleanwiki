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

    > This is confirmed by the [go style guide](https://code.google.com/p/go-wiki/wiki/CodeReviewComments#Pass_Values)

    > [Effective go](http://golang.org/doc/effective_go.html#pointers_vs_values)
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
