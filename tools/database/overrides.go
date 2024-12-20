package database

import (
	"regexp"

	"github.com/wowsims/cata/sim/core/proto"
	"github.com/wowsims/cata/sim/core/stats"
)

var OtherItemIdsToFetch = []string{
	// Hallow's End Ilvl bumped rings
	"211817",
	"211844",
	"211847",
	"211850",
	"211851",
}

var ItemOverrides = []*proto.UIItem{
	// TODO: Remove these when blizzard fixes the item data
	{ /** Alysra's Razor (H) */ Id: 71427, WeaponDamageMin: 610, WeaponDamageMax: 1134, WeaponSpeed: 1.4},
	{ /** Arathar, the Eye of Flame (H) */ Id: 71611, WeaponDamageMin: 1865, WeaponDamageMax: 3465, WeaponSpeed: 3.0},
	{ /** Arbalest of Erupting Fury (H) */ Id: 71414, WeaponDamageMin: 2015, WeaponDamageMax: 3024, WeaponSpeed: 3.0},
	{ /** Avool's Incendiary Shanker */ Id: 71779, WeaponDamageMin: 695, WeaponDamageMax: 1291, WeaponSpeed: 1.8},
	{ /** Avool's Incendiary Shanker (H) */ Id: 71778, WeaponDamageMin: 784, WeaponDamageMax: 1458, WeaponSpeed: 1.8},
	{ /** Deflecting Star (H) */ Id: 71592, WeaponDamageMin: 1410, WeaponDamageMax: 2116, WeaponSpeed: 2.1},
	{ /** Entrail Disgorger */ Id: 71787, WeaponDamageMin: 618, WeaponDamageMax: 927, WeaponSpeed: 1.4},
	{ /** Entrail Disgorger (H) */ Id: 71786, WeaponDamageMin: 697, WeaponDamageMax: 1046, WeaponSpeed: 1.4},
	{ /** Eye of Purification */ Id: 71776, WeaponDamageMin: 463, WeaponDamageMax: 861, WeaponSpeed: 2.4},
	{ /** Eye of Purification (H) */ Id: 71777, WeaponDamageMin: 523, WeaponDamageMax: 972, WeaponSpeed: 2.4},
	{ /** Feeding Frenzy (H) */ Id: 71441, WeaponDamageMin: 896, WeaponDamageMax: 1345, WeaponSpeed: 1.8},
	{ /** Funeral Pyre (H) */ Id: 71409, WeaponDamageMin: 597, WeaponDamageMax: 897, WeaponSpeed: 2.4},
	{ /** Gatecrasher (H) */ Id: 71454, WeaponDamageMin: 1133, WeaponDamageMax: 2105, WeaponSpeed: 2.6},
	{ /** Giantslicer (H) */ Id: 71593, WeaponDamageMin: 1410, WeaponDamageMax: 2116, WeaponSpeed: 2.1},
	{ /** Ko'gun, Hammer of the Firelord (H) */ Id: 71615, WeaponDamageMin: 484, WeaponDamageMax: 899, WeaponSpeed: 2.1},
	{ /** Morningstar Shard (H) */ Id: 71568, WeaponDamageMin: 1343, WeaponDamageMax: 2016, WeaponSpeed: 2.0},
	{ /** Sho'ravon, Greatstaff of Annihilation */ Id: 71798, WeaponDamageMin: 560, WeaponDamageMax: 840, WeaponSpeed: 2.4},
	{ /** Skullstealer Greataxe (H) */ Id: 71445, WeaponDamageMin: 2418, WeaponDamageMax: 3628, WeaponSpeed: 3.6},
	{ /** Smoldering Censer of Purity */ Id: 71775, WeaponDamageMin: 640, WeaponDamageMax: 960, WeaponSpeed: 2.9},
	{ /** Smoldering Censer of Purity (H) */ Id: 71774, WeaponDamageMin: 722, WeaponDamageMax: 1084, WeaponSpeed: 2.9},
	{ /** Volcanospike (H) */ Id: 71422, WeaponDamageMin: 348, WeaponDamageMax: 648, WeaponSpeed: 1.6},
	{ /** Zoid's Firelit Greatsword */ Id: 71780, WeaponDamageMin: 2142, WeaponDamageMax: 3214, WeaponSpeed: 3.6},
	{ /** Zoid's Firelit Greatsword (H) */ Id: 71781, WeaponDamageMin: 2418, WeaponDamageMax: 3628, WeaponSpeed: 3.6},

	{ /** Raz's Pauldrons */ Id: 56318, SocketBonus: stats.Stats{stats.HitRating: 20}.ToProtoArray()},

	{ /** Destruction Holo-gogs */ Id: 32494, ClassAllowlist: []proto.Class{proto.Class_ClassMage, proto.Class_ClassPriest, proto.Class_ClassWarlock}},
	{ /** Gadgetstorm Goggles */ Id: 32476, ClassAllowlist: []proto.Class{proto.Class_ClassShaman}},
	{ /** Magnified Moon Specs */ Id: 32480, ClassAllowlist: []proto.Class{proto.Class_ClassDruid}},
	{ /** Quad Deathblow X44 Goggles */ Id: 34353, ClassAllowlist: []proto.Class{proto.Class_ClassDruid, proto.Class_ClassRogue}},
	{ /** Hyper-Magnified Moon Specs */ Id: 35182, ClassAllowlist: []proto.Class{proto.Class_ClassDruid}},
	{ /** Lightning Etched Specs */ Id: 34355, ClassAllowlist: []proto.Class{proto.Class_ClassShaman}},
	{ /** Annihilator Holo-Gogs */ Id: 34847, ClassAllowlist: []proto.Class{proto.Class_ClassMage, proto.Class_ClassPriest, proto.Class_ClassWarlock}},

	// Balance T9 "of Conquest" Alliance set
	{Id: 48158, SetName: "Malfurion's Regalia"},
	{Id: 48159, SetName: "Malfurion's Regalia"},
	{Id: 48160, SetName: "Malfurion's Regalia"},
	{Id: 48161, SetName: "Malfurion's Regalia"},
	{Id: 48162, SetName: "Malfurion's Regalia"},

	// Death Knight T9 "of Conquest" Horde set
	{Id: 48501, SetName: "Koltira's Battlegear"},
	{Id: 48502, SetName: "Koltira's Battlegear"},
	{Id: 48503, SetName: "Koltira's Battlegear"},
	{Id: 48504, SetName: "Koltira's Battlegear"},
	{Id: 48505, SetName: "Koltira's Battlegear"},

	// Death Knight T9 "of Conquest" Tank Horde set
	{Id: 48558, SetName: "Koltira's Plate"},
	{Id: 48559, SetName: "Koltira's Plate"},
	{Id: 48560, SetName: "Koltira's Plate"},
	{Id: 48561, SetName: "Koltira's Plate"},
	{Id: 48562, SetName: "Koltira's Plate"},

	// Valorous T8 Sets
	{Id: 45375, Phase: 2},
	{Id: 45381, Phase: 2},
	{Id: 45382, Phase: 2},
	{Id: 45376, Phase: 2},
	{Id: 45370, Phase: 2},
	{Id: 45371, Phase: 2},
	{Id: 45383, Phase: 2},
	{Id: 45372, Phase: 2},
	{Id: 45377, Phase: 2},
	{Id: 45384, Phase: 2},
	{Id: 45379, Phase: 2},
	{Id: 45385, Phase: 2},
	{Id: 45380, Phase: 2},
	{Id: 45373, Phase: 2},
	{Id: 45374, Phase: 2},
	{Id: 45391, Phase: 2},
	{Id: 45386, Phase: 2},
	{Id: 45340, Phase: 2},
	{Id: 45335, Phase: 2},
	{Id: 45336, Phase: 2},
	{Id: 45341, Phase: 2},
	{Id: 45337, Phase: 2},
	{Id: 45342, Phase: 2},
	{Id: 45338, Phase: 2},
	{Id: 45343, Phase: 2},
	{Id: 45339, Phase: 2},
	{Id: 45344, Phase: 2},
	{Id: 45419, Phase: 2},
	{Id: 45417, Phase: 2},
	{Id: 45420, Phase: 2},
	{Id: 45421, Phase: 2},
	{Id: 45422, Phase: 2},
	{Id: 45387, Phase: 2},
	{Id: 45392, Phase: 2},
	{Id: 46131, Phase: 2},
	{Id: 45365, Phase: 2},
	{Id: 45367, Phase: 2},
	{Id: 45369, Phase: 2},
	{Id: 45368, Phase: 2},
	{Id: 45388, Phase: 2},
	{Id: 45393, Phase: 2},
	{Id: 46313, Phase: 2},
	{Id: 45351, Phase: 2},
	{Id: 45355, Phase: 2},
	{Id: 45345, Phase: 2},
	{Id: 45356, Phase: 2},
	{Id: 45346, Phase: 2},
	{Id: 45347, Phase: 2},
	{Id: 45357, Phase: 2},
	{Id: 45352, Phase: 2},
	{Id: 45358, Phase: 2},
	{Id: 45348, Phase: 2},
	{Id: 45359, Phase: 2},
	{Id: 45349, Phase: 2},
	{Id: 45353, Phase: 2},
	{Id: 45354, Phase: 2},
	{Id: 45394, Phase: 2},
	{Id: 45395, Phase: 2},
	{Id: 45389, Phase: 2},
	{Id: 45360, Phase: 2},
	{Id: 45361, Phase: 2},
	{Id: 45362, Phase: 2},
	{Id: 45363, Phase: 2},
	{Id: 45364, Phase: 2},
	{Id: 45390, Phase: 2},
	{Id: 45429, Phase: 2},
	{Id: 45424, Phase: 2},
	{Id: 45430, Phase: 2},
	{Id: 45425, Phase: 2},
	{Id: 45426, Phase: 2},
	{Id: 45431, Phase: 2},
	{Id: 45427, Phase: 2},
	{Id: 45432, Phase: 2},
	{Id: 45428, Phase: 2},
	{Id: 45433, Phase: 2},
	{Id: 45396, Phase: 2},
	{Id: 45397, Phase: 2},
	{Id: 45398, Phase: 2},
	{Id: 45399, Phase: 2},
	{Id: 45400, Phase: 2},
	{Id: 45413, Phase: 2},
	{Id: 45412, Phase: 2},
	{Id: 45406, Phase: 2},
	{Id: 45414, Phase: 2},
	{Id: 45401, Phase: 2},
	{Id: 45411, Phase: 2},
	{Id: 45402, Phase: 2},
	{Id: 45408, Phase: 2},
	{Id: 45409, Phase: 2},
	{Id: 45403, Phase: 2},
	{Id: 45415, Phase: 2},
	{Id: 45410, Phase: 2},
	{Id: 45404, Phase: 2},
	{Id: 45405, Phase: 2},
	{Id: 45416, Phase: 2},

	// Other items Wowhead has the wrong phase listed for
	// Ick's loot table from Pit of Saron
	{Id: 49812, Phase: 4},
	{Id: 49808, Phase: 4},
	{Id: 49811, Phase: 4},
	{Id: 49807, Phase: 4},
	{Id: 49810, Phase: 4},
	{Id: 49809, Phase: 4},

	// Drape of Icy Intent
	{Id: 45461, Phase: 2},

	// Valentine's day event rewards
	{Id: 51804, Phase: 2},
	{Id: 51805, Phase: 2},
	{Id: 51806, Phase: 2},
	{Id: 51807, Phase: 2},
	{Id: 51808, Phase: 2},

	// Cata pre-patch event items
	{Id: 53492, Phase: 5},

	// Heirloom Dwarven Handcannon, Wowhead partially glitchs out and shows us some other lvl calc for this
	{Id: 44093, Stats: stats.Stats{stats.CritRating: 30, stats.ResilienceRating: 13, stats.AttackPower: 34}.ToProtoArray()},
}

