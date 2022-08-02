package gambas

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestSeriesPrint(t *testing.T) {
	type seriesPrintTest struct {
		arg1 Series
	}
	seriesPrintTests := []seriesPrintTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a", "red"}},
						{1, []interface{}{"b", "yellow"}},
						{2, []interface{}{"c", "red"}},
					},
					[]string{"ID", "Color"},
				},
				"Fruit",
				"string",
			},
		},
	}

	for _, test := range seriesPrintTests {
		test.arg1.Print()
	}
}

func TestSeriesPrintRange(t *testing.T) {
	type printRangeTest struct {
		arg1 Series
	}
	printRangeTests := []printRangeTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a", "red"}},
						{1, []interface{}{"b", "yellow"}},
						{2, []interface{}{"c", "red"}},
					},
					[]string{"ID", "Color"},
				},
				"Fruit",
				"string",
			},
		},
	}

	for _, test := range printRangeTests {
		test.arg1.PrintRange(1, 3)
	}
}

func TestSeriesHead(t *testing.T) {
	type headTest struct {
		arg1 Series
	}
	headTests := []headTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{"Index"},
				},
				"Fruit",
				"string",
			},
		},
	}

	for _, test := range headTests {
		test.arg1.Head(2)
	}
}

func TestSeriesTail(t *testing.T) {
	type tailTest struct {
		arg1 Series
	}
	tailTests := []tailTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{"RangeIndex"},
				},
				"People",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{"Index"},
				},
				"Fruit",
				"string",
			},
		},
	}

	for _, test := range tailTests {
		test.arg1.Tail(2)
	}
}

func BenchmarkSeriesAt(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	names := [][]interface{}{
		{"Avery Bradley"},
		{"Jae Crowder"},
		{"John Holland"},
		{"R.J. Hunter"},
		{"Jonas Jerebko"},
		{"Amir Johnson"},
		{"Jordan Mickey"},
		{"Kelly Olynyk"},
		{"Terry Rozier"},
		{"Marcus Smart"},
		{"Jared Sullinger"},
		{"Isaiah Thomas"},
		{"Evan Turner"},
		{"James Young"},
		{"Tyler Zeller"},
		{"Bojan Bogdanovic"},
		{"Markel Brown"},
		{"Wayne Ellington"},
		{"Rondae Hollis-Jefferson"},
		{"Jarrett Jack"},
		{"Sergey Karasev"},
		{"Sean Kilpatrick"},
		{"Shane Larkin"},
		{"Brook Lopez"},
		{"Chris McCullough"},
		{"Willie Reed"},
		{"Thomas Robinson"},
		{"Henry Sims"},
		{"Donald Sloan"},
		{"Thaddeus Young"},
		{"Arron Afflalo"},
		{"Lou Amundson"},
		{"Thanasis Antetokounmpo"},
		{"Carmelo Anthony"},
		{"Jose Calderon"},
		{"Cleanthony Early"},
		{"Langston Galloway"},
		{"Jerian Grant"},
		{"Robin Lopez"},
		{"Kyle O'Quinn"},
		{"Kristaps Porzingis"},
		{"Kevin Seraphin"},
		{"Lance Thomas"},
		{"Sasha Vujacic"},
		{"Derrick Williams"},
		{"Tony Wroten"},
		{"Elton Brand"},
		{"Isaiah Canaan"},
		{"Robert Covington"},
		{"Joel Embiid"},
		{"Jerami Grant"},
		{"Richaun Holmes"},
		{"Carl Landry"},
		{"Kendall Marshall"},
		{"T.J. McConnell"},
		{"Nerlens Noel"},
		{"Jahlil Okafor"},
		{"Ish Smith"},
		{"Nik Stauskas"},
		{"Hollis Thompson"},
		{"Christian Wood"},
		{"Bismack Biyombo"},
		{"Bruno Caboclo"},
		{"DeMarre Carroll"},
		{"DeMar DeRozan"},
		{"James Johnson"},
		{"Cory Joseph"},
		{"Kyle Lowry"},
		{"Lucas Nogueira"},
		{"Patrick Patterson"},
		{"Norman Powell"},
		{"Terrence Ross"},
		{"Luis Scola"},
		{"Jason Thompson"},
		{"Jonas Valanciunas"},
		{"Delon Wright"},
		{"Leandro Barbosa"},
		{"Harrison Barnes"},
		{"Andrew Bogut"},
		{"Ian Clark"},
		{"Stephen Curry"},
		{"Festus Ezeli"},
		{"Draymond Green"},
		{"Andre Iguodala"},
		{"Shaun Livingston"},
		{"Kevon Looney"},
		{"James Michael McAdoo"},
		{"Brandon Rush"},
		{"Marreese Speights"},
		{"Klay Thompson"},
		{"Anderson Varejao"},
		{"Cole Aldrich"},
		{"Jeff Ayres"},
		{"Jamal Crawford"},
		{"Branden Dawson"},
		{"Jeff Green"},
		{"Blake Griffin"},
		{"Wesley Johnson"},
		{"DeAndre Jordan"},
		{"Luc Richard Mbah a Moute"},
		{"Chris Paul"},
		{"Paul Pierce"},
		{"Pablo Prigioni"},
		{"JJ Redick"},
		{"Austin Rivers"},
		{"C.J. Wilcox"},
		{"Brandon Bass"},
		{"Tarik Black"},
		{"Anthony Brown"},
		{"Kobe Bryant"},
		{"Jordan Clarkson"},
		{"Roy Hibbert"},
		{"Marcelo Huertas"},
		{"Ryan Kelly"},
		{"Larry Nance Jr."},
		{"Julius Randle"},
		{"D'Angelo Russell"},
		{"Robert Sacre"},
		{"Louis Williams"},
		{"Metta World Peace"},
		{"Nick Young"},
		{"Eric Bledsoe"},
		{"Devin Booker"},
		{"Chase Budinger"},
		{"Tyson Chandler"},
		{"Archie Goodwin"},
		{"John Jenkins"},
		{"Brandon Knight"},
		{"Alex Len"},
		{"Jon Leuer"},
		{"Phil Pressey"},
		{"Ronnie Price"},
		{"Mirza Teletovic"},
		{"P.J. Tucker"},
		{"T.J. Warren"},
		{"Alan Williams"},
		{"Quincy Acy"},
		{"James Anderson"},
		{"Marco Belinelli"},
		{"Caron Butler"},
		{"Omri Casspi"},
		{"Willie Cauley-Stein"},
		{"Darren Collison"},
		{"DeMarcus Cousins"},
		{"Seth Curry"},
		{"Duje Dukan"},
		{"Rudy Gay"},
		{"Kosta Koufos"},
		{"Ben McLemore"},
		{"Eric Moreland"},
		{"Rajon Rondo"},
		{"Cameron Bairstow"},
		{"Aaron Brooks"},
		{"Jimmy Butler"},
		{"Mike Dunleavy"},
		{"Cristiano Felicio"},
		{"Pau Gasol"},
		{"Taj Gibson"},
		{"Justin Holiday"},
		{"Doug McDermott"},
		{"Nikola Mirotic"},
		{"E'Twaun Moore"},
		{"Joakim Noah"},
		{"Bobby Portis"},
		{"Derrick Rose"},
		{"Tony Snell"},
		{"Matthew Dellavedova"},
		{"Channing Frye"},
		{"Kyrie Irving"},
		{"LeBron James"},
		{"Richard Jefferson"},
		{"Dahntay Jones"},
		{"James Jones"},
		{"Sasha Kaun"},
		{"Kevin Love"},
		{"Jordan McRae"},
		{"Timofey Mozgov"},
		{"Iman Shumpert"},
		{"J.R. Smith"},
		{"Tristan Thompson"},
		{"Mo Williams"},
		{"Joel Anthony"},
		{"Aron Baynes"},
		{"Steve Blake"},
		{"Lorenzo Brown"},
		{"Reggie Bullock"},
		{"Kentavious Caldwell-Pope"},
		{"Spencer Dinwiddie"},
		{"Andre Drummond"},
		{"Tobias Harris"},
		{"Darrun Hilliard"},
		{"Reggie Jackson"},
		{"Stanley Johnson"},
		{"Jodie Meeks"},
		{"Marcus Morris"},
		{"Anthony Tolliver"},
		{"Lavoy Allen"},
		{"Rakeem Christmas"},
		{"Monta Ellis"},
		{"Paul George"},
		{"George Hill"},
		{"Jordan Hill"},
		{"Solomon Hill"},
		{"Ty Lawson"},
		{"Ian Mahinmi"},
		{"C.J. Miles"},
		{"Glenn Robinson III"},
		{"Rodney Stuckey"},
		{"Myles Turner"},
		{"Shayne Whittington"},
		{"Joe Young"},
		{"Giannis Antetokounmpo"},
		{"Jerryd Bayless"},
		{"Michael Carter-Williams"},
		{"Jared Cunningham"},
		{"Tyler Ennis"},
		{"John Henson"},
		{"Damien Inglis"},
		{"O.J. Mayo"},
		{"Khris Middleton"},
		{"Greg Monroe"},
		{"Steve Novak"},
		{"Johnny O'Bryant III"},
		{"Jabari Parker"},
		{"Miles Plumlee"},
		{"Greivis Vasquez"},
		{"Rashad Vaughn"},
		{"Justin Anderson"},
		{"J.J. Barea"},
		{"Jeremy Evans"},
		{"Raymond Felton"},
		{"Devin Harris"},
		{"David Lee"},
		{"Wesley Matthews"},
		{"JaVale McGee"},
		{"Salah Mejri"},
		{"Dirk Nowitzki"},
		{"Zaza Pachulia"},
		{"Chandler Parsons"},
		{"Dwight Powell"},
		{"Charlie Villanueva"},
		{"Deron Williams"},
		{"Trevor Ariza"},
		{"Michael Beasley"},
		{"Patrick Beverley"},
		{"Corey Brewer"},
		{"Clint Capela"},
		{"Sam Dekker"},
		{"Andrew Goudelock"},
		{"James Harden"},
		{"Montrezl Harrell"},
		{"Dwight Howard"},
		{"Terrence Jones"},
		{"K.J. McDaniels"},
		{"Donatas Motiejunas"},
		{"Josh Smith"},
		{"Jason Terry"},
		{"Jordan Adams"},
		{"Tony Allen"},
		{"Chris Andersen"},
		{"Matt Barnes"},
		{"Vince Carter"},
		{"Mike Conley"},
		{"Bryce Cotton"},
		{"Jordan Farmar"},
		{"Marc Gasol"},
		{"JaMychal Green"},
		{"P.J. Hairston"},
		{"Jarell Martin"},
		{"Ray McCallum"},
		{"Xavier Munford"},
		{"Zach Randolph"},
		{"Lance Stephenson"},
		{"Alex Stepheson"},
		{"Brandan Wright"},
		{"Alexis Ajinca"},
		{"Ryan Anderson"},
		{"Omer Asik"},
		{"Luke Babbitt"},
		{"Norris Cole"},
		{"Dante Cunningham"},
		{"Anthony Davis"},
		{"Bryce Dejean-Jones"},
		{"Toney Douglas"},
		{"James Ennis"},
		{"Tyreke Evans"},
		{"Tim Frazier"},
		{"Alonzo Gee"},
		{"Eric Gordon"},
		{"Jordan Hamilton"},
		{"Jrue Holiday"},
		{"Orlando Johnson"},
		{"Kendrick Perkins"},
		{"Quincy Pondexter"},
		{"LaMarcus Aldridge"},
		{"Kyle Anderson"},
		{"Matt Bonner"},
		{"Boris Diaw"},
		{"Tim Duncan"},
		{"Manu Ginobili"},
		{"Danny Green"},
		{"Kawhi Leonard"},
		{"Boban Marjanovic"},
		{"Kevin Martin"},
		{"Andre Miller"},
		{"Patty Mills"},
		{"Tony Parker"},
		{"Jonathon Simmons"},
		{"David West"},
		{"Kent Bazemore"},
		{"Tim Hardaway Jr."},
		{"Kirk Hinrich"},
		{"Al Horford"},
		{"Kris Humphries"},
		{"Kyle Korver"},
		{"Paul Millsap"},
		{"Mike Muscala"},
		{"Lamar Patterson"},
		{"Dennis Schroder"},
		{"Mike Scott"},
		{"Thabo Sefolosha"},
		{"Tiago Splitter"},
		{"Walter Tavares"},
		{"Jeff Teague"},
		{"Nicolas Batum"},
		{"Troy Daniels"},
		{"Jorge Gutierrez"},
		{"Tyler Hansbrough"},
		{"Aaron Harrison"},
		{"Spencer Hawes"},
		{"Al Jefferson"},
		{"Frank Kaminsky III"},
		{"Michael Kidd-Gilchrist"},
		{"Jeremy Lamb"},
		{"Courtney Lee"},
		{"Jeremy Lin"},
		{"Kemba Walker"},
		{"Marvin Williams"},
		{"Cody Zeller"},
		{"Chris Bosh"},
		{"Luol Deng"},
		{"Goran Dragic"},
		{"Gerald Green"},
		{"Udonis Haslem"},
		{"Joe Johnson"},
		{"Tyler Johnson"},
		{"Josh McRoberts"},
		{"Josh Richardson"},
		{"Amar'e Stoudemire"},
		{"Dwyane Wade"},
		{"Briante Weber"},
		{"Hassan Whiteside"},
		{"Justise Winslow"},
		{"Dorell Wright"},
		{"Dewayne Dedmon"},
		{"Evan Fournier"},
		{"Aaron Gordon"},
		{"Mario Hezonja"},
		{"Ersan Ilyasova"},
		{"Brandon Jennings"},
		{"Devyn Marble"},
		{"Shabazz Napier"},
		{"Andrew Nicholson"},
		{"Victor Oladipo"},
		{"Elfrid Payton"},
		{"Jason Smith"},
		{"Nikola Vucevic"},
		{"C.J. Watson"},
		{"Alan Anderson"},
		{"Bradley Beal"},
		{"Jared Dudley"},
		{"Jarell Eddie"},
		{"Drew Gooden"},
		{"Marcin Gortat"},
		{"JJ Hickson"},
		{"Nene Hilario"},
		{"Markieff Morris"},
		{"Kelly Oubre Jr."},
		{"Otto Porter Jr."},
		{"Ramon Sessions"},
		{"Garrett Temple"},
		{"Marcus Thornton"},
		{"John Wall"},
		{"Darrell Arthur"},
		{"D.J. Augustin"},
		{"Will Barton"},
		{"Wilson Chandler"},
		{"Kenneth Faried"},
		{"Danilo Gallinari"},
		{"Gary Harris"},
		{"Nikola Jokic"},
		{"Joffrey Lauvergne"},
		{"Mike Miller"},
		{"Emmanuel Mudiay"},
		{"Jameer Nelson"},
		{"Jusuf Nurkic"},
		{"JaKarr Sampson"},
		{"Axel Toupane"},
		{"Nemanja Bjelica"},
		{"Gorgui Dieng"},
		{"Kevin Garnett"},
		{"Tyus Jones"},
		{"Zach LaVine"},
		{"Shabazz Muhammad"},
		{"Adreian Payne"},
		{"Nikola Pekovic"},
		{"Tayshaun Prince"},
		{"Ricky Rubio"},
		{"Damjan Rudez"},
		{"Greg Smith"},
		{"Karl-Anthony Towns"},
		{"Andrew Wiggins"},
		{"Steven Adams"},
		{"Nick Collison"},
		{"Kevin Durant"},
		{"Randy Foye"},
		{"Josh Huestis"},
		{"Serge Ibaka"},
		{"Enes Kanter"},
		{"Mitch McGary"},
		{"Nazr Mohammed"},
		{"Anthony Morrow"},
		{"Cameron Payne"},
		{"Andre Roberson"},
		{"Kyle Singler"},
		{"Dion Waiters"},
		{"Russell Westbrook"},
		{"Cliff Alexander"},
		{"Al-Farouq Aminu"},
		{"Pat Connaughton"},
		{"Allen Crabbe"},
		{"Ed Davis"},
		{"Maurice Harkless"},
		{"Gerald Henderson"},
		{"Chris Kaman"},
		{"Meyers Leonard"},
		{"Damian Lillard"},
		{"C.J. McCollum"},
		{"Luis Montero"},
		{"Mason Plumlee"},
		{"Brian Roberts"},
		{"Noah Vonleh"},
		{"Trevor Booker"},
		{"Trey Burke"},
		{"Alec Burks"},
		{"Dante Exum"},
		{"Derrick Favors"},
		{"Rudy Gobert"},
		{"Gordon Hayward"},
		{"Rodney Hood"},
		{"Joe Ingles"},
		{"Chris Johnson"},
		{"Trey Lyles"},
		{"Shelvin Mack"},
		{"Raul Neto"},
		{"Tibor Pleiss"},
		{"Jeff Withey"},
	}
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].At(names[rand.Intn(len(names))])
	}
}

