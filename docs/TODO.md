# TODO

## Chettlr, Blog

The Chettlr blog is my chess blog. Here's the TODO list items for it:

- [ ] Find a way to generate the "<chess-carousels>" from a Lichess game or analysis board.
  - An extension may be the way forward here, unless we invest time now and create infrastructure to parse PGNs and feed them to an engine or intermediary that can export to FEN format(s).
- [ ] Actually write article(s).
- [ ] Find a way to auto-deploy articles from Zettlr to the Website.

## Chettlr, Analysis 

- [ ] Before worrying about the front-end, we need to actually have a functionality to display.
    - [ ] From the front-end, pull the last game from a given user into Lichess.
    - [ ] Parse the PGN, feeding the moves into an intermediary (or directly to an engine).
    - [ ] Collect data and analyze it, providing it for display purposes back to the user in the front-end.

## Chettlr, Puzzles

This assumes that the PGN to intermediary has already been solved and now we build on top of that and provide the following puzzle functionality to the user:

- [ ] Anti-Puzzle(s), puzzles where there is no tactical motif and the user must showcase an understanding of _when not to look for tactics_.