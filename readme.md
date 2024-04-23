# Investment Growth Calculator

## Descrição
O "Investment Growth Calculator" é uma ferramenta de linha de comando desenvolvida em Go que permite aos usuários calcular o crescimento futuro de um investimento com base em contribuições mensais e uma taxa de rentabilidade anual. O aplicativo é destinado a investidores que desejam planejar seus investimentos e visualizar o potencial de crescimento ao longo dos anos.

## Instalação

### Pré-requisitos
Para rodar este programa, você precisará do Go instalado em sua máquina. Você pode baixar e instalar o Go em [https://golang.org/dl/](https://golang.org/dl/).

### Baixando o Programa
Clone o repositório para sua máquina local usando:

```
git clone https://github.com/gabrielcaiana/go-investiment-growth-calculator
```

cd investment-growth-calculator


## Uso
Para usar o programa, execute o seguinte comando no terminal dentro do diretório do projeto:

```
go run main.go
```

## Utilizando via docker

Para buildar a imagem
```
docker build -t go-investiment .
```

Para rodar o container de forma interativa

```
docker run -it go-investiment
```

Você será solicitado a inserir o valor inicial do investimento, o valor mensal recorrente, a rentabilidade anual em percentual, e a duração do investimento em anos.

## Rodando os Testes
Para executar os testes unitários associados ao programa, use o comando:

```
go test
```

Isso testará a lógica principal de cálculo do programa para garantir que tudo está funcionando como esperado.

## Contribuições
Contribuições são sempre bem-vindas! Se você tem alguma sugestão para melhorar o aplicativo ou quer adicionar novas funcionalidades, sinta-se à vontade para criar um fork do repositório e submeter suas mudanças por meio de um pull request.


