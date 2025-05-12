# 🚀 ValidationAPP

[![Stars](https://img.shields.io/github/stars/KaduSantanaDev/ValidationAPP?style=social)](https://github.com/KaduSantanaDev/ValidationAPP/stargazers)
[![Issues](https://img.shields.io/github/issues/KaduSantanaDev/ValidationAPP)](https://github.com/KaduSantanaDev/ValidationAPP/issues)
[![Contributors](https://img.shields.io/github/contributors/KaduSantanaDev/ValidationAPP)](https://github.com/KaduSantanaDev/ValidationAPP/graphs/contributors)

## Índice

- [Descrição](#descrição)  
- [Visão Geral e Features](#visão-geral-e-features)  
- [Arquitetura](#arquitetura)  
- [Tecnologias](#tecnologias)  
- [Pré-requisitos](#pré-requisitos)  
- [Instalação & Execução](#instalação--execução)  
- [Variáveis de Ambiente](#variáveis-de-ambiente)  
- [Uso da API](#uso-da-api)  

## Descrição

O **ValidationAPP** é um sistema que valida eletronicamente documentos de identidade (RG) usando OCR e microserviços com comunicação via Apache Kafka.  

### Visão Geral e Features

- Validação OCR de imagens de RG (via Python + Tesseract)  
- API Gateway em Go com arquitetura hexagonal  
- Comunicação assíncrona via Apache Kafka  
- Persistência em PostgreSQL  
- UI em React para cadastro e consulta de clientes  

## Tecnologias

- **Back-End**: Go, go-chi, Kafka, PostgreSQL, Python, Pytesseract, Docker  
- **Front-End**: React, JavaScript, Bootstrap, React Router

## Arquitetura
![Diagrama de Arquitetura](imgs/arch.png)

## Pré-requisitos

- [Docker](https://www.docker.com/) ≥ 20.x  
- [docker-compose](https://docs.docker.com/compose/) ≥ 1.29  
## Instalação & Execução
```bash
git clone https://github.com/KaduSantanaDev/ValidationAPP.git
cd ValidationAPP

# Subir todos os serviços em background
docker-compose up --build -d
```
### A aplicação estará disponível em:
- UI React: http://localhost:5173
- API: http://localhost:3031/api/v1/clients

## Variáveis de Ambiente

| Chave         | Descrição                | Valor Padrão |
| ------------- | ------------------------ | ------------ |
| DB\_HOST      | Host do Postgres         | `postgres`   |
| DB\_PORT      | Porta do Postgres        | `5432`       |
| DB\_USER      | Usuário do Postgres      | `root`       |
| DB\_PASSWORD  | Senha do Postgres        | `secret`     |
| DB\_NAME      | Nome do banco            | `knowyourfan`|

## Uso da API
### Criar cliente
```bash
curl -X POST http://localhost:3031/api/v1/clients \
  -F "name=Fulano da Silva" \
  -F "email=fulano@example.com" \
  -F "cpf=12345678900" \
  -F "document=@/caminho/para/rg.jpg" \
  -F "street=Rua das Flores" \
  -F "number=123" \
  -F "complement=Apto 45" \
  -F "neighborhood=Centro" \
  -F "city=São Paulo" \
  -F "state=SP" \
  -F "cep=01001000"
```
### Listar todos os clientes
```bash
curl http://localhost:3031/api/v1/clients
```
### Buscar por ID

```bash
curl http://localhost:3031/api/v1/clients?id=123
```