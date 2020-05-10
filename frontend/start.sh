cd /app/js
JSPATH=`find . -name 'app.*.js'`
JSNAME=`echo ${JSPATH#*/}`
SOURCE_WS_URL="ws://ws.appboot.com:8888/appboot"
# https://clubmate.fi/replace-strings-in-files-with-the-sed-bash-command/
sed -i "s|${SOURCE_WS_URL}|${WS_URL}|g" $JSNAME

cd /app
nginx -g 'daemon off;'