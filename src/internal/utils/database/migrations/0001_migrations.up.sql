CREATE TABLE task_statuses (
    id SERIAL PRIMARY KEY,
    status TEXT NOT NULL
);

INSERT INTO task_statuses (status) VALUES ('new'), ('in_progress'), ('done');

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_status FOREIGN KEY (status_id) REFERENCES task_statuses(id) ON DELETE RESTRICT
);

INSERT INTO tasks (title, description, status_id) VALUES
    ('Разработка API', 'Создать эндпоинты для управления задачами', 1),
    ('Написание документации', 'Описать все методы API', 2),
    ('Тестирование функционала', 'Проверить работу всех эндпоинтов', 3),
    ('Оптимизация базы данных', 'Настроить индексы и оптимизировать запросы', 1);
