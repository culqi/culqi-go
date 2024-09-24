package culqi

import (
	"fmt"
	"strconv"
)

// CulqiValidation contains methods for validating card data
type SubscriptionValidation struct{}

// NewCulqiValidation creates a new instance of CulqiValidation
func NewSubscriptionValidation() *SubscriptionValidation {
	return &SubscriptionValidation{}
}

func (t *SubscriptionValidation) Create(data map[string]interface{}) error {
	requiredFields := []string{"card_id", "plan_id", "tyc"}
	err := additionalValidation(data, requiredFields)
	if err != nil {
		return NewCustomError(err.Error())
	} else {
		card_id, ok := data["card_id"].(string)
		countCaracter_card_id := len(card_id)
		if !ok || countCaracter_card_id != 25 {
			return NewCustomError("El campo 'card_id' es inválido. La longitud debe ser de 25.")
		}
		if err := ValidateStringStart(data["card_id"].(string), "crd"); err != nil {
			return err
		}

		plan_id, ok := data["plan_id"].(string)
		countCaracter_plan_id := len(plan_id)
		if !ok || countCaracter_plan_id != 25 {
			return NewCustomError("El campo 'plan_id' es inválido. La longitud debe ser de 25.")
		}
		if err := ValidateStringStart(data["plan_id"].(string), "pln"); err != nil {
			return err
		}

		tyc, ok := data["tyc"].(bool)
		if !ok {
			return NewCustomError("El campo 'tyc' es inválido o está vacío. El valor debe ser un booleano.")
		}

		fmt.Print(tyc)

		if metadata, ok := data["metadata"]; ok {
			if valueMetadata, typeMetada := metadata.(map[string]interface{}); typeMetada {
				err := validateMetadataSchema(valueMetadata)
				if err != nil {
					return NewCustomError("Enviaste el campo metadata con un formato incorrecto.")
				}
			} else {
				return NewCustomError("Enviaste el campo metadata con un formato incorrecto.")
			}
		}
	}
	return nil
}

