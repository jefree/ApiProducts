package main

import (
  _ "products/routers"

  "github.com/astaxie/beego"

  "github.com/astaxie/beego/orm"
  _ "github.com/lib/pq"
  "github.com/astaxie/beego/plugins/cors"
)

func configDB() {
  orm.RegisterDriver("postgres", orm.DRPostgres)
  orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("db_config"))
}

func main() {
  if beego.BConfig.RunMode == "dev" {
    beego.BConfig.WebConfig.DirectoryIndex = true
    beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
  }
  // CORS for https://foo.* origins, allowing:
  // - PUT and PATCH methods
  // - Origin header
  // // - Credentials share

  beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options {
    AllowOrigins: []string{"*"},
    AllowMethods: []string{"PUT", "PATCH"},
    AllowHeaders: []string{"Origin"},
    ExposeHeaders: []string{"Content-Length"},
    AllowCredentials: true,
  }))


  configDB()
  beego.Run()
}
