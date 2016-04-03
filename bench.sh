#!/usr/bin/env bash

set -e

frameworks=(beego echo/standard echo/fasthttp go-gin goji martini)

for f in ${frameworks[@]}; do
  echo "benchmarking $f..."
  go build $f/server.go
  ./server &
  pid=$!
  sleep 2
  wrk -t2 -c20 -d20s http://localhost:8080/teams/x-men/members/batman
  echo "benchmarking with pipleline..."
  wrk -t2 -c20 -d20s http://localhost:8080 -s pipeline.lua /teams/x-men/members/batman 64
  echo "stopping $f ($pid)..."
  echo $pid, pid
  kill -9 $pid &> /dev/null
  sleep 2
  rm server
done
