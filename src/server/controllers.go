package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"utils"
)

func accountSendVerificationCode(c *gin.Context) {
	PhoneNumber := c.Query("phone_number")
	resp := sendVerificationCode(PhoneNumber)
	c.JSON(resp.Status, resp.Data)
}

func accountRegister(c *gin.Context) {
	PhoneNumber := c.PostForm("phone_number")
	Password := c.PostForm("password")
	UserName := c.PostForm("username")
	Code := c.PostForm("code")
	resp := registerNewUser(PhoneNumber, Password, UserName, Code)
	c.JSON(resp.Status, resp.Data)
}

func accountLogin(c *gin.Context) {
	PhoneNumber := c.Query("phone_number")
	Password := c.Query("password")
	resp := loginUser(PhoneNumber, Password)
	c.JSON(resp.Status, resp.Data)
}

func accountLogout(c *gin.Context) {
	SessionId := c.Query("session_id")
	resp := logoutUser(SessionId)
	c.JSON(resp.Status, resp.Data)
}

func accountAlterInformation(c *gin.Context) {
	userId, _ := c.Get("user_id")
	UserName := c.PostForm("username")
	Gender := c.PostForm("gender")
	Height, _ := strconv.ParseFloat(c.PostForm("height"), 64)
	Weight, _ := strconv.ParseFloat(c.PostForm("weight"), 64)
	Area := c.PostForm("area")
	Job := c.PostForm("job")
	Age, _ := strconv.Atoi(c.PostForm("age"))
	resp := alterUserInformation(userId.(int), UserName, Gender, Height, Weight, Area, Job, Age)
	c.JSON(resp.Status, resp.Data)
}

func accountGetUserInformation(c *gin.Context) {
	userId, _ := c.Get("user_id")
	OtherUserId := c.Query("other_user_id")
	var resp responseBody
	if OtherUserId == "" {
		resp = getUserInfoFromUserId(userId.(int))
	} else {
		otherUserId, _ := strconv.Atoi(OtherUserId)
		resp = getOtherUserInformationFromOtherUserId(userId.(int), otherUserId)
	}
	c.JSON(resp.Status, resp.Data)
}

func accountAlterPassword(c *gin.Context) {
	PhoneNumber := c.PostForm("phone_number")
	Code := c.PostForm("code")
	NewPassword := c.PostForm("password")
	resp := alterPassword(PhoneNumber, Code, NewPassword)
	c.JSON(resp.Status, resp.Data)
}

func accountGetUserPrivacySetting(c *gin.Context) {
	userId, _ := c.Get("user_id")
	resp := getUserPrivacySetting(userId.(int))
	c.JSON(resp.Status, resp.Data)
}

func accountAlterUserPrivacySetting(c *gin.Context) {
	userId, _ := c.Get("user_id")
	ShowPhoneNumber := c.PostForm("show_phone_number") == "true"
	ShowGender := c.PostForm("show_gender") == "true"
	ShowAge := c.PostForm("show_age") == "true"
	ShowHeight := c.PostForm("show_height") == "true"
	ShowWeight := c.PostForm("show_weight") == "true"
	ShowArea := c.PostForm("show_area") == "true"
	ShowJob := c.PostForm("show_job") == "true"
	resp := alterUserPrivacy(userId.(int), ShowPhoneNumber, ShowGender,
		ShowAge, ShowHeight, ShowWeight, ShowArea, ShowJob)
	c.JSON(resp.Status, resp.Data)
}

func accountGetUserFollowingList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserFollowingList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func accountGetUserFollowerList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserFollowerList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func accountFollowUser(c *gin.Context) {
	userId, _ := c.Get("user_id")
	OtherUserId := c.PostForm("other_user_id")
	otherUserId, _ := strconv.Atoi(OtherUserId)
	resp := followUser(userId.(int), otherUserId)
	c.JSON(resp.Status, resp.Data)
}