// Keep these sorted by item ID.
var ItemAllowList = map[int32]struct{}{
	//Shaman Dungeon Set 3 Tidefury
	27510: {}, // Tidefury Gauntlets
	27802: {}, // Tidefury Shoulderguards
	27909: {}, // Tidefury Kilt
	28231: {}, // Tidefury Chestpiece
	28349: {}, // Tidefury Helm

	45703: {}, // Spark of Hope
}

// Keep these sorted by item ID.
var ItemDenyList = map[int32]struct{}{
	17782: {}, // talisman of the binding shard
	17783: {}, // talisman of the binding fragment
	17802: {}, // Deprecated version of Thunderfury
	18582: {},
	18583: {},
	18584: {},
	24265: {},
	32384: {},
	32421: {},
	32422: {},
	33482: {},
	33350: {},
	34576: {}, // Battlemaster's Cruelty
	34577: {}, // Battlemaster's Depreavity
	34578: {}, // Battlemaster's Determination
	34579: {}, // Battlemaster's Audacity
	34580: {}, // Battlemaster's Perseverence

	38694: {}, // "Family" Shoulderpads heirloom
	45084: {}, // 'Book of Crafting Secrets' heirloom

	// '10 man' onyxia head rewards
	49312: {},
	49313: {},
	49314: {},

	50251: {}, // 'one hand shadows edge'
	53500: {}, // Tectonic Plate

	48880: {}, // DK's Tier 9 Duplicates
	48881: {}, // DK's Tier 9 Duplicates
	48882: {}, // DK's Tier 9 Duplicates
	48883: {}, // DK's Tier 9 Duplicates
	48884: {}, // DK's Tier 9 Duplicates
	48885: {}, // DK's Tier 9 Duplicates
	48886: {}, // DK's Tier 9 Duplicates
	48887: {}, // DK's Tier 9 Duplicates
	48888: {}, // DK's Tier 9 Duplicates
	48889: {}, // DK's Tier 9 Duplicates
	48890: {}, // DK's Tier 9 Duplicates
	48891: {}, // DK's Tier 9 Duplicates
	48892: {}, // DK's Tier 9 Duplicates
	48893: {}, // DK's Tier 9 Duplicates
	48894: {}, // DK's Tier 9 Duplicates
	48895: {}, // DK's Tier 9 Duplicates
	48896: {}, // DK's Tier 9 Duplicates
	48897: {}, // DK's Tier 9 Duplicates
	48898: {}, // DK's Tier 9 Duplicates
	48899: {}, // DK's Tier 9 Duplicates
	68710: {}, // Stump of Time Duplicate (Not available ingame)
	68711: {}, // Mandala of Stirring Patterns Duplicate
	68712: {}, // Impatience of Youth Duplicate
	68713: {}, // Mirror of Broken Images Duplicate
	65104: {}, // DONTUSEUnheeded Warning
	65015: {}, // DONTUSEFury of Angerforge

	232548: {}, // The Horseman's Sinister Saber - 353
	232544: {}, // The Horseman's Horrific Helmet - 353
	232536: {}, // Band of Ghoulish Glee - 353
	232537: {}, // The Horseman's Signet - 353
	232540: {}, // Wicked Witch's Ring - 353
	232538: {}, // Seal of the Petrified Pumpkin - 353

	71331: {}, // Direbrew's Bloodied Shanker - 365
	71332: {}, // Tremendous Tankard O' Terror - 365
	71333: {}, // Bubblier Brightbrew Charm - 365
	71334: {}, // Bitterer Balebrew Charm - 365
	71335: {}, // Coren's Chilled Chromium Coaster - 365
	71336: {}, // Petrified Pickled Egg - 365
	71337: {}, // Mithril Stopwatch - 365
	71338: {}, // Brawler's Trophy - 365

	// T11 BoE items which are not available in the game
	65005: {}, // Claws of Agony - 372
	65006: {}, // Claws of Torment - 372
	65008: {}, // Shadowforge's Lightbound Smock - 372
	65009: {}, // Hide of Chromaggus - 372
	65010: {}, // Ironstar's Impenetrable Cover - 372
	65011: {}, // Corehammer's Riveted Girdle - 372
	65012: {}, // Treads of Savage Beatings - 372
	65013: {}, // Maldo's Sword Cane - 372
	65014: {}, // Maimgor's Bite - 372
	65016: {}, // Theresa's Booklight - 372
	65097: {}, // Bracers of the Dark Pool - 372
	65098: {}, // Crossfire Carbine - 372
	65099: {}, // Tsanga's Helm - 372
	65100: {}, // Phase-Twister Leggings - 372
	65101: {}, // Heaving Plates of Protection - 372
	65102: {}, // Chelley's Staff of Dark Mending - 372
	65103: {}, // Soul Blade - 372

	// Firelands "upgraded" items which are not available in the game
	69184: {}, // Stay of Execution - 391
	69185: {}, // Rune of Zeth - 391
	69198: {}, // Fiery Quintessence - 391
	69199: {}, // Ancient Petrified Seed - 391
	69200: {}, // Essence of the Eternal Flame - 391
	71388: {}, // Sleek Flamewrath Cloak - 391
	71389: {}, // Rippling Flamewrath Drape - 391
	71390: {}, // Flowing Flamewrath Cape - 391
	71391: {}, // Bladed Flamewrath Cover - 391
	71392: {}, // Durable Flamewrath Greatcloak - 391
	71393: {}, // Embereye Belt - 391
	71394: {}, // Flamebinding Girdle - 391
	71395: {}, // Firescar Sash - 391
	71396: {}, // Firearrow Belt - 391
	71397: {}, // Firemend Cinch - 391
	71398: {}, // Belt of the Seven Seals - 391
	71399: {}, // Cinch of the Flaming Ember - 391
	71400: {}, // Girdle of the Indomitable Flame - 391
	71565: {}, // Necklace of Smoke Signals - 391
	71566: {}, // Splintered Brimstone Seal - 391
	71569: {}, // Flamebinder Bracers, - 391
	71570: {}, // Bracers of Forked Lightning, - 391
	71571: {}, // Emberflame Bracers, - 391
	71572: {}, // Firesoul Wristguards, - 391
	71573: {}, // Amulet of Burning Brilliance - 391
	71574: {}, // Crystalline Brimstone Ring - 391
	71576: {}, // Firemind Pendant - 391
	71578: {}, // Soothing Brimstone Circle - 391
	71581: {}, // Smolderskull Bindings, - 391
	71582: {}, // Bracers of Misting Ash, - 391
	71583: {}, // Bracers of Imperious Truths - 391
	71584: {}, // Gigantiform Bracers, - 391
	71585: {}, // Bracers of Regal Force, - 391
	71586: {}, // Stoneheart Choker - 391
	71588: {}, // Serrated Brimstone Signet - 391
	71589: {}, // Stoneheart Necklace - 391
	71591: {}, // Deflecting Brimstone Band - 391
}

