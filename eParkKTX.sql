USE eparkktx;

-- Bảng USER
CREATE TABLE User (
    UserID VARCHAR(20) PRIMARY KEY,
    user_password VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    DoB DATE,
    gender VARCHAR(10),
    phone_number VARCHAR(20) NOT NULL,
    Role ENUM('Dormitory_Management', 'Park_Management', 'Student') NOT NULL
);

-- Bảng BQL_KTX
CREATE TABLE Dormitory_Management (
    UserID VARCHAR(20) PRIMARY KEY,
    building VARCHAR(20) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES User(UserID)
        ON DELETE CASCADE
);

-- Bảng QL_Baixe
CREATE TABLE Park_Management (
    UserID VARCHAR(20) PRIMARY KEY,
    park_name VARCHAR(20) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES User(UserID)
        ON DELETE CASCADE
);

-- Bảng Sinhvien
CREATE TABLE Student (
    UserID VARCHAR(20) PRIMARY KEY,
    school VARCHAR(30),
    room VARCHAR(20),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
        ON DELETE CASCADE
);

-- Bảng NHA_XE
CREATE TABLE Garage (
    GarageID VARCHAR(20) PRIMARY KEY,
    GarageName VARCHAR(100) NOT NULL,
    size INT NOT NULL,
    UserID VARCHAR(20) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Park_Management(UserID)
        ON DELETE CASCADE
);

-- Bảng XE
CREATE TABLE Vehicle (
    Number_plate VARCHAR(30) PRIMARY KEY,
    vehicle_type VARCHAR(50) NOT NULL,
    register_date DATE NOT NULL,
    color VARCHAR(20) NOT NULL,
    ID_Student VARCHAR(20) NOT NULL,
    ID_ParkManagement VARCHAR(20) NOT NULL,
    FOREIGN KEY (ID_Student) REFERENCES Student(UserID)
        ON DELETE CASCADE,
    FOREIGN KEY (ID_ParkManagement) REFERENCES Park_Management(UserID)
        ON DELETE CASCADE
);

-- Bảng HOP_DONG
CREATE TABLE Contract (
    Contract_ID VARCHAR(20) PRIMARY KEY,
    Start_Date DATE NOT NULL,
    End_Date DATE NOT NULL,
    Contract_Type VARCHAR(30) NOT NULL,
    Cost BIGINT NOT NULL,
    ID_ParkManagement VARCHAR(20) NOT NULL,
    ID_DormitoryManagement VARCHAR(20),
    Number_Plate VARCHAR(30) NOT NULL,
    FOREIGN KEY (ID_ParkManagement) REFERENCES Park_Management(UserID)
        ON DELETE CASCADE,
    FOREIGN KEY (ID_DormitoryManagement) REFERENCES Dormitory_Management(UserID)
        ON DELETE SET NULL,
    FOREIGN KEY (Number_Plate) REFERENCES Vehicle(Number_plate)
        ON DELETE CASCADE
);

-- Bảng Lich_su_ra_vao
CREATE TABLE EE_History (
    time_date DATETIME NOT NULL,
    status ENUM('in','out') NOT NULL,
    Number_plate VARCHAR(30) NOT NULL,
    ID_Garage VARCHAR(20) NOT NULL,
    PRIMARY KEY (time_date, Number_plate),
    FOREIGN KEY (Number_plate) REFERENCES Vehicle(Number_plate)
        ON DELETE CASCADE,
    FOREIGN KEY (ID_Garage) REFERENCES Garage(GarageID)
        ON DELETE CASCADE
);

-- Kiểm tra các bảng đã tạo
SHOW TABLES;
