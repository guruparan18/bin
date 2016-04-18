SERV_FILE=$HOME/cfg/sshfiles.txt
get_file () {
        echo " Getting file $1 from $serv_name"
        scp -p $serv_name:$1 .
}
send_file () {
        dos2unix $1
        chmod 666 $1
        echo " Sending file $1 to $serv_name"
        scp -Cp $1 $serv_name:/tmp/
}

if [[ $# -eq 0 ]]; then
 echo "Usage: go [get|send] servername /path/to/file"
 echo "  * get  needs full remote path and copies the file to the current local path"
 echo "  * send needs full remote path and copies the file from the current local path"  
 cat -n $SERV_FILE
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
