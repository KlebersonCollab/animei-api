// Struct para representar um episódio
package entity

type Episodio struct {
	ID        int    `json:"id"`
	IDAnime   int    `json:"id_anime"`
	Temporada int    `json:"temporada"`
	Episodio  int    `json:"episodio"`
	URLThumb  string `json:"urlthumb"`
	Intro     string `json:"intro"`
	URLVideo  string `json:"urlvideo"`
}
