# Ponderada 1
A ponderada é uma ferramenta de simulação de dispositivos IoT que emprega o protocolo MQTT por meio da biblioteca Eclipse Paho e da linguagem de programação Go.

Para simular o dispositivo IoT, optou-se pelo [Sensor de Radiação Solar](https://sigmasensors.com.br/produtos/sensor-de-radiacao-solar-sem-fio-hobonet-rxw-lib-900), que possui uma faixa de medição de 0 a 1280 W/m² e com uma taxa de registro de pelo menos 1 minuto.

# Execução do projeto
É necesário clonar o projeto e entrar na ambiente usando os segunites comandos:

```
https://github.com/LucaSarhan/Modulo9.git
cd ambiente_clonado/Modulo9/pond1
```

Vai ser necesário abrir 3 terminais dentro do peojeto. Caso necesário use esse comando para entrar no local nos terminais novos:

```
cd ambiente_clonado/Modulo9/pond1
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

# Funcionamento comprovado
Segue o [link do projeto](https://drive.google.com/file/d/19Xb20oNHNnaYbpXYjZWyF1yJ1bFiQNjs/view?usp=sharing) rodando corretamente
