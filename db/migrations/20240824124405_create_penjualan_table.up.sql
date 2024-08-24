-- Migration to create Marketing table

CREATE TABLE marketing (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

-- Migration to create Penjualan table

CREATE TABLE Penjualan (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    TransactionNumber VARCHAR(255) NOT NULL,
    MarketingID INT,
    Date DATE,
    CargoFee DECIMAL(10,2),
    TotalBalance DECIMAL(10,2),
    GrandTotal DECIMAL(10,2),
    FOREIGN KEY (MarketingID) REFERENCES marketing(ID) ON DELETE SET NULL
);


-- Migration script to create Pembayaran table

CREATE TABLE Pembayaran (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    MarketingID INT NOT NULL,
    Amount DECIMAL(10,2) NOT NULL,
    PaymentDate DATE NOT NULL,
    Status VARCHAR(255) NOT NULL,
    FOREIGN KEY (MarketingID) REFERENCES marketing(ID) ON DELETE CASCADE
);

