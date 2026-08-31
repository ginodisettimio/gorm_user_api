# GOrm User API

API REST para gestión de usuarios implementada en Go utilizando **GORM** y **SQLite**, bajo el patrón de diseño Repository.

---

## Arquitectura

* **`internal/model`**: Definición de entidades y esquemas de base de datos.
* **`internal/store`**: Capa de persistencia y consultas a la base de datos (Repository).
* **`internal/service`**: Lógica de negocio y validaciones.
* **`internal/transport`**: Handlers HTTP y procesamiento de peticiones/respuestas.
* **`main.go`**: Configuración de dependencias, base de datos y arranque del servidor.

---

## Requisitos

* **Go** (ver >= 1.27)
* **GCC / Build Essentials** (`gorm.io/driver/sqlite`)

---

## Ejecución

1. **Clonar el repositorio:**
   ```bash
   git clone https://github.com/ginodisettimio/gorm_user_api
   cd gorm_user_api

2. **Instalar dependencias:**
    ```bash
    go mod tidy

3. **Ejecutar**
    ```bash
    go run main.go

4. **Compilar (opcional)**
    ```
    go build .