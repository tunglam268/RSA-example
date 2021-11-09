# RSA-example
test file
 ```go run main.go```

## Hệ mật mã khóa đối xứng
Là những hệ mật được sử dụng chung 1 khóa trong quá trình mã hóa và mã hóa. Do đó khóa phải được giữ bí mật tuyện đối.
Một số hệ mật mã khóa đối xứng hiện đại mà mình thấy hay được sử dụng là DES, AES, RC4, RC5,... 

### Hệ mật sẽ bao gồm:
- Bản rõ (plaintext-M): bản tin được sinh ra bởi bên gửi 
- Bản mật (ciphertext-C): bản tin che giấu thông tin của bản rõ, được gửi tới bên nhận qua một kênh không bí mật 
- Khóa (Ks): nó là giá trị ngẫu nhiên và bí mật được chia sẻ giữa các bên trao đổi thông tin và được tạo ra từ: 
  - Bên thứ 3 được tin cậy tạo và phân phối tới bên gửi và bên nhận 
  - Hoặc, bên gửi tạo ra và chuyển cho bên nhận 
- Mã hóa (encrypt-E): C = E(KS, M) 
- Giải mã (decrypt): M = D(KS, C) = D(KS, E(KS, M)) 

### Cơ chế hoạt động :
- Người gửi sử dụng khóa chung (Ks) để mã hóa thông tin rồi gửi cho nguời nhận. 
- Người nhận nhận được thông tin đó sẽ dùng chính khóa chung (Ks) để giải mã. 

## Hệ mật mã khóa bất đối xứng
Ở hệ mật này thay vì nguời dùng dùng chung 1 khóa như ở hệ mật mã khóa đối xứng thì ở đây sẽ dùng 1 cặp khóa có tên là public key và private key.
Hệ mật mã khóa bất đối xứng mình thấy được dùng nhiều nhất là RSA.

### Hệ mật sẽ bao gồm: 
- Bản rõ (plaintext-M): bản tin được sinh ra bởi bên gửi 
- Bản mật (ciphertext-C): bản tin che giấu thông tin của bản rõ, được gửi tới bên nhận qua một kênh không bí mật 
- Khóa: Bên nhận có 1 cặp khóa: 
  - Khóa công khai (Kub) : công bố cho tất cả mọi người biết (kể cả hacker) 
  - Khóa riêng (Krb) : bên nhận giữ bí mật, không chia sẻ cho bất kỳ ai 
- Mã hóa (encrypt-E): C = E(Kub, M) 
- Giải mã (decrypt): M = D(Krb, C) = D(Krb, E(Kub, M)) 

Yêu cầu đối với cặp khóa (Kub, Krb) là:
- Hoàn toàn ngẫu nhiên 
- Có quan hệ về mặt toán học 1-1. 
- Nếu chỉ có giá trị của Kub không thể tính được Krb. 
- Krb phải được giữ mật hoàn toàn. 

### Cơ chế hoạt động :
- Người gửi(A) gửi thông tin đã được mã hóa bằng khóa công khai (Kub) của người nhận(B) thông qua kênh truyền tin không bí mật 
- Người nhận(B) nhận được thông tin đó sẽ giải mã bằng khóa riêng (Krb) của mình. 
- Hacker cũng sẽ biết khóa công khai (Kub) của B tuy nhiên do không có khóa riêng (Krb) nên Hacker không thể xem được thông tin gửi 

