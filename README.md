# pwmake

Yet Another Password Generator. This one has two selling points:
It tries to be like another, existing (and, most probably, superior)
password generator (pwgen). And, it is implemented in Go. Point one
isn't that exciting, since it basically spells "So Why Did You Do
It Again?", so it all comes down to point two: It is implemented
in Go, and it is a vehicle to learn the language to me.


## What would you need?

Probably a Go compiler. On an Ubuntu system, you can install one
using

	sudo apt-get install golang

This should work on any modern Debian-derived systems. If not, go to
http://golang.org/ and download a Go distribution for your system.

For the time being, you should have a working make or rake installed
to start building, but since Go comes with its own build system, you
could as well figure out how the provided Makefile fires up the build
and build using the Go-provided wrapper. Or write a script which does
the same.


## What is that application?

The application pwmake generates passwords. Nothing fancy right now,
but maybe we throw in fancy stuff later. Basically, it creates a
series of random numbers, maps them to a set of symbols (really just
"characters" for now), strings them together and prints them out.

However, there's already one thing to learn here: The ultimative best
passwords are generated in the most simply way. Anything which puts
in structure (to remember them better) will weaken the password.
Just sayin'.


## What IDE should you use?

Well, lots of people would say [vim](http://vim.org). However, one
very good alternative to using plain `vim` is
[IDEA](https://www.jetbrains.com/idea/), which is nice, friendly,
powerful, and comes with a
[vim plugin](https://plugins.jetbrains.com/plugin/164). Even better,
there is a nice
[Go plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin)
which makes life better for those who occasionally dip into Go.


## What's with the Makefile and Rakefile?

These are purely decorative. The Go community feels they are done with
`make` and surely don't like `rake`. However, I like the idea of
building a statically linked version of this project with five key
strokes, so it's still in here.
