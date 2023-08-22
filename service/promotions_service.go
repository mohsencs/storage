package service

import (
	"context"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"storage/model"
	"storage/repo"
	"storage/utils"
	"strconv"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)

type service struct {
	repo  repo.PromotionRepository
	cache *bigcache.BigCache
}

func NewPromotionService(repo repo.PromotionRepository) PromotionService {
	cache, _ := bigcache.New(context.Background(), getCacheConfig())

	s := service{
		repo:  repo,
		cache: cache,
	}

	go printCacheStats(cache)

	return &s
}

func (s *service) UploadPromotion(ctx *gin.Context, file *multipart.FileHeader) {
	// go updateNewPromotions()

	f, err := file.Open()
	if err != nil {
		utils.PrintError(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer f.Close()

	var promotions []model.Promotion

	err = gocsv.UnmarshalWithoutHeaders(f, &promotions)
	if err != nil {
		utils.PrintError(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed when fetch data from file. "})
		return
	}

	s.repo.RemoveAllPromotions()
	s.resetCache()

	chunkSize := 10000
	chunkSizeEnv := os.Getenv("BULK_INSERT_MAX_SIZE")
	if chunkSizeEnv != "" {
		chunkSize, _ = strconv.Atoi(chunkSizeEnv)
	}
	for i := 0; i < len(promotions); i += chunkSize {
		end := i + chunkSize
		if end > len(promotions) {
			end = len(promotions)
		}

		err = s.repo.UploadPromotion(ctx, promotions[i:end])
		if err != nil {
			utils.PrintError(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed when insert to db. "})
			return
		}
	}

}

func updateNewPromotions() {
	log.Print("starting to inserte new input csv to db.")

}

func (s *service) GetPromotionById(ctx *gin.Context, id string) (*model.Promotion, error) {

	if value := s.cacheGet(id); value != nil {
		return value, nil
	}

	promotino, err := s.repo.GetById(ctx.Request.Context(), id)
	if err != nil {
		utils.PrintError(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed in get promotino with id: " + id})
		return nil, err
	}

	s.cacheSet(id, promotino)
	return promotino, nil
}

func getCacheConfig() bigcache.Config {
	config := bigcache.DefaultConfig(10 * time.Minute)
	config.MaxEntriesInWindow = 10000
	config.HardMaxCacheSize = 32 // MB
	return config
}

func (s *service) cacheGet(key string) *model.Promotion {
	if value, err := s.cache.Get(key); err == nil {
		var prom model.Promotion
		json.Unmarshal(value, &prom)

		return &prom
	}

	return nil
}

func (s *service) cacheSet(key string, value *model.Promotion) {
	v, _ := json.Marshal(value)
	s.cache.Set(key, v)
}

func (s *service) resetCache() {
	if err := s.cache.Reset(); err != nil {
		utils.PrintError(err)
	}
}

func printCacheStats(cache *bigcache.BigCache) {
	for range time.Tick(time.Minute) {
		log.Printf("cache stats: %+v,	length: %d\n", cache.Stats(), cache.Len())
	}
}
