# Spel

Volgens mij is er wel een aardige analogie te maken zonder aan computers
te denken.

Stel je een spel fabrikant voor en een use case waarin een klant een web
formulier kan invullen op een nieuw onderdeel op te sturen.

-   spel onderdelen zit in het domein, speelt ook een rol in productie

-   afhandelen van een verzoek zit in `use cases`

-   de relatie met het post bedrijf dat gaat bezorgen zit in
    infrastructuur.

-   de enveloppen waarin wordt verstuurd worden ingepakt in interfaces.
    We bedenken niet iets zelf om ze in te versturen (boterhamzakjes)
    maar we confirmeren aan de wensen van het post bedrijf: een stevige
    envelop.

    We zorgen er voor dat infrastructuur een stevige envelop kan
    organiseren.

* * * * *

Stel je "Anna" voor. Ze heeft een spel gekocht dat ze met plezier
speelt. Op een moment raakt ze een onderdeel van het spel kwijt.
Gelukkig zit er een kaart bij de gebruiksaanwijzing. Hierop kan je je
naam en adres invullen, de naam van het spel en een beschrijving van het
onderdeel. Dit doet ze en ze stuurt de kaart per post op (adres is
voorbedrukt, antwoordnummer).

> In deze beschrijving ga ik gebruik maken van het principe "je moet nu
> de best mogelijke beslissing maken, later kan deze verbeteren". Dus,
> geen twijfels in de hoofdtekst. Twijfel kan in commentaar, zoals dit
> stuk. Dit moet er toe leiden dat de hoofdtekst zo concreet mogelijk is
> en daardoor vlot op te pakken. (Het *"ik besluit nu"* principe)[^1]

-   De kaart wordt bij "de speelgoedfabriek" afgeleverd in de centrale
    postkamer waar alle post binnenkomt.

-   In de eerste schifting wordt bepaald dat de kaart van Anna door de
    afdeling "Vervanging" moet worden afgehandeld.

    > Er wordt hier alleen maar gekeken naar welke afdeling wat kan met
    > de bedoeling van deze kaart. In dit geval is dat gemakkelijk omdat
    > de kaart een standaard vervangingskaart is. Soms schrijven mensen
    > een brief om te vragen naar een vervangend onderdeel, ook deze
    > worden door "de speelgoedfabriek" netjes afgehandeld.

    > In het verhaal ben je geneigd dat iemand de klus doorgeeft aan een
    > andere persoon. Wellicht is het meer "appropriate" dat de andere
    > persoon *aanbied* de klus voor de eerste te klaren. Wellicht past
    > dat meer in het idee dat een dienst wordt aangeboden, zoals een
    > package ook een dienst aanbiedt. Tegelijkertijd moet de afnemer
    > van een dienst de aanbieder kennen (via een import), vanuit dat
    > perspectief is het weer beter om de beeldspraak te houden zoals
    > die nu is.

-   De kaart komt op de afdeling "Vervanging" binnen. Daar wordt gekeken
    of alle informatie op de kaart staat.

    > In termen van een functie worden alle velden van de kaart naar de
    > argumenten van een functie in de use-case gemapt.

    > Zo nodig kan een eerste controle op de gegevens plaatsvinden.
    > Spelnummers bestaan altijd uit een reeks cijfers.

    > Door nu in de eerste kaart te vragen om een beschrijving in plaats
    > van een onderdeelnummer heb ik nu wel een beetje een probleem.
    > Waar vindt de conversie van beschrijving naar artikelnummer
    > plaats? Ik *besluit nu* dat ik dit de usecase laat doen.

    Met de informatie op de kaart wordt nu een officieel
    vervangingsverzoek gedaan.

    > De buitenste twee cirkels horen niet tot de kern van de
    > applicatie. Zoals de naam al suggereert zit hier de aansluiting
    > naar de "buitenwereld" (inclusief andere software als databases)
    > en de vertaling van "buitenwereld" naar kern van de applicatie in.

    > Met de vertaling naar een "officieel verzoek" zitten we op een
    > bestaande procedure in de kern van de applicatie (use cases in dit
    > geval).

    Tot nu toe was het nog steeds mogelijk geweest dat "de
    speelgoedfabriek" Anna een vriendelijk antwoord had gestuurd met
    daarin de boodschap dat ze niet geholpen kon worden omdat de
    informatie onvolledig is.

