# Website Monitor 🔍

Sistema de monitoreo de sitios web en tiempo real. Un proyecto modular que verifica continuamente la disponibilidad y rendimiento de URLs registradas.

## 🎯 Idea General

El sistema está diseñado como una arquitectura de microservicios con responsabilidades bien separadas:

### 1. **Config Service** 🔧
- **Responsabilidad**: Gestionar la configuración de sitios a monitorear
- **Funcionalidades**:
  - Registrar URLs para monitoreo
  - Establecer frecuencia de revisión (ej: cada 10 segundos, cada 30 segundos)
  - Almacenar y recuperar configuración de sitios
  - API REST para CRUD de sitios
- **Ejemplo**: 
  ```
  POST /sites
    {
    "url": "https://www.ejemplo.com",
    "reviewTime": 10
    }
  ```

### 2. **Pinger Service** ❤️
- **Responsabilidad**: El corazón del sistema - ejecuta el monitoreo real
- **Funcionalidades**:
  - Obtiene las URLs registradas del Config Service
  - Realiza peticiones HTTP periódicamente
  - Mide tiempos de respuesta
  - Detecta fallos y cambios de estado
  - Registra resultados de cada ping

### 3. **History Service** 📊 (Próximo)
- **Responsabilidad**: Persistencia histórica de los pings
- **Funcionalidades**:
  - Almacenar cada resultado de ping
  - Consultar histórico por sitio
  - Análisis de disponibilidad (uptime)

### 4. **Alert Service** 🔔 (Opcional - Futuro)
- **Responsabilidad**: Notificaciones cuando ocurren problemas
- **Funcionalidades**:
  - Detectar cambios de estado (UP → DOWN)
  - Enviar alertas (consola, email, webhooks, etc.)

## 📁 Estructura del Proyecto

```
website-monitor/
├── config-service/          # Servicio de configuración
│   ├── cmd/api/             # Punto de entrada
│   ├── internal/
│   │   ├── adapters/        # Handlers HTTP, repositorios
│   │   ├── application/     # Lógica de negocio
│   │   ├── domain/          # Modelos y puertos
│   │   └── infrastructure/  # Conexiones BD
│   └── go.mod
│
└── pinger-service/          # Servicio de monitoreo
    ├── cmd/                 # Punto de entrada
    ├── internal/
    │   ├── adapters/        # Scheduler, clientes HTTP
    │   ├── application/     # Lógica de servicio
    │   ├── domain/          # Modelos y puertos
    │   └── infrastructure/  # Persistencia
    └── go.mod
```

## 🏗️ Arquitectura

El proyecto sigue **Arquitectura Hexagonal** (Ports & Adapters) con:
- **Domain**: Modelos centrales (`Site`, `PingResult`)
- **Ports**: Interfaces de entrada/salida
- **Adapters**: Implementaciones concretas (HTTP, BD, Scheduler)
- **Application**: Lógica de negocio de servicios

## 📈 Estado Actual

⚠️ **En desarrollo**
- [x] Estructura base de servicios
- [x] Config Service con modelos
- [x] Pinger Service con scheduler
- [ ] Persistencia real (bases de datos)
- [ ] History Service
- [ ] Alert Service
- [ ] Tests
- [ ] Documentación API

## 🚀 Próximos Pasos

1. Implementar persistencia en PostgreSQL
2. Crear Historia de pings
3. Agregar sistema de alertas
4. Documentación Swagger/OpenAPI
5. Docker & Docker Compose
6. Tests unitarios e integración

---

**Última actualización**: Abril 2026
