package reader

import "com.vorto.vehiclerouter/pkg/models"

// LoadsFileReader is a generic interface for reading loads data stored in files
type LoadsFileReader interface {
	Read(filePath string) ([]models.Load, error)
}
