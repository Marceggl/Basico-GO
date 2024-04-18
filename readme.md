# Começando com GOLang

Estou começando a aprender a linguagem GO seguindo a documentação do site, a estrutura das pastas deste projeto podem ficar confusas

Documentação que estou seguindo: https://go.dev/doc/

# 1. importar dependencias

É necessário gerenciar as depencias utilizadas no código pelo seu próprio módulo que é definido como um arquivo "go.mod"

- Iniciar track de dependencias

```go mod init <exemplo/moddulo>```

É possivel publicar um módulo no github e importa-lo adicionando-o no arquivo mod omo github.com/meu-modulo

- Módulos ainda locais

Para utilizar módulos ainda locais, é necessário utilizar este comando abaixo

```go mod edit -replace example.com/greetings=../greetings```

Isto especifica que example.com/greetings deve ser trocado por ../greetings para conseguir localizar a dependencia, assim o arquivo .mod é modifica e ficará assim:

```
    module example.com/hello

    go 1.16

    replace example.com/greetings => ../greetings
```

- Para sincronizar o módulo use

```go mod tidy```

Após o comando ser executado, o módulo aparecerá como:

```
    module example.com/hello

    go 1.16

    replace example.com/greetings => ../greetings

    require example.com/greetings v0.0.0-00010101000000-000000000000
```

