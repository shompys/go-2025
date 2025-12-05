# Teoria

## Arrange /Given

Requisitos que debe cumplir el código principal
En la preparación se establecen las condiciones iniciales.

## Act / When

Proceso de creación, donde vamos acumulando los resultados que analizaremos
En la acción, se ejexuta el código bajo prueba

## Assert / Then

Momento donde validamos si los resultados son correctos o incorrectos
En la comprobación, se verifica si los resultados son los esperados.

```go
require.Equal(t, 123, 123, "deberían ser iguales") //validar Igualdad
require.NotEqual(t, 123, 456, "no deberían ser iguales") //validar desigualdad
require.Nil(t, object)//validar Nulo esperado(bueno para errores)
require.NotNill(t, object)//validar no nulo esperado (Bueno para cuando esperamos algo)
require.Error(t, err) // valida que un error no sea de tipo nil
```
