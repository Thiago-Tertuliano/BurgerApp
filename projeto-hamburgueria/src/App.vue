<script setup>
// Importa funções reativas do Vue.js
import { ref, computed, onMounted } from 'vue'
// Importa componentes da aplicação
import Header from './components/Header.vue'
import Menu from './components/Menu.vue'
import Kitchen from './components/Kitchen.vue'
// Importa configuração da API
import API_URL from './api'

// ===== ESTADO GLOBAL DA APLICAÇÃO =====
// Array reativo que armazena todos os pedidos
const orders = ref([])
// Estado reativo que controla se a cozinha está aberta ou fechada
const isKitchenOpen = ref(false)
// Estado reativo que controla qual tela está ativa (menu ou cozinha)
const currentStep = ref('menu') // menu, kitchen

// ===== FUNÇÕES DE GESTÃO DE PEDIDOS =====
// Função assíncrona para adicionar novos pedidos
const addToOrders = async (item) => {
  try {
    // Determinar product_id baseado no tipo de item
    let productId = 1 // ID padrão para produtos customizados
    
    // Se o item tem ID numérico, usa ele
    if (item.id && typeof item.id === 'number') {
      productId = item.id
    } else if (item.id && typeof item.id === 'string' && item.id.startsWith('custom-')) {
      // Para lanches customizados, usar ID 1
      productId = 1
    }
    
    // Preparar dados do pedido para enviar ao backend
    const orderData = {
      customer_name: 'Cliente',
      table_number: 1,
      items: [{
        product_id: productId,
        ingredients: item.ingredients || '',
        quantity: 1,
        notes: ''
      }],
      notes: ''
    }

    // Enviar pedido para o backend via API
    const response = await fetch(`${API_URL}/orders`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(orderData)
    })

    if (response.ok) {
      const result = await response.json()
      console.log('Pedido criado:', result)
      
      // Adicionar ao estado local para exibição imediata
      const newOrder = {
        id: result.order_id,
        name: item.name,
        description: item.description,
        price: item.price,
        status: 'preparando',
        time: new Date().toLocaleTimeString('pt-BR', { 
          hour: '2-digit', 
          minute: '2-digit' 
        }),
        createdAt: new Date()
      }
      orders.value.push(newOrder)
    } else {
      const errorData = await response.json()
      console.error('Erro ao criar pedido:', errorData)
    }
  } catch (error) {
    console.error('Erro ao enviar pedido:', error)
  }
}

// Função assíncrona para atualizar status dos pedidos
const updateOrder = async (updateData) => {
  try {
    // Converter status do frontend para o formato do backend
    let backendStatus = updateData.status
    if (updateData.status === 'pronto') {
      backendStatus = 'ready'
    } else if (updateData.status === 'entregue') {
      backendStatus = 'delivered'
    }
    
    // Atualizar status no backend via API
    const response = await fetch(`${API_URL}/orders/${updateData.id}/status`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ status: backendStatus })
    })

    if (response.ok) {
      // Atualizar estado local após sucesso no backend
      const orderIndex = orders.value.findIndex(order => order.id === updateData.id)
      if (orderIndex !== -1) {
        orders.value[orderIndex].status = updateData.status
      }
    } else {
      const errorData = await response.json()
      console.error('Erro ao atualizar status:', errorData)
    }
  } catch (error) {
    console.error('Erro ao atualizar pedido:', error)
  }
}

// ===== COMPUTED PROPERTIES =====
// Propriedade computada que conta pedidos ativos (em preparo)
const activeOrders = computed(() => {
  return orders.value.filter(order => order.status === 'preparando').length
})

// ===== FUNÇÕES DE CONTROLE DA COZINHA =====
// Função para alternar visibilidade da cozinha
const toggleKitchen = () => {
  isKitchenOpen.value = !isKitchenOpen.value
}

// Função para fechar a cozinha
const closeKitchen = () => {
  console.log('Fechando cozinha')
  isKitchenOpen.value = false
}

// ===== HOOK DE CICLO DE VIDA =====
// Carregar pedidos existentes ao iniciar a aplicação
onMounted(async () => {
  try {
    // Buscar pedidos em preparo do backend
    const response = await fetch(`${API_URL}/orders?status=preparing`)
    if (response.ok) {
      const backendOrders = await response.json()
      // Converter formato do backend para o formato do frontend
      orders.value = backendOrders.map(order => {
        // Converter status do backend para o formato do frontend
        let frontendStatus = order.status
        if (order.status === 'ready') {
          frontendStatus = 'pronto'
        } else if (order.status === 'delivered') {
          frontendStatus = 'entregue'
        } else if (order.status === 'preparing') {
          frontendStatus = 'preparando'
        }
        
        return {
          id: order.id,
          name: `Pedido #${order.id}`,
          description: `Mesa ${order.table_number} - ${order.customer_name}`,
          price: order.total_amount,
          status: frontendStatus,
          time: new Date(order.created_at).toLocaleTimeString('pt-BR', { 
            hour: '2-digit', 
            minute: '2-digit' 
          }),
          createdAt: new Date(order.created_at)
        }
      })
    }
  } catch (error) {
    console.error('Erro ao carregar pedidos:', error)
  }
})
</script>

<template>
  <!-- Container principal da aplicação -->
  <div id="app">
    <!-- Componente Header com props e eventos -->
    <Header 
      :active-orders="activeOrders"
      @toggle-kitchen="toggleKitchen"
    />

    <!-- Conteúdo principal da aplicação -->
    <main class="main-content">
      <!-- Seção do menu - só mostra quando currentStep é 'menu' -->
      <div v-if="currentStep === 'menu'" class="menu-section">
        <Menu @add-to-cart="addToOrders" />
      </div>
    </main>

    <!-- Cozinha lateral - só mostra quando isKitchenOpen é true -->
    <Kitchen 
      v-if="isKitchenOpen"
      :orders="orders"
      @close="closeKitchen"
      @update-order="updateOrder"
    />
  </div>
</template>

<style>
/* Reset CSS básico */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* Estilos do body */
body {
  font-family: 'Inter', sans-serif;
  background: var(--bg-light);
  color: var(--text-dark);
  min-height: 100vh;
}

/* Container principal da aplicação */
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* Conteúdo principal */
.main-content {
  flex: 1;
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
}

/* Seção do menu */
.menu-section {
  width: 100%;
}

/* Responsividade para dispositivos móveis */
@media (max-width: 768px) {
  .main-content {
    padding: 10px;
  }
}
</style>
