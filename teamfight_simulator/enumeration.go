package main

type ItemName int

const (
	BFSword ItemName = iota
	ChainVest
	FryingPan
	GiantsBelt
	NeedlesslyLargeRod
	NegatronCloak
	RecurveBow
	SparringGloves
	Spatula
	TearOfTheGoddess

	AdaptiveHelm
	ArchangelsStaff
	Bloodthirster
	BlueBuff
	BrambleVest
	Crownguard
	Deathblade
	DragonsClaw
	EdgeOfNight
	Evenshroud
	GargoyleStoneplate
	GiantSlayer
	Guardbreaker
	GuinsoosRageblade
	HandsOfJustice
	HextechGunblade
	InfinityEdge
	IonicSpark
	JeweledGauntlet
	LastWhisper
	Morellonomicon
	NashorsTooth
	ProtectorsVow
	Quicksilver
	RabadonsDeathcap
	RedBuff
	Redemption
	RunaansHurricane
	SpearOfShojin
	StatikkShiv
	SteadfastHeart
	SteraksGage
	SunfireCape
	ThiefsGloves
	TitansResolve
	WarmogsArmor

	AmbusherEmblem
	ArtilleristEmblem
	AutomataEmblem
	BlackRoseEmblem
	BruiserEmblem
	ConquerorEmblem
	EnforcerEmblem
	ExperimentEmblem
	FamilyEmblem
	FirelightEmblem
	PitFighterEmblem
	QuickstrikerEmblem
	RebelEmblem
	SentinelEmblem
	SorcererEmblem
	VisionaryEmblem

	TactitiansCape
	TactitiansCrown
	TactitiansShield

	AccomplicesGloves
	AegisOfTheLegion
	BansheesWill
	ChaliceOfPower
	CrestOfCinders
	KnightsVow
	LocketOfTheIronSolari
	MoonstoneRenewer
	NeedlesslyBigGem
	ObsidianCleaver
	RanduinsOmen
	Spite
	TheEternalFlame
	UnstableTreasureChest
	VirtueOfTheMartyr
	ZekesHerald
	Zephyr
	ZzRotPortal

	AnimaVisage
	BlightingJewel
	DeathsDefiance
	DeathfireGrasp
	Fishbones
	GamblersBlade
	GoldCollector
	HorizonFocus
	HullCrusher
	InfinityForce
	InnervatingLocket
	LichBane
	LightshieldCrest
	LudensTempest
	Manazane
	Mittens
	MogulsMail
	ProwlersClaw
	RapidFirecannon
	SeekersArmguard
	SilvermereDawn
	SnipersFocus
	SpectralCutlass
	SuspiciousTrenchcoat
	TalismanOfAscension
	TrickstersGlass
	UnendingDespair
	WitsEnd
	ZhonyasParadox
)

