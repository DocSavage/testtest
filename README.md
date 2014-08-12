Simple exploration of scoping of *_test.go files
====================

Before playing with this code, I thought the following:

Code in *_test.go acts normally *except* it is not compiled into the package unless you run "go test".
_
After playing with this code, it seems that:

* Exported package identifiers within *_test.go files aren't visible outside the package.

* Initialization functions "init()" within *_test.go files can change exported package variables but these changes are only visible within the test files.


