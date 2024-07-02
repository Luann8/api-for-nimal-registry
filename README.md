### Executar a API

1. **Navegar até o diretório do projeto**:
   No terminal, navegue até o diretório onde o projeto está localizado:
   ```sh
   cd /caminho/do/projeto
   ```

2. **Executar o servidor**:
Execute o seguinte comando para iniciar o servidor na porta 8000:
```sh
Copiar código
go run main.go
```

3. **Testar a API**:
Você pode usar ferramentas como curl ou Postman para testar a API.

4. **Registrar um novo animal**
Para registrar um novo animal, utilize o seguinte comando curl:

```sh
curl -X POST http://localhost:8000/api/animal -H "Content-Type: application/json" -d '{
    "id": "1",
    "brinco_id": "12345",
    "lote": "A",
    "manejo": "SISBOV",
    "nascimento": "2020-01",
    "idade": 24,
    "raca": "Nelore",
    "sexo": "Macho",
    "categoria": "Bezerro",
    "cria": "Sim",
    "mae": "Mãe1",
    "pai": "Pai1",
    "peso_anterior": 200.5,
    "peso_atual": 250.0,
    "medicamentos": "Vermífugo",
    "observacoes": "Nenhuma"
}'
```