func TestSeriesAt(t *testing.T) {
	type atTest struct {
		arg1     Series
		arg2     []interface{}
		expected interface{}
	}
	atTests := []atTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
				"string",
			},
			[]interface{}{0},
			"alice",
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
				"string",
			},
			[]interface{}{"b"},
			"banana",
		},
	}

	for _, test := range atTests {
		output, err := test.arg1.At(test.arg2...)
		if output != test.expected || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSeriesIAt(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	names := [][]interface{}{
		{"Avery Bradley"},
		{"Jae Crowder"},
		{"John Holland"},
		{"R.J. Hunter"},
		{"Jonas Jerebko"},
		{"Amir Johnson"},
		{"Jordan Mickey"},
		{"Kelly Olynyk"},
		{"Terry Rozier"},
		{"Marcus Smart"},
		{"Jared Sullinger"},
		{"Isaiah Thomas"},
		{"Evan Turner"},
		{"James Young"},
		{"Tyler Zeller"},
		{"Bojan Bogdanovic"},
		{"Markel Brown"},
		{"Wayne Ellington"},
		{"Rondae Hollis-Jefferson"},
		{"Jarrett Jack"},
		{"Sergey Karasev"},
		{"Sean Kilpatrick"},
		{"Shane Larkin"},
		{"Brook Lopez"},
		{"Chris McCullough"},
		{"Willie Reed"},
		{"Thomas Robinson"},
		{"Henry Sims"},
		{"Donald Sloan"},
		{"Thaddeus Young"},
		{"Arron Afflalo"},
		{"Lou Amundson"},
		{"Thanasis Antetokounmpo"},
		{"Carmelo Anthony"},
		{"Jose Calderon"},
		{"Cleanthony Early"},
		{"Langston Galloway"},
		{"Jerian Grant"},
		{"Robin Lopez"},
		{"Kyle O'Quinn"},
		{"Kristaps Porzingis"},
		{"Kevin Seraphin"},
		{"Lance Thomas"},
		{"Sasha Vujacic"},
		{"Derrick Williams"},
		{"Tony Wroten"},
		{"Elton Brand"},
		{"Isaiah Canaan"},
		{"Robert Covington"},
		{"Joel Embiid"},
		{"Jerami Grant"},
		{"Richaun Holmes"},
		{"Carl Landry"},
		{"Kendall Marshall"},
		{"T.J. McConnell"},
		{"Nerlens Noel"},
		{"Jahlil Okafor"},
		{"Ish Smith"},
		{"Nik Stauskas"},
		{"Hollis Thompson"},
		{"Christian Wood"},
		{"Bismack Biyombo"},
		{"Bruno Caboclo"},
		{"DeMarre Carroll"},
		{"DeMar DeRozan"},
		{"James Johnson"},
		{"Cory Joseph"},
		{"Kyle Lowry"},
		{"Lucas Nogueira"},
		{"Patrick Patterson"},
		{"Norman Powell"},
		{"Terrence Ross"},
		{"Luis Scola"},
		{"Jason Thompson"},
		{"Jonas Valanciunas"},
		{"Delon Wright"},
		{"Leandro Barbosa"},
		{"Harrison Barnes"},
		{"Andrew Bogut"},
		{"Ian Clark"},
		{"Stephen Curry"},
		{"Festus Ezeli"},
		{"Draymond Green"},
		{"Andre Iguodala"},
		{"Shaun Livingston"},
		{"Kevon Looney"},
		{"James Michael McAdoo"},
		{"Brandon Rush"},
		{"Marreese Speights"},
		{"Klay Thompson"},
		{"Anderson Varejao"},
		{"Cole Aldrich"},
		{"Jeff Ayres"},
		{"Jamal Crawford"},
		{"Branden Dawson"},
		{"Jeff Green"},
		{"Blake Griffin"},
		{"Wesley Johnson"},
		{"DeAndre Jordan"},
		{"Luc Richard Mbah a Moute"},
		{"Chris Paul"},
		{"Paul Pierce"},
		{"Pablo Prigioni"},
		{"JJ Redick"},
		{"Austin Rivers"},
		{"C.J. Wilcox"},
		{"Brandon Bass"},
		{"Tarik Black"},
		{"Anthony Brown"},
		{"Kobe Bryant"},
		{"Jordan Clarkson"},
		{"Roy Hibbert"},
		{"Marcelo Huertas"},
		{"Ryan Kelly"},
		{"Larry Nance Jr."},
		{"Julius Randle"},
		{"D'Angelo Russell"},
		{"Robert Sacre"},
		{"Louis Williams"},
		{"Metta World Peace"},
		{"Nick Young"},
		{"Eric Bledsoe"},
		{"Devin Booker"},
		{"Chase Budinger"},
		{"Tyson Chandler"},
		{"Archie Goodwin"},
		{"John Jenkins"},
		{"Brandon Knight"},
		{"Alex Len"},
		{"Jon Leuer"},
		{"Phil Pressey"},
		{"Ronnie Price"},
		{"Mirza Teletovic"},
		{"P.J. Tucker"},
		{"T.J. Warren"},
		{"Alan Williams"},
		{"Quincy Acy"},
		{"James Anderson"},
		{"Marco Belinelli"},
		{"Caron Butler"},
		{"Omri Casspi"},
		{"Willie Cauley-Stein"},
		{"Darren Collison"},
		{"DeMarcus Cousins"},
		{"Seth Curry"},
		{"Duje Dukan"},
		{"Rudy Gay"},
		{"Kosta Koufos"},
		{"Ben McLemore"},
		{"Eric Moreland"},
		{"Rajon Rondo"},
		{"Cameron Bairstow"},
		{"Aaron Brooks"},
		{"Jimmy Butler"},
		{"Mike Dunleavy"},
		{"Cristiano Felicio"},
		{"Pau Gasol"},
		{"Taj Gibson"},
		{"Justin Holiday"},
		{"Doug McDermott"},
		{"Nikola Mirotic"},
		{"E'Twaun Moore"},
		{"Joakim Noah"},
		{"Bobby Portis"},
		{"Derrick Rose"},
		{"Tony Snell"},
		{"Matthew Dellavedova"},
		{"Channing Frye"},
		{"Kyrie Irving"},
		{"LeBron James"},
		{"Richard Jefferson"},
		{"Dahntay Jones"},
		{"James Jones"},
		{"Sasha Kaun"},
		{"Kevin Love"},
		{"Jordan McRae"},
		{"Timofey Mozgov"},
		{"Iman Shumpert"},
		{"J.R. Smith"},
		{"Tristan Thompson"},
		{"Mo Williams"},
		{"Joel Anthony"},
		{"Aron Baynes"},
		{"Steve Blake"},
		{"Lorenzo Brown"},
		{"Reggie Bullock"},
		{"Kentavious Caldwell-Pope"},
		{"Spencer Dinwiddie"},
		{"Andre Drummond"},
		{"Tobias Harris"},
		{"Darrun Hilliard"},
		{"Reggie Jackson"},
		{"Stanley Johnson"},
		{"Jodie Meeks"},
		{"Marcus Morris"},
		{"Anthony Tolliver"},
		{"Lavoy Allen"},
		{"Rakeem Christmas"},
		{"Monta Ellis"},
		{"Paul George"},
		{"George Hill"},
		{"Jordan Hill"},
		{"Solomon Hill"},
		{"Ty Lawson"},
		{"Ian Mahinmi"},
		{"C.J. Miles"},
		{"Glenn Robinson III"},
		{"Rodney Stuckey"},
		{"Myles Turner"},
		{"Shayne Whittington"},
		{"Joe Young"},
		{"Giannis Antetokounmpo"},
		{"Jerryd Bayless"},
		{"Michael Carter-Williams"},
		{"Jared Cunningham"},
		{"Tyler Ennis"},
		{"John Henson"},
		{"Damien Inglis"},
		{"O.J. Mayo"},
		{"Khris Middleton"},
		{"Greg Monroe"},
		{"Steve Novak"},
		{"Johnny O'Bryant III"},
		{"Jabari Parker"},
		{"Miles Plumlee"},
		{"Greivis Vasquez"},
		{"Rashad Vaughn"},
		{"Justin Anderson"},
		{"J.J. Barea"},
		{"Jeremy Evans"},
		{"Raymond Felton"},
		{"Devin Harris"},
		{"David Lee"},
		{"Wesley Matthews"},
		{"JaVale McGee"},
		{"Salah Mejri"},
		{"Dirk Nowitzki"},
		{"Zaza Pachulia"},
		{"Chandler Parsons"},
		{"Dwight Powell"},
		{"Charlie Villanueva"},
		{"Deron Williams"},
		{"Trevor Ariza"},
		{"Michael Beasley"},
		{"Patrick Beverley"},
		{"Corey Brewer"},
		{"Clint Capela"},
		{"Sam Dekker"},
		{"Andrew Goudelock"},
		{"James Harden"},
		{"Montrezl Harrell"},
		{"Dwight Howard"},
		{"Terrence Jones"},
		{"K.J. McDaniels"},
		{"Donatas Motiejunas"},
		{"Josh Smith"},
		{"Jason Terry"},
		{"Jordan Adams"},
		{"Tony Allen"},
		{"Chris Andersen"},
		{"Matt Barnes"},
		{"Vince Carter"},
		{"Mike Conley"},
		{"Bryce Cotton"},
		{"Jordan Farmar"},
		{"Marc Gasol"},
		{"JaMychal Green"},
		{"P.J. Hairston"},
		{"Jarell Martin"},
		{"Ray McCallum"},
		{"Xavier Munford"},
		{"Zach Randolph"},
		{"Lance Stephenson"},
		{"Alex Stepheson"},
		{"Brandan Wright"},
		{"Alexis Ajinca"},
		{"Ryan Anderson"},
		{"Omer Asik"},
		{"Luke Babbitt"},
		{"Norris Cole"},
		{"Dante Cunningham"},
		{"Anthony Davis"},
		{"Bryce Dejean-Jones"},
		{"Toney Douglas"},
		{"James Ennis"},
		{"Tyreke Evans"},
		{"Tim Frazier"},
		{"Alonzo Gee"},
		{"Eric Gordon"},
		{"Jordan Hamilton"},
		{"Jrue Holiday"},
		{"Orlando Johnson"},
		{"Kendrick Perkins"},
		{"Quincy Pondexter"},
		{"LaMarcus Aldridge"},
		{"Kyle Anderson"},
		{"Matt Bonner"},
		{"Boris Diaw"},
		{"Tim Duncan"},
		{"Manu Ginobili"},
		{"Danny Green"},
		{"Kawhi Leonard"},
		{"Boban Marjanovic"},
		{"Kevin Martin"},
		{"Andre Miller"},
		{"Patty Mills"},
		{"Tony Parker"},
		{"Jonathon Simmons"},
		{"David West"},
		{"Kent Bazemore"},
		{"Tim Hardaway Jr."},
		{"Kirk Hinrich"},
		{"Al Horford"},
		{"Kris Humphries"},
		{"Kyle Korver"},
		{"Paul Millsap"},
		{"Mike Muscala"},
		{"Lamar Patterson"},
		{"Dennis Schroder"},
		{"Mike Scott"},
		{"Thabo Sefolosha"},
		{"Tiago Splitter"},
		{"Walter Tavares"},
		{"Jeff Teague"},
		{"Nicolas Batum"},
		{"Troy Daniels"},
		{"Jorge Gutierrez"},
		{"Tyler Hansbrough"},
		{"Aaron Harrison"},
		{"Spencer Hawes"},
		{"Al Jefferson"},
		{"Frank Kaminsky III"},
		{"Michael Kidd-Gilchrist"},
		{"Jeremy Lamb"},
		{"Courtney Lee"},
		{"Jeremy Lin"},
		{"Kemba Walker"},
		{"Marvin Williams"},
		{"Cody Zeller"},
		{"Chris Bosh"},
		{"Luol Deng"},
		{"Goran Dragic"},
		{"Gerald Green"},
		{"Udonis Haslem"},
		{"Joe Johnson"},
		{"Tyler Johnson"},
		{"Josh McRoberts"},
		{"Josh Richardson"},
		{"Amar'e Stoudemire"},
		{"Dwyane Wade"},
		{"Briante Weber"},
		{"Hassan Whiteside"},
		{"Justise Winslow"},
		{"Dorell Wright"},
		{"Dewayne Dedmon"},
		{"Evan Fournier"},
		{"Aaron Gordon"},
		{"Mario Hezonja"},
		{"Ersan Ilyasova"},
		{"Brandon Jennings"},
		{"Devyn Marble"},
		{"Shabazz Napier"},
		{"Andrew Nicholson"},
		{"Victor Oladipo"},
		{"Elfrid Payton"},
		{"Jason Smith"},
		{"Nikola Vucevic"},
		{"C.J. Watson"},
		{"Alan Anderson"},
		{"Bradley Beal"},
		{"Jared Dudley"},
		{"Jarell Eddie"},
		{"Drew Gooden"},
		{"Marcin Gortat"},
		{"JJ Hickson"},
		{"Nene Hilario"},
		{"Markieff Morris"},
		{"Kelly Oubre Jr."},
		{"Otto Porter Jr."},
		{"Ramon Sessions"},
		{"Garrett Temple"},
		{"Marcus Thornton"},
		{"John Wall"},
		{"Darrell Arthur"},
		{"D.J. Augustin"},
		{"Will Barton"},
		{"Wilson Chandler"},
		{"Kenneth Faried"},
		{"Danilo Gallinari"},
		{"Gary Harris"},
		{"Nikola Jokic"},
		{"Joffrey Lauvergne"},
		{"Mike Miller"},
		{"Emmanuel Mudiay"},
		{"Jameer Nelson"},
		{"Jusuf Nurkic"},
		{"JaKarr Sampson"},
		{"Axel Toupane"},
		{"Nemanja Bjelica"},
		{"Gorgui Dieng"},
		{"Kevin Garnett"},
		{"Tyus Jones"},
		{"Zach LaVine"},
		{"Shabazz Muhammad"},
		{"Adreian Payne"},
		{"Nikola Pekovic"},
		{"Tayshaun Prince"},
		{"Ricky Rubio"},
		{"Damjan Rudez"},
		{"Greg Smith"},
		{"Karl-Anthony Towns"},
		{"Andrew Wiggins"},
		{"Steven Adams"},
		{"Nick Collison"},
		{"Kevin Durant"},
		{"Randy Foye"},
		{"Josh Huestis"},
		{"Serge Ibaka"},
		{"Enes Kanter"},
		{"Mitch McGary"},
		{"Nazr Mohammed"},
		{"Anthony Morrow"},
		{"Cameron Payne"},
		{"Andre Roberson"},
		{"Kyle Singler"},
		{"Dion Waiters"},
		{"Russell Westbrook"},
		{"Cliff Alexander"},
		{"Al-Farouq Aminu"},
		{"Pat Connaughton"},
		{"Allen Crabbe"},
		{"Ed Davis"},
		{"Maurice Harkless"},
		{"Gerald Henderson"},
		{"Chris Kaman"},
		{"Meyers Leonard"},
		{"Damian Lillard"},
		{"C.J. McCollum"},
		{"Luis Montero"},
		{"Mason Plumlee"},
		{"Brian Roberts"},
		{"Noah Vonleh"},
		{"Trevor Booker"},
		{"Trey Burke"},
		{"Alec Burks"},
		{"Dante Exum"},
		{"Derrick Favors"},
		{"Rudy Gobert"},
		{"Gordon Hayward"},
		{"Rodney Hood"},
		{"Joe Ingles"},
		{"Chris Johnson"},
		{"Trey Lyles"},
		{"Shelvin Mack"},
		{"Raul Neto"},
		{"Tibor Pleiss"},
		{"Jeff Withey"},
	}
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].IAt(rand.Intn(len(names)))
	}
}

