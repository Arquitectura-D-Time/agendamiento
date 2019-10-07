DROP DATABASE IF exists agendamiento;
CREATE DATABASE agendamiento CHARACTER SET utf8;

USE agendamiento;

CREATE TABLE Horario (
	IDtutoria BIGINT NOT NULL,
    IDtutor BIGINT,
    NombreMateria VARCHAR(30),
    Fecha VARCHAR(30),
    HoraInicio VARCHAR(30),
    HoraFinal VARCHAR(30),
    Cupos BIGINT
);

CREATE TABLE Agendadas (
	IDtutoria BIGINT NOT NULL,
    IDalumno BIGINT
);