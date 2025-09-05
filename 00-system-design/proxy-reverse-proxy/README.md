# Proxy

## Apa itu Proxy?

Penghubung/Middleman dari Client ke Server.

<img width="1108" height="227" alt="image" src="https://github.com/user-attachments/assets/35e1c2a5-b48a-42e2-8fd6-3c1b47009baf" />

#### Use-case of Proxy:

<ol>
  <li>Receive request from client and relay to server.</li>
</ol>

#### Advantages of using Proxy:

<ol>
  <li>Filter request apa aja yang masuk dari client.</li>
  <li>Log Request (mencatat request apa saja yang masuk).</li>
  <li>Transform Request:
    <ul>
      <li>Add/Remove Header</li>
      <li>Encrypt/Decrypt</li>
      <li>Compression</li>
    </ul>
  </li>
</ol>

#### Jenis-jenis Proxy:

Proxy dibagi jadi 2 jenis:<br/>
<ol>
  <li>Forward Proxy</li>
  <li>Reverse Proxy</li>
</ol>

### Apa itu Forward Proxy?

Literally proxy yang tadi kita bahas diatas, yaitu proxy dari client ke server. <br/>
Contoh kasusnya:<br/>
> Multiple Client interact with one proxy terus diterusin requestnya ke server.

<img width="1110" height="565" alt="image" src="https://github.com/user-attachments/assets/a0e7f962-bb74-49c6-8177-1ef983b1683e" />

Forward Proxy = Proxy yang ada di depan client. Ketika client-client membuat request ke backend server, proxy ini akan intercept requestnya, dan instead client yang berkomunikasi ke backend, ini proxynya yang berkomunikasi ke backend. <br/>
Analogi sederhananya semacam:

<img width="923" height="483" alt="image" src="https://github.com/user-attachments/assets/83262f15-ac4e-4f90-8414-feb7b6e19f55" />

#### Advantages of Forward Proxy:

<ol>
  <li>Block access to certain content.</li>
  <li>Allow access to geo-restricted content.</li>
  <li>Provides anonymity (chef ga kenal sama customernya).</li>
</ol>

#### Disadvantages of Forward Proxy:

<ol>
  <li>Ada few information dari Customer yang tetap terbawa.</li>
</ol>

### Apa itu Reverse Proxy?

instead of proxy yang ada di client-side, reverse-proxy ada di server-side.

<img width="1889" height="651" alt="image" src="https://github.com/user-attachments/assets/705b6b8b-3bd9-47a4-877e-74aa51b1df96" />

Reverse Proxy diibaratin sebagai perwakilan dari server kita. Dia **ensures that no client communicate directly with server.**<br/>
Beda dengan forward proxy yang dianalogikan sebagai <b>"waiters"</b>, Reverse Proxy dapat kita analogikan sebagai Head chef-nya dan server sebagai asisten kokinya. Head chef makesured no client communicate ke asisten kokinya.<br/>

#### Advantages of Reverse Proxy:

<ol>
  <li>Load balancing: memastikan distribusi request secara merata.</li>
  <li>Caching</li>
  <li>SSL Encryption</li>
  <li>Improve Scalability</li>
  <li>Improve Security: Karena kita bisa pasang authentication pada reverse-proxynya.</li>
</ol>

#### Forward Proxy vs Reverse Proxy:

<ol>
  <li>Forward Proxy: Customer (client) -> Waiter (proxy)</li>
  <li>Reverse Proxy: Head Chef (reverse-proxy) -> Asisten Chef (server)</li>
</ol>

Forward Proxy dan Reverse Proxy bukan sesuatu yang harus dipilih. Tapi lebih ke fitur yang dipakai dengan use-case tertentu.<br/>

ex.)

<ol>
  <li>Forward Proxy => mengakses request secara anonimus.</li>
  <li>Reverse Proxy => memastikan request terdistribusi dengan baik ke server.</li>
</ol>

### Reverse Proxy vs Load Balancer

<ol>
  <li>Load Balancer => is useful when we have <b>multiple server</b>.</li>
  <li>Reverse Proxy => can be useful with just <b>one server</b>. Karena bisa dipakai untuk encryption, dsbnya.</li>
</ol>

<b>Reverse Proxy can act as load balancer, but not the other way around.</b><br/>
Beberapa contoh dari proxy:
<ol>
  <li>Nginx</li>
  <li>Traefik</li>
</ol>

##### Q: Apakah kita bisa pakai both reverse proxy and load balancer at the same time?
##### A: bisa-bisa aja
<img width="2161" height="651" alt="image" src="https://github.com/user-attachments/assets/2789d3b8-c11a-402a-b19b-ac6a650acc88" />

### Reverse Proxy vs API Gateway

API Gateway Usecase:
<ol>
  <li>
    <b>Routing</b><br/>
    ex.) https://.../product/1
  </li>
  <li><b>Service Discovery</b></li>
  <li>
    <b>Circuit Breaker</b><br/>
    Ketika server mati, nah dia bisa tutup request secara otomatis. Ibaratnya, kalau ada circuit breaker, request ga perlu routing secara terus-menerus.
  </li>
  <li>
    <b>Rate Limiting</b><br/>
    Membatasi X request dalam Y waktu.
  </li>
</ol>
