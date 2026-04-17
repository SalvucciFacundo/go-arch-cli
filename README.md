# Go-Architect CLI (go-arch) 🚀

**Go-Architect CLI** es una herramienta de scaffolding profesional, agnóstica y multiplataforma diseñada para estandarizar la creación de proyectos en Go. Inspirada en la eficiencia del Angular CLI, permite generar arquitecturas limpias y mantenibles en segundos.

## ✨ Características Principales

- 🏗️ **Arquitecturas Soportadas**: Minimalista, Standard y Hexagonal (Ports & Adapters).
- 🔌 **Agnosticismo Total**: Independiente de la base de datos (PostgreSQL, MySQL, MongoDB) y del IDE.
- ⚡ **Hot-Reload**: Integración nativa con `Air` para un ciclo de desarrollo ultra rápido.
- 🛠️ **Generadores de Componentes**: Crea servicios, repositorios y handlers con un solo comando.
- 🧪 **QA Ready**: Generación automática de tests y mocks manuales para fomentar TDD.
- 💻 **Multiplataforma**: Soporte total para Windows (ANSI), Linux y macOS.

## 🚀 Instalación

```bash
go install github.com/tu-usuario/go-arch@latest
```

## 🛠️ Comandos Rápidos

### 1. Configurar Entorno
Detecta tu sistema operativo e instala Go y las dependencias necesarias.
```bash
go-arch setup
```

### 2. Crear un Nuevo Proyecto
Asistente interactivo para configurar nombre, arquitectura y base de datos.
```bash
go-arch new mi-proyecto
```

### 3. Ejecutar con Hot-Reload
Lanza el servidor de desarrollo.
```bash
go-arch serve
```

### 4. Generar Componentes
Genera piezas de código que se ubican automáticamente según la arquitectura del proyecto.
```bash
go-arch generate service Product
go-arch generate handler Product
```

## 📐 Arquitecturas Disponibles

- **Minimalist**: Para microservicios o scripts de un solo archivo.
- **Standard**: Estructura clásica de Go para proyectos medianos.
- **Hexagonal**: El estándar de oro para aplicaciones empresariales desacopladas del mundo exterior.

## 🤝 Contribuciones

¡Las PR son bienvenidas! Si tenés sugerencias para nuevas plantillas o drivers de base de datos, no dudes en abrir un issue.

---
Desarrollado con ❤️ para la comunidad de Go.
