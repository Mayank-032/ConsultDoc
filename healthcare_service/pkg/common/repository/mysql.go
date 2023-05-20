package repository

import (
	"context"
	"database/sql"
	"errors"
	"healthcare-service/domain"
	"healthcare-service/domain/entity"
	"healthcare-service/domain/interfaces/repository"
	"log"
)

type CommonRepo struct {
	DB *sql.DB
}

func NewCommonRepo(db *sql.DB) repository.ICommonRepository {
	return CommonRepo {
		DB: db,
	}
}

func (cr CommonRepo) GetDoctorsList(ctx context.Context, address entity.Address) ([]entity.Doctor, error) {
	conn, err := cr.DB.Conn(ctx)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_db_connect\n\n", err.Error())
		return nil, errors.New("unable to db connect")
	}
	defer conn.Close()

	sqlQuery := "SELECT id, name, phone, ST_X(location), ST_Y(location), appointment_fees, FROM doctors WHERE ST_Distance_Sphere(location, POINT(?, ?)) <= 10000 and status = ?"
	args := []interface{}{address.Latitude, address.Longitude, domain.StatusActive}
	rows, err := cr.DB.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_execute_query\n\n", err)
		return nil, errors.New("unable to execute query")
	}

	var doctors []entity.Doctor
	for rows.Next() {
		var doctor entity.Doctor
		err = rows.Scan(&doctor.Id, &doctor.Name, &doctor.Phone, &doctor.Address.Latitude, &doctor.Address.Longitude, &doctor.Fees)
		if err != nil {
			log.Printf("Error: %v\n, unable_to_scan_data_from_database\n\n", err)
			return nil, errors.New("unable to scan data from database")
		}
		doctors = append(doctors, doctor)
	}
	return doctors, nil
}

func (cr CommonRepo) GetDoctorSlots(ctx context.Context, id int) ([]entity.Slot, error) {
	conn, err := cr.DB.Conn(ctx)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_db_connect\n\n", err.Error())
		return nil, errors.New("unable to db connect")
	}
	defer conn.Close()

	sqlQuery := "SELECT start_time, end_time WHERE doctor_id = ? and status = ?"
	args := []interface{}{id, domain.StatusActive}
	rows, err := cr.DB.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		log.Printf("Error: %v\n, unable_to_execute_query\n\n", err)
		return nil, errors.New("unable to execute query")
	}

	var slots []entity.Slot
	for rows.Next() {
		var slot entity.Slot
		err = rows.Scan(&slot.StartTime, &slot.EndTime)
		if err != nil {
			log.Printf("Error: %v\n, unable_to_scan_data_from_database\n\n", err)
			return nil, errors.New("unable to scan data from database")
		}
		slots = append(slots, slot)
	}
	return slots, nil
}