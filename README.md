# go-docker-env-test
go를 docker환경에서 개발하도록 hot-reload적용
remote dubugging 가능토록 수정 예정

실행 :<br/>
docker build --tag golang-docker-pivonachi .<br/>
docker run -p 9999:9999 -v $(pwd):/app -d golang-docker-pivonachi 

확인 :
http://localhost:9999/fib?num=10
