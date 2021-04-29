package telebotapi

type Message struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat     Chat
	Date     int    `json:"date"`
	Text     string `json:"text"`
	Entities []struct {
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		Type   string `json:"type"`
	} `json:"entities"`
	Document Document `json:"document"`
}

type Document struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
}

type FileResponse struct {
	Ok bool `json:"ok"`
	Result File `json:"result"`
}

type File struct {
	FileId       string `json:"file_id"`
	FilePath     string `json:"file_path"`
}


type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Update struct {
	UpdateID int `json:"update_id"`
	Message  Message
}

type Response struct {
	Ok     bool `json:"ok"`
	Result []Update
}
