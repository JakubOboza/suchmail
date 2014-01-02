package main

import (
  "fmt"
  "net/smtp"
  "flag"
  "strings"
  "io/ioutil"
  "os"
)

var from       = flag.String("from", "from@default.com", "from this email we will send message")
var toAddr     = flag.String("to", "to@default.com", "to this email we will send message")
var server     = flag.String("addr", "smtp.gmail.com:587", "mail server to use")
var authserver = flag.String("authaddr", "smtp.gmail.com", "auth server (usually same as addr)")
var password   = flag.String("password", "", "your email password")
var ident      = flag.String("ident", "", "identity")

func main() {

  flag.Parse()

  to := strings.Split(*toAddr, ",")

  auth := smtp.PlainAuth(
          *ident,
          *from,
          *password,
          *authserver,
  )

  msg, inerr := ioutil.ReadAll(os.Stdin)
  if inerr != nil {
    fmt.Print(inerr)
  }

  err := smtp.SendMail(*server, auth, *from, to, msg)
  if err != nil {
    fmt.Print(err)
  }

}