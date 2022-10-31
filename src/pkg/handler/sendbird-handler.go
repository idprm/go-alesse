package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/pkg/config"
	"github.com/idprm/go-alesse/src/pkg/model"
)

type User struct {
	UserId     string `json:"user_id"`
	Nickname   string `json:"nickname"`
	ProfileUrl string `json:"profile_url"`
}

type Group struct {
	Name        string   `json:"name"`
	ChannelUrl  string   `json:"channel_url"`
	CoverUrl    string   `json:"cover_url"`
	ChannelType string   `json:"channel_type"`
	CustomType  string   `json:"custom_type"`
	UserIds     []string `json:"user_ids"`
	InviterId   string   `json:"inviter_id"`
	OperatorIds []string `json:"operator_ids"`
}

type UserLeave struct {
	UserIds []string `json:"user_ids"`
}

type AutoMessage struct {
	MessageType string `json:"message_type"`
	UserId      string `json:"user_id"`
	Message     string `json:"message"`
}

type ErrorResponse struct {
	ChannelUrl string `json:"channel_url"`
	Error      bool   `json:"error"`
}

func SendbirdCreateUser(user model.User) (string, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/users"

	userId := user.Msisdn
	nickName := user.Name

	data := User{
		UserId:     userId,
		Nickname:   nickName,
		ProfileUrl: "https://www.a-lesse.com/images/logo/logo.png",
	}

	payload, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string([]byte(body)), nil
}

func SendbirdGetUser(user model.User) (string, bool, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/users/" + user.Msisdn

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", true, errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", true, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", true, errors.New(err.Error())
	}

	var errorResponse ErrorResponse
	err = json.Unmarshal(body, &errorResponse)

	return string([]byte(body)), errorResponse.Error, nil
}

func SendbirdDeleteUser(user model.User) (string, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/users/" + user.Msisdn

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string([]byte(body)), nil
}

func SendbirdCreateGroupChannel(doctor model.Doctor, user model.User) (string, string, string, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/group_channels"

	users := []string{doctor.Username, user.Msisdn}
	operators := []string{"alesse"}

	now := time.Now()
	date := now.Format("200601021504")

	group := Group{
		Name:        user.Name + " - " + doctor.Name,
		ChannelUrl:  user.Msisdn + "_" + date,
		ChannelType: "group_messaging",
		CustomType:  "chat_with_doctor",
		UserIds:     users,
		InviterId:   "alesse",
		OperatorIds: operators,
	}

	payload, _ := json.Marshal(group)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		return "", "", "", errors.New(err.Error())
	}

	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)

	if err != nil {
		return "", "", "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	type resGroup struct {
		Name       string `json:"name"`
		ChannelUrl string `json:"channel_url"`
	}

	var result resGroup

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", errors.New(err.Error())
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", "", errors.New(err.Error())
	}

	return string([]byte(body)), result.Name, result.ChannelUrl, nil
}

func SendbirdGetGroupChannel(chat model.Chat) (string, bool, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/group_channels/" + chat.ChannelUrl

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", true, errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", true, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", true, errors.New(err.Error())
	}

	var errorResponse ErrorResponse

	return string([]byte(body)), errorResponse.Error, nil
}

func SendbirdDeleteGroupChannel(channel model.Chat) (string, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/group_channels/" + channel.ChannelUrl

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string([]byte(body)), nil
}

func SendbirdLeaveGroupChannel(channel string, user string) (string, bool, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/group_channels/" + channel + "/leave"

	userLeave := UserLeave{
		UserIds: []string{user},
	}

	payload, _ := json.Marshal(userLeave)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")
	if err != nil {
		return "", true, errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", true, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", true, errors.New(err.Error())
	}

	var errorResponse ErrorResponse
	err = json.Unmarshal(body, &errorResponse)

	return string([]byte(body)), errorResponse.Error, nil
}

func SendbirdAutoMessageDoctor(channel string, doctor model.Doctor, user model.User) (string, error) {
	url := "https://api-" + config.ViperEnv("SB_APP_ID") + ".sendbird.com/v3/group_channels/" + channel + "/messages"

	var conf model.Config
	database.Datasource.DB().Where("name", "AUTO_MESSAGE_SENDBIRD").First(&conf)
	replaceMessage := strings.Replace(conf.Value, "@v1", doctor.Name, 1)

	message := AutoMessage{
		MessageType: "MESG",
		UserId:      doctor.Username,
		Message:     replaceMessage,
	}

	payload, _ := json.Marshal(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		return "", errors.New(err.Error())
	}

	req.Header.Set("Api-Token", config.ViperEnv("SB_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string([]byte(body)), nil
}
