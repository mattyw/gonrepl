gonrepl
=======

An nrepl client written in go

```
$ go install github.com/mattyw/gonrepl
$ gonrepl localhost:nreplport "(+ 1 2)"
$ gonrepl localhost: nreplport "(map (fn [x] (* x 2) [1 2 3 4 5])"
class clojure.lang.LispReader$ReaderException
$ gonrepl localhost: nreplport "(map (fn [x] (* x 2)) [1 2 3 4 5])"
(2 4 6 8 10)
```
