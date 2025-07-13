---
title: Golang Intro
date: 2025-7-12
description: Learning the Go programming language
categories: ['go', 'programming', 'language', 'experiment']
draft: false # Change to true to not render the post in on the website
---

### Setup
1. Downloaded & installed Go from [go.dev](https://go.dev)
2. Installed VSCode extension
    - [<img src="https://github.com/RoarkCats/roarkcats.github.io/blob/main/posts/golang/govscode.png?raw=true" width=25%>](https://marketplace.visualstudio.com/items?itemName=golang.Go)
3. Created standard workspace folder path `../<user>/go/src/<projects>` for my modules
4. Created an example module `../example/` with a `main.go` file
    - Don't think the file *has* to be called 'main', but if it runs, there must be a file within the module that starts with `package main` on the first line (you also put that in `main.go` ofc)

### Guides & Docs
The docs for Go are good, but these helped me more personally

- [Learn Go in 12 Minutes](https://www.youtube.com/watch?v=C8LgvuEBraI)
- [Go by Example](https://gobyexample.com/)
- [Go Packages](https://pkg.go.dev)

### Notes
Go at first seems like a language mix of Py and JS, however, after coding in it for a bit, I might throw C into the mix as it does feel like a lower level language with how it makes you code ad think in certain ways. While it does contain a lot of modern/high level niceties, it's still quite a simple language, ex arrays pretty much only have `append`, `len`, and `range` (+`sort`/`reverse` for slices)

I found it interesting, and slightly bothersome, how (by default) the official VSCode extension formats your program every time you save. Pretty sure this can be configured

It may be noted that running/programming Go appears to make antivirus programs freak out each time its run/saved(?), never had any problems but it was kinda annoying, does make sense as Go is a compiled language that is often used by attackers apparently. Anyhow I added my Go folder to Norton's exception list to get rid of the hassle. (Yes yes norton bad xD)

I like how easy it is to compile Go into an exe/binary for portability and ease of use! Also its speed :>

You can view my first testing script [here](https://github.com/RoarkCats/roarkcats.github.io/blob/main/posts/golang/main.go) on GitHub