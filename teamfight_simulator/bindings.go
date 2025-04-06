package main

func RegisterTrait(champion *Champion, traitName string, traitCount int) {
	switch traitName {

	}
}

func RegisterAbility(unit Unit) {

	switch unit.Name() {
	case "Akali":
		RegisterAbilityAkali(unit)
	case "Ambessa":
		RegisterAbilityAmbessa(unit)
	case "Amumu":
		RegisterAbilityAmumu(unit)
	case "Blitzcrank":
		RegisterAbilityBlitzcrank(unit)
	case "Caitlyn":
		RegisterAbilityCaitlyn(unit)
	case "Camille":
		RegisterAbilityCamille(unit)
	case "Cassiopeia":
		RegisterAbilityCassiopeia(unit)
	case "Corki":
		RegisterAbilityCorki(unit)
	case "Darius":
		RegisterAbilityDarius(unit)
	case "Dr. Mundo":
		RegisterAbilityDrMundo(unit)
	case "Draven":
		RegisterAbilityDraven(unit)
	case "Ekko":
		RegisterAbilityEkko(unit)
	case "Elise":
		RegisterAbilityElise(unit)
	case "Ezreal":
		RegisterAbilityEzreal(unit)
	case "Gangplank":
		RegisterAbilityGangplank(unit)
	case "Garen":
		RegisterAbilityGaren(unit)
	case "Heimerdinger":
		RegisterAbilityHeimerdinger(unit)
	case "Illaoi":
		RegisterAbilityIllaoi(unit)
	case "Irelia":
		RegisterAbilityIrelia(unit)
	case "Jayce":
		RegisterAbilityJayce(unit)
	case "Jinx":
		RegisterAbilityJinx(unit)
	case "Kog'Maw":
		RegisterAbilityKogMaw(unit)
	case "LeBlanc":
		RegisterAbilityLeBlanc(unit)
	case "Leona":
		RegisterAbilityLeona(unit)
	case "Loris":
		RegisterAbilityLoris(unit)
	case "Lux":
		RegisterAbilityLux(unit)
	case "Maddie":
		RegisterAbilityMaddie(unit)
	case "Malzahar":
		RegisterAbilityMalzahar(unit)
	case "Minion":
		RegisterAbilityMinion(unit)
	case "Mordekaiser":
		RegisterAbilityMordekaiser(unit)
	case "Morgana":
		RegisterAbilityMorgana(unit)
	case "Nami":
		RegisterAbilityNami(unit)
	case "Nocturne":
		RegisterAbilityNocturne(unit)
	case "Nunu & Willump":
		RegisterAbilityNunuWillump(unit)
	case "Powder":
		RegisterAbilityPowder(unit)
	case "Rell":
		RegisterAbilityRell(unit)
	case "Renata Glasc":
		RegisterAbilityRenataGlasc(unit)
	case "Renni":
		RegisterAbilityRenni(unit)
	case "Rumble":
		RegisterAbilityRumble(unit)
	case "Scar":
		RegisterAbilityScar(unit)
	case "Sett":
		RegisterAbilitySett(unit)
	case "Sevika":
		RegisterAbilitySevika(unit)
	case "Silco":
		RegisterAbilitySilco(unit)
	case "Singed":
		RegisterAbilitySinged(unit)
	case "Smeech":
		RegisterAbilitySmeech(unit)
	case "Steb":
		RegisterAbilitySteb(unit)
	case "Swain":
		RegisterAbilitySwain(unit)
	case "Tristana":
		RegisterAbilityTristana(unit)
	case "Trundle":
		RegisterAbilityTrundle(unit)
	case "Twisted Fate":
		RegisterAbilityTwistedFate(unit)
	case "Twitch":
		RegisterAbilityTwitch(unit)
	case "Urgot":
		RegisterAbilityUrgot(unit)
	case "Vander":
		RegisterAbilityVander(unit)
	case "Vex":
		RegisterAbilityVex(unit)
	case "Vi":
		RegisterAbilityVi(unit)
	case "Violet":
		RegisterAbilityViolet(unit)
	case "Vladimir":
		RegisterAbilityVladimir(unit)
	case "Zeri":
		RegisterAbilityZeri(unit)
	case "Ziggs":
		RegisterAbilityZiggs(unit)
	case "Zoe":
		RegisterAbilityZoe(unit)
	case "Zyra":
		RegisterAbilityZyra(unit)
	}
}

