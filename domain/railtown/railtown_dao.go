package railtown

import (
	"net/http"
	db "railtown/datasources/mysql"
	"railtown/logger"
	"railtown/utills/responses"
)

const (
	queryInsertIntoRailTown = "INSERT INTO railtown (`current`, `location`) VALUES(?, ?);"
)

func (rtr *RailTownRES) Save() *responses.Response {
	stmt, err := db.Client.Prepare(queryInsertIntoRailTown)
	if err != nil {
		logger.Error("error when trying to prepare save into railtown", err)
		return responses.NewBadRequestError("Internal Error", "try again...", http.StatusInternalServerError)
	}
	defer stmt.Close()

	result, saveErr := stmt.Exec(rtr.Current, rtr.Location)
	if saveErr != nil {
		logger.Error("error when trying to save railtown", saveErr)
		return responses.NewBadRequestError("railtown Inserted before", "Please make sure to enter new railtown", http.StatusBadRequest)
	}

	rtr.Id, err = result.LastInsertId()
	if err != nil {
		return responses.NewInternalServerError("NewInternalServerError", "Please try again later...")
	}

	return nil
}
