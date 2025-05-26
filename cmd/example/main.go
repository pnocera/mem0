package main

import (
	"encoding/json"
	"fmt"
	"log"

	"mem0/graphs" // Assuming module name is 'mem0' as per go.mod
	"mem0/vectorstores"
)

func main() {
	fmt.Println("--- Graphs Package Example ---")

	// 1. GraphStoreConfig Unmarshalling
	graphConfigJSON := `{
		"provider": "neo4j",
		"config": {
			"url": "bolt://localhost:7687",
			"username": "neo4j",
			"password": "password",
			"database": "neo4j",
			"base_label": true
		},
		"custom_prompt": "custom graph prompt"
	}`
	var graphCfg graphs.GraphStoreConfig
	err := json.Unmarshal([]byte(graphConfigJSON), &graphCfg)
	if err != nil {
		log.Fatalf("Error unmarshalling GraphStoreConfig: %v", err)
	}
	fmt.Printf("GraphStoreConfig Provider: %s\n", graphCfg.Provider)
	if neo4jConf, ok := graphCfg.Config.(*graphs.Neo4jConfig); ok {
		fmt.Printf("Neo4j URL: %s, BaseLabel: %t\n", neo4jConf.URL, neo4jConf.BaseLabel)
	} else {
		fmt.Println("Could not assert graphCfg.Config to *graphs.Neo4jConfig")
	}
	fmt.Printf("Custom Graph Prompt: %s\n", graphCfg.CustomPrompt)

	// Validate the unmarshalled config
	err = graphCfg.Validate()
	if err != nil {
		fmt.Printf("GraphStoreConfig validation error: %v\n", err)
	} else {
		fmt.Println("GraphStoreConfig validated successfully.")
	}

	// 2. Accessing an exported tool definition
	fmt.Printf("AddMemoryToolGraph Name: %s\n", graphs.AddMemoryToolGraph.Function.Name)
	fmt.Printf("AddMemoryToolGraph Description: %s\n", graphs.AddMemoryToolGraph.Function.Description)

	// 3. Calling GetDeleteMessages
	systemMsg, userMsg := graphs.GetDeleteMessages("graph_mem1, graph_mem2", "delete graph_mem1", "user123")
	fmt.Printf("GetDeleteMessages System: %s\n", systemMsg)
	fmt.Printf("GetDeleteMessages User: %s\n", userMsg)

	fmt.Println("\n--- Vectorstores Package Example ---")

	// 1. VectorStoreConfig Unmarshalling
	vectorConfigJSON := `{
		"provider": "qdrant",
		"config": {
			"address": "http://localhost:6333",
			"api_key": "some_api_key",
			"collection_name": "my_vectors"
		}
	}`
	var vectorCfg vectorstores.VectorStoreConfig
	err = json.Unmarshal([]byte(vectorConfigJSON), &vectorCfg)
	if err != nil {
		log.Fatalf("Error unmarshalling VectorStoreConfig: %v", err)
	}
	fmt.Printf("VectorStoreConfig Provider: %s\n", vectorCfg.Provider)
	if qdrantConf, ok := vectorCfg.Config.(*vectorstores.QdrantConfig); ok {
		fmt.Printf("Qdrant Address: %s, Collection: %s\n", qdrantConf.Address, qdrantConf.CollectionName)
	} else {
		fmt.Println("Could not assert vectorCfg.Config to *vectorstores.QdrantConfig")
	}

	// Validate the unmarshalled config
	err = vectorCfg.Validate()
	if err != nil {
		fmt.Printf("VectorStoreConfig validation error: %v\n", err)
	} else {
		fmt.Println("VectorStoreConfig validated successfully.")
	}

	// 2. Instantiate QdrantStore and call a method
	qStore := vectorstores.QdrantStore{}
	collections, err := qStore.ListCollections()
	if err != nil {
		fmt.Printf("QdrantStore.ListCollections() error: %v\n", err)
	} else {
		fmt.Printf("QdrantStore.ListCollections() result: %v\n", collections)
	}

	// Example of using another method
	searchResult, err := qStore.GetVector("test_collection", "vector_id_1")
	if err != nil {
		fmt.Printf("QdrantStore.GetVector() error: %v\n", err)
	} else {
		fmt.Printf("QdrantStore.GetVector() result: %v\n", searchResult)
	}
}