func TestSeriesIAt(t *testing.T) {
	type iatTest struct {
		arg1     Series
		arg2     int
		expected interface{}
	}
	iatTests := []iatTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
				"string",
			},
			0,
			"alice",
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
				"string",
			},
			1,
			"banana",
		},
	}

	for _, test := range iatTests {
		output, err := test.arg1.IAt(test.arg2)
		if output != test.expected || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSeriesLoc(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	names := [][]interface{}{
		{"Avery Bradley"},
		{"Jae Crowder"},
		{"John Holland"},
		{"R.J. Hunter"},
		{"Jonas Jerebko"},
		{"Amir Johnson"},
		{"Jordan Mickey"},
		{"Kelly Olynyk"},
		{"Terry Rozier"},
		{"Marcus Smart"},
		{"Jared Sullinger"},
		{"Isaiah Thomas"},
		{"Evan Turner"},
		{"James Young"},
		{"Tyler Zeller"},
		{"Bojan Bogdanovic"},
		{"Markel Brown"},
		{"Wayne Ellington"},
		{"Rondae Hollis-Jefferson"},
		{"Jarrett Jack"},
		{"Sergey Karasev"},
		{"Sean Kilpatrick"},
		{"Shane Larkin"},
		{"Brook Lopez"},
		{"Chris McCullough"},
		{"Willie Reed"},
		{"Thomas Robinson"},
		{"Henry Sims"},
		{"Donald Sloan"},
		{"Thaddeus Young"},
		{"Arron Afflalo"},
		{"Lou Amundson"},
		{"Thanasis Antetokounmpo"},
		{"Carmelo Anthony"},
		{"Jose Calderon"},
		{"Cleanthony Early"},
		{"Langston Galloway"},
		{"Jerian Grant"},
		{"Robin Lopez"},
		{"Kyle O'Quinn"},
		{"Kristaps Porzingis"},
		{"Kevin Seraphin"},
		{"Lance Thomas"},
		{"Sasha Vujacic"},
		{"Derrick Williams"},
		{"Tony Wroten"},
		{"Elton Brand"},
		{"Isaiah Canaan"},
		{"Robert Covington"},
		{"Joel Embiid"},
		{"Jerami Grant"},
		{"Richaun Holmes"},
		{"Carl Landry"},
		{"Kendall Marshall"},
		{"T.J. McConnell"},
		{"Nerlens Noel"},
		{"Jahlil Okafor"},
		{"Ish Smith"},
		{"Nik Stauskas"},
		{"Hollis Thompson"},
		{"Christian Wood"},
		{"Bismack Biyombo"},
		{"Bruno Caboclo"},
		{"DeMarre Carroll"},
		{"DeMar DeRozan"},
		{"James Johnson"},
		{"Cory Joseph"},
		{"Kyle Lowry"},
		{"Lucas Nogueira"},
		{"Patrick Patterson"},
		{"Norman Powell"},
		{"Terrence Ross"},
		{"Luis Scola"},
		{"Jason Thompson"},
		{"Jonas Valanciunas"},
		{"Delon Wright"},
		{"Leandro Barbosa"},
		{"Harrison Barnes"},
		{"Andrew Bogut"},
		{"Ian Clark"},
		{"Stephen Curry"},
		{"Festus Ezeli"},
		{"Draymond Green"},
		{"Andre Iguodala"},
		{"Shaun Livingston"},
		{"Kevon Looney"},
		{"James Michael McAdoo"},
		{"Brandon Rush"},
		{"Marreese Speights"},
		{"Klay Thompson"},
		{"Anderson Varejao"},
		{"Cole Aldrich"},
		{"Jeff Ayres"},
		{"Jamal Crawford"},
		{"Branden Dawson"},
		{"Jeff Green"},
		{"Blake Griffin"},
		{"Wesley Johnson"},
		{"DeAndre Jordan"},
		{"Luc Richard Mbah a Moute"},
		{"Chris Paul"},
		{"Paul Pierce"},
		{"Pablo Prigioni"},
		{"JJ Redick"},
		{"Austin Rivers"},
		{"C.J. Wilcox"},
		{"Brandon Bass"},
		{"Tarik Black"},
		{"Anthony Brown"},
		{"Kobe Bryant"},
		{"Jordan Clarkson"},
		{"Roy Hibbert"},
		{"Marcelo Huertas"},
		{"Ryan Kelly"},
		{"Larry Nance Jr."},
		{"Julius Randle"},
		{"D'Angelo Russell"},
		{"Robert Sacre"},
		{"Louis Williams"},
		{"Metta World Peace"},
		{"Nick Young"},
		{"Eric Bledsoe"},
		{"Devin Booker"},
		{"Chase Budinger"},
		{"Tyson Chandler"},
		{"Archie Goodwin"},
		{"John Jenkins"},
		{"Brandon Knight"},
		{"Alex Len"},
		{"Jon Leuer"},
		{"Phil Pressey"},
		{"Ronnie Price"},
		{"Mirza Teletovic"},
		{"P.J. Tucker"},
		{"T.J. Warren"},
		{"Alan Williams"},
		{"Quincy Acy"},
		{"James Anderson"},
		{"Marco Belinelli"},
		{"Caron Butler"},
		{"Omri Casspi"},
		{"Willie Cauley-Stein"},
		{"Darren Collison"},
		{"DeMarcus Cousins"},
		{"Seth Curry"},
		{"Duje Dukan"},
		{"Rudy Gay"},
		{"Kosta Koufos"},
		{"Ben McLemore"},
		{"Eric Moreland"},
		{"Rajon Rondo"},
		{"Cameron Bairstow"},
		{"Aaron Brooks"},
		{"Jimmy Butler"},
		{"Mike Dunleavy"},
		{"Cristiano Felicio"},
		{"Pau Gasol"},
		{"Taj Gibson"},
		{"Justin Holiday"},
		{"Doug McDermott"},
		{"Nikola Mirotic"},
		{"E'Twaun Moore"},
		{"Joakim Noah"},
		{"Bobby Portis"},
		{"Derrick Rose"},
		{"Tony Snell"},
		{"Matthew Dellavedova"},
		{"Channing Frye"},
		{"Kyrie Irving"},
		{"LeBron James"},
		{"Richard Jefferson"},
		{"Dahntay Jones"},
		{"James Jones"},
		{"Sasha Kaun"},
		{"Kevin Love"},
		{"Jordan McRae"},
		{"Timofey Mozgov"},
		{"Iman Shumpert"},
		{"J.R. Smith"},
		{"Tristan Thompson"},
		{"Mo Williams"},
		{"Joel Anthony"},
		{"Aron Baynes"},
		{"Steve Blake"},
		{"Lorenzo Brown"},
		{"Reggie Bullock"},
		{"Kentavious Caldwell-Pope"},
		{"Spencer Dinwiddie"},
		{"Andre Drummond"},
		{"Tobias Harris"},
		{"Darrun Hilliard"},
		{"Reggie Jackson"},
		{"Stanley Johnson"},
		{"Jodie Meeks"},
		{"Marcus Morris"},
		{"Anthony Tolliver"},
		{"Lavoy Allen"},
		{"Rakeem Christmas"},
		{"Monta Ellis"},
		{"Paul George"},
		{"George Hill"},
		{"Jordan Hill"},
		{"Solomon Hill"},
		{"Ty Lawson"},
		{"Ian Mahinmi"},
		{"C.J. Miles"},
		{"Glenn Robinson III"},
		{"Rodney Stuckey"},
		{"Myles Turner"},
		{"Shayne Whittington"},
		{"Joe Young"},
		{"Giannis Antetokounmpo"},
		{"Jerryd Bayless"},
		{"Michael Carter-Williams"},
		{"Jared Cunningham"},
		{"Tyler Ennis"},
		{"John Henson"},
		{"Damien Inglis"},
		{"O.J. Mayo"},
		{"Khris Middleton"},
		{"Greg Monroe"},
		{"Steve Novak"},
		{"Johnny O'Bryant III"},
		{"Jabari Parker"},
		{"Miles Plumlee"},
		{"Greivis Vasquez"},
		{"Rashad Vaughn"},
		{"Justin Anderson"},
		{"J.J. Barea"},
		{"Jeremy Evans"},
		{"Raymond Felton"},
		{"Devin Harris"},
		{"David Lee"},
		{"Wesley Matthews"},
		{"JaVale McGee"},
		{"Salah Mejri"},
		{"Dirk Nowitzki"},
		{"Zaza Pachulia"},
		{"Chandler Parsons"},
		{"Dwight Powell"},
		{"Charlie Villanueva"},
		{"Deron Williams"},
		{"Trevor Ariza"},
		{"Michael Beasley"},
		{"Patrick Beverley"},
		{"Corey Brewer"},
		{"Clint Capela"},
		{"Sam Dekker"},
		{"Andrew Goudelock"},
		{"James Harden"},
		{"Montrezl Harrell"},
		{"Dwight Howard"},
		{"Terrence Jones"},
		{"K.J. McDaniels"},
		{"Donatas Motiejunas"},
		{"Josh Smith"},
		{"Jason Terry"},
		{"Jordan Adams"},
		{"Tony Allen"},
		{"Chris Andersen"},
		{"Matt Barnes"},
		{"Vince Carter"},
		{"Mike Conley"},
		{"Bryce Cotton"},
		{"Jordan Farmar"},
		{"Marc Gasol"},
		{"JaMychal Green"},
		{"P.J. Hairston"},
		{"Jarell Martin"},
		{"Ray McCallum"},
		{"Xavier Munford"},
		{"Zach Randolph"},
		{"Lance Stephenson"},
		{"Alex Stepheson"},
		{"Brandan Wright"},
		{"Alexis Ajinca"},
		{"Ryan Anderson"},
		{"Omer Asik"},
		{"Luke Babbitt"},
		{"Norris Cole"},
		{"Dante Cunningham"},
		{"Anthony Davis"},
		{"Bryce Dejean-Jones"},
		{"Toney Douglas"},
		{"James Ennis"},
		{"Tyreke Evans"},
		{"Tim Frazier"},
		{"Alonzo Gee"},
		{"Eric Gordon"},
		{"Jordan Hamilton"},
		{"Jrue Holiday"},
		{"Orlando Johnson"},
		{"Kendrick Perkins"},
		{"Quincy Pondexter"},
		{"LaMarcus Aldridge"},
		{"Kyle Anderson"},
		{"Matt Bonner"},
		{"Boris Diaw"},
		{"Tim Duncan"},
		{"Manu Ginobili"},
		{"Danny Green"},
		{"Kawhi Leonard"},
		{"Boban Marjanovic"},
		{"Kevin Martin"},
		{"Andre Miller"},
		{"Patty Mills"},
		{"Tony Parker"},
		{"Jonathon Simmons"},
		{"David West"},
		{"Kent Bazemore"},
		{"Tim Hardaway Jr."},
		{"Kirk Hinrich"},
		{"Al Horford"},
		{"Kris Humphries"},
		{"Kyle Korver"},
		{"Paul Millsap"},
		{"Mike Muscala"},
		{"Lamar Patterson"},
		{"Dennis Schroder"},
		{"Mike Scott"},
		{"Thabo Sefolosha"},
		{"Tiago Splitter"},
		{"Walter Tavares"},
		{"Jeff Teague"},
		{"Nicolas Batum"},
		{"Troy Daniels"},
		{"Jorge Gutierrez"},
		{"Tyler Hansbrough"},
		{"Aaron Harrison"},
		{"Spencer Hawes"},
		{"Al Jefferson"},
		{"Frank Kaminsky III"},
		{"Michael Kidd-Gilchrist"},
		{"Jeremy Lamb"},
		{"Courtney Lee"},
		{"Jeremy Lin"},
		{"Kemba Walker"},
		{"Marvin Williams"},
		{"Cody Zeller"},
		{"Chris Bosh"},
		{"Luol Deng"},
		{"Goran Dragic"},
		{"Gerald Green"},
		{"Udonis Haslem"},
		{"Joe Johnson"},
		{"Tyler Johnson"},
		{"Josh McRoberts"},
		{"Josh Richardson"},
		{"Amar'e Stoudemire"},
		{"Dwyane Wade"},
		{"Briante Weber"},
		{"Hassan Whiteside"},
		{"Justise Winslow"},
		{"Dorell Wright"},
		{"Dewayne Dedmon"},
		{"Evan Fournier"},
		{"Aaron Gordon"},
		{"Mario Hezonja"},
		{"Ersan Ilyasova"},
		{"Brandon Jennings"},
		{"Devyn Marble"},
		{"Shabazz Napier"},
		{"Andrew Nicholson"},
		{"Victor Oladipo"},
		{"Elfrid Payton"},
		{"Jason Smith"},
		{"Nikola Vucevic"},
		{"C.J. Watson"},
		{"Alan Anderson"},
		{"Bradley Beal"},
		{"Jared Dudley"},
		{"Jarell Eddie"},
		{"Drew Gooden"},
		{"Marcin Gortat"},
		{"JJ Hickson"},
		{"Nene Hilario"},
		{"Markieff Morris"},
		{"Kelly Oubre Jr."},
		{"Otto Porter Jr."},
		{"Ramon Sessions"},
		{"Garrett Temple"},
		{"Marcus Thornton"},
		{"John Wall"},
		{"Darrell Arthur"},
		{"D.J. Augustin"},
		{"Will Barton"},
		{"Wilson Chandler"},
		{"Kenneth Faried"},
		{"Danilo Gallinari"},
		{"Gary Harris"},
		{"Nikola Jokic"},
		{"Joffrey Lauvergne"},
		{"Mike Miller"},
		{"Emmanuel Mudiay"},
		{"Jameer Nelson"},
		{"Jusuf Nurkic"},
		{"JaKarr Sampson"},
		{"Axel Toupane"},
		{"Nemanja Bjelica"},
		{"Gorgui Dieng"},
		{"Kevin Garnett"},
		{"Tyus Jones"},
		{"Zach LaVine"},
		{"Shabazz Muhammad"},
		{"Adreian Payne"},
		{"Nikola Pekovic"},
		{"Tayshaun Prince"},
		{"Ricky Rubio"},
		{"Damjan Rudez"},
		{"Greg Smith"},
		{"Karl-Anthony Towns"},
		{"Andrew Wiggins"},
		{"Steven Adams"},
		{"Nick Collison"},
		{"Kevin Durant"},
		{"Randy Foye"},
		{"Josh Huestis"},
		{"Serge Ibaka"},
		{"Enes Kanter"},
		{"Mitch McGary"},
		{"Nazr Mohammed"},
		{"Anthony Morrow"},
		{"Cameron Payne"},
		{"Andre Roberson"},
		{"Kyle Singler"},
		{"Dion Waiters"},
		{"Russell Westbrook"},
		{"Cliff Alexander"},
		{"Al-Farouq Aminu"},
		{"Pat Connaughton"},
		{"Allen Crabbe"},
		{"Ed Davis"},
		{"Maurice Harkless"},
		{"Gerald Henderson"},
		{"Chris Kaman"},
		{"Meyers Leonard"},
		{"Damian Lillard"},
		{"C.J. McCollum"},
		{"Luis Montero"},
		{"Mason Plumlee"},
		{"Brian Roberts"},
		{"Noah Vonleh"},
		{"Trevor Booker"},
		{"Trey Burke"},
		{"Alec Burks"},
		{"Dante Exum"},
		{"Derrick Favors"},
		{"Rudy Gobert"},
		{"Gordon Hayward"},
		{"Rodney Hood"},
		{"Joe Ingles"},
		{"Chris Johnson"},
		{"Trey Lyles"},
		{"Shelvin Mack"},
		{"Raul Neto"},
		{"Tibor Pleiss"},
		{"Jeff Withey"},
	}
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].Loc(names[rand.Intn(len(names))])
	}
}

