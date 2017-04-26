# World counter

wordcounter - word counter command line utilitary

## Setup

If you have go (golang.org) on your system:
```
    $ go get github.com/ademilly/wordcounter
```

Else... yeah no, I'm not bothering serving up binaries for every arch :D

## Usage

Output a word count file from file content provided through absolute path via command line argument or through stdin using the `-` placeholder.
If given multiple files, will return the aggregated count of words.

```
    $ wordcount /path/to/some/file
    $ cat /path/to/some/file | wordcount -
    $ wordcount /path/to/some/file /path/to/some/other/file
```
