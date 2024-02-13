package inicializer

import "github.com/guiddodg/go-jwt/internal/domain/model"

func FixtureLoad() {
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		panic("error while trying to load fixtures")
	}
}
