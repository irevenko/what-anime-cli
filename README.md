# What Anime CLI ❓🖼

<p align="center">> This application is basically a 🍬 wrapper around
<a href="https://github.com/soruly/trace.moe">trace.moe </a></p>
<p align="center">PREVIEW</p>
<p align="center"><img src="./anime_images/record/demo.gif" width="700"></p>

# Installation 🔨

```bash
go install github.com/irevenko/what-anime-cli@latest
```

```bash
what-anime-cli file ani.png
```

[AUR package](https://aur.archlinux.org/packages/what-anime-cli-git)

# Usage 🖥

### Get Anime By Image File 🗃

`what-anime file anime.jpg`

### Get Anime By Image Link 🔗

`what-anime link https://anime.com/image.png` <br>

### Potential troubles

Just escpape the link with quotes
<b>"</b>https://anime.com.anime.png&...<b>"</b>

### Supported image extensions 🖼

I've tested these:

- jpg
- png
- jfif
- webp
- gif

If you have found other's working formats please create an issue

# Contributing 🤝

Contributions, issues and feature requests are welcome! 👍 <br> Feel free to
check [open issues](https://github.com/irevenko/what-anime-cli/issues).

# Quick Start 🚀

`git clone https://github.com/irevenko/what-anime-cli.git` <br>
`cd what-anime-cli` <br> `go get -d ./...` <br> `go run main.go` <br>

# What I Learned 🧠

- How to build CLI using Go
- Go project structure
- Go basics (modules, working with images, making HTTP requests)

# License 📑

(c) 2021 Ilya Revenko. [MIT License](https://tldrlegal.com/license/mit-license)