func TestSeriesLoc(t *testing.T) {
	type locTest struct {
		arg1          Series
		arg2          [][]interface{}
		expected      Series
		expectedError error
	}
	locTests := []locTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
				"string",
			},
			[][]interface{}{{0}, {1}},
			Series{
				[]interface{}{"alice", "bob"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
					},
					[]string{""},
				},
				"People",
				"string",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
				"string",
			},
			[][]interface{}{{"b"}, {"c"}},
			Series{
				[]interface{}{"banana", "cherry"},
				IndexData{
					[]Index{
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
				"string",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"female"}},
			Series{
				[]interface{}{"clara", "anna"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"male", "volleyball"}},
			Series{
				[]interface{}{"brian"},
				IndexData{
					[]Index{{1, []interface{}{"male", "volleyball"}}},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"male"}, {"volleyball"}},
			Series{},
			fmt.Errorf("no data found for index [volleyball]"),
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"volleyball"}},
			Series{},
			fmt.Errorf("no data found for index [volleyball]"),
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"female"}, {"male"}},
			Series{
				[]interface{}{"clara", "anna", "brian", "dorian", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"female"}, {"volleyball"}},
			Series{},
			fmt.Errorf("no data found for index [volleyball]"),
		},
	}

	for _, test := range locTests {
		output, err := test.arg1.Loc(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSeriesLocItems(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	names := [][]interface{}{
		{"Avery Bradley"},
		{"Jae Crowder"},
		{"John Holland"},
		{"R.J. Hunter"},
		{"Jonas Jerebko"},
		{"Amir Johnson"},
		{"Jordan Mickey"},
		{"Kelly Olynyk"},
		{"Terry Rozier"},
		{"Marcus Smart"},
		{"Jared Sullinger"},
		{"Isaiah Thomas"},
		{"Evan Turner"},
		{"James Young"},
		{"Tyler Zeller"},
		{"Bojan Bogdanovic"},
		{"Markel Brown"},
		{"Wayne Ellington"},
		{"Rondae Hollis-Jefferson"},
		{"Jarrett Jack"},
		{"Sergey Karasev"},
		{"Sean Kilpatrick"},
		{"Shane Larkin"},
		{"Brook Lopez"},
		{"Chris McCullough"},
		{"Willie Reed"},
		{"Thomas Robinson"},
		{"Henry Sims"},
		{"Donald Sloan"},
		{"Thaddeus Young"},
		{"Arron Afflalo"},
		{"Lou Amundson"},
		{"Thanasis Antetokounmpo"},
		{"Carmelo Anthony"},
		{"Jose Calderon"},
		{"Cleanthony Early"},
		{"Langston Galloway"},
		{"Jerian Grant"},
		{"Robin Lopez"},
		{"Kyle O'Quinn"},
		{"Kristaps Porzingis"},
		{"Kevin Seraphin"},
		{"Lance Thomas"},
		{"Sasha Vujacic"},
		{"Derrick Williams"},
		{"Tony Wroten"},
		{"Elton Brand"},
		{"Isaiah Canaan"},
		{"Robert Covington"},
		{"Joel Embiid"},
		{"Jerami Grant"},
		{"Richaun Holmes"},
		{"Carl Landry"},
		{"Kendall Marshall"},
		{"T.J. McConnell"},
		{"Nerlens Noel"},
		{"Jahlil Okafor"},
		{"Ish Smith"},
		{"Nik Stauskas"},
		{"Hollis Thompson"},
		{"Christian Wood"},
		{"Bismack Biyombo"},
		{"Bruno Caboclo"},
		{"DeMarre Carroll"},
		{"DeMar DeRozan"},
		{"James Johnson"},
		{"Cory Joseph"},
		{"Kyle Lowry"},
		{"Lucas Nogueira"},
		{"Patrick Patterson"},
		{"Norman Powell"},
		{"Terrence Ross"},
		{"Luis Scola"},
		{"Jason Thompson"},
		{"Jonas Valanciunas"},
		{"Delon Wright"},
		{"Leandro Barbosa"},
		{"Harrison Barnes"},
		{"Andrew Bogut"},
		{"Ian Clark"},
		{"Stephen Curry"},
		{"Festus Ezeli"},
		{"Draymond Green"},
		{"Andre Iguodala"},
		{"Shaun Livingston"},
		{"Kevon Looney"},
		{"James Michael McAdoo"},
		{"Brandon Rush"},
		{"Marreese Speights"},
		{"Klay Thompson"},
		{"Anderson Varejao"},
		{"Cole Aldrich"},
		{"Jeff Ayres"},
		{"Jamal Crawford"},
		{"Branden Dawson"},
		{"Jeff Green"},
		{"Blake Griffin"},
		{"Wesley Johnson"},
		{"DeAndre Jordan"},
		{"Luc Richard Mbah a Moute"},
		{"Chris Paul"},
		{"Paul Pierce"},
		{"Pablo Prigioni"},
		{"JJ Redick"},
		{"Austin Rivers"},
		{"C.J. Wilcox"},
		{"Brandon Bass"},
		{"Tarik Black"},
		{"Anthony Brown"},
		{"Kobe Bryant"},
		{"Jordan Clarkson"},
		{"Roy Hibbert"},
		{"Marcelo Huertas"},
		{"Ryan Kelly"},
		{"Larry Nance Jr."},
		{"Julius Randle"},
		{"D'Angelo Russell"},
		{"Robert Sacre"},
		{"Louis Williams"},
		{"Metta World Peace"},
		{"Nick Young"},
		{"Eric Bledsoe"},
		{"Devin Booker"},
		{"Chase Budinger"},
		{"Tyson Chandler"},
		{"Archie Goodwin"},
		{"John Jenkins"},
		{"Brandon Knight"},
		{"Alex Len"},
		{"Jon Leuer"},
		{"Phil Pressey"},
		{"Ronnie Price"},
		{"Mirza Teletovic"},
		{"P.J. Tucker"},
		{"T.J. Warren"},
		{"Alan Williams"},
		{"Quincy Acy"},
		{"James Anderson"},
		{"Marco Belinelli"},
		{"Caron Butler"},
		{"Omri Casspi"},
		{"Willie Cauley-Stein"},
		{"Darren Collison"},
		{"DeMarcus Cousins"},
		{"Seth Curry"},
		{"Duje Dukan"},
		{"Rudy Gay"},
		{"Kosta Koufos"},
		{"Ben McLemore"},
		{"Eric Moreland"},
		{"Rajon Rondo"},
		{"Cameron Bairstow"},
		{"Aaron Brooks"},
		{"Jimmy Butler"},
		{"Mike Dunleavy"},
		{"Cristiano Felicio"},
		{"Pau Gasol"},
		{"Taj Gibson"},
		{"Justin Holiday"},
		{"Doug McDermott"},
		{"Nikola Mirotic"},
		{"E'Twaun Moore"},
		{"Joakim Noah"},
		{"Bobby Portis"},
		{"Derrick Rose"},
		{"Tony Snell"},
		{"Matthew Dellavedova"},
		{"Channing Frye"},
		{"Kyrie Irving"},
		{"LeBron James"},
		{"Richard Jefferson"},
		{"Dahntay Jones"},
		{"James Jones"},
		{"Sasha Kaun"},
		{"Kevin Love"},
		{"Jordan McRae"},
		{"Timofey Mozgov"},
		{"Iman Shumpert"},
		{"J.R. Smith"},
		{"Tristan Thompson"},
		{"Mo Williams"},
		{"Joel Anthony"},
		{"Aron Baynes"},
		{"Steve Blake"},
		{"Lorenzo Brown"},
		{"Reggie Bullock"},
		{"Kentavious Caldwell-Pope"},
		{"Spencer Dinwiddie"},
		{"Andre Drummond"},
		{"Tobias Harris"},
		{"Darrun Hilliard"},
		{"Reggie Jackson"},
		{"Stanley Johnson"},
		{"Jodie Meeks"},
		{"Marcus Morris"},
		{"Anthony Tolliver"},
		{"Lavoy Allen"},
		{"Rakeem Christmas"},
		{"Monta Ellis"},
		{"Paul George"},
		{"George Hill"},
		{"Jordan Hill"},
		{"Solomon Hill"},
		{"Ty Lawson"},
		{"Ian Mahinmi"},
		{"C.J. Miles"},
		{"Glenn Robinson III"},
		{"Rodney Stuckey"},
		{"Myles Turner"},
		{"Shayne Whittington"},
		{"Joe Young"},
		{"Giannis Antetokounmpo"},
		{"Jerryd Bayless"},
		{"Michael Carter-Williams"},
		{"Jared Cunningham"},
		{"Tyler Ennis"},
		{"John Henson"},
		{"Damien Inglis"},
		{"O.J. Mayo"},
		{"Khris Middleton"},
		{"Greg Monroe"},
		{"Steve Novak"},
		{"Johnny O'Bryant III"},
		{"Jabari Parker"},
		{"Miles Plumlee"},
		{"Greivis Vasquez"},
		{"Rashad Vaughn"},
		{"Justin Anderson"},
		{"J.J. Barea"},
		{"Jeremy Evans"},
		{"Raymond Felton"},
		{"Devin Harris"},
		{"David Lee"},
		{"Wesley Matthews"},
		{"JaVale McGee"},
		{"Salah Mejri"},
		{"Dirk Nowitzki"},
		{"Zaza Pachulia"},
		{"Chandler Parsons"},
		{"Dwight Powell"},
		{"Charlie Villanueva"},
		{"Deron Williams"},
		{"Trevor Ariza"},
		{"Michael Beasley"},
		{"Patrick Beverley"},
		{"Corey Brewer"},
		{"Clint Capela"},
		{"Sam Dekker"},
		{"Andrew Goudelock"},
		{"James Harden"},
		{"Montrezl Harrell"},
		{"Dwight Howard"},
		{"Terrence Jones"},
		{"K.J. McDaniels"},
		{"Donatas Motiejunas"},
		{"Josh Smith"},
		{"Jason Terry"},
		{"Jordan Adams"},
		{"Tony Allen"},
		{"Chris Andersen"},
		{"Matt Barnes"},
		{"Vince Carter"},
		{"Mike Conley"},
		{"Bryce Cotton"},
		{"Jordan Farmar"},
		{"Marc Gasol"},
		{"JaMychal Green"},
		{"P.J. Hairston"},
		{"Jarell Martin"},
		{"Ray McCallum"},
		{"Xavier Munford"},
		{"Zach Randolph"},
		{"Lance Stephenson"},
		{"Alex Stepheson"},
		{"Brandan Wright"},
		{"Alexis Ajinca"},
		{"Ryan Anderson"},
		{"Omer Asik"},
		{"Luke Babbitt"},
		{"Norris Cole"},
		{"Dante Cunningham"},
		{"Anthony Davis"},
		{"Bryce Dejean-Jones"},
		{"Toney Douglas"},
		{"James Ennis"},
		{"Tyreke Evans"},
		{"Tim Frazier"},
		{"Alonzo Gee"},
		{"Eric Gordon"},
		{"Jordan Hamilton"},
		{"Jrue Holiday"},
		{"Orlando Johnson"},
		{"Kendrick Perkins"},
		{"Quincy Pondexter"},
		{"LaMarcus Aldridge"},
		{"Kyle Anderson"},
		{"Matt Bonner"},
		{"Boris Diaw"},
		{"Tim Duncan"},
		{"Manu Ginobili"},
		{"Danny Green"},
		{"Kawhi Leonard"},
		{"Boban Marjanovic"},
		{"Kevin Martin"},
		{"Andre Miller"},
		{"Patty Mills"},
		{"Tony Parker"},
		{"Jonathon Simmons"},
		{"David West"},
		{"Kent Bazemore"},
		{"Tim Hardaway Jr."},
		{"Kirk Hinrich"},
		{"Al Horford"},
		{"Kris Humphries"},
		{"Kyle Korver"},
		{"Paul Millsap"},
		{"Mike Muscala"},
		{"Lamar Patterson"},
		{"Dennis Schroder"},
		{"Mike Scott"},
		{"Thabo Sefolosha"},
		{"Tiago Splitter"},
		{"Walter Tavares"},
		{"Jeff Teague"},
		{"Nicolas Batum"},
		{"Troy Daniels"},
		{"Jorge Gutierrez"},
		{"Tyler Hansbrough"},
		{"Aaron Harrison"},
		{"Spencer Hawes"},
		{"Al Jefferson"},
		{"Frank Kaminsky III"},
		{"Michael Kidd-Gilchrist"},
		{"Jeremy Lamb"},
		{"Courtney Lee"},
		{"Jeremy Lin"},
		{"Kemba Walker"},
		{"Marvin Williams"},
		{"Cody Zeller"},
		{"Chris Bosh"},
		{"Luol Deng"},
		{"Goran Dragic"},
		{"Gerald Green"},
		{"Udonis Haslem"},
		{"Joe Johnson"},
		{"Tyler Johnson"},
		{"Josh McRoberts"},
		{"Josh Richardson"},
		{"Amar'e Stoudemire"},
		{"Dwyane Wade"},
		{"Briante Weber"},
		{"Hassan Whiteside"},
		{"Justise Winslow"},
		{"Dorell Wright"},
		{"Dewayne Dedmon"},
		{"Evan Fournier"},
		{"Aaron Gordon"},
		{"Mario Hezonja"},
		{"Ersan Ilyasova"},
		{"Brandon Jennings"},
		{"Devyn Marble"},
		{"Shabazz Napier"},
		{"Andrew Nicholson"},
		{"Victor Oladipo"},
		{"Elfrid Payton"},
		{"Jason Smith"},
		{"Nikola Vucevic"},
		{"C.J. Watson"},
		{"Alan Anderson"},
		{"Bradley Beal"},
		{"Jared Dudley"},
		{"Jarell Eddie"},
		{"Drew Gooden"},
		{"Marcin Gortat"},
		{"JJ Hickson"},
		{"Nene Hilario"},
		{"Markieff Morris"},
		{"Kelly Oubre Jr."},
		{"Otto Porter Jr."},
		{"Ramon Sessions"},
		{"Garrett Temple"},
		{"Marcus Thornton"},
		{"John Wall"},
		{"Darrell Arthur"},
		{"D.J. Augustin"},
		{"Will Barton"},
		{"Wilson Chandler"},
		{"Kenneth Faried"},
		{"Danilo Gallinari"},
		{"Gary Harris"},
		{"Nikola Jokic"},
		{"Joffrey Lauvergne"},
		{"Mike Miller"},
		{"Emmanuel Mudiay"},
		{"Jameer Nelson"},
		{"Jusuf Nurkic"},
		{"JaKarr Sampson"},
		{"Axel Toupane"},
		{"Nemanja Bjelica"},
		{"Gorgui Dieng"},
		{"Kevin Garnett"},
		{"Tyus Jones"},
		{"Zach LaVine"},
		{"Shabazz Muhammad"},
		{"Adreian Payne"},
		{"Nikola Pekovic"},
		{"Tayshaun Prince"},
		{"Ricky Rubio"},
		{"Damjan Rudez"},
		{"Greg Smith"},
		{"Karl-Anthony Towns"},
		{"Andrew Wiggins"},
		{"Steven Adams"},
		{"Nick Collison"},
		{"Kevin Durant"},
		{"Randy Foye"},
		{"Josh Huestis"},
		{"Serge Ibaka"},
		{"Enes Kanter"},
		{"Mitch McGary"},
		{"Nazr Mohammed"},
		{"Anthony Morrow"},
		{"Cameron Payne"},
		{"Andre Roberson"},
		{"Kyle Singler"},
		{"Dion Waiters"},
		{"Russell Westbrook"},
		{"Cliff Alexander"},
		{"Al-Farouq Aminu"},
		{"Pat Connaughton"},
		{"Allen Crabbe"},
		{"Ed Davis"},
		{"Maurice Harkless"},
		{"Gerald Henderson"},
		{"Chris Kaman"},
		{"Meyers Leonard"},
		{"Damian Lillard"},
		{"C.J. McCollum"},
		{"Luis Montero"},
		{"Mason Plumlee"},
		{"Brian Roberts"},
		{"Noah Vonleh"},
		{"Trevor Booker"},
		{"Trey Burke"},
		{"Alec Burks"},
		{"Dante Exum"},
		{"Derrick Favors"},
		{"Rudy Gobert"},
		{"Gordon Hayward"},
		{"Rodney Hood"},
		{"Joe Ingles"},
		{"Chris Johnson"},
		{"Trey Lyles"},
		{"Shelvin Mack"},
		{"Raul Neto"},
		{"Tibor Pleiss"},
		{"Jeff Withey"},
	}
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].LocItems(names[rand.Intn(len(names))])
	}
}

