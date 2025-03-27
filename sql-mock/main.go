package main

import (
	"fmt"
	"strconv"
	"strings"
)

// User संरचना डेटा का निरूपण करता है
type User struct {
	ID   int
	Name string
	Age  int
}

// Database एक सरल सूची है जो डेटा को स्टोर करती है
var Database = []User{
	{ID: 1, Name: "Alice", Age: 30},
	{ID: 2, Name: "Bob", Age: 25},
	{ID: 3, Name: "Charlie", Age: 35},
}

// parseCondition विश्लेषण करता है कि क्या दिया गया उपयोगकर्ता शर्त को पूरा करता है
func parseCondition(user User, condition string) bool {
	parts := strings.Split(condition, "=")
	if len(parts) != 2 {
		fmt.Println("Invalid condition format:", condition)
		return false
	}

	field := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	value = strings.Trim(value, "'")

	switch field {
	case "id":
		id, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting ID:", err)
			return false
		}
		return user.ID == id
	case "name":
		return user.Name == value
	case "age":
		age, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error converting Age:", err)
			return false
		}
		return user.Age == age
	default:
		fmt.Println("Unsupported field:", field)
		return false
	}
}

// ExecuteQuery फंक्शन SQL क्वेरीज़ का प्रोसेसिंग करता है
func ExecuteQuery(query string) {
	fmt.Println("Executing query:", query)
	var affectedRows int

	switch strings.ToLower(strings.Split(query, " ")[0]) {
	case "select":
		affectedRows = Select(query)
	case "insert":
		affectedRows = Insert(query)
	case "update":
		affectedRows = Update(query)
	case "delete":
		affectedRows = Delete(query)
	default:
		fmt.Println("Unsupported query type")
		return
	}

	fmt.Printf("Operation affected %d row(s).\n", affectedRows)
	fmt.Println("Current Database:", Database)
}

// Select क्वेरी व्यवस्थापन
func Select(query string) int {
	whereIndex := strings.Index(strings.ToLower(query), "where")
	var condition string
	if whereIndex != -1 {
		condition = query[whereIndex+5:]
	}

	var users []User
	for _, user := range Database {
		if condition == "" || parseCondition(user, condition) {
			users = append(users, user)
		}
	}
	fmt.Println("Select Result:", users)
	return len(users)
}

// Insert क्वेरी व्यवस्थापन
func Insert(query string) int {
	values := extractValuesFromInsertQuery(query)
	fmt.Println("Extracted values:", values)

	if len(values) >= 3 {
		id, err1 := strconv.Atoi(values[0])
		name := values[1]
		age, err2 := strconv.Atoi(values[2])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting values:", err1, err2)
			return 0
		}

		newUser := User{ID: id, Name: name, Age: age}
		Database = append(Database, newUser)
		return 1
	}
	return 0
}

// extractValuesFromInsertQuery INSERT क्वेरी से मान निकालता है
func extractValuesFromInsertQuery(query string) []string {
	// Find the last pair of parentheses (which corresponds to the VALUES clause)
	valuesStart := strings.LastIndex(query, "(")
	valuesEnd := strings.LastIndex(query, ")")

	if valuesStart == -1 || valuesEnd == -1 || valuesEnd <= valuesStart {
		fmt.Println("Could not find VALUES clause")
		return []string{}
	}

	valueStr := query[valuesStart+1 : valuesEnd]
	values := strings.Split(valueStr, ",")

	// स्पेस और उद्धरण हटाएं
	for i := range values {
		values[i] = strings.TrimSpace(strings.Trim(values[i], "'"))
	}

	return values
}

// Update क्वेरी व्यवस्थापन
func Update(query string) int {
	setIndex := strings.Index(strings.ToLower(query), "set")
	whereIndex := strings.Index(strings.ToLower(query), "where")

	if setIndex == -1 || whereIndex == -1 {
		fmt.Println("Invalid UPDATE query format")
		return 0
	}

	setClause := query[setIndex+3 : whereIndex]
	setParts := strings.Split(setClause, "=")
	if len(setParts) != 2 {
		fmt.Println("Invalid SET clause")
		return 0
	}

	field := strings.TrimSpace(setParts[0])
	value := strings.TrimSpace(strings.Trim(setParts[1], "'"))

	condition := query[whereIndex+5:]

	affectedRows := 0
	for i := range Database {
		if parseCondition(Database[i], condition) {
			switch field {
			case "name":
				Database[i].Name = value
			case "age":
				age, err := strconv.Atoi(value)
				if err == nil {
					Database[i].Age = age
				} else {
					fmt.Println("Error converting age:", err)
				}
			}
			affectedRows++
		}
	}
	return affectedRows
}

// Delete क्वेरी व्यवस्थापन
func Delete(query string) int {
	whereIndex := strings.Index(strings.ToLower(query), "where")
	if whereIndex == -1 {
		fmt.Println("No WHERE clause in DELETE query")
		return 0
	}

	condition := query[whereIndex+5:]

	var updatedDatabase []User
	affectedRows := 0

	for _, user := range Database {
		if !parseCondition(user, condition) {
			updatedDatabase = append(updatedDatabase, user)
		} else {
			affectedRows++
		}
	}

	Database = updatedDatabase
	return affectedRows
}

func main() {
	fmt.Println("Initial Database:", Database)

	ExecuteQuery("SELECT * FROM User WHERE id=1")
	ExecuteQuery("INSERT INTO User (id, name, age) VALUES (4, 'David', 40)")
	ExecuteQuery("SELECT * FROM User WHERE id=1")
	ExecuteQuery("UPDATE User SET name='Bob Updated' WHERE id=2 OR id=5")
	ExecuteQuery("DELETE FROM User WHERE id=3")
}

// I have to mock the CRUD operations (insert, read, update, delete) of a SQL database where it will take SQL queries (select, insert, update, delete) as input and process the data accordingly. Assuming the demo table already exists, it will have one or more `rows` and the queries will be only for a single row as of now which will be inserted, updated, deleted and retrieved a row. Based on the query, it will decide on how the operation will be performed and on which row. I want to implement two approaches in `Go` programming language for this problem so that data processing can be handled efficiently and not use any database to store the data, implement a data structure. I want to perform all the 4 operations only on 'row-level' right now