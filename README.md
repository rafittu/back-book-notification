# 📖 Back-end da aplicação Book Notification

###

<br>

Este projeto foi criado para enviar notificações diárias via SMS, utilizando AWS Lambda, S3 Bucket, SNS e CloudWatch. O objetivo é enviar uma página aleatória do livro **A Arte de Viver**, de Epicteto, como uma reflexão diária sobre virtude, felicidade e sabedoria. Cada ensinamento ocupa uma página e, ao longo do tempo, o código envia uma nova lição do livro, inspirando o destinatário com o "ensinamento do dia".

O fluxo de trabalho envolve o envio das informações para um bucket S3 e a configuração de um cron job via CloudWatch para garantir que as notificações sejam enviadas diariamente, sem falhas.

<br>

## Tecnologias e Funcionalidades

<br>

- **Golang** como linguagem de programação de código aberto;
- **AWS S3 Bucket** para armazenamento de dados JSON sobre páginas lidas, com controle de acesso via políticas **IAM**;
- **AWS SNS** como serviço de mensagens utilizado para enviar notificações SMS ao usuário;
- **AWS CloudWatch** configurado para disparar a execução da **Lambda** diariamente;

<br>

## Instalação

### Pré-requisitos

  - Conta na AWS
  - Acesso ao Console da AWS para configurar os serviços Lambda, SNS, S3 e CloudWatch
  - A AWS CLI ou SDK para interagir com a infraestrutura da AWS

### Passos para implantação
 
  1. **Configurar o Bucket S3**:
       - Crie um bucket S3 para armazenar os arquivos JSON
       - Configure a política de acesso adequada no IAM para permitir que a Lambda escreva no S3

    
  2. **Configurar o SNS**:
       - Crie um tópico SNS para enviar as notificações
       - Configure as assinaturas do SNS para os destinatários desejados


  3. **Criar a Função Lambda**:
       - No Console da AWS, crie uma função Lambda
       - Preencha o arquivo `.env` com as informações do S3 e SNS, em seguida carregue o código (zipado ou via repositório)
       - Configure as permissões necessárias (IAM role)

    
  4. **Configuração do CloudWatch**:
       - No Console da AWS, crie uma regra no CloudWatch Events
       - Configure para acionar a Lambda todos os dias no seu horário de preferência


  5. **Compilação e Deploy**:
        - Caso precise compilar o código Go, use o seguinte comando:

          ```bash
           GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./cmd/main.go
           zip deployment.zip bootstrap
          ``` 

<br>

##

<p align="right">
  <a href="https://www.linkedin.com/in/rafittu/">Rafael Ribeiro 🚀</a>
</p>
