# CDN

## Apa itu CDN?

CDN, enables every machines from different locations to get the resources with short amount of time.

<img width="741" height="135" alt="image" src="https://github.com/user-attachments/assets/3ebc6d98-ddaa-490c-8fba-17dcf5f652ae" />

## How does CDN works?

CDN sendiri adalah server jaringan yang tersebar di berbagai lokasi. Dia dapat mempercepat proses pendapatan suatu resources dengan membuat cache dari resources tersebut.

<img width="720" height="223" alt="image" src="https://github.com/user-attachments/assets/4d6777e0-fcbb-4afe-ba68-e68050ca4c18" />

Jadi, berdasarkan CDN ini, kita harus selalu host static assets di CDN.

## Ada 2 cara how content are distributed into CDN:

- push CDN
- pull CDN

### Push CDN

Receives new content whenever changes occur on the server, Suitables for sites with small traffic. (Setiap ada perubahan di server, CDN secara otomatis mendapatkan dan mencatat perubahan tersebut.)

### Pull CDN

CDN Cache is updated based on request. Suitable for sites with heavy traffic. (Perubahan di CDN di catat ketika ada perintah/request tertentu saja, tidak setiap saat seperti push CDN yang mencatat setiap perubahan ke dalam Cache CDN, cocok untuk website dengan heavy traffic)

## Disadvantages:

1. Extra Changes can be expensive.
2. Location & Majority audience located in country without CDN, data may have to travel further.
3. Restrictions: Chance that some countries/organization don't allow CDN.

### How Does CDN Works in Details

<img width="3237" height="2868" alt="cabdb67b-7b7f-423f-a9d5-7dce167d88cb_3237x2868" src="https://github.com/user-attachments/assets/b9175599-eccc-4870-a31c-55728ef82b73" />

1. Bob mengetik `www.myshop.com` di browser. Browser secara otomatis mencari `www.myshop.com` di local CDN Cache (Cache Browser kita).

2. Kalau domain name gaada di local DNS Cache (Cache browser), Browsernya menggunakan DNS Resolver (Cloudfare contohnya) yang ada di Internet Service Provider (Indihome, Biznet).

3. DNS Resolvernya secara rekursif mencari nama domain di internet (bro surves the internet) sampai ketemu. 

4. Kalau ga pake CDN, DNS Resolver-nya bakal return IP Address dari website kita yang udah diubah dalam bentuk DNS langsung `www.myshop.com`. Tapi dengan CDN, DNS Resolver bakal return IP Address dari website kita yang udah dalam bentuk DNS CDN `www.myshop.cdn.com` (the domain name of the CDN server).

5. DNS Resolver bakal nyuruh `www.myshop.cdn.com` buat nunjuk ke website aslinya.

6. Disini, dia bakal lewat ke website dari load balancernya dulu (The authoritative name server returns the domain name for the load balancer of CDN `www.myshop.lb.com`.)

7. DNS Resolver bakal nyuruh CDN load balancer buat resolve `www.myshop.lb.com` untuk memilih server CDN yang optimal (Optimal Edge Server) yaitu server yang paling dekat ke lokasi User.

8. CDN load balancer bakal return server’s IP address buat `www.myshop.lb.com` yang paling optimal.

9. Sekarang, kita bisa visit website dari browser kita yang merupakan website dari lokasi CDN terdekat ke lokasi kita.  

10. Browser bakal minta CDN edge server buat load content website kita. Ada 2 jenis konten yang disimpan di CDN Website: static contents dan dynamic contents. Biasanya CDN nyimpen static pages, pictures, videos.

11. Kalau CDN Server kita ga punya konten yang dipengenin, dia bakal ke regional CDN Server terdekat buat lakuin step 4, Kalau regional CDN masih ga punya, dia minta ke CDN pusat.
