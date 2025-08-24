# Docker

## I. Pengenalan Container

### Virtual Machine

<ol>
  <li>Dalam sebuah VM, kita bisa install operation system-nya dan sebuah VM itu wajib mempunyai sebuah Operating System terlepas digunakan atau tidaknya Operating System tersebut.</li>
  <li>Masalahnya, ketika menggunakan VM, dia bakal makan banyak resources karena boot system untuk OS ketika restart VM-nya.</li>
  <li>Intinya, VM tu makan banyak storage (OS + App) + Memory (RAM).</li>
</ol>

<img width="1193" height="456" alt="image" src="https://github.com/user-attachments/assets/0b646857-f9be-4ac6-baad-b919e1ec0d61" />

Tiap VM harus punya OS-nya masing-masing. Selain storage yang udah kepake sama aplikasi, ini ada OS yang tiap VM wajib punya untuk run aplikasinya. Oleh karena itu, penggunaan VM untuk host aplikasi cenderung berat dan sulit.  

### Container

<ol>
  <li>Berbeda dengan VM, Container sendiri berfokus pada sisi aplikasinya.</li>
  <li>Jadi container sendiri sebenarnya berjalan diatas aplikasi Container Manager yang berjalan di sistem operasi.</li>
  <li>Beda dengan VM yang jalan di atas Hypervisor yang manage lebih dari 1 VM.</li>
  <li>Yang membedakan dengan VM adalah, pada Container, kita bisa mem-package aplikasi dan dependency-nya tanpa harus menggabungkan sistem operasi.</li>
  <li>Container akan menggunakan sistem operasi host dimana Container Manager nya berjalan, jadi kalau misal container di install di Linux, maka dia akan menggunakan OS bawaan dari Container Managernya yaitu Linux.</li>
  <li>Beda dengan VM, Kalau kita pakai VM, kita harus install lagi OS khusus untuk VM tersebut.</li>
  <li>Jadi Container tu di sharing lah ibaratnya untuk Operating Systemnya.</li>
  <li>Karena itu Container tidak butuh sistem operasi sendiri.</li>
  <li>Ukuran Container biasanya hanya hitungan MB, berbeda dengan VM yang bisa sampai GB karena di dalamnya ada sistem operasinya.</li>
  <li>Container bisa menggunakan sistem operasi bawaannya.</li>
</ol>

<img width="591" height="518" alt="image" src="https://github.com/user-attachments/assets/f7ac1935-1361-48c2-bc41-24a61e1a1b24" />

<p style="text-align: justify;">
  Nah, anggep Infrastructure itu Laptop kita, nah Laptop kita itu OS-nya Windows, lalu kita install Container Manager yaitu Docker di laptop kita. Nah, dengan Container, kita bisa membuat Container tersebut isinya aplikasi-aplikasinya aja tanpa harus menginstall OS. Dan Container antar Container itu saling terisolasi, jadi App A dan App B tidak akan mempengaruhi satu sama lain.
</p>

## II. Pengenalan Docker

<ul>
  <li>Docker adalah salah satu implementasi dari Container Manager dimana kita bisa manage Container-container yang ada dan memasukkan aplikasi ke dalam Container tersebut.</li>
</ul>

### Docker Architecture

<ol>
  <li>Saat kita menginstall Docker, Docker menggunakan arsitektur Client dan juga Server. Maksudnya apa? Jadi kalau kita install Docker, dalam Docker ada 2 aplikasi yaitu aplikasi Client yang kita gunakan, dan aplikasi Server yang digunakan untuk manage Docker-nya.</li>
  <li>Jadi ketika kita memanage dari aplikasi Client, dia akan berkomunikasi ke Server yang namanya itu Docker Daemon.</li>
  <li>Nah, ketika kita memberikan command-command dalam Docker tersebut, maka perintah/command tersebut akan dikirimkan ke Docker Daemonnya.</li>
  <li>Jadi waktu kita run command di Docker Client, kita harus menjalankan Docker Servernya (Biasanya, Docker Dekstop harus dibuka dulu kalau di Laptop).</li>
</ol>

<img width="1233" height="651" alt="image" src="https://github.com/user-attachments/assets/7961c533-aa17-44ba-8e2e-3116d380b08c" />