-   Er is een moment waarop iemand begint met het wegwerken van de
    binnengekomen officiele vervangingsverzoeken. Ook het verzoek van
    Anna komt aan bod.

    > En toch is dit een vreemd onderscheid. De reden de vertaling van
    > kaart naar officieel verzoek uit elkaar te trekken is dat je de
    > afhandeling van het verzoek niet wilt belasten met het
    > interpreteren van de informatie op de kaart. Op enig moment gaat
    > "de speelgoedfabriek" ook verzoeken op hun website afhandelen. Je
    > wilt dan niet de use case code moeten aanpassen. Het is veel beter
    > een extra controller toe te voegen die de informatie van de
    > webpagina leest en deze vertaalt naar exact dezelfde functie
    > aanroep die ook van de controller afkomt welke de kaart
    > interpreteert. Dit is een voorbeeld waarin bestaande code niet mag
    > worden aangepast.

    > Het mooie van zo'n officieel verzoek is dus dat niet meer is te
    > zien of deze van een kaart afkomt of van de telefooncentrale.

    > Een bijkomend voordeel is dat deze de opzet de mogelijkheid biedt
    > om verzoeken te bundelen en op een later moment een heel aantal
    > tegelijk af te handelen (je ziet mogelijkheden voor "concurrency"
    > ontstaan).

Er moeten in deze fase twee dingen gebeuren:

1.  Op basis van de gegevens wordt het artikelnummer bepaald en uit de
    voorraad opgehaald (aangenomen dat het nog op voorraad is).

2.  Het artikel moet worden ingepakt in een geaddresseerde envelop en
    opgestuurd.

We gaan eerst het artikel uit de voorraad halen.

> Het interessante is dat er ook andere processen zijn waarbij een
> artikel uit een voorraad gehaald moeten kunnen worden. Dit deelproces
> is onderdeel van de vervangingsprocedure maar weet van zichzelf niet
> dat het bij dat proces hoort.

> Conceptueel is het ophalen van het artikel een van de lastigste
> onderdelen uit dit verhaal. Dat komt doordat dit proces op domein
> niveau is gedefinieerd, op use case niveau wordt gebruikt en op
> interface en infrastructuur niveau wordt geïmplementeerd.

> > Formeel is de implementatie op "interfaces" niveau, maar op zijn
> > beurt wordt daar weer infrastructuur code bij uitgevoerd, nl. het
> > daadwerkelijke ophalen uit de database, euh, het magazijn.

-   Degene die het verzoek van Anna afhandelt belt "Domein" met de vraag
    of deze het artikel met genoemde artikelnummer brengt.

    > Waarom hier bellen? Het is van belang te snappen dat in deze use
    > case code een functie van het domein wordt geroepen. Dat is op
    > zich niet anders dan alle voorgaande acties waar in een cirkel een
    > functie op een binnenliggende functie wordt geroepen.

    > Tegelijk zitten we met het conceptueel lastige dat de beller denkt
    > het domein te bellen. Het is dan ook vastgelegd dat dit één van de
    > verantwoordelijkheden is van het domein. Terecht mag je ze met
    > deze vraag lastig vallen.

    > Het lastige is dat het Domein de uitvoering van deze opdracht
    > overlaat aan een functie in "interface". Het is in de code nog
    > erger, het Domein weet officieel niet eens / bepaalt niet *wie*
    > die opdracht uitvoert.

    > Wie de opdracht uitvoert wordt bepaald door de "wiring" in de
    > applicatie. Ergens zit een super afdeling ("main") die op een
    > goede dag heeft besloten dat de afhandeling van deze vraag niet
    > meer gedaan wordt door Piet in zijn grijze stofjas, maar door een
    > aantal magazijnmedewerkers die een daartoe bestemde opleiding met
    > succes hebben gevolgd en werken in het nieuwe magazijn dat door
    > een professioneel bedrijf is opgezet en ingericht.

    Al snel komt iemand met het artikel binnen. Voordat hij weer wil
    vertrekken kunnen we hem nog net even een kop koffie aanbieden en we
    vragen geïnteresseerd hoe hij weet dat hij dit artikel moest zoeken
    en hier brengen.

    Dat verhaal hoor je later.

Het tweede gedeelte bestaat uit het opsturen van het artikel.

> Een vraag die hier naar boven komt is of het adres van Anna in het
> domein moet worden opgslagen. Nou gaat dit verhaal niet echt over de
> vraag of een speelgoedfabriek adressen moet opslaan. Ik *beslis nu*
> dat in "de speelgoedfabriek" het gaat om goede spellen alleen en dat
> er op domein niveau geen interesse is in adressen van klanten.

