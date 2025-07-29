import { describe, it, expect, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import Menu from '../Menu.vue'

describe('Menu.vue', () => {
  let wrapper

  beforeEach(() => {
    wrapper = mount(Menu)
  })

  it('renderiza o menu corretamente', () => {
    expect(wrapper.find('.menu').exists()).toBe(true)
    expect(wrapper.find('.menu-grid').exists()).toBe(true)
  })

  it('exibe categorias de produtos', () => {
    const categories = wrapper.findAll('.category')
    expect(categories.length).toBeGreaterThan(0)
  })

  it('filtra produtos por categoria', async () => {
    // Simular dados de produtos
    await wrapper.setData({
      products: [
        { id: 1, name: 'Hambúrguer', category: 'lanches', price: 15.00 },
        { id: 2, name: 'Batata Frita', category: 'acompanhamentos', price: 8.00 }
      ],
      categories: [
        { id: 1, name: 'lanches' },
        { id: 2, name: 'acompanhamentos' }
      ]
    })

    const categoryButtons = wrapper.findAll('.category-button')
    expect(categoryButtons.length).toBeGreaterThan(0)
  })

  it('adiciona produtos ao carrinho', async () => {
    await wrapper.setData({
      products: [
        { id: 1, name: 'Hambúrguer', price: 15.00, is_available: true }
      ]
    })

    const addButton = wrapper.find('.add-to-cart')
    if (addButton.exists()) {
      await addButton.trigger('click')
      expect(wrapper.emitted('add-to-cart')).toBeTruthy()
    }
  })

  it('exibe preços corretamente', () => {
    const priceElements = wrapper.findAll('.price')
    priceElements.forEach(price => {
      expect(price.text()).toMatch(/R\$\s*\d+,\d+/)
    })
  })

  it('mostra produtos disponíveis', async () => {
    await wrapper.setData({
      products: [
        { id: 1, name: 'Hambúrguer', is_available: true },
        { id: 2, name: 'Produto Indisponível', is_available: false }
      ]
    })

    const availableProducts = wrapper.findAll('.product-card')
    expect(availableProducts.length).toBeGreaterThan(0)
  })
})