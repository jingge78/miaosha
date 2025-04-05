package staff

import (
	"context"
	"github.com/google/uuid"
	"miaosha-jjl/common/model"
	"miaosha-jjl/common/pkg"
	"miaosha-jjl/common/proto/staff"
	"strconv"
)

type StaffService struct {
	staff.UnimplementedStaffServer
}

func (s *StaffService) CreateStaff(ctx context.Context, in *staff.StaffRequest) (*staff.StaffResponse, error) {
	uid := uuid.New().String()
	m := model.EbSystemStoreStaff{
		Uid:          uint32(in.Uid),
		Avatar:       in.Avatar,
		StoreId:      uint32(in.StoreId),
		StaffName:    string(in.StaffName),
		Phone:        in.Phone,
		VerifyStatus: uint32(in.VerifyStatus),
		Status:       uint32(in.Status),
		AddTime:      uint32(in.AddTime),
	}
	err := m.StaffIntegralog()
	if err != nil {
		return nil, err
	}
	pay := pkg.NewPay()
	price := strconv.FormatFloat(float64(in.StaffName), 'f', 2, 6)
	s2 := pay.Pay(in.Phone, uid, price)
	return &staff.StaffResponse{Result: s2}, err

}
