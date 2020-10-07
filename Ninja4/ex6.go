package main

import (
	"fmt"
)

func main() {

	counties := make([]string, 50, 120)

	counties = []string{
		"Bath and North East Somerset",
		"Bedfordshire",
		"Berkshire",
		"Bristol",
		"Buckinghamshire",
		"Cambridgeshire",
		"Cheshire",
		"Cornwall",
		"County Durham",
		"Cumbria",
		"Derbyshire",
		"Devon",
		"Dorset",
		"East Riding of Yorkshire",
		"East Sussex",
		"Essex",
		"Gloucestershire",
		"Greater London",
		"Greater Manchester",
		"Hampshire",
		"Herefordshire",
		"Hertfordshire",
		"Isle of Wight",
		"Isles of Scilly",
		"Kent",
		"Lancashire",
		"Leicestershire",
		"Lincolnshire",
		"Merseyside",
		"Norfolk",
		"North Somerset",
		"North Yorkshire",
		"Northamptonshire",
		"Northumberland",
		"Nottinghamshire",
		"Oxfordshire",
		"Rutland",
		"Shropshire",
		"Somerset",
		"South Gloucestershire",
		"South Yorkshire",
		"Staffordshire",
		"Suffolk",
		"Surrey",
		"Tyne & Wear",
		"Warwickshire",
		"West Midlands",
		"West Sussex",
		"West Yorkshire",
		"Wiltshire",
		"Worcestershire"}

	fmt.Println(len(counties))
	fmt.Println(cap(counties))

	for i := 0; i < len(counties); i++ {
		fmt.Println("Index is:", i, " and county is ", counties[i])
	}
}
