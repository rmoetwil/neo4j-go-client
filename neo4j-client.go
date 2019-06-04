package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func simpleTest() error {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		result  neo4j.Result
		err     error
	)

	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "neo4j", "")); err != nil {
		return err // handle error
	}
	// handle driver lifetime based on your application lifetime requirements
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per application
	defer driver.Close()

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer session.Close()

	result, err = session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
		"id":   1,
		"name": "Item 1",
	})
	if err != nil {
		return err // handle error
	}

	for result.Next() {
		fmt.Printf("Created Item with Id = '%d' and Name = '%s'\n", result.Record().GetByIndex(0).(int64), result.Record().GetByIndex(1).(string))
	}
	if err = result.Err(); err != nil {
		return err // handle error
	}

	return nil
}

func main() {
	var err = simpleTest()
	if err != nil {
		fmt.Println(err)
	}
}
