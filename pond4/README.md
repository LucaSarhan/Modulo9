# Ponderada 4
A ponderada 4 é sobre o HiveMQ. Essa ponderada mudou o broker utilizado do mosquitto, um broker local, para o HiveMQ, um broker na web e adciona uma camada de autenticação TLS. Para connectar ao HiveMq foi necessário criar um cluster que permite publicar e subscrever ao mesmo. 

# Configuração do ambiente

Vai ser necessário inserir suas credenciais do cluster no codigo. Para isso vai ser necessário cria 2 arquivos .env um na pasta publisher e outro no subscriber. Em ambos vai ser necessário colocar suas credencias da seguinte forma:

```
BROKER_ADDR = "seu_address_do_cluster"
HIVE_USER = "<nome-de-usuario>"
HIVE_PSWD = "<senha-cadastrada>"
```

# Execução do Projeto

É necesário clonar o projeto e entrar na ambiente usando os segunites comandos:

```
https://github.com/LucaSarhan/Modulo9.git
cd ambiente_clonado/Modulo9/pond4
```

Vai ser necessário abrir 2 terminais dentro do peojeto. Caso necessário use esse comando para entrar no local nos terminais novos:

```
cd ambiente_clonado/Modulo9/pond4
```

Nos 2 terminais roda o seguinte comando para acessar o binário do go:

```
source .bashrc
```

Como a conexão com o cluster no HiveMQ é feito através do codigo não tem necessidade de executar nenhum comando a respeito.

Em 1 dos terminais roda os seguintes comandos:

```
cd publisher
```

```
go run .
```

No terminal outro terminal roda os seguintes comandos:

```
cd subscriber
```
```
go run .
```

# Funcionamento comprovado

Link da [ponderada](https://drive.google.com/file/d/1O4Y3A5_WiylkM4ldF5C3ctA4nuueZPXC/view?usp=sharing) rodando corretamente