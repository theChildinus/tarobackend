package utils

type EpcCtx struct {
	XmlnsEpml         string `json:"@xmlns:epml"`
	XmlnsXsi          string `json:"@xmlns:xsi"`
	XsiSchemaLocation string `json:"@xsi:schemaLocation"`
	Epc               struct {
		Name         string      `json:"@name"`
		ServicesName interface{} `json:"@ServicesName"`
		ProcCateID   string      `json:"@procCateId"`
		ProcCateName string      `json:"@procCateName"`
		Sort         interface{} `json:"@sort"`
		Event        []struct {
			Topic      string `json:"@topic"`
			TopicID    string `json:"@topicId"`
			TopicTitle string `json:"@topicTitle"`
			ID         string `json:"@id"`
			Name       string `json:"name"`
			Rule       struct {
				RuleID    string `json:"@ruleId"`
				RuleTitle string `json:"@ruleTitle"`
				Text      string `json:"#text"`
			} `json:"rule,omitempty"`
		} `json:"event"`
		Function []struct {
			ID       string `json:"@id"`
			Name     string `json:"name"`
			Funccate struct {
				ServiceCateID    string `json:"@serviceCateId"`
				ServiceCateTitle string `json:"@serviceCateTitle"`
			} `json:"funccate"`
			Serviceoperation struct {
				ServiceName  string `json:"@serviceName"`
				ServiceID    string `json:"@serviceId"`
				ServiceTitle string `json:"@serviceTitle"`
			} `json:"serviceoperation"`
			Form struct {
				FormID string `json:"@formId"`
			} `json:"form"`
			FieldPerms []struct {
				Name      string `json:"name"`
				Fieldname string `json:"fieldname"`
				Datatype  string `json:"datatype"`
				Labelname string `json:"labelname"`
				Inputtype string `json:"inputtype"`
				Permtype  string `json:"permtype"`
			} `json:"fieldPerms"`
		} `json:"function"`
		Ou []struct {
			OuName string `json:"@ouName"`
			ID     string `json:"@id"`
			Roles  string `json:"roles"`
		} `json:"ou"`
		Iu []struct {
			IuName string `json:"@iuName"`
			ID     string `json:"@id"`
			Table  string `json:"table"`
		} `json:"iu"`
		And struct {
			ID      string      `json:"@id"`
			AndName interface{} `json:"@andName"`
		} `json:"and"`
		Arc []struct {
			ID   string `json:"@id"`
			Flow struct {
				Source string `json:"@source"`
				Target string `json:"@target"`
			} `json:"flow"`
		} `json:"arc"`
	} `json:"epc"`
}
