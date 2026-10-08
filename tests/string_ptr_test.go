package fmt_test

import (
	std_fmt "fmt"
	"testing"
	. "webtyp.com/fmt"
)

func TestStringPointer(t *testing.T) {
	tests := []struct {
		name          string
		initialValue  string
		transform     func(*Conv) *Conv
		expectedValue string
	}{
		{
			name:         "Convert to lowercase with string pointer",
			initialValue: "HELLO WORLD",
			transform: func(t *Conv) *Conv {
				return t.ToLower()
			},
			expectedValue: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create string pointer with initial value
			originalPtr := tt.initialValue

			// Convert using string pointer and apply changes
			tt.transform(Convert(&originalPtr)).Apply()

			// Check if original pointer was updated correctly
			if originalPtr != tt.expectedValue {
				t.Errorf("\noriginalPtr = %q\nwant %q", originalPtr, tt.expectedValue)
			}
		})
	}
}

// Estos ejemplos ilustran cómo usar los punteros a strings para evitar asignaciones adicionales
func Example_stringPointerBasic() {
	// Creamos una variable string que queremos modificar
	myText := "HELLO World"

	// En lugar de crear una nueva variable con el resultado,
	// modificamos directamente la variable original usando Apply()
	Convert(&myText).ToLower().Apply()

	// La variable original ha sido modificada
	std_fmt.Println(myText)
	// Output: hello world
}

func Example_stringPointerCamelCase() {
	// Ejemplo de uso con múltiples transformaciones
	originalText := "el murcielago rapido"

	// Las transformaciones modifican la variable original directamente
	// usando el método Apply() para actualizar el puntero
	Convert(&originalText).CamelUp().Apply()

	std_fmt.Println(originalText)
	// Output: ElMurcielagoRapido
}

func Example_stringPointerEfficiency() {
	// En aplicaciones de alto rendimiento, reducir asignaciones de memoria
	// puede ser importante para evitar la presión sobre el garbage collector
	// Método tradicional (crea nuevas asignaciones de memoria)
	traditionalText := "Texto con ACENTOS"
	processedText := Convert(traditionalText).ToLower().String()
	std_fmt.Println(processedText)

	// Método con punteros (modifica directamente la variable original)
	directText := "Otro TEXTO con ACENTOS"
	Convert(&directText).ToLower().Apply()
	std_fmt.Println(directText)

	// Output:
	// texto con acentos
	// otro texto con acentos
}
