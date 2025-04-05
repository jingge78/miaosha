package client

import (
	"context"
	"google.golang.org/grpc"
	"miaosha-jjl/common/proto/staff"
)

type StaffHandler func(ctx context.Context, client staff.StaffClient) (interface{}, error)

func StaffClient(ctx context.Context, handler StaffHandler) (interface{}, error) {
	dial, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer dial.Close()
	client := staff.NewStaffClient(dial)
	return handler(ctx, client)
}

func StaffEnd(ctx context.Context, i *staff.StaffRequest) (*staff.StaffResponse, error) {
	client, err := StaffClient(ctx, func(ctx context.Context, client staff.StaffClient) (interface{}, error) {
		createStatff, err := client.StaffSend(ctx, i)
		if err != nil {
			return nil, err
		}
		return createStatff, err
	})
	if err != nil {
		return nil, err
	}
	return client.(*staff.StaffResponse), err
}
