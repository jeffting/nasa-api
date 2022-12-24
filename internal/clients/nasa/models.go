package nasa

type images struct {
	Photos []photo `json:"photos"`
}

type photo struct {
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
}
