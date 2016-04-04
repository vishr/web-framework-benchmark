#!/usr/bin/env bash

set -e

frameworks=(beego echo/standard echo/fasthttp gin goji martini)

for f in ${frameworks[@]}; do
  echo "benchmarking $f..."
  go build $f/server.go
  ./server &
  pid=$!
  sleep 2
  wrk -t2 -c20 -d10s http://localhost:8080/teams/x-men/members/wolverine
  echo "benchmarking with pipleline..."
  wrk -t2 -c20 -d10s http://localhost:8080 -s pipeline.lua /teams/x-men/members/wolverine 64
  echo "stopping $f ($pid)..."
  kill -9 $pid &> /dev/null
  sleep 2
  rm server
done
