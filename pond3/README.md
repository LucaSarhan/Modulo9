# Relatório

Este relatório apresenta uma análise detalhada dos fundamentos de segurança da informação aplicados a ambientes de Internet das Coisas (IoT). Exploraremos conceitos essenciais, como a proteção de dados em trânsito e em armazenamento, bem como a autenticação e autorização da comunicação. Além disso, discutiremos a relevância e a eficácia de cada uma dessas estratégias de proteção, destacando sua importância para garantir a segurança e integridade das informações em ambientes IoT.

Pilares do CIA Triad:

Integridade
Disponibilidade
Confidencialidade

**Pergunta:** 

O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do browser? Algum pilar do CIA Triad é violado com isso?

**Resposta:** 

Quando a segunda sessão do browser é conectado no ClientID a primeira sessão vai ser desconectada automaticamente. Isso pode resultar em uma violação do pilar de Integridade. O pilar de Integridade refere-se à garantia de que os dados permaneçam íntegros, precisos e confiáveis ao longo do tempo e durante todo o processo de manipulação. Quando o mesmo ClientID é usado em diferentes máquinas ou sessões do navegador, pode haver uma falta de integridade nos dados associados a esse identificador. Isso ocorre porque não há garantia de que as alterações feitas em uma máquina ou sessão sejam refletidas de forma consistente em todas as outras instâncias.

**Pergunta:** 

Com os parâmetros de ```resources```, algum pilar do CIA Triad pode ser facilmente violado?

**Resposta:** 

Os parâmetros de recursos definidos estão relacionados à alocação de recursos de computação para o contêiner MQTT5, como CPU e memória. Embora esses parâmetros possam influenciar o desempenho e a disponibilidade do serviço, não estão diretamente ligados aos pilares. Por exemplo, se 