#!/bin/bash

echo "register docker container texlive2017..."

BASEDIR=`pwd`

# build remote invokerc
go get github.com/go-ini/ini
mkdir -p $BASEDIR/bin
go build -o $BASEDIR/bin/rinvokerc $BASEDIR/src/rinvokerc.go

#make alias to pdflatex
ln -s $BASEDIR/bin/rinvokerc $BASEDIR/bin/pdflatex

# prompt for name of file or directory
echo -n "Please input your latex work dir: "
read WORKDIR

# check if it exists and is readable
if [ ! -r "$WORKDIR" ]
then
    echo "$WORKDIR is not readable";
    # if not, exit with an exit code != 0
    exit 2;
fi

cd $WORKDIR

# run docker container
docker run -p 8000:8000 -v `pwd`:`pwd` --name texlive2017  -d --rm lazydomino/texlive2017


echo "####################"
echo "PUT this line into your ~/.bashrc"
echo "export PATH=$BASEDIR/bin:\$PATH"

echo "####################"

echo "PUT these lines into your ~/.rinvokerc configure file:"
echo "[pdflatex]"
echo "hostIP=localhost"
echo "port=8000"
echo "workDir=$WORKDIR"
echo "####################"

echo ""
echo ""

echo "Remember to use 'docker stop texlive2017' to stop the docker container!"
