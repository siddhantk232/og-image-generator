
# og image generator

Generate open graph images for you awesome blogs.

> Disclaimer: This is created for my personal use but can be used in other projects
> with similar requirements.

## The Problem

I wanted a solution that will automatically create open graph images for the 
posts I write in [siddhant.codes](https://siddhant.codes). The goal was to take 
the yaml frontmatter from my posts, generate a png format image and store it in
a certain directory, preferably `./src/assets/og/{og_output}.png`.

## The Solution

I created this cli program that does only one job from requirements stated above. 
It creates a beautiful png image based on the args passed to it.

## Installation

The only way to get this program on your computer is by running:

```
go get github.com/siddhantk232/og-image-generator
```

> This should work on windows computer also but I haven't tested it.

## Usage

Run `og-image-generator` with no args and it will generate a sample png file
named `out.png` in the current directory.

This is intended to be used as a git hook. See the [`sample-git-hook.sh`](sample-git-hook.sh) 
for an example.

### Options

`-title`

default: "Sample Post Title"

description: "Title text"

`desc`

default: "Description of the post in about 25 words. This string could be your og description also!"

description: "Description text"

`date`

default: "30 April 2021"

description: "Date text"

`readtime`

default: "4 min read"

description: "Reading time text"

`out`

default: "out.png"

description: "Name of the output file"

`dpi`

default: 72

description: "screen resolution in dots per inch"

## TODO:

- add font size arg (with max limit)
- add font arg (currently uses Montsterrat only)
- add more background patterns
