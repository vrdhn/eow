#!/bin/sh



do_rel ( )
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


do_git ( )
{
    url=$2
    sha=$(git ls-remote -q $url refs/heads/master | cut -c1-8)
    ver=git-$1-$sha
    date=$(date -I -u)
    
    
    dir=$(pwd)/inst/espeakng/$ver@$date

    bld=espeak-ng-$ver-$date
    
    
    rm -fr $file $dir $bld &&
	mkdir -p $dir &&
	git clone --depth 1 $url $bld &&
	cd $bld &&
	./autogen.sh > /dev/null &&
	./configure --prefix=$dir --disable-shared > /dev/null &&
	make > /dev/null &&
	make install > /dev/null &&
	cd .. &&
	rm -fr $file $bld
    
    
}


do_rel 1.49.1 2017-01-21
do_rel 1.49.0 2017-09-26
do_git espeakng https://github.com/espeak-ng/espeak-ng.git
do_git vrdhn https://github.com/vrdhn/espeak-ng.git 


