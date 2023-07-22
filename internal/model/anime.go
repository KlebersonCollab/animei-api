// Struct para representar um anime
package entity

type Anime struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	URLCapa   string  `json:"urlcapa"`
	Tipo      string  `json:"tipo"`
	Intro     string  `json:"intro"`
	URLBanner string  `json:"urlbanner"`
	Nota      float64 `json:"nota"`
}
