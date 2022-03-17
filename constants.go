package main

// techLevels - lists the tech levels supported by the app
var (
	techLevels = []string{"A", "B", "C", "D", "E", "F", "G"}

	noneString = "0"

	// tons - lists the ship tonnage amounts supported by the app
	tons = []string{
		"95", "100", "200", "300", "400", "500", "600", "700", "800", "900", "1000",
		"1100", "1200", "1200", "1300", "1400", "1500", "1600", "1700", "1800", "1900", "2000",
		"4000", "6000", "8000", "10000", "20000", "40000", "60000", "80000", "100000",
		"250000", "500000", "1000000",
	}
)

type hullSelection struct {
	code  string
	tons  int
	price int
}

var hullSelectionCode = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C",
	"D", "E", "F", "G", "H", // Skipping I
	"J", "K", "L", "M", "N", // Skipping O
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func getIndexFromHull(dString string) int {
	for resultInt, dMatch := range hullSelectionCode {
		if dMatch == dString {
			return resultInt
		}
	}
	return -1
}

var hullSelections = []hullSelection{
	{hullSelectionCode[0], 95, 2},           // 1
	{hullSelectionCode[1], 100, 8},          // 2
	{hullSelectionCode[2], 200, 12},         // 3
	{hullSelectionCode[3], 300, 16},         // 4
	{hullSelectionCode[4], 400, 32},         // 5
	{hullSelectionCode[5], 500, 48},         // 6
	{hullSelectionCode[6], 600, 64},         // 7
	{hullSelectionCode[7], 700, 80},         // 8
	{hullSelectionCode[8], 800, 90},         // 9
	{hullSelectionCode[9], 900, 100},        // A
	{hullSelectionCode[10], 1000, 120},      // B
	{hullSelectionCode[11], 1100, 120},      // C
	{hullSelectionCode[12], 1200, 140},      // D
	{hullSelectionCode[13], 1300, 160},      // E
	{hullSelectionCode[14], 1400, 180},      // F
	{hullSelectionCode[15], 1500, 200},      // G
	{hullSelectionCode[16], 1600, 220},      // H
	{hullSelectionCode[17], 1700, 240},      // J
	{hullSelectionCode[18], 1800, 260},      // K
	{hullSelectionCode[19], 1900, 300},      // L
	{hullSelectionCode[20], 2000, 300},      // M
	{hullSelectionCode[21], 4000, 750},      // N
	{hullSelectionCode[22], 6000, 1100},     // P
	{hullSelectionCode[23], 8000, 1400},     // Q
	{hullSelectionCode[24], 10000, 1750},    // R
	{hullSelectionCode[25], 20000, 4250},    // S
	{hullSelectionCode[26], 40000, 10200},   // T
	{hullSelectionCode[27], 60000, 14400},   // U
	{hullSelectionCode[28], 80000, 18000},   // V
	{hullSelectionCode[29], 100000, 21600},  // W
	{hullSelectionCode[30], 250000, 65000},  // X
	{hullSelectionCode[31], 500000, 150000}, // Y
	{hullSelectionCode[32], 100000, 375000}, // Z
}

var hullSelectionCfgs = []string{"streamlined", "distributed", "standard"}

type hullSelectionConfiguration struct {
	cfg    string
	cost   float32
	scoops bool
}

var hullSelectionConfigurations = []hullSelectionConfiguration{
	{hullSelectionCfgs[0], 1.1, true},
	{hullSelectionCfgs[1], 0.9, false},
	{hullSelectionCfgs[2], 1.0, true},
}

type offering struct {
	offer     string
	tlMin     int
	costPerMC float32
}

var hullSelectionOfferings = []offering{
	{"Reflec", 10, 0.1},
	{"Self-Sealing", 9, 0.01},
	{"Stealth", 11, 0.1},
}

type armor struct {
	tlMin  int
	factor float32
}

var (
	armors = []armor{
		{7, 4.0},
		{10, 3.0},
		{12, 2.0},
		{14, 1.0},
		{16, 0.6},
	}
	weaponLevel = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31",
		"32", "33", "34", "35", "36", "37", "38", "39", "40",
	}
	vehicleCount = []string{
		"0", "1", "2", "3", "4",
	}
	computer = [16]int{
		1, 2, 3, 4, 5, 7, 9, 11, 13, 15, 17, 20, 23, 26, 30, 35,
	}
)

var TrvIndex = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

const (
	defaultDriveCode  = "B"
	defaultHullCode   = "2"
	defaultTons       = 200
	defaultMaxhardpts = 2
	defaultHullPrice  = 8
	roundUp           = 0.999999

	// Vehicles
	vehicleTypes      = "Vehicles, Boats & Fighters"
	wATVType          = "Wheeled ATV"
	tATVType          = "Tracked ATV"
	airRaftType       = "Air Raft"
	speederType       = "Speeder"
	gCarrierType      = "GCarrier"
	shipsLaunchType   = "Ship's Launch"
	shipsBoatType     = "Ship's Boat"
	pinnaceType       = "Pinnace"
	cutterType        = "Modular Cutter"
	slowBoatType      = "Slow Boat"
	slowPinnaceType   = "Slow Pinnace"
	shuttleType       = "Shuttle"
	lightFighterType  = "Light Fighter"
	mediumFighterType = "Medium Fighter"
	heavyFighterType  = "Heavy Fighter"
)
