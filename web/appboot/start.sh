cd /app/assets
JSPATH=`find . -name 'index.*.js'`
JSNAME=`echo ${JSPATH#*/}`
DEFAULT_API_URL="ws://127.0.0.1:8000/ws"
DEFAULT_STATIC_URL="http://127.0.0.1:8000"

if [ -z "${API_URL}" ]
then
    echo "API_URL is empty, use the default value ws://127.0.0.1:8000/ws"
else
    # https://clubmate.fi/replace-strings-in-files-with-the-sed-bash-command/
    sed -i "s|${DEFAULT_API_URL}|${API_URL}|g" $JSNAME
fi

if [ -z "${STATIC_URL}" ]
then
    echo "STATIC_URL is empty, use the default value http://127.0.0.1:8000"
else
    # https://clubmate.fi/replace-strings-in-files-with-the-sed-bash-command/
    sed -i "s|${DEFAULT_STATIC_URL}|${STATIC_URL}|g" $JSNAME
fi

cd /app
nginx -g 'daemon off;'
