<div>
<h1 align="center"> TurmaFinder </h1>
</div>
<p aligin="center">
<img alt="GitHub all releases" src="https://img.shields.io/github/downloads/alternativeon/turmafinder/total?style=for-the-badge">
<img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/alternativeon/turmafinder?style=for-the-badge">
<img alt="GitHub" src="https://img.shields.io/github/license/alternativeon/turmafinder?style=for-the-badge">
<img src="https://img.shields.io/badge/vers%C3%A3o-windows-success?style=for-the-badge&logo=windows">
</p>
Descubra em qual turma você está esse ano.

## Como saber em qual turma estou?
Veja a seção de [downloads/releases](https://github.com/alternativeon/turmafinder) para baixar a ferramenta.

## Como funciona?
Descobri isso acidentalmente enquanto fazia webscraping. O positivo on faz uma requisão para `https://apihub.positivoon.com.br/api/NivelEnsino` com sua token de acesso, e a resposta disso é sua turma. Mesmo a escola dissendo que "as turmas não foram formadas", elas estão já presentes no positivo.

Essa informação provavelmente é usada na página inicial, mostrado em "aluno do xº ano do ensino fundamental/medio", estranhamente a turma não é mostrada, possivelmente é usada em outras coisas (como determinar as salas virtuais, por exemplo).

## FAQ
> **Pergunta**: O Chrome disse que o arquivo não costuma ser transferido por download e pode ser perigoso. Você vai me hackear?

**Resposta**: Aplicativos que não estão na "lista branca" do Chrome costumam receber esses alertas, mas são falsos postivos, uma maneira de resolver isso é assinando o código digitalmente, porém por agora é inviavel, por ser algo relativamente caro. E não, eu não tenho intenções, não quero e nem posso te hackear, o código-fonte da aplicação está nesse repositorio de maneira que você possa ler sem necessidade de autorização previa minha.
##### [_Saiba mais sobre os avisos do Chrome_](https://support.google.com/chrome/answer/6261569?hl=pt-BR&dark=1#zippy&zippy=#:~:text=Incomum:,desconhecido%20e%20possivelmente%20perigoso.)

****

> **P:** Como posso construir o arquivo?

**R:** [Baixe a linguagem Go](https://go.dev) e faça o download desse repositório. Na pasta do repositorio, abra uma janela do CMD ou Powershell e digite: `go mod tidy` (para resolver as dependencias) e depois `go build`. Após a compilação, um arquivo chamado `turmafinder.exe` deve ser gerado.

****

> **P:** Como contribuo para o projeto?

**R:** Abra uma issue ou crie uma pull request :)

****

## Licensa
Este projeto é licenciado para você sobre a licença [BSD 3 Clause](https://choosealicense.com/licenses/bsd-3-clause/).
