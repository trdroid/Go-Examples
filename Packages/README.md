### Package naming conventions

The package naming convention is to name the package after the directory containing it, 
so that a package name can easily be identified by the name of directory from which the package is being imported.

**Package main**

The package named *main* is treated specially by the *go* command which compiles it into a binary executable if it also finds a function called *main()*. 
Any executable program in Go must have the *main* package and the *main()* function. The *main()* function is the entry point of a program.

When the Go compiler encounters a package called *main*, it tries to find a function called *main()* to create a binary executable. 
The executable will be named after the directory that contains the *main* package.
If in case the Go compiler does not find the *main()* function, it does not create a binary executable.



