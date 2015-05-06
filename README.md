# language-identifier
Golang program to identify unknown languages using multilingual lexicons

## Purpose
To use word lists (lexicons) in various languages to identify unknown languages in a text file.

## Status
Ready to use

## Installation
Use go install.

## Usage
Usage:

    language-identifier [-lex=DIRNAME] -text=FILENAME

Options:

    -lex=DIRNAME    Name of lexicon directory [default=value of LEX_DATA environment variable]
    -text=FILENAME  Name of file to check
language-identifier -text [filename] -lex [dirname]

The -lex argument specifies the name of the lexicon directory that contains the files. It is optional if the `LEX_DATA`
environment variable is set to the location of the data directory in the lexicon repository, for example:

export LEXICON_DATA=~/go/src/github.com/BluntSporks/lexicon/data

## Programming notes
See [abbreviation](https://www.github.com/BluntSporks/abbreviation) for a list of abbreviations used.

