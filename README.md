# race
[![CircleCI](https://circleci.com/gh/xxorde/race.svg?style=svg)](https://circleci.com/gh/xxorde/race)
[![Build Status](https://travis-ci.org/xxorde/race.svg?branch=master)](https://travis-ci.org/xxorde/race)

Litle test to show: os/exec: data race between StdoutPipe and Wait

Code was found here: https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/

I am having the same issue here, but the code is more complex: https://github.com/xxorde/pgglaskugel/


## Try out

```bash
go run --race main.go
```

or

```
make
```

## Output from CircleCI
```bash
go run --race main.go
docker build out | hello world
==================
WARNING: DATA RACE
Write at 0x00c420010a00 by main goroutine:
  os.(*file).close()
      /usr/local/go/src/os/file_unix.go:143 +0x124
  os.(*File).Close()
      /usr/local/go/src/os/file_unix.go:132 +0x55
  os/exec.(*Cmd).closeDescriptors()
      /usr/local/go/src/os/exec/exec.go:262 +0x67
  os/exec.(*Cmd).Wait()
      /usr/local/go/src/os/exec/exec.go:447 +0x2bd
  main.main()
      /home/ubuntu/.go_project/src/github.com/xxorde/race/main.go:37 +0x257

Previous read at 0x00c420010a00 by goroutine 6:
  os.(*File).read()
      /usr/local/go/src/os/file_unix.go:228 +0x50
  os.(*File).Read()
      /usr/local/go/src/os/file.go:101 +0x6f
  bufio.(*Scanner).Scan()
      /usr/local/go/src/bufio/scan.go:208 +0x526
  main.main.func1()
      /home/ubuntu/.go_project/src/github.com/xxorde/race/main.go:26 +0x47

Goroutine 6 (finished) created at:
  main.main()
      /home/ubuntu/.go_project/src/github.com/xxorde/race/main.go:29 +0x21e
==================
Found 1 data race(s)
exit status 66

go run --race main.go returned exit code 1

Action failed: go run --race main.go
```
