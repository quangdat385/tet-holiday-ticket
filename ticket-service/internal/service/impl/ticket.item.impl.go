package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/quangdat385/holiday-ticket/ticket-service/global"
	consts "github.com/quangdat385/holiday-ticket/ticket-service/internal/const"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
	"go.uber.org/zap"
)

type sTicketItem struct {
	// implementation interface here
	r                *database.Queries
	distributedCache service.IRedisCache
	localCache       service.ILocalCache // Video Go 39: Add local cache
	// Add a default return for unhandled cases
}

func NewTicketItemImpl(r *database.Queries, redisCache service.IRedisCache, localCache service.ILocalCache) *sTicketItem {
	return &sTicketItem{
		r:                r,
		distributedCache: redisCache,
		localCache:       localCache,
	}
}
func (s *sTicketItem) CreateTicketItem(ctx context.Context, input model.TicketItemInPut) (out model.TicketItemsOutput, err error) {
	// 1. create ticket item in mysql
	result, err := s.r.InsertTicketItem(ctx, database.InsertTicketItemParams{
		Name:            input.TicketName,
		Description:     sql.NullString{String: input.Description, Valid: input.Description != ""},
		TicketID:        int64(input.TicketId),
		TrainID:         int64(input.TrainID),
		SeatClass:       database.PreGoTicketItem99999SeatClass(input.SeatClass),
		StockInitial:    int32(input.StockInitial),
		StockAvailable:  int32(input.StockAvailable),
		DepartureTime:   input.DepartureTime,
		IsStockPrepared: input.IsStockPrepared,
		PriceOriginal:   fmt.Sprintf("%f", input.PriceOriginal),
		PriceFlash:      fmt.Sprintf("%f", input.PriceFlash),
		SaleStartTime:   input.SaleStartTime,
		SaleEndTime:     input.SaleEndTime,
		Status:          int32(input.Status),
		ActivityID:      int64(input.ActivityId),
	})

	if err != nil {
		return out, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil || lastInsertId == 0 {
		return out, nil
	}
	ticketItem, err := s.r.GetTicketItemById(ctx, lastInsertId)
	if err != nil {
		if err == sql.ErrNoRows {
			return out, err
		}
		return out, fmt.Errorf("get ticket item by id failed: %w", err)
	}
	out = mapper.ToTicketItemDTO(ticketItem)
	return out, nil
}
func (s *sTicketItem) UpdateTicketItem(ctx context.Context, input model.UpdateTicketItemInPut) (out model.TicketItemsOutput, err error) {
	// 1. find ticket item by id
	ticketItem, err := s.r.GetTicketItemById(ctx, int64(input.TicketItemId))
	if err != nil {
		return out, err
	}
	if ticketItem.ID == 0 {
		return out, fmt.Errorf("%d with id = %d", response.FoundTicketErrCodeStatus, input.TicketItemId)
	}

	// 2. check input data if not empty, then update ticket item
	if input.TicketName != "" {
		ticketItem.Name = input.TicketName
	}
	if input.Description != "" {
		ticketItem.Description = sql.NullString{String: input.Description, Valid: true}
	}
	if input.TicketId != 0 {
		ticketItem.TicketID = int64(input.TicketId)
	}
	if input.TrainID != 0 {
		ticketItem.TrainID = int64(input.TrainID)
	}
	if input.SeatClass != "" {
		ticketItem.SeatClass = database.PreGoTicketItem99999SeatClass(input.SeatClass)
	}
	if input.StockInitial != 0 {
		ticketItem.StockInitial = int32(input.StockInitial)
	}
	if input.StockAvailable != 0 {
		ticketItem.StockAvailable = int32(input.StockAvailable)
	}
	if !input.DepartureTime.IsZero() {
		ticketItem.DepartureTime = input.DepartureTime
	}
	if input.IsStockPrepared {
		ticketItem.IsStockPrepared = input.IsStockPrepared
	}
	if input.PriceOriginal != 0 {
		ticketItem.PriceOriginal = fmt.Sprintf("%f", input.PriceOriginal)
	}
	if input.PriceFlash != 0 {
		ticketItem.PriceFlash = fmt.Sprintf("%f", input.PriceFlash)
	}
	if !input.SaleStartTime.IsZero() {
		ticketItem.SaleStartTime = input.SaleStartTime
	}
	if !input.SaleEndTime.IsZero() {
		ticketItem.SaleEndTime = input.SaleEndTime
	}
	if input.Status != 0 {
		ticketItem.Status = int32(input.Status)
	}
	if input.ActivityId != 0 {
		ticketItem.ActivityID = int64(input.ActivityId)
	}
	// 3. update ticket item in mysql
	result, err := s.r.UpdateTicketItem(ctx, database.UpdateTicketItemParams{
		ID:              int64(input.TicketItemId),
		Name:            ticketItem.Name,
		Description:     ticketItem.Description,
		TicketID:        ticketItem.TicketID,
		TrainID:         ticketItem.TrainID,
		SeatClass:       ticketItem.SeatClass,
		StockInitial:    ticketItem.StockInitial,
		StockAvailable:  ticketItem.StockAvailable,
		DepartureTime:   ticketItem.DepartureTime,
		IsStockPrepared: ticketItem.IsStockPrepared,
		PriceOriginal:   ticketItem.PriceOriginal,
		PriceFlash:      ticketItem.PriceFlash,
		SaleStartTime:   ticketItem.SaleStartTime,
		SaleEndTime:     ticketItem.SaleEndTime,
		Status:          ticketItem.Status,
		ActivityID:      ticketItem.ActivityID,
	})
	if err != nil {
		return out, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil || lastInsertId == 0 {
		return out, err
	}
	out = mapper.ToTicketItemDTO(ticketItem)
	// 4. update cache
	return out, nil
}

func (s *sTicketItem) GetTicketItemById(ctx context.Context, ticketId int, version int) (out model.TicketItemsOutput, err error) {

	// 1. get ticket item from local cache
	fmt.Println("START GET TICKET >>>>>> WITH TICKETID -> | ", ticketId)
	out, err = s.getTicketItemFromLocalCache(ctx, ticketId) // version...
	if err != nil {
		return out, fmt.Errorf("%w with id = %d -> err: %w", response.ErrCouldNotGetTicketErr, ticketId, err)
	}
	if (out != model.TicketItemsOutput{}) {
		if version == 0 || version <= out.Version {
			return out, nil
		}
	}
	// 1 get cache from distributed cache
	out, err = s.getTicketItemFromDistributedCache(ctx, ticketId)
	if err != nil {
		return out, fmt.Errorf("%w with id = %d -> err: %w", response.ErrCouldNotGetTicketErr, ticketId, err)
	}

	if (out != model.TicketItemsOutput{}) {
		return out, nil
	}

	// out, err = s.getTicketItemFromDatabase(ctx, ticketId)
	out, err = s.getTicketItemFromDatabaseLock(ctx, ticketId)
	if err != nil {
		return out, fmt.Errorf("%w with id = %d -> err: %w", response.ErrCouldNotGetTicketErr, ticketId, err)
	}
	// fmt.Println("11 -RESPONSE TICKET ITEM MYSQL -> CHECK DATA TICKET WITH ID -> ", ticketId)
	// global.Logger.Info("11 -RESPONSE TICKET ITEM MYSQL+++")
	return out, nil
}

func (s *sTicketItem) getTicketItemFromDatabaseLock(ctx context.Context, ticketId int) (out model.TicketItemsOutput, err error) {
	lockKey := "lock:ticketItem:" + strconv.Itoa(ticketId)
	// 1. lock
	err = s.distributedCache.WithDistributedLock(ctx, lockKey, 5, func(ctx context.Context) error {
		global.Logger.Info("LOCK ACQUIRED → QUERY DATABASE →", zap.Any("ticketId", ticketId))

		// 2. get data from mysql
		ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketId))
		if err != nil {
			return err
		}
		out = mapper.ToTicketItemDTO(ticketItem)
		out.Version = int(time.Now().UnixMilli()) // set version to current time
		ticketItemCacheJSON, err := json.Marshal(out)
		if err != nil {
			return fmt.Errorf("marshal json failed: %w", err)
		}

		err = global.Rdb.Set(
			ctx, s.getKeyTicketItemCache(ticketId),
			ticketItemCacheJSON,
			time.Duration(consts.TICKET_CACHE_DURATION)*time.Minute,
		).Err()
		if err != nil {
			return fmt.Errorf("save redis failed: %w", err)
		}

		isSuccess := s.localCache.SetWithTTL(ctx, s.getKeyTicketItemCache(ticketId), ticketItem)
		if !isSuccess {
			return fmt.Errorf("save local cache failed")
		}
		return nil
	})

	return
}

