#!/bin/sh


ip=94.177.245.217
dr=/home/ttsso

rsync bin/eow ttsso@$ip:$dr/
rsync -a get-espeakng.sh css images templates ttsso@$ip:$dr/
