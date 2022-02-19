package services

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	g "github.com/chryscloud/video-edge-ai-proxy/globals"
	"github.com/chryscloud/video-edge-ai-proxy/models"
	"github.com/dgraph-io/badger/v2"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Printf("Setup failed: :%v\n", err.Error())
	}
	retCode := m.Run()

	err = teardown()
	if err != nil {
		fmt.Printf("Teardown failed: :%v\n", err.Error())
	}
	os.Exit(retCode)
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

var storage *Storage

func setup() error {
	db, err := badger.Open(badger.DefaultOptions