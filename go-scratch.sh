#!/bin/bash

go-scratch () {
    mkdir -p -- "$HOME/projekt/go-scratch/$1" &&
        cd -P -- "$HOME/projekt/go-scratch/$1" &&
        go mod init scratch
}
