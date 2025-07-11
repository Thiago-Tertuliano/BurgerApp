// Importa utilitários do Node.js para manipular URLs
import { fileURLToPath, URL } from 'node:url'

// Importa função para definir configuração do Vite
import { defineConfig } from 'vite'
// Importa plugin do Vue.js para o Vite
import vue from '@vitejs/plugin-vue'
// Importa plugin para Vue DevTools
import vueDevTools from 'vite-plugin-vue-devtools'

// Configuração do Vite - https://vite.dev/config/
export default defineConfig({
  // Array de plugins que serão usados
  plugins: [
    vue(), // Plugin principal do Vue.js
    vueDevTools(), // Plugin para ferramentas de desenvolvimento
  ],
  // Configurações de resolução de módulos
  resolve: {
    // Aliases para facilitar imports
    alias: {
      // '@' aponta para a pasta src - permite imports como @/components/Header.vue
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
