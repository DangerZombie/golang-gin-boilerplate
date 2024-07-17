package service_user

import (
	"github.com/DangerZombie/golang-gin-boilerplate/helper/message"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"
	"github.com/DangerZombie/golang-gin-boilerplate/model/request"
	"github.com/DangerZombie/golang-gin-boilerplate/model/response"
	"go.uber.org/zap"
)

// swagger:operation GET /user/profile User UserProfileRequest
// User profile
//
// ---
// security:
// - Bearer: []
// responses:
//   '200':
//     description: Success Response.
//     schema:
//       properties:
//         meta:
//           $ref: '#/definitions/MetaSingleSuccessResponse'
//         data:
//           properties:
//             record:
//               $ref: '#/definitions/UserProfileResponse'
//           type: object

func (s *userServiceImpl) UserProfile(req request.UserProfileRequest) (res response.UserProfileResponse, msg message.Message, errMsg map[string]string) {
	logger := s.logger.With(zap.String("UserService", "UserProfile"))
	errMsg = map[string]string{}

	if req.Id == "" {
		logger.Error("log", zap.String("error", "field cannot be empty"))
		errMsg["id"] = "field cannot be empty"
		return res, message.FailedMsg, errMsg
	}

	tx := s.baseRepo.GetBegin()
	defer func() {
		if msg != message.SuccessMsg {
			s.baseRepo.BeginRollback(tx)
		} else {
			s.baseRepo.BeginCommit(tx)
		}
	}()

	findUserByIdInput := parameter.FindUserByIdInput{
		Id: req.Id,
	}

	user, err := s.userRepo.FindUserById(tx, findUserByIdInput)
	if err != nil {
		logger.Error("log", zap.String("error", err.Error()))
		errMsg["id"] = "user id invalid"
		return res, message.FailedMsg, errMsg
	}

	res = response.UserProfileResponse{
		Id:       user.Id,
		Nickname: user.Nickname,
	}

	return res, message.SuccessMsg, nil
}