> Desalnietemin kan het zelfs in "de speelgoedfabriek" zo zijn dat het
> adres van Anna (tijdelijk) moet worden opgelagen. Dit gebeurt ook in
> het clean architecture verhaal: "users" zijn daar in de use case code
> gedefinieerd. Interfaces heeft er op zijn beurt weer geen weet van dat
> "customers" domein zijn en "users" "use case".

> Volgens mij stuit ik nu op een verschil tussen de fabriekanalogie en
> de echte implementatie: in de fabriekanalogie raakt de usecase nu de
> controle kwijt door het onderdeel in combinatie met adres naar de
> interface te sturen die de interne representatie omzet naar de externe
> representatie: een brief geadresseerd volgens de regels en voldoende
> gefrankeerd. Deze brief gaat vervolgens naar de postkamer om verstuurd
> te worden (de analogie voor infrastructuur).

> Ik twijfel nog een beetje over de implementatie in de applicatie. In
> een synchrone applicatie krijgt de oorspronkelijke controller wellicht
> weer de controle en retouneert het resultaat. Dit is vermoedelijk niet
> hoe het echt gaat. Zelfs in een synchrone applicatie is er geen sprake
> van een directe return. Net als in lht zet de invoer iets in werking
> waar even later een bericht komt op basis waarvan de view update en
> zich aanpast aan de nieuwe situatie. De front-end is stateless en laat
> als zodanig niet de return waarde zien maar synchroniseert met de
> model representatie.

Degene die het verzoek van Anna afhandelt geeft de postkamer de opdracht
het onderdeel naar Anna te sturen.

-   In de postkamer wordt gezorgd voor correcte adressering en
    frankering en de post wordt naar de ruimte gebracht waar alle
    uitgaande post wordt verzameld om door de ptt te worden opgehaald.

-   Een dag later ontvangt Anna het onderdeel en ze kan haar spel weer
    spelen.

Je hebt het verhaal van de magazijn medewerker nog tegoed.

> De essentie van dit verhaal is de organisatie van kern-processen
> zonder dat de manier waarop voor iedereen bekend moet zijn (het gaat
> niet zozeer om bekend zijn, het gaat nog er nog meer om dat bij een
> verandering van magazijnsysteem zo goed als niemand mee hoeft te
> veranderen). In het verhaal zou je kunnen zeggen dan er ooit een
> moment was waarop iedereen maar een beetje deed. Behalve dat er iets
> was als een soort magazijn (een chaos) hielden mensen ook hun eigen
> voorraadje bij. Het gaat er juist om dat er één georganiseerde
> voorraad is waarvan iedereen gebruik kan maken en waardoor de voorraad
> te controleren is. Uiteindelijk is de belangrijkste reden nog wel dat
> alle gebruiker *geen weet mogen (nee, **hoeven**) hebben **hoe** de
> artikelen in het magazijn worden opgeslagen*. Het magazijn moet
> vervangbaar zijn door een ander systeem zonder dat iedereen ingelicht
> moeten worden. Een verandering van magazijn systeem moet maar op één
> plek tot een verandering leiden. In termen van het gebouw, je moet
> niet een compleet andere bedrading door het hele gebouw hebben wanneer
> het magazijn systeem wijzigt. (Een aanleg van een buizensysteem?)

> Nog een andere gedachte die de moeite waard is om na te gaan.
> Verschillende ringen bevatten meerdere functies die eigenlijk niets
> met elkaar van doen hebben, behalve dat ze in dezelfde cirkel zitten.
> Zo hebben verschillende use cases niets met elkaar van doen, behalve
> dat ze van hetzelfde domein gebruik maken. Is het nu ook zo dat deze
> verschillende afdelingen binnen één cirkel **niet** met elkaar mogen
> communiceren? Dus, dat wanneer degen die het artikel nodig heeft voor
> Anna's verzoek alleen via een binnengelegen cirkel uiteindelijk blijkt
> te communiceren met de afdeling die eigenlijk binnen de cirkel een
> buur had kunnen zijn?

-   Het telefonisch verzoek kwam op een andere plek binnen dan dat de
    beller dacht. Het kwam bij de magazijnmedewerkers binnen.

