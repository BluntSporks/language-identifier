# language-identifier
Golang program to identify unknown languages using multilingual lexicons

## Purpose
To use word lists (lexicons) in various languages to identify unknown languages in a text file.

## Status
Ready to use

## Installation
Use go install.

## Usage
language-identifier -text [filename] -lex [dirname]

The -text argument specifies the name of the file to check and is a required argument.

The -lex argument specifies the name of the lexicon directory that contains the files. It is optional if the `LEX_DATA`
environment variable is set to the location of the data directory in the lexicon repository, for example:

export LEXICON_DATA=~/go/src/github.com/BluntSporks/lexicon/data

## Short forms
See [short-names](https://www.github.com/BluntSporks/short-names) for a list of abbreviations used.

