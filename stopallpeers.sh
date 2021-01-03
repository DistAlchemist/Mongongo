#/bin/bash 

MGPATH="mongongo"
BINPATH="bin/mg-server"
USERNAME="2019211170"
for ((i=0;i<5;i=i+1));do
    ssh thumm0$(($i+1)) \
        "ps -U ${USERNAME} -x | grep $BINPATH | grep -v grep | awk '{print $1}' | xargs kill &"&
done