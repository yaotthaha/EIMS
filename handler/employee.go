package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type EmployeeGETRequest struct {
	Auth
	UserID       uint64 `json:"user_id"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Age          uint8  `json:"age"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Position     string `json:"position"`
	Marry        string `json:"marry"`
	Education    string `json:"education"`
	JoinTime     uint64 `json:"join_time"`
	DepartmentID uint64 `json:"department_id"`
}

func (r *RunnginCtx) handlerEmployee(ginCtx *gin.Context) {
	switch ginCtx.Request.Method {
	case http.MethodGet: // Get Employee Info
		r.handlerEmployeeGET(ginCtx)
	default:
		ginCtx.JSON(http.StatusMethodNotAllowed, errUnknownMethod.ToGinJSON())
	}
}

func (r *RunnginCtx) handlerEmployeeGET(ginCtx *gin.Context) {
	var req EmployeeGETRequest
	err := ginCtx.ShouldBindJSON(&req)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, errUnknownMethod.ToGinJSON())
		return
	}

	whereSlice := make(map[string]any)
	parseRequest(&req, "json", func(key string, value any, fieldValue string) bool {
		switch fieldValue {
		case "user_id", "join_time":
			if v, ok := value.(uint64); !ok || v == 0 {
				return true
			} else {
				whereSlice[fieldValue] = v
			}
		case "name", "sex", "phone", "email", "position", "marry", "education":
			if v, ok := value.(string); !ok || v == "" {
				return true
			} else {
				whereSlice[fieldValue] = v
			}
		case "age":
			if v, ok := value.(uint8); !ok || v == 0 {
				return true
			} else {
				whereSlice[fieldValue] = v
			}
		case "department_id":
			if v, ok := value.(uint64); !ok || v == 0 {
				return true
			} else {
				whereSlice[fieldValue] = v
			}
		}
		return true
	})
	whereSqlSlice := make([]string, 0)
	whereQuerySlice := make([]any, 0)
	for k, v := range whereSlice {
		whereSqlSlice = append(whereSqlSlice, fmt.Sprintf("%s = ?", k))
		whereQuerySlice = append(whereQuerySlice, v)
	}

	var employeeInfo EmployeeInfo
	sqlCtx, sqlCancel := context.WithTimeout(r.Ctx, sqlExecTimeout)
	err = r.DB.WithContext(sqlCtx).Select("*").Table("employee").Joins("JOIN department on department.department_id = employee.department_id").Where(strings.Join(whereSqlSlice, " and "), whereQuerySlice...).Scan(&employeeInfo).Error
	sqlCancel()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, errServerUnknownError.ToGinJSON())
		return
	}
}