<p style="text-align: justify;">
  Jadi inti dari gambar ini, kita bakal ada 2 aplikasi yaitu Client dan Server. Dimana setiap perintah yang kita gunakan dalam Client Application, perintah tersebut akan dieksekusi di Docker Daemon (Server Application). Contohnya seperti membuat Container, Download Images, dan sebagainya.
</p>

### Docker Registry

<ol>
  <li>Apa itu Docker Registry? Ini adalah tempat kita menempatkan Docker Image.</li>
  <li>Dengan Docker Registry, kita bisa menyimpan Docker Image yang kita buat atau bahkan Docker Image milik orang lain disini.</li>
  <li>Docker Image yang ada di Docker Registry, nanti bisa di run oleh Docker Daemon kita.</li>
</ol>

<img width="886" height="575" alt="image" src="https://github.com/user-attachments/assets/64cdb7ba-b88b-4d99-9ec8-856ddef78f3a" />

### Docker Image

<ol>
  <li>Docker Image itu mirip sama installer aplikasi, dimana dalam Docker Image ada aplikasi dan semua dependency yang dibutuhkan aplikasi tersebut.</li>
  <li>Ini yang membedakan dengan VM, kalau VM ada aplikasi, OS, dan dependency, sementara Docker Image cuman ada aplikasi dan dependency-nya.</li>
  <li>Cara lihat Docker Image di Server: <b>docker image ls</b></li>
  <li>Cara download Docker Image dari Docker Registry (download redis image dari Docker Hub): <b>docker pull redis:latest</b></li>
  <li>Cara menghapus Docker Image: <b>docker image rm nama-image:tag</b></li>
</ol>  

### Docker Container

<ol>
  <li>Kalau Docker Image itu installer aplikasi-nya, nah Docker Container ini adalah hasil dari installer aplikasi-nya, yaitu aplikasi-nya itu sendiri.</li>
  <li>Berbeda dengan laptop kita yang kalau install aplikasi itu cuman 1 kali, kalau Docker berbeda.</li>
  <li>Satu Docker Image itu bisa digunakan untuk membuat beberapa Docker Container, asalkan nama Docker Containernya berbeda.</li>
  <li>Jadi misal, kita sebenernya bisa punya multiple Docker Image buat running Redis asalkan nama containernya berbeda.</li>
  <li>Ketika kita sudah membuat Container, kita tidak bisa menghapus image-nya yang ada dalam Docker Container tersebut.</li>
  <li>Ga kek aplikasi, kita bisa hapus installer aplikasinya tanpa menghapus aplikasinya.</li>
  <li>Nah, kalau Docker Container, Docker Image yang merupakan installer aplikasinya, itu gabisa dihapus kecuali kita hapus Docker Containernya langsung.</li>
  <li>Hal ini karena sebenernya Docker Container yang merupakan hasil dari installer aplikasi (Docker Image) itu tidak membuat copy hasil dari Docker Image tersebut melainkan tetap menggunakan isi dari Docker Imagenya.</li>
  <li>Cara melihat Container yang ada dalam Docker Daemon: <b>docker container ls -a</b></li>
  <li>Cara buat Docker Container: <b>docker container create --name namaContainer namaImage:tag</b></li>
  <li>Cara run Docker Container: <b>docker container start namaContainer</b></li>
  <li>Cara stop a running Docker Container: <b>docker container stop namaContainer</b></li>
  <li>Cara menghapus Docker Container: <b>docker container rm namaContainer</b></li>
  <li><b>Notes: semua ini bisa diatur pake Docker Dekstop jadi gaush pusing-pusing apalin ni command la ya...</b></li>
  <li><b>Notes: port yang ada dalam sebuah Container, tidak akan mempengaruhi satu sama yang lain. Misal, docker container yang punya redis akan running di port 6379, tapi kita masih bisa menggunakan port tersebut di laptop kita tanpa harus khawatir port tersebut occupied.</b></li>
</ol>

#### Container Log

<ol>
  <li>Kadang saat terjadi masalah dengan aplikasi di Container, kita ingin melihat log dari aplikasinya.</li>
  <li>Hal ini dilakukan untuk debugging atau melihat detail masalah dari aplikasi kita.</li>
  <li>Caranya: <b>docker container logs namaContainer</b></li>
  <li>Atau pake aja tu Docker Dekstop 😎</li>
</ol>

#### Container Exec
