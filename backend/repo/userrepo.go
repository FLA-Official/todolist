package repo

import (
	"errors"
	"todolist/model"
)

type UserRepo interface {
	CreateUser(user *model.User) error
	GetUserByID(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int) error
	ListUsers() ([]model.User, error)
}

// ID        int       `json:"userid"`
// Username  string    `json:"username"`
// Fullname  string    `json:"fullname"`
// Gmail     string    `json:"gmail"`
// Password  string    `json:"password"`
// Role      string    `json:"role"` // developer, superviser, Admin, viewer
// Createdat time.Time `json:"createdat"`

type userRepo struct {
	userlist []*model.User
}

func NewUserRepo() UserRepo {
	repo := &userRepo{}
	return repo
}

func (u *userRepo) CreateUser(user *model.User) error {
	//adding ID
	user.ID = len(u.userlist) + 1
	//Adding time

	err := user.Validate()
	if err != nil {
		return err
	} else {
		for _, users := range u.userlist {
			if users.Gmail == user.Gmail {
				return errors.New("There is already a user under this gmail")
			} else if users.Username == user.Username {
				return errors.New("This username is not available")
			} else {
				u.userlist = append(u.userlist, user)
			}
		}
	}

	return nil
}

func (u *userRepo) GetUserByID(id int) (*model.User, error) {
	for _, user := range u.userlist {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (u *userRepo) GetUserByEmail(email string) (*model.User, error) {
	for _, user := range u.userlist {
		if user.Gmail == email {
			return user, nil
		}
	}

	return nil, errors.New("There is no user under this gmail")

}

func (u *userRepo) UpdateUser(user *model.User) error {
	for idx, users := range u.userlist {
		if users.ID == user.ID {
			users.CreatedAt = user.CreatedAt
			u.userlist[idx] = user
			return nil
		}
	}

	return errors.New("Task not found")
}

func (u *userRepo) DeleteUser(id int) error {
	var tempList []*model.User

	for _, user := range u.userlist {
		if user.ID != id {
			tempList = append(tempList, user)
		}
	}
	//to maintain ID order 1,2,3,4....
	for i := 0; i < len(tempList); i++ {
		tempList[i].ID = i + 1
	}

	u.userlist = tempList

	return nil
}

func (u *userRepo) ListUsers() ([]model.User, error) {
	var users []model.User

	if len(u.userlist) == 0 {
		return nil, nil
	} else {
		for _, user := range u.userlist {
			users = append(users, *user)
		}
		return users, nil
	}

}
