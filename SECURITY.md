# Segurança do Projeto BurgerApp

## Configuração de Segurança

### GitHub Actions Security Workflows

O projeto possui workflows dedicados para segurança:

#### 1. **Security Scan (Principal)** (`security.yml`)
- Scan completo do sistema de arquivos
- Análise de dependências Node.js e Go
- Foco em vulnerabilidades CRÍTICAS, ALTAS e MÉDIAS
- Execução em push, pull request e diariamente

#### 2. **Security Scan (Simplificado)** (`security-simple.yml`)
- Scan único do sistema de arquivos
- Configuração mais robusta com verificações de erro
- Melhor tratamento de falhas
- Execução em push, pull request e diariamente

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

#### Erro: "unknown command 'dep' for 'trivy'"

**Causa**: O comando `dep` não existe no Trivy. O comando correto é `fs` (filesystem).

**Solução**: 
1. Usar `scan-type: 'fs'` em vez de `scan-type: 'dep'`
2. O Trivy automaticamente detecta dependências no filesystem
3. Configurar severidade adequada: `severity: 'CRITICAL,HIGH,MEDIUM'`

#### Erro: "Path does not exist: trivy-results.sarif"

**Causa**: O arquivo SARIF não foi gerado devido a falha no scan.

**Solução**: 
1. Adicionar verificações de existência do arquivo
2. Usar condicionais `if: always() && steps.check-scan.outputs.exists == 'true'`
3. Implementar logs detalhados para debug

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