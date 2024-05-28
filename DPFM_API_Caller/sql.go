package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-station-train-time-table-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-station-train-time-table-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeadersByStation":
			func() {
				header = c.HeadersByStation(mtx, input, output, errs, log)
			}()
		case "HeadersByRailwayLine":
			func() {
				header = c.HeadersByRailwayLine(mtx, input, output, errs, log)
			}()
		case "HeadersByRailwayLineWeekdayType":
			func() {
				header = c.HeadersByRailwayLineWeekdayType(mtx, input, output, errs, log)
			}()
		case "HeadersByRailwayLineWeekdayTypeExpressType":
			func() {
				header = c.HeadersByRailwayLineWeekdayTypeExpressType(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:		header,
		Item:		item,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("header.Station = %d", input.Header.Station)

	where := fmt.Sprintf("%s\nAND header.RailwayLine = %d", where, input.Header.RailwayLine)

	where := fmt.Sprintf("%s\nAND header.TimeTableVersion = %d", where, input.Header.TimeTableVersion)

	where := fmt.Sprintf("%s\nAND header.WeekdayType = \"%s\"", where, input.Header.WeekdayType)

	where := fmt.Sprintf("%s\nAND header.OutboundOrInbound = \"%s\"", where, input.Header.OutboundOrInbound)

	where := fmt.Sprintf("%s\nAND header.TimeTableItem = %d", where, input.Header.TimeTableItem)

	where := fmt.Sprintf("%s\nAND header.ExpressType = \"%s\"", where, input.Header.ExpressType)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_train_time_table_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByRailwayLine(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.RailwayLine = %d", input.Header.RailwayLine)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_train_time_table_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.Station ASC, header.TimeTableVersion ASC, header.WeekdayType ASC, header.StationTrainTimeTableID ASC, header.ExpressType ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByStation(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Station = %d", input.Header.Station)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_train_time_table_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.RailwayLine ASC, header.TimeTableVersion ASC, header.WeekdayType ASC, header.StationTrainTimeTableID ASC, header.ExpressType ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByRailwayLineWeekdayType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Station = %d", input.Header.Station)

	where := fmt.Sprintf("%s\nAND header.RailwayLine = %d", where, input.Header.RailwayLine)

	where := fmt.Sprintf("%s\nAND header.WeekdayType = \"%s\"", where, input.Header.WeekdayType)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_train_time_table_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.TimeTableVersion ASC, header.StationTrainTimeTableID ASC, header.ExpressType ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByRailwayLineWeekdayTypeExpressType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.Station = %d", input.Header.Station)

	where := fmt.Sprintf("%s\nAND header.RailwayLine = %d", where, input.Header.RailwayLine)

	where := fmt.Sprintf("%s\nAND header.WeekdayType = \"%s\"", where, input.Header.WeekdayType)

	where := fmt.Sprintf("%s\nAND header.ExpressType = \"%s\"", where, input.Header.ExpressType)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_station_train_time_table_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.TimeTableVersion ASC, header.StationTrainTimeTableID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