-   De magazijnmedewerker heeft het nummer ingevoerd in het
    administratie systeem van het magazijn. Het automatische
    magazijnsysteem pikte het onderdeel op uit het nieuwe onbemande
    systeem.

    > Vroeger deed iedereen maar wat in het magazijn. Mensen liepen in
    > en uit. Dingen werden zonder te melden meegenomen en andere
    > rotzooi werd "tijdelijk" in het magazijn gestald. Niemand had meer
    > overzicht wat er nu precies was en wat aangevuld moest worden.

    Nadat het artikel was afgeleverd bij de magazijnmederwerker heeft
    deze het meteen naar de bellende medewerker gebracht.

* * * * *

We willen weten hoe software in elkaar gezet kan worden. We kunnen
daarvoor software bestuderen maar in de praktijk blijkt dat het best wel
eens lastig is om te begrijpen waarom de verschillende onderdelen
samenwerken op de manier waarop het is gedaan. Ik heb een ander
voorstel.

We gaan een kijkje nemen in "De Speelgoedfabriek". De Speelgoedfabriek
is een bedrijf waarin mensen werken met een passie voor leuke spellen,
gewoon, bordspellen.

Het is met een paar mensen in een kleine ruimte begonnen en alles werd
zelf gedaan: ontwerp, productie en verkoop. Doordat ze nogal gegroeid
zijn, zijn ze in de afgelopen jaren tegen allerlei problemen aangelopen.
Langzamerhand zijn de meeste van die problemen één voor één aangepakt en
onlangs is De Speelgoedfabriek verhuisd naar een compleet nieuw pand dat
is ingericht op een manier die past bij de manier van werken die is
gevonden bij het oplossen van al die problemen.

Voor we De Speelgoedfabriek binnengaan gaan we eerst kennismaken met
Anna. Anna is fan van de spellen van De Speelgoedfabriek vanaf het
allereerste uur.

Anna heeft pas een spel gekocht dat ze met plezier speelt. Vorige week
is ze een onderdeel van het spel kwijtgeraakt. Gelukkig zit er een kaart
bij de gebruiksaanwijzing. Hierop kan je je naam en adres invullen, de
naam van het spel en een beschrijving van het onderdeel. Dit doet ze en
ze stuurt de kaart per post op.

Wij gaan De Speelgoedfabriek volgen in het proces waarin voor Anna een
nieuw onderdeel wordt gezorgd.

Oke, op een goede dag doet Anna de kaart op de bus.

's Avonds leegt de post de brievenbus en de dag erna bezorgd een
postbode de post bij De Speelgoedfabriek, waaronder de kaart van Anna.

Bob werkt op de postkamer van De Speelgoedfabriek en doet een eerste
schifting van de binnengekomen post. De kaart van Anna komt op het
stapeltje met andere vervangingsverzoeken.

Het is Carl (van Klantenservice?) die de vervangingsverzoeken afhandelt.
Carl controleert of alle informatie leesbaar op het verzoek staat. Ook
doet hij een eerste controle. Spelnummers zijn altijd een getal tussen
100 en 1000. Soms is wel iets ingevuld waarvan onmogelijk is af te
leiden welk spel bedoeld wordt.

Carl neemt ook wel eens de telefoon aan. Sommige klanten weten het
telefoonnummer van De Speelgoedfabriek te achterhalen en bellen met het
verzoek om een vervangend onderdeel. Bij De Speelgoedfabriek is de klant
koning en zo mogelijk wordt iedereen geholpen.

Vroeger was het wel eens zo dat degene die de de verzoeken uiteindelijk
afhandelt (dat is niet Carl, vandaag is dat de taak van Dani, maar dat
komt zo) een ratjetoe in de handen gedrukt kreeg met handgeschreven
brieven, onleesbare standaardkaartjes en onvolledige kladjes van iemand
die de telefoon aannam terwijl deze eigenlijk met wat anders bezig was.
Er ging veel te veel tijd verloren en er werden teveel fouten gemaakt.
Medewerkers van De Speelgoedfabriek werden ook wel eens boos omdat ze
eigenlijk druk waren met wat anders.

Precies om dat te voorkomen is nu besloten dat slechts een paar mensen
verzoeken in behandeling nemen en deze leesbaar en volledig op zo'n rood
standaard kaartje zetten. Wanneer niet alle informatie er is, dan wordt
zo mogelijk een vriendelijk kaartje naar de klant gestuurd met daarin de
vraag nogmaals een verzoek te doen met de volledige informatie. Voor het
verzoek van Anna is dat niet nodig.

