/**
 * @Author: Hongker
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2020/6/17 23:40
 */

package request

type IdRequest struct {
	Id string `json:"id" binding:"required" comment:"编号"`
}