// Item icons to include in the DB, so they don't need to be separately loaded in the UI.
var ExtraItemIcons = []int32{
	// Pet foods
	33874,
	43005,

	// Demonic Rune
	12662,

	// Food IDs
	27655,
	27657,
	27658,
	27664,
	33052,
	33825,
	33872,
	34753,
	34754,
	34756,
	34758,
	34767,
	34769,
	42994,
	42995,
	42996,
	42998,
	42999,
	43000,
	43015,
	62290,
	62649,
	62671,
	62670,
	62661,
	62665,
	62668,
	62669,
	62664,
	62666,
	62667,
	62662,
	62663,

	// Flask IDs
	13512,
	22851,
	22853,
	22854,
	22861,
	22866,
	33208,
	40079,
	44939,
	46376,
	46377,
	46378,
	46379,

	// Elixer IDs
	40072,
	40078,
	40097,
	40109,
	44328,
	44332,

	// Elixer IDs
	13452,
	13454,
	22824,
	22827,
	22831,
	22833,
	22834,
	22835,
	22840,
	28103,
	28104,
	31679,
	32062,
	32067,
	32068,
	39666,
	40068,
	40070,
	40073,
	40076,
	44325,
	44327,
	44329,
	44330,
	44331,
	9088,
	9224,

	// Potions / In Battle Consumes
	13442,
	20520,
	22105,
	22788,
	22828,
	22832,
	22837,
	22838,
	22839,
	22849,
	31677,
	33447,
	33448,
	36892,
	40093,
	40211,
	40212,
	40536,
	40771,
	41119,
	41166,
	42545,
	42641,

	// Poisons
	43231,
	43233,
	43235,

	// Thistle Tea
	7676,

	// Scrolls
	37094,
	43466,
	43464,
	37092,
	37098,
	43468,

	// Drums
	49633,
	49634,
}

