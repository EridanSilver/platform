package handlers

import "context"

func (s *SomeHandlerService) Hand2(ctx context.Context, req interface{}) (res interface{}, err error) {
	return "hand2test", nil
}
