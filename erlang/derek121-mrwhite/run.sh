#!/bin/sh -e

exec erl -pa _build/default/lib/mrwhite/ebin -noshell \
  -s mrwhite main "$@" \
  -s erlang halt
