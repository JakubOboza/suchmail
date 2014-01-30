# usage
This is a simple tool to send emails from command line. You can use it like this:

`./suchmail -to=<recipent email@example.com> -from=<your email@example.com> -password=<enter your password here> < src/github.com/jakuboboza/suchmail/main.go`

This command will send contents of this file via email.

You can also run it with -h to see optional arguments you can use.

* If you need more than one recipent just use , eg. `one@example.com,two@example.com,three@example.com` don't use space!
* If you will not pipe output to it, you will get black input and it will wait for characters until CTRL-D is pressed.


