<template>
  <!-- Cabeçalho da aplicação -->
  <header class="header">
    <div class="header-content">
      <!-- Logo e informações da marca -->
      <div class="logo">
        <div class="logo-icon">🍔</div>
        <div class="logo-text">
          <h1>BurgerApp</h1>
          <p>Os melhores hambúrgueres da cidade</p>
        </div>
      </div>
      
      <!-- Informações da cozinha e botão de acesso -->
      <div class="kitchen-info">
        <!-- Botão para abrir/fechar a cozinha -->
        <button class="kitchen-button" @click="$emit('toggle-kitchen')">
          <span class="kitchen-icon">👨‍🍳</span>
          <span class="kitchen-text">Cozinha</span>
          <!-- Badge com número de pedidos ativos -->
          <span v-if="activeOrders > 0" class="orders-count">{{ activeOrders }}</span>
        </button>
        <!-- Status da cozinha - mostra quantos pedidos estão em preparo -->
        <div v-if="activeOrders > 0" class="kitchen-status">
          {{ activeOrders }} pedido{{ activeOrders > 1 ? 's' : '' }} em preparo
        </div>
      </div>
    </div>
  </header>
</template>

<script setup>
// Define as props que o componente recebe do componente pai
defineProps({
  // Número de pedidos ativos (em preparo)
  activeOrders: {
    type: Number,
    default: 0
  }
})

// Define os eventos que o componente pode emitir
defineEmits(['toggle-kitchen'])
</script>

<style scoped>
/* Estilos do cabeçalho */
.header {
  background: var(--bg-white);
  border-bottom: 1px solid var(--border-light);
  box-shadow: var(--shadow-sm);
  padding: 1rem 0;
}

/* Container do conteúdo do cabeçalho */
.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* Estilos do logo */
.logo {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.logo-icon {
  font-size: 2.5rem;
  filter: drop-shadow(0 2px 4px rgba(0,0,0,0.1));
}

.logo-text h1 {
  font-family: 'Inter', sans-serif;
  font-size: 1.875rem;
  font-weight: 700;
  color: var(--text-dark);
  margin: 0;
  letter-spacing: -0.025em;
}

.logo-text p {
  font-size: 0.875rem;
  color: var(--text-light);
  margin: 0;
  font-weight: 500;
}

/* Estilos da seção da cozinha */
.kitchen-info {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* Botão da cozinha */
.kitchen-button {
  background: var(--accent-orange);
  color: var(--primary-light);
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: var(--radius-sm);
  cursor: pointer;
  font-size: 0.875rem;
  font-weight: 600;
  font-family: 'Inter', sans-serif;
  transition: all 0.2s ease;
  position: relative;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.kitchen-button:hover {
  background: var(--accent-gold);
  transform: translateY(-1px);
  box-shadow: var(--shadow-md);
}

/* Badge com contador de pedidos */
.orders-count {
  position: absolute;
  top: -8px;
  right: -8px;
  background: #ef4444;
  color: var(--primary-light);
  border-radius: 50%;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.75rem;
  font-weight: 600;
  border: 2px solid var(--primary-light);
}

/* Status da cozinha */
.kitchen-status {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--text-light);
  font-family: 'Inter', sans-serif;
}

/* Responsividade para dispositivos móveis */
@media (max-width: 768px) {
  .header-content {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }
  
  .logo {
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .logo-text h1 {
    font-size: 1.5rem;
  }
  
  .kitchen-button {
    padding: 0.6rem 1.2rem;
    font-size: 0.8rem;
  }
}
</style> 