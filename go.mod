module github.com/aryan-jain/pokedex

go 1.22.1

require internal/poke_api v1.0.0

replace internal/poke_api => ./internal/poke_api

require internal/repl v1.0.0

replace internal/repl => ./internal/repl

require internal/cache v1.0.0

replace internal/cache => ./internal/cache
