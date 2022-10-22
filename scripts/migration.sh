
cd ${PWD}/src/migrations/goose
MYSQL_USER="root"
MYSQL_PASS="secret"
MYSQL_HOST="0.0.0.0"
MYSQL_PORT="3306"
MYSQL_DB="shop_car"

goose mysql "${MYSQL_USER}:${MYSQL_PASS}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DB}?charset=utf8" up