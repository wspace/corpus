#!/bin/sh

usage() {
  echo 'Usage: <command> [...]

Commands:

    whitespace -- Whitespace interpreter
    mkws       -- Whitespace assembler' >&2

  exit 2
}

if [ $# = 0 ]; then
  usage
fi

command="$1"
shift
case "$command" in
whitespace)
  exec python whitespace.py "$@"
  ;;
mkws)
  exec python mkws.py "$@"
  ;;
*)
  usage
  ;;
esac
