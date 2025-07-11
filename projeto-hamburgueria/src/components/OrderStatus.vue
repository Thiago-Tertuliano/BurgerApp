<template>
  <div class="order-status-card">
    <h2 class="status-title">Status do Pedido</h2>
    <div class="status-steps">
      <div v-for="(step, idx) in steps" :key="step.id" class="status-step" :class="{active: idx <= currentStep}">
        <div class="step-icon">{{ step.icon }}</div>
        <div class="step-label">{{ step.label }}</div>
      </div>
    </div>
    <div class="status-message">
      <span>{{ steps[currentStep].message }}</span>
    </div>
    <div class="status-actions">
      <button v-if="currentStep < steps.length - 1" class="modern-button" @click="nextStep">
        Avançar
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const steps = [
  { id: 'preparando', label: 'Preparando', icon: '👨‍🍳', message: 'Seu pedido está sendo preparado!' },
  { id: 'pronto', label: 'Pronto', icon: '✅', message: 'Seu pedido está pronto para retirada!' },
  { id: 'entregue', label: 'Entregue', icon: '🛵', message: 'Pedido entregue! Bom apetite!' }
]
const currentStep = ref(0)
function nextStep() {
  if (currentStep.value < steps.length - 1) currentStep.value++
}
</script>

<style scoped>
.order-status-card {
  background: var(--bg-white);
  border: 1px solid var(--border-light);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-md);
  padding: 2rem;
  max-width: 500px;
  margin: 2rem auto;
  text-align: center;
}

.status-title {
  font-family: 'Inter', sans-serif;
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-dark);
  margin-bottom: 2rem;
}

.status-steps {
  display: flex;
  justify-content: space-between;
  margin: 2rem 0 1rem 0;
  position: relative;
}

.status-steps::before {
  content: '';
  position: absolute;
  top: 1.5rem;
  left: 0;
  right: 0;
  height: 2px;
  background: var(--border-light);
  z-index: 1;
}

.status-step {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  z-index: 2;
  opacity: 0.5;
  transition: opacity 0.3s ease;
}

.status-step.active {
  opacity: 1;
}

.step-icon {
  font-size: 2rem;
  margin-bottom: 0.5rem;
  background: var(--bg-white);
  padding: 0.5rem;
  border-radius: 50%;
  border: 2px solid var(--border-light);
  transition: all 0.3s ease;
}

.status-step.active .step-icon {
  border-color: var(--accent-orange);
  background: var(--accent-orange);
  color: var(--primary-light);
}

.step-label {
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--text-dark);
  font-family: 'Inter', sans-serif;
}

.status-message {
  font-size: 1.125rem;
  color: var(--text-dark);
  margin-bottom: 1.5rem;
  font-family: 'Inter', sans-serif;
  font-weight: 500;
}

.status-actions {
  margin-top: 1rem;
}

@media (max-width: 768px) {
  .order-status-card {
    padding: 1.5rem;
    margin: 1rem;
  }
  
  .status-steps {
    flex-direction: column;
    gap: 1rem;
  }
  
  .status-steps::before {
    display: none;
  }
}
</style> 