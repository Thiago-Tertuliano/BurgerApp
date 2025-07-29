<template>
  <!-- Overlay da cozinha - fundo escuro que cobre toda a tela -->
  <div class="kitchen-overlay" @click="handleOverlayClick">
    <!-- Sidebar da cozinha - painel lateral -->
    <div class="kitchen-sidebar" @click.stop>
      <!-- Cabeçalho da cozinha -->
      <div class="kitchen-header">
        <h2>👨‍🍳 Cozinha</h2>
        <button class="close-btn" @click="handleClose" type="button">✕</button>
      </div>

      <!-- Conteúdo da cozinha -->
      <div class="kitchen-content">
        <!-- Estado vazio - quando não há pedidos -->
        <div v-if="orders.length === 0" class="empty-kitchen">
          <div class="empty-icon">🍽️</div>
          <p>Nenhum pedido em preparo</p>
          <p>Os pedidos aparecerão aqui quando forem feitos!</p>
        </div>

        <!-- Lista de pedidos -->
        <div v-else class="orders-list">
          <!-- Card de cada pedido -->
          <div v-for="order in orders" :key="order.id" class="order-card">
            <!-- Cabeçalho do pedido -->
            <div class="order-header">
              <h3>{{ order.name }}</h3>
              <span class="order-time">{{ order.time }}</span>
            </div>

            <!-- Descrição do pedido -->
            <div class="order-description">
              {{ order.description }}
            </div>

            <!-- Status do pedido -->
            <div class="order-status">
              <!-- Indicador de status com ícone e texto -->
              <div class="status-indicator" :class="order.status">
                <span class="status-icon">{{
                  getStatusIcon(order.status)
                }}</span>
                <span class="status-text">{{
                  getStatusText(order.status)
                }}</span>
              </div>

              <!-- Barra de progresso -->
              <div class="progress-bar">
                <div
                  class="progress-fill"
                  :style="{ width: getProgressWidth(order.status) }"
                ></div>
              </div>
            </div>

            <!-- Ações do pedido -->
            <div class="order-actions">
              <!-- Botão para marcar como pronto - só aparece se status é 'preparando' -->
              <button
                v-if="order.status === 'preparando'"
                class="action-btn ready-btn"
                @click="markAsReady(order.id)"
              >
                Marcar como Pronto
              </button>
              <!-- Botão para entregar - só aparece se status é 'pronto' -->
              <button
                v-if="order.status === 'pronto'"
                class="action-btn delivered-btn"
                @click="markAsDelivered(order.id)"
              >
                Entregar
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Importa funções reativas do Vue.js
import { onMounted, onUnmounted } from "vue";
// Importa configuração da API
import API_URL from "../api";

// ===== PROPS =====
// Define as props que o componente recebe do componente pai
defineProps({
  // Array de pedidos
  orders: {
    type: Array,
    default: () => [],
  },
});

// ===== EMITS =====
// Define os eventos que o componente pode emitir
const emit = defineEmits(["close", "update-order"]);

// ===== FUNÇÕES DE CONTROLE =====
// Função para fechar a cozinha
function handleClose() {
  console.log("Botão de fechar clicado");
  emit("close");
}

// Função para fechar ao clicar no overlay
function handleOverlayClick() {
  console.log("Overlay clicado");
  emit("close");
}

// ===== FUNÇÕES DE STATUS =====
// Função para obter ícone baseado no status
function getStatusIcon(status) {
  const icons = {
    preparando: "👨‍🍳",
    pronto: "✅",
    entregue: "🛵",
  };
  return icons[status] || "⏳";
}

// Função para obter texto baseado no status
function getStatusText(status) {
  const texts = {
    preparando: "Preparando",
    pronto: "Pronto para entrega",
    entregue: "Entregue",
  };
  return texts[status] || "Aguardando";
}

// Função para obter largura da barra de progresso baseada no status
function getProgressWidth(status) {
  const progress = {
    preparando: "60%",
    pronto: "100%",
    entregue: "100%",
  };
  return progress[status] || "0%";
}

// ===== FUNÇÕES DE ATUALIZAÇÃO =====
// Função para marcar pedido como pronto
async function markAsReady(orderId) {
  try {
    // Atualizar status no backend
    const response = await fetch(`${API_URL}/orders/${orderId}/status`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ status: "ready" }),
    });

    if (response.ok) {
      // Emitir evento para atualizar no componente pai
      emit("update-order", { id: orderId, status: "pronto" });
    } else {
      console.error("Erro ao marcar como pronto");
    }
  } catch (error) {
    console.error("Erro ao atualizar status:", error);
  }
}

// Função para marcar pedido como entregue
async function markAsDelivered(orderId) {
  try {
    // Atualizar status no backend
    const response = await fetch(`${API_URL}/orders/${orderId}/status`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ status: "delivered" }),
    });

    if (response.ok) {
      // Emitir evento para atualizar no componente pai
      emit("update-order", { id: orderId, status: "entregue" });
    } else {
      console.error("Erro ao marcar como entregue");
    }
  } catch (error) {
    console.error("Erro ao atualizar status:", error);
  }
}