func TestSeriesLocItems(t *testing.T) {
	type locItemsTest struct {
		arg1          Series
		arg2          [][]interface{}
		expected      []interface{}
		expectedError error
	}
	locItemsTests := []locItemsTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
				"string",
			},
			[][]interface{}{{0}, {1}},
			[]interface{}{"alice", "bob"},
			nil,
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
				"string",
			},
			[][]interface{}{{"b"}, {"c"}},
			[]interface{}{"banana", "cherry"},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"female"}},
			[]interface{}{"clara", "anna"},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"male", "volleyball"}},
			[]interface{}{"brian"},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"male"}, {"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"female"}, {"male"}},
			[]interface{}{"clara", "anna", "brian", "dorian", "michael"},
			nil,
		},
		{
			Series{
				[]interface{}{"clara", "brian", "dorian", "anna", "michael"},
				IndexData{
					[]Index{
						{0, []interface{}{"female", "basketball"}},
						{1, []interface{}{"male", "volleyball"}},
						{2, []interface{}{"male", "basketball"}},
						{3, []interface{}{"female", "volleyball"}},
						{4, []interface{}{"male", "swimming"}},
					},
					[]string{"sex", "sports"},
				},
				"People",
				"string",
			},
			[][]interface{}{{"female"}, {"volleyball"}},
			nil,
			fmt.Errorf("no data found for index [volleyball]"),
		},
	}

	for _, test := range locItemsTests {
		output, err := test.arg1.LocItems(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) || (fmt.Sprint(err) != fmt.Sprint(test.expectedError)) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSeriesILoc(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	names := [][]interface{}{
		{"Avery Bradley"},
		{"Jae Crowder"},
		{"John Holland"},
		{"R.J. Hunter"},
		{"Jonas Jerebko"},
		{"Amir Johnson"},
		{"Jordan Mickey"},
		{"Kelly Olynyk"},
		{"Terry Rozier"},
		{"Marcus Smart"},
		{"Jared Sullinger"},
		{"Isaiah Thomas"},
		{"Evan Turner"},
		{"James Young"},
		{"Tyler Zeller"},
		{"Bojan Bogdanovic"},
		{"Markel Brown"},
		{"Wayne Ellington"},
		{"Rondae Hollis-Jefferson"},
		{"Jarrett Jack"},
		{"Sergey Karasev"},
		{"Sean Kilpatrick"},
		{"Shane Larkin"},
		{"Brook Lopez"},
		{"Chris McCullough"},
		{"Willie Reed"},
		{"Thomas Robinson"},
		{"Henry Sims"},
		{"Donald Sloan"},
		{"Thaddeus Young"},
		{"Arron Afflalo"},
		{"Lou Amundson"},
		{"Thanasis Antetokounmpo"},
		{"Carmelo Anthony"},
		{"Jose Calderon"},
		{"Cleanthony Early"},
		{"Langston Galloway"},
		{"Jerian Grant"},
		{"Robin Lopez"},
		{"Kyle O'Quinn"},
		{"Kristaps Porzingis"},
		{"Kevin Seraphin"},
		{"Lance Thomas"},
		{"Sasha Vujacic"},
		{"Derrick Williams"},
		{"Tony Wroten"},
		{"Elton Brand"},
		{"Isaiah Canaan"},
		{"Robert Covington"},
		{"Joel Embiid"},
		{"Jerami Grant"},
		{"Richaun Holmes"},
		{"Carl Landry"},
		{"Kendall Marshall"},
		{"T.J. McConnell"},
		{"Nerlens Noel"},
		{"Jahlil Okafor"},
		{"Ish Smith"},
		{"Nik Stauskas"},
		{"Hollis Thompson"},
		{"Christian Wood"},
		{"Bismack Biyombo"},
		{"Bruno Caboclo"},
		{"DeMarre Carroll"},
		{"DeMar DeRozan"},
		{"James Johnson"},
		{"Cory Joseph"},
		{"Kyle Lowry"},
		{"Lucas Nogueira"},
		{"Patrick Patterson"},
		{"Norman Powell"},
		{"Terrence Ross"},
		{"Luis Scola"},
		{"Jason Thompson"},
		{"Jonas Valanciunas"},
		{"Delon Wright"},
		{"Leandro Barbosa"},
		{"Harrison Barnes"},
		{"Andrew Bogut"},
		{"Ian Clark"},
		{"Stephen Curry"},
		{"Festus Ezeli"},
		{"Draymond Green"},
		{"Andre Iguodala"},
		{"Shaun Livingston"},
		{"Kevon Looney"},
		{"James Michael McAdoo"},
		{"Brandon Rush"},
		{"Marreese Speights"},
		{"Klay Thompson"},
		{"Anderson Varejao"},
		{"Cole Aldrich"},
		{"Jeff Ayres"},
		{"Jamal Crawford"},
		{"Branden Dawson"},
		{"Jeff Green"},
		{"Blake Griffin"},
		{"Wesley Johnson"},
		{"DeAndre Jordan"},
		{"Luc Richard Mbah a Moute"},
		{"Chris Paul"},
		{"Paul Pierce"},
		{"Pablo Prigioni"},
		{"JJ Redick"},
		{"Austin Rivers"},
		{"C.J. Wilcox"},
		{"Brandon Bass"},
		{"Tarik Black"},
		{"Anthony Brown"},
		{"Kobe Bryant"},
		{"Jordan Clarkson"},
		{"Roy Hibbert"},
		{"Marcelo Huertas"},
		{"Ryan Kelly"},
		{"Larry Nance Jr."},
		{"Julius Randle"},
		{"D'Angelo Russell"},
		{"Robert Sacre"},
		{"Louis Williams"},
		{"Metta World Peace"},
		{"Nick Young"},
		{"Eric Bledsoe"},
		{"Devin Booker"},
		{"Chase Budinger"},
		{"Tyson Chandler"},
		{"Archie Goodwin"},
		{"John Jenkins"},
		{"Brandon Knight"},
		{"Alex Len"},
		{"Jon Leuer"},
		{"Phil Pressey"},
		{"Ronnie Price"},
		{"Mirza Teletovic"},
		{"P.J. Tucker"},
		{"T.J. Warren"},
		{"Alan Williams"},
		{"Quincy Acy"},
		{"James Anderson"},
		{"Marco Belinelli"},
		{"Caron Butler"},
		{"Omri Casspi"},
		{"Willie Cauley-Stein"},
		{"Darren Collison"},
		{"DeMarcus Cousins"},
		{"Seth Curry"},
		{"Duje Dukan"},
		{"Rudy Gay"},
		{"Kosta Koufos"},
		{"Ben McLemore"},
		{"Eric Moreland"},
		{"Rajon Rondo"},
		{"Cameron Bairstow"},
		{"Aaron Brooks"},
		{"Jimmy Butler"},
		{"Mike Dunleavy"},
		{"Cristiano Felicio"},
		{"Pau Gasol"},
		{"Taj Gibson"},
		{"Justin Holiday"},
		{"Doug McDermott"},
		{"Nikola Mirotic"},
		{"E'Twaun Moore"},
		{"Joakim Noah"},
		{"Bobby Portis"},
		{"Derrick Rose"},
		{"Tony Snell"},
		{"Matthew Dellavedova"},
		{"Channing Frye"},
		{"Kyrie Irving"},
		{"LeBron James"},
		{"Richard Jefferson"},
		{"Dahntay Jones"},
		{"James Jones"},
		{"Sasha Kaun"},
		{"Kevin Love"},
		{"Jordan McRae"},
		{"Timofey Mozgov"},
		{"Iman Shumpert"},
		{"J.R. Smith"},
		{"Tristan Thompson"},
		{"Mo Williams"},
		{"Joel Anthony"},
		{"Aron Baynes"},
		{"Steve Blake"},
		{"Lorenzo Brown"},
		{"Reggie Bullock"},
		{"Kentavious Caldwell-Pope"},
		{"Spencer Dinwiddie"},
		{"Andre Drummond"},
		{"Tobias Harris"},
		{"Darrun Hilliard"},
		{"Reggie Jackson"},
		{"Stanley Johnson"},
		{"Jodie Meeks"},
		{"Marcus Morris"},
		{"Anthony Tolliver"},
		{"Lavoy Allen"},
		{"Rakeem Christmas"},
		{"Monta Ellis"},
		{"Paul George"},
		{"George Hill"},
		{"Jordan Hill"},
		{"Solomon Hill"},
		{"Ty Lawson"},
		{"Ian Mahinmi"},
		{"C.J. Miles"},
		{"Glenn Robinson III"},
		{"Rodney Stuckey"},
		{"Myles Turner"},
		{"Shayne Whittington"},
		{"Joe Young"},
		{"Giannis Antetokounmpo"},
		{"Jerryd Bayless"},
		{"Michael Carter-Williams"},
		{"Jared Cunningham"},
		{"Tyler Ennis"},
		{"John Henson"},
		{"Damien Inglis"},
		{"O.J. Mayo"},
		{"Khris Middleton"},
		{"Greg Monroe"},
		{"Steve Novak"},
		{"Johnny O'Bryant III"},
		{"Jabari Parker"},
		{"Miles Plumlee"},
		{"Greivis Vasquez"},
		{"Rashad Vaughn"},
		{"Justin Anderson"},
		{"J.J. Barea"},
		{"Jeremy Evans"},
		{"Raymond Felton"},
		{"Devin Harris"},
		{"David Lee"},
		{"Wesley Matthews"},
		{"JaVale McGee"},
		{"Salah Mejri"},
		{"Dirk Nowitzki"},
		{"Zaza Pachulia"},
		{"Chandler Parsons"},
		{"Dwight Powell"},
		{"Charlie Villanueva"},
		{"Deron Williams"},
		{"Trevor Ariza"},
		{"Michael Beasley"},
		{"Patrick Beverley"},
		{"Corey Brewer"},
		{"Clint Capela"},
		{"Sam Dekker"},
		{"Andrew Goudelock"},
		{"James Harden"},
		{"Montrezl Harrell"},
		{"Dwight Howard"},
		{"Terrence Jones"},
		{"K.J. McDaniels"},
		{"Donatas Motiejunas"},
		{"Josh Smith"},
		{"Jason Terry"},
		{"Jordan Adams"},
		{"Tony Allen"},
		{"Chris Andersen"},
		{"Matt Barnes"},
		{"Vince Carter"},
		{"Mike Conley"},
		{"Bryce Cotton"},
		{"Jordan Farmar"},
		{"Marc Gasol"},
		{"JaMychal Green"},
		{"P.J. Hairston"},
		{"Jarell Martin"},
		{"Ray McCallum"},
		{"Xavier Munford"},
		{"Zach Randolph"},
		{"Lance Stephenson"},
		{"Alex Stepheson"},
		{"Brandan Wright"},
		{"Alexis Ajinca"},
		{"Ryan Anderson"},
		{"Omer Asik"},
		{"Luke Babbitt"},
		{"Norris Cole"},
		{"Dante Cunningham"},
		{"Anthony Davis"},
		{"Bryce Dejean-Jones"},
		{"Toney Douglas"},
		{"James Ennis"},
		{"Tyreke Evans"},
		{"Tim Frazier"},
		{"Alonzo Gee"},
		{"Eric Gordon"},
		{"Jordan Hamilton"},
		{"Jrue Holiday"},
		{"Orlando Johnson"},
		{"Kendrick Perkins"},
		{"Quincy Pondexter"},
		{"LaMarcus Aldridge"},
		{"Kyle Anderson"},
		{"Matt Bonner"},
		{"Boris Diaw"},
		{"Tim Duncan"},
		{"Manu Ginobili"},
		{"Danny Green"},
		{"Kawhi Leonard"},
		{"Boban Marjanovic"},
		{"Kevin Martin"},
		{"Andre Miller"},
		{"Patty Mills"},
		{"Tony Parker"},
		{"Jonathon Simmons"},
		{"David West"},
		{"Kent Bazemore"},
		{"Tim Hardaway Jr."},
		{"Kirk Hinrich"},
		{"Al Horford"},
		{"Kris Humphries"},
		{"Kyle Korver"},
		{"Paul Millsap"},
		{"Mike Muscala"},
		{"Lamar Patterson"},
		{"Dennis Schroder"},
		{"Mike Scott"},
		{"Thabo Sefolosha"},
		{"Tiago Splitter"},
		{"Walter Tavares"},
		{"Jeff Teague"},
		{"Nicolas Batum"},
		{"Troy Daniels"},
		{"Jorge Gutierrez"},
		{"Tyler Hansbrough"},
		{"Aaron Harrison"},
		{"Spencer Hawes"},
		{"Al Jefferson"},
		{"Frank Kaminsky III"},
		{"Michael Kidd-Gilchrist"},
		{"Jeremy Lamb"},
		{"Courtney Lee"},
		{"Jeremy Lin"},
		{"Kemba Walker"},
		{"Marvin Williams"},
		{"Cody Zeller"},
		{"Chris Bosh"},
		{"Luol Deng"},
		{"Goran Dragic"},
		{"Gerald Green"},
		{"Udonis Haslem"},
		{"Joe Johnson"},
		{"Tyler Johnson"},
		{"Josh McRoberts"},
		{"Josh Richardson"},
		{"Amar'e Stoudemire"},
		{"Dwyane Wade"},
		{"Briante Weber"},
		{"Hassan Whiteside"},
		{"Justise Winslow"},
		{"Dorell Wright"},
		{"Dewayne Dedmon"},
		{"Evan Fournier"},
		{"Aaron Gordon"},
		{"Mario Hezonja"},
		{"Ersan Ilyasova"},
		{"Brandon Jennings"},
		{"Devyn Marble"},
		{"Shabazz Napier"},
		{"Andrew Nicholson"},
		{"Victor Oladipo"},
		{"Elfrid Payton"},
		{"Jason Smith"},
		{"Nikola Vucevic"},
		{"C.J. Watson"},
		{"Alan Anderson"},
		{"Bradley Beal"},
		{"Jared Dudley"},
		{"Jarell Eddie"},
		{"Drew Gooden"},
		{"Marcin Gortat"},
		{"JJ Hickson"},
		{"Nene Hilario"},
		{"Markieff Morris"},
		{"Kelly Oubre Jr."},
		{"Otto Porter Jr."},
		{"Ramon Sessions"},
		{"Garrett Temple"},
		{"Marcus Thornton"},
		{"John Wall"},
		{"Darrell Arthur"},
		{"D.J. Augustin"},
		{"Will Barton"},
		{"Wilson Chandler"},
		{"Kenneth Faried"},
		{"Danilo Gallinari"},
		{"Gary Harris"},
		{"Nikola Jokic"},
		{"Joffrey Lauvergne"},
		{"Mike Miller"},
		{"Emmanuel Mudiay"},
		{"Jameer Nelson"},
		{"Jusuf Nurkic"},
		{"JaKarr Sampson"},
		{"Axel Toupane"},
		{"Nemanja Bjelica"},
		{"Gorgui Dieng"},
		{"Kevin Garnett"},
		{"Tyus Jones"},
		{"Zach LaVine"},
		{"Shabazz Muhammad"},
		{"Adreian Payne"},
		{"Nikola Pekovic"},
		{"Tayshaun Prince"},
		{"Ricky Rubio"},
		{"Damjan Rudez"},
		{"Greg Smith"},
		{"Karl-Anthony Towns"},
		{"Andrew Wiggins"},
		{"Steven Adams"},
		{"Nick Collison"},
		{"Kevin Durant"},
		{"Randy Foye"},
		{"Josh Huestis"},
		{"Serge Ibaka"},
		{"Enes Kanter"},
		{"Mitch McGary"},
		{"Nazr Mohammed"},
		{"Anthony Morrow"},
		{"Cameron Payne"},
		{"Andre Roberson"},
		{"Kyle Singler"},
		{"Dion Waiters"},
		{"Russell Westbrook"},
		{"Cliff Alexander"},
		{"Al-Farouq Aminu"},
		{"Pat Connaughton"},
		{"Allen Crabbe"},
		{"Ed Davis"},
		{"Maurice Harkless"},
		{"Gerald Henderson"},
		{"Chris Kaman"},
		{"Meyers Leonard"},
		{"Damian Lillard"},
		{"C.J. McCollum"},
		{"Luis Montero"},
		{"Mason Plumlee"},
		{"Brian Roberts"},
		{"Noah Vonleh"},
		{"Trevor Booker"},
		{"Trey Burke"},
		{"Alec Burks"},
		{"Dante Exum"},
		{"Derrick Favors"},
		{"Rudy Gobert"},
		{"Gordon Hayward"},
		{"Rodney Hood"},
		{"Joe Ingles"},
		{"Chris Johnson"},
		{"Trey Lyles"},
		{"Shelvin Mack"},
		{"Raul Neto"},
		{"Tibor Pleiss"},
		{"Jeff Withey"},
	}
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].ILoc(0, rand.Intn(len(names)))
	}
}

