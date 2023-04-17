
> ## Andamento Curso
- Introdução
	- [x] O que é a Linguagem GO
	- [x] Histórico
	- [x] O que GO NÃO é
	- [x] Motivações
	- [x] Código fonte
- Configurando ambiente
	- [x] Instalação
	- [x] Principais pastas
	- [x] Go no VSCode
- Fundação
	- [x] Entendendo a primeira linha
	- [x] Declaração e atribuição
	- [x] Criação de tipos
	- [x] Importando fmt e tipagem
	- [x] Percorrendo Arrays
	- [x] Slices
	- [x] Maps
	- [x] Funções
	- [x] Funções variádicas
	- [x] Closures
	- [x] Iniciando com Structs
	- [x] Composição de Structs
	- [x] Métodos em Structs
	- [x] Interfaces
	- [x] Ponteiros
	- [x] Quando usar ponteiros
	- [x] Ponteiros e Structs
	- [x] Interfaces vazias
	- [x] Type assertation
	- [ ] Generics
	- [ ] Pacotes e módulos parte 1
	- [ ] Pacotes e módulos parte 2
	- [ ] Pacotes e módulos parte 3
	- [ ] Instalando pacotes
	- [ ] For
	- [ ] Condicionais
	- [ ] Compilando projetos
- Pacotes importantes
	- [ ] Manipulação de arquivos
	- [ ] Realizando chamada HTTP
	- [ ] Defer
	- [ ] Trabalhando com json
	- [ ] Busca CEP
	- [ ] Iniciando com HTTP
	- [ ] Manipulando Headers
	- [ ] Criando função para BuscaCEP
	- [ ] Finalizando resposta para o servidor HTTP
	- [ ] ServeMux
	- [ ] FileServer
	- [ ] Iniciando com templates
	- [ ] Template.Must
	- [ ] Templates com arquivo externo
	- [ ] Templates com webserver
	- [ ] Compondo templates
	- [ ] Mapeando funções nos templates
	- [ ] HttpClient com Timeout
	- [ ] Trabalhando com Post
	- [ ] Customizando Request
	- [ ] Trabalhando com HTTP usando Contextos
	- [ ] Documentação para Template
- Context
	- [ ] Introdução aos contextos
	- [ ] Entendendo conceitos básicos
	- [ ] Context utilizando server HTTP
	- [ ] Context no lado do Client
	- [ ] Context WithValue
- Banco de dados
	- [ ] Introdução a banco de dados
	- [ ] Preparando base de código
	- [ ] Inserindo dados no banco
	- [ ] Alterando dados no banco
	- [ ] Trabalhando com QueryRow
	- [ ] Selecionando múltiplos registros
	- [ ] Removendo registro
	- [ ] Iniciando com GORM
	- [ ] Configurando e criando operações
	- [ ] Realizando primeiras consultas
	- [ ] Realizando consultas com where
	- [ ] Alterando e removendo registros
	- [ ] Trabalhando com Soft delete
	- [ ] Belongs to
	- [ ] Has One
	- [ ] Has many
	- [ ] Pegadinha no has many
	- [ ] Many to Many
	- [ ] Lock otimista vs Pessimista
	- [ ] Desafio Client-Server-API
- Packaging
	- [ ] Introduzindo go mod
	- [ ] Acessando pacotes criados
	- [ ] Exportação de objetos
	- [ ] Entendendo go mod
	- [ ] Trabalhando com go mod replace
	- [ ] Usando workspaces
- Testing
	- [ ] Iniciando com testes automatizados
	- [ ] Testando em batch
	- [ ] Verificando cobertura de código
	- [ ] Trabalhando com benchmarking
	- [ ] Fuzzing
	- [ ] Iniciando com Testify
	- [ ] Trabalhando com Mocks
