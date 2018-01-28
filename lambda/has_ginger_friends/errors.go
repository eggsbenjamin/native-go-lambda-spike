package main

type ErrNoGingerFriends struct {
	Name string
}

func (e ErrNoGingerFriends) Error() string {
	return "no ginger friends : " + e.Name
}
