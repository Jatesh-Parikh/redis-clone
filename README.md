# Go Redis Clone

Built a Redis-inspired in-memory key-value server in Go using TCP networking and the RESP protocol. Implemented concurrent client handling with goroutines and channels, Redis-compatible GET/SET, CLIENT, and HELLO commands, and a thread-safe key-value store using sync.RWMutex.

> This project is based on a tutorial by [Anthony GG](https://www.youtube.com/watch?v=LMrxfWB6sbQ&list=PLHsjm_W8kcWZLPDxUplr8yredk95F-KIx).