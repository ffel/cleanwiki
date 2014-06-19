# Clean Wiki

A case study to try [clean architecture](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)
on this [golang wiki tutorial](http://golang.org/doc/articles/wiki/).

# Next

-	~~Make `FindByTitle` more idiomatic go.  Let it return an error.~~

	A disadvantage of functions that return more than one value is
	that these cannot be used "in place".  Code becomes more verbose.

-	Basic *interface* and / or *infrastructure* implementation for
	reading and writing notes.

	It turns out that implementing the page repository "leaks"
	domain.Page into the infrastructure package.

	This contrasts a little bit my previous attempts to prevent
	leaking from `usecase.Add` and `usecase.Read`.

	On the other hand, we may have succeeded to prevent leaking
	in the repository test, which partly mimics the use of this
	code in infrastructure or main.

-	Proceed with the [tutorial](http://golang.org/doc/articles/wiki/#tmp_3).