// Raid buffs / debuffs
var SharedSpellsIcons = []int32{
	// Revitalize, Rejuv, WG
	48545,
	26982,
	53251,

	// Registered CD's
	49016,
	57933,
	64382,
	10060,
	16190,
	29166,
	53530,
	33206,
	2825,
	54758,

	// Raid Buffs
	43002,
	57567,
	54038,

	48470,
	17051,

	25898,
	25899,

	48942,
	20140,
	8071,
	16293,

	48161,
	14767,

	8075,
	52456,
	57623,

	48073,

	48934,
	20045,
	47436,

	53138,
	30808,
	19506,

	31869,
	31583,
	34460,

	57472,
	50720,

	53648,

	47440,
	12861,
	47982,
	18696,

	48938,
	20245,
	5675,
	16206,

	17007,
	34300,
	29801,

	55610,
	8512,
	29193,

	48160,
	31878,
	53292,
	54118,
	44561,

	24907,
	48396,
	51470,

	3738,
	47240,
	57722,
	8227,

	54043,
	48170,
	31025,
	31035,
	6562,
	31033,
	53307,
	16840,
	54648,

	// Raid Debuffs
	8647,
	47467,
	55749,

	770,
	33602,
	702,
	18180,
	56631,
	53598,

	26016,
	47437,
	12879,
	48560,
	16862,
	55487,

	48566,
	46855,
	57386,

	30706,
	20337,
	58410,

	47502,
	12666,
	55095,
	51456,
	53696,
	48485,

	3043,
	29859,
	58413,
	65855,

	17800,
	17803,
	12873,
	28593,

	33198,
	51161,
	48511,
	1490,

	20271,
	53408,

	11374,
	15235,

	27013,

	58749,
	49071,

	30708,
}

