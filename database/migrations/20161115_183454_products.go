package main

import (
  "github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Products_20161115_183454 struct {
  migration.Migration
}

// DO NOT MODIFY
func init() {
  m := &Products_20161115_183454{}
  m.Created = "20161115_183454"
  migration.Register("Products_20161115_183454", m)
}

// Run the migrations
func (m *Products_20161115_183454) Up() {
  // use m.SQL("CREATE TABLE ...") to make schema update
  m.SQL("CREATE TABLE products(id serial primary key,name TEXT NOT NULL,description TEXT NOT NULL,price integer DEFAULT NULL)")
}

// Reverse the migrations
func (m *Products_20161115_183454) Down() {
  // use m.SQL("DROP TABLE ...") to reverse schema update
  m.SQL("DROP TABLE products")
}
