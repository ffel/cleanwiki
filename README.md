# Clean Wiki

A case study to try [clean architecture](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/)
on this [golang wiki tutorial](http://golang.org/doc/articles/wiki/).

# Next

-	~~Make `FindByTitle` more idiomatic go.  Let it return an error.~~

	A disadvantage of functions that return more than one value is
	that these cannot be used "in place".  Code becomes more verbose.

-	Basic *interface* and / or *infrastructure* implementation for
	reading and writing notes.

-	Proceed with the [tutorial](http://golang.org/doc/articles/wiki/#tmp_3).
