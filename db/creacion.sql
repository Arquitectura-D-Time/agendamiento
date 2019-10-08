CREATE DATABASE IF NOT EXISTS agendamiento;

USE agendamiento;

CREATE TABLE IF NOT EXISTS Horario (
	IDtutoria BIGINT NOT NULL,
    IDtutor BIGINT,
    NombreMateria VARCHAR(30),
    Fecha VARCHAR(30),
    HoraInicio VARCHAR(30),
    HoraFinal VARCHAR(30),
    Cupos BIGINT,
    PRIMARY KEY(IDtutoria),
    INDEX(IDtutoria)
)ENGINE=InnoDB DEFAULT CHARACTER SET = utf8;

DESCRIBE Horario;

CREATE TABLE IF NOT EXISTS Agendadas (
	IDtutoria BIGINT NOT NULL,
    IDalumno BIGINT,
    PRIMARY KEY(IDtutoria),
    INDEX(IDtutoria)
)ENGINE=InnoDB DEFAULT CHARACTER SET = utf8;

DESCRIBE Agendadas;

ALTER USER 'Fernando'@'%' IDENTIFIED WITH mysql_native_password BY '2123';
FLUSH PRIVILEGES;
