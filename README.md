Для проверки rest-api онлайн библиотеки песен:
  
  1.Склонируйте репозиторий:
  ```bash
   git clone https://github.com/himka0-0/Online_song_library.git
   cd Online_song_library
  ```
  2.Зайдите в .env:
  
    Необходимо вставить url стороннего сервиса, согласно ТЗ пункту 2;
    Если будете использовать докер, то всё настроено;
    Если ручное тестирование, то DB_HOST=db надо изменить на DB_HOST=localhost;
    
  3.Запустите docker-compose:
  ```bash
   docker-compose up --build
  ```
  4.Зайдите в браузер и напишите:
  http://localhost:8080/swagger/index.html
  Вам откроется свагер, готовый для тестирования (всё описано в папке docs)
