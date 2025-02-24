===========================================================

1: Como criar um arquivo .go pelo CMD

A- Iniciar o CMD pelo administrador

B- Navegar até a pasta desejada

C- Comando para abrir o arquivo:
fsutil file createnew nomedoarquivo.go 0

D- Comando para abrir o arquivo:
nomedoarquivo.go  


===========================================================

2: Executar um programa go pelo terminal

A- navegar até o diretorio

B- Comando para executar o arquivo:
go run nomedoarquivo.go

===========================================================

3: GERAR O EXECUTÁVEL

A- Comando gerar o arquivo binário:
go build -o nomedoarquivo.go


===========================================================

4: Criar o módulo para trabalhar com pacotes

Anotação- funções escritas com letra minuscula são invisíveis para pacotes externos

A- Comando gerar go.mod
go mod init github.com/DanyloHenrique/go