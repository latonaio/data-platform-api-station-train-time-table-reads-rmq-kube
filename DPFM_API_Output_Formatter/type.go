package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header	*[]Header	`json:"Header"`
}

type Header struct {
	Station						int		`json:"Station"`
	RailwayLine					int		`json:"RailwayLine"`
	TimeTableVersion			int		`json:"TimeTableVersion"`
	WeekdayType					string	`json:"WeekdayType"`
	OutboundOrInbound			string	`json:"OutboundOrInbound"`
	TimeTableItem				int		`json:"TimeTableItem"`
	ExpressType					string	`json:"ExpressType"`
	DepartingTime				string	`json:"DepartingTime"`
	ArrivingTime				*string	`json:"ArrivingTime"`
	Description					string	`json:"Description"`
	OperationRemarks			*string	`json:"OperationRemarks"`
	IsSuspended					*bool	`json:"IsSuspended"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	CreateUser					int		`json:"CreateUser"`
	LastChangeUser				int		`json:"LastChangeUser"`
	IsReleased					*bool	`json:"IsReleased"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
