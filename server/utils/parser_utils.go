package utils

import (
	"errors"
	"net/url"
	"strings"

	"github.com/chryscloud/video-edge-ai-proxy/models"
)

// ParseRTMPKey takes the complete RTMP url and extracts the streaming key
func ParseRTMPKey(rtmpURL string) (string, error) {
	u, err := url.Parse(rtmpURL)
	if err != nil {
		return "", err
	}
	if u.Scheme != "rtmp" {
		return "", err
	}
	splitted := strings.Split(u.Path, "/")
	if len(splitted) > 0 {
		key := splitted[len(splitted)-1]
		return key, nil

	}
	return "", errors.New("failed to parse RTMP key")
}

func StringPairsToVarPairs(stringPairs []string) []*models.VarPair {
	varPairs := make([]*models.VarPair, 0)
	if len(stringPairs) > 0 {
		for _, str := range stringPairs {
			if strings.Contains(str, "=") {
				pair := &models.VarPair{}
				split := strings.Split(str,