// ===== HOOKS DE CICLO DE VIDA =====
// Atualizar pedidos periodicamente quando componente é montado
onMounted(() => {
  // Função para atualizar pedidos
  const updateOrders = async () => {
    try {
      // Buscar pedidos em preparo do backend
      const response = await fetch(`${API_URL}/orders?status=preparing`);
      if (response.ok) {
        const backendOrders = await response.json();
        // Converter formato do backend para o formato do frontend
        const convertedOrders = backendOrders.map((order) => ({
          id: order.id,
          name: `Pedido #${order.id}`,
          description: `Mesa ${order.table_number} - ${order.customer_name}`,
          price: order.total_amount,
          status: order.status,
          time: new Date(order.created_at).toLocaleTimeString("pt-BR", {
            hour: "2-digit",
            minute: "2-digit",
          }),
          createdAt: new Date(order.created_at),
        }));

        // Emitir evento para atualizar pedidos no App.vue
        emit("update-orders", convertedOrders);
      }
    } catch (error) {
      console.error("Erro ao atualizar pedidos:", error);
    }
  };

  // Atualizar a cada 10 segundos
  const interval = setInterval(updateOrders, 10000);

  // Limpar intervalo quando componente for desmontado
  onUnmounted(() => {
    clearInterval(interval);
  });
});
</script>

<style scoped>
/* Overlay da cozinha - fundo escuro */
.kitchen-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(5px);
  z-index: 1000;
  display: flex;
  justify-content: flex-end;
  animation: slideIn 0.3s ease-out;
  cursor: pointer;
}

/* Sidebar da cozinha */
.kitchen-sidebar {
  width: 450px;
  max-width: 90vw;
  height: 100vh;
  background: var(--bg-white);
  border-left: 1px solid var(--border-light);
  box-shadow: var(--shadow-lg);
  display: flex;
  flex-direction: column;
  animation: slideIn 0.3s ease-out;
  pointer-events: auto;
}

/* Cabeçalho da cozinha */
.kitchen-header {
  background: var(--accent-orange);
  color: var(--primary-light);
  padding: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid var(--border-light);
}

.kitchen-header h2 {
  font-family: "Inter", sans-serif;
  font-size: 1.25rem;
  font-weight: 700;
  margin: 0;
}

/* Botão de fechar */
.close-btn {
  background: none;
  border: none;
  color: var(--primary-light);
  font-size: 1.25rem;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 50%;
  transition: background 0.2s;
  position: relative;
  z-index: 1001;
  min-width: 40px;
  min-height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  pointer-events: auto;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.2);
}

.close-btn:active {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(0.95);
}

/* Conteúdo da cozinha */
.kitchen-content {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
}

/* Estado vazio da cozinha */
.empty-kitchen {
  text-align: center;
  padding: 3rem 1rem;
  color: var(--text-light);
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  opacity: 0.5;
}

/* Lista de pedidos */
.orders-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* Card de pedido */
.order-card {
  background: var(--bg-white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-md);
  padding: 1.5rem;
  box-shadow: var(--shadow-sm);
  transition: all 0.2s ease;
}

/* Cabeçalho do pedido */
.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.order-header h3 {
  font-family: "Inter", sans-serif;
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--text-dark);
  margin: 0;
}

.order-time {
  font-family: "Inter", sans-serif;
  font-size: 0.875rem;
  color: var(--text-light);
}

/* Descrição do pedido */
.order-description {
  font-family: "Inter", sans-serif;
  font-size: 0.875rem;
  color: var(--text-light);
  margin-bottom: 1rem;
  line-height: 1.5;
}

/* Status do pedido */
.order-status {
  margin-bottom: 1rem;
}

/* Indicador de status */
.status-indicator {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.status-icon {
  font-size: 1.25rem;
}

.status-text {
  font-family: "Inter", sans-serif;
  font-size: 0.875rem;
  font-weight: 600;
}

/* Cores do status */
.status-indicator.preparando .status-text {
  color: var(--accent-orange);
}

.status-indicator.pronto .status-text {
  color: var(--accent-green);
}

.status-indicator.entregue .status-text {
  color: var(--text-light);
}

/* Barra de progresso */
.progress-bar {
  width: 100%;
  height: 4px;
  background: var(--border-light);
  border-radius: 2px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: var(--accent-orange);
  transition: width 0.3s ease;
}

/* Ações do pedido */
.order-actions {
  display: flex;
  gap: 0.5rem;
}

/* Botões de ação */
.action-btn {
  flex: 1;
  padding: 0.75rem;
  border: none;
  border-radius: var(--radius-sm);
  font-family: "Inter", sans-serif;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 0.875rem;
}

/* Botão "Marcar como Pronto" */
.ready-btn {
  background: var(--accent-green);
  color: var(--primary-light);
}

.ready-btn:hover {
  background: var(--accent-green-dark);
}

/* Botão "Entregar" */
.delivered-btn {
  background: var(--accent-blue);
  color: var(--primary-light);
}

.delivered-btn:hover {
  background: var(--accent-blue-dark);
}

/* Animação de entrada */
@keyframes slideIn {
  from {
    transform: translateX(100%);
  }
  to {
    transform: translateX(0);
  }
}

/* Responsividade para dispositivos móveis */
@media (max-width: 768px) {
  .kitchen-sidebar {
    width: 100vw;
  }

  .order-card {
    padding: 1rem;
  }

  .order-actions {
    flex-direction: column;
  }
}
</style>
