package genmath

import (
	"math"
)

const (
	DEG2RAD  = 0.01745329251994329576923690768488612713442871888541725456097191 // Tau / 360
	RAD2DEG  = 57.2957795130823208767981548141051703324054724665643215491602439 // 360 / Tau
	PI       = 3.14159265358979323846264338327950288419716939937510582097494459 // Circumference / Diameter
	TAU      = 6.28318530717958647692528676655900576839433879875021164194988918 // 2 * Pi
	E        = 2.71828182845904523536028747135266249775724709369995957496696763 // Euler's constant (e)
	PHI      = 1.61803398874989484820458683436563811772030917980576286213544862 // Phi (golden ratio)
	LOG_E_2  = 0.69314718055994530941723212145817656807550013436025525412068001 // Log base e of 2
	LOG_2_E  = 1.44269504088896340735992468100189213742664595415298593413544941 // Log base 2 of e
	LOG_2_10 = 3.32192809488736234787031942948939017586483139302458061205475640 // Log base 2 of 10
	LOG_10_2 = 0.30102999566398119521373889472449302676818988146210854131042746 // Log base 10 of 2
	LOG_E_10 = 2.30258509299404568401799145468436420760110148862877297603332790 // Log base e of 10
	LOG_10_E = 0.43429448190325182765112891891660508229439700580366656611445378 // Log base 10 of e
	SQRT_2   = 1.41421356237309504880168872420969807856967187537694807317667974 // Square root of 2
	SQRT_E   = 1.64872127070012814684865078781416357165377610071014801157507931 // Square root of e
	SQRT_PHI = 1.27201964951406896425242246173749149171560804184009624861664038 // Square root of Phi (golden ratio)
	SQRT_PI  = 1.77245385090551602729816748334114518279754945612238712821380779 // Square root of Pi
	SQRT_TAU = 2.50662827463100050241576528481104525300698674060993831662992358 // Square root of Tau

	MAX_U8    = math.MaxUint8
	MAX_I8    = math.MaxInt8
	MIN_I8    = math.MinInt8
	MAX_U16   = math.MaxUint16
	MAX_I16   = math.MaxInt16
	MIN_I16   = math.MinInt16
	MAX_U32   = math.MaxUint32
	MAX_I32   = math.MaxInt32
	MIN_I32   = math.MinInt32
	MAX_U64   = math.MaxUint64
	MAX_I64   = math.MaxInt64
	MIN_I64   = math.MinInt64
	MAX_F32   = math.MaxFloat32
	SMALL_F32 = math.SmallestNonzeroFloat32
	MAX_F64   = math.MaxFloat64
	SMALL_F64 = math.SmallestNonzeroFloat64
)

type Integer interface {
	~uint | ~int | ~uint8 | ~int8 | ~uint16 | ~int16 | ~uint32 | ~int32 | ~uint64 | ~int64 | ~uintptr
}

type Float interface {
	~float32 | float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Real interface {
	Integer | Float
}

type Number interface {
	Integer | Float | Complex
}

type Bool interface {
	~bool
}

func Abs[T Real](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

func Min[T Real](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Real](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Clamp[T Real](min, val, max T) T {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

func IMod[T Real](val, div T) T {
	negV, negD := val < 0, div < 0
	if negV {
		val = -val
	}
	if negD {
		div = -div
	}
	uVal := uint64(val)
	uDiv := uint64(div)
	uMod := uVal % uDiv
	mod := T(uMod)
	if (negV && !negD) || (!negV && negD) {
		mod = -mod
	}
	return mod
}

func FMod[T Real](val, div T) T {
	negV, negD := val < 0, div < 0
	if negV {
		val = -val
	}
	if negD {
		div = -div
	}
	fVal, fDiv := float64(val), float64(div)
	fMod := math.Mod(fVal, fDiv)
	if (negV && !negD) || (!negV && negD) {
		fMod = -fMod
	}
	return T(fMod)
}

func Pow[T Real](val, exp T) T {
	fVal, fExp := float64(val), float64(exp)
	fPow := math.Pow(fVal, fExp)
	return T(fPow)
}

func Root[T Real](val, root T) T {
	return Pow(val, 1/root)
}

func Log[T Real](base, val T) T {
	fVal, fBase := float64(val), float64(base)
	fLog := math.Log(fVal) / math.Log(fBase)
	return T(fLog)
}

func Cos[T Real](radians T) T {
	fVal := float64(radians)
	fCos := math.Cos(fVal)
	return T(fCos)
}

func Sin[T Real](radians T) T {
	fVal := float64(radians)
	fSin := math.Sin(fVal)
	return T(fSin)
}

func Tan[T Real](radians T) T {
	fVal := float64(radians)
	fTan := math.Tan(fVal)
	return T(fTan)
}

func ACos[T Real](cos T) T {
	fVal := float64(cos)
	fRad := math.Acos(fVal)
	return T(fRad)
}

func ASin[T Real](sin T) T {
	fVal := float64(sin)
	fRad := math.Asin(fVal)
	return T(fRad)
}

func ATan[T Real](tan T) T {
	fVal := float64(tan)
	fRad := math.Atan(fVal)
	return T(fRad)
}

func CosDeg[T Real](degrees T) T {
	fVal := float64(degrees) * DEG2RAD
	fCos := math.Cos(fVal)
	return T(fCos)
}

func SinDeg[T Real](degrees T) T {
	fVal := float64(degrees) * DEG2RAD
	fSin := math.Sin(fVal)
	return T(fSin)
}

func TanDeg[T Real](degrees T) T {
	fVal := float64(degrees) * DEG2RAD
	fTan := math.Tan(fVal)
	return T(fTan)
}

func ACosDeg[T Real](cos T) T {
	fVal := float64(cos)
	fRad := math.Acos(fVal)
	return T(fRad * RAD2DEG)
}

func ASinDeg[T Real](sin T) T {
	fVal := float64(sin)
	fRad := math.Asin(fVal)
	return T(fRad * RAD2DEG)
}

func ATanDeg[T Real](tan T) T {
	fVal := float64(tan)
	fRad := math.Atan(fVal)
	return T(fRad * RAD2DEG)
}

func Sign[T Real](val T) T {
	if val < 0 {
		return -T(1)
	}
	return T(1)
}

func Ciel[T Real](val T) T {
	fVal := float64(val)
	fCiel := math.Ceil(fVal)
	return T(fCiel)
}

func Floor[T Real](val T) T {
	fVal := float64(val)
	fFloor := math.Floor(fVal)
	return T(fFloor)
}

func Round[T Real](val T) T {
	fVal := float64(val)
	fRound := math.Round(fVal)
	return T(fRound)
}

func ZeroOrVal[C Bool, T Real](condition C, val T) T {
	if condition {
		return val
	}
	return T(0)
}

func Lerp[T Real](start T, end T, amount float64) T {
	diff := end - start
	lerp := float64(diff) * amount
	return start + T(lerp)
}

func Range[T Real](start T, end T, val T) float64 {
	diff := end - start
	lerp := val - start
	percent := float64(lerp) / float64(diff)
	return percent
}
