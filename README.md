# dojo-test-go

## Instructions ##

```sh
# Clone the repository
git@github.com:klasrak/go-meli-test-dojo.git

# Install/fix dependencies
go mod tidy

# Run API
make run

# Run tests
make test
```

## cURL ##

**GET Starships**
```curl
curl --request GET \
  --url http://localhost:3000/api/v1/starships
```

**GET Starship by ID**
```curl
curl --request GET \
  --url http://localhost:3000/api/v1/starships/9
```

**GET People list**
```curl
curl --request GET \
  --url http://localhost:3000/api/v1/people
```

**GET People by ID**
```curl
curl --request GET \
  --url http://localhost:3000/api/v1/people/1
```
