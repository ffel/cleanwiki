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
> en daardoor vlot op te pakken. (Het *"ik besluit nu"* principe)

-   De kaart wordt bij "de speelgoedfabriek" afgeleverd in de centrale
    postkamer waar alle post binnenkomt.

-   In de eerste schifting wordt bepaald dat de kaart van Anna door de
    afdeling "Vervanging" moet worden afgehandeld.

    > Er wordt hier alleen maar gekeken naar welke afdeling wat kan met
    > de bedoeling van deze kaart. In dit geval is dat gemakkelijk omdat
    > de kaart een standaard vervangingskaart is. Soms schrijven mensen
    > een brief om te vragen naar een vervangend onderdeel, ook deze
    > worden door "de speelgoedfabriek" netjes afgehandeld.

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

> Waarom hier bellen? Het is van belang te snappen dat in deze use case
> code een functie van het domein wordt geroepen. Dat is op zich niet
> anders dan alle voorgaande acties waar in een cirkel een functie op
> een binnenliggende functie wordt geroepen.

> Tegelijk zitten we met het conceptueel lastige dat de beller denkt het
> domein te bellen. Het is dan ook vastgelegd dat dit één van de
> verantwoordelijkheden is van het domein. Terecht mag je ze met deze
> vraag lastig vallen.

> Het lastige is dat het Domein de uitvoering van deze opdracht overlaat
> aan een functie in "interface". Het is in de code nog erger, het
> Domein weet officieel niet eens / bepaalt niet *wie* die opdracht
> uitvoert.

> Wie de opdracht uitvoert wordt bepaald door de "wiring" in de
> applicatie. Ergens zit een super afdeling ("main") die op een goede
> dag heeft besloten dat de afhandeling van deze vraag niet meer gedaan
> wordt door Piet in zijn grijze stofjas, maar door een aantal
> magazijnmedewerkers die een daartoe bestemde opleiding met succes
> hebben gevolgd en werken in het nieuwe magazijn dat door een
> professioneel bedrijf is opgezet en ingericht.

    Al snel komt iemand met het artikel binnen.  Voordat hij weer wil vertrekken
    kunnen we hem nog net even een kop koffie aanbieden en we vragen geïnteresseerd
    hoe hij weet dat hij dit artikel moest zoeken en hier brengen.

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