func SubscriptionListValidation(data map[string]interface{}) error {

	if _, existsFrom := data["creation_date_from"]; existsFrom {
		if _, existsTo := data["creation_date_to"]; existsTo {
			date_from, _ := strconv.ParseInt(data["creation_date_from"].(string), 10, 64)
			date_to, _ := strconv.ParseInt(data["creation_date_to"].(string), 10, 64)
			err := ValidateDateFilter(date_from, date_to)
			if err != nil {
				return err
			}
		}
	}

	if plan_id, ok := data["plan_id"]; ok {
		if val_plan_id, oktype := plan_id.(string); oktype {
			countCaracter_plan_id := len(val_plan_id)
			if !ok || countCaracter_plan_id != 25 {
				return NewCustomError("El campo 'plan_id' es inválido. La longitud debe ser de 25.")
			}
			if err := ValidateStringStart(data["plan_id"].(string), "pln"); err != nil {
				return NewCustomError("El campo 'plan_id' es inválido. La longitud debe ser de 25.")
			}

		} else {
			return NewCustomError("El campo 'plan_id' es inválido. La longitud debe ser de 25.")
		}
	}

	if limit, ok := data["limit"]; ok {
		if valuesLimit, typemLimit := limit.(string); typemLimit {
			valLimit, err := strconv.ParseFloat(valuesLimit, 64)
			if err != nil {
				return NewCustomError("El filtro 'limit' admite valores en el rango 1 a 100")
			}
			typeLimit := validateIsInteger(valLimit)
			if !typeLimit || !validateInRange(valLimit, 1, 100) {
				return NewCustomError("El filtro 'limit' admite valores en el rango 1 a 100")
			}

		} else {
			return NewCustomError("El filtro 'limit' admite valores en el rango 1 a 100")
		}
	}

	if before, ok := data["before"]; ok {
		if valuesBefore, typemBefore := before.(string); typemBefore {
			countCaracterBefore := len(valuesBefore)
			if countCaracterBefore != 25 {
				return NewCustomError("El campo 'before' es inválido. La longitud debe ser de 25 caracteres")
			}

		} else {
			return NewCustomError("El campo 'before' es inválido. La longitud debe ser de 25 caracteres")
		}
	}

	if after, ok := data["after"]; ok {
		if valueAfter, typemAfter := after.(string); typemAfter {
			countCaracterAfter := len(valueAfter)
			if countCaracterAfter != 25 {
				return NewCustomError("El campo 'after' es inválido. La longitud debe ser de 25 caracteres")
			}

		} else {
			return NewCustomError("El campo 'after' es inválido. La longitud debe ser de 25 caracteres")
		}
	}

	if creation_date_from, ok := data["creation_date_from"]; ok {
		if value_creation_date_from, type_creation_date_from := creation_date_from.(string); type_creation_date_from {
			countCaracter_creation_date_from := len(value_creation_date_from)
			if !(countCaracter_creation_date_from == 10 || countCaracter_creation_date_from == 13) {
				return NewCustomError("El campo 'creation_date_from' debe tener una longitud de 10 o 13 caracteres.")
			}

		} else {
			return NewCustomError("El campo 'creation_date_from' debe tener una longitud de 10 o 13 caracteres.")
		}
	}

	if creation_date_to, ok := data["creation_date_to"]; ok {
		if value_creation_date_to, type_creation_date_to := creation_date_to.(string); type_creation_date_to {
			countCaracter_creation_date_to := len(value_creation_date_to)
			if !(countCaracter_creation_date_to == 10 || countCaracter_creation_date_to == 13) {
				return NewCustomError("El campo 'creation_date_to' debe tener una longitud de 10 o 13 caracteres.")
			}

		} else {
			return NewCustomError("El campo 'creation_date_to' debe tener una longitud de 10 o 13 caracteres.")
		}
	}

	if status, ok := data["status"]; ok {
		if valuesStatus, typeStatus := status.(string); typeStatus {
			allowedStatusPlanValues := []float64{1, 2}
			valStatus, err := strconv.ParseFloat(valuesStatus, 64)
			if err != nil {
				return NewCustomError("El filtro 'status' tiene un valor inválido o está vacío. Estos son los únicos valores permitidos: 1, 2.")
			}
			if !validateFloat64InArray(valStatus, allowedStatusPlanValues) || !ok {
				return NewCustomError("El filtro 'status' tiene un valor inválido o está vacío. Estos son los únicos valores permitidos: 1, 2.")
			}
		} else {
			return NewCustomError("El filtro 'status' tiene un valor inválido o está vacío. Estos son los únicos valores permitidos: 1, 2.")
		}
	}

	return nil
}

func (t *SubscriptionValidation) Update(data map[string]interface{}, id string) error {
	countCaracterId := len(id)
	if id == "" || countCaracterId != 25 {
		return NewCustomError("El campo 'id' es inválido. La longitud debe ser de 25 caracteres.")
	}

	requiredFields := []string{"card_id"}
	err := additionalValidation(data, requiredFields)
	if err != nil {
		return NewCustomError(err.Error())
	} else {
		card_id, ok := data["card_id"].(string)
		countCaracter_card_id := len(card_id)
		if !ok || countCaracter_card_id != 25 {
			return NewCustomError("El campo 'card_id' es inválido. La longitud debe ser de 25.")
		}
		if err := ValidateStringStart(data["card_id"].(string), "crd"); err != nil {
			return err
		}

		if metadata, ok := data["metadata"]; ok {
			if valueMetadata, typeMetada := metadata.(map[string]interface{}); typeMetada {
				err := validateMetadataSchema(valueMetadata)
				if err != nil {
					return NewCustomError("Enviaste el campo metadata con un formato incorrecto.")
				}
			} else {
				return NewCustomError("Enviaste el campo metadata con un formato incorrecto.")
			}
		}
	}
	return nil
}

func (t *SubscriptionValidation) ValidateId(id string) error {
	countCaracterId := len(id)
	if id == "" || countCaracterId != 25 {
		return NewCustomError("El campo 'id' es inválido. La longitud debe ser de 25 caracteres.")
	}
	return nil
}
