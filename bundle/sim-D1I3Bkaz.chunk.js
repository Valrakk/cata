import{l as e,n as l,o as s,q as a,s as t,R as n,a7 as p,ak as o,ae as i,Q as d,U as c,V as r,al as I,W as m,w as u,Y as S,G as h,T as g,af as f,J as T}from"./preset_utils-D8uJf-0T.chunk.js";import{R as y}from"./suggest_reforges_action-bRNPSoFg.chunk.js";import{a4 as v,a5 as A,a7 as O,G as k,bI as P,bJ as R,bK as C,bN as w,bM as q,ac as b,af as G,ad as L,ae as V,ag as M,aj as F,al as B,am as H,b as D,ah as j,a6 as N,a as W,ax as K,S as x,ay as U,az as E,b1 as _,ak as z,aq as J,F as Q,R as Y,aA as X}from"./detailed_results-DuTE03lA.chunk.js";import{W as Z,P as $,D as ee,a as le,b as se}from"./presets-zlDKdp8C.chunk.js";const ae={type:"TypeAPL",prepullActions:[{action:{castSpell:{spellId:{spellId:30146}}},doAtValue:{const:{val:"-10s"}}},{action:{castSpell:{spellId:{spellId:77801}}},doAtValue:{const:{val:"-4.0s"}}},{action:{castSpell:{spellId:{spellId:74434}}},doAtValue:{const:{val:"-4.0s"}}},{action:{castSpell:{spellId:{spellId:691}}},doAtValue:{const:{val:"-4.0s"}}},{action:{castSpell:{spellId:{otherId:"OtherActionPotion"}}},doAtValue:{const:{val:"-2.5s"}}},{action:{castSpell:{spellId:{spellId:59672}}},doAtValue:{const:{val:"-2.5s"}}},{action:{castSpell:{spellId:{spellId:29722}}},doAtValue:{const:{val:"-2.5s"}}},{action:{castSpell:{spellId:{spellId:348}}},doAtValue:{const:{val:"-0.7s"}}},{action:{castAllStatBuffCooldowns:{statType1:3,statType2:14,statType3:-1}},doAtValue:{const:{val:"-0.7s"}}},{action:{castSpell:{spellId:{spellId:82174}}},doAtValue:{const:{val:"-0.7s"}}},{action:{castSpell:{spellId:{spellId:33697}}},doAtValue:{const:{val:"-0.7s"}}}],priorityList:[{action:{castAllStatBuffCooldowns:{statType1:3,statType2:14,statType3:-1}}},{action:{autocastOtherCooldowns:{}}},{action:{condition:{and:{vals:[{auraIsKnown:{auraId:{spellId:89937}}},{auraIsActive:{auraId:{spellId:89937}}}]}},castSpell:{spellId:{spellId:77799}}}},{action:{condition:{and:{vals:[{or:{vals:[{allTrinketStatProcsActive:{statType1:3,statType2:14,statType3:11}},{cmp:{op:"OpLe",lhs:{trinketProcsMinRemainingTime:{statType1:3,statType2:14,statType3:11}},rhs:{const:{val:"1s"}}}},{and:{vals:[{auraIsActive:{auraId:{itemId:58091}}},{or:{vals:[{and:{vals:[{auraIsActive:{auraId:{spellId:75170}}},{auraIsActive:{auraId:{spellId:74241}}},{cmp:{op:"OpLe",lhs:{auraRemainingTime:{auraId:{itemId:58091}}},rhs:{const:{val:"1s"}}}}]}},{cmp:{op:"OpLe",lhs:{remainingTime:{}},rhs:{const:{val:"120s"}}}}]}},{anyTrinketStatProcsActive:{statType1:3,statType2:14,statType3:11}}]}},{and:{vals:[{auraIsActive:{auraId:{itemId:58091}}},{cmp:{op:"OpLe",lhs:{auraRemainingTime:{auraId:{itemId:58091}}},rhs:{const:{val:"3s"}}}}]}}]}},{spellCanCast:{spellId:{spellId:18540}}}]}},strictSequence:{actions:[{castSpell:{spellId:{spellId:18540}}}]}}},{action:{condition:{auraIsActive:{auraId:{spellId:59672}}},castSpell:{spellId:{itemId:58091}}}},{action:{condition:{and:{}},sequence:{name:"burst",actions:[{castSpell:{spellId:{spellId:77801}}},{castSpell:{spellId:{spellId:74434}}},{castSpell:{spellId:{spellId:691}}},{castSpell:{spellId:{spellId:59672}}},{castSpell:{spellId:{spellId:603}}},{castSpell:{spellId:{spellId:348}}},{castSpell:{spellId:{spellId:172}}}]}}},{action:{condition:{and:{vals:[{sequenceIsComplete:{sequenceName:"burst"}},{spellCanCast:{spellId:{spellId:77801}}},{spellCanCast:{spellId:{spellId:59672}}}]}},resetSequence:{sequenceName:"burst"}}},{action:{condition:{and:{vals:[{spellIsReady:{spellId:{spellId:59672}}},{cmp:{op:"OpGe",lhs:{spellTimeToReady:{spellId:{spellId:77801}}},rhs:{const:{val:"20"}}}},{cmp:{op:"OpGe",lhs:{remainingTime:{}},rhs:{const:{val:"60"}}}}]}},castSpell:{spellId:{spellId:59672}}}},{action:{condition:{and:{vals:[{spellIsReady:{spellId:{spellId:77801}}},{cmp:{op:"OpGe",lhs:{spellTimeToReady:{spellId:{spellId:59672}}},rhs:{remainingTime:{}}}}]}},castSpell:{spellId:{spellId:77801}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpLe",lhs:{currentMana:{}},rhs:{const:{val:"67000"}}}},{cmp:{op:"OpLe",lhs:{sequenceTimeToReady:{sequenceName:"burst"}},rhs:{const:{val:"10s"}}}}]}},castSpell:{spellId:{spellId:1454}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpGe",lhs:{sequenceTimeToReady:{sequenceName:"burst"}},rhs:{const:{val:"121s"}}}},{cmp:{op:"OpGt",lhs:{currentTime:{}},rhs:{const:{val:"60s"}}}}]}},castSpell:{spellId:{spellId:33697}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpLe",lhs:{sequenceTimeToReady:{sequenceName:"burst"}},rhs:{spellCastTime:{spellId:{spellId:30146}}}}},{cmp:{op:"OpLe",lhs:{spellCpm:{spellId:{spellId:30146}}},rhs:{const:{val:"1"}}}},{cmp:{op:"OpLe",lhs:{spellTimeToReady:{spellId:{spellId:77801}}},rhs:{const:{val:"10"}}}},{cmp:{op:"OpLe",lhs:{spellTimeToReady:{spellId:{spellId:59672}}},rhs:{const:{val:"10"}}}},{cmp:{op:"OpGe",lhs:{remainingTime:{}},rhs:{const:{val:"30"}}}}]}},castSpell:{spellId:{spellId:30146}}}},{action:{condition:{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:348,tag:1}}},rhs:{math:{op:"OpAdd",lhs:{spellCastTime:{spellId:{spellId:348}}},rhs:{dotTickFrequency:{spellId:{spellId:348,tag:1}}}}}}},castSpell:{spellId:{spellId:348}}}},{action:{condition:{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:172}}},rhs:{dotTickFrequency:{spellId:{spellId:172}}}}},castSpell:{spellId:{spellId:172}}}},{action:{condition:{spellIsReady:{spellId:{spellId:71521}}},castSpell:{spellId:{spellId:71521}}}},{action:{condition:{spellIsReady:{spellId:{spellId:47897}}},castSpell:{spellId:{spellId:47897}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:603}}},rhs:{dotTickFrequency:{spellId:{spellId:603}}}}},{cmp:{op:"OpGt",lhs:{remainingTime:{}},rhs:{math:{op:"OpMul",lhs:{dotTickFrequency:{spellId:{spellId:603}}},rhs:{const:{val:"2"}}}}}}]}},castSpell:{spellId:{spellId:603}}}},{action:{castSpell:{spellId:{spellId:50589}}}},{action:{condition:{and:{vals:[{auraIsKnown:{auraId:{spellId:71165}}},{auraIsActive:{auraId:{spellId:71165}}}]}},castSpell:{spellId:{spellId:29722}}}},{action:{condition:{and:{vals:[{auraIsKnown:{auraId:{spellId:63167}}},{auraIsActive:{auraId:{spellId:63167}}}]}},castSpell:{spellId:{spellId:6353}}}},{action:{castSpell:{spellId:{spellId:29722}}}},{action:{castSpell:{spellId:{spellId:1454}}}}]},te={type:"TypeAPL",prepullActions:[{action:{castSpell:{spellId:{spellId:30146}}},doAtValue:{const:{val:"-10s"}}},{action:{castSpell:{spellId:{spellId:77801}}},doAtValue:{const:{val:"-4.0s"}}},{action:{castSpell:{spellId:{spellId:74434}}},doAtValue:{const:{val:"-4.0s"}}},{action:{castSpell:{spellId:{spellId:691}}},doAtValue:{const:{val:"-4.0s"}}},{action:{castSpell:{spellId:{otherId:"OtherActionPotion"}}},doAtValue:{const:{val:"-2.5s"}}},{action:{castSpell:{spellId:{spellId:59672}}},doAtValue:{const:{val:"-2.5s"}}},{action:{castSpell:{spellId:{spellId:686}}},doAtValue:{const:{val:"-2.5s"}}},{action:{castSpell:{spellId:{spellId:348}}},doAtValue:{const:{val:"-0.7s"}}},{action:{castAllStatBuffCooldowns:{statType1:3,statType2:14,statType3:-1}},doAtValue:{const:{val:"-0.7s"}}},{action:{castSpell:{spellId:{spellId:82174}}},doAtValue:{const:{val:"-0.7s"}}},{action:{castSpell:{spellId:{spellId:33697}}},doAtValue:{const:{val:"-0.7s"}}}],priorityList:[{action:{castAllStatBuffCooldowns:{statType1:3,statType2:14,statType3:-1}}},{action:{autocastOtherCooldowns:{}}},{action:{condition:{and:{vals:[{or:{vals:[{allTrinketStatProcsActive:{statType1:3,statType2:14,statType3:11}},{cmp:{op:"OpLe",lhs:{trinketProcsMinRemainingTime:{statType1:3,statType2:14,statType3:11}},rhs:{const:{val:"1s"}}}},{and:{vals:[{auraIsActive:{auraId:{itemId:58091}}},{or:{vals:[{and:{vals:[{auraIsActive:{auraId:{spellId:75170}}},{auraIsActive:{auraId:{spellId:74241}}},{cmp:{op:"OpLe",lhs:{auraRemainingTime:{auraId:{itemId:58091}}},rhs:{const:{val:"1s"}}}}]}},{cmp:{op:"OpLe",lhs:{remainingTime:{}},rhs:{const:{val:"120s"}}}}]}},{anyTrinketStatProcsActive:{statType1:3,statType2:14,statType3:11}}]}},{and:{vals:[{auraIsActive:{auraId:{itemId:58091}}},{cmp:{op:"OpLe",lhs:{auraRemainingTime:{auraId:{itemId:58091}}},rhs:{const:{val:"3s"}}}}]}}]}},{spellCanCast:{spellId:{spellId:18540}}}]}},strictSequence:{actions:[{castSpell:{spellId:{spellId:18540}}}]}}},{action:{condition:{and:{}},sequence:{name:"burst",actions:[{castSpell:{spellId:{spellId:77801}}},{castSpell:{spellId:{spellId:74434}}},{castSpell:{spellId:{spellId:691}}},{castSpell:{spellId:{spellId:59672}}},{castSpell:{spellId:{spellId:603}}},{castSpell:{spellId:{spellId:348}}},{castSpell:{spellId:{spellId:172}}}]}}},{action:{condition:{and:{vals:[{sequenceIsComplete:{sequenceName:"burst"}},{spellCanCast:{spellId:{spellId:77801}}},{spellCanCast:{spellId:{spellId:59672}}}]}},resetSequence:{sequenceName:"burst"}}},{action:{condition:{and:{vals:[{auraIsActive:{auraId:{spellId:59672}}},{auraIsKnown:{auraId:{spellId:59672}}}]}},castSpell:{spellId:{itemId:58091}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpLe",lhs:{currentMana:{}},rhs:{const:{val:"67000"}}}},{cmp:{op:"OpLe",lhs:{sequenceTimeToReady:{sequenceName:"burst"}},rhs:{const:{val:"10s"}}}}]}},castSpell:{spellId:{spellId:1454}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpGe",lhs:{sequenceTimeToReady:{sequenceName:"burst"}},rhs:{const:{val:"121s"}}}},{cmp:{op:"OpGt",lhs:{currentTime:{}},rhs:{const:{val:"60s"}}}}]}},castSpell:{spellId:{spellId:33697}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpLe",lhs:{sequenceTimeToReady:{sequenceName:"burst"}},rhs:{spellCastTime:{spellId:{spellId:30146}}}}},{cmp:{op:"OpLe",lhs:{spellCpm:{spellId:{spellId:30146}}},rhs:{const:{val:"1"}}}}]}},castSpell:{spellId:{spellId:30146}}}},{action:{condition:{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:348,tag:1}}},rhs:{math:{op:"OpAdd",lhs:{spellCastTime:{spellId:{spellId:348}}},rhs:{dotTickFrequency:{spellId:{spellId:348,tag:1}}}}}}},castSpell:{spellId:{spellId:348}}}},{action:{condition:{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:172}}},rhs:{dotTickFrequency:{spellId:{spellId:172}}}}},castSpell:{spellId:{spellId:172}}}},{action:{condition:{spellIsReady:{spellId:{spellId:71521}}},castSpell:{spellId:{spellId:71521}}}},{action:{condition:{and:{vals:[{auraIsKnown:{auraId:{spellId:89937}}},{auraIsActive:{auraId:{spellId:89937}}}]}},castSpell:{spellId:{spellId:77799}}}},{action:{condition:{spellIsReady:{spellId:{spellId:47897}}},castSpell:{spellId:{spellId:47897}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:603}}},rhs:{dotTickFrequency:{spellId:{spellId:603}}}}},{cmp:{op:"OpGt",lhs:{remainingTime:{}},rhs:{math:{op:"OpMul",lhs:{dotTickFrequency:{spellId:{spellId:603}}},rhs:{const:{val:"2"}}}}}}]}},castSpell:{spellId:{spellId:603}}}},{action:{castSpell:{spellId:{spellId:50589}}}},{action:{condition:{and:{vals:[{auraIsKnown:{auraId:{spellId:17941}}},{auraIsActive:{auraId:{spellId:17941}}}]}},castSpell:{spellId:{spellId:686}}}},{action:{condition:{and:{vals:[{auraIsKnown:{auraId:{spellId:71165}}},{auraIsActive:{auraId:{spellId:71165}}}]}},castSpell:{spellId:{spellId:29722}}}},{action:{condition:{and:{vals:[{auraIsKnown:{auraId:{spellId:63167}}},{auraIsActive:{auraId:{spellId:63167}}}]}},castSpell:{spellId:{spellId:6353}}}},{action:{castSpell:{spellId:{spellId:686}}}},{action:{castSpell:{spellId:{spellId:1454}}}}]},ne={items:[{id:60237,enchant:4207,gems:[68780,52217],reforging:144},{id:69882,randomSuffix:-114},{id:65263,enchant:4200,gems:[52207],reforging:154},{id:60232,enchant:4115,gems:[52208],reforging:151},{id:65262,enchant:4102,gems:[52207,52217],reforging:144},{id:65138,enchant:4257,gems:[0],reforging:151},{id:65259,enchant:4107,gems:[52207,0],reforging:144},{id:65034,gems:[52207,52207],reforging:144},{id:65261,enchant:4112,gems:[52208,52207],reforging:140},{id:65069,enchant:4104,gems:[52207],reforging:165},{id:65123,reforging:151},{id:65373,randomSuffix:-114,reforging:138},{id:65053,reforging:144},{id:62047,reforging:165},{id:68132,randomSuffix:-114,enchant:4097},{id:65133,enchant:4091,reforging:154},{id:59460,reforging:167}]},pe={items:[{id:71595,enchant:4207,gems:[68780,52207],reforging:151},{id:71472,gems:[52207],reforging:165},{id:71598,enchant:4200,gems:[52207],reforging:154},{id:71434,enchant:4115,reforging:144},{id:71597,enchant:4102,gems:[52207,52207],reforging:154},{id:71471,enchant:4257,gems:[0],reforging:144},{id:71614,enchant:4107,gems:[52207,0],reforging:147},{id:71613,gems:[52207,52207],reforging:151},{id:71596,enchant:4112,gems:[52207,52207],reforging:144},{id:71447,enchant:4104,gems:[52207],reforging:144},{id:71217,gems:[52207]},{id:71449,reforging:144},{id:69110},{id:62047,reforging:165},{id:71086,enchant:4097,gems:[52207,52207,52207],reforging:154},{},{id:71473,reforging:144}]},oe=e("Pre-raid",{items:[]}),ie=e("P1 - BIS",ne),de=e("P3 - BIS",pe),ce=l("Shadow Bolt",te),re=l("Incinerate",ae),Ie=s("Default",v.fromMap({[A.StatIntellect]:1.27,[A.StatSpellPower]:1,[A.StatHitRating]:.92,[A.StatCritRating]:.51,[A.StatHasteRating]:2.75,[A.StatMasteryRating]:.76})),me={name:"Shadow bolt",data:O.create({talentsString:"-3312222300310212211-33202",glyphs:k.create({prime1:P.GlyphOfImmolate,prime2:P.GlyphOfCorruption,prime3:P.GlyphOfMetamorphosis,major1:R.GlyphOfShadowBolt,major2:R.GlyphOfLifeTap,major3:R.GlyphOfFelhunter,minor1:C.GlyphOfDrainSoul,minor2:C.GlyphOfRitualOfSouls,minor3:C.GlyphOfUnendingBreath})})},ue={name:"Incinerate",data:O.create({talentsString:"003-3312222300310212211-03202",glyphs:k.create({prime1:P.GlyphOfImmolate,prime2:P.GlyphOfIncinerate,prime3:P.GlyphOfMetamorphosis,major1:R.GlyphOfSoulstone,major2:R.GlyphOfLifeTap,major3:R.GlyphOfSoulLink,minor1:C.GlyphOfDrainSoul,minor2:C.GlyphOfRitualOfSouls,minor3:C.GlyphOfUnendingBreath})})},Se=w.create({classOptions:{summon:q.Felguard,detonateSeed:!1,prepullMastery:0,prepullPostSnapshotMana:0}}),he=b.create({defaultPotion:G.VolcanicPotion,prepopPotion:G.VolcanicPotion,flask:L.FlaskOfTheDraconicMind,food:V.FoodSeafoodFeast,tinkerHands:M.TinkerHandsSynapseSprings,explosiveBigDaddy:!1}),ge=F.create({arcaneBrilliance:!0,bloodlust:!0,markOfTheWild:!0,icyTalons:!0,moonkinForm:!0,leaderOfThePack:!0,powerWordFortitude:!0,strengthOfEarthTotem:!0,trueshotAura:!0,wrathOfAirTotem:!0,demonicPact:!0,blessingOfKings:!0,blessingOfMight:!0,communion:!0}),fe=B.create({vampiricTouch:!0,darkIntent:!0}),Te=H.create({bloodFrenzy:!0,sunderArmor:!0,ebonPlaguebringer:!0,mangle:!0,criticalMass:!1,demoralizingShout:!0,frostFever:!0}),ye={distanceFromTarget:25,profession1:D.Engineering,profession2:D.Tailoring,channelClipDelay:150,duration:180,durationVariation:30,darkIntentUptime:90},ve=a("Shadow Bolt",{talents:me,rotation:ce}),Ae=a("Incinerate",{talents:ue,rotation:re}),Oe=[{unitStat:j.fromPseudoStat(N.PseudoStatSpellHastePercent),presets:new Map([...Z.find((e=>e.unitStat.equalsPseudoStat(N.PseudoStatSpellHastePercent))).presets,["16-tick - Immo Aura",3.35918],["8-tick - Immolate:Inferno",7.12373],["17-tick - Immo Aura",9.95053],["18-tick - Immo Aura",16.61809],["9-tick - Immolate:Inferno",21.43291],["19-tick - Immo Aura",23.38064],["20-tick - Immo Aura",29.95453],["10-tick - Immolate:Inferno",35.71591],["21-tick - Immo Aura",36.70542],["22-tick - Immo Aura",43.3692],["11-tick - Immolate:Inferno",50.03752],["23-tick - Immo Aura",50.03753],["24-tick - Immo Aura",56.6171],["25-tick - Immo Aura",63.26533],["12-tick - Immolate:Inferno",64.24857],["26-tick - Immo Aura",69.92356],["27-tick - Immo Aura",76.52254],["13-tick - Immolate:Inferno",78.6246],["28-tick - Immo Aura",83.31809],["29-tick - Immo Aura",89.93356],["14-tick - Immolate:Inferno",92.86404],["30-tick - Immo Aura",96.65687],["31-tick - Immo Aura",103.45884],["15-tick - Immolate:Inferno",107.11082],["32-tick - Immo Aura",109.86363]])}],ke=t(x.SpecDemonologyWarlock,{cssClass:"demonology-warlock-sim-ui",cssScheme:W.getCssClass(W.Warlock),knownIssues:[],epStats:[A.StatIntellect,A.StatSpellPower,A.StatHitRating,A.StatCritRating,A.StatHasteRating,A.StatMasteryRating],epReferenceStat:A.StatSpellPower,displayStats:j.createDisplayStatArray([A.StatHealth,A.StatMana,A.StatStamina,A.StatIntellect,A.StatSpellPower,A.StatMasteryRating,A.StatMP5],[N.PseudoStatSpellHitPercent,N.PseudoStatSpellCritPercent,N.PseudoStatSpellHastePercent]),defaults:{gear:ie.gear,epWeights:Ie.epWeights,statCaps:(new v).withPseudoStat(N.PseudoStatSpellHitPercent,17),softCapBreakpoints:(()=>{const e=[],l=K.get(x.SpecDemonologyWarlock);for(let a=19;a<=200;a++)e.push(a/l*X);const s=U.fromStat(A.StatMasteryRating,{breakpoints:e,capType:E.TypeThreshold,postCapEPs:[0]});return[U.fromPseudoStat(N.PseudoStatSpellHastePercent,{breakpoints:[16.65,25],capType:E.TypeSoftCap,postCapEPs:[.64*_,.61*_]}),s]})(),consumes:he,talents:me.data,specOptions:Se,raidBuffs:ge,partyBuffs:z.create({}),individualBuffs:fe,debuffs:Te,other:ye},playerIconInputs:[$()],includeBuffDebuffInputs:[n,p,o,i,d,c,r,I,m],excludeBuffDebuffInputs:[],petConsumeInputs:[],otherInputs:{inputs:[ee(),le,se,u,S,h,g,f]},itemSwapSlots:[J.ItemSlotMainHand,J.ItemSlotOffHand,J.ItemSlotRanged],encounterPicker:{showExecuteProportion:!1},presets:{epWeights:[Ie],talents:[me,ue],rotations:[ce,re],gear:[oe,ie,de],builds:[ve,Ae]},autoRotation:e=>ce.rotation.rotation,raidSimPresets:[{spec:x.SpecDemonologyWarlock,talents:me.data,specOptions:Se,consumes:he,defaultFactionRaces:{[Q.Unknown]:Y.RaceUnknown,[Q.Alliance]:Y.RaceHuman,[Q.Horde]:Y.RaceOrc},defaultGear:{[Q.Unknown]:{},[Q.Alliance]:{1:oe.gear,2:ie.gear,3:de.gear},[Q.Horde]:{1:oe.gear,2:ie.gear,3:de.gear}},otherDefaults:ye}]});class Pe extends T{constructor(e,l){super(e,l,ke),l.sim.waitForInit().then((()=>{new y(this,{statSelectionPresets:Oe})}))}}export{Pe as D};
