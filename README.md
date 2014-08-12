Exploring scoping of *_test.go files
====================

Before playing with this code, I thought the following:

Code in foo_test.go acts normally *except* it is not compiled into the package unless you run "go test".  I had the idea that "go test" builds a binary that includes all code, including test files, then runs the init() functions and Test*** functions.

After playing with this code and [discussing it on golang-nuts](https://groups.google.com/d/msg/golang-nuts/Mt6fjVZzQ84/7kMcxMvjChEJ), it appears:

Each package gets compiled into test-build and non-test-build versions.  Tests are run on a package-by-package basis, where *only the test-build is used for the current package* and all imports use non-test-build versions.

This explains why:

* Exported package identifiers within *_test.go files aren't visible outside the package.

* Initialization functions "init()" within *_test.go files can change exported package variables but these changes are only visible within the test files.


