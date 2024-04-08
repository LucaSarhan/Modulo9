# Ponderada 6 - kafka e confluente

Nessa ponderada decidir ir um pouco além e decidir integrar com um banco de dados. A ponderada com Kafka funciona da seguinte maneira: o HiveMQ atua como broker para permitir a conexão entre o publisher e o subscriber. O publisher publica as informações no tópico do Kafka, enquanto o subscriber recebe essas informações. A API captura os dados do subscriber e os envia para o banco de dados. O Metabase está conectado diretamente ao banco de dados, possibilitando atualizações em tempo real dos dados provenientes do Kafka.

# Execução do Projeto

É necesário clonar o projeto e entrar na ambiente usando os segunites comandos:

```
https://github.com/LucaSarhan/Modulo9.git
```

Caso queira ver o banco de dados funcionando vai ser necessário abrir 4 terminais dentro do projeto. Se não só 3 terminais vão ser necessários. Caso necessário use esse comando para entrar no local nos terminais novos:

```
cd ambiente_clonado/Modulo9/pond5
```

Nos 3 terminais roda o seguinte comando para acessar o binário do go:

```
source .bashrc
```

Parte para ver o banco de dados funcional: No terminal que não foi rodado o comando anterior roda o seguinte comando para ter acesso ao metabase atraves de um container do docker e posteriormente segue os passos que o metabase proporciona. Caso voce tenha um container com o nome metabase vai ser necessário remover-lo:

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

Retornando para o principal da atividade: Em 1 dos terminais, rode os seguintes comandos:

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

Link da [ponderada](https://drive.google.com/file/d/1TnTwtHW_m-m1Vt8CCXjYCM1wQ71cFwdO/view?usp=sharing) rodando corretamente