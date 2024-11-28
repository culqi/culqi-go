package culqi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	utils "github.com/culqi/culqi-go/utils/validation"
)

const (
	planURL = baseURL + "/recurrent/plans"
)

// Create método para crear un plan
func CreatePlan(body []byte, encryptionData ...byte) (int, string, error) {
	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		return 0, "", fmt.Errorf("error al decodificar JSON: %w", err)
	}

	if data == nil {
		return 0, "", errors.New("el cuerpo JSON no se deserializó correctamente")
	}

	validator := utils.NewPlanValidation()
	if err := validator.Create(data); err != nil {
		return ErrorBadRequest, "", err
	}

	createURL := planURL + "/create"

	statusCode, res, err := Create(createURL, body, encryptionData...)
	if err != nil {
		return 0, "", fmt.Errorf("error al crear el plan: %w", err)
	}

	return statusCode, res, nil
}

// GetByID método para obtener un plan por id
func GetByIDPlan(id string, body []byte) (int, string, error) {
	validator := utils.NewPlanValidation()
	if err := validator.ValidateId(id); err != nil {
		return ErrorBadRequest, "", err
	}
	err := utils.ValidateStringStart(id, "pln")
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := GetById(planURL, id, body)
	return statusCode, res, err
}

// GetAll método para obtener la lista de los planes
func GetAllPlan(queryParams url.Values, body []byte) (int, string, error) {
	data := make(map[string]interface{})
	for key, values := range queryParams {
		if len(values) > 0 {
			data[key] = values[0]
		}
	}

	err := utils.PlanListValidation(data)
	if err != nil {
		return ErrorBadRequest, "", err
	}
	statusCode, res, err := GetAll(planURL, queryParams, body)
	return statusCode, res, err
}

// Update método para agregar o remplazar información a los valores de la metadata de un plan
func UpdatePlan(id string, body []byte, encryptionData ...byte) (int, string, error) {
	err := utils.ValidateStringStart(id, "pln")
	if err != nil {
		return ErrorBadRequest, "", err
	}

	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		return 0, "", fmt.Errorf("error al decodificar JSON: %w", err)
	}

	if data == nil {
		return 0, "", errors.New("el cuerpo JSON no se deserializó correctamente")
	}

	validator := utils.NewPlanValidation()
	if err := validator.Update(data, id); err != nil {
		return ErrorBadRequest, "", err
	}

	statusCode, res, err := Update(planURL, id, body, encryptionData...)
	return statusCode, res, err
}

// Delete método para eliminar un plan por id
func DeletePlan(id string, body []byte) (int, string, error) {
	validator := utils.NewPlanValidation()
	if err := validator.ValidateId(id); err != nil {
		return ErrorBadRequest, "", err
	}
	err := utils.ValidateStringStart(id, "pln")
	if err != nil {
		return ErrorBadRequest, "", err
	}

	statusCode, res, err := Delete(planURL, id, body)
	return statusCode, res, err
}
