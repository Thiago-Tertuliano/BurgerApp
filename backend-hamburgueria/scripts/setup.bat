@echo off
echo 🚀 Configurando Backend da Hamburgueria...

REM Verificar se Go está instalado
go version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Go não está instalado. Por favor, instale o Go 1.21 ou superior.
    pause
    exit /b 1
)

echo ✅ Go encontrado
go version

REM Verificar se PostgreSQL está instalado
psql --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ PostgreSQL não está instalado. Por favor, instale o PostgreSQL.
    pause
    exit /b 1
)

echo ✅ PostgreSQL encontrado

REM Instalar dependências
echo 📦 Instalando dependências...
go mod tidy

REM Criar arquivo .env se não existir
if not exist .env (
    echo 📝 Criando arquivo .env...
    copy config.env.example .env
    echo ✅ Arquivo .env criado. Edite as configurações conforme necessário.
) else (
    echo ✅ Arquivo .env já existe.
)

echo.
echo 🎉 Setup concluído!
echo.
echo 📋 Próximos passos:
echo 1. Configure o banco de dados PostgreSQL
echo 2. Edite o arquivo .env com suas configurações
echo 3. Execute: go run main.go
echo.
echo 📚 Para mais informações, consulte o README.md
pause 