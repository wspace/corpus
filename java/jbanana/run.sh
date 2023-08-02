#!/bin/sh -e

usage() {
  echo 'Usage: <command> [...]

Commands:

    ws,  Whitespace -- Interpreter
    wsa, WsAsm      -- Assembler
    wsd, Wsd        -- Debugger
    wsi, WsInject   -- Injector' >&2

  exit 2
}

if [ $# = 0 ]; then
  usage
fi

command="$1"
shift
case "$command" in
ws|Whitespace)
  exec java -cp whitespace.jar org.codeberg.jbanana.whitespace.Whitespace "$@"
  ;;
wsa|WsAsm)
  exec java -cp whitespace.jar org.codeberg.jbanana.whitespace.WsAsm "$@"
  ;;
wsd|Wsd)
  exec java -cp whitespace.jar org.codeberg.jbanana.whitespace.Wsd "$@"
  ;;
wsi|WsInject)
  exec java -cp whitespace.jar org.codeberg.jbanana.whitespace.WsInject "$@"
  ;;
*)
  usage
  ;;
esac