func ItemToString(itemName ItemName) string {
	switch itemName {

	case BFSword:
		return "B.F. Sword"
	case ChainVest:
		return "Chain Vest"
	case FryingPan:
		return "Frying Pan"
	case GiantsBelt:
		return "Giant's Belt"
	case NeedlesslyLargeRod:
		return "Needlessly Large Rod"
	case NegatronCloak:
		return "Negatron Cloak"
	case RecurveBow:
		return "Recurve Bow"
	case SparringGloves:
		return "Sparring Gloves"
	case Spatula:
		return "Spatula"
	case TearOfTheGoddess:
		return "Tear of the Godess"
	case AdaptiveHelm:
		return "Adaptive Helm"
	case ArchangelsStaff:
		return "Archangel's Staff"
	case Bloodthirster:
		return "Bloodthirster"
	case BlueBuff:
		return "Blue Buff"
	case BrambleVest:
		return "Bramble Vest"
	case Crownguard:
		return "Crownguard"
	case Deathblade:
		return "Deathblade"
	case DragonsClaw:
		return "Dragon's Claw"
	case EdgeOfNight:
		return "Edge of Night"
	case Evenshroud:
		return "Evenshroud"
	case GargoyleStoneplate:
		return "Gargoyle Stoneplate"
	case GiantSlayer:
		return "Giant Slayer"
	case Guardbreaker:
		return "Guiardbreaker"
	case GuinsoosRageblade:
		return "Guinsoo's Rageblade"
	case HandsOfJustice:
		return "Hands of Justice"
	case HextechGunblade:
		return "Hextech Gunblade"
	case InfinityEdge:
		return "Infinity Edge"
	case IonicSpark:
		return "Ionic Spark"
	case JeweledGauntlet:
		return "Jeweled Gauntlet"
	case LastWhisper:
		return "Last Whisper"
	case Morellonomicon:
		return "Morellonomicon"
	case NashorsTooth:
		return "Nahsor's Tooth"
	case ProtectorsVow:
		return "Protector's Vow"
	case Quicksilver:
		return "Quicksilver"
	case RabadonsDeathcap:
		return "Rabadon's Deathcap"
	case RedBuff:
		return "Red Buff"
	case Redemption:
		return "Redemption"
	case RunaansHurricane:
		return "Runaan's Hurricane"
	case SpearOfShojin:
		return "Spear of Shojin"
	case StatikkShiv:
		return "Statikk Shiv"
	case SteadfastHeart:
		return "Steadfast Heart"
	case SteraksGage:
		return "Sterak's Gage"
	case SunfireCape:
		return "Sunfire Cape"
	case ThiefsGloves:
		return "Thief's Gloves"
	case TitansResolve:
		return "Titan's Resolve"
	case WarmogsArmor:
		return "Warmog's Armor"

	case AmbusherEmblem:
		return "Ambusher Emblem"
	case ArtilleristEmblem:
		return "Artillerist Emblem"
	case AutomataEmblem:
		return "Automata Emblem"
	case BlackRoseEmblem:
		return "Black Rose Emblem"
	case BruiserEmblem:
		return "Bruiser Emblem"
	case ConquerorEmblem:
		return "Conqueror Emblem"
	case EnforcerEmblem:
		return "Enforcer Emblem"
	case ExperimentEmblem:
		return "Experiment Emblem"
	case FamilyEmblem:
		return "Family Emblem"
	case FirelightEmblem:
		return "Firelight Emblem"
	case PitFighterEmblem:
		return "Pit Fighter Emblem"
	case QuickstrikerEmblem:
		return "Quickstriker Emblem"
	case RebelEmblem:
		return "Rebel Emblem"
	case SentinelEmblem:
		return "Sentinel Emblem"
	case SorcererEmblem:
		return "Sorcerer Emblem"
	case VisionaryEmblem:
		return "Visionary Emblem"

	case TactitiansCape:
		return "Tactitian's Cape"
	case TactitiansCrown:
		return "Tactician's Crown"
	case TactitiansShield:
		return "Tactician's Shield"

	case AccomplicesGloves:
		return "Accomplice's Gloves"
	case AegisOfTheLegion:
		return "Aegis of the Legion"
	case BansheesWill:
		return "Banshee's Will"
	case ChaliceOfPower:
		return "Chalice of Power"
	case CrestOfCinders:
		return "Crest of Cinders"
	case KnightsVow:
		return "Knights Vow"
	case LocketOfTheIronSolari:
		return "Locket of the Iron Solari"
	case MoonstoneRenewer:
		return "Moonstone Renewer"
	case NeedlesslyBigGem:
		return "Needlessly Big Gem"
	case ObsidianCleaver:
		return "Obsidian Cleaver"
	case RanduinsOmen:
		return "Randuin's Omen"
	case Spite:
		return "Sptire"
	case TheEternalFlame:
		return "The Eternal Flame"
	case UnstableTreasureChest:
		return "Unstable Treasure Chest"
	case VirtueOfTheMartyr:
		return "Virtue of the Martyr"
	case ZekesHerald:
		return "Zeke's Herald"
	case Zephyr:
		return "Zephyr"
	case ZzRotPortal:
		return "Zz'Rot Portal"

	case AnimaVisage:
		return "Anima Visage"
	case BlightingJewel:
		return "Blighting Jewel"
	case DeathsDefiance:
		return "Death's Defiance"
	case DeathfireGrasp:
		return "Deathfire Grasp"
	case Fishbones:
		return "Fishbones"
	case GamblersBlade:
		return "Gambler's Blade"
	case GoldCollector:
		return "Gold Collector"
	case HorizonFocus:
		return "Horizon Focus"
	case HullCrusher:
		return "Hull Crusher"
	case InfinityForce:
		return "Infinity Force"
	case InnervatingLocket:
		return "Innervating Locket"
	case LichBane:
		return "Lich Bane"
	case LightshieldCrest:
		return "Lightshield Crest"
	case LudensTempest:
		return "Luden's Tempest"
	case Manazane:
		return "Manzane"
	case Mittens:
		return "Mittens"
	case MogulsMail:
		return "Mogul's Mail"
	case ProwlersClaw:
		return "Prowler's Claw"
	case RapidFirecannon:
		return "Rapid Firecannon"
	case SeekersArmguard:
		return "Seeker's Armguard"
	case SilvermereDawn:
		return "Silvermere Dawn"
	case SnipersFocus:
		return "Sniper's Focus"
	case SpectralCutlass:
		return "Spectral Cutlass"
	case SuspiciousTrenchcoat:
		return "Suspicious Trenchcoat"
	case TalismanOfAscension:
		return "Talisman of Ascension"
	case TrickstersGlass:
		return "Trickster's Glass"
	case UnendingDespair:
		return "Unending Despair"
	case WitsEnd:
		return "Wit's End"
	case ZhonyasParadox:
		return "Zhonya's Paradox"

	default:
		return "Unknown Items"
	}

}

