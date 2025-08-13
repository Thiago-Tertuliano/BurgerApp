# Documentação da API - BurgerApp

## Introdução

Esta API permite gerenciar pedidos, clientes e produtos de uma hamburgueria.

---

## Endpoints

### Autenticação

#### `POST /login`
Autentica um usuário e retorna um token JWT.

**Request Body:**
```json
{
  "username": "usuario",
  "password": "senha"
}
```

**Response:**
```json
{
  "token": "jwt-token"
}
```

---

### Produtos

#### `GET /produtos`
Lista todos os produtos disponíveis.

**Response:**
```json
[
  {
    "id": 1,
    "nome": "Hambúrguer Clássico",
    "preco": 25.00,
    "descricao": "Pão, carne, queijo, alface, tomate"
  }
]
```

#### `POST /produtos`
Adiciona um novo produto.

**Request Body:**
```json
{
  "nome": "Hambúrguer Vegano",
  "preco": 28.00,
  "descricao": "Pão integral, hambúrguer de grão de bico, alface, tomate"
}
```

---

### Pedidos

#### `GET /pedidos`
Lista todos os pedidos.

**Response:**
```json
[
  {
    "id": 1,
    "cliente": "João",
    "produtos": [1, 2],
    "status": "em preparo"
  }
]
```

#### `POST /pedidos`
Cria um novo pedido.

**Request Body:**
```json
{
  "cliente": "Maria",
  "produtos": [1, 3]
}
```

---

### Clientes

#### `GET /clientes`
Lista todos os clientes.

**Response:**
```json
[
  {
    "id": 1,
    "nome": "João",
    "email": "joao@email.com"
  }
]
```

#### `POST /clientes`
Adiciona um novo cliente.

**Request Body:**
```json
{
  "nome": "Maria",
  "email": "maria@email.com"
}
```

---

## Códigos de Status

- `200 OK` – Sucesso
- `201 Created` – Recurso criado
- `400 Bad Request` – Erro na requisição
- `401 Unauthorized` – Não autorizado
- `404 Not Found` – Não encontrado

---

## Observações

- Todos os endpoints que alteram dados exigem autenticação via JWT.
- Os exemplos acima são ilustrativos. Consulte o