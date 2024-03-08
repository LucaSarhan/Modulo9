# Prova
Link da atividade [prática](https://drive.google.com/file/d/1qPJOmd9c3yQubNYq1S9k1mk0CegN_Y6Y/view?usp=sharing) funcionando corretamente

# Execução da atividade
É necesário clonar o projeto e entrar na ambiente usando os segunites comandos:

```
git clone https://github.com/LucaSarhan/Modulo9.git
cd ambiente_clonado/Modulo9/pond1
```

Vai ser necesário abrir 3 terminais dentro do peojeto. Caso necesário use esse comando para entrar no local nos terminais novos:

```
cd ambiente_clonado/Modulo9/prova1
```

Nos 3 terminais roda o seguinte comando para acessar o binário do go:

```
source .bashrc
```

Para rodar o broker, utilize esse comando em um terminal:

```
mosquitto -c mosquitto.conf
```

Em 1 dos terminais roda os seguintes comandos:

```
cd publisher
go run .
```

No terminal restante dos terminais roda os seguintes comandos:

```
cd subscriber
go run .
```