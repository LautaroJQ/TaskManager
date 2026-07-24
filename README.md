# TaskManager

TaskManager es una aplicación de línea de comandos (CLI) desarrollada en Go utilizando el framework **Cobra** para administrar tareas de forma simple desde la terminal.

Actualmente permite crear, listar y eliminar tareas, almacenándolas de manera persistente en un archivo JSON.

## Características

- Crear tareas.
- Listar todas las tareas.
- Eliminar tareas.
- Persistencia mediante JSON.
- Duración configurable para cada tarea.
- Fecha de creación y fecha de vencimiento automáticas.

## Tecnologías utilizadas

- Go
- Cobra CLI
- JSON para persistencia

## Instalación

Clonar el repositorio:

```bash
git clone https://github.com/LautaroJQ/TaskManager.git
cd TaskManager
```

Instalar las dependencias:

```bash
go mod tidy
```

Compilar el proyecto:

```bash
go build -o task
```

También puede ejecutarse directamente con:

```bash
go run .
```

## Uso

### Agregar una tarea

```bash
task add "Aprender Cobra"
```

Con descripción y duración:

```bash
task add "Finalizar proyecto" \
    --description "Completar la CLI" \
    --duration 10
```

### Listar tareas

```bash
task list
```

### Eliminar una tarea

```bash
task remove 3
```

## Estructura del proyecto

```
TaskManager
├── cmd/          # Comandos de Cobra
├── models/       # Modelos del dominio
├── storage/      # Persistencia de datos
├── main.go
├── go.mod
└── tasks.json
```

## Próximas funcionalidades

- [ ] Marcar tareas como completadas
- [ ] Editar tareas
- [ ] Buscar tareas
- [ ] Filtros por estado
- [ ] Ordenamiento
- [ ] Tests unitarios
- [ ] Persistencia con SQLite

## Aprendizajes

Este proyecto fue desarrollado con el objetivo de practicar:

- Organización de proyectos en Go.
- Desarrollo de aplicaciones CLI con Cobra.
- Manejo de archivos JSON.
- Modelado de estructuras.
- Manejo de fechas utilizando `time.Time`.
- Separación de responsabilidades entre dominio, comandos y persistencia.

## Licencia

Este proyecto se distribuye bajo la licencia MIT.