# Website Monitor

Proyecto pequeño para practicar:

- Arquitectura Hexagonal (Ports and Adapters)
- Microservicios
- Comunicación entre servicios HTTP

La idea es entender bien los conceptos con un sistema sencillo pero suficiente para ver responsabilidades separadas, contratos entre servicios y flujo completo de monitoreo.

## Servicios actuales (3)

### 1) Config Service (puerto 8080)
Responsable de registrar y listar los sitios a monitorear.

### 2) Pinger Service (worker - puerto 8081)
Responsable de:

- Pedir sitios al Config Service
- Hacer los pings HTTP
- Enviar cada resultado al History Service

Nota: este servicio ya no guarda historial local ni expone endpoint de resultados.

### 3) History Service (puerto 8082)
Responsable de:

- Guardar resultados de ping
- Exponer resumen histórico por URL (uptime, total checks, último estado)

## Flujo

1. Registras sitios en Config Service.
2. Pinger Service consulta esos sitios periódicamente.
3. Pinger envía cada resultado a History Service.
4. Consultas el resumen en History Service.

## Ejemplo para crear sitio (Config Service)

POST http://localhost:8080/sites

```json
{
  "url": "https://www.noexiste123450.com",
  "reviewTime": 40
}
```

## Ejemplo para consultar historial/resumen (History Service)

GET http://localhost:8082/results

Respuesta ejemplo:

```json
[
  {
    "url": "https://www.instagram.com",
    "uptime": 100,
    "total_checks": 21,
    "last_status": "UP",
    "last_checked": "2026-04-23T09:38:01.8455988-05:00"
  },
  {
    "url": "https://www.facebook.com",
    "uptime": 100,
    "total_checks": 21,
    "last_status": "UP",
    "last_checked": "2026-04-23T09:38:01.8516164-05:00"
  },
  {
    "url": "https://example.com",
    "uptime": 100,
    "total_checks": 1553,
    "last_status": "UP",
    "last_checked": "2026-04-23T09:25:33.4930173-05:00"
  },
  {
    "url": "https://www.noexiste123450.com",
    "uptime": 0,
    "total_checks": 21,
    "last_status": "DOWN",
    "last_checked": "2026-04-23T09:38:01.7141895-05:00"
  },
  {
    "url": "https://www.youtube.com",
    "uptime": 100,
    "total_checks": 21,
    "last_status": "UP",
    "last_checked": "2026-04-23T09:38:01.808978-05:00"
  },
  {
    "url": "https://www.google.com",
    "uptime": 100,
    "total_checks": 21,
    "last_status": "UP",
    "last_checked": "2026-04-23T09:38:01.8285973-05:00"
  }
]
```

## Estructura general

```text
website-monitor/
|- config-service/
|- pinger-service/
|- history-service/
```

## Estado actual

- Arquitectura base de microservicios funcionando
- Config -> Pinger -> History funcionando
- Persistencia en memoria (sin base de datos)

## Futuro

La siguiente evolución natural es un cuarto microservicio de alertas/notificaciones (por ejemplo cuando un sitio cambia de UP a DOWN).

---

Ultima actualizacion: Abril 2026
