#!/bin/sh -e

usage() {
  echo 'Usage: <command> [...]

Commands:

    gmh,    AlpacaVM    -- GrassMudHorse interpreter
    gmhd,   JOTCompiler -- GrassMudHorse disassembler
    ws2gmh, WS2GMH      -- Whitespace-to-GrassMudHorse converter
    gmherl, ErlangVM    -- GrassMudHorse interpreter in Erlang' >&2

  exit 2
}

if [ $# = 0 ]; then
  usage
fi

command="$1"
shift
case "$command" in
gmh|AlpacaVM)
  exec java -cp bin cn.icybear.GrassMudHorse.AlpacaVM "$@"
  ;;
gmhd|JOTCompiler)
  exec java -cp bin cn.icybear.GrassMudHorse.JOTCompiler "$@"
  ;;
ws2gmh|WS2GMH)
  exec java -cp bin cn.icybear.GrassMudHorse.WS2GMH "$@"
  ;;
gmherl|ErlangVM)
  if [ $# != 1 ]; then
    echo 'Usage: grass_mud_horse <gmh_file>' >&2
    exit 2
  fi
  exec erl -noshell -s grass_mud_horse compile_run "$1" -s init stop
  ;;
*)
  usage
  ;;
esac
