#!/bin/bash
##############################################################
# File Name: build_compress.sh
# Version: 1.0.0
# Author: shiwei-Du
# Created Time : 2020-02-15 11:26:36
# Website: https://www.dushiwei.cn
##############################################################

#use upx compress flag
USE_UPX_COMPRESS=0
BUILD_PACKAGENAME=api-server-demo

# format echo
print_style() {
    if [ "$2" == "info" ] ; then
        COLOR="96m"
    elif [ "$2" == "success" ] ; then
        COLOR="92m"
    elif [ "$2" == "warning" ] ; then
        COLOR="93m"
    elif [ "$2" == "danger" ] ; then
        COLOR="91m"
    else #default color
        COLOR="0m"
    fi
    STARTCOLOR="\e[$COLOR"
    ENDCOLOR="\e[0m"
    printf "$STARTCOLOR%b$ENDCOLOR" "$1"
    echo 
}

# display options
display_options () {
    printf "Available first options:\n";
    print_style "   show" "info"; 
    printf "%-100s\n" "         Display All tables name"
}

# comnand judge
hasCommandByType(){
    if type $1 2>/dev/null;
    then
        return 1
    else
        return 0
    fi
}

# use upx compress
use_upx(){
	print_style "***Use upx Compress beging***" "info"
    upx $BUILD_PACKAGENAME
}

# install fail
install_fail(){
	print_style "***Install upx Fail***" "danger"
}

# install success
install_success(){
	print_style "***Install upx Success***" "success"
}

# install result
install_result(){
	if [[ $? == 0 ]];then
		install_success
		use_upx
	else
		install_fail
	fi
}

# install
install(){
	sys=$1
	mac="Darwin"
	centos=”Centos“
	if [[ $sys =~ $mac ]];then
		brew install upx
		install_result 
	elif [[ $sys =~ $centos ]]; then
		#statements
		yum -y install upx
		install_result
	else
		print_style "***Please install upx by yourself***" "info"
	fi
}

# go build
go build -ldflags '-w -s' .

hasCommandByType upx
returnVue=$?

# echo $returnVue
if [[ $returnVue != 0 ]] ;
then
    use_upx
elif [[ $USE_UPX_COMPRESS == 1 ]]; then
	sys=`uname  -a`
	install $sys
fi