import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Header from '../Header.vue'

describe('Header.vue', () => {
  const wrapper = mount(Header, {
    props: {
      activeOrders: 2
    }
  })

  it('renderiza o cabeçalho corretamente', () => {
    expect(wrapper.find('.header').exists()).toBe(true)
    expect(wrapper.find('.logo').exists()).toBe(true)
    expect(wrapper.find('.kitchen-info').exists()).toBe(true)
  })

  it('exibe o título da aplicação', () => {
    const title = wrapper.find('.logo-text h1')
    expect(title.exists()).toBe(true)
    expect(title.text()).toBe('BurgerApp')
  })

  it('tem navegação funcional', () => {
    const kitchenButton = wrapper.find('.kitchen-button')
    expect(kitchenButton.exists()).toBe(true)
    expect(kitchenButton.text()).toContain('Cozinha')
  })

  it('aplica estilos corretos', () => {
    const header = wrapper.find('.header')
    expect(header.exists()).toBe(true)
  })

  it('é responsivo', () => {
    const headerContent = wrapper.find('.header-content')
    expect(headerContent.exists()).toBe(true)
  })

  it('exibe contador de pedidos ativos', () => {
    const ordersCount = wrapper.find('.orders-count')
    expect(ordersCount.exists()).toBe(true)
    expect(ordersCount.text()).toBe('2')
  })

  it('emite evento ao clicar no botão da cozinha', async () => {
    const kitchenButton = wrapper.find('.kitchen-button')
    await kitchenButton.trigger('click')
    expect(wrapper.emitted('toggle-kitchen')).toBeTruthy()
  })
})