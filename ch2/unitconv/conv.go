package unitconv

//Convert Feets to Metres
func FToM(f Feets) Metres { return Metres(f / 0.3048) }

//Convert Metres to Feets
func MToF(m Metres) Feets { return Feets(m * 0.3048) }

//Convert Pounds to Kilograms
func PToK(p Pounds) Kilograms { return Kilograms(p * 2.2046226) }

//Convert Kilograms to Pounds
func KToP(k Kilograms) Pounds { return Pounds(k * 0.45359237) }
