### Goroutines

Goroutines are functions that can run concurrently with other goroutines. 

**Where are goroutines run**

The Go runtime has a set of configured logical processors where each logical processor is bound to a single OS thread. 
The Go runtime automatically schedules the execution of goroutines against these logical processors. 
A logical processor/OS thread can run more than one goroutine concurrently.

*In comparison with threads*

* More than one goroutine can be executed on a thread
* Goroutines consume less memory than threads and because of minimal overhead, it is perfectly fine to spawn tens of thousand of them in an application

*In comparison with Java's main thread*

* The main thread in Java is designated to run the *main* method which is the entry point of the program. In go, the entry point of the program is run in a goroutine.
* The main thread in Java can run concurrently with other threads and similarly the "main" goroutine can run concurrently with other goroutines.
