/**
 * @Author: Hongker
 * @Description:
 * @File:  upstream
 * @Version: 1.0.0
 * @Date: 2021/4/3 23:07
 */

package impl

import (
	"github.com/ebar-go/ego/component/mysql"
	"github.com/ebar-go/gateway/internal/domain/entity"
	"github.com/ebar-go/gateway/internal/domain/repository"
)

type upstreamRepoImpl struct {
	db mysql.Database
}

func (impl upstreamRepoImpl) FindByRouter(router string) (*entity.UpstreamEntity, error) {
	item := new(entity.UpstreamEntity)
	if err := impl.db.GetInstance().
		Model(item).Where("router = ?", router).
		First(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func newUpstreamRepo(db mysql.Database) repository.UpstreamRepo {
	return &upstreamRepoImpl{db: db}
}

