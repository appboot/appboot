cd /app/js
JSPATH=`find . -name 'app.*.js'`
JSNAME=`echo ${JSPATH#*/}`
DEFAULT_API_HOST="http://api.appboot.com:8000"
API_HOST="http://${HOST_ADDRESS}:8000"
# https://clubmate.fi/replace-strings-in-files-with-the-sed-bash-command/
sed -i "s|${DEFAULT_API_HOST}|${API_HOST}|g" $JSNAME

cd /app
nginx -g 'daemon off;'