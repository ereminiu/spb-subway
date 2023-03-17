# spb-subway

(scroll down for english description)

Симулятор метро Санкт Петербурга. *Yandex metro clone

## Функции
0) Отрисовывает карту метро
1) Находит самый быстрый путь между заданными станциями 
2) Отрисовывает оптимальный маршрут на экран и выводит его время
3) Исправление ошибок ввода и неправильной раскладки клавиатуры

### 0. Отображение карты метро

![image](https://user-images.githubusercontent.com/111375726/226001931-1626532b-4c38-4973-90ad-c2603d3cf7f2.png)

HomeHandler (handlers/handlers.go) загружает файл stations.json, получает станции и расстояния между ними, затем отобржает их в index.tmpl строит неориентированный взвешенный граф по нему. Также HomeHandler сохраняет названия станций и сортирует их по алфавитному порядку. 

структура stations.json:

```json
"st_g0": {
  "Name": "Беговая",
  "X": 93,
  "Y": 140,
  "Adj": ["st_g1"],
  "Dist": [4],
  "Line": "green",
  "Shift": "left",
  "Dx": 0,
  "Dy": 0
}
```

Name, X, Y, Line означает название станции, цвет ветки, позиция по X и Y соответственно. Adj - станции, на которые можно проехать от данной станции, Dist - время проезда до них, Shift - с какой стороны находится текст относительно X, Y. Dx, Dy - смешение текста по X, Y. 

### 1. Поиск самого быстрого пути между станциями 
FindRouteHandler (handlers/handlers.go) считывает две строки from, to из формы, находит подходящую станцию (см п.3), находит оптимальный маршрут и его время. 

Кратчайший путь находится с помощью алгоритма Дейкстры (dijkstra/dijkstra.go), который принимает номер станций в алфавитном порядке, возвращает минимальное время и массив из номеров станций, которые входят в этот путь. 

| From      | To           | Best route                                                                    | Duration |
|-----------|--------------|:-----------------------------------------------------------------------------:|---------:|
|Горьковская|Обводный канал|Горьковская, Невский, проспект, Сенная, Садовая, Звенигородская, Обводный канал| 18 минут |

Пример. from="Горьковская" to="Обводный канал" время маршрута - 18 минут, маршрут (route):[Горьковская, Невский, проспект, Сенная, Садовая, Звенигородская, Обводный канал]

### 2. Отображение кратчайшего маршрута.

![image](https://user-images.githubusercontent.com/111375726/226002641-9811083e-3962-4493-bf20-741bdd26c923.png)

ShowRouteHandler (handlers/handlers.go) строит карту метро, к ребрам не входящим в кратчайший маршрут добавляется эффект прозрачности 80%. 

### 3. Исправление ошибок ввода и неправильной раскладки клавиатуры
..* utils.EnToRu(s) - приминает строку, возвращает строку с замененными английскими буквами на русские, находящиеся на тех же клавищах клавиатуры. 

..* getIdx(s) - возвращает максимально "похожую" на s строку из всех названий станций. Строка s1 больше похоже на строку s, чем s2, если расстояние Левенштейна (utils.EditDistance()) между s1 и s меньше, чем между s2 и s. 