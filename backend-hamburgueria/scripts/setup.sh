#!/bin/bash

echo "🚀 Configurando Backend da Hamburgueria..."

# Verificar se Go está instalado
if ! command -v go &> /dev/null; then
    echo "❌ Go não está instalado. Por favor, instale o Go 1.21 ou superior."
    exit 1
fi

echo "✅ Go encontrado: $(go version)"

# Verificar se PostgreSQL está instalado
if ! command -v psql &> /dev/null; then
    echo "❌ PostgreSQL não está instalado. Por favor, instale o PostgreSQL."
    exit 1
fi

echo "✅ PostgreSQL encontrado"

# Instalar dependências
echo "📦 Instalando dependências..."
go mod tidy

# Criar arquivo .env se não existir
if [ ! -f .env ]; then
    echo "📝 Criando arquivo .env..."
    cp config.env.example .env
    echo "✅ Arquivo .env criado. Edite as configurações conforme necessário."
else
    echo "✅ Arquivo .env já existe."
fi

echo ""
echo "🎉 Setup concluído!"
echo ""
echo "📋 Próximos passos:"
echo "1. Configure o banco de dados PostgreSQL"
echo "2. Edite o arquivo .env com suas configurações"
echo "3. Execute: go run main.go"
echo ""
echo "📚 Para mais informações, consulte o README.md" 