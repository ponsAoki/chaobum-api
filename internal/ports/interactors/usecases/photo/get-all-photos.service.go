package port

import entity_port "chaobum-api/internal/ports/domains/entities"

type GetAllPhotosService interface {
	Handle() ([]entity_port.PhotoPort, error)
}