- APIs	
	- [ ] Falando sobre APIs
	- [ ] Estruturando diretórios
	- [ ] Criando arquivo de configuração
	- [ ] Finalizando configuração
	- [ ] Outras possibilidades de configuração
	- [ ] Criando entidade User
	- [ ] Testando User
	- [ ] Criando entidade Product
	- [ ] Testando Product
	- [ ] Criando UserDB
	- [ ] Testando criação do usuário
	- [ ] Finalizando teste de UserDB
	- [ ] Criando principais métodos de ProductDB
	- [ ] Criando FindAll com paginação
	- [ ] Testando ProductDB
	- [ ] Finalizando testes de ProductDB
	- [ ] Criando handler para criar produto
	- [ ] Testando endpoint de criação de Product
	- [ ] Ajustando package para os handlers
	- [ ] Web frameoworks vs Frameworks
	- [ ] Roteadores
	- [ ] Buscando e alterando products
	- [ ] Listando e removendo Products
	- [ ] Criando user
	- [ ] Falando sobre JWT
	- [ ] Gerando token JWT
	- [ ] Protegendo products
	- [ ] Criando e trabalhando com middlewares
	- [ ] Iniciando com documentação da API
	- [ ] Gerando primeira documentação
	- [ ] Acessando swagger pela primeira vez
	- [ ] Documentando criação do User
	- [ ] Trabalhando na geração do access token
	- [ ] Criando e listando products
	- [ ] Finalizando docs
- Multithreading
	- [ ] Iniciando com Processos
	- [ ] Introdução a concorrência e Mutex
	- [ ] Concorrência vs Paralelismo vs Go
	- [ ] Multithreading
	- [ ] Schedulers
	- [ ] Go e suas green threads
	- [ ] Iniciando com Go Routines
	- [ ] Trabalhando com Wait Groups
	- [ ] Problema simples de concorrência
	- [ ] Entendendo Mutex e Operações Atômicas
	- [ ] Iniciando com Channels
	- [ ] Forever
	- [ ] Iterando com range
	- [ ] Range With WaitGroup
	- [ ] Channel Directions
	- [ ] Criando um Load Balancer
	- [ ] Trabalhando com Select
	- [ ] Canais com Buffer
	- [ ] Desafio Multithreading
- Manipulando eventos
	- [ ] Introdução a eventos
	- [ ] Elementos táticos de um contexto de eventos
	- [ ] Criando Interfaces da solução
	- [ ] Registrando eventos
	- [ ] Criando suite de testes
	- [ ] Testando Register
	- [ ] Testando Registro de Handlers Repetidos
	- [ ] Implementando e testando método Clear
	- [ ] Implementando e testando método Has
	- [ ] Implementando método Dispatch
	- [ ] Revisitando slices
	- [ ] Removendo handlers
	- [ ] Adicionando go routine no event dispatcher
	- [ ] Utilizando go routines no Dispatcher
	- [ ] Instalando RabbitMQ
	- [ ] Entendendo o RabbitMQ
	- [ ] Consumindo mensagens do RabbitMQ
	- [ ] Criando consumidor na pratica
	- [ ] Produzindo e consumindo mensagens
- Módulos privados
	- [ ] Introdução
	- [ ] Criando repositório privado
	- [ ] Configurando GOPRIVATE
	- [ ] Errata
	- [ ] Autenticando no Bitbucket
	- [ ] Go Proxy vs Vendor
- GraphQL
	- [ ] Introdução ao GraphQL
	- [ ] Gerando esqueleto do servidor GraphQL
	- [ ] Criando Schema GraphQL
	- [ ] Gerando esqueleto de nossa aplicação
	- [ ] Criando resolver para Category
	- [ ] Persistindo Category via Playground
	- [ ] Fazendo queries de Category
	- [ ] Implementando CourseDB
	- [ ] Criando resolver de CreateCourse
	- [ ] Implementando QueryCourses
	- [ ] Dados encadeados
	- [ ] Finalizando encadeamento de categorias
	- [ ] gqlgen
- gRPC
	- [ ] Iniciando com gRPC
	- [ ] Conceitos iniciais
	- [ ] gRPC HTTP2 e Protocol Buffers
	- [ ] Formatos de comunicação
	- [ ] REST vs gRPC
	- [ ] gRPC vs Protocol Buffers
	- [ ] Instalando compilador e plugins
	- [ ] Fazendo setup do projeto
	- [ ] Criando nosso protofile
	- [ ] Fazendo geração de código com protoc
	- [ ] Implementando CreateCategory
	- [ ] Criando servidor gRPC
	- [ ] Interagindo com evans
	- [ ] Criando categoryList no protofile
	- [ ] Listando categories
	- [ ] Buscando uma categoria
	- [ ] Trabalhando com stream
	- [ ] Trabalhando com streams bidirecionais
