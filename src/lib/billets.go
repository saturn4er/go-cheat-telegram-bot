package lib

import (
	"path/filepath"
	"fmt"
	"regexp"
	"io/ioutil"
	"encoding/json"
	"os"
)

type BilletInfo struct {
	Id        string
	Hash      string
	Questions []string
}

//func (bi *BilletInfo) ToMessages(chatId int64) []tgbotapi.Chattable {
//	result := []tgbotapi.Chattable{}
//	result = append(result, tgbotapi.NewMessage(chatId, fmt.Sprintf("Название: %s", bi.Name)))
//	for _, value := range bi.Data {
//		switch value.Type {
//		case QuestionDataText:
//			result = append(result, tgbotapi.NewMessage(chatId, value.Data))
//		case QuestionDataPhoto:
//			result = append(result, (tgbotapi.NewPhotoShare(chatId, value.Data)))
//		case QuestionDataDocument:
//			result = append(result, (tgbotapi.NewDocumentShare(chatId, value.Data)))
//		}
//	}
//	return result
//}
func (bi *BilletInfo) Save() error {
	data, err := json.Marshal(bi)
	if err != nil {
		return err
	}
	questionDir := filepath.Join("./billets/", bi.Hash)
	os.MkdirAll(questionDir, 0700)
	err = ioutil.WriteFile(filepath.Join(questionDir, "/info.json"), data, 0700)
	return err
}
func GetBilletsList() ([]BilletInfo, error) {
	result := []BilletInfo{}
	infos, err := filepath.Glob("./billets/*/info.json")
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	if len(infos) == 0 {
		return result, nil
	}
	regxp := regexp.MustCompile("billets/(.+?)/info\\.json")
	for _, value := range infos {
		file, err := ioutil.ReadFile(value)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		info := BilletInfo{}
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