package paquetes

import (
	"fmt"
	"time"
)

func TimePackage() string {

	now := time.Now() // 2025-12-10 15:48:25.390779 -0300 -03 m=+0.000145543
	// Acceder a los valores individuales de la fecha y hora
	// now.Year()
	// now.Month()
	// now.Day()
	// now.Hour()
	// now.Minute()
	// now.Second()
	// now.Nanosecond()
	// now.Location()
	// now.Zone()
	// now.IsZero()
	// now.Unix()
	// now.UnixMilli()
	// now.UnixMicro()
	customTime := time.Date(1991, 9, 29, 13, 00, 00, 00, time.UTC)

	time2 := customTime.Add(time.Hour * 2)

	// comparaciones retornan booleanos

	// now.After(customTime)
	// now.Equal(customTime)
	// now.Compare(customTime)
	// now.Compare(customTime)

	// diferencia de tiempo
	// diff := now.Sub(customTime)
	// diff.Hours()
	// diff.Minutes()
	// diff.Seconds()
	// diff.Milliseconds()
	// diff.Microseconds()
	// diff.Nanoseconds()

	timeResta2 := customTime.Add(-time.Hour * 2)

	return fmt.Sprintln("\nTIME PACKAGE:", "\nnow:", now, "\ncustomTime: ", customTime, "\ncon 2 hs mas: ", time2, "\ncon 2 menos: ", timeResta2)
}