func accountIgnoreUser(c *gin.Context) {
	userId, _ := c.Get("user_id")
	OtherUserId := c.PostForm("other_user_id")
	otherUserId, _ := strconv.Atoi(OtherUserId)
	resp := ignoreUser(userId.(int), otherUserId)
	c.JSON(resp.Status, resp.Data)
}

func schoolGetArticle(c *gin.Context) {
	userId, _ := c.Get("user_id")
	ArticleId := c.Query("article_id")
	articleId, _ := strconv.Atoi(ArticleId)
	resp := getArticle(userId.(int), articleId)
	c.JSON(resp.Status, resp.Data)
}

func schoolGetArticlePage(c *gin.Context) {
	ArticleId := c.Param("id")
	articleId, _ := strconv.Atoi(ArticleId)
	content, err := getArticlePage(articleId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "article.tmpl", gin.H{
		"content": template.HTML(content),
	})
}

func schoolGetArticleList(c *gin.Context) {
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getArticleList(beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func schoolPublishArticleComment(c *gin.Context) {
	userId, _ := c.Get("user_id")
	ArticleId := c.PostForm("article_id")
	articleId, _ := strconv.Atoi(ArticleId)
	Content := c.PostForm("content")
	resp := createArticleComment(userId.(int), articleId, Content)
	c.JSON(resp.Status, resp.Data)
}

func schoolGetArticleCommentList(c *gin.Context) {
	ArticleId := c.Query("article_id")
	articleId, _ := strconv.Atoi(ArticleId)
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getArticleCommentList(articleId, beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func schoolSearchArticle(c *gin.Context) {
	SearchContent := c.Query("content")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getSearchArticleList(SearchContent, beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func schoolGetUserCollectedArticleList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserCollectedArticleList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func schoolGetUserArticleCommentList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserArticleCommentList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func schoolCollectArticle(c *gin.Context) {
	userId, _ := c.Get("user_id")
	ArticleId := c.PostForm("article_id")
	articleId, _ := strconv.Atoi(ArticleId)
	resp := addCollectedArticle(userId.(int), articleId)
	c.JSON(resp.Status, resp.Data)
}

func schoolCancelCollectArticle(c *gin.Context) {
	userId, _ := c.Get("user_id")
	ArticleId := c.PostForm("article_id")
	articleId, _ := strconv.Atoi(ArticleId)
	resp := removeCollectedArticle(userId.(int), articleId)
	c.JSON(resp.Status, resp.Data)
}

func schoolValueArticle(c *gin.Context) {
	ArticleCommentId := c.PostForm("article_comment_id")
	articleCommentId, _ := strconv.Atoi(ArticleCommentId)
	Value := c.PostForm("value")
	value, _ := strconv.Atoi(Value)
	resp := valueArticleComment(articleCommentId, value)
	c.JSON(resp.Status, resp.Data)
}

func bbsPublishTopic(c *gin.Context) {
	userId, _ := c.Get("user_id")
	Content := c.PostForm("content")
	resp := publishTopic(userId.(int), Content)
	c.JSON(resp.Status, resp.Data)
}

func bbsPublishTopicLordReply(c *gin.Context) {
	userId, _ := c.Get("user_id")
	TopicId := c.PostForm("topic_id")
	topicId, _ := strconv.Atoi(TopicId)
	Content := c.PostForm("content")
	resp := publishTopicLordReply(userId.(int), topicId, Content)
	c.JSON(resp.Status, resp.Data)
}

func bbsPublishTopicLayerReply(c *gin.Context) {
	userId, _ := c.Get("user_id")
	TopicLordReplyId := c.PostForm("topic_lord_reply_id")
	topicLordReplyId, _ := strconv.Atoi(TopicLordReplyId)
	Content := c.PostForm("content")
	resp := publishTopicLayerReply(userId.(int), topicLordReplyId, Content)
	c.JSON(resp.Status, resp.Data)
}

func bbsRemoveTopic(c *gin.Context) {
	TopicId := c.PostForm("topic_id")
	topicId, _ := strconv.Atoi(TopicId)
	resp := removeTopic(topicId)
	c.JSON(resp.Status, resp.Data)
}

func bbsRemoveTopicLordReply(c *gin.Context) {
	TopicLordReplyId := c.PostForm("topic_lord_reply_id")
	topicLordReplyId, _ := strconv.Atoi(TopicLordReplyId)
	resp := removeTopicLordReply(topicLordReplyId)
	c.JSON(resp.Status, resp.Data)
}

func bbsGetLatestTopicList(c *gin.Context) {
	TopicList := c.Query("topic_id_list")
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getLatestTopicList(TopicList, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func bbsGetTopic(c *gin.Context) {
	userId, _ := c.Get("user_id")
	TopicId := c.Query("topic_id")
	topicId, _ := strconv.Atoi(TopicId)
	resp := getTopic(userId.(int), topicId)
	c.JSON(resp.Status, resp.Data)
}

func bbsGetTopicLordReplyList(c *gin.Context) {
	TopicId := c.Query("topic_id")
	topicId, _ := strconv.Atoi(TopicId)
	BeginFloor := c.Query("begin_floor")
	beginFloor, _ := strconv.Atoi(BeginFloor)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getTopicLordReplyList(topicId, beginFloor, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func bbsGetTopicLayerReplyList(c *gin.Context) {
	TopicLordReplyId := c.Query("topic_lord_reply_id")
	topicLordReplyId, _ := strconv.Atoi(TopicLordReplyId)
	BeginFloor := c.Query("begin_floor")
	beginFloor, _ := strconv.Atoi(BeginFloor)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getTopicLayerReplyList(topicLordReplyId, beginFloor, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func bbsValueTopic(c *gin.Context) {
	TopicId := c.PostForm("topic_id")
	topicId, _ := strconv.Atoi(TopicId)
	Value := c.PostForm("value")
	value, _ := strconv.Atoi(Value)
	resp := valueTopic(topicId, value)
	c.JSON(resp.Status, resp.Data)
}

func bbsValueTopicLordReply(c *gin.Context) {
	TopicLordReplyId := c.PostForm("topic_lord_reply_id")
	topicLordReplyId, _ := strconv.Atoi(TopicLordReplyId)
	Value := c.PostForm("value")
	value, _ := strconv.Atoi(Value)
	resp := valueTopicLordReply(topicLordReplyId, value)
	c.JSON(resp.Status, resp.Data)
}

func bbsValueTopicLayerReply(c *gin.Context) {
	TopicLayerReplyId := c.PostForm("topic_layer_reply_id")
	topicLayerReplyId, _ := strconv.Atoi(TopicLayerReplyId)
	Value := c.PostForm("value")
	value, _ := strconv.Atoi(Value)
	resp := valueTopicLayerReply(topicLayerReplyId, value)
	c.JSON(resp.Status, resp.Data)
}

func bbsCollectTopic(c *gin.Context) {
	userId, _ := c.Get("user_id")
	TopicId := c.PostForm("topic_id")
	topicId, _ := strconv.Atoi(TopicId)
	resp := addCollectedTopic(userId.(int), topicId)
	c.JSON(resp.Status, resp.Data)
}

func bbsCancelCollectTopic(c *gin.Context) {
	userId, _ := c.Get("user_id")
	TopicId := c.PostForm("topic_id")
	topicId, _ := strconv.Atoi(TopicId)
	resp := removeCollectedTopic(userId.(int), topicId)
	c.JSON(resp.Status, resp.Data)
}

func bbsSearchTopic(c *gin.Context) {
	SearchContent := c.Query("content")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getSearchTopicList(SearchContent, beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func bbsGetUserCollectedTopicList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserCollectedTopicList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func bbsGetUserPublishedTopicList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserPublishedTopicList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func bbsGetUserTopicReplyList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserReplyList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func homeCheckInUser(c *gin.Context) {
	userId, _ := c.Get("user_id")
	resp := checkinUser(userId.(int))
	c.JSON(resp.Status, resp.Data)
}

func homeGetUserFamilyMemberList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	resp := getUserFamilyMemberList(userId.(int))
	c.JSON(resp.Status, resp.Data)
}

func homeLinkNewFamilyMember(c *gin.Context) {
	userId, _ := c.Get("user_id")
	CallName := c.PostForm("call_name")
	PhoneNumber := c.PostForm("phone_number")
	Code := c.PostForm("code")
	resp := linkNewFamilyMember(userId.(int), CallName, PhoneNumber, Code)
	c.JSON(resp.Status, resp.Data)
}

func homeRecordBloodSugar(c *gin.Context) {
	userId, _ := c.Get("user_id")
	Period := c.PostForm("period")
	BloodSugarValue := c.PostForm("blood_sugar_value")
	RecordTime := c.PostForm("record_time")
	RecordDate := c.PostForm("record_date")
	recordDate := utils.DateTimeParser(RecordDate)
	resp := recordBloodSugar(userId.(int), BloodSugarValue, Period, RecordTime, recordDate)
	c.JSON(resp.Status, resp.Data)
}

func homeGetBloodSugarRecordList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getBloodSugarRecordList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func homeGetBloodSugarRecord(c *gin.Context) {
	userId, _ := c.Get("user_id")
	RecordDate := c.Query("record_date")
	recordDate := utils.DateTimeParser(RecordDate)
	resp := getBloodSugarRecord(userId.(int), recordDate)
	c.JSON(resp.Status, resp.Data)
}

func homeRecordHealth(c *gin.Context) {
	userId, _ := c.Get("user_id")
	Insulin := c.PostForm("insulin")
	SportTime := c.PostForm("sport_time")
	Weight := c.PostForm("weight")
	BloodPressure := c.PostForm("blood_pressure")
	RecordTime := c.PostForm("record_time")
	RecordDate := c.PostForm("record_date")
	recordDate := utils.DateTimeParser(RecordDate)
	resp := recordHealth(userId.(int), Insulin, SportTime, Weight, BloodPressure, RecordTime, recordDate)
	c.JSON(resp.Status, resp.Data)

}

func homeGetHealthRecordList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getHealthRecordList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func homeParseBloodSugarRecordVoiceInput(c *gin.Context) {
	audio := c.PostForm("audio")
	resp := parseBloodSugarRecordVoiceInput(audio)
	c.JSON(resp.Status, resp.Data)
}

func homeParseHealthRecordVoiceInput(c *gin.Context) {
	audio := c.PostForm("audio")
	resp := parseHealthRecordVoiceInput(audio)
	c.JSON(resp.Status, resp.Data)
}

func homeSugarGuideWebsocket(c *gin.Context) {
	sugarGuideWebsocket(c)
}

func homeWeeklyNewspaper(c *gin.Context) {
	userId, _ := c.Get("user_id")
	resp := getWeeklyNewspaper(userId.(int))
	fmt.Println(resp)
	if resp.Status != http.StatusOK {
		c.JSON(resp.Status, resp.Data)
	} else {
		if resp.Data["code"].(int) != 0 {
			c.JSON(resp.Status, resp.Data)
		} else {
			plan := resp.Data["data"]
			cerealsValue := plan.(gin.H)["diet"].(gin.H)["cereals"].(int)
			fruitValue := plan.(gin.H)["diet"].(gin.H)["fruit"].(int)
			meatValue := plan.(gin.H)["diet"].(gin.H)["meat"].(int)
			milkValue := plan.(gin.H)["diet"].(gin.H)["milk"].(int)
			fatValue := plan.(gin.H)["diet"].(gin.H)["fat"].(int)
			vegetablesValue := plan.(gin.H)["diet"].(gin.H)["vegetables"].(int)
			c.HTML(http.StatusOK, "newspaper.tmpl", gin.H{
				// diet
				"cerealsValue": cerealsValue,
				"cereals":      fmt.Sprintf("%d%s%d%s", cerealsValue*25, "克大米, 麦片，面食， 面包, ", cerealsValue*50, "克马铃薯"),
				"fruitValue":   fruitValue,
				"fruit": fmt.Sprintf("%d%s%d%s%d%s%d%s", fruitValue*500, "克西瓜, +", fruitValue*300, "克草莓，",
					fruitValue*200, "克葡萄，橙子，橘子", fruitValue*150, "克香蕉,荔枝"),
				"meatValue": meatValue,
				"meat": fmt.Sprintf("%d%s%d%s%d%s%d%s%d%s%d%s", meatValue*25, "克大豆，", meatValue*20, "克腐竹，",
					meatValue*60, "克鸡蛋，", meatValue*50, "克鸭肉,猪肉，", meatValue*80, "克草鱼，", meatValue*100, "克鲫鱼"),
				"milkValue": milkValue,
				"milk":      fmt.Sprintf("%d%s%d%s", milkValue*20, "克奶粉,", milkValue*160, "克牛奶，羊奶"),
				"fatValue":  fatValue,
				"fat": fmt.Sprintf("%s%d%s%d%s", "（一汤勺为准）", fatValue*10, "克花生油，豆油，黄油，菜籽油,",
					fatValue*15, "克核桃,杏仁,花生"),
				"vegetablesValue": vegetablesValue,
				"vegetables": fmt.Sprintf("%d%s%d%s%d%s%d%s", vegetablesValue*500, "克白菜，韭菜，西红柿，冬瓜，茄子，丝瓜,",
					vegetablesValue*200, "克胡萝卜,", vegetablesValue*150, "克山药，", vegetablesValue*70, "克毛豆"),
				//sport
				"sport1": plan.(gin.H)["sport"].(gin.H)["sport1"],
				"sport2": plan.(gin.H)["sport"].(gin.H)["sport2"],
				"sport3": plan.(gin.H)["sport"].(gin.H)["sport3"],
				"sport4": plan.(gin.H)["sport"].(gin.H)["sport4"],
				"time1":  plan.(gin.H)["sport"].(gin.H)["time1"],
				"time2":  plan.(gin.H)["sport"].(gin.H)["time2"],
				"time3":  plan.(gin.H)["sport"].(gin.H)["time3"],
				"time4":  plan.(gin.H)["sport"].(gin.H)["time4"],
				"week1":  plan.(gin.H)["sport"].(gin.H)["week1"],
				"week2":  plan.(gin.H)["sport"].(gin.H)["week2"],
				"week3":  plan.(gin.H)["sport"].(gin.H)["week3"],
				"week4":  plan.(gin.H)["sport"].(gin.H)["week4"],
				// control
				"min1":   fmt.Sprintf("%.1f", plan.(gin.H)["control"].(gin.H)["min1"].(float64)),
				"min2":   fmt.Sprintf("%.1f", plan.(gin.H)["control"].(gin.H)["min2"].(float64)),
				"max1":   fmt.Sprintf("%.1f", plan.(gin.H)["control"].(gin.H)["max1"].(float64)),
				"max2":   fmt.Sprintf("%.1f", plan.(gin.H)["control"].(gin.H)["max2"].(float64)),
				"sleep1": fmt.Sprintf("%.1f", plan.(gin.H)["control"].(gin.H)["sleep1"].(float64)),
				"sleep2": fmt.Sprintf("%.1f", plan.(gin.H)["control"].(gin.H)["sleep2"].(float64)),
				// level
				"level0": plan.(gin.H)["level"].(gin.H)["0"],
				"level1": plan.(gin.H)["level"].(gin.H)["1"],
				"level2": plan.(gin.H)["level"].(gin.H)["2"],
				"level3": plan.(gin.H)["level"].(gin.H)["3"],
				"level4": plan.(gin.H)["level"].(gin.H)["4"],
				"level5": plan.(gin.H)["level"].(gin.H)["5"],
				"level6": plan.(gin.H)["level"].(gin.H)["6"],
			})
		}
	}
}

func socialSendMessageUser(c *gin.Context) {
	userId, _ := c.Get("user_id")
	Content := c.PostForm("content")
	TargetUserId := c.PostForm("target_user_id")
	targetUserId, _ := strconv.Atoi(TargetUserId)
	resp := sendMessageToUser(userId.(int), Content, targetUserId)
	c.JSON(resp.Status, resp.Data)
}

func socialGetLatestMessageU2uList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	TargetUserId := c.Query("target_user_id")
	targetUserId, _ := strconv.Atoi(TargetUserId)
	LatestMessageId := c.Query("latest_message_id")
	latestMessageId, _ := strconv.Atoi(LatestMessageId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getLatestMessageU2uList(userId.(int), targetUserId, latestMessageId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func socialGetHistoryMessageU2uList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	TargetUserId := c.Query("target_user_id")
	targetUserId, _ := strconv.Atoi(TargetUserId)
	OldestMessageId := c.Query("oldest_message_id")
	oldestMessageId, _ := strconv.Atoi(OldestMessageId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getHistoryMessageU2uList(userId.(int), targetUserId, oldestMessageId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func socialGetUserJoinGroupList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	BeginId := c.Query("begin_id")
	beginId, _ := strconv.Atoi(BeginId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getUserJoinGroupList(userId.(int), beginId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func socialCreateGroup(c *gin.Context) {
	userId, _ := c.Get("user_id")
	GroupName := c.PostForm("group_name")
	GroupMembers := c.PostForm("group_members")
	resp := createGroup(userId.(int), GroupName, GroupMembers)
	c.JSON(resp.Status, resp.Data)
}

func socialSendMessageInGroup(c *gin.Context) {
	userId, _ := c.Get("user_id")
	GroupId := c.PostForm("group_id")
	groupId, _ := strconv.Atoi(GroupId)
	Content := c.PostForm("content")
	resp := sendMessageInGroup(userId.(int), groupId, Content)
	c.JSON(resp.Status, resp.Data)
}

func socialGetHistoryMessageInGroupList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	GroupId := c.Query("group_id")
	groupId, _ := strconv.Atoi(GroupId)
	OldestMessageId := c.Query("oldest_message_id")
	oldestMessageId, _ := strconv.Atoi(OldestMessageId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getHistoryMessageInGroupList(userId.(int), groupId, oldestMessageId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func socialGetLatestMessageInGroupList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	GroupId := c.Query("group_id")
	groupId, _ := strconv.Atoi(GroupId)
	LatestMessageId := c.Query("latest_message_id")
	latestMessageId, _ := strconv.Atoi(LatestMessageId)
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getLatestMessageInGroupList(userId.(int), groupId, latestMessageId, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func socialGetMessageList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	ExistList := c.Query("exist_list")
	NeedNumber := c.Query("need_number")
	needNumber, _ := strconv.Atoi(NeedNumber)
	resp := getMessageList(userId.(int), ExistList, needNumber)
	c.JSON(resp.Status, resp.Data)
}

func socialGetUserInGroup(c *gin.Context) {
	userId, _ := c.Get("user_id")
	GroupId := c.Query("group_id")
	groupId, _ := strconv.Atoi(GroupId)
	resp := getUserListInGroup(userId.(int), groupId)
	c.JSON(resp.Status, resp.Data)
}

func socialMemberLevelGroup(c *gin.Context) {
	GroupId := c.PostForm("group_id")
	groupId, _ := strconv.Atoi(GroupId)
	MemberId := c.PostForm("member_id")
	memberId, _ := strconv.Atoi(MemberId)
	resp := removeMemberInGroup(groupId, memberId)
	c.JSON(resp.Status, resp.Data)
}

func socialRemoveGroup(c *gin.Context) {
	userId, _ := c.Get("user_id")
	GroupId := c.PostForm("group_id")
	groupId, _ := strconv.Atoi(GroupId)
	resp := removeGroup(userId.(int), groupId)
	c.JSON(resp.Status, resp.Data)
}

func socialRecommendUser(c *gin.Context) {
	userId, _ := c.Get("user_id")
	resp := recommendUser(userId.(int))
	c.JSON(resp.Status, resp.Data)
}
