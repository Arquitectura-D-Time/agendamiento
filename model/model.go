package model

type Horario struct {
	IDtutoria     int64  `json':"idtutoria"`
	IDtutor       int64  `json':"idtutor"`
	NombreMateria string `json':"nombremateria"`
	Fecha         string `json':"fecha"`
	HoraInicial   string `json':"horainicial"`
	HoraFinal     string `json':"horafinal"`
	Cupos         int64  `json':"cupos"`
}

type Agendadas struct {
	IDtutoria int64 `json':"idtutoria"`
	IDalumno  int64 `json':"idalumno"`
}
