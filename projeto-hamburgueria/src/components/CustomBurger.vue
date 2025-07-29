<template>
  <!-- Card do lanche customizado -->
  <div class="custom-burger-card">
    <!-- Formulário para montar o lanche -->
    <form @submit.prevent="sendToKitchen">
      <!-- Grid de ingredientes -->
      <div class="ingredients-grid">
        <!-- Seção de pão -->
        <div class="custom-section">
          <label class="section-label">🍞 Pão:</label>
          <select v-model="selected.bread" class="custom-select">
            <option
              v-for="option in options.bread"
              :key="option.id"
              :value="option"
            >
              {{ option.name }} (+R$ {{ option.price.toFixed(2) }})
            </option>
          </select>
        </div>

        <!-- Seção de carne -->
        <div class="custom-section">
          <label class="section-label">🥩 Carne:</label>
          <select v-model="selected.meat" class="custom-select">
            <option
              v-for="option in options.meat"
              :key="option.id"
              :value="option"
            >
              {{ option.name }} (+R$ {{ option.price.toFixed(2) }})
            </option>
          </select>
        </div>

        <!-- Seção de queijo -->
        <div class="custom-section">
          <label class="section-label">🧀 Queijo:</label>
          <select v-model="selected.cheese" class="custom-select">
            <option
              v-for="option in options.cheese"
              :key="option.id"
              :value="option"
            >
              {{ option.name }} (+R$ {{ option.price.toFixed(2) }})
            </option>
          </select>
        </div>

        <!-- Seção de molho -->
        <div class="custom-section">
          <label class="section-label">🥫 Molho:</label>
          <select v-model="selected.sauce" class="custom-select">
            <option
              v-for="option in options.sauce"
              :key="option.id"
              :value="option"
            >
              {{ option.name }} (+R$ {{ option.price.toFixed(2) }})
            </option>
          </select>
        </div>
      </div>

      <!-- Botão para enviar para a cozinha -->
      <button class="send-to-kitchen-btn" type="submit">
        👨‍🍳 Enviar para Cozinha
      </button>
    </form>
  </div>
</template>

<script setup>
// Importa funções reativas do Vue.js
import { ref, computed, onMounted } from "vue";
// Importa configuração da API
import API_URL from "../api";

// Define os eventos que o componente pode emitir
const emit = defineEmits(["add-to-cart"]);

// ===== ESTADO REATIVO =====
// Array reativo para armazenar todos os ingredientes
const ingredients = ref([]);
// Objeto reativo para organizar ingredientes por categoria
const options = ref({
  bread: [],
  meat: [],
  cheese: [],
  sauce: [],
});

// Objeto reativo para armazenar ingredientes selecionados
const selected = ref({
  bread: null,
  meat: null,
  cheese: null,
  sauce: null,
});

// ===== HOOK DE CICLO DE VIDA =====
// Buscar ingredientes da API quando o componente é montado
onMounted(async () => {
  try {
    // Buscar todos os ingredientes da API
    const res = await fetch(`${API_URL}/ingredients`);
    ingredients.value = await res.json();

    // Separar ingredientes por categoria
    options.value.bread = ingredients.value.filter(
      (i) => i.category === "pao" || i.category === "pão",
    );
    options.value.meat = ingredients.value.filter(
      (i) => i.category === "carne",
    );
    options.value.cheese = ingredients.value.filter(
      (i) => i.category === "queijo",
    );
    options.value.sauce = ingredients.value.filter(
      (i) => i.category === "molho" || i.category === "molhos",
    );

    // Selecionar primeiros itens por padrão
    if (options.value.bread.length > 0)
      selected.value.bread = options.value.bread[0];
    if (options.value.meat.length > 0)
      selected.value.meat = options.value.meat[0];
    if (options.value.cheese.length > 0)
      selected.value.cheese = options.value.cheese[0];
    if (options.value.sauce.length > 0)
      selected.value.sauce = options.value.sauce[0];
  } catch (error) {
    console.error("Erro ao carregar ingredientes:", error);
  }
});

// ===== COMPUTED PROPERTIES =====
// Propriedade computada que calcula o preço total do lanche
const totalPrice = computed(() => {
  let total = 0;
  if (selected.value.bread) total += selected.value.bread.price;
  if (selected.value.meat) total += selected.value.meat.price;
  if (selected.value.cheese) total += selected.value.cheese.price;
  if (selected.value.sauce) total += selected.value.sauce.price;
  return total;
});

// ===== FUNÇÕES =====
// Função para enviar o lanche customizado para a cozinha
function sendToKitchen() {
  // Validação: pelo menos pão e carne devem ser selecionados
  if (!selected.value.bread || !selected.value.meat) {
    alert("Selecione pelo menos pão e carne!");
    return;
  }

  // Criar objeto do lanche customizado
  const customBurger = {
    name: `Lanche Personalizado`,
    description: `${selected.value.bread.name}, ${selected.value.meat.name}, ${selected.value.cheese.name}, ${selected.value.sauce.name}`,
    price: totalPrice.value,
    image_url: "/brownie.jfif", // Usar uma imagem local como placeholder
    ingredients: JSON.stringify({
      bread: selected.value.bread,
      meat: selected.value.meat,
      cheese: selected.value.cheese,
      sauce: selected.value.sauce,
    }),
  };

  // Emitir evento para enviar para a cozinha
  emit("add-to-cart", customBurger);

  // Mostrar confirmação
  alert("🍔 Lanche enviado para a cozinha!");
}
</script>

<style scoped>
/* Card do lanche customizado */
.custom-burger-card {
  background: var(--bg-white);
  border-radius: var(--radius-lg);
  padding: 2rem;
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--border-light);
}

/* Grid de ingredientes */
.ingredients-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

/* Seção de cada ingrediente */
.custom-section {
  margin-bottom: 1rem;
}

/* Label da seção */
.section-label {
  display: block;
  font-family: "Inter", sans-serif;
  font-weight: 600;
  color: var(--text-dark);
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
}

/* Select de ingredientes */
.custom-select {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid var(--border-light);
  border-radius: var(--radius-sm);
  font-family: "Inter", sans-serif;
  font-size: 0.875rem;
  background: var(--bg-white);
  color: var(--text-dark);
  transition: border-color 0.2s ease;
}

.custom-select:focus {
  outline: none;
  border-color: var(--accent-orange);
  box-shadow: 0 0 0 2px rgba(255, 107, 53, 0.1);
}

/* Botão para enviar para a cozinha */
.send-to-kitchen-btn {
  width: 100%;
  background: var(--accent-green);
  color: var(--primary-light);
  border: none;
  padding: 1rem;
  border-radius: var(--radius-sm);
  font-family: "Inter", sans-serif;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 1rem;
}

.send-to-kitchen-btn:hover {
  background: var(--accent-green-dark);
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

/* Responsividade para dispositivos móveis */
@media (max-width: 768px) {
  .custom-burger-card {
    padding: 1rem;
  }

  .ingredients-grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }

  .checkbox-group {
    grid-template-columns: 1fr;
  }
}
</style>
