# AnimeAPI em Go (Golang)

![Golang](https://media1.giphy.com/media/en4M5qpoxaOyUcDmYU/200w.webp?cid=ecf05e47l291nbrx71kziekjpx4kvv94k4yehjl7gkqpml3o&ep=v1_gifs_search&rid=200w.webp&ct=g)

Este projeto é uma API simples para gerenciar informações de animes, temporadas e episódios. Desenvolvido em Golang, utiliza o banco de dados MySQL para armazenar os dados.

## Funcionalidades

- Listar todos os animes
- Obter detalhes de um anime específico
- Criar um novo anime
- Atualizar informações de um anime existente
- Deletar um anime
- Deletar todos os animes

- Listar todos os episódios
- Obter detalhes de um episódio específico
- Criar um novo episódio
- Atualizar informações de um episódio existente
- Deletar um episódio
- Deletar todos os episódios

## Endpoints

### Animes

- **GET** `/api/v1/animes` - Listar todos os animes
- **GET** `/api/v1/animes?page=1&perPage=5` - Listar todos os animes de modo paginado (**Novo**)
- **GET** `/api/v1/animes/{id}` - Obter detalhes de um anime específico
- **POST** `/api/v1/animes` - Criar um novo anime
- **PUT** `/api/v1/animes/{id}` - Atualizar informações de um anime existente
- **DELETE** `/api/v1/animes/{id}` - Deletar um anime
- **DELETE** `/api/v1/animes` - Deletar todos os animes

### Episódios

- **GET** `/api/v1/episodios` - Listar todos os episódios
- **GET** `/api/v1/episodios/{id}` - Obter detalhes de um episódio específico
- **POST** `/api/v1/episodios` - Criar um novo episódio
- **PUT** `/api/v1/episodios/{id}` - Atualizar informações de um episódio existente
- **DELETE** `/api/v1/episodios/{id}` - Deletar um episódio
- **DELETE** `/api/v1/episodios` - Deletar todos os episódios

## Autenticação

Para testar os endpoints, é necessário fornecer um token de autenticação válido no cabeçalho da solicitação. O token de autenticação pode ser configurado no arquivo de ambiente com a variável `AUTH_TOKEN`.

## Pré-requisitos

- Go (Golang)
- MySQL

## Executando a API

Para iniciar a API, execute o seguinte comando:

go run ./app/main.go

A API estará disponível em `http://localhost:8000`.

## Contribuição

Sinta-se à vontade para contribuir com melhorias e correções de bugs. Abra um pull request e ficarei feliz em revisá-lo.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