- Upload S3
	- [ ] Entendendo o problema referente a upload
	- [ ] Gerando arquivos exemplo
	- [ ] Configurando AWS session
	- [ ] Desenvolvendo função de upload
	- [ ] Finalizando primeira implementação
	- [ ] Criando credenciais na AWS
	- [ ] Fazendo upload de forma serial
	- [ ] Realizando uploads usando go routines
	- [ ] Limitando quantidade máxima de upload
	- [ ] Fazendo retentativas de erro
- Cobra CLI
	- [ ] Introdução sobre CLI
	- [ ] Setup básico da aplicação
	- [ ] Inicializando projeto cobra
	- [ ] Criando nossos primeiros comandos
	- [ ] Comandos encadeados
	- [ ] Flags locais vs globais
	- [ ] Manipulando flags
	- [ ] Flags mudando valor por referência
	- [ ] Entendendo hooks
	- [ ] Trabalhando com banco de dados
	- [ ] Inversão de controle ao executar comandos
- SQLC
	- [ ] Mais sobre banco de dados
	- [ ] Trabalhando com migrations
	- [ ] Golang-migrate
	- [ ] Falando sobre SQLX
	- [ ] Iniciando com SQLC
	- [ ] Criando CRUD
	- [ ] Finalizando CRUD
	- [ ] Iniciando com transações
	- [ ] Implementando método para transação
	- [ ] Vendo a transação funcionar
	- [ ] Ajustando tipo de campo
	- [ ] Exibindo dados com Join
	- [ ] Repositório do SQLC
- UOW
	- [ ] Link do projeto
	- [ ] Apresentando caso
	- [ ] Entendendo inconsistência
	- [ ] Entendendo Unit of Work
	- [ ] Criando interface do UOW
	- [ ] Registrando repositórios
	- [ ] Implementando principais métodos
	- [ ] Implementando GetRepository
	- [ ] Criando novo caso de uso
	- [ ] Testando implementação com uow
- DI
	- [ ] Introdução sobre DI
	- [ ] Criando situação problema
	- [ ] Exibindo árvore de dependencias
	- [ ] Bibliotecas de DI
	- [ ] Wire na prática
	- [ ] Trabalhando com sets e Interfaces
	- [ ] Sem pontas soltas
	- [ ] Google Wire
- Clean Architecture
	- [ ] Introdução
	- [ ] Origem da Clean Architecture
	- [ ] Pontos importantes sobre arquitetura
	- [ ] Keep options open
	- [ ] Use Cases
	- [ ] O fluxo dos Use Cases
	- [ ] Limites arquiteturais
	- [ ] Input vs Output
	- [ ] Entendendo DTOs
	- [ ] Presenters
	- [ ] Entities vs DDD
	- [ ] Nosso objetivo
	- [ ] Entity e UseCase
	- [ ] Repositorios e EventHandler
	- [ ] Webserver
	- [ ] gRPC
	- [ ] GraphQL
	- [ ] Configuração com Viper
	- [ ] RabbitMQ e Google Wire
	- [ ] Executando servidores
	- [ ] Tudo funcionando
	- [ ] Sobre o desafio
	- [ ] Código fonte
- Deploy com Docker e Kubernetes
	- [ ] O que faremos
	- [ ] Criando projeto exemplo
	- [ ] Criando Dockerfile
	- [ ] Go com Docker em modo dev
	- [ ] Entendendo processo de build
	- [ ] Otimizando geração do executável
	- [ ] Gerando imagem otimizada
	- [ ] C GO e seus impactos
	- [ ] Criando cluster kubernetes com Kind
	- [ ] Criando primeiro deployment
	- [ ] Criando service no kubernetes
	- [ ] Probes
	- [ ] Palavras finais