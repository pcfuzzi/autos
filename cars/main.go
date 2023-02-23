package main

import "fmt"

type Cars struct {
	Marke         string
	Model         string
	Baujahr       int
	PS            int
	Farbe         string
	Verkaufspreis float64
	Sitze         int
	Türen         int
	Klimaanlage   string
	Verbrauch     float64
	Traumfahrzeug bool
}

func (c *Cars) BerechneVerbrauch(ortschaft, ueberland, autobahn float64) {
	c.Verbrauch = (ortschaft + ueberland + autobahn) / 2.0
}

func main() {
	// Erstellen eines neuen Auto-Objekts
	auto1 := Cars{
		Marke:         "Bentley",
		Model:         "Continental GT Speed",
		Baujahr:       2021,
		PS:            635,
		Farbe:         "Arnage Peacock Metallic",
		Verkaufspreis: 248810.000,
		Sitze:         5,
		Türen:         2,
		Klimaanlage:   "Ja",
		Traumfahrzeug: true,
	}
	// Erstellen eines weiteren Auto-Objekts
	auto2 := Cars{
		Marke:         "Bugatti",
		Model:         "Chiron Super Sport",
		Baujahr:       2020,
		PS:            1600,
		Farbe:         "JET GREY",
		Verkaufspreis: 3850000.000,
		Sitze:         2,
		Türen:         2,
		Klimaanlage:   " Ja ",
		Traumfahrzeug: false,
	}

	// Berechnung des Verbrauchs
	auto2.BerechneVerbrauch(40.310, 22.150, 41.705)

	// Ausgabe der Eigenschaften des Auto-Objekts
	fmt.Printf("Marke: %s\nModel: %s\nBaujahr: %d\nPS: %d\nFarbe: %s\nVerkaufspreis: %.2f EUR\nSitze: %d\nTüren: %d\nKlimaanlage: %s\nVerbrauch: %.2f l/100km\nTraumfahrzeug: %t\n",
		auto2.Marke, auto2.Model, auto2.Baujahr, auto2.PS, auto2.Farbe, auto2.Verkaufspreis, auto2.Sitze, auto2.Türen, auto2.Klimaanlage, auto2.Verbrauch, auto2.Traumfahrzeug)

	// Berechnung des Verbrauchs
	auto1.BerechneVerbrauch(17.7, 6.9, 12.2)

	// Ausgabe der Eigenschaften des Auto-Objekts
	fmt.Printf("Marke: %s\nModel: %s\nBaujahr: %d\nPS: %d\nFarbe: %s\nVerkaufspreis: %.2f EUR\nSitze: %d\nTüren: %d\nKlimaanlage: %s\nVerbrauch: %.2f l/100km\nTraumfahrzeug: %t\n",
		auto1.Marke, auto1.Model, auto1.Baujahr, auto1.PS, auto1.Farbe, auto1.Verkaufspreis, auto1.Sitze, auto1.Türen, auto1.Klimaanlage, auto1.Verbrauch, auto1.Traumfahrzeug)
}
