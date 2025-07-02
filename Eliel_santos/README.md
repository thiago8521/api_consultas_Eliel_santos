# API de Agendamento de Consultas Médicas – SaúdeMais

Implementação em Go (Gin + GORM + JWT) referente ao desafio técnico.

## Variáveis de ambiente

Copie `.env.example` para `.env` e ajuste conforme seu ambiente.

## Subindo a stack com Docker Compose

```bash
# Na raiz do repositório
docker-compose up -d
```

Isso sobe:
* **PostgreSQL** em `localhost:5432`
* **Adminer** em `localhost:8080`

Em outro terminal, com Go instalado, execute:

```bash
cd seu-nome
go run .
```

A API ficará em `http://localhost:8080`.

## Endpoints

| Método | Rota | Protegido | Descrição |
|--------|------|-----------|-----------|
| POST | /register | ❌ | Cadastro de paciente |
| POST | /login | ❌ | Login e geração de JWT |
| POST | /appointments | ✔️ | Agendar consulta |
| GET | /appointments | ✔️ | Listar consultas |
| DELETE | /appointments/:id | ✔️ | Cancelar consulta |

Mensagens de erro seguem o formato do desafio.
