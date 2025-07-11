# 🍔 BurgerApp - Sistema de Pedidos Frontend

## 📋 Descrição

Sistema de pedidos para hamburgueria desenvolvido em **Vue.js 3** com **Composition API**. O projeto permite visualizar o cardápio, montar lanches personalizados e gerenciar pedidos em tempo real através de uma interface moderna e responsiva.

## 🚀 Tecnologias Utilizadas

### Core
- **Vue.js 3** - Framework JavaScript progressivo
- **Composition API** - Sistema de reatividade do Vue 3
- **Vite** - Build tool e dev server moderno
- **CSS3** - Estilização com variáveis CSS e animações

### Dependências
- `vue` ^3.5.17 - Framework principal
- `@vitejs/plugin-vue` ^6.0.0 - Plugin Vue para Vite
- `vite` ^7.0.0 - Bundler e dev server
- `vite-plugin-vue-devtools` ^7.7.7 - Ferramentas de desenvolvimento

## 📁 Estrutura do Projeto

```
projeto-hamburgueria/
├── public/                 # Arquivos estáticos
│   ├── favicon.ico        # Ícone da aplicação
│   └── *.jpg/*.jfif       # Imagens dos produtos
├── src/                    # Código fonte
│   ├── assets/            # Recursos estáticos
│   │   ├── base.css       # Estilos base
│   │   ├── main.css       # Estilos principais
│   │   └── outback-theme.css # Tema personalizado
│   ├── components/        # Componentes Vue
│   │   ├── Header.vue     # Cabeçalho da aplicação
│   │   ├── Menu.vue       # Cardápio de produtos
│   │   ├── Kitchen.vue    # Interface da cozinha
│   │   └── CustomBurger.vue # Montador de lanches
│   ├── api.js             # Configuração da API
│   ├── App.vue            # Componente raiz
│   └── main.js            # Ponto de entrada
├── index.html             # HTML base
├── package.json           # Dependências e scripts
├── vite.config.js         # Configuração do Vite
└── README.md              # Este arquivo
```

## 🎯 Funcionalidades

### 1. **Cardápio Interativo**
- Visualização de produtos por categoria
- Filtros dinâmicos
- Cards com imagens e preços
- Interface responsiva

### 2. **Montador de Lanches Personalizados**
- Seleção de ingredientes por categoria
- Cálculo automático de preços
- Validação de ingredientes obrigatórios
- Interface intuitiva

### 3. **Sistema de Pedidos**
- Adição de produtos ao carrinho
- Envio automático para o backend
- Contador de pedidos ativos
- Sincronização em tempo real

### 4. **Interface da Cozinha**
- Painel lateral responsivo
- Lista de pedidos em tempo real
- Controle de status (preparando → pronto → entregue)
- Barra de progresso visual
- Atualização automática a cada 10 segundos

## 🔧 Configuração e Instalação

### Pré-requisitos
- **Node.js** (versão 16 ou superior)
- **npm** ou **yarn**
- Backend rodando na porta 8080

### Instalação

1. **Clone o repositório**
```bash
git clone <url-do-repositorio>
cd projeto-hamburgueria
```

2. **Instale as dependências**
```bash
npm install
```

3. **Configure a API**
Edite o arquivo `src/api.js` se necessário:
```javascript
const API_URL = 'http://localhost:8080/api';
```

4. **Execute o projeto**
```bash
npm run dev
```

5. **Acesse a aplicação**
Abra `http://localhost:5173` no navegador

## 📖 Como Usar

### Desenvolvimento
```bash
npm run dev          # Inicia servidor de desenvolvimento
npm run build        # Gera build de produção
npm run preview      # Preview do build de produção
```

### Scripts Disponíveis
- `dev` - Servidor de desenvolvimento com hot reload
- `build` - Build otimizado para produção
- `preview` - Serve os arquivos de produção localmente

## 🏗️ Arquitetura

### Componentes Principais

#### 1. **App.vue** - Componente Raiz
- Gerencia estado global da aplicação
- Controla comunicação entre componentes
- Integra com API do backend
- Gerencia ciclo de vida dos pedidos

#### 2. **Header.vue** - Cabeçalho
- Logo e informações da marca
- Botão de acesso à cozinha
- Contador de pedidos ativos
- Interface responsiva

#### 3. **Menu.vue** - Cardápio
- Lista produtos por categoria
- Filtros dinâmicos
- Integração com CustomBurger
- Cards de produtos interativos

#### 4. **CustomBurger.vue** - Montador
- Seleção de ingredientes
- Cálculo de preços
- Validação de dados
- Interface de formulário

