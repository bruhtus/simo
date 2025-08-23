# Simo

Simo is a simple pomodoro timer CLI that can be used in status bar,
such as `i3status`.

## Installation

To install the this CLI, we can do:
```sh
go install github.com/bruhtus/simo@latest
```

Or if we want to use specific version, we can do:
```sh
go install github.com/bruhtus/simo@v1.0.0-rc1
```

If using `go install` does not work or we want to change the source code, we
can clone the repo and use `go build` like this:
```sh
git clone https://github.com/bruhtus/simo.git
cd simo
go build
```

## Usage

The available subcommands are:
- status: to show the current time of on going session.
- focus: to start the focus session of pomodoro.
- break: to start the break session of pomodoro.
- pause: to toggle pause or unpause the on going session.
- reset: to reset or stop the on going session.

The default focus session time is 50 minutes and break session is 10 minutes.
We can change the focus or break session time using flag `-t` like this:
```sh
simo focus -t 25m
simo break -t 5m
```

For notification, simo using `notify-send`. If you want to use other
notification command, you need to change the source code and build it again.

To enable notification, we need to provide flag `-n` in the CLI like this:
```sh
simo focus -t 25m -n
simo break -t 5m -n
```

## References

- https://github.com/mskelton/pomo-go
- https://github.com/gen2brain/beeep
- https://zetcode.com/golang/exec-command/
- https://medium.com/@omidahn/building-cli-applications-with-go-1d2fb73bb24e
- https://github.com/NanXiao/golang-101-hacks/blob/master/posts/go-build-vs-go-install.md
- https://www.geeksforgeeks.org/go-language/main-and-init-function-in-golang/
- https://stackoverflow.com/a/38551362 (golang runtime.Caller() example)
- https://stackoverflow.com/a/67905164 (golang FlagSet example)
- https://stackoverflow.com/a/23551970 (when to use pointer or values in function)
