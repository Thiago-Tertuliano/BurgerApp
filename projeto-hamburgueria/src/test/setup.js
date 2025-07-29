import { config } from '@vue/test-utils'

// Configuração global para testes
config.global.stubs = {
  // Stubs para componentes externos se necessário
}

// Mock para localStorage
Object.defineProperty(window, 'localStorage', {
  value: {
    getItem: vi.fn(),
    setItem: vi.fn(),
    removeItem: vi.fn(),
    clear: vi.fn()
  },
  writable: true
})

// Mock para fetch
global.fetch = vi.fn()

// Mock para console para evitar logs nos testes
global.console = {
  ...console,
  log: vi.fn(),
  warn: vi.fn(),
  error: vi.fn()
}