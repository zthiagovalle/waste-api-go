# waste-api-go

## 💻 Projeto
Uma api em golang que le os dados de uma planilha publica no google sheets, converte para uma struct e retornar em json.

Esse projeto serve de base para ser consumido no front-end feito no repositório https://github.com/zthiagovalle/recicle-municipio.

Os dados na planilha estão assim:
<img width="1434" alt="image" src="https://user-images.githubusercontent.com/46169077/227675474-5aba53b7-708d-4eee-902c-7e427dd329ed.png">

Existe uma opção para compartilhar na web, indo em File -> Share -> Publish to web
<img width="878" alt="image" src="https://user-images.githubusercontent.com/46169077/227675729-3a35df19-49c8-4b15-8975-395b02f7faf9.png">

Uma das opções de compartilhamento é em csv
<img width="725" alt="image" src="https://user-images.githubusercontent.com/46169077/227675966-f780515c-f8dd-4981-88e7-6f5cddfda2f7.png">

Feito isso voce pode de uma forma muito simples ter um armazenamento facil, gratis e simples de uma forma muito rapida.
Lembrando que todos conseguem visualizar os dados, então não armazene dados sensiveis.

###  📓 Requirementos
Docker

## 🚀 Como rodar
Após clonar o projeto, na raiz do projeto rode os comandos.
- docker-compose build
- docker-compose up -d

Feito isso a api vai estar no ar no localhost:3000/
![image](https://user-images.githubusercontent.com/46169077/227675404-6bb25de0-37c5-47ff-86c9-24a2fd3b73fc.png)

feito isso, só fazer um get que irá retornar os dados da planilha:
![image](https://user-images.githubusercontent.com/46169077/227676810-85d2c998-0a9e-49f8-bae9-d99662254a7d.png)
