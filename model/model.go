package model

// URLRequest — структура для входящего запроса на сокращение URL
// ShortURL — структура для ответа с короткой ссылкой
// URLMapping — структура для хранения соответствия shortID и longURL

type URLRequest struct {
	URL string `json:"url"`
}

type ShortURL struct {
	ShortURL string `json:"short_url"`
}

type URLMapping struct {
	ShortID string
	LongURL string
}
