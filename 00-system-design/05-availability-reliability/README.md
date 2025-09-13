# Availability

## Apa itu Availability?

Availability indicate the time a system remains operational. (**Berapa lama sistem kita tetap beroperasional?**).

Cara menghitung Availability:

$$\ \text{availability} = \frac{uptime}{uptime+downtime} \$$

Jadi misal:

Availability = 90%, berarti:

- 1 year = 36 days down
- 1 month = 72 hours down
- 1 week = 16 hours down

Jadi, apakah availability 90% itu bagus? <br/>
Tidak, bayangkan aja dalam 1 bulan ada 3 hari full down. Oleh karena itu ada yang namanya:

### Nine's of Availability

- 99,99% (four nines)
- 99,999% (five nines)
- ...
- 99,9999999 (nine nines *max*) --ini berarti--> 1 year = 32 ms down.

Contoh:

$$\ \text{availability} = \text{99,9 \%} \$$

jadi:
- 1 year = 8 hours down
- 1 month = 40 minutes down
- 1 week = 10 minutes down

Tapi kalau kita punya services, yang consist of multiple component, where each components prone to failure, kita gabisa nyebut availability-nya langsung 3 nines, tapi overall availability diatur dari apakah component-nya are **sequence or parallel**.

### Sequence

Apa itu sequence? Sequence itu ketika ada 2 component yang punya availability masing-masing dalam 1 rangkaian.

ex.) **Component A manggil Component B (dependent)**

<img width="794" height="210" alt="image" src="https://github.com/user-attachments/assets/347200cf-ed28-4ff1-83ec-86c74d4da77d" />

Meskipun each component punya three nines availability, karena dia sequence, overall availability-nya akan decrease. Cara hitungnya:

$$\ \text{Availability(total)} = Availability(Foo) * Availability(Bar) \$$
$$\ \text{Availability(total)} = 99,9 * 99,9 \$$
$$\ \text{Availability(total)} = 99,8 \$$

### Parallel

Apa yang dimaksud dengan Parallel Component? 2 component yang ga rely satu sama lain.

ex.) **Component A dan Component B (independent)**

<img width="794" height="210" alt="image" src="https://github.com/user-attachments/assets/081eed7e-a90b-495e-8e98-aba96ab71b62" />

Kalau dia parallel, overall availability-nya akan **increases**. Cara hitung:

$$\ \text{Availability(total)} = 1 - (1 - \text{Availability(foo)}) * (1 - \text{Availability(bar)}) = 99,9999 \$$

# Reliability

## Apa itu Reliability?

- If a service is reliable, it is available.
- IF a service is available, it doesn't mean it's reliable.

### Gimana cara kita meningkatkan reliability dari sebuah server?

Ada 2 cara:

1. Vertical Scaling: Ibaratnya kita naikin spec computer. (**Expand system availability, by adding more power to existing machine.**)
2. Horizontal Scaling: Instead of naikin spec computer, kita perbanyak jumlah computer-nya. (**Adding more machines/instances to existing machine.**)

### Vertical Scaling vs Horizontal Scaling

#### Vertical Scaling:

<ol>
  <li>(+) Simple to implement (ibarat RAM tinggal pasang aja).</li>
  <li>(+) Easier to manage.</li>
  <li>(+) Data consistent (Storage sama).</li>
  <li>(-) Single point of Failure (kalau laptop mati, ywd kelar tu aplikasi).</li>
  <li>(-) Harder to Upgrade (when system already complex, misal RAM udah 1 TB di total 4 slot, mau nambah 10 TB lagi jadi susah, mending beli laptop baru).</li>
  <li>(-) Risk of High Downtime.</li>
</ol>

#### Horizontal Scaling:

<ol>
  <li>(+) Increased Redundancy (menghindari single point of failure, jadi kalau 1 server mati, ada server lain yang bisa menggantikan).</li>
  <li>(+) Easier to Upgrade (tinggal beli/nambah laptop, tanpa harus ribet mikirin masang RAM dkk).</li>
  <li>(+) Better Fault Tolerance (If system fails, there's no loss in uptime).</li>
  <li>(-) Increase Complexity (susah diimplement karena harus **sync** antar server).</li>
  <li>(-) Data Inconsistency (data antar server bisa berbeda secara langsung karena harus ada sync terlebih dahulu - eventually consistent data).</li>
</ol>

#### FAQ: Kenapa sistem yang available ga selalu reliable?

Ibarat cons dalam horizontal scaling, kita bisa akses mesinnya, tapi data antar mesin tidak konsisten.