## Mã hóa đối xứng vs. Mã hóa bất đối xứng
Mã hóa đối xứng là một trong 2 phương pháp chính trong mã hóa dữ liệu trên các hệ thống máy tính hiện đại. Phương pháp còn lại là mã hóa bất đối xứng hay còn gọi là mật mã khóa công khai. Trên thực tế, điểm khác biệt chính giữa 2 phương pháp này chính là việc các hệ thống bất đối xứng thường sử dụng 2 khóa khác nhau. Trong đó, 1 khóa sẽ được công khai chia sẻ (khóa công khai), khóa còn lại được giữ bí mật (khóa bí mật).
<br>
Việc sử dụng tới 2 loại khóa thay vì 1 cũng tạo ra nhiều điểm khác biệt về mặt tính năng giữa 2 loại mã hóa này. Các thuật toán bất đối xứng thường phức tạp và chậm hơn so với các thuật toán đối xứng. Bởi khóa công khai và khóa bí mật sử dụng trong mã hóa bất đối xứng thường có liên quan tới toán học về một mức độ nào đó, vì thế, các khóa này bản thân nó sẽ dài hơn đáng kể so với các khóa đối xứng khác để có thể đạt được mức độ bảo mật tương đương.

## Chữ ký số
Chữ ký số (còn được biết tới với cái tên Token USB) là loại chữ ký điện tử sử dụng để thay thế cho chữ ký thường bằng tay trên các thiết bị điện tử số, các văn bản và tài liệu số. Thông thường các tài liệu này được dùng để kê khai, nộp thuế qua mạng, khai hải quan điện tử và các giao dịch số khác.<br>
Thường chúng ta sẽ hiểu chữ ký số và chữ ký điện tử là một. Tuy nhiên không hẳn là vậy bởi khái niệm chữ ký điện tử rất rộng và chữ ký số chỉ là một loại phổ biến trong chữ ký điện tử.<br>
Chữ ký số bao gồm khóa bí mật và khóa công khai được mã hóa dữ liệu bao gồm các thông tin về doanh nghiệp và mã số thuế doanh nghiệp.

### Các khái niệm quan trọng trong chữ ký số: 
Có một số khái niệm khi tìm hiểu và sử dụng chữ ký số mà bạn cần biết, đó là:
- Khóa bí mật: Khóa được dùng để tạo chữ ký số; 
- Khóa công khai: Khóa được sử dụng để kiểm tra chữ ký số, được tạo bởi khóa bí mật tương ứng trong cặp khóa; 
- Ký số: Việc đưa khóa bí mật vào một chương trình phần mềm để tự động tạo và gắn chữ ký số vào thông điệp dữ liệu; 
- Người ký: thuê bao dùng đúng khóa bí mật của mình để ký số vào một thông điệp dữ liệu dưới tên của mình; 
- Người nhận: Các tổ chức, cá nhân nhận được thông điệp dữ liệu được ký số bởi người ký, sử dụng chứng thư số của người ký đó để kiểm tra chữ ký số trong thông điệp dữ liệu nhận được và tiến hành các hoạt động, giao dịch có liên quan; 

### Chữ ký số dùng để làm gì? 
- Thay thế cho chữ ký tay trong các giao dịch thương mại điện tử trên môi trường số, giúp hoạt động giao dịch có thể diễn ra nhanh chóng và tiết kiệm thời gian; 
- Với các cá nhân, chữ ký số có giá trị pháp lý tương đương chữ ký tay; 
- Với các tổ chức, doanh nghiệp, chữ ký số có giá trị tương đương con dấu và chữ ký của người đại diện pháp luật; 
- Chữ ký số giúp bạn ký trong các giao dịch thư điện tử, ký vào mail để xác nhận người gửi thư cho khách hàng; 
- Chữ ký số có thể giúp bạn đầu tư chứng khoán trực tiếp, mua hàng, thanh toán và chuyển tiền trực tiếp một cách bảo mật, an toàn; 
- Không cần phải in ấn các tờ kê khai hay đến cơ quan thuế để giải quyết khi thực hiện kê khai thuế trực tuyến hoặc thông quan trực tuyến; 
- Đóng bảo hiểm; 
- Ký hợp đồng điện tới với khách hàng trực tuyến thông qua hợp đồng điện tử; 
- Có thể sử dụng chữ ký số với các ứng dụng quản lý của doanh nghiệp với mức độ tin cậy, bảo mật và tính xác thực cao; 
