# Titan

**Titan** é uma linguagem de script experimental, fortemente inspirada no Lua, mas com tipagem estática simples e escrita 100% em Go.

O projeto busca simplicidade, segurança e extensibilidade. Titan é pensado para ser um ponto de partida para estudo e evolução: começa como interpretador, mas com visão futura para suporte a `async/await`, waitgroups e compilação JIT ou AOT.

---

## ✨ Exemplo de código

```lua
local num1: int = 10
local num2: float = 20.0
local num3 = num1 + num2
print(num3) -- saída esperada: 30.0

local pessoa: TitanTable = {
  nome: string = "Ana",
  idade: int = 25
}

print(pessoa.nome) -- saída esperada: Ana
```

---

## 🎯 Objetivos

- Sintaxe leve, parecida com Lua
- Escrita em Go puro, sem dependências externas
- Tipagem estática simples: `int`, `float`, `string`, `bool`
- `TitanTable` para estrutura dinâmica de chave-valor
- Execução direta de arquivos `.titan`
- Estrutura limpa para aprendizado de interpretadores
- Rumo à concorrência com `async` e `await`

---

## ✅ Funcionalidades da versão 1

- [x] Interpretador via linha de comando
- [x] Leitura e execução de arquivos `.titan`
- [x] Tipos suportados:
  - `int`
  - `float`
  - `string`
  - `bool`
  - `table` → via `TitanTable`
- [x] Operações:
  - Soma (`+`)
  - Concatenação de strings (`..`)
- [x] Função `print(...)`
- [ ] Suporte básico a **`TitanTable`**:
  - Criação com `{ chave = valor }`
  - Acesso por `obj.chave` ou `obj["chave"]`
- [ ] Operadores relacionais (`==`, `<`, `>`)
- [ ] Condições (`if`, `else`)
- [ ] Laços (`while`, `for`)
- [ ] Funções definidas pelo usuário

---

## 🚀 Futuro (v2 e além)

- [ ] `async`, `await` e `waitgroup`
- [ ] Funções concorrentes com semântica clara
- [ ] Compilação para bytecode intermediário
- [ ] Suporte a módulos/imports
- [ ] JIT / AOT opcional
- [ ] Exportação como `.exe`, `.wasm`, etc

#### Exemplo futuro planejado:

```lua
async function process()
  waitgroup_add(2)

  async task1()
    print("tarefa 1")
    waitgroup_done()
  end

  async task2()
    print("tarefa 2")
    waitgroup_done()
  end

  waitgroup_wait()
  print("fim")
end

await process()
```

---

## 📁 Estrutura do Projeto

```
titan/
│
├── main.go               # Entrada do interpretador
│
├── lexer/                # Tokenização
│   ├── lexer.go
│   ├── token.go
│
├── parser/               # Parser + AST
│   ├── parser.go
│   ├── ast.go
│
├── interpreter/          # Execução e ambiente de variáveis
│   ├── interpreter.go
│   ├── environment.go
│   ├── types.go          # Definições como TitanValue, TitanTable
│
├── examples/             # Exemplos de scripts `.titan`
│   └── test.titan
│
├── go.mod
└── README.md
```

---

## 📦 TitanTable

`TitanTable` é a estrutura de dados que representa objetos, arrays e dicionários, de forma semelhante às *tables* do Lua. Ela permite armazenar pares `chave: valor` com acesso dinâmico.

Exemplo:

```lua
local config = {
  host = "localhost",
  port = 8080
}

print(config.host)
```

Internamente no Go, é implementada como:

```go
type TitanTable struct {
  fields map[string]TitanValue
}
```

---

## ▶️ Como executar

### Requisitos:
- [Go instalado](https://golang.org/dl/)

### Executando um script:

```bash
go run main.go examples/test.titan
```

---

## 🪪 Licença
MIT
