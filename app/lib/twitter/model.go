package twitter

type User struct {
	ID         string `json:"id_str"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	ImageURL   string `json:"profile_image_url_https"`
}

type NewPost struct {
	Status string `json:"status"`
}

type Post struct {
	Text string `json:"text"`
}

type UserTimeline struct {
	Posts []Post
}
