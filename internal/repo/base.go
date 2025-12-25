package repo

import (
	"context"
	"go-server-starter/pkg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
*
* BaseRepo 是基础仓库接口，定义了基础的增删改查操作
*
* note:
* - ZeroFields 结构体中的零值会被更新到数据库中
* - NonZeroFields 结构体中的零值不会被更新到数据库中
 */
type BaseRepo[T any] interface {
	// create 创建数据
	Create(ctx context.Context, entity *T) error
	// create batch 创建批量数据
	CreateBatch(ctx context.Context, entities []*T) error
	// update by map 使用map更新数据
	UpdateByMap(ctx context.Context, id uint64, entity map[string]any) error
	// update by zero fields 使用零字段更新数据
	UpdateByZeroFields(ctx context.Context, id uint64, entity *T) error
	// update by non zero fields 使用非零字段更新数据
	UpdateByNonZeroFields(ctx context.Context, id uint64, entity *T) error
	// update batch by ids with map 使用map更新批量数据
	UpdateBatchByIDsWithMap(ctx context.Context, ids []uint64, entity map[string]any) error
	// update batch by ids with zero fields 使用零字段更新批量数据
	UpdateBatchByIDsWithZeroFields(ctx context.Context, ids []uint64, entity *T) error
	// update batch by ids with non zero fields 使用非零字段更新批量数据
	UpdateBatchByIDsWithNonZeroFields(ctx context.Context, ids []uint64, entity *T) error
	// update by opts and zero fields 使用opts和零字段更新数据
	UpdateByOptsAndZeroFields(ctx context.Context, where QueryOption, entity *T) error
	// update by opts and non zero fields 使用opts和非零字段更新数据
	UpdateByOptsAndNonZeroFields(ctx context.Context, where QueryOption, entity *T) error

	// delete 删除数据
	SoftDelete(ctx context.Context, id uint64) error
	// soft delete by ids 软删除批量数据
	SoftDeleteByIDs(ctx context.Context, ids []uint64) error
	// hard delete 硬删除数据
	HardDelete(ctx context.Context, id uint64) error
	// hard delete by ids 硬删除批量数据
	HardDeleteByIDs(ctx context.Context, ids []uint64) error

	// get by id 查询单个数据
	GetByID(ctx context.Context, id uint64, opts ...QueryOption) (*T, error)
	// get by ids 查询批量数据
	GetByIDs(ctx context.Context, ids []uint64, opts ...QueryOption) ([]*T, error)
	// get one 查询单个数据
	GetOne(ctx context.Context, opts ...QueryOption) (*T, error)
	// get many 查询多个数据
	GetMany(ctx context.Context, opts ...QueryOption) ([]*T, error)
	// get table 查询分页数据
	GetTable(ctx context.Context, page int, pageSize int, opts ...QueryOption) ([]*T, int64, error)
}

type BaseRepoImpl[T any] struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewBaseRepo[T any](db *gorm.DB, logger *zap.Logger) BaseRepo[T] {
	return &BaseRepoImpl[T]{db: db, logger: logger}
}

func (r *BaseRepoImpl[T]) GetByID(ctx context.Context, id uint64, opts ...QueryOption) (*T, error) {
	db := ApplyQueryOptions(r.db.WithContext(ctx), opts...)
	var entity T
	if err := db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepoImpl[T]) GetByIDs(ctx context.Context, ids []uint64, opts ...QueryOption) ([]*T, error) {
	db := ApplyQueryOptions(r.db.WithContext(ctx), opts...)
	var entities []*T
	if err := db.Find(&entities, ids).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseRepoImpl[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *BaseRepoImpl[T]) CreateBatch(ctx context.Context, entities []*T) error {
	return r.db.WithContext(ctx).CreateInBatches(entities, 1000).Error
}

func (r *BaseRepoImpl[T]) UpdateByMap(ctx context.Context, id uint64, entity map[string]any) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id = ?", id).Updates(entity).Error
}

func (r *BaseRepoImpl[T]) UpdateByZeroFields(ctx context.Context, id uint64, entity *T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id = ?", id).Select("*").Updates(entity).Error
}

func (r *BaseRepoImpl[T]) UpdateByNonZeroFields(ctx context.Context, id uint64, entity *T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id = ?", id).Updates(entity).Error
}

func (r *BaseRepoImpl[T]) UpdateBatchByIDsWithMap(ctx context.Context, ids []uint64, entity map[string]any) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id IN ?", ids).Updates(entity).Error
}

func (r *BaseRepoImpl[T]) UpdateBatchByIDsWithZeroFields(ctx context.Context, ids []uint64, entity *T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id IN ?", ids).Select("*").Updates(entity).Error
}

func (r *BaseRepoImpl[T]) UpdateBatchByIDsWithNonZeroFields(ctx context.Context, ids []uint64, entity *T) error {
	return r.db.WithContext(ctx).Model(new(T)).Where("id IN ?", ids).Updates(entity).Error
}

func (r *BaseRepoImpl[T]) UpdateByOptsAndZeroFields(ctx context.Context, where QueryOption, entity *T) error {
	db := ApplyQueryOptions(r.db.WithContext(ctx), where)
	return db.Model(new(T)).Select("*").Updates(entity).Error
}

func (r *BaseRepoImpl[T]) UpdateByOptsAndNonZeroFields(ctx context.Context, where QueryOption, entity *T) error {
	db := ApplyQueryOptions(r.db.WithContext(ctx), where)
	return db.Model(new(T)).Updates(entity).Error
}

func (r *BaseRepoImpl[T]) SoftDelete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(new(T), id).Error
}

func (r *BaseRepoImpl[T]) SoftDeleteByIDs(ctx context.Context, ids []uint64) error {
	return r.db.WithContext(ctx).Delete(new(T), ids).Error
}

func (r *BaseRepoImpl[T]) HardDelete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Unscoped().Delete(new(T), id).Error
}

func (r *BaseRepoImpl[T]) HardDeleteByIDs(ctx context.Context, ids []uint64) error {
	return r.db.WithContext(ctx).Unscoped().Delete(new(T), ids).Error
}

func (r *BaseRepoImpl[T]) GetOne(ctx context.Context, opts ...QueryOption) (*T, error) {
	var entity T
	db := ApplyQueryOptions(r.db.WithContext(ctx), opts...)
	if err := db.First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepoImpl[T]) GetMany(ctx context.Context, opts ...QueryOption) ([]*T, error) {
	var entities = make([]*T, 0)
	db := ApplyQueryOptions(r.db.WithContext(ctx), opts...)
	if err := db.Find(&entities).Error; err != nil {
		return entities, err
	}
	return entities, nil
}

func (r *BaseRepoImpl[T]) GetTable(ctx context.Context, page int, pageSize int, opts ...QueryOption) ([]*T, int64, error) {
	page, pageSize = utils.NormalizePageAndPageSize(page, pageSize)
	var total int64
	var entities = make([]*T, 0)
	db := ApplyQueryOptions(r.db.WithContext(ctx), opts...)
	var entity T
	if err := db.Model(&entity).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Find(&entities).Error; err != nil {
		return entities, 0, err
	}
	return entities, total, nil
}
