CREATE DATABASE AlertsMachine;
USE AlertsMachine;
IF NOT EXISTS (SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'receivedAlerts')
BEGIN
    CREATE TABLE receivedAlerts (
        id INT PRIMARY KEY IDENTITY(1,1),
        alertId NVARCHAR(255),
        occurrences NVARCHAR(255),
        machineLinearTime BIGINT,
        bus NVARCHAR(255),
        time NVARCHAR(255),
        color NVARCHAR(255),
        severity NVARCHAR(255),
        acknowledgementStatus NVARCHAR(255),
        ignored BIT,
        invisible BIT,
        rel NVARCHAR(255),
        uri NVARCHAR(255),
        durationSeconds NVARCHAR(255),
        engineHoursReading FLOAT,
        suspectParameterName NVARCHAR(255),
        failureModeIndicator NVARCHAR(255),
        definitionBus NVARCHAR(255),
        sourceAddress NVARCHAR(255),
        threeLetterAcronym NVARCHAR(255),
        definitionId NVARCHAR(255),
        description NVARCHAR(255),
        definitionRel NVARCHAR(255),
        definitionUri NVARCHAR(255),
        lat FLOAT,
        lon FLOAT
    );
END
-- Inserir os dados na tabela principal
INSERT INTO receivedAlerts (
    alertId,
    occurrences, 
    machineLinearTime, 
    bus, 
    time,
    color, 
    severity, 
    acknowledgementStatus, 
    ignored, 
    invisible,
    rel, 
    uri, 
    durationSeconds, 
    engineHoursReading,
    suspectParameterName, 
    failureModeIndicator, 
    definitionBus, 
    sourceAddress, 
    threeLetterAcronym, 
    definitionId, 
    description, 
    definitionRel, 
    definitionUri, 
    lat, 
    lon
) VALUES (
    '123456789',  -- O valor '1234567' corresponde a "Id" no JSON
    '1',  -- O valor '1' corresponde a "occurrences" no JSON
    '36951157',  -- O valor '36951157' corresponde a "machineLinearTime" no JSON
    '0',  -- O valor '0' corresponde a "bus" no JSON
    '2020-05-13T19:15:11.000Z',  -- O valor '2020-05-13T19:15:11.000Z' corresponde a "time" no JSON
    'BLUE',  -- O valor 'BLUE' corresponde a "color" no JSON
    'INFO',  -- O valor 'INFO' corresponde a "severity" no JSON
    'N',  -- O valor 'N' corresponde a "acknowledgementStatus" no JSON
    0,  -- O valor 0 corresponde a "ignored" no JSON
    0,  -- O valor 0 corresponde a "invisible" no JSON
    'machine',  -- O valor 'machine' corresponde a "rel" no JSON
    'https://sandboxapi.deere.com/platform/machines/123456',  -- O valor corresponde a "uri" no JSON
    '3',  -- O valor '3' corresponde a "durationSeconds" no JSON
    '668.25',  -- O valor '668.25' corresponde a "engineHoursReading" no JSON
    '524019',  -- O valor '524019' corresponde a "suspectParameterName" no JSON
    '31',  -- O valor '31' corresponde a "failureModeIndicator" no JSON
    '0',  -- O valor '0' corresponde a "definitionBus" no JSON
    '140',  -- O valor '140' corresponde a "sourceAddress" no JSON
    'AIC',  -- O valor 'AIC' corresponde a "threeLetterAcronym" no JSON
    '1234567',  -- O valor '1234567' corresponde a "definitionId" no JSON
    'Other AIC 524019.31 Reverser lever left in incorrect position. - Return to park to attempt recovery.',  -- O valor corresponde a "description" no JSON
    'self',  -- O valor 'self' corresponde a "definitionRel" no JSON
    'https://sandboxapi.deere.com/platform/api/alertTypes/diagnosticTroubleCodeAlert/definitions/1328105',  -- O valor corresponde a "definitionUri" no JSON
    '41.123456',  -- O valor '41.123456' corresponde a "lat" no JSON
    '-90.234567'  -- O valor '-90.234567' corresponde a "lon" no JSON
);
-- Tabela para Notificações
CREATE TABLE Notificacoes (
    id INT IDENTITY(1,1) PRIMARY KEY,
    [key] NVARCHAR(255),
    [value] INT
);
-- Tabela para Encaminhamento Ativo
CREATE TABLE EncaminhamentoAtivo (
    idEncaminhamento INT IDENTITY(1,1) PRIMARY KEY,
    idEvento INT,
    idUsuarioEncaminhador INT,
    idUsuario INT,
    motivo NVARCHAR(255),
    idEmpresa INT,
    encaminhamentoAtivo BIT,
    dataInclusao DATETIME,
    dataAlteracao DATETIME,
    usuarioInc NVARCHAR(255),
    usuarioAlt NVARCHAR(255),
    origemRetorno INT
);
-- Tabela para Veículo
CREATE TABLE Veiculo (
    idVeiculo NVARCHAR(255) PRIMARY KEY,
    idCliente INT,
    idEmpresa INT,
    placa NVARCHAR(255),
    cor NVARCHAR(255),
    chassi NVARCHAR(255),
    modelo NVARCHAR(255)
);
-- Tabela para Evento
CREATE TABLE Evento (
	idEvento NVARCHAR(255) PRIMARY KEY,
    idCliente INT,
    idEmpresa INT,
    idGrupoEvento NVARCHAR(255),
    idTipoEvento NVARCHAR(255),
    formaContato NVARCHAR(255),
    midia NVARCHAR(255),
    temperatura NVARCHAR(255),
    marca NVARCHAR(255),
    modelo NVARCHAR(255),
    idAgente INT,
    observacao NVARCHAR(255),
    dataProximaAcao DATETIME
);
-- Tabela para Interface de Negociação
CREATE TABLE InterfaceNegociacao (
    idInterfaceNegociacao INT IDENTITY(1,1) PRIMARY KEY,
    idNotificacoes INT,
    idEncaminhamentoAtivo INT,
    idVeiculo NVARCHAR(255),
    idEvento NVARCHAR(255),
    FOREIGN KEY (idNotificacoes) REFERENCES Notificacoes(id),
    FOREIGN KEY (idEncaminhamentoAtivo) REFERENCES EncaminhamentoAtivo(idEncaminhamento),
    FOREIGN KEY (idVeiculo) REFERENCES Veiculo(idVeiculo),
    FOREIGN KEY (idEvento) REFERENCES Evento(idEvento)
);