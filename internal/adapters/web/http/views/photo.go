package view

type PhotoInput struct {
	ShootingDate string `json:"shootingDate"`
}

type DownloadImageFileInput struct {
	ImageUrl string `json:"imageUrl"`
}
