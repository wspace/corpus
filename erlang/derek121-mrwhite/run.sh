#!/bin/sh -e

exec erl -pa . -noshell \
  -s mrwhite main "$@" \
  -s erlang halt
