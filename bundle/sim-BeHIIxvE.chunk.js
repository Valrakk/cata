import{l as e,o as a,s as t,w as s,J as n}from"./preset_utils-D8uJf-0T.chunk.js";import{a4 as i,a5 as d,a7 as c,bH as r,bz as m,ac as g,af as h,ad as o,ae as l,a as S,ah as u,a6 as p,aj as f,ak as P,al as R,am as k,aH as w,S as O,F,R as H}from"./detailed_results-DuTE03lA.chunk.js";import{S as I,T as M}from"./inputs-CWxnBPtq.chunk.js";const y={items:[{id:40510,enchant:3820,gems:[41401,39998]},{id:44662,gems:[40051]},{id:40513,enchant:3810,gems:[39998]},{id:44005,enchant:3831,gems:[40027]},{id:40508,enchant:2381,gems:[39998,40051]},{id:40209,enchant:2332,gems:[0]},{id:40564,enchant:3246,gems:[0]},{id:40327,gems:[39998]},{id:40512,enchant:3721,gems:[39998,40027]},{id:39734,enchant:3244},{id:40399},{id:40375},{id:37111},{id:40685},{id:40395,enchant:3834},{id:40401,enchant:1128},{id:40709}]},T={items:[{id:46201,enchant:3820,gems:[41401,40027]},{id:45443,gems:[40027]},{id:46204,enchant:3810,gems:[45883]},{id:45486,enchant:3831,gems:[40051]},{id:45867,enchant:2381,gems:[40051,39998]},{id:45460,enchant:2332,gems:[40027,0]},{id:46199,enchant:3246,gems:[40051,0]},{id:45151,gems:[39998]},{id:46202,enchant:3721,gems:[39998,40027]},{id:45615,enchant:3232,gems:[39998,40027]},{id:45614,gems:[40051]},{id:46046,gems:[40051]},{id:45535},{id:45466},{id:46017,enchant:3834},{id:45470,enchant:1128,gems:[40027]},{id:45114}]},b={items:[{id:46201,enchant:3820,gems:[41401,40134]},{id:47468,gems:[40128]},{id:46204,enchant:3810,gems:[40155]},{id:47551,enchant:3831,gems:[40155]},{id:46198,enchant:2381,gems:[40155,40155]},{id:45460,enchant:2332,gems:[40155,0]},{id:46199,enchant:3246,gems:[40128,0]},{id:47426,gems:[40155,40155,40155]},{id:47479,enchant:3721,gems:[40155,40155,40155]},{id:47456,enchant:3232,gems:[40155,40155]},{id:45614,gems:[40128]},{id:47439,gems:[40128]},{id:45535},{id:47059},{id:46017,enchant:3834},{id:47448,enchant:1128,gems:[40128]},{id:47665}]},A={items:[{id:51247,enchant:3820,gems:[41401,40155]},{id:50724,gems:[40128]},{id:51245,enchant:3810,gems:[40179]},{id:50628,enchant:3831,gems:[40128]},{id:51249,enchant:2381,gems:[40128,40128]},{id:50687,enchant:2332,gems:[40155,0]},{id:50703,enchant:3246,gems:[40128,40128,0]},{id:50613,gems:[40128,40128,40128]},{id:51246,enchant:3721,gems:[40155,40128]},{id:50699,enchant:3232,gems:[40128,40128]},{id:50664,gems:[40128]},{id:50400,gems:[40128]},{id:50366},{id:47059},{id:46017,enchant:3834},{id:50616,enchant:1128,gems:[40128]},{id:47665}]},B=e("PreRaid",{items:[{id:42555,enchant:3820,gems:[41401,40017]},{id:40681},{id:37875,enchant:3838,gems:[40017]},{id:37291,enchant:3859},{id:44180,enchant:2381},{id:37788,enchant:3758,gems:[0]},{id:37623,enchant:3604,gems:[0]},{id:40693,gems:[40051,0]},{id:37791,enchant:3721},{id:44202,enchant:4223,gems:[40105]},{id:44283},{id:37694},{id:37111},{id:37657},{id:37169,enchant:3834},{id:40700,enchant:1128},{id:40709}]}),W=e("P1 Preset",y),C=e("P2 Preset",T),j=e("P3 Preset",b),x=e("P4 Preset",A),D=a("P1",i.fromMap({[d.StatIntellect]:.22,[d.StatSpirit]:.05,[d.StatSpellPower]:1,[d.StatCritRating]:.67,[d.StatHasteRating]:1.29,[d.StatMP5]:.08})),E={name:"Tank Healing",data:c.create({})},U={name:"Raid Healing",data:c.create({})},v=r.create({classOptions:{shield:m.WaterShield},earthShieldPPM:0}),_=g.create({defaultPotion:h.RunicManaInjector,flask:o.FlaskOfTheFrostWyrm,food:l.FoodFishFeast}),z=t(O.SpecRestorationShaman,{cssClass:"restoration-shaman-sim-ui",cssScheme:S.getCssClass(S.Shaman),knownIssues:[],warnings:[],epStats:[d.StatIntellect,d.StatSpirit,d.StatSpellPower,d.StatCritRating,d.StatHasteRating,d.StatMP5,d.StatMasteryRating],epReferenceStat:d.StatSpellPower,displayStats:u.createDisplayStatArray([d.StatHealth,d.StatMana,d.StatStamina,d.StatIntellect,d.StatSpirit,d.StatSpellPower,d.StatMP5,d.StatMasteryRating],[p.PseudoStatSpellCritPercent,p.PseudoStatSpellHastePercent]),defaults:{gear:W.gear,epWeights:D.epWeights,consumes:_,talents:U.data,specOptions:v,raidBuffs:f.create({arcaneBrilliance:!0,bloodlust:!0,markOfTheWild:!0,icyTalons:!0,moonkinForm:!0,leaderOfThePack:!0,powerWordFortitude:!0,strengthOfEarthTotem:!0,trueshotAura:!0,wrathOfAirTotem:!0,demonicPact:!0,blessingOfKings:!0,blessingOfMight:!0,communion:!0}),partyBuffs:P.create({}),individualBuffs:R.create({vampiricTouch:!0}),debuffs:k.create({curseOfElements:!0,shadowAndFlame:!0})},playerIconInputs:[I()],includeBuffDebuffInputs:[],excludeBuffDebuffInputs:[],otherInputs:{inputs:[s]},customSections:[M],encounterPicker:{showExecuteProportion:!1},presets:{epWeights:[D],talents:[U,E],rotations:[],gear:[B,W,C,j,x]},autoRotation:e=>w.create(),raidSimPresets:[{spec:O.SpecRestorationShaman,talents:U.data,specOptions:v,consumes:_,defaultFactionRaces:{[F.Unknown]:H.RaceUnknown,[F.Alliance]:H.RaceDraenei,[F.Horde]:H.RaceOrc},defaultGear:{[F.Unknown]:{},[F.Alliance]:{1:W.gear,2:C.gear,3:j.gear,4:x.gear},[F.Horde]:{1:W.gear,2:C.gear,3:j.gear,4:x.gear}}}]});class G extends n{constructor(e,a){super(e,a,z)}}export{G as R};
