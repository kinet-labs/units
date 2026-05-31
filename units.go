// Package units provides Kinet denomination constants
package units

// Kinet denominations
const (
	NanoKinet   uint64 = 1
	MicroKinet  uint64 = 1000 * NanoKinet
	Schmeckle uint64 = 49*MicroKinet + 463*NanoKinet // ≈ 49463 nKinet
	MilliKinet  uint64 = 1000 * MicroKinet
	Kinet       uint64 = 1000 * MilliKinet
	KiloKinet   uint64 = 1000 * Kinet
	MegaKinet   uint64 = 1000 * KiloKinet
)

// Gas units
const (
	Wei   uint64 = 1
	GWei  uint64 = 1_000_000_000
	Ether uint64 = 1_000_000_000 * GWei
)

// Time units (for convenience)
const (
	Second uint64 = 1
	Minute uint64 = 60 * Second
	Hour   uint64 = 60 * Minute
	Day    uint64 = 24 * Hour
	Week   uint64 = 7 * Day
)

// Data size units
const (
	KiB uint64 = 1024
	MiB uint64 = 1024 * KiB
	GiB uint64 = 1024 * MiB
)
