#!/bin/bash
echo 'Hi this will build suchmail. Important to not is that it will change your GOPATH!'
echo 'before we start this is your current GOPATH'
echo $GOPATH
echo 'Ok! lets go!'
set -x

export GOPATH=`pwd`

cd src/github.com/jakuboboza/suchmail
go install

set +x

echo 'Job Done?! Check bin/suchemail, you can see options by calling ./bin/suchemail -h'
echo "Such success, much email"