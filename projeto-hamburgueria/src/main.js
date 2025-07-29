// Importa os estilos CSS globais da aplicação
import "./assets/main.css";

// Importa função do Vue.js para criar a aplicação
import { createApp } from "vue";
// Importa o componente raiz da aplicação
import App from "./App.vue";

// Cria a instância da aplicação Vue e monta no elemento #app
// createApp() retorna uma instância da aplicação
// .mount('#app') conecta a aplicação ao elemento HTML com id="app"
createApp(App).mount("#app");
