package domain

import "github.com/YacineMK/DTQ/services/trip-service/pkg/types"

type RideFareModel struct {
	ID          string                 `bson:"id"`
	PackageSlug string                 `bson:"packageSlug"`
	Route       *types.OsrmApiResponse `bson:"route"`
	TotalPrice  float64                `bson:"totalPrice"`
}
