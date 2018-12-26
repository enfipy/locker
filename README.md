# Lightweight Golang Locker

Utility `locker` - simple and lightweight Golang locker based on map of RWMutexes

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

To lock only read:

```
locker := locker.Initialize()
locker.RLock("key")
<code>
locker.RUnlock("key")
```

## Test:

To run tests:

```
go test
```

## Links:

[Good explanation](https://stackoverflow.com/a/19168242/10052381) of sync package and why is RWMutex with RLock are useful

[Reddit](https://www.reddit.com/r/golang/comments/a9j0we/enfipylocker_lightweight_named_locker_based_on/) post
