# industrialTechnology
Практические занятия по предмету "Технологии индустриального программирования" 1 семестр.  
Автор: Балуева Мария, ЭФМО-02-24

## Практика 10. 
### 10.1. Создание базы данных интернет-магазина кондитерский изделий  

CREATE SCHEMA IF NOT EXISTS bakery;
SET search_path TO bakery;

-- Таблица accessGroup
CREATE TABLE accessGroup (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256)
);

-- Таблица city
CREATE TABLE city (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256)
);

-- Таблица address
CREATE TABLE address (
    id BIGINT PRIMARY KEY,
    cityId BIGINT REFERENCES city(id),
    street VARCHAR(256),
    houseNumber INT,
    buildingNumber VARCHAR(32),
    index VARCHAR(32)
);

-- Таблица bank
CREATE TABLE bank (
    bic VARCHAR(256) PRIMARY KEY,
    name VARCHAR(256),
    addressId BIGINT REFERENCES address(id)
);

-- Таблица payAccount
CREATE TABLE payAccount (
    number VARCHAR(256) PRIMARY KEY,
    clientId BIGINT,
    bankBic VARCHAR(256) REFERENCES bank(bic)
);

-- Таблица client
CREATE TABLE client (
    id BIGINT PRIMARY KEY,
    surname VARCHAR(256),
    name VARCHAR(256),
    patronymic VARCHAR(256),
    email VARCHAR(32),
    phoneNumber VARCHAR(32)
);

-- Таблица clientAddress
CREATE TABLE clientAddress (
    clientId BIGINT REFERENCES client(id),
    addressId BIGINT REFERENCES address(id),
    PRIMARY KEY (clientId, addressId)
);

-- Таблица app_user
CREATE TABLE app_user (
    id BIGINT PRIMARY KEY,
    login VARCHAR(32),
    password VARCHAR(32),
    empId BIGINT,
    clientId BIGINT REFERENCES client(id),
    createDate TIMESTAMP,
    isActive BOOLEAN
);

-- Таблица userAccess
CREATE TABLE userAccess (
    userId BIGINT REFERENCES app_user(id),
    groupId BIGINT REFERENCES accessGroup(id),
    PRIMARY KEY (userId, groupId)
);

-- Таблица job_position
CREATE TABLE job_position (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256)
);

-- Таблица employee
CREATE TABLE employee (
    id BIGINT PRIMARY KEY,
    positionId BIGINT REFERENCES job_position(id),
    surname VARCHAR(256),
    name VARCHAR(256),
    patronymic VARCHAR(256),
    phoneNumber VARCHAR(32),
    email VARCHAR(32),
    startDate TIMESTAMP,
    endDate TIMESTAMP,
    number VARCHAR(32),
    photoLink VARCHAR(256)
);

-- Таблица report
CREATE TABLE report (
    id BIGINT PRIMARY KEY,
    createDate TIMESTAMP,
    patternId INT REFERENCES pattern(id),
    fileLink VARCHAR(256),
    empId BIGINT REFERENCES employee(id)
);

-- Таблица pattern
CREATE TABLE pattern (
    id INT PRIMARY KEY,
    name VARCHAR(256),
    fileLink VARCHAR(256)
);

-- Таблица paymentType
CREATE TABLE paymentType (
    id INT PRIMARY KEY,
    name VARCHAR(256)
);

-- Таблица payment
CREATE TABLE payment (
    id BIGINT PRIMARY KEY,
    orderId BIGINT,
    date TIMESTAMP,
    sum FLOAT,
    payTypeId INT REFERENCES paymentType(id)
);

-- Таблица orderStatus
CREATE TABLE orderStatus (
    id INT PRIMARY KEY,
    name VARCHAR(32)
);

-- Таблица warehouse
CREATE TABLE warehouse (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256),
    addressId BIGINT REFERENCES address(id)
);

-- Таблица manufacture
CREATE TABLE manufacture (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256),
    addressId BIGINT REFERENCES address(id),
    email VARCHAR(256),
    phoneNumber VARCHAR(32)
);

-- Таблица category
CREATE TABLE category (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256)
);

-- Таблица subcategoryProduct
CREATE TABLE subcategoryProduct (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256)
);

-- Таблица packageType
CREATE TABLE packageType (
    id INT PRIMARY KEY,
    name VARCHAR(256)
);

-- Таблица OkeiDict
CREATE TABLE OkeiDict (
    id INT PRIMARY KEY,
    code INT,
    name VARCHAR(256)
);

-- Таблица product
CREATE TABLE product (
    id BIGINT PRIMARY KEY,
    categoryId BIGINT REFERENCES category(id),
    name VARCHAR(256),
    description VARCHAR(4000),
    proteins FLOAT,
    fats FLOAT,
    carbohydrates FLOAT,
    calories FLOAT,
    unWeight FLOAT,
    weight INT,
    packaged INT REFERENCES packageType(id),
    countTypePack INT,
    cost FLOAT,
    OkeiId INT REFERENCES OkeiDict(id),
    inStore BOOLEAN,
    photoLink VARCHAR(256)
);

-- Таблица priceHistory
CREATE TABLE priceHistory (
    productId BIGINT REFERENCES product(id),
    startDate TIMESTAMP,
    endDate TIMESTAMP,
    cost FLOAT,
    PRIMARY KEY (productId, startDate)
);

-- Таблица productInOrder
CREATE TABLE productInOrder (
    productId BIGINT REFERENCES product(id),
    orderId BIGINT,
    count INT,
    cost FLOAT,
    PRIMARY KEY (productId, orderId)
);

-- Таблица productInWarehouse
CREATE TABLE productInWarehouse (
    productId BIGINT REFERENCES product(id),
    validTo TIMESTAMP,
    count INT,
    warehouseId BIGINT REFERENCES warehouse(id),
    status BOOLEAN,
    PRIMARY KEY (productId, warehouseId)
);

-- Таблица manufactureProduct
CREATE TABLE manufactureProduct (
    productId BIGINT REFERENCES product(id),
    manufactureId BIGINT REFERENCES manufacture(id),
    PRIMARY KEY (productId, manufactureId)
);

-- Таблица order
CREATE TABLE "order" (
    id BIGINT PRIMARY KEY,
    name VARCHAR(256),
    createDate TIMESTAMP,
    addressId BIGINT REFERENCES address(id),
    statusId INT REFERENCES orderStatus(id),
    clientId BIGINT REFERENCES client(id),
    sumOrder FLOAT,
    isPay BOOLEAN,
    comment VARCHAR(2000),
    endDate TIMESTAMP,
    delStartDate TIMESTAMP,
    delEndDate TIMESTAMP,
    respEmpId BIGINT REFERENCES employee(id)
);

-- Таблица productInBasket
CREATE TABLE productInBasket (
    productId BIGINT REFERENCES product(id),
    clientId BIGINT REFERENCES client(id),
    count INT,
    PRIMARY KEY (productId, clientId)
);
