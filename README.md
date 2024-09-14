# industrialTechnology
Практические занятия по предмету "Технологии индустриального программирования" 1 семестр. Автор: Балуева Мария, ЭФМО-02-24\n

Практика 1. 
1.1. Логическая модель данных предметной области "Интернет-магазин кондитерских изделий" 

![логическая модель данных (1)](https://github.com/user-attachments/assets/10fde5bf-a512-4455-9811-5b9e7434b6d9)

Описание сущностей:
Модель содержит следующие основные сущности: 
•	продукция - информация о продаваемой продукции; 
•	история цен - историчность цен на продаваемую продукцию;
•	склады - информация о местах хранения продукции; 
•	заказы - информация о заказах пользователей; 
•	оплаты - информация об оплатах заказов; 
•	адреса - информация об адресах доставки и складов; 
•	банки - информация о банках; 
•	корреспондентские счета банков - информация о корреспондентских счетах банков;
•	юридические лица - информация о клиентах-юридических лицах;
•	расчетные счета юридических лиц - информация о расчетных счетах юридических лиц;
•	физические лица - информация о клиентах-физических лицах;
• счета физических лиц - информация о счетах физических лиц;
• клиенты - информация о клиентах и их типах;
•	документы - информация об используемых документах;
•	сотрудники - информация о сотрудниках организации;
•	отчеты - информация о формируемых отчетах;
• пользователи - информация о пользователях.
Сущности-справочники предполагают наполнение значениями до ввода системы в эксплуатацию. Они необходимы для исключения ручного ввода повторяющихся или установленных в организации значений некоторых атрибутов основных сущностей. К таким сущностям относятся:
•	подкатегории продукции - справочник подкатегорий продукции;
•	категории продукции - справочник категорий продукции;
•	справочник ОКЕИ - справочник единиц измерения продукции;
•	виды упаковок - справочник видов упаковок;
•	способы оплаты - справочник способов оплаты;
•	города - справочник городов;
•	улицы - справочник улиц;
•	статусы заказа - справочник статусов заказов;
•	должности - справочник должностей;
•	шаблоны - справочник шаблонов документов;
• типы клиентов - справочник типов клиентов;
• группы доступа - справочник групп доступа пользователей.
Некоторые сущности связаны друг с другом отношением многие ко многим, например, продукция и заказы, продукция и склады, а также юридические лица и физические лица. Для реализации таких связей выделены вспомогательные таблицы:
•	продукция на складе - связь между продукцией и складами хранения;
•	контактные лица юридического лица - связь между юридическими лицами и их представителями;
•	продукция в заказе - связь между продукцией и заказами;
• доступы пользователей - связь между пользователями и их группами доступа;
• производители продукции - связь между продукцией и ее производителями.

1.2. Физическая модель данных

![физическая модель данных (1)](https://github.com/user-attachments/assets/0377dbfb-2e05-4736-824c-82257648387a)

1.3. Диаграмма use-case

![варианты использования](https://github.com/user-attachments/assets/78fca2ed-def3-45b1-996f-38171ed8aad8)

1.4. План проекта по инкрементной модели:

1. **Анализ требований**
- Выявление требований к системе.
- Создание спецификации требований, основанной на функциональных блоках схемы, таких как:
     - Просмотр продукции
     - Добавление в корзину/избранное
     - Оформление заказа
     - Взаимодействие с документами, создание отчетов
     - Рабочие процессы для менеджеров и руководителей.
- Разделение требований на инкременты.

2. **Проектирование архитектуры**
   - Проектирование общей архитектуры системы.
   - Определение ключевых модулей (каталог продукции, корзина, личный кабинет, система управления заказами, документы и отчеты).
   - Выбор технологий и базовых компонентов.

3. **Первый инкремент: Основной функционал пользователя (Клиент)**
   **Функционал:**
   - Регистрация и авторизация (вход в личный кабинет).
   - Просмотр каталога продукции.
   - Добавление товаров в корзину и оформление заказа.
   **Этапы:**
   - Реализация базы данных для хранения данных о продуктах и клиентах.
   - Разработка базового интерфейса для просмотра каталога и управления корзиной.
   - Тестирование первичной версии системы.

4. **Второй инкремент: Дополнительный функционал пользователя (Клиент)**
   **Функционал:**
   - Просмотр истории заказов.
   - Изменение личных данных.
   **Этапы:**
   - Добавление функционала для работы с историей заказов.
   - Тестирование изменения данных пользователя и сохранения их в системе.
   - Интеграция личного кабинета с системой учета заказов.

5. **Третий инкремент: Функционал для менеджера**
   **Функционал:**
   - Просмотр данных по заказам и клиентам.
   - Работа с документами: создание счета на оплату, подписание договоров, загрузка документов.
   **Этапы:**
   - Создание интерфейса и API для менеджера.
   - Интеграция с функционалом обработки заказов.
   - Обеспечение работы с документацией и загрузки данных.
   - Тестирование взаимодействия между клиентом и менеджером.

6. **Четвертый инкремент: Функционал для руководителя**
   **Функционал:**
   - Формирование отчетности.
   - Просмотр данных о сотрудниках и продукции.
   - Загрузка сертификатов.
   **Этапы:**
   - Реализация отчетов по данным из базы.
   - Разработка интерфейса для управления данными сотрудников и продукции.
   - Обеспечение выгрузки отчетов и сертификатов.
   - Тестирование и проверка всей системы.

7. **Пятый инкремент: Завершение системы**
   **Функционал:**
   - Интеграция всех модулей.
   - Обеспечение безопасности данных и взаимодействия с внешними системами.
   **Этапы:**
   - Полное тестирование всех функциональных блоков.
   - Тестирование безопасности.

Финальные этапы:
- **Внедрение**: Развертывание системы на реальных серверах, обучение сотрудников (менеджеров и руководителей).
- **Поддержка и доработка**: После завершения проекта продолжается мониторинг системы, устранение багов, возможная интеграция новых функций по мере появления новых требований.
