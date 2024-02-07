package intelliflomodels

import (
	"encoding/json"
)

func (wp *WebhookPayload) UnmarshalJSON(data []byte) error {
	type Alias WebhookPayload
	aux := &struct {
		InstalledFor   map[string]interface{} `json:"installedFor,omitempty"`
		UninstalledFor map[string]interface{} `json:"uninstalledFor,omitempty"`
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
	return nil
}
