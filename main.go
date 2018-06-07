package main

import "github.com/HAL-RO-Developer/caseTeamB_server/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}