package user

func ToEntity(userDto UserDto) (userEntity User) {
	userEntity = User{
		Name:     userDto.Name,
		Username: userDto.Username,
		Password: userDto.Password,
	}
	return
}

func ToEntities(userDto []UserDto) (users []User) {
	for _, item := range userDto {
		users = append(users, User{
			Name:     item.Name,
			Username: item.Username,
			Password: item.Password,
		})
	}
	return
}

func ToDto(userEntity User) (userDto UserDto) {
	userDto = UserDto{
		ID:       userEntity.ID,
		Name:     userEntity.Name,
		Username: userEntity.Username,
		Password: userEntity.Password,
	}
	return
}

func ToDtos(userEntity []User) (userDtos []UserDto) {
	for _, item := range userEntity {
		userDtos = append(userDtos, UserDto{
			ID:       item.ID,
			Name:     item.Name,
			Username: item.Username,
			Password: item.Password,
		})
	}
	return
}