type AugmentName int

const (
	Backup AugmentName = iota
	BandOfThievesI
	BeggarsCanBeChoosers
	BladeDance
	BlisteringStrikes
	BranchingOut
	BulkyBuddiesI
	CalledShot
	CaretakersAlly
	ClimbTheLadderI
	CombatMedic
	ComponentBuffet
	Corrosion
	CraftedCrafting
	DelayedStart
	DiversifiedPortfolio
	DiversifiedPortfolioPlus
	Dummify
	EyeForAnEye
	EyeForAnEyePlus
	FindYourCenter
	FireSale
	FireSalePlus
	GlassCannonI
	GoodForSomethingI
	HeadStart
	HealthIsWealthI
	IHopeThisWorks
	IronAssets
	ItemCollectorI
	ItemGrabBagI
	Kingslayer
	LategameSpecialist
	LatentForge
	Lineup
	LoneHero
	LunchMoney
	MadChemist
	ManaFlowI
	MentorshipI
	MissedConnections
	OneForAllI
	OneTwoFive
	OneTwosThree
	OverEncumbered
	PandorasBench
	PandorasItems
	PatienceIsAVirtue
	Placebo
	PlaceboPlus
	PowerUp
	PumpingUpI
	Recombobulator
	RerollTransfer
	RestartMission
	RiggedShop
	RiggedShopPlus
	RiskyMoves
	RollingForDaysI
	SilverSpoon
	SpoilsOfWarI
	SuperstarsI
	SupportMining
	SupportMiningPlus
	Survivor
	TeamBuilding
	TeamingUpI
	TitanicTitan
	Trolling
	Underdogs
	YoungAndWildAndFree

	AcademicResearch
	AcademyCrest
	AdrenalineBurst
	AerialWarfare
	AGoldenFind
	AllThatShimmers
	AllThatShimmersPlus
	AMagicRoll
	AmbusherCrest
	AnotherAnomaly
	ArcaneRetribution
	ArtilleristCrest
	AutomataCrest
	BadLuckProtection
	BalancedBudget
	BalancedBudgetPlus
	BigGrabBag
	BlackRoseCrest
	BlazingSoulI
	BRB
	BronzeForLifeI
	BruiserCrest
	BrutalRevenge
	BuiltDifferent
	BulkyByddiesII
	CaretakersFavor
	Category5
	ChemBaronCrest
	ClearMind
	ClimbTheLadderII
	ClockworkAccelerator
	CloningFacility
	ClutteredMind
	ConquerorCrest
	CookingPot
	CrimsonPact
	CrownGuarded
	CrownsWill
	Domination
	DominatorCrest
	DragonsSpirit
	DuoQueue
	EnforcerCrest
	Epoch
	EpochPlus
	ExperimentCrest
	FamilyCrest
	FirelightCrest
	ForbiddenMagic
	ForwardThinking
	FracturedCrystals
	GlassCannonII
	GlovesOff
	GoldForDummies
	Golemify
	HealthIsWealthII
	HeavilySmash
	HeroicGrabBag
	HighVoltage
	InspiringEpitaph
	InvestmentStrategyI
	ItemCollectorII
	LawEnforcement
	LittleBuddies
	LongDistancePals
	LootExplosion
	MacesWill
	MaliciousMonetization
	ManaflowII
	MentorshipII
	Moonlight
	NobleSacrifice
	NoScoutNoPivot
	NotToday
	NoxianGuillotine
	OneForAllII
	Overheal
	PaintTheTownBlue
	PairOfFours
	PandorasItemsII
	PatientStudy
	PiercingLotusI
	Pilfer
	PitFighterCrest
	PortableForge
	PoweredShields
	PrizeFighter
	PumpingUpII
	Pyromaniac
	QuickstrikerCrest
	RainingGold
	RainingGoldPlus
	RebelCrest
	ReinFourCement
	Replication
	RocketCollection
	RocketCollectionPlus
	SalvageBin
	SalvageBinPlus
	SatedSpellweaver
	Scapegoat
	Scavenger
	ScoreboardScrapper
	ScrapChest
	SentinelCrest
	ShieldBash
	ShopGlitch
	Slammin
	SlamminPlus
	SniperCrest
	SnipersNest
	SorcererCrest
	SpearsWill
	SpiritLink
	SpoilOfWarII
	StarryNightPlus
	SuperstarsII
	SupportCache
	TeamingUpII
	TheMutationSurvives
	ThornPlatedArmor
	TombRaiderI
	TopOfTheScrapHeap
	TowerDefense
	TradeSector
	TrainingArc
	TraitBetrayal
	TraitMartialLaw
	TraitMenaces
	TraitReunion
	TraitSisters
	TraitTracker
	TraitUnlikelyDuo
	TrifectaI
	TwoMuchValue
	UnleashTheBeast
	VampiricVitality
	VisionaryCrest
	VoidCaller
	WanderingTrainerI
	WarForTheUndercity
	Warpath
	WatcherCrest
	WelcomeToThePlayground
	WhatDoesntKillYou
	WhyNotBoth
	WorthTheWaitI

	AcademyCrown
	AChangeOfFate
	AmbusherCrown
	AnExaltedAdventure
	AngerIssues
	Artifactory
	ArtilleristCrown
	AtWhatCost
	AutomataCrown
	BeltOverflow
	BirthdayPresent
	BlackRoseCrown
	BlazingSoulII
	BlindingSpeed
	BronzeForLifeII
	BruiserCrown
	BuildABud
	BulkyBuddiesIII
	BuriedTreasuresIII
	CalculatedEnhancement
	CallToChaos
	CaretakersChosen
	ChemBaronCrown
	ConquerorCrown
	Coronation
	DarkAlleyDealings
	DominatorCrown
	DualPurpose
	EnforcerCrown
	ExpectedUnexpectedness
	ExperimentCrown
	FamilyCrown
	FinalPolish
	FirelightCrown
	Flexible
	FurryOfBlows
	GhostOfFriendsPast
	GloriousEvolution
	GoingLong
	GreaterMoonlight
	HardCommit
	HedgeFund
	ImmovableObject
	ImTheCarryNow
	InvestedPlus
	InvestedPlusPlus
	InvestmentStrategyII
	LevelUp
	LivingForge
	LuckyGloves
	LuckyGlovesPlus
	MaxCap
	NewRecruit
	OneBuffTwoBuff
	PandorasItemsIII
	PhreakyFriday
	PhreakyFridayPlus
	PiercingLotusII
	PrismaticPipeline
	PrismaticTicket
	PumpingUpIII
	QualityOverQuantity
	QuickstrikerCrown
	RadiantRefactor
	RadiantRelics
	RebelCrown
	RollTheDice
	ScrapCrown
	SentinelCrown
	ShimmerscaleEssence
	ShoppingSpree
	SniperCrown
	SorcererCrown
	SpoilsOfWarIII
	Sponging
	SubscriptionService
	SwordOverflow
	TacticiansKitchen
	TheGoldenEgg
	TiniestTitan
	TiniestTitanPlus
	TombRaiderII
	TraitGeniuses
	TraitWhatCouldHaveBeen
	TrifectaII
	UpwardMobility
	VisionaryCrown
	VoidSwarm
	WanderingTrainerII
	WandOverflow
	WatcherCrown
	WhatYouTrulyAre
	WorthTheWaitII
)
