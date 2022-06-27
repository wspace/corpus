#!/bin/sh -e

usage() {
  echo 'Usage: <command> [...]

Commands:

    AlpacaVM,    gmh    -- GrassMudHorse interpreter
    JOTCompiler, gmhd   -- GrassMudHorse disassembler
    WS2GMH,      ws2gmh -- Whitespace-to-GrassMudHorse converter
    ErlangVM,    gmherl -- GrassMudHorse interpreter in Erlang' >&2

  exit 2
}

if [ $# = 0 ]; then
  usage
fi

command="$1"
shift
case "$command" in
AlpacaVM|gmh)
  exec java -cp bin cn.icybear.GrassMudHorse.AlpacaVM "$@"
  ;;
JOTCompiler|gmhd)
  exec java -cp bin cn.icybear.GrassMudHorse.JOTCompiler "$@"
  ;;
WS2GMH|ws2gmh)
  exec java -cp bin cn.icybear.GrassMudHorse.WS2GMH "$@"
  ;;
ErlangVM|gmherl)
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
