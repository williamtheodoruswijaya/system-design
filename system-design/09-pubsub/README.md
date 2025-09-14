# Publish-Subscribe

* Producer --> Publisher (Send Request)
* Consumer --> Subscriber (Handle request)

Kalau Message Queue, 1 Request cuman bisa di handle oleh 1 consumer (1 producer, 1 consumer), kalau PubSub, 1 request bisa dihandle oleh banyak subscriber.

<img width="722" height="356" alt="image" src="https://github.com/user-attachments/assets/a5ffcd73-362f-4ab5-a2df-16a62e5c19af" />

PubSub technically remove the queue system, instead kita langsung lempar request ke `n` subscriber. Producer tinggal lempar topic biar `n` subscriber tersebut bisa tinggal baca aja dengan tiap subscriber memiliki kecepatan proses request masing-masing tanpa mempengaruhi satu sama lain termasuk producernya sendiri. PubSub Model basically adalah service-to-service asynchronous communication where each subscriber can get the messages via topic.

## Advantages

1. **Independent Scaling**: develop and scale each publisher or subscriber independently.
2. **Simplify Communication**: removing point-to-point connections and replace it with a single connection. (Dengan menggunakan topic, publisher cukup build single connection ke topic nanti tinggal topic yang handle oleh each subscriber.)
3. **Less Error Prone**: Publisher simply post message-nya ke topic, making it have a small chances for error.
4. **Reduce Delivery Latency**.
5. **Google PubSub guarantee at-least once delivery by making every message seolah-olah di duplicate.**
6. **Code Decoupling between Publisher and Subscriber.**
