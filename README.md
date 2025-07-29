# Go API

API RESTful desenvolvida em **Go** com conexão ao **SQL Server**.

---

## Como instalar e rodar

### 1. Clone o repositório
```bash
git clone https://github.com/CDSCarlos/go-api.git
cd go-api

### 2. Configure as variáveis de ambiente. Crie um arquivo .env na raiz do projeto com as informações do banco

DB_HOST=localhost
DB_PORT=1433
DB_USER=sa
DB_PASS=your_password
DB_NAME=go_api_db

### 3. Instale as dependências
go mod tidy

### 4. Rode a aplicação
go run main.go

### 5. Acesse no navegador ou via API client
http://localhost:8080