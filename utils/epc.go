package utils

import "encoding/xml"

type EpcCtx struct {
	XmlName           xml.Name `xml:"epml"`
	XmlnsEpml         string   `json:"@xmlns:epml" xml:"epml,attr"`
	XmlnsXsi          string   `json:"@xmlns:xsi" xml:"xsi,attr"`
	XsiSchemaLocation string   `json:"@xsi:schemaLocation" xml:"schemaLocation,attr"`
	Epc               struct {
		Name         string `json:"@name" xml:"name,attr"`
		ServicesName string `json:"@ServicesName" xml:"ServicesName,attr"`
		ProcCateID   string `json:"@procCateId" xml:"procCateId,attr"`
		ProcCateName string `json:"@procCateName" xml:"procCateName,attr"`
		Sort         string `json:"@sort" xml:"sort,attr"`
		Event        []struct {
			Topic      string `json:"@topic" xml:"topic,attr"`
			TopicID    string `json:"@topicId" xml:"topicId,attr"`
			TopicTitle string `json:"@topicTitle" xml:"topicTitle,attr"`
			ID         string `json:"@id" xml:"id,attr"`
			Name       string `json:"name" xml:"name"`
			Rule       struct {
				RuleID    string `json:"@ruleId" xml:"ruleId,attr"`
				RuleTitle string `json:"@ruleTitle" xml:"ruleTitle,attr"`
				Text      string `json:"#text" xml:",chardata"`
			} `json:"rule,omitempty" xml:"rule,omitempty"`
		} `json:"event" xml:"event"`
		Function []struct {
			ID       string `json:"@id" xml:"id,attr"`
			Name     string `json:"name" xml:"name"`
			Funccate struct {
				ServiceCateID    string `json:"@serviceCateId" xml:"serviceCateId"`
				ServiceCateTitle string `json:"@serviceCateTitle" xml:"serviceCateTitle"`
			} `json:"funccate" xml:"funccate"`
			Serviceoperation struct {
				ServiceName  string `json:"@serviceName" xml:"serviceName,attr"`
				ServiceID    string `json:"@serviceId" xml:"serviceId,attr"`
				ServiceTitle string `json:"@serviceTitle" xml:"serviceTitle,attr"`
			} `json:"serviceoperation" xml:"serviceoperation"`
			Form struct {
				FormID string `json:"@formId" xml:"formId,attr"`
			} `json:"form" xml:"form"`
			FieldPerms []struct {
				Name      string `json:"name" xml:"name"`
				Fieldname string `json:"fieldname" xml:"fieldname"`
				Datatype  string `json:"datatype" xml:"datatype"`
				Labelname string `json:"labelname" xml:"labelname"`
				Inputtype string `json:"inputtype" xml:"inputtype"`
				Permtype  string `json:"permtype" xml:"permtype"`
			} `json:"fieldPerms" xml:"fieldPerms"`
		} `json:"function" xml:"function"`
		Ou []struct {
			OuName string `json:"@ouName" xml:"ouName,attr"`
			ID     string `json:"@id" xml:"id,attr"`
			Roles  string `json:"roles" xml:"roles"`
		} `json:"ou" xml:"ou"`
		Iu []struct {
			IuName string `json:"@iuName" xml:"iuName,attr"`
			ID     string `json:"@id" xml:"id,attr"`
			Table  string `json:"table" xml:"table"`
		} `json:"iu" xml:"iu"`
		And struct {
			ID      string `json:"@id" xml:"id,attr"`
			AndName string `json:"@andName" xml:"andName,attr"`
		} `json:"and" xml:"and"`
		Arc []struct {
			ID   string `json:"@id" xml:"id,attr"`
			Flow struct {
				Source string `json:"@source" xml:"source,attr"`
				Target string `json:"@target" xml:"target,attr"`
			} `json:"flow" xml:"flow"`
		} `json:"arc" xml:"arc"`
	} `json:"epc" xml:"epc"`
}
