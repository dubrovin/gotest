# gotest

Для запуска сервиса требуется поставить:
```
go get github.com/valyala/fasthttp
go get github.com/qiangxue/fasthttp-routing
```

Запустить двумя способами:
1)
```
go build
./gotest
```

2)
```
go run main.go // запустить с параметрами по умолчанию
go run main.go --addr :8090 --dir dir --ttl 60s
```


По умолчанию сервис запускается со следующими параметрами:
```
addr = flag.String("addr", ":8080", "http service address")
dir  = flag.String("dir", "storage", "storage dir path")
ttl  = flag.String("ttl", "10s", "time to live")
```

После запуска сервис будет доступен по указанному адресу в настройках
По умолчанию: 127.0.0.1:8500
Доступный список URLs:
```
"api/list" - получить весь список загруженных zip файлов
"api/upload" - загрузить zip файл
```

Примеры запросов:

```
curl http://127.0.0.1:8080/api/list

curl -i -X POST -H "Content-Type: multipart/form-data" -F "data=@zip.zip" http://127.0.0.1:8080/api/upload
```

В файле ab_log.txt представлен вывод apache benchmark:
```
ab -n 10000 -c 10 http://127.0.0.1:8080/api/list > ab_log.txt
```