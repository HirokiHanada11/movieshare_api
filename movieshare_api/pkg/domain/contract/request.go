package contract

type (
	// MoviePostRequestBody is request body for creating a movie.
	MoviePostRequestBody struct {
		UserID         string `json:"userId"`
		UserName       string `json:"userName"`
		Title          string `json:"title"`
		Overview       string `json:"overview"`
		Genre          string `json:"genre"`
		YouTubeTitleID string `json:"youtubeTitleId"`
	}

	// MoviePostRequestBody is request body for creating a movie.
	MoviePutRequestBody struct {
		GrinningScore *uint32 `json:"grinningScore"`
	}
	
	// LoginPostRequestBody is the structure of request body for creating a session token.
	LoginPostRequestBody struct {
		IdToken string `json:"idToken"`
	} 
)
