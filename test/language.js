/**
 * null--表示当前语言去除这个文案
 */
// 公共
var common = {
    INSA: {
        productName: 'SmartOption',
        copy: 'Salin',
        currencySymbol: 'Rp', //货币符号
        noData: 'Tidak ada data',
        copySuccess: 'Berhasil salin',
        requestErr0: 'Server memiliki kesalahan internal, silakan hubungi layanan pelanggan.',
        requestErr1: 'Jaringan tidak normal, silakan coba lagi nanti.',
        requestErr2: 'Pengecualian sistem.',
    },
    EN: {
        productName: 'SmartOption',
        copy: 'Copy',
        currencySymbol: '₹', //货币符号
        noData: 'No data',
        copySuccess: 'Copy success',
        requestErr0: 'The server generated an internal error, please contact the customer service.',
        requestErr1: 'Network anomalies, please try again later.',
        requestErr2: 'System exception.',
    },
    CH: {
        productName: 'SmartOption',
        copy: 'Copy',
        currencySymbol: '₹', //货币符号
        noData: 'No data',
        copySuccess: 'Copy success',
        requestErr0: 'The server generated an internal error, please contact the customer service.',
        requestErr1: 'Network anomalies, please try again later.',
        requestErr2: 'System exception.',
    },
    MX: {
        productName: 'SmartOption',
        copy: 'Copy',
        currencySymbol: '$', //货币符号
        noData: 'No data',
        copySuccess: 'Copy success',
        requestErr0: 'The server generated an internal error, please contact the customer service.',
        requestErr1: 'Network anomalies, please try again later.',
        requestErr2: 'System exception.',
    }
};
// 新闻网页
var newsPage = {
    INSA: {
        news0: {
            activeText1: 'Isi ulang',
            activeText2: 'Peringkat',
            activeText3: 'Bonus',
            // 新闻详情
            newsPhotoUrl: '../static/images/${LKEY}/news0.png',
            newsTitle: 'SmartOption - Trading Online Untuk Pemula',
            newsInformation: 'SmartOption adalah platform trading yang aman dan nyaman dengan aset pasar keuangan paling populer: mata uang euro / dolar, pound / dolar dan banyak lainnya, serta aset lainnya.',
            detailInformationObj0Title: 'SmartOption adalah platform trading yang aman dan nyaman dengan aset pasar keuangan paling populer: mata uang euro / dolar, pound / dolar dan banyak lainnya, serta aset lainnya.',
            detailInformationObj1Title: 'Fitur utama:',
            detailInformationObj2Title: 'Tutorial yang ramah untuk pemula',
            detailInformationObj2Text0: 'Mempermudah pemula untuk belajar trading dengan cepat. Setelah belajar, Anda akan mendapatkan hal-hal yang lebih menyenangkan dan diharapkan.',
            detailInformationObj3Title: 'Trading yang fleksibel',
            detailInformationObj3Text0: 'Kami menawarkan platform perdagangan 24 / 7 dengan beberapa aset dan beberapa aset transaksi di luar bursa bahkan mendukung perdagangan akhir pekan.',
            detailInformationObj4Title: 'Penarikan yang nyaman',
            detailInformationObj4Text0: 'Lebih dari 10 jenis kartu Bank, VA, QRIS, E-wallet telah didukung.',
            detailInformationObj5Title: '24/7 Layanan pelanggan',
            detailInformationObj5Text0: 'Tim dukungan teknis beroperasi sepanjang waktu. Gunakan layanan obrolan, email, atau panggilan suara untuk mengajukan pertanyaan dan mendapatkan jawaban cepat dari pakar kami. Kami peduli dengan setiap pelanggan.',
            detailInformationObj6Title: 'Eksekusi Transaksi Real Time',
            detailInformationObj6Text0: 'Kutipan secara real time dan eksekusi pesanan tercepat.',
            detailInformationObj7Img: '../static/images/${LKEY}/news0.png',
            detailInformationObj8Title: null, //印尼文去除文案
            detailInformationObj9Title: null, //印尼文去除文案
            creatTime: '25-7-2022'
        }
    },
    EN: {
        news0: {
            activeText1: 'Top up',
            activeText2: 'Ranking',
            activeText3: 'Bonus',
            newsPhotoUrl: '../static/images/${LKEY}/news0.png',
            newsTitle: 'SmartOption - Online Trading For Beginners',
            newsInformation: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj0Title: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj1Title: null, //英文去除这个文案
            detailInformationObj2Title: 'Friendly Tutorials For Beginners',
            detailInformationObj2Text0: 'It will be easy for beginners to learn trading quickly. After learning, you will get more fun and expected things.',
            detailInformationObj3Title: 'Flexible Trading',
            detailInformationObj3Text0: 'Bank card and any kinds of e-wallets are supported.',
            detailInformationObj4Title: 'Convenient withdrawal',
            detailInformationObj4Text0: 'Bank card and any one of more than 10 kinds of e-wallets are supported',
            detailInformationObj5Title: '24 / 7 Customer Service',
            detailInformationObj5Text0: 'The technical support team operates around the clock. Use chat services, email to ask your questions and get a quick answer from our experts. We care about every customer.',
            detailInformationObj6Title: 'Company: SmartOption SPACE EXCHANGE INC.',
            detailInformationObj6Text0: null, //英文去除这个文案
            detailInformationObj7Img: '../static/images/${LKEY}/news0.png',
            detailInformationObj8Title: 'Address:7514 BATH ROAD MISSISSAUGA, ON, CANADA L4T1L2.',
            detailInformationObj9Title: 'Email：sssmartoption@gmail.com.',
            creatTime: '25-7-2022'
        }
    },
    MX: {
        news0: {
            activeText1: 'Top up',
            activeText2: 'Ranking',
            activeText3: 'Bonus',
            newsPhotoUrl: '../static/images/${LKEY}/news0.png',
            newsTitle: 'SmartOption - Online Trading For Beginners',
            newsInformation: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj0Title: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj1Title: null, //英文去除这个文案
            detailInformationObj2Title: 'Friendly Tutorials For Beginners',
            detailInformationObj2Text0: 'It will be easy for beginners to learn trading quickly. After learning, you will get more fun and expected things.',
            detailInformationObj3Title: 'Flexible Trading',
            detailInformationObj3Text0: 'Bank card and any kinds of e-wallets are supported.',
            detailInformationObj4Title: 'Convenient withdrawal',
            detailInformationObj4Text0: 'Bank card and any one of more than 10 kinds of e-wallets are supported',
            detailInformationObj5Title: '24 / 7 Customer Service',
            detailInformationObj5Text0: 'The technical support team operates around the clock. Use chat services, email to ask your questions and get a quick answer from our experts. We care about every customer.',
            detailInformationObj6Title: 'Company: SmartOption SPACE EXCHANGE INC.',
            detailInformationObj6Text0: null, //英文去除这个文案
            detailInformationObj7Img: '../static/images/${LKEY}/news0.png',
            detailInformationObj8Title: 'Address:7514 BATH ROAD MISSISSAUGA, ON, CANADA L4T1L2.',
            detailInformationObj9Title: 'Email：sssmartoption@gmail.com.',
            creatTime: '25-7-2022'
        }
    }
};
// 新闻详情
var newsPageInfo = {
    INSA: {
        title: 'About',
        news0: {
            // 新闻详情
            newsPhotoUrl: '../static/images/${LKEY}/news0.png',
            newsTitle: 'SmartOption - Trading Online Untuk Pemula',
            newsInformation: 'SmartOption adalah platform trading yang aman dan nyaman dengan aset pasar keuangan paling populer: mata uang euro / dolar, pound / dolar dan banyak lainnya, serta aset lainnya.',
            detailInformationObj0Title: 'SmartOption adalah platform trading yang aman dan nyaman dengan aset pasar keuangan paling populer: mata uang euro / dolar, pound / dolar dan banyak lainnya, serta aset lainnya.',
            detailInformationObj1Title: 'Fitur utama:',
            detailInformationObj2Title: 'Tutorial yang ramah untuk pemula',
            detailInformationObj2Text0: 'Mempermudah pemula untuk belajar trading dengan cepat. Setelah belajar, Anda akan mendapatkan hal-hal yang lebih menyenangkan dan diharapkan.',
            detailInformationObj3Title: 'Trading yang fleksibel',
            detailInformationObj3Text0: 'Kami menawarkan platform perdagangan 24 / 7 dengan beberapa aset dan beberapa aset transaksi di luar bursa bahkan mendukung perdagangan akhir pekan.',
            detailInformationObj4Title: 'Penarikan yang nyaman',
            detailInformationObj4Text0: 'Lebih dari 10 jenis kartu Bank, VA, QRIS, E-wallet telah didukung.',
            detailInformationObj5Title: '24/7 Layanan pelanggan',
            detailInformationObj5Text0: 'Tim dukungan teknis beroperasi sepanjang waktu. Gunakan layanan obrolan, email, atau panggilan suara untuk mengajukan pertanyaan dan mendapatkan jawaban cepat dari pakar kami. Kami peduli dengan setiap pelanggan.',
            detailInformationObj6Title: 'Eksekusi Transaksi Real Time',
            detailInformationObj6Text0: 'Kutipan secara real time dan eksekusi pesanan tercepat.',
            detailInformationObj7Img: '../static/images/${LKEY}/news0.png',
            detailInformationObj8Title: null, //印尼文去除文案
            detailInformationObj9Title: null, //印尼文去除文案
            creatTime: '25-7-2022'
        },
        news1: {
            newsPhotoUrl: '../static/images/news1.jpg',
            newsTitle: 'Cara membuat rencana investasi',
            newsInformation: '',
            detailInformationObj0Title: 'Modal Anda harus bekerja untuk Anda. Ini adalah aturan umum bagi mereka yang menginginkan penghasilan tambahan. Tentu saja, Anda hanya dapat menghemat sebagian, tetapi ini tidak akan memungkinkan Anda untuk mendapatkan jumlah yang cukup besar dengan cepat. Karena itu, Anda harus mempertimbangkan investasi. Sebelum Anda mulai, buatlah rencana yang akan membantu Anda menentukan di mana, berapa banyak, dan berapa lama Anda harus menginvestasikan dana Anda.',
            detailInformationObj1Title: 'Dengan melewatkan langkah ini, Anda berisiko membuat keputusan yang salah, yang akan menyebabkan kerugian. Lagi pula, Anda tidak ingin menjadi salah satu dari 90% investor pemula yang gagal meningkatkan atau mendapatkan kembali investasi mereka.',
            detailInformationObj2Title: 'Mari temukan apa rencana investasi terbaik.',
            detailInformationObj3Title: '1. Tetapkan tujuan',
            detailInformationObj3Text0: 'Untuk mendapatkan penghasilan tambahan, Anda harus memiliki tujuan. Tujuan yang tepat akan menentukan berapa lama Anda harus berinvestasi dan jenis investasi apa yang memungkinkan Anda mencapainya. Apakah Anda ingin berinvestasi karena Anda memiliki daftar keinginan yang ketat dan ingin mencoret poin demi poin setiap posisi? Itu akan menjadi motivasi yang hebat.',
            detailInformationObj4Title: '2. Nilai seberapa banyak Anda dapat berinvestasi',
            detailInformationObj4Text0: 'Sebelum Anda mulai berinvestasi, Anda harus memahami berapa banyak yang dapat Anda belanjakan. Jumlahnya harus terjangkau, jadi Anda tidak menderita. Ada banyak aturan untuk membantu Anda memutuskan berapa banyak yang akan diinvestasikan. Yang paling bisa diterima adalah 30-30-40. Artinya, Anda dapat membelanjakan 30% dari penghasilan Anda untuk kebutuhan umum. Anda harus menghabiskan 30% lagi untuk hutang. Jika Anda tidak memiliki utang, bagilah jumlah ini, masukkan 15% untuk pengeluaran harian dan 15% untuk investasi.',
            detailInformationObj4Text1: 'Kemudian, tambahan 40% dari total pendapatan Anda harus digunakan untuk investasi. Saat menentukan jumlah untuk investasi, Anda harus mengingat masalah darurat. Oleh karena itu, belanjakan sebanyak 40%, agar Anda tetap merasa nyaman.',
            detailInformationObj5Title: '3. Tentukan berapa banyak waktu yang Anda miliki',
            detailInformationObj5Text0: 'Anda harus menentukan seberapa cepat Anda ingin mencapai tujuan tertentu. Ini akan membantu Anda memilih jenis rencana investasi yang Anda butuhkan. Misalnya, rencana investasi terbaik untuk pendapatan bulanan dapat mencakup deposito bank dengan tingkat bunga tertentu. Pendapatan akan tergantung pada tingkat bank dan jumlah yang Anda masukkan ke dalam akun Anda.',
            detailInformationObj6Title: 'Anda dapat mempertimbangkan perdagangan sebagai semacam pendekatan investasi. Jika Anda seorang trader aktif, Anda akan memiliki kesempatan untuk mendapatkan penghasilan tambahan setiap bulannya.',
            detailInformationObj7Img: '../static/images/news1.jpg',
            creatTime: '1-8-2022'
        },
        news2: {
            newsPhotoUrl: '../static/images/news2.jpg',
            newsTitle: 'Tips Mudah untuk Meningkatkan Manajemen Risiko Anda',
            newsInformation: '',
            detailInformationObj0Title: 'Apakah perdagangan berisiko? Ya, tentu saja, tetapi di bawah ini, kami telah mengumpulkan beberapa cara terbaik untuk mengurangi risiko Anda.',
            detailInformationObj1Title: 'Tentukan entri dan keluar Anda sebelum perdagangan',
            detailInformationObj1Text0: 'Pertama, langkah penting yang akan digunakan oleh setiap trader profesional adalah mengetahui dengan tepat di mana mereka ingin mengambil keuntungan atau menghentikan perdagangan. Ini akan membantu Anda menentukan risiko Anda: rasio imbalan (lebih lanjut tentang itu dalam satu detik) dan menghentikan Anda dari meyakinkan diri sendiri bahwa perdagangan yang kalah tiba-tiba akan berbalik dan menjadi menguntungkan. Hal ini jarang terjadi dan biasanya mengakibatkan kerugian yang signifikan.',
            detailInformationObj1Text1: 'Sebaliknya, ambil untung dan hentikan kerugian saat Anda membuka perdagangan; apalagi, Anda harus memiliki alasan yang dapat dibenarkan untuk target tersebut. Jangan pernah menggunakan target tetap (misalnya 50 pips). Menetapkan target secara sewenang-wenang adalah kesalahan pemula yang tidak pernah mengarah pada profitabilitas jangka panjang dan sering kali menghasilkan stop loss yang terlalu lebar atau memotong keuntungan sebelum langkah besar yang sebenarnya.',
            detailInformationObj2Title: 'Pahami risiko: rasio imbalan',
            detailInformationObj2Text0: 'Risiko: imbalan sangat penting untuk perdagangan manajemen risiko tingkat lanjut. Intinya, risk:reward mengacu pada rasio antara berapa banyak akun Anda yang Anda risikokan dan berapa banyak yang bisa Anda peroleh. Rasio 1:3 akan menunjukkan bahwa Anda mempertaruhkan $10 untuk menghasilkan $30, misalnya.',
            detailInformationObj2Text1: 'Winrate dan risiko: rasio imbalan terkait erat dan harus dipahami untuk berdagang dengan benar. Jika risiko rata-rata: imbalan Anda adalah 1:1, Anda hanya perlu memenangkan 51% dari perdagangan Anda untuk secara teknis menguntungkan; jika rata-rata 1:2, ini turun menjadi 33%, dan seterusnya. Ketika Anda memahami bagaimana keduanya terkait, Anda dapat memeriksa tingkat kemenangan Anda dan menentukan risiko minimum yang tepat: rasio imbalan yang Anda butuhkan untuk menjadi menguntungkan.',
            detailInformationObj3Title: 'Jaga agar risiko Anda tetap rendah',
            detailInformationObj3Text0: 'Nasihat yang sering diulang bagi mereka yang ingin memahami bagaimana melakukan manajemen risiko dalam perdagangan adalah menjaga jumlah keseluruhan saldo perdagangan Anda pada risiko rendah. Yang kami maksud dengan ini adalah jangan pernah mengambil risiko lebih dari 3% dari akun Anda, idealnya 1% atau 2%. Risiko konstan 1% berarti Anda membutuhkan 100 kerugian berturut-turut untuk menghapus saldo Anda sepenuhnya.',
            detailInformationObj4Img: '../static/images/news2.jpg',
            creatTime: '8-8-2022'
        },
        news3: {
            newsPhotoUrl: '../static/images/news3.png',
            newsTitle: 'Price Action Untuk Pemula',
            newsInformation: '',
            detailInformationObj0Title: 'Untuk trader baru, Price Action adalah cara terbaik untuk memulai jalur trading di apk SmartOption. Itu pendapat saya setelah belajar dan trading dengan ratusan metode yang berbeda. Semuanya belum membuahkan hasil yang baik, kecuali Price Action. Itu sebabnya saya menyarankan Anda untuk belajar tentang Price Action, bukan indikator.',
            detailInformationObj1Title: 'Saya tidak menulis ini untuk mengkritik penggunaan indikator sebagai salah atau tidak efektif. Indikator masih memiliki fitur yang sangat berguna yang tidak dapat digantikan oleh aksi harga.',
            detailInformationObj2Title: 'Namun, untuk pemula, indikator bukanlah tempat untuk memulai perjalanan trading mereka.Mereka akan membuat mereka rentan terhadap ilusi tentang kemampuan mereka, dan tentu saja, kerugian akan mengikuti.',
            detailInformationObj3Title: '1. Mengapa pemula harus memulai trading dengan Price Action?',
            detailInformationObj3Text0: 'Banyak trader sukses masih menggunakan indikator teknis dalam strategi trading mereka. Itu karena mereka telah berpengalaman cukup lama untuk memahami sepenuhnya sifat dari indikator yang mereka gunakan.',
            detailInformationObj3Text1: 'Dan mereka dengan mudah melihat kekurangan dari indikator teknikal sehingga tidak bisa terjebak oleh pasar. Sementara itu, pedagang baru tidak. Pemula dapat terus-menerus tertipu oleh indikator sampai mereka tidak punya uang lagi dan menyerah.Tindakan Harga, bila diteliti dengan benar, akan membantu pedagang baru memahami bagaimana pasar pada dasarnya bekerja. Karena Price Action hanya mengambil harga sebagai satu-satunya sumber input “bahan baku”.',
            detailInformationObj3Text2: 'Trader yang menggunakan Price Action tidak akan suka menangkap puncak dan palung atau melawan tren. Mereka senang memiliki dasar pengetahuan yang kuat tentang cara menganalisis struktur pasar. Selain itu, mereka tidak perlu mencari cawan suci atau indikator teknis utama.',
            detailInformationObj3Text3: null,
            detailInformationObj4Title: '2. Dimulai dengan Price Action akan menghemat waktu trader',
            detailInformationObj4Text0: 'Perjalanan untuk menjadi trader yang menguntungkan adalah perjalanan yang panjang. Untuk mempersingkat waktu itu, Anda perlu fokus mempelajari hal-hal yang paling penting dan perlu. Jangan buang waktu Anda berkeliaran di jejaring sosial mencari cawan suci dengan sia-sia.',
            detailInformationObj4Text1: 'Investor baru perlu fokus untuk memahami bagaimana pasar bekerja, dengan mengamati dan mencoba memahami tindakan selanjutnya. Price Action akan membantu kami melakukan hal itu. Saya yakin Anda akan segera melihat perilaku pasar pada tingkat fundamental lebih cepat daripada menemukan indikator yang Anda sukai.',
            detailInformationObj5Title: '3. Price Action mempromosikan pemikiran dan sikap yang baik dari seorang trader',
            detailInformationObj5Text0: 'Intinya, indikator membuat trader lebih fokus pada sinyal masuk. Sangat penting untuk memasukkan pesanan ketika ada sinyal jika mereka tidak ingin kehilangan kesempatan untuk menghasilkan uang. Ini adalah awal yang sangat buruk bagi trader baru. Mereka "diajari" oleh indikator kapan harus memasukkan pesanan.',
            detailInformationObj5Text1: 'Jika Anda baru dalam perdagangan, hal pertama yang harus dipelajari adalah kapan tidak memasuki perdagangan. Price Action melakukan hal itu.',
            detailInformationObj5Text2: 'Trader perlu mengamati, mempelajari, dan melihat apa yang terjadi di pasar. Kemudian jalankan order perdagangan mengikuti pergerakan harga saat ini untuk mendapatkan keuntungan. Tindakan apakah akan memasuki order atau tidak harus dilakukan setelah trader memahami cerita pasar. Price Action mengajarkan trader cara membaca cerita yang ingin disampaikan pasar, alih-alih memberikan sinyal yang tidak jelas.',
            detailInformationObj5Text3: 'Dengan kata lain, strategi Price Action menekankan kemampuan analitis dan kesabaran daripada sinyal Beli – Jual yang mengikuti pola tertentu.',
            detailInformationObj6Img: '../static/images/news3.png',
            creatTime: '15-8-2022'
        },
        news4: {
            newsPhotoUrl: '../static/images/news4.png',
            newsTitle: 'Langkah Menjadi Trader Di Pasar',
            detailInformationObj0Title: 'Untuk menjadi ahli di bidang apa pun, Anda perlu meluangkan waktu, tenaga, belajar, dan berlatih. Jika Anda ingin menjadi trader profesional dengan Price Action di SmartOption, itu juga membutuhkan banyak usaha dari diri Anda sendiri.',
            detailInformationObj1Title: 'Namun, seluruh proses latihan itu dapat diringkas dalam 3 langkah, yang akan saya tunjukkan dalam artikel ini.',
            detailInformationObj2Title: 'Artikel ini seperti langkah persiapan untuk perencanaan trading. Anda harus tahu langkah apa yang akan Anda lalui di masa depan dan apa setiap langkah yang harus dilakukan. Dari sana, Anda dapat dengan percaya diri menggunakan Price Action untuk berdagang di SmartOption.',
            detailInformationObj3Title: 'Langkah 1: Belajar menghargai keindahan Aksi Harga',
            detailInformationObj3Text0: 'Banyak pedagang mulai mempelajari Tindakan Harga dengan kesalahpahaman dan akhirnya frustrasi. Price Action hanyalah sebuah metode, bukan cawan suci. Masih memiliki kelemahan tertentu yang perlu diwaspadai pengguna, untuk membatasi dan meminimalkan risiko. Jika Anda memilih Price Action sebagai alat untuk menghasilkan keuntungan, Anda harus menghargai keindahannya selain kerugian tersebut.',
            detailInformationObj3Text1: 'Kekuatan Price Action berasal dari kesederhanaan. Satu fokus pada harga memberi kita informasi yang kita butuhkan untuk membuat keputusan perdagangan yang baik.',
            detailInformationObj3Text2: 'Pertama, Anda harus menentukan apakah Anda cocok untuk Price Action atau tidak.',
            detailInformationObj3Text3: 'Jika Anda menyukai kesederhanaan, maka pilih Price Action.',
            detailInformationObj3Text4: 'Jika Anda menyukai rumus, otomatisasi, statistik, angka, maka pilih indikator teknis.',
            detailInformationObj3Text5: 'Setelah itu ditentukan, Anda juga harus mengetahui kebenaran yang pahit.Hanya Price Action saja tidak cukup untuk membantu Anda menghasilkan keuntungan.Keuntungan datang dari modal dan manajemen emosional, bukan hanya analisis.',
            detailInformationObj4Title: 'Langkah 2: Mulailah mempelajari dasar-dasarnya sebelum berdagang dengan Price Action',
            detailInformationObj4Text0: 'Mari belajar dari tingkat dasar hingga tingkat lanjut secara metodis. Pertama, mulailah dengan pengetahuan yang perlu Anda ketahui di pasar seperti tren, analisis teknis.',
            detailInformationObj4Text1: 'Jangan mencoba memulai dengan hal- hal yang membutuhkan banyak pengalaman seperti pola Inside bar, Fakey, Engulfing, Pin bar… Atau menggunakannya untuk berdagang dengan keinginan menghasilkan uang, yang sangat berbahaya.Dengan dasar yang kuat dalam Price Action, Anda dapat memahami konsep yang lebih kompleks dengan mudah.Anda juga akan belajar bagaimana melihat logika dan alasan di balik setiap pola Price Action tanpa kesulitan mengingatnya.',
            detailInformationObj5Title: 'Langkah 3: Ganti indikator yang digunakan',
            detailInformationObj5Text0: 'Jika Anda telah berdagang dengan indikator sebelumnya, ini bisa menjadi sedikit sulit. Namun coba menghilangkan indikator yang “berlebihan” dan tidak berfungsi. Hapus untuk fokus melatih kemampuan membaca dan memahami Tindakan Harga dengan lebih cepat. Ingat kita tidak harus menyingkirkan barang-barang penting. Namun kami menghilangkan hal-hal yang tidak berguna yang tidak banyak berguna untuk analisis. Jika Anda memilih price action sebagai metode utama Anda, maka Anda perlu merampingkan grafik Anda dengan menghapus indikator yang menghalangi pandangan Anda dan mengurangi kemampuan Anda untuk melihat peluang terbaik.',
            detailInformationObj6Img: '../static/images/news4.png',
            creatTime: '22-8-2022'
        }
    },
    EN: {
        title: 'About',
        news0: {
            newsPhotoUrl: '../static/images/${LKEY}/news0.png',
            newsTitle: 'SmartOption - Online Trading For Beginners',
            newsInformation: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj0Title: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj1Title: null, //英文去除这个文案
            detailInformationObj2Title: 'Friendly Tutorials For Beginners',
            detailInformationObj2Text0: 'It will be easy for beginners to learn trading quickly. After learning, you will get more fun and expected things.',
            detailInformationObj3Title: 'Flexible Trading',
            detailInformationObj3Text0: 'Bank card and any kinds of e-wallets are supported.',
            detailInformationObj4Title: 'Convenient withdrawal',
            detailInformationObj4Text0: 'Bank card and any one of more than 10 kinds of e-wallets are supported',
            detailInformationObj5Title: '24 / 7 Customer Service',
            detailInformationObj5Text0: 'The technical support team operates around the clock. Use chat services, email to ask your questions and get a quick answer from our experts. We care about every customer.',
            detailInformationObj6Title: 'Company: SmartOption SPACE EXCHANGE INC.',
            detailInformationObj6Text0: null, //英文去除这个文案
            detailInformationObj7Img: '../static/images/${LKEY}/news0.png',
            detailInformationObj8Title: 'Address:7514 BATH ROAD MISSISSAUGA, ON, CANADA L4T1L2.',
            detailInformationObj9Title: 'Email：sssmartoption@gmail.com.',
            creatTime: '25-7-2022'
        },
        news1: {
            newsPhotoUrl: '../static/images/news1.jpg',
            newsTitle: 'How to create an investment plan',
            newsInformation: '',
            detailInformationObj0Title: 'Your capital should work for you. It’s a common rule for those who want an additional income. Of course, you can just save some, but this won’t allow you to earn a considerable sum quickly. Therefore, you should consider investments. Before you start, create a plan that will help you determine where, how much, and for how long you should invest your funds. ',
            detailInformationObj1Title: 'By skipping this step, you risk making the wrong decisions, which will lead to losses. After all, you don’t want to be among 90% of first-time investors who fail to increase or regain their investments.',
            detailInformationObj2Title: 'Let’s discover what the best investment plan is. ',
            detailInformationObj3Title: '1. Set goals',
            detailInformationObj3Text0: 'To earn additional income, you need to have a goal. A precise purpose will determine how long you should invest and what type of investment will allow you to reach it. Do you want to invest because you have a strict wishlist and want to cross out point by point every position? That would be a great motivation. ',
            detailInformationObj4Title: '2. Assess how much you can invest',
            detailInformationObj4Text0: 'Before you start investing, you should understand how much you can spend. The sum should be affordable, so you don’t suffer. There are many rules to help you decide how much to invest. The most acceptable one is 30-30-40. It means that you can spend 30% of your income on general needs. You should spend another 30% on debts. If you don’t have any debts, divide this sum, putting 15% toward daily spending and 15% toward investments. ',
            detailInformationObj4Text1: 'Then, an additional 40% of your total income should go to investments. When determining a sum for investing, you should remember emergency issues. Therefore, spend as much as 40%, so you still feel comfortable. ',
            detailInformationObj5Title: '3. Define how much time you have',
            detailInformationObj5Text0: 'You should determine how fast you want to reach a particular goal. This will help you choose what kind of investment plan you need. For instance, the best investment plan for monthly income may include bank deposits with a specific interest rate. The income will depend on the bank rate and the amount you put into your account. ',
            detailInformationObj6Title: 'You can consider trading as a kind of investment approach. If you are an active trader, you will have an opportunity to get additional income every month.',
            detailInformationObj7Img: '../static/images/news1.jpg',
            creatTime: '1-8-2022'
        },
        news2: {
            newsPhotoUrl: '../static/images/news2.jpg',
            newsTitle: 'Easy tips to improve your risk management',
            detailInformationObj0Title: 'Is trading a risk? Yes, sure, but below, we’ve collected several best ways you can reduce your risk.',
            detailInformationObj1Title: 'Define your entries and exits before the trade',
            detailInformationObj1Text0: 'Firstly, a critical step that every professional trader will employ is knowing exactly where they want to take profit or stop out of a trade. It will help you define your risk:reward ratio (more on that in a second) and stop you from convincing yourself that a losing trade will suddenly turn around and become profitable. This rarely happens and usually results in a significant loss.',
            detailInformationObj1Text1: 'Instead, take profits and stop losses when you open the trade; moreover, you should have a justifiable reason for those targets. Never use fixed targets (e.g. 50 pips). Arbitrarily setting targets is a rookie mistake that never leads to long-term profitability and often results in having a stop loss that’s unnecessarily wide or cutting profits before the real big move.',
            detailInformationObj2Title: 'Understand risk:reward ratios',
            detailInformationObj2Text0: 'Risk:reward is essential to advanced risk management trading. In essence, risk:reward refers to the ratio between how much of your account you risk and how much you could gain. A 1:3 ratio would indicate that you are risking $10 to make $30, for example.',
            detailInformationObj2Text1: 'Winrate and risk:reward ratios are inextricably linked and must be understood to trade correctly. If your average risk:reward is 1:1, you only need to win 51% of your trades to technically be profitable; if the average is 1:2, this drops to 33%, and so on. When you understand how the two are linked, you can examine your winrate and determine the exact minimum risk:reward ratio you need to be profitable.',
            detailInformationObj3Title: 'Keep your risk low',
            detailInformationObj3Text0: 'An often repeated piece of advice to those looking to understand how to do risk management in trading is to keep the amount of your overall trading balance at risk low. What we mean by this is never risk more than 3% of your account, ideally 1% or 2%. A constant 1% risk means you’d need 100 consecutive losses to wipe out your balance completely.',
            detailInformationObj4Img: '../static/images/news2.jpg',
            creatTime: '8-8-2022'
        },
        news3: {
            newsPhotoUrl: '../static/images/news3.png',
            newsTitle: 'Price Action For User',
            detailInformationObj0Title: 'For new traders, Price Action is the best way to start the trading path in SmartOption. That’s my opinion after learning and trading with hundreds of different methods. All of them have not yielded good results, except for Price Action. That’s why I advise you to learn about Price Action, not indicators.',
            detailInformationObj1Title: 'I do not write this to criticize the use of indicators as wrong or ineffective. Indicators still have extremely useful features that price action cannot replace.',
            detailInformationObj2Title: 'However, for beginners, indicators are not the place to start their trading journey. They will make them vulnerable to illusions about their abilities, and of course, losses will follow.',
            detailInformationObj3Title: 'Why should beginners start trading with Price Action?',
            detailInformationObj3Text0: 'Many successful traders are still using technical indicators in their trading strategies. It’s because they have experienced long enough to understand the full nature of the indicator they use.',
            detailInformationObj3Text1: 'And they easily see the shortcomings of technical indicators so they can’t be trapped by the market. Meanwhile, new traders do not. Beginners could constantly be deceived by the indicator until they have no money left and give up.',
            detailInformationObj3Text2: 'Price Action, when properly researched, will help new traders understand how the market basically works. Because Price Action only takes price as the only input “raw material” source.',
            detailInformationObj3Text3: 'Traders using Price Action will not like catching peaks and troughs or going against the trend. They love to have solid knowledge foundations on how to analyze the market structure. Plus, they don’t need to go looking for the holy grail or the ultimate technical indicator.',
            detailInformationObj4Title: 'Starting with Price Action will save traders time',
            detailInformationObj4Text0: 'The journey to becoming a profitable trader is a long one. To shorten that time, you need to focus on learning the most important and necessary things. Don’t waste your time wandering the social networks searching for the holy grail in vain.',
            detailInformationObj4Text1: 'New investors need to focus on understanding how the market works, by observing and trying to understand its next actions. Price Action will help us do just that. I’m sure you’ll quickly see the market’s behavior on a fundamental level faster than finding an indicator you like.',
            detailInformationObj5Title: 'Price Action promotes good thinking and attitude of a trader',
            detailInformationObj5Text0: 'In short, Price Action will help traders have more realistic expectations.Traders need to observe, learn, and see what’s going on in the market. Then execute a trade order following the current price movement to make a profit. The action whether to enter the order or not will have to take place after the trader understands the market story. Price Action teaches traders how to read the story the market wants to tell, instead of giving vague signals.',
            detailInformationObj5Text1: 'In other words, the Price Action strategy emphasizes analytical ability and patience instead of Buy – Sell signals following a certain pattern.',
            detailInformationObj5Text2: null,
            detailInformationObj5Text3: null,
            detailInformationObj6Img: '../static/images/news3.png',
            creatTime: '15-8-2022',
        },
        news4: {
            newsPhotoUrl: '../static/images/news4.png',
            newsTitle: 'Steps To Becoming Traders For Market',
            detailInformationObj0Title: 'To become an expert in any field, you need to put in time, effort, study and practice. If you want to become a professional trader with the Price Action in SmartOption, it also takes a lot of effort from yourself.',
            detailInformationObj1Title: 'However, the whole process of that practice can be summed up in 3 steps, which I will show you in this article.',
            detailInformationObj2Title: 'This article is like a preparatory step for trading planning. You must know what steps you will go through in the future and what each step must be done. From there, you can confidently use Price Action to trade in SmartOption.',
            detailInformationObj3Title: 'Step 1: Learn to appreciate the beauty of Price Action',
            detailInformationObj3Text0: 'Many traders start learning Price Action with misconceptions and end up frustrated. Price Action is just a method, not the holy grail. It still has certain disadvantages that users need to be aware of, in order to limit and minimize risks. If you choose Price Action as a tool to make profits, you must appreciate its beauty besides those disadvantages.',
            detailInformationObj3Text1: 'Price Action’s power comes from simplicity. A single focus on price gives us the information we need to make sound trading decisions.',
            detailInformationObj3Text2: 'First, you must determine if you are suitable for Price Action or not.',
            detailInformationObj3Text3: 'If you like simplicity, then choose SmartOption.',
            detailInformationObj3Text4: 'If you like formulas, automation, statistics, numbers, then choose technical indicators.',
            detailInformationObj3Text5: 'Once that is determined, you must also know the harsh truth.Only Price Action isn’t enough to help you make profits.Profits come from capital and emotional management, not just analysis.',
            detailInformationObj4Title: 'Step 2: Start learning the basics before trading with Price Action',
            detailInformationObj4Text0: 'Let’s learn from the basic stage to the advanced in a methodical way. First, start with the knowledge you need to know in the market such as trends, technical analysis.',
            detailInformationObj4Text1: 'Do not try to start with things that require a lot of experience such as Inside bar, Fakey, Engulfing, Pin bar patterns… Or use them to trade with the desire to make money, which is very dangerous.With a solid foundation in Price Action, you can understand more complex concepts with ease. You will also learn how to see the logic and reason behind each Price Action pattern without having trouble remembering them. ',
            detailInformationObj5Title: 'Step 3: Replace the indicators in use',
            detailInformationObj5Text0: 'If you have traded with the indicator before, this can be a bit difficult. But try to remove the indicators that are “redundant” and do not work. Remove to focus on training the ability to read and understand Price Action more quickly. Remember we don’t have to get rid of the important stuffs. But we are removing the superfluous things that do not bring much use for the analysis. If you choose Price Action as your main method, then you need to streamline your chart by removing the indicators that block your view and reduce your ability to spot the best opportunities.',
            detailInformationObj6Img: '../static/images/news4.png',
            creatTime: '22-8-2022'
        }
    },
    CH: {
        title: 'About',
        news0: {
            newsPhotoUrl: '../static/images/${LKEY}/news0.png',
            newsTitle: 'SmartOption - Online Trading For Beginners',
            newsInformation: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj0Title: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj1Title: null, //英文去除这个文案
            detailInformationObj2Title: 'Friendly Tutorials For Beginners',
            detailInformationObj2Text0: 'It will be easy for beginners to learn trading quickly. After learning, you will get more fun and expected things.',
            detailInformationObj3Title: 'Flexible Trading',
            detailInformationObj3Text0: 'Bank card and any kinds of e-wallets are supported.',
            detailInformationObj4Title: 'Convenient withdrawal',
            detailInformationObj4Text0: 'Bank card and any one of more than 10 kinds of e-wallets are supported',
            detailInformationObj5Title: '24 / 7 Customer Service',
            detailInformationObj5Text0: 'The technical support team operates around the clock. Use chat services, email to ask your questions and get a quick answer from our experts. We care about every customer.',
            detailInformationObj6Title: 'Company: SmartOption SPACE EXCHANGE INC.',
            detailInformationObj6Text0: null, //英文去除这个文案
            detailInformationObj7Img: '../static/images/${LKEY}/news0.png',
            detailInformationObj8Title: 'Address:7514 BATH ROAD MISSISSAUGA, ON, CANADA L4T1L2.',
            detailInformationObj9Title: 'Email：sssmartoption@gmail.com.',
            creatTime: '25-7-2022'
        },
        news1: {
            newsPhotoUrl: '../static/images/news1.jpg',
            newsTitle: 'How to create an investment plan',
            newsInformation: '',
            detailInformationObj0Title: 'Your capital should work for you. It’s a common rule for those who want an additional income. Of course, you can just save some, but this won’t allow you to earn a considerable sum quickly. Therefore, you should consider investments. Before you start, create a plan that will help you determine where, how much, and for how long you should invest your funds. ',
            detailInformationObj1Title: 'By skipping this step, you risk making the wrong decisions, which will lead to losses. After all, you don’t want to be among 90% of first-time investors who fail to increase or regain their investments.',
            detailInformationObj2Title: 'Let’s discover what the best investment plan is. ',
            detailInformationObj3Title: '1. Set goals',
            detailInformationObj3Text0: 'To earn additional income, you need to have a goal. A precise purpose will determine how long you should invest and what type of investment will allow you to reach it. Do you want to invest because you have a strict wishlist and want to cross out point by point every position? That would be a great motivation. ',
            detailInformationObj4Title: '2. Assess how much you can invest',
            detailInformationObj4Text0: 'Before you start investing, you should understand how much you can spend. The sum should be affordable, so you don’t suffer. There are many rules to help you decide how much to invest. The most acceptable one is 30-30-40. It means that you can spend 30% of your income on general needs. You should spend another 30% on debts. If you don’t have any debts, divide this sum, putting 15% toward daily spending and 15% toward investments. ',
            detailInformationObj4Text1: 'Then, an additional 40% of your total income should go to investments. When determining a sum for investing, you should remember emergency issues. Therefore, spend as much as 40%, so you still feel comfortable. ',
            detailInformationObj5Title: '3. Define how much time you have',
            detailInformationObj5Text0: 'You should determine how fast you want to reach a particular goal. This will help you choose what kind of investment plan you need. For instance, the best investment plan for monthly income may include bank deposits with a specific interest rate. The income will depend on the bank rate and the amount you put into your account. ',
            detailInformationObj6Title: 'You can consider trading as a kind of investment approach. If you are an active trader, you will have an opportunity to get additional income every month.',
            detailInformationObj7Img: '../static/images/news1.jpg',
            creatTime: '1-8-2022'
        },
        news2: {
            newsPhotoUrl: '../static/images/news2.jpg',
            newsTitle: 'Easy tips to improve your risk management',
            detailInformationObj0Title: 'Is trading a risk? Yes, sure, but below, we’ve collected several best ways you can reduce your risk.',
            detailInformationObj1Title: 'Define your entries and exits before the trade',
            detailInformationObj1Text0: 'Firstly, a critical step that every professional trader will employ is knowing exactly where they want to take profit or stop out of a trade. It will help you define your risk:reward ratio (more on that in a second) and stop you from convincing yourself that a losing trade will suddenly turn around and become profitable. This rarely happens and usually results in a significant loss.',
            detailInformationObj1Text1: 'Instead, take profits and stop losses when you open the trade; moreover, you should have a justifiable reason for those targets. Never use fixed targets (e.g. 50 pips). Arbitrarily setting targets is a rookie mistake that never leads to long-term profitability and often results in having a stop loss that’s unnecessarily wide or cutting profits before the real big move.',
            detailInformationObj2Title: 'Understand risk:reward ratios',
            detailInformationObj2Text0: 'Risk:reward is essential to advanced risk management trading. In essence, risk:reward refers to the ratio between how much of your account you risk and how much you could gain. A 1:3 ratio would indicate that you are risking $10 to make $30, for example.',
            detailInformationObj2Text1: 'Winrate and risk:reward ratios are inextricably linked and must be understood to trade correctly. If your average risk:reward is 1:1, you only need to win 51% of your trades to technically be profitable; if the average is 1:2, this drops to 33%, and so on. When you understand how the two are linked, you can examine your winrate and determine the exact minimum risk:reward ratio you need to be profitable.',
            detailInformationObj3Title: 'Keep your risk low',
            detailInformationObj3Text0: 'An often repeated piece of advice to those looking to understand how to do risk management in trading is to keep the amount of your overall trading balance at risk low. What we mean by this is never risk more than 3% of your account, ideally 1% or 2%. A constant 1% risk means you’d need 100 consecutive losses to wipe out your balance completely.',
            detailInformationObj4Img: '../static/images/news2.jpg',
            creatTime: '8-8-2022'
        },
        news3: {
            newsPhotoUrl: '../static/images/news3.png',
            newsTitle: 'Price Action For User',
            detailInformationObj0Title: 'For new traders, Price Action is the best way to start the trading path in SmartOption. That’s my opinion after learning and trading with hundreds of different methods. All of them have not yielded good results, except for Price Action. That’s why I advise you to learn about Price Action, not indicators.',
            detailInformationObj1Title: 'I do not write this to criticize the use of indicators as wrong or ineffective. Indicators still have extremely useful features that price action cannot replace.',
            detailInformationObj2Title: 'However, for beginners, indicators are not the place to start their trading journey. They will make them vulnerable to illusions about their abilities, and of course, losses will follow.',
            detailInformationObj3Title: 'Why should beginners start trading with Price Action?',
            detailInformationObj3Text0: 'Many successful traders are still using technical indicators in their trading strategies. It’s because they have experienced long enough to understand the full nature of the indicator they use.',
            detailInformationObj3Text1: 'And they easily see the shortcomings of technical indicators so they can’t be trapped by the market. Meanwhile, new traders do not. Beginners could constantly be deceived by the indicator until they have no money left and give up.',
            detailInformationObj3Text2: 'Price Action, when properly researched, will help new traders understand how the market basically works. Because Price Action only takes price as the only input “raw material” source.',
            detailInformationObj3Text3: 'Traders using Price Action will not like catching peaks and troughs or going against the trend. They love to have solid knowledge foundations on how to analyze the market structure. Plus, they don’t need to go looking for the holy grail or the ultimate technical indicator.',
            detailInformationObj4Title: 'Starting with Price Action will save traders time',
            detailInformationObj4Text0: 'The journey to becoming a profitable trader is a long one. To shorten that time, you need to focus on learning the most important and necessary things. Don’t waste your time wandering the social networks searching for the holy grail in vain.',
            detailInformationObj4Text1: 'New investors need to focus on understanding how the market works, by observing and trying to understand its next actions. Price Action will help us do just that. I’m sure you’ll quickly see the market’s behavior on a fundamental level faster than finding an indicator you like.',
            detailInformationObj5Title: 'Price Action promotes good thinking and attitude of a trader',
            detailInformationObj5Text0: 'In short, Price Action will help traders have more realistic expectations.Traders need to observe, learn, and see what’s going on in the market. Then execute a trade order following the current price movement to make a profit. The action whether to enter the order or not will have to take place after the trader understands the market story. Price Action teaches traders how to read the story the market wants to tell, instead of giving vague signals.',
            detailInformationObj5Text1: 'In other words, the Price Action strategy emphasizes analytical ability and patience instead of Buy – Sell signals following a certain pattern.',
            detailInformationObj5Text2: null,
            detailInformationObj5Text3: null,
            detailInformationObj6Img: '../static/images/news3.png',
            creatTime: '15-8-2022',
        },
        news4: {
            newsPhotoUrl: '../static/images/news4.png',
            newsTitle: 'Steps To Becoming Traders For Market',
            detailInformationObj0Title: 'To become an expert in any field, you need to put in time, effort, study and practice. If you want to become a professional trader with the Price Action in SmartOption, it also takes a lot of effort from yourself.',
            detailInformationObj1Title: 'However, the whole process of that practice can be summed up in 3 steps, which I will show you in this article.',
            detailInformationObj2Title: 'This article is like a preparatory step for trading planning. You must know what steps you will go through in the future and what each step must be done. From there, you can confidently use Price Action to trade in SmartOption.',
            detailInformationObj3Title: 'Step 1: Learn to appreciate the beauty of Price Action',
            detailInformationObj3Text0: 'Many traders start learning Price Action with misconceptions and end up frustrated. Price Action is just a method, not the holy grail. It still has certain disadvantages that users need to be aware of, in order to limit and minimize risks. If you choose Price Action as a tool to make profits, you must appreciate its beauty besides those disadvantages.',
            detailInformationObj3Text1: 'Price Action’s power comes from simplicity. A single focus on price gives us the information we need to make sound trading decisions.',
            detailInformationObj3Text2: 'First, you must determine if you are suitable for Price Action or not.',
            detailInformationObj3Text3: 'If you like simplicity, then choose SmartOption.',
            detailInformationObj3Text4: 'If you like formulas, automation, statistics, numbers, then choose technical indicators.',
            detailInformationObj3Text5: 'Once that is determined, you must also know the harsh truth.Only Price Action isn’t enough to help you make profits.Profits come from capital and emotional management, not just analysis.',
            detailInformationObj4Title: 'Step 2: Start learning the basics before trading with Price Action',
            detailInformationObj4Text0: 'Let’s learn from the basic stage to the advanced in a methodical way. First, start with the knowledge you need to know in the market such as trends, technical analysis.',
            detailInformationObj4Text1: 'Do not try to start with things that require a lot of experience such as Inside bar, Fakey, Engulfing, Pin bar patterns… Or use them to trade with the desire to make money, which is very dangerous.With a solid foundation in Price Action, you can understand more complex concepts with ease. You will also learn how to see the logic and reason behind each Price Action pattern without having trouble remembering them. ',
            detailInformationObj5Title: 'Step 3: Replace the indicators in use',
            detailInformationObj5Text0: 'If you have traded with the indicator before, this can be a bit difficult. But try to remove the indicators that are “redundant” and do not work. Remove to focus on training the ability to read and understand Price Action more quickly. Remember we don’t have to get rid of the important stuffs. But we are removing the superfluous things that do not bring much use for the analysis. If you choose Price Action as your main method, then you need to streamline your chart by removing the indicators that block your view and reduce your ability to spot the best opportunities.',
            detailInformationObj6Img: '../static/images/news4.png',
            creatTime: '22-8-2022'
        }
    },
    MX: {
        title: 'About',
        news0: {
            newsPhotoUrl: '../static/images/${LKEY}/news0.png',
            newsTitle: 'SmartOption - Online Trading For Beginners',
            newsInformation: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj0Title: 'SmartOption is a safe and convenient trading platform with the most popular financial market assets: currencies EUR/USD, GBP/USD and many others, as well as other assets.',
            detailInformationObj1Title: null, //英文去除这个文案
            detailInformationObj2Title: 'Friendly Tutorials For Beginners',
            detailInformationObj2Text0: 'It will be easy for beginners to learn trading quickly. After learning, you will get more fun and expected things.',
            detailInformationObj3Title: 'Flexible Trading',
            detailInformationObj3Text0: 'Bank card and any kinds of e-wallets are supported.',
            detailInformationObj4Title: 'Convenient withdrawal',
            detailInformationObj4Text0: 'Bank card and any one of more than 10 kinds of e-wallets are supported',
            detailInformationObj5Title: '24 / 7 Customer Service',
            detailInformationObj5Text0: 'The technical support team operates around the clock. Use chat services, email to ask your questions and get a quick answer from our experts. We care about every customer.',
            detailInformationObj6Title: 'Company: SmartOption SPACE EXCHANGE INC.',
            detailInformationObj6Text0: null, //英文去除这个文案
            detailInformationObj7Img: '../static/images/${LKEY}/news0.png',
            detailInformationObj8Title: 'Address:7514 BATH ROAD MISSISSAUGA, ON, CANADA L4T1L2.',
            detailInformationObj9Title: 'Email：sssmartoption@gmail.com.',
            creatTime: '25-7-2022'
        },
        news1: {
            newsPhotoUrl: '../static/images/news1.jpg',
            newsTitle: 'How to create an investment plan',
            newsInformation: '',
            detailInformationObj0Title: 'Your capital should work for you. It’s a common rule for those who want an additional income. Of course, you can just save some, but this won’t allow you to earn a considerable sum quickly. Therefore, you should consider investments. Before you start, create a plan that will help you determine where, how much, and for how long you should invest your funds. ',
            detailInformationObj1Title: 'By skipping this step, you risk making the wrong decisions, which will lead to losses. After all, you don’t want to be among 90% of first-time investors who fail to increase or regain their investments.',
            detailInformationObj2Title: 'Let’s discover what the best investment plan is. ',
            detailInformationObj3Title: '1. Set goals',
            detailInformationObj3Text0: 'To earn additional income, you need to have a goal. A precise purpose will determine how long you should invest and what type of investment will allow you to reach it. Do you want to invest because you have a strict wishlist and want to cross out point by point every position? That would be a great motivation. ',
            detailInformationObj4Title: '2. Assess how much you can invest',
            detailInformationObj4Text0: 'Before you start investing, you should understand how much you can spend. The sum should be affordable, so you don’t suffer. There are many rules to help you decide how much to invest. The most acceptable one is 30-30-40. It means that you can spend 30% of your income on general needs. You should spend another 30% on debts. If you don’t have any debts, divide this sum, putting 15% toward daily spending and 15% toward investments. ',
            detailInformationObj4Text1: 'Then, an additional 40% of your total income should go to investments. When determining a sum for investing, you should remember emergency issues. Therefore, spend as much as 40%, so you still feel comfortable. ',
            detailInformationObj5Title: '3. Define how much time you have',
            detailInformationObj5Text0: 'You should determine how fast you want to reach a particular goal. This will help you choose what kind of investment plan you need. For instance, the best investment plan for monthly income may include bank deposits with a specific interest rate. The income will depend on the bank rate and the amount you put into your account. ',
            detailInformationObj6Title: 'You can consider trading as a kind of investment approach. If you are an active trader, you will have an opportunity to get additional income every month.',
            detailInformationObj7Img: '../static/images/news1.jpg',
            creatTime: '1-8-2022'
        },
        news2: {
            newsPhotoUrl: '../static/images/news2.jpg',
            newsTitle: 'Easy tips to improve your risk management',
            detailInformationObj0Title: 'Is trading a risk? Yes, sure, but below, we’ve collected several best ways you can reduce your risk.',
            detailInformationObj1Title: 'Define your entries and exits before the trade',
            detailInformationObj1Text0: 'Firstly, a critical step that every professional trader will employ is knowing exactly where they want to take profit or stop out of a trade. It will help you define your risk:reward ratio (more on that in a second) and stop you from convincing yourself that a losing trade will suddenly turn around and become profitable. This rarely happens and usually results in a significant loss.',
            detailInformationObj1Text1: 'Instead, take profits and stop losses when you open the trade; moreover, you should have a justifiable reason for those targets. Never use fixed targets (e.g. 50 pips). Arbitrarily setting targets is a rookie mistake that never leads to long-term profitability and often results in having a stop loss that’s unnecessarily wide or cutting profits before the real big move.',
            detailInformationObj2Title: 'Understand risk:reward ratios',
            detailInformationObj2Text0: 'Risk:reward is essential to advanced risk management trading. In essence, risk:reward refers to the ratio between how much of your account you risk and how much you could gain. A 1:3 ratio would indicate that you are risking $10 to make $30, for example.',
            detailInformationObj2Text1: 'Winrate and risk:reward ratios are inextricably linked and must be understood to trade correctly. If your average risk:reward is 1:1, you only need to win 51% of your trades to technically be profitable; if the average is 1:2, this drops to 33%, and so on. When you understand how the two are linked, you can examine your winrate and determine the exact minimum risk:reward ratio you need to be profitable.',
            detailInformationObj3Title: 'Keep your risk low',
            detailInformationObj3Text0: 'An often repeated piece of advice to those looking to understand how to do risk management in trading is to keep the amount of your overall trading balance at risk low. What we mean by this is never risk more than 3% of your account, ideally 1% or 2%. A constant 1% risk means you’d need 100 consecutive losses to wipe out your balance completely.',
            detailInformationObj4Img: '../static/images/news2.jpg',
            creatTime: '8-8-2022'
        },
        news3: {
            newsPhotoUrl: '../static/images/news3.png',
            newsTitle: 'Price Action For User',
            detailInformationObj0Title: 'For new traders, Price Action is the best way to start the trading path in SmartOption. That’s my opinion after learning and trading with hundreds of different methods. All of them have not yielded good results, except for Price Action. That’s why I advise you to learn about Price Action, not indicators.',
            detailInformationObj1Title: 'I do not write this to criticize the use of indicators as wrong or ineffective. Indicators still have extremely useful features that price action cannot replace.',
            detailInformationObj2Title: 'However, for beginners, indicators are not the place to start their trading journey. They will make them vulnerable to illusions about their abilities, and of course, losses will follow.',
            detailInformationObj3Title: 'Why should beginners start trading with Price Action?',
            detailInformationObj3Text0: 'Many successful traders are still using technical indicators in their trading strategies. It’s because they have experienced long enough to understand the full nature of the indicator they use.',
            detailInformationObj3Text1: 'And they easily see the shortcomings of technical indicators so they can’t be trapped by the market. Meanwhile, new traders do not. Beginners could constantly be deceived by the indicator until they have no money left and give up.',
            detailInformationObj3Text2: 'Price Action, when properly researched, will help new traders understand how the market basically works. Because Price Action only takes price as the only input “raw material” source.',
            detailInformationObj3Text3: 'Traders using Price Action will not like catching peaks and troughs or going against the trend. They love to have solid knowledge foundations on how to analyze the market structure. Plus, they don’t need to go looking for the holy grail or the ultimate technical indicator.',
            detailInformationObj4Title: 'Starting with Price Action will save traders time',
            detailInformationObj4Text0: 'The journey to becoming a profitable trader is a long one. To shorten that time, you need to focus on learning the most important and necessary things. Don’t waste your time wandering the social networks searching for the holy grail in vain.',
            detailInformationObj4Text1: 'New investors need to focus on understanding how the market works, by observing and trying to understand its next actions. Price Action will help us do just that. I’m sure you’ll quickly see the market’s behavior on a fundamental level faster than finding an indicator you like.',
            detailInformationObj5Title: 'Price Action promotes good thinking and attitude of a trader',
            detailInformationObj5Text0: 'In short, Price Action will help traders have more realistic expectations.Traders need to observe, learn, and see what’s going on in the market. Then execute a trade order following the current price movement to make a profit. The action whether to enter the order or not will have to take place after the trader understands the market story. Price Action teaches traders how to read the story the market wants to tell, instead of giving vague signals.',
            detailInformationObj5Text1: 'In other words, the Price Action strategy emphasizes analytical ability and patience instead of Buy – Sell signals following a certain pattern.',
            detailInformationObj5Text2: null,
            detailInformationObj5Text3: null,
            detailInformationObj6Img: '../static/images/news3.png',
            creatTime: '15-8-2022',
        },
        news4: {
            newsPhotoUrl: '../static/images/news4.png',
            newsTitle: 'Steps To Becoming Traders For Market',
            detailInformationObj0Title: 'To become an expert in any field, you need to put in time, effort, study and practice. If you want to become a professional trader with the Price Action in SmartOption, it also takes a lot of effort from yourself.',
            detailInformationObj1Title: 'However, the whole process of that practice can be summed up in 3 steps, which I will show you in this article.',
            detailInformationObj2Title: 'This article is like a preparatory step for trading planning. You must know what steps you will go through in the future and what each step must be done. From there, you can confidently use Price Action to trade in SmartOption.',
            detailInformationObj3Title: 'Step 1: Learn to appreciate the beauty of Price Action',
            detailInformationObj3Text0: 'Many traders start learning Price Action with misconceptions and end up frustrated. Price Action is just a method, not the holy grail. It still has certain disadvantages that users need to be aware of, in order to limit and minimize risks. If you choose Price Action as a tool to make profits, you must appreciate its beauty besides those disadvantages.',
            detailInformationObj3Text1: 'Price Action’s power comes from simplicity. A single focus on price gives us the information we need to make sound trading decisions.',
            detailInformationObj3Text2: 'First, you must determine if you are suitable for Price Action or not.',
            detailInformationObj3Text3: 'If you like simplicity, then choose SmartOption.',
            detailInformationObj3Text4: 'If you like formulas, automation, statistics, numbers, then choose technical indicators.',
            detailInformationObj3Text5: 'Once that is determined, you must also know the harsh truth.Only Price Action isn’t enough to help you make profits.Profits come from capital and emotional management, not just analysis.',
            detailInformationObj4Title: 'Step 2: Start learning the basics before trading with Price Action',
            detailInformationObj4Text0: 'Let’s learn from the basic stage to the advanced in a methodical way. First, start with the knowledge you need to know in the market such as trends, technical analysis.',
            detailInformationObj4Text1: 'Do not try to start with things that require a lot of experience such as Inside bar, Fakey, Engulfing, Pin bar patterns… Or use them to trade with the desire to make money, which is very dangerous.With a solid foundation in Price Action, you can understand more complex concepts with ease. You will also learn how to see the logic and reason behind each Price Action pattern without having trouble remembering them. ',
            detailInformationObj5Title: 'Step 3: Replace the indicators in use',
            detailInformationObj5Text0: 'If you have traded with the indicator before, this can be a bit difficult. But try to remove the indicators that are “redundant” and do not work. Remove to focus on training the ability to read and understand Price Action more quickly. Remember we don’t have to get rid of the important stuffs. But we are removing the superfluous things that do not bring much use for the analysis. If you choose Price Action as your main method, then you need to streamline your chart by removing the indicators that block your view and reduce your ability to spot the best opportunities.',
            detailInformationObj6Img: '../static/images/news4.png',
            creatTime: '22-8-2022'
        }
    }
};
// Bonus活动
var bonusActivities = {
    INSA: {
        title: 'Bonus',
        act_tips0: '1. Bonus adalah bonus yang disediakan secara gratis dari platform, bisa digunakan untuk trading di akun nyata, keuntungan bisa ditarik.',
        act_tips1: '2. Riwayat pencairan bisa dilihat di Bonus Saya.',
        act_tips2: 'Banyak Event Bonus silakan terus ikuti halaman ini.',
        confire: 'Saya paham',
    },
    EN: {
        title: 'Bonus',
        act_tips0: '1.Bonus are gifts of our platform, It can be used as free trading money, you can withdraw the bonus profit as well.',
        act_tips1: '2.The release record can be viewed in [Me]-[Bonus].',
        act_tips2: 'For more Bonus activities, please continue to pay attention to this page.',
        confire: 'Got It',
    },
    CH: {
        title: 'Bonus',
        act_tips0: '1.Bonus are gifts of our platform, It can be used as free trading money, you can withdraw the bonus profit as well.',
        act_tips1: '2.The release record can be viewed in [Me]-[Bonus].',
        act_tips2: 'For more Bonus activities, please continue to pay attention to this page.',
        confire: 'Got It',
    },
    MX: {
        title: 'Bonus',
        act_tips0: '1.Bonus are gifts of our platform, It can be used as free trading money, you can withdraw the bonus profit as well.',
        act_tips1: '2.The release record can be viewed in [Me]-[Bonus].',
        act_tips2: 'For more Bonus activities, please continue to pay attention to this page.',
        confire: 'Got It',
    }
};
// 邀请好友
var inviteFriends = {
    INSA: {
        title: 'Undang teman',
        bannerTitleH2: 'UNDANG TEMAN <br/> UNTUK MENGHASILKAN UANG',
        bannerTitleH3: 'ANDA MENDAPATKAN CASHBACK DAN BONUS TANPA BATAS',
        inviteTitle: 'Undang teman',
        inviteName: 'Kode undangan',
        shareText: 'Bagikan Google play',
        invitedUsers: 'Pengundang',
        toBeRecharged: 'Belum isi ulang',
        toBetraded: 'Belum trading',
        trading: 'Trading',
        times: 'kali',
        invitationSucceeded: 'Berhasil undang',
        rules: 'Aturan',
        inviteReward: '20.000',
        rules1: '1.  Anda bisa mendapatkan bonus sebesar <span ref="currencySymbol" class="currencySymbol"></span><span ref="inviteReward" class="inviteReward"></span> jika pengguna yang anda undang melakukan pemesanan 20 kali di akun nyatanya',
        rules2: '2. Tidak ada batasan jumlah orang, lebih banyak undangan, lebih banyak hadiah',
        rules3: '3. Waktu pencairan bonus: jam 08:00 keesokan harinya setelah memenuhi persyaratan',
        rules4: '4. Kode undangan hanya berlaku ketika dimasukkan waktu daftar',
        contactService: 'Hubungi customer service jika mempunyai pertanyaan',
    },
    EN: {
        title: 'Invite friends',
        bannerTitleH2: 'INVITE YOUR FRIENDS',
        bannerTitleH3: 'GET CASHBACKS AND UNLIMITED BONUS',
        inviteTitle: 'Invite friends',
        inviteName: 'Referral code',
        shareText: 'Share Google Play Link',
        invitedUsers: 'Invited Users',
        toBeRecharged: 'Awaiting top-up',
        toBetraded: 'Awaiting trading',
        trading: 'Trading',
        times: 'times',
        invitationSucceeded: 'Invitation completed',
        rules: 'Rules',
        inviteReward: '100',
        rules1: '1. You can get <span ref="currencySymbol" class="currencySymbol"></span><span ref="inviteReward" class="inviteReward"></span> bonus if your invited users trade 20 times in the real account.',
        rules2: '2. Unlimited invitation number of people, the more people you invite, the more bonus you get.',
        rules3: '3. Bonus distribution time : 8:00 a.m of the next day after completing requirements.',
        rules4: '4. Invitation code is only valid when filled during registration',
        contactService: 'Please contact customer service for any questions.',
    },
    MX: {
        title: 'Invite friends',
        bannerTitleH2: 'INVITE YOUR FRIENDS ',
        bannerTitleH3: 'GET BONUS NO LIMIT!',
        inviteTitle: 'Invite friends',
        inviteName: 'Referral code',
        shareText: 'Share Google Play Link',
        invitedUsers: 'Invited Users',
        toBeRecharged: 'Awaiting top-up',
        toBetraded: 'Awaiting trading',
        trading: 'Trading',
        times: 'times',
        invitationSucceeded: 'Invitation completed',
        rules: 'Rules',
        inviteReward: '100',
        rules1: '1.The users you invited need to trade 20 times,then you can get <span ref="currencySymbol" class="currencySymbol"></span><span ref="inviteReward" class="inviteReward"></span> bouns.',
        rules2: '2. Unlimited invitation number of people, the more people you invite, the more bonus you get.',
        rules3: '3. Bonus distribution time : 8:00 a.m of the next day after completing requirements.',
        rules4: '4. Invitation code is only valid when filled during registration',
        contactService: 'Please contact customer service for any questions.',
    }
};
// 交易指南
var tradingGuide = {
    INSA: {
        title: 'Panduan tutorial',
        step1TipsText: 'Selamat datang di <br /> SmartOption!',
        step1Title0: 'Kami menyediakan platform trading yang aman dan nyaman bagi pemula.',
        step1Title1: 'Layanan pelanggan 24 jam siap melayani anda dengan senang hati.',
        step1text: 'Ini adalah latihan trading pertama kali kamu, silakan ikuti kami!',
        step1Start: 'Start',
        step1JumpOver: 'Skip',
        step2Title: 'Step 1',
        step2Text: 'Silakan cek aset dan margin keuntungan. <br />Klik <span>"▼"</span> untuk memilih lebih banyak aset.',
        step2Next: 'Next',
        step3Title: 'Step 2',
        step3Text: 'Pilih <span>"+"</span> atau <span>"-"</span> untuk mengatur jumlah transaksi.',
        step4LeftTitle: 'Beli turun',
        step4BoxText: 'Dapatkan untung 75% setelah 1 menit',
        step4RightTitle: 'Beli naik',
        step4Title: 'Step 3',
        step4Text: 'Prediksi naik atau turun setelah 1 menit, dan pilih naik <span>"↗"</span> atau turun <span>"↘"</span>. <br />Hasilkan keuntungan 75% jika Anda memilih arah yang benar.',
        step5Text: '<span>Bonus</span> adalah bonus yang disediakan secara gratis dari platform. Bisa digunakan untuk trading di akun nyata. keuntungan bisa ditarik.',
        step6Text: 'Kami menyiapkan <span>${common[LKEY].currencySymbol} 1.000.000</span> untuk Anda di <span>akun Demo.</span> <br />Ayo mulai perjalanan investasi Anda!',
        step6TryNow: 'Try Now!'
    },
    EN: {
        title: 'Tutorials',
        step1TipsText: 'Welcome to <br /> SmartOption!',
        step1Title0: 'We offer safe and convenient trading platform for beginners.',
        step1Title1: '24/7 online service, we are always glad to assist.',
        step1text: 'Here is your first practice trading, come with us!',
        step1Start: 'Start',
        step1JumpOver: 'Skip',
        step2Title: 'Step 1',
        step2Text: 'Check asset & profit ratio here.<br />Click <span>"▼"</span> to choose more asset.',
        step2Next: 'Next',
        step3Title: 'Step 2',
        step3Text: 'Click <span>"+"</span> or <span>"-"</span> to set your trading amount.',
        step4LeftTitle: 'Down',
        step4BoxText: 'Earn 75% profit after 1min',
        step4RightTitle: 'Up',
        step4Title: 'Step 3',
        step4Text: 'Predict rise or fall after 1 min, and select up <span>"↗"</span> or down <span>"↘"</span>. <br />Earn 75% profit if you choose the right direction.',
        step5Text: '<span>Bonus</span> are gifts of our platform. Can be used as free trading money. <br />You can withdraw the bonus profit as well.',
        step6Text: 'We prepare <span>${common.EN.currencySymbol} 10.000</span> for you on the <span>Demo account</span>. <br />Let`s start your investment journey!',
        step6TryNow: 'Try Now!',
    },
    MX: {
        title: 'Tutorials',
        step1TipsText: 'Welcome to <br /> SmartOption!',
        step1Title0: 'We offer safe and convenient trading platform for beginners.',
        step1Title1: '24/7 online service, we are always glad to assist.',
        step1text: 'Here is your first practice trading, come with us!',
        step1Start: 'Start',
        step1JumpOver: 'Skip',
        step2Title: 'Step 1',
        step2Text: 'Check asset & profit ratio here.<br />Click <span>"▼"</span> to choose more asset.',
        step2Next: 'Next',
        step3Title: 'Step 2',
        step3Text: 'Click <span>"+"</span> or <span>"-"</span> to set your trading amount.',
        step4LeftTitle: 'Down',
        step4BoxText: 'Earn 75% profit after 1min',
        step4RightTitle: 'Up',
        step4Title: 'Step 3',
        step4Text: 'Predict rise or fall after 1 min, and select up <span>"↗"</span> or down <span>"↘"</span>. <br />Earn 75% profit if you choose the right direction.',
        step5Text: '<span>Bonus</span> are gifts of our platform. Can be used as free trading money. <br />You can withdraw the bonus profit as well.',
        step6Text: 'We prepare <span>${common.MX.currencySymbol} 100.000</span> for you on the <span>Demo account</span>. <br />Let`s start your investment journey!',
        step6TryNow: 'Try Now!',
    }
};
// 亏损返还
var lossReturn = {
    INSA: {
        title: 'Pengembalian kerugian',
        contentText0: 'KAMI AKAN',
        contentText1: 'MENANGGUNG KERUGIAN ANDA!',
        contentText2: 'BENEFIT KHUSUS SELURUH JARINGAN!',
        contentBoxText: 'TIDAK ADA BATAS ATAS JUMLAH!',
        contentTitle: 'Bonus',
        contentTips: 'Peraturan:',
        contentTipsText0: 'Maksimum pengembalian* 15% dari jumlah kerugian aktual hari itu, dibayarkan sebagai bonus.',
        contentTipsText1: '(*kerugian aktual = total kerugian hari itu - total profit hari itu)',
        contentTipsText2: 'Waktu pembayaran bonus : Pukul 8:00 pagi hari berikutnya',
    },
    EN: {
        title: 'Compensation of the loss',
        contentText0: null,// 英语不需要这个值
        contentText1: 'WE RECOVER YOUR LOSSES!',
        contentText2: 'NETWORK-WIDE <br /> EXCLUSIVE BENEFITS!',
        contentBoxText: 'No Amount Limit!',
        contentTitle: 'Bonus',
        contentTips: 'Rules:',
        contentTipsText0: 'Maximum Return* 15% of actual loss amount of the day, will be paid as bonus.',
        contentTipsText1: '(*actual loss = total loss of the day - total profit of the day)',
        contentTipsText2: 'Bonus payment time : 8:00 a.m of the next day',
    },
    MX: {
        title: 'Compensation of the loss',
        contentText0: null,// 英语不需要这个值
        contentText1: 'WE RECOVER YOUR LOSSES!',
        contentText2: 'NETWORK-WIDE <br /> EXCLUSIVE BENEFITS!',
        contentBoxText: 'No Amount Limit!',
        contentTitle: 'Bonus',
        contentTips: 'Rules:',
        contentTipsText0: 'Maximum Return* 15% of actual loss amount of the day, will be paid as bonus.',
        contentTipsText1: '(*actual loss = total loss of the day - total profit of the day)',
        contentTipsText2: 'Bonus payment time : 8:00 a.m of the next day',
    }
};
// 付款方式说明
var paymentInfo = {
    INSA: {
        title: 'Informasi isi ulang',
        contTitle: 'Deskripsi cara pembayaran',
        otherPay: 'Pilih metode pembayaran lain?',
        contTips: 'Jika Anda belum menerima pembayaran setelah 30 menit, Silakan hubungi kami, Anda dapat melihatnya di [Saya]-[Riwayat isi ulang] setelah diisi ulang',
        aldPay0Title: 'BCA',
        aldPay0Channel: 'BCA',
        aldPay0Text: [
            '<span class="under_line">ATM BCA</span>',
            '1. Masukan kartu ATM',
            '2. Pilih Bahasa Indonesia',
            '3. Masukan PIN ATM dengan Benar',
            '4. Cari Menu Transfer',
            '5. Pilih Transfer ke lainnya >Transfer > ke Rek BCA',
            '6. Masukan nomor virtual akun yang dituju <span class="payMoney"></span>',
            '7. Konfirmasi ulang transfer, jikabenar maka pilih Benar',
            '8. Transfer sesama bank akandiproses',
            '<span class="under_line">BCA Mobile</span>',
            '1. Menu m-Transfer',
            '2. Pilih Transfer – BCA Virtual Account',
            '3. Input nomor virtual account <span class="payCode"></span>, klik OK apabila telah sesuai',
            '4. Klik OK apabila telah sesuai',
            '<span class="under_line">KLIK BCA INDIVIDUAL</span>',
            '1. Menu Transfer Dana',
            '2. Pilih Transfer ke BCA Virtual Account',
            '3. Input nomor virtual account <span class="payCode"></span>',
            '4. Input nomor respon KeyBCA APPLI 1',
            '<span class="under_line">KLIK BCA BISNIS</span>',
            '1. Menu Transfer Dana',
            '2. Pilih Daftar Transfer - Tambah',
            '3. Pilih Jenis Tujuan ke BCA Virtual Account',
            '4. Input nomor virtual account <span class="payCode"></span>, dan klik tombol kirim apabila telah sesuai',
            '5. Pilih Transfer Dana ke BCA Virtual Account',
            '6. Pilih Otorisasi Transaksi dan klik tombol setuju apabila telah sesuai',
            '7. Input nomor respon KeyBCA',
        ],
        aldPay1Title: 'Bank Mandiri',
        aldPay1Channel: 'MANDIRI',
        aldPay1Text: [
            '<span class="title">Channel Bank Mandiri</span>',
            '1. Pilih Multipayment',
            '2. Pilih i-Pay (70017)',
            '3. Masukkan nomor <span class="payCode"></span> dan nomor jumlah',
            '4. Lakukan pembayaran',
            '<span class="title">Channel non-Bank Mandiri</span>',
            '1. Pilih Transfer Bank',
            '2. Pilih Bank Mandiri (008)',
            '3. Masukkan 88017 diikuti dengan Nomor <span class="payCode"></span>',
            '4. Lakukan Pembayaran',
        ],
        aldPay2Title: 'BRI',
        aldPay2Channel: 'BRI',
        aldPay2Text: [
            '<span class="under_line">ATM BRI</span>',
            '1. Pilih menu Transaksi Lain',
            '2. Pilih menu Pembayaran',
            '3. Pilih menu Lainnya',
            '4. Pilih menu BRIVA',
            '5. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '6. Pilih Ya untuk memproses pembayaran',
            '<span class="under_line">M-Banking BRI</span>',
            '1. Masuk ke aplikasi BRI Mobile dan pilih Mobile Banking BRI',
            '2. Pilih menu Info',
            '3. Pilih menu BRIVA',
            '4. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '5. Masukkan PIN Mobile/SMS Banking BRI',
            '6. Anda akan mendapatkan notifikasi pembayaran melalui SMS',
            '<span class="under_line">Internet Banking BRI</span>',
            '1. Login pada halaman Internet Banking BRI',
            '2. Pilih menu Pembayaran',
            '3. Pilih menu BRIVA',
            '4. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '5. Masukkan password Internet Banking BRI',
            '6. Masukkan mToken Internet Banking BRI',
            '7. Anda akan mendapatkan notifikasi pembayaran',
        ],
        aldPay3Title: 'BNI',
        aldPay3Channel: 'BNI',
        aldPay3Text: [
            '<span class="under_line">BNI ATM</span>',
            '1. Masukkan kartu debit/ATM Anda',
            '2. Masukkan PIN Anda',
            '3. Pilih Menu Lain > Transfer',
            '4. Pilih rekening asal dan pilih rekening tujuan ke rekening BNI',
            '5. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '6. Masukkan jumlah nominal',
            '7. Ikuti petunjuk hingga transaksi selesai',
            '<span class="under_line">BNI Mobile (mBanking)</span>',
            '1. Pilih Transfer > Antar Rekening BNI',
            '2. Pilih Rekening Tujuan > Input Rekening Baru. Masukkan nomor rekening dengan nomor Virtual Account Anda dan klik Lanjut, kemudian klik Lanjut lagi.',
            '3. Masukkan jumlah pembayaran sejumlah tagihan Anda, lalu klik Lanjutkan.',
            '4. Periksa detail konfirmasi. Pastikan Nama Rekening Tujuan adalah nama penerima Anda dan nominal transaksi sudah benar. Jika benar, masukkan password transaksi dan klik Lanjut.',
            '5. Ikuti petunjuk hingga transaksi selesai',
            '<span class="under_line">iBanking</span>',
            '1. Pilih Transaksi > Info & Administrasi Transfer > Atur Rekening Tujuan > Tambah Rekening Tujuan dan klik OK.',
            '2. Pilih Kode Network & Bank: Transfer Antar Rek. BNI. Masukkan nomor rekening dengan nomor Virtual Account Anda dan klik Lanjut. Isi data lainnya dan klik Lanjutkan.',
            '3. Cek detail konfirmasi. Pastikan Nama Rekening Tujuan adalah nama penerima Anda dan nominal transaksi sudah benar. Masukkan PIN BNI e-Secure Anda dan klik Proses.',
            '4. Ikuti petunjuk hingga transaksi selesai',
        ],

        aldPay4Title: 'Bank Permata / Permata Syariah',
        aldPay4Channel: 'PERMATA',
        aldPay4Text: [
            '<span class="under_line">Permata ATM</span>',
            '1. Masukkan kartu debit/ATM Anda',
            '2. Masukkan PIN Anda',
            '3. Pilih "Transaksi Lainnya"',
            '4. Pilih "Virtual Account"',
            '5. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '6. Masukkan jumlah nominal',
            '7. Ikuti petunjuk hingga transaksi selesai',
            '<span class="under_line">Permata Mobile</span>',
            '1. Pilih menu "Pembayaran Tagihan"',
            '2. Pilih menu "Virtual Account"',
            '3. Daftarkan nomor Virtual Account jika belum terdaftar. Pilih "Tagihan Anda" > "Daftar Tagihan Baru"',
            '4. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '5. Pastikan semua data ringkasan terisi, tekan tombol Lanjut jika sudah sesuai',
            '6. Masukkan response code yang dikirim ke nomor ponsel Anda',
            '7. Ikuti petunjuk hingga transaksi selesai',
            '<span class="under_line">PermataNET</span>',
            '1. Pilih menu "Pembayaran"',
            '2. Pilih menu "Pembayaran Tagihan"',
            '3. Pilih "Virtual Account"',
            '4. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '5. Masukkan response code yang dikirim ke nomor ponsel Anda',
            '6. Ikuti petunjuk hingga transaksi selesai',
        ],

        aldPay5Title: 'DANA',
        aldPay5Channel: 'DANA',
        aldPay5Text: [
            'Bayar di halaman DANA',
            'ID Transaksi: <span class="payCode"></span>',
        ],

        aldPay6Title: 'OVO',
        aldPay6Channel: 'OVO',
        aldPay6Text: [
            'Bayar di aplikasi OVO',
            'ID Transaksi: <span class="payCode"></span>',
            'Anda akan menerima pesan pembayaran dari OVO. ',
            'Klik pesan dan selesaikan pembayaran. ',
            'Setelah dibayar, klik "Cek Status Pembayaran"',
        ],

        aldPay7Title: 'CIMB Niaga / CIMB Niaga Syariah',
        aldPay7Channel: 'CIMB_NIAGA',
        aldPay7Text: [
            '<span class="under_line">CIMB Clicks</span>',
            '1. Login ke CIMB clicks (<a href="https://www.cimbclicks.co.id/" style="color: #3062ff;">https://www.cimbclicks.co.id/</a>)',
            '2. Pilih menu Bayar Tagihan',
            '3. Pilih Rekening Sumber dan Jenis Pembayaran > Virtual Account',
            '4. Masukkan nomor virtual account CIMB Anda',
            '5. Nomor, nama virtual account dan jumlah billing ditampilkan pada layar',
            '6. Masukkan 6 digit mPIN dan tekan tombol Submit',
            '7. Konfirmasi pembayaran ditampilkan pada layar',
            '<span class="under_line">CIMB Mobile - Go Mobile</span>',
            '1. Login ke Go Mobile',
            '2. Pilih menu TRANSFER',
            '3. Pilih Transfer to Other CIMB Niaga Account',
            '4. Pilih rekening sumber anda : CASA atau Rekening Ponsel',
            '5. Pilih CASA dan masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '6. Masukkan jumlah pembayaran sesuai tagihan',
            '7. Masukkan Mobile Banking PIN',
            '8. Konfirmasi pembayaran ditampilkan pada layar',
            '<span class="under_line">ATM CIMB</span>',
            '1. Masukkan Kartu ATM dan PIN CIMB Anda',
            '2. Pilih menu Pembayaran > Lanjut > Virtual Account',
            '3. Masukkan nomor Virtual Account <span class="payCode"></span> Anda',
            '4. Pilih rekening debit',
            '5. Nomor, nama virtual account dan jumlah billing ditampilkan pada layar',
            '6. Pilih OK untuk melakukan pembayaran',
            '7. Konfirmasi pembayaran ditampilkan pada layar',
        ],
        wypPay8Title: 'OVO WALLET',
        wypPay8Channel: 'OVO',
        wypPay8Text: [
            '1. Harap tetap buka aplikasi OVO;',
            '2. Masukkan nomor telepon (ID OVO) ',
            '3. Klik tombol "Lanjutkan ke aplikasi OVO',
            '4. Buka aplikasi OVO untuk mengkonfirmasi transaksi dan membayar dalam 55 detik',
            '(Catatan: Pembayaran akan kedaluwarsa dalam 55 detik)',
            '5. Selesai'
        ],

        wypPay9Title: 'ALFA GROUP',
        wypPay9Channel: 'ALFAGROUP',
        wypPay9Text: [
            '1. Catat kode pembayaran lalu datang ke gerai Alfamart, Alfa Midi, Alfa Express, Lawson atau DAN+DAN terdekat ',
            '2. Beritahukan ke kasir bahwa anda ingin melakukan pembayaran ke PLASAMALL ',
            '3. Jika kasir tidak mengetahui mengenai pembayaran PLASAMALL, informasikan ke kasir untuk membuka terminal e-transaction, pilih "no. 2 menu Pembayaran atau cari pada menu search" ',
            '4. Minta kasir untuk menekan "Shift" dan "?" lalu ketik PLASAMALL, setelah itu klik OK ',
            '5. Minta kasir untuk pilih "no.1 PLASAMALL - Pembayaran PLASAMALL " ',
            '6. Kasir akan menanyakan kode pembayaran. Berikan kode pembayaran anda.contoh 10013xxxxxxx. Kasir akan melakukan konfirmasi data konsumen berupa Nama Merchant, Nama Konsumen, dan Nominal. Lakukan pembayaran ke kasir sejumlah nominal yang disebutkan',
            '7. Terima struk sebagai bukti pembayaran sudah sukses dilakukan. Notifikasi pembayaran akan langsung diterima oleh Merchant ',
            '8. Selesai',
        ],

        wypPay10Title: 'PERMATA VA',
        wypPay10Channel: 'PERMATA',
        wypPay10Text: [
            '<span class="under_line">1. ATM PERMATA </span>',
            '• Pilih menu Transaksi Lainnya ',
            '• Pilih Pembayaran ',
            '• Pilih Pembayaran Lainnya ',
            '• Pilih Virtual Account ',
            '• Input nomor virtual account ',
            '• Pada layar akan tampil konfirmasi pembayaran',

            '<span class="under_line">2. MOBILE INTERNET </span>',
            '• Masukan User ID & Password ',
            '• Pilih Pembayaran Tagihan ',
            '• Pilih virtual account ',
            '• Input nomor virtual account ',
            '• Input nominal pembayaran ',
            '• Apabila data sudah sesuai masukan otentikasi transaksi/Token',


            '<span class="under_line">3. ATM BERSAMA, ALTO & PRIMA  </span>',
            '• Pilih menu Transfer ',
            '• Pilih Transfer ke Bank Lain ',
            '    <b>Jika melalui ATM Bersama & Alto : </b>',
            '    Input kode PermataBank ',
            '    (013) +16 digit kode bayar ',
            '    <b>Jika melalui ATM Prima : </b>',
            '    Masukan kode PermataBank(013)tekan BENAR, dilanjutkan dengan 16 digit kode bayar ',
            '• Input nominal pembayaran ',
            '• Pada layar akan tampil konfirmasi Nama & Nominal Pembayaran',
        ],

        wypPay11Title: 'BRI VA',
        wypPay11Channel: 'BRI',
        wypPay11Text: [
            '<span class="under_line">1. BRImo  </span>',
            '• Login pada aplikasi BRImo (input username dan password) ',
            '• Pilih menu BRIVA ',
            '• Pilih sumber dana ',
            '• Input kode BRIVA ',
            '• Konfirmasi dan proses pembayaran ',
            '<span2. BRI INET BANKING </span>',
            '• Input username dan password ',
            '• Pilih transfer ',
            '• Input kode BRIVA ',
            '• Konfirmasi dan proses pembayaran',

            '<span class="under_line">3. ATM BRI </span>',
            '• Input kartu dan PIN ',
            '• Pilih menu Transaksi Lain > Pembayaran > Lainnya > BRIVA ',
            '• Input kode BRIVA ',
            '• Konfirmasi dan proses pembayaran ',
            '<span class="under_line">4. ATM BANK LAINNYA</span>',
            '• Input kartu dan PIN ',
            '• Pilih menu Transaksi Lain > Transfer > Ke Rek Bank Lain ',
            '• Input kode bank 002 ',
            '• Input nominal transfer ',
            '• Input kode BRIVA ',
            '• Pilih rekening yang akan didebet ',
            '• Konfirmasi dan proses pembayaran',

            '<span class="under_line">5. MOBILE BANKING BRI  </span>',
            '• Pilih menu Info > Info BRIVA  ',
            '• Input kode BRIVA  ',
            '• Input password  ',
            '• Konfirmasi dan proses pembayaran  ',
            '<span class="under_line">6. MELALUI TELLER BRI </span>',
            '• Mengisi form sesuai ketentuan BANK BRI  ',
            '• Kasih form dan uang kepada teller  ',
            '• Terima struk sebagai bukti pembayaran sudah sukses dilakukan ',

        ],
        wypPay12Title: 'CIMB VA',
        wypPay12Channel: 'CIMB_NIAGA',
        wypPay12Text: [
            '<span class="under_line">1. ATM CIMB </span>',
            '• Menu Transfer ',
            '• Pilih Antar Rekening CIMB NIAGA ',
            '• Input nominal transfer ',
            '• Input nomor virtual account di rekening tujuan ',
            '• Konfirmasi dan proses pembayaran ',
            '<span class="under_line">2. CIMB CLICKS </span>',
            '• Menu Transfer ',
            '• Input nominal transfer ',
            '• Pilih Other Account – Bank CIMB Niaga / Rekening Ponsel ',
            '• Rekening penerima diisi dengan nomor virtual account ',
            '• Input mPin ',
            '• Konfirmasi dan proses pembayaran',
            '<span class="under_line">3. GO MOBILE </span>',
            '• Menu Transfer ',
            '• Pilih Transfer to Other CIMB Niaga Account ',
            '• Input nomor virtual account dan nominal pembayaran ',
            '• Input PIN ',
            '• Konfirmasi dan proses pembayaran ',
            '<span class="under_line">4. INTERNET BANKING LAINNYA </span>',
            '• Menu Transfer ',
            '• Pilih Rekening Bank Lain',
            '• Input nomor virtual account pada rekening tujuan ',
            '• Input nominal pembayaran ',
            '• Input bank penerima dengan CIMB Niaga ',
            '• Jenis transfer yang digunakan adalah online/realtime ',
            '• Konfirmasi dan proses pembayaran',
            '<span class="under_line">5. ATM BANK LAINNYA </span>',
            '• Menu Transfer ',
            '• Pilih Rekening Bank Lain ',
            '• Input kode bank 022 ',
            '• Input nomor virtual account pada rekening tujuan ',
            '• Input nominal pembayaran ',
            '• Konfirmasi dan proses pembayaran',

        ],

        wypPay13Title: 'MANDIRI VA',
        wypPay13Channel: 'MANDIRI',
        wypPay13Text: [
            '<span class="under_line">1. ATM MANDIRI Pilih menu Bayar/Beli </span>',
            'Pilih menu Multipayment ',
            'Input kode perusahaan Bank Mandiri （5 digit terdepan kode pembayaran） tekan Benar ',
            'Input nomor virtual account kemudian tekan Benar ',
            'Pilih Nomor 1 atau sesuai tagihan yang akan dibayar lalu tekan YA',

            '<span class="under_line">1. ATM MANDIRI Pilih menu Bayar/Beli </span>',
            'Pilih menu Multipayment ',
            'Input kode perusahaan Bank Mandiri （5 digit terdepan kode pembayaran） tekan Benar ',
            'Input nomor virtual account kemudian tekan Benar ',
            'Pilih Nomor 1 atau sesuai tagihan yang akan dibayar lalu tekan YA',

            '<span class="under_line">2. ONLINE BANKING </span>',
            '• Input User ID & Password ',
            '• Pilih menu Pembayaran ',
            '• Pilih menu Multipayment ',
            '• Pada Kolom Penyedia Jasa, pilih penyedia layanan yang sesuai (Masukkan 5 digit terdepan kode pembayaran) ',
            '• Input nomor virtual account ',
            '• Pilih Lanjut ',
            '• Web : ',
            'a. OTP dikirimkan ke nomor HP. ',
            'b. OTP yang didapatkan diinput di Token untuk mendapatkan challenge code ',
            'c. Chanllenge code input ke Mandiri online, lalu klik Lanjut ',
            '• Apps : ',
            'a. Setelah muncul konfirmasi transaksi pilih Kirim ',
            'b. Input MPIN (6 digit)',

        ],
        wypPay14Title: 'BNI VA',
        wypPay14Channel: 'BNI',
        wypPay14Text: [
            '<span class="under_line">1. ATM BNI</span>',
            '• Menu Lainnya',
            '• Pilih Transfer',
            '• Pilih Rekening Tabungan',
            '• Pilih Ke Rekening BNI',
            '• Input nomor virtual account',
            '• Input nominal transfer',
            '• Konfirmasi pemindahbukuan',
            '<span class="under_line">2. INTERNET BANKING BNI</span>',
            '• Menu Transfer',
            '• Pilih Tambah Rekening Favorit',
            '• Pilih Antar Rekening BNI',
            '• Bila menggunakan desktop untuk menambah rekening: menu Transaksi, lalu Atur',
            'Rekening Tujuan, Tambah Rekening Tujuan',
            '• Masukkan nama dan nomor virtual account',
            '• Input kode otentikasi token',
            '• Rekening tujuan berhasil ditambahkan',
            '• Kembali ke menu Transfer',
            '• Pilih Transfer Antar Rekening BNI',
            '• Pilih Rekening Tujuan',
            '• Pilih rekening debet dan input nominal transfer',
            '• Input kode otentikasi token',
            '<span class="under_line">3. MOBILE BANKING BNI</span>',
            '• Menu Transfer',
            '• Pilih Antar Rekening BNI dan Input Rekening Baru',
            '• Input rekening debet, nomor virtual account di rekening tujuan, dan nominal transfer',
            '• Input password dan konfirmasi transaksi',
            '<span class="under_line">4. INTERNET BANKING LAINNYA</span>',
            '• Menu Transfer',
            '• Pilih Rekening Bank Lain',
            '• Input nomor virtual account pada rekening tujuan',
            '• Input nominal pembayaran',
            '• Input bank penerima dengan BNI',
            '• Jenis transfer yang digunakan adalah online/realtime',
            '• Konfirmasi dan proses pembayaran',
            '<span class="under_line">5. ATM BANK LAINNYA</span>',
            '• Menu Transfer',
            '• Pilih Rekening Bank Lain',
            '• Input kode bank 009',
            '• Input nomor virtual account pada rekening tujuan',
            '• Input nominal pembayaran',
            '• Konfirmasi dan proses pembayaran',
        ],

        cashyPay0Title: 'BRI VA',
        cashyPay0Channel: 'BRI',
        cashyPay0Text: [
            '<span class="under_line">1、Mobile Banking</span>',
            'LANGKAH 1: MASUK KE AKUN ANDA ',
            '1. Buka aplikasi BRI Mobile Banking, masukkan USER ID dan PIN anda ',
            '2. Pilih "Pembayaran" dan pilih "Briva" ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Masukkan Nomor Virtual Account anda dan jumlah yang ingin anda bayarkan ',
            '2. Masukkan PIN Mobile Banking BRI ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Setelah transaksi anda selesai, invoice ini akan diupdate secara otomatis. Proses ini mungkin memakan waktu hingga 5 menit',

            '<span class="under_line">2、Internet Banking</span>',
            'LANGKAH 1: MASUK KE AKUN ANDA',
            '1. Buka situs <a href="https://ib.bri.co.id/ib-bri/"  target="_blank">https://ib.bri.co.id/ib-bri/</a> , dan masukkan USER ID dan PASSWORD anda ',
            '2. Pilih "Pembayaran" dan pilih "Briva" ',
            'LANGKAH 2: DETAIL PEMBAYARAN',
            '1. Masukkan Nomor Virtual Account 92001988572642231 dan jumlah yang ingin anda bayarkan ',
            '2. Masukkan password anda kemudian masukkan mToken internet banking ',
            'LANGKAH 3: TRANSAKSI BERHASIL',
            '1. Setelah transaksi anda selesai, invoice ini akan diupdate secara otomatis. Proses in',

            '<span class="under_line">3、ATM</span>',
            'LANGKAH 1: TEMUKAN ATM TERDEKAT',
            '1. Masukkan kartu, kemudian pilih bahasa dan masukkan PIN anda ',
            '2. Pilih "Transaksi Lain" dan pilih "Pembayaran" ',
            '3. Pilih menu "Lainnya" dan pilih "Briva" ',
            'LANGKAH 2: DETAIL PEMBAYARAN',
            '1. Masukkan Nomor Virtual Account dan jumlah yang ingin anda bayarkan ',
            '2. Periksa data transaksi dan tekan "YA" ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Setelah transaksi anda selesai, invoice ini akan diupdate secara otomatis. Proses ini mungkin memakan waktu hingga 5 menit'

        ],
        cashyPay1Title: 'ATM Bersama VA',
        cashyPay1Channel: 'BERSAMA',
        cashyPay1Text: [
            '<span class="under_line">1、Internet Banking</span>',
            'LANGKAH 1: MASUK KE AKUN ANDA',
            '1. Masuk ke <a href="https://www.atmbersama.com/"  target="_blank">https://www.atmbersama.com/</a> ',
            '2. Pilih "Transfer ke Bank Lain" ',
            '3. Masukkan Kode Bank dan nomor rekening virtual',
            '4. Masukkan Jumlah Pembayaran',
            '5. Konfirmasi untuk menyelesaikan proses pembayaran',

            '<span class="under_line">1、ATM</span>',
            'LANGKAH 1: TEMUKAN ATM TERDEKAT ',
            '1. Masukkan Kartu Debit Anda dengan logo "ATM Bersama" ',
            '2. Pilih "Transfer ke Bank Lain',
            '3. Masukkan Kode Bank  dan nomor rekening virtual',
            '4. Masukkan Jumlah Pembayaran',
            '5. Konfirmasi untuk menyelesaikan proses pembayaran',
        ],
        cashyPay2Title: 'Permata VA',
        cashyPay2Channel: 'PERMATA',
        cashyPay2Text: [
            '<span class="under_line">1、Mobile Banking</span>',
            'LANGKAH 1: MASUK KE AKUN ANDA ',
            '1. Buka aplikasi PermataMobile Internet ',
            '2. Masukkan User ID dan Password ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Pilih "Pembayaran Tagihan" ',
            '2. Pilih "Virtual Account" ',
            '3. Masukkan Nomor Virtual Account Anda',
            '4. Masukkan otentikasi transaksi/token ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Transaksi Anda telah selesai ',
            '2. Setelah transaksi anda selesai, invoice ini akan diupdate secara otomatis. Proses ini mungkin memakan waktu hingga 5 menit',

            '<span class="under_line">2、Internet Banking</span>',
            'LANGKAH 1: MASUK KE AKUN ANDA ',
            '1. Buka situs <a href="https://www.permatanet.com/"  target="_blank">https://www.permatanet.com/ </a> ',
            '2. Masukkan User ID dan Password ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Pilih "Pembayaran Tagihan" ',
            '2. Pilih "Virtual Account" ',
            '3. Masukk Nomor Virtual Account ',
            '4. Periksa kembali detail pembayaran anda ',
            '5. Masukkan otentikasi transaksi/token ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Transaksi Anda telah selesai ',
            '2. Setelah transaksi anda selesai, invoice ini akan diupdate secara otomatis. Proses ini mungkin memakan waktu hingga 5 menit ',

            '<span class="under_line">3、ATM </span>',
            'LANGKAH 1: TEMUKAN ATM TERDEKAT',
            '1. Masukkan kartu ATM Permata anda ',
            '2. Masukkan PIN ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Pilih menu "Transaksi Lainnya" ',
            '2. Pilih menu "Pembayaran" ',
            '3. Pilih menu "Pembayaran Lainnya" ',
            '4. Pilih menu "Virtual Account" ',
            '5. Masukkan Nomor Virtual Account',
            '6. Lalu pilih rekening debet yang akan digunakan ',
            '7. Konfirmasi detail transaksi anda ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Transaksi Anda telah selesai ',
            '2. Setelah transaksi anda selesai, invoice ini akan diupdate secara otomatis. Proses ini mungkin memakan waktu hingga 5 menit'

        ],


        cashyPay3Title: 'Mandiri VA',
        cashyPay3Channel: 'MANDIRI',
        cashyPay3Text: [
            '<span class="under_line">1、Mobile Banking</span>',
            'Old Livin by Mandiri',
            'LANGKAH 1: MASUK KE AKUN ANDA',
            '1. Buka aplikasi Mandiri Online, masukkan USERNAME dan PASSWORD anda',
            '2. Pilih "Bayar" ',
            '3. Pilih "Multipayment" ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Pilih "6 digit pertama Nomor Virtual Account" (untuk closed VA) sebagai penyedia jasa ',
            '2. Masukkan Nomor Virtual Account ',
            '3. Tekan Lanjut ',
            '4. Tinjau dan konfirmasi detail transaksi anda, lalu tekan Konfirmasi ',
            '5. Selesaikan transaksi dengan memasukkan MPIN anda ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Setelah transaksi pembayaran Anda selesai, simpan bukti pembayaran ',
            '2. Invoice ini akan diperbarui secara otomatis. Ini bisa memakan waktu hingga 5 menit',
            'New Livin by Mandiri',
            'LANGKAH 1: MASUK KE AKUN ANDA ',
            '1. Buka aplikasi Livin by Mandiri, masukkan PASSWORD atau lakukan verifikasi wajah',
            '2. Pilih "Bayar" ',
            '3. Cari “6 digit pertama Nomor Virtual Account” (untuk Closed VA) ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Pilih "6 digit pertama Nomor Virtual Account" (untuk closed VA) sebagai penyedia jasa ',
            '2. Masukkan Nomor Virtual Account',
            '3. Masukkan nominal yang ingin dibayarkan (Jika VA merupakan closed VA, maka nominal akan otomatis terisi) ',
            '4. Tinjau dan konfirmasi detail transaksi anda, lalu tekan Konfirmasi ',
            '5. Selesaikan transaksi dengan memasukkan MPIN anda ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Setelah transaksi pembayaran Anda selesai, simpan bukti pembayaran ',
            '2. Invoice ini akan diperbarui secara otomatis. Ini bisa memakan waktu hingga 5 menit',

            '<span class="under_line">2、Internet Banking</span>',
            'LANGKAH 1: MASUK KE AKUN ANDA ',
            '1. Buka situs Mandiri Internet Banking <a href="https://ibank.bankmandiri.co.id"  target="_blank">https://ibank.bankmandiri.co.id</a>  ',
            '2. Masuk menggunakan USER ID dan PASSWORD anda ',
            '3. Buka halaman beranda, kemudian pilih "Pembayaran" ',
            '4. Pilih "Multi Payment" ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Pilih "6 digit pertama Nomor Virtual Account" (untuk closed VA)  sebagai penyedia jasa',
            '2. Masukkan Nomor Virtual Account',
            '3. Lalu pilih Lanjut ',
            '4. Apabila semua detail benar tekan "KONFIRMASI" ',
            '5. Masukkan PIN / Challenge Code Token ',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Setelah transaksi pembayaran Anda selesai, simpan bukti pembayaran ',
            '2. Invoice ini akan diperbarui secara otomatis. Ini bisa memakan waktu hingga 5 menit ',

            '<span class="under_line">3、ATM</span>',
            'LANGKAH 1: TEMUKAN ATM TERDEKAT ',
            '1. Masukkan ATM dan tekan "Bahasa Indonesia" ',
            '2. Masukkan PIN, lalu tekan "Benar" ',
            '3. Pilih "Pembayaran", lalu pilih "Multi Payment" ',
            'LANGKAH 2: DETAIL PEMBAYARAN ',
            '1. Masukkan kode perusahaan "6 digit pertama Nomor Virtual Account" untuk closed amount VA, lalu tekan "BENAR" ',
            '2. Masukkan Nomor Virtual Account, lalu tekan "BENAR"',
            '3. Untuk open amount VA, masukkan nominal yang ingin di transfer, lalu tekan "BENAR"',
            '4. Informasi pelanggan akan ditampilkan, pilih nomor 1 sesuai dengan nominal pembayaran kemudian tekan "YA" ',
            '5. Konfirmasi pembayaran akan muncul, tekan "YES", untuk melanjutkan',
            'LANGKAH 3: TRANSAKSI BERHASIL ',
            '1. Simpan bukti transaksi anda ',
            '2. Transaksi anda berhasil ',
            '3. Setelah transaksi anda selesai, invoice ini akan diupdate secara otomatis. Proses ini mungkin memakan waktu hingga 5 menit '

        ],
        cashyPay4Title: 'AlfaGroup VA',
        cashyPay4Channel: 'ALFAGROUP',
        cashyPay4Text: [
            '<span class="under_line">Alfamart Group (Alfamart, Afamidi, Dan+Dan, Lawson).</span>',
            'Pembayaran dengan uang tunai ',
            '1. Kunjungi gerai Alfamart/Alfamidi/Lawson/DAN+DAN terdekat sebelum kode pembayaran kedaluwarsa.',
            '2. Informasikan ke Kasir untuk melakukan pembayaran ke produk/servis anda dengan menyatakan nama Brand Anda.',
            '3. Berikan informasi KODE PEMBAYARAN yang tertera di Aplikasi atau Website kepada kasir.',
            '4. Jika anda memiliki barcode, dapat memberikan barcode tersebut ke kasir untuk di scan/pindai.',
            '5. Pastikan jumlah pembayaran di tempat anda berbelanja sama dengan jumlah yang diberitahukan oleh kasir.',
            '6 Informasikan pada kasir jika anda ingin membayar dengan uang tunai.',
            '7. Konfirmasi pembayaran dengan memberikan uang tunai pada kasir.',
            '8. Minta struk pembayaran sebagai bukti transaksi dari Alfamart.',
            '9. Pembayaran Alfamart anda sudah selesai.',

            ' Pembayaran dengan Kartu Debit/Prepaid/E-Wallet ',
            '1. Kunjungi gerai Alfamart/Alfamidi/Lawson/DAN+DAN terdekat sebelum kode pembayaran kedaluwarsa.',
            '2. Informasikan ke Kasir untuk melakukan pembayaran ke produk/servis anda dengan menyatakan nama Brand Anda.',
            '3. Berikan informasi KODE PEMBAYARAN yang tertera di Aplikasi atau Website kepada kasir.',
            '4. Jika anda memiliki barcode, dapat memberikan barcode tersebut ke kasir untuk di scan/pindai.',
            '5. Pastikan jumlah pembayaran di tempat anda berbelanja sama dengan jumlah yang diberitahukan oleh kasir.',
            '6. Informasikan pada kasir jika anda ingin membayar dengan Kartu Debit/Prepaid Card atau E-wallet.',
            '7. Konfirmasi pembayaran dengan memberikan Kartu Debit/Prepaid atau informasi E-wallet pada kasir.',
            '8. Minta struk pembayaran sebagai bukti transaksi dari Alfamart.',
            '9. Pembayaran Alfamart anda sudah selesai.',

            'Pembayaran gabungan uang tunai dan Kartu Debit/Prepaid/E-Wallet',
            '1. Kunjungi gerai Alfamart/Alfamidi/Lawson/DAN+DAN terdekat sebelum kode pembayaran kedaluwarsa.',
            '2. Informasikan ke Kasir untuk melakukan pembayaran ke produk/servis anda dengan menyatakan nama Brand Anda.',
            '3. Berikan informasi KODE PEMBAYARAN yang tertera di Aplikasi atau Website kepada kasir.',
            '4. Jika anda memiliki barcode, dapat memberikan barcode tersebut ke kasir untuk di scan/pindai.',
            '5. Pastikan jumlah pembayaran di tempat anda berbelanja sama dengan jumlah yang diberitahukan oleh kasir.',
            '6. Informasikan pada kasir jika anda ingin membayar dengan kombinasi/gabungan uang tunai dan Kartu Debit/Prepaid Card atau E-wallet.',
            '7. Konfirmasi pembayaran dengan memberikan uang tunai serta Kartu Debit/Prepaid atau informasi E-wallet pada kasir.',
            '8. Minta struk pembayaran sebagai bukti transaksi dari Alfamart.',
            '9. Pembayaran Alfamart anda sudah selesai.',
        ],
        cashyPay5Title: 'QRIS',
        cashyPay5Channel: 'QRIS',
        cashyPay5Text: [
            'QRIS dapat di-scan dan dibayar melalui semua aplikasi penyelenggara yang diberi izin oleh Bank Indonesia.'
        ]

    },

    EN: {
        title: 'Deposit information',
        contTitle: 'Payment method introduction',
        otherPay: 'Choose other Payment method',
        contTips: 'If your account have not received the deposit after 30 minutes, Please contact us. You can view it in[Me]-[Deposit Record] after deposit.',
        aldPay0Title: 'BCA',
        aldPay0Channel: 'BCA',
        aldPay0Text: [
            '<span class="under_line">ATM BCA</span>',
            '1. Insert the ATM card',
            '2. Select Indonesian language',
            '3. Enter the correct ATM PIN',
            '4. Go to the transfer menu',
            '5. Enter the beneficiary account number',
            '6. Enter the payment amount <span class="payMoney"></span>',
            '7. Reconfirm the information and amount, then select "Yes"',
            '8. Bank transfers will be processed',
            '<span class="under_line">BCA Mobile</span>',
            '1) Select m-Transfer',
            '2) Select Transfer – BCA Virtual Account',
            '3. Enter the virtual account number <span class="payCode"></span>, click OK after confirming',
            '4. Click OK after confirmation',
            '<span class="under_line">CLICK INDIVIDUAL BCA</span>',
            '1. Select Money Transfer',
            '2. Select Transfer to BCA Virtual Account',
            '3. Enter the virtual account number <span class="payCode"></span>',
            '4. Enter the response number of KeyBCA APPLI 1',
            '<span class="under_line">CLICK BCA BISNIS</span>',
            '1. Select Money Transfer',
            '2. Select Transfer List - Add new',
            '3. Select the destination type as BCA virtual account',
            '4. Enter the virtual account number <span class="payCode"></span>, click send after confirmation',
            '5. Select transfer to BCA virtual account',
            '6. Select transaction authorization, and click Agree after confirmation',
            '7. Enter the response number of KeyBCA',
        ],
        aldPay1Title: 'Bank Mandiri',
        aldPay1Channel: 'MANDIRI',
        aldPay1Text: [
            '<span class="title">Bank Mandiri Channels</span>',
            '1. Select Multipayment',
            '2. Select i-Pay (70017)',
            '3. Enter the number <span class="payCode"></span> and the amount',
            '4. Proceed your payment',
            '<span class="title">non-Bank Mandiri channel</span>',
            '1. Select Bank Transfer',
            '2. Select Bank Mandiri (008)',
            '3. Enter 88017 followed by the number <span class="payCode"></span>',
            '4. Proceed your payment',
        ],

        aldPay2Title: 'BRI',
        aldPay2Channel: 'BRI',
        aldPay2Text: [
            '<span class="under_line">ATM BRI</span>',
            '1. Select Other Transactions',
            '2. Select Payment',
            '3. Select Others',
            '4. Select BRIVA',
            '5. Enter Virtual account number <span class="payCode"></span>',
            '6. Select "Yes" to proceed transaction',
            '<span class="under_line">M-Banking BRI</span>',
            '1. Enter the BRI mobile app and select BRI Mobile Banking',
            '2. Select Information',
            '3. Select BRIVA',
            '4. Enter your virtual account number <span class="payCode"></span>',
            "5. Enter BRI's mobile PIN/SMS Banking",
            '6. You will be notified of payment via SMS',
            '<span class="under_line">Internet Banking BRI</span>',
            '1. Log in to the BRI Internet Banking page',
            '2. Select Payment',
            '3. Select BRIVA',
            '4. Enter your virtual account number <span class="payCode"></span>',
            '5. Enter password for Internet Banking',
            '6. Enter mToken (token) online banking',
            '7. You will be notified of payment',
        ],
        aldPay3Title: 'BNI',
        aldPay3Channel: 'BNI',
        aldPay3Text: [
            '<span class="under_line">BNI ATM</span>',
            '1. Insert your debit/ATM card',
            '2. Enter your PIN',
            '3. Select Other Menu > Transfer',
            '4. Select the source account and select the beneficiary BNI account',
            '5. Enter your Virtual Account number <span class="payCode"></span>',
            '6. Enter the amount',
            '7. Follow the instructions to complete transaction',
            '<span class="under_line">BNI Mobile (mBanking)</span>',
            '1. Select Transfer > Between BNI Accounts',
            '2. Select Destination Account > Input New Account. Enter your Virtual Account number and click Next, then click Next again.',
            '3. Enter the payment amount for your bill, then click Continue.',
            "4. Check the confirmation details. Make sure the beneficiary account and the transaction amount is correct. Enter the transaction password and click Continue if it's correct.",
            '5. Follow the instructions to complete the transaction.',
            '<span class="under_line">iBanking</span>',
            '1. Select Transaction > Transfer Info & Administration > Set Destination Account > Add  Beneficiary Account and click OK.',
            '2. Select Network & Bank Code: Transfer Between BNI Accounts. Enter your Virtual Account number and click Next. Fill in the rest of the data and click Continue.',
            '3. Check the confirmation details. Make sure the beneficiary account and the transaction amount is correct. Enter your BNI e-Secure PIN and click Process.',
            '4. Follow the instructions to complete the transaction',
        ],
        aldPay4Title: 'Bank Permata / Permata Syariah',
        aldPay4Channel: 'PERMATA',
        aldPay4Text: [
            '<span class="under_line">Permata ATM</span>',
            '1. Insert your debit/ATM card',
            '2. Enter your PIN',
            "3. Select 'Other Transactions'",
            "4. Select 'Virtual Accounts'",
            '5. Enter your Virtual Account number <span class="payCode"></span>',
            '6. Enter the payment amount',
            '7. Follow the instructions to complete the transaction',
            '<span class="under_line">Permata Mobile</span>',
            '1. Select the "Bill Payment" option',
            '2. Select virtual account',
            "3. Register your virtual account if it was not already registered. Select 'Your Bill' > 'Register a New Bill'",
            '4. Enter your virtual account number <span class="payCode"></span>',
            '5. Make sure to fill in all required data and click the Continue button',
            '6. Enter the response code sent to your mobile number',
            '7. Follow the instructions to complete the transaction',
            '<span class="under_line">PermataNET</span>',
            '1. Select “Payment”menu',
            '2. Select “Bill payment”',
            '3. Select virtual account',
            '4. Enter your virtual number <span class="payCode"></span>',
            '5. Enter the response code sent to your mobile number',
            '6. Follow the instructions to complete the transaction',
        ],
        aldPay5Title: 'DANA',
        aldPay5Channel: 'DANA',
        aldPay5Text: [
            'Pay in DANA page',
            'Transaction ID: <span class="payCode"></span>',
        ],
        aldPay6Title: 'OVO',
        aldPay6Channel: 'OVO',
        aldPay6Text: [
            'Pay in OVO app',
            'Transaction ID: <span class="payCode"></span>',
            'You will receive a payment message from OVO.',
            'Click order and complete payment. ',
            "Click 'Check Payment Status' after paid.",
        ],
        aldPay7Title: 'CIMB Niaga / CIMB Niaga Syariah',
        aldPay7Channel: 'CIMB_NIAGA',
        aldPay7Text: [
            '<span class="under_line">CIMB Clicks</span>',
            '1. Log in to CIMB Click (<a href="https://www.cimbclicks.co.id/" style="color: #3062ff;">https://www.cimbclicks.co.id/</a>)',
            '2. Select “Bill payment',
            '3. Select Source Account and Payment Type > Virtual Account',
            '4. Enter your CIMB virtual account number',
            '5. Your number, virtual account number and billing amount will be displayed on the screen',
            '6. Enter the 6-digit PIN code and click the Confirm button',
            '7. Confirm the payment displayed on the screen',
            '<span class="under_line">CIMB Mobile - Go Mobile</span>',
            '1. Log in to Go Mobile',
            '2. Select "Transfer"',
            '3. Select transfer to other CIMB Niaga account',
            '4. Select your source account: CASA or Mobile Account',
            '5. Select CASA and enter your virtual account number <span class="payCode"></span>',
            '6. Enter the billing amount',
            '7. Enter your mobile banking PIN',
            '8. Confirm the payment displayed on the screen',
            '<span class="under_line">ATM CIMB</span>',
            '1. Insert ATM card and enter your CIMB PIN',
            '2. Select Payment > Continue > Virtual Account',
            '3. Enter your virtual account number <span class="payCode"></span>',
            '4. Select Debit Account',
            '5. Your number, virtual account name and payment amount will be displayed on the screen',
            '6. Select OK to confirm payment',
            '7. Confirm the payment displayed on the screen',
        ],

        wypPay9Title: 'ALFA Group',
        wypPay9Channel: 'ALFAGROUP',
        wypPay9Text: [
            '1. Note the payment code and come to Alfamart, Alfa Midi, Alfa Express,Lawson or DAN+DAN closest ',
            '2. Tell the cashier that you want to pay to PLASAMALL ',
            '3. If the cashier does not know about PLASAMALL payment, inform the cashier to open the e- transaction terminal, select "no. 2 Payment menu or search for it on the menu search " ',
            '4. Ask the cashier to press "Shift" and "?" then type PLASAMALL, then click OK ',
            '5. Ask the cashier to select "no.1 PLASAMAL L- PLASAMALL Payment" ',
            '6. The cashier will ask for the payment code. Provide your payment code, e.g. 10013xxxxxxx. The cashier will confirm consumer data in the form of Merchant Name, Consumer Name, and Nominal. Make a payment to the cashier at the nominal amount specified ',
            '7. Receive receipt as proof of payment has been successfully carried out. Payment notification will be immediately accepted by the Merchant ',
            '8. Done',
        ],
        wypPay10Title: 'PERMATA VA',
        wypPay10Channel: 'PERMATA',
        wypPay10Text: [
            '<span class="under_line">1. ATM PERMATA  </span>',
            '• Select the Other Transaction menu ',
            '• Select Payment ',
            '• Choose Other Payments ',
            '• Select Virtual Account ',
            '• Enter the virtual account number ',
            '• A payment confirmation screen will appear ',
            '<span class="under_line"2. MOBILE INTERNET </span>',
            '• Enter User ID & Password ',
            '• Select Bill Payment ',
            '• Select a virtual account ',
            '• Enter the virtual account number ',
            '• Input nominal payment ',
            '• If the data matches the transaction / Token authentication input ',
            '<span class="under_line"3. ATM BERSAMA, ALTO& PRIMA </span>',
            '• Select the Transfer menu ',
            '• Choose Transfer to Other Bank If through ATM Bersama& Alto: Enter Permata Bank code (013) +16 digit pay code If via ATM Prima: Enter the PermataBank code (013) press BENAR, followed by the 16 digit pay code ',
            '• Input nominal payment ',
            '• The payment name & nominal payment confirmation screen will appear',
        ],
        wypPay11Title: 'BRI VA',
        wypPay11Channel: 'BRI',
        wypPay11Text: [
            '<span class="under_line">1. BRImo  </span>',
            '• Login to the BRImo application (enter username and password) ',
            '• Select the BRIVA menu ',
            '• Select a source of funds ',
            '• Input BRIVA code ',
            '• Confirmation and payment process ',
            '<span class="under_line">2. BRI INET BANKING  </span>',
            '• Input card and PIN ',
            '• Select the menu Other Transactions> Payment> Others> BRIVA ',
            '• Input the BRIVA code ',
            '• Confirmation and payment process ',
            '<span class="under_line">3. ATMBRI • Input card and PIN ',
            '• Select the menu Other Transactions> Payment> Others> BRIVA ',
            '• Input the BRIVA code ',
            '• Confirmation and payment process ',
            '<span class="under_line">4. ATM BANK LAINNYA  </span>',
            '• Input card and PIN ',
            '• Select the menu Other Transactions> Transfer> To Other Bank Accounts ',
            '• Input bank code 002 ',
            '• Input nominal transfer ',
            '• Input the BRIVA code ',
            '• Select the account to be debited ',
            '• Confirmation and process payments',

            '<span class="under_line">5. MOBILE BANKING BRI  </span>',
            '• Select the Info menu> BRIVA Info ',
            '• Input BRIVA code ',
            '• Input password ',
            '• Confirmation and payment process ',
            '<span class="under_line">6. BRI BY TELLER  </span>',
            '• Fill out the form according to BANK BRI regulations ',
            '• Give forms and money to teller ',
            '• Receive the receipt as proof of successful payment',

        ],
        wypPay12Title: 'CIMB VA',
        wypPay12Channel: 'CIMB_NIAGA',
        wypPay12Text: [
            '<span class="under_line">1. CIMB ATM </span>',
            "• Transfer menu ",
            "• Select Between CIMB NIAGA Accounts ",
            "• Input nominal transfer ",
            "• Input the virtual account number in the destination account ",
            "• Confirmation and payment process ",
            '<span class="under_line">2. CIMB CLICKS • Transfer menu </span>',
            "• Input nominal transfer ",
            "• Select Other Account - Bank CIMB Niaga / Rekening Ponsel ",
            "• The recipient's account is filled with a virtual account number ",
            "•Input MPIN ",
            "• Confirmation and payment process ",
            '<span class="under_line">3. GO MOBILE </span>',
            "• Transfer menu ",
            "• Select Transfer to Other CIMB Niaga Account ",
            "• Input virtual account number and payment nominal ",
            "• PIN input ",
            "• Confirmation and payment process ",
            '<span class="under_line">4. OTHER INTERNET BANKING </span>',
            "• Transfer menu ",
            "• Choose Another Bank Account ",
            "• Input the virtual account number in the destination account ",
            "• Input nominal payment ",
            "• Enter beneficiary bank input with CIMB Niaga ",
            "• The type of transfer used is online / realtime ",
            "• Confirmation and payment process ",
            '<span class="under_line">5. OTHER BANK ATMs </span>',
            "• Transfer menu ",
            "• Choose Another Bank Account ",
            "• Enter bank code 022 ",
            "• Input the virtual account number in the destination account ",
            "• Input nominal payment ",
            "• Confirmation",

        ],

        wypPay13Title: 'MANDIRI VA',
        wypPay13Channel: 'MANDIRI',
        wypPay13Text: [
            '<span class="under_line">1. ATM MANDIRI </span>',
            '• Select the Pay / Buy menu ',
            '• Select the Multipayment menu ',
            '• Enter the Bank Mandiricompany code(5 digits before payment code), press Correct ',
            '• Enter the virtual account number then press Correct ',
            '• Select Number 1 or according to the bill to be paid then press YES ',
            '<span class="under_line">2. ONLINE BANKING </span>',
            '• Enter User ID & Password ',
            '• Select the Payment menu ',
            '• Select the Multipayment menu ',
            '• In the Service Provider Column, select the corresponding service provider(Enter the first 5 digits of the payment code) ',
            '• Enter the virtual account number ',
            '• Select Continue ',
            '• Web: ',
            'a. OTP sent to mobile number. ',
            'b. The OTP obtained is inputted inToken to getchallenge code ',
            'c. Chanllenge code input to Mandiri online, then click Continue ',
            '• Apps: ',
            'a. After the transaction confirmation appears Select Send ',
            'b. Input MPIN (6 digits)',

        ],

    },
    MX: {
        title: '',
        contTitle: 'Instrucciones de pago',
        otherPay: '',
        contTips: '',
        xPay0Title: 'SPEI',
        xPay0Channel: 'SPEI',
        xPay0Text: [
            '<span class="under_line">SPEI</span>',
            '1. Abra la aplicación bancaria en su móvil.',
            '2. Seleccione "Transferencia rápida".',
            '3. Introduzca la CLABE de 18 dígitos.',
            '4. En "Importe", rellene la cantidad de pago correspondiente.',
            '5. Haga clic en "Continuar',
            '6. Confirme que los datos sean correctos y seleccione "Transferencia".',
            '7. Intoduzca su señal o contraseña para completar el pago.',
            '8. La transferencia está completa.',
        ],
    }
};
// 充值信息
var rechargeInfo = {
    INSA: {
        title: 'Informasi isi ulang',
        regTitle0: 'Akun Nyata',
        regTitle1: 'Nominal deposit',
        regTitle2: 'Cara pembayaran',
        regTitle3: 'Nomor pesanan bisnis',
    },
    EN: {
        title: 'Deposit information',
        regTitle0: 'Demo Account',
        regTitle1: 'Deposit amount',
        regTitle2: 'Payment method',
        regTitle3: 'Merchant order number',
    },
    MX: {
        title: 'Deposit information',
        regTitle0: 'Demo Account',
        regTitle1: 'Deposit amount',
        regTitle2: 'Payment method',
        regTitle3: 'Merchant order number',
    }
};
// 疑难解答
var troubleshoot = {
    INSA: {
        title: 'Masalah umum',
        tabArr0Title: 'Isi ulang',
        tabArr1Title: 'Trading',
        tabArr2Title: 'Tarik tunai',
        contentRechargeArr0Title: '1. Apakah ada biaya administrasi yang dikenakan pada saat isi ulang?',
        contentRechargeArr0Text: 'Tidak, perusahaan tidak akan mengambil biaya penyimpanan Anda.',
        contentRechargeArr1Title: '2. Berapa jumlah minimum isi ulang?',
        contentRechargeArr1Text: 'Silakan isi ulang berdasarkan tampilan saat pengisian ulang.',
        contentRechargeArr2Title: '3. Bagaimana cara menyelesaikan isi ulang?',
        contentRechargeArr2Text: 'Untuk menyetorkan penyimpanan ke akun Anda, Anda harus memilih tampilan "Penyimpanan" di menu. Pilih jumlah dana, ada banyak cara untuk mengisi ulang. Biasanya dana masuk dalam 1 menit. Jika Anda memiliki pertanyaan, silakan hubungi Layanan Pelanggan, kami akan menanganinya sesegera mungkin.',
        contentTradeArr0Title: '1. Disaat aplikasi muncul tampilan sistem eror apa yang harus Saya lakukan?',
        contentTradeArr0Text: 'Jika ada kesalahan teknis, pertama-tama Anda disarankan untuk menghapus cache dan cookie di browser Anda. Anda juga harus memeriksa apakah aplikasi yang sedang Anda gunakan sekarang adalah versi terbaru. Jika cara-cara ini dan pemeriksaan internet tidak membantu, dan masih ada error, silakan menghubungi Layanan Pelanggan.',
        contentTradeArr1Title: '2. Disaat apakah waktu trading paling aktif dalam sehari?',
        contentTradeArr1Text: 'Kegiatan transaksi mata uang di tentukan Eropa dan Amerika, contoh,waktu aktif di Eropa adalah dari jam 06:00 AM sampai jam 03:00 PM UTC. Waktu aktif Amerika adalah dari jam 01:00 AM sampai jam 10:00 PM UTC.',
        contentTradeArr2Title: '3. Bagaimana cara Saya memeriksa riwayat aliran transaksi Saya?',
        contentTradeArr2Text: 'Anda dapat melihat riwayat transaksi Anda melalui "Riwayat Pesanan".',
        contentTradeArr3Title: '4. Apakah ada akun demo?',
        contentTradeArr3Text: 'Kami akan menyiapkan 1.000.000 mata uang virtual di akun demo Anda, Anda dapat menggunakannya untuk mensimulasikan transaksi di platform.',
        contentWithdrawArr0Title: '1. Berapa jumlah penarikan minimum?',
        contentWithdrawArr0Text: 'Silakan tarik dana berdasarkan tampilan informasi yang diminta pada saat penarikan.',
        contentWithdrawArr1Title: '2. Kapan uang akan dipotong dari saldo?',
        contentWithdrawArr1Text: 'Setelah Anda mengajukan penarikan dana, jumlah penarikan dana akan ditetapkan. Setelah berhasil penarikan dana ke akun Anda, jumlah dana yang ditarik akan dipotong dari saldo Anda.',
        contentWithdrawArr2Title: '3. Bagaimana cara saya penarikan dana?',
        contentWithdrawArr2Text: 'Untuk penarikan dana, Anda perlu mengklik tombol "Penarikan Dana" di menu, pilih jumlah dana, dan klik "Konfirmasi penarikan dana". Biasanya, 90% permintaan akan diproses dalam waktu 24 jam atau diterima pada hari kerja berikutnya. Waktu penarikan maksimum adalah 5 hari kerja, tetapi kami akan selalu memproses permintaan penarikan Anda sesegera mungkin.',
        contentWithdrawArr3Title: '4. Apakah saya bisa menentukan informasi tentang dompet digital atau akun sistem pembayaran saya?',
        contentWithdrawArr3Text: 'Platform ini mendukung penarikan dana langsung ke rekening bank Anda. Untuk penarikan dana, Anda perlu mengisi data rekening bank pribadi Anda, kemudian mengajukan permintaan pada menu.',
        contentWithdrawArr4Title: '5. Apakah ada biaya yang dikenakan pada saat penarikan dana?',
        contentWithdrawArr4Text: 'Ya, tetapi kami tidak mengenakan biaya layanan apapun, semua biaya dikumpulkan oleh perusahaan pembayaran digital.',
    },
    EN: {
        title: 'FAQ',
        tabArr0Title: 'Deposit',
        tabArr1Title: 'Trading',
        tabArr2Title: 'Withdrawal',
        contentRechargeArr0Title: '1. Is there any extra fee charges for deposit?',
        contentRechargeArr0Text: 'No, the company will not charge you for your deposit.',
        contentRechargeArr1Title: '2. What is the minimum deposit amount?',
        contentRechargeArr1Text: 'Minimum: ${common.EN.currencySymbol}100.',
        contentRechargeArr2Title: '3. How do I complete a deposit?',
        contentRechargeArr2Text: 'You need to select the "Deposit" section in the dashboard in order to make a deposit to your account. Select the amount, there are many deposit options. That will take 1 minute to complete deposit.',
        contentTradeArr0Title: '1. What should I do when system error pops up in the app?',
        contentTradeArr0Text: 'If there is a technical error, it is first recommended to clear the cache and cookies in your browser. You should also check wheather you are using the latest version of the product. If these actions and network checks do not help and the error persists, please contact customer service.',
        contentTradeArr1Title: '2. What is the most active trading time of the day?',
        contentTradeArr1Text: "Currency trading activity depends on Europe and the United States, for example, Europe's active time is from 06: 00 AM to 03: 00 PM UTC.Active hours in the US are from 1: 00 AM to 10: 00 PM UTC.",
        contentTradeArr2Title: '3. How do I view my trading record?',
        contentTradeArr2Text: 'You can view your trading records in"History".',
        contentTradeArr3Title: '4. Is there a demo account?',
        contentTradeArr3Text: 'We will prepare 100,000 virtual currency in your demo account, you can use it to simulate trading on the platform.',
        contentWithdrawArr0Title: '1. What is the minimum withdrawal amount?',
        contentWithdrawArr0Text: 'Please refer to the instruction prompted when withdrawing.',
        contentWithdrawArr1Title: '2. When will the money be deducted from the balance?',
        contentWithdrawArr1Text: 'The withdrawal amount will be locked after you apply for withdrawal. The withdrawal amount will be deducted from the balance after successful withdrawal to your account.',
        contentWithdrawArr2Title: '3. How do I do a withdrawal?',
        contentWithdrawArr2Text: 'You need to click the "Withdraw" button in the dashboard in order to do a withdrawal, select desired amount, and click "Confirm Withdrawal". Typically, 90% of requests are processed within 24 hours or credited in the next working day. The maximum withdrawal time is 5 working days, but we will always process your withdrawal request as soon as possible.',
        contentWithdrawArr3Title: '4. Can I decide specify information about my e-wallet or payment system account?',
        contentWithdrawArr3Text: 'The platform supports direct withdrawal to your bank card. To do a withdrawal, you need to fill in your personal bank card and submit a request on the dashboard.',
        contentWithdrawArr4Title: '5. Are there any withdrawal fees?',
        contentWithdrawArr4Text: 'Yes, but we do not charge any service fees, all fees are collected by the payment company.',
    },
    MX: {
        title: 'FAQ',
        tabArr0Title: 'Deposit',
        tabArr1Title: 'Trading',
        tabArr2Title: 'Withdrawal',
        contentRechargeArr0Title: '1. Is there any extra fee charges for deposit?',
        contentRechargeArr0Text: 'No, the company will not charge you for your deposit.',
        contentRechargeArr1Title: '2. What is the minimum deposit amount?',
        contentRechargeArr1Text: 'Minimum: ${common.MX.currencySymbol}100.',
        contentRechargeArr2Title: '3. How do I complete a deposit?',
        contentRechargeArr2Text: 'You need to select the "Deposit" section in the dashboard in order to make a deposit to your account. Select the amount, there are many deposit options. That will take 1 minute to complete deposit.',
        contentTradeArr0Title: '1. What should I do when system error pops up in the app?',
        contentTradeArr0Text: 'If there is a technical error, it is first recommended to clear the cache and cookies in your browser. You should also check wheather you are using the latest version of the product. If these actions and network checks do not help and the error persists, please contact customer service.',
        contentTradeArr1Title: '2. What is the most active trading time of the day?',
        contentTradeArr1Text: "Currency trading activity depends on Europe and the United States, for example, Europe's active time is from 06: 00 AM to 03: 00 PM UTC.Active hours in the US are from 1: 00 AM to 10: 00 PM UTC.",
        contentTradeArr2Title: '3. How do I view my trading record?',
        contentTradeArr2Text: 'You can view your trading records in"History".',
        contentTradeArr3Title: '4. Is there a demo account?',
        contentTradeArr3Text: 'We will prepare 100,000 virtual currency in your demo account, you can use it to simulate trading on the platform.',
        contentWithdrawArr0Title: '1. What is the minimum withdrawal amount?',
        contentWithdrawArr0Text: 'Please refer to the instruction prompted when withdrawing.',
        contentWithdrawArr1Title: '2. When will the money be deducted from the balance?',
        contentWithdrawArr1Text: 'The withdrawal amount will be locked after you apply for withdrawal. The withdrawal amount will be deducted from the balance after successful withdrawal to your account.',
        contentWithdrawArr2Title: '3. How do I do a withdrawal?',
        contentWithdrawArr2Text: 'You need to click the "Withdraw" button in the dashboard in order to do a withdrawal, select desired amount, and click "Confirm Withdrawal". Typically, 90% of requests are processed within 24 hours or credited in the next working day. The maximum withdrawal time is 5 working days, but we will always process your withdrawal request as soon as possible.',
        contentWithdrawArr3Title: '4. Can I decide specify information about my e-wallet or payment system account?',
        contentWithdrawArr3Text: 'The platform supports direct withdrawal to your bank card. To do a withdrawal, you need to fill in your personal bank card and submit a request on the dashboard.',
        contentWithdrawArr4Title: '5. Are there any withdrawal fees?',
        contentWithdrawArr4Text: 'Yes, but we do not charge any service fees, all fees are collected by the payment company.',
    }
};
// 服务条款
var termsOfService = {
    INSA: {
        title: 'Perjanjian persyaratan',
        termsArr0Title: 'Isian: ',
        termsArr0TextArr: [
            {text: '1. Ketentuan Umum dan Subyek Perjanjian'},
            {text: '2. Istilah dan definisi'},
            {text: '3. Komunikasi dan Penyediaan Informasi'},
            {text: '4. Ketentuan penggunaan layanan SmartOption'},
            {text: '5. Klaim dan Penyelesaian Sengketa'},
            {text: '6. Peraturan pemerintah'},
            {text: '7. Keadaan Kahar'},
            {text: '8. Tanggung Jawab Para Pihak'},
            {text: '9. Jangka Waktu dan Pengakhiran Perjanjian'},
            {text: '10. Ketentuan akhir'},
            {text: '11. Daftar Negara Tempat Perusahaan Tidak Melakukan Bisnis'},
        ],
        termsArr1Title: '1. Ketentuan Umum dan Subyek Perjanjian',
        termsArr1TextArr: [
            {
                text: '1.1. Perjanjian Layanan ini telah dibuat oleh Aplikasi SmartOption (selanjutnya - SmartOption), dan orang yang mengisi formulir pendaftaran di situs web SmartOption atau di platform perdagangan dan telah menerima persyaratan Perjanjian Layanan ini dan lampirannya pada saat pendaftaran (selanjutnya disebut Klien).',
            },
            {
                text: '1.2. Perjanjian adalah undangan untuk membuat penawaran yang diposting di situs web SmartOption yang harus diperlakukan sebagai penawaran untuk masuk ke dalam Perjanjian ini dengan ketentuan yang ditetapkan di dalamnya. Penawaran yang diposting tidak bersifat publik. SmartOption atas kebijakannya sendiri dapat menolak untuk membuat Perjanjian dengan siapa pun tanpa menjelaskan alasan penolakan atau, jika pendaftaran dilakukan, untuk mengakhiri hubungan kontrak dan memblokir akses ke terminal perdagangan. Pendaftaran Klien di Situs Web SmartOption atau di terminal perdagangan dianggap sebagai penerimaan penuh dan tanpa syarat dari ketentuan Perjanjian. Segera setelah SmartOption menerima pembayaran untuk mengisi kembali akun perdagangan Klien, setiap transaksi Klien menggunakan terminal perdagangan atau Dasbor menjadi subjek dari Perjanjian ini.',
            },
            {
                text: '1.3. Klien harus hati-hati meninjau persyaratan Perjanjian. Dengan menerima persyaratan Perjanjian ini, Klien menyetujui persyaratan semua lampiran yang tercantum di atas, termasuk persyaratan pada subdomain situs web SmartOption yang dapat diakses oleh Klien dan mengonfirmasi bahwa dia adalah orang dewasa yang merupakan orang yang cakap secara hukum dan bukan merupakan penduduk negara di mana perdagangan opsi dapat dianggap ilegal. Klien juga menyatakan dan menjamin kepada SmartOption bahwa:',
                child: [
                    '1.3.1. Semua informasi yang diberikan selama pendaftaran Klien dan selama pelaksanaan Perjanjian, adalah benar, akurat, andal, dan lengkap dalam segala hal, dan Klien melengkapi formulir pendaftarannya sendiri;',
                    '1.3.2. Klien memiliki kepribadian hukum untuk masuk ke dalam Perjanjian ini, untuk mengajukan pertanyaan dan memberikan perintah, dan untuk menggunakan haknya dan untuk memenuhi kewajiban sesuai dengan ketentuan Perjanjian;',
                    '1.3.3. Klien akan melakukan transaksi perdagangan dan non-perdagangan secara pribadi, atas namanya sendiri dan atas biaya sendiri dan tidak akan menggunakan dana yang dipinjam dari Klien SmartOption lain atau dari pihak ketiga untuk melakukan transaksi. Klien akan dipandu oleh prinsip-prinsip integritas, kejujuran, dan rasionalitas; Klien tidak akan melakukan tindakan yang terkoordinasi dengan Klien SmartOption lainnya yang bertujuan merusak SmartOption; Klien tidak akan menggunakan fitur teknis dari pembaruan aliran kutipan di terminal perdagangan dan tidak akan menggunakan kesalahan perangkat lunak, cacat, dan kerentanan yang ditemukannya di terminal perdagangan untuk mengekstrak pendapatan dan tidak akan mendistribusikan informasi tentang kerentanan kepada pihak ketiga.Klien tidak akan menggunakan metode atau cara yang tidak adil dan tidak jujur ​​untuk melakukan perdagangan (transaksi) dengan SmartOption; Klien tidak akan menggunakan informasi orang dalam atau rahasia atau informasi lainnya, sebagai akibat dari penggunaannya yang mungkin menguntungkan Klien saat berdagang dengan SmartOption dan/atau yang dapat merusak SmartOption;',
                    '1.3.4. Klien akan mematuhi norma-norma hukum, khususnya, namun tidak terbatas pada, norma-norma internasional yang ditujukan untuk mengendalikan perdagangan ilegal, penipuan keuangan, dan pencucian uang;',
                    '1.3.5.Klien tidak akan menggunakan terminal perdagangan atau situs web untuk berkolusi dalam aktivitas keuangan ilegal atau transaksi ilegal lainnya;',
                    '1.3.6.Uang yang terdaftar oleh Klien di akun SmartOption memiliki asal yang sah.Klien secara hukum memiliki uang dan memiliki hak untuk menggunakan dan mengelolanya.Akun Klien tidak akan diisi ulang dari instrumen pembayaran pihak ketiga.Klien tidak akan mengisi kembali akun Klien pihak ketiga atau menarik uang dari akun Klien ke instrumen pembayaran pihak ketiga.',
                    '1.3.7.Tidak ada tindakan Klien berdasarkan Perjanjian yang akan melanggar hukum, peraturan, hak, peraturan daerah, atau aturan dan regulasi yang berlaku untuk Klien atau di yurisdiksi tempat tinggalnya atau ketentuan perjanjian lain yang mengikat Klien atau melibatkan setiap aset Klien',
                    '1.3.8.Untuk melakukan transaksi, Klien akan menggunakan data akun dari Terminal Perdagangan miliknya dan tidak akan mentransfer data akun ke pihak ketiga dan tidak akan menggunakan data akun klien SmartOption untuk operasi perdagangan dan(atau) non- perdagangan;',
                    '1.3.9.Klien bukan pegawai federal atau kotamadya, pegawai lembaga federal atau kotamadya, pegawai organisasi federal atau kotamadya, organisasi yang negara bagian memiliki kepentingan didalamnya; Klien bukan orang penting secara politik, anggota keluarga atau kerabat orang penting secara politik; Klien bukan orang yang terkait erat dengan orang yang signifikan secara politik; Klien bukanlah orang yang terhubung dengan Amerika Serikat atau dengan negara lain di mana SmartOption tidak beroperasi.Istilah yang digunakan dalam paragraf ini ditafsirkan dan diterapkan oleh SmartOption atas kebijakannya sendiri sesuai dengan norma undang- undang internasional dan / atau undang - undang negara bagian tertentu, istilah dan definisi yang diterima secara umum, praktik bisnis yang lazim.',
                ]
            },
            {
                text: '1.4. Subyek Perjanjian adalah definisi kondisi umum di mana Para Pihak melakukan transaksi (perdagangan), isi dan prosedur yang diatur dalam Perjanjian ini. SmartOption secara sepihak menetapkan dan dapat mengubah atas kebijakannya sendiri kondisi penting untuk transaksi (perdagangan), dapat membatasi jumlah perdagangan yang dieksekusi secara bersamaan, dan membatasi jumlah perdagangan yang dapat dilakukan Klien dalam jangka waktu yang ditetapkan oleh SmartOption, dapat menetapkan batasan lain untuk perdagangan atas kebijakannya sendiri dan secara sepihak.',
            },
            {
                text: '1.5. SmartOption dapat melibatkan pihak ketiga untuk melakukan Perjanjian ini. SmartOption tidak bertanggung jawab atas layanan yang disediakan oleh pihak ketiga tersebut.',
            }
        ],
        termsArr2Title: '2. Istilah dan definisi',
        termsArr2TextArr: [
            {
                text: '2.1. Aset - instrumen keuangan yang mendasari sebuah opsi. Aset dapat berupa saham perusahaan, indeks saham, pasangan mata uang (nilai tukar satu mata uang ke mata uang lain), barang yang dikutip di bursa komoditas, opsi komoditas, dll.',
            },
            {
                text: '2.2. Saldo Akun Klien - jumlah uang di akun Klien.',
            },
            {
                text: '2.3. Opsi - instrumen keuangan derivatif (perdagangan), yang terdiri dari dua transaksi: Melakukan perdagangan dan Menutup perdagangan, sebagai akibat dari pemenuhan yang mana Klien menerima Penghasilan atau kehilangan jumlah Perdagangan.SmartOption tidak, bagaimanapun, memberikan Klien kemampuan untuk menanggung kewajiban yang lebih besar dari jumlah Perdagangan (SmartOption tidak memberikan leverage kepada Klien). SmartOption dan Klien melakukan perdagangan dengan opsi Lebih Tinggi/Lebih Rendah - opsi, tingkat yang ditargetkan sama dengan harga aset atau harga lain yang ditentukan oleh Klien atas saran SmartOption saat perdagangan dilakukan, hanya pembayaran yang dilakukan jika perubahan harga ke atas dipilih dan harga aset lebih besar dari level yang ditargetkan saat opsi kedaluwarsa atau jika perubahan harga ke bawah dipilih dan harga aset kurang dari level yang ditargetkan saat opsi kedaluwarsa. Jika harga aset sama dengan level yang ditargetkan saat opsi kedaluwarsa, Jumlah perdagangan dikembalikan ke Klien.',
            },
            {
                text: '2.4. Bonus - uang tunai virtual yang dikreditkan oleh SmartOption ke Akun Klien ketika Klien memenuhi ketentuan yang ditetapkan oleh SmartOption. Ketika Klien melakukan transaksi perdagangan, dana yang disetorkan oleh Klien harus digunakan terlebih dahulu dan, hanya setelah dana tersebut benar-benar habis, Klien berhak menggunakan Bonus untuk transaksi perdagangan lebih lanjut. Secara umum, Klien tidak dapat menarik jumlah Bonus ke Akun Eksternal Klien. Jika Klien menarik dana yang sebelumnya dikreditkan ke akun Eksternalnya, SmartOption berhak mendebit seluruh jumlah Bonus yang diperoleh sebelumnya dari Akun Klien. SmartOption berhak untuk menetapkan jumlah perdagangan minimum dan persyaratan lain di mana Bonus tidak didebit dari Akun Klien dan dapat ditarik ke akun Eksternal Klien. SmartOption dapat menerbitkan aturan kredit dan debit, aturan untuk menggunakan dan mengonversi Bonus menjadi dana nyata, dan aturan untuk transaksi lain dengan Bonus dan dana Klien yang diperoleh Klien bersama dengan Bonus di situs webnya. ',
            },
            {
                text: '2.5. Situs Web SmartOption - situs internet di alamat (nama domain) SmartOption.com, dan situs internet lainnya yang dirujuk oleh SmartOption dalam Perjanjian atau di Situs Web SmartOption.',
            },
            {
                text: '2.6. Akun Eksternal Klien - akun Klien saat ini di lembaga pemberi pinjaman, akun (dompet) pada sistem pembayaran elektronik.',
            },
            {
                text: '2.7. Waktu Kedaluwarsa Opsi - waktu ketika kondisi pembayaran pendapatan pada opsi dicentang.',
            },
            {
                text: '2.8. Jam Perdagangan - waktu saat perdagangan aset dapat dilakukan.',
            },
            {
                text: '2.9. Penarikan Dana - transfer uang dari Akun Klien ke Akun Eksternal Klien.',
            },
            {
                text: '2.10. Pendapatan - remunerasi tetap yang dikreditkan ke Akun Klien jika kondisi opsi terpenuhi. Jumlah Pendapatan ditentukan pada saat perdagangan dieksekusi. Pendapatan dikreditkan ke akun Klien dari dana jaminan (kompensasi, khusus) SmartOption (dana cadangan), yang dibentuk dengan cara yang ditentukan oleh SmartOption atas biayanya sendiri.',
            },
            {
                text: '2.11. Eksekusi perdagangan - transaksi perdagangan di mana kondisi penting dari perdagangan opsi disetujui oleh Klien dan SmartOption. Setelah kondisi penting disepakati, perdagangan opsi dibuka.',
            },
            {
                text: '2.12. Menutup perdagangan - transaksi perdagangan di mana perdagangan terbuka pada opsi dilikuidasi pada waktu kedaluwarsa opsi yang telah disepakati oleh Para Pihak saat perdagangan dilakukan. Penutupan awal perdagangan - menutup perdagangan atas inisiatif Klien sebelum ditutup pada waktu kedaluwarsa opsi. Penutupan awal perdagangan hanya dimungkinkan jika SmartOption memiliki kemampuan teknis untuk melakukannya. SmartOption dapat secara sepihak atas kebijakannya sendiri memberikan atau menolak untuk memberikan hak kepada Klien untuk melakukan penutupan perdagangan lebih awal.',
            },
            {
                text: '2.13. Situasi Luar Biasa - kondisi pasar yang tidak normal dan kondisi luar biasa lainnya yang dijelaskan dalam Peraturan tentang Transaksi Perdagangan.',
            },
            {
                text: '2.14. Rasio pembayaran - persentase yang menentukan jumlah Pendapatan yang ditetapkan oleh SmartOption tergantung pada aset opsi dan kondisi opsi lainnya.',
            },
            {
                text: '2.15. Kutipan Aset - harga untuk satu unit Aset.',
            },
            {
                text: '2.16. Entri log - entri dalam database yang dibuat oleh server SmartOption yang, dengan akurasi hingga milidetik atau, jika tidak ada kemampuan teknis, dengan akurasi hingga detik, mendokumentasikan semua permintaan dan pesanan Klien serta hasil pemrosesannya . Setiap kontak Klien dengan terminal perdagangan dan dasbor didokumentasikan dalam entri log. Data server ini merupakan sumber informasi utama dan digunakan oleh para pihak sebagai bukti jika timbul perselisihan terkait pelaksanaan Perjanjian. Informasi dari entri Log server SmartOption memiliki prioritas tanpa syarat di atas semua argumen lain selama penyelesaian perselisihan, termasuk informasi dari file Log terminal perdagangan Klien. SmartOption berhak untuk tidak menyimpan entri Log.',
            },
            {
                text: '2.17. Arah Harga - kondisi penting dari perdagangan opsi yang menentukan pembayaran perdagangan. Sebuah harga dapat berubah arah Lebih Tinggi atau Lebih Rendah.',
            },
            {
                text: '2.18. Transaksi Non-perdagangan - setiap transaksi Klien untuk menyetor dana ke Akun Klien atau menarik dana dari Akun Klien.',
            },
            {
                text: '2.19. Transaksi - Transaksi perdagangan dan non-perdagangan Klien.',
            },
            {
                text: '2.20. Perdagangan Terbuka - perdagangan opsi setelah pelaksanaan perdagangan dan sebelum penutupan perdagangan, yang belum ditentukan apakah pembayaran akan dilakukan.',
            },
            {
                text: '2.21. Umpan Harga - rangkaian harga yang ditampilkan di Terminal Perdagangan.',
            },
            {
                text: '2.22. Penyedia layanan pembayaran - perusahaan yang menyediakan layanan transfer dana.',
            },
            {
                text: '2.23. Pembayaran berulang - transaksi berulang secara berkala untuk mengisi kembali Saldo akun Klien tanpa mengharuskan yang berikutnya untuk memasukkan kembali informasi kartu kreditnya.',
            },
            {
                text: '2.24. Server SmartOption - perangkat lunak SmartOption yang digunakan untuk memproses dan menyimpan informasi tentang permintaan klien untuk transaksi perdagangan dan non-perdagangan, untuk memberikan kepada Klien informasi realtime tentang kutipan, untuk memperhitungkan transaksi perdagangan dan non-perdagangan, untuk memantau pemenuhan kondisi perdagangan dan membatasi transaksi perdagangan, dan untuk menentukan hasil keuangan perdagangan.',
            },
            {
                text: '2.25. Metode Penarikan - salah satu metode Penarikan Dana yang ditawarkan kepada Klien yang dipasang di terminal perdagangan dan di Dasbor.',
            },
            {
                text: '2.26. Jumlah Perdagangan adalah jumlah yang dibayarkan oleh Klien ke SmartOption saat perdagangan dilakukan. Keuntungan Klien pada perdagangan opsi jika kondisi perdagangan terpenuhi ditentukan oleh prosedur yang ditetapkan dalam Bagian 2.10. dari Perjanjian ini.',
            },
            {
                text: '2.27. Akun SmartOption - akun penyelesaian SmartOption di lembaga keuangan, akun (dompet) dalam sistem pembayaran elektronik, dan akun lainnya, termasuk akun Agen Pembayaran.',
            },
            {
                text: '2.28. Ketentuan Esensial dari transaksi perdagangan (kondisi penting perdagangan) - ketentuan yang mengatur pembayaran pendapatan dari perdagangan ke Klien oleh SmartOption.',
            },
            {
                text: '2.29. Akun Klien (Akun Perdagangan) - akun khusus dalam sistem akuntansi SmartOption di mana dana yang ditransfer oleh Klien untuk melakukan perdagangan diposting; dari mana jumlah perdagangan didebit saat perdagangan dieksekusi; dan kemana pendapatan dikreditkan ketika perdagangan ditutup dan kondisi penting perdagangan terpenuhi. Klien hanya berhak memiliki satu akun Klien. Dalam kasus pelanggaran aturan ini, SmartOption berhak untuk menolak Klien dalam layanan lebih lanjut (rendering layanan), untuk mengakhiri Perjanjian ini, dan untuk memblokir kemungkinan lebih lanjut untuk melakukan transaksi tanpa penjelasan dan tanpa pembayaran dana dari akun klien. Bukan merupakan pelanggaran terhadap klausul Perjanjian ini ketika SmartOption, jika mungkin dan secara sepihak, memberi Klien hak untuk menggunakan beberapa mata uang dalam akun Klien, serta hak untuk menggunakan akun Klien dalam hubungan antara SmartOption dan Klien, yang diatur secara bersamaan oleh Perjanjian ini dan perjanjian lainnya, yang dibuat antara SmartOption dan Klien, di mana SmartOption memiliki keleluasaan untuk memberi Klien hak untuk menggunakan akun Klien untuk terlibat dalam transaksi perdagangan yang tidak diatur dalam Perjanjian ini.',
            },
            {
                text: '2.30. Transaksi perdagangan - prosedur untuk membuat dan menutup perdagangan dengan opsi antara SmartOption dan Klien. Transaksi perdagangan dilakukan di tempat pendaftaran SmartOption. Tidak ada pengiriman fisik aset yang terjadi selama transaksi perdagangan. Jumlah perdagangan untuk transaksi perdagangan didebit dari saldo akun Klien setelah perdagangan dilakukan. Pendapatan dari transaksi perdagangan dikreditkan ke akun Klien segera setelah perdagangan ditutup.',
            },
            {
                text: '2.31. Terminal Perdagangan adalah perangkat lunak di mana Klien dapat memperoleh informasi real-time tentang kutipan harga untuk melakukan transaksi perdagangan dan non-perdagangan dan menerima pesan dari SmartOption. Masuk ke Terminal Perdagangan dilindungi oleh kata sandi yang ditetapkan Klien saat pendaftaran di situs SmartOption. Semua pesanan dan pertanyaan yang dilakukan melalui Terminal Perdagangan dianggap dilakukan secara pribadi oleh Klien. Klien dari negara yang undang-undangnya melarang perdagangan opsi atau derivatif bebas lainnya dan karyawan, afiliasi, agen, dan perwakilan lain dari SmartOption dan kerabat mereka dilarang menggunakan Terminal Perdagangan. Bagian dari Terminal Perdagangan yang dengannya Klien dapat melakukan transaksi non-perdagangan dalam Perjanjian ini dapat disebut Dasbor.',
            },
            {
                text: '2.32. Wilayah Merah - keadaan perdagangan opsi terbuka jika, pada harga aset saat ini, pendapatan tidak dapat dibayarkan pada perdagangan itu.',
            },
            {
                text: '2.33. Layanan sekali klik - memungkinkan untuk mengisi kembali saldo akun Klien dari kartu bank (pembayaran) Klien tanpa memasukkan data kartu bank (pembayaran) pemegang kartu bank lagi. ',
            },
            {
                text: '2.34. Tingkat yang ditargetkan - tingkat harga aset sehubungan dengan hasil perdagangan yang dihitung.',
            },
            {
                text: '2.35. File cookie - kumpulan data kecil termasuk pengidentifikasi unik anonim yang dikirim ke browser web komputer atau telepon seluler Klien (selanjutnya disebut "perangkat") dari server SmartOption (situs web) dan disimpan di perangkat Klien. Klien dapat mengonfigurasi browser web untuk memblokir akses cookie ke perangkat. Saat Klien mengunjungi Situs Web SmartOption, halaman yang dilihat dan file cookie diunduh ke perangkat Klien. File cookie yang disimpan di perangkat dapat digunakan untuk identifikasi anonim jika Klien mengunjungi kembali situs web SmartOption dan untuk menentukan halaman situs web yang paling populer di antara Klien kami. File cookie yang disimpan di hard drive perangkat Klien memungkinkan SmartOption membuat situs web yang paling ramah pengguna dan efisien untuk Klien SmartOption, memberikan kami kesempatan untuk mengidentifikasi preferensi Klien kami.',
            },
            {
                text: '2.36. Sinyal perdagangan adalah informasi tentang keadaan pasar yang dikumpulkan oleh SmartOption berdasarkan kesimpulan analitis bahwa SmartOption berhak untuk memberikan kebijaksanaannya kepada beberapa atau semua Klien sehubungan dengan indikator pasar tertentu. Sinyal perdagangan bukan merupakan penawaran dan bukan merupakan rekomendasi eksplisit SmartOption kepada Klien untuk terlibat dalam transaksi perdagangan atau melakukan perdagangan. SmartOption tidak bertanggung jawab atas keakuratan Sinyal Perdagangan atau untuk transaksi perdagangan Klien dan perdagangan berdasarkan Sinyal Perdagangan. Klien dapat, atas kebijaksanaannya sendiri, mempertimbangkan atau mengabaikan sinyal Perdagangan saat terlibat dalam transaksi perdagangan dan perdagangan.',
            },
            {
                text: '2.37. Quote - harga Aset saat ini yang ditampilkan di Terminal Perdagangan. Istilah yang digunakan dalam Perjanjian ini dan definisi yang tidak ada dalam Klausul ini harus ditafsirkan sesuai dengan praktik bisnis biasa dan kebiasaan yang berlaku sehubungan dengan kesimpulan perjanjian dengan derivatif keuangan.',
            },
        ],
        termsArr3Title: '3. Komunikasi dan Penyediaan Informasi',
        termsArr3TextArr: [
            {
                text: '3.1. Untuk berkomunikasi dengan Klien, SmartOption dapat menggunakan:',
                child: [
                    '· email;',
                    '· fax; ',
                    '· telepon; ',
                    '· SMS; ',
                    '· Surat yang dikirimkan via pos; ',
                    '· berbagai jenis pesan yang dikirim ke Klien di terminal perdagangan, di Dasbor, jendela browser, dll. (pemberitahuan push, pengingat, pesan layanan, dll.);',
                    '· pengumuman di situs web SmartOption.',
                ]
            },
            {
                text: '3.2. Untuk segera berkomunikasi dengan Klien untuk menyelesaikan masalah yang terkait dengan transaksi Klien, SmartOption akan menggunakan informasi kontak Klien yang dimasukkan saat Klien mendaftarkan atau memodifikasi akun sesuai dengan klausul 4.5 pada Perjanjian. Klien setuju untuk menerima pesan dari SmartOption setiap saat.',
            },
            {
                text: '3.3. Setiap korespondensi (dokumen, pemberitahuan, konfirmasi, pengumuman, laporan, dll.) dianggap telah diterima oleh Klien:',
                child: [
                    '1) satu (1) jam setelah dikirim ke alamat email (e-mail);',
                    '2) segera setelah mereka difaks;',
                    '3) segera setelah panggilan telepon berakhir;',
                    '4) segera setelah pesan SMS dikirim;',
                    '5) tujuh (7) hari kalender setelah dikirimkan via pos;',
                    '6) segera setelah pengumuman diposting di situs web SmartOption.',
                ]
            },
            {
                text: '3.4. Klien juga dapat menghubungi SmartOption melalui email di sssmartoption@gmail.com, dan alamat email lainnya serta nomor telepon yang tercantum dalam Perjanjian ini dan di situs web SmartOption.',
            },
            {
                text: '3.5. Klien memahami dan menyetujui bahwa, jika perilaku Klien selama percakapan dengan perwakilan SmartOption tidak pantas, SmartOption berhak untuk mengakhiri Perjanjian ini secara sepihak.',
            },
            {
                text: '3.6. SmartOption dapat menggunakan informasi kontak yang disediakan oleh Klien untuk mengirim materi informasi, pemasaran, dan iklan, serta pesan layanan dan untuk menyelesaikan tugas lainnya. SmartOption akan menentukan frekuensi pengiriman pesan ke Klien atas kebijakannya sendiri. Jika Klien tidak ingin menerima dari pesan informasi (dan lainnya) dari SmartOption, dia harus berhenti berlangganan dengan mengklik tautan Berhenti Berlangganan (jika format pesan menyediakan kemampuan ini) atau dengan menghubungi departemen dukungan klien.',
            },
        ],
        termsArr4Title: '4. Ketentuan penggunaan layanan SmartOption.',
        termsArr4TextArr: [
            {
                text: '4.1. Selama pendaftaran, Klien berjanji untuk memberikan informasi identifikasi yang benar dan dapat diandalkan sesuai dengan persyaratan formulir pendaftaran Klien.',
            },
            {
                text: '4.2. Setelah pendaftaran berhasil, Klien akan diberikan akses ke terminal perdagangan, kemampuan untuk mentransfer dana ke Akun Klien (menempatkan deposit di Akun Klien untuk dapat melakukan perdagangan opsi) dan untuk melakukan operasi lainnya. ',
            },
            {
                text: '4.3. Klien harus segera memberi tahu SmartOption tentang perubahan identifikasi dan informasi kontak (dalam waktu 7 (tujuh) hari setelah perubahan itu) dengan membuat perubahan yang sesuai di Terminal Perdagangan atau dengan cara lain apa pun yang ditawarkan oleh SmartOption. Untuk mengidentifikasi Klien dan untuk memeriksa asal dana Klien setiap saat setelah pendaftaran, SmartOption berhak untuk meminta dan Klien berkewajiban untuk menyediakan dalam waktu 7 (tujuh) hari setelah permintaan diterima, setiap dokumen identifikasi (ini dokumen termasuk dokumen identitas; dokumen yang mengonfirmasi alamat tempat tinggal, dokumen yang menunjukkan status keuangan Klien, dan dokumen lain atas kebijaksanaan SmartOption). SmartOption berhak untuk menangguhkan transaksi non-perdagangan dan/atau perdagangan di akun Klien jika ternyata informasi identifikasi Klien tidak benar atau tidak akurat; jika Klien belum memberikan dokumen yang diminta, SmartOption berhak memblokir akses Klien ke terminal perdagangan hingga prosedur identifikasi klien selesai. SmartOption juga berhak meminta Klien untuk menjalani prosedur identifikasi dengan kunjungan pribadi ke agen resmi SmartOption dan memberikan dokumen, daftarnya ditentukan oleh SmartOption atas kebijakannya sendiri.',
            },
            {
                text: '4.4. Masuk ke terminal perdagangan dilindungi kata sandi:',
                child: [
                    '4.4.1.Klien mengonfirmasi dan menyetujui bahwa akses ke terminal perdagangan akan dilindungi dengan kata sandi yang ditetapkan oleh Klien sendiri saat pendaftaran.Klien tidak dapat mentransfer kata sandi terminal perdagangan ke pihak ketiga.',
                    '4.4.2.Klien bertanggung jawab penuh atas perlindungan kata sandi dan pencegahan akses pihak ketiga yang tidak sah ke sana.',
                    '4.4.3.Semua pesanan melalui terminal perdagangan dengan kata sandi Klien akan dianggap telah dibuat oleh Klien kecuali SmartOption menetapkan sebaliknya.',
                    '4.4.4.Setiap orang yang mendapatkan akses ke terminal perdagangan dengan memasukkan kata sandi Klien akan diidentifikasi sebagai Klien kecuali ditentukan lain oleh SmartOption.',
                    '4.4.5.SmartOption tidak bertanggung jawab atas kerugian apa pun yang mungkin diderita Klien jika terjadi pencurian, kehilangan, atau pengungkapan kata sandi kepada pihak ketiga atau dalam kasus penggunaan data pendaftaran yang tidak sah oleh pihak ketiga.',
                ]
            },
            {
                text: '4.5. Klien dapat mengubah kata sandi ke terminal perdagangan sendiri atau menggunakan prosedur pemulihan kata sandi yang ditetapkan oleh SmartOption.',
            },
        ],
        termsArr5Title: '5. Klaim dan Penyelesaian Sengketa',
        termsArr5TextArr: [
            {
                text: '5.1. Para Pihak telah sepakat bahwa mereka akan melakukan segala upaya untuk menyelesaikan semua perselisihan antara SmartOption dan Klien terkait dengan transaksi, pembayaran, dan tindakan lain berdasarkan Perjanjian ini melalui negosiasi.',
            },
            {
                text: '5.2. Jika timbul perselisihan, Klien dapat mengajukan klaim/keluhan ke SmartOption dan mengirimkan pernyataan atau pemberitahuan ke SmartOption. Semua klaim/keluhan/pernyataan/pemberitahuan terkait transaksi yang dilakukan oleh Klien harus disampaikan sesuai dengan persyaratan sebagai berikut:',
                child: [
                    '5.2.1.klaim/keluhan/pernyataan/pemberitahuan harus disampaikan secara tertulis;',
                    '5.2.2. klaim/keluhan/pernyataan/pemberitahuan harus memuat informasi berikut: Nama belakang, nama depan, nama tengah (jika ada), email Klien, Nomor Rekening Klien, tanggal dan waktu perselisihan, deskripsi singkat perselisihan, tuntutan Klien; jumlah yang diklaim dan perhitungan yang dapat dipertanggungjawabkan (jika klaim dapat dinilai dengan uang); keadaan yang menjadi dasar klaim dan bukti yang mendukung keadaan tersebut, termasuk referensi ke klausul Perjanjian ini (lampiran di sini) yang dilanggar menurut pendapat Klien; daftar dokumen dan bukti lain yang dilampirkan pada klaim (pengaduan) yang disahkan oleh Klien; informasi lain yang diperlukan untuk menyelesaikan sengketa;',
                    '5.2.3. klaim/keluhan/pernyataan/pemberitahuan harus dikirimkan oleh Klien dalam waktu lima (5) hari kerja setelah kejadian yang menjadi dasar pengajuan klaim (keluhan) yang bersangkutan. Klien setuju bahwa keterlambatan dalam pengajuan klaim (keluhan) adalah alasan penolakan untuk mempertimbangkannya;',
                    '5.2.4. klaim/keluhan/pernyataan/pemberitahuan dapat dikirim melalui email ke sssmartoption@gmail.com, melalui surat terdaftar atau bersertifikat. Klaim/keluhan/pernyataan/pemberitahuan yang dibuat dan dikirim dalam bentuk lain tidak akan dipertimbangkan.',
                ]
            },
            {
                text: '5.3. Klaim/keluhan/pernyataan/pemberitahuan tidak boleh memuat:',
                child: [
                    'a) evaluasi emosional dari perselisihan;',
                    'b) pernyataan ofensif yang ditujukan kepada SmartOption;',
                    'c) kata-kata kotor.',
                ]
            },
            {
                text: '5.4. Untuk menanggapi klaim/keluhan/pernyataan/pemberitahuan, SmartOption dapat meminta dokumen dan informasi tambahan dari Klien. Klaim/keluhan/pernyataan/pemberitahuan akan ditinjau berdasarkan data yang diberikan oleh Klien dan entri log dari server SmartOption. Entri log dari Server SmartOption selalu menang atas bukti dan bukti lainnya. SmartOption tidak bertanggung jawab atas perdagangan yang tidak selesai dan tidak akan mengkompensasi kerusakan finansial atau kerugian moral yang diderita oleh Klien sehubungan dengan apa yang dianggap Klien sebagai kerugian yang hilang. Saat mempertimbangkan perselisihan, referensi Klien ke informasi dari perusahaan dan situs web lain tidak dipertimbangkan.',
            },
            {
                text: '5.5. SmartOption dapat menolak klaim/keluhan/pernyataan/pemberitahuan jika ketentuan bagian ini dilanggar.',
            },
            {
                text: '5.6. SmartOption harus mempertimbangkan klaim/keluhan/permohonan/banding dalam waktu tidak lebih dari 10 hari kerja setelah tanggal pengajuan. Istilah ini tidak termasuk waktu untuk penyediaan dokumen tambahan oleh Klien atas permintaan SmartOption.',
            },
            {
                text: '5.7. Selain ketentuan yang diatur dalam pasal 5.2.-5.6. dari prosedur penyelesaian perselisihan, Klien dapat mengajukan klaim ke pengadilan, asalkan dia telah terlebih dahulu mengajukan klaim ke SmartOption sesuai dengan prosedur penyelesaian perselisihan ini. Prosedur klaim untuk penyelesaian perselisihan akan dianggap diikuti jika: ',
                child: [
                    'a) bentuk dan isi klaim memenuhi persyaratan klausul 5.2.1., 5.2.2., dan 5.2.3.; ',
                    'b) klaim dikirim ke alamat pendaftaran SmartOption melalui pos terdaftar atau bersertifikat;',
                    'c) Klien memiliki konfirmasi penerimaan klaim oleh SmartOption; ',
                    'd) batas waktu untuk menanggapi klaim telah berakhir. Waktu respons klaim - enam puluh (60) hari kalender setelah diterima oleh SmartOption.',
                ]
            },
            {
                text: '5.8. Jika terjadi perselisihan, SmartOption berhak untuk sepenuhnya atau sebagian memblokir transaksi di Akun Klien sampai perselisihan diselesaikan atau sampai Para Pihak mencapai kesepakatan sementara.',
            },
        ],
        termsArr6Title: '6. Peraturan pemerintah',
        termsArr6TextArr: [
            {
                text: '6.1. Klien dengan tegas:',
                child: [
                    'a) setuju bahwa pengadilan negara tempat pendaftaran SmartOption memiliki yurisdiksi eksklusif untuk melakukan proses hukum apa pun sehubungan dengan Perjanjian ini;',
                    'b) tunduk pada yurisdiksi pengadilan negara tempat pendaftaran SmartOption;',
                    'c) mengesampingkan setiap banding sehubungan dengan proses di salah satu pengadilan tersebut;',
                    'd) setuju untuk tidak membuat klaim tidak nyaman mengenai lokasi uji coba dan tidak menyatakan bahwa negara lokasi uji coba tidak memiliki yurisdiksi hukum atas Klien.',
                ]
            },
        ],
        termsArr7Title: '7. Force Majeure ',
        termsArr7TextArr: [
            {
                text: '7.1. Jika SmartOption memiliki alasan yang cukup, ia dapat mengklaim peristiwa force majeure. Peristiwa force majeure termasuk (tanpa batasan):',
                child: [
                    'a) setiap tindakan, peristiwa atau kejadian (termasuk, namun tidak terbatas pada, pemogokan, kerusuhan atau perselisihan sipil, aksi teroris, perang, bencana alam, kecelakaan, kebakaran, banjir, badai, pemadaman listrik, gangguan dalam pengoperasian komunikasi peralatan, perangkat lunak atau peralatan elektronik, pengoperasian yang salah dari segala jenis peralatan atau perangkat lunak, ketidakstabilan dalam aliran kuotasi, gangguan dalam pengoperasian atau ketidakstabilan penyedia likuiditas, dll.), yang, menurut pendapat wajar SmartOption, menyebabkan destabilisasi pasar atau pasar untuk satu atau lebih aset (instrumen);',
                    'b) penangguhan pekerjaan, likuidasi atau penutupan pasar apa pun atau tidak adanya peristiwa apa pun yang menjadi dasar kutipan SmartOption, atau pengenaan pembatasan atau persyaratan perdagangan khusus atau non-standar di pasar mana pun, atau sehubungan dengan acara semacam itu.',
                ]
            },
            {
                text: '7.2. Jika SmartOption telah menetapkan bahwa suatu peristiwa force majeure terjadi, SmartOption berhak (tanpa mengurangi hak-hak SmartOption lainnya) untuk mengambil langkah-langkah berikut tanpa pemberitahuan tertulis sebelumnya dan setiap saat: a) membatalkan setiap atau semua perdagangan yang secara langsung atau tidak langsung disebabkan oleh force majeure; (b) menangguhkan atau mengubah penerapan salah satu atau semua ketentuan Perjanjian selama peristiwa force majeure membuat SmartOption tidak mungkin untuk mematuhi ketentuan - ketentuan ini; c) mengambil atau, sebaliknya, tidak mengambil tindakan apa pun sehubungan dengan SmartOption, Klien, atau klien lain, jika SmartOption secara wajar menganggapnya tepat dalam situasi tersebut.',
            },
            {
                text: '7.3 SmartOption tidak bertanggung jawab atas pelanggaran (pembebasan yang tidak semestinya) atas kewajiban jika peristiwa force majeure mengganggu pelepasan tersebut.',
            },
        ],
        termsArr8Title: '8. Tanggung Jawab Para Pihak',
        termsArr8TextArr: [
            {
                text: '8.1. Tanggung jawab Para Pihak dalam Perjanjian ini ditentukan oleh ketentuan - ketentuan Perjanjian dan lampirannya.',
            },
            {
                text: '8.2. SmartOption bertanggung jawab hanya untuk kerugian nyata yang disebabkan oleh Klien sebagai akibat dari pelanggaran yang disengaja oleh SmartOption atas kewajibannya di bawah ini. SmartOption bertanggung jawab atas tindakan perwakilannya, departemen, dan agen pembayarannya seperti halnya tindakannya sendiri.',
            },
            {
                text: '8.3. Klien bertanggung jawab kepada SmartOption atas kerugian yang ditimbulkan oleh SmartOption yang menjadi tanggung jawab Klien, termasuk:',
                child: [
                    'a) untuk kerusakan yang disebabkan sebagai akibat dari kegagalan Klien untuk memberikan (atau keterlambatan penyediaan) dokumen apa pun yang harus diberikan kepada SmartOption berdasarkan Perjanjian ini dan lampirannya dan untuk kerusakan yang terjadi kepada SmartOption karena kesalahan penyajian informasi yang terkandung dalam dokumen yang disediakan oleh Klien;',
                    'b) untuk kerusakan yang terjadi pada SmartOption karena penyalahgunaan layanan yang diberikan oleh SmartOption kepada Klien, termasuk kerusakan (kerugian) yang disebabkan oleh SmartOption oleh penggunaan algoritma transaksi robotik dan otomatis dan/atau perangkat lunak khusus dan perangkat lainnya, perangkat, metode, dan teknik yang memfasilitasi atau berkontribusi terhadap pelanggaran prinsip integritas, kejujuran, dan kewajaran dalam pelaksanaan transaksi;',
                    'c) fatau kerusakan yang disebabkan oleh Klien sebagai akibat dari tindakan yang dikoordinasikan dengan klien lain dari SmartOption dan/atau afiliasi Klien yang bertujuan menyebabkan kerugian SmartOption dan untuk kerusakan lain yang disebabkan oleh SmartOption oleh Klien dari penggunaan metode dan teknik tidak adil dan tidak jujur ​lainnya untuk melakukan perdagangan (transaksi) dengan SmartOption, termasuk menggunakan bonus. Dalam situasi apa pun "Afiliasi klien" berarti orang-orang dalam hubungan kekerabatan dalam tingkat apa pun; perkawinan, kemitraan, atau hubungan lainnya; bertempat tinggal di alamat yang sama dengan orang tersebut; orang yang menggunakan perangkat yang sama; orang yang terlibat sebagai Klien SmartOption oleh mitra atau Klien yang sama dari SmartOption; dan orang-orang yang melakukan kegiatan bersama dengan atau tanpa pembentukan badan hukum. SmartOption berhak untuk memperluas daftar situasi dan atribut di mana Klien dan pihak ketiga dapat dikenali sebagai afiliasi;',
                    'd) jika ada bukti yang cukup untuk menunjukkan bahwa Klien secara tidak sah mencoba menggunakan perangkat lunak yang disediakan oleh SmartOption dan dana yang ditransfer ke akun SmartOption;',
                    'e) untuk kerusakan yang terjadi pada SmartOption sebagai akibat dari penggalian pendapatan dari penggunaan fitur teknis dari pembaruan aliran kutipan di terminal perdagangan dan ekstraksi pendapatan dari penggunaan kesalahan perangkat lunak dan kerentanan di terminal perdagangan;',
                    'f) untuk kerusakan yang disebabkan oleh SmartOption oleh Klien yang menggunakan informasi orang dalam, rahasia, atau informasi lain yang memberikan keuntungan apa pun kepada Klien dalam menyelesaikan perdagangan dengan SmartOption. SmartOption berhak untuk mendebit kerugian ini dari akun Klien dan/atau akun orang lain (jika diketahui bahwa akun ini adalah milik Klien (atau kaki tangan Klien) menggunakan peralatan dan alat teknis dan lainnya dari SmartOption ). SmartOption juga memiliki hak untuk memblokir transaksi lebih lanjut di terminal perdagangan dan dasbor untuk Klien sehubungan dengan siapa SmartOption memiliki alasan dan kecurigaan yang cukup untuk mengklasifikasikan tindakan mereka (termasuk tindakan bersama dengan Klien lain) yang bertujuan untuk menyebabkan kerusakan ke SmartOption dan mendebit dana dari akun Klien untuk SmartOption.',
                ]
            },
            {
                text: '8.4. Jika Klien melanggar Perjanjian ini, SmartOption dapat, atas pilihannya:',
                child: [
                    '8.4.1. Meninjau jumlah kewajiban keuangan SmartOptions kepada Klien dan buat perubahan pada data (saldo) Akun Klien.',
                    '8.4.2. Menangguhkan layanan ke Klien, blokir akses ke terminal perdagangan. Jika SmartOption memblokir akses Klien ke terminal perdagangan, Klien harus mengambil semua langkah yang diperlukan dan wajar untuk memperbaiki alasan mengapa akses ke terminal perdagangan diblokir. Jika Klien tidak mengambil tindakan atau tindakan apa pun untuk mengatasi alasan mengapa akses diblokir dalam waktu 30 (tiga puluh) hari, SmartOption dapat menarik semua dana dari akun perdagangan. SmartOption memiliki hak tetapi tidak berkewajiban untuk menyetor ulang semua dana yang ditarik ke akun perdagangan Klien jika Klien memenuhi semua persyaratan yang diperlukan untuk mencabut pemblokiran pada akun perdagangan.',
                ]
            },
            {
                text: '8.5.Jika Klien melanggar ketentuan apa pun dari Perjanjian dan bagian integralnya yang tercantum dalam klausul 1.2, termasuk penolakan untuk menjalani pemeriksaan yang diperlukan dan penolakan untuk memberikan informasi yang diperlukan, SmartOption berhak untuk mengakhiri Perjanjian; untuk membatalkan setiap transaksi Klien; untuk menutup satu, beberapa, atau semua perdagangan Klien kapan saja, atas kebijakannya sendiri; dan untuk berhenti memberikan layanan kepada Klien dan mengembalikan dana kepada Klien atau tidak atas kebijakannya sendiri. Pelanggaran apa pun terhadap persyaratan yang tercantum di bagian ini, membuat Pelanggan kehilangan hak untuk meminta pembayaran atau pengembalian dana dari SmartOption:',
                child: [
                    '8.5.1. Jika SmartOption mengakhiri Perjanjian dengan Klien karena melanggar Perjanjian, Klien tidak berhak membuka akun baru, termasuk dengan memasukkan data pihak ketiga selama pendaftaran. Jika SmartOption mendeteksi pelanggaran Klien sebagaimana diatur dalam klausul ini, konsekuensi yang ditetapkan dalam bagian 8.5. dari Perjanjian akan berlaku.'
                ]
            },
            {
                text: '8.6. SmartOption tidak bertanggung jawab kepada Klien atas segala kerusakan, kerugian, kehilangan keuntungan, kehilangan peluang (tanpa batasan, karena kemungkinan fluktuasi pasar), pengeluaran atau kerusakan yang ditanggung oleh Klien sebagai akibat dari pelaksanaan perdagangan sesuai dengan ketentuan dari Perjanjian.',
            },
            {
                text: '8.7. SmartOption tidak bertanggung jawab jika terjadi perbedaan antara informasi yang ditampilkan di terminal perdagangan Klien dan informasi di Server SmartOption ketika hasil keuangan dari perdagangan Klien ditentukan. Untuk menghilangkan perbedaan tersebut, SmartOption akan menyesuaikan data di terminal perdagangan sesuai dengan informasi yang tersedia di Server SmartOption.',
            },
            {
                text: '8.8. SmartOption tidak bertanggung jawab atas kerusakan Klien jika kerusakan tersebut diakibatkan oleh serangan peretas, kecelakaan (kegagalan) jaringan komputer, jaringan komunikasi, saluran listrik atau sistem telekomunikasi, dll., yang langsung digunakan untuk menentukan kondisi penting dari transaksi Klien atau memastikan prosedur operasi SmartOption lainnya yang terjadi bukan karena kesalahan SmartOption.',
            },
            {
                text: '8.9. SmartOption tidak bertanggung jawab atas kegagalan teknis dan/atau gangguan dalam operasi terminal perdagangan yang terjadi sebagai akibat dari serangan peretas, kecelakaan (kegagalan) jaringan komputer, jaringan komunikasi, saluran listrik atau sistem telekomunikasi, dll., atau untuk kerugian Klien akibat kegagalan dan/atau gangguan tersebut.',
            },
            {
                text: '8.10. SmartOption tidak bertanggung jawab atas hasil perdagangan yang diputuskan oleh Klien berdasarkan bahan analitik yang disediakan oleh SmartOption dan/atau pihak ketiga. Klien telah diberitahu bahwa transaksi yang dilakukan di bawah ini menanggung risiko tidak menerima pendapatan yang diharapkan dan risiko kehilangan sebagian atau seluruh uang yang disimpan olehnya di Akun Klien. Klien mengakui bahwa kecuali ada penipuan, pelanggaran kewajiban yang disengaja, atau kelalaian berat dari pihak SmartOption, SmartOption tidak akan bertanggung jawab atas kerugian, pengeluaran, biaya, dan kerusakan apa pun dari Klien yang diakibatkan oleh ketidakakuratan informasi yang diberikan kepada Klien, termasuk, namun tidak terbatas pada, informasi tentang transaksi perdagangan Klien. SmartOption berhak untuk membatalkan atau menutup perdagangan Klien berdasarkan ketentuan yang ditetapkan dalam Perjanjian ini; namun demikian, semua transaksi yang dilakukan oleh Klien sebagai akibat dari informasi yang tidak akurat atau kesalahan ini tetap berlaku dan harus dipenuhi baik oleh Klien maupun SmartOption.',
            },
            {
                text: '8.11.SmartOption tidak bertanggung jawab atas kerugian apa pun yang mungkin diderita Klien jika terjadi pencurian, kehilangan, atau pengungkapan kata sandinya ke terminal perdagangan kepada pihak ketiga. Klien bertanggung jawab penuh untuk melindungi kata sandi dan menjaganya dari akses pihak ketiga yang tidak sah.',
            },
            {
                text: '8.12. SmartOption tidak bertanggung jawab atas pelanggaran (pembebasan yang tidak semestinya) dari kewajiban di bawah ini jika disebabkan oleh peristiwa force majeure atau kondisi luar biasa lainnya yang dinyatakan dalam Perjanjian atau lampiran Perjanjian.',
            },
            {
                text: '8.13. SmartOption tidak bertanggung jawab atas kerugian tidak langsung, khusus, sewenang-wenang, atau hukuman yang diderita oleh Klien, termasuk, namun tidak terbatas pada, kehilangan keuntungan, kehilangan tabungan yang diharapkan, atau kehilangan pendapatan, bahkan jika Klien diberitahu oleh SmartOption tentang kemungkinan kerusakan tersebut. Kerusakan moral tidak dikompensasi.',
            },
            {
                text: '8.14. SmartOption berhak untuk mempertimbangkan pelanggaran Klien setiap saat, terlepas dari waktu pelanggaran dilakukan, dan, jika pelanggaran ditemukan, untuk mengambil tindakan sesuai dengan Perjanjian ini.',
            },
        ],
        termsArr9Title: '9. Jangka Waktu dan Pengakhiran Perjanjian',
        termsArr9TextArr: [
            {
                text: '9.1. Perjanjian ini mulai berlaku pada kesimpulannya (pendaftaran klien di situs web atau di terminal perdagangan SmartOption) dan berlaku hingga waktu yang tidak ditentukan.',
            },
            {
                text: '9.2. Salah satu Pihak dapat mengakhiri Perjanjian ini secara sepihak:',
                child: [
                    '9.2.1. Perjanjian diakhiri atas inisiatif SmartOption pada tanggal yang ditentukan dalam pemberitahuan yang dikirim oleh SmartOption kepada Klien;',
                    '9.2.2. Perjanjian diakhiri atas inisiatif Klien dalam lima (5) hari kerja setelah SmartOption menerima pemberitahuan tertulis kepada Klien yang berisi pernyataan penghentian Perjanjian, asalkan Klien tidak memiliki kewajiban yang belum terpenuhi di bawah ini. Klien harus mengirimkan Pemberitahuan Penghentian ke alamat SmartOption yang diberikan dalam klausul 1.1. Perjanjian ini, atau ke alamat email sssmartoption@gmail.com.',
                ]
            },
            {
                text: '9.3. Perjanjian ini dianggap diakhiri sehubungan dengan Para Pihak ketika kewajiban timbal balik antara Klien dan SmartOption sehubungan dengan transaksi yang dilakukan sebelumnya telah terpenuhi dan semua hutang dari masing-masing Pihak dilunasi.',
            },
        ],
        termsArr10Title: '10. Ketentuan akhir',
        termsArr10TextArr: [
            {
                text: '10.1. Amandemen dan penambahan pada Perjanjian ini dan lampirannya dibuat oleh SmartOption secara sepihak. Semua amandemen dan tambahan yang dibuat oleh SmartOption dan tidak terkait dengan keadaan yang ditentukan di sini mulai berlaku pada tanggal yang ditentukan oleh SmartOption.',
            },
            {
                text: '10.2. Amandemen dan tambahan yang dibuat oleh SmartOption pada Perjanjian ini dan lampirannya karena amandemen undang-undang dan peraturan yang mengatur subjek perjanjian ini dan aturan dan kontrak sistem perdagangan yang digunakan oleh SmartOption untuk memenuhi kewajibannya di bawah ini mulai berlaku bersamaan dengan amandemen dalam dokumen-dokumen tersebut di atas.',
            },
            {
                text: '10.3. Ketika amandemen dan suplemen yang dibuat oleh SmartOption mulai berlaku, mereka akan berlaku sama untuk semua Klien, termasuk mereka yang menandatangani Perjanjian sebelum tanggal efektifnya.',
            },
            {
                text: '10.4. Untuk memastikan bahwa Klien yang menandatangani Perjanjian mengetahui amandemen dan penambahan, Klien harus mengunjungi Situs Web SmartOption atau platform perdagangan sendiri atau melalui orang yang berwenang setidaknya sekali seminggu untuk menemukan informasi tentang setiap amandemen dan/atau tambahan.',
            },
            {
                text: '10.5. Klien (individu) yang memberikan data pribadinya ke SmartOption dalam bentuk apa pun dan dengan cara apa pun (saat melakukan tindakan apa pun di situs web SmartOption, melalui rekanan SmartOption, dll.) dengan demikian memberikan SmartOption dan mitranya persetujuannya untuk pemrosesan data pribadinya secara otomatis dan tidak otomatis untuk tujuan pemenuhan Perjanjian ini, pelaksanaan kampanye iklan, penyediaan materi iklan, informasi, dan pemasaran, informasi tentang kampanye. dan acara yang diadakan oleh SmartOption, dan untuk tujuan lain yang ditentukan oleh SmartOption, yaitu: mengumpulkan, mencatat, mensistematisasikan, mengakumulasi, menyimpan, menyesuaikan (memperbarui, mengubah), mengekstrak, mentransfer (menyebarluaskan, memberikan akses), mendepersonalisasi, memblokir, menghapus, menghancurkan, dan mentransfer data pribadi lintas batas. Persetujuan diberikan untuk jangka waktu 75 tahun (atau sampai berakhirnya periode penyimpanan untuk informasi atau dokumen relevan yang berisi informasi tersebut yang ditentukan sesuai dengan undang-undang saat ini di lokasi SmartOption). Persetujuan dapat ditarik sesuai dengan kebijakan Privasi. SmartOption menjamin kerahasiaan data pribadi yang diberikan oleh Klien, tidak termasuk keadaan yang ditetapkan oleh hukum yang berlaku dan force majeure.',
            },
            {
                text: '10.6. Klien dapat menggunakan informasi yang diberikan kepada Klien baik secara lisan atau tertulis yang diposting oleh SmartOption atau oleh pihak ketiga, akses yang diberikan kepadanya sebagai bagian dari layanan di bawah ini hanya untuk transaksi berdasarkan Perjanjian ini. Klien tidak boleh menyebarkan, mengubah, atau melengkapi informasi yang disebutkan di atas, atau menyimpannya dalam arsip terpisah. Dalam keadaan apa pun, cakupan kewenangan yang diberikan kepada Klien sehubungan dengan informasi yang diposting oleh pihak ketiga tidak dapat melebihi cakupan kewenangan yang diperoleh SmartOption dari pihak ketiga. SmartOption tidak menjamin bahwa informasi yang diposting oleh pihak ketiga dapat diandalkan, akurat, atau relevan dan diberikan secara berkelanjutan tanpa gangguan. SmartOption juga tidak bertanggung jawab atas hasil transaksi (kerugian, kehilangan keuntungan, kehilangan pendapatan, cedera pada niat baik, dll.) yang diputuskan oleh Klien untuk dilakukan berdasarkan informasi yang dikomunikasikan kepada Klien baik secara lisan maupun tertulis oleh SmartOption atau oleh pihak ketiga.',
            },
            {
                text: '10.7. SmartOption dapat sepenuhnya atau sebagian mengalihkan hak dan kewajiban berdasarkan Perjanjian ini dan lampirannya kepada pihak ketiga jika orang tersebut berjanji untuk memenuhi persyaratan perjanjian ini. Pengalihan hak dan kewajiban ini tidak memerlukan pemberitahuan sebelumnya kepada Klien oleh SmartOption dan harus dilakukan pada saat informasi yang memadai dipublikasikan di situs web SmartOption.',
            },
            {
                text: '10.8. Klien tidak berhak untuk mengalihkan haknya sendiri, memaksakan kewajibannya sendiri atau mentransfer hak atau kewajibannya di bawah ini dengan cara lain apa pun tanpa persetujuan tertulis sebelumnya dari SmartOption. Jika kondisi ini dilanggar, setiap penugasan, pengenaan atau pengalihan tersebut akan dianggap batal demi hukum.',
            },
            {
                text: '10.9. SmartOption, mitranya atau afiliasi lainnya mungkin memiliki keuntungan material, hubungan hukum atau pengaturan sehubungan dengan transaksi di platform perdagangan atau di Dasbor atau keuntungan material, hubungan hukum atau pengaturan yang bertentangan dengan kepentingan Klien . Misalnya, SmartOption dapat: a) bertindak sebagai rekanan sehubungan dengan aset apa pun; b) menyarankan mitra lain dari SmartOption sebagai rekanan untuk operasi perdagangan; c) memberikan rekomendasi dan memberikan layanan kepada mitranya atau klien lain dari SmartOption sehubungan dengan aset yang mereka minati, meskipun faktanya bertentangan dengan kepentingan Klien.',
            },
            {
                text: '10.10. Klien setuju dan mengizinkan SmartOption untuk bertindak sehubungan dengan Klien dan untuk Klien sebagaimana yang dianggap tepat oleh SmartOption, terlepas dari potensi konflik kepentingan atau adanya beberapa kepentingan material sehubungan dengan transaksi apa pun di terminal perdagangan atau di Dashboard tanpa pemberitahuan sebelumnya dari Klien. Adanya konflik kepentingan atau keuntungan material sehubungan dengan transaksi apa pun di terminal perdagangan atau di Dasbor tidak boleh memengaruhi penyediaan layanan kepada Klien oleh karyawan SmartOption. Dari waktu ke waktu SmartOption dapat bertindak atas nama Klien dengan pihak-pihak yang dengannya SmartOption atau salah satu pihak terkaitnya memiliki perjanjian untuk menerima barang atau jasa. SmartOption menjamin bahwa perjanjian ini dibuat sejauh mungkin untuk kepentingan Klien, misalnya, perjanjian ini memungkinkan untuk mengakses informasi dan layanan lain yang tidak dapat diakses.',
            },
            {
                text: '10.11. Jika pengadilan dengan yurisdiksi yang tepat menyatakan ketentuan apa pun dalam Perjanjian (atau bagian mana pun dari ketentuan apa pun) batal, ketentuan tersebut akan diperlakukan sebagai bagian terpisah dari Perjanjian dan ini tidak akan memengaruhi kekuatan hukum dari sisa Perjanjian.',
            },
            {
                text: '10.12. SmartOption dapat sewaktu - waktu menangguhkan layanan kepada Klien  (pemberitahuan sebelumnya kepada Klien tidak diperlukan).',
            },
            {
                text: '10.13. Dalam situasi yang tidak dijelaskan dalam Perjanjian, SmartOption akan bertindak sesuai dengan praktik terbaik pasar berdasarkan prinsip kejujuran dan keadilan.',
            },
            {
                text: '10.14. SmartOption dapat menyiapkan dan menggunakan teks Perjanjian dan lampirannya dalam bahasa selain bahasa Inggris. Jika ada kontradiksi antara teks Perjanjian ini dan lampirannya dalam bahasa Inggris dan teks terkait dalam bahasa lain, teks dalam bahasa Inggris yang akan berlaku. Teks Perjanjian yang dipublikasikan di situs web SmartOption berlaku di atas teks Perjanjian yang diterbitkan di tempat lain.',
            },
            {
                text: '10.15. Klien diberikan hak terbatas dan non-eksklusif untuk menggunakan Terminal Perdagangan hanya untuk tujuan yang ditetapkan dalam Perjanjian ini. Jika Perjanjian diakhiri dengan alasan apa pun, hak Klien untuk menggunakan Terminal Perdagangan yang ditetapkan dalam klausul ini akan diakhiri pada saat Perjanjian diakhiri.',
            },
            {
                text: '10.16. Klien setuju bahwa SmartOption tidak dapat menjamin pengoperasian Terminal Perdagangan yang terus menerus tanpa gangguan dan baik secara teknis, dan oleh karena itu Klien menerima perangkat lunak ini apa adanya. SmartOption tidak bertanggung jawab kepada Klien atas kegagalan teknis dalam operasi Terminal Perdagangan.',
            },
            {
                text: '10.17. Semua istilah yang digunakan dalam Perjanjian ini dan bagian-bagiannya yang tidak terpisahkan, baik dalam huruf besar maupun huruf kecil, akan memiliki kepentingan yang sama, kecuali jika tidak mengikuti esensi kewajiban.',
            },
        ],
        termsArr11Title: '11. Daftar Negara di mana SmartOption Tidak Melakukan Bisnis',
        termsArr11TextArr: [
            {
                text: '11.1. TSmartOption tidak menjalankan bisnis di negara-negara berikut dan tidak menyediakan layanan kepada orang-orang yang terkait dengan negara bagian tersebut: Gibraltar, Pulau Man, Guernsey, Jersey, Australia, China, Kanada, Amerika Serikat, Jepang, Austria, Belgia, Bulgaria, Kroasia, Siprus, Lichtenstein, Republik Ceko, Denmark, Estonia, Finlandia, Prancis, Jerman, Yunani, Hongaria, Irlandia, Islandia, Italia, Israel, Latvia, Lithuania, Luksemburg, Malta, Belanda, Norwegia, Selandia Baru, Polandia, Portugal, Rumania, Slovakia, Slovenia, Spanyol, Swiss, Swedia, Republik Islam Iran dan Inggris;',
            },
            {
                text: '11.2. Orang-orang yang terkait dengan suatu negara didefinisikan sebagai orang-orang:',
                child: [
                    '11.2.1. memiliki kewarganegaraan, visa tinggal, atau dokumen serupa lainnya dari negara di mana SmartOption tidak menjalankan bisnis;',
                    '11.2.2. tinggal di negara di mana SmartOption tidak melakukan bisnis, menjadi penduduk negara di mana SmartOption tidak melakukan bisnis atau memiliki alamat surat atau alamat tempat tinggal di negara di mana SmartOption tidak melakukan bisnis;',
                    '11.2.3. yang tempat lahirnya adalah negara di mana SmartOption tidak melakukan bisnis;',
                    '11.2.4. memiliki nomor telepon dengan kode negara negara di mana SmartOption tidak menjalankan bisnis; memiliki alamat IP yang terkait dengan negara tempat SmartOption tidak menjalankan bisnis;',
                    '11.2.5. memiliki koneksi lain dengan negara di mana SmartOption tidak menjalankan bisnis sebagaimana ditentukan oleh SmartOption atas kebijakannya sendiri.',
                ]
            },
            {
                text: '11.3. Jika ditemukan bahwa SmartOption menyediakan layanan kepada orang-orang yang terkait dengan negara-negara di mana SmartOption tidak menjalankan bisnis, SmartOption dapat menerapkan konsekuensi yang tercantum dalam klausul 8.5 Perjanjian.',
            },
        ],
    },
    EN: {
        title: 'Terms & Conditions',
        termsArr0Title: 'Contents: ',
        termsArr0TextArr: [
            {text: '1. General Provisions and Subject of the Agreement'},
            {text: '2. Terms and Definitions'},
            {text: '3. Communications and Provision of Information'},
            {text: '4. Terms of use of the SmartOption services'},
            {text: '5. Claims and Dispute Resolution'},
            {text: '6. Governing Law'},
            {text: '7. Force Majeure'},
            {text: '8. Liability of the Parties'},
            {text: '9. Term and Termination of the Agreement'},
            {text: '10. Final Provisions'},
            {text: '11. List of Countries in which the Company Does Not Do Business'},
        ],
        termsArr1Title: '1. General Provisions and Subject of the Agreement',
        termsArr1TextArr: [
            {
                text: "1.1. This Service Agreement has been made by SmartOption App (hereinafter - SmartOption), and the natural person who filled in the registration form on the SmartOption's website or in the trading platform and accepted the terms of this Service Agreement and of the appendices hereto at the moment of registration (hereinafter Client)."
            },
            {
                text: "1.2. The Agreement is an invitation to make offers posted on the SmartOption's website that must be treated as an offer to enter into this Agreement on the terms set forth in it. The posted offer is not public. The SmartOption at its sole discretion may refuse to enter into an Agreement with anyone without explaining the reasons for refusal or, if registration took place, to terminate contractual relations and to block access to the trading terminal. The Client's registration on the SmartOption's Website or on the trading terminal is considered as full and unconditional acceptance of the terms of the Agreement. As soon as the SmartOption receives a payment to replenish the Client's trading account, each Client transaction using the trading terminal or Dashboard becomes the subject of this Agreement."
            },
            {
                text: "1.3. The Client must carefully review the terms of the Agreement. By accepting the terms of this Agreement, the Client agrees to the terms of all annexes to it listed above, including the terms on subdomains of the SmartOption's website that are accessible by the Client and confirms that he/she is an adult who is a legally capable person and is not a resident of a country where trading in options may be deemed illegal."
            },
            {
                text: 'The Client also represents and warrants to the SmartOption that:',
                child: [
                    "1.3.1. All information provided during Client registration and during the performance of Agreement, is true, accurate, reliable and complete in all respects, and the Client completed the registration form him/herself; ",
                    "1.3.2. The Client possesses legal personality to enter into this Agreement, to make inquiries and give orders, and to exercise his rights and to fulfill obligations in accordance with the terms of the Agreement; ",
                    "1.3.3. The Client will carry out trading and non-trading transactions personally, on his own behalf and at his own expense and will not use funds borrowed from other SmartOption Clients or from third parties to carry out transactions. The Client will be guided by the principles of integrity, honesty, and rationality; the Client will not take actions coordinated with other SmartOption Clients aimed at damaging the SmartOption; the Client will not use technical features of the quote stream update on the trading terminal and will not use software errors, defects, and vulnerabilities he discovers in the trading terminal to extract income and will not distribute the information about vulnerabilities to the third parties. The Client will not use unfair and dishonest methods or ways of making trades (transactions) with the SmartOption; the Client will not use insider or confidential information or any other information, as a result of the use of which the Client might benefit when trading with the SmartOption and/or that might damage the SmartOption; ",
                    "1.3.4. The Client will adhere to legal norms, in particular, but not limited to, international norms aimed at controlling illegal trade, financial fraud, and money laundering; ",
                    "1.3.5. The Client will not use the trading terminal or website to collude in illegal financial activities or any other illegal transaction; ",
                    "1.3.6. The money listed by the Client on the SmartOption's account has legal origin. The Client legally owns the money and has the right to use and manage it. The Client's account will not be replenished from third party payment instruments. ",
                    "The Client will not replenish third party Client accounts or withdraw money from the Client's account to third party payment instruments. ",
                    "1.3.7. No actions of the Client pursuant to the Agreement will violate any law, regulation, right, bylaws, or rules and regulations applicable to the Client or in the jurisdiction where he/she resides or the provisions of any other agreement binding on the Client or involving any assets of the Client ",
                    "1.3.8. To carry out transactions, the Client will use account data from Trading Terminal belonging to him and will not transfer account data to third parties and will not use account data of the SmartOption's clients for trading and (or) non-trading operations; ",
                    "1.3.9. The Client is not a federal or municipal employee, an employee of a federal or municipal institution, an employee of a federal or municipal organization, an organization in whose capital the state has a prevailing interest; the Client is not a politically significant person, a family member or a relative of a politically significant person; the Client is not a person closely connected with a politically significant person; the Client is not a person connected with the United States or with another country in which the SmartOption does not operate. The terms used in this paragraph are interpreted and applied by the SmartOption at its own discretion in accordance with the norms of international legislation and/or legislation of a particular state, generally accepted terms and definitions, customary business practices. ",
                ]
            },
            {
                text: '1.4. The subject of the Agreement is the definition of the general conditions under which the Parties carry out transactions (trades), the content and procedure for which are set forth in this Agreement. The SmartOption unilaterally sets and may alter at its sole discretion essential conditions for a transaction (trade), may limit the number of trades executed simultaneously, and limit the number of trades that a Client may make within the time period set by the SmartOption, may set other limits for trades at its sole discretion and unilaterally.'
            },
            {
                text: '1.5. The SmartOption may engage third parties to perform this Agreement. The SmartOption is not responsible for services provided by such third parties.'
            },
        ],
        termsArr2Title: '2. Terms and Definitions',
        termsArr2TextArr: [
            {
                text: ".1. Asset - the underlying financial instrument of an option. An asset may be company stock, a stock index, currency pair (the rate of an exchange of one currency into another currency), goods quoted on a commodity exchange, commodity options, etc."
            },
            {
                text: "2.2. Client's Account Balance - the amount of money in the Client's account."
            },
            {
                text: "2.3. Option - a derivative financial instrument (trade), consisting of two transactions: Making a trade and Closing a trade, as a result of the fulfillment of which the Client either receives Income or loses the Trade amount. The SmartOption does not, however, provide a Client with the ability to assume obligations greater than the Trade amount (the SmartOption does not provide leverage to the Client). The SmartOption and the Client make a trade with an Higher/Lower option - an option, the targeted level of which equals the asset price or another price determined by the Client at the SmartOption's suggestion when the trade is made, only payment on which is made if an upward price change was selected and the asset price was greater than the targeted level when the option expired or if a downward price change was selected and the asset price was less than the targeted level when the option expired. If the asset price is the same as the targeted level when the option expires, the trade Amount is returned to the Client"
            },
            {
                text: "2.4. Bonus - virtual cash credited by the SmartOption to the Client's Account when the Client meets the conditions established by the SmartOption. When the Client executes trading transactions, funds deposited by the Client must be used first and, only after those funds are completely exhausted, does the Client have the right to use the Bonus for further trading transactions. In general, the Client cannot withdraw the Bonus amount to the Client's External Account. If the Client withdraws funds that were previously credited to his External account, the SmartOption has the right to debit the entire amount of previously accrued Bonuses from the Client's Account. The SmartOption has the right to set a minimum trading amount and other terms under which the Bonus is not debited from the Client's Account and may be withdrawn to the Client's External account. The SmartOption may publish credit and debit rules, rules for using and converting Bonuses to real funds, and rules for other transactions with Bonuses and Client funds that the Client accrued along with Bonuses on its website."
            },
            {
                text: "2.5. SmartOption Website - an internet site at the address (domain name) SmartOption.com, and other internet sites which the SmartOption references in the Agreement or on the SmartOption Website."
            },
            {
                text: "2.6. Client External Account - the Client's current account at a lending institution, an account (wallet) on an electronic payment system."
            },
            {
                text: "2.7. Option Expiration Time - the time when the condition for payout of income on the option is checked."
            },
            {
                text: "2.8. Trading Hours - the time when trading in an asset may be done."
            },
            {
                text: "2.9. Funds Withdrawal - transfer of money from the Clients Account to the Clients External Account."
            },
            {
                text: "2.10. Income - fixed remuneration credited to the Client's Account if the conditions of the option are met. The Income amount is determined at the time that the trade is executed. Income is credited to the Clients account from the SmartOption's guaranteed (compensatory, special) fund (reserve fund), which is formed in the manner determined by the SmartOption at its own expense."
            },
            {
                text: "2.11. Execution of a trade - a trading transaction in which the essential conditions of an option trade are agreed by the Client and the SmartOption. After the essential conditions are agreed, the option trade is open."
            },
            {
                text: "2.12. Closing a trade - a trading transaction in which an open trade on an option is liquidated at the option expiration time agreed by the Parties when the trade is made. Early closing of a trade - closing a trade at the Client's initiative before it is closed at the option expiry time. Early closing of a trade is possible only if the SmartOption has the technical capability to do so. The SmartOption may unilaterally at its own discretion provide or refuse to grant the Client the right to make an early closing of a trade."
            },
            {
                text: "2.13. Exceptional Situations - abnormal market conditions and other exceptional conditions described in the Regulation on Trading Transactions."
            },
            {
                text: "2.14. Payout ratio - the percentage that determines the Income amount set by the SmartOption depending on the option asset and other conditions of the option."
            },
            {
                text: "2.15. Asset Quote - the price for one unit of the Asset."
            },
            {
                text: "2.16. Log entry - an entry in the database made by the SmartOption server that, with accuracy to the millisecond or, in the case of absence of technical capability, with accuracy to the second, documents all Client inquiries and orders and the result of their processing. Each Client contact with the trading terminal and dashboard is documented in a log entry. These server data are the primary source of information and are used by the parties as evidence if disputes related to the performance of the Agreement arise. Information from a SmartOption server Log entry has unconditional priority over all other arguments during settlement of a dispute, including over information from the Client's trading terminal Log file. The SmartOption reserves the right not to keep Log entries."
            },
            {
                text: "2.17. Price Direction - an essential condition of an option trade that determines the trade payout. A price may change direction Higher or Lower."
            },
            {
                text: "2.18. Non-trading Transaction - any Client transaction to deposit funds on the Client's Account or withdraw funds from the Client's Account."
            },
            {
                text: "2.19. Transactions - The Client's trading and non-trading transactions."
            },
            {
                text: "2.20. Open Trade - an option trade after the execution of the trade and before trade closure, for which it is not determined yet whether a payout will be made."
            },
            {
                text: "2.21. Price Feed - a succession of prices displayed on the Trading Terminal."
            },
            {
                text: "2.22. Payment service provider - a company that provides funds transfer services."
            },
            {
                text: "2.23. Recurring payment - a periodically repeating transaction to replenish the Balance of the Client's account without requiring the latter to reenter his or her credit card information."
            },
            {
                text: "2.24. SmartOption Server - the SmartOption's software used to process and store information about client requests for trading and non-trading transactions, to provide to the Client realtime information about quotes, to account for trading and non-trading transactions, to monitor the fulfillment of trade conditions and limit trading transactions, and to determine the financial result of trades."
            },
            {
                text: "2.25. Withdrawal Method - one of the Funds Withdrawal methods offered to the Client that is posted on the trading terminal and in the Dashboard."
            },
            {
                text: "2.26. The Trade Amount is the amount paid by the Client to the SmartOption when the trade is made. The Client's profit on an option trade if the trade conditions are fulfilled is determined by the procedure set forth in Section 2.10. of this Agreement."
            },
            {
                text: "2.27. SmartOption's Account - the SmartOption's settlement account at a financial institution, an account (wallet) in an electronic payment system, and other accounts, including accounts of Payment Agents."
            },
            {
                text: "2.28. Essential Conditions of a trading transaction (essential conditions of a trade) - the conditions that govern the payout of income from a trade to the Client by the SmartOption."
            },
            {
                text: "2.29. Client account (Trading account) - a special account in the SmartOption's accounting system in which funds transferred by the Client to make trades are posted; from which the trade amount is debited when the trade is executed; and to which income is credited when a trade is closed and the essential conditions of a trade are met. The Client has the right to have only one Client account. In case of violation of this rule, the SmartOption has the right to refuse the Client in further service (rendering of services), to terminate this Agreement, and to block further possibility to carry out transactions without explanation and without payment of funds from the Client's account. It is not a violation of this clause of the Agreement when the SmartOption, if possible and unilaterally, grants the Client the right to use several currencies within the Client's account, as well as the right to use the Client's account in relations between the SmartOption and the Client, which are governed simultaneously by this Agreement and other agreements, concluded between the SmartOption and the Client, under which the SmartOption has the discretion to grant the Client the right to use the Client's account to engage in trading transactions which are not prescribed in this Agreement."
            },
            {
                text: "2.30. Trading transactions - procedures to make and close trades with options between the SmartOption and a Client. Trading transactions are carried out at the place of the SmartOption's registration. No physical delivery of assets takes place during trading transactions. Trade amounts for trading transactions are debited from the Client's account balance after a trade is made. The Income on trading transactions are credited to the Client's account immediately after a trade closes."
            },
            {
                text: "2.31. The Trading Terminal is software through which the Client can obtain real-time information about quotes to carry out trading and non-trading transactions and receive messages from the SmartOption. Entry to the Trading Terminal is protected by a password that the Client sets at registration on the SmartOption's site. All orders and inquiries carried out through the Trading Terminal are deemed to be carried out personally by the Client. Clients from countries whose laws prohibit trading in options or other over-the-counter derivatives and the employees, affiliates, agents and other representatives of the SmartOption and their relatives are prohibited from using the Trading Terminal. The part of the Trading Terminal with which a Client can carry out nontrading transactions in this Agreement may be called the Dashboard."
            },
            {
                text: "2.32. Red Territory - the state of an open option trade if, at the current asset price, income cannot be paid out on that trade."
            },
            {
                text: "2.33. 1-Click service - makes it possible to replenish a Client's account balance from the Client's bank (payment) cards without entering the bank (payment) card data of the bank card holder again."
            },
            {
                text: "2.34. Targeted level - the level of asset prices with respect to which the trade result is calculated."
            },
            {
                text: "2.35. Cookie file - a small data set including an anonymous unique identifier which is sent to the web browser of the Client's computer or mobile telephone (hereinafter 'device') from the SmartOption server (website) and is stored on the Client's device. The Client may configure a web browser to block cookie access to the device. When a Client visits the SmartOption Website, the viewed pages and cookie files are downloaded onto the Client's device. Cookie files stored on the device may be used for anonymous identification if the Client revisits the SmartOption's website and to determine the website pages that are the most popular among our Clients. Cookie files stored on the hard drive of the Client's device enable the SmartOption to create the most user-friendly and efficient website for SmartOption Clients, providing us an opportunity to identify our Clients' preferences."
            },
            {
                text: "2.36. Trading signals are information about the state of the market gathered by the SmartOption based on analytical conclusions that the SmartOption is entitled to provide at its discretion to some or all Clients with respect to certain market indicators. Trading signals are not an offer and are not the SmartOption's explicit recommendation to Clients to engage in trading transactions or to make trades. The SmartOption is not liable for the accuracy of Trading Signals or for Client trading transactions and trades based on Trading Signals. The Client may, at his discretion, consider or ignore Trading signals when engaging in trading transactions and trades."
            },
            {
                text: "2.37. Quote - the current Asset price displayed on the Trading Terminal. Terms used in this Agreement and definitions absent in this Clause shall be interpreted in accordance with usual and customary business practices applicable in regard to conclusion of agreements with financial derivatives."
            },
        ],
        termsArr3Title: '3. Communications and Provision of Information',
        termsArr3TextArr: [
            {
                text: '3.1. To communicate with the Client, the SmartOption may use:',
                child: [
                    '· e-mail;',
                    '· fax;',
                    '· telephone;',
                    '· SMS;',
                    '· letters sent by post;',
                    '· different types of messages sent to the Client on the trading terminal, in the Dashboard, browser window, etc. (push notifications, reminders, service messages, etc.);',
                    '· announcements on the SmartOption website.',
                ]
            },
            {
                text: "3.2. To promptly communicate with the Client to resolve problems related to Client transactions, the SmartOption will use the Client's contact information entered when the Client registered or modified the account pursuant to clause 4.5 of the Agreement. The Client agrees to accept message from the SmartOption at any time. ",
            },
            {
                text: "3.3. Any correspondence (documents, notices, confirmations, announcements, reports etc.) are deemed received by the Client:",
                child: [
                    "1) one (1) hour after they are sent to the email address (e-mail);",
                    "2) immediately after they are faxed;",
                    "3) immediately after the telephone call ends;",
                    "4) immediately after an SMS message is sent;",
                    "5) seven (7) calendar days after they are mailed",
                    "6) immediately after an announcement is posted on the SmartOption website.",
                ]
            },
            {
                text: "3.4. The Client may also contact the SmartOption by e-mail at sssmartoption@gmail.com, and other email addresses and the telephone numbers listed in this Agreement and on the SmartOption website."
            },
            {
                text: "3.5. The Client understands and agrees that, if the Client's behavior during conversations with SmartOption representative is inappropriate, the SmartOption reserves the right to unilaterally terminate this Agreement."
            },
            {
                text: "3.6. The SmartOption may use contact information provided by the Client to send informational, marketing, and advertising materials, and service messages and to resolve other tasks. The SmartOption will determine the frequency with which it sends messages to the Client at its sole discretion. If the Client wishes not to receive from informational (and other) messages from the SmartOption, he/she must unsubscribe by clicking on the Unsubscribe link (if the message format provides this ability) or by contacting the client support department."
            },
        ],
        termsArr4Title: '4. Terms of use of the SmartOption services.',
        termsArr4TextArr: [
            {
                text: "4.1. During registration the Client undertakes to provide correct and reliable identification information in accordance with requirements of the Client registration form."
            },
            {
                text: "4.2. After successful registration, the Client will be granted access to the trading terminal, the ability to transfer funds to the Clients Account (placing a deposit in the Clients Account to be able to make option trades) and to perform other operations."
            },
            {
                text: "4.3. The Client must promptly inform the SmartOption about changes in identification and contact information (within 7 (seven) days after that change) by making the appropriate changes on the Trading Terminal or in any other way offered by the SmartOption. To identify the Client and to inspect the origin of the Client's funds at any time after registration, the SmartOption has the right to ask and the Client obliges to provide within 7 (seven) days after the request is received, any identification documents (these documents include identity documents; documents confirming the residential address, documents indicating the Client's financial standing, and other documents at the SmartOption's discretion). The SmartOption reserves the right to suspend non-trading and/or trading transactions on the Client's account if it finds that the Client's identification information is incorrect or inaccurate; if the Client has not provided the requested documents, the SmartOption has the right to block the Client's access to the trading terminal until the client's identification procedure is completed. The SmartOption also has the right to demand the Client to undergo the identification procedure by a personal visit to a SmartOption-authorized agent and provide documents, the list of which is determined by the SmartOption at its discretion."
            },
            {
                text: "4.4. Login to the trading terminal is password protected.",
                child: [
                    "4.4.1. The Client confirms and agrees that access to the trading terminal will be protected with a password set by the Client himself/herself during registration. The Client may not transfer the trading terminal password to third parties. ",
                    "4.4.2. The Client assumes full responsibility for password protection and prevention of unauthorized third-party access to it. ",
                    "4.4.3. All orders through the trading terminal with the Client's password will be deemed to have been made by the Client unless the SmartOption establishes otherwise. ",
                    "4.4.4. Any person who obtains access to the trading terminal by entering the Client's password will be identified as the Client unless otherwise specified by the SmartOption. ",
                    "4.4.5. The SmartOption assumes no liability for any losses the Client may suffer in case of theft, loss or disclosure of the password to third parties or in case of unauthorized use of registration data by third parties. ",
                ]
            },
            {
                text: "4.5. The Client may change the password to the trading terminal himself/herself or use the procedure for password recovery set by the SmartOption."
            },
        ],
        termsArr5Title: "5. Claims and Dispute Resolution",
        termsArr5TextArr: [
            {
                text: "5.1. The Parties have agreed that they will make every effort to settle all disputes between the SmartOption and the Client related to transactions, payouts and other actions hereunder by means of negotiations."
            },
            {
                text: "5.2. If a dispute arises, the Client may submit a claim/complaint to the SmartOption and send a statement or notice to the SmartOption. All claims/complaints/statements/ notices related to transactions performed by the Client must be submitted in accordance with the following requirements:",
                child: [
                    "5.2.1. claims/complaints/statements/notices must be submitted in writing;",
                    "5.2.2. claims/complaints/statements/notices must contain the following information: Surname, first name, middle name (if any), Clients e-mail, Clients Account Number, date and time of the dispute, brief description of the dispute, Clients demands; claimed amount and justified calculation thereof (if the claim can be valued in money); circumstances that are the grounds for the claim and evidence supporting those circumstances, including reference to the clause of this Agreement (annexes hereto) that were breached in the Clients opinion; a list of documents and other evidence attached to the claim (complaint) authenticated by the Client; other information necessary to settle the dispute;",
                    "5.2.3. claims/complaints/statements/notices must be sent by the Client within five (5) business days after the event that was the basis for submission of the relevant claim (complaint). The Client agrees that a delay in submission of the claim (complaint) is grounds for refusal to consider it;",
                    "5.2.4. claims/complaints/statements/notices may be sent by e-mail to sssmartoption@gmail.com, by registered or certified mail. Claims/ complaints/statements/notices made and sent in another form will not be considered.",
                ]
            },
            {
                text: "5.3. Claims/complaints/statements/notices must not contain:",
                child: [
                    "a) an emotional evaluation of the dispute;",
                    "b) offensive statements addressed to the SmartOption;",
                    "c) profanity.",
                ]
            },
            {
                text: "5.4. To respond to a claim/complaint/statement/notice, the SmartOption may request additional documents and information from the Client. A claim/complaint/statement/ notice will be reviewed on the basis of data provided by the Client and log entries from the SmartOption server. Log entries from the SmartOption Server always prevail over other evidence and proof. The SmartOption assumes no liability for incomplete trades and will not compensate any financial damages or moral harm suffered by the Client with respect to what the Client considers to be lost profit. When considering disputes, the Clients references to information from other companies and websites are not considered.",
            },
            {
                text: "5.5. The SmartOption may reject a claim/complaint/statement/notice if the terms of this section are violated.",
            },
            {
                text: "5.6. The SmartOption must consider a claim/complaint/application/appeal within no more than 10 working days after the submission date. This term does not include time for provision of additional documents by the Client upon the SmartOption's request.",
            },
            {
                text: "5.7. In addition to provision set forth in clauses 5.2.-5.6. of the dispute resolution procedure, the Client may file a claim to a court, provided that he/she has first submitted the claim to the SmartOption in accordance with this dispute resolution procedure. The claim procedure for dispute resolution will be deemed followed if: a) the form and content of the claim meet the requirements of clauses 5.2.1., 5.2.2., and 5.2.3.; b) the claim is sent to the SmartOption's registration address by registered or certified mail; c) the Client has a confirmation of claim receipt by the SmartOption; d) the deadline for responding to the claim has expired. Claim response time - sixty (60) calendar days after it is received by the SmartOption.",
            },
            {
                text: "5.9. In case of any disputes, the SmartOption reserves the right to fully or partially block transactions in the Client's Account until the dispute is settled or until the Parties come to an interim agreement.",
            },
        ],
        termsArr6Title: "6. Governing Law",
        termsArr6TextArr: [
            {
                text: "6.1. The Client expressly:",
                child: [
                    "a) agrees that the courts of the country of the SmartOption's registration have exclusive jurisdiction to conduct any legal proceedings with respect to this Agreement;",
                    "b) submits to the jurisdiction of the courts of the country of the SmartOption's registration;",
                    "c) waives any appeals with respect to proceedings in any of such courts;",
                    "d) agrees to make no claims regarding trials location as non-convenient and not to declare that the trials location state has no legal jurisdiction over the Client.",
                ]
            },
        ],
        termsArr7Title: "7. Force Majeure",
        termsArr7TextArr: [
            {
                text: "7.1. If the SmartOption has sufficient grounds, it may claim force majeure events. Force majeure events include (without limitation):",
                child: [
                    "(a) any action, event or occurrence (including, but not limited to, any strike, riots or civil strife, terrorist acts, wars, natural disasters, accidents, fires, floods, storms, power outages, interruptions in the operation of communication equipment, software or electronic equipment, incorrect operation of any kind of equipment or software, the instability in the quote stream, interruptions in the operation of or the instability of liquidity providers, etc.), which, in the SmartOption's reasonable opinion, led to the destabilization of the market or the markets for one or more assets (instruments);",
                    "b) the suspension of work, the liquidation or closure of any market or the absence of any event on which the SmartOption bases quotes, or the imposition of restrictions or special or non-standard terms of trade in any market, or in respect of any such event.",
                ]
            },
            {
                text: "7.2. If the SmartOption has established that a force majeure event occurred, the SmartOption has the right (without prejudice to the SmartOption's other rights) to take any of the following steps without prior written notification and at any time: a) cancel any or all trades the result of which is directly or indirectly caused by force majeure; (b) suspend or amend the application of one or all provisions of the Agreement as long as the force majeure event makes it impossible for the SmartOption to comply with these provisions; c) take or, on the contrary, not take any action in respect of the SmartOption, the Client or other clients, if the SmartOption reasonably deems it appropriate under the circumstances.",
            },
            {
                text: "7.3 The SmartOption assumes no liability for breach (improper discharge) of obligations if force majeure events interfered with that discharge.",
            },
        ],
        termsArr8Title: "8. Liability of the Parties",
        termsArr8TextArr: [
            {
                text: "8.1. The liability of the Parties to this Agreement is determined by the terms of the Agreement and its annexes.",
            },
            {
                text: "8.2. The SmartOption assumes liability only for real damages caused to the Client as a result of the SmartOption's deliberate breach of its obligations hereunder. The SmartOption is as responsible for the actions of its representatives, departments, and payment agents as it is for its own actions.",
            },
            {
                text: "8.3. The Client assumes liability to the SmartOption for losses incurred by the SmartOption for which the Client is responsible, including:",
                child: [
                    "a) for damage caused as a result of the Client's failure to provide (or late provision) of any documents that must be provided to the SmartOption under this Agreement and its annexes and for damage caused to SmartOption because of any misstatement of information contained in documents provided by the Client;",
                    "b) for damage caused to the SmartOption because of abuse of services provided by the SmartOption to the Client, including damage (losses) caused to the SmartOption by the use of robotic and automated transaction algorithms and/or special software tools and other tools, devices, methods and techniques that facilitate or contribute to the violation of the principle of integrity, honesty and fairness in the execution of transactions;",
                    "c) for damage caused by the Client as a result of the actions coordinated with other clients of the SmartOption and/or affiliates of the Client aimed at causing the SmartOption losses and for other damage caused to the SmartOption by the Client from the use of other unfair and dishonest methods and techniques for making trades (transactions) with the SmartOption, including using bonuses. In any situation 'Client affiliates' means persons in a kinship relationship of any degree; marital, partnership, or other relationships; residing at the same address as the person; persons using the same devices; persons engaged as a SmartOption Client by the same partner or Client of the SmartOption; and persons engaged in any joint activity with or without the formation of a legal entity. The SmartOption reserves the right to expand the list of situations and attributes in which the Client and third parties may be recognized as affiliated;",
                    "d) if there is sufficient evidence to suggest that the Client unlawfully attempted to use the software provided by the SmartOption and the funds transferred to the SmartOption's account;",
                    "e) for damage caused to the SmartOption as a result of extracting income from the use of the technical features of the quote stream update on the trading terminal and of extracting income from the use of software errors and vulnerabilities in the trading terminal;",
                    "f) for damage caused to the SmartOption by the Client's use of insider, confidential, or other information that provided the Client with any kind of advantages in concluding trades with the SmartOption. The SmartOption has the right to debit these losses from the Client's account and/or the accounts of other persons (if it is established that these accounts belong to the Client (or the Client's accomplices) using the SmartOption's technical and other equipment and tools). The SmartOption also has the right to block the further transactions on the trading terminal and the dashboard for Clients in respect to whom the SmartOption has sufficient grounds and suspicions to classify their actions (including joint actions with other Client's) as aimed at causing damage to the SmartOption and to debit the funds from the Client's account in favor of the SmartOption.",
                ]
            },
            {
                text: "8.4. If the Client breaches this Agreement, the SmartOption may, at its option:",
                child: [
                    "8.4.1. Review the amount of the SmartOptions financial obligations to the Client and make changes to the data (balance) of the Clients Account.",
                    "8.4.2. Suspend services to the Client, block access to the trading terminal. Should the SmartOption block the Clients access to the trading terminal, the Client must take all necessary and reasonable steps to remedy the reasons why access to the trading terminal was blocked. If the Client does not take any measures or actions to address the reasons why access was blocked within 30 (thirty) days, the SmartOption may withdraw all funds from the trading account. The SmartOption has a right but is not obliged to redeposit all funds that were withdrawn into the Clients trading account if the Client satisfies all the requirements needed to lift the block on the trading account.",
                ]
            },
            {
                text: "8.5. If the Client breaches any terms of the Agreement and its integral parts listed in clause 1.2, including refusal to undergo the necessary checks and refusal to provide the necessary information, the SmartOption has the right to terminate the Agreement; to void any Client transaction; to close one, several, or all Client trades at any time, at its own discretion; and to stop rendering services to the Client and return funds to the Client or not at its discretion. Any violation of the terms listed in this section, deprives the Customer the right to demand payment or refund from the SmartOption.",
                child: [
                    "8.5.1. If the SmartOption terminates the Agreement with the Client for breach of the Agreement, the Client has no right to open a new account, including by entering third party data during the registration. If the SmartOption detects the Client's breach as set forth in this clause, the consequences stipulated in the section 8.5. of the Agreement will apply.",
                ]
            },
            {
                text: "8.6. The SmartOption assumes no liability to the Client for any damages, losses, lost profit, lost opportunities (without limitations, due to possible market fluctuations), expenses or damages incurred by the Client as a result of the execution of trades pursuant to the terms of the Agreement.",
            },
            {
                text: "8.7. The SmartOption assumes no liability in case of a discrepancy between the information represented on the Clients trading terminal and the information on the SmartOption Server when the financial result of the Clients trades is determined. To eliminate such discrepancy, the SmartOption will adjust the data on the trading terminal in accordance with the information available on the SmartOption Server.",
            },
            {
                text: "8.8. The SmartOption assumes no liability for the Clients damages if those damages resulted from hacker attacks, accidents (failures) of computer networks, communication networks, power lines or telecommunication systems, etc., directly used to determine essential conditions of the Clients transactions or ensure other SmartOption operating procedures that occurred through no fault of the SmartOption.",
            },
            {
                text: "8.9. The SmartOption assumes no liability for technical failures and/or interruptions in trading terminal operation that occurred as a result of hacker attacks, accidents (failures) of computer networks, communication networks, power lines or telecommunication systems, etc., or for the Clients losses resulting from such failure and/or interruption.",
            },
            {
                text: "8.10. The SmartOption assumes no liability for the results of trades that the Client decided to make on the basis of analytical materials provided by the SmartOption and/ or third parties. The Client has been informed that transactions made hereunder bear the risk of not receiving the expected income and the risk of loss of some or all of the money deposited by him/her in the Clients Account. The Client acknowledges that unless there is fraud, deliberate breach of obligations or gross negligence on the part of the SmartOption, the SmartOption will not be liable for any losses, expenses, costs and damages of the Client resulting from inaccuracy of information provided to the Client, including, but not limited to, information about the Clients trading transactions. The SmartOption reserves the right to cancel or close any Client trade under the conditions set forth in this Agreement; nevertheless, all transactions made by the Client as a result of this inaccurate information or an error remain in force and must be fulfilled both by the Client and the SmartOption.",
            },
            {
                text: "8.11. The SmartOption assumes no liability for any losses the Client may suffer in case of theft, loss or disclosure of his/her password to the trading terminal to third parties. The Client assumes full responsibility for protecting the password and safeguarding it against unauthorized third-party access.",
            },
            {
                text: "8.12. The SmartOption assumes no liability for breach (improper discharge) of the obligations hereunder if it was caused by force majeure events or other exceptional conditions stated in the Agreement or appendices to the Agreement.",
            },
            {
                text: "8.13. The SmartOption assumes no liability for any indirect, special, arbitrary, or punitive damages suffered by the Client, including, but not limited to, lost profit, loss of expected savings or loss of income, even if the Client was informed by the SmartOption about the possibility of such damages. Moral harm is not compensated.",
            },
            {
                text: "8.14. The SmartOption reserves the right to consider the Clients breaches at any time, regardless of the time when the breach was made, and, if breaches are discovered, to take measures in accordance with this Agreement.",
            },
        ],
        termsArr9Title: "9. Term and Termination of the Agreement",
        termsArr9TextArr: [
            {
                text: "9.1. This Agreement comes into force upon its conclusion (client registration on the website or on the SmartOption's trading terminal) and is valid sine die.",
            },
            {
                text: "9.2. Either Party may terminate this Agreement unilaterally:",
                child: [
                    "9.2.1. The Agreement is terminated on the initiative of the SmartOption as of the date specified in the notice sent by the SmartOption to the Client;",
                    "9.2.2. The Agreement is terminated at the initiative of the Client in five (5) business days after the SmartOption receives the Clients written notice containing the statement of termination of the Agreement, provided the Client has no unfulfilled obligations hereunder. The Client must send the Termination Notification to the SmartOption's address provided in clause 1.1. of this Agreement, or to the email address sssmartoption@gmail.com.",
                ]
            },
            {
                text: "9.3. This Agreement is considered terminated with respect to the Parties when the mutual obligations of the Client and of the SmartOption with respect to previously made transactions are fulfilled and all debts of each Party are repaid.",
            },
        ],
        termsArr10Title: "10. Final Provisions",
        termsArr10TextArr: [
            {
                text: "10.1. Amendments and additions to this Agreement and to the annexes hereto are made by the SmartOption unilaterally. All amendments and additions made by the SmartOption and not related to the circumstances specified herein come into force on the date specified by the SmartOption.",
            },
            {
                text: "10.2. Amendments and additions made by the SmartOption to this Agreement and to its annexes because of amendments to legislation and regulations governing the subject hereof and to rules and contracts of trading systems used by the SmartOption to discharge its obligations hereunder come into force simultaneously with amendments in the aforementioned documents.",
            },
            {
                text: "10.3. When the amendments and supplements made by the SmartOption come into force, they will apply equally to all Clients, including those who concluded the Agreement before their effective date.",
            },
            {
                text: "10.4. To ensure that a Client who entered into the Agreement is aware of amendments and additions, the Client must visit the SmartOption's Website or the trading platform him/herself or through authorized persons at least once a week to find information about any amendments and/or additions.",
            },
            {
                text: "10.5. The Client (individual) providing his/her personal data to the SmartOption in any form and in any way (when performing any actions on the SmartOption's website, through the SmartOption's counterparties, etc.) thereby gives the SmartOption and its partners his/her consent for automated and non-automated processing of his/her personal data for the purpose of the fulfillment of this Agreement, the implementation of advertising campaigns, provision to him/her of advertising, informational, and marketing materials, information about campaigns and events held by the SmartOption, and for other purposes determined by the SmartOption, namely: to collect, record, systematize, accumulate, store, adjust (update, amend), extract, transfer (disseminate, grant access), depersonalize, block, delete, destroy and transfer across borders personal data. Consent is given for a period of 75 years (or until expiry of the retention periods for the relevant information or documents containing that information determined in accordance with the current legislation of the SmartOption's location). Consent may be withdrawn in accordance with Privacy policy. The SmartOption guarantees the confidentiality of the personal data provided by the Client, excluding the circumstances set by the applicable law and force majeure.",
            },
            {
                text: "10.6. The Client may use information that was provided to the Client either orally or in writing that is posted by the SmartOption or by third parties, access to which was granted to him/her as part of the services hereunder only for transactions under this Agreement. The Client may not disseminate, alter, or supplement the aforementioned information, or store it in separate archives. In any event, the scope of powers granted to the Client with respect to information posted by third parties cannot exceed the scope of powers obtained by the SmartOption from the third party. The SmartOption does not guarantee that information posted by third parties is reliable, accurate, or relevant and is provided on an ongoing basis without interruptions. Nor is the SmartOption liable for the results of transactions (losses, lost profit, lost income, injury to goodwill, etc.) that the Client decided to perform on the basis of information that was communicated to the Client either orally or in writing by the SmartOption or by third parties.",
            },
            {
                text: "10.7. The SmartOption may fully or partially transfer the rights and obligations under this Agreement and the annexes hereto to a third party if such person undertakes to fulfill the terms hereof. This transfer of rights and obligations does not require prior notification of the Client by the SmartOption and shall be made at the time the adequate information is published on the SmartOption's website.",
            },
            {
                text: "10.8. The Client is not entitled to assign his/her own rights, impose his/her own obligations or transfer rights or obligations hereunder in any other way without prior written consent of the SmartOption. If this condition is breached, any such assignment, imposition or transfer will be deemed null and void.",
            },
            {
                text: "10.9. The SmartOption, its partners or any other affiliates may have a material benefit, legal relation or arrangement with respect to a transaction on the trading platform or in the Dashboard or a material benefit, legal relation or arrangement that is in conflict with the Clients interests. For example, the SmartOption may: a) act as a counterparty with respect to any asset; b) suggest another partner of the SmartOption as a counterparty for a trading operation; c) give recommendations and render services to its partners or other clients of the SmartOption with respect to assets they are interested in, despite the fact that this is in conflict with the Clients interests.",
            },
            {
                text: "10.10. The Client agrees and authorizes the SmartOption to act with respect to the Client and for the Client as the SmartOption finds appropriate, despite a potential conflict of interests or the existence of some material interest with respect to any transaction on the trading terminal or in the Dashboard without prior notification of the Client. The existence of a conflict of interest or material benefit with respect to any transaction on the trading terminal or in the Dashboard must not affect the provision of services to the Client by SmartOption employees. From time to time the SmartOption may act on behalf of the Client with parties with whom the SmartOption or any of its related parties have an agreement to receive goods or services. The SmartOption guarantees that these agreements are made as far as possible for the benefit of the Client, for example, these agreements make it possible to access information and other services that would otherwise be inaccessible.",
            },
            {
                text: "10.11. If a court of proper jurisdiction declares any provision of the Agreement (or any part of any provision) void, that provision will be treated as a separate part of the Agreement and this will not affect the legal force of the rest of the Agreement.",
            },
            {
                text: "10.12. The SmartOption may suspend services to the Client at any time (prior notice to the Client is not required).",
            },
            {
                text: "10.13. In situations not described in the Agreement, the SmartOption will act according to market best practices based on principles of honesty and fairness.",
            },
            {
                text: "10.14. The SmartOption may prepare and use texts of the Agreement and its annexes in languages other than English. If there are contradictions between the text of this Agreement and its annexes in English and the corresponding texts in other languages, the text in English will prevail. The text of the Agreement published on the SmartOption website prevails over the text of the Agreement published elsewhere.",
            },
            {
                text: "10.15. The Client is granted a limited and non-exclusive right to use the Trading Terminal only for the purposes stipulated in this Agreement. If the Agreement is terminated on any grounds, the Client's rights to use the Trading Terminal stipulated in this clause shall be terminated at the moment of the termination of the Agreement.",
            },
            {
                text: "10.16. The Client agrees that the SmartOption cannot guarantee continuous uninterrupted and technically sound operation of the Trading Terminal, and therefore the Client accepts this software as is. The SmartOption assumes no liability to the Client for technical failures in Trading Terminal operation.",
            },
            {
                text: "10.17. All terms used in this Agreement and its integral parts, whether they are in uppercase or lowercase characters, shall be of equal importance, unless otherwise follows from the essence of the obligation.",
            },
        ],
        termsArr11Title: "11. List of Countries in which the SmartOption Does Not Do Business",
        termsArr11TextArr: [
            {
                text: "11.1. The SmartOption does not do business in the following countries and does not provide services to persons associated with those states: Gibraltar, the Isle of man, Guernsey, Jersey, Australia, China, Canada, the United States, Japan, Austria, Belgium, Bulgaria, Croatia, Cyprus, Lichtenstein, Czech Republic, Denmark, Estonia, Finland, France, Germany, Greece, Hungary, Ireland, Iceland, Italy, Israel, Latvia, Lithuania, Luxembourg, Malta, Netherlands, Norway, New Zealand, Poland, Portugal, Romania, Slovakia, Slovenia, Spain, Switzerland, Sweden, Islamic Republic of Iran and United Kingdom;",
            },
            {
                text: "11.2. Persons associated with a country are defined as persons:",
                child: [
                    "11.2.1. having citizenship, a residence visa, or other similar document from a country in which the SmartOption does not do business;",
                    "11.2.2. residing in a country in which the SmartOption does not do business, being a resident of a country in which the SmartOption does not do business or having mailing addresses or residence addresses in a country in which the SmartOption does not do business;",
                    "11.2.3. whose places of birth are countries in which the SmartOption does not do business;",
                    "11.2.4. having a telephone number with the country code of a country in which the SmartOption does not do business; having an IP address associated with a country in which the SmartOption does not do business;",
                    "11.2.5. having another connection with a country in which the SmartOption does not do business as defined by the SmartOption at its discretion.",
                ]
            },
            {
                text: "11.3. If it is found that the SmartOption is providing services to persons associated with the countries in which the SmartOption does not do business, the SmartOption may apply the consequences listed in clause 8.5 of the Agreement.",
            },
        ]
    },
    MX: {
        title: 'Terms & Conditions',
        termsArr0Title: 'Contents: ',
        termsArr0TextArr: [
            {text: '1. General Provisions and Subject of the Agreement'},
            {text: '2. Terms and Definitions'},
            {text: '3. Communications and Provision of Information'},
            {text: '4. Terms of use of the SmartOption services'},
            {text: '5. Claims and Dispute Resolution'},
            {text: '6. Governing Law'},
            {text: '7. Force Majeure'},
            {text: '8. Liability of the Parties'},
            {text: '9. Term and Termination of the Agreement'},
            {text: '10. Final Provisions'},
            {text: '11. List of Countries in which the Company Does Not Do Business'},
        ],
        termsArr1Title: '1. General Provisions and Subject of the Agreement',
        termsArr1TextArr: [
            {
                text: "1.1. This Service Agreement has been made by SmartOption App (hereinafter - SmartOption), and the natural person who filled in the registration form on the SmartOption's website or in the trading platform and accepted the terms of this Service Agreement and of the appendices hereto at the moment of registration (hereinafter Client)."
            },
            {
                text: "1.2. The Agreement is an invitation to make offers posted on the SmartOption's website that must be treated as an offer to enter into this Agreement on the terms set forth in it. The posted offer is not public. The SmartOption at its sole discretion may refuse to enter into an Agreement with anyone without explaining the reasons for refusal or, if registration took place, to terminate contractual relations and to block access to the trading terminal. The Client's registration on the SmartOption's Website or on the trading terminal is considered as full and unconditional acceptance of the terms of the Agreement. As soon as the SmartOption receives a payment to replenish the Client's trading account, each Client transaction using the trading terminal or Dashboard becomes the subject of this Agreement."
            },
            {
                text: "1.3. The Client must carefully review the terms of the Agreement. By accepting the terms of this Agreement, the Client agrees to the terms of all annexes to it listed above, including the terms on subdomains of the SmartOption's website that are accessible by the Client and confirms that he/she is an adult who is a legally capable person and is not a resident of a country where trading in options may be deemed illegal."
            },
            {
                text: 'The Client also represents and warrants to the SmartOption that:',
                child: [
                    "1.3.1. All information provided during Client registration and during the performance of Agreement, is true, accurate, reliable and complete in all respects, and the Client completed the registration form him/herself; ",
                    "1.3.2. The Client possesses legal personality to enter into this Agreement, to make inquiries and give orders, and to exercise his rights and to fulfill obligations in accordance with the terms of the Agreement; ",
                    "1.3.3. The Client will carry out trading and non-trading transactions personally, on his own behalf and at his own expense and will not use funds borrowed from other SmartOption Clients or from third parties to carry out transactions. The Client will be guided by the principles of integrity, honesty, and rationality; the Client will not take actions coordinated with other SmartOption Clients aimed at damaging the SmartOption; the Client will not use technical features of the quote stream update on the trading terminal and will not use software errors, defects, and vulnerabilities he discovers in the trading terminal to extract income and will not distribute the information about vulnerabilities to the third parties. The Client will not use unfair and dishonest methods or ways of making trades (transactions) with the SmartOption; the Client will not use insider or confidential information or any other information, as a result of the use of which the Client might benefit when trading with the SmartOption and/or that might damage the SmartOption; ",
                    "1.3.4. The Client will adhere to legal norms, in particular, but not limited to, international norms aimed at controlling illegal trade, financial fraud, and money laundering; ",
                    "1.3.5. The Client will not use the trading terminal or website to collude in illegal financial activities or any other illegal transaction; ",
                    "1.3.6. The money listed by the Client on the SmartOption's account has legal origin. The Client legally owns the money and has the right to use and manage it. The Client's account will not be replenished from third party payment instruments. ",
                    "The Client will not replenish third party Client accounts or withdraw money from the Client's account to third party payment instruments. ",
                    "1.3.7. No actions of the Client pursuant to the Agreement will violate any law, regulation, right, bylaws, or rules and regulations applicable to the Client or in the jurisdiction where he/she resides or the provisions of any other agreement binding on the Client or involving any assets of the Client ",
                    "1.3.8. To carry out transactions, the Client will use account data from Trading Terminal belonging to him and will not transfer account data to third parties and will not use account data of the SmartOption's clients for trading and (or) non-trading operations; ",
                    "1.3.9. The Client is not a federal or municipal employee, an employee of a federal or municipal institution, an employee of a federal or municipal organization, an organization in whose capital the state has a prevailing interest; the Client is not a politically significant person, a family member or a relative of a politically significant person; the Client is not a person closely connected with a politically significant person; the Client is not a person connected with the United States or with another country in which the SmartOption does not operate. The terms used in this paragraph are interpreted and applied by the SmartOption at its own discretion in accordance with the norms of international legislation and/or legislation of a particular state, generally accepted terms and definitions, customary business practices. ",
                ]
            },
            {
                text: '1.4. The subject of the Agreement is the definition of the general conditions under which the Parties carry out transactions (trades), the content and procedure for which are set forth in this Agreement. The SmartOption unilaterally sets and may alter at its sole discretion essential conditions for a transaction (trade), may limit the number of trades executed simultaneously, and limit the number of trades that a Client may make within the time period set by the SmartOption, may set other limits for trades at its sole discretion and unilaterally.'
            },
            {
                text: '1.5. The SmartOption may engage third parties to perform this Agreement. The SmartOption is not responsible for services provided by such third parties.'
            },
        ],
        termsArr2Title: '2. Terms and Definitions',
        termsArr2TextArr: [
            {
                text: ".1. Asset - the underlying financial instrument of an option. An asset may be company stock, a stock index, currency pair (the rate of an exchange of one currency into another currency), goods quoted on a commodity exchange, commodity options, etc."
            },
            {
                text: "2.2. Client's Account Balance - the amount of money in the Client's account."
            },
            {
                text: "2.3. Option - a derivative financial instrument (trade), consisting of two transactions: Making a trade and Closing a trade, as a result of the fulfillment of which the Client either receives Income or loses the Trade amount. The SmartOption does not, however, provide a Client with the ability to assume obligations greater than the Trade amount (the SmartOption does not provide leverage to the Client). The SmartOption and the Client make a trade with an Higher/Lower option - an option, the targeted level of which equals the asset price or another price determined by the Client at the SmartOption's suggestion when the trade is made, only payment on which is made if an upward price change was selected and the asset price was greater than the targeted level when the option expired or if a downward price change was selected and the asset price was less than the targeted level when the option expired. If the asset price is the same as the targeted level when the option expires, the trade Amount is returned to the Client"
            },
            {
                text: "2.4. Bonus - virtual cash credited by the SmartOption to the Client's Account when the Client meets the conditions established by the SmartOption. When the Client executes trading transactions, funds deposited by the Client must be used first and, only after those funds are completely exhausted, does the Client have the right to use the Bonus for further trading transactions. In general, the Client cannot withdraw the Bonus amount to the Client's External Account. If the Client withdraws funds that were previously credited to his External account, the SmartOption has the right to debit the entire amount of previously accrued Bonuses from the Client's Account. The SmartOption has the right to set a minimum trading amount and other terms under which the Bonus is not debited from the Client's Account and may be withdrawn to the Client's External account. The SmartOption may publish credit and debit rules, rules for using and converting Bonuses to real funds, and rules for other transactions with Bonuses and Client funds that the Client accrued along with Bonuses on its website."
            },
            {
                text: "2.5. SmartOption Website - an internet site at the address (domain name) SmartOption.com, and other internet sites which the SmartOption references in the Agreement or on the SmartOption Website."
            },
            {
                text: "2.6. Client External Account - the Client's current account at a lending institution, an account (wallet) on an electronic payment system."
            },
            {
                text: "2.7. Option Expiration Time - the time when the condition for payout of income on the option is checked."
            },
            {
                text: "2.8. Trading Hours - the time when trading in an asset may be done."
            },
            {
                text: "2.9. Funds Withdrawal - transfer of money from the Clients Account to the Clients External Account."
            },
            {
                text: "2.10. Income - fixed remuneration credited to the Client's Account if the conditions of the option are met. The Income amount is determined at the time that the trade is executed. Income is credited to the Clients account from the SmartOption's guaranteed (compensatory, special) fund (reserve fund), which is formed in the manner determined by the SmartOption at its own expense."
            },
            {
                text: "2.11. Execution of a trade - a trading transaction in which the essential conditions of an option trade are agreed by the Client and the SmartOption. After the essential conditions are agreed, the option trade is open."
            },
            {
                text: "2.12. Closing a trade - a trading transaction in which an open trade on an option is liquidated at the option expiration time agreed by the Parties when the trade is made. Early closing of a trade - closing a trade at the Client's initiative before it is closed at the option expiry time. Early closing of a trade is possible only if the SmartOption has the technical capability to do so. The SmartOption may unilaterally at its own discretion provide or refuse to grant the Client the right to make an early closing of a trade."
            },
            {
                text: "2.13. Exceptional Situations - abnormal market conditions and other exceptional conditions described in the Regulation on Trading Transactions."
            },
            {
                text: "2.14. Payout ratio - the percentage that determines the Income amount set by the SmartOption depending on the option asset and other conditions of the option."
            },
            {
                text: "2.15. Asset Quote - the price for one unit of the Asset."
            },
            {
                text: "2.16. Log entry - an entry in the database made by the SmartOption server that, with accuracy to the millisecond or, in the case of absence of technical capability, with accuracy to the second, documents all Client inquiries and orders and the result of their processing. Each Client contact with the trading terminal and dashboard is documented in a log entry. These server data are the primary source of information and are used by the parties as evidence if disputes related to the performance of the Agreement arise. Information from a SmartOption server Log entry has unconditional priority over all other arguments during settlement of a dispute, including over information from the Client's trading terminal Log file. The SmartOption reserves the right not to keep Log entries."
            },
            {
                text: "2.17. Price Direction - an essential condition of an option trade that determines the trade payout. A price may change direction Higher or Lower."
            },
            {
                text: "2.18. Non-trading Transaction - any Client transaction to deposit funds on the Client's Account or withdraw funds from the Client's Account."
            },
            {
                text: "2.19. Transactions - The Client's trading and non-trading transactions."
            },
            {
                text: "2.20. Open Trade - an option trade after the execution of the trade and before trade closure, for which it is not determined yet whether a payout will be made."
            },
            {
                text: "2.21. Price Feed - a succession of prices displayed on the Trading Terminal."
            },
            {
                text: "2.22. Payment service provider - a company that provides funds transfer services."
            },
            {
                text: "2.23. Recurring payment - a periodically repeating transaction to replenish the Balance of the Client's account without requiring the latter to reenter his or her credit card information."
            },
            {
                text: "2.24. SmartOption Server - the SmartOption's software used to process and store information about client requests for trading and non-trading transactions, to provide to the Client realtime information about quotes, to account for trading and non-trading transactions, to monitor the fulfillment of trade conditions and limit trading transactions, and to determine the financial result of trades."
            },
            {
                text: "2.25. Withdrawal Method - one of the Funds Withdrawal methods offered to the Client that is posted on the trading terminal and in the Dashboard."
            },
            {
                text: "2.26. The Trade Amount is the amount paid by the Client to the SmartOption when the trade is made. The Client's profit on an option trade if the trade conditions are fulfilled is determined by the procedure set forth in Section 2.10. of this Agreement."
            },
            {
                text: "2.27. SmartOption's Account - the SmartOption's settlement account at a financial institution, an account (wallet) in an electronic payment system, and other accounts, including accounts of Payment Agents."
            },
            {
                text: "2.28. Essential Conditions of a trading transaction (essential conditions of a trade) - the conditions that govern the payout of income from a trade to the Client by the SmartOption."
            },
            {
                text: "2.29. Client account (Trading account) - a special account in the SmartOption's accounting system in which funds transferred by the Client to make trades are posted; from which the trade amount is debited when the trade is executed; and to which income is credited when a trade is closed and the essential conditions of a trade are met. The Client has the right to have only one Client account. In case of violation of this rule, the SmartOption has the right to refuse the Client in further service (rendering of services), to terminate this Agreement, and to block further possibility to carry out transactions without explanation and without payment of funds from the Client's account. It is not a violation of this clause of the Agreement when the SmartOption, if possible and unilaterally, grants the Client the right to use several currencies within the Client's account, as well as the right to use the Client's account in relations between the SmartOption and the Client, which are governed simultaneously by this Agreement and other agreements, concluded between the SmartOption and the Client, under which the SmartOption has the discretion to grant the Client the right to use the Client's account to engage in trading transactions which are not prescribed in this Agreement."
            },
            {
                text: "2.30. Trading transactions - procedures to make and close trades with options between the SmartOption and a Client. Trading transactions are carried out at the place of the SmartOption's registration. No physical delivery of assets takes place during trading transactions. Trade amounts for trading transactions are debited from the Client's account balance after a trade is made. The Income on trading transactions are credited to the Client's account immediately after a trade closes."
            },
            {
                text: "2.31. The Trading Terminal is software through which the Client can obtain real-time information about quotes to carry out trading and non-trading transactions and receive messages from the SmartOption. Entry to the Trading Terminal is protected by a password that the Client sets at registration on the SmartOption's site. All orders and inquiries carried out through the Trading Terminal are deemed to be carried out personally by the Client. Clients from countries whose laws prohibit trading in options or other over-the-counter derivatives and the employees, affiliates, agents and other representatives of the SmartOption and their relatives are prohibited from using the Trading Terminal. The part of the Trading Terminal with which a Client can carry out nontrading transactions in this Agreement may be called the Dashboard."
            },
            {
                text: "2.32. Red Territory - the state of an open option trade if, at the current asset price, income cannot be paid out on that trade."
            },
            {
                text: "2.33. 1-Click service - makes it possible to replenish a Client's account balance from the Client's bank (payment) cards without entering the bank (payment) card data of the bank card holder again."
            },
            {
                text: "2.34. Targeted level - the level of asset prices with respect to which the trade result is calculated."
            },
            {
                text: "2.35. Cookie file - a small data set including an anonymous unique identifier which is sent to the web browser of the Client's computer or mobile telephone (hereinafter 'device') from the SmartOption server (website) and is stored on the Client's device. The Client may configure a web browser to block cookie access to the device. When a Client visits the SmartOption Website, the viewed pages and cookie files are downloaded onto the Client's device. Cookie files stored on the device may be used for anonymous identification if the Client revisits the SmartOption's website and to determine the website pages that are the most popular among our Clients. Cookie files stored on the hard drive of the Client's device enable the SmartOption to create the most user-friendly and efficient website for SmartOption Clients, providing us an opportunity to identify our Clients' preferences."
            },
            {
                text: "2.36. Trading signals are information about the state of the market gathered by the SmartOption based on analytical conclusions that the SmartOption is entitled to provide at its discretion to some or all Clients with respect to certain market indicators. Trading signals are not an offer and are not the SmartOption's explicit recommendation to Clients to engage in trading transactions or to make trades. The SmartOption is not liable for the accuracy of Trading Signals or for Client trading transactions and trades based on Trading Signals. The Client may, at his discretion, consider or ignore Trading signals when engaging in trading transactions and trades."
            },
            {
                text: "2.37. Quote - the current Asset price displayed on the Trading Terminal. Terms used in this Agreement and definitions absent in this Clause shall be interpreted in accordance with usual and customary business practices applicable in regard to conclusion of agreements with financial derivatives."
            },
        ],
        termsArr3Title: '3. Communications and Provision of Information',
        termsArr3TextArr: [
            {
                text: '3.1. To communicate with the Client, the SmartOption may use:',
                child: [
                    '· e-mail;',
                    '· fax;',
                    '· telephone;',
                    '· SMS;',
                    '· letters sent by post;',
                    '· different types of messages sent to the Client on the trading terminal, in the Dashboard, browser window, etc. (push notifications, reminders, service messages, etc.);',
                    '· announcements on the SmartOption website.',
                ]
            },
            {
                text: "3.2. To promptly communicate with the Client to resolve problems related to Client transactions, the SmartOption will use the Client's contact information entered when the Client registered or modified the account pursuant to clause 4.5 of the Agreement. The Client agrees to accept message from the SmartOption at any time. ",
            },
            {
                text: "3.3. Any correspondence (documents, notices, confirmations, announcements, reports etc.) are deemed received by the Client:",
                child: [
                    "1) one (1) hour after they are sent to the email address (e-mail);",
                    "2) immediately after they are faxed;",
                    "3) immediately after the telephone call ends;",
                    "4) immediately after an SMS message is sent;",
                    "5) seven (7) calendar days after they are mailed",
                    "6) immediately after an announcement is posted on the SmartOption website.",
                ]
            },
            {
                text: "3.4. The Client may also contact the SmartOption by e-mail at sssmartoption@gmail.com, and other email addresses and the telephone numbers listed in this Agreement and on the SmartOption website."
            },
            {
                text: "3.5. The Client understands and agrees that, if the Client's behavior during conversations with SmartOption representative is inappropriate, the SmartOption reserves the right to unilaterally terminate this Agreement."
            },
            {
                text: "3.6. The SmartOption may use contact information provided by the Client to send informational, marketing, and advertising materials, and service messages and to resolve other tasks. The SmartOption will determine the frequency with which it sends messages to the Client at its sole discretion. If the Client wishes not to receive from informational (and other) messages from the SmartOption, he/she must unsubscribe by clicking on the Unsubscribe link (if the message format provides this ability) or by contacting the client support department."
            },
        ],
        termsArr4Title: '4. Terms of use of the SmartOption services.',
        termsArr4TextArr: [
            {
                text: "4.1. During registration the Client undertakes to provide correct and reliable identification information in accordance with requirements of the Client registration form."
            },
            {
                text: "4.2. After successful registration, the Client will be granted access to the trading terminal, the ability to transfer funds to the Clients Account (placing a deposit in the Clients Account to be able to make option trades) and to perform other operations."
            },
            {
                text: "4.3. The Client must promptly inform the SmartOption about changes in identification and contact information (within 7 (seven) days after that change) by making the appropriate changes on the Trading Terminal or in any other way offered by the SmartOption. To identify the Client and to inspect the origin of the Client's funds at any time after registration, the SmartOption has the right to ask and the Client obliges to provide within 7 (seven) days after the request is received, any identification documents (these documents include identity documents; documents confirming the residential address, documents indicating the Client's financial standing, and other documents at the SmartOption's discretion). The SmartOption reserves the right to suspend non-trading and/or trading transactions on the Client's account if it finds that the Client's identification information is incorrect or inaccurate; if the Client has not provided the requested documents, the SmartOption has the right to block the Client's access to the trading terminal until the client's identification procedure is completed. The SmartOption also has the right to demand the Client to undergo the identification procedure by a personal visit to a SmartOption-authorized agent and provide documents, the list of which is determined by the SmartOption at its discretion."
            },
            {
                text: "4.4. Login to the trading terminal is password protected.",
                child: [
                    "4.4.1. The Client confirms and agrees that access to the trading terminal will be protected with a password set by the Client himself/herself during registration. The Client may not transfer the trading terminal password to third parties. ",
                    "4.4.2. The Client assumes full responsibility for password protection and prevention of unauthorized third-party access to it. ",
                    "4.4.3. All orders through the trading terminal with the Client's password will be deemed to have been made by the Client unless the SmartOption establishes otherwise. ",
                    "4.4.4. Any person who obtains access to the trading terminal by entering the Client's password will be identified as the Client unless otherwise specified by the SmartOption. ",
                    "4.4.5. The SmartOption assumes no liability for any losses the Client may suffer in case of theft, loss or disclosure of the password to third parties or in case of unauthorized use of registration data by third parties. ",
                ]
            },
            {
                text: "4.5. The Client may change the password to the trading terminal himself/herself or use the procedure for password recovery set by the SmartOption."
            },
        ],
        termsArr5Title: "5. Claims and Dispute Resolution",
        termsArr5TextArr: [
            {
                text: "5.1. The Parties have agreed that they will make every effort to settle all disputes between the SmartOption and the Client related to transactions, payouts and other actions hereunder by means of negotiations."
            },
            {
                text: "5.2. If a dispute arises, the Client may submit a claim/complaint to the SmartOption and send a statement or notice to the SmartOption. All claims/complaints/statements/ notices related to transactions performed by the Client must be submitted in accordance with the following requirements:",
                child: [
                    "5.2.1. claims/complaints/statements/notices must be submitted in writing;",
                    "5.2.2. claims/complaints/statements/notices must contain the following information: Surname, first name, middle name (if any), Clients e-mail, Clients Account Number, date and time of the dispute, brief description of the dispute, Clients demands; claimed amount and justified calculation thereof (if the claim can be valued in money); circumstances that are the grounds for the claim and evidence supporting those circumstances, including reference to the clause of this Agreement (annexes hereto) that were breached in the Clients opinion; a list of documents and other evidence attached to the claim (complaint) authenticated by the Client; other information necessary to settle the dispute;",
                    "5.2.3. claims/complaints/statements/notices must be sent by the Client within five (5) business days after the event that was the basis for submission of the relevant claim (complaint). The Client agrees that a delay in submission of the claim (complaint) is grounds for refusal to consider it;",
                    "5.2.4. claims/complaints/statements/notices may be sent by e-mail to sssmartoption@gmail.com, by registered or certified mail. Claims/ complaints/statements/notices made and sent in another form will not be considered.",
                ]
            },
            {
                text: "5.3. Claims/complaints/statements/notices must not contain:",
                child: [
                    "a) an emotional evaluation of the dispute;",
                    "b) offensive statements addressed to the SmartOption;",
                    "c) profanity.",
                ]
            },
            {
                text: "5.4. To respond to a claim/complaint/statement/notice, the SmartOption may request additional documents and information from the Client. A claim/complaint/statement/ notice will be reviewed on the basis of data provided by the Client and log entries from the SmartOption server. Log entries from the SmartOption Server always prevail over other evidence and proof. The SmartOption assumes no liability for incomplete trades and will not compensate any financial damages or moral harm suffered by the Client with respect to what the Client considers to be lost profit. When considering disputes, the Clients references to information from other companies and websites are not considered.",
            },
            {
                text: "5.5. The SmartOption may reject a claim/complaint/statement/notice if the terms of this section are violated.",
            },
            {
                text: "5.6. The SmartOption must consider a claim/complaint/application/appeal within no more than 10 working days after the submission date. This term does not include time for provision of additional documents by the Client upon the SmartOption's request.",
            },
            {
                text: "5.7. In addition to provision set forth in clauses 5.2.-5.6. of the dispute resolution procedure, the Client may file a claim to a court, provided that he/she has first submitted the claim to the SmartOption in accordance with this dispute resolution procedure. The claim procedure for dispute resolution will be deemed followed if: a) the form and content of the claim meet the requirements of clauses 5.2.1., 5.2.2., and 5.2.3.; b) the claim is sent to the SmartOption's registration address by registered or certified mail; c) the Client has a confirmation of claim receipt by the SmartOption; d) the deadline for responding to the claim has expired. Claim response time - sixty (60) calendar days after it is received by the SmartOption.",
            },
            {
                text: "5.9. In case of any disputes, the SmartOption reserves the right to fully or partially block transactions in the Client's Account until the dispute is settled or until the Parties come to an interim agreement.",
            },
        ],
        termsArr6Title: "6. Governing Law",
        termsArr6TextArr: [
            {
                text: "6.1. The Client expressly:",
                child: [
                    "a) agrees that the courts of the country of the SmartOption's registration have exclusive jurisdiction to conduct any legal proceedings with respect to this Agreement;",
                    "b) submits to the jurisdiction of the courts of the country of the SmartOption's registration;",
                    "c) waives any appeals with respect to proceedings in any of such courts;",
                    "d) agrees to make no claims regarding trials location as non-convenient and not to declare that the trials location state has no legal jurisdiction over the Client.",
                ]
            },
        ],
        termsArr7Title: "7. Force Majeure",
        termsArr7TextArr: [
            {
                text: "7.1. If the SmartOption has sufficient grounds, it may claim force majeure events. Force majeure events include (without limitation):",
                child: [
                    "(a) any action, event or occurrence (including, but not limited to, any strike, riots or civil strife, terrorist acts, wars, natural disasters, accidents, fires, floods, storms, power outages, interruptions in the operation of communication equipment, software or electronic equipment, incorrect operation of any kind of equipment or software, the instability in the quote stream, interruptions in the operation of or the instability of liquidity providers, etc.), which, in the SmartOption's reasonable opinion, led to the destabilization of the market or the markets for one or more assets (instruments);",
                    "b) the suspension of work, the liquidation or closure of any market or the absence of any event on which the SmartOption bases quotes, or the imposition of restrictions or special or non-standard terms of trade in any market, or in respect of any such event.",
                ]
            },
            {
                text: "7.2. If the SmartOption has established that a force majeure event occurred, the SmartOption has the right (without prejudice to the SmartOption's other rights) to take any of the following steps without prior written notification and at any time: a) cancel any or all trades the result of which is directly or indirectly caused by force majeure; (b) suspend or amend the application of one or all provisions of the Agreement as long as the force majeure event makes it impossible for the SmartOption to comply with these provisions; c) take or, on the contrary, not take any action in respect of the SmartOption, the Client or other clients, if the SmartOption reasonably deems it appropriate under the circumstances.",
            },
            {
                text: "7.3 The SmartOption assumes no liability for breach (improper discharge) of obligations if force majeure events interfered with that discharge.",
            },
        ],
        termsArr8Title: "8. Liability of the Parties",
        termsArr8TextArr: [
            {
                text: "8.1. The liability of the Parties to this Agreement is determined by the terms of the Agreement and its annexes.",
            },
            {
                text: "8.2. The SmartOption assumes liability only for real damages caused to the Client as a result of the SmartOption's deliberate breach of its obligations hereunder. The SmartOption is as responsible for the actions of its representatives, departments, and payment agents as it is for its own actions.",
            },
            {
                text: "8.3. The Client assumes liability to the SmartOption for losses incurred by the SmartOption for which the Client is responsible, including:",
                child: [
                    "a) for damage caused as a result of the Client's failure to provide (or late provision) of any documents that must be provided to the SmartOption under this Agreement and its annexes and for damage caused to SmartOption because of any misstatement of information contained in documents provided by the Client;",
                    "b) for damage caused to the SmartOption because of abuse of services provided by the SmartOption to the Client, including damage (losses) caused to the SmartOption by the use of robotic and automated transaction algorithms and/or special software tools and other tools, devices, methods and techniques that facilitate or contribute to the violation of the principle of integrity, honesty and fairness in the execution of transactions;",
                    "c) for damage caused by the Client as a result of the actions coordinated with other clients of the SmartOption and/or affiliates of the Client aimed at causing the SmartOption losses and for other damage caused to the SmartOption by the Client from the use of other unfair and dishonest methods and techniques for making trades (transactions) with the SmartOption, including using bonuses. In any situation 'Client affiliates' means persons in a kinship relationship of any degree; marital, partnership, or other relationships; residing at the same address as the person; persons using the same devices; persons engaged as a SmartOption Client by the same partner or Client of the SmartOption; and persons engaged in any joint activity with or without the formation of a legal entity. The SmartOption reserves the right to expand the list of situations and attributes in which the Client and third parties may be recognized as affiliated;",
                    "d) if there is sufficient evidence to suggest that the Client unlawfully attempted to use the software provided by the SmartOption and the funds transferred to the SmartOption's account;",
                    "e) for damage caused to the SmartOption as a result of extracting income from the use of the technical features of the quote stream update on the trading terminal and of extracting income from the use of software errors and vulnerabilities in the trading terminal;",
                    "f) for damage caused to the SmartOption by the Client's use of insider, confidential, or other information that provided the Client with any kind of advantages in concluding trades with the SmartOption. The SmartOption has the right to debit these losses from the Client's account and/or the accounts of other persons (if it is established that these accounts belong to the Client (or the Client's accomplices) using the SmartOption's technical and other equipment and tools). The SmartOption also has the right to block the further transactions on the trading terminal and the dashboard for Clients in respect to whom the SmartOption has sufficient grounds and suspicions to classify their actions (including joint actions with other Client's) as aimed at causing damage to the SmartOption and to debit the funds from the Client's account in favor of the SmartOption.",
                ]
            },
            {
                text: "8.4. If the Client breaches this Agreement, the SmartOption may, at its option:",
                child: [
                    "8.4.1. Review the amount of the SmartOptions financial obligations to the Client and make changes to the data (balance) of the Clients Account.",
                    "8.4.2. Suspend services to the Client, block access to the trading terminal. Should the SmartOption block the Clients access to the trading terminal, the Client must take all necessary and reasonable steps to remedy the reasons why access to the trading terminal was blocked. If the Client does not take any measures or actions to address the reasons why access was blocked within 30 (thirty) days, the SmartOption may withdraw all funds from the trading account. The SmartOption has a right but is not obliged to redeposit all funds that were withdrawn into the Clients trading account if the Client satisfies all the requirements needed to lift the block on the trading account.",
                ]
            },
            {
                text: "8.5. If the Client breaches any terms of the Agreement and its integral parts listed in clause 1.2, including refusal to undergo the necessary checks and refusal to provide the necessary information, the SmartOption has the right to terminate the Agreement; to void any Client transaction; to close one, several, or all Client trades at any time, at its own discretion; and to stop rendering services to the Client and return funds to the Client or not at its discretion. Any violation of the terms listed in this section, deprives the Customer the right to demand payment or refund from the SmartOption.",
                child: [
                    "8.5.1. If the SmartOption terminates the Agreement with the Client for breach of the Agreement, the Client has no right to open a new account, including by entering third party data during the registration. If the SmartOption detects the Client's breach as set forth in this clause, the consequences stipulated in the section 8.5. of the Agreement will apply.",
                ]
            },
            {
                text: "8.6. The SmartOption assumes no liability to the Client for any damages, losses, lost profit, lost opportunities (without limitations, due to possible market fluctuations), expenses or damages incurred by the Client as a result of the execution of trades pursuant to the terms of the Agreement.",
            },
            {
                text: "8.7. The SmartOption assumes no liability in case of a discrepancy between the information represented on the Clients trading terminal and the information on the SmartOption Server when the financial result of the Clients trades is determined. To eliminate such discrepancy, the SmartOption will adjust the data on the trading terminal in accordance with the information available on the SmartOption Server.",
            },
            {
                text: "8.8. The SmartOption assumes no liability for the Clients damages if those damages resulted from hacker attacks, accidents (failures) of computer networks, communication networks, power lines or telecommunication systems, etc., directly used to determine essential conditions of the Clients transactions or ensure other SmartOption operating procedures that occurred through no fault of the SmartOption.",
            },
            {
                text: "8.9. The SmartOption assumes no liability for technical failures and/or interruptions in trading terminal operation that occurred as a result of hacker attacks, accidents (failures) of computer networks, communication networks, power lines or telecommunication systems, etc., or for the Clients losses resulting from such failure and/or interruption.",
            },
            {
                text: "8.10. The SmartOption assumes no liability for the results of trades that the Client decided to make on the basis of analytical materials provided by the SmartOption and/ or third parties. The Client has been informed that transactions made hereunder bear the risk of not receiving the expected income and the risk of loss of some or all of the money deposited by him/her in the Clients Account. The Client acknowledges that unless there is fraud, deliberate breach of obligations or gross negligence on the part of the SmartOption, the SmartOption will not be liable for any losses, expenses, costs and damages of the Client resulting from inaccuracy of information provided to the Client, including, but not limited to, information about the Clients trading transactions. The SmartOption reserves the right to cancel or close any Client trade under the conditions set forth in this Agreement; nevertheless, all transactions made by the Client as a result of this inaccurate information or an error remain in force and must be fulfilled both by the Client and the SmartOption.",
            },
            {
                text: "8.11. The SmartOption assumes no liability for any losses the Client may suffer in case of theft, loss or disclosure of his/her password to the trading terminal to third parties. The Client assumes full responsibility for protecting the password and safeguarding it against unauthorized third-party access.",
            },
            {
                text: "8.12. The SmartOption assumes no liability for breach (improper discharge) of the obligations hereunder if it was caused by force majeure events or other exceptional conditions stated in the Agreement or appendices to the Agreement.",
            },
            {
                text: "8.13. The SmartOption assumes no liability for any indirect, special, arbitrary, or punitive damages suffered by the Client, including, but not limited to, lost profit, loss of expected savings or loss of income, even if the Client was informed by the SmartOption about the possibility of such damages. Moral harm is not compensated.",
            },
            {
                text: "8.14. The SmartOption reserves the right to consider the Clients breaches at any time, regardless of the time when the breach was made, and, if breaches are discovered, to take measures in accordance with this Agreement.",
            },
        ],
        termsArr9Title: "9. Term and Termination of the Agreement",
        termsArr9TextArr: [
            {
                text: "9.1. This Agreement comes into force upon its conclusion (client registration on the website or on the SmartOption's trading terminal) and is valid sine die.",
            },
            {
                text: "9.2. Either Party may terminate this Agreement unilaterally:",
                child: [
                    "9.2.1. The Agreement is terminated on the initiative of the SmartOption as of the date specified in the notice sent by the SmartOption to the Client;",
                    "9.2.2. The Agreement is terminated at the initiative of the Client in five (5) business days after the SmartOption receives the Clients written notice containing the statement of termination of the Agreement, provided the Client has no unfulfilled obligations hereunder. The Client must send the Termination Notification to the SmartOption's address provided in clause 1.1. of this Agreement, or to the email address sssmartoption@gmail.com.",
                ]
            },
            {
                text: "9.3. This Agreement is considered terminated with respect to the Parties when the mutual obligations of the Client and of the SmartOption with respect to previously made transactions are fulfilled and all debts of each Party are repaid.",
            },
        ],
        termsArr10Title: "10. Final Provisions",
        termsArr10TextArr: [
            {
                text: "10.1. Amendments and additions to this Agreement and to the annexes hereto are made by the SmartOption unilaterally. All amendments and additions made by the SmartOption and not related to the circumstances specified herein come into force on the date specified by the SmartOption.",
            },
            {
                text: "10.2. Amendments and additions made by the SmartOption to this Agreement and to its annexes because of amendments to legislation and regulations governing the subject hereof and to rules and contracts of trading systems used by the SmartOption to discharge its obligations hereunder come into force simultaneously with amendments in the aforementioned documents.",
            },
            {
                text: "10.3. When the amendments and supplements made by the SmartOption come into force, they will apply equally to all Clients, including those who concluded the Agreement before their effective date.",
            },
            {
                text: "10.4. To ensure that a Client who entered into the Agreement is aware of amendments and additions, the Client must visit the SmartOption's Website or the trading platform him/herself or through authorized persons at least once a week to find information about any amendments and/or additions.",
            },
            {
                text: "10.5. The Client (individual) providing his/her personal data to the SmartOption in any form and in any way (when performing any actions on the SmartOption's website, through the SmartOption's counterparties, etc.) thereby gives the SmartOption and its partners his/her consent for automated and non-automated processing of his/her personal data for the purpose of the fulfillment of this Agreement, the implementation of advertising campaigns, provision to him/her of advertising, informational, and marketing materials, information about campaigns and events held by the SmartOption, and for other purposes determined by the SmartOption, namely: to collect, record, systematize, accumulate, store, adjust (update, amend), extract, transfer (disseminate, grant access), depersonalize, block, delete, destroy and transfer across borders personal data. Consent is given for a period of 75 years (or until expiry of the retention periods for the relevant information or documents containing that information determined in accordance with the current legislation of the SmartOption's location). Consent may be withdrawn in accordance with Privacy policy. The SmartOption guarantees the confidentiality of the personal data provided by the Client, excluding the circumstances set by the applicable law and force majeure.",
            },
            {
                text: "10.6. The Client may use information that was provided to the Client either orally or in writing that is posted by the SmartOption or by third parties, access to which was granted to him/her as part of the services hereunder only for transactions under this Agreement. The Client may not disseminate, alter, or supplement the aforementioned information, or store it in separate archives. In any event, the scope of powers granted to the Client with respect to information posted by third parties cannot exceed the scope of powers obtained by the SmartOption from the third party. The SmartOption does not guarantee that information posted by third parties is reliable, accurate, or relevant and is provided on an ongoing basis without interruptions. Nor is the SmartOption liable for the results of transactions (losses, lost profit, lost income, injury to goodwill, etc.) that the Client decided to perform on the basis of information that was communicated to the Client either orally or in writing by the SmartOption or by third parties.",
            },
            {
                text: "10.7. The SmartOption may fully or partially transfer the rights and obligations under this Agreement and the annexes hereto to a third party if such person undertakes to fulfill the terms hereof. This transfer of rights and obligations does not require prior notification of the Client by the SmartOption and shall be made at the time the adequate information is published on the SmartOption's website.",
            },
            {
                text: "10.8. The Client is not entitled to assign his/her own rights, impose his/her own obligations or transfer rights or obligations hereunder in any other way without prior written consent of the SmartOption. If this condition is breached, any such assignment, imposition or transfer will be deemed null and void.",
            },
            {
                text: "10.9. The SmartOption, its partners or any other affiliates may have a material benefit, legal relation or arrangement with respect to a transaction on the trading platform or in the Dashboard or a material benefit, legal relation or arrangement that is in conflict with the Clients interests. For example, the SmartOption may: a) act as a counterparty with respect to any asset; b) suggest another partner of the SmartOption as a counterparty for a trading operation; c) give recommendations and render services to its partners or other clients of the SmartOption with respect to assets they are interested in, despite the fact that this is in conflict with the Clients interests.",
            },
            {
                text: "10.10. The Client agrees and authorizes the SmartOption to act with respect to the Client and for the Client as the SmartOption finds appropriate, despite a potential conflict of interests or the existence of some material interest with respect to any transaction on the trading terminal or in the Dashboard without prior notification of the Client. The existence of a conflict of interest or material benefit with respect to any transaction on the trading terminal or in the Dashboard must not affect the provision of services to the Client by SmartOption employees. From time to time the SmartOption may act on behalf of the Client with parties with whom the SmartOption or any of its related parties have an agreement to receive goods or services. The SmartOption guarantees that these agreements are made as far as possible for the benefit of the Client, for example, these agreements make it possible to access information and other services that would otherwise be inaccessible.",
            },
            {
                text: "10.11. If a court of proper jurisdiction declares any provision of the Agreement (or any part of any provision) void, that provision will be treated as a separate part of the Agreement and this will not affect the legal force of the rest of the Agreement.",
            },
            {
                text: "10.12. The SmartOption may suspend services to the Client at any time (prior notice to the Client is not required).",
            },
            {
                text: "10.13. In situations not described in the Agreement, the SmartOption will act according to market best practices based on principles of honesty and fairness.",
            },
            {
                text: "10.14. The SmartOption may prepare and use texts of the Agreement and its annexes in languages other than English. If there are contradictions between the text of this Agreement and its annexes in English and the corresponding texts in other languages, the text in English will prevail. The text of the Agreement published on the SmartOption website prevails over the text of the Agreement published elsewhere.",
            },
            {
                text: "10.15. The Client is granted a limited and non-exclusive right to use the Trading Terminal only for the purposes stipulated in this Agreement. If the Agreement is terminated on any grounds, the Client's rights to use the Trading Terminal stipulated in this clause shall be terminated at the moment of the termination of the Agreement.",
            },
            {
                text: "10.16. The Client agrees that the SmartOption cannot guarantee continuous uninterrupted and technically sound operation of the Trading Terminal, and therefore the Client accepts this software as is. The SmartOption assumes no liability to the Client for technical failures in Trading Terminal operation.",
            },
            {
                text: "10.17. All terms used in this Agreement and its integral parts, whether they are in uppercase or lowercase characters, shall be of equal importance, unless otherwise follows from the essence of the obligation.",
            },
        ],
        termsArr11Title: "11. List of Countries in which the SmartOption Does Not Do Business",
        termsArr11TextArr: [
            {
                text: "11.1. The SmartOption does not do business in the following countries and does not provide services to persons associated with those states: Gibraltar, the Isle of man, Guernsey, Jersey, Australia, China, Canada, the United States, Japan, Austria, Belgium, Bulgaria, Croatia, Cyprus, Lichtenstein, Czech Republic, Denmark, Estonia, Finland, France, Germany, Greece, Hungary, Ireland, Iceland, Italy, Israel, Latvia, Lithuania, Luxembourg, Malta, Netherlands, Norway, New Zealand, Poland, Portugal, Romania, Slovakia, Slovenia, Spain, Switzerland, Sweden, Islamic Republic of Iran and United Kingdom;",
            },
            {
                text: "11.2. Persons associated with a country are defined as persons:",
                child: [
                    "11.2.1. having citizenship, a residence visa, or other similar document from a country in which the SmartOption does not do business;",
                    "11.2.2. residing in a country in which the SmartOption does not do business, being a resident of a country in which the SmartOption does not do business or having mailing addresses or residence addresses in a country in which the SmartOption does not do business;",
                    "11.2.3. whose places of birth are countries in which the SmartOption does not do business;",
                    "11.2.4. having a telephone number with the country code of a country in which the SmartOption does not do business; having an IP address associated with a country in which the SmartOption does not do business;",
                    "11.2.5. having another connection with a country in which the SmartOption does not do business as defined by the SmartOption at its discretion.",
                ]
            },
            {
                text: "11.3. If it is found that the SmartOption is providing services to persons associated with the countries in which the SmartOption does not do business, the SmartOption may apply the consequences listed in clause 8.5 of the Agreement.",
            },
        ]
    }
};
// 隐私协议
var userAgreement = {
    INSA: {
        title: 'Kebijakan Privasi',
        privacyArr0TextArr: [
            {
                text: 'Kebijakan ini mencakup bagaimana SmartOption ("kami") melindungi dan menggunakan informasi yang dikumpulkan tentang pengguna.SmartOption melindungi privasi semua pengguna yang mengunjungi atau menggunakan Layanan SmartOption(secara kolektif atau individual, "Anda" atau "pengguna").'
            },
            {
                text: 'SmartOption membangun aplikasi SmartOption sebagai aplikasi Gratis. LAYANAN ini disediakan oleh SmartOption tanpa biaya dan dimaksudkan untuk digunakan apa adanya.'
            },
            {
                text: 'Halaman ini digunakan untuk memberi tahu pengunjung mengenai kebijakan kami mengenai pengumpulan, penggunaan, dan pengungkapan Informasi Pribadi jika ada yang memutuskan untuk menggunakan Layanan kami.'
            },
            {
                text: 'Jika Anda memilih untuk menggunakan Layanan kami, maka Anda menyetujui pengumpulan dan penggunaan informasi sehubungan dengan kebijakan ini. Informasi Pribadi yang kami kumpulkan digunakan untuk menyediakan dan meningkatkan Layanan. Kami tidak akan menggunakan atau membagikan informasi Anda dengan siapa pun kecuali sebagaimana dijelaskan dalam Kebijakan Privasi ini.'
            },
            {
                text: 'Istilah yang digunakan dalam Kebijakan Privasi ini memiliki arti yang sama seperti dalam <a class= "termsOfService" href = "./termsOfService.html?notApp=true" > Terms & Conditions</a > kami, yang dapat diakses di SmartOption kecuali ditentukan lain dalam Kebijakan Privasi ini.'
            }
        ],
        privacyArr1Titel: 'I. Pengumpulan dan Penggunaan Informasi',
        privacyArr1TextArr: [
            {
                text: '1. Untuk pengalaman yang lebih baik, saat menggunakan Layanan kami, kami mungkin meminta Anda untuk memberikan kami informasi pribadi tertentu yang dapat diidentifikasi.Informasi yang kami minta akan disimpan di perangkat Anda dan tidak dikumpulkan oleh kami dengan cara apa pun.'
            },
            {
                text: '2. Aplikasi ini menggunakan layanan pihak ketiga yang dapat mengumpulkan informasi yang digunakan untuk mengidentifikasi Anda.'
            },
            {
                text: '3. Tautan ke kebijakan privasi penyedia layanan pihak ketiga yang digunakan oleh aplikasi'
            },
            {
                indent: true,
                text: '<a href="https://policies.google.com/privacy">Google Play Services</a>'
            },
            {
                indent: true,
                text: '<a href="https://www.facebook.com/about/privacy/update/printable">Facebook</a>'
            },
            {
                text: '4. Kami dapat mengumpulkan, menggunakan, menyimpan, dan mentransfer berbagai jenis data tentang Anda yang telah kami kelompokkan sebagai berikut: '
            },
            {
                indent: true,
                text: '· Data Kontak mencakup alamat penagihan, alamat email, dan nomor telepon.'
            },
            {
                indent: true,
                text: '· Data Keuangan mencakup rekening bank, detail kartu pembayaran.'
            },
            {
                indent: true,
                text: '· Data Transaksi mencakup detail tentang transaksi yang Anda lakukan, detail tentang pembayaran, penarikan, riwayat perdagangan, laba, saldo, metode jumlah setoran dan penarikan, dan detail lainnya terkait dengan layanan yang telah Anda gunakan melalui Situs Web kami.'
            },
            {
                indent: true,
                text: '· Data Teknis mencakup alamat protokol internet (IP), data login Anda, jenis dan versi browser, setelan dan lokasi zona waktu, jenis dan versi plugin browser, sistem operasi dan platform, serta teknologi lain pada perangkat yang Anda gunakan untuk mengakses Situs Web dan penggunaan cookie yang disimpan di perangkat Anda.'
            },
            {
                indent: true,
                text: '· Data Profil mencakup detail akun Klien Anda, nama pengguna dan kata sandi, transaksi yang Anda lakukan, minat, preferensi, umpan balik, dan/atau informasi yang diterima melalui interaksi Anda dengan kami selama memberikan layanan dan tanggapan survei kami.'
            },
            {
                indent: true,
                text: '· Data Penggunaan mencakup informasi tentang cara Anda menggunakan Situs Web, produk dan layanan, tanggal pendaftaran, kategori akun, jumlah keluhan, jumlah permintaan yang diajukan, dan riwayat IP.'
            },
            {
                indent: true,
                text: '· Data Pemasaran dan Komunikasi mencakup preferensi Anda dalam menerima pemasaran dari kami dan preferensi komunikasi Anda.'
            },
            {
                indent: true,
                text: '· Data Perbankan mencakup detail jumlah dompet pembayaran dan/atau informasi kartu bank Anda, termasuk informasi bank penerbit, nomor kartu, nama pemegang kartu.'
            },
            {
                indent: true,
                text: '· Data di KYC (Kenali pelanggan Anda) termasuk informasi dokumen identitas, Kartu Identitas.'
            },
            {
                indent: true,
                text: '· Data Lokasi mencakup detail lokasi Anda yang sebenarnya saat berinteraksi dengan Situs kami (misalnya, serangkaian parameter yang menentukan setelan regional antarmuka Anda, yaitu negara tempat tinggal, zona waktu, dan bahasa antarmuka).'
            },
            {
                text: '5. Kami akan mengumpulkan informasi pribadi termasuk nama, email, dan gambar profil tentang Anda dalam keadaan berikut: '
            },
            {
                indent: true,
                text: '· saat Anda membuka atau masuk ke akun SmartOption dengan menggunakan akun penyedia layanan pihak ketiga. Informasi ini dapat ditinjau kapan saja dengan masuk ke akun Anda.'
            },
            {
                indent: true,
                text: '· Saat Anda berkomunikasi dengan kami: mungkin perlu untuk memberikan informasi tambahan tentang Anda dan transaksi Anda dalam komunikasi.Selain itu, kami dapat menghubungi Anda untuk meminta bukti identitas pribadi, seperti gambar dokumen hukum, untuk memastikan bahwa transaksi Anda sah atau sejauh diperlukan untuk mematuhi kewajiban hukum kami sebagai lembaga keuangan.'
            },
        ],
        privacyArr2Titel: 'II. Bagaimana kami memproses data informasi Anda:',
        privacyArr2TextArr: [
            {
                text: '1. Pemrosesan data Anda dilakukan oleh SmartOption dengan mengikuti prinsip keabsahan, keadilan, transparansi, dan selalu mengikuti tujuan pemrosesan data yang dimaksudkan, prinsip minimalisasi data, akurasi, penyimpanan data terbatas, integritas data, kerahasiaan, dan akuntabilitas.'
            },
            {
                text: '2. Data Agregat mencakup data statistik atau demografis untuk tujuan apa pun. Data tersebut dapat berasal dari data Anda tetapi tidak dapat dianggap sebagai data pribadi dalam undang- undang karena tidak akan secara langsung atau tidak langsung mengungkapkan identitas Anda.Contoh Data Gabungan tersebut dapat berupa bahwa kami menggabungkan Data Penggunaan Anda untuk menghitung persentase pengguna yang mengakses fitur situs web dan / atau layanan / preferensi produk tertentu.',
            },
            {
                text: 'Meskipun demikian, jika SmartOption menggabungkan Data Agregat dengan data sedemikian rupa sehingga hasil akhirnya dapat dengan cara apa pun mengidentifikasi subjek data, SmartOption akan memperlakukan data gabungan tersebut sebagai data yang akan diperlakukan sesuai dengan ketentuan yang tercantum di sini.',
            },
            {
                text: '3. Kami memproses semua data berdasarkan dasar hukum berikut:',
            },
            {
                indent: true,
                text: '· Untuk memenuhi kewajiban hukum kami;',
            },
            {
                indent: true,
                text: '· Untuk kinerja kewajiban kontraktual kami terhadap pengguna;',
            },
            {
                indent: true,
                text: '· Untuk tujuan menjaga kepentingan sah kami dan kepentingan Anda serta hak-hak dasar tidak mengesampingkan kepentingan - kepentingan tersebut; dan/atau',
            },
            {
                indent: true,
                text: '· Atas dasar persetujuan Anda.',
            },
        ],
        privacyArr3Titel: 'III. Tujuan pengumpulan data informasi Anda:',
        privacyArr3TextArr: [
            {
                indent: true,
                text: '· Kami hanya menggunakan data Anda untuk alasan berikut:'
            },
            {
                indent: true,
                text: '· Untuk menawarkan layanan SmartOption: Kami akan menggunakan informasi pribadi Anda sebagaimana diperlukan untuk menyediakan layanan SmartOption, termasuk sejauh yang diperlukan untuk memproses transaksi yang Anda lakukan, memverifikasi identitas Anda, mengautentikasi akses Anda ke akun SmartOption dan berkomunikasi dengan Anda tentang layanan.'
            },
            {
                indent: true,
                text: '· Untuk mengelola risiko dan melindungi Anda, situs web dan layanan SmartOption: Pencegahan penipuan adalah bagian penting dalam menyediakan layanan pembayaran, dan kami menggunakan informasi Anda untuk membantu kami mendeteksi dan mencegah penipuan.'
            },
            {
                indent: true,
                text: '· Untuk mematuhi kewajiban hukum kami, termasuk semua undang-undang anti pencucian uang yang berlaku, undang-undang anti pendanaan teroris, peraturan layanan keuangan, dan kewajiban kontraktual kami kepada mitra pihak ketiga yang menyediakan atau membantu menyediakan metode pembayaran apa pun yang digunakan dalam layanan kami.'
            },
            {
                indent: true,
                text: '· Untuk kebutuhan bisnis sah kami lainnya: Kami mungkin perlu menggunakan informasi pribadi Anda untuk menegakkan kontrak dan persyaratan layanan kami dan untuk memantau aktivitas di situs.'
            },
        ],
        privacyArr4Titel: 'IV. Pilihan Anda tentang informasi Anda',
        privacyArr4TextArr: [
            {
                text: 'Anda berhak untuk'
            },
            {
                indent: true,
                text: '· Meminta kami mengubah informasi pribadi apa pun yang menjadi milik Anda.'
            },
            {
                indent: true,
                text: '· Meminta kami menghapus informasi pribadi apa pun tentang Anda.'
            },
            {
                indent: true,
                text: '· Untuk meminta perubahan pada informasi pribadi Anda atau menghapus informasi pribadi Anda, hubungi kami.'
            },
        ],
        privacyArr5Titel: 'V. Informasi Pengguna SmartOption',
        privacyArr5TextArr: [
            {
                text: 'SmartOption tidak akan, kapanpun, meminta informasi kartu kredit lengkap, ID akun dan kata sandi atau nomor identifikasi nasional dalam email yang tidak aman atau tidak diminta atau melalui komunikasi telepon.Jangan pernah mengungkapkan kata sandi Anda untuk alasan apa pun.'
            },
        ],
        privacyArr6Titel: 'VI. Data Log',
        privacyArr6TextArr: [
            {
                text: 'Kami ingin memberi tahu Anda bahwa setiap kali Anda menggunakan Layanan kami, jika terjadi kesalahan dalam aplikasi, kami mengumpulkan data dan informasi(melalui produk pihak ketiga) di ponsel Anda yang disebut Data Log.Data Log ini dapat mencakup informasi seperti alamat Protokol Internet(“IP”) perangkat Anda, nama perangkat, versi sistem operasi, konfigurasi aplikasi saat menggunakan Layanan kami, waktu dan tanggal penggunaan Layanan oleh Anda, dan statistik lainnya.'
            },
        ],
        privacyArr7Titel: 'VII. Cookie',
        privacyArr7TextArr: [
            {
                text: '1. Cookie adalah file dengan sejumlah kecil data yang biasanya digunakan sebagai pengidentifikasi unik anonim. Ini dikirim ke browser Anda dari situs web yang Anda kunjungi dan disimpan di memori internal perangkat Anda.'
            },
            {
                text: '2. Layanan ini tidak menggunakan “cookie” ini secara eksplisit. Namun, aplikasi dapat menggunakan kode dan perpustakaan pihak ketiga yang menggunakan "cookies" untuk mengumpulkan informasi dan meningkatkan layanan mereka. Anda memiliki Pilihan untuk menerima atau menolak cookie ini dan mengetahui kapan cookie dikirim ke perangkat Anda. Jika Anda memilih untuk menolak cookie kami, Anda mungkin tidak dapat menggunakan beberapa bagian dari Layanan ini.'
            },
        ],
        privacyArr8Titel: 'VIII. Penyedia layanan',
        privacyArr8TextArr: [
            {
                text: '1. Kami dapat mempekerjakan perusahaan dan individu pihak ketiga karena alasan berikut:'
            },
            {
                indent: true,
                text: '· Untuk memfasilitasi layanan kami;'
            },
            {
                indent: true,
                text: '· Untuk menyediakan layanan atas nama kami;'
            },
            {
                indent: true,
                text: '· Untuk melakukan layanan terkait layanan; atau'
            },
            {
                indent: true,
                text: '·Untuk membantu kami dalam menganalisis bagaimana layanan kami digunakan.'
            },
            {
                text: '2. Kami ingin memberi tahu pengguna Layanan ini bahwa pihak ketiga ini memiliki akses ke Informasi Pribadi Anda. Alasannya adalah untuk melakukan tugas yang diberikan kepada mereka atas nama kita.Namun, mereka berkewajiban untuk tidak mengungkapkan atau menggunakan informasi tersebut untuk tujuan lain.'
            },
        ],
        privacyArr9Titel: 'IX. Keamanan',
        privacyArr9TextArr: [
            {
                text: 'Kami menghargai kepercayaan Anda dalam memberikan Informasi Pribadi Anda kepada kami, oleh karena itu kami berusaha untuk menggunakan cara yang dapat diterima secara komersial untuk melindunginya.Tetapi ingat bahwa tidak ada metode transmisi melalui internet, atau metode penyimpanan elektronik yang 100% aman dan andal, dan kami tidak dapat menjamin keamanan mutlaknya.'
            },
        ],
        privacyArr10Titel: 'X. Tautan ke Situs Lain',
        privacyArr10TextArr: [
            {
                text: 'Layanan ini mungkin berisi tautan ke situs lain. Jika Anda mengklik tautan pihak ketiga, Anda akan diarahkan ke situs itu. Perhatikan bahwa situs eksternal ini tidak Dioperasikan oleh kami. Oleh karena itu, kami sangat menyarankan Anda untuk meninjau Kebijakan Privasi situs web ini. Kami tidak memiliki kendali atas dan tidak bertanggung jawab atas konten, kebijakan privasi, atau praktik situs atau layanan pihak ketiga mana pun.'
            },
        ],
        privacyArr11Titel: 'XI. Privasi Anak-anak',
        privacyArr11TextArr: [
            {
                text: 'Berusia dewasa sesuai dengan yurisdiksi tempat dia tinggal atau bertempat tinggal, memiliki kompetensi hukum dan waras.'
            },
        ],
        privacyArr12Titel: 'XII. Perubahan pada Kebijakan Privasi Ini.',
        privacyArr12TextArr: [
            {
                text: 'Kami dapat memperbarui Kebijakan Privasi kami dari waktu ke waktu. Oleh karena itu, Anda disarankan untuk meninjau halaman ini secara berkala untuk setiap perubahan. Kami akan memberi tahu Anda tentang perubahan apa pun dengan memposting Kebijakan Privasi baru di halaman ini. Perubahan ini efektif segera setelah diposting di halaman ini.'
            },
        ],
        privacyArr13Titel: 'XIII. Hubungi kami',
        privacyArr13TextArr: [
            {
                text: 'Jika Anda memiliki pertanyaan atau saran tentang Kebijakan Privasi kami, jangan ragu untuk menghubungi kami di sssmartoption@gmail.com.'
            },
        ],
    },
    EN: {
        title: 'Privacy Policy',
        privacyArr0TextArr: [
            {
                text: 'This policy covers how SmartOption ("we") protects and uses the information collected about the user. SmartOption protects the privacy of all users who visit or use the SmartOption Service (collectively or individually, "you" or "users").',
            },
            {
                text: 'SmartOption built the SmartOption app as a Free app. This SERVICE is provided by SmartOption at no cost and is intended for use as is.',
            },
            {
                text: 'This page is used to inform visitors regarding our policies with the collection, use, and disclosure of Personal Information if anyone decided to use our Service.',
            },
            {
                text: 'If you choose to use our Service, then you agree to the collection and use of information in relation to this policy. The Personal Information that we collect is used for providing and improving the Service. We will not use or share your information with anyone except as described in this Privacy Policy.',
            },
            {
                text: 'The terms used in this Privacy Policy have the same meanings as in our Terms & Conditions, which is accessible at SmartOption unless otherwise defined in this Privacy Policy.',
            }
        ],
        privacyArr1Titel: 'I. Information Collection and Use',
        privacyArr1TextArr: [
            {
                text: '1. For a better experience, while using our Service, we may require you to provide us with certain personally identifiable information. The information that we request will be retained on your device and is not collected by us in any way.',
            },
            {
                text: '2. The app does use third party services that may collect information used to identify you.',
            },
            {
                text: '3. Link to privacy policy of third-party service providers used by the app',
            },
            {
                indent: true,
                text: '<a href="https://policies.google.com/privacy">Google Play Services</a>'
            },
            {
                indent: true,
                text: '<a href="https://www.facebook.com/about/privacy/update/printable">Facebook</a>'
            },
            {
                text: '4. We may collect, use, store and transfer different kinds of data about you which we have grouped together as follows:'
            },
            {
                indent: true,
                text: '· Identity Data includes first name, last name, and patronymic (if available), date of birth gender, ID and copy of photo.'
            },
            {
                indent: true,
                text: '· Contact Data includes billing address, email address and telephone numbers.'
            },
            {
                indent: true,
                text: '· Financial Data includes bank account, payment card details.'
            },
            {
                indent: true,
                text: '· Transaction Data includes details about the transactions performed by you, details about payments, withdrawals, trading history, profit, balance, deposited and withdrawal amount methods, and any other details in relation to the services you have made use of through our Website.'
            },
            {
                indent: true,
                text: '· Technical Data includes internet protocol (IP) address, your login data, browser type and version, time zone setting and location, browser plug-in type and versions, operating system and platform, and other technologies on the devices you use to access the Website and use of cookies stored on your device.'
            },
            {
                indent: true,
                text: "· Profile Data includes your Client's account details, username and password, transactions made by you, your interests, preferences, feedback and/or information received through your interaction with us within the course of providing our services and survey responses."
            },
            {
                indent: true,
                text: '· Usage Data includes information about how you use the Website, products and services, registration date, account category, number of complaints, number of requests filed and IP history.'
            },
            {
                indent: true,
                text: '· Marketing and Communication Data includes your preferences in receiving marketing from us and your communication preferences.'
            },
            {
                indent: true,
                text: '· Banking Data includes details of the number of payment wallets and/or information of your bank card, including information of the issuing bank, card number, cardholder name.'
            },
            {
                indent: true,
                text: '· Data in KYC (Know your customer) includes identity document information, Identity Card.'
            },
            {
                indent: true,
                text: '· Location Data includes details on your actual location when interacting with our Website (for example, a set of parameters that determine regional settings of your interface, namely residency country, time zone, and the interface language).'
            },
            {
                text: '5. We will collect the personal information including your name, email, and profile picture about you in the following circumstances:'
            },
            {
                indent: true,
                text: "· when you open or log into a SmartOption account by using a third-party service providers' account. This information can be reviewed at any time by logging into your account."
            },
            {
                indent: true,
                text: '· When you communicate with us: it may be necessary to provide additional information about you and your transaction in the communication. In addition, we may contact you to request proof of personal identity, such as an image of a legal document, to ensure that your transaction is valid or insofar as it is necessary to comply with our legal obligations as a financial institution.'
            },
        ],
        privacyArr2Titel: 'II. How we process your information data:',
        privacyArr2TextArr: [
            {
                text: '1. Processing of your data is carried out by SmartOption following the principles of lawfulness, fairness, transparency, and always adhering to the intended purpose of data processing, the principle of data minimization, accuracy, limited data storage, data integrity, confidentiality and accountability.',
            },
            {
                text: '2. Aggregated Data includes statistical or demographic data for any purpose. Such data can be derived from your data but may not be considered personal data in law as it will not directly or indirectly reveal your identity. An example of such Aggregated Data could be that we aggregate your Usage Data to calculate the percentage of users accessing a specific website feature and/or services/product preference.',
            },
            {
                text: 'Notwithstanding the above, if SmartOption combines Aggregate Data with data in a way that the end result can in any way identify the data subject, SmartOption shall treat such combined data as data which will be treated as per the provisions herein contained.',
            },
            {
                text: '3. We process all data based on the following legal basis:',
            },
            {
                indent: true,
                text: '· For compliance with our legal obligations;',
            },
            {
                indent: true,
                text: '· For the performance of our contractual obligations towards the users;',
            },
            {
                indent: true,
                text: '· For the purposes of safeguarding our legitimate interests and your interests and fundamental rights do not override those interests; and/or',
            },
            {
                indent: true,
                text: '· On the basis of your consent.',
            },
        ],
        privacyArr3Titel: 'III. Purpose of collecting your information data:',
        privacyArr3TextArr: [
            {
                text: 'We only use your data for the following reasons:',
            },
            {
                indent: true,
                text: '· To offer the SmartOption service: We will use your personal information as necessary to provide the SmartOption service, including to the extent necessary to process the transaction you initiated, verify your identity, authenticate your access to a SmartOption account and communicate with you about the service.',
            },
            {
                indent: true,
                text: '· To manage risks and protect you, the SmartOption website and service: Fraud prevention is a critical part of providing payment services, and we use your information to help us detect and prevent fraud.',
            },
            {
                indent: true,
                text: '· To comply with our legal obligations, including all applicable anti-money laundering laws, anti-terrorist financing laws, financial services regulations, and our contractual obligations to third party partners that provide or help provide any payment method used in our service.',
            },
            {
                indent: true,
                text: '· For our other legitimate business needs: We may need to use your personal information to enforce our contracts and terms of service and to monitor activity on the site.',
            },
        ],
        privacyArr4Titel: 'IV. Your choices about your information',
        privacyArr4TextArr: [
            {
                text: 'You have the right to:',
            },
            {
                indent: true,
                text: '· Request that we change any personal information that belongs to you.',
            },
            {
                indent: true,
                text: '· Request that we delete any personal information concerning you.',
            },
            {
                indent: true,
                text: '· To request a change to your personal information or delete your personal information, contact us.',
            },
        ],
        privacyArr5Titel: 'V. SmartOption User Information',
        privacyArr5TextArr: [
            {
                text: 'SmartOption will not, at any time, request full credit card information, account ID and password or national identification numbers in an unsecured or unsolicited email or by telephone communication. Never disclose your passwords for any reason.',
            }
        ],
        privacyArr6Titel: 'VI. Log Data',
        privacyArr6TextArr: [
            {
                text: 'We want to inform you that whenever you use our Service, in a case of an error in the app we collect data and information (through third party products) on your phone called Log Data. This Log Data may include information such as your device Internet Protocol (“IP”) address, device name, Operating system version, the configuration of the app when utilizing my Service, the time and date of your use of the Service, and other statistics.',
            }
        ],
        privacyArr7Titel: 'VII. Cookies',
        privacyArr7TextArr: [
            {
                text: "1. Cookies are files with a small amount of data that are commonly used as anonymous unique identifiers. These are sent to your browser from the websites that you visit and are stored on your device's internal memory.",
            },
            {
                text: '2. This Service does not use these “cookies” explicitly. However, the app may use third party code and libraries that use “cookies” to collect information and improve their services. You have the Option to either accept or refuse these cookies and know when a cookie is being sent to your device. If you choose to refuse our cookies, you may not be able to use some portions of this Service.',
            },
        ],
        privacyArr8Titel: 'VIII. Service Providers',
        privacyArr8TextArr: [
            {
                text: '1. We may employ third-party companies and individuals due to the following reasons:',
            },
            {
                indent: true,
                text: '· To facilitate our Service;',
            },
            {
                indent: true,
                text: '· To provide the Service on our behalf;',
            },
            {
                indent: true,
                text: '· To perform Service-related services; or',
            },
            {
                indent: true,
                text: '· To assist us in analyzing how our Service is used.',
            },
            {
                text: '2. We want to inform users of this Service that these third parties have access to your Personal Information. The reason is to perform the tasks assigned to them on our behalf. However, they are obligated not to disclose or use the information for any other purpose.'
            }
        ],
        privacyArr9Titel: 'IX. Security',
        privacyArr9TextArr: [
            {
                text: 'We value your trust in providing us your Personal Information, thus we are striving to use commercially acceptable means of protecting it. But remember that no method of transmission over the internet, or method of electronic storage is 100% secure and reliable, and we cannot guarantee its absolute security.',
            }
        ],
        privacyArr10Titel: 'X. Links to Other Sites',
        privacyArr10TextArr: [
            {
                text: 'This Service may contain links to other sites. If you click on a third-party link, you will be directed to that site. Note that these external sites are not Operated by me. Therefore, we strongly advise you to review the Privacy Policy of these websites. We have no control over and assume no responsibility for the content, privacy policies, or practices of any third-party sites or services.',
            }
        ],
        privacyArr11Titel: "XI. Children's Privacy",
        privacyArr11TextArr: [
            {
                text: 'Being of age of maturity in accordance with the jurisdiction he/she resides in or is a resident of, is of legal competence and of sound mind.XII. Changes to This Privacy Policy.',
            }
        ],
        privacyArr12Titel: 'XII. Changes to This Privacy Policy',
        privacyArr12TextArr: [
            {
                text: 'We may update our Privacy Policy from time to time. Thus, you are advised to review this page periodically for any changes. We will notify you of any changes by posting the new Privacy Policy on this page. These changes are effective immediately after they are posted on this page.',
            }
        ],
        privacyArr13Titel: 'XIII. Contact Us',
        privacyArr13TextArr: [
            {
                text: 'If you have any questions or suggestions about my Privacy Policy, do not hesitate to contact me at sssmartoption@gmail.com.',
            }
        ],
    },
    MX: {
        title: 'Privacy Policy',
        privacyArr0TextArr: [
            {
                text: 'This policy covers how SmartOption ("we") protects and uses the information collected about the user. SmartOption protects the privacy of all users who visit or use the SmartOption Service (collectively or individually, "you" or "users").',
            },
            {
                text: 'SmartOption built the SmartOption app as a Free app. This SERVICE is provided by SmartOption at no cost and is intended for use as is.',
            },
            {
                text: 'This page is used to inform visitors regarding our policies with the collection, use, and disclosure of Personal Information if anyone decided to use our Service.',
            },
            {
                text: 'If you choose to use our Service, then you agree to the collection and use of information in relation to this policy. The Personal Information that we collect is used for providing and improving the Service. We will not use or share your information with anyone except as described in this Privacy Policy.',
            },
            {
                text: 'The terms used in this Privacy Policy have the same meanings as in our Terms & Conditions, which is accessible at SmartOption unless otherwise defined in this Privacy Policy.',
            }
        ],
        privacyArr1Titel: 'I. Information Collection and Use',
        privacyArr1TextArr: [
            {
                text: '1. For a better experience, while using our Service, we may require you to provide us with certain personally identifiable information. The information that we request will be retained on your device and is not collected by us in any way.',
            },
            {
                text: '2. The app does use third party services that may collect information used to identify you.',
            },
            {
                text: '3. Link to privacy policy of third-party service providers used by the app',
            },
            {
                indent: true,
                text: '<a href="https://policies.google.com/privacy">Google Play Services</a>'
            },
            {
                indent: true,
                text: '<a href="https://www.facebook.com/about/privacy/update/printable">Facebook</a>'
            },
            {
                text: '4. We may collect, use, store and transfer different kinds of data about you which we have grouped together as follows:'
            },
            {
                indent: true,
                text: '· Identity Data includes first name, last name, and patronymic (if available), date of birth gender, ID and copy of photo.'
            },
            {
                indent: true,
                text: '· Contact Data includes billing address, email address and telephone numbers.'
            },
            {
                indent: true,
                text: '· Financial Data includes bank account, payment card details.'
            },
            {
                indent: true,
                text: '· Transaction Data includes details about the transactions performed by you, details about payments, withdrawals, trading history, profit, balance, deposited and withdrawal amount methods, and any other details in relation to the services you have made use of through our Website.'
            },
            {
                indent: true,
                text: '· Technical Data includes internet protocol (IP) address, your login data, browser type and version, time zone setting and location, browser plug-in type and versions, operating system and platform, and other technologies on the devices you use to access the Website and use of cookies stored on your device.'
            },
            {
                indent: true,
                text: "· Profile Data includes your Client's account details, username and password, transactions made by you, your interests, preferences, feedback and/or information received through your interaction with us within the course of providing our services and survey responses."
            },
            {
                indent: true,
                text: '· Usage Data includes information about how you use the Website, products and services, registration date, account category, number of complaints, number of requests filed and IP history.'
            },
            {
                indent: true,
                text: '· Marketing and Communication Data includes your preferences in receiving marketing from us and your communication preferences.'
            },
            {
                indent: true,
                text: '· Banking Data includes details of the number of payment wallets and/or information of your bank card, including information of the issuing bank, card number, cardholder name.'
            },
            {
                indent: true,
                text: '· Data in KYC (Know your customer) includes identity document information, Identity Card.'
            },
            {
                indent: true,
                text: '· Location Data includes details on your actual location when interacting with our Website (for example, a set of parameters that determine regional settings of your interface, namely residency country, time zone, and the interface language).'
            },
            {
                text: '5. We will collect the personal information including your name, email, and profile picture about you in the following circumstances:'
            },
            {
                indent: true,
                text: "· when you open or log into a SmartOption account by using a third-party service providers' account. This information can be reviewed at any time by logging into your account."
            },
            {
                indent: true,
                text: '· When you communicate with us: it may be necessary to provide additional information about you and your transaction in the communication. In addition, we may contact you to request proof of personal identity, such as an image of a legal document, to ensure that your transaction is valid or insofar as it is necessary to comply with our legal obligations as a financial institution.'
            },
        ],
        privacyArr2Titel: 'II. How we process your information data:',
        privacyArr2TextArr: [
            {
                text: '1. Processing of your data is carried out by SmartOption following the principles of lawfulness, fairness, transparency, and always adhering to the intended purpose of data processing, the principle of data minimization, accuracy, limited data storage, data integrity, confidentiality and accountability.',
            },
            {
                text: '2. Aggregated Data includes statistical or demographic data for any purpose. Such data can be derived from your data but may not be considered personal data in law as it will not directly or indirectly reveal your identity. An example of such Aggregated Data could be that we aggregate your Usage Data to calculate the percentage of users accessing a specific website feature and/or services/product preference.',
            },
            {
                text: 'Notwithstanding the above, if SmartOption combines Aggregate Data with data in a way that the end result can in any way identify the data subject, SmartOption shall treat such combined data as data which will be treated as per the provisions herein contained.',
            },
            {
                text: '3. We process all data based on the following legal basis:',
            },
            {
                indent: true,
                text: '· For compliance with our legal obligations;',
            },
            {
                indent: true,
                text: '· For the performance of our contractual obligations towards the users;',
            },
            {
                indent: true,
                text: '· For the purposes of safeguarding our legitimate interests and your interests and fundamental rights do not override those interests; and/or',
            },
            {
                indent: true,
                text: '· On the basis of your consent.',
            },
        ],
        privacyArr3Titel: 'III. Purpose of collecting your information data:',
        privacyArr3TextArr: [
            {
                text: 'We only use your data for the following reasons:',
            },
            {
                indent: true,
                text: '· To offer the SmartOption service: We will use your personal information as necessary to provide the SmartOption service, including to the extent necessary to process the transaction you initiated, verify your identity, authenticate your access to a SmartOption account and communicate with you about the service.',
            },
            {
                indent: true,
                text: '· To manage risks and protect you, the SmartOption website and service: Fraud prevention is a critical part of providing payment services, and we use your information to help us detect and prevent fraud.',
            },
            {
                indent: true,
                text: '· To comply with our legal obligations, including all applicable anti-money laundering laws, anti-terrorist financing laws, financial services regulations, and our contractual obligations to third party partners that provide or help provide any payment method used in our service.',
            },
            {
                indent: true,
                text: '· For our other legitimate business needs: We may need to use your personal information to enforce our contracts and terms of service and to monitor activity on the site.',
            },
        ],
        privacyArr4Titel: 'IV. Your choices about your information',
        privacyArr4TextArr: [
            {
                text: 'You have the right to:',
            },
            {
                indent: true,
                text: '· Request that we change any personal information that belongs to you.',
            },
            {
                indent: true,
                text: '· Request that we delete any personal information concerning you.',
            },
            {
                indent: true,
                text: '· To request a change to your personal information or delete your personal information, contact us.',
            },
        ],
        privacyArr5Titel: 'V. SmartOption User Information',
        privacyArr5TextArr: [
            {
                text: 'SmartOption will not, at any time, request full credit card information, account ID and password or national identification numbers in an unsecured or unsolicited email or by telephone communication. Never disclose your passwords for any reason.',
            }
        ],
        privacyArr6Titel: 'VI. Log Data',
        privacyArr6TextArr: [
            {
                text: 'We want to inform you that whenever you use our Service, in a case of an error in the app we collect data and information (through third party products) on your phone called Log Data. This Log Data may include information such as your device Internet Protocol (“IP”) address, device name, Operating system version, the configuration of the app when utilizing my Service, the time and date of your use of the Service, and other statistics.',
            }
        ],
        privacyArr7Titel: 'VII. Cookies',
        privacyArr7TextArr: [
            {
                text: "1. Cookies are files with a small amount of data that are commonly used as anonymous unique identifiers. These are sent to your browser from the websites that you visit and are stored on your device's internal memory.",
            },
            {
                text: '2. This Service does not use these “cookies” explicitly. However, the app may use third party code and libraries that use “cookies” to collect information and improve their services. You have the Option to either accept or refuse these cookies and know when a cookie is being sent to your device. If you choose to refuse our cookies, you may not be able to use some portions of this Service.',
            },
        ],
        privacyArr8Titel: 'VIII. Service Providers',
        privacyArr8TextArr: [
            {
                text: '1. We may employ third-party companies and individuals due to the following reasons:',
            },
            {
                indent: true,
                text: '· To facilitate our Service;',
            },
            {
                indent: true,
                text: '· To provide the Service on our behalf;',
            },
            {
                indent: true,
                text: '· To perform Service-related services; or',
            },
            {
                indent: true,
                text: '· To assist us in analyzing how our Service is used.',
            },
            {
                text: '2. We want to inform users of this Service that these third parties have access to your Personal Information. The reason is to perform the tasks assigned to them on our behalf. However, they are obligated not to disclose or use the information for any other purpose.'
            }
        ],
        privacyArr9Titel: 'IX. Security',
        privacyArr9TextArr: [
            {
                text: 'We value your trust in providing us your Personal Information, thus we are striving to use commercially acceptable means of protecting it. But remember that no method of transmission over the internet, or method of electronic storage is 100% secure and reliable, and we cannot guarantee its absolute security.',
            }
        ],
        privacyArr10Titel: 'X. Links to Other Sites',
        privacyArr10TextArr: [
            {
                text: 'This Service may contain links to other sites. If you click on a third-party link, you will be directed to that site. Note that these external sites are not Operated by me. Therefore, we strongly advise you to review the Privacy Policy of these websites. We have no control over and assume no responsibility for the content, privacy policies, or practices of any third-party sites or services.',
            }
        ],
        privacyArr11Titel: "XI. Children's Privacy",
        privacyArr11TextArr: [
            {
                text: 'Being of age of maturity in accordance with the jurisdiction he/she resides in or is a resident of, is of legal competence and of sound mind.XII. Changes to This Privacy Policy.',
            }
        ],
        privacyArr12Titel: 'XII. Changes to This Privacy Policy',
        privacyArr12TextArr: [
            {
                text: 'We may update our Privacy Policy from time to time. Thus, you are advised to review this page periodically for any changes. We will notify you of any changes by posting the new Privacy Policy on this page. These changes are effective immediately after they are posted on this page.',
            }
        ],
        privacyArr13Titel: 'XIII. Contact Us',
        privacyArr13TextArr: [
            {
                text: 'If you have any questions or suggestions about my Privacy Policy, do not hesitate to contact me at sssmartoption@gmail.com.',
            }
        ],
    },
};
// 交易页面数据
var tradePage = {
    INSA: {
        realAccount: 'Nyata',
        bonus: 'Bonus',
        virAccount: 'Demo',
        buyDown: 'Beli turun',
        buyUp: 'Beli naik',
        deposit: 'Deposit',
        oneMinuteProfit: 'Dapatkan untung {rate}% setelah 1 menit',
        OneClickTransaction: 'Trading Sekali Klik',
        buyPointPrice: 'Beli indeks',
        buyAmounts: 'Saldo beli',
        tradingGain: 'Keuntungan',
        sureAmounts: 'Konfirmasi Saldo',
        beginTime: 'Waktu Mulai',
        endTime: 'Akhir waktu',
        tradingResult: 'Hasil Trading',
        startIndex: 'Harga Mulai',
        endTradingTime: 'Waktu Akhir',
        endPointPrice: 'Harga Akhir',
        determine: 'Konfirmasi',
        selectAssets: 'Pilih aset',
        assets: 'Asset',
        profitRate: 'Untung',
        all: 'Jumlah',
        low: 'Minimal',
        languageSet_tips: {
            title: 'Judul',
            refresh: 'Segarkan',
            confirm: 'Konfirmasi',
            Data_Exception: 'Dapatkan data abnormal, silakan coba lagi nanti.',
            Asset_DataException: 'Data aset abnormal sekarang, silakan pilih aset yang lain.',
            Insufficient_balance: 'Saldo Anda tidak cukup untuk membayar transaksi ini.',
            Insufficient_balance1: 'Saldo tidak cukup',
            Network_exception: 'Internet tidak normal, silakan cek internet saat ini.',
            Loading_Delay: 'Kelambatan pemuatan data, silakan coba lagi nanti.',
            Amount_maxtrade_new: 'Jumlah transaksi user baru tidak boleh lebih tinggi dari',
            Amount_maxtrade: 'Maksimum jumlah transaksi tunggal  ',
            Amount_mintrade: 'Maksimum jumlah transaksi tunggal ',
            trade_max_tips: 'Total transaksi anda hari ini melebihi ',
            Deposit: 'Deposit',
            buDown1: 'Earn ',
            buDown2: ' % profit after 1min',
            singleTransactionLimit: 'Jumlah transaksi maksimum santuan {currencySymbol} {money}',
        },
    },
    EN: {
        realAccount: 'Real',
        bonus: 'Bonus',
        virAccount: 'Demo',
        buyDown: 'LOWER',
        buyUp: 'HIGHER',
        deposit: 'Deposit',
        oneMinuteProfit: 'Earn {rate}% profit after 1min',
        OneClickTransaction: 'One Click',
        buyPointPrice: 'Buy index',
        buyAmounts: 'Amount',
        tradingGain: 'Profit',
        sureAmounts: 'Confirm',
        beginTime: 'Opening Time',
        endTime: 'End Time',
        tradingResult: 'Trading Result',
        startIndex: 'Opening Price',
        endTradingTime: 'Closing Time',
        endPointPrice: 'Closing Price',
        determine: 'Confirm',
        selectAssets: 'Choose asset',
        assets: 'Asset',
        profitRate: 'Profit',
        all: 'Total',
        low: 'Minimal',
        languageSet_tips: {
            title: 'Title',
            refresh: 'Refresh',
            confirm: 'Confirm',
            Data_Exception: 'Data error, please try again later.',
            Asset_DataException: 'The asset is abnormal, please select other asset.',
            Insufficient_balance: 'Balance is not enough to do the transaction.',
            Insufficient_balance1: 'Balance is not enough',
            Network_exception: 'The network is abnormal, please check the current network.',
            Loading_Delay: 'Data loading delay, please try again.',
            Amount_maxtrade_new: 'The first transaction amount cannot be higher than ',
            Amount_maxtrade: 'Maximum Trading Amount ',
            Amount_mintrade: 'Minimum Trading Amount ',
            trade_max_tips: 'You have exceeded the limit of transactions today ',
            Deposit: 'Deposit',
            buDown1: 'Dapatkan untung ',
            buDown2: ' % setelah 1 menit',
            singleTransactionLimit: 'Maximun trading amount a time: {currencySymbol} {money}',
        },
    },
    MX: {
        realAccount: 'Real',
        bonus: 'Bonus',
        virAccount: 'Demo',
        buyDown: 'LOWER',
        buyUp: 'HIGHER',
        deposit: 'Deposit',
        oneMinuteProfit: 'Earn {rate}% profit after 1min',
        OneClickTransaction: 'One Click',
        buyPointPrice: 'Buy index',
        buyAmounts: 'Amount',
        tradingGain: 'Profit',
        sureAmounts: 'Confirm',
        beginTime: 'Opening Time',
        endTime: 'End Time',
        tradingResult: 'Trading Result',
        startIndex: 'Opening Price',
        endTradingTime: 'Closing Time',
        endPointPrice: 'Closing Price',
        determine: 'Confirm',
        selectAssets: 'Choose asset',
        assets: 'Asset',
        profitRate: 'Profit',
        all: 'Total',
        low: 'Minimal',
        languageSet_tips: {
            title: 'Title',
            refresh: 'Refresh',
            confirm: 'Confirm',
            Data_Exception: 'Data error, please try again later.',
            Asset_DataException: 'The asset is abnormal, please select other asset.',
            Insufficient_balance: 'Balance is not enough to do the transaction.',
            Insufficient_balance1: 'Balance is not enough',
            Network_exception: 'The network is abnormal, please check the current network.',
            Loading_Delay: 'Data loading delay, please try again.',
            Amount_maxtrade_new: 'The first transaction amount cannot be higher than ',
            Amount_maxtrade: 'Maximum Trading Amount ',
            Amount_mintrade: 'Minimum Trading Amount ',
            trade_max_tips: 'You have exceeded the limit of transactions today ',
            Deposit: 'Deposit',
            buDown1: 'Dapatkan untung ',
            buDown2: ' % setelah 1 menit',
            singleTransactionLimit: 'Maximun trading amount a time: {currencySymbol} {money}',
        },
    },
    CH: {
        realAccount: 'Real',
        bonus: 'Bonus',
        virAccount: 'Demo',
        buyDown: 'LOWER',
        buyUp: 'HIGHER',
        deposit: 'Deposit',
        oneMinuteProfit: 'Earn {rate}% profit after 1min',
        OneClickTransaction: 'One Click',
        buyPointPrice: 'Buy index',
        buyAmounts: 'Amount',
        tradingGain: 'Profit',
        sureAmounts: 'Confirm',
        beginTime: 'Opening Time',
        endTime: 'End Time',
        tradingResult: 'Trading Result',
        startIndex: 'Opening Price',
        endTradingTime: 'Closing Time',
        endPointPrice: 'Closing Price',
        determine: 'Confirm',
        selectAssets: 'Choose asset',
        assets: 'Asset',
        profitRate: 'Profit',
        all: 'Total',
        low: 'Minimal',
        languageSet_tips: {
            title: 'Title',
            refresh: 'Refresh',
            confirm: 'Confirm',
            Data_Exception: 'Data error, please try again later.',
            Asset_DataException: 'The asset is abnormal, please select other asset.',
            Insufficient_balance: 'Balance is not enough to do the transaction.',
            Insufficient_balance1: 'Balance is not enough',
            Network_exception: 'The network is abnormal, please check the current network.',
            Loading_Delay: 'Data loading delay, please try again.',
            Amount_maxtrade_new: 'The first transaction amount cannot be higher than ',
            Amount_maxtrade: 'Maximum Trading Amount ',
            Amount_mintrade: 'Minimum Trading Amount ',
            trade_max_tips: 'You have exceeded the limit of transactions today ',
            Deposit: 'Deposit',
            buDown1: 'Dapatkan untung ',
            buDown2: ' % setelah 1 menit',
            singleTransactionLimit: 'Maximun trading amount a time: {currencySymbol} {money}',
        },
    },
};