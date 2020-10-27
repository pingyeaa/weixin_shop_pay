#!/usr/bin/expect

set PORT 22
set HOST 47.112.229.178
set USER root
set PASSWORD Ty2020!#!\$2030
set FILE ./hdzs-api-go/hdzs-api_linux_amd64

set PORT1 22
set HOST1 47.115.120.240
set USER1 root
set PASSWORD1 Ty2020!#!\$2030

# 删除服务器的二进制文件
spawn ssh -p $PORT $USER@$HOST "rm -rf /data/webapp/hdzs-api-8202/hdzs-api_linux_amd64"
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD\r" }
}
expect EOF

# 将新的二进制文件传入服务器
spawn scp $FILE $USER@$HOST:/data/webapp/hdzs-api-8202/
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD\r" }
}
expect EOF

# 查询当前服务的进程ID
spawn ssh -p $PORT $USER@$HOST "ps -ef | grep \"/data/webapp/hdzs-api-8202/hdzs-api_linux_amd64\""
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD\r" }
}
expect EOF

# 平滑重启服务
spawn ssh -p $PORT $USER@$HOST "kill -n USR1 `pgrep -f /data/webapp/hdzs-api-8202/hdzs-api_linux_amd64`"
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD\r" }
}
expect EOF

# 删除服务器的二进制文件
spawn ssh -p $PORT1 $USER1@$HOST1 "rm -rf /data/webapp/hdzs-api-8202/hdzs-api_linux_amd64"
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD1\r" }
}
expect EOF

# 将新的二进制文件传入服务器
spawn scp $FILE $USER1@$HOST1:/data/webapp/hdzs-api-8202/
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD1\r" }
}
expect EOF

# 查询当前服务的进程ID
spawn ssh -p $PORT1 $USER1@$HOST1 "ps -ef | grep \"/data/webapp/hdzs-api-8202/hdzs-api_linux_amd64\""
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD1\r" }
}
expect EOF

# 平滑重启服务
spawn ssh -p $PORT1 $USER1@$HOST1 "kill -n USR1 `pgrep -f /data/webapp/hdzs-api-8202/hdzs-api_linux_amd64`"
expect {
  "yes/no" {send "yes\r";exp_continue;}
  "*password:*" { send "$PASSWORD1\r" }
}
expect EOF