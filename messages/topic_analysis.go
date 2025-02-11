package messages

type TopicAnalysisTask struct {
	Type         string `json:"type"`           // Тип сообщения
	RedditPostID int    `json:"reddit_post_id"` // ID поста на реддите
}

// Константа для типа сообщения
const TopicAnalysisTaskType = "topic_analysis"
