# WFI - wait for it

As a (Sys|Dev)Ops, I'm oftenly annoyed to run multilple `nc -v` or `telnet` commands when I'm troubleshooting a network issue.

Yes, `watch` command exists. But where's the fun in that.

I wrote this little piece of software that loops until a connection is succeeded.

## Build

`go build -o bin/wfi cmd/wfi.go`

## Usage

```
NAME:
   wait for it - try to join a host on a port until connection succeeded

USAGE:
   wfi <protocol> <host> <port>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Example

In a terminal window, I'm running :

```
➜  wfi tcp localhost 3306
(nothing happen...)
```

Then in another window, I'm running:

`docker run -d --name mariatest -p 3306:3306 --env MARIADB_ROOT_PASSWORD=my-secret-pw  mariadb:latest`

And in the first window :

```
2024/03/03 12:05:25 connection succeeded
➜
```

## License

MIT Licence (a copy is in this repository)
