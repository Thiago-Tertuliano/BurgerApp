# Segurança do Projeto BurgerApp

## Configuração de Segurança

### GitHub Actions Security Workflow

O projeto possui um workflow dedicado para segurança (`security.yml`) que executa:

1. **Scan de Vulnerabilidades** (Trivy)
   - Scan do sistema de arquivos
   - Foco em vulnerabilidades CRÍTICAS, ALTAS e MÉDIAS
   - Execução em push, pull request e diariamente

2. **Scan de Dependências** (Trivy)
   - Análise de dependências Node.js
   - Análise de dependências Go
   - Detecção de vulnerabilidades conhecidas

### Permissões Configuradas

O workflow de segurança possui as seguintes permissões:

```yaml
permissions:
  contents: read
  security-events: write
  actions: read
  pull-requests: read
```

### Resolução de Problemas

#### Erro: "Resource not accessible by integration"

**Causa**: O GitHub Actions não tem permissão para acessar a API de Code Scanning.

**Solução**: 
1. Adicionar permissões explícitas no workflow
2. Garantir que `security-events: write` está configurado
3. Verificar se não é um fork (forks têm limitações de segurança)

#### Configuração Recomendada

```yaml
# No workflow
permissions:
  contents: read
  security-events: write
  actions: read
  pull-requests: read

# No job específico
jobs:
  security-scan:
    permissions:
      contents: read
      security-events: write
      actions: read
```

### Integração com GitHub Security

Os resultados dos scans são enviados para:
- **GitHub Security Tab**: Visualização de vulnerabilidades
- **Code Scanning**: Análise detalhada de problemas de segurança
- **Dependabot**: Atualizações automáticas de dependências

### Monitoramento

- **Scans Diários**: Executados automaticamente às 2:00 UTC
- **Scans em PR**: Executados em cada pull request
- **Scans em Push**: Executados em pushes para main/develop

### Ferramentas Utilizadas

- **Trivy**: Scanner de vulnerabilidades
- **SARIF**: Formato de relatório de segurança
- **GitHub CodeQL**: Upload de resultados para análise

### Próximos Passos

1. Configurar Dependabot para atualizações automáticas
2. Implementar CodeQL para análise de código
3. Configurar alertas de segurança
4. Implementar políticas de branch protection 