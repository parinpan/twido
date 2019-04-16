package dataprovider

type User struct {
	ID         int    `json:"id"`
	IDStr      string `json:"id_str"`
	ScreenName string `json:"screen_name"`
}

type VideoVariant struct {
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
	Bitrate     int    `json:"bitrate"`
}

type VideoInfo struct {
	AspectRatio    []int          `json:"aspect_ratio"`
	DurationMillis int            `json:"duration_millis"`
	Variants       []VideoVariant `json:"variants"`
}

type Media struct {
	ID        int       `json:"id"`
	IDStr     string    `json:"id_str"`
	VideoInfo VideoInfo `json:"video_info"`
}

type ExtendedEntities struct {
	Media []Media `json:"media"`
}

type Status struct {
	ID                   int              `json:"id"`
	IDStr                string           `json:"id_str"`
	Text                 string           `json:"text"`
	CreatedAt            string           `json:"created_at"`
	InReplyToStatusID    int              `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string           `json:"in_reply_to_status_id_str"`
	User                 User             `json:"user"`
	ExtendedEntities     ExtendedEntities `json:"extended_entities"`
}

type TwitterSearch struct {
	Statuses []Status `json:"statuses"`
}
