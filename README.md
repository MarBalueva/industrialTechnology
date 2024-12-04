# industrialTechnology  
Практические занятия по предмету "Технологии индустриального программирования" 1 семестр.  
Автор: Балуева Мария, ЭФМО-02-24  

## Практика 13
## Асинхронная обработка и задачи в фоновом режиме
Ссылка на код: https://github.com/MarBalueva/industrialTechnology/blob/pr13/main.go  

### 13.1. Логирование действий в консоль для каждого этапа выполнения задач  
POST http://localhost:8080/tasks  
![alt text](/img/image.png)  

### 13.2. Отмена задач  
POST http://localhost:8080/tasks  
![alt text](image-1.png)  

POST http://localhost:8080/tasks/20241205010619/cancel  
![alt text](image.png)  

### 13.3. Ограничение на количество одновременно выполняющихся задач  
POST http://localhost:8080/tasks  
![alt text](image-2.png)  

Попытка создать новую задачу  
![alt text](image-3.png)  

Сообщение о завершении всех задач  
![alt text](image-4.png)  