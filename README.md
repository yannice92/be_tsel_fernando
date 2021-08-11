#  DEPENDENCIES
* github.com/go-sql-driver/mysql
* github.com/joho/godotenv
* github.com/julienschmidt/httprouter
# How To Build & Run
* copy .env.example dan paste menjadi .env
* sesuaikan .env dengan config yg ada
* jalankan docker-compose up
* import DB ke dlm docker
`
  docker exec -i telkomselbe_db_1 mysql -uroot -ppassword123 < dump-telkomselbe.sql
`
* import postman collection and environment
# Assumption
1. login untuk rest API di handle dari API GW jadi saya menggunakan customer ID untuk get metadata user nya
2. asumsi untuk view referral & reward ada di bulan berjalan dgn rule sbagai berikut:

min 1 max 1 = 2 GB

min 2 max 5 = 12 GB

min 6 max 9999999 = 20 GB

if total >= min && total <= max
# To Do
1. mungkin akan dibuat 2 service, untuk service customer dan referral
2. untuk rule mungkin akan di load di redis dulu sehingga tidak query ke DB terus menerus
3. blm mengconsider running di multi instance