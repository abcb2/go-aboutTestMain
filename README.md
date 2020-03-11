# go-sampleTest
TestMainを使った場合全部のテストを実行するが、`-run` オプションをつけることでテストケースの指定も可能

# 指定なし

```
$ go test -v -count=1 ./
----------START
=== RUN   TestStudent
=== RUN   TestStudent/SayHello
=== RUN   TestStudent/Bye
--- PASS: TestStudent (0.00s)
    --- PASS: TestStudent/SayHello (0.00s)
    --- PASS: TestStudent/Bye (0.00s)
PASS
----------FINISH
ok  	github.com/abcb2/sampleTest	0.069s
```

# 指定あり

```
$ go test -v -count=1 ./ -run TestStudent/Bye
----------START
=== RUN   TestStudent
=== RUN   TestStudent/Bye
--- PASS: TestStudent (0.00s)
    --- PASS: TestStudent/Bye (0.00s)
PASS
----------FINISH
ok  	github.com/abcb2/sampleTest	0.209s

$ go test -v -count=1 ./ -run TestStudent/SayHello
----------START
=== RUN   TestStudent
=== RUN   TestStudent/SayHello
--- PASS: TestStudent (0.00s)
    --- PASS: TestStudent/SayHello (0.00s)
PASS
----------FINISH
ok  	github.com/abcb2/sampleTest	0.067s
```

# subdirectory
フォルダ階層の認識してくれないのでtestディレクトリに1階層でまとめることにした。

`test/foo` 以下のテストは実行されない