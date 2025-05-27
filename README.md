📦 go-redis-echo
Proyek sederhana menggunakan Golang, Redis, dan Echo framework untuk menyimpan dan mengambil data user menggunakan Redis Hash.

🚀 Fitur

Insert data ke Redis dengan query parameter

Get data dari Redis berdasarkan ID

Menggunakan Echo sebagai web framework

Redis sebagai penyimpanan sementara (in-memory database)

➕ Insert Data
GET /insert?id=1&name=Muhammad Rivaldhi&age=20

Response:
{
  "data": "1",
  "status": "Successfully inserted data to Redis"
}

🔍 Get Data
GET /get-redis?id=1

Response:

{
  "data": {
    "name": "Muhammad Rivaldhi",
    "age": "20"
  },
  "status": "Successfully retrieved data from Redis"
}

✅ Dependencies

Echo
Go Redis v9

