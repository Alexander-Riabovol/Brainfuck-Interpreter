echo -n "Enter your Brainfuck code: "
read code

memory=(0 0 0)
ptr=0

loop_queue=(0)

code_len=${#code}

for (( i=0; i<code_len; i++ )); do
    
    command="${code:$i:1}"
    
    if [ $command == '+' ]; then
    ((memory[$ptr]++))
    elif [ $command == '-' ]; then
    ((memory[$ptr]--))
    elif [ $command == '>' ]; then
    ((ptr++))
    elif [ $command == '<' ]; then
    ((ptr--))
    elif [ $command == '.' ]; then
    printf "\x$(printf %x ${memory[$ptr]})"
    elif [ $command == ',' ]; then
    read -n 1 input 
    memory[$ptr] = $(printf "%d" "'$input")
    elif [ $command == '[' ]; then
    loop_queue[${#loop_queue[@]}]=$i
    elif [ $command == ']' ]; then
        if [ ${memory[$ptr]} == 0 ]; then
            unset loop_queue[$((${#loop_queue[@]} - 1))]
        else
            i=${loop_queue[$((${#loop_queue[@]} - 1))]}
        fi
    else 
    echo 'Unknown command: ' $command
    fi

done