package project_errors

import "errors"

var (
	// Bank.
	ErrInsertBank   = errors.New("err_insert_bank")
	ErrUpdateBank   = errors.New("err_update_bank")
	ErrValidateBank = errors.New("err_bank_validate")
	// CRUD.
	ErrToFind = errors.New("err_to_find")
	// Marshal && Unmarshal.
	ErrToUnmarshal = errors.New("err_to_unmarshal")
	// Snowflake.
	ErrToCreateSnowflake   = errors.New("err_to_create_snowflake")
	ErrToValidateSnowflake = errors.New("err_to_validate_snow_flake")
)
