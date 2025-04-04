package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml: "plant"`
	Id      int      `xml: "id,attr"`
	Name    string   `xml: name`
	Origin  []string `xml: origin`
}

func (self *Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", self.Id, self.Name, self.Origin)
}

type PlantCollection struct {
	XMLName xml.Name `xml: "nesting"`
	Plants  []Plant  `xml: "parent>child>plant"`
}

func _xml() {
	coffee := &Plant{Id: 27, Name: "Coffee", Origin: []string{"Ethiopia", "Brazil"}}
	coffeeXml, _ := xml.Marshal(coffee)
	fmt.Println(string(coffeeXml))
	fmt.Println(xml.Header + string(coffeeXml))

	indentedCoffeeXml, _ := xml.MarshalIndent(coffee, " ", "	")
	fmt.Println(string(indentedCoffeeXml))

	var unmarshalledCoffee Plant
	_error := xml.Unmarshal(coffeeXml, &unmarshalledCoffee)
	if _error != nil {
		panic(_error)
	}
	fmt.Println(unmarshalledCoffee)

	tomato := &Plant{Id: 81, Name: "Tomato", Origin: []string{"Mexico", "California"}}
	collection := &PlantCollection{Plants: []Plant{*tomato, *coffee}}
	collectionXml, _ := xml.MarshalIndent(collection, " ", "	")
	fmt.Println(string(collectionXml))
}
