package requests

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
