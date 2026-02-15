function ct_logs() {
    curl -s 'https://www.gstatic.com/ct/log_list/v3/log_list.json' | jq -r '.operators[].logs[].url' | sed 's/^/  - /'
}

function update_env() {
    local urls="$1"
    
    grep -vE "^ct_logs:|^  -" env.yaml > env.tmp
    echo "ct_logs:" >> env.tmp
    echo "$urls" >> env.tmp
    mv env.tmp env.yaml
}

urls=$(ct_logs)
update_env "$urls"