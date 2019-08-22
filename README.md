# word_detector

## that is this?

Fast search for words in huge log files.

## use case

When migrating a database, I wanted to determine whether all query logs included incompatible queries. 

## method

Here we expand the word into on-memory and search in a single thread.
If the log file is large, it should be multi threaded.
If a word file changes frequently, you should create a inverted index.


## how to use 

WORD_FILE=/path/to/word.txt TARGET_FILE=/path/to/log.txt go run cli/main.go
