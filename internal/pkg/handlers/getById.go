package handlers

import "context"

func (s *SomeHandlerService) GetByID(context.Context, interface{}) (interface{}, error) {
	return "data", nil
}
