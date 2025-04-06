package main

type Event struct {
	Type    EventType
	Payload []int
	// Thornskin
	//
}

func NewEvent(eventType EventType) *Event {
	return &Event{
		eventType,
		make([]int, 0),
	}
}

func NewEventPayload(eventType EventType, payload ...int) *Event {
	return &Event{
		eventType,
		payload,
	}
}

type EventType int

const (
	OnAlways EventType = iota

	OnPositionChange
	// Sunfire Cape

	OnCombatStart // Done
	// Adaptive Helm
	// Crownguard
	// Evenshroud
	// Hands of Justice
	// Quicksilver
	// Thiefs Gloves

	OnOverheal

	OnTick
	// Evenshroud (Keeps track of enemies within 2 hexes)
	// Ionic Spark

	// TODO: Make this an Interval Event
	// Archangel's Staff
	// Dragonsclaw
	// Redemption
	// Sunfire Cape
	// Crownguard (One-time)
	// Quicksilver (One-time)
	// Tactician's Cape (One-time)

	OnAttacked
	// Bramble Vest

	OnAttack
	// Guinsoo's
	// Runaan's Hurricane
	// Spear of Shojin
	// Statikk Shiv
	// Titan's Resolve

	OnHealthChange // Done
	// Edge of Night
	// Bloodthirster
	// Protector's Vow
	// Steadfast Heart
	// Sterak's Gage
	// Tactician's Shield

	OnReceivingDamage // Done
	// Titan's Resolve

	OnDealingDamage // Done
	// Giantslayer
	// Guardbreaker
	// Hextech Gunblade
	// Last Whisper
	// Morellonomicon
	// Red Buff

	OnCast

	OnCastEnd
	// Nashor's tooth

	OnTakedown // Done
	// Blue buff

	OnKill // Done

	OnDeath // Done

	OnTargeterChange
	// Gargoyle Stoneplate

	OnEnemyAbilityCast
	// Ionic Spark

	OnCombatEnd
	// Tactician's Crown

	OnDamageShield
	// Guardbreaker

	OnEffectExpiration // TODO: implement?????

)
