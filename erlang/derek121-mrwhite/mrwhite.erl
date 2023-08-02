-module(mrwhite).

-export([main/1, main/0]).

main([convert, ws, Filename]) ->
  T = mrwhite_from_whitespace:convert({file_in, Filename}, return),
  io:format("~s", [T]);
main([convert, text, Filename]) ->
  W = mrwhite_from_text:convert({file_in, Filename}, return),
  io:format("~s", [W]);
main([run, ws, Filename]) ->
  mrwhite_from_whitespace:run({file_in, Filename});
main([run, text, Filename]) ->
  mrwhite_from_text:run({file_in, Filename});
main(_) ->
  usage().
main() ->
  usage().

usage() ->
  io:format("Usage: mrwhite (convert | run) (ws | text) <file>\n"),
  halt(1).
