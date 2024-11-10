# industrialTechnology  
Практические занятия по предмету "Технологии индустриального программирования" 1 семестр.  
Автор: Балуева Мария, ЭФМО-02-24  

## Практика 8  
## Создание простого REST API на языке Go с использованием фреймворка Gin.  

### 8.1. Получение всех продуктов  
GET http://localhost:8080/products  
![alt text](image.png)  

### 8.2. Получение продукта по ID  
GET http://localhost:8080/products/1  
![alt text](image-1.png)  

### 8.3. Создание нового продукта
POST http://localhost:8080/products  
![alt text](image-2.png)  

GET http://localhost:8080/products/4  
![alt text](image-3.png)  

### 8.4. Обновление существующего продукта
PUT http://localhost:8080/products/4  
![alt text](image-4.png)  

GET http://localhost:8080/products/4  
![alt text](image-5.png)  

### 8.5. Удаление продукта
DELETE http://localhost:8080/products/4  
![alt text](image-6.png)  

GET http://localhost:8080/products/4  
![alt text](image-7.png)  

## Работа с корзиной

### 8.6. Получение всех продуктов в корзине  
GET http://localhost:8080/cart  
![alt text](image-8.png)  

### 8.7. Добавление продукта в корзину
POST http://localhost:8080/cart  
![alt text](image-9.png)  

GET http://localhost:8080/cart  
![alt text](image-10.png)  

### 8.8. Удаление продукта из корзины
DELETE http://localhost:8080/cart/2  
![alt text](image-11.png)  

GET http://localhost:8080/cart  
![alt text](image-12.png)  