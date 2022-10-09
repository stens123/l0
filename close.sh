#!/bin/bash

fuser -k 4222/tcp &> /dev/null &
fuser -k 3333/tcp &> /dev/null &
