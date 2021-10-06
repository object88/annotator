//go:build test

package tools

//go:generate go build -o ../bin/mockgen ../vendor/github.com/golang/mock/mockgen
//go:generate mkdir -p ../mocks
//go:generate ../bin/mockgen -destination=../mocks/mock_admission_webhookprocessor.go -package=mocks github.com/object88/tugboat-annotator/pkg/admission WebhookProcessor

import (
	_ "github.com/golang/mock/gomock"
)
