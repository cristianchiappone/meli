# Meli Magneto

#### NOTA:
Se supuso el caso de encontrar una secuencia de cuatro letras iguales, ya sea de forma oblicua, horizontal o vertical, los elementos pertenecientes a ella no se podrán volver a utilizar para otra secuencia en la misma dirección.
Es decir que que si se encuentra una secuencia A A A A A, solo se utilizará (A A A A) A y no A (A A A A)

#### Desafios

- Nivel 1:
	El codigo fuente se encuentra /magneto/mutant.go
- Nivel 2 y 3:
	El codigo fuente se encuentra /magneto/api_rest.go

#### URL API REST
- Servicio en donde se puede detectar si un humano es mutante o no
	http://35.199.125.222:9212/mutant 
- Servicio con las estadísticas de las verificaciones de ADN
	http://35.199.125.222:9212/stats
##### Retorno de errores
```
	DB error: error al conectarse a la base de datos con la configuración proporcionada
	Sequence error:  Secuencia inválida
	Error decoding sequence: error al realizar Marshal de secuencia
	Query error: error al realizar la consulta en la base de datos
	Function error: error en la función isMutant
```
#### Requisitos para ejecutar localmente
- Instalar go: https://golang.org/dl/
- Instalar MySQL y correr su servicio en el puerto por defecto(3306)
- Descargar dependencias:
```
	go get github.com/go-sql-driver/mysql
	go get github.com/gorilla/mux
```
- Descargar el repositorio actual
- Crear un nuevo esquema llamado 'magneto' ejecutando:
```
	create database magneto;
```
- Importar el backup que se encuentra en /magneto/sequence.sql
El mismo contiene más de 100.000 casos de prueba clasificados
```
	mysql -u usuario -p magneto < sequence.sql
```
- Modificar la configuración de la API REST en /magneto/api_rest.go
```
	var host = "http://yourhost"
	var port = "yourport"
	var connectionString = "user:password@tcp(127.0.0.1:3306)/magneto?charset=utf8&parseTime=True&loc=Local"
```
- Ejecutar el siguiente comando dentro de /magneto/
```
	go run api_rest.go mutant.go stat.go sequence.go
```
#### Code coverage
- Ejecutar los siguientes comandos:
Realiza el code coverage y lo muestra en formato HTML
```
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
```
Devolverá:
```
	PASS
	coverage: 100.0% of statements
```
