package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/go-ego/gse"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

// 由于云函数一次运行后会在后台一直运行（可能会有休眠），
// 所以可以全局加载分词字典，多次调用可直接使用字典
var s Similarity

func init() {
	s, _ = New("./dictionary.txt")
}

// DefineEvent 入参参数
type DefineEvent struct {
	// test event define
	Body string `json:"body"`
}

// RequestData 请求体结构
type RequestData struct {
	Data []struct {
		ID     int    `json:"id"`
		First  string `json:"first"`
		Second string `json:"second"`
	} `json:"data"`
}

// ResponseData 响应体结构
type ResponseData struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data []RespItem `json:"data"`
}

// RespItem 响应体data里的子结构
type RespItem struct {
	ID          int     `json:"id"`
	Probability float64 `json:"probability"`
}

// Similarity 存放字符串分词后字典
type Similarity struct {
	S1  map[string]float64
	S2  map[string]float64
	Seg gse.Segmenter
}

// New 工厂函数
func New(filePath string) (Similarity, error) {
	s := Similarity{}
	//全局加载一次字典
	err := s.Seg.LoadDict(filePath)
	if err != nil {
		return Similarity{}, errors.New("构建相似度对象失败")
	}
	return s, nil
}

// SimiCos 余弦相似度
func (s Similarity) SimiCos(s1, s2 string) (float64, error) {
	//清空s1  s2 内的数据，防止第二次调用时数据污染
	s.S1 = map[string]float64{}
	s.S2 = map[string]float64{}

	//计算两字符串分词
	s1List := s.cut(s1)
	s2List := s.cut(s2)

	//所有词汇
	s3Map := map[string]bool{}
	for _, v := range s1List {
		s3Map[v] = true
	}
	for _, v := range s2List {
		s3Map[v] = true
	}

	//	计算词频
	for _, v := range s1List {
		s.S1[v] += 1
	}
	for _, v := range s2List {
		s.S2[v] += 1
	}

	//计算余弦相似度
	var a float64
	var b float64
	var c float64

	for k, _ := range s3Map {
		a += s.S1[k] * s.S2[k]
		b += s.S1[k] * s.S1[k]
		c += s.S2[k] * s.S2[k]
	}

	var d float64
	d = a / (math.Sqrt(b) * math.Sqrt(c))
	e, err := strconv.ParseFloat(fmt.Sprintf("%.2f", d), 64)
	return e, err
}

// 分词
func (s Similarity) cut(s1 string) []string {
	text1 := []byte(s1)
	segments := s.Seg.Segment(text1)
	return gse.ToSlice(segments)
}

// ***************************************************

// Scf 入口函数
func Scf(event DefineEvent) (ResponseData, error) {
	req := &RequestData{}

	resp := ResponseData{
		Code: 2901,
		Msg:  "参数错误",
	}

	fmt.Println("输入原始数据 ->")
	fmt.Println(event)

	// 入参
	err := json.Unmarshal([]byte(event.Body), req)
	if err != nil {
		fmt.Println("json反序列入参失败->", err)
		return resp, nil
	}

	fmt.Println("解析输入内容 ->")
	fmt.Println(req)

	// 查询相似度

	for _, item := range req.Data {
		simi, _ := s.SimiCos(item.First, item.Second)
		temp := RespItem{
			ID:          item.ID,
			Probability: simi,
		}
		resp.Data = append(resp.Data, temp)
	}

	resp.Code = 2000
	resp.Msg = "成功"

	fmt.Println("输出 ->")
	fmt.Println(resp)

	return resp, nil
}

// 腾讯云函数入口
func main() {
	// Make the handler available for Remote Procedure Call by Cloud Function
	cloudfunction.Start(Scf)

	// d := DefineEvent{}
	// d.Body = `{"data":[{"id":1,"first":"我爱中国","second":"我爱祖国"}]}`
	// Scf(d)

}
