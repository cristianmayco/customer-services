# Customer Services API

## **Descrição**
A `Customer Services API` é uma aplicação RESTful desenvolvida em Go utilizando o framework Gin. A API permite gerenciar clientes e serviços, protegendo os endpoints com autenticação baseada em JWT (JSON Web Token).

---

## **Pré-requisitos**
- Go 1.18 ou superior
- Biblioteca Gin: `github.com/gin-gonic/gin`
- Biblioteca JWT: `github.com/golang-jwt/jwt/v5`

---

## **Instalação**
1. Instale as dependências:
   ```bash
   go mod tidy
   ```

2. Execute a aplicação:
   ```bash
   go run cmd/main.go
   ```

A aplicação estará disponível em `http://localhost:8080`.

---

## **Endpoints**

### **1. Login**
#### **POST /login**
Gera um token JWT para autenticação.

- **Requisição:**
  ```json
  {
    "username": "admin",
    "password": "admin"
  }
  ```

- **Resposta (200 - Sucesso):**
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

- **Resposta (401 - Credenciais inválidas):**
  ```json
  {
    "error": "Invalid credentials"
  }
  ```

---

### **2. Clientes**
Todos os endpoints abaixo exigem o token JWT no cabeçalho `Authorization` no formato:

```
Authorization: Bearer <token>
```

#### **GET /customers**
Retorna todos os clientes.

- **Resposta (200):**
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "customer_identity": "123456789",
      "status": "active",
      "create_at": "2025-03-15 10:00:00",
      "update_at": "2025-03-15 10:00:00"
    }
  ]
  ```

---

#### **GET /customers/:id**
Retorna um cliente pelo ID.

- **Resposta (200):**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "customer_identity": "123456789",
    "status": "active",
    "create_at": "2025-03-15 10:00:00",
    "update_at": "2025-03-15 10:00:00"
  }
  ```

- **Resposta (404 - Cliente não encontrado):**
  ```json
  {
    "error": "Customer not found"
  }
  ```

---

#### **POST /customers**
Cria um novo cliente.

- **Requisição:**
  ```json
  {
    "name": "Jane Doe",
    "customer_identity": "987654321",
    "status": "active"
  }
  ```

- **Resposta (201 - Criado):**
  ```json
  {
    "id": 2,
    "name": "Jane Doe",
    "customer_identity": "987654321",
    "status": "active",
    "create_at": "2025-03-15 10:30:00",
    "update_at": "2025-03-15 10:30:00"
  }
  ```

---

#### **PUT /customers/:id**
Atualiza um cliente existente.

- **Requisição:**
  ```json
  {
    "name": "Jane Smith",
    "customer_identity": "987654321",
    "status": "inactive"
  }
  ```

- **Resposta (200 - Atualizado):**
  ```json
  {
    "id": 2,
    "name": "Jane Smith",
    "customer_identity": "987654321",
    "status": "inactive",
    "create_at": "2025-03-15 10:30:00",
    "update_at": "2025-03-15 11:00:00"
  }
  ```

- **Resposta (404 - Cliente não encontrado):**
  ```json
  {
    "error": "Customer not found"
  }
  ```

---

#### **DELETE /customers/:id**
Remove um cliente pelo ID.

- **Resposta (200 - Removido):**
  ```json
  {
    "message": "Customer deleted"
  }
  ```

- **Resposta (404 - Cliente não encontrado):**
  ```json
  {
    "error": "Customer not found"
  }
  ```

---

## **Exemplo de Fluxo Completo**

1. **Login para obter o token JWT:**
   ```bash
   curl -X POST http://localhost:8080/login \
   -H "Content-Type: application/json" \
   -d '{"username": "admin", "password": "admin"}'
   ```

2. **Usar o token para acessar os endpoints protegidos:**
   ```bash
   curl -X GET http://localhost:8080/customers \
   -H "Authorization: Bearer <token>"
   ```

---

## **Notas**
- Substitua `"secret"` no código pelo valor de uma chave segura para produção.
- Use HTTPS para proteger a comunicação entre cliente e servidor.
- Implemente um banco de dados para persistir os dados de clientes e usuários.

---

## **Licença**
Este projeto está licenciado sob a licença MIT.
```
