# ktx2png - A simple tool that converts ktx files to png


A wrapper around PVRTexToolCLI to convert ktx to png. KTX is a format for storing textures for OpenGL and OpenGL ES applications.

## Installation
Build from src:
```
    git clone https://github.com/raboafo/ktx2png && cd ktx2png
    go build .
```

## Usage:

```terminal
In all situations, destination path == source path if -o option is omitted. Any existing file with the same name is overwritten.
// Convert texture.ktx, store as texture.png
ktx2png -i texture.ktx -o texture.png
ktx2png -i texture.ktx

// Convert all `.ktx` files in directory `path/to/ktx`, store output pngs in path/to/png. Subdirectories are ignored.
ktx2png -i path/to/ktx -o path/to/png

```

## TODO:
  * Add support for linux and macos.
