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
