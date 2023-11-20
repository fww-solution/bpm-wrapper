package dtoticket

type RequestRedeemTicketToBPM struct {
	CaseID     int64    `json:"case_id"`
	IdNumbers  []string `json:"id_numbers"`
	CodeTicket string   `json:"code_ticket"`
}

type RequestUpdateTicket struct {
	CodeTicket         string `json:"code_ticket"`
	IsBoardingPass     bool   `json:"is_boarding_pass"`
	IsEligibleToFlight bool   `json:"is_eligible_to_flight"`
}