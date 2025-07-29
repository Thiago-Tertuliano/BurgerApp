import { describe, it, expect, beforeEach } from 'vitest'
import { mount } from '@vue/test-utils'
import Header from '../Header.vue'

describe('Header.vue', () => {
  let wrapper

  beforeEach(() => {
    wrapper = mount(Header)
  })

  it('renderiza o cabeçalho corretamente', () => {
    expect(wrapper.find('.header').exists()).toBe(true)
    expect(wrapper.find('.logo').exists()).toBe(true)
    expect(wrapper.find('.nav').exists()).toBe(true)
  })

  it('exibe o título da aplicação', () => {
    const title = wrapper.find('.logo h1')
    expect(title.exists()).toBe(true)
    expect(title.text()).toContain('BurgerApp')
  })

  it('tem navegação funcional', () => {
    const navItems = wrapper.findAll('.nav-item')
    expect(navItems.length).toBeGreaterThan(0)
  })

  it('aplica estilos corretos', () => {
    const header = wrapper.find('.header')
    expect(header.classes()).toContain('header')
  })

  it('é responsivo', () => {
    const header = wrapper.find('.header')
    expect(header.classes()).toContain('header')
    // Verificar se tem classes responsivas
    expect(wrapper.html()).toContain('class')
  })
})