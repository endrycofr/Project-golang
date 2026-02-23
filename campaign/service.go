package campaign

import (
	miniostorage "bwastartup/minio"
	"context"
	"fmt"
	"io"

	minio "github.com/minio/minio-go/v7" // alias untuk hindari conflict

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
	UploadMinio(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) (string, error)
}

type service struct {
	repository Repository
	minio      *miniostorage.MinioStorage
}

func NewService(repository Repository, minio *miniostorage.MinioStorage) *service {
	return &service{repository, minio}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil

	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil

}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}
	if campaign.UserID != inputData.User.ID {
		return campaign, fmt.Errorf("unauthorized to update this campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}
	if campaign.UserID != input.User.ID {
		return CampaignImage{}, fmt.Errorf("unauthorized to update this campaign")
	}

	isPrimary := 0

	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{
		CampaignID: input.CampaignID,
		IsPrimary:  isPrimary,
		FileName:   fileLocation, // simpan URL dari MinIO
	}
	// campaignImage.CampaignID = input.CampaignID
	// campaignImage.IsPrimary = isPrimary
	// campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return newCampaignImage, nil
}

// UploadMinio mengunggah file ke MinIO dan mengembalikan URL aksesnya
func (s *service) UploadMinio(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) (string, error) {
	_, err := s.minio.Client.PutObject(ctx, s.minio.BucketName, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload object: %w", err)
	}

	filename := fmt.Sprintf("https://%s/%s/%s",
		s.minio.Client.EndpointURL().Host,
		s.minio.BucketName,
		objectName,
	)
	return filename, nil
}

//save campaign image dengan upload ke MinIO
