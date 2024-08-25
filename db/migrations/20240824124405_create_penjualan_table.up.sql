-- Create Marketing table with ID starting from 1
CREATE TABLE marketing (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
) AUTO_INCREMENT = 1;

-- Create Penjualan table with ID starting from 1
CREATE TABLE Penjualan (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    TransactionNumber VARCHAR(255) NOT NULL UNIQUE,
    MarketingID INT,
    Date DATE NOT NULL,
    CargoFee DECIMAL(10,2) DEFAULT 0,
    TotalBalance DECIMAL(10,2) NOT NULL,
    GrandTotal DECIMAL(20,2) NOT NULL,
    FOREIGN KEY (MarketingID) REFERENCES marketing(ID) ON DELETE SET NULL,
    INDEX (TransactionNumber)
) AUTO_INCREMENT = 1;

-- Create Pembayaran table with ID starting from 1
CREATE TABLE Pembayaran (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    MarketingID INT NOT NULL,
    Amount DECIMAL(10,2) NOT NULL CHECK (Amount > 0),
    PaymentDate DATE NOT NULL,
    Status VARCHAR(255) NOT NULL,
    FOREIGN KEY (MarketingID) REFERENCES marketing(ID) ON DELETE CASCADE,
    INDEX (Status)
) AUTO_INCREMENT = 1;
