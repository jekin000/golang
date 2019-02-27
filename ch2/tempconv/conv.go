package tempconv

//CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//KToC exe 2.1
func KToC(k Kelvins) Celsius { return Celsius(Celsius(k) + AbsoluteZeroC) }

//CToK exe 2.1
func CToK(c Celsius) Kelvins { return Kelvins(c - AbsoluteZeroC) }
