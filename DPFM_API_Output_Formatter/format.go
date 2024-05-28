package dpfm_api_output_formatter

import (
	"data-platform-api-station-train-time-table-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.Station,
			&pm.RailwayLine,
			&pm.TimeTableVersion,
			&pm.WeekdayType,
			&pm.OutboundOrInbound,
			&pm.TimeTableItem,
			&pm.ExpressType,
			&pm.DepartingTime,
			&pm.ArrivingTime,
			&pm.Description,
			&pm.OperationRemarks,
			&pm.IsSuspended,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			Station:						data.Station,
			RailwayLine:					data.RailwayLine,
			TimeTableVersion:				data.TimeTableVersion,
			WeekdayType:					data.WeekdayType,
			OutboundOrInbound:				data.OutboundOrInbound,
			TimeTableItem:					data.TimeTableItem,
			ExpressType:					data.ExpressType,
			DepartingTime:					data.DepartingTime,
			ArrivingTime:					data.ArrivingTime,
			Description:					data.Description,
			OperationRemarks:				data.OperationRemarks,
			IsSuspended:					data.IsSuspended,
			ValidityStartDate:				data.ValidityStartDate,
			ValidityEndDate:				data.ValidityEndDate,
			CreationDate:					data.CreationDate,
			CreationTime:					data.CreationTime,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
			CreateUser:						data.CreateUser,
			LastChangeUser:					data.LastChangeUser,
			IsReleased:						data.IsReleased,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
