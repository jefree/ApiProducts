package main

import (
  _ "products/routers"

  "github.com/astaxie/beego"

  "github.com/astaxie/beego/orm"
  _ "github.com/lib/pq"
)

func configDB() {
  orm.RegisterDriver("postgres", orm.DRPostgres)
  orm.RegisterDataBase("default", "postgres", "postgres://advocates:aguacat3s@localhost:5432/products")
}

func main() {
  if beego.BConfig.RunMode == "dev" {
    beego.BConfig.WebConfig.DirectoryIndex = true
    beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
  }
  configDB()
  beego.Run()
}
