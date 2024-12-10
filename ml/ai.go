package ml

import (
	"context"
	"fmt"
	"log"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/structpb"
)

// GetEventRecommendations fetches event recommendations for a user
func GetEventRecommendations(userID string) {
	ctx := context.Background()

	client, err := aiplatform.NewPredictionClient(ctx, option.WithCredentialsFile("path/to/your-service-account.json"))
	if err != nil {
		log.Fatalf("Failed to create prediction client: %v", err)
	}
	defer client.Close()

	// Vertex AI endpoint for the deployed recommendation model
	endpoint := "projects/your-project-id/locations/us-central1/endpoints/your-endpoint-id"

	// Create a payload for the prediction request
	instance, err := structpb.NewStruct(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		log.Fatalf("Failed to create instance payload: %v", err)
	}

	req := &aiplatformpb.PredictRequest{
		Endpoint:  endpoint,
		Instances: []*structpb.Value{structpb.NewStructValue(instance)},
	}

	// Make the prediction request
	resp, err := client.Predict(ctx, req)
	if err != nil {
		log.Fatalf("Prediction request failed: %v", err)
	}

	// Print the recommendations
	for _, prediction := range resp.Predictions {
		fmt.Printf("Recommended Event: %v\n", prediction)
	}
}
