# todo

Simple todo list manager for the CLI.

## Build

```sh
go build -o todo main.go
```

## Add to path

Just put the binary somewhere in your path.
I've added `~/.config/todo` to my path, so it's as simple as:

```sh
cp ./todo ~/.config/todo
```

## Run

```sh
todo help

todo add -d 2025-04-01 Take the bins out

todo ls

todo done 1
```

## Notes

`todo` will add a config.json file and a todos.json file under ~/.config/todo/
Make sure those files are available to be used.
