CREATE TABLE IF NOT EXISTS
    nfe (
        ID int,
        CreatedAt timestamp,
        UpdatedAt timestamp,
        DeletedAt timestamp,
        RawXml text,
	    NFeId varchar(255),
	    CNPJ   varchar(255),
	    VNF    varchar(255)
    );