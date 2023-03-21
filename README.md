# spb-subway

(scroll down for english description)

Симулятор метро Санкт Петербурга. *Yandex metro clone

## Функции
0) Отрисовывает карту метро
1) Находит самый короткий путь между заданными станциями 
2) Отрисовывает оптимальный маршрут на экран и выводит его время
3) Исправляет ошибки ввода и неправильную раскладку клавиатуры

### 0. Отображение карты метро

![image](https://user-images.githubusercontent.com/111375726/226001931-1626532b-4c38-4973-90ad-c2603d3cf7f2.png)

HomeHandler (handlers/handlers.go) загружает файл assets/stations.json, получает информацию про станции и расстояния между ними, затем отобржает их в index.tmpl, строит неориентированный взвешенный граф по нему. Также HomeHandler сохраняет названия станций и сортирует их в алфавитном порядке. 

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

Name, X, Y, Line означает название станции, позиция по X Y и цвет ветки соответственно. Adj - станции, на которые можно проехать от данной станции, Dist - время проезда до них, Shift - с какой стороны находится текст относительно X, Y. Dx, Dy - смешение текста по X, Y. 

### 1. Поиск самого короткого пути между станциями 
FindRouteHandler (handlers/handlers.go) считывает две строки from, to из формы index.tmpl, находит подходящую станцию (см п.3), находит оптимальный маршрут и его время. 

Кратчайший путь находится с помощью алгоритма Дейкстры (dijkstra/dijkstra.go), который принимает номера станций в алфавитном порядке, возвращает минимальное время и массив из номеров станций, которые входят в этот путь. 

пример.
| From      | To           | Best route                                                                    | Duration |
|-----------|--------------|:-----------------------------------------------------------------------------:|---------:|
|Горьковская|Обводный канал|Горьковская, Невский, проспект, Сенная, Садовая, Звенигородская, Обводный канал| 18 минут |

### 2. Отображение кратчайшего маршрута.

![image](https://user-images.githubusercontent.com/111375726/226002641-9811083e-3962-4493-bf20-741bdd26c923.png)

ShowRouteHandler (handlers/handlers.go) строит карту метро, к ребрам не входящим в кратчайший маршрут добавляется эффект прозрачности 80%. 

### 3. Исправление ошибок ввода и неправильной раскладки клавиатуры
..* utils.EnToRu(s) - приминает строку, возвращает строку с замененными английскими буквами на русские, находящиеся на тех же клавищах клавиатуры. 

..* getIdx(s) - возвращает максимально "похожую" на s строку из всех названий станций. Строка s1 больше похоже на строку s, чем s2, если расстояние Левенштейна (utils.EditDistance()) между s1 и s меньше, чем между s2 и s. 

#### English desription.

Saint Petersburg subway simulator. *Yandex metro clone

## Functions 
0) Draw subway map
1) Find the shortes path between given stations
2) Draw this route
3) Correct mistakes in input

### 0. Subway stations drawing 

![image](https://user-images.githubusercontent.com/111375726/226001931-1626532b-4c38-4973-90ad-c2603d3cf7f2.png)

HomeHandler (handlers/handlers.go) loads asstes/stations.json file witch is describing stations, builds graph of stations and distances between them and draw it. Also HomeHandler saves stations names and sort them in alphabetic order. 

stations.json structure: 

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

Name, X, Y, Line stands for station name, X Y position and color of branch respectively. Adj is array of stations connecting with this one. Dist - time to reach them, Shift - offset side of the text. Dx, Dy - text shift by X and Y. 

### 1. Search the shortest route between given stations
FindRouteHandler (handlers/handlers.go) reads two string "from", "to" from the form in index.tmpl page, find the station you mean (p.3) and, finally, search the best route between from and to. 

Dijkstra algorithm (dijkstra/dijkstra.go) return the shortest path between two nodes. 

### 2. Shortest route Visualization 

![image](https://user-images.githubusercontent.com/111375726/226002641-9811083e-3962-4493-bf20-741bdd26c923.png)

ShowRouteHandler (handlers/handlers.go) builds subway map. Edges wich are not in shortest path are drawing transparent. 

### 3. Input mistakes correction

getIdx(s) - returns the most similar stations name to s. Similarity between string s and string t is the Levenshtein distance from s to t. Levenshtein distance is a string metric for measuring the difference between two strings. 
