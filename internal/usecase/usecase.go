package usecase

import (
	"bpm-wrapper/internal/adapter"
	"bpm-wrapper/internal/config"
	"bpm-wrapper/internal/data/dto"
	dtobooking "bpm-wrapper/internal/data/dto_booking"
	dtonotification "bpm-wrapper/internal/data/dto_notification"
	dtopassenger "bpm-wrapper/internal/data/dto_passenger"
	dtopayment "bpm-wrapper/internal/data/dto_payment"
	dtoticket "bpm-wrapper/internal/data/dto_ticket"
	"bpm-wrapper/internal/repository"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

type usecase struct {
	adapter adapter.Adapter
	cfg     *config.BonitaConfig
	redis   *redis.Client
	pub     message.Publisher
	repo    repository.Repository
}
type Usecase interface {
	// StartProcess
	StartProcess(processName string, version string, body interface{}) (string, error)
	// StopProcess
	StopProcess(token string, processId string) error
	// Update Human Process
	UpdateHumanProcess(processId string, variables []byte) error
	// Assign Human Task
	AssignHumanTask(taskID int64, caseID int64, userID string) error
	// // ExecuteHumanTask
	ExecuteHumanTask(taskID string, caseID int64, variables []byte) error
	// GetTaskID
	GetTaskID(taskName string, caseID int64) (string, error)
	// SaveWorkflow
	SaveWorkflow(body *dto.SaveWorkflowRequest) error

	// Passenger
	UpdatePassenger(body dtopassenger.Passenger) error
	StartProcessPassenger(processName string, version string, body dto.StartProcessPassengerRequest) (string, error)

	// Booking
	GenerateInvoice(body dtopayment.GenerateInvoiceRequest) error
	StartProcessBooking(processName string, version string, body dtobooking.StartProcessBookingRequest) (string, error)
	DoPayment(body *dtopayment.DoPaymentRequest) error
	UpdatePayment(body *dtopayment.RequestUpdatePayment) error
	UpdateBooking(body *dtobooking.RequestUpdateBooking) error

	// Ticket
	UpdateTicket(body *dtoticket.RequestUpdateTicket) error
	RedeemTicket(body *dtoticket.RequestRedeemTicketToBPM) error

	// Notification
	SendEmailNotification(body *dtonotification.SendEmailRequest) error
	SendNotification(body *dtonotification.Request) error
}

func New(adapter *adapter.Adapter, cfg *config.BonitaConfig, redis *redis.Client, pub message.Publisher, repo repository.Repository) Usecase {
	return &usecase{
		adapter: *adapter,
		cfg:     cfg,
		redis:   redis,
		pub:     pub,
		repo:    repo,
	}
}
