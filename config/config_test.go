package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Crear un archivo de configuración temporal

	content := []byte(`
app:
  api: v1beta1
  port: 8080
database:
  name: test-db
  user: test
  password: changepass
  host: db.example.org:3306
`)

	tmpfile, err := os.CreateTemp("./", "test*.yaml")
	if err != nil {
		t.Fatalf("No se pudo crear el archivo temporal: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatalf("No se pudo escribir en el archivo temporal: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("No se pudo cerrar el archivo temporal: %v", err)
	}

	// Inicializar viper con el archivo de configuración temporal
	InitViper(tmpfile.Name())

	// Cargar la configuración
	config := LoadConfig()

	// Verificar los valores cargados
	if config.App.ApiVersion != "v1beta1" {
		t.Errorf("App api version 'v1beta1' was expected but have '%s'", config.App.ApiVersion)
	}
	if config.App.Port != 8080 {
		t.Errorf("App port 8080 was expected but have %d", config.App.Port)
	}

	if config.Db.Host != "db.example.org:3306" {
		t.Errorf("Host db.example.org:3306 was expected but have '%s'", config.Db.Host)
	}

	if config.Db.Name != "test-db" {
		t.Errorf("Database name 'test-db' was expected, but have '%s'", config.Db.Name)
	}

	if config.Db.User != "test" {
		t.Errorf("user 'test' was expected but have '%s'", config.Db.User)
	}

	if config.Db.Passwd != "changepass" {
		t.Errorf("db password 'changepass' was expected, but have '%s'", config.Db.Passwd)
	}
}
