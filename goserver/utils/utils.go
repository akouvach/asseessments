package utils

// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go

import (
	"fmt"
	"reflect"
)

type I interface{}

var dict = map[string]string{
	"Hello!":                 "Hallo!",
	"What's up?":             "Was geht?",
	"translate this":         "übersetze dies",
	"point here":             "zeige hier her",
	"translate this as well": "übersetze dies auch...",
	"and one more":           "und noch eins",
	"deep":                   "tief",
}

func ExtractDataFromJson(datos map[string]interface{},
	propiedad string,
	obj interface{}) string {
	var ok, encontrado bool
	var rdo string

	t := reflect.TypeOf(obj)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
	// v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		encontrado = false
		if field.Name == propiedad {
			encontrado = true
			/* Encontré la propiedad... voy a devolver lo que tiene en el nombre json */
			campoJson := field.Tag.Get("json")
			rdo, ok = datos[campoJson].(string)
			if !ok {
				fmt.Println("Lo encontré pero hubo una falla en convertirlo a string.... ",
					campoJson, rdo, campoJson, datos[campoJson], reflect.TypeOf(rdo))
				return "-1"
			}
			// fmt.Println("campo json", campoJson, rdo)
			break
		}

		// fmt.Println("\n sigo buscando..", propiedad, "-->", field.Name, field.Type, field.Tag.Get("json"))
	}

	if !encontrado {
		fmt.Println(propiedad, " no encontrada ")
	}
	// 	var x float64 = 3.4
	// v := reflect.ValueOf(x)
	// fmt.Println("type:", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value:", v.Float())

	fmt.Println("resultado", rdo, reflect.TypeOf(rdo))

	return rdo
	// switch t {
	// case "int":
	// 	return rdo.(int)
	// case "string":
	// 	return rdo.(string)
	// default:
	// 	// freebsd, openbsd,
	// 	// plan9, windows...
	// 	return rdo
	// }

}

func MostrarEstructura(obj interface{}) {
	// Wrap the original in a reflect.Value
	t := reflect.TypeOf(obj)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)
	// v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println("\n ", field.Name, field.Type)

	}

}

func Translate(obj interface{}) interface{} {
	// Wrap the original in a reflect.Value
	original := reflect.ValueOf(obj)

	copy := reflect.New(original.Type()).Elem()
	translateRecursive(copy, original)

	// Remove the reflection wrapper
	return copy.Interface()
}

func translateRecursive(copy, original reflect.Value) {
	switch original.Kind() {
	// The first cases handle nested structures and translate them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		// To get the actual value of the original we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		originalValue := original.Elem()
		// Check if the pointer is nil
		if !originalValue.IsValid() {
			return
		}
		// Allocate a new object and set the pointer to it
		copy.Set(reflect.New(originalValue.Type()))
		// Unwrap the newly created pointer
		translateRecursive(copy.Elem(), originalValue)

	// If it is an interface (which is very similar to a pointer), do basically the
	// same as for the pointer. Though a pointer is not the same as an interface so
	// note that we have to call Elem() after creating a new object because otherwise
	// we would end up with an actual pointer
	case reflect.Interface:
		// Get rid of the wrapping interface
		originalValue := original.Elem()
		// Create a new object. Now new gives us a pointer, but we want the value it
		// points to, so we have to call Elem() to unwrap it
		copyValue := reflect.New(originalValue.Type()).Elem()
		translateRecursive(copyValue, originalValue)
		copy.Set(copyValue)

	// If it is a struct we translate each field
	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			translateRecursive(copy.Field(i), original.Field(i))
		}

	// If it is a slice we create a new slice and translate each element
	case reflect.Slice:
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i += 1 {
			translateRecursive(copy.Index(i), original.Index(i))
		}

	// If it is a map we create a new map and translate each value
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			// New gives us a pointer, but again we want the value
			copyValue := reflect.New(originalValue.Type()).Elem()
			translateRecursive(copyValue, originalValue)
			copy.SetMapIndex(key, copyValue)
		}

	// Otherwise we cannot traverse anywhere so this finishes the the recursion

	// If it is a string translate it (yay finally we're doing what we came for)
	case reflect.String:
		translatedString := dict[original.Interface().(string)]
		copy.SetString(translatedString)

	// And everything else will simply be taken from the original
	default:
		copy.Set(original)
	}

}

// func SendMail() {
// 	from := mail.NewEmail("Eduardo", "akouvach@yahoo.com")
// 	subject := "Sending with SendGrid is Fun"
// 	to := mail.NewEmail("Andres", "akouvach@gmail.com")
// 	plainTextContent := "and easy to do anywhere, even with Go"
// 	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
// 	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
// 	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
// 	fmt.Println("api:", os.Getenv("SENDGRID_API_KEY"))
// 	response, err := client.Send(message)
// 	if err != nil {
// 		log.Println("Error:", err)
// 	} else {
// 		fmt.Println("Respuestas-------------------------------")
// 		fmt.Println(response.StatusCode)
// 		fmt.Println(response.Body)
// 		fmt.Println(response.Headers)
// 	}
// }