// If any of these match the item name, don't include it.
var DenyListNameRegexes = []*regexp.Regexp{
	regexp.MustCompile(`30 Epic`),
	regexp.MustCompile(`130 Epic`),
	regexp.MustCompile(`63 Blue`),
	regexp.MustCompile(`63 Green`),
	regexp.MustCompile(`66 Epic`),
	regexp.MustCompile(`90 Epic`),
	regexp.MustCompile(`90 Green`),
	regexp.MustCompile(`Boots 1`),
	regexp.MustCompile(`Boots 2`),
	regexp.MustCompile(`Boots 3`),
	regexp.MustCompile(`Bracer 1`),
	regexp.MustCompile(`Bracer 2`),
	regexp.MustCompile(`Bracer 3`),
	regexp.MustCompile(`DB\d`),
	regexp.MustCompile(`DEPRECATED`),
	regexp.MustCompile(`Deprecated: Keanna`),
	regexp.MustCompile(`Indalamar`),
	regexp.MustCompile(`Monster -`),
	regexp.MustCompile(`NEW`),
	regexp.MustCompile(`PH`),
	regexp.MustCompile(`QR XXXX`),
	regexp.MustCompile(`TEST`),
	regexp.MustCompile(`Test`),
	regexp.MustCompile(`zOLD`),
	regexp.MustCompile(`Archaic Spell`),
	regexp.MustCompile(`Well Repaired`),
}

// Allows manual overriding for Gem fields in case WowHead is wrong.
var GemOverrides = []*proto.UIGem{
	{Id: 33131, Stats: stats.Stats{stats.AttackPower: 32, stats.RangedAttackPower: 32}.ToProtoArray()},
}
var GemAllowList = map[int32]struct{}{
	//22459: {}, // Void Sphere
	//36766: {}, // Bright Dragon's Eye
	//36767: {}, // Solid Dragon's Eye
}
var GemDenyList = map[int32]struct{}{
	// pvp non-unique gems not in game currently.
	32735: {},
	34142: {}, // Infinite Sphere
	34143: {}, // Chromatic Sphere
	35489: {},
	37430: {}, // Solid Sky Sapphire (Unused)
	38545: {},
	38546: {},
	38547: {},
	38548: {},
	38549: {},
	38550: {},
	63696: {},
	63697: {},
}
