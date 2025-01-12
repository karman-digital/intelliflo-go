package webhookmodels

import (
	"encoding/json"
	"errors"
	"strconv"
)

func (wp *WebhookPayload) UnmarshalJSON(data []byte) error {
	type Alias WebhookPayload
	aux := &struct {
		InstalledFor   map[string]any `json:"installedFor,omitempty"`
		UninstalledFor map[string]any `json:"uninstalledFor,omitempty"`
		TenantId       any            `json:"tenantId"`
		UserId         any            `json:"userId"`
		*Alias
	}{
		Alias: (*Alias)(wp),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.InstalledFor != nil {
		wp.InstalledFor = aux.InstalledFor
	} else if aux.UninstalledFor != nil {
		wp.InstalledFor = aux.UninstalledFor
	}
	switch v := aux.TenantId.(type) {
	case string:
		id, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("tenantId string could not be converted to int")
		}
		wp.TenantId = id
	case float64:
		wp.TenantId = int(v)
	default:
		wp.TenantId = 0
	}
	switch v := aux.UserId.(type) {
	case string:
		id, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("userId string could not be converted to int")
		}
		wp.UserId = id
	case float64:
		wp.UserId = int(v)
	default:
		wp.UserId = 0
	}
	return nil
}
