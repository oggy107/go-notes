package main

// Mutexes (mutual exclusion) allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.

// Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type and its two methods:

// .Lock()
// .Unlock()
// We can protect a block of code by surrounding it with a call to Lock and Unlock as shown on the protected() method below.

// It's good practice to structure the protected code within a function so that defer can be used to ensure that we never forget to unlock the mutex.

// func protected(){
//     mux.Lock()
//     defer mux.Unlock()
//     // the rest of the function is protected
//     // any other calls to `mux.Lock()` will block
// }
// Mutexes are powerful. Like most powerful things, they can also cause many bugs if used carelessly.

// RW Mutex (Read write mutex)
// The standard library also exposes a sync.RWMutex

// In addition to these methods:

// Lock()
// Unlock()
// The sync.RWMutex also has these methods:

// RLock()
// RUnlock()
// The sync.RWMutex can help with performance if we have a read-intensive process. Many goroutines can safely read from the map at the same time (multiple Rlock() calls can happen simultaneously). However, only one goroutine can hold a Lock() and all RLock()'s will also be excluded.

// so to clear things suppose we have two goroutines that are accesstion same map
// reading and reading => safe
// reading and writing => not safe
// writing and reading => not safe
// writing and writing => extremely unsafe
