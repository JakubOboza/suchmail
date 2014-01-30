package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

var from = flag.String("from", "from@default.com", "from this email we will send message")
var toAddr = flag.String("to", "to@default.com", "to this email we will send message")
var server = flag.String("addr", "smtp.gmail.com:587", "mail server to use")
var authserver = flag.String("authaddr", "smtp.gmail.com", "auth server (usually same as addr)")
var password = flag.String("password", "", "your email password")
var ident = flag.String("ident", "", "identity")

func main() {

	flag.Parse()

	to := strings.Split(*toAddr, ",")

	auth := smtp.PlainAuth(
		*ident,
		*from,
		*password,
		*authserver,
	)


  var in *os.File
  var errf error

  switch name := flag.Arg(0); {
  case name == "":
        in = os.Stdin
  default:
        in, errf = os.Open(name)
        if errf != nil {
                fmt.Print(errf)
        }
  }

	msg, inerr := ioutil.ReadAll(in)  // os.Stdin
	if inerr != nil {
		fmt.Print(inerr)
	}

	err := smtp.SendMail(*server, auth, *from, to, msg)
	if err != nil {
		fmt.Print(err)
	}

}
