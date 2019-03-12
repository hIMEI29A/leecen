#!/usr/bin/env bash

# Functions
###########

# Gets latest release
get_latest_release() {
  curl --silent "https://api.github.com/repos/hIMEI29A/leecen/releases/latest" | 
    grep '"tag_name":' |                                            
    sed -E 's/.*"([^"]+)".*/\1/'                                    
}

# Install
install() {
  wget https://github.com/hIMEI29A/leecen/releases/download/$1/\
      leecen-"$1"-amd64.deb && sudo dpkg -i leecen-"$1"-amd64.deb
}

mkdir $HOME/.leecen
mkdir $HOME/.leecen/templates

cp -r ./templates/headers  $HOME/.leecen/templates/headers
cp -r ./templates/licenses  $HOME/.leecen/templates/licenses
# cd
# rm -r leecen