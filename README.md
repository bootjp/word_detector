# word_detector

## that is this?

Fast search for words in huge log files.

## use case

When migrating a database, I wanted to determine whether all query logs included incompatible queries. 

## method

Here we expand the word on-memory and search in a single thread
If the log file is larger, it should be multi threaded.
Inverted indexes should be created when word files change frequently.
