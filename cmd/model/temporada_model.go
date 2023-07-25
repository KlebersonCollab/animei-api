// Arquivo model/anime.go
package model

type Temporada struct {
	ID        int        `json:"id"`
	IDAnime   int        `json:"id_anime"`
	Numero    int        `json:"numero"`
	Episodios []Episodio `json:"episodios"`
}
