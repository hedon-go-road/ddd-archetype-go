package domain

import "context"

// OpenAIGateway is a required dependency for the domain to generate content by OpenAI.
type OpenAIGateway interface {
	// GenerateContent generates a content for a diary.
	GenerateContent(ctx context.Context, uid, input string) (string, error)
}
