package usecase

import (
	dtobooking "bpm-wrapper/internal/data/dto_booking"
	dtopayment "bpm-wrapper/internal/data/dto_payment"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/goccy/go-json"
)

// GenerateInvoice implements Usecase.
func (u *usecase) GenerateInvoice(body dtopayment.GenerateInvoiceRequest) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	id := watermill.NewUUID()

	// publish to message broker
	err = u.pub.Publish("generate_invoice_from_bpm", message.NewMessage(id, payload))
	if err != nil {
		return err
	}

	return nil
}

// StartProcessBooking implements Usecase.
func (u *usecase) StartProcessBooking(processName string, version string, body dtobooking.StartProcessBookingRequest) (string, error) {
	token, err := u.loginUser()
	if err != nil {
		return "", err
	}

	processId, err := u.adapter.FindProcess(&token, processName, version)
	if err != nil {
		return "", err
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	result, err := u.adapter.CreateProcessInstance(&token, processId, jsonBody)
	if err != nil {
		return "", err
	}
	fmt.Println("case id", result)
	return result, nil
}

// DoPayment implements Usecase.
func (u *usecase) DoPayment(body *dtopayment.DoPaymentRequest) error {
	token, err := u.loginUser()
	if err != nil {
		log.Println("Error Login", err)
		return err
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Error Marshal", err)
		return err
	}

	latestTask, err := u.repo.FindLatestTaskByCaseID(body.CaseID)
	if err != nil {
		log.Println("Error FindLatestTaskByCaseID", err)
		return err
	}

	task, err := u.adapter.FindTaskByName(&token, body.CaseID, latestTask.TaskName)
	if err != nil {
		log.Println("Error FindTaskByName", err)
		return err
	}

	err = u.adapter.ExecuteTask(&token, task.ID, jsonBody)
	if err != nil {
		log.Println("Error ExecuteTask", err)
		return err
	}

	return nil
}

// UpdatePayment implements Usecase.
func (u *usecase) UpdatePayment(body *dtopayment.RequestUpdatePayment) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	id := watermill.NewUUID()

	// publish to message broker
	err = u.pub.Publish("update_payment_from_bpm", message.NewMessage(id, jsonBody))
	if err != nil {
		return err
	}

	return nil
}
