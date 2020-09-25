package handlers

import "context"

func (s *SomeHandlerService) Hand(ctx context.Context, req interface{}) (res interface{}, err error) {
	return "Welcome", nil
}
