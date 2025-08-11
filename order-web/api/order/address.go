package order

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"io/ioutil"
	"localword/order-web/global"
	"localword/order-web/models/dbmodel"
	"math/big"
	"net/http"
	"os"
	"strings"
)

type Tips struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Address  string `json:"address"`
}

type TipsStruct struct {
	Status   string `json:"status"`
	Infocode string `json:"infocode"`
	Tips     []Tips `json:"tips"`
}
type WordParams struct {
	W     string `json:"w"`
	User  string `json:"user"`
	Soure string `json:"soure"`
}
type BodyWord struct {
	Bdword []string `json:"bdword"`
	Status int      `json:"status"`
}

type WordItem struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}

func WordBaidu(c *gin.Context) {
	word := c.PostForm("word")
	word = strings.TrimSpace(word)
	if word == "" {
		return
	}
	tableIndex := GetTableIndex(word)
	tableName := fmt.Sprintf("seo_word_%d", tableIndex)

	words := WordExtendFromBaidu(word)
	//fmt.Println(tableName, words)

	var seoWordDatas []dbmodel.SeoWord
	global.GlobalDB.Select("id", "title", "extend_title").Table(tableName).Where("title=?", word).Find(&seoWordDatas)
	wordData := make([]WordItem, 0)
	addWords := make([]dbmodel.SeoWord, 0)
	for _, extendword := range words {
		if extendword == word {
			continue
		}
		existens := false
		for _, seo := range seoWordDatas {
			if extendword == seo.ExtendTitle {
				existens = true
				wordData = append(wordData, WordItem{Id: seo.Id, Title: extendword})
				break
			}
		}
		if !existens {
			addWords = append(addWords, dbmodel.SeoWord{Title: word, ExtendTitle: extendword})
		}
	}
	if len(addWords) > 0 {
		global.GlobalDB.Table(tableName).Create(&addWords)
		for _, additem := range addWords {
			wordData = append(wordData, WordItem{Id: additem.Id, Title: additem.ExtendTitle})
		}
	}
	var ginres = gin.H{}
	ginres["bdword"] = wordData
	ginres["tableindex"] = tableIndex
	c.JSON(http.StatusOK, ginres)
}
func ExtendWordById(c *gin.Context) {
	id := c.PostForm("id")
	tableindex := c.PostForm("tableindex")
	if id == "" && tableindex == "" {
		return
	}
	tableIndex := cast.ToInt(tableindex)
	tableName := fmt.Sprintf("seo_word_%d", tableIndex)

	var seoWordDatas dbmodel.SeoWord
	global.GlobalDB.Select("id", "title", "extend_title").Table(tableName).Where("id=?", id).Find(&seoWordDatas)
	title := seoWordDatas.ExtendTitle
	var ginres = gin.H{}
	ginres["title"] = title
	c.JSON(http.StatusOK, ginres)
}

func containsStringInSliceField(slice []dbmodel.SeoWord, target string) bool {
	for _, item := range slice {
		if strings.Contains(item.ExtendTitle, target) {
			return true
		}
	}
	return false
}

func WordExtendFromBaidu(word string) []string {

	user := WordParams{
		W:     word,
		User:  "admin987457",
		Soure: "baidu",
	}
	body, _ := json.Marshal(user)
	fmt.Println(user)
	resp, err := http.Post("https://mysql.topfd.cc/o/v1/bdword/getwordbykey", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyAnswer, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(bodyAnswer)
		var bodyRes BodyWord
		json.Unmarshal(bodyAnswer, &bodyRes)
		//fmt.Println(bodyRes.Bdword)
		return bodyRes.Bdword
	} else {
		//fmt.Println("Get failed with error: ", resp.Status)
		return []string{}
	}

	//global.GlobalDB.Limit(1).Where("ftype=?", ftypeInt).Order("rand()").Find(&res)
	//fmt.Println(res)
	//var ginres = gin.H{}
	//ginres["randata"] = bodyRes.Bdword
	//c.JSON(http.StatusOK, ginres)
}

func GetTable(title string) string {
	//hasher := md5.New()
	//hasher.Write([]byte(title))
	//hash := hex.EncodeToString(hasher.Sum(nil))
	//// 将哈希值转化为整数并取模
	//var sum int
	//for _, c := range hash {
	//	sum += int(c)
	//}
	hash := sha256.Sum256([]byte(title))
	// 将哈希值转换为一个大整数
	bigInt := new(big.Int).SetBytes(hash[:])
	// 使用1000对大整数取模，得到表的编号（从0到999）
	tableNumber := new(big.Int).Mod(bigInt, big.NewInt(1000)).Int64()
	tableIndex := int(tableNumber) + 1
	//fmt.Println(title, tableIndex)
	tableName := fmt.Sprintf("seo_word_%d", tableIndex)
	return tableName
}
func GetTableIndex(title string) int {
	//hasher := md5.New()
	//hasher.Write([]byte(title))
	//hash := hex.EncodeToString(hasher.Sum(nil))
	//// 将哈希值转化为整数并取模
	//var sum int
	//for _, c := range hash {
	//	sum += int(c)
	//}
	hash := sha256.Sum256([]byte(title))
	// 将哈希值转换为一个大整数
	bigInt := new(big.Int).SetBytes(hash[:])
	// 使用1000对大整数取模，得到表的编号（从0到999）
	tableNumber := new(big.Int).Mod(bigInt, big.NewInt(1000)).Int64()
	tableIndex := int(tableNumber) + 1
	return tableIndex
}

func ReadPass(dirName string) []string {
	fp, err := os.Open(dirName)
	if err != nil {
		fmt.Println(err) //打开文件错误
		return nil
	}
	buf := bufio.NewScanner(fp)
	var passList []string
	for {
		if !buf.Scan() {
			break //文件读完了,退出for
		}
		line := buf.Text() //获取每一行
		passList = append(passList, line)
	}
	return passList
}
