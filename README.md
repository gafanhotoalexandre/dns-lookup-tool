# DNS Lookup Tool

Uma ferramenta de linha de comando simples e eficiente para consultas DNS, permitindo buscar endereços IP e servidores de nomes de domínios.

## 🚀 Funcionalidades

- **Busca de IPs**: Encontra todos os endereços IPv4 e IPv6 de um domínio
- **Busca de Servidores NS**: Lista os servidores de nomes (Name Servers) de um domínio
- **Interface amigável**: Comandos simples e output formatado
- **Validação**: Tratamento de erros e validação de entrada

## 📋 Pré-requisitos

- Go 1.24.1 ou superior
- Conexão com a internet para consultas DNS

## 🔧 Instalação

1. Clone o repositório:
```bash
git clone https://github.com/gafanhotoalexandre/dns-lookup-tool.git
cd dns-lookup-tool
```

2. Baixe as dependências:
```bash
go mod tidy
```

3. Compile o projeto:
```bash
go build
```

## 💻 Como usar

### Buscar endereços IP

```bash
# Usando o domínio padrão
./dns-lookup ip

# Especificando um domínio
./dns-lookup ip --host google.com

# Usando alias
./dns-lookup i --host github.com
```

### Buscar servidores de nomes

```bash
# Usando o domínio padrão
./dns-lookup servidores

# Especificando um domínio
./dns-lookup servidores --host google.com

# Usando aliases
./dns-lookup ns --host github.com
./dns-lookup s --host example.com
```

### Ajuda

```bash
# Ajuda geral
./dns-lookup help

# Ajuda para comandos específicos
./dns-lookup help ip
./dns-lookup help servidores
```

## 📖 Exemplos de saída

### Busca de IPs
```
$ ./dns-lookup ip --host google.com
Buscando IPs para: google.com
----------------------------------------
1. 142.250.191.78 (IPv4)
2. 2800:3f0:4004:c15::65 (IPv6)
```

### Busca de servidores NS
```
$ ./dns-lookup servidores --host google.com
Buscando servidores NS para: google.com
----------------------------------------
1. ns1.google.com.
2. ns2.google.com.
3. ns3.google.com.
4. ns4.google.com.
```

## 🛠️ Tecnologias utilizadas

- **Go 1.24.1**: Linguagem de programação
- **urfave/cli**: Framework para aplicações de linha de comando
- **net**: Biblioteca padrão do Go para operações de rede

## 📝 Comandos disponíveis

| Comando | Aliases | Descrição |
|---------|---------|-----------|
| `ip` | `i` | Busca endereços IP de um domínio |
| `servidores` | `ns`, `s` | Busca servidores de nomes de um domínio |

## 🎯 Flags disponíveis

| Flag | Valor padrão | Descrição |
|------|-------------|-----------|
| `--host` | `alemartins.dev.br` | Nome do host/domínio para consulta |

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

**Alexandre Martins**
- Website: [alemartins.dev.br](https://alemartins.dev.br)
- Email: alexandrevmartinsdelima@gmail.com
- GitHub: [@gafanhotoalexandre](https://github.com/gafanhotoalexandre)

---

⭐ Se este projeto te ajudou, considere dar uma estrela no repositório!