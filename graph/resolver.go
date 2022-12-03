package graph

import (
	"github.com/KnightHacks/knighthacks_hackathon/repository"
	"github.com/KnightHacks/knighthacks_shared/auth"
	"github.com/KnightHacks/knighthacks_shared/azure_blob"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repository      repository.Repository
	Auth            *auth.Auth
	AzureBlobClient *azure_blob.AzureBlobClient
}
