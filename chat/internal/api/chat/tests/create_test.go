package tests

import (
    "chat/internal/api/chat"
    "chat/internal/repo/mocks"
    "chat/internal/service"
    serv "chat/internal/service/chat"
    desc "chat/pkg/chat_v1"
    "context"
    "testing"

    "github.com/brianvoe/gofakeit"
    "github.com/gojuno/minimock/v3"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
    "reflect"
)

func TestCreate(t *testing.T) {
    type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService

    type args struct {
        ctx context.Context
        req *desc.CreateChatRequest
    }

    var (
        ctx      = context.Background()
        mc       = minimock.NewController(t)
        id       = gofakeit.Int32()
        name     = gofakeit.BeerName()
        length   = gofakeit.Number(1, 10)
        userIds  = make([]int32, length)
    )

    for i := 0; i < length; i++ {
        userIds[i] = gofakeit.Int32()
    }

    req := &desc.CreateChatRequest{
        Name:    name,
        UserIds: userIds,
    }

    tests := []struct {
        name            string
        args            args
        want            *emptypb.Empty
        err             error
        chatServiceMock chatServiceMockFunc
    }{
        {
            name: "success case",
            args: args{
                ctx: ctx,
                req: req,
            },
            want: &emptypb.Empty{},
            err:  nil,
            chatServiceMock: func(mc *minimock.Controller) service.ChatService {
                mock := mocks.NewRepoMock(mc)
                mock.CreateChatMock.Expect(name).Return(id, nil)
                for _, userId := range userIds {
                    mock.CreateChatUserMock.Expect(id, userId).Return(nil)
                }
                svc := serv.NewChatService(mock, nil)
                return svc
            },
        },
        {
            name: "missing name case",
            args: args{
                ctx: ctx,
                req: &desc.CreateChatRequest{
                    Name: "",
                },
            },
            want: &emptypb.Empty{},
            err:  status.Error(codes.InvalidArgument, "name is required"),
            chatServiceMock: func(mc *minimock.Controller) service.ChatService {
                return nil // Не нужен мок для этой проверки
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            chatService := tt.chatServiceMock(mc)
            impl := &chat.ChatImplementation{ChatService: chatService}
            got, err := impl.Create(tt.args.ctx, tt.args.req)

            // Проверка ошибок
            if (err != nil) != (tt.err != nil) || (err != nil && err.Error() != tt.err.Error()) {
                t.Errorf("Create() error = %v, wantErr = %v", err, tt.err)
                return
            }

            // Проверка результатов
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Create() got = %v, want = %v", got, tt.want)
            }
        })
    }
}
