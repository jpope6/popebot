package engine

// Precalcate all possible attack sets on all squares,
// for all possible sets of blockers (on the appropriate
// rank/file or diagonals).
// Store the attack sets in a hash table and look them up as needed.

var bishopMagics [64]uint64 = [64]uint64{
	54045411736421449,
	3467949876944314368,
	1170054334914564,
	9386668213868429344,
	2612653001713909776,
	143074487959556,
	144405463509778640,
	2314868913337286662,
	1161933171206914308,
	5102013192899726,
	2306410683698270464,
	2343024120192851072,
	75437501641129984,
	8942543438848,
	576498144559498370,
	9313461918075293712,
	18041957234770120,
	9007906867522064,
	18577417326039296,
	72198365919854592,
	291045211837106176,
	4728814801970530308,
	70385941350476,
	72199439661474050,
	1227301284920249602,
	4904472908247872790,
	149533851001120,
	576746625460996128,
	4686136700865880064,
	92330389433092353,
	23648296635597120,
	1232552543291392,
	2348574017462305,
	74311661594870048,
	13835098222891107329,
	2201179652224,
	1152992149307130112,
	144397766860079688,
	1299994464869172242,
	2306124761249416192,
	23089753846974488,
	7512568235569455104,
	18579548627011656,
	1126038738339843,
	672935777931282,
	2323875584167379009,
	51793734359057408,
	2594777699881387520,
	1139147805032960,
	4683780038219268372,
	1161232523493384,
	72706313822208,
	653022014721892608,
	10408980038721605632,
	18070782873796608,
	4972541478955057216,
	10151825236041728,
	8521389665689600,
	14992505291550049312,
	290271107551376,
	2594076822951969037,
	270224911207891468,
	22817078453010948,
	153157608243496480,
}

var rookMagics [64]uint64 = [64]uint64{
	2918332697048780928,
	54044432483225602,
	3206580529027613000,
	324264121107222528,
	432358827626201096,
	144119654858752008,
	504403793954218240,
	108086684901442816,
	10555321290359856,
	73253863225839626,
	9369879830954247040,
	5188287542587297793,
	289637785395269889,
	4614500785407920208,
	1442277815999267236,
	291045144457675786,
	306262916611457058,
	1170938652432286848,
	1225261674207911936,
	576743327060205580,
	2252349703979136,
	563499751179264,
	1152922608413704704,
	13512997914017857,
	351884523110496,
	4611756390394892289,
	10394325534313029632,
	9313461623736238208,
	5226990463546036232,
	567350155805184,
	284240936185921,
	2270500101333060,
	36028934730547264,
	1193454038830563330,
	140875002810368,
	72092847146274048,
	4512397909820416,
	11241547688623212548,
	2269426426610177,
	4512949200749569,
	1765420262378029056,
	650911420596224,
	27022972707602450,
	2252907916300288,
	18295907847045126,
	10133116375040128,
	577639433099280520,
	9511638697972989988,
	6971572361690611840,
	58547070035820608,
	141149814653568,
	13511907521595648,
	9147953923489920,
	1407684159209728,
	4621396913713905792,
	11711048431922381312,
	4611721481982148609,
	3458928342135415298,
	35184510506241,
	9368086462124261389,
	288511919881977861,
	1873779031699105833,
	5764643842386034820,
	862571370643778,
}

type MagicEntry struct {
	Mask  Bitboard
	Magic uint64
	index uint8
}

func getMagicIndex(entry *MagicEntry, blockers Bitboard) {
	blockers &= entry.Mask
	// var key uint64 = (uint64(blockers) * bishopMagics[])
}
