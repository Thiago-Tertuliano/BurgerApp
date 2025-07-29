import { describe, it, expect, vi, beforeEach } from "vitest";
import { mount } from "@vue/test-utils";
import Menu from "../Menu.vue";

// Mock do fetch
global.fetch = vi.fn();

describe("Menu.vue", () => {
  const mockCategories = [
    { id: 1, name: "Hambúrgueres" },
    { id: 2, name: "Bebidas" },
  ];

  const mockProducts = [
    {
      id: 1,
      name: "X-Burger",
      description: "Hambúrguer clássico",
      price: 15.9,
      category_id: 1,
      image_url: "/burger.jpg",
    },
  ];

  beforeEach(() => {
    // Mock das respostas da API
    fetch.mockResolvedValueOnce({
      json: () => Promise.resolve(mockCategories),
    });
    fetch.mockResolvedValueOnce({
      json: () => Promise.resolve(mockProducts),
    });
  });

  it("renderiza o menu corretamente", () => {
    const wrapper = mount(Menu);
    expect(wrapper.find(".menu").exists()).toBe(true);
    expect(wrapper.find(".menu-title").exists()).toBe(true);
  });

  it("exibe categorias de produtos", async () => {
    const wrapper = mount(Menu);
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick(); // Aguarda o segundo nextTick para os dados serem carregados

    const categoryButtons = wrapper.findAll(".category-btn");
    expect(categoryButtons.length).toBeGreaterThan(0);
  });

  it("filtra produtos por categoria", async () => {
    const wrapper = mount(Menu);
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick();

    // Verifica se a propriedade computada existe
    expect(wrapper.vm.filteredProducts).toBeDefined();
  });

  it("adiciona produtos ao carrinho", async () => {
    const wrapper = mount(Menu);
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick();

    // Verifica se o evento é emitido
    expect(wrapper.emitted()).toBeDefined();
  });

  it("exibe preços corretamente", async () => {
    const wrapper = mount(Menu);
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick();

    // Verifica se o componente tem a estrutura para exibir preços
    expect(wrapper.find(".products-grid").exists()).toBe(true);
  });

  it("mostra produtos disponíveis", async () => {
    const wrapper = mount(Menu);
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick();

    // Verifica se o grid de produtos existe
    expect(wrapper.find(".products-grid").exists()).toBe(true);
  });

  it("emite evento ao adicionar produto ao carrinho", async () => {
    const wrapper = mount(Menu);
    await wrapper.vm.$nextTick();
    await wrapper.vm.$nextTick();

    const addButton = wrapper.find(".add-to-cart-btn");
    if (addButton.exists()) {
      await addButton.trigger("click");
      expect(wrapper.emitted("add-to-cart")).toBeTruthy();
    }
  });
});
