package gateway

import (
	"context"
	"fmt"
)

type OpenAIGatewayImpl struct{}

func NewOpenAIGatewayImpl() *OpenAIGatewayImpl {
	return &OpenAIGatewayImpl{}
}

func (g *OpenAIGatewayImpl) GenerateContent(ctx context.Context, uid, input string) (string, error) {
	// TODO: use openai api to generate content
	return fmt.Sprintf("uid: %s\norigin: %s\ngenerated: %s",
		uid, input,
		"generated content for "+uid,
	), nil
}
