create a makefile as follows:

- in directory web/sqirvy-query, create a makefile that will build the web app and the server
- the makefile should have the following targets:

  - build - build the web app and the server. place the built files in directory web/sqirvy-query/build
  - run - run the server
  - clean - remove the built files

- in directory cmd/sqirvy-query, create a makefile that will build the command line app
- the makefile should have the following targets:

  - build - build the command line app. place the built binary in directory cmd//build
  - test - use the cmd go test .
  - clean - remove the built binary

- in directory cmd/gemini, create a makefile that will build the command line app

  - the makefile should have the following targets:
    - build - build the command line app. place the built binary in directory cmd/build
    - test - use the cmd go test .
    - clean - remove the built binary

- in directory cmd/openai, create a makefile that will build the command line app

  - the makefile should have the following targets:
    - build - build the command line app. place the built binary in directory cmd/build
    - test - use the cmd go test .
    - clean - remove the built binary

- in directory cmd/anthropic, create a makefile that will build the command line app

  - the makefile should have the following targets:
    - build - build the command line app. place the built binary in directory cmd/build
    - test - use the cmd go test .
    - clean - remove the built binary

- in directory pkg/api, create a makefile that will run tests

  - create an empty 'build' make rule
  - test: run go test . in the pkg/api directory
  - create an empty 'clean' make rule

- in directory cmd, create a makefile that invokes the lower level makefiles
  - build, test and clean rules for each lower level makefile
- in directory web, create a makefile that invokes the lower level makefiles
  - build, test and clean rules for each lower level makefile
- in the top level directory, create a makefile that invokes the lower level makefiles
  - build, test and clean rules for each lower level makefile
