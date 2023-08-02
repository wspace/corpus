#!/bin/sh -e

usage() {
  echo 'Usage: <command> [...]

Commands:

    WhitespaceProgramRunner      -- Interpreter
    ListingFileProgramRunner     -- Interpreter for listing files
    WhitespaceCompilerDebugger   -- Assembler
    CreateSimpleOutputFromString -- Text to Whitespace converter' >&2

  exit 2
}

if [ $# = 0 ]; then
  usage
fi

command="$1"
shift
case "$command" in
WhitespaceProgramRunner)
  exec java -cp whitespacesdk.jar uk.co.mash.whitespace.executable.WhitespaceProgramRunner "$@"
  ;;
ListingFileProgramRunner)
  exec java -cp whitespacesdk.jar uk.co.mash.whitespace.executable.ListingFileProgramRunner "$@"
  ;;
WhitespaceCompilerDebugger)
  exec java -cp whitespacesdk.jar uk.co.mash.whitespace.executable.WhitespaceCompilerDebugger "$@"
  ;;
CreateSimpleOutputFromString)
  exec java -cp whitespacesdk.jar uk.co.mash.whitespace.executable.CreateSimpleOutputFromString "$@"
  ;;
*)
  usage
  ;;
esac
