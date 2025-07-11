<template>
  <!-- Container principal do menu -->
  <div class="menu">
    <h2 class="menu-title">🍔 Cardápio</h2>
    
    <!-- Categorias de produtos - botões para filtrar produtos -->
    <div class="categories">
      <!-- Botões das categorias normais (excluindo 'Ingredientes') -->
      <button 
        v-for="category in categories.filter(c => c.name !== 'Ingredientes')" 
        :key="category.id"
        :class="['category-btn', { active: selectedCategory === category.id }]"
        @click="selectedCategory = category.id"
      >
        {{ category.name }}
      </button>
      <!-- Botão especial para montar lanche personalizado -->
      <button 
        :class="['category-btn', { active: selectedCategory === 'custom' }]"
        @click="selectedCategory = 'custom'"
      >
        🎯 Monte seu Lanche
      </button>
    </div>
    
    <!-- Seção de Montar Lanche - só aparece quando 'custom' está selecionado -->
    <div v-if="selectedCategory === 'custom'" class="custom-burger-section">
      <CustomBurger @add-to-cart="addCustomBurger" />
    </div>
    
    <!-- Grid de produtos normais - só aparece quando categoria normal está selecionada -->
    <div v-else class="products-grid">
      <!-- Card de cada produto -->
      <div 
        v-for="product in filteredProducts" 
        :key="product.id"
        class="product-card"
      >
        <!-- Imagem do produto -->
        <div class="product-image">
          <img :src="product.image_url" :alt="product.name">
        </div>
        <!-- Informações do produto -->
        <div class="product-info">
          <h3 class="product-name">{{ product.name }}</h3>
          <p class="product-description">{{ product.description }}</p>
          <!-- Rodapé do card com preço e botão -->
          <div class="product-footer">
            <span class="product-price">R$ {{ product.price.toFixed(2) }}</span>
            <button 
              class="add-to-cart-btn"
              @click="$emit('add-to-cart', product)"
            >
              ➕ Adicionar
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Importa funções reativas do Vue.js
import { ref, computed, onMounted } from 'vue'
// Importa componente para montar lanche personalizado
import CustomBurger from './CustomBurger.vue'
// Importa configuração da API
import API_URL from '../api'

// Define os eventos que o componente pode emitir
const emit = defineEmits(['add-to-cart'])

// ===== ESTADO REATIVO =====
// Array reativo para armazenar categorias
const categories = ref([])
// Array reativo para armazenar produtos
const products = ref([])
// Estado reativo para categoria selecionada
const selectedCategory = ref(null)

// ===== HOOK DE CICLO DE VIDA =====
// Buscar dados da API quando o componente é montado
onMounted(async () => {
  try {
    // Buscar categorias da API
    const resCat = await fetch(`${API_URL}/categories`)
    categories.value = await resCat.json()
    
    // Buscar produtos da API
    const resProd = await fetch(`${API_URL}/products`)
    products.value = await resProd.json()
    
    // Selecionar a primeira categoria por padrão
    if (categories.value.length > 0) {
      selectedCategory.value = categories.value[0].id
    }
  } catch (error) {
    console.error('Erro ao carregar dados:', error)
  }
})

// ===== COMPUTED PROPERTIES =====
// Propriedade computada que filtra produtos pela categoria selecionada
const filteredProducts = computed(() => {
  return products.value.filter(p => p.category_id === selectedCategory.value)
})

// ===== FUNÇÕES =====
// Função para adicionar lanche customizado ao carrinho
const addCustomBurger = (customBurger) => {
  // Gera ID único para o lanche customizado
  customBurger.id = 'custom-' + Date.now()
  // Emite evento para o componente pai
  emit('add-to-cart', customBurger)
}
</script>

<style scoped>
/* Container principal do menu */
.menu {
  background: var(--bg-white);
  border-radius: var(--radius-lg);
  padding: 2rem;
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border-light);
}

/* Título do menu */
.menu-title {
  font-family: 'Inter', sans-serif;
  font-size: 1.875rem;
  font-weight: 700;
  margin-bottom: 2rem;
  color: var(--text-dark);
  text-align: center;
  letter-spacing: -0.025em;
}

/* Container dos botões de categoria */
.categories {
  display: flex;
  justify-content: center;
  gap: 0.75rem;
  margin-bottom: 2rem;
  flex-wrap: wrap;
}

/* Estilos dos botões de categoria */
.category-btn {
  background: var(--bg-light);
  border: 1px solid var(--border-light);
  color: var(--text-light);
  padding: 0.75rem 1.5rem;
  border-radius: var(--radius-sm);
  font-size: 0.875rem;
  font-family: 'Inter', sans-serif;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 500;
}

.category-btn.active, .category-btn:hover {
  background: var(--accent-orange);
  color: var(--primary-light);
  border-color: var(--accent-orange);
}

/* Seção do lanche customizado */
.custom-burger-section {
  max-width: 800px;
  margin: 0 auto;
}

/* Grid de produtos */
.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1.5rem;
}

/* Card de produto */
.product-card {
  background: var(--bg-white);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border-light);
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: all 0.2s ease;
}

.product-card:hover {
  box-shadow: var(--shadow-md);
  transform: translateY(-2px);
}

/* Imagem do produto */
.product-image img {
  width: 120px;
  height: 120px;
  object-fit: cover;
  border-radius: 50%;
  margin-bottom: 1rem;
  border: 3px solid var(--border-light);
}

/* Informações do produto */
.product-info {
  width: 100%;
}

.product-name {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  color: var(--text-dark);
  font-family: 'Inter', sans-serif;
}

.product-description {
  font-size: 0.875rem;
  color: var(--text-light);
  margin-bottom: 1rem;
  line-height: 1.5;
}

/* Rodapé do card com preço e botão */
.product-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
}

.product-price {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--accent-orange);
  font-family: 'Inter', sans-serif;
}

.add-to-cart-btn {
  background: var(--accent-orange);
  color: var(--primary-light);
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: var(--radius-sm);
  font-family: 'Inter', sans-serif;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.875rem;
}

.add-to-cart-btn:hover {
  background: var(--accent-orange-dark);
  transform: translateY(-1px);
}

/* Responsividade para dispositivos móveis */
@media (max-width: 768px) {
  .menu {
    padding: 1rem;
  }
  
  .categories {
    gap: 0.5rem;
  }
  
  .category-btn {
    padding: 0.5rem 1rem;
    font-size: 0.8rem;
  }
  
  .products-grid {
    grid-template-columns: 1fr;
  }
  
  .product-footer {
    flex-direction: column;
    gap: 0.5rem;
  }
}
</style>