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


