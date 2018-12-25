# Golang Locker

Utility `locker` - simple Golang locker based on map of mutex.

## Installation:

To install via `go get` (needs `golang` version 1.11+ installed):

```
go get github.com/enfipy/locker
```

## Usage:

To lock by key:

```
locker := locker.Initialize()
locker.Lock("key")
<code>
locker.Unlock("key")
```
