package chat

import (
	"context"
	"sync"
	model "chat/internal/model/chat"
	userProto "chat/userclient/pkg/user_v1"
)

func (s *chatService) Create(ctx context.Context, info *model.ChatInfo) error {
	chatId, err := s.storage.CreateChat(info.Name)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(info.UserIds))

	for _, userId := range info.UserIds {
		wg.Add(1)
		go func(userId int32) {
			defer wg.Done()
			_, err := s.clientUser.GetUser(ctx, &userProto.GetUserRequest{UserId: userId})
			if err != nil {
				errCh <- err
				return
			}
			err = s.storage.CreateChatUser(chatId, userId)
			if err != nil {
				errCh <- err
			}
		}(userId)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
