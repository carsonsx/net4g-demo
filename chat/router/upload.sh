#!/usr/bin/env bash
##!/usr/bin/expect -f
set timeout -1
#spawn ssh root@192.168.56.201
#expect "password:"
#send "root\r"
#interact

#ssh root@192.168.56.201 'rm -rf ~/golang/net4g-chat'
ssh root@192.168.56.201 'mkdir -p ~/golang/net4g-chat/router/vendor/github.com/carsonsx'
scp -r ./ root@192.168.56.201:~/golang/net4g-chat/router
scp -r ../../net4g/ root@192.168.56.201:~/golang/net4g-chat/router/vendor/github.com/carsonsx/
scp -r ../../log4g/ root@192.168.56.201:~/golang/net4g-chat/router/vendor/github.com/carsonsx/
scp -r ../../vendor/ root@192.168.56.201:~/golang/net4g-chat/router/
#ssh root@192.168.56.201 'find ~/golang/net4g-chat/ -name ".git"  | xargs rm -rf'
#ssh root@192.168.56.201 'find ~/golang/net4g-chat/ -name ".DS_Store"  | xargs rm -rf'