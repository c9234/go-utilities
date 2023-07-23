package internal

type Follower struct {
	Name       string
	ScreenName string
	Url        string
}

type RawFollowersJson struct {
	Data struct {
		User struct {
			Result struct {
				Timeline struct {
					Timeline struct {
						Instructions []*Instruction `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

type Instruction struct {
	Type    string   `json:"type"`
	Entries []*Entry `json:"entries,omitempty"`
}

type Entry struct {
	Content struct {
		EntryType   string `json:"entryType"`
		ItemContent struct {
			UserResults struct {
				Result struct {
					Id     string `json:"id"`
					Legacy struct {
						Name        string `json:"name"`
						ScreenName  string `json:"screen_name"`
						Description string `json:"description"`
					} `json:"legacy"`
				} `json:"result"`
			} `json:"user_results"`
		} `json:"itemContent"`
	} `json:"content"`
}
