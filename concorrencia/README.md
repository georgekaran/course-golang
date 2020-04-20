# Concorrência

Neste seção será trabalhada sobre o tema de concorrência dentro de GO.
***
## Concorrência vs Paralelismo

Go é uma linguagem concorrente.

**Paralelismo** - Executar código simultamente em processadores físicos diferentes.

**Concorrência** - Intercalar (administrar) vários processos ao mesmo tempo e 
isso pode ocorre em um único procesador físico.

**Concorrência** - É uma forma de estruturar o seu programa e pode ser vista como a 
composição de processos (tipicamente funções) que executam de forma independente.

#### Concorrência viabiliza paralelismo.

É possível que a concorrência seja melhor que o paralelismo!
Paralelismo é exige muito mais do SO e do hardware.