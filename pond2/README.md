# Ponderada 2
A ponderada são alguns testes que foram criados para garantir a qualidade so simulador feito na ponderada anterior. Os testes visam garantir um certo nivel de qualidade para os seguintes aspectos: a taxa de registro, mensagens enviada pelo broker e mensagens lida pelo subscriber.

# Execução do projeto
É necesário clonar o projeto e entrar na ambiente usando os segunites comandos:

```
https://github.com/LucaSarhan/Modulo9.git
cd ambiente_clonado/Modulo9/pond2
```

Vai ser necesário abrir 3 terminais dentro do peojeto. Caso necesário use esse comando para entrar no local nos terminais novos:

```
cd ambiente_clonado/Modulo9/pond2
```

Nos 3 terminais roda o seguinte comando para acessar o binário do go:

```
source .bashrc
```

Para rodar o broker, utilize esse comando em um terminal:

```
mosquitto -c mosquitto.conf
```

Em 1 dos terminais não utilizado roda os seguintes comandos:

```
cd publisher_test
```

```
go test -v
```

No terminal restante dos terminais roda os seguintes comandos:

```
cd subscriber_test
```
```
go test -v
```

# Funcionamento comprovado
Segue o [link do projeto](https://drive.google.com/file/d/1KgFQBZkbOuU0Gg7NtmoiLwkBJBduChJw/view?usp=sharing) rodando corretamente