func RegisterItem(itemName ItemName, isRadiant bool, unit Unit) {
	switch itemName {
	case BFSword:
		RegisterBFSword(unit)
	case ChainVest:
		RegisterChainVest(unit)

	case GiantsBelt:
		RegisterGiantsBelt(unit)
	case NeedlesslyLargeRod:
		RegisterNeedlesslyLargeRod(unit)
	case NegatronCloak:
		RegisterNegatronCloak(unit)
	case RecurveBow:
		RegisterRecurveBow(unit)
	case SparringGloves:
		RegisterSparringGloves(unit)

	case TearOfTheGoddess:
		RegisterTearOfTheGoddess(unit)

	case AdaptiveHelm:
		RegisterAdaptiveHelm(unit, isRadiant)
	case ArchangelsStaff:
		RegisterArchangelsStaff(unit, isRadiant)
	case Bloodthirster:
		RegisterBloodthirster(unit, isRadiant)
	case BlueBuff:
		RegisterBlueBuff(unit, isRadiant)
	case BrambleVest:
		RegisterBrambleVest(unit, isRadiant)
	case Crownguard:
		RegisterCrownguard(unit, isRadiant)
	case Deathblade:
		RegisterDeathblade(unit, isRadiant)
	case DragonsClaw:
		RegisterDragonsClaw(unit, isRadiant)
	case EdgeOfNight:
		RegisterEdgeOfNight(unit, isRadiant)
	case Evenshroud:
		RegisterEvenshroud(unit, isRadiant)
	case GargoyleStoneplate:
		RegisterGargoyleStoneplate(unit, isRadiant)
	case GiantSlayer:
		RegisterGiantSlayer(unit, isRadiant)
	case Guardbreaker:
		RegisterGuardbreaker(unit, isRadiant)
	case GuinsoosRageblade:
		RegisterGuinsoosRageblade(unit, isRadiant)
	case HandsOfJustice:
		RegisterHandsOfJustice(unit, isRadiant)
	case HextechGunblade:
		RegisterHextechGunblade(unit, isRadiant)
	case InfinityEdge:
		RegisterInfinityEdge(unit, isRadiant)
	case IonicSpark:
		RegisterIonicSpark(unit, isRadiant)
	case JeweledGauntlet:
		RegisterJeweledGauntlet(unit, isRadiant)
	case LastWhisper:
		RegisterLastWhisper(unit, isRadiant)
	case Morellonomicon:
		RegisterMorellonomicon(unit, isRadiant)
	case NashorsTooth:
		RegisterNashorsTooth(unit, isRadiant)
	case ProtectorsVow:
		RegisterProtectorsVow(unit, isRadiant)
	case Quicksilver:
		RegisterQuicksilver(unit, isRadiant)
	case RabadonsDeathcap:
		RegisterRabadonsDeathcap(unit, isRadiant)
	case RedBuff:
		RegisterRedBuff(unit, isRadiant)
	case Redemption:
		RegisterRedemption(unit, isRadiant)
	case RunaansHurricane:
		RegisterRunaansHurricane(unit, isRadiant)
	case SpearOfShojin:
		RegisterSpearOfShojin(unit, isRadiant)
	case StatikkShiv:
		RegisterStatikkShiv(unit, isRadiant)
	case SteadfastHeart:
		RegisterSteadfastHeart(unit, isRadiant)
	case SteraksGage:
		RegisterSteraksGage(unit, isRadiant)
	case SunfireCape:
		RegisterSunfireCape(unit, isRadiant)
	case ThiefsGloves:
		RegisterThiefsGloves(unit, isRadiant)
	case TitansResolve:
		RegisterTitansResolve(unit, isRadiant)
	case WarmogsArmor:
		RegisterWarmogsArmor(unit, isRadiant)

	default:
		panic("Unknown Item in RegisterItem()")
	}

}

func RegisterAugment(augmentName AugmentName, champion *UnitData) {
	// TODO: This has to be clever because there are 310 augments
}
