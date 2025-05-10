package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const AdmissionRespStatusErr = "error"

// AdmissionRequest 定义请求参数结构体
type AdmissionRequest struct {
	URL            string `json:"url"`
	Year           string `json:"year"`
	Province       string `json:"province"`
	AdmissionType  string `json:"admission_type"`
	UniversityName string `json:"university_name"`
}

// AdmissionData 定义响应数据中的 data 字段子项
type AdmissionScore struct {
	Year             string `json:"year"`
	Province         string `json:"province"`
	AdmissionType    string `json:"admission_type"`
	SubjectCategory  string `json:"subject_category"`
	MajorName        string `json:"major_name"`
	HighestScore     string `json:"highest_score"`
	HighestScoreRank string `json:"highest_score_rank"`
	LowestScore      string `json:"lowest_score"`
	LowestScoreRank  string `json:"lowest_score_rank"`
}

// AdmissionResponse 定义响应参数结构体
type AdmissionResponse struct {
	Status  string           `json:"status"`
	Data    []AdmissionScore `json:"data"`
	Message string           `json:"message"`
}

// FetchAdmissionScoreData 发送 POST 请求获取招生数据
func FetchAdmissionScoreData(url string, req AdmissionRequest) (AdmissionResponse, error) {
	var resp AdmissionResponse

	// 将请求参数序列化为 JSON
	body, err := json.Marshal(req)
	if err != nil {
		return resp, err
	}

	// 创建 HTTP POST 请求
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return resp, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()

	// 解析响应
	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return resp, err
	}

	return resp, nil
}

func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}
