
# Weather App

Weather App é um serviço em Go que permite consultar o clima de uma determinada localização com base no CEP. Ele utiliza APIs externas para obter os dados de localização e temperatura e está implantado no Google Cloud Run.

---

## **Endpoint Disponível**

O serviço está disponível no seguinte endpoint público do Google Cloud Run:

- **Base URL:** [https://pg-lab1-681802062904.us-central1.run.app](https://pg-lab1-681802062904.us-central1.run.app)

---

## **Exemplos de Uso**

### **Consulta de Clima por CEP**

Faça uma requisição GET para o seguinte endpoint:

- **Endpoint:** `/weather/:cep`
- **Exemplo de CEP:** `01001000` (São Paulo - SP)

#### **Exemplo de Requisição**

```bash
curl -X GET "https://pg-lab1-681802062904.us-central1.run.app/weather/01001000"
```

#### **Resposta de Sucesso (HTTP 200)**

```json
{
  "temp_C": 25.0,
  "temp_F": 77.0,
  "temp_K": 298.0
}
```

#### **Erros Possíveis**

1. **CEP Inválido (menos de 8 dígitos):**
   - **Status:** `422 Unprocessable Entity`
   - **Resposta:**

     ```json
     {
       "message": "invalid zipcode"
     }
     ```

2. **CEP Não Encontrado:**
   - **Status:** `404 Not Found`
   - **Resposta:**

     ```json
     {
       "message": "can not find zipcode"
     }
     ```

3. **Erro ao Buscar Temperatura:**
   - **Status:** `500 Internal Server Error`
   - **Resposta:**

     ```json
     {
       "message": "temperature service error"
     }
     ```

---

## **Tecnologias Utilizadas**

- **Linguagem:** Go (Golang)
- **Framework Web:** Gin
- **APIs Utilizadas:**
  - [ViaCEP](https://viacep.com.br/) (Para busca de localização)
  - [WeatherAPI](https://www.weatherapi.com/) (Para busca de temperatura)
- **Containerização:** Docker
- **Cloud:** Google Cloud Run
- **Gerenciamento de Dependências:** Go Modules
- **Teste:** Testify

---

## **Configuração Local**

### **Pré-requisitos**

- Go 1.20+
- Docker
- Google Cloud SDK (opcional para deploy)

### **Passos para Executar Localmente**

1. Clone o repositório:

   ```bash
   git clone https://github.com/username/repository.git
   cd repository
   ```

2. Instale as dependências:

   ```bash
   go mod tidy
   ```

3. Configure as variáveis de ambiente:
   - Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

     ```env
     WEATHER_API_KEY=your_weather_api_key_here
     ```

4. Execute o servidor:

   ```bash
   go run main.go
   ```

5. Acesse a aplicação em [http://localhost:8080](http://localhost:8080).

---

## **Deploy**

Este serviço foi implantado no **Google Cloud Run**. Para refazer o deploy, siga os passos:

1. Autentique-se no Google Cloud:

   ```bash
   gcloud auth login
   gcloud auth configure-docker
   ```

2. Construa e envie a imagem Docker:

   ```bash
   docker build -t gcr.io/PROJECT-ID/weather-app:latest .
   docker push gcr.io/PROJECT-ID/weather-app:latest
   ```

3. Faça o deploy no Cloud Run:

   ```bash
   gcloud run deploy weather-app      --image gcr.io/PROJECT-ID/weather-app:latest      --platform managed      --region us-central1      --allow-unauthenticated
   ```

---

## **Licença**

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
