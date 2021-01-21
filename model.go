package main

// AppInfo sturct that defines info of app
type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Rgb struct that defines red green blue value of each linr
type Rgb struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

// Line stuct defines the train line
type Line struct {
	Name        string `json:"name"`
	LineID      string `json:"lineId"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
	Rgb         Rgb    `json:"rgb"`
}

// Stop struct defines the info of each stop
type Stop struct {
	StopID                 int    `json:"stopID" bson:"STOP_ID"`
	DirectionID            string `json:"directionID" bson:"DIRECTION_ID"`
	StopName               string `json:"stopName" bson:"STOP_NAME"`
	StationName            string `json:"stationName" bson:"STATION_NAME"`
	StationDescriptiveName string `json:"stationDescriptiveName" bson:"STATION_DESCRIPTIVE_NAME"`
	Ada                    bool   `json:"ADA"`
	Red                    bool   `json:"RED"`
	Blue                   bool   `json:"BLUE"`
	Brown                  bool   `json:"BRN"`
	Purple                 bool   `json:"P"`
	PurpleExpress          bool   `json:"Pexp"`
	Yellow                 bool   `json:"Y"`
	Pink                   bool   `json:"Pnk"`
	Orange                 bool   `json:"O"`
	Location               string `json:"Location"`
}

// Ctatt Root element
type Ctatt struct {
	Ctatt Tatt
}

//Tatt ETA array
type Tatt struct {
	Eta []ETA `json:"ETA"`
}

// ETA single eta value
type ETA struct {
	ArrivalTime     string `json:"arrt"`
	DestinationName string `json:"destNm"`
	DestinationStop string `json:"destSt"`
	RunNumber       string `json:"rn"`
	Route           string `json:"rt"`
	IsApproaching   string `json:"isApp"`
	IsDelayed       string `json:"isDly"`
	StationName     string `json:"staNm"`
	StopDescription string `json:"stpDe"`
	PredictedTime   string `json:"prdt"`
}
