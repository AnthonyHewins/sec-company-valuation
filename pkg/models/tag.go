package models

type Tag struct{
	Tag      string `csv:"tag"`
	Version  string `csv:"version"`
	Custom   bool   `csv:"custom"`
	Abstract bool   `csv:"abstract"`
	Datatype string `csv:"datatype"`
	Iord     string `csv:"iord"`
	Crdr     string `csv:"crdr"`
	Tlabel   string `csv:"tlabel"`
	Doc      string `csv:"doc"`
}
