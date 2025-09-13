# Load Balancer

## Apa itu Load Balancer?

<img width="861" height="361" alt="image" src="https://github.com/user-attachments/assets/feedd938-43d6-4cd0-a451-b8ea39eb1171" />

Fungsi utama:<br/>
Distribute incoming network traffic across multiple servers ensuring high availability and **scalability**.

### Q: Gimana Load Balancer bisa improve Scalability?

<img width="1385" height="362" alt="image" src="https://github.com/user-attachments/assets/3c8cfa7c-8c3b-4dc2-b629-6ff2c15150ec" />

#### Conclusions:

Intinya, kita bisa pakai Load Balancer ke semua network distributions, bahkan proses read/write ke database yang punya banyak replication.

### Q: Why'd we need Load Balancer? Why not make it into only one server?

Server can be overworked, which can reduce performance. Therefore, we need multiple server, **but how'd we distribute the request?** **USE LOAD BALANCER!!**

## Routing Algorithms

<img width="3735" height="3573" alt="0251-lb-algorithms" src="https://github.com/user-attachments/assets/531174ed-858b-4cee-afee-8f26a4883290" />

<ol>
  <li>
    <b>Round Robin</b><br/>
    Distribute requests in rotation (abis ke server 1, ke server 2, ke server 3, dst...)
  </li>
  <li>
    <b>Sticky Round Robin</b><br/>
    Ini improvement dari Round Robin biasa, kalau Request dari Alice pertama-tama udah diarahkan ke server 1, maka semua request selanjutnya dari Alice akan diteruskan ke server 1.
  </li>
  <li>
    <b>Weighted Round Robin</b><br/>
    Assign Weight to each server (Round Robin + If command to check weight for server where weight = how many request one server can handle)
  </li>
  <li>
    <b>Hashing: hash client ip address.</b>
  </li>
  <li>
    <b>Least Connections</b><br/>
    Request will be sent to server with fewest connections.
  </li>
  <li>
    <b>Least Response Time</b><br/>
    Request will be sent to server with least response time.<br/>
    Kita tau kalau semakin banyak request terhadap sebuah server, semakin berat performance-nya. Artinya, response time-nya juga meningkat. Nah algoritma ini akan mendistribusikan request ke sever dengan response time tercepat.
  </li>
</ol>

## Disadvantages of Load Balancer

If a Load Balancer fails (single point of failure), **that's why we have redundant load balancer (n load balancer).**<br/>
Jadi sebenarnya, diagram Load balancer kita gambarin kayak begini:

<img width="995" height="492" alt="image" src="https://github.com/user-attachments/assets/3d886d5a-9d3f-4543-93ce-c5659f002221" />

ex.) Load Balancer:

- Azure Load Balancing
- Amazon Elastic Load Balancing
- Digital Ocean
- Nginx

## Advantages of Load Balancer

<ol>
  <li><b>Autoscaling:</b> Add/Remove instances automatically.</li>
  <li><b>Encryption:</b> Encrypted Connection (SSL) -> Kalau pake Load Balancer bisa jadi https dari http kek pas pake AWS.</li>
  <li><b>Health Check:</b> Deteksi server mati/engga (periodically) otomatis matiin server, terus spawn new server.</li>
  <li><b>Compression:</b> Compress size biar ga gede.</li>
  <li><b>Sticky Session:</b> Assign same user/device to same resources.</li>
  <li><b>Persistence Connections (Websocket):</b> allow server to create persistent connection to client.</li>
</ol>
