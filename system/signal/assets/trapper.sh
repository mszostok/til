#!/usr/bin/env bash

_term() {
  echo "child: received signal, ignoring it for 30 sec"
  sleep 30
}

trap _term SIGTERM SIGINT

main () {
	echo "Starting child process - sleep 60s"
	sleep 60
}

main
