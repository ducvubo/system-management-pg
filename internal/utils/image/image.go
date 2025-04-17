package image

import (
	"encoding/json"
	"fmt"
)

type ImageData struct {
	ImageCloud  string `json:"image_cloud"`
	ImageCustom string `json:"image_custom"`
}

func BuildImageJson(Image ImageData) (string, error) {
	jsonBytes, err := json.Marshal(Image)
	if err != nil {
		return "", fmt.Errorf("lỗi chuyển JSON: %w", err)
	}

	return string(jsonBytes), nil
}
