package db

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	// User Basic data
	UserName    string
	PhoneNumber string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Gender      string
	Age         int

	// User Account data
	Exp                  int       `gorm:"not null;default:'0'"`
	Level                int       `gorm:"not null;default:'1'"`
	HeadPortraitUrl      string    ``
	FollowingUsers       []*User   `gorm:"many2many:user_following_ships;association_jointable_foreignkey:following_user_id"`
	FollowerUsers        []*User   `gorm:"many2many:user_follower_ships;association_jointable_foreignkey:follower_user_id"`
	CollectedArticles    []Article `gorm:"many2many:user_collected_article"`
	CollectedTopics      []Topic   `gorm:"many2many:user_collected_topic"`
	UserPrivacySetting   UserPrivacySetting
	UserPrivacySettingID uint

	// User healthy data
	Height float64
	Weight float64
	Area   string
	Job    string
}

type UserPrivacySetting struct {
	gorm.Model
	ShowPhoneNumber bool
	ShowGender      bool
	ShowAge         bool
	ShowHeight      bool
	ShowWeight      bool
	ShowArea        bool
	ShowJob         bool
}

type ArticleLabel struct {
	gorm.Model
	Value    string
	Articles []*Article `gorm:"many2many:articles_to_labels;"`
}

type Article struct {
	gorm.Model
	Title           string          `gorm:"not null"`
	Content         string          `gorm:"not null;type:text;"`
	Labels          []*ArticleLabel `gorm:"many2many:articles_to_labels;"`
	CoverImageUrl   string
	ReadCount       int `gorm:"not null;default:'0'"`
	ArticleComments []ArticleComment
}

type ArticleComment struct {
	gorm.Model
	Content       string `gorm:"not null;type:text"`
	ThumbsUpCount int    `gorm:"not null;default:'0'"`
	User          User
	UserID        int
	ArticleID     int
}

type Topic struct {
	gorm.Model
	Content       string `gorm:"not null;type:text"`
	User          User
	PictureUrls   string
	ThumbsUpCount int `gorm:"not null;default:'0'"`
	LordReplies   []TopicLordReply
}

type TopicLordReply struct {
	gorm.Model
	Content       string `gorm:"not null;type:text"`
	User          User
	PictureUrls   string
	ThumbsUpCount int `gorm:"not null;default:'0'"`
	TopicID       uint
	LayerReplies  []TopicLayerReply
}

type TopicLayerReply struct {
	gorm.Model
	Content          string `gorm:"not null;type:text"`
	User             User
	ThumbsUpCount    int `gorm:"not null;default:'0'"`
	TopicLordReplyID uint
}
