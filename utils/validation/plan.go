package culqi

import (
	"regexp"
	"strconv"
)

type PlanValidation struct{}

func NewPlanValidation() *PlanValidation {
	return &PlanValidation{}
}

var regex = `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([-.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

func (t *PlanValidation) Create(data map[string]interface{}) error {
	requiredFields := []string{"short_name", "description", "amount", "currency", "interval_unit_time",
		"interval_count", "initial_cycles", "name"}
	err := additionalValidation(data, requiredFields)
	if err != nil {
		return NewCustomError(err.Error())
	} else {
		allowedDeviceValues := []float64{1, 2, 3, 4, 5, 6}
		intervalUnitTime, ok := data["interval_unit_time"].(float64)
		if !validateIsInteger(intervalUnitTime) || !validateFloat64InArray(intervalUnitTime, allowedDeviceValues) || !ok {
			return NewCustomError("El campo 'interval_unit_time' tiene un valor inválido o está vacío. Estos son los únicos valores permitidos: [ 1, 2, 3, 4, 5, 6]")
		}

		intervalCount, ok := data["interval_count"].(float64)
		if !ok || !validateIsInteger(intervalCount) || !validateInRange(intervalUnitTime, 0, 9999) {
			return NewCustomError("El campo 'interval_count' solo admite valores numéricos en el rango 0 a 9999.")
		}

		amount, ok := data["amount"].(float64)
		if !ok || !validateIsInteger(amount) {
			return NewCustomError("El campo 'amount' es inválido o está vacío, debe tener un valor numérico entero.")
		}

		currency, ok := data["currency"].(string)
		if !ok || currency == "" {
			return NewCustomError("El campo 'currency' es inválido o está vacío, el código de la moneda en tres letras (Formato ISO 4217). Culqi actualmente soporta las siguientes monedas: ['PEN','USD'].")

		}

		if err := validateEnumCurrency(data["currency"].(string)); err != nil {
			return NewCustomError(err.Error())
		}

		name, ok := data["name"].(string)
		countCaracterName := len(name)
		if !ok || !validateInRange(float64(countCaracterName), 5, 50) {
			return NewCustomError("El campo 'name' es inválido o está vacío. El valor debe tener un rango de 5 a 50 caracteres.")
		}

		description, ok := data["description"].(string)
		countCaracterDescription := len(description)
		if !ok || !validateInRange(float64(countCaracterDescription), 5, 250) {
			return NewCustomError("El campo 'description' es inválido o está vacío. El valor debe tener un rango de 5 a 250 caracteres.")
		}

		short_name, ok := data["short_name"].(string)
		countCaracterShort_name := len(short_name)
		if !ok || !validateInRange(float64(countCaracterShort_name), 5, 50) {
			return NewCustomError("El campo 'short_name' es inválido o está vacío. El valor debe tener un rango de 5 a 50 caracteres.")
		}

		requiredFieldsInitialCycles := []string{"has_initial_charge", "count", "amount", "interval_unit_time"}
		err := additionalValidation(data["initial_cycles"].(map[string]interface{}), requiredFieldsInitialCycles, "initial_cycles")
		if err != nil {
			return NewCustomError(err.Error())

		} else {
			if err := validateInitialCycles(data["initial_cycles"].(map[string]interface{}), data["currency"].(string)); err != nil {
				return NewCustomError(err.Error())
			}
		}

		if image, ok := data["image"]; ok {
			if valueImage, typeImage := image.(string); typeImage {
				countCaracterImage := len(valueImage)
				re := regexp.MustCompile(regex)

				if !validateInRange(float64(countCaracterImage), 5, 250) || !re.MatchString(valueImage) {
					return NewCustomError("El campo 'image' es inválido o está vacío. El valor debe ser una cadena y debe ser una URL válida.")
				}
			} else {
				return NewCustomError("El campo 'image' es inválido o está vacío. El valor debe ser una cadena y debe ser una URL válida.")
			}
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

func PlanListValidation(data map[string]interface{}) error {
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

	if min_amount, ok := data["min_amount"]; ok {
		if valuesMin_amount, typemMin_amount := min_amount.(string); typemMin_amount {
			valMinAmount, err := strconv.ParseFloat(valuesMin_amount, 64)
			if err != nil {
				return NewCustomError("El filtro 'min_amount' es invalido, debe tener un valor numérico entero.")
			}
			if !validateIsInteger(valMinAmount) {
				return NewCustomError("El filtro 'min_amount' es invalido, debe tener un valor numérico entero.")
			}

		} else {
			return NewCustomError("El filtro 'min_amount' es invalido, debe tener un valor numérico entero.")
		}
	}

	if max_amount, ok := data["max_amount"]; ok {
		if valuesMax_amount, typemMax_amount := max_amount.(string); typemMax_amount {
			valMaxAmount, err := strconv.ParseFloat(valuesMax_amount, 64)
			if err != nil {
				return NewCustomError("El filtro 'max_amount' es invalido, debe tener un valor numérico entero.")
			}
			if !validateIsInteger(valMaxAmount) {
				return NewCustomError("El filtro 'max_amount' es invalido, debe tener un valor numérico entero.")
			}

		} else {
			return NewCustomError("El filtro 'max_amount' es invalido, debe tener un valor numérico entero.")
		}
	}

	if limit, ok := data["limit"]; ok {
		if valuesLimit, typemLimit := limit.(string); typemLimit {
			valLimit, err := strconv.ParseFloat(valuesLimit, 64)
			if err != nil {
				return NewCustomError("El filtro 'limit' admite valores en el rango 1 a 100")
			}
			if !validateIsInteger(valLimit) || !validateInRange(valLimit, 1, 100) {
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

func (t *PlanValidation) Update(data map[string]interface{}, id string) error {
	countCaracterId := len(id)
	if id == "" || countCaracterId != 25 {
		return NewCustomError("El campo 'id' es inválido. La longitud debe ser de 25 caracteres.")
	}

	if short_name, ok := data["short_name"]; ok {
		if valueShort_name, typeShort_name := short_name.(string); typeShort_name {
			countCaracterShort_name := len(valueShort_name)
			if !ok || !validateInRange(float64(countCaracterShort_name), 5, 50) {
				return NewCustomError("El campo 'short_name' es inválido o está vacío. El valor debe tener un rango de 5 a 50 caracteres.")
			}
		} else {
			return NewCustomError("El campo 'short_name' es inválido o está vacío. El valor debe tener un rango de 5 a 50 caracteres.")
		}
	}

	if name, ok := data["name"]; ok {
		if valueName, typeName := name.(string); typeName {
			countCaracterName := len(valueName)
			if !ok || !validateInRange(float64(countCaracterName), 5, 50) {
				return NewCustomError("El campo 'name' es inválido o está vacío. El valor debe tener un rango de 5 a 50 caracteres.")
			}
		} else {
			return NewCustomError("El campo 'name' es inválido o está vacío. El valor debe tener un rango de 5 a 50 caracteres.")
		}
	}

	if description, ok := data["description"]; ok {
		if valueDescription, typeDescription := description.(string); typeDescription {
			countCaracterDescription := len(valueDescription)
			if !ok || !validateInRange(float64(countCaracterDescription), 5, 250) {
				return NewCustomError("El campo 'description' es inválido o está vacío. El valor debe tener un rango de 5 a 250 caracteres.")
			}
		} else {
			return NewCustomError("El campo 'description' es inválido o está vacío. El valor debe tener un rango de 5 a 250 caracteres.")
		}
	}

	if image, ok := data["image"]; ok {
		if valueImage, typeImage := image.(string); typeImage {
			countCaracterImage := len(valueImage)
			re := regexp.MustCompile(regex)

			if !validateInRange(float64(countCaracterImage), 5, 250) || !re.MatchString(valueImage) {
				return NewCustomError("El campo 'image' es inválido o está vacío. El valor debe ser una cadena y debe ser una URL válida.")
			}
		} else {
			return NewCustomError("El campo 'image' es inválido o está vacío. El valor debe ser una cadena y debe ser una URL válida.")
		}
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

	if status, ok := data["status"]; ok {
		if valueStatus, typeStatus := status.(float64); typeStatus {
			allowedStatusPlanValues := []float64{1, 2}
			if !validateFloat64InArray(valueStatus, allowedStatusPlanValues) || !ok {
				return NewCustomError("El campo 'status' tiene un valor inválido o está vacío. Estos son los únicos valores permitidos: [ 1, 2 ]")
			}
		} else {
			return NewCustomError("El campo 'status' tiene un valor inválido o está vacío. Estos son los únicos valores permitidos: [ 1, 2 ]")
		}
	}

	return nil
}

func (t *PlanValidation) ValidateId(id string) error {
	countCaracterId := len(id)
	if id == "" || countCaracterId != 25 {
		return NewCustomError("El campo 'id' es inválido. La longitud debe ser de 25 caracteres.")
	}
	return nil
}
