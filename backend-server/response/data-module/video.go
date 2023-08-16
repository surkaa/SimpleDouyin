// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    video, err := UnmarshalVideo(bytes)
//    bytes, err = video.Marshal()

package data_module

import "encoding/json"

func UnmarshalVideo(data []byte) (Video, error) {
	var r Video
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Video) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Video 视频信息返回体模型
type Video struct {
	Author        UserResponseModule `json:"author"`         // 视频作者信息
	CommentCount  int64              `json:"comment_count"`  // 视频的评论总数
	CoverURL      string             `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64              `json:"favorite_count"` // 视频的点赞总数
	ID            int64              `json:"id"`             // 视频唯一标识
	IsFavorite    bool               `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string             `json:"play_url"`       // 视频播放地址
	Title         string             `json:"title"`          // 视频标题
}
