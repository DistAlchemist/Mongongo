#/bin/bash

while [[ "$#" -gt 0 ]]; do
    case $1 in
        -n) COPY=1 ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

MGPATH="mongongo"
BINPATH="bin/mg-server"
for ((i=0;i<5;i=i+1));do
    if [ -z "$COPY" ]
    then
        ssh thumm0$(($i+1)) "mkdir -p $MGPATH/bin"
        scp -r $BINPATH thumm0$(($i+1)):~/${MGPATH}/bin/
    fi
    ssh thumm0$(($i+1)) "cd $MGPATH && ($BINPATH > $MGPATH.log 2>&1&)" &
done
