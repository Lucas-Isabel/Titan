# 🐉 Titan

**Titan** é uma linguagem de script experimental, fortemente inspirada em Lua, com tipagem estática simples, escrita 100% em Go.

Ela foi projetada para ser **simples, segura e extensível**, evoluindo em etapas claras: começando como interpretador, passando por transpiler com verificação de tipos e chegando até suporte a `async/await`, concorrência e até mesmo compilação para bytecode ou binário.

---

## ✨ Exemplo de código Titan

```lua
let num1: int = 10
let num2: float = 20.0
let num3 = num1 + num2
print(num3) -- saída esperada: 30.0

let pessoa: TitanTable = {
  nome: string = "Ana",
  idade: int = 25
}

print(pessoa.nome) -- saída esperada: Ana
```

---

## 🧱 Etapas do Projeto

### 1. **Execução de Lua com GopherLua**
- Base mínima para rodar scripts Lua em Go.
- Permite evoluir passo a passo com o runtime seguro e extensível.

### 2. **Transpiler: Titan → Lua**
- Lê código com tipagem (`Titan`) e converte para Lua puro.
- Remove anotações como `: int`, mantendo o código Lua executável.

### 3. **Verificador de Tipos (análise estática)**
- Antes de rodar, o código é verificado para garantir tipos corretos.
- Tipagem para variáveis, funções, parâmetros e retorno.
- Erros emitidos em tempo de compilação.

### 4. **Novas Funcionalidades**
- `async`, `await` e waitgroups.
- Controle de concorrência integrado à linguagem.
- Exportação para bytecode, binário, ou até WebAssembly (WASM).

---

## 🎯 Objetivos

- Sintaxe leve e familiar para quem conhece Lua
- Escrita em Go puro (sem dependências C)
- Tipagem estática simples: `int`, `float`, `string`, `bool`
- Estrutura dinâmica com `TitanTable`
- Ideal para aprendizado de compiladores e linguagens
- Rumo à concorrência e paralelismo com semântica clara

---

## ✅ Funcionalidades Atuais

- [x] Execução de scripts `.titan`
- [x] Tipagem básica: `int`, `float`, `string`, `bool`
- [ ] `TitanTable` com acesso por ponto e string
- [ ] Tipagem em functions: retorno, parametros e chamadas

---

## 🚀 Futuro da linguagem

- [ ] `async function` e `await`
- [ ] `waitgroup` para controle de tarefas assíncronas
- [ ] Analisador semântico completo e escopos aninhados
- [ ] Compilação para bytecode
- [ ] Exportação de `lua` para `.exe`

### Exemplo Futuro:

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
├── main.go               # Entrada principal
│
├── runtime/              # Execução com GopherLua
│   └── vm.go
│
├── transpiler/           # Titan para Lua
│   └── transpile.go
│
├── typechecker/          # Verificador de tipos
│   └── checker.go
│
├── examples/             # Scripts Titan para teste
│   └── hello.titan
│
├── go.mod
└── README.md
```

---

## ▶️ Como executar

### Pré-requisitos:
- [Go instalado](https://golang.org/dl/)

### Execução:

```bash
go run main.go examples/hello.titan
```

---

## 🪪 Licença

MIT
