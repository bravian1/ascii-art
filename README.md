

# ASCII Art

## Overview
The ASCII Art is a Go program designed to convert strings into graphical representations using ASCII characters. It's capable of handling input with numbers, letters, spaces, special characters, and newline (`\n`) characters.

## Installation
Clone the repository and navigate to the project directory:

```bash
git clone https://learn.zone01kisumu.ke/git/bnyatoro/ascii-art
cd ascii-art
```

Run the program as shown in the usage section.

## Features
- Converts strings into ASCII art representations.
- Supports multiple banner formats including `shadow`, `standard`, and `thinkertoy`.
- Each ASCII character has a height of 8 lines.
- Characters are separated by a newline character.

## Limitations
- The program is designed to work with printable ASCII characters only.
- Characters outside the range of space (` `) to tilde (`~`) are not supported and may not render correctly.

## Prerequisites
- Go programming language
- ASCII banner files in the correct format

## Usage
To use the ASCII Art Generator, compile and run the program with a string argument. Here are some examples:

```bash
go run . "Hello" | cat -e
```
```console
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
```                              
```bash
go run . "\n" | cat -e
```
```console
$
```
```bash
go run . "HeLlO" | cat -e
```
```console
 _    _          _        _    ____   $
| |  | |        | |      | |  / __ \  $
| |__| |   ___  | |      | | | |  | | $
|  __  |  / _ \ | |      | | | |  | | $
| |  | | |  __/ | |____  | | | |__| | $
|_|  |_|  \___| |______| |_|  \____/  $
                                      $
                                      $
 ```                                     
```bash
go run . "Hello There" | cat -e
```
```console
 _    _          _   _               _______   _                           $
| |  | |        | | | |             |__   __| | |                          $
| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  $
|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ $
| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ $
|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| $
                                                                           $
                                                                           $

```



## Contributing
Contributions to the ASCII Art Generator are welcome. Please ensure that your code adheres to Go's best practices and includes unit tests for new features.

