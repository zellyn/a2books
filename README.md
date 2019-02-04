# transcriber

[![No Maintenance Intended](http://unmaintained.tech/badge.svg)](http://unmaintained.tech/)

Transcriber is a personal side-project of mine. The immediate goal is
to re-transcribe Jim Sather's _Understanding the Apple II_ and
_Understanding the Apple IIe_ books, which have been immeasurably
useful to me and other emulator writers. But of course I'd like to
build a generally useful tool in the process, so the big goal is to
create a useful platform for transcribing any old technical book.

It being a side project, I'm also doing a bit of experimenting with
technologies that I wouldn't sensibly use at work. The backend is
fairly boring, written in Go, with XML flat files as the data format,
because diffs between versions will be readable as text, and the
coordination method for contributing changes might just be git pull
requests. The backend exists, although the schema covers only basic
bibliographic data right now.

The protocol between backend and frontend is GraphQL, because I'd like
to get a bit of practical experience with it, rather than just hearing
hype or reading enough to get a basic philosophy. The backend uses the
fantastic [gqlgen](https://gqlgen.com/) library to generate Go code.

The frontend is likely to be written in [Elm](https://elm-lang.org/)
(remember, there's no innovation budget on side projects!), except for
the part where you drag rectangles around on images to select parts of
the page for transcription. Elm is bad at that kind of thing, so the
intention is to encapsulate that logic into a Web Component. That will
live (initially) at
[github.com/zellyn/rekt](https://github.com/zellyn/rekt), and is in
fact the part that I'm playing with now: learning how to resize and
rotate rectangles in SVG.

There is no timeline for any of this. I'm in a new role at work, and
our second kid is due towards the end of March 2019, so we'll see :-)
