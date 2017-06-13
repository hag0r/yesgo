# yesgo
Go implementation of [yes tool](https://www.gnu.org/software/coreutils/manual/html_node/yes-invocation.html#yes-invocation)

This [reddit thread](https://www.reddit.com/r/unix/comments/6gxduc/how_is_gnu_yes_so_fast/) showed that the `GNU yes` tool
is extremely fast.

This is a implementation of the `yes` tool in Go. It's extremely fast and maybe even a bit faster than GNU yes.

## Install
    go get github.com/sthagedorn/yesgo
    $ yesgo | pv > /dev/null
    ... [5,66GiB/s] ...
