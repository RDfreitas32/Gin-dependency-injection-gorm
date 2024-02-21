# Aplicação de Gerenciamento de Tarefas em Go

Esta é uma aplicação simples de gerenciamento de tarefas criada em Go, usando o framework Gin para criar uma API RESTful e o GORM ORM para interagir com um banco de dados SQLite.
## Uso

A API oferece os seguintes endpoints para gerenciar tarefas:

- `POST /tasks`: Cria uma nova tarefa.
- `GET /tasks`: Retorna todas as tarefas.
- `GET /tasks/:id`: Retorna uma tarefa específica pelo ID.
- `PUT /tasks/:id`: Atualiza uma tarefa existente pelo ID.
- `DELETE /tasks/:id`: Exclui uma tarefa existente pelo ID.

Para usar a API, você pode enviar solicitações HTTP para esses endpoints usando um cliente REST como cURL ou Postman.
