/**
 * @Author: Hongker
 * @Description:
 * @File:  inject
 * @Version: 1.0.0
 * @Date: 2021/4/3 22:27
 */

package impl

import "go.uber.org/dig"

func Inject(container *dig.Container) {
	_ = container.Provide(newDispatcherHandler)
}
