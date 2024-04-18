# Começando com GOLang

Estou começando a aprender a linguagem GO seguindo a documentação do site, a estrutura das pastas deste projeto podem ficar confusas

- Iniciar track de dependencias

```go mod init exemple/hello```

- Para módulos que ainda estão locais na máquina, é preciso rodar o seguinte comando

```go mod edit -replace example.com/greetings=../greetings```

Isto especifica que example.com/greetings deve ser trocado por ../greetings para conseguir localizar a dependencia, assim o arquivo .mod é modifica e ficará assim:

```
    module example.com/hello

    go 1.16

    replace example.com/greetings => ../greetings
```

- Para sincronizar o módulo use

```go mod tidy```