Eenmaal per dag handelt één persoon alle rode standaardkaartjes af. Zo
komt Dani (daar is ze dan) het rode kaartje tegen met het verzoek van
Anna in het nette handschrift van Carl.

De taak van Dani bestaat uit twee onderdelen. Eerst moet ze het missende
onderdeel vinden en vervolgens in een brief of doosje te doen om deze
naar Anna te kunnen sturen.

Dat zoeken van het onderdeel was vroeger echt een probleem. Er was een
ruimte waar onderdelen verzameld werden. Iedereen liep in en uit en soms
zette medewerkers de oude rommel waar ze zelf van af wilden in dat
magazijn. Al met al kwam het veel te vaak voor dat er naar onderdelen
gezocht werd die ergens in een hoek lagen maar niet gevonden werd. Het
liep zelfs zo uit de hand dat sommige medewerkers een eigen voorraadje
aanlegden om zo klanten te kunnen helpen. Dat leidde er dan weer toe dat
onderdelen wel in huis waren maar dat ze lang niet altijd gevonden
werden of pas na lang zoeken als ze nodig waren.

Om al deze problemen te voorkomen, is het beheer van de voorraad nu
anders aangepast. Sterker nog, er wordt al nagedacht over een nog groter
en efficiënter magazijn. Maar één ding staat vast. Het is belangrijk dat
er één duidelijke procedure is waarmee mensen een artikel kunnen
opvragen en dat die procedure niet verandert op moment dat met een ander
type magazijn gewerkt gaat worden.

> Hier komt Emma

Er is daarom het volgende gedaan. Spelonderdelen zijn uiteraard cruciaal
in De Speelgoedfabriek. Er gebeurt hier veel meer dan alleen maar
vervangende onderdelen opsturen. Een medewerker die vandaag de
vervangingsverzoeken afhandelt moet op dezelfde manier aan onderdelen
komen als de dag erna als deze bestellingen van nieuwe spellen afhandelt
(of werkt aan de vertaling van speluitleg).

Het ophalen van spelonderdelen is volledig gestandaardiseerd. Ieder
onderdeel heeft een nummer die in een catalogus is op te zoeken. Overal
waar nodig zijn terminals waar zo'n onderdeelnummer kan worden
ingevoerd. Wanneer je dat doet wordt even later het artikel gebracht.

> Je kan in dit verhaal Emma introduceren die in de MT van De
> Speelgoedfabriek werkt. Zij heeft de procedure bepaald waarop
> artikelen kunnen worden opgehaald. Dat geeft meer "smoel" aan het
> domein.

> > Ed en Fred worden dan Fiona en Geoff

Dani voert het artikelnummer van het artikel dat Anna vervangen wil
hebben in in zo'n terminal. Dit verzoek komt meteen binnen bij Ed. Ed
heeft de beschikking over het magazijnsysteem en zij zoek meteen uit
waar in het magazijn het artikel te vinden is.

Ed laat het artikel door Fred ophalen (Fred heeft een uitzendbaantje en
haalt artikelen uit het magazijn). Fred heeft van Ed het nummer van de
stellingkast gekregen waar volgens het systeem nog meer dan genoeg van
die onderdelen liggen. Nadat Fred Ed het onderdeel heeft gegeven brengt
Ed het naar Dani. Ed hoopt even koffie met Dani te kunnen drinken, dat
is altijd gezellig.

> Een aantal nachtelijke gedachten

> Het gaat er in deze analogie niet zozeer om wat de juiste oplossing
> is, het gaat erom dat er een context is waarin een aantal mogelijke
> oplossingen voor een probleem vergeleken kunnen worden en zo voor- en
> nadelen tegen elkaar kunnen worden afgewogen. Het gaat dus om een
> kader te scheppen waardoor je een zo goed mogelijke oplossing
> ontwikkeld.

> Je kan er ook zo tegenaankijken. Het gaat er om dat je problemen
> oplost door duidelijke rollen te definieren. Door te voorkomen dat je
> onduidelijkheid krijgt door overlap van verantwoordelijkheden. Door er
> voor te zorgen dat problemen tussen wal en schip vallen doordat
> niemand de verantwoordelijkheid heeft of neemt er iets aan te doen.

> De hele opzet van De Speelgoedfabriek is om snel aan te kunnen passen
> aan wisselenden omstandigheden. En je krijgt snelheid door in geval
> van een probleem op zo weinig mogelijk plekken een aanpassing te
> moeten doen.

