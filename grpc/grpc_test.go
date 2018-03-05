package grpc

import (
	"context"
	"testing"

	"encoding/json"
	"log"

	"github.com/Sharykhin/gl-mail-grpc"
	"github.com/Sharykhin/gl-mail-manager/entity"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type mockClient struct {
	mock.Mock
}

func (m *mockClient) CreateFailMail(ctx context.Context, in *api.FailMailRequest, opts ...grpc.CallOption) (*api.FailMailResponse, error) {
	ret := m.Called(ctx, in)
	res, err := ret.Get(0), ret.Get(1)
	if err != nil {
		return nil, err.(error)
	}

	return res.(*api.FailMailResponse), nil
}

func (m *mockClient) GetFailMails(ctx context.Context, in *api.FailMailFilter, opts ...grpc.CallOption) (api.FailMail_GetFailMailsClient, error) {
	return nil, nil
}

func (m *mockClient) CountFailMails(ctx context.Context, in *api.Empty, opts ...grpc.CallOption) (*api.CountResponse, error) {
	return nil, nil
}

func TestCreateFailMail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mm := entity.MailMessage{
			Action: "test",
			Payload: map[string]interface{}{
				"to": "test@test.com",
			},
		}
		reason := "test reason"

		b, err := json.Marshal(mm.Payload)
		if err != nil {
			log.Fatalf("Could not marshal: %v", err)
		}

		fmr := api.FailMailRequest{
			Action:  mm.Action,
			Payload: b,
			Reason:  reason,
		}

		res := api.FailMailResponse{
			ID:      1,
			Action:  "test",
			Payload: b,
			Reason:  "test reason",
		}

		m := new(mockClient)
		m.On("CreateFailMail", context.Background(), &fmr).Return(&res, nil).Once()
		client = m

		row, err := CreateFailedMail(mm, reason)
		if err != nil {
			t.Errorf("expected error nil but got: %v", err)
		}
		m.AssertExpectations(t)

		assert.Equal(t, mm.Action, row.Action)
		assert.Equal(t, b, row.Payload)
		assert.Equal(t, reason, row.Reason)
	})

	t.Run("error", func(t *testing.T) {
		mm := entity.MailMessage{
			Action: "test",
			Payload: map[string]interface{}{
				"to": "test@test.com",
			},
		}
		reason := "test reason"

		b, err := json.Marshal(mm.Payload)
		if err != nil {
			log.Fatalf("Could not marshal: %v", err)
		}

		fmr := api.FailMailRequest{
			Action:  mm.Action,
			Payload: b,
			Reason:  reason,
		}

		errEx := errors.New("test error")

		m := new(mockClient)
		m.On("CreateFailMail", context.Background(), &fmr).Return(nil, errEx).Once()
		client = m

		row, err := CreateFailedMail(mm, reason)
		m.AssertExpectations(t)

		assert.Nil(t, row)
		assert.NotNil(t, err)
		assert.Equal(t, errEx.Error(), err.Error())
	})
}
