# Ponderada 7 - desafio de 1 bilhão de linhas

# Criando o arquivo com 1B delinhas

Primeiro instale os requirements:
```
python3 -m pip install -r requirements.txt
```

O script `createMeasurements.py` vai criar o arquivo com 1 bilhão de linhas:

```
python3 createMeasurements.py
```

Para rodar o arquivo com os calculos:
```
python3 calculateAveragePolars.py
```
# Explicação do codigo do arquivo createMeasurments.py

**Imports:** O script importa os módulos necessários, como argparse, time, numpy, polars e tqdm. Esses são usados para análise de argumentos, medição de tempo, geração de números aleatórios, manipulação de dados e visualização de progresso, respectivamente.

**Classe CreateMeasurement:**

Define uma classe CreateMeasurement responsável por gerar dados de temperatura.
A classe contém uma constante STATIONS, que é uma tupla de tuplas contendo nomes de estações e suas temperaturas médias.
Inicializa com um gerador de números aleatórios (rng do NumPy).

**Método generate_batch:**

Este método gera um lote de dados de temperatura.
Parâmetros:
std_dev: Desvio padrão para os dados de temperatura.
records: Número de registros (medições de temperatura) a serem gerados.
Amostra estações com reposição e as embaralha.
Em seguida, gera dados de temperatura desenhando números aleatórios de uma distribuição normal com médias correspondentes às temperaturas médias das estações selecionadas e o desvio padrão fornecido.
Retorna um polars.DataFrame com colunas names (nomes das estações) e temperature.

**Método generate_measurement_file:**

Este método gera um arquivo de medição com dados de temperatura.
Parâmetros:
file_name: Nome do arquivo de saída.
records: Número total de registros de temperatura a serem gerados.
sep: Separador usado no arquivo de saída.
std_dev: Desvio padrão para os dados de temperatura.
Divide o número total de registros em lotes (o tamanho padrão do lote é 10.000.000) e gera dados lote por lote.
Cada lote é gerado usando o método generate_batch.
Escreve os dados gerados no arquivo de saída no formato CSV.

**Seção de Análise de Argumentos:**

Define um tipo de argumento personalizado min_records para garantir que o número de registros seja um número inteiro positivo.
Configura um argparse.ArgumentParser para lidar com argumentos da linha de comando.
Define dois argumentos:
-o/--output: Especifica o nome do arquivo de saída.
-r/--records: Especifica o número total de registros a serem gerados.
Analisa os argumentos da linha de comando usando parser.parse_args().

# Explicação do codigo do arquivo calculateAveragePolars.py

**Leitura do Arquivo de Dados:**

Utiliza a função pl.scan_csv para ler o arquivo CSV "measurements.txt".
Define o separador como ;.
Indica que o arquivo não tem cabeçalho (has_header=False).
Usa uma função lambda para definir os nomes das colunas como "station_name" e "measurement".

**Agrupamento de Dados:**

O dataframe df resultante é agrupado pelo nome da estação ("station_name").
Utiliza o método .agg() para calcular estatísticas resumidas para cada grupo, incluindo o valor mínimo, médio e máximo das medições de temperatura.
Os resultados são renomeados usando .alias() para fornecer nomes mais significativos.
Por fim, os grupos são classificados pelo nome da estação e coletados em um stream (transmissão) usando .collect(streaming=True).

**Impressão dos Resultados Finais:**

Itera sobre os resultados agrupados, formatando e imprimindo os dados.
Cada iteração extrai os valores da estação, do mínimo, da média e do máximo da medição de temperatura.
Os dados são impressos no formato "estação=min/med/max, ", onde min, med e max representam o valor mínimo, médio e máximo, respectivamente, das medições de temperatura para cada estação.
O caractere de retrocesso (\b) é utilizado para remover a vírgula e o espaço em branco extras no final da impressão.
O resultado final é uma sequência no formato {estação=min/med/max, ...} representando as estatísticas resumidas para cada estação.

# Funcionamento comprovado

Link da [ponderada](https://drive.google.com/file/d/1uzo3Nr67_4d_wbUPcq_1LZjzA3IFpFUr/view?usp=sharing) rodando corretamente