> Alarm! Beveiliging nodig. Kan dat nu in de structuur worden ingepast
> zonder alle bestaande componenten aan te moeten passen (oké, één
> component). Ik ben geneigd na te gaan denken over de vraag of de toe
> te voegen beveiliging in "infrastructuur" komt, of in "interfaces" of
> misschien toch "use cases" (zoals in clean architecture verhaal). Het
> antwoord is denk ik niet zondermeer "dit probleem los je altijd op
> deze manier op". Vanuit het idee dat je zo weinig mogelijk moet
> veranderen zou je een compleet nieuwe use case kunnen *toevoegen* (dat
> is heel wat anders als *veranderen* van bestaande rollen). We zouden
> Bob de taak kunnen geven om het proces te starten of Anna om een nieuw
> onderdeel mag vragen. Misschien is het beter om het Carl te laten
> doen. Misschien komen we wel tot de conclusie dat er een fundamentele
> fout zit in de manier waarop Dani haar werk doet en moeten we daar
> aanpassen.

> Dus, nogmaals, de architectuur is niet alleen het antwoord, het is
> misschien nog veel meer het proces dat leidt tot een antwoord.

> In deze analogie loopt het adres van Anna mee. Dat lijkt me overbodig.
> Je kan de individuele rollen of processen nog atomairder maken door
> Bob de verantwoordelijkheid te geven de communicatie met de
> buitenwereld te verzorgen.

> Op deze manier scheidt Bob het adres van Anna en de taak die moet
> worden gedaan. Om precies te zijn:

> 1.  Bob zorgt dat Carl een vervangingsverzoek kan gaan uitvoeren.

> 2.  Het adres van Anna is geen onderdeel van dat verzoek, Bob zorgt er
>     voor dat wanneer er op een moment een gevulde envelop voor Anna
>     komt dat haar adres er weer op komt.

> 3.  Er zit nog een interessant aspect aan die envelop. Bob weet hoe je
>     iets goed en voordelig kan versturen. Carl heeft op een moment
>     iets om te versturen. Bob zal uit zichzelf nooit iets versturen,
>     dat is niet zijn taak. Carl moet wel iets versturen maar wil zich
>     niet druk maken over hoe dit zo efficiënt mogelijk kan.

> We hebben hier dezelfde situatie als met het magazijn. In een van de
> binnencirkels is besloten dat onderdelen en hele spellen verzonden
> moeten kunnen worden. Er is een structuur opgezet waardoor mensen in
> de binnencirkel op een uniforme wijze items kunnen verzenden zonder
> dat deze weet hoeven te hebben over wat de beste manier is.

> Om dit op te lossen is het mogelijk dat Carl uiteindelijk het
> onderdeel doet in een envelop die hij ook van de postkamer krijgt
> (niet per se Bob) en dat het kaartje van Bob aan deze envelop
> bevestigd wordt.

> Waar ik geloof ik heen wil is dat er een interface is die door Bob
> wordt geïmplementeerd. In deze analogie zorgt Bob er dan voor dat Carl
> de juiste type envelop krijgt om het onderdeel van Anna in te doen.

> Dit systeem is helemaal niet handig wanneer een medewerker uit de
> binnencirkel een brief wil sturen aan een of andere instantie. Nou is
> dit wel een andere use case, maar het is toch wel aardig wanneer
> dezelfde rol verdeling kan gebruiken.

> Een oplossing hiervoor zou kunnen zijn dat medewerkers die brieven
> sturen deze nooit zelf hoeven te printen. Eenmaal ingevoerd zou de
> postkamer voor printen, de envelop, addresering en frankering kunnen
> zorgen.

> Je ziet bijna als vanzelf het kanban systeem in deze organisatie
> ontstaan. Op moment dat je voor iedere overdracht een kaartje hebt is
> dat feitelijk ook zo. Als vanzelf krijg je ook een concurrent
> oplossing lijkt het wel waarin kaartjes over channels verstuurd kunnen
> worden. Voeg nog een beetje Kaizen toe en je bent lekker bezig.

* * * * *

[^1]: Eigenlijk zijn er twee redenen voor een dergelijk commentaar. Ik
    gebruik het nu voor mijn eigen "definitieve tekst" te
    becommentarieren, maar het kan net zo goed gebruikt worden in een
    definitieve tekst om onderscheid te maken tussen kern tekst en
    context.