// get data from database
func (s *sTicketItem) DecreaseStock(ctx context.Context, ticketId int, quantity int) (out int, code int) {
	newStock, errCode := s.decreaseStockCacheByLuaScript(ctx, ticketId, quantity)
	if errCode == 1 {
		return out, 1
	}
	if errCode == 2 {
		return out, 2
	}

	result, err := s.r.UpdateTicketItemStock(ctx, database.UpdateTicketItemStockParams{
		ID:               int64(ticketId),
		StockAvailable:   int32(newStock),
		StockAvailable_2: int32(newStock + quantity),
	})
	if err != nil {
		return out, 1
	}
	_, lastErr := result.LastInsertId()
	if lastErr != nil {
		return out, 1
	}
	return newStock, 3
}

// get data from redis distributed

func (s *sTicketItem) getTicketItemFromDistributedCache(ctx context.Context, ticketId int) (out model.TicketItemsOutput, err error) {

	// ticketItemCache, err := global.Rdb.Get(ctx, s.getKeyTicketItemCache(ticketId)).Result()
	// if err != nil {
	// 	if errors.Is(err, redis.Nil) {
	// 		// Trả về lỗi riêng khi key không tồn tại
	// 		return out, nil
	// 	}
	// 	return out, fmt.Errorf("failed to get ticket item cache: %v", err)
	// }
	fmt.Println("04 - DISTRIBUTED CACHE -> CHECK DATA TICKET WITH ID -> ", ticketId)

	ticketItemCache, err := s.distributedCache.Get(ctx, s.getKeyTicketItemCache(ticketId))
	if err != nil {
		return out, fmt.Errorf("failed to get ticket item cache: %v", err)
	}
	if ticketItemCache == "" {
		fmt.Println("05 - DISTRIBUTED CACHE: NOT FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId)
		return out, nil
	}

	if err := json.Unmarshal([]byte(ticketItemCache), &out); err != nil {
		return out, fmt.Errorf("parse redis data failed: %v", err)
	}
	fmt.Println("06 - DISTRIBUTED CACHE: FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId, ticketItemCache)

	// put to local cache
	s.localCache.SetWithTTL(ctx, s.getKeyTicketItemCache(ticketId), out)
	return out, nil
}

// get data from local cache
func (s *sTicketItem) getTicketItemFromLocalCache(ctx context.Context, ticketId int) (out model.TicketItemsOutput, err error) {
	// global.Logger.Info("getTicketItemFromLocalCache with ticketId: ", zap.Int("TicketId", ticketId))
	// var out model.TicketItemsOutput
	fmt.Println("01 - LOCAL CACHE -> CHECK DATA TICKET WITH ID -> ", ticketId)

	ticketItemLocalCache, isFound := s.localCache.Get(ctx, s.getKeyTicketItemCache(ticketId))
	if !isFound {
		// global.Logger.Info("getTicketItemFromLocalCache with ticketId is notfound: ", zap.Int("TicketId", ticketId))
		// fmt.Println("getTicketItemFromLocalCache with ticketId is notfound: ", ticketId)
		fmt.Println("02 - LOCAL CACHE: NOT FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId)
		return out, nil
	}
	// fmt.Println(">>>", ticketItemLocalCache)
	// fmt.Printf(">>> Value: %+v, Type: %s\n", ticketItemLocalCache, reflect.TypeOf(ticketItemLocalCache))

	fmt.Println("03 - LOCAL CACHE: FOUND -> CHECK DATA TICKET WITH ID -> ", ticketId)

	// // Type Assertion to string
	jsonTicketString, ok := ticketItemLocalCache.(string)
	if !ok {
		fmt.Printf("ERROR: Local cache item with key %d is not a string\n", ticketId)
	}

	if err := json.Unmarshal([]byte(jsonTicketString), &out); err != nil {
		return out, fmt.Errorf("parse redis data failed: %v", err)
	}

	return out, nil
}

// generate key cache
func (s *sTicketItem) getKeyTicketItemCache(ticketId int) string {
	return "PRO_TICKET:ITEM:" + strconv.Itoa(ticketId)
}
func (s *sTicketItem) getKeyStockItemCahe(ticketId int) string {
	return "TICKET::" + strconv.Itoa(ticketId) + "::STOCK"
}

func (s *sTicketItem) decreaseStockCacheByLuaScript(ctx context.Context, ticketId int, quantity int) (number int, codeStatus int) {
	// Lua script to atomically decrease the stock
	luaScript := `
		local currentStock = redis.call("GET", KEYS[1])
		if currentStock then
			currentStock = tonumber(currentStock)
			if currentStock >= ARGV[1] then
				local newStock = currentStock - ARGV[1]
				redis.call("SET", KEYS[1], newStock)
				return newStock
			else
				return -1 -- Not enough stock
			end
		else
			return -1 -- Key does not exist
		end
	`

	// Execute the Lua script
	result, err := global.Rdb.Eval(ctx, luaScript, []string{s.getKeyStockItemCahe(ticketId)}, quantity).Result()
	if err != nil {
		return 0, 1
	}

	newStock, ok := result.(int64)
	if !ok {
		return 0, 1
	}

	if newStock == -1 {
		return 0, 2
	}
	return int(newStock), 3
}
func (s *sTicketItem) DeleteTicketItem(ctx context.Context, ticketItemId int) (err error) {
	// 1. find ticket item by id
	ticketItem, err := s.r.GetTicketItemById(ctx, int64(ticketItemId))
	if err != nil {
		return err
	}
	if ticketItem.ID == 0 {
		return fmt.Errorf("%d with id = %d", response.FoundTicketErrCodeStatus, ticketItemId)
	}
	// 2. delete ticket item in mysql
	_, err = s.r.DeleteTicketItem(ctx, int64(ticketItemId))
	if err != nil {
		return fmt.Errorf("delete ticket item failed: %w", err)
	}
	return nil
}
