// Arquivo model/anime.go
package model

type Episodio struct {
	ID          int    `json:"id"`
	IDTemporada int    `json:"id_temporada"`
	Numero      int    `json:"numero"`
	URLThumb    string `json:"urlthumb"`
	Intro       string `json:"intro"`
	URLVideo    string `json:"urlvideo"`
}
