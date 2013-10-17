gonrepl
=======

An nrepl client written in go

gonrepl has two modes: 

gonrepl host:port code

``` bash
$ go install github.com/mattyw/gonrepl
$ gonrepl localhost:nreplport "(+ 1 2)"
$ gonrepl localhost: nreplport "(map (fn [x] (* x 2) [1 2 3 4 5])"
class clojure.lang.LispReader$ReaderException
$ gonrepl localhost: nreplport "(map (fn [x] (* x 2)) [1 2 3 4 5])"
(2 4 6 8 10)
```

or taking the code from stdin

``` bash
| gonrepl host:port
```

This second mode makes it ideal for using with acme
