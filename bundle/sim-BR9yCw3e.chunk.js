import{m as e,k as t,l as s,n as a,o as l,a6 as n,a4 as i,a5 as o,s as p,W as r,t as d,a7 as c,w as g,Y as m,T as I,F as h,G as S,J as u}from"./preset_utils-D8uJf-0T.chunk.js";import{R as f}from"./suggest_reforges_action-bRNPSoFg.chunk.js";import{aQ as v,aR as T,T as A,a4 as O,a5 as y,a6 as R,a7 as P,G as w,aT as k,aU as G,aW as C,H,ac as V,af as W,ad as E,ae as x,b as F,a as b,ah as D,aj as L,ak as J,al as j,am as U,aG as B,aH as M,aI as N,S as _,F as K,R as z}from"./detailed_results-DuTE03lA.chunk.js";import{s as Q}from"./apl_utils-CsETylHS.chunk.js";import{s as Y,P as $,a as q,A as X,N as Z}from"./shared-BeHk80VZ.chunk.js";const ee=e({fieldName:"sniperTrainingUptime",label:"ST Uptime (%)",labelTooltip:"Uptime for the Sniper Training talent, as a percent of the fight duration.",percent:!0,showWhen:e=>e.getTalents().sniperTraining>0,changeEmitter:e=>A.onAny([e.specOptionsChangeEmitter,e.talentsChangeEmitter])}),te={inputs:[t({fieldName:"type",label:"Type",values:[{name:"Single Target",value:v.SingleTarget},{name:"AOE",value:v.Aoe}]}),t({fieldName:"sting",label:"Sting",labelTooltip:"Maintains the selected Sting on the primary target.",values:[{name:"None",value:T.NoSting},{name:"Serpent Sting",value:T.SerpentSting}],showWhen:e=>e.getSimpleRotation().type==v.SingleTarget})]},se={type:"TypeAPL",prepullActions:[{action:{castSpell:{spellId:{spellId:13165}}},doAtValue:{const:{val:"-10s"}}},{action:{castSpell:{spellId:{spellId:1130}}},doAtValue:{const:{val:"-5s"}}},{action:{castSpell:{spellId:{otherId:"OtherActionPotion"}}},doAtValue:{const:{val:"-1s"}}},{action:{castSpell:{spellId:{spellId:13812}}},doAtValue:{const:{val:"-1s"}}}],priorityList:[{action:{condition:{cmp:{op:"OpGt",lhs:{currentTime:{}},rhs:{const:{val:"1s"}}}},autocastOtherCooldowns:{}}},{action:{condition:{auraIsActive:{auraId:{spellId:77769}}},castSpell:{spellId:{spellId:13812}}}},{action:{condition:{spellIsReady:{spellId:{spellId:13812}}},castSpell:{spellId:{spellId:77769}}}},{action:{castSpell:{spellId:{spellId:2643}}}},{action:{castSpell:{spellId:{spellId:53351}}}},{action:{condition:{and:{vals:[{auraIsActive:{auraId:{spellId:56343}}},{or:{vals:[{not:{val:{dotIsActive:{spellId:{spellId:53301}}}}},{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:53301}}},rhs:{math:{op:"OpAdd",lhs:{spellTravelTime:{spellId:{spellId:53301}}},rhs:{const:{val:"1s"}}}}}}]}}]}},castSpell:{spellId:{spellId:53301}}}},{action:{castSpell:{spellId:{spellId:77767}}}}]},ae={type:"TypeAPL",prepullActions:[{action:{castSpell:{spellId:{spellId:13812}}},doAtValue:{const:{val:"-25s"}}},{action:{castSpell:{spellId:{otherId:"OtherActionPotion"}}},doAtValue:{const:{val:"-1.4s"}}},{action:{castSpell:{spellId:{spellId:77767}}},doAtValue:{const:{val:"-1.4s"}}},{action:{castSpell:{spellId:{spellId:13165}}},doAtValue:{const:{val:"-10s"}}},{action:{castSpell:{spellId:{spellId:1130}}},doAtValue:{const:{val:"-3s"}}},{action:{castSpell:{spellId:{spellId:53517}}},doAtValue:{const:{val:"-4s"}},hide:!0},{action:{triggerIcd:{auraId:{spellId:97125}}},doAtValue:{const:{val:"-40s"}},hide:!0}],priorityList:[{action:{condition:{cmp:{op:"OpGt",lhs:{currentTime:{}},rhs:{const:{val:"3s"}}}},autocastOtherCooldowns:{}}},{action:{condition:{or:{vals:[{isExecutePhase:{threshold:"E20"}},{cmp:{op:"OpLe",lhs:{remainingTime:{}},rhs:{const:{val:"25s"}}}}]}},castSpell:{spellId:{itemId:58145}}}},{hide:!0,action:{condition:{not:{val:{dotIsActive:{spellId:{spellId:1978}}}}},castSpell:{spellId:{spellId:2643}}}},{action:{condition:{not:{val:{dotIsActive:{spellId:{spellId:1978}}}}},castSpell:{spellId:{spellId:1978}}}},{action:{castSpell:{spellId:{spellId:53351}}}},{action:{condition:{and:{vals:[{auraIsActive:{auraId:{spellId:56343}}},{dotIsActive:{spellId:{spellId:53301}}},{auraIsActive:{auraId:{spellId:99060}}}]}},castSpell:{spellId:{spellId:3044}}}},{action:{condition:{and:{vals:[{dotIsActive:{spellId:{spellId:53301}}},{cmp:{op:"OpGe",lhs:{dotRemainingTime:{spellId:{spellId:53301}}},rhs:{math:{op:"OpAdd",lhs:{spellTravelTime:{spellId:{spellId:53301}}},rhs:{const:{val:"1000ms"}}}}}},{spellCanCast:{spellId:{spellId:53301}}},{cmp:{op:"OpLe",lhs:{dotRemainingTime:{spellId:{spellId:53301}}},rhs:{math:{op:"OpAdd",lhs:{spellTravelTime:{spellId:{spellId:53301}}},rhs:{const:{val:"1050ms"}}}}}}]}},wait:{duration:{const:{val:"51ms"}}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpLt",lhs:{spellTimeToReady:{spellId:{spellId:53301}}},rhs:{const:{val:"250ms"}}}},{cmp:{op:"OpGt",lhs:{currentFocus:{}},rhs:{const:{val:"44"}}}}]}},waitUntil:{condition:{spellIsReady:{spellId:{spellId:53301}}}}}},{action:{condition:{or:{vals:[{not:{val:{dotIsActive:{spellId:{spellId:53301}}}}},{cmp:{op:"OpLt",lhs:{dotRemainingTime:{spellId:{spellId:53301}}},rhs:{math:{op:"OpAdd",lhs:{spellTravelTime:{spellId:{spellId:53301}}},rhs:{const:{val:"1000ms"}}}}}}]}},castSpell:{spellId:{spellId:53301}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpGe",lhs:{remainingTime:{}},rhs:{const:{val:"8s"}}}},{spellIsReady:{spellId:{spellId:3674}}}]}},castSpell:{spellId:{spellId:3674}}}},{action:{condition:{cmp:{op:"OpGe",lhs:{currentFocus:{}},rhs:{const:{val:"80"}}}},castSpell:{spellId:{spellId:3044}}}},{action:{condition:{and:{vals:[{cmp:{op:"OpGe",lhs:{currentFocus:{}},rhs:{const:{val:"40"}}}},{cmp:{op:"OpLe",lhs:{remainingTime:{}},rhs:{const:{val:"8s"}}}},{cmp:{op:"OpGe",lhs:{spellTimeToReady:{spellId:{spellId:53301}}},rhs:{const:{val:"1s"}}}}]}},castSpell:{spellId:{spellId:3044}}}},{hide:!0,action:{condition:{or:{vals:[{isExecutePhase:{threshold:"E20"}},{cmp:{op:"OpLe",lhs:{remainingTime:{}},rhs:{const:{val:"25s"}}}}]}},castSpell:{spellId:{spellId:3045}}}},{action:{castSpell:{spellId:{spellId:77767}}}}]},le={items:[{id:65206,enchant:4209,gems:[68778,52209],reforging:165},{id:69880,randomSuffix:-135,reforging:151},{id:65208,enchant:4204,gems:[52212],reforging:166},{id:69884,randomSuffix:-135,enchant:1099},{id:65204,enchant:4102,gems:[52212,52220]},{id:65028,enchant:4258,gems:[52212]},{id:65205,enchant:3222,gems:[52212,52212]},{id:65132,gems:[52212,52212]},{id:60230,enchant:3823,gems:[52212,52220,52209]},{id:65063,enchant:4105,gems:[52220]},{id:65082},{id:65367,randomSuffix:-133},{id:65140},{id:65026},{id:65139,enchant:4227,reforging:167},{},{id:65058,enchant:4175,reforging:167}]},ne={items:[{id:71503,enchant:4209,gems:[68778,52209],reforging:154},{id:71610,reforging:152},{id:71403,randomSuffix:-294,enchant:4204,gems:[52258],reforging:165},{id:71415,enchant:4100,gems:[52258,52258],reforging:137},{id:71501,enchant:4102,gems:[52212,52209]},{id:71561,enchant:4258,gems:[52212],reforging:152},{id:71502,enchant:4107,gems:[52212,52212],reforging:151},{id:71255,gems:[52212,52212],reforging:152},{id:71504,enchant:4126,gems:[52212,52220],reforging:152},{id:71457,enchant:4105,gems:[52212]},{id:71216,gems:[52212],reforging:152},{id:71401,reforging:152},{id:69150},{id:69112},{id:71466,enchant:4227},{},{id:71611,enchant:4267,reforging:151}]},ie=s("SV PreRaid Preset",{items:[{id:59456,enchant:4209,gems:[68778,59478,59493]},{id:52350,gems:[52212]},{id:64712,enchant:4204,gems:[52212],reforging:152},{id:56315,enchant:1099},{id:56564,enchant:4063,gems:[52258],reforging:152},{id:63479,enchant:4071,gems:[0]},{id:64709,enchant:3222,gems:[52212,0],reforging:137},{id:56539,gems:[52212,52212],reforging:165},{id:56386,enchant:4126,gems:[52258,52258]},{id:62385,enchant:4076,gems:[52212],reforging:166},{id:52348,gems:[52212],reforging:167},{id:62362,reforging:166},{id:68709,reforging:166},{id:56328,reforging:137},{id:55066,enchant:4227,reforging:165},{},{id:59367,enchant:4175,gems:[52212],reforging:151}]}),oe=s("SV T11 Preset",le),pe=s("SV T12 Preset",ne),re=a("SV",ae),de=a("AOE",se),ce=l("SV P1",O.fromMap({[y.StatStamina]:.5,[y.StatAgility]:3.27,[y.StatRangedAttackPower]:1,[y.StatHitRating]:2.16,[y.StatCritRating]:1.17,[y.StatHasteRating]:.89,[y.StatMasteryRating]:.88},{[R.PseudoStatRangedDps]:3.75})),ge=l("SV P3 (T12 4 set)",O.fromMap({[y.StatStamina]:.5,[y.StatAgility]:3.37,[y.StatRangedAttackPower]:1,[y.StatHitRating]:2.56,[y.StatCritRating]:1.27,[y.StatHasteRating]:.94,[y.StatMasteryRating]:1.04},{[R.PseudoStatRangedDps]:4.16})),me={name:"Survival",data:P.create({talentsString:"03-2302-03203203023022121311",glyphs:w.create({prime1:k.GlyphOfExplosiveShot,prime2:k.GlyphOfKillShot,prime3:k.GlyphOfArcaneShot,major1:G.GlyphOfDisengage,major2:G.GlyphOfRaptorStrike,major3:G.GlyphOfTrapLauncher})})},Ie=C.create({classOptions:{useHuntersMark:!0,petType:H.Wolf,petTalents:n,petUptime:1},sniperTrainingUptime:.9}),he=V.create({defaultPotion:W.PotionOfTheTolvir,prepopPotion:W.PotionOfTheTolvir,flask:E.FlaskOfTheWinds,defaultConjured:i.value,food:x.FoodSeafoodFeast,tinkerHands:o.value}),Se={distanceFromTarget:24,duration:240,durationVariation:20,profession1:F.Engineering,profession2:F.Jewelcrafting},ue=p(_.SpecSurvivalHunter,{cssClass:"survival-hunter-sim-ui",cssScheme:b.getCssClass(b.Hunter),knownIssues:[],warnings:[],epStats:[y.StatStamina,y.StatAgility,y.StatRangedAttackPower,y.StatHitRating,y.StatCritRating,y.StatHasteRating,y.StatMasteryRating],epPseudoStats:[R.PseudoStatRangedDps],epReferenceStat:y.StatRangedAttackPower,displayStats:D.createDisplayStatArray([y.StatHealth,y.StatStamina,y.StatAgility,y.StatRangedAttackPower,y.StatMasteryRating],[R.PseudoStatPhysicalHitPercent,R.PseudoStatPhysicalCritPercent,R.PseudoStatRangedHastePercent]),modifyDisplayStats:e=>Y(e),defaults:{gear:pe.gear,epWeights:ge.epWeights,statCaps:(new O).withPseudoStat(R.PseudoStatPhysicalHitPercent,8),other:Se,consumes:he,talents:me.data,specOptions:Ie,raidBuffs:L.create({arcaneBrilliance:!0,bloodlust:!0,markOfTheWild:!0,icyTalons:!0,moonkinForm:!0,leaderOfThePack:!0,powerWordFortitude:!0,strengthOfEarthTotem:!0,trueshotAura:!0,wrathOfAirTotem:!0,demonicPact:!0,blessingOfKings:!0,blessingOfMight:!0,communion:!0}),partyBuffs:J.create({}),individualBuffs:j.create({vampiricTouch:!0}),debuffs:U.create({sunderArmor:!0,curseOfElements:!0,savageCombat:!0,bloodFrenzy:!0})},playerIconInputs:[$()],rotationInputs:te,petConsumeInputs:[],includeBuffDebuffInputs:[r,d,c],excludeBuffDebuffInputs:[],otherInputs:{inputs:[q(),X(),Z(),ee,g,m,I,h,S]},encounterPicker:{showExecuteProportion:!1},presets:{epWeights:[ce,ge],talents:[me],rotations:[re,de],gear:[pe,oe,ie]},autoRotation:e=>e.sim.encounter.targets.length>=4?de.rotation.rotation:re.rotation.rotation,simpleRotation:(e,t,s)=>{const[a,l]=Q(s),n=B.fromJsonString(`{"condition":{"cmp":{"op":"OpGt","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"6s"}}}},"multidot":{"spellId":{"spellId":49001},"maxDots":${t.multiDotSerpentSting?3:1},"maxOverlap":{"const":{"val":"0ms"}}}}`),i=B.fromJsonString('{"condition":{"auraShouldRefresh":{"auraId":{"spellId":3043},"maxOverlap":{"const":{"val":"0ms"}}}},"castSpell":{"spellId":{"spellId":3043}}}'),o=B.fromJsonString('{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":49067}}}}},"castSpell":{"spellId":{"tag":1,"spellId":49067}}}'),p=B.fromJsonString('{"castSpell":{"spellId":{"spellId":58434}}}'),r=B.fromJsonString('{"castSpell":{"spellId":{"spellId":61006}}}'),d=B.fromJsonString('{"castSpell":{"spellId":{"spellId":49050}}}'),c=B.fromJsonString('{"castSpell":{"spellId":{"spellId":49048}}}'),g=B.fromJsonString('{"castSpell":{"spellId":{"spellId":49052}}}'),m=B.fromJsonString('{"castSpell":{"spellId":{"spellId":63672}}}'),I=B.fromJsonString('{"condition":{"not":{"val":{"dotIsActive":{"spellId":{"spellId":60053}}}}},"castSpell":{"spellId":{"spellId":60053}}}'),h=B.fromJsonString('{"condition":{"dotIsActive":{"spellId":{"spellId":60053}}},"castSpell":{"spellId":{"spellId":60052}}}');return t.type==v.Aoe?l.push(...[t.sting==T.ScorpidSting?i:null,t.sting==T.SerpentSting?n:null,t.trapWeave?o:null,p].filter((e=>e))):l.push(...[r,I,t.allowExplosiveShotDownrank?h:null,t.trapWeave?o:null,t.sting==T.ScorpidSting?i:null,t.sting==T.SerpentSting?n:null,m,d,c,g].filter((e=>e))),M.create({prepullActions:a,priorityList:l.map((e=>N.create({action:e})))})},raidSimPresets:[{spec:_.SpecSurvivalHunter,talents:me.data,specOptions:Ie,consumes:he,defaultFactionRaces:{[K.Unknown]:z.RaceUnknown,[K.Alliance]:z.RaceWorgen,[K.Horde]:z.RaceTroll},defaultGear:{[K.Unknown]:{},[K.Alliance]:{1:ie.gear},[K.Horde]:{1:ie.gear}},otherDefaults:Se}]});class fe extends u{constructor(e,t){super(e,t,ue),t.sim.waitForInit().then((()=>{new f(this,{getEPDefaults:e=>e.getGear().getItemSetCount("Lightning-Charged Battlegear")>=4?ce.epWeights:e.getGear().getItemSetCount("Flamewaker's Battlegear")>=4?ge.epWeights:ce.epWeights})}))}}export{fe as S};
