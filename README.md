Run these command in **bash**

#### Generating our rootCA file.
- `openssl req -newkey rsa:2048 -nodes -x509 -days 365 -out ca.crt -keyout ca.key -subj "/C=BD/ST=Dhaka/L=Dhaka/O=Appscodee, Inc./CN=Appscodea.test Root CA" `

#### Generating server certificate.
- `openssl genrsa -out server.key 2048`
- `openssl req -new -key server.key -out server.csr -subj "/C=BD/ST=Dhaka/L=Dhaka/O=Appscode, Inc./CN=pritam.test"`
- `openssl x509 -req -extfile <(printf "subjectAltName=DNS:pritam.test,DNS:pritam.test") -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt`

#### Generating the client certificate.
- `openssl genrsa -out client.key 2048`
- `openssl req -new -key client.key -out client.csr -subj "/C=BD/ST=Dhaka/L=Dhaka/O=Orange, Inc./CN=pratim.test"`
- `openssl x509 -req -extfile <(printf "subjectAltName=DNS:pratim.test,DNS:pratim.test") -in client.csr -CA ca.crt -CAkey ca.key -out client.crt -days 365 -sha256 -CAcreateserial`