func TestSeriesILoc(t *testing.T) {
	type ilocTest struct {
		arg1     Series
		arg2     int
		arg3     int
		expected []interface{}
	}
	ilocTests := []ilocTest{
		{
			Series{
				[]interface{}{"alice", "bob", "charlie"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"People",
				"string",
			},
			0,
			2,
			[]interface{}{"alice", "bob"},
		},
		{
			Series{
				[]interface{}{"apple", "banana", "cherry"},
				IndexData{
					[]Index{
						{0, []interface{}{"a"}},
						{1, []interface{}{"b"}},
						{2, []interface{}{"c"}},
					},
					[]string{""},
				},
				"Fruit",
				"string",
			},
			0,
			1,
			[]interface{}{"apple"},
		},
	}

	for _, test := range ilocTests {
		output, err := test.arg1.ILoc(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSeriesCount(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].Count()
	}
}

func TestSeriesCount(t *testing.T) {
	type countTest struct {
		arg1     Series
		expected StatsResult
	}
	countTests := []countTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Count",
				4.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Count",
				3.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Count",
				0.0,
				nil,
			},
		},
	}
	for _, test := range countTests {
		output := test.arg1.Count()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors(), cmpopts.EquateNaNs()) {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func BenchmarkSeriesMean(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Mean()
	}
}

