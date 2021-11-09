# RSA-example
test file
 ```go run main.go```

##Hệ mật mã khóa đối xứng
Là những hệ mật được sử dụng chung 1 khóa trong quá trình mã hóa và mã hóa. Do đó khóa phải được giữ bí mật tuyện đối.
Một số hệ mật mã khóa đối xứng hiện đại mà mình thấy hay được sử dụng là DES, AES, RC4, RC5,... 

###Hệ mật sẽ bao gồm:
- Bản rõ (plaintext-M): bản tin được sinh ra bởi bên gửi 
- Bản mật (ciphertext-C): bản tin che giấu thông tin của bản rõ, được gửi tới bên nhận qua một kênh không bí mật 
- Khóa (Ks): nó là giá trị ngẫu nhiên và bí mật được chia sẻ giữa các bên trao đổi thông tin và được tạo ra từ: 
  - Bên thứ 3 được tin cậy tạo và phân phối tới bên gửi và bên nhận 
  - Hoặc, bên gửi tạo ra và chuyển cho bên nhận 
- Mã hóa (encrypt-E): C = E(KS, M) 
- Giải mã (decrypt): M = D(KS, C) = D(KS, E(KS, M)) 
