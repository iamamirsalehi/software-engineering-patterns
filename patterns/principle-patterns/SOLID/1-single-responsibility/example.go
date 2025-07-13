package userservice

type UserService struct{
  userRepo repo.UserRepository 
}

func New(userRepo repo.UserRepository) *UserService{
  return &{userRepo: userRepo}
}

func (us *UserService) Deactive(ctx context.Content, userID int) error {
  user, err := us.userRepo.FindByID(ctx, userID)
  if err != nil{
   return err 
  }
  
  err = user.Deactive()
  if err != nil{
    return err
  }
  
  return us.userRepo.Save(ctx, user)
}


func (us *UserService) Active(ctx context.Content, userID int) error {
  user, err := us.userRepo.FindByID(ctx, userID)
  if err != nil{
   return err 
  }
  
  err = user.Active()
  if err != nil{
    return err
  }
  
  return us.userRepo.Save(ctx, user)
}