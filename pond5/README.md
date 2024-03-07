# Ponderada 

A ponderada funciona da seguinte maneira: O HiveMQ funciona como o broker para que o subscriber e o publisher podem se conectar. O publisher publica as informação e o subscriber recebe a informação. A api pega a informação do subscriber e manda para o banco de dados. O metabase é conectado diretamente no banco podendo ser atualizado em tempo real.

# Execução do Projeto

É necesário clonar o projeto e entrar na ambiente usando os segunites comandos:

```
https://github.com/LucaSarhan/Modulo9.git
```

Vai ser necessário abrir 4 terminais dentro do projeto. Caso necessário use esse comando para entrar no local nos terminais novos:

```
cd ambiente_clonado/Modulo9/pond5
```

Em 3 terminais roda o seguinte comando para acessar o binário do go:

```
source .bashrc
```

No terminal que não foi rodado o comando anterior roda o seguinte comando para ter acesso ao metabase atraves de um container do docker e posteriormente segue os passos que o metabase proporciona. Caso voce tenha um container com o nome metabase vai ser necessário remover-lo:

```
cd database
```
```
sudo docker run -d -p 3000:3000 -v ~<caminho-para-o-repositorio>/Modulo9/pond5/database:/database --name metabase metabase/metabase
```

No passo do metabase para conectar o banco de dados vai ser necessário colocar esse caminho:

```
database/database.db
```

Em 1 dos terminais, rode os seguintes comandos:

```
cd api
```
```
go run .
```

No outro terminal, rode os seguintes comandos:

```
cd publisher
```
```
go run .
```

No terminal final, rode o seguinte comando

```
cd subscriber
```
```
go run .
```


# Funcionamento comprovado

Link da [ponderada](https://drive.google.com/file/d/14ureZIxT_-BuiznJSt5TFTx_YVTvPCg-/view?usp=sharing) rodando corretamente