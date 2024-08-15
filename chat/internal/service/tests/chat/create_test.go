package chat

import (
	model "chat/internal/model/chat"
	"chat/internal/repo"
	repoMocks "chat/internal/repo/mocks"
	chatServ "chat/internal/service/chat"
	userProto "chat/userclient/pkg/user_v1"
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type client struct {}

func (c *client) CreateUser(ctx context.Context, req *userProto.CreateUserRequest, other ...grpc.CallOption) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (c *client) GetUser(ctx context.Context, req *userProto.GetUserRequest, other ...grpc.CallOption) (*userProto.GetUserResponse, error) {
	return &userProto.GetUserResponse{}, nil
}

func TestCreate(t *testing.T) {
	type chatRepoMockFunc func(mc *minimock.Controller) repo.Repo

	type args struct {
		ctx context.Context
		req *model.ChatInfo
	}

	var (
		ctx    = context.Background()
		mc     = minimock.NewController(t)
		id     = gofakeit.Int32()
		name   = gofakeit.Name()
		length = gofakeit.Number(1, 10)
		userIds = make([]int32, length)
		client = &client{}
	)
	for i := 0; i < length; i++ {
		userIds[i] = gofakeit.Int32()
	}

	req := &model.ChatInfo{
		Name:    name,
		UserIds: userIds,
	}

	tests := []struct {
		name        string
		args        args
		err         error
		chatRepoMock chatRepoMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			chatRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := repoMocks.NewRepoMock(mc)
				mock.CreateChatMock.Expect(req.Name).Return(id, nil)
				for _, userId := range req.UserIds {
					mock.CreateChatUserMock.Expect(id, userId).Return(nil)
				}
				return mock
			},
		},
		{
			name: "repository error",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: fmt.Errorf("repo error"),
			chatRepoMock: func(mc *minimock.Controller) repo.Repo {
				mock := repoMocks.NewRepoMock(mc)
				mock.CreateChatMock.Expect(req.Name).Return(0, fmt.Errorf("repo error"))
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			chatRepoMock := tt.chatRepoMock(mc)
			service := chatServ.NewChatService(chatRepoMock, client)
			err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