func TestSeriesMean(t *testing.T) {
	type meanTest struct {
		arg1     Series
		expected StatsResult
	}
	meanTests := []meanTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Mean",
				math.NaN(),
				fmt.Errorf("data is not float64: Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Mean",
				24.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Mean",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
	}

	for _, test := range meanTests {
		output := test.arg1.Mean()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesMedian(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Median()
	}
}

func TestSeriesMedian(t *testing.T) {
	type medianTest struct {
		arg1     Series
		expected StatsResult
	}
	medianTests := []medianTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Median",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Median",
				23.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Median",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Median",
				175.85,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Median",
				178.7,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Median",
				171.5,
				nil,
			},
		},
	}

	for _, test := range medianTests {
		output := test.arg1.Median()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesStd(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Std()
	}
}

func TestSeriesStd(t *testing.T) {
	type stdTest struct {
		arg1     Series
		expected StatsResult
	}
	stdTests := []stdTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Std",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Std",
				5.568,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Std",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Std",
				7.913,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Std",
				9.601,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Std",
				10.182,
				nil,
			},
		},
	}
	for _, test := range stdTests {
		output := test.arg1.Std()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesMin(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Min()
	}
}

func TestSeriesMin(t *testing.T) {
	type minTest struct {
		arg1     Series
		expected StatsResult
	}
	minTests := []minTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Min",
				math.NaN(),
				fmt.Errorf("data is not a float64: %v", "Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Min",
				19.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Min",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Min",
				164.3,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Min",
				164.3,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Min",
				164.3,
				nil,
			},
		},
	}
	for _, test := range minTests {
		output := test.arg1.Min()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesMax(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Max()
	}
}

func TestSeriesMax(t *testing.T) {
	type maxTest struct {
		arg1     Series
		expected StatsResult
	}
	maxTests := []maxTest{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Max",
				math.NaN(),
				fmt.Errorf("data is not a float64: %v", "Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Max",
				30.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Max",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Max",
				182.5,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Max",
				182.5,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Max",
				178.7,
				nil,
			},
		},
	}
	for _, test := range maxTests {
		output := test.arg1.Max()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesQ1(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Q1()
	}
}

