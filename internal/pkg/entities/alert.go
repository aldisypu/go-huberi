package entities

type Alert struct {
	ID              uint   `json:"id"`
	BackgroundColor string `json:"background_color"`
	HighlightColor  string `json:"highlight_color"`
	TextColor       string `json:"text_color"`
	Duration        int16  `json:"duration"`
	IsUnit          bool   `json:"is_unit"`
	IsBorder        bool   `json:"is_border"`
	CreatedAt       int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt       int64  `gorm:"autoCreateTime:milli;autoUpdateTime:milli" json:"updated_at"`
}
