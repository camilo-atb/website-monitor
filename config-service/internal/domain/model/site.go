package model

import "time"

type MonitoredURL struct {
	ID           int           `json:"id"`
	URL          string        `json:"url"`
	ReviewTime   time.Duration `json:"reviewTime"`
	CreationDate time.Time     `json:"creationDate"`
	ModifyDate   time.Time     `json:"modifyDate"`
}

/*
🔌 Input port
Crear sitio
Actualizar
Eliminar
Listar

💾 Output port
Guardar en DB
Consultar
Eliminar
*/
