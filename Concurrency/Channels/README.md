### Channels

Concurrent modification of data(global variables or shared memory) by processes, threads or goroutines running concurrently is a huge problem.
Unsynchronized modifications to shared data can be prevented by using locks to enable synchronized access/changes.

Channels enforce a pattern to provide safety for shared data from concurrent modification by ensuring

* data modification by ONLY one goroutine at any point in time.

*Caveat*

Channels provide data protection ONLY if copies of original data are exchanged between goroutines through the channels.
In such cases, the goroutines have their own copy of data and any modifications done to the data affects only the local copy. 
If in cases, pointers to data are exchanged between goroutines through the channels, the channels do not offer any data protection, 
instead access to shared data should be synchronized across goroutines.

Goroutines and channels can be used in scenarios where different processes have to access/modify data sequentially.