func TestSeriesQ1(t *testing.T) {
	type q1Test struct {
		arg1     Series
		expected StatsResult
	}
	q1Tests := []q1Test{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Q1",
				math.NaN(),
				fmt.Errorf("data is not a float64: %v", "Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Q1",
				19.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Q1",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q1",
				168.65,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q1",
				164.3,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q1",
				164.3,
				nil,
			},
		},
	}
	for _, test := range q1Tests {
		output := test.arg1.Q1()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesQ2(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Q2()
	}
}

func TestSeriesQ2(t *testing.T) {
	type q2Test struct {
		arg1     Series
		expected StatsResult
	}
	q2Tests := []q2Test{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Q2",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Q2",
				23.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Q2",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q2",
				175.85,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q2",
				178.7,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q2",
				171.5,
				nil,
			},
		},
	}
	for _, test := range q2Tests {
		output := test.arg1.Q2()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesQ3(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Q3()
	}
}

func TestSeriesQ3(t *testing.T) {
	type q3Test struct {
		arg1     Series
		expected StatsResult
	}
	q3Tests := []q3Test{
		{
			Series{
				[]interface{}{"Avery", "Bradley", "Candice", "Diana"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Name",
				"string",
			},
			StatsResult{
				"Q3",
				math.NaN(),
				fmt.Errorf("data is not float64: %v", "Avery"),
			},
		},
		{
			Series{
				[]interface{}{30.0, 23.0, 19.0},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Age",
				"float64",
			},
			StatsResult{
				"Q3",
				30.0,
				nil,
			},
		},
		{
			Series{
				[]interface{}{},
				IndexData{},
				"Empty",
				"",
			},
			StatsResult{
				"Q3",
				math.NaN(),
				fmt.Errorf("no elements in this column"),
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, 173.0, 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q3",
				180.6,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, 182.5, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q3",
				182.5,
				nil,
			},
		},
		{
			Series{
				[]interface{}{164.3, math.NaN(), 178.7},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"Height",
				"float64",
			},
			StatsResult{
				"Q3",
				178.7,
				nil,
			},
		},
	}
	for _, test := range q3Tests {
		output := test.arg1.Q3()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}), cmpopts.EquateErrors()) {
			if math.IsNaN(output.Result) {
				continue
			} else {
				t.Fatalf("expected %v, got %v", test.expected, output)
			}
		}
	}
}

func BenchmarkSeriesDescribe(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[4].Describe()
	}
}

func TestSeriesDescribe(t *testing.T) {
	type describeTest struct {
		arg1     Series
		expected []StatsResult
	}
	describeTests := []describeTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			nil,
		},
		{
			Series{
				[]interface{}{123.123, 456.456, 789.789},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"float64",
			},
			[]StatsResult{
				{"Count", 3, nil},
				{"Mean", 456.456, nil},
				{"Median", 456.456, nil},
				{"Std", 333.333, nil},
				{"Min", 123.123, nil},
				{"Max", 789.789, nil},
				{"Q1", 123.123, nil},
				{"Q2", 456.456, nil},
				{"Q3", 789.789, nil},
			},
		},
	}
	for _, test := range describeTests {
		output, err := test.arg1.Describe()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSeriesValueCounts(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].ValueCounts()
	}
}

func TestSeriesValueCounts(t *testing.T) {
	type valueCountsTest struct {
		arg1     Series
		expected Series
	}
	valueCountsTests := []valueCountsTest{
		{
			Series{
				[]interface{}{"Amazon", "Amazon", "Google", "Apple", "Apple", "Apple", "Facebook"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
						{4, []interface{}{4}},
						{5, []interface{}{5}},
						{6, []interface{}{6}},
					},
					[]string{"RangeIndex"},
				},
				"Workplaces",
				"string",
			},
			Series{
				[]interface{}{2, 3, 1, 1},
				IndexData{
					[]Index{
						{0, []interface{}{"Amazon"}},
						{1, []interface{}{"Apple"}},
						{2, []interface{}{"Facebook"}},
						{3, []interface{}{"Google"}},
					},
					[]string{"Data"},
				},
				"Unique Value Count of Workplaces",
				"int",
			},
		},
	}
	for _, test := range valueCountsTests {
		output, err := test.arg1.ValueCounts()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{}), cmpopts.IgnoreFields(Index{}, "id")) || (!cmp.Equal(output, Series{}, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkSeriesRenameCol(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].RenameCol("NewName")
	}
}

func TestSeriesRenameCol(t *testing.T) {
	type renameColTest struct {
		arg1     Series
		arg2     string
		expected Series
	}
	renameColTests := []renameColTest{
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
			),
			"Names",
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Names",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
			),
		},
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
			"Names",
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Names",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
		},
	}

	for _, test := range renameColTests {
		test.arg1.RenameCol(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func BenchmarkSeriesRenameIndex(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].RenameIndex(map[string]string{"Name": "NewName"})
	}
}

func TestSeriesRenameIndex(t *testing.T) {
	type renameIndexTest struct {
		arg1     Series
		arg2     map[string]string
		expected Series
	}
	renameIndexTests := []renameIndexTest{
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"}},
			),
			map[string]string{"Name": "Names"},
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Names"},
				},
			),
		},
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
			map[string]string{"Name": "Names"},
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Names", "Age"},
				},
			),
		},
		{
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
			map[string]string{"Names": "Name"},
			func(data []interface{}, name string, index *IndexData) Series {
				newSer, err := NewSeries(data, name, index)
				if err != nil {
					t.Error(err)
				}
				return newSer
			}(
				[]interface{}{"Avery", "Bradley", "Candice"},
				"Name",
				&IndexData{
					[]Index{{0, []interface{}{"Avery", 13}}, {1, []interface{}{"Bradley", 25}}, {2, []interface{}{"Candice", 29}}},
					[]string{"Name", "Age"},
				},
			),
		},
	}

	for _, test := range renameIndexTests {
		err := test.arg1.RenameIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) || err != nil {
			if fmt.Sprint(err)[:21] == "index does not exist:" {
				continue
			}
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func BenchmarkSeriesSortByIndex(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[1].SortByIndex(true)
	}
}

func TestSeriesSortByIndex(t *testing.T) {
	type sortByIndexTest struct {
		arg1     Series
		arg2     bool
		expected Series
	}
	sortByIndexTests := []sortByIndexTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col2",
				"string",
			},
			true,
			Series{
				[]interface{}{"d", "a", "c", "b"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col2",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col3",
				"string",
			},
			false,
			Series{
				[]interface{}{"b", "c", "a", "d"},
				IndexData{
					[]Index{
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{1, []interface{}{1}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col3",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", math.NaN()}},
						{1, []interface{}{"Female", 34}},
						{2, []interface{}{"NaN", 40}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"Emily", "Gina", "Michael", "William", "Chris", "Alice"},
				IndexData{
					[]Index{
						{1, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", math.NaN()}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
						{2, []interface{}{"NaN", 40}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", math.NaN()}},
						{1, []interface{}{"Female", 34}},
						{2, []interface{}{"NaN", 40}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 19}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"Emily", "Gina", "Michael", "William", "Chris", "Alice"},
				IndexData{
					[]Index{
						{1, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", math.NaN()}},
						{3, []interface{}{"Male", 19}},
						{4, []interface{}{"Male", 19}},
						{5, []interface{}{"Male", 50}},
						{2, []interface{}{"NaN", 40}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
		},
	}
	for _, test := range sortByIndexTests {
		test.arg1.SortByIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func BenchmarkSeriesSortByGivenIndex(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].SortByGivenIndex(testDf.index)
	}
}

func TestSeriesSortByGivenIndex(t *testing.T) {
	type sortByGivenIndexTest struct {
		arg1     Series
		arg2     IndexData
		expected Series
	}
	sortByGivenIndexTests := []sortByGivenIndexTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			IndexData{
				[]Index{{0, []interface{}{0}}, {3, []interface{}{3}}, {1, []interface{}{1}}, {2, []interface{}{2}}},
				[]string{""},
			},
			Series{
				[]interface{}{"a", "d", "b", "c"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{3, []interface{}{3}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			IndexData{
				[]Index{{1, []interface{}{1}}, {3, []interface{}{3}}, {2, []interface{}{2}}, {0, []interface{}{0}}},
				[]string{""},
			},
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			IndexData{
				[]Index{{3, []interface{}{3}}, {2, []interface{}{2}}, {1, []interface{}{1}}, {0, []interface{}{0}}},
				[]string{""},
			},
			Series{
				[]interface{}{"b", "c", "a", "d"},
				IndexData{
					[]Index{
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{1, []interface{}{1}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			IndexData{
				[]Index{{3, []interface{}{"Female", 16}}, {4, []interface{}{"Female", 34}}, {0, []interface{}{"Female", 40}}, {1, []interface{}{"Male", 19}}, {2, []interface{}{"Male", 25}}, {5, []interface{}{"Male", 50}}},
				[]string{"Sex", "Age"},
			},
			Series{
				[]interface{}{"Gina", "Emily", "Alice", "Michael", "William", "Chris"},
				IndexData{
					[]Index{
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
		},
	}

	for _, test := range sortByGivenIndexTests {
		test.arg1.SortByGivenIndex(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func BenchmarkSeriesSortByValues(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[1].SortByValues(true)
	}
}

func TestSeriesSortByValues(t *testing.T) {
	type sortByValuesTest struct {
		arg1     Series
		arg2     bool
		expected Series
	}
	sortByValuesTests := []sortByValuesTest{
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"d", "a", "c", "b"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{1, []interface{}{1}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"a", "b", "c", "d"},
				IndexData{
					[]Index{
						{1, []interface{}{1}},
						{3, []interface{}{3}},
						{2, []interface{}{2}},
						{0, []interface{}{0}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
			false,
			Series{
				[]interface{}{"d", "c", "b", "a"},
				IndexData{
					[]Index{
						{0, []interface{}{0}},
						{2, []interface{}{2}},
						{3, []interface{}{3}},
						{1, []interface{}{1}},
					},
					[]string{""},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"Alice", "Chris", "Emily", "Gina", "Michael", "William"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{5, []interface{}{"Male", 50}},
						{4, []interface{}{"Female", 34}},
						{3, []interface{}{"Female", 16}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{"Alice", "NaN", "William", "NaN", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			true,
			Series{
				[]interface{}{"Alice", "Chris", "Emily", "William", "NaN", "NaN"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{5, []interface{}{"Male", 50}},
						{4, []interface{}{"Female", 34}},
						{2, []interface{}{"Male", 25}},
						{1, []interface{}{"Male", 19}},
						{3, []interface{}{"Female", 16}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
		},
		{
			Series{
				[]interface{}{64.0, math.NaN(), 66.1, 57.24, 96.432, math.NaN()},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"float64",
			},
			true,
			Series{
				[]interface{}{57.24, 64.0, 66.1, 96.432, math.NaN(), math.NaN()},
				IndexData{
					[]Index{
						{3, []interface{}{"Female", 16}},
						{0, []interface{}{"Female", 40}},
						{2, []interface{}{"Male", 25}},
						{4, []interface{}{"Female", 34}},
						{1, []interface{}{"Male", 19}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"float64",
			},
		},
	}
	for _, test := range sortByValuesTests {
		test.arg1.SortByValues(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}

func BenchmarkSeriesIndexHasDuplicateValues(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.series[0].IndexHasDuplicateValues()
	}
}

func TestSeriesIndexHasDuplicateValues(t *testing.T) {
	type indexHasDuplicateValuesTest struct {
		arg1     Series
		expected bool
	}
	indexHasDuplicateValuesTests := []indexHasDuplicateValuesTest{
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 16}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			false,
		},
		{
			Series{
				[]interface{}{"Alice", "Michael", "William", "Gina", "Emily", "Chris"},
				IndexData{
					[]Index{
						{0, []interface{}{"Female", 40}},
						{1, []interface{}{"Male", 19}},
						{2, []interface{}{"Male", 25}},
						{3, []interface{}{"Female", 40}},
						{4, []interface{}{"Female", 34}},
						{5, []interface{}{"Male", 50}},
					},
					[]string{"Sex", "Age"},
				},
				"col1",
				"string",
			},
			true,
		},
	}
	for _, test := range indexHasDuplicateValuesTests {
		output, err := test.arg1.IndexHasDuplicateValues()
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(Series{}, IndexData{}, Index{})) || (output != true && err != nil) {
			t.Fatalf("expected %v, got %v", test.expected, test.arg1)
		}
	}
}
