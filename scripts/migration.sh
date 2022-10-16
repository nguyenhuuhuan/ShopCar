
cd ${PWD}/src/migrations/goose
MYSQL_USER="root"
MYSQL_PASS="elnino19031999"
MYSQL_HOST="localhost"
MYSQL_PORT="3306"
MYSQL_DB="shop_car"

goose mysql "${MYSQL_USER}:${MYSQL_PASS}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DB}?charset=utf8" up