#### 5. **Kitchen.vue** - Cozinha
- Painel lateral responsivo
- Lista de pedidos em tempo real
- Controle de status
- Atualização automática

### Estado da Aplicação

```javascript
// Estado Global (App.vue)
const orders = ref([])           // Array de pedidos
const isKitchenOpen = ref(false) // Visibilidade da cozinha
const currentStep = ref('menu')  // Tela ativa
```

### Comunicação com Backend

```javascript
// Endpoints utilizados
GET  /api/categories     // Lista categorias
GET  /api/products       // Lista produtos
GET  /api/ingredients    // Lista ingredientes
POST /api/orders         // Criar pedido
PUT  /api/orders/:id/status // Atualizar status
```

## 🎨 Design System

### Paleta de Cores
- **Primárias**: `#1a1a1a`, `#ffffff`
- **Destaque**: `#ff6b35` (laranja), `#f7931e` (dourado)
- **Sucesso**: `#10b981` (verde)
- **Informação**: `#3b82f6` (azul)
- **Texto**: `#2d2d2d` (escuro), `#6b7280` (claro)

### Tipografia
- **Fonte**: Inter (Google Fonts)
- **Pesos**: 400, 500, 600, 700
- **Tamanhos**: 14px, 16px, 18px, 24px, 30px

### Componentes
- **Cards**: Bordas arredondadas, sombras suaves
- **Botões**: Estados hover, transições suaves
- **Formulários**: Validação visual, feedback imediato

## 📱 Responsividade

### Breakpoints
- **Desktop**: > 768px
- **Tablet**: 768px - 1024px
- **Mobile**: < 768px

### Adaptações Mobile
- Layout em coluna única
- Botões maiores para touch
- Espaçamentos otimizados
- Navegação simplificada

## 🔄 Fluxo de Dados

### 1. **Carregamento Inicial**
```
App.vue → API → Carrega pedidos existentes → Atualiza estado
```

### 2. **Novo Pedido**
```
Menu/CustomBurger → App.vue → API → Backend → Atualiza estado
```

### 3. **Atualização de Status**
```
Kitchen.vue → App.vue → API → Backend → Atualiza estado
```

### 4. **Sincronização Automática**
```
Kitchen.vue → API (a cada 10s) → Atualiza pedidos
```

## 🛠️ Desenvolvimento

### Estrutura de Componentes
```vue
<template>
  <!-- Template HTML -->
</template>

<script setup>
// Composition API
// Imports, estado, computed, métodos
</script>

<style scoped>
/* Estilos específicos do componente */
</style>
```

### Padrões Utilizados
- **Composition API** para lógica reativa
- **Props** para comunicação pai-filho
- **Emits** para comunicação filho-pai
- **Computed** para valores derivados
- **Lifecycle hooks** para efeitos colaterais

### Boas Práticas
- Componentes pequenos e focados
- Reutilização de código
- Nomenclatura consistente
- Comentários explicativos
- Responsividade mobile-first

## 🚀 Deploy

### Build de Produção
```bash
npm run build
```

### Arquivos Gerados
- `dist/` - Arquivos otimizados
- `dist/index.html` - HTML principal
- `dist/assets/` - CSS e JS minificados

### Servidor de Produção
```bash
npm run preview
```

## 🔧 Configurações

### Vite
```javascript
// vite.config.js
export default defineConfig({
  plugins: [vue(), vueDevTools()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
```

### Variáveis de Ambiente
```bash
# .env (se necessário)
VITE_API_URL=http://localhost:8080/api
```

## 📚 Recursos de Aprendizado

### Vue.js 3
- [Documentação Oficial](https://vuejs.org/)
- [Composition API](https://vuejs.org/guide/extras/composition-api-faq.html)
- [Guia de Estilo](https://vuejs.org/style-guide/)

### Vite
- [Documentação Oficial](https://vitejs.dev/)
- [Guia de Configuração](https://vitejs.dev/config/)

### CSS
- [CSS Variables](https://developer.mozilla.org/en-US/docs/Web/CSS/Using_CSS_custom_properties)
- [Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)
- [Grid](https://css-tricks.com/snippets/css/complete-guide-grid/)

## 🤝 Contribuição

### Como Contribuir
1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

### Padrões de Código
- Use ESLint e Prettier
- Siga o guia de estilo do Vue.js
- Escreva testes quando possível
- Documente mudanças importantes

## 👥 Autores

- **Desenvolvedor**: Thiago Matos Tertuliano
- **Data**: Julho 2025
- **Versão**: 1.0.0

---

**🍔 BurgerApp** - Transformando a experiência de pedidos de hambúrgueres!
