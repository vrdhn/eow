#!/bin/sh



dofor ( )
{
    
    ver=$1
    date=$2
    
    
    dir=$(pwd)/inst/espeakng/$ver@$date

    file=espeak-ng-$ver.tar.gz
    url=https://github.com/espeak-ng/espeak-ng/releases/download/$ver/$file
    bld=espeak-ng-$ver

    #https://github.com/espeak-ng/espeak-ng/archive/1.49.0.tar.gz    
    file=$ver.tar.gz
    url=https://github.com/espeak-ng/espeak-ng/archive/$file
    bld=espeak-ng-$ver
    
    
    rm -fr $file $dir $bld &&
	mkdir -p $dir &&
	wget -q  $url &&
	(tar zxf $file || tar xf $file) &&
	cd $bld &&
	./autogen.sh > /dev/null &&
	./configure --prefix=$dir --disable-shared > /dev/null &&
	make > /dev/null &&
	make install > /dev/null &&
	cd .. &&
	rm -fr $file $bld
    
    
}


dofor 1.49.1 2017-Jan-21
dofor 1.49.0 2017-Sep-26



