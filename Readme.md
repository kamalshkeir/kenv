# A minimalistic pkg to load environment variables to strct

## get it :
```sh
go get -u github.com/kamalshkeir/kenv@latest
```
## Example:

##### Struct to Fill from env
```go
type GlobalConfig struct {
	Host       string `kenv:"HOST|localhost"` // DEFAULT to 'localhost': if HOST not found in env
	Port       string `kenv:"PORT|9313"`
	Embed struct {
		Static    bool `kenv:"EMBED_STATIC|false"`
		Templates bool `kenv:"EMBED_TEMPLATES|false"`
	}
	Db struct {
		Name     string `kenv:"DB_NAME|db"`
		Type     string `kenv:"DB_TYPE"` // REEQUIRED: this env var is required, you will have error if empty
		DSN      string `kenv:"DB_DSN|"` // NOT REQUIRED: if DB_DSN not found it's not required, it's ok to stay empty
	}
	Smtp struct {
		Email string `kenv:"SMTP_EMAIL|"`
		Pass  string `kenv:"SMTP_PASS|"`
		Host  string `kenv:"SMTP_HOST|"`
		Port  string `kenv:"SMTP_PORT|"`
	}
	Profiler   bool   `kenv:"PROFILER|false"`
	Docs       bool   `kenv:"DOCS|false"`
	Logs       bool   `kenv:"LOGS|false"`
	Monitoring bool   `kenv:"MONITORING|false"`
}
```

##### Fill
```go
kenv.Load(".env") // load env files and add to env vars
// the command:
err := kenv.Fill(&Config) // fill struct with env vars loaded before
```

