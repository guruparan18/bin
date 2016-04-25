SERV_FILE=$HOME/cfg/sshfiles.txt
get_file () {
        echo " Getting file $1 from $serv_name"
        scp -p $serv_name:$1 .
}
send_file () {
        fl=$(basename $1)
        dr=$(dirname $1)
        dos2unix $fl
        chmod 666 $fl

        if [[ $dr = "." ]]; then
          dr="/tmp/y317522/"
        fi

        echo " Sending file $fl to $serv_name:$dr"
        scp -Cp $fl $serv_name:$dr
}

if [[ $# -eq 0 ]]; then
 echo "Usage: go [get|send] servername /path/to/file"
 echo "  * get  needs full remote path and copies the file to the current local path"
 echo "  * send needs full remote path and copies the file from the current local path"  
 awk '{printf("%-10s %-50s\n", $2, $1)}' $SERV_FILE
 exit
fi

serv_name=$(grep -i $2 $SERV_FILE | awk '{print $1}' )
if [[ $1 = "get" ]]
then
  get_file "$3"
elif [[ $1 = "send" ]]
then
  send_file "$3"
fi
