#!/bin/bash

appdir=/home/yezhiming/binary/email
appdirHistory=/home/yezhiming/binary/email/history
appdirError=/home/yezhiming/binary/email/error
appdirLogs=/home/yezhiming/binary/email/logs
runname=iqoption-go-email
putfile=binary-go-email
port=7107
app=test-iqoption-email

if [ ! -d $appdir ]; then
    mkdir -p $appdir
fi

if [ ! -d $appdirHistory ]; then
    mkdir -p $appdirHistory
fi

if [ ! -d $appdirError ]; then
    mkdir -p $appdirError
fi

if [ ! -d $appdirLogs ]; then
    mkdir -p $appdirLogs
fi


date=`date "+%Y-%m-%d-%H-%M-%S"`

   echo "请确认上传的目录为$appdir,上传包的名称为$putfile"
   sleep 2

   echo "开始更新项目"
   ls -l $appdir | grep $runname | awk '{print $6,$7,$8,$9}'

   echo "1、删除history目录下旧备份"
   sleep 1
   cd $appdir/history/
   rm -rf $runname

   echo "2、停止运行项目代码"
   sleep 1
   app_pid=`ps -ef | grep $runname | grep $port | awk '{print $2}'`
   echo "old PID $app_pid"
   kill -9 $app_pid

   echo "3、移动旧代码到history目录"
   sleep 1
   mv $appdir/$runname $appdirHistory/$runname"_"$date

   echo "4、重命名新版本代码名称"
   sleep 1
   mv $appdir/$putfile $appdir/$runname

   echo "5、启动项目代码"
   cd $appdir
   chmod 755 *
   nohup $appdir/$runname --APP_PORT=$port >/dev/null 2>$appdirError/$runname.log & echo $! >$appdirLogs/$runname-running.pid

   app_new_pid=`ps -ef | grep $runname | grep $port | awk '{print $2}'`
   echo "new PID $app_new_pid"

   if [ $app_new_pid -gt 0 ]
      then
      echo "$runname运行成功"
      echo "$date 代码已更新，$runname已运行" >>$appdirLogs/update.log
      echo -e "$date $app已更新 " | mail -v -s '更新代码通知' 'yezhiming@yuctime.com' >/dev/null 2>&1
#     ls -l $appdir | grep $runname  | awk '{print $6,$7,$8,$9}'
   else
      echo "$runname运行失败，请检查原因"
      echo "$date $runname运行失败，请检查原因" >>$appdirLogs/update.log
      echo -e "$date $app更新失败，请检查原因 " | mail -v -s '更新代码通知' 'yezhiming@yuctime.com' >/dev/null 2>&1
      exit 0
   fi
