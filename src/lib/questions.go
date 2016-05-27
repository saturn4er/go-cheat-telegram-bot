package lib

import (
	"strings"
	"path/filepath"
	"fmt"
	"regexp"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
)

const (
	QuestionDataText int = iota
	QuestionDataPhoto
	QuestionDataDocument
)

type QuestionData struct {
	Type int
	Data string
}
type QuestionInfo struct {
	Name string
	Hash string
	Data []QuestionData
}

func (bi *QuestionInfo) MatchString(str string) bool {
	if strings.Contains(strings.ToLower(bi.Name), strings.ToLower(str)) {
		return true
	}
	for _, value := range bi.Data {
		if value.Type == QuestionDataText {
			if strings.Contains(strings.ToLower(value.Data), strings.ToLower(str)) {
				return true
			}
		}
	}
	return false
}
func (bi *QuestionInfo) ToMessages(chatId int64) []tgbotapi.Chattable {
	result := []tgbotapi.Chattable{}
	result = append(result, tgbotapi.NewMessage(chatId, fmt.Sprintf("Название: %s", bi.Name)))
	for _, value := range bi.Data {
		switch value.Type {
		case QuestionDataText:
			result = append(result, tgbotapi.NewMessage(chatId, value.Data))
		case QuestionDataPhoto:
			result = append(result, (tgbotapi.NewPhotoShare(chatId, value.Data)))
		case QuestionDataDocument:
			result = append(result, (tgbotapi.NewDocumentShare(chatId, value.Data)))
		}
	}
	return result
}

func GetQuestionsList() ([]QuestionInfo, error) {
	result := []QuestionInfo{}
	infos, err := filepath.Glob("./questions/*/info.json")
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	if len(infos) == 0 {
		return result, nil
	}
	regxp := regexp.MustCompile("questions/(.+?)/info\\.json")
	for _, value := range infos {
		file, err := ioutil.ReadFile(value)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		info := QuestionInfo{}
		err = json.Unmarshal([]byte(file), &info)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		hash := regxp.FindAllStringSubmatch(value, -1)[0][1]
		info.Hash = hash
		result = append(result, info)
	}
	return result, nil
}