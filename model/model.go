package model

type Horario struct {
	IDtutoria     int64  `json':"IDtutoria"`
	IDtutor       int64  `json':"IDtutor"`
	NombreMateria string `json':"NombreMateria"`
	Fecha         string `json':"Fecha"`
	HoraInicial   string `json':"HoraInicial"`
	HoraFinal     string `json':"HoraFinal"`
	Cupos         int64  `json':"Cupos"`
}

type Agendadas struct {
	IDtutoria int64 `json':"IDtutoria"`
	IDalumno  int64 `json':"IDalumno"`
}
