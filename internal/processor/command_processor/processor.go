package command_processor

import (
	"github.com/pkg/errors"
	"github.com/zigal0/sport_bot/internal/domain"
)

type CommandProcessor struct {
	userRepo UserRepo
}

func NewCommandProcessor(userRepo UserRepo) *CommandProcessor {
	return &CommandProcessor{
		userRepo: userRepo,
	}
}

func (cp *CommandProcessor) GetUser(id int64) (domain.User, error) {
	user, err := cp.userRepo.GetUser(id)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "CommandProcessor.GetUser")
	}

	return user, nil
}

func (cp *CommandProcessor) AddUser(user domain.User) error {
	err := cp.userRepo.AddUser(user)
	if err != nil {
		return errors.Wrap(err, "CommandProcessor.AddUser")
	}

	return nil
}

func (cp *CommandProcessor) UpdateUser(user domain.User) error {
	err := cp.userRepo.UpdateUser(user)
	if err != nil {
		return errors.Wrap(err, "CommandProcessor.UpdateUser")
	}

	return nil
}

func (cp *CommandProcessor) DeleteUser(id int64) error {
	err := cp.userRepo.DeleteUser(id)
	if err != nil {
		return errors.Wrap(err, "CommandProcessor.DeleteUser")
	}

	return nil
}
