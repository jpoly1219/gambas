package gambas

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestDataFramePrint(t *testing.T) {
	type printTest struct {
		arg1 DataFrame
	}
	printTests := []printTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
		},
	}

	for _, test := range printTests {
		test.arg1.Print()
	}
}

func TestDataFramePrintRange(t *testing.T) {
	type printRangeTest struct {
		arg1 DataFrame
	}
	printRangeTests := []printRangeTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
		},
	}

	for _, test := range printRangeTests {
		test.arg1.PrintRange(1, 2)
	}
}

func TestDataFrameHead(t *testing.T) {
	type headTest struct {
		arg1 DataFrame
	}
	headTests := []headTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
		},
	}

	for _, test := range headTests {
		test.arg1.Head(2)
	}
}

func TestDataFrameTail(t *testing.T) {
	type tailTest struct {
		arg1 DataFrame
	}
	tailTests := []tailTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
		},
	}

	for _, test := range tailTests {
		test.arg1.Tail(3)
	}
}

func BenchmarkDataFrameLocRows(b *testing.B) {
	nbaDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	neoDf, err := ReadCsv("testfiles/neo_v2.csv", []string{"id"})
	if err != nil {
		b.Error(err)
	}
	nbaCol, err := readCsvColIntoData("testfiles/nba.csv", "Name")
	if err != nil {
		b.Error(err)
	}
	neoCol, err := readCsvColIntoData("testfiles/neo_v2.csv", "id")
	if err != nil {
		b.Error(err)
	}
	dfList := []DataFrame{nbaDf, neoDf}
	colList := [][][]interface{}{nbaCol, neoCol}

	for i := 0; i < 2; i++ {
		b.Run(fmt.Sprintf("case_%d", i), func(b *testing.B) {
			df := dfList[i]
			col := colList[i]
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				df.LocRows(col[rand.Intn(len(col))])
			}
		})
	}

}

func TestDataFrameLocRows(t *testing.T) {
	type dataframeLocRowsTest struct {
		arg1     DataFrame
		arg2     [][]interface{}
		expected DataFrame
	}
	dataframeLocRowsTests := []dataframeLocRowsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[][]interface{}{{"Avery"}},
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{19},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}},
							[]string{"Name"},
						},
						"Age",
						"int",
					},
					{
						[]interface{}{"Male"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder"}},
			DataFrame{
				[]Series{
					{
						[]interface{}{"Jae Crowder"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{"Boston Celtics"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Team",
						"string",
					},
					{
						[]interface{}{99.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Number",
						"float64",
					},
					{
						[]interface{}{"SF"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Position",
						"string",
					},
					{
						[]interface{}{25.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"6-6"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Height",
						"string",
					},
					{
						[]interface{}{235.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Weight",
						"float64",
					},
					{
						[]interface{}{"Marquette"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"College",
						"string",
					},
					{
						[]interface{}{6796117.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}},
							[]string{"Name"},
						},
						"Salary",
						"float64",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Jae Crowder"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder"}, {"Avery Bradley"}},
			DataFrame{
				[]Series{
					{
						[]interface{}{"Jae Crowder", "Avery Bradley"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Team",
						"string",
					},
					{
						[]interface{}{99.0, 0.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Number",
						"float64",
					},
					{
						[]interface{}{"SF", "PG"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Position",
						"string",
					},
					{
						[]interface{}{25.0, 25.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"6-6", "6-2"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Height",
						"string",
					},
					{
						[]interface{}{235.0, 180.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Weight",
						"float64",
					},
					{
						[]interface{}{"Marquette", "Texas"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"College",
						"string",
					},
					{
						[]interface{}{6796117.0, 7730337.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
							[]string{"Name"},
						},
						"Salary",
						"float64",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Jae Crowder"}}, {0, []interface{}{"Avery Bradley"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder", 25.0}, {"Avery Bradley", 25.0}},
			DataFrame{
				[]Series{
					{
						[]interface{}{"Jae Crowder", "Avery Bradley"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Team",
						"string",
					},
					{
						[]interface{}{99.0, 0.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Number",
						"float64",
					},
					{
						[]interface{}{"SF", "PG"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Position",
						"string",
					},
					{
						[]interface{}{25.0, 25.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"6-6", "6-2"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Height",
						"string",
					},
					{
						[]interface{}{235.0, 180.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Weight",
						"float64",
					},
					{
						[]interface{}{"Marquette", "Texas"},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"College",
						"string",
					},
					{
						[]interface{}{6796117.0, 7730337.0},
						IndexData{
							[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
							[]string{"Name", "Age"},
						},
						"Salary",
						"float64",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Jae Crowder", 25.0}}, {0, []interface{}{"Avery Bradley", 25.0}}},
					[]string{"Name", "Age"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder", 22.0}, {"Avery Bradley", 25.0}},
			DataFrame{},
		},
	}

	for _, test := range dataframeLocRowsTests {
		output, err := test.arg1.LocRows(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameLocRowsItems(b *testing.B) {
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
		testDf.LocRowsItems(names[rand.Intn(len(names))])
	}
}

func TestDataFrameLocRowsItems(t *testing.T) {
	type dataframeLocRowsItemsTest struct {
		arg1     DataFrame
		arg2     [][]interface{}
		expected [][]interface{}
	}
	dataframeLocRowsItemsTests := []dataframeLocRowsItemsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[][]interface{}{{"Avery"}},
			[][]interface{}{{"Avery", 19, "Male"}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder"}},
			[][]interface{}{{"Jae Crowder", "Boston Celtics", 99.0, "SF", 25.0, "6-6", 235.0, "Marquette", 6796117.0}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder"}, {"Avery Bradley"}},
			[][]interface{}{
				{"Jae Crowder", "Boston Celtics", 99.0, "SF", 25.0, "6-6", 235.0, "Marquette", 6796117.0},
				{"Avery Bradley", "Boston Celtics", 0.0, "PG", 25.0, "6-2", 180.0, "Texas", 7730337.0},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder", 25.0}, {"Avery Bradley", 25.0}},
			[][]interface{}{
				{"Jae Crowder", "Boston Celtics", 99.0, "SF", 25.0, "6-6", 235.0, "Marquette", 6796117.0},
				{"Avery Bradley", "Boston Celtics", 0.0, "PG", 25.0, "6-2", 180.0, "Texas", 7730337.0},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdflocrows1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[][]interface{}{{"Jae Crowder", 22.0}, {"Avery Bradley", 25.0}},
			nil,
		},
	}

	for _, test := range dataframeLocRowsItemsTests {
		output, err := test.arg1.LocRowsItems(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameLocCols(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.LocCols("Name")
	}
}

func TestDataFrameLocCols(t *testing.T) {
	type dataframeLocColsTest struct {
		arg1     DataFrame
		arg2     []string
		expected DataFrame
	}
	dataframeLocColsTests := []dataframeLocColsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]string{"Age"},
			DataFrame{
				[]Series{
					{
						[]interface{}{19, 27, 22},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"int",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccols1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[]string{"Position"},
			DataFrame{
				[]Series{
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Position",
						"string",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Avery Bradley"}},
						{1, []interface{}{"Jae Crowder"}},
						{2, []interface{}{"John Holland"}},
						{3, []interface{}{"R.J. Hunter"}},
					},
					[]string{"Name"},
				},
				[]string{"Position"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccols1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[]string{"Age", "College", "Name"},
			DataFrame{
				[]Series{
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"College",
						"string",
					},
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						IndexData{
							[]Index{
								{0, []interface{}{"Avery Bradley"}},
								{1, []interface{}{"Jae Crowder"}},
								{2, []interface{}{"John Holland"}},
								{3, []interface{}{"R.J. Hunter"}},
							},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"Avery Bradley"}},
						{1, []interface{}{"Jae Crowder"}},
						{2, []interface{}{"John Holland"}},
						{3, []interface{}{"R.J. Hunter"}},
					},
					[]string{"Name"},
				},
				[]string{"Age", "College", "Name"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccols1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[]string{"Avery Bradley"},
			DataFrame{},
		},
	}
	for _, test := range dataframeLocColsTests {
		output, err := test.arg1.LocCols(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameLocColsItems(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.LocColsItems("Name")
	}
}

func TestDataFrameLocColsItems(t *testing.T) {
	type dataframeLocColsItemsTest struct {
		arg1     DataFrame
		arg2     []string
		expected [][]interface{}
	}
	dataframeLocColsItemsTests := []dataframeLocColsItemsTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]string{"Age"},
			[][]interface{}{{19, 27, 22}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccolsitems1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[]string{"Position"},
			[][]interface{}{{"PG", "SF", "SG", "SG"}},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccolsitems1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[]string{"Age", "College", "Name"},
			[][]interface{}{
				{25.0, 25.0, 27.0, 22.0},
				{"Texas", "Marquette", "Boston University", "Georgia State"},
				{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloccolsitems1.csv", []string{"Name", "Age"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[]string{"Avery Bradley"},
			nil,
		},
	}
	for _, test := range dataframeLocColsItemsTests {
		output, err := test.arg1.LocColsItems(test.arg2...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (output != nil && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameLoc(b *testing.B) {
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
		testDf.Loc([]string{"Name"}, names[rand.Intn(len(names))])
	}
}

func TestDataFrameLoc(t *testing.T) {
	type dataframeLocTest struct {
		arg1     DataFrame
		arg2     []string
		arg3     [][]interface{}
		expected DataFrame
	}
	dataframeLocTests := []dataframeLocTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			[]string{"Age"},
			[][]interface{}{{"Bradley"}},
			DataFrame{
				[]Series{
					{
						[]interface{}{27},
						IndexData{
							[]Index{{1, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Age",
						"int",
					},
				},
				IndexData{
					[]Index{{1, []interface{}{"Bradley"}}},
					[]string{"Name"},
				},
				[]string{"Age"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfloc1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			[]string{"Age", "Name"},
			[][]interface{}{{"John Holland"}},
			DataFrame{
				[]Series{
					{
						[]interface{}{27.0},
						IndexData{
							[]Index{{2, []interface{}{"John Holland"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"John Holland"},
						IndexData{
							[]Index{{2, []interface{}{"John Holland"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
				},
				IndexData{
					[]Index{{2, []interface{}{"John Holland"}}},
					[]string{"Name"},
				},
				[]string{"Age", "Name"},
			},
		},
	}
	for _, test := range dataframeLocTests {
		output, err := test.arg1.Loc(test.arg2, test.arg3...)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColAdd(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColAdd("Salary", rand.Float64())
	}
}

func TestDataFrameColAdd(t *testing.T) {
	type colAddTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colAddTests := []colAddTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{24.0, 32.0, 27.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colAddTests {
		output, err := test.arg1.ColAdd(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColSub(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColSub("Salary", rand.Float64())
	}
}

func TestDataFrameColSub(t *testing.T) {
	type colSubTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colSubTests := []colSubTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{14.0, 22.0, 17.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colSubTests {
		output, err := test.arg1.ColSub(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColMul(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColMul("Salary", rand.Float64())
	}
}

func TestDataFrameColMul(t *testing.T) {
	type colMulTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colMulTests := []colMulTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			2.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{38.0, 54.0, 44.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colMulTests {
		output, err := test.arg1.ColMul(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColDiv(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColDiv("Salary", rand.Float64())
	}
}

func TestDataFrameColDiv(t *testing.T) {
	type colDivTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colDivTests := []colDivTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{3.8, 5.4, 4.4},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colDivTests {
		output, err := test.arg1.ColDiv(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColMod(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColMod("Salary", rand.Float64())
	}
}

func TestDataFrameColMod(t *testing.T) {
	type colModTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colModTests := []colModTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			5.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{4.0, 2.0, 2.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colModTests {
		output, err := test.arg1.ColMod(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColGt(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColGt("Salary", rand.Float64())
	}
}

func TestDataFrameColGt(t *testing.T) {
	type colGtTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colGtTests := []colGtTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			25.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{false, true, false},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"bool",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colGtTests {
		output, err := test.arg1.ColGt(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColLt(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColLt("Salary", rand.Float64())
	}
}

func TestDataFrameColLt(t *testing.T) {
	type colLtTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colLtTests := []colLtTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			22.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{true, false, false},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"bool",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colLtTests {
		output, err := test.arg1.ColLt(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameColEq(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.ColEq("Salary", rand.Float64())
	}
}

func TestDataFrameColEq(t *testing.T) {
	type colEqTest struct {
		arg1     DataFrame
		arg2     string
		arg3     float64
		expected DataFrame
	}

	colEqTests := []colEqTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19, 27, 22}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Name",
			5.0,
			DataFrame{},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			19.0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{true, false, false},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"bool",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}
	for _, test := range colEqTests {
		output, err := test.arg1.ColEq(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameNewCol(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	data := make([]interface{}, len(testDf.series[0].data))
	for i := 0; i < len(testDf.series[0].data); i++ {
		data[i] = rand.Float64()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.NewCol("New Column", data)
	}
}

func TestDataFrameNewCol(t *testing.T) {
	type newColTest struct {
		arg1     DataFrame
		arg2     string
		arg3     []interface{}
		expected DataFrame
	}
	newColTests := []newColTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Nationality",
			[]interface{}{"USA", "UK", "Canada"},
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
					{
						[]interface{}{"USA", "UK", "Canada"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Nationality",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex", "Nationality"},
			},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age+5",
			[]interface{}{"", "", ""},
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
					{
						[]interface{}{"", "", ""},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age+5",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex", "Age+5"},
			},
		},
	}
	for _, test := range newColTests {
		output, err := test.arg1.NewCol(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameNewDerivedCol(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.NewDerivedCol("New Column", "Salary")
	}
}

func TestDataFrameNewDerivedCol(t *testing.T) {
	type newDerivedColTest struct {
		arg1     DataFrame
		arg2     string
		arg3     string
		expected DataFrame
	}
	newDerivedColTests := []newDerivedColTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"NewAge",
			"Age",
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"NewAge",
						"float64",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex", "NewAge"},
			},
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"NewCol",
			"Doesn't Exist",
			DataFrame{},
		},
	}
	for _, test := range newDerivedColTests {
		output, err := test.arg1.NewDerivedCol(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || (!cmp.Equal(output, DataFrame{}, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) && err != nil) {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameRenameCol(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.RenameCol(map[string]string{"Name": "New Name", "Age": "New Age"})
	}
}

func TestDataFrameRenameCol(t *testing.T) {
	type renameColTest struct {
		arg1     DataFrame
		arg2     map[string]string
		expected DataFrame
	}
	renameColTests := []renameColTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			map[string]string{"Name": "Names"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Names", "Age", "Sex"}, []string{"Names"}),
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			map[string]string{"Age": "HowOld"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "HowOld", "Sex"}, []string{"Name"}),
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name", "Sex"}),
			map[string]string{"Name": "Names"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Names", "Age", "Sex"}, []string{"Names", "Sex"}),
		},
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name", "Sex"}),
			map[string]string{"Names": "Name"},
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Avery", "Bradley", "Candice"}, {19.0, 27.0, 22.0}, {"Male", "Male", "Female"}}, []string{"Name", "Age", "Sex"}, []string{"Name", "Sex"}),
		},
	}

	for _, test := range renameColTests {
		err := test.arg1.RenameCol(test.arg2)
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			if fmt.Sprint(err)[:22] == "column does not exist:" || fmt.Sprint(err)[:21] == "index does not exist:" {
				continue
			}
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func TestDataFrameMergeDfsVertically(t *testing.T) {
	type mergeDfsVerticallyTest struct {
		arg1     DataFrame
		arg2     DataFrame
		expected DataFrame
	}
	mergeDfsVerticallyTests := []mergeDfsVerticallyTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/mergeDfsVertically/1src.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/mergeDfsVertically/1target.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/mergeDfsVertically/1exp.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/mergeDfsVertically/2src.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/mergeDfsVertically/2target.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/mergeDfsVertically/2exp.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
		},
	}
	for _, test := range mergeDfsVerticallyTests {
		output, err := test.arg1.MergeDfsVertically(test.arg2)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v,\ngot %v,\nerror %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameSortByIndex(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.SortByIndex(true)
	}
}

func TestDataFrameSortByIndex(t *testing.T) {
	type sortByIndexTest struct {
		arg1     DataFrame
		arg2     bool
		expected DataFrame
	}
	sortByIndexTests := []sortByIndexTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Bradley", "Candice", "Avery"}, {27.0, 22.0, 19.0}, {"Male", "Female", "Male"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			true,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Bradley", "Candice"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{19.0, 27.0, 22.0},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Male", "Female"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{2, []interface{}{"Avery"}}, {0, []interface{}{"Bradley"}}, {1, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
		// {
		// 	func() DataFrame {
		// 		newDf, err := ReadCsv("./testfiles/nba.csv", []string{"Team"})
		// 		if err != nil {
		// 			t.Error(err)
		// 		}
		// 		return *newDf
		// 	}(),
		// 	true,
		// 	func() DataFrame {
		// 		newDf, err := ReadCsv("./testfiles/nba.csv", []string{"Team"})
		// 		if err != nil {
		// 			t.Error(err)
		// 		}
		// 		return *newDf
		// 	}(),
		// },
	}
	for _, test := range sortByIndexTests {
		test.arg1.Print()
		err := test.arg1.SortByIndex(test.arg2)
		test.arg1.Print()
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func BenchmarkDataFrameSortByValues(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.SortByValues("Name", true)
	}
}

func TestDataFrameSortByValues(t *testing.T) {
	type sortByValuesTest struct {
		arg1     DataFrame
		arg2     string
		arg3     bool
		expected DataFrame
	}
	sortByValuesTests := []sortByValuesTest{
		{
			func(data [][]interface{}, columns []string, indexCols []string) DataFrame {
				newDf, err := NewDataFrame(data, columns, indexCols)
				if err != nil {
					t.Error(err)
				}
				return newDf
			}([][]interface{}{{"Bradley", "Candice", "Avery"}, {27.0, 22.0, 19.0}, {"Male", "Female", "Male"}}, []string{"Name", "Age", "Sex"}, []string{"Name"}),
			"Age",
			true,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery", "Candice", "Bradley"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{19.0, 22.0, 27.0},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"Male", "Female", "Male"},
						IndexData{
							[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
							[]string{"Name"},
						},
						"Sex",
						"string",
					},
				},
				IndexData{
					[]Index{{2, []interface{}{"Avery"}}, {1, []interface{}{"Candice"}}, {0, []interface{}{"Bradley"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Age", "Sex"},
			},
		},
	}

	for _, test := range sortByValuesTests {
		test.arg1.Print()
		err := test.arg1.SortByValues(test.arg2, test.arg3)
		test.arg1.Print()
		if !cmp.Equal(test.arg1, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{})) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func BenchmarkDataFrameDropNaN(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.DropNaN(0)
	}
}

func TestDataFrameDropNaN(t *testing.T) {
	type dropNaNTest struct {
		arg1     DataFrame
		arg2     int
		expected DataFrame
	}
	dropNaNTests := []dropNaNTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			0,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "R.J. Hunter"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Team",
						"string",
					},
					{
						[]interface{}{0.0, 99.0, 28.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Number",
						"float64",
					},
					{
						[]interface{}{"PG", "SF", "SG"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Position",
						"string",
					},
					{
						[]interface{}{25.0, 25.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Height",
						"string",
					},
					{
						[]interface{}{180.0, 235.0, 185.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Weight",
						"float64",
					},
					{
						[]interface{}{"Texas", "Marquette", "Georgia State"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"College",
						"string",
					},
					{
						[]interface{}{7730337.0, 6796117.0, 1148640.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Salary",
						"float64",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {3, []interface{}{"R.J. Hunter"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Number", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdropnan2.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			1,
			DataFrame{
				[]Series{
					{
						[]interface{}{"Avery Bradley", "Jae Crowder", "John Holland", "R.J. Hunter"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Name",
						"string",
					},
					{
						[]interface{}{"Boston Celtics", "Boston Celtics", "Boston Celtics", "Boston Celtics"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Team",
						"string",
					},
					{
						[]interface{}{"PG", "SF", "SG", "SG"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Position",
						"string",
					},
					{
						[]interface{}{25.0, 25.0, 27.0, 22.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Age",
						"float64",
					},
					{
						[]interface{}{"6-2", "6-6", "6-5", "6-5"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Height",
						"string",
					},
					{
						[]interface{}{180.0, 235.0, 205.0, 185.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Weight",
						"float64",
					},
					{
						[]interface{}{"Texas", "Marquette", "Boston University", "Georgia State"},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"College",
						"string",
					},
					{
						[]interface{}{7730337.0, 6796117.0, 5000000.0, 1148640.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
							[]string{"Name"},
						},
						"Salary",
						"float64",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery Bradley"}}, {1, []interface{}{"Jae Crowder"}}, {2, []interface{}{"John Holland"}}, {3, []interface{}{"R.J. Hunter"}}},
					[]string{"Name"},
				},
				[]string{"Name", "Team", "Position", "Age", "Height", "Weight", "College", "Salary"},
			},
		},
	}

	for _, test := range dropNaNTests {
		output, err := test.arg1.DropNaN(test.arg2)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, test.arg1, err)
		}
	}
}

func BenchmarkDataFramePivot(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.Pivot("Sex", "Height")
	}
}

func TestDataFramePivot(t *testing.T) {
	type pivotTest struct {
		arg1     DataFrame
		arg2     string
		arg3     string
		expected DataFrame
	}
	pivotTests := []pivotTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfpivot1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			"Sex",
			"Height",
			DataFrame{
				[]Series{
					{
						[]interface{}{172.0, 180.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Male",
						"float64",
					},
					{
						[]interface{}{math.NaN(), math.NaN(), 165.0},
						IndexData{
							[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
							[]string{"Name"},
						},
						"Female",
						"float64",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Avery"}}, {1, []interface{}{"Bradley"}}, {2, []interface{}{"Candice"}}},
					[]string{"Name"},
				},
				[]string{"Male", "Female"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfpivot2.csv", []string{"Time"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			"Fruit",
			"Color",
			DataFrame{
				[]Series{
					{
						[]interface{}{"Red", ""},
						IndexData{
							[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
							[]string{"Time"},
						},
						"Apple",
						"string",
					},
					{
						[]interface{}{"Yellow", "Yellow"},
						IndexData{
							[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
							[]string{"Time"},
						},
						"Banana",
						"string",
					},
					{
						[]interface{}{"", "Red"},
						IndexData{
							[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
							[]string{"Time"},
						},
						"Cherry",
						"string",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"12:00"}}, {1, []interface{}{"12:01"}}},
					[]string{"Time"},
				},
				[]string{"Apple", "Banana", "Cherry"},
			},
		},
	}

	for _, test := range pivotTests {
		output, err := test.arg1.Pivot(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFramePivotTable(b *testing.B) {
	testDf, err := ReadCsv("testfiles/nba.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testDf.PivotTable("Team", "Height", "Salary", Mean)
	}
}

func TestDataFramePivotTable(t *testing.T) {
	type pivotTableTest struct {
		arg1     DataFrame
		arg2     string
		arg3     string
		arg4     string
		arg5     StatsFunc
		expected DataFrame
	}
	pivotTableTests := []pivotTableTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/testdfpivottable1.csv", []string{"Name"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			"Team",
			"Height",
			"Salary",
			Mean,
			DataFrame{
				[]Series{
					{
						[]interface{}{math.NaN(), 1500000.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"5-11",
						"float64",
					},
					{
						[]interface{}{6912869.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"5-9",
						"float64",
					},
					{
						[]interface{}{5000000.0, 958633.333},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-10",
						"float64",
					},
					{
						[]interface{}{math.NaN(), 1140240.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-11",
						"float64",
					},
					{
						[]interface{}{4777348.5, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-2",
						"float64",
					},
					{
						[]interface{}{math.NaN(), 2697445.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-3",
						"float64",
					},
					{
						[]interface{}{3431040.0, 817107.5},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-4",
						"float64",
					},
					{
						[]interface{}{1148640.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-5",
						"float64",
					},
					{
						[]interface{}{4272978.5, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-6",
						"float64",
					},
					{
						[]interface{}{3425510.0, 1467660.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-7",
						"float64",
					},
					{
						[]interface{}{1170960.0, 7330732.5},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-8",
						"float64",
					},
					{
						[]interface{}{7284630.0, math.NaN()},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"6-9",
						"float64",
					},
					{
						[]interface{}{2391067.5, 19689000.0},
						IndexData{
							[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
							[]string{"Team"},
						},
						"7-0",
						"float64",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"Boston Celtics"}}, {1, []interface{}{"Brooklyn Nets"}}},
					[]string{"Team"},
				},
				[]string{"5-11", "5-9", "6-10", "6-11", "6-2", "6-3", "6-4", "6-5", "6-6", "6-7", "6-8", "6-9", "7-0"},
			},
		},
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/airquality.csv", []string{"Time"})
				if err != nil {
					t.Error(err)
				}
				return newDf
			}(),
			"location",
			"parameter",
			"value",
			Mean,
			DataFrame{
				[]Series{
					{
						[]interface{}{26.951, 29.374, 29.740},
						IndexData{
							[]Index{{0, []interface{}{"BETR801"}}, {1, []interface{}{"FR04014"}}, {2, []interface{}{"London Westminster"}}},
							[]string{"location"},
						},
						"no2",
						"float64",
					},
					{
						[]interface{}{23.169, math.NaN(), 13.444},
						IndexData{
							[]Index{{0, []interface{}{"BETR801"}}, {1, []interface{}{"FR04014"}}, {2, []interface{}{"London Westminster"}}},
							[]string{"location"},
						},
						"pm25",
						"float64",
					},
				},
				IndexData{
					[]Index{{0, []interface{}{"BETR801"}}, {1, []interface{}{"FR04014"}}, {2, []interface{}{"London Westminster"}}},
					[]string{"location"},
				},
				[]string{"no2", "pm25"},
			},
		},
	}

	for _, test := range pivotTableTests {
		output, err := test.arg1.PivotTable(test.arg2, test.arg3, test.arg4, test.arg5)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

func BenchmarkDataFrameMelt(b *testing.B) {
	testDf, err := ReadCsv("testfiles/airquality.csv", []string{"Name"})
	if err != nil {
		b.Error(err)
	}
	dfPivoted, err := testDf.PivotTable("location", "parameter", "value", Mean)
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		dfPivoted.Melt("parameter", "value")
	}
}

func TestDataFrameMelt(t *testing.T) {
	type meltTest struct {
		arg1     DataFrame
		arg2     string
		arg3     string
		expected DataFrame
	}
	meltTests := []meltTest{
		{
			func() DataFrame {
				newDf, err := ReadCsv("./testfiles/airquality.csv", []string{"Time"})
				if err != nil {
					t.Error(err)
				}
				dfPivoted, _ := newDf.PivotTable("location", "parameter", "value", Mean)
				return dfPivoted
			}(),
			"parameter",
			"value",
			DataFrame{
				[]Series{
					{
						[]interface{}{"BETR801", "BETR801", "FR04014", "FR04014", "London Westminster", "London Westminster"},
						IndexData{
							[]Index{
								{0, []interface{}{"BETR801"}},
								{1, []interface{}{"BETR801"}},
								{2, []interface{}{"FR04014"}},
								{3, []interface{}{"FR04014"}},
								{4, []interface{}{"London Westminster"}},
								{5, []interface{}{"London Westminster"}},
							},
							[]string{"location"},
						},
						"location",
						"string",
					},
					{
						[]interface{}{"no2", "pm25", "no2", "pm25", "no2", "pm25"},
						IndexData{
							[]Index{
								{0, []interface{}{"BETR801"}},
								{1, []interface{}{"BETR801"}},
								{2, []interface{}{"FR04014"}},
								{3, []interface{}{"FR04014"}},
								{4, []interface{}{"London Westminster"}},
								{5, []interface{}{"London Westminster"}},
							},
							[]string{"location"},
						},
						"parameter",
						"string",
					},
					{
						[]interface{}{26.951, 23.169, 29.374, math.NaN(), 29.740, 13.444},
						IndexData{
							[]Index{
								{0, []interface{}{"BETR801"}},
								{1, []interface{}{"BETR801"}},
								{2, []interface{}{"FR04014"}},
								{3, []interface{}{"FR04014"}},
								{4, []interface{}{"London Westminster"}},
								{5, []interface{}{"London Westminster"}},
							},
							[]string{"location"},
						},
						"value",
						"float64",
					},
				},
				IndexData{
					[]Index{
						{0, []interface{}{"BETR801"}},
						{1, []interface{}{"BETR801"}},
						{2, []interface{}{"FR04014"}},
						{3, []interface{}{"FR04014"}},
						{4, []interface{}{"London Westminster"}},
						{5, []interface{}{"London Westminster"}},
					},
					[]string{"location"},
				},
				[]string{"location", "parameter", "value"},
			},
		},
	}
	for _, test := range meltTests {
		output, err := test.arg1.Melt(test.arg2, test.arg3)
		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}), cmpopts.EquateNaNs()) || err != nil {
			t.Fatalf("expected %v, got %v, error %v", test.expected, output, err)
		}
	}
}

// func TestDataFrameGroupBy(t *testing.T) {
// 	type groupByTest struct {
// 		arg1     DataFrame
// 		arg2     []string
// 		expected *GroupBy
// 	}
// 	groupByTests := []groupByTest{
// 		{
// 			func() DataFrame {
// 				df, err := ReadCsv("./testfiles/airquality.csv", []string{"date.utc"})
// 				if err != nil {
// 					t.Error(err)
// 				}
// 				return *df
// 			}(),
// 			[]string{"parameter", "location"},
// 			&GroupBy{
// 				func() *DataFrame {
// 					df, err := ReadCsv("./testfiles/airquality.csv", []string{"date.utc"})
// 					if err != nil {
// 						t.Error(err)
// 					}
// 					return df
// 				}(),
// 				map[string][]interface{}{
// 					"57b88dc64ef3f516950d6b78f19ab0df836341799c104fb6b23dddf99fe44d2690d4a9b1a7c7e7831493fa9421ffe8021dd4e3e6c43c753c587d5d1df4d9e377": {
// 						"2019-06-17 08:00:00+00:00", "2019-06-17 07:00:00+00:00", "2019-06-17 06:00:00+00:00", "2019-06-17 05:00:00+00:00", "2019-06-17 04:00:00+00:00", "2019-06-17 03:00:00+00:00", "2019-06-17 02:00:00+00:00", "2019-06-17 01:00:00+00:00", "2019-06-16 01:00:00+00:00", "2019-06-15 01:00:00+00:00", "2019-06-14 09:00:00+00:00", "2019-06-13 01:00:00+00:00", "2019-06-12 01:00:00+00:00", "2019-06-11 01:00:00+00:00", "2019-06-10 01:00:00+00:00", "2019-06-09 01:00:00+00:00", "2019-06-05 01:00:00+00:00", "2019-06-01 01:00:00+00:00", "2019-05-31 01:00:00+00:00", "2019-05-30 01:00:00+00:00", "2019-05-29 01:00:00+00:00", "2019-05-28 01:00:00+00:00", "2019-05-27 01:00:00+00:00", "2019-05-26 01:00:00+00:00", "2019-05-25 01:00:00+00:00", "2019-05-24 01:00:00+00:00", "2019-05-23 01:00:00+00:00", "2019-05-22 01:00:00+00:00", "2019-05-21 01:00:00+00:00", "2019-05-20 15:00:00+00:00", "2019-05-20 14:00:00+00:00", "2019-05-20 13:00:00+00:00", "2019-05-20 12:00:00+00:00", "2019-05-20 11:00:00+00:00", "2019-05-20 10:00:00+00:00", "2019-05-20 09:00:00+00:00", "2019-05-20 08:00:00+00:00", "2019-05-20 07:00:00+00:00", "2019-05-20 06:00:00+00:00", "2019-05-20 05:00:00+00:00", "2019-05-20 04:00:00+00:00", "2019-05-20 03:00:00+00:00", "2019-05-20 02:00:00+00:00", "2019-05-20 01:00:00+00:00", "2019-05-20 00:00:00+00:00", "2019-05-19 23:00:00+00:00", "2019-05-19 22:00:00+00:00", "2019-05-19 21:00:00+00:00", "2019-05-19 20:00:00+00:00", "2019-05-19 19:00:00+00:00", "2019-05-19 18:00:00+00:00", "2019-05-19 17:00:00+00:00", "2019-05-19 16:00:00+00:00", "2019-05-19 15:00:00+00:00", "2019-05-19 14:00:00+00:00", "2019-05-19 13:00:00+00:00", "2019-05-19 12:00:00+00:00", "2019-05-19 11:00:00+00:00", "2019-05-19 10:00:00+00:00", "2019-05-19 09:00:00+00:00", "2019-05-19 08:00:00+00:00", "2019-05-19 07:00:00+00:00", "2019-05-19 06:00:00+00:00", "2019-05-19 05:00:00+00:00", "2019-05-19 04:00:00+00:00", "2019-05-19 03:00:00+00:00", "2019-05-19 02:00:00+00:00", "2019-05-19 01:00:00+00:00", "2019-05-19 00:00:00+00:00", "2019-05-18 23:00:00+00:00", "2019-05-18 22:00:00+00:00", "2019-05-18 21:00:00+00:00", "2019-05-18 20:00:00+00:00", "2019-05-18 19:00:00+00:00", "2019-05-18 18:00:00+00:00", "2019-05-18 01:00:00+00:00", "2019-05-16 01:00:00+00:00", "2019-05-15 02:00:00+00:00", "2019-05-15 01:00:00+00:00", "2019-05-14 02:00:00+00:00", "2019-05-14 01:00:00+00:00", "2019-05-13 02:00:00+00:00", "2019-05-13 01:00:00+00:00", "2019-05-12 02:00:00+00:00", "2019-05-12 01:00:00+00:00", "2019-05-11 02:00:00+00:00", "2019-05-11 01:00:00+00:00", "2019-05-10 02:00:00+00:00", "2019-05-10 01:00:00+00:00", "2019-05-09 02:00:00+00:00", "2019-05-09 01:00:00+00:00", "2019-05-08 02:00:00+00:00", "2019-05-08 01:00:00+00:00", "2019-05-07 02:00:00+00:00", "2019-05-07 01:00:00+00:00", "2019-05-06 02:00:00+00:00", "2019-05-06 01:00:00+00:00", "2019-05-05 02:00:00+00:00", "2019-05-05 01:00:00+00:00", "2019-05-04 02:00:00+00:00", "2019-05-04 01:00:00+00:00", "2019-05-03 02:00:00+00:00", "2019-05-03 01:00:00+00:00", "2019-05-02 02:00:00+00:00", "2019-05-02 01:00:00+00:00", "2019-05-01 02:00:00+00:00", "2019-05-01 01:00:00+00:00", "2019-04-30 02:00:00+00:00", "2019-04-30 01:00:00+00:00", "2019-04-29 02:00:00+00:00", "2019-04-29 01:00:00+00:00", "2019-04-28 02:00:00+00:00", "2019-04-28 01:00:00+00:00", "2019-04-27 02:00:00+00:00", "2019-04-27 01:00:00+00:00", "2019-04-26 02:00:00+00:00", "2019-04-26 01:00:00+00:00", "2019-04-25 02:00:00+00:00", "2019-04-25 01:00:00+00:00", "2019-04-22 01:00:00+00:00", "2019-04-21 02:00:00+00:00", "2019-04-21 01:00:00+00:00", "2019-04-19 01:00:00+00:00", "2019-04-18 02:00:00+00:00", "2019-04-17 03:00:00+00:00", "2019-04-17 02:00:00+00:00", "2019-04-17 01:00:00+00:00", "2019-04-16 02:00:00+00:00", "2019-04-16 01:00:00+00:00", "2019-04-15 15:00:00+00:00", "2019-04-15 14:00:00+00:00", "2019-04-15 13:00:00+00:00", "2019-04-15 12:00:00+00:00", "2019-04-15 11:00:00+00:00", "2019-04-15 10:00:00+00:00", "2019-04-15 09:00:00+00:00", "2019-04-15 08:00:00+00:00", "2019-04-15 07:00:00+00:00", "2019-04-15 06:00:00+00:00", "2019-04-15 05:00:00+00:00", "2019-04-15 04:00:00+00:00", "2019-04-15 03:00:00+00:00", "2019-04-15 02:00:00+00:00", "2019-04-15 01:00:00+00:00", "2019-04-12 02:00:00+00:00", "2019-04-12 01:00:00+00:00", "2019-04-11 02:00:00+00:00", "2019-04-11 01:00:00+00:00", "2019-04-10 02:00:00+00:00", "2019-04-10 01:00:00+00:00", "2019-04-09 13:00:00+00:00", "2019-04-09 12:00:00+00:00", "2019-04-09 11:00:00+00:00", "2019-04-09 10:00:00+00:00", "2019-04-09 09:00:00+00:00", "2019-04-09 08:00:00+00:00", "2019-04-09 07:00:00+00:00", "2019-04-09 06:00:00+00:00", "2019-04-09 05:00:00+00:00", "2019-04-09 04:00:00+00:00", "2019-04-09 03:00:00+00:00", "2019-04-09 02:00:00+00:00", "2019-04-09 01:00:00+00:00",
// 					},
// 					"4f014d3c9993903073a9e5e6a32a2cb8773f0c285c4900f44e81a309ebb21399ec0998fdea2dc8ae3ce9fdfdaf49cdc12956783e512aaa0c7c145bf311335127": {
// 						"2019-06-21 00:00:00+00:00", "2019-06-20 23:00:00+00:00", "2019-06-20 22:00:00+00:00", "2019-06-20 21:00:00+00:00", "2019-06-20 20:00:00+00:00", "2019-06-20 19:00:00+00:00", "2019-06-20 18:00:00+00:00", "2019-06-20 17:00:00+00:00", "2019-06-20 16:00:00+00:00", "2019-06-20 15:00:00+00:00", "2019-06-20 14:00:00+00:00", "2019-06-20 13:00:00+00:00", "2019-06-19 10:00:00+00:00", "2019-06-19 09:00:00+00:00", "2019-06-18 22:00:00+00:00", "2019-06-18 21:00:00+00:00", "2019-06-18 20:00:00+00:00", "2019-06-18 19:00:00+00:00", "2019-06-18 08:00:00+00:00", "2019-06-18 07:00:00+00:00", "2019-06-18 06:00:00+00:00", "2019-06-18 05:00:00+00:00", "2019-06-18 04:00:00+00:00", "2019-06-18 03:00:00+00:00", "2019-06-18 02:00:00+00:00", "2019-06-18 01:00:00+00:00", "2019-06-18 00:00:00+00:00", "2019-06-17 23:00:00+00:00", "2019-06-17 22:00:00+00:00", "2019-06-17 21:00:00+00:00", "2019-06-17 20:00:00+00:00", "2019-06-17 19:00:00+00:00", "2019-06-17 18:00:00+00:00", "2019-06-17 17:00:00+00:00", "2019-06-17 16:00:00+00:00", "2019-06-17 15:00:00+00:00", "2019-06-17 14:00:00+00:00", "2019-06-17 13:00:00+00:00", "2019-06-17 12:00:00+00:00", "2019-06-17 11:00:00+00:00", "2019-06-17 10:00:00+00:00", "2019-06-17 09:00:00+00:00", "2019-06-17 08:00:00+00:00", "2019-06-17 07:00:00+00:00", "2019-06-17 06:00:00+00:00", "2019-06-17 05:00:00+00:00", "2019-06-17 04:00:00+00:00", "2019-06-17 03:00:00+00:00", "2019-06-17 02:00:00+00:00", "2019-06-17 01:00:00+00:00", "2019-06-17 00:00:00+00:00", "2019-06-16 23:00:00+00:00", "2019-06-16 22:00:00+00:00", "2019-06-16 21:00:00+00:00", "2019-06-16 20:00:00+00:00", "2019-06-16 19:00:00+00:00", "2019-06-16 18:00:00+00:00", "2019-06-16 17:00:00+00:00", "2019-06-16 16:00:00+00:00", "2019-06-16 15:00:00+00:00", "2019-06-16 14:00:00+00:00", "2019-06-16 13:00:00+00:00", "2019-06-16 12:00:00+00:00", "2019-06-16 11:00:00+00:00", "2019-06-16 10:00:00+00:00", "2019-06-16 09:00:00+00:00", "2019-06-16 08:00:00+00:00", "2019-06-16 07:00:00+00:00", "2019-06-16 06:00:00+00:00", "2019-06-16 05:00:00+00:00", "2019-06-16 04:00:00+00:00", "2019-06-16 03:00:00+00:00", "2019-06-16 02:00:00+00:00", "2019-06-16 01:00:00+00:00", "2019-06-16 00:00:00+00:00", "2019-06-15 23:00:00+00:00", "2019-06-15 22:00:00+00:00", "2019-06-15 21:00:00+00:00", "2019-06-15 20:00:00+00:00", "2019-06-15 19:00:00+00:00", "2019-06-15 18:00:00+00:00", "2019-06-15 17:00:00+00:00", "2019-06-15 16:00:00+00:00", "2019-06-15 15:00:00+00:00", "2019-06-15 14:00:00+00:00", "2019-06-15 13:00:00+00:00", "2019-06-15 12:00:00+00:00", "2019-06-15 11:00:00+00:00", "2019-06-15 10:00:00+00:00", "2019-06-15 09:00:00+00:00", "2019-06-15 08:00:00+00:00", "2019-06-15 07:00:00+00:00", "2019-06-15 06:00:00+00:00", "2019-06-15 02:00:00+00:00", "2019-06-15 01:00:00+00:00", "2019-06-15 00:00:00+00:00", "2019-06-14 23:00:00+00:00", "2019-06-14 22:00:00+00:00", "2019-06-14 21:00:00+00:00", "2019-06-14 20:00:00+00:00", "2019-06-14 19:00:00+00:00", "2019-06-14 18:00:00+00:00", "2019-06-14 17:00:00+00:00", "2019-06-14 16:00:00+00:00", "2019-06-14 15:00:00+00:00", "2019-06-14 14:00:00+00:00", "2019-06-14 13:00:00+00:00", "2019-06-14 12:00:00+00:00", "2019-06-14 11:00:00+00:00", "2019-06-14 10:00:00+00:00", "2019-06-14 09:00:00+00:00", "2019-06-14 08:00:00+00:00", "2019-06-14 07:00:00+00:00", "2019-06-14 06:00:00+00:00", "2019-06-14 05:00:00+00:00", "2019-06-14 04:00:00+00:00", "2019-06-14 03:00:00+00:00", "2019-06-14 02:00:00+00:00", "2019-06-14 01:00:00+00:00", "2019-06-14 00:00:00+00:00", "2019-06-13 23:00:00+00:00", "2019-06-13 22:00:00+00:00", "2019-06-13 21:00:00+00:00", "2019-06-13 20:00:00+00:00", "2019-06-13 19:00:00+00:00", "2019-06-13 18:00:00+00:00", "2019-06-13 17:00:00+00:00", "2019-06-13 16:00:00+00:00", "2019-06-13 15:00:00+00:00", "2019-06-13 14:00:00+00:00", "2019-06-13 13:00:00+00:00", "2019-06-13 12:00:00+00:00", "2019-06-13 11:00:00+00:00", "2019-06-13 10:00:00+00:00", "2019-06-13 09:00:00+00:00", "2019-06-13 08:00:00+00:00", "2019-06-13 07:00:00+00:00", "2019-06-13 06:00:00+00:00", "2019-06-13 05:00:00+00:00", "2019-06-13 04:00:00+00:00", "2019-06-13 03:00:00+00:00", "2019-06-13 02:00:00+00:00", "2019-06-13 01:00:00+00:00", "2019-06-13 00:00:00+00:00", "2019-06-12 23:00:00+00:00", "2019-06-12 22:00:00+00:00", "2019-06-12 21:00:00+00:00", "2019-06-12 20:00:00+00:00", "2019-06-12 19:00:00+00:00", "2019-06-12 18:00:00+00:00", "2019-06-12 17:00:00+00:00", "2019-06-12 16:00:00+00:00", "2019-06-12 15:00:00+00:00", "2019-06-12 14:00:00+00:00", "2019-06-12 13:00:00+00:00", "2019-06-12 12:00:00+00:00", "2019-06-12 11:00:00+00:00", "2019-06-12 10:00:00+00:00", "2019-06-12 09:00:00+00:00", "2019-06-12 08:00:00+00:00", "2019-06-12 07:00:00+00:00", "2019-06-12 06:00:00+00:00", "2019-06-12 05:00:00+00:00", "2019-06-12 04:00:00+00:00", "2019-06-12 03:00:00+00:00", "2019-06-12 02:00:00+00:00", "2019-06-12 01:00:00+00:00", "2019-06-12 00:00:00+00:00", "2019-06-11 23:00:00+00:00", "2019-06-11 22:00:00+00:00", "2019-06-11 21:00:00+00:00", "2019-06-11 20:00:00+00:00", "2019-06-11 19:00:00+00:00", "2019-06-11 18:00:00+00:00", "2019-06-11 17:00:00+00:00", "2019-06-11 16:00:00+00:00", "2019-06-11 15:00:00+00:00", "2019-06-11 14:00:00+00:00", "2019-06-11 13:00:00+00:00", "2019-06-11 12:00:00+00:00", "2019-06-11 11:00:00+00:00", "2019-06-11 10:00:00+00:00", "2019-06-11 09:00:00+00:00", "2019-06-11 08:00:00+00:00", "2019-06-11 07:00:00+00:00", "2019-06-11 06:00:00+00:00", "2019-06-11 05:00:00+00:00", "2019-06-11 04:00:00+00:00", "2019-06-11 03:00:00+00:00", "2019-06-11 02:00:00+00:00", "2019-06-11 01:00:00+00:00", "2019-06-11 00:00:00+00:00", "2019-06-10 23:00:00+00:00", "2019-06-10 22:00:00+00:00", "2019-06-10 21:00:00+00:00", "2019-06-10 20:00:00+00:00", "2019-06-10 19:00:00+00:00", "2019-06-10 18:00:00+00:00", "2019-06-10 17:00:00+00:00", "2019-06-10 16:00:00+00:00", "2019-06-10 15:00:00+00:00", "2019-06-10 14:00:00+00:00", "2019-06-10 13:00:00+00:00", "2019-06-10 12:00:00+00:00", "2019-06-10 11:00:00+00:00", "2019-06-10 10:00:00+00:00", "2019-06-10 09:00:00+00:00", "2019-06-10 08:00:00+00:00", "2019-06-10 07:00:00+00:00", "2019-06-10 06:00:00+00:00", "2019-06-10 05:00:00+00:00", "2019-06-10 04:00:00+00:00", "2019-06-10 03:00:00+00:00", "2019-06-10 02:00:00+00:00", "2019-06-10 01:00:00+00:00", "2019-06-10 00:00:00+00:00", "2019-06-09 23:00:00+00:00", "2019-06-09 22:00:00+00:00", "2019-06-09 21:00:00+00:00", "2019-06-09 20:00:00+00:00", "2019-06-09 19:00:00+00:00", "2019-06-09 18:00:00+00:00", "2019-06-09 17:00:00+00:00", "2019-06-09 16:00:00+00:00", "2019-06-09 15:00:00+00:00", "2019-06-09 14:00:00+00:00", "2019-06-09 13:00:00+00:00", "2019-06-09 12:00:00+00:00", "2019-06-09 11:00:00+00:00", "2019-06-09 10:00:00+00:00", "2019-06-09 09:00:00+00:00", "2019-06-09 08:00:00+00:00", "2019-06-09 07:00:00+00:00", "2019-06-09 06:00:00+00:00", "2019-06-09 05:00:00+00:00", "2019-06-09 04:00:00+00:00", "2019-06-09 03:00:00+00:00", "2019-06-09 02:00:00+00:00", "2019-06-09 01:00:00+00:00", "2019-06-09 00:00:00+00:00", "2019-06-08 23:00:00+00:00", "2019-06-08 22:00:00+00:00", "2019-06-08 21:00:00+00:00", "2019-06-08 18:00:00+00:00", "2019-06-08 17:00:00+00:00", "2019-06-08 16:00:00+00:00", "2019-06-08 15:00:00+00:00", "2019-06-08 14:00:00+00:00", "2019-06-08 13:00:00+00:00", "2019-06-08 12:00:00+00:00", "2019-06-08 11:00:00+00:00", "2019-06-08 10:00:00+00:00", "2019-06-08 09:00:00+00:00", "2019-06-08 08:00:00+00:00", "2019-06-08 07:00:00+00:00", "2019-06-08 06:00:00+00:00", "2019-06-08 05:00:00+00:00", "2019-06-08 04:00:00+00:00", "2019-06-08 03:00:00+00:00", "2019-06-08 02:00:00+00:00", "2019-06-08 01:00:00+00:00", "2019-06-08 00:00:00+00:00", "2019-06-07 23:00:00+00:00", "2019-06-07 22:00:00+00:00", "2019-06-07 21:00:00+00:00", "2019-06-07 20:00:00+00:00", "2019-06-07 19:00:00+00:00", "2019-06-07 18:00:00+00:00", "2019-06-07 17:00:00+00:00", "2019-06-07 16:00:00+00:00", "2019-06-07 15:00:00+00:00", "2019-06-07 14:00:00+00:00", "2019-06-07 13:00:00+00:00", "2019-06-07 12:00:00+00:00", "2019-06-07 11:00:00+00:00", "2019-06-07 10:00:00+00:00", "2019-06-07 09:00:00+00:00", "2019-06-07 08:00:00+00:00", "2019-06-07 07:00:00+00:00", "2019-06-07 06:00:00+00:00", "2019-06-06 14:00:00+00:00", "2019-06-06 13:00:00+00:00", "2019-06-06 12:00:00+00:00", "2019-06-06 11:00:00+00:00", "2019-06-06 10:00:00+00:00", "2019-06-06 09:00:00+00:00", "2019-06-06 08:00:00+00:00", "2019-06-06 07:00:00+00:00", "2019-06-06 06:00:00+00:00", "2019-06-06 05:00:00+00:00", "2019-06-06 04:00:00+00:00", "2019-06-06 03:00:00+00:00", "2019-06-06 02:00:00+00:00", "2019-06-06 01:00:00+00:00", "2019-06-06 00:00:00+00:00", "2019-06-05 23:00:00+00:00", "2019-06-05 22:00:00+00:00", "2019-06-05 21:00:00+00:00", "2019-06-05 20:00:00+00:00", "2019-06-05 19:00:00+00:00", "2019-06-05 18:00:00+00:00", "2019-06-05 17:00:00+00:00", "2019-06-05 16:00:00+00:00", "2019-06-05 15:00:00+00:00", "2019-06-05 14:00:00+00:00", "2019-06-05 13:00:00+00:00", "2019-06-05 12:00:00+00:00", "2019-06-05 11:00:00+00:00", "2019-06-05 10:00:00+00:00", "2019-06-05 09:00:00+00:00", "2019-06-05 08:00:00+00:00", "2019-06-05 07:00:00+00:00", "2019-06-05 06:00:00+00:00", "2019-06-05 05:00:00+00:00", "2019-06-05 04:00:00+00:00", "2019-06-05 03:00:00+00:00", "2019-06-05 02:00:00+00:00", "2019-06-05 01:00:00+00:00", "2019-06-05 00:00:00+00:00", "2019-06-04 23:00:00+00:00", "2019-06-04 22:00:00+00:00", "2019-06-04 21:00:00+00:00", "2019-06-04 20:00:00+00:00", "2019-06-04 19:00:00+00:00", "2019-06-04 18:00:00+00:00", "2019-06-04 17:00:00+00:00", "2019-06-04 16:00:00+00:00", "2019-06-04 15:00:00+00:00", "2019-06-04 14:00:00+00:00", "2019-06-04 13:00:00+00:00", "2019-06-04 12:00:00+00:00", "2019-06-04 11:00:00+00:00", "2019-06-04 10:00:00+00:00", "2019-06-04 09:00:00+00:00", "2019-06-04 08:00:00+00:00", "2019-06-04 07:00:00+00:00", "2019-06-04 06:00:00+00:00", "2019-06-04 05:00:00+00:00", "2019-06-04 04:00:00+00:00", "2019-06-04 03:00:00+00:00", "2019-06-04 02:00:00+00:00", "2019-06-04 01:00:00+00:00", "2019-06-04 00:00:00+00:00", "2019-06-03 23:00:00+00:00", "2019-06-03 22:00:00+00:00", "2019-06-03 21:00:00+00:00", "2019-06-03 20:00:00+00:00", "2019-06-03 19:00:00+00:00", "2019-06-03 18:00:00+00:00", "2019-06-03 17:00:00+00:00", "2019-06-03 16:00:00+00:00", "2019-06-03 15:00:00+00:00", "2019-06-03 14:00:00+00:00", "2019-06-03 13:00:00+00:00", "2019-06-03 12:00:00+00:00", "2019-06-03 11:00:00+00:00", "2019-06-03 10:00:00+00:00", "2019-06-03 09:00:00+00:00", "2019-06-03 08:00:00+00:00", "2019-06-03 07:00:00+00:00", "2019-06-03 06:00:00+00:00", "2019-06-03 05:00:00+00:00", "2019-06-03 04:00:00+00:00", "2019-06-03 03:00:00+00:00", "2019-06-03 02:00:00+00:00", "2019-06-03 01:00:00+00:00", "2019-06-03 00:00:00+00:00", "2019-06-02 23:00:00+00:00", "2019-06-02 22:00:00+00:00", "2019-06-02 21:00:00+00:00", "2019-06-02 20:00:00+00:00", "2019-06-02 19:00:00+00:00", "2019-06-02 18:00:00+00:00", "2019-06-02 17:00:00+00:00", "2019-06-02 16:00:00+00:00", "2019-06-02 15:00:00+00:00", "2019-06-02 14:00:00+00:00", "2019-06-02 13:00:00+00:00", "2019-06-02 12:00:00+00:00", "2019-06-02 11:00:00+00:00", "2019-06-02 10:00:00+00:00", "2019-06-02 09:00:00+00:00", "2019-06-02 08:00:00+00:00", "2019-06-02 07:00:00+00:00", "2019-06-02 06:00:00+00:00", "2019-06-02 05:00:00+00:00", "2019-06-02 04:00:00+00:00", "2019-06-02 03:00:00+00:00", "2019-06-02 02:00:00+00:00", "2019-06-02 01:00:00+00:00", "2019-06-02 00:00:00+00:00", "2019-06-01 23:00:00+00:00", "2019-06-01 22:00:00+00:00", "2019-06-01 21:00:00+00:00", "2019-06-01 20:00:00+00:00", "2019-06-01 19:00:00+00:00", "2019-06-01 18:00:00+00:00", "2019-06-01 17:00:00+00:00", "2019-06-01 16:00:00+00:00", "2019-06-01 15:00:00+00:00", "2019-06-01 14:00:00+00:00", "2019-06-01 13:00:00+00:00", "2019-06-01 12:00:00+00:00", "2019-06-01 11:00:00+00:00", "2019-06-01 10:00:00+00:00", "2019-06-01 09:00:00+00:00", "2019-06-01 08:00:00+00:00", "2019-06-01 07:00:00+00:00", "2019-06-01 06:00:00+00:00", "2019-06-01 02:00:00+00:00", "2019-06-01 01:00:00+00:00", "2019-06-01 00:00:00+00:00", "2019-05-31 23:00:00+00:00", "2019-05-31 22:00:00+00:00", "2019-05-31 21:00:00+00:00", "2019-05-31 20:00:00+00:00", "2019-05-31 19:00:00+00:00", "2019-05-31 18:00:00+00:00", "2019-05-31 17:00:00+00:00", "2019-05-31 16:00:00+00:00", "2019-05-31 15:00:00+00:00", "2019-05-31 14:00:00+00:00", "2019-05-31 13:00:00+00:00", "2019-05-31 12:00:00+00:00", "2019-05-31 11:00:00+00:00", "2019-05-31 10:00:00+00:00", "2019-05-31 09:00:00+00:00", "2019-05-31 08:00:00+00:00", "2019-05-31 07:00:00+00:00", "2019-05-31 06:00:00+00:00", "2019-05-31 05:00:00+00:00", "2019-05-31 04:00:00+00:00", "2019-05-31 03:00:00+00:00", "2019-05-31 02:00:00+00:00", "2019-05-31 01:00:00+00:00", "2019-05-31 00:00:00+00:00", "2019-05-30 23:00:00+00:00", "2019-05-30 22:00:00+00:00", "2019-05-30 21:00:00+00:00", "2019-05-30 20:00:00+00:00", "2019-05-30 19:00:00+00:00", "2019-05-30 18:00:00+00:00", "2019-05-30 17:00:00+00:00", "2019-05-30 16:00:00+00:00", "2019-05-30 15:00:00+00:00", "2019-05-30 14:00:00+00:00", "2019-05-30 13:00:00+00:00", "2019-05-30 12:00:00+00:00", "2019-05-30 11:00:00+00:00", "2019-05-30 10:00:00+00:00", "2019-05-30 09:00:00+00:00", "2019-05-30 08:00:00+00:00", "2019-05-30 07:00:00+00:00", "2019-05-30 06:00:00+00:00", "2019-05-30 05:00:00+00:00", "2019-05-30 04:00:00+00:00", "2019-05-30 03:00:00+00:00", "2019-05-30 02:00:00+00:00", "2019-05-30 01:00:00+00:00", "2019-05-30 00:00:00+00:00", "2019-05-29 23:00:00+00:00", "2019-05-29 22:00:00+00:00", "2019-05-29 21:00:00+00:00", "2019-05-29 20:00:00+00:00", "2019-05-29 19:00:00+00:00", "2019-05-29 18:00:00+00:00", "2019-05-29 17:00:00+00:00", "2019-05-29 16:00:00+00:00", "2019-05-29 15:00:00+00:00", "2019-05-29 14:00:00+00:00", "2019-05-29 13:00:00+00:00", "2019-05-29 12:00:00+00:00", "2019-05-29 11:00:00+00:00", "2019-05-29 10:00:00+00:00", "2019-05-29 09:00:00+00:00", "2019-05-29 08:00:00+00:00", "2019-05-29 07:00:00+00:00", "2019-05-29 06:00:00+00:00", "2019-05-29 05:00:00+00:00", "2019-05-29 04:00:00+00:00", "2019-05-29 03:00:00+00:00", "2019-05-29 02:00:00+00:00", "2019-05-29 01:00:00+00:00", "2019-05-29 00:00:00+00:00", "2019-05-28 23:00:00+00:00", "2019-05-28 22:00:00+00:00", "2019-05-28 21:00:00+00:00", "2019-05-28 20:00:00+00:00", "2019-05-28 19:00:00+00:00", "2019-05-28 18:00:00+00:00", "2019-05-28 17:00:00+00:00", "2019-05-28 16:00:00+00:00", "2019-05-28 15:00:00+00:00", "2019-05-28 14:00:00+00:00", "2019-05-28 13:00:00+00:00", "2019-05-28 12:00:00+00:00", "2019-05-28 11:00:00+00:00", "2019-05-28 10:00:00+00:00", "2019-05-28 09:00:00+00:00", "2019-05-28 08:00:00+00:00", "2019-05-28 07:00:00+00:00", "2019-05-28 06:00:00+00:00", "2019-05-28 05:00:00+00:00", "2019-05-28 04:00:00+00:00", "2019-05-28 03:00:00+00:00", "2019-05-28 02:00:00+00:00", "2019-05-28 01:00:00+00:00", "2019-05-28 00:00:00+00:00", "2019-05-27 23:00:00+00:00", "2019-05-27 22:00:00+00:00", "2019-05-27 21:00:00+00:00", "2019-05-27 20:00:00+00:00", "2019-05-27 19:00:00+00:00", "2019-05-27 18:00:00+00:00", "2019-05-27 17:00:00+00:00", "2019-05-27 16:00:00+00:00", "2019-05-27 15:00:00+00:00", "2019-05-27 14:00:00+00:00", "2019-05-27 13:00:00+00:00", "2019-05-27 12:00:00+00:00", "2019-05-27 11:00:00+00:00", "2019-05-27 10:00:00+00:00", "2019-05-27 09:00:00+00:00", "2019-05-27 08:00:00+00:00", "2019-05-27 07:00:00+00:00", "2019-05-27 06:00:00+00:00", "2019-05-27 05:00:00+00:00", "2019-05-27 04:00:00+00:00", "2019-05-27 03:00:00+00:00", "2019-05-27 02:00:00+00:00", "2019-05-27 01:00:00+00:00", "2019-05-27 00:00:00+00:00", "2019-05-26 23:00:00+00:00", "2019-05-26 22:00:00+00:00", "2019-05-26 21:00:00+00:00", "2019-05-26 20:00:00+00:00", "2019-05-26 19:00:00+00:00", "2019-05-26 18:00:00+00:00", "2019-05-26 17:00:00+00:00", "2019-05-26 16:00:00+00:00", "2019-05-26 15:00:00+00:00", "2019-05-26 14:00:00+00:00", "2019-05-26 13:00:00+00:00", "2019-05-26 12:00:00+00:00", "2019-05-26 11:00:00+00:00", "2019-05-26 10:00:00+00:00", "2019-05-26 09:00:00+00:00", "2019-05-26 08:00:00+00:00", "2019-05-26 07:00:00+00:00", "2019-05-26 06:00:00+00:00", "2019-05-26 05:00:00+00:00", "2019-05-26 04:00:00+00:00", "2019-05-26 03:00:00+00:00", "2019-05-26 02:00:00+00:00", "2019-05-26 01:00:00+00:00", "2019-05-26 00:00:00+00:00", "2019-05-25 23:00:00+00:00", "2019-05-25 22:00:00+00:00", "2019-05-25 21:00:00+00:00", "2019-05-25 20:00:00+00:00", "2019-05-25 19:00:00+00:00", "2019-05-25 18:00:00+00:00", "2019-05-25 17:00:00+00:00", "2019-05-25 16:00:00+00:00", "2019-05-25 15:00:00+00:00", "2019-05-25 14:00:00+00:00", "2019-05-25 13:00:00+00:00", "2019-05-25 12:00:00+00:00", "2019-05-25 11:00:00+00:00", "2019-05-25 10:00:00+00:00", "2019-05-25 09:00:00+00:00", "2019-05-25 08:00:00+00:00", "2019-05-25 07:00:00+00:00", "2019-05-25 06:00:00+00:00", "2019-05-25 02:00:00+00:00", "2019-05-25 01:00:00+00:00", "2019-05-25 00:00:00+00:00", "2019-05-24 23:00:00+00:00", "2019-05-24 22:00:00+00:00", "2019-05-24 21:00:00+00:00", "2019-05-24 20:00:00+00:00", "2019-05-24 19:00:00+00:00", "2019-05-24 18:00:00+00:00", "2019-05-24 17:00:00+00:00", "2019-05-24 16:00:00+00:00", "2019-05-24 15:00:00+00:00", "2019-05-24 14:00:00+00:00", "2019-05-24 13:00:00+00:00", "2019-05-24 12:00:00+00:00", "2019-05-24 11:00:00+00:00", "2019-05-24 10:00:00+00:00", "2019-05-24 09:00:00+00:00", "2019-05-24 08:00:00+00:00", "2019-05-24 07:00:00+00:00", "2019-05-24 06:00:00+00:00", "2019-05-24 05:00:00+00:00", "2019-05-24 04:00:00+00:00", "2019-05-24 03:00:00+00:00", "2019-05-24 02:00:00+00:00", "2019-05-24 01:00:00+00:00", "2019-05-24 00:00:00+00:00", "2019-05-23 23:00:00+00:00", "2019-05-23 22:00:00+00:00", "2019-05-23 21:00:00+00:00", "2019-05-23 20:00:00+00:00", "2019-05-23 19:00:00+00:00", "2019-05-23 18:00:00+00:00", "2019-05-23 17:00:00+00:00", "2019-05-23 16:00:00+00:00", "2019-05-23 15:00:00+00:00", "2019-05-23 14:00:00+00:00", "2019-05-23 13:00:00+00:00", "2019-05-23 12:00:00+00:00", "2019-05-23 11:00:00+00:00", "2019-05-23 10:00:00+00:00", "2019-05-23 09:00:00+00:00", "2019-05-23 08:00:00+00:00", "2019-05-23 07:00:00+00:00", "2019-05-23 06:00:00+00:00", "2019-05-23 05:00:00+00:00", "2019-05-23 04:00:00+00:00", "2019-05-23 03:00:00+00:00", "2019-05-23 02:00:00+00:00", "2019-05-23 01:00:00+00:00", "2019-05-23 00:00:00+00:00", "2019-05-22 23:00:00+00:00", "2019-05-22 22:00:00+00:00", "2019-05-22 21:00:00+00:00", "2019-05-22 20:00:00+00:00", "2019-05-22 19:00:00+00:00", "2019-05-22 18:00:00+00:00", "2019-05-22 17:00:00+00:00", "2019-05-22 16:00:00+00:00", "2019-05-22 15:00:00+00:00", "2019-05-22 14:00:00+00:00", "2019-05-22 13:00:00+00:00", "2019-05-22 12:00:00+00:00", "2019-05-22 11:00:00+00:00", "2019-05-22 10:00:00+00:00", "2019-05-22 09:00:00+00:00", "2019-05-22 08:00:00+00:00", "2019-05-22 07:00:00+00:00", "2019-05-22 06:00:00+00:00", "2019-05-22 05:00:00+00:00", "2019-05-22 04:00:00+00:00", "2019-05-22 03:00:00+00:00", "2019-05-22 02:00:00+00:00", "2019-05-22 01:00:00+00:00", "2019-05-22 00:00:00+00:00", "2019-05-21 23:00:00+00:00", "2019-05-21 22:00:00+00:00", "2019-05-21 21:00:00+00:00", "2019-05-21 20:00:00+00:00", "2019-05-21 19:00:00+00:00", "2019-05-21 18:00:00+00:00", "2019-05-21 17:00:00+00:00", "2019-05-21 16:00:00+00:00", "2019-05-21 15:00:00+00:00", "2019-05-21 14:00:00+00:00", "2019-05-21 13:00:00+00:00", "2019-05-21 12:00:00+00:00", "2019-05-21 11:00:00+00:00", "2019-05-21 10:00:00+00:00", "2019-05-21 09:00:00+00:00", "2019-05-21 08:00:00+00:00", "2019-05-21 07:00:00+00:00", "2019-05-21 06:00:00+00:00", "2019-05-21 05:00:00+00:00", "2019-05-21 04:00:00+00:00", "2019-05-21 03:00:00+00:00", "2019-05-21 02:00:00+00:00", "2019-05-21 01:00:00+00:00", "2019-05-21 00:00:00+00:00", "2019-05-20 23:00:00+00:00", "2019-05-20 22:00:00+00:00", "2019-05-20 21:00:00+00:00", "2019-05-20 20:00:00+00:00", "2019-05-20 19:00:00+00:00", "2019-05-20 18:00:00+00:00", "2019-05-20 17:00:00+00:00", "2019-05-20 16:00:00+00:00", "2019-05-20 15:00:00+00:00", "2019-05-20 14:00:00+00:00", "2019-05-20 13:00:00+00:00", "2019-05-20 12:00:00+00:00", "2019-05-20 11:00:00+00:00", "2019-05-20 10:00:00+00:00", "2019-05-20 09:00:00+00:00", "2019-05-20 08:00:00+00:00", "2019-05-20 07:00:00+00:00", "2019-05-20 06:00:00+00:00", "2019-05-20 05:00:00+00:00", "2019-05-20 04:00:00+00:00", "2019-05-20 03:00:00+00:00", "2019-05-20 02:00:00+00:00", "2019-05-20 01:00:00+00:00", "2019-05-20 00:00:00+00:00", "2019-05-19 23:00:00+00:00", "2019-05-19 22:00:00+00:00", "2019-05-19 21:00:00+00:00", "2019-05-19 20:00:00+00:00", "2019-05-19 19:00:00+00:00", "2019-05-19 18:00:00+00:00", "2019-05-19 17:00:00+00:00", "2019-05-19 16:00:00+00:00", "2019-05-19 15:00:00+00:00", "2019-05-19 14:00:00+00:00", "2019-05-19 13:00:00+00:00", "2019-05-19 12:00:00+00:00", "2019-05-19 11:00:00+00:00", "2019-05-19 10:00:00+00:00", "2019-05-19 09:00:00+00:00", "2019-05-19 08:00:00+00:00", "2019-05-19 07:00:00+00:00", "2019-05-19 06:00:00+00:00", "2019-05-19 05:00:00+00:00", "2019-05-19 04:00:00+00:00", "2019-05-19 03:00:00+00:00", "2019-05-19 02:00:00+00:00", "2019-05-19 01:00:00+00:00", "2019-05-19 00:00:00+00:00", "2019-05-18 23:00:00+00:00", "2019-05-18 22:00:00+00:00", "2019-05-18 21:00:00+00:00", "2019-05-18 20:00:00+00:00", "2019-05-18 19:00:00+00:00", "2019-05-18 18:00:00+00:00", "2019-05-18 17:00:00+00:00", "2019-05-18 16:00:00+00:00", "2019-05-18 15:00:00+00:00", "2019-05-18 14:00:00+00:00", "2019-05-18 13:00:00+00:00", "2019-05-18 12:00:00+00:00", "2019-05-18 11:00:00+00:00", "2019-05-18 10:00:00+00:00", "2019-05-18 09:00:00+00:00", "2019-05-18 08:00:00+00:00", "2019-05-18 07:00:00+00:00", "2019-05-18 06:00:00+00:00", "2019-05-18 05:00:00+00:00", "2019-05-18 04:00:00+00:00", "2019-05-18 03:00:00+00:00", "2019-05-18 02:00:00+00:00", "2019-05-18 01:00:00+00:00", "2019-05-18 00:00:00+00:00", "2019-05-17 23:00:00+00:00", "2019-05-17 22:00:00+00:00", "2019-05-17 21:00:00+00:00", "2019-05-17 20:00:00+00:00", "2019-05-17 19:00:00+00:00", "2019-05-17 18:00:00+00:00", "2019-05-17 17:00:00+00:00", "2019-05-17 16:00:00+00:00", "2019-05-17 15:00:00+00:00", "2019-05-17 14:00:00+00:00", "2019-05-17 13:00:00+00:00", "2019-05-17 12:00:00+00:00", "2019-05-17 11:00:00+00:00", "2019-05-17 10:00:00+00:00", "2019-05-17 09:00:00+00:00", "2019-05-17 08:00:00+00:00", "2019-05-17 07:00:00+00:00", "2019-05-17 06:00:00+00:00", "2019-05-17 05:00:00+00:00", "2019-05-17 04:00:00+00:00", "2019-05-17 03:00:00+00:00", "2019-05-17 02:00:00+00:00", "2019-05-17 01:00:00+00:00", "2019-05-17 00:00:00+00:00", "2019-05-16 23:00:00+00:00", "2019-05-16 22:00:00+00:00", "2019-05-16 21:00:00+00:00", "2019-05-16 20:00:00+00:00", "2019-05-16 19:00:00+00:00", "2019-05-16 18:00:00+00:00", "2019-05-16 17:00:00+00:00", "2019-05-16 16:00:00+00:00", "2019-05-16 15:00:00+00:00", "2019-05-16 14:00:00+00:00", "2019-05-16 13:00:00+00:00", "2019-05-16 12:00:00+00:00", "2019-05-16 11:00:00+00:00", "2019-05-16 10:00:00+00:00", "2019-05-16 09:00:00+00:00", "2019-05-16 08:00:00+00:00", "2019-05-16 07:00:00+00:00", "2019-05-16 05:00:00+00:00", "2019-05-16 04:00:00+00:00", "2019-05-16 03:00:00+00:00", "2019-05-16 02:00:00+00:00", "2019-05-16 01:00:00+00:00", "2019-05-16 00:00:00+00:00", "2019-05-15 23:00:00+00:00", "2019-05-15 22:00:00+00:00", "2019-05-15 21:00:00+00:00", "2019-05-15 20:00:00+00:00", "2019-05-15 19:00:00+00:00", "2019-05-15 18:00:00+00:00", "2019-05-15 17:00:00+00:00", "2019-05-15 16:00:00+00:00", "2019-05-15 15:00:00+00:00", "2019-05-15 14:00:00+00:00", "2019-05-15 13:00:00+00:00", "2019-05-15 12:00:00+00:00", "2019-05-15 11:00:00+00:00", "2019-05-15 10:00:00+00:00", "2019-05-15 09:00:00+00:00", "2019-05-15 08:00:00+00:00", "2019-05-15 07:00:00+00:00", "2019-05-15 06:00:00+00:00", "2019-05-15 05:00:00+00:00", "2019-05-15 04:00:00+00:00", "2019-05-15 03:00:00+00:00", "2019-05-15 02:00:00+00:00", "2019-05-15 01:00:00+00:00", "2019-05-15 00:00:00+00:00", "2019-05-14 23:00:00+00:00", "2019-05-14 22:00:00+00:00", "2019-05-14 21:00:00+00:00", "2019-05-14 20:00:00+00:00", "2019-05-14 19:00:00+00:00", "2019-05-14 18:00:00+00:00", "2019-05-14 17:00:00+00:00", "2019-05-14 16:00:00+00:00", "2019-05-14 15:00:00+00:00", "2019-05-14 14:00:00+00:00", "2019-05-14 13:00:00+00:00", "2019-05-14 12:00:00+00:00", "2019-05-14 11:00:00+00:00", "2019-05-14 10:00:00+00:00", "2019-05-14 09:00:00+00:00", "2019-05-14 08:00:00+00:00", "2019-05-14 07:00:00+00:00", "2019-05-14 06:00:00+00:00", "2019-05-14 05:00:00+00:00", "2019-05-14 04:00:00+00:00", "2019-05-14 03:00:00+00:00", "2019-05-14 02:00:00+00:00", "2019-05-14 01:00:00+00:00", "2019-05-14 00:00:00+00:00", "2019-05-13 23:00:00+00:00", "2019-05-13 22:00:00+00:00", "2019-05-13 21:00:00+00:00", "2019-05-13 20:00:00+00:00", "2019-05-13 19:00:00+00:00", "2019-05-13 18:00:00+00:00", "2019-05-13 17:00:00+00:00", "2019-05-13 16:00:00+00:00", "2019-05-13 15:00:00+00:00", "2019-05-13 14:00:00+00:00", "2019-05-13 13:00:00+00:00", "2019-05-13 12:00:00+00:00", "2019-05-13 11:00:00+00:00", "2019-05-13 10:00:00+00:00", "2019-05-13 09:00:00+00:00", "2019-05-13 08:00:00+00:00", "2019-05-13 07:00:00+00:00", "2019-05-13 06:00:00+00:00", "2019-05-13 05:00:00+00:00", "2019-05-13 04:00:00+00:00", "2019-05-13 03:00:00+00:00", "2019-05-13 02:00:00+00:00", "2019-05-13 01:00:00+00:00", "2019-05-13 00:00:00+00:00", "2019-05-12 23:00:00+00:00", "2019-05-12 22:00:00+00:00", "2019-05-12 21:00:00+00:00", "2019-05-12 20:00:00+00:00", "2019-05-12 19:00:00+00:00", "2019-05-12 18:00:00+00:00", "2019-05-12 17:00:00+00:00", "2019-05-12 16:00:00+00:00", "2019-05-12 15:00:00+00:00", "2019-05-12 14:00:00+00:00", "2019-05-12 13:00:00+00:00", "2019-05-12 12:00:00+00:00", "2019-05-12 11:00:00+00:00", "2019-05-12 10:00:00+00:00", "2019-05-12 09:00:00+00:00", "2019-05-12 08:00:00+00:00", "2019-05-12 07:00:00+00:00", "2019-05-12 06:00:00+00:00", "2019-05-12 05:00:00+00:00", "2019-05-12 04:00:00+00:00", "2019-05-12 03:00:00+00:00", "2019-05-12 02:00:00+00:00", "2019-05-12 01:00:00+00:00", "2019-05-12 00:00:00+00:00", "2019-05-11 23:00:00+00:00", "2019-05-11 22:00:00+00:00", "2019-05-11 21:00:00+00:00", "2019-05-11 20:00:00+00:00", "2019-05-11 19:00:00+00:00", "2019-05-11 18:00:00+00:00", "2019-05-11 17:00:00+00:00", "2019-05-11 16:00:00+00:00", "2019-05-11 15:00:00+00:00", "2019-05-11 14:00:00+00:00", "2019-05-11 13:00:00+00:00", "2019-05-11 12:00:00+00:00", "2019-05-11 11:00:00+00:00", "2019-05-11 10:00:00+00:00", "2019-05-11 09:00:00+00:00", "2019-05-11 08:00:00+00:00", "2019-05-11 07:00:00+00:00", "2019-05-11 06:00:00+00:00", "2019-05-11 02:00:00+00:00", "2019-05-11 01:00:00+00:00", "2019-05-11 00:00:00+00:00", "2019-05-10 23:00:00+00:00", "2019-05-10 22:00:00+00:00", "2019-05-10 21:00:00+00:00", "2019-05-10 20:00:00+00:00", "2019-05-10 19:00:00+00:00", "2019-05-10 18:00:00+00:00", "2019-05-10 17:00:00+00:00", "2019-05-10 16:00:00+00:00", "2019-05-10 15:00:00+00:00", "2019-05-10 14:00:00+00:00", "2019-05-10 13:00:00+00:00", "2019-05-10 12:00:00+00:00", "2019-05-10 11:00:00+00:00", "2019-05-10 10:00:00+00:00", "2019-05-10 09:00:00+00:00", "2019-05-10 08:00:00+00:00", "2019-05-10 07:00:00+00:00", "2019-05-10 06:00:00+00:00", "2019-05-10 05:00:00+00:00", "2019-05-10 04:00:00+00:00", "2019-05-10 03:00:00+00:00", "2019-05-10 02:00:00+00:00", "2019-05-10 01:00:00+00:00", "2019-05-10 00:00:00+00:00", "2019-05-09 23:00:00+00:00", "2019-05-09 22:00:00+00:00", "2019-05-09 21:00:00+00:00", "2019-05-09 20:00:00+00:00", "2019-05-09 19:00:00+00:00", "2019-05-09 18:00:00+00:00", "2019-05-09 17:00:00+00:00", "2019-05-09 16:00:00+00:00", "2019-05-09 15:00:00+00:00", "2019-05-09 14:00:00+00:00", "2019-05-09 13:00:00+00:00", "2019-05-09 12:00:00+00:00", "2019-05-09 11:00:00+00:00", "2019-05-09 10:00:00+00:00", "2019-05-09 09:00:00+00:00", "2019-05-09 08:00:00+00:00", "2019-05-09 07:00:00+00:00", "2019-05-09 06:00:00+00:00", "2019-05-09 05:00:00+00:00", "2019-05-09 04:00:00+00:00", "2019-05-09 03:00:00+00:00", "2019-05-09 02:00:00+00:00", "2019-05-09 01:00:00+00:00", "2019-05-09 00:00:00+00:00", "2019-05-08 23:00:00+00:00", "2019-05-08 22:00:00+00:00", "2019-05-08 21:00:00+00:00", "2019-05-08 20:00:00+00:00", "2019-05-08 19:00:00+00:00", "2019-05-08 18:00:00+00:00", "2019-05-08 17:00:00+00:00", "2019-05-08 16:00:00+00:00", "2019-05-08 15:00:00+00:00", "2019-05-08 14:00:00+00:00", "2019-05-08 13:00:00+00:00", "2019-05-08 12:00:00+00:00", "2019-05-08 11:00:00+00:00", "2019-05-08 10:00:00+00:00", "2019-05-08 09:00:00+00:00", "2019-05-08 08:00:00+00:00", "2019-05-08 07:00:00+00:00", "2019-05-08 06:00:00+00:00", "2019-05-08 05:00:00+00:00", "2019-05-08 04:00:00+00:00", "2019-05-08 03:00:00+00:00", "2019-05-08 02:00:00+00:00", "2019-05-08 01:00:00+00:00", "2019-05-08 00:00:00+00:00", "2019-05-07 23:00:00+00:00", "2019-05-07 22:00:00+00:00", "2019-05-07 21:00:00+00:00", "2019-05-07 20:00:00+00:00", "2019-05-07 19:00:00+00:00", "2019-05-07 18:00:00+00:00", "2019-05-07 17:00:00+00:00", "2019-05-07 16:00:00+00:00", "2019-05-07 15:00:00+00:00", "2019-05-07 14:00:00+00:00", "2019-05-07 13:00:00+00:00", "2019-05-07 12:00:00+00:00", "2019-05-07 11:00:00+00:00", "2019-05-07 10:00:00+00:00", "2019-05-07 09:00:00+00:00", "2019-05-07 08:00:00+00:00", "2019-05-07 07:00:00+00:00", "2019-05-07 06:00:00+00:00", "2019-05-07 05:00:00+00:00", "2019-05-07 04:00:00+00:00", "2019-05-07 03:00:00+00:00", "2019-05-07 02:00:00+00:00", "2019-05-07 01:00:00+00:00", "2019-05-07 00:00:00+00:00", "2019-05-06 23:00:00+00:00", "2019-05-06 22:00:00+00:00", "2019-05-06 21:00:00+00:00", "2019-05-06 20:00:00+00:00", "2019-05-06 19:00:00+00:00", "2019-05-06 18:00:00+00:00", "2019-05-06 17:00:00+00:00", "2019-05-06 16:00:00+00:00", "2019-05-06 15:00:00+00:00", "2019-05-06 14:00:00+00:00", "2019-05-06 13:00:00+00:00", "2019-05-06 12:00:00+00:00", "2019-05-06 11:00:00+00:00", "2019-05-06 10:00:00+00:00", "2019-05-06 09:00:00+00:00", "2019-05-06 08:00:00+00:00", "2019-05-06 07:00:00+00:00", "2019-05-06 06:00:00+00:00", "2019-05-06 05:00:00+00:00", "2019-05-06 04:00:00+00:00", "2019-05-06 03:00:00+00:00", "2019-05-06 02:00:00+00:00", "2019-05-06 01:00:00+00:00", "2019-05-06 00:00:00+00:00", "2019-05-05 23:00:00+00:00", "2019-05-05 22:00:00+00:00", "2019-05-05 21:00:00+00:00", "2019-05-05 20:00:00+00:00", "2019-05-05 19:00:00+00:00", "2019-05-05 18:00:00+00:00", "2019-05-05 17:00:00+00:00", "2019-05-05 16:00:00+00:00", "2019-05-05 15:00:00+00:00", "2019-05-05 14:00:00+00:00", "2019-05-05 13:00:00+00:00", "2019-05-05 12:00:00+00:00", "2019-05-05 11:00:00+00:00", "2019-05-05 10:00:00+00:00", "2019-05-05 09:00:00+00:00", "2019-05-05 08:00:00+00:00", "2019-05-05 07:00:00+00:00", "2019-05-05 06:00:00+00:00", "2019-05-05 05:00:00+00:00", "2019-05-05 04:00:00+00:00", "2019-05-05 03:00:00+00:00", "2019-05-05 02:00:00+00:00", "2019-05-05 01:00:00+00:00", "2019-05-05 00:00:00+00:00", "2019-05-04 23:00:00+00:00", "2019-05-04 22:00:00+00:00", "2019-05-04 21:00:00+00:00", "2019-05-04 20:00:00+00:00", "2019-05-04 19:00:00+00:00", "2019-05-04 18:00:00+00:00", "2019-05-04 17:00:00+00:00", "2019-05-04 16:00:00+00:00", "2019-05-04 15:00:00+00:00", "2019-05-04 14:00:00+00:00", "2019-05-04 13:00:00+00:00", "2019-05-04 12:00:00+00:00", "2019-05-04 11:00:00+00:00", "2019-05-04 10:00:00+00:00", "2019-05-04 09:00:00+00:00", "2019-05-04 08:00:00+00:00", "2019-05-04 07:00:00+00:00", "2019-05-04 06:00:00+00:00", "2019-05-04 05:00:00+00:00", "2019-05-04 04:00:00+00:00", "2019-05-04 03:00:00+00:00", "2019-05-04 02:00:00+00:00", "2019-05-04 01:00:00+00:00", "2019-05-04 00:00:00+00:00", "2019-05-03 23:00:00+00:00", "2019-05-03 22:00:00+00:00", "2019-05-03 21:00:00+00:00", "2019-05-03 20:00:00+00:00", "2019-05-03 19:00:00+00:00", "2019-05-03 18:00:00+00:00", "2019-05-03 17:00:00+00:00", "2019-05-03 16:00:00+00:00", "2019-05-03 15:00:00+00:00", "2019-05-03 14:00:00+00:00", "2019-05-03 13:00:00+00:00", "2019-05-03 12:00:00+00:00", "2019-05-03 11:00:00+00:00", "2019-05-03 10:00:00+00:00", "2019-05-03 09:00:00+00:00", "2019-05-03 08:00:00+00:00", "2019-05-03 07:00:00+00:00", "2019-05-03 06:00:00+00:00", "2019-05-03 05:00:00+00:00", "2019-05-03 04:00:00+00:00", "2019-05-03 03:00:00+00:00", "2019-05-03 02:00:00+00:00", "2019-05-03 01:00:00+00:00", "2019-05-03 00:00:00+00:00", "2019-05-02 23:00:00+00:00", "2019-05-02 22:00:00+00:00", "2019-05-02 21:00:00+00:00", "2019-05-02 20:00:00+00:00", "2019-05-02 19:00:00+00:00", "2019-05-02 18:00:00+00:00", "2019-05-02 17:00:00+00:00", "2019-05-02 16:00:00+00:00", "2019-05-02 15:00:00+00:00", "2019-05-02 14:00:00+00:00", "2019-05-02 13:00:00+00:00", "2019-05-02 12:00:00+00:00", "2019-05-02 11:00:00+00:00", "2019-05-02 10:00:00+00:00", "2019-05-02 09:00:00+00:00", "2019-05-02 08:00:00+00:00", "2019-05-02 07:00:00+00:00", "2019-05-02 06:00:00+00:00", "2019-05-02 05:00:00+00:00", "2019-05-02 04:00:00+00:00", "2019-05-02 03:00:00+00:00", "2019-05-02 02:00:00+00:00", "2019-05-02 01:00:00+00:00", "2019-05-02 00:00:00+00:00", "2019-05-01 23:00:00+00:00", "2019-05-01 22:00:00+00:00", "2019-05-01 21:00:00+00:00", "2019-05-01 20:00:00+00:00", "2019-05-01 19:00:00+00:00", "2019-05-01 18:00:00+00:00", "2019-05-01 17:00:00+00:00", "2019-05-01 16:00:00+00:00", "2019-05-01 15:00:00+00:00", "2019-05-01 14:00:00+00:00", "2019-05-01 13:00:00+00:00", "2019-05-01 12:00:00+00:00", "2019-05-01 11:00:00+00:00", "2019-05-01 10:00:00+00:00", "2019-05-01 09:00:00+00:00", "2019-05-01 08:00:00+00:00", "2019-05-01 07:00:00+00:00", "2019-05-01 06:00:00+00:00", "2019-05-01 05:00:00+00:00", "2019-05-01 04:00:00+00:00", "2019-05-01 03:00:00+00:00", "2019-05-01 02:00:00+00:00", "2019-05-01 01:00:00+00:00", "2019-05-01 00:00:00+00:00", "2019-04-30 23:00:00+00:00", "2019-04-30 22:00:00+00:00", "2019-04-30 21:00:00+00:00", "2019-04-30 20:00:00+00:00", "2019-04-30 19:00:00+00:00", "2019-04-30 18:00:00+00:00", "2019-04-30 17:00:00+00:00", "2019-04-30 16:00:00+00:00", "2019-04-30 15:00:00+00:00", "2019-04-30 14:00:00+00:00", "2019-04-30 13:00:00+00:00", "2019-04-30 12:00:00+00:00", "2019-04-30 11:00:00+00:00", "2019-04-30 10:00:00+00:00", "2019-04-30 09:00:00+00:00", "2019-04-30 08:00:00+00:00", "2019-04-30 07:00:00+00:00", "2019-04-30 06:00:00+00:00", "2019-04-30 05:00:00+00:00", "2019-04-30 04:00:00+00:00", "2019-04-30 03:00:00+00:00", "2019-04-30 02:00:00+00:00", "2019-04-30 01:00:00+00:00", "2019-04-30 00:00:00+00:00", "2019-04-29 23:00:00+00:00", "2019-04-29 22:00:00+00:00", "2019-04-29 21:00:00+00:00", "2019-04-29 20:00:00+00:00", "2019-04-29 19:00:00+00:00", "2019-04-29 18:00:00+00:00", "2019-04-29 17:00:00+00:00", "2019-04-29 16:00:00+00:00", "2019-04-29 15:00:00+00:00", "2019-04-29 14:00:00+00:00", "2019-04-29 13:00:00+00:00", "2019-04-29 12:00:00+00:00", "2019-04-29 11:00:00+00:00", "2019-04-29 10:00:00+00:00", "2019-04-29 09:00:00+00:00", "2019-04-29 08:00:00+00:00", "2019-04-29 07:00:00+00:00", "2019-04-29 06:00:00+00:00", "2019-04-29 05:00:00+00:00", "2019-04-29 04:00:00+00:00", "2019-04-29 03:00:00+00:00", "2019-04-29 02:00:00+00:00", "2019-04-29 01:00:00+00:00", "2019-04-29 00:00:00+00:00", "2019-04-28 23:00:00+00:00", "2019-04-28 22:00:00+00:00", "2019-04-28 21:00:00+00:00", "2019-04-28 20:00:00+00:00", "2019-04-28 19:00:00+00:00", "2019-04-28 18:00:00+00:00", "2019-04-28 17:00:00+00:00", "2019-04-28 16:00:00+00:00", "2019-04-28 15:00:00+00:00", "2019-04-28 14:00:00+00:00", "2019-04-28 13:00:00+00:00", "2019-04-28 12:00:00+00:00", "2019-04-28 11:00:00+00:00", "2019-04-28 10:00:00+00:00", "2019-04-28 09:00:00+00:00", "2019-04-28 08:00:00+00:00", "2019-04-28 07:00:00+00:00", "2019-04-28 06:00:00+00:00", "2019-04-28 05:00:00+00:00", "2019-04-28 04:00:00+00:00", "2019-04-28 03:00:00+00:00", "2019-04-28 02:00:00+00:00", "2019-04-28 01:00:00+00:00", "2019-04-28 00:00:00+00:00", "2019-04-27 23:00:00+00:00", "2019-04-27 22:00:00+00:00", "2019-04-27 21:00:00+00:00", "2019-04-27 20:00:00+00:00", "2019-04-27 19:00:00+00:00", "2019-04-27 18:00:00+00:00", "2019-04-27 17:00:00+00:00", "2019-04-27 16:00:00+00:00", "2019-04-27 15:00:00+00:00", "2019-04-27 14:00:00+00:00", "2019-04-27 13:00:00+00:00", "2019-04-27 12:00:00+00:00", "2019-04-27 11:00:00+00:00", "2019-04-27 10:00:00+00:00", "2019-04-27 09:00:00+00:00", "2019-04-27 08:00:00+00:00", "2019-04-27 07:00:00+00:00", "2019-04-27 06:00:00+00:00", "2019-04-27 05:00:00+00:00", "2019-04-27 04:00:00+00:00", "2019-04-27 03:00:00+00:00", "2019-04-27 02:00:00+00:00", "2019-04-27 01:00:00+00:00", "2019-04-27 00:00:00+00:00", "2019-04-26 23:00:00+00:00", "2019-04-26 22:00:00+00:00", "2019-04-26 21:00:00+00:00", "2019-04-26 20:00:00+00:00", "2019-04-26 19:00:00+00:00", "2019-04-26 18:00:00+00:00", "2019-04-26 17:00:00+00:00", "2019-04-26 16:00:00+00:00", "2019-04-26 15:00:00+00:00", "2019-04-26 14:00:00+00:00", "2019-04-26 13:00:00+00:00", "2019-04-26 12:00:00+00:00", "2019-04-26 11:00:00+00:00", "2019-04-26 10:00:00+00:00", "2019-04-26 09:00:00+00:00", "2019-04-26 08:00:00+00:00", "2019-04-26 07:00:00+00:00", "2019-04-26 06:00:00+00:00", "2019-04-26 05:00:00+00:00", "2019-04-26 04:00:00+00:00", "2019-04-26 03:00:00+00:00", "2019-04-26 02:00:00+00:00", "2019-04-26 01:00:00+00:00", "2019-04-26 00:00:00+00:00", "2019-04-25 23:00:00+00:00", "2019-04-25 22:00:00+00:00", "2019-04-25 21:00:00+00:00", "2019-04-25 20:00:00+00:00", "2019-04-25 19:00:00+00:00", "2019-04-25 18:00:00+00:00", "2019-04-25 17:00:00+00:00", "2019-04-25 16:00:00+00:00", "2019-04-25 15:00:00+00:00", "2019-04-25 14:00:00+00:00", "2019-04-25 13:00:00+00:00", "2019-04-25 12:00:00+00:00", "2019-04-25 11:00:00+00:00", "2019-04-25 10:00:00+00:00", "2019-04-25 09:00:00+00:00", "2019-04-25 08:00:00+00:00", "2019-04-25 07:00:00+00:00", "2019-04-25 06:00:00+00:00", "2019-04-25 05:00:00+00:00", "2019-04-25 04:00:00+00:00", "2019-04-25 03:00:00+00:00", "2019-04-25 02:00:00+00:00", "2019-04-25 01:00:00+00:00", "2019-04-25 00:00:00+00:00", "2019-04-24 23:00:00+00:00", "2019-04-24 22:00:00+00:00", "2019-04-24 21:00:00+00:00", "2019-04-24 20:00:00+00:00", "2019-04-24 19:00:00+00:00", "2019-04-24 18:00:00+00:00", "2019-04-24 17:00:00+00:00", "2019-04-24 16:00:00+00:00", "2019-04-24 15:00:00+00:00", "2019-04-24 14:00:00+00:00", "2019-04-24 13:00:00+00:00", "2019-04-24 12:00:00+00:00", "2019-04-24 11:00:00+00:00", "2019-04-24 10:00:00+00:00", "2019-04-24 09:00:00+00:00", "2019-04-24 08:00:00+00:00", "2019-04-24 07:00:00+00:00", "2019-04-24 06:00:00+00:00", "2019-04-24 05:00:00+00:00", "2019-04-24 04:00:00+00:00", "2019-04-24 03:00:00+00:00", "2019-04-24 02:00:00+00:00", "2019-04-24 01:00:00+00:00", "2019-04-24 00:00:00+00:00", "2019-04-23 23:00:00+00:00", "2019-04-23 22:00:00+00:00", "2019-04-23 21:00:00+00:00", "2019-04-23 20:00:00+00:00", "2019-04-23 19:00:00+00:00", "2019-04-23 18:00:00+00:00", "2019-04-23 17:00:00+00:00", "2019-04-23 16:00:00+00:00", "2019-04-23 15:00:00+00:00", "2019-04-23 14:00:00+00:00", "2019-04-23 13:00:00+00:00", "2019-04-23 12:00:00+00:00", "2019-04-23 11:00:00+00:00", "2019-04-23 10:00:00+00:00", "2019-04-23 09:00:00+00:00", "2019-04-23 08:00:00+00:00", "2019-04-23 07:00:00+00:00", "2019-04-23 06:00:00+00:00", "2019-04-23 05:00:00+00:00", "2019-04-23 04:00:00+00:00", "2019-04-23 03:00:00+00:00", "2019-04-23 02:00:00+00:00", "2019-04-23 01:00:00+00:00", "2019-04-23 00:00:00+00:00", "2019-04-22 23:00:00+00:00", "2019-04-22 22:00:00+00:00", "2019-04-22 21:00:00+00:00", "2019-04-22 20:00:00+00:00", "2019-04-22 19:00:00+00:00", "2019-04-22 18:00:00+00:00", "2019-04-22 17:00:00+00:00", "2019-04-22 16:00:00+00:00", "2019-04-22 15:00:00+00:00", "2019-04-22 14:00:00+00:00", "2019-04-22 13:00:00+00:00", "2019-04-22 12:00:00+00:00", "2019-04-22 11:00:00+00:00", "2019-04-22 10:00:00+00:00", "2019-04-22 09:00:00+00:00", "2019-04-22 08:00:00+00:00", "2019-04-22 07:00:00+00:00", "2019-04-22 06:00:00+00:00", "2019-04-22 05:00:00+00:00", "2019-04-22 04:00:00+00:00", "2019-04-22 03:00:00+00:00", "2019-04-22 02:00:00+00:00", "2019-04-22 01:00:00+00:00", "2019-04-22 00:00:00+00:00", "2019-04-21 23:00:00+00:00", "2019-04-21 22:00:00+00:00", "2019-04-21 21:00:00+00:00", "2019-04-21 20:00:00+00:00", "2019-04-21 19:00:00+00:00", "2019-04-21 18:00:00+00:00", "2019-04-21 17:00:00+00:00", "2019-04-21 16:00:00+00:00", "2019-04-21 15:00:00+00:00", "2019-04-21 14:00:00+00:00", "2019-04-21 13:00:00+00:00", "2019-04-21 12:00:00+00:00", "2019-04-21 11:00:00+00:00", "2019-04-21 10:00:00+00:00", "2019-04-21 09:00:00+00:00", "2019-04-21 08:00:00+00:00", "2019-04-21 07:00:00+00:00", "2019-04-21 06:00:00+00:00", "2019-04-21 05:00:00+00:00", "2019-04-21 04:00:00+00:00", "2019-04-21 03:00:00+00:00", "2019-04-21 02:00:00+00:00", "2019-04-21 01:00:00+00:00", "2019-04-21 00:00:00+00:00", "2019-04-20 23:00:00+00:00", "2019-04-20 22:00:00+00:00", "2019-04-20 21:00:00+00:00", "2019-04-20 20:00:00+00:00", "2019-04-20 19:00:00+00:00", "2019-04-20 18:00:00+00:00", "2019-04-20 17:00:00+00:00", "2019-04-20 16:00:00+00:00", "2019-04-20 15:00:00+00:00", "2019-04-20 14:00:00+00:00", "2019-04-20 13:00:00+00:00", "2019-04-20 12:00:00+00:00", "2019-04-20 11:00:00+00:00", "2019-04-20 10:00:00+00:00", "2019-04-20 09:00:00+00:00", "2019-04-20 08:00:00+00:00", "2019-04-20 07:00:00+00:00", "2019-04-20 06:00:00+00:00", "2019-04-20 05:00:00+00:00", "2019-04-20 04:00:00+00:00", "2019-04-20 03:00:00+00:00", "2019-04-20 02:00:00+00:00", "2019-04-20 01:00:00+00:00", "2019-04-20 00:00:00+00:00", "2019-04-19 23:00:00+00:00", "2019-04-19 22:00:00+00:00", "2019-04-19 21:00:00+00:00", "2019-04-19 20:00:00+00:00", "2019-04-19 19:00:00+00:00", "2019-04-19 18:00:00+00:00", "2019-04-19 17:00:00+00:00", "2019-04-19 16:00:00+00:00", "2019-04-19 15:00:00+00:00", "2019-04-19 14:00:00+00:00", "2019-04-19 13:00:00+00:00", "2019-04-19 12:00:00+00:00", "2019-04-19 11:00:00+00:00", "2019-04-19 10:00:00+00:00", "2019-04-19 09:00:00+00:00", "2019-04-19 08:00:00+00:00", "2019-04-19 07:00:00+00:00", "2019-04-19 06:00:00+00:00", "2019-04-19 05:00:00+00:00", "2019-04-19 04:00:00+00:00", "2019-04-19 03:00:00+00:00", "2019-04-19 02:00:00+00:00", "2019-04-19 01:00:00+00:00", "2019-04-19 00:00:00+00:00", "2019-04-18 23:00:00+00:00", "2019-04-18 22:00:00+00:00", "2019-04-18 21:00:00+00:00", "2019-04-18 20:00:00+00:00", "2019-04-18 19:00:00+00:00", "2019-04-18 18:00:00+00:00", "2019-04-18 17:00:00+00:00", "2019-04-18 16:00:00+00:00", "2019-04-18 15:00:00+00:00", "2019-04-18 14:00:00+00:00", "2019-04-18 13:00:00+00:00", "2019-04-18 12:00:00+00:00", "2019-04-18 11:00:00+00:00", "2019-04-18 10:00:00+00:00", "2019-04-18 09:00:00+00:00", "2019-04-18 08:00:00+00:00", "2019-04-18 07:00:00+00:00", "2019-04-18 06:00:00+00:00", "2019-04-18 05:00:00+00:00", "2019-04-18 04:00:00+00:00", "2019-04-18 03:00:00+00:00", "2019-04-18 02:00:00+00:00", "2019-04-18 01:00:00+00:00", "2019-04-18 00:00:00+00:00", "2019-04-17 23:00:00+00:00", "2019-04-17 22:00:00+00:00", "2019-04-17 21:00:00+00:00", "2019-04-17 20:00:00+00:00", "2019-04-17 19:00:00+00:00", "2019-04-17 18:00:00+00:00", "2019-04-17 17:00:00+00:00", "2019-04-17 16:00:00+00:00", "2019-04-17 15:00:00+00:00", "2019-04-17 14:00:00+00:00", "2019-04-17 13:00:00+00:00", "2019-04-17 12:00:00+00:00", "2019-04-17 11:00:00+00:00", "2019-04-17 10:00:00+00:00", "2019-04-17 09:00:00+00:00", "2019-04-17 08:00:00+00:00", "2019-04-17 07:00:00+00:00", "2019-04-17 06:00:00+00:00", "2019-04-17 05:00:00+00:00", "2019-04-17 04:00:00+00:00", "2019-04-17 03:00:00+00:00", "2019-04-17 02:00:00+00:00", "2019-04-17 01:00:00+00:00", "2019-04-17 00:00:00+00:00", "2019-04-16 23:00:00+00:00", "2019-04-16 22:00:00+00:00", "2019-04-16 21:00:00+00:00", "2019-04-16 20:00:00+00:00", "2019-04-16 19:00:00+00:00", "2019-04-16 18:00:00+00:00", "2019-04-16 17:00:00+00:00", "2019-04-16 16:00:00+00:00", "2019-04-16 15:00:00+00:00", "2019-04-16 14:00:00+00:00", "2019-04-16 13:00:00+00:00", "2019-04-16 12:00:00+00:00", "2019-04-16 11:00:00+00:00", "2019-04-16 10:00:00+00:00", "2019-04-16 09:00:00+00:00", "2019-04-16 08:00:00+00:00", "2019-04-16 07:00:00+00:00", "2019-04-16 06:00:00+00:00", "2019-04-16 05:00:00+00:00", "2019-04-16 04:00:00+00:00", "2019-04-16 03:00:00+00:00", "2019-04-16 02:00:00+00:00", "2019-04-16 01:00:00+00:00", "2019-04-16 00:00:00+00:00", "2019-04-15 23:00:00+00:00", "2019-04-15 22:00:00+00:00", "2019-04-15 21:00:00+00:00", "2019-04-15 20:00:00+00:00", "2019-04-15 19:00:00+00:00", "2019-04-15 18:00:00+00:00", "2019-04-15 17:00:00+00:00", "2019-04-15 16:00:00+00:00", "2019-04-15 15:00:00+00:00", "2019-04-15 14:00:00+00:00", "2019-04-15 13:00:00+00:00", "2019-04-15 12:00:00+00:00", "2019-04-15 11:00:00+00:00", "2019-04-15 10:00:00+00:00", "2019-04-15 09:00:00+00:00", "2019-04-15 08:00:00+00:00", "2019-04-15 07:00:00+00:00", "2019-04-15 06:00:00+00:00", "2019-04-15 05:00:00+00:00", "2019-04-15 04:00:00+00:00", "2019-04-15 03:00:00+00:00", "2019-04-15 02:00:00+00:00", "2019-04-15 01:00:00+00:00", "2019-04-15 00:00:00+00:00", "2019-04-14 23:00:00+00:00", "2019-04-14 22:00:00+00:00", "2019-04-14 21:00:00+00:00", "2019-04-14 20:00:00+00:00", "2019-04-14 19:00:00+00:00", "2019-04-14 18:00:00+00:00", "2019-04-14 17:00:00+00:00", "2019-04-14 16:00:00+00:00", "2019-04-14 15:00:00+00:00", "2019-04-14 14:00:00+00:00", "2019-04-14 13:00:00+00:00", "2019-04-14 12:00:00+00:00", "2019-04-14 11:00:00+00:00", "2019-04-14 10:00:00+00:00", "2019-04-14 09:00:00+00:00", "2019-04-14 08:00:00+00:00", "2019-04-14 07:00:00+00:00", "2019-04-14 06:00:00+00:00", "2019-04-14 05:00:00+00:00", "2019-04-14 04:00:00+00:00", "2019-04-14 03:00:00+00:00", "2019-04-14 02:00:00+00:00", "2019-04-14 01:00:00+00:00", "2019-04-14 00:00:00+00:00", "2019-04-13 23:00:00+00:00", "2019-04-13 22:00:00+00:00", "2019-04-13 21:00:00+00:00", "2019-04-13 20:00:00+00:00", "2019-04-13 19:00:00+00:00", "2019-04-13 18:00:00+00:00", "2019-04-13 17:00:00+00:00", "2019-04-13 16:00:00+00:00", "2019-04-13 15:00:00+00:00", "2019-04-13 14:00:00+00:00", "2019-04-13 13:00:00+00:00", "2019-04-13 12:00:00+00:00", "2019-04-13 11:00:00+00:00", "2019-04-13 10:00:00+00:00", "2019-04-13 09:00:00+00:00", "2019-04-13 08:00:00+00:00", "2019-04-13 07:00:00+00:00", "2019-04-13 06:00:00+00:00", "2019-04-13 05:00:00+00:00", "2019-04-13 04:00:00+00:00", "2019-04-13 03:00:00+00:00", "2019-04-13 02:00:00+00:00", "2019-04-13 01:00:00+00:00", "2019-04-13 00:00:00+00:00", "2019-04-12 23:00:00+00:00", "2019-04-12 22:00:00+00:00", "2019-04-12 21:00:00+00:00", "2019-04-12 20:00:00+00:00", "2019-04-12 19:00:00+00:00", "2019-04-12 18:00:00+00:00", "2019-04-12 17:00:00+00:00", "2019-04-12 16:00:00+00:00", "2019-04-12 15:00:00+00:00", "2019-04-12 14:00:00+00:00", "2019-04-12 13:00:00+00:00", "2019-04-12 12:00:00+00:00", "2019-04-12 11:00:00+00:00", "2019-04-12 10:00:00+00:00", "2019-04-12 09:00:00+00:00", "2019-04-12 08:00:00+00:00", "2019-04-12 07:00:00+00:00", "2019-04-12 06:00:00+00:00", "2019-04-12 05:00:00+00:00", "2019-04-12 04:00:00+00:00", "2019-04-12 03:00:00+00:00", "2019-04-12 02:00:00+00:00", "2019-04-12 01:00:00+00:00", "2019-04-12 00:00:00+00:00", "2019-04-11 23:00:00+00:00", "2019-04-11 22:00:00+00:00", "2019-04-11 21:00:00+00:00", "2019-04-11 20:00:00+00:00", "2019-04-11 19:00:00+00:00", "2019-04-11 18:00:00+00:00", "2019-04-11 17:00:00+00:00", "2019-04-11 16:00:00+00:00", "2019-04-11 15:00:00+00:00", "2019-04-11 14:00:00+00:00", "2019-04-11 13:00:00+00:00", "2019-04-11 12:00:00+00:00", "2019-04-11 11:00:00+00:00", "2019-04-11 10:00:00+00:00", "2019-04-11 09:00:00+00:00", "2019-04-11 08:00:00+00:00", "2019-04-11 07:00:00+00:00", "2019-04-11 06:00:00+00:00", "2019-04-11 05:00:00+00:00", "2019-04-11 04:00:00+00:00", "2019-04-11 03:00:00+00:00", "2019-04-11 02:00:00+00:00", "2019-04-11 01:00:00+00:00", "2019-04-11 00:00:00+00:00", "2019-04-10 23:00:00+00:00", "2019-04-10 22:00:00+00:00", "2019-04-10 21:00:00+00:00", "2019-04-10 20:00:00+00:00", "2019-04-10 19:00:00+00:00", "2019-04-10 18:00:00+00:00", "2019-04-10 17:00:00+00:00", "2019-04-10 16:00:00+00:00", "2019-04-10 15:00:00+00:00", "2019-04-10 14:00:00+00:00", "2019-04-10 13:00:00+00:00", "2019-04-10 12:00:00+00:00", "2019-04-10 11:00:00+00:00", "2019-04-10 10:00:00+00:00", "2019-04-10 09:00:00+00:00", "2019-04-10 08:00:00+00:00", "2019-04-10 07:00:00+00:00", "2019-04-10 06:00:00+00:00", "2019-04-10 05:00:00+00:00", "2019-04-10 04:00:00+00:00", "2019-04-10 03:00:00+00:00", "2019-04-10 02:00:00+00:00", "2019-04-10 01:00:00+00:00", "2019-04-10 00:00:00+00:00", "2019-04-09 23:00:00+00:00", "2019-04-09 22:00:00+00:00", "2019-04-09 21:00:00+00:00", "2019-04-09 20:00:00+00:00", "2019-04-09 19:00:00+00:00", "2019-04-09 18:00:00+00:00", "2019-04-09 17:00:00+00:00", "2019-04-09 16:00:00+00:00", "2019-04-09 15:00:00+00:00", "2019-04-09 14:00:00+00:00", "2019-04-09 13:00:00+00:00", "2019-04-09 12:00:00+00:00", "2019-04-09 11:00:00+00:00", "2019-04-09 10:00:00+00:00", "2019-04-09 09:00:00+00:00", "2019-04-09 08:00:00+00:00", "2019-04-09 07:00:00+00:00", "2019-04-09 06:00:00+00:00", "2019-04-09 05:00:00+00:00", "2019-04-09 04:00:00+00:00", "2019-04-09 03:00:00+00:00", "2019-04-09 02:00:00+00:00", "2019-04-09 01:00:00+00:00",
// 					},
// 					"5a19d02982add20b2d3c0bc86795c532f97f8ad4642548abe3e854fdfbd68612c4b9d89b608c3423a3b10c39290f6522d59858517d29b4272959ff4a96759770": {
// 						"2019-06-17 11:00:00+00:00", "2019-06-17 10:00:00+00:00", "2019-06-17 09:00:00+00:00", "2019-06-17 08:00:00+00:00", "2019-06-17 07:00:00+00:00", "2019-06-17 06:00:00+00:00", "2019-06-17 05:00:00+00:00", "2019-06-17 04:00:00+00:00", "2019-06-17 03:00:00+00:00", "2019-06-17 02:00:00+00:00", "2019-06-17 01:00:00+00:00", "2019-06-17 00:00:00+00:00", "2019-06-16 23:00:00+00:00", "2019-06-16 21:00:00+00:00", "2019-06-16 20:00:00+00:00", "2019-06-16 19:00:00+00:00", "2019-06-16 18:00:00+00:00", "2019-06-16 17:00:00+00:00", "2019-06-16 16:00:00+00:00", "2019-06-16 15:00:00+00:00", "2019-06-16 14:00:00+00:00", "2019-06-16 13:00:00+00:00", "2019-06-16 12:00:00+00:00", "2019-06-16 11:00:00+00:00", "2019-06-16 10:00:00+00:00", "2019-06-16 09:00:00+00:00", "2019-06-16 08:00:00+00:00", "2019-06-16 07:00:00+00:00", "2019-06-16 06:00:00+00:00", "2019-06-16 05:00:00+00:00", "2019-06-16 04:00:00+00:00", "2019-06-16 03:00:00+00:00", "2019-06-16 02:00:00+00:00", "2019-06-16 01:00:00+00:00", "2019-06-16 00:00:00+00:00", "2019-06-15 23:00:00+00:00", "2019-06-15 22:00:00+00:00", "2019-06-15 21:00:00+00:00", "2019-06-15 20:00:00+00:00", "2019-06-15 19:00:00+00:00", "2019-06-15 18:00:00+00:00", "2019-06-15 17:00:00+00:00", "2019-06-15 16:00:00+00:00", "2019-06-15 15:00:00+00:00", "2019-06-15 14:00:00+00:00", "2019-06-15 13:00:00+00:00", "2019-06-15 12:00:00+00:00", "2019-06-15 11:00:00+00:00", "2019-06-15 10:00:00+00:00", "2019-06-15 09:00:00+00:00", "2019-06-15 08:00:00+00:00", "2019-06-15 07:00:00+00:00", "2019-06-15 06:00:00+00:00", "2019-06-15 05:00:00+00:00", "2019-06-15 04:00:00+00:00", "2019-06-15 00:00:00+00:00", "2019-06-14 23:00:00+00:00", "2019-06-14 22:00:00+00:00", "2019-06-14 21:00:00+00:00", "2019-06-14 20:00:00+00:00", "2019-06-14 19:00:00+00:00", "2019-06-14 18:00:00+00:00", "2019-06-14 17:00:00+00:00", "2019-06-14 16:00:00+00:00", "2019-06-14 15:00:00+00:00", "2019-06-14 14:00:00+00:00", "2019-06-14 13:00:00+00:00", "2019-06-14 12:00:00+00:00", "2019-06-14 11:00:00+00:00", "2019-06-14 10:00:00+00:00", "2019-06-14 09:00:00+00:00", "2019-06-14 08:00:00+00:00", "2019-06-14 07:00:00+00:00", "2019-06-14 06:00:00+00:00", "2019-06-14 05:00:00+00:00", "2019-06-14 04:00:00+00:00", "2019-06-14 03:00:00+00:00", "2019-06-14 02:00:00+00:00", "2019-06-14 00:00:00+00:00", "2019-06-13 23:00:00+00:00", "2019-06-13 22:00:00+00:00", "2019-06-13 21:00:00+00:00", "2019-06-13 20:00:00+00:00", "2019-06-13 19:00:00+00:00", "2019-06-13 18:00:00+00:00", "2019-06-13 17:00:00+00:00", "2019-06-13 16:00:00+00:00", "2019-06-13 15:00:00+00:00", "2019-06-13 14:00:00+00:00", "2019-06-13 13:00:00+00:00", "2019-06-13 12:00:00+00:00", "2019-06-13 11:00:00+00:00", "2019-06-13 10:00:00+00:00", "2019-06-13 09:00:00+00:00", "2019-06-13 08:00:00+00:00", "2019-06-13 07:00:00+00:00", "2019-06-13 06:00:00+00:00", "2019-06-13 05:00:00+00:00", "2019-06-13 04:00:00+00:00", "2019-06-13 03:00:00+00:00", "2019-06-13 02:00:00+00:00", "2019-06-13 00:00:00+00:00", "2019-06-12 23:00:00+00:00", "2019-06-12 21:00:00+00:00", "2019-06-12 20:00:00+00:00", "2019-06-12 19:00:00+00:00", "2019-06-12 18:00:00+00:00", "2019-06-12 17:00:00+00:00", "2019-06-12 16:00:00+00:00", "2019-06-12 15:00:00+00:00", "2019-06-12 14:00:00+00:00", "2019-06-12 13:00:00+00:00", "2019-06-12 12:00:00+00:00", "2019-06-12 11:00:00+00:00", "2019-06-12 10:00:00+00:00", "2019-06-12 09:00:00+00:00", "2019-06-12 08:00:00+00:00", "2019-06-12 07:00:00+00:00", "2019-06-12 06:00:00+00:00", "2019-06-12 05:00:00+00:00", "2019-06-12 04:00:00+00:00", "2019-06-12 03:00:00+00:00", "2019-06-12 00:00:00+00:00", "2019-06-11 23:00:00+00:00", "2019-06-11 22:00:00+00:00", "2019-06-11 21:00:00+00:00", "2019-06-11 20:00:00+00:00", "2019-06-11 19:00:00+00:00", "2019-06-11 18:00:00+00:00", "2019-06-11 17:00:00+00:00", "2019-06-11 16:00:00+00:00", "2019-06-11 15:00:00+00:00", "2019-06-11 14:00:00+00:00", "2019-06-11 13:00:00+00:00", "2019-06-11 12:00:00+00:00", "2019-06-11 11:00:00+00:00", "2019-06-11 10:00:00+00:00", "2019-06-11 09:00:00+00:00", "2019-06-11 08:00:00+00:00", "2019-06-11 07:00:00+00:00", "2019-06-11 06:00:00+00:00", "2019-06-11 05:00:00+00:00", "2019-06-11 04:00:00+00:00", "2019-06-11 03:00:00+00:00", "2019-06-11 02:00:00+00:00", "2019-06-11 01:00:00+00:00", "2019-06-11 00:00:00+00:00", "2019-06-10 23:00:00+00:00", "2019-06-10 22:00:00+00:00", "2019-06-10 21:00:00+00:00", "2019-06-10 20:00:00+00:00", "2019-06-10 19:00:00+00:00", "2019-06-10 18:00:00+00:00", "2019-06-10 17:00:00+00:00", "2019-06-10 16:00:00+00:00", "2019-06-10 15:00:00+00:00", "2019-06-10 14:00:00+00:00", "2019-06-10 13:00:00+00:00", "2019-06-10 12:00:00+00:00", "2019-06-10 11:00:00+00:00", "2019-06-10 10:00:00+00:00", "2019-06-10 09:00:00+00:00", "2019-06-10 08:00:00+00:00", "2019-06-10 07:00:00+00:00", "2019-06-10 06:00:00+00:00", "2019-06-10 05:00:00+00:00", "2019-06-10 04:00:00+00:00", "2019-06-10 03:00:00+00:00", "2019-06-10 02:00:00+00:00", "2019-06-10 01:00:00+00:00", "2019-06-10 00:00:00+00:00", "2019-06-09 23:00:00+00:00", "2019-06-09 21:00:00+00:00", "2019-06-09 20:00:00+00:00", "2019-06-09 19:00:00+00:00", "2019-06-09 18:00:00+00:00", "2019-06-09 17:00:00+00:00", "2019-06-09 16:00:00+00:00", "2019-06-09 15:00:00+00:00", "2019-06-09 14:00:00+00:00", "2019-06-09 13:00:00+00:00", "2019-06-09 12:00:00+00:00", "2019-06-09 11:00:00+00:00", "2019-06-09 10:00:00+00:00", "2019-06-09 09:00:00+00:00", "2019-06-09 08:00:00+00:00", "2019-06-09 07:00:00+00:00", "2019-06-09 06:00:00+00:00", "2019-06-09 05:00:00+00:00", "2019-06-09 04:00:00+00:00", "2019-06-09 03:00:00+00:00", "2019-06-09 02:00:00+00:00", "2019-06-09 01:00:00+00:00", "2019-06-09 00:00:00+00:00", "2019-06-08 23:00:00+00:00", "2019-06-08 21:00:00+00:00", "2019-06-08 20:00:00+00:00", "2019-06-08 19:00:00+00:00", "2019-06-08 18:00:00+00:00", "2019-06-08 17:00:00+00:00", "2019-06-08 16:00:00+00:00", "2019-06-08 15:00:00+00:00", "2019-06-08 14:00:00+00:00", "2019-06-08 13:00:00+00:00", "2019-06-08 12:00:00+00:00", "2019-06-08 11:00:00+00:00", "2019-06-08 10:00:00+00:00", "2019-06-08 09:00:00+00:00", "2019-06-08 08:00:00+00:00", "2019-06-08 07:00:00+00:00", "2019-06-08 06:00:00+00:00", "2019-06-08 05:00:00+00:00", "2019-06-08 04:00:00+00:00", "2019-06-08 03:00:00+00:00", "2019-06-08 02:00:00+00:00", "2019-06-08 00:00:00+00:00", "2019-06-07 23:00:00+00:00", "2019-06-07 21:00:00+00:00", "2019-06-07 20:00:00+00:00", "2019-06-07 19:00:00+00:00", "2019-06-07 18:00:00+00:00", "2019-06-07 17:00:00+00:00", "2019-06-07 16:00:00+00:00", "2019-06-07 15:00:00+00:00", "2019-06-07 14:00:00+00:00", "2019-06-07 13:00:00+00:00", "2019-06-07 12:00:00+00:00", "2019-06-07 11:00:00+00:00", "2019-06-07 10:00:00+00:00", "2019-06-07 09:00:00+00:00", "2019-06-07 08:00:00+00:00", "2019-06-07 07:00:00+00:00", "2019-06-07 06:00:00+00:00", "2019-06-07 05:00:00+00:00", "2019-06-07 04:00:00+00:00", "2019-06-07 03:00:00+00:00", "2019-06-07 02:00:00+00:00", "2019-06-07 01:00:00+00:00", "2019-06-07 00:00:00+00:00", "2019-06-06 23:00:00+00:00", "2019-06-06 22:00:00+00:00", "2019-06-06 21:00:00+00:00", "2019-06-06 20:00:00+00:00", "2019-06-06 19:00:00+00:00", "2019-06-06 18:00:00+00:00", "2019-06-06 17:00:00+00:00", "2019-06-06 16:00:00+00:00", "2019-06-06 15:00:00+00:00", "2019-06-06 14:00:00+00:00", "2019-06-06 13:00:00+00:00", "2019-06-06 12:00:00+00:00", "2019-06-06 11:00:00+00:00", "2019-06-06 10:00:00+00:00", "2019-06-06 09:00:00+00:00", "2019-06-06 08:00:00+00:00", "2019-06-06 07:00:00+00:00", "2019-06-06 06:00:00+00:00", "2019-06-06 05:00:00+00:00", "2019-06-06 04:00:00+00:00", "2019-06-06 03:00:00+00:00", "2019-06-06 02:00:00+00:00", "2019-06-06 00:00:00+00:00", "2019-06-05 23:00:00+00:00", "2019-06-05 22:00:00+00:00", "2019-06-05 21:00:00+00:00", "2019-06-05 20:00:00+00:00", "2019-06-05 19:00:00+00:00", "2019-06-05 18:00:00+00:00", "2019-06-05 17:00:00+00:00", "2019-06-05 16:00:00+00:00", "2019-06-05 15:00:00+00:00", "2019-06-05 14:00:00+00:00", "2019-06-05 13:00:00+00:00", "2019-06-05 12:00:00+00:00", "2019-06-05 11:00:00+00:00", "2019-06-05 10:00:00+00:00", "2019-06-05 09:00:00+00:00", "2019-06-05 08:00:00+00:00", "2019-06-05 07:00:00+00:00", "2019-06-05 06:00:00+00:00", "2019-06-05 05:00:00+00:00", "2019-06-05 04:00:00+00:00", "2019-06-05 03:00:00+00:00", "2019-06-05 02:00:00+00:00", "2019-06-05 01:00:00+00:00", "2019-06-05 00:00:00+00:00", "2019-06-04 23:00:00+00:00", "2019-06-04 22:00:00+00:00", "2019-06-04 21:00:00+00:00", "2019-06-04 20:00:00+00:00", "2019-06-04 19:00:00+00:00", "2019-06-04 18:00:00+00:00", "2019-06-04 17:00:00+00:00", "2019-06-04 16:00:00+00:00", "2019-06-04 15:00:00+00:00", "2019-06-04 14:00:00+00:00", "2019-06-04 13:00:00+00:00", "2019-06-04 12:00:00+00:00", "2019-06-04 11:00:00+00:00", "2019-06-04 10:00:00+00:00", "2019-06-04 09:00:00+00:00", "2019-06-04 08:00:00+00:00", "2019-06-04 07:00:00+00:00", "2019-06-04 06:00:00+00:00", "2019-06-04 05:00:00+00:00", "2019-06-04 04:00:00+00:00", "2019-06-04 03:00:00+00:00", "2019-06-04 02:00:00+00:00", "2019-06-04 01:00:00+00:00", "2019-06-04 00:00:00+00:00", "2019-06-03 23:00:00+00:00", "2019-06-03 22:00:00+00:00", "2019-06-03 21:00:00+00:00", "2019-06-03 20:00:00+00:00", "2019-06-03 19:00:00+00:00", "2019-06-03 18:00:00+00:00", "2019-06-03 17:00:00+00:00", "2019-06-03 16:00:00+00:00", "2019-06-03 15:00:00+00:00", "2019-06-03 14:00:00+00:00", "2019-06-03 13:00:00+00:00", "2019-06-03 12:00:00+00:00", "2019-06-03 11:00:00+00:00", "2019-06-03 10:00:00+00:00", "2019-06-03 09:00:00+00:00", "2019-06-03 08:00:00+00:00", "2019-06-03 07:00:00+00:00", "2019-06-03 06:00:00+00:00", "2019-06-03 05:00:00+00:00", "2019-06-03 04:00:00+00:00", "2019-06-03 03:00:00+00:00", "2019-06-03 02:00:00+00:00", "2019-06-03 01:00:00+00:00", "2019-06-03 00:00:00+00:00", "2019-06-02 23:00:00+00:00", "2019-06-02 22:00:00+00:00", "2019-06-02 21:00:00+00:00", "2019-06-02 20:00:00+00:00", "2019-06-02 19:00:00+00:00", "2019-06-02 18:00:00+00:00", "2019-06-02 17:00:00+00:00", "2019-06-02 16:00:00+00:00", "2019-06-02 15:00:00+00:00", "2019-06-02 14:00:00+00:00", "2019-06-02 13:00:00+00:00", "2019-06-02 12:00:00+00:00", "2019-06-02 11:00:00+00:00", "2019-06-02 10:00:00+00:00", "2019-06-02 09:00:00+00:00", "2019-06-02 08:00:00+00:00", "2019-06-02 07:00:00+00:00", "2019-06-02 06:00:00+00:00", "2019-06-02 05:00:00+00:00", "2019-06-02 04:00:00+00:00", "2019-06-02 03:00:00+00:00", "2019-06-02 02:00:00+00:00", "2019-06-02 01:00:00+00:00", "2019-06-02 00:00:00+00:00", "2019-06-01 23:00:00+00:00", "2019-06-01 22:00:00+00:00", "2019-06-01 21:00:00+00:00", "2019-06-01 20:00:00+00:00", "2019-06-01 19:00:00+00:00", "2019-06-01 18:00:00+00:00", "2019-06-01 17:00:00+00:00", "2019-06-01 16:00:00+00:00", "2019-06-01 15:00:00+00:00", "2019-06-01 14:00:00+00:00", "2019-06-01 13:00:00+00:00", "2019-06-01 12:00:00+00:00", "2019-06-01 11:00:00+00:00", "2019-06-01 10:00:00+00:00", "2019-06-01 09:00:00+00:00", "2019-06-01 08:00:00+00:00", "2019-06-01 07:00:00+00:00", "2019-06-01 06:00:00+00:00", "2019-06-01 05:00:00+00:00", "2019-06-01 04:00:00+00:00", "2019-06-01 03:00:00+00:00", "2019-06-01 02:00:00+00:00", "2019-06-01 01:00:00+00:00", "2019-06-01 00:00:00+00:00", "2019-05-31 23:00:00+00:00", "2019-05-31 22:00:00+00:00", "2019-05-31 21:00:00+00:00", "2019-05-31 20:00:00+00:00", "2019-05-31 19:00:00+00:00", "2019-05-31 18:00:00+00:00", "2019-05-31 17:00:00+00:00", "2019-05-31 16:00:00+00:00", "2019-05-31 15:00:00+00:00", "2019-05-31 14:00:00+00:00", "2019-05-31 13:00:00+00:00", "2019-05-31 12:00:00+00:00", "2019-05-31 11:00:00+00:00", "2019-05-31 10:00:00+00:00", "2019-05-31 09:00:00+00:00", "2019-05-31 08:00:00+00:00", "2019-05-31 07:00:00+00:00", "2019-05-31 06:00:00+00:00", "2019-05-31 05:00:00+00:00", "2019-05-31 04:00:00+00:00", "2019-05-31 03:00:00+00:00", "2019-05-31 02:00:00+00:00", "2019-05-31 01:00:00+00:00", "2019-05-31 00:00:00+00:00", "2019-05-30 23:00:00+00:00", "2019-05-30 22:00:00+00:00", "2019-05-30 21:00:00+00:00", "2019-05-30 20:00:00+00:00", "2019-05-30 19:00:00+00:00", "2019-05-30 18:00:00+00:00", "2019-05-30 17:00:00+00:00", "2019-05-30 16:00:00+00:00", "2019-05-30 15:00:00+00:00", "2019-05-30 14:00:00+00:00", "2019-05-30 13:00:00+00:00", "2019-05-30 12:00:00+00:00", "2019-05-30 11:00:00+00:00", "2019-05-30 10:00:00+00:00", "2019-05-30 09:00:00+00:00", "2019-05-30 08:00:00+00:00", "2019-05-30 07:00:00+00:00", "2019-05-30 06:00:00+00:00", "2019-05-30 05:00:00+00:00", "2019-05-30 04:00:00+00:00", "2019-05-30 03:00:00+00:00", "2019-05-30 02:00:00+00:00", "2019-05-30 01:00:00+00:00", "2019-05-30 00:00:00+00:00", "2019-05-29 23:00:00+00:00", "2019-05-29 22:00:00+00:00", "2019-05-29 21:00:00+00:00", "2019-05-29 20:00:00+00:00", "2019-05-29 19:00:00+00:00", "2019-05-29 18:00:00+00:00", "2019-05-29 17:00:00+00:00", "2019-05-29 16:00:00+00:00", "2019-05-29 15:00:00+00:00", "2019-05-29 14:00:00+00:00", "2019-05-29 13:00:00+00:00", "2019-05-29 12:00:00+00:00", "2019-05-29 11:00:00+00:00", "2019-05-29 10:00:00+00:00", "2019-05-29 09:00:00+00:00", "2019-05-29 08:00:00+00:00", "2019-05-29 07:00:00+00:00", "2019-05-29 06:00:00+00:00", "2019-05-29 05:00:00+00:00", "2019-05-29 04:00:00+00:00", "2019-05-29 03:00:00+00:00", "2019-05-29 02:00:00+00:00", "2019-05-29 01:00:00+00:00", "2019-05-29 00:00:00+00:00", "2019-05-28 23:00:00+00:00", "2019-05-28 21:00:00+00:00", "2019-05-28 20:00:00+00:00", "2019-05-28 19:00:00+00:00", "2019-05-28 18:00:00+00:00", "2019-05-28 17:00:00+00:00", "2019-05-28 16:00:00+00:00", "2019-05-28 15:00:00+00:00", "2019-05-28 14:00:00+00:00", "2019-05-28 13:00:00+00:00", "2019-05-28 12:00:00+00:00", "2019-05-28 11:00:00+00:00", "2019-05-28 10:00:00+00:00", "2019-05-28 09:00:00+00:00", "2019-05-28 08:00:00+00:00", "2019-05-28 07:00:00+00:00", "2019-05-28 06:00:00+00:00", "2019-05-28 05:00:00+00:00", "2019-05-28 04:00:00+00:00", "2019-05-28 03:00:00+00:00", "2019-05-28 02:00:00+00:00", "2019-05-28 01:00:00+00:00", "2019-05-28 00:00:00+00:00", "2019-05-27 23:00:00+00:00", "2019-05-27 22:00:00+00:00", "2019-05-27 21:00:00+00:00", "2019-05-27 20:00:00+00:00", "2019-05-27 19:00:00+00:00", "2019-05-27 18:00:00+00:00", "2019-05-27 17:00:00+00:00", "2019-05-27 16:00:00+00:00", "2019-05-27 15:00:00+00:00", "2019-05-27 14:00:00+00:00", "2019-05-27 13:00:00+00:00", "2019-05-27 12:00:00+00:00", "2019-05-27 11:00:00+00:00", "2019-05-27 10:00:00+00:00", "2019-05-27 09:00:00+00:00", "2019-05-27 08:00:00+00:00", "2019-05-27 07:00:00+00:00", "2019-05-27 06:00:00+00:00", "2019-05-27 05:00:00+00:00", "2019-05-27 04:00:00+00:00", "2019-05-27 03:00:00+00:00", "2019-05-27 02:00:00+00:00", "2019-05-27 01:00:00+00:00", "2019-05-27 00:00:00+00:00", "2019-05-26 23:00:00+00:00", "2019-05-26 22:00:00+00:00", "2019-05-26 21:00:00+00:00", "2019-05-26 20:00:00+00:00", "2019-05-26 19:00:00+00:00", "2019-05-26 18:00:00+00:00", "2019-05-26 17:00:00+00:00", "2019-05-26 16:00:00+00:00", "2019-05-26 15:00:00+00:00", "2019-05-26 14:00:00+00:00", "2019-05-26 13:00:00+00:00", "2019-05-26 12:00:00+00:00", "2019-05-26 11:00:00+00:00", "2019-05-26 10:00:00+00:00", "2019-05-26 09:00:00+00:00", "2019-05-26 08:00:00+00:00", "2019-05-26 07:00:00+00:00", "2019-05-26 06:00:00+00:00", "2019-05-26 05:00:00+00:00", "2019-05-26 04:00:00+00:00", "2019-05-26 03:00:00+00:00", "2019-05-26 02:00:00+00:00", "2019-05-26 01:00:00+00:00", "2019-05-26 00:00:00+00:00", "2019-05-25 23:00:00+00:00", "2019-05-25 22:00:00+00:00", "2019-05-25 21:00:00+00:00", "2019-05-25 20:00:00+00:00", "2019-05-25 19:00:00+00:00", "2019-05-25 18:00:00+00:00", "2019-05-25 17:00:00+00:00", "2019-05-25 16:00:00+00:00", "2019-05-25 15:00:00+00:00", "2019-05-25 14:00:00+00:00", "2019-05-25 13:00:00+00:00", "2019-05-25 12:00:00+00:00", "2019-05-25 11:00:00+00:00", "2019-05-25 10:00:00+00:00", "2019-05-25 09:00:00+00:00", "2019-05-25 08:00:00+00:00", "2019-05-25 07:00:00+00:00", "2019-05-25 06:00:00+00:00", "2019-05-25 05:00:00+00:00", "2019-05-25 04:00:00+00:00", "2019-05-25 03:00:00+00:00", "2019-05-25 02:00:00+00:00", "2019-05-25 01:00:00+00:00", "2019-05-25 00:00:00+00:00", "2019-05-24 23:00:00+00:00", "2019-05-24 22:00:00+00:00", "2019-05-24 21:00:00+00:00", "2019-05-24 20:00:00+00:00", "2019-05-24 19:00:00+00:00", "2019-05-24 18:00:00+00:00", "2019-05-24 17:00:00+00:00", "2019-05-24 16:00:00+00:00", "2019-05-24 15:00:00+00:00", "2019-05-24 14:00:00+00:00", "2019-05-24 13:00:00+00:00", "2019-05-24 12:00:00+00:00", "2019-05-24 11:00:00+00:00", "2019-05-24 10:00:00+00:00", "2019-05-24 09:00:00+00:00", "2019-05-24 08:00:00+00:00", "2019-05-24 07:00:00+00:00", "2019-05-24 06:00:00+00:00", "2019-05-24 05:00:00+00:00", "2019-05-24 04:00:00+00:00", "2019-05-24 03:00:00+00:00", "2019-05-24 02:00:00+00:00", "2019-05-24 00:00:00+00:00", "2019-05-23 23:00:00+00:00", "2019-05-23 22:00:00+00:00", "2019-05-23 21:00:00+00:00", "2019-05-23 20:00:00+00:00", "2019-05-23 19:00:00+00:00", "2019-05-23 18:00:00+00:00", "2019-05-23 17:00:00+00:00", "2019-05-23 16:00:00+00:00", "2019-05-23 15:00:00+00:00", "2019-05-23 14:00:00+00:00", "2019-05-23 13:00:00+00:00", "2019-05-23 12:00:00+00:00", "2019-05-23 11:00:00+00:00", "2019-05-23 10:00:00+00:00", "2019-05-23 09:00:00+00:00", "2019-05-23 08:00:00+00:00", "2019-05-23 07:00:00+00:00", "2019-05-23 06:00:00+00:00", "2019-05-23 05:00:00+00:00", "2019-05-23 04:00:00+00:00", "2019-05-23 03:00:00+00:00", "2019-05-23 02:00:00+00:00", "2019-05-23 01:00:00+00:00", "2019-05-23 00:00:00+00:00", "2019-05-22 23:00:00+00:00", "2019-05-22 22:00:00+00:00", "2019-05-22 21:00:00+00:00", "2019-05-22 20:00:00+00:00", "2019-05-22 19:00:00+00:00", "2019-05-22 18:00:00+00:00", "2019-05-22 17:00:00+00:00", "2019-05-22 16:00:00+00:00", "2019-05-22 15:00:00+00:00", "2019-05-22 14:00:00+00:00", "2019-05-22 13:00:00+00:00", "2019-05-22 12:00:00+00:00", "2019-05-22 11:00:00+00:00", "2019-05-22 10:00:00+00:00", "2019-05-22 09:00:00+00:00", "2019-05-22 08:00:00+00:00", "2019-05-22 07:00:00+00:00", "2019-05-22 06:00:00+00:00", "2019-05-22 05:00:00+00:00", "2019-05-22 04:00:00+00:00", "2019-05-22 03:00:00+00:00", "2019-05-22 02:00:00+00:00", "2019-05-22 01:00:00+00:00", "2019-05-22 00:00:00+00:00", "2019-05-21 23:00:00+00:00", "2019-05-21 22:00:00+00:00", "2019-05-21 21:00:00+00:00", "2019-05-21 20:00:00+00:00", "2019-05-21 19:00:00+00:00", "2019-05-21 18:00:00+00:00", "2019-05-21 17:00:00+00:00", "2019-05-21 16:00:00+00:00", "2019-05-21 15:00:00+00:00", "2019-05-21 14:00:00+00:00", "2019-05-21 13:00:00+00:00", "2019-05-21 12:00:00+00:00", "2019-05-21 11:00:00+00:00", "2019-05-21 10:00:00+00:00", "2019-05-21 09:00:00+00:00", "2019-05-21 08:00:00+00:00", "2019-05-21 07:00:00+00:00", "2019-05-21 06:00:00+00:00", "2019-05-21 05:00:00+00:00", "2019-05-21 04:00:00+00:00", "2019-05-21 03:00:00+00:00", "2019-05-21 02:00:00+00:00", "2019-05-21 01:00:00+00:00", "2019-05-21 00:00:00+00:00", "2019-05-20 23:00:00+00:00", "2019-05-20 22:00:00+00:00", "2019-05-20 21:00:00+00:00", "2019-05-20 20:00:00+00:00", "2019-05-20 19:00:00+00:00", "2019-05-20 18:00:00+00:00", "2019-05-20 17:00:00+00:00", "2019-05-20 16:00:00+00:00", "2019-05-20 15:00:00+00:00", "2019-05-20 14:00:00+00:00", "2019-05-20 13:00:00+00:00", "2019-05-20 12:00:00+00:00", "2019-05-20 11:00:00+00:00", "2019-05-20 10:00:00+00:00", "2019-05-20 09:00:00+00:00", "2019-05-20 08:00:00+00:00", "2019-05-20 07:00:00+00:00", "2019-05-20 06:00:00+00:00", "2019-05-20 05:00:00+00:00", "2019-05-20 04:00:00+00:00", "2019-05-20 03:00:00+00:00", "2019-05-20 02:00:00+00:00", "2019-05-20 01:00:00+00:00", "2019-05-20 00:00:00+00:00", "2019-05-19 23:00:00+00:00", "2019-05-19 22:00:00+00:00", "2019-05-19 21:00:00+00:00", "2019-05-19 20:00:00+00:00", "2019-05-19 19:00:00+00:00", "2019-05-19 18:00:00+00:00", "2019-05-19 17:00:00+00:00", "2019-05-19 16:00:00+00:00", "2019-05-19 15:00:00+00:00", "2019-05-19 14:00:00+00:00", "2019-05-19 13:00:00+00:00", "2019-05-19 12:00:00+00:00", "2019-05-19 11:00:00+00:00", "2019-05-19 10:00:00+00:00", "2019-05-19 09:00:00+00:00", "2019-05-19 08:00:00+00:00", "2019-05-19 07:00:00+00:00", "2019-05-19 06:00:00+00:00", "2019-05-19 05:00:00+00:00", "2019-05-19 04:00:00+00:00", "2019-05-19 03:00:00+00:00", "2019-05-19 02:00:00+00:00", "2019-05-19 01:00:00+00:00", "2019-05-19 00:00:00+00:00", "2019-05-18 23:00:00+00:00", "2019-05-18 22:00:00+00:00", "2019-05-18 21:00:00+00:00", "2019-05-18 20:00:00+00:00", "2019-05-18 19:00:00+00:00", "2019-05-18 18:00:00+00:00", "2019-05-18 17:00:00+00:00", "2019-05-18 16:00:00+00:00", "2019-05-18 15:00:00+00:00", "2019-05-18 14:00:00+00:00", "2019-05-18 13:00:00+00:00", "2019-05-18 12:00:00+00:00", "2019-05-18 11:00:00+00:00", "2019-05-18 10:00:00+00:00", "2019-05-18 09:00:00+00:00", "2019-05-18 08:00:00+00:00", "2019-05-18 07:00:00+00:00", "2019-05-18 06:00:00+00:00", "2019-05-18 05:00:00+00:00", "2019-05-18 04:00:00+00:00", "2019-05-18 03:00:00+00:00", "2019-05-18 02:00:00+00:00", "2019-05-18 01:00:00+00:00", "2019-05-18 00:00:00+00:00", "2019-05-17 23:00:00+00:00", "2019-05-17 22:00:00+00:00", "2019-05-17 21:00:00+00:00", "2019-05-17 20:00:00+00:00", "2019-05-17 19:00:00+00:00", "2019-05-17 18:00:00+00:00", "2019-05-17 17:00:00+00:00", "2019-05-17 16:00:00+00:00", "2019-05-17 15:00:00+00:00", "2019-05-17 14:00:00+00:00", "2019-05-17 13:00:00+00:00", "2019-05-17 12:00:00+00:00", "2019-05-17 11:00:00+00:00", "2019-05-17 10:00:00+00:00", "2019-05-17 09:00:00+00:00", "2019-05-17 08:00:00+00:00", "2019-05-17 07:00:00+00:00", "2019-05-17 06:00:00+00:00", "2019-05-17 05:00:00+00:00", "2019-05-17 04:00:00+00:00", "2019-05-17 03:00:00+00:00", "2019-05-17 02:00:00+00:00", "2019-05-17 01:00:00+00:00", "2019-05-17 00:00:00+00:00", "2019-05-16 23:00:00+00:00", "2019-05-16 22:00:00+00:00", "2019-05-16 21:00:00+00:00", "2019-05-16 20:00:00+00:00", "2019-05-16 19:00:00+00:00", "2019-05-16 18:00:00+00:00", "2019-05-16 17:00:00+00:00", "2019-05-16 16:00:00+00:00", "2019-05-16 15:00:00+00:00", "2019-05-16 14:00:00+00:00", "2019-05-16 13:00:00+00:00", "2019-05-16 12:00:00+00:00", "2019-05-16 11:00:00+00:00", "2019-05-16 10:00:00+00:00", "2019-05-16 09:00:00+00:00", "2019-05-16 08:00:00+00:00", "2019-05-16 07:00:00+00:00", "2019-05-16 06:00:00+00:00", "2019-05-16 05:00:00+00:00", "2019-05-16 04:00:00+00:00", "2019-05-16 03:00:00+00:00", "2019-05-16 02:00:00+00:00", "2019-05-16 01:00:00+00:00", "2019-05-16 00:00:00+00:00", "2019-05-15 23:00:00+00:00", "2019-05-15 22:00:00+00:00", "2019-05-15 21:00:00+00:00", "2019-05-15 20:00:00+00:00", "2019-05-15 19:00:00+00:00", "2019-05-15 18:00:00+00:00", "2019-05-15 17:00:00+00:00", "2019-05-15 16:00:00+00:00", "2019-05-15 15:00:00+00:00", "2019-05-15 14:00:00+00:00", "2019-05-15 13:00:00+00:00", "2019-05-15 12:00:00+00:00", "2019-05-15 11:00:00+00:00", "2019-05-15 10:00:00+00:00", "2019-05-15 09:00:00+00:00", "2019-05-15 08:00:00+00:00", "2019-05-15 07:00:00+00:00", "2019-05-15 06:00:00+00:00", "2019-05-15 05:00:00+00:00", "2019-05-15 04:00:00+00:00", "2019-05-15 03:00:00+00:00", "2019-05-15 02:00:00+00:00", "2019-05-15 00:00:00+00:00", "2019-05-14 23:00:00+00:00", "2019-05-14 22:00:00+00:00", "2019-05-14 21:00:00+00:00", "2019-05-14 20:00:00+00:00", "2019-05-14 19:00:00+00:00", "2019-05-14 18:00:00+00:00", "2019-05-14 17:00:00+00:00", "2019-05-14 16:00:00+00:00", "2019-05-14 15:00:00+00:00", "2019-05-14 14:00:00+00:00", "2019-05-14 13:00:00+00:00", "2019-05-14 12:00:00+00:00", "2019-05-14 11:00:00+00:00", "2019-05-14 10:00:00+00:00", "2019-05-14 09:00:00+00:00", "2019-05-14 08:00:00+00:00", "2019-05-14 07:00:00+00:00", "2019-05-14 06:00:00+00:00", "2019-05-14 05:00:00+00:00", "2019-05-14 04:00:00+00:00", "2019-05-14 03:00:00+00:00", "2019-05-14 02:00:00+00:00", "2019-05-14 01:00:00+00:00", "2019-05-14 00:00:00+00:00", "2019-05-13 23:00:00+00:00", "2019-05-13 22:00:00+00:00", "2019-05-13 21:00:00+00:00", "2019-05-13 20:00:00+00:00", "2019-05-13 19:00:00+00:00", "2019-05-13 18:00:00+00:00", "2019-05-13 17:00:00+00:00", "2019-05-13 16:00:00+00:00", "2019-05-13 15:00:00+00:00", "2019-05-13 14:00:00+00:00", "2019-05-13 13:00:00+00:00", "2019-05-13 12:00:00+00:00", "2019-05-13 11:00:00+00:00", "2019-05-13 10:00:00+00:00", "2019-05-13 09:00:00+00:00", "2019-05-13 08:00:00+00:00", "2019-05-13 07:00:00+00:00", "2019-05-13 06:00:00+00:00", "2019-05-13 05:00:00+00:00", "2019-05-13 04:00:00+00:00", "2019-05-13 03:00:00+00:00", "2019-05-13 02:00:00+00:00", "2019-05-13 01:00:00+00:00", "2019-05-13 00:00:00+00:00", "2019-05-12 23:00:00+00:00", "2019-05-12 22:00:00+00:00", "2019-05-12 21:00:00+00:00", "2019-05-12 20:00:00+00:00", "2019-05-12 19:00:00+00:00", "2019-05-12 18:00:00+00:00", "2019-05-12 17:00:00+00:00", "2019-05-12 16:00:00+00:00", "2019-05-12 15:00:00+00:00", "2019-05-12 14:00:00+00:00", "2019-05-12 13:00:00+00:00", "2019-05-12 12:00:00+00:00", "2019-05-12 11:00:00+00:00", "2019-05-12 10:00:00+00:00", "2019-05-12 09:00:00+00:00", "2019-05-12 08:00:00+00:00", "2019-05-12 07:00:00+00:00", "2019-05-12 06:00:00+00:00", "2019-05-12 05:00:00+00:00", "2019-05-12 04:00:00+00:00", "2019-05-12 03:00:00+00:00", "2019-05-12 02:00:00+00:00", "2019-05-12 01:00:00+00:00", "2019-05-12 00:00:00+00:00", "2019-05-11 23:00:00+00:00", "2019-05-11 22:00:00+00:00", "2019-05-11 21:00:00+00:00", "2019-05-11 20:00:00+00:00", "2019-05-11 19:00:00+00:00", "2019-05-11 18:00:00+00:00", "2019-05-11 17:00:00+00:00", "2019-05-11 16:00:00+00:00", "2019-05-11 15:00:00+00:00", "2019-05-11 09:00:00+00:00", "2019-05-11 08:00:00+00:00", "2019-05-11 07:00:00+00:00", "2019-05-11 06:00:00+00:00", "2019-05-11 05:00:00+00:00", "2019-05-11 04:00:00+00:00", "2019-05-11 03:00:00+00:00", "2019-05-11 02:00:00+00:00", "2019-05-11 01:00:00+00:00", "2019-05-11 00:00:00+00:00", "2019-05-10 23:00:00+00:00", "2019-05-10 22:00:00+00:00", "2019-05-10 21:00:00+00:00", "2019-05-10 20:00:00+00:00", "2019-05-10 19:00:00+00:00", "2019-05-10 18:00:00+00:00", "2019-05-10 17:00:00+00:00", "2019-05-10 16:00:00+00:00", "2019-05-10 15:00:00+00:00", "2019-05-10 14:00:00+00:00", "2019-05-10 13:00:00+00:00", "2019-05-10 12:00:00+00:00", "2019-05-10 11:00:00+00:00", "2019-05-10 10:00:00+00:00", "2019-05-10 09:00:00+00:00", "2019-05-10 08:00:00+00:00", "2019-05-10 07:00:00+00:00", "2019-05-10 06:00:00+00:00", "2019-05-10 05:00:00+00:00", "2019-05-10 04:00:00+00:00", "2019-05-10 03:00:00+00:00", "2019-05-10 02:00:00+00:00", "2019-05-10 01:00:00+00:00", "2019-05-10 00:00:00+00:00", "2019-05-09 23:00:00+00:00", "2019-05-09 22:00:00+00:00", "2019-05-09 21:00:00+00:00", "2019-05-09 20:00:00+00:00", "2019-05-09 19:00:00+00:00", "2019-05-09 18:00:00+00:00", "2019-05-09 17:00:00+00:00", "2019-05-09 16:00:00+00:00", "2019-05-09 15:00:00+00:00", "2019-05-09 14:00:00+00:00", "2019-05-09 13:00:00+00:00", "2019-05-09 12:00:00+00:00", "2019-05-09 11:00:00+00:00", "2019-05-09 10:00:00+00:00", "2019-05-09 09:00:00+00:00", "2019-05-09 08:00:00+00:00", "2019-05-09 07:00:00+00:00", "2019-05-09 06:00:00+00:00", "2019-05-09 05:00:00+00:00", "2019-05-09 04:00:00+00:00", "2019-05-09 03:00:00+00:00", "2019-05-09 02:00:00+00:00", "2019-05-09 00:00:00+00:00", "2019-05-08 23:00:00+00:00", "2019-05-08 21:00:00+00:00", "2019-05-08 20:00:00+00:00", "2019-05-08 19:00:00+00:00", "2019-05-08 18:00:00+00:00", "2019-05-08 17:00:00+00:00", "2019-05-08 16:00:00+00:00", "2019-05-08 15:00:00+00:00", "2019-05-08 14:00:00+00:00", "2019-05-08 13:00:00+00:00", "2019-05-08 12:00:00+00:00", "2019-05-08 11:00:00+00:00", "2019-05-08 10:00:00+00:00", "2019-05-08 09:00:00+00:00", "2019-05-08 08:00:00+00:00", "2019-05-08 07:00:00+00:00", "2019-05-08 06:00:00+00:00", "2019-05-08 05:00:00+00:00", "2019-05-08 04:00:00+00:00", "2019-05-08 03:00:00+00:00", "2019-05-08 02:00:00+00:00", "2019-05-08 01:00:00+00:00", "2019-05-08 00:00:00+00:00", "2019-05-07 23:00:00+00:00", "2019-05-07 21:00:00+00:00", "2019-05-07 20:00:00+00:00", "2019-05-07 19:00:00+00:00", "2019-05-07 18:00:00+00:00", "2019-05-07 17:00:00+00:00", "2019-05-07 16:00:00+00:00", "2019-05-07 15:00:00+00:00", "2019-05-07 14:00:00+00:00", "2019-05-07 13:00:00+00:00", "2019-05-07 12:00:00+00:00", "2019-05-07 11:00:00+00:00", "2019-05-07 10:00:00+00:00", "2019-05-07 09:00:00+00:00", "2019-05-07 08:00:00+00:00", "2019-05-07 07:00:00+00:00", "2019-05-07 06:00:00+00:00", "2019-05-07 04:00:00+00:00", "2019-05-07 03:00:00+00:00", "2019-05-07 02:00:00+00:00", "2019-05-07 01:00:00+00:00", "2019-05-06 23:00:00+00:00", "2019-05-06 22:00:00+00:00", "2019-05-06 21:00:00+00:00", "2019-05-06 20:00:00+00:00", "2019-05-06 19:00:00+00:00", "2019-05-06 18:00:00+00:00", "2019-05-06 17:00:00+00:00", "2019-05-06 16:00:00+00:00", "2019-05-06 15:00:00+00:00", "2019-05-06 14:00:00+00:00", "2019-05-06 13:00:00+00:00", "2019-05-06 12:00:00+00:00", "2019-05-06 11:00:00+00:00", "2019-05-06 10:00:00+00:00", "2019-05-06 09:00:00+00:00", "2019-05-06 08:00:00+00:00", "2019-05-06 07:00:00+00:00", "2019-05-06 06:00:00+00:00", "2019-05-06 05:00:00+00:00", "2019-05-06 04:00:00+00:00", "2019-05-06 03:00:00+00:00", "2019-05-06 02:00:00+00:00", "2019-05-06 01:00:00+00:00", "2019-05-06 00:00:00+00:00", "2019-05-05 23:00:00+00:00", "2019-05-05 22:00:00+00:00", "2019-05-05 21:00:00+00:00", "2019-05-05 20:00:00+00:00", "2019-05-05 19:00:00+00:00", "2019-05-05 18:00:00+00:00", "2019-05-05 17:00:00+00:00", "2019-05-05 16:00:00+00:00", "2019-05-05 15:00:00+00:00", "2019-05-05 14:00:00+00:00", "2019-05-05 13:00:00+00:00", "2019-05-05 12:00:00+00:00", "2019-05-05 11:00:00+00:00", "2019-05-05 10:00:00+00:00", "2019-05-05 09:00:00+00:00", "2019-05-05 08:00:00+00:00", "2019-05-05 07:00:00+00:00", "2019-05-05 06:00:00+00:00", "2019-05-05 05:00:00+00:00", "2019-05-05 04:00:00+00:00", "2019-05-05 03:00:00+00:00", "2019-05-05 02:00:00+00:00", "2019-05-05 01:00:00+00:00", "2019-05-05 00:00:00+00:00", "2019-05-04 23:00:00+00:00", "2019-05-04 22:00:00+00:00", "2019-05-04 21:00:00+00:00", "2019-05-04 20:00:00+00:00", "2019-05-04 19:00:00+00:00", "2019-05-04 18:00:00+00:00", "2019-05-04 17:00:00+00:00", "2019-05-04 16:00:00+00:00", "2019-05-04 15:00:00+00:00", "2019-05-04 14:00:00+00:00", "2019-05-04 13:00:00+00:00", "2019-05-04 12:00:00+00:00", "2019-05-04 11:00:00+00:00", "2019-05-04 10:00:00+00:00", "2019-05-04 09:00:00+00:00", "2019-05-04 08:00:00+00:00", "2019-05-04 07:00:00+00:00", "2019-05-04 06:00:00+00:00", "2019-05-04 05:00:00+00:00", "2019-05-04 04:00:00+00:00", "2019-05-04 03:00:00+00:00", "2019-05-04 02:00:00+00:00", "2019-05-04 01:00:00+00:00", "2019-05-04 00:00:00+00:00", "2019-05-03 23:00:00+00:00", "2019-05-03 22:00:00+00:00", "2019-05-03 21:00:00+00:00", "2019-05-03 20:00:00+00:00", "2019-05-03 19:00:00+00:00", "2019-05-03 18:00:00+00:00", "2019-05-03 17:00:00+00:00", "2019-05-03 16:00:00+00:00", "2019-05-03 15:00:00+00:00", "2019-05-03 14:00:00+00:00", "2019-05-03 13:00:00+00:00", "2019-05-03 12:00:00+00:00", "2019-05-03 11:00:00+00:00", "2019-05-03 10:00:00+00:00", "2019-05-03 09:00:00+00:00", "2019-05-03 08:00:00+00:00", "2019-05-03 07:00:00+00:00", "2019-05-03 06:00:00+00:00", "2019-05-03 05:00:00+00:00", "2019-05-03 04:00:00+00:00", "2019-05-03 03:00:00+00:00", "2019-05-03 02:00:00+00:00", "2019-05-03 01:00:00+00:00", "2019-05-03 00:00:00+00:00", "2019-05-02 23:00:00+00:00", "2019-05-02 22:00:00+00:00", "2019-05-02 21:00:00+00:00", "2019-05-02 20:00:00+00:00", "2019-05-02 19:00:00+00:00", "2019-05-02 18:00:00+00:00", "2019-05-02 17:00:00+00:00", "2019-05-02 16:00:00+00:00", "2019-05-02 15:00:00+00:00", "2019-05-02 14:00:00+00:00", "2019-05-02 13:00:00+00:00", "2019-05-02 12:00:00+00:00", "2019-05-02 11:00:00+00:00", "2019-05-02 10:00:00+00:00", "2019-05-02 09:00:00+00:00", "2019-05-02 08:00:00+00:00", "2019-05-02 07:00:00+00:00", "2019-05-02 06:00:00+00:00", "2019-05-02 05:00:00+00:00", "2019-05-02 04:00:00+00:00", "2019-05-02 03:00:00+00:00", "2019-05-02 02:00:00+00:00", "2019-05-02 01:00:00+00:00", "2019-05-02 00:00:00+00:00", "2019-05-01 23:00:00+00:00", "2019-05-01 22:00:00+00:00", "2019-05-01 21:00:00+00:00", "2019-05-01 20:00:00+00:00", "2019-05-01 19:00:00+00:00", "2019-05-01 18:00:00+00:00", "2019-05-01 17:00:00+00:00", "2019-05-01 16:00:00+00:00", "2019-05-01 15:00:00+00:00", "2019-05-01 14:00:00+00:00", "2019-05-01 13:00:00+00:00", "2019-05-01 12:00:00+00:00", "2019-05-01 11:00:00+00:00", "2019-05-01 10:00:00+00:00", "2019-05-01 09:00:00+00:00", "2019-05-01 08:00:00+00:00", "2019-05-01 07:00:00+00:00", "2019-05-01 06:00:00+00:00", "2019-05-01 05:00:00+00:00", "2019-05-01 04:00:00+00:00", "2019-05-01 03:00:00+00:00", "2019-05-01 00:00:00+00:00", "2019-04-30 23:00:00+00:00", "2019-04-30 22:00:00+00:00", "2019-04-30 21:00:00+00:00", "2019-04-30 20:00:00+00:00", "2019-04-30 19:00:00+00:00", "2019-04-30 18:00:00+00:00", "2019-04-30 17:00:00+00:00", "2019-04-30 16:00:00+00:00", "2019-04-30 15:00:00+00:00", "2019-04-30 14:00:00+00:00", "2019-04-30 13:00:00+00:00", "2019-04-30 12:00:00+00:00", "2019-04-30 11:00:00+00:00", "2019-04-30 10:00:00+00:00", "2019-04-30 09:00:00+00:00", "2019-04-30 08:00:00+00:00", "2019-04-30 07:00:00+00:00", "2019-04-30 06:00:00+00:00", "2019-04-30 05:00:00+00:00", "2019-04-30 04:00:00+00:00", "2019-04-30 03:00:00+00:00", "2019-04-30 02:00:00+00:00", "2019-04-30 01:00:00+00:00", "2019-04-30 00:00:00+00:00", "2019-04-29 23:00:00+00:00", "2019-04-29 22:00:00+00:00", "2019-04-29 21:00:00+00:00", "2019-04-29 20:00:00+00:00", "2019-04-29 19:00:00+00:00", "2019-04-29 18:00:00+00:00", "2019-04-29 17:00:00+00:00", "2019-04-29 16:00:00+00:00", "2019-04-29 15:00:00+00:00", "2019-04-29 14:00:00+00:00", "2019-04-29 13:00:00+00:00", "2019-04-29 12:00:00+00:00", "2019-04-29 11:00:00+00:00", "2019-04-29 10:00:00+00:00", "2019-04-29 09:00:00+00:00", "2019-04-29 08:00:00+00:00", "2019-04-29 07:00:00+00:00", "2019-04-29 06:00:00+00:00", "2019-04-29 05:00:00+00:00", "2019-04-29 04:00:00+00:00", "2019-04-29 03:00:00+00:00", "2019-04-29 02:00:00+00:00", "2019-04-29 01:00:00+00:00", "2019-04-29 00:00:00+00:00", "2019-04-28 23:00:00+00:00", "2019-04-28 22:00:00+00:00", "2019-04-28 21:00:00+00:00", "2019-04-28 20:00:00+00:00", "2019-04-28 19:00:00+00:00", "2019-04-28 18:00:00+00:00", "2019-04-28 17:00:00+00:00", "2019-04-28 16:00:00+00:00", "2019-04-28 15:00:00+00:00", "2019-04-28 14:00:00+00:00", "2019-04-28 13:00:00+00:00", "2019-04-28 12:00:00+00:00", "2019-04-28 11:00:00+00:00", "2019-04-28 10:00:00+00:00", "2019-04-28 09:00:00+00:00", "2019-04-27 13:00:00+00:00", "2019-04-27 12:00:00+00:00", "2019-04-27 11:00:00+00:00", "2019-04-27 10:00:00+00:00", "2019-04-27 09:00:00+00:00", "2019-04-27 08:00:00+00:00", "2019-04-27 07:00:00+00:00", "2019-04-27 06:00:00+00:00", "2019-04-27 05:00:00+00:00", "2019-04-27 04:00:00+00:00", "2019-04-27 03:00:00+00:00", "2019-04-27 02:00:00+00:00", "2019-04-27 00:00:00+00:00", "2019-04-26 23:00:00+00:00", "2019-04-26 22:00:00+00:00", "2019-04-26 21:00:00+00:00", "2019-04-26 20:00:00+00:00", "2019-04-26 19:00:00+00:00", "2019-04-26 18:00:00+00:00", "2019-04-26 17:00:00+00:00", "2019-04-26 16:00:00+00:00", "2019-04-26 15:00:00+00:00", "2019-04-26 14:00:00+00:00", "2019-04-26 13:00:00+00:00", "2019-04-26 12:00:00+00:00", "2019-04-26 11:00:00+00:00", "2019-04-26 10:00:00+00:00", "2019-04-26 09:00:00+00:00", "2019-04-26 08:00:00+00:00", "2019-04-26 07:00:00+00:00", "2019-04-26 06:00:00+00:00", "2019-04-26 05:00:00+00:00", "2019-04-26 04:00:00+00:00", "2019-04-26 03:00:00+00:00", "2019-04-26 02:00:00+00:00", "2019-04-26 01:00:00+00:00", "2019-04-26 00:00:00+00:00", "2019-04-25 23:00:00+00:00", "2019-04-25 22:00:00+00:00", "2019-04-25 21:00:00+00:00", "2019-04-25 20:00:00+00:00", "2019-04-25 19:00:00+00:00", "2019-04-25 18:00:00+00:00", "2019-04-25 17:00:00+00:00", "2019-04-25 16:00:00+00:00", "2019-04-25 15:00:00+00:00", "2019-04-25 14:00:00+00:00", "2019-04-25 13:00:00+00:00", "2019-04-25 12:00:00+00:00", "2019-04-25 11:00:00+00:00", "2019-04-25 10:00:00+00:00", "2019-04-25 09:00:00+00:00", "2019-04-25 08:00:00+00:00", "2019-04-25 07:00:00+00:00", "2019-04-25 06:00:00+00:00", "2019-04-25 05:00:00+00:00", "2019-04-25 04:00:00+00:00", "2019-04-25 03:00:00+00:00", "2019-04-25 02:00:00+00:00", "2019-04-25 00:00:00+00:00", "2019-04-24 23:00:00+00:00", "2019-04-24 22:00:00+00:00", "2019-04-24 21:00:00+00:00", "2019-04-24 20:00:00+00:00", "2019-04-24 19:00:00+00:00", "2019-04-24 18:00:00+00:00", "2019-04-24 17:00:00+00:00", "2019-04-24 16:00:00+00:00", "2019-04-24 15:00:00+00:00", "2019-04-24 14:00:00+00:00", "2019-04-24 13:00:00+00:00", "2019-04-24 12:00:00+00:00", "2019-04-24 11:00:00+00:00", "2019-04-24 10:00:00+00:00", "2019-04-24 09:00:00+00:00", "2019-04-24 08:00:00+00:00", "2019-04-24 07:00:00+00:00", "2019-04-24 06:00:00+00:00", "2019-04-24 05:00:00+00:00", "2019-04-24 04:00:00+00:00", "2019-04-24 03:00:00+00:00", "2019-04-24 02:00:00+00:00", "2019-04-24 00:00:00+00:00", "2019-04-23 23:00:00+00:00", "2019-04-23 22:00:00+00:00", "2019-04-23 21:00:00+00:00", "2019-04-23 20:00:00+00:00", "2019-04-23 19:00:00+00:00", "2019-04-23 18:00:00+00:00", "2019-04-23 17:00:00+00:00", "2019-04-23 16:00:00+00:00", "2019-04-23 15:00:00+00:00", "2019-04-23 14:00:00+00:00", "2019-04-23 13:00:00+00:00", "2019-04-23 12:00:00+00:00", "2019-04-23 11:00:00+00:00", "2019-04-23 10:00:00+00:00", "2019-04-23 09:00:00+00:00", "2019-04-23 08:00:00+00:00", "2019-04-23 07:00:00+00:00", "2019-04-23 06:00:00+00:00", "2019-04-23 05:00:00+00:00", "2019-04-23 04:00:00+00:00", "2019-04-23 03:00:00+00:00", "2019-04-23 02:00:00+00:00", "2019-04-23 01:00:00+00:00", "2019-04-23 00:00:00+00:00", "2019-04-22 23:00:00+00:00", "2019-04-22 22:00:00+00:00", "2019-04-22 21:00:00+00:00", "2019-04-22 20:00:00+00:00", "2019-04-22 19:00:00+00:00", "2019-04-22 18:00:00+00:00", "2019-04-22 17:00:00+00:00", "2019-04-22 16:00:00+00:00", "2019-04-22 15:00:00+00:00", "2019-04-22 14:00:00+00:00", "2019-04-22 13:00:00+00:00", "2019-04-22 12:00:00+00:00", "2019-04-22 11:00:00+00:00", "2019-04-22 10:00:00+00:00", "2019-04-22 09:00:00+00:00", "2019-04-22 08:00:00+00:00", "2019-04-22 07:00:00+00:00", "2019-04-22 06:00:00+00:00", "2019-04-22 05:00:00+00:00", "2019-04-22 04:00:00+00:00", "2019-04-22 03:00:00+00:00", "2019-04-22 02:00:00+00:00", "2019-04-22 01:00:00+00:00", "2019-04-22 00:00:00+00:00", "2019-04-21 23:00:00+00:00", "2019-04-21 22:00:00+00:00", "2019-04-21 21:00:00+00:00", "2019-04-21 20:00:00+00:00", "2019-04-21 19:00:00+00:00", "2019-04-21 18:00:00+00:00", "2019-04-21 17:00:00+00:00", "2019-04-21 16:00:00+00:00", "2019-04-21 15:00:00+00:00", "2019-04-21 14:00:00+00:00", "2019-04-21 13:00:00+00:00", "2019-04-21 12:00:00+00:00", "2019-04-21 11:00:00+00:00", "2019-04-21 10:00:00+00:00", "2019-04-21 09:00:00+00:00", "2019-04-21 08:00:00+00:00", "2019-04-21 07:00:00+00:00", "2019-04-21 06:00:00+00:00", "2019-04-21 05:00:00+00:00", "2019-04-21 04:00:00+00:00", "2019-04-21 03:00:00+00:00", "2019-04-21 02:00:00+00:00", "2019-04-21 01:00:00+00:00", "2019-04-21 00:00:00+00:00", "2019-04-20 23:00:00+00:00", "2019-04-20 22:00:00+00:00", "2019-04-20 21:00:00+00:00", "2019-04-20 20:00:00+00:00", "2019-04-20 19:00:00+00:00", "2019-04-20 18:00:00+00:00", "2019-04-20 17:00:00+00:00", "2019-04-20 16:00:00+00:00", "2019-04-20 15:00:00+00:00", "2019-04-20 14:00:00+00:00", "2019-04-20 13:00:00+00:00", "2019-04-20 12:00:00+00:00", "2019-04-20 11:00:00+00:00", "2019-04-20 10:00:00+00:00", "2019-04-20 09:00:00+00:00", "2019-04-20 08:00:00+00:00", "2019-04-20 07:00:00+00:00", "2019-04-20 06:00:00+00:00", "2019-04-20 05:00:00+00:00", "2019-04-20 04:00:00+00:00", "2019-04-20 03:00:00+00:00", "2019-04-20 02:00:00+00:00", "2019-04-20 01:00:00+00:00", "2019-04-20 00:00:00+00:00", "2019-04-19 23:00:00+00:00", "2019-04-19 22:00:00+00:00", "2019-04-19 21:00:00+00:00", "2019-04-19 20:00:00+00:00", "2019-04-19 19:00:00+00:00", "2019-04-19 18:00:00+00:00", "2019-04-19 17:00:00+00:00", "2019-04-19 16:00:00+00:00", "2019-04-19 15:00:00+00:00", "2019-04-19 14:00:00+00:00", "2019-04-19 13:00:00+00:00", "2019-04-19 12:00:00+00:00", "2019-04-19 11:00:00+00:00", "2019-04-19 10:00:00+00:00", "2019-04-19 09:00:00+00:00", "2019-04-19 08:00:00+00:00", "2019-04-19 07:00:00+00:00", "2019-04-19 06:00:00+00:00", "2019-04-19 05:00:00+00:00", "2019-04-19 04:00:00+00:00", "2019-04-19 03:00:00+00:00", "2019-04-19 02:00:00+00:00", "2019-04-19 00:00:00+00:00", "2019-04-18 23:00:00+00:00", "2019-04-18 22:00:00+00:00", "2019-04-18 21:00:00+00:00", "2019-04-18 20:00:00+00:00", "2019-04-18 19:00:00+00:00", "2019-04-18 18:00:00+00:00", "2019-04-18 17:00:00+00:00", "2019-04-18 16:00:00+00:00", "2019-04-18 15:00:00+00:00", "2019-04-18 14:00:00+00:00", "2019-04-18 13:00:00+00:00", "2019-04-18 12:00:00+00:00", "2019-04-18 11:00:00+00:00", "2019-04-18 10:00:00+00:00", "2019-04-18 09:00:00+00:00", "2019-04-18 08:00:00+00:00", "2019-04-18 07:00:00+00:00", "2019-04-18 06:00:00+00:00", "2019-04-18 05:00:00+00:00", "2019-04-18 04:00:00+00:00", "2019-04-18 03:00:00+00:00", "2019-04-18 02:00:00+00:00", "2019-04-18 01:00:00+00:00", "2019-04-18 00:00:00+00:00", "2019-04-17 23:00:00+00:00", "2019-04-17 22:00:00+00:00", "2019-04-17 21:00:00+00:00", "2019-04-17 20:00:00+00:00", "2019-04-17 19:00:00+00:00", "2019-04-17 18:00:00+00:00", "2019-04-17 17:00:00+00:00", "2019-04-17 16:00:00+00:00", "2019-04-17 15:00:00+00:00", "2019-04-17 14:00:00+00:00", "2019-04-17 13:00:00+00:00", "2019-04-17 12:00:00+00:00", "2019-04-17 11:00:00+00:00", "2019-04-17 10:00:00+00:00", "2019-04-17 09:00:00+00:00", "2019-04-17 08:00:00+00:00", "2019-04-17 07:00:00+00:00", "2019-04-17 06:00:00+00:00", "2019-04-17 05:00:00+00:00", "2019-04-17 04:00:00+00:00", "2019-04-17 03:00:00+00:00", "2019-04-17 02:00:00+00:00", "2019-04-17 00:00:00+00:00", "2019-04-16 23:00:00+00:00", "2019-04-16 22:00:00+00:00", "2019-04-16 21:00:00+00:00", "2019-04-16 20:00:00+00:00", "2019-04-16 19:00:00+00:00", "2019-04-16 18:00:00+00:00", "2019-04-16 17:00:00+00:00", "2019-04-16 15:00:00+00:00", "2019-04-16 14:00:00+00:00", "2019-04-16 13:00:00+00:00", "2019-04-16 12:00:00+00:00", "2019-04-16 11:00:00+00:00", "2019-04-16 10:00:00+00:00", "2019-04-16 09:00:00+00:00", "2019-04-16 08:00:00+00:00", "2019-04-16 07:00:00+00:00", "2019-04-16 06:00:00+00:00", "2019-04-16 05:00:00+00:00", "2019-04-16 04:00:00+00:00", "2019-04-16 03:00:00+00:00", "2019-04-16 02:00:00+00:00", "2019-04-16 00:00:00+00:00", "2019-04-15 23:00:00+00:00", "2019-04-15 22:00:00+00:00", "2019-04-15 21:00:00+00:00", "2019-04-15 20:00:00+00:00", "2019-04-15 19:00:00+00:00", "2019-04-15 18:00:00+00:00", "2019-04-15 17:00:00+00:00", "2019-04-15 16:00:00+00:00", "2019-04-15 15:00:00+00:00", "2019-04-15 14:00:00+00:00", "2019-04-15 13:00:00+00:00", "2019-04-15 12:00:00+00:00", "2019-04-15 11:00:00+00:00", "2019-04-15 10:00:00+00:00", "2019-04-15 09:00:00+00:00", "2019-04-15 08:00:00+00:00", "2019-04-15 07:00:00+00:00", "2019-04-15 06:00:00+00:00", "2019-04-15 05:00:00+00:00", "2019-04-15 04:00:00+00:00", "2019-04-15 03:00:00+00:00", "2019-04-15 02:00:00+00:00", "2019-04-15 01:00:00+00:00", "2019-04-15 00:00:00+00:00", "2019-04-14 23:00:00+00:00", "2019-04-14 22:00:00+00:00", "2019-04-14 21:00:00+00:00", "2019-04-14 20:00:00+00:00", "2019-04-14 19:00:00+00:00", "2019-04-14 18:00:00+00:00", "2019-04-14 17:00:00+00:00", "2019-04-14 16:00:00+00:00", "2019-04-14 15:00:00+00:00", "2019-04-14 14:00:00+00:00", "2019-04-14 13:00:00+00:00", "2019-04-14 12:00:00+00:00", "2019-04-14 11:00:00+00:00", "2019-04-14 10:00:00+00:00", "2019-04-14 09:00:00+00:00", "2019-04-14 08:00:00+00:00", "2019-04-14 07:00:00+00:00", "2019-04-14 06:00:00+00:00", "2019-04-14 05:00:00+00:00", "2019-04-14 04:00:00+00:00", "2019-04-14 03:00:00+00:00", "2019-04-14 02:00:00+00:00", "2019-04-14 01:00:00+00:00", "2019-04-14 00:00:00+00:00", "2019-04-13 23:00:00+00:00", "2019-04-13 22:00:00+00:00", "2019-04-13 21:00:00+00:00", "2019-04-13 20:00:00+00:00", "2019-04-13 19:00:00+00:00", "2019-04-13 18:00:00+00:00", "2019-04-13 17:00:00+00:00", "2019-04-13 16:00:00+00:00", "2019-04-13 15:00:00+00:00", "2019-04-13 14:00:00+00:00", "2019-04-13 13:00:00+00:00", "2019-04-13 12:00:00+00:00", "2019-04-13 11:00:00+00:00", "2019-04-13 10:00:00+00:00", "2019-04-13 09:00:00+00:00", "2019-04-13 08:00:00+00:00", "2019-04-13 07:00:00+00:00", "2019-04-13 06:00:00+00:00", "2019-04-13 05:00:00+00:00", "2019-04-13 04:00:00+00:00", "2019-04-13 03:00:00+00:00", "2019-04-13 02:00:00+00:00", "2019-04-13 01:00:00+00:00", "2019-04-13 00:00:00+00:00", "2019-04-12 23:00:00+00:00", "2019-04-12 22:00:00+00:00", "2019-04-12 21:00:00+00:00", "2019-04-12 20:00:00+00:00", "2019-04-12 19:00:00+00:00", "2019-04-12 18:00:00+00:00", "2019-04-12 17:00:00+00:00", "2019-04-12 16:00:00+00:00", "2019-04-12 15:00:00+00:00", "2019-04-12 14:00:00+00:00", "2019-04-12 13:00:00+00:00", "2019-04-12 12:00:00+00:00", "2019-04-12 11:00:00+00:00", "2019-04-12 10:00:00+00:00", "2019-04-12 09:00:00+00:00", "2019-04-12 08:00:00+00:00", "2019-04-12 07:00:00+00:00", "2019-04-12 06:00:00+00:00", "2019-04-12 05:00:00+00:00", "2019-04-12 04:00:00+00:00", "2019-04-12 03:00:00+00:00", "2019-04-12 00:00:00+00:00", "2019-04-11 23:00:00+00:00", "2019-04-11 22:00:00+00:00", "2019-04-11 21:00:00+00:00", "2019-04-11 20:00:00+00:00", "2019-04-11 19:00:00+00:00", "2019-04-11 18:00:00+00:00", "2019-04-11 17:00:00+00:00", "2019-04-11 16:00:00+00:00", "2019-04-11 15:00:00+00:00", "2019-04-11 14:00:00+00:00", "2019-04-11 13:00:00+00:00", "2019-04-11 12:00:00+00:00", "2019-04-11 11:00:00+00:00", "2019-04-11 10:00:00+00:00", "2019-04-11 09:00:00+00:00", "2019-04-11 08:00:00+00:00", "2019-04-11 07:00:00+00:00", "2019-04-11 06:00:00+00:00", "2019-04-11 05:00:00+00:00", "2019-04-11 04:00:00+00:00", "2019-04-11 03:00:00+00:00", "2019-04-11 02:00:00+00:00", "2019-04-11 00:00:00+00:00", "2019-04-10 23:00:00+00:00", "2019-04-10 22:00:00+00:00", "2019-04-10 21:00:00+00:00", "2019-04-10 20:00:00+00:00", "2019-04-10 19:00:00+00:00", "2019-04-10 18:00:00+00:00", "2019-04-10 17:00:00+00:00", "2019-04-10 16:00:00+00:00", "2019-04-10 15:00:00+00:00", "2019-04-10 14:00:00+00:00", "2019-04-10 13:00:00+00:00", "2019-04-10 12:00:00+00:00", "2019-04-10 11:00:00+00:00", "2019-04-10 10:00:00+00:00", "2019-04-10 09:00:00+00:00", "2019-04-10 08:00:00+00:00", "2019-04-10 07:00:00+00:00", "2019-04-10 06:00:00+00:00", "2019-04-10 05:00:00+00:00", "2019-04-10 04:00:00+00:00", "2019-04-10 03:00:00+00:00", "2019-04-10 02:00:00+00:00", "2019-04-10 01:00:00+00:00", "2019-04-10 00:00:00+00:00", "2019-04-09 23:00:00+00:00", "2019-04-09 22:00:00+00:00", "2019-04-09 21:00:00+00:00", "2019-04-09 20:00:00+00:00", "2019-04-09 19:00:00+00:00", "2019-04-09 18:00:00+00:00", "2019-04-09 17:00:00+00:00", "2019-04-09 16:00:00+00:00", "2019-04-09 15:00:00+00:00", "2019-04-09 14:00:00+00:00", "2019-04-09 13:00:00+00:00", "2019-04-09 12:00:00+00:00", "2019-04-09 11:00:00+00:00", "2019-04-09 10:00:00+00:00", "2019-04-09 09:00:00+00:00", "2019-04-09 08:00:00+00:00", "2019-04-09 07:00:00+00:00", "2019-04-09 06:00:00+00:00", "2019-04-09 05:00:00+00:00", "2019-04-09 04:00:00+00:00", "2019-04-09 03:00:00+00:00", "2019-04-09 02:00:00+00:00",
// 					},
// 					"c367804c2f44cd55a157c97d086a1b146f7a1f5949e1fbff4a01762ff17466d8ffc389ce5dbf62437b7d1041aba3aa33919ba40b7528f92542538127a1fe73b0": {
// 						"2019-06-18 06:00:00+00:00", "2019-06-17 08:00:00+00:00", "2019-06-17 07:00:00+00:00", "2019-06-17 06:00:00+00:00", "2019-06-17 05:00:00+00:00", "2019-06-17 04:00:00+00:00", "2019-06-17 03:00:00+00:00", "2019-06-17 02:00:00+00:00", "2019-06-17 01:00:00+00:00", "2019-06-16 01:00:00+00:00", "2019-06-15 01:00:00+00:00", "2019-06-14 09:00:00+00:00", "2019-06-13 01:00:00+00:00", "2019-06-12 01:00:00+00:00", "2019-06-11 01:00:00+00:00", "2019-06-10 01:00:00+00:00", "2019-06-09 01:00:00+00:00", "2019-06-08 01:00:00+00:00", "2019-06-06 01:00:00+00:00", "2019-06-05 01:00:00+00:00", "2019-06-04 01:00:00+00:00", "2019-06-03 01:00:00+00:00", "2019-06-02 01:00:00+00:00", "2019-06-01 01:00:00+00:00", "2019-05-31 01:00:00+00:00", "2019-05-30 01:00:00+00:00", "2019-05-29 01:00:00+00:00", "2019-05-28 01:00:00+00:00", "2019-05-27 01:00:00+00:00", "2019-05-26 01:00:00+00:00", "2019-05-25 01:00:00+00:00", "2019-05-24 01:00:00+00:00", "2019-05-23 01:00:00+00:00", "2019-05-22 01:00:00+00:00", "2019-05-21 01:00:00+00:00", "2019-05-20 17:00:00+00:00", "2019-05-20 16:00:00+00:00", "2019-05-20 15:00:00+00:00", "2019-05-20 14:00:00+00:00", "2019-05-20 13:00:00+00:00", "2019-05-20 12:00:00+00:00", "2019-05-20 11:00:00+00:00", "2019-05-20 10:00:00+00:00", "2019-05-20 09:00:00+00:00", "2019-05-20 08:00:00+00:00", "2019-05-20 07:00:00+00:00", "2019-05-20 06:00:00+00:00", "2019-05-20 05:00:00+00:00", "2019-05-20 04:00:00+00:00", "2019-05-20 03:00:00+00:00", "2019-05-20 02:00:00+00:00", "2019-05-20 01:00:00+00:00", "2019-05-19 21:00:00+00:00", "2019-05-19 20:00:00+00:00", "2019-05-19 19:00:00+00:00", "2019-05-19 18:00:00+00:00", "2019-05-19 17:00:00+00:00", "2019-05-19 16:00:00+00:00", "2019-05-19 15:00:00+00:00", "2019-05-19 14:00:00+00:00", "2019-05-19 13:00:00+00:00", "2019-05-19 12:00:00+00:00", "2019-05-19 11:00:00+00:00", "2019-05-19 10:00:00+00:00", "2019-05-19 09:00:00+00:00", "2019-05-19 08:00:00+00:00", "2019-05-19 07:00:00+00:00", "2019-05-19 06:00:00+00:00", "2019-05-19 05:00:00+00:00", "2019-05-19 04:00:00+00:00", "2019-05-19 03:00:00+00:00", "2019-05-19 02:00:00+00:00", "2019-05-19 01:00:00+00:00", "2019-05-19 00:00:00+00:00", "2019-05-18 23:00:00+00:00", "2019-05-18 22:00:00+00:00", "2019-05-18 21:00:00+00:00", "2019-05-18 20:00:00+00:00", "2019-05-18 19:00:00+00:00", "2019-05-18 18:00:00+00:00", "2019-05-18 01:00:00+00:00", "2019-05-17 01:00:00+00:00", "2019-05-16 01:00:00+00:00", "2019-05-15 02:00:00+00:00", "2019-05-15 01:00:00+00:00", "2019-05-14 02:00:00+00:00", "2019-05-14 01:00:00+00:00", "2019-05-13 02:00:00+00:00", "2019-05-13 01:00:00+00:00", "2019-05-12 02:00:00+00:00", "2019-05-12 01:00:00+00:00", "2019-05-11 02:00:00+00:00", "2019-05-11 01:00:00+00:00", "2019-05-10 02:00:00+00:00", "2019-05-10 01:00:00+00:00", "2019-05-09 02:00:00+00:00", "2019-05-09 01:00:00+00:00", "2019-05-08 02:00:00+00:00", "2019-05-08 01:00:00+00:00", "2019-05-07 02:00:00+00:00", "2019-05-07 01:00:00+00:00", "2019-05-06 02:00:00+00:00", "2019-05-06 01:00:00+00:00", "2019-05-05 02:00:00+00:00", "2019-05-05 01:00:00+00:00", "2019-05-04 02:00:00+00:00", "2019-05-04 01:00:00+00:00", "2019-05-03 02:00:00+00:00", "2019-05-03 01:00:00+00:00", "2019-05-02 02:00:00+00:00", "2019-05-02 01:00:00+00:00", "2019-05-01 02:00:00+00:00", "2019-05-01 01:00:00+00:00", "2019-04-30 02:00:00+00:00", "2019-04-30 01:00:00+00:00", "2019-04-29 02:00:00+00:00", "2019-04-29 01:00:00+00:00", "2019-04-28 02:00:00+00:00", "2019-04-28 01:00:00+00:00", "2019-04-27 02:00:00+00:00", "2019-04-27 01:00:00+00:00", "2019-04-26 02:00:00+00:00", "2019-04-26 01:00:00+00:00", "2019-04-25 02:00:00+00:00", "2019-04-25 01:00:00+00:00", "2019-04-24 02:00:00+00:00", "2019-04-24 01:00:00+00:00", "2019-04-23 02:00:00+00:00", "2019-04-23 01:00:00+00:00", "2019-04-22 02:00:00+00:00", "2019-04-22 01:00:00+00:00", "2019-04-21 02:00:00+00:00", "2019-04-21 01:00:00+00:00", "2019-04-20 02:00:00+00:00", "2019-04-20 01:00:00+00:00", "2019-04-19 01:00:00+00:00", "2019-04-18 02:00:00+00:00", "2019-04-18 01:00:00+00:00", "2019-04-17 03:00:00+00:00", "2019-04-17 02:00:00+00:00", "2019-04-17 01:00:00+00:00", "2019-04-16 02:00:00+00:00", "2019-04-16 01:00:00+00:00", "2019-04-15 15:00:00+00:00", "2019-04-15 14:00:00+00:00", "2019-04-15 13:00:00+00:00", "2019-04-15 12:00:00+00:00", "2019-04-15 11:00:00+00:00", "2019-04-15 10:00:00+00:00", "2019-04-15 09:00:00+00:00", "2019-04-15 08:00:00+00:00", "2019-04-15 07:00:00+00:00", "2019-04-15 06:00:00+00:00", "2019-04-15 05:00:00+00:00", "2019-04-15 04:00:00+00:00", "2019-04-15 03:00:00+00:00", "2019-04-15 02:00:00+00:00", "2019-04-15 01:00:00+00:00", "2019-04-12 02:00:00+00:00", "2019-04-12 01:00:00+00:00", "2019-04-11 02:00:00+00:00", "2019-04-11 01:00:00+00:00", "2019-04-10 02:00:00+00:00", "2019-04-10 01:00:00+00:00", "2019-04-09 13:00:00+00:00", "2019-04-09 12:00:00+00:00", "2019-04-09 11:00:00+00:00", "2019-04-09 10:00:00+00:00", "2019-04-09 09:00:00+00:00", "2019-04-09 08:00:00+00:00", "2019-04-09 07:00:00+00:00", "2019-04-09 06:00:00+00:00", "2019-04-09 05:00:00+00:00", "2019-04-09 04:00:00+00:00", "2019-04-09 03:00:00+00:00", "2019-04-09 02:00:00+00:00", "2019-04-09 01:00:00+00:00",
// 					},
// 					"6e2143717dc23a5087b440b39ee01182896c37449603628e1e2f67c3e80b7eefe9a84fdb3a664a7106120dd9fcc26a5ba3460610b41a1fa085f61fee971e1b1b": {
// 						"2019-06-21 00:00:00+00:00", "2019-06-20 23:00:00+00:00", "2019-06-20 22:00:00+00:00", "2019-06-20 21:00:00+00:00", "2019-06-20 20:00:00+00:00", "2019-06-20 19:00:00+00:00", "2019-06-20 18:00:00+00:00", "2019-06-20 17:00:00+00:00", "2019-06-20 16:00:00+00:00", "2019-06-20 15:00:00+00:00", "2019-06-19 13:00:00+00:00", "2019-06-19 12:00:00+00:00", "2019-06-19 11:00:00+00:00", "2019-06-19 00:00:00+00:00", "2019-06-18 23:00:00+00:00", "2019-06-18 22:00:00+00:00", "2019-06-18 21:00:00+00:00", "2019-06-18 11:00:00+00:00", "2019-06-18 10:00:00+00:00", "2019-06-18 09:00:00+00:00", "2019-06-18 08:00:00+00:00", "2019-06-18 07:00:00+00:00", "2019-06-18 06:00:00+00:00", "2019-06-18 05:00:00+00:00", "2019-06-18 04:00:00+00:00", "2019-06-18 03:00:00+00:00", "2019-06-18 02:00:00+00:00", "2019-06-18 00:00:00+00:00", "2019-06-17 23:00:00+00:00", "2019-06-17 22:00:00+00:00", "2019-06-17 21:00:00+00:00", "2019-06-17 20:00:00+00:00", "2019-06-17 19:00:00+00:00", "2019-06-17 18:00:00+00:00", "2019-06-17 17:00:00+00:00", "2019-06-17 16:00:00+00:00", "2019-06-17 15:00:00+00:00", "2019-06-17 14:00:00+00:00", "2019-06-17 13:00:00+00:00", "2019-06-17 12:00:00+00:00", "2019-06-17 11:00:00+00:00", "2019-06-17 10:00:00+00:00", "2019-06-17 09:00:00+00:00", "2019-06-17 08:00:00+00:00", "2019-06-17 07:00:00+00:00", "2019-06-17 06:00:00+00:00", "2019-06-17 05:00:00+00:00", "2019-06-17 04:00:00+00:00", "2019-06-17 03:00:00+00:00", "2019-06-17 02:00:00+00:00", "2019-06-17 01:00:00+00:00", "2019-06-17 00:00:00+00:00", "2019-06-16 23:00:00+00:00", "2019-06-16 21:00:00+00:00", "2019-06-16 20:00:00+00:00", "2019-06-16 19:00:00+00:00", "2019-06-16 18:00:00+00:00", "2019-06-16 17:00:00+00:00", "2019-06-16 16:00:00+00:00", "2019-06-16 15:00:00+00:00", "2019-06-16 14:00:00+00:00", "2019-06-16 13:00:00+00:00", "2019-06-16 12:00:00+00:00", "2019-06-16 11:00:00+00:00", "2019-06-16 10:00:00+00:00", "2019-06-16 09:00:00+00:00", "2019-06-16 08:00:00+00:00", "2019-06-16 07:00:00+00:00", "2019-06-16 06:00:00+00:00", "2019-06-16 05:00:00+00:00", "2019-06-16 04:00:00+00:00", "2019-06-16 03:00:00+00:00", "2019-06-16 02:00:00+00:00", "2019-06-16 01:00:00+00:00", "2019-06-16 00:00:00+00:00", "2019-06-15 23:00:00+00:00", "2019-06-15 22:00:00+00:00", "2019-06-15 21:00:00+00:00", "2019-06-15 20:00:00+00:00", "2019-06-15 19:00:00+00:00", "2019-06-15 18:00:00+00:00", "2019-06-15 17:00:00+00:00", "2019-06-15 16:00:00+00:00", "2019-06-15 15:00:00+00:00", "2019-06-15 14:00:00+00:00", "2019-06-15 13:00:00+00:00", "2019-06-15 12:00:00+00:00", "2019-06-15 11:00:00+00:00", "2019-06-15 10:00:00+00:00", "2019-06-15 09:00:00+00:00", "2019-06-15 08:00:00+00:00", "2019-06-15 07:00:00+00:00", "2019-06-15 06:00:00+00:00", "2019-06-15 05:00:00+00:00", "2019-06-15 04:00:00+00:00", "2019-06-15 00:00:00+00:00", "2019-06-14 23:00:00+00:00", "2019-06-14 22:00:00+00:00", "2019-06-14 21:00:00+00:00", "2019-06-14 20:00:00+00:00", "2019-06-14 19:00:00+00:00", "2019-06-14 18:00:00+00:00", "2019-06-14 17:00:00+00:00", "2019-06-14 16:00:00+00:00", "2019-06-14 15:00:00+00:00", "2019-06-14 14:00:00+00:00", "2019-06-14 13:00:00+00:00", "2019-06-14 12:00:00+00:00", "2019-06-14 11:00:00+00:00", "2019-06-14 10:00:00+00:00", "2019-06-14 09:00:00+00:00", "2019-06-14 08:00:00+00:00", "2019-06-14 07:00:00+00:00", "2019-06-14 06:00:00+00:00", "2019-06-14 05:00:00+00:00", "2019-06-14 04:00:00+00:00", "2019-06-14 03:00:00+00:00", "2019-06-14 02:00:00+00:00", "2019-06-14 00:00:00+00:00", "2019-06-13 23:00:00+00:00", "2019-06-13 22:00:00+00:00", "2019-06-13 21:00:00+00:00", "2019-06-13 20:00:00+00:00", "2019-06-13 19:00:00+00:00", "2019-06-13 18:00:00+00:00", "2019-06-13 17:00:00+00:00", "2019-06-13 16:00:00+00:00", "2019-06-13 15:00:00+00:00", "2019-06-13 14:00:00+00:00", "2019-06-13 13:00:00+00:00", "2019-06-13 12:00:00+00:00", "2019-06-13 11:00:00+00:00", "2019-06-13 10:00:00+00:00", "2019-06-13 09:00:00+00:00", "2019-06-13 08:00:00+00:00", "2019-06-13 07:00:00+00:00", "2019-06-13 06:00:00+00:00", "2019-06-13 05:00:00+00:00", "2019-06-13 04:00:00+00:00", "2019-06-13 03:00:00+00:00", "2019-06-13 02:00:00+00:00", "2019-06-13 00:00:00+00:00", "2019-06-12 23:00:00+00:00", "2019-06-12 21:00:00+00:00", "2019-06-12 20:00:00+00:00", "2019-06-12 19:00:00+00:00", "2019-06-12 18:00:00+00:00", "2019-06-12 17:00:00+00:00", "2019-06-12 16:00:00+00:00", "2019-06-12 15:00:00+00:00", "2019-06-12 14:00:00+00:00", "2019-06-12 13:00:00+00:00", "2019-06-12 12:00:00+00:00", "2019-06-12 11:00:00+00:00", "2019-06-12 10:00:00+00:00", "2019-06-12 09:00:00+00:00", "2019-06-12 08:00:00+00:00", "2019-06-12 07:00:00+00:00", "2019-06-12 06:00:00+00:00", "2019-06-12 05:00:00+00:00", "2019-06-12 04:00:00+00:00", "2019-06-12 03:00:00+00:00", "2019-06-12 00:00:00+00:00", "2019-06-11 23:00:00+00:00", "2019-06-11 22:00:00+00:00", "2019-06-11 21:00:00+00:00", "2019-06-11 20:00:00+00:00", "2019-06-11 19:00:00+00:00", "2019-06-11 18:00:00+00:00", "2019-06-11 17:00:00+00:00", "2019-06-11 16:00:00+00:00", "2019-06-11 15:00:00+00:00", "2019-06-11 14:00:00+00:00", "2019-06-11 13:00:00+00:00", "2019-06-11 12:00:00+00:00", "2019-06-11 11:00:00+00:00", "2019-06-11 10:00:00+00:00", "2019-06-11 09:00:00+00:00", "2019-06-11 08:00:00+00:00", "2019-06-11 07:00:00+00:00", "2019-06-11 06:00:00+00:00", "2019-06-11 05:00:00+00:00", "2019-06-11 04:00:00+00:00", "2019-06-11 03:00:00+00:00", "2019-06-11 02:00:00+00:00", "2019-06-11 01:00:00+00:00", "2019-06-11 00:00:00+00:00", "2019-06-10 23:00:00+00:00", "2019-06-10 22:00:00+00:00", "2019-06-10 21:00:00+00:00", "2019-06-10 20:00:00+00:00", "2019-06-10 19:00:00+00:00", "2019-06-10 18:00:00+00:00", "2019-06-10 17:00:00+00:00", "2019-06-10 16:00:00+00:00", "2019-06-10 15:00:00+00:00", "2019-06-10 14:00:00+00:00", "2019-06-10 13:00:00+00:00", "2019-06-10 12:00:00+00:00", "2019-06-10 11:00:00+00:00", "2019-06-10 10:00:00+00:00", "2019-06-10 09:00:00+00:00", "2019-06-10 08:00:00+00:00", "2019-06-10 07:00:00+00:00", "2019-06-10 06:00:00+00:00", "2019-06-10 05:00:00+00:00", "2019-06-10 04:00:00+00:00", "2019-06-10 03:00:00+00:00", "2019-06-10 02:00:00+00:00", "2019-06-10 01:00:00+00:00", "2019-06-10 00:00:00+00:00", "2019-06-09 23:00:00+00:00", "2019-06-09 21:00:00+00:00", "2019-06-09 20:00:00+00:00", "2019-06-09 19:00:00+00:00", "2019-06-09 18:00:00+00:00", "2019-06-09 17:00:00+00:00", "2019-06-09 16:00:00+00:00", "2019-06-09 15:00:00+00:00", "2019-06-09 14:00:00+00:00", "2019-06-09 13:00:00+00:00", "2019-06-09 12:00:00+00:00", "2019-06-09 11:00:00+00:00", "2019-06-09 10:00:00+00:00", "2019-06-09 09:00:00+00:00", "2019-06-09 08:00:00+00:00", "2019-06-09 07:00:00+00:00", "2019-06-09 06:00:00+00:00", "2019-06-09 05:00:00+00:00", "2019-06-09 04:00:00+00:00", "2019-06-09 03:00:00+00:00", "2019-06-09 02:00:00+00:00", "2019-06-09 01:00:00+00:00", "2019-06-09 00:00:00+00:00", "2019-06-08 23:00:00+00:00", "2019-06-08 21:00:00+00:00", "2019-06-08 20:00:00+00:00", "2019-06-08 19:00:00+00:00", "2019-06-08 18:00:00+00:00", "2019-06-08 17:00:00+00:00", "2019-06-08 16:00:00+00:00", "2019-06-08 15:00:00+00:00", "2019-06-08 14:00:00+00:00", "2019-06-08 13:00:00+00:00", "2019-06-08 12:00:00+00:00", "2019-06-08 11:00:00+00:00", "2019-06-08 10:00:00+00:00", "2019-06-08 09:00:00+00:00", "2019-06-08 08:00:00+00:00", "2019-06-08 07:00:00+00:00", "2019-06-08 06:00:00+00:00", "2019-06-08 05:00:00+00:00", "2019-06-08 04:00:00+00:00", "2019-06-08 03:00:00+00:00", "2019-06-08 02:00:00+00:00", "2019-06-08 00:00:00+00:00", "2019-06-07 23:00:00+00:00", "2019-06-07 21:00:00+00:00", "2019-06-07 20:00:00+00:00", "2019-06-07 19:00:00+00:00", "2019-06-07 18:00:00+00:00", "2019-06-07 17:00:00+00:00", "2019-06-07 16:00:00+00:00", "2019-06-07 15:00:00+00:00", "2019-06-07 14:00:00+00:00", "2019-06-07 13:00:00+00:00", "2019-06-07 12:00:00+00:00", "2019-06-07 11:00:00+00:00", "2019-06-07 10:00:00+00:00", "2019-06-07 09:00:00+00:00", "2019-06-07 08:00:00+00:00", "2019-06-07 07:00:00+00:00", "2019-06-07 06:00:00+00:00", "2019-06-07 05:00:00+00:00", "2019-06-07 04:00:00+00:00", "2019-06-07 03:00:00+00:00", "2019-06-07 02:00:00+00:00", "2019-06-07 01:00:00+00:00", "2019-06-07 00:00:00+00:00", "2019-06-06 23:00:00+00:00", "2019-06-06 22:00:00+00:00", "2019-06-06 21:00:00+00:00", "2019-06-06 20:00:00+00:00", "2019-06-06 19:00:00+00:00", "2019-06-06 18:00:00+00:00", "2019-06-06 17:00:00+00:00", "2019-06-06 16:00:00+00:00", "2019-06-06 15:00:00+00:00", "2019-06-06 14:00:00+00:00", "2019-06-06 13:00:00+00:00", "2019-06-06 12:00:00+00:00", "2019-06-06 11:00:00+00:00", "2019-06-06 10:00:00+00:00", "2019-06-06 09:00:00+00:00", "2019-06-06 08:00:00+00:00", "2019-06-06 07:00:00+00:00", "2019-06-06 06:00:00+00:00", "2019-06-06 05:00:00+00:00", "2019-06-06 04:00:00+00:00", "2019-06-06 03:00:00+00:00", "2019-06-06 02:00:00+00:00", "2019-06-06 00:00:00+00:00", "2019-06-05 23:00:00+00:00", "2019-06-05 22:00:00+00:00", "2019-06-05 21:00:00+00:00", "2019-06-05 20:00:00+00:00", "2019-06-05 19:00:00+00:00", "2019-06-05 18:00:00+00:00", "2019-06-05 17:00:00+00:00", "2019-06-05 16:00:00+00:00", "2019-06-05 15:00:00+00:00", "2019-06-05 14:00:00+00:00", "2019-06-05 13:00:00+00:00", "2019-06-05 12:00:00+00:00", "2019-06-05 11:00:00+00:00", "2019-06-05 10:00:00+00:00", "2019-06-05 09:00:00+00:00", "2019-06-05 08:00:00+00:00", "2019-06-05 07:00:00+00:00", "2019-06-05 06:00:00+00:00", "2019-06-05 05:00:00+00:00", "2019-06-05 04:00:00+00:00", "2019-06-05 03:00:00+00:00", "2019-06-05 02:00:00+00:00", "2019-06-05 01:00:00+00:00", "2019-06-05 00:00:00+00:00", "2019-06-04 23:00:00+00:00", "2019-06-04 22:00:00+00:00", "2019-06-04 21:00:00+00:00", "2019-06-04 20:00:00+00:00", "2019-06-04 19:00:00+00:00", "2019-06-04 18:00:00+00:00", "2019-06-04 17:00:00+00:00", "2019-06-04 16:00:00+00:00", "2019-06-04 15:00:00+00:00", "2019-06-04 14:00:00+00:00", "2019-06-04 13:00:00+00:00", "2019-06-04 12:00:00+00:00", "2019-06-04 11:00:00+00:00", "2019-06-04 10:00:00+00:00", "2019-06-04 09:00:00+00:00", "2019-06-04 08:00:00+00:00", "2019-06-04 07:00:00+00:00", "2019-06-04 06:00:00+00:00", "2019-06-04 05:00:00+00:00", "2019-06-04 04:00:00+00:00", "2019-06-04 03:00:00+00:00", "2019-06-04 02:00:00+00:00", "2019-06-04 01:00:00+00:00", "2019-06-04 00:00:00+00:00", "2019-06-03 23:00:00+00:00", "2019-06-03 22:00:00+00:00", "2019-06-03 21:00:00+00:00", "2019-06-03 20:00:00+00:00", "2019-06-03 19:00:00+00:00", "2019-06-03 18:00:00+00:00", "2019-06-03 17:00:00+00:00", "2019-06-03 16:00:00+00:00", "2019-06-03 15:00:00+00:00", "2019-06-03 14:00:00+00:00", "2019-06-03 13:00:00+00:00", "2019-06-03 12:00:00+00:00", "2019-06-03 11:00:00+00:00", "2019-06-03 10:00:00+00:00", "2019-06-03 09:00:00+00:00", "2019-06-03 08:00:00+00:00", "2019-06-03 07:00:00+00:00", "2019-06-03 06:00:00+00:00", "2019-06-03 05:00:00+00:00", "2019-06-03 04:00:00+00:00", "2019-06-03 03:00:00+00:00", "2019-06-03 02:00:00+00:00", "2019-06-03 01:00:00+00:00", "2019-06-03 00:00:00+00:00", "2019-06-02 23:00:00+00:00", "2019-06-02 22:00:00+00:00", "2019-06-02 21:00:00+00:00", "2019-06-02 20:00:00+00:00", "2019-06-02 19:00:00+00:00", "2019-06-02 18:00:00+00:00", "2019-06-02 17:00:00+00:00", "2019-06-02 16:00:00+00:00", "2019-06-02 15:00:00+00:00", "2019-06-02 14:00:00+00:00", "2019-06-02 13:00:00+00:00", "2019-06-02 12:00:00+00:00", "2019-06-02 11:00:00+00:00", "2019-06-02 10:00:00+00:00", "2019-06-02 09:00:00+00:00", "2019-06-02 08:00:00+00:00", "2019-06-02 07:00:00+00:00", "2019-06-02 06:00:00+00:00", "2019-06-02 05:00:00+00:00", "2019-06-02 04:00:00+00:00", "2019-06-02 03:00:00+00:00", "2019-06-02 02:00:00+00:00", "2019-06-02 01:00:00+00:00", "2019-06-02 00:00:00+00:00", "2019-06-01 23:00:00+00:00", "2019-06-01 22:00:00+00:00", "2019-06-01 21:00:00+00:00", "2019-06-01 20:00:00+00:00", "2019-06-01 19:00:00+00:00", "2019-06-01 18:00:00+00:00", "2019-06-01 17:00:00+00:00", "2019-06-01 16:00:00+00:00", "2019-06-01 15:00:00+00:00", "2019-06-01 14:00:00+00:00", "2019-06-01 13:00:00+00:00", "2019-06-01 12:00:00+00:00", "2019-06-01 11:00:00+00:00", "2019-06-01 10:00:00+00:00", "2019-06-01 09:00:00+00:00", "2019-06-01 08:00:00+00:00", "2019-06-01 07:00:00+00:00", "2019-06-01 06:00:00+00:00", "2019-06-01 05:00:00+00:00", "2019-06-01 04:00:00+00:00", "2019-06-01 03:00:00+00:00", "2019-06-01 02:00:00+00:00", "2019-06-01 01:00:00+00:00", "2019-06-01 00:00:00+00:00", "2019-05-31 23:00:00+00:00", "2019-05-31 22:00:00+00:00", "2019-05-31 21:00:00+00:00", "2019-05-31 20:00:00+00:00", "2019-05-31 19:00:00+00:00", "2019-05-31 18:00:00+00:00", "2019-05-31 17:00:00+00:00", "2019-05-31 16:00:00+00:00", "2019-05-31 15:00:00+00:00", "2019-05-31 14:00:00+00:00", "2019-05-31 13:00:00+00:00", "2019-05-31 12:00:00+00:00", "2019-05-31 11:00:00+00:00", "2019-05-31 10:00:00+00:00", "2019-05-31 09:00:00+00:00", "2019-05-31 08:00:00+00:00", "2019-05-31 07:00:00+00:00", "2019-05-31 06:00:00+00:00", "2019-05-31 05:00:00+00:00", "2019-05-31 04:00:00+00:00", "2019-05-31 03:00:00+00:00", "2019-05-31 02:00:00+00:00", "2019-05-31 01:00:00+00:00", "2019-05-31 00:00:00+00:00", "2019-05-30 23:00:00+00:00", "2019-05-30 22:00:00+00:00", "2019-05-30 21:00:00+00:00", "2019-05-30 20:00:00+00:00", "2019-05-30 19:00:00+00:00", "2019-05-30 18:00:00+00:00", "2019-05-30 17:00:00+00:00", "2019-05-30 16:00:00+00:00", "2019-05-30 15:00:00+00:00", "2019-05-30 14:00:00+00:00", "2019-05-30 13:00:00+00:00", "2019-05-30 12:00:00+00:00", "2019-05-30 11:00:00+00:00", "2019-05-30 10:00:00+00:00", "2019-05-30 09:00:00+00:00", "2019-05-30 08:00:00+00:00", "2019-05-30 07:00:00+00:00", "2019-05-30 06:00:00+00:00", "2019-05-30 05:00:00+00:00", "2019-05-30 04:00:00+00:00", "2019-05-30 03:00:00+00:00", "2019-05-30 02:00:00+00:00", "2019-05-30 01:00:00+00:00", "2019-05-30 00:00:00+00:00", "2019-05-29 23:00:00+00:00", "2019-05-29 22:00:00+00:00", "2019-05-29 21:00:00+00:00", "2019-05-29 20:00:00+00:00", "2019-05-29 19:00:00+00:00", "2019-05-29 18:00:00+00:00", "2019-05-29 17:00:00+00:00", "2019-05-29 16:00:00+00:00", "2019-05-29 15:00:00+00:00", "2019-05-29 14:00:00+00:00", "2019-05-29 13:00:00+00:00", "2019-05-29 12:00:00+00:00", "2019-05-29 11:00:00+00:00", "2019-05-29 10:00:00+00:00", "2019-05-29 09:00:00+00:00", "2019-05-29 08:00:00+00:00", "2019-05-29 07:00:00+00:00", "2019-05-29 06:00:00+00:00", "2019-05-29 05:00:00+00:00", "2019-05-29 04:00:00+00:00", "2019-05-29 03:00:00+00:00", "2019-05-29 02:00:00+00:00", "2019-05-29 01:00:00+00:00", "2019-05-29 00:00:00+00:00", "2019-05-28 23:00:00+00:00", "2019-05-28 21:00:00+00:00", "2019-05-28 20:00:00+00:00", "2019-05-28 19:00:00+00:00", "2019-05-28 18:00:00+00:00", "2019-05-28 17:00:00+00:00", "2019-05-28 16:00:00+00:00", "2019-05-28 15:00:00+00:00", "2019-05-28 14:00:00+00:00", "2019-05-28 13:00:00+00:00", "2019-05-28 12:00:00+00:00", "2019-05-28 11:00:00+00:00", "2019-05-28 10:00:00+00:00", "2019-05-28 09:00:00+00:00", "2019-05-28 08:00:00+00:00", "2019-05-28 07:00:00+00:00", "2019-05-28 06:00:00+00:00", "2019-05-28 05:00:00+00:00", "2019-05-28 04:00:00+00:00", "2019-05-28 03:00:00+00:00", "2019-05-28 02:00:00+00:00", "2019-05-28 01:00:00+00:00", "2019-05-28 00:00:00+00:00", "2019-05-27 23:00:00+00:00", "2019-05-27 22:00:00+00:00", "2019-05-27 21:00:00+00:00", "2019-05-27 20:00:00+00:00", "2019-05-27 19:00:00+00:00", "2019-05-27 18:00:00+00:00", "2019-05-27 17:00:00+00:00", "2019-05-27 16:00:00+00:00", "2019-05-27 15:00:00+00:00", "2019-05-27 14:00:00+00:00", "2019-05-27 13:00:00+00:00", "2019-05-27 12:00:00+00:00", "2019-05-27 11:00:00+00:00", "2019-05-27 10:00:00+00:00", "2019-05-27 09:00:00+00:00", "2019-05-27 08:00:00+00:00", "2019-05-27 07:00:00+00:00", "2019-05-27 06:00:00+00:00", "2019-05-27 05:00:00+00:00", "2019-05-27 04:00:00+00:00", "2019-05-27 03:00:00+00:00", "2019-05-27 02:00:00+00:00", "2019-05-27 01:00:00+00:00", "2019-05-27 00:00:00+00:00", "2019-05-26 23:00:00+00:00", "2019-05-26 22:00:00+00:00", "2019-05-26 21:00:00+00:00", "2019-05-26 20:00:00+00:00", "2019-05-26 19:00:00+00:00", "2019-05-26 18:00:00+00:00", "2019-05-26 17:00:00+00:00", "2019-05-26 16:00:00+00:00", "2019-05-26 15:00:00+00:00", "2019-05-26 14:00:00+00:00", "2019-05-26 13:00:00+00:00", "2019-05-26 12:00:00+00:00", "2019-05-26 11:00:00+00:00", "2019-05-26 10:00:00+00:00", "2019-05-26 09:00:00+00:00", "2019-05-26 08:00:00+00:00", "2019-05-26 07:00:00+00:00", "2019-05-26 06:00:00+00:00", "2019-05-26 05:00:00+00:00", "2019-05-26 04:00:00+00:00", "2019-05-26 03:00:00+00:00", "2019-05-26 02:00:00+00:00", "2019-05-26 01:00:00+00:00", "2019-05-26 00:00:00+00:00", "2019-05-25 23:00:00+00:00", "2019-05-25 22:00:00+00:00", "2019-05-25 21:00:00+00:00", "2019-05-25 20:00:00+00:00", "2019-05-25 19:00:00+00:00", "2019-05-25 18:00:00+00:00", "2019-05-25 17:00:00+00:00", "2019-05-25 16:00:00+00:00", "2019-05-25 15:00:00+00:00", "2019-05-25 14:00:00+00:00", "2019-05-25 13:00:00+00:00", "2019-05-25 12:00:00+00:00", "2019-05-25 11:00:00+00:00", "2019-05-25 10:00:00+00:00", "2019-05-25 09:00:00+00:00", "2019-05-25 08:00:00+00:00", "2019-05-25 07:00:00+00:00", "2019-05-25 06:00:00+00:00", "2019-05-25 05:00:00+00:00", "2019-05-25 04:00:00+00:00", "2019-05-25 03:00:00+00:00", "2019-05-25 02:00:00+00:00", "2019-05-25 01:00:00+00:00", "2019-05-25 00:00:00+00:00", "2019-05-24 23:00:00+00:00", "2019-05-24 22:00:00+00:00", "2019-05-24 21:00:00+00:00", "2019-05-24 20:00:00+00:00", "2019-05-24 19:00:00+00:00", "2019-05-24 18:00:00+00:00", "2019-05-24 17:00:00+00:00", "2019-05-24 16:00:00+00:00", "2019-05-24 15:00:00+00:00", "2019-05-24 14:00:00+00:00", "2019-05-24 13:00:00+00:00", "2019-05-24 12:00:00+00:00", "2019-05-24 11:00:00+00:00", "2019-05-24 10:00:00+00:00", "2019-05-24 09:00:00+00:00", "2019-05-24 08:00:00+00:00", "2019-05-24 07:00:00+00:00", "2019-05-24 06:00:00+00:00", "2019-05-24 05:00:00+00:00", "2019-05-24 04:00:00+00:00", "2019-05-24 03:00:00+00:00", "2019-05-24 02:00:00+00:00", "2019-05-24 00:00:00+00:00", "2019-05-23 23:00:00+00:00", "2019-05-23 22:00:00+00:00", "2019-05-23 21:00:00+00:00", "2019-05-23 20:00:00+00:00", "2019-05-23 19:00:00+00:00", "2019-05-23 18:00:00+00:00", "2019-05-23 17:00:00+00:00", "2019-05-23 16:00:00+00:00", "2019-05-23 15:00:00+00:00", "2019-05-23 14:00:00+00:00", "2019-05-23 13:00:00+00:00", "2019-05-23 12:00:00+00:00", "2019-05-23 11:00:00+00:00", "2019-05-23 10:00:00+00:00", "2019-05-23 09:00:00+00:00", "2019-05-23 08:00:00+00:00", "2019-05-23 07:00:00+00:00", "2019-05-23 06:00:00+00:00", "2019-05-23 05:00:00+00:00", "2019-05-23 04:00:00+00:00", "2019-05-23 03:00:00+00:00", "2019-05-23 02:00:00+00:00", "2019-05-23 01:00:00+00:00", "2019-05-23 00:00:00+00:00", "2019-05-22 23:00:00+00:00", "2019-05-22 22:00:00+00:00", "2019-05-22 21:00:00+00:00", "2019-05-22 20:00:00+00:00", "2019-05-22 19:00:00+00:00", "2019-05-22 18:00:00+00:00", "2019-05-22 17:00:00+00:00", "2019-05-22 16:00:00+00:00", "2019-05-22 15:00:00+00:00", "2019-05-22 14:00:00+00:00", "2019-05-22 13:00:00+00:00", "2019-05-22 12:00:00+00:00", "2019-05-22 11:00:00+00:00", "2019-05-22 10:00:00+00:00", "2019-05-22 09:00:00+00:00", "2019-05-22 08:00:00+00:00", "2019-05-22 07:00:00+00:00", "2019-05-22 06:00:00+00:00", "2019-05-22 05:00:00+00:00", "2019-05-22 04:00:00+00:00", "2019-05-22 03:00:00+00:00", "2019-05-22 02:00:00+00:00", "2019-05-22 01:00:00+00:00", "2019-05-22 00:00:00+00:00", "2019-05-21 23:00:00+00:00", "2019-05-21 22:00:00+00:00", "2019-05-21 21:00:00+00:00", "2019-05-21 20:00:00+00:00", "2019-05-21 19:00:00+00:00", "2019-05-21 18:00:00+00:00", "2019-05-21 17:00:00+00:00", "2019-05-21 16:00:00+00:00", "2019-05-21 15:00:00+00:00", "2019-05-21 14:00:00+00:00", "2019-05-21 13:00:00+00:00", "2019-05-21 12:00:00+00:00", "2019-05-21 11:00:00+00:00", "2019-05-21 10:00:00+00:00", "2019-05-21 09:00:00+00:00", "2019-05-21 08:00:00+00:00", "2019-05-21 07:00:00+00:00", "2019-05-21 06:00:00+00:00", "2019-05-21 05:00:00+00:00", "2019-05-21 04:00:00+00:00", "2019-05-21 03:00:00+00:00", "2019-05-21 02:00:00+00:00", "2019-05-21 01:00:00+00:00", "2019-05-21 00:00:00+00:00", "2019-05-20 23:00:00+00:00", "2019-05-20 22:00:00+00:00", "2019-05-20 21:00:00+00:00", "2019-05-20 20:00:00+00:00", "2019-05-20 19:00:00+00:00", "2019-05-20 18:00:00+00:00", "2019-05-20 17:00:00+00:00", "2019-05-20 16:00:00+00:00", "2019-05-20 15:00:00+00:00", "2019-05-20 14:00:00+00:00", "2019-05-20 13:00:00+00:00", "2019-05-20 12:00:00+00:00", "2019-05-20 11:00:00+00:00", "2019-05-20 10:00:00+00:00", "2019-05-20 09:00:00+00:00", "2019-05-20 08:00:00+00:00", "2019-05-20 07:00:00+00:00", "2019-05-20 06:00:00+00:00", "2019-05-20 05:00:00+00:00", "2019-05-20 04:00:00+00:00", "2019-05-20 03:00:00+00:00", "2019-05-20 02:00:00+00:00", "2019-05-20 01:00:00+00:00", "2019-05-20 00:00:00+00:00", "2019-05-19 23:00:00+00:00", "2019-05-19 22:00:00+00:00", "2019-05-19 21:00:00+00:00", "2019-05-19 20:00:00+00:00", "2019-05-19 19:00:00+00:00", "2019-05-19 18:00:00+00:00", "2019-05-19 17:00:00+00:00", "2019-05-19 16:00:00+00:00", "2019-05-19 15:00:00+00:00", "2019-05-19 14:00:00+00:00", "2019-05-19 13:00:00+00:00", "2019-05-19 12:00:00+00:00", "2019-05-19 11:00:00+00:00", "2019-05-19 10:00:00+00:00", "2019-05-19 09:00:00+00:00", "2019-05-19 08:00:00+00:00", "2019-05-19 07:00:00+00:00", "2019-05-19 06:00:00+00:00", "2019-05-19 05:00:00+00:00", "2019-05-19 04:00:00+00:00", "2019-05-19 03:00:00+00:00", "2019-05-19 02:00:00+00:00", "2019-05-19 01:00:00+00:00", "2019-05-19 00:00:00+00:00", "2019-05-18 23:00:00+00:00", "2019-05-18 22:00:00+00:00", "2019-05-18 21:00:00+00:00", "2019-05-18 20:00:00+00:00", "2019-05-18 19:00:00+00:00", "2019-05-18 18:00:00+00:00", "2019-05-18 17:00:00+00:00", "2019-05-18 16:00:00+00:00", "2019-05-18 15:00:00+00:00", "2019-05-18 14:00:00+00:00", "2019-05-18 13:00:00+00:00", "2019-05-18 12:00:00+00:00", "2019-05-18 11:00:00+00:00", "2019-05-18 10:00:00+00:00", "2019-05-18 09:00:00+00:00", "2019-05-18 08:00:00+00:00", "2019-05-18 07:00:00+00:00", "2019-05-18 06:00:00+00:00", "2019-05-18 05:00:00+00:00", "2019-05-18 04:00:00+00:00", "2019-05-18 03:00:00+00:00", "2019-05-18 02:00:00+00:00", "2019-05-18 01:00:00+00:00", "2019-05-18 00:00:00+00:00", "2019-05-17 23:00:00+00:00", "2019-05-17 22:00:00+00:00", "2019-05-17 21:00:00+00:00", "2019-05-17 20:00:00+00:00", "2019-05-17 19:00:00+00:00", "2019-05-17 18:00:00+00:00", "2019-05-17 17:00:00+00:00", "2019-05-17 16:00:00+00:00", "2019-05-17 15:00:00+00:00", "2019-05-17 14:00:00+00:00", "2019-05-17 13:00:00+00:00", "2019-05-17 12:00:00+00:00", "2019-05-17 11:00:00+00:00", "2019-05-17 10:00:00+00:00", "2019-05-17 09:00:00+00:00", "2019-05-17 08:00:00+00:00", "2019-05-17 07:00:00+00:00", "2019-05-17 06:00:00+00:00", "2019-05-17 05:00:00+00:00", "2019-05-17 04:00:00+00:00", "2019-05-17 03:00:00+00:00", "2019-05-17 02:00:00+00:00", "2019-05-17 01:00:00+00:00", "2019-05-17 00:00:00+00:00", "2019-05-16 23:00:00+00:00", "2019-05-16 22:00:00+00:00", "2019-05-16 21:00:00+00:00", "2019-05-16 20:00:00+00:00", "2019-05-16 19:00:00+00:00", "2019-05-16 18:00:00+00:00", "2019-05-16 17:00:00+00:00", "2019-05-16 16:00:00+00:00", "2019-05-16 15:00:00+00:00", "2019-05-16 14:00:00+00:00", "2019-05-16 13:00:00+00:00", "2019-05-16 12:00:00+00:00", "2019-05-16 11:00:00+00:00", "2019-05-16 10:00:00+00:00", "2019-05-16 09:00:00+00:00", "2019-05-16 08:00:00+00:00", "2019-05-16 07:00:00+00:00", "2019-05-16 06:00:00+00:00", "2019-05-16 05:00:00+00:00", "2019-05-16 04:00:00+00:00", "2019-05-16 03:00:00+00:00", "2019-05-16 02:00:00+00:00", "2019-05-16 01:00:00+00:00", "2019-05-16 00:00:00+00:00", "2019-05-15 23:00:00+00:00", "2019-05-15 22:00:00+00:00", "2019-05-15 21:00:00+00:00", "2019-05-15 20:00:00+00:00", "2019-05-15 19:00:00+00:00", "2019-05-15 18:00:00+00:00", "2019-05-15 17:00:00+00:00", "2019-05-15 16:00:00+00:00", "2019-05-15 15:00:00+00:00", "2019-05-15 14:00:00+00:00", "2019-05-15 13:00:00+00:00", "2019-05-15 12:00:00+00:00", "2019-05-15 11:00:00+00:00", "2019-05-15 10:00:00+00:00", "2019-05-15 09:00:00+00:00", "2019-05-15 08:00:00+00:00", "2019-05-15 07:00:00+00:00", "2019-05-15 06:00:00+00:00", "2019-05-15 05:00:00+00:00", "2019-05-15 04:00:00+00:00", "2019-05-15 03:00:00+00:00", "2019-05-15 02:00:00+00:00", "2019-05-15 00:00:00+00:00", "2019-05-14 23:00:00+00:00", "2019-05-14 22:00:00+00:00", "2019-05-14 21:00:00+00:00", "2019-05-14 20:00:00+00:00", "2019-05-14 19:00:00+00:00", "2019-05-14 18:00:00+00:00", "2019-05-14 17:00:00+00:00", "2019-05-14 16:00:00+00:00", "2019-05-14 15:00:00+00:00", "2019-05-14 14:00:00+00:00", "2019-05-14 13:00:00+00:00", "2019-05-14 12:00:00+00:00", "2019-05-14 11:00:00+00:00", "2019-05-14 10:00:00+00:00", "2019-05-14 09:00:00+00:00", "2019-05-14 08:00:00+00:00", "2019-05-14 07:00:00+00:00", "2019-05-14 06:00:00+00:00", "2019-05-14 05:00:00+00:00", "2019-05-14 04:00:00+00:00", "2019-05-14 03:00:00+00:00", "2019-05-14 02:00:00+00:00", "2019-05-14 01:00:00+00:00", "2019-05-14 00:00:00+00:00", "2019-05-13 23:00:00+00:00", "2019-05-13 22:00:00+00:00", "2019-05-13 21:00:00+00:00", "2019-05-13 20:00:00+00:00", "2019-05-13 19:00:00+00:00", "2019-05-13 18:00:00+00:00", "2019-05-13 17:00:00+00:00", "2019-05-13 16:00:00+00:00", "2019-05-13 15:00:00+00:00", "2019-05-13 14:00:00+00:00", "2019-05-13 13:00:00+00:00", "2019-05-13 12:00:00+00:00", "2019-05-13 11:00:00+00:00", "2019-05-13 10:00:00+00:00", "2019-05-13 09:00:00+00:00", "2019-05-13 08:00:00+00:00", "2019-05-13 07:00:00+00:00", "2019-05-13 06:00:00+00:00", "2019-05-13 05:00:00+00:00", "2019-05-13 04:00:00+00:00", "2019-05-13 03:00:00+00:00", "2019-05-13 02:00:00+00:00", "2019-05-13 01:00:00+00:00", "2019-05-13 00:00:00+00:00", "2019-05-12 23:00:00+00:00", "2019-05-12 22:00:00+00:00", "2019-05-12 21:00:00+00:00", "2019-05-12 20:00:00+00:00", "2019-05-12 19:00:00+00:00", "2019-05-12 18:00:00+00:00", "2019-05-12 17:00:00+00:00", "2019-05-12 16:00:00+00:00", "2019-05-12 15:00:00+00:00", "2019-05-12 14:00:00+00:00", "2019-05-12 13:00:00+00:00", "2019-05-12 12:00:00+00:00", "2019-05-12 11:00:00+00:00", "2019-05-12 10:00:00+00:00", "2019-05-12 09:00:00+00:00", "2019-05-12 08:00:00+00:00", "2019-05-12 07:00:00+00:00", "2019-05-12 06:00:00+00:00", "2019-05-12 05:00:00+00:00", "2019-05-12 04:00:00+00:00", "2019-05-12 03:00:00+00:00", "2019-05-12 02:00:00+00:00", "2019-05-12 01:00:00+00:00", "2019-05-12 00:00:00+00:00", "2019-05-11 23:00:00+00:00", "2019-05-11 22:00:00+00:00", "2019-05-11 21:00:00+00:00", "2019-05-11 20:00:00+00:00", "2019-05-11 19:00:00+00:00", "2019-05-11 18:00:00+00:00", "2019-05-11 17:00:00+00:00", "2019-05-11 16:00:00+00:00", "2019-05-11 15:00:00+00:00", "2019-05-11 09:00:00+00:00", "2019-05-11 08:00:00+00:00", "2019-05-11 07:00:00+00:00", "2019-05-11 06:00:00+00:00", "2019-05-11 05:00:00+00:00", "2019-05-11 04:00:00+00:00", "2019-05-11 03:00:00+00:00", "2019-05-11 02:00:00+00:00", "2019-05-11 01:00:00+00:00", "2019-05-11 00:00:00+00:00", "2019-05-10 23:00:00+00:00", "2019-05-10 22:00:00+00:00", "2019-05-10 21:00:00+00:00", "2019-05-10 20:00:00+00:00", "2019-05-10 19:00:00+00:00", "2019-05-10 18:00:00+00:00", "2019-05-10 17:00:00+00:00", "2019-05-10 16:00:00+00:00", "2019-05-10 15:00:00+00:00", "2019-05-10 14:00:00+00:00", "2019-05-10 13:00:00+00:00", "2019-05-10 12:00:00+00:00", "2019-05-10 11:00:00+00:00", "2019-05-10 10:00:00+00:00", "2019-05-10 09:00:00+00:00", "2019-05-10 08:00:00+00:00", "2019-05-10 07:00:00+00:00", "2019-05-10 06:00:00+00:00", "2019-05-10 05:00:00+00:00", "2019-05-10 04:00:00+00:00", "2019-05-10 03:00:00+00:00", "2019-05-10 02:00:00+00:00", "2019-05-10 01:00:00+00:00", "2019-05-10 00:00:00+00:00", "2019-05-09 23:00:00+00:00", "2019-05-09 22:00:00+00:00", "2019-05-09 21:00:00+00:00", "2019-05-09 20:00:00+00:00", "2019-05-09 19:00:00+00:00", "2019-05-09 18:00:00+00:00", "2019-05-09 17:00:00+00:00", "2019-05-09 16:00:00+00:00", "2019-05-09 15:00:00+00:00", "2019-05-09 14:00:00+00:00", "2019-05-09 13:00:00+00:00", "2019-05-09 12:00:00+00:00", "2019-05-09 11:00:00+00:00", "2019-05-09 10:00:00+00:00", "2019-05-09 09:00:00+00:00", "2019-05-09 08:00:00+00:00", "2019-05-09 07:00:00+00:00", "2019-05-09 06:00:00+00:00", "2019-05-09 05:00:00+00:00", "2019-05-09 04:00:00+00:00", "2019-05-09 03:00:00+00:00", "2019-05-09 02:00:00+00:00", "2019-05-09 00:00:00+00:00", "2019-05-08 23:00:00+00:00", "2019-05-08 21:00:00+00:00", "2019-05-08 20:00:00+00:00", "2019-05-08 19:00:00+00:00", "2019-05-08 18:00:00+00:00", "2019-05-08 17:00:00+00:00", "2019-05-08 16:00:00+00:00", "2019-05-08 15:00:00+00:00", "2019-05-08 14:00:00+00:00", "2019-05-08 13:00:00+00:00", "2019-05-08 12:00:00+00:00", "2019-05-08 11:00:00+00:00", "2019-05-08 10:00:00+00:00", "2019-05-08 09:00:00+00:00", "2019-05-08 08:00:00+00:00", "2019-05-08 07:00:00+00:00", "2019-05-08 06:00:00+00:00", "2019-05-08 05:00:00+00:00", "2019-05-08 04:00:00+00:00", "2019-05-08 03:00:00+00:00", "2019-05-08 02:00:00+00:00", "2019-05-08 01:00:00+00:00", "2019-05-08 00:00:00+00:00", "2019-05-07 23:00:00+00:00", "2019-05-07 21:00:00+00:00", "2019-05-07 20:00:00+00:00", "2019-05-07 19:00:00+00:00", "2019-05-07 18:00:00+00:00", "2019-05-07 17:00:00+00:00", "2019-05-07 16:00:00+00:00", "2019-05-07 15:00:00+00:00", "2019-05-07 14:00:00+00:00", "2019-05-07 13:00:00+00:00", "2019-05-07 12:00:00+00:00", "2019-05-07 11:00:00+00:00", "2019-05-07 10:00:00+00:00", "2019-05-07 09:00:00+00:00", "2019-05-07 08:00:00+00:00", "2019-05-07 07:00:00+00:00", "2019-05-07 06:00:00+00:00", "2019-05-07 04:00:00+00:00", "2019-05-07 03:00:00+00:00", "2019-05-07 02:00:00+00:00", "2019-05-07 01:00:00+00:00", "2019-05-06 23:00:00+00:00", "2019-05-06 22:00:00+00:00", "2019-05-06 21:00:00+00:00", "2019-05-06 20:00:00+00:00", "2019-05-06 19:00:00+00:00", "2019-05-06 18:00:00+00:00", "2019-05-06 17:00:00+00:00", "2019-05-06 16:00:00+00:00", "2019-05-06 15:00:00+00:00", "2019-05-06 14:00:00+00:00", "2019-05-06 13:00:00+00:00", "2019-05-06 12:00:00+00:00", "2019-05-06 11:00:00+00:00", "2019-05-06 10:00:00+00:00", "2019-05-06 09:00:00+00:00", "2019-05-06 08:00:00+00:00", "2019-05-06 07:00:00+00:00", "2019-05-06 06:00:00+00:00", "2019-05-06 05:00:00+00:00", "2019-05-06 04:00:00+00:00", "2019-05-06 03:00:00+00:00", "2019-05-06 02:00:00+00:00", "2019-05-06 01:00:00+00:00", "2019-05-06 00:00:00+00:00", "2019-05-05 23:00:00+00:00", "2019-05-05 22:00:00+00:00", "2019-05-05 21:00:00+00:00", "2019-05-05 20:00:00+00:00", "2019-05-05 19:00:00+00:00", "2019-05-05 18:00:00+00:00", "2019-05-05 17:00:00+00:00", "2019-05-05 16:00:00+00:00", "2019-05-05 15:00:00+00:00", "2019-05-05 14:00:00+00:00", "2019-05-05 13:00:00+00:00", "2019-05-05 12:00:00+00:00", "2019-05-05 11:00:00+00:00", "2019-05-05 10:00:00+00:00", "2019-05-05 09:00:00+00:00", "2019-05-05 08:00:00+00:00", "2019-05-05 07:00:00+00:00", "2019-05-05 06:00:00+00:00", "2019-05-05 05:00:00+00:00", "2019-05-05 04:00:00+00:00", "2019-05-05 03:00:00+00:00", "2019-05-05 02:00:00+00:00", "2019-05-05 01:00:00+00:00", "2019-05-05 00:00:00+00:00", "2019-05-04 23:00:00+00:00", "2019-05-04 22:00:00+00:00", "2019-05-04 21:00:00+00:00", "2019-05-04 20:00:00+00:00", "2019-05-04 19:00:00+00:00", "2019-05-04 18:00:00+00:00", "2019-05-04 17:00:00+00:00", "2019-05-04 16:00:00+00:00", "2019-05-04 15:00:00+00:00", "2019-05-04 14:00:00+00:00", "2019-05-04 13:00:00+00:00", "2019-05-04 12:00:00+00:00", "2019-05-04 11:00:00+00:00", "2019-05-04 10:00:00+00:00", "2019-05-04 09:00:00+00:00", "2019-05-04 08:00:00+00:00", "2019-05-04 07:00:00+00:00", "2019-05-04 06:00:00+00:00", "2019-05-04 05:00:00+00:00", "2019-05-04 04:00:00+00:00", "2019-05-04 03:00:00+00:00", "2019-05-04 02:00:00+00:00", "2019-05-04 01:00:00+00:00", "2019-05-04 00:00:00+00:00", "2019-05-03 23:00:00+00:00", "2019-05-03 22:00:00+00:00", "2019-05-03 21:00:00+00:00", "2019-05-03 20:00:00+00:00", "2019-05-03 19:00:00+00:00", "2019-05-03 18:00:00+00:00", "2019-05-03 17:00:00+00:00", "2019-05-03 16:00:00+00:00", "2019-05-03 15:00:00+00:00", "2019-05-03 14:00:00+00:00", "2019-05-03 13:00:00+00:00", "2019-05-03 12:00:00+00:00", "2019-05-03 11:00:00+00:00", "2019-05-03 10:00:00+00:00", "2019-05-03 09:00:00+00:00", "2019-05-03 08:00:00+00:00", "2019-05-03 07:00:00+00:00", "2019-05-03 06:00:00+00:00", "2019-05-03 05:00:00+00:00", "2019-05-03 04:00:00+00:00", "2019-05-03 03:00:00+00:00", "2019-05-03 02:00:00+00:00", "2019-05-03 01:00:00+00:00", "2019-05-03 00:00:00+00:00", "2019-05-02 23:00:00+00:00", "2019-05-02 22:00:00+00:00", "2019-05-02 21:00:00+00:00", "2019-05-02 20:00:00+00:00", "2019-05-02 19:00:00+00:00", "2019-05-02 18:00:00+00:00", "2019-05-02 17:00:00+00:00", "2019-05-02 16:00:00+00:00", "2019-05-02 15:00:00+00:00", "2019-05-02 14:00:00+00:00", "2019-05-02 13:00:00+00:00", "2019-05-02 12:00:00+00:00", "2019-05-02 11:00:00+00:00", "2019-05-02 10:00:00+00:00", "2019-05-02 09:00:00+00:00", "2019-05-02 08:00:00+00:00", "2019-05-02 07:00:00+00:00", "2019-05-02 06:00:00+00:00", "2019-05-02 05:00:00+00:00", "2019-05-02 04:00:00+00:00", "2019-05-02 03:00:00+00:00", "2019-05-02 02:00:00+00:00", "2019-05-02 01:00:00+00:00", "2019-05-02 00:00:00+00:00", "2019-05-01 23:00:00+00:00", "2019-05-01 22:00:00+00:00", "2019-05-01 21:00:00+00:00", "2019-05-01 20:00:00+00:00", "2019-05-01 19:00:00+00:00", "2019-05-01 18:00:00+00:00", "2019-05-01 17:00:00+00:00", "2019-05-01 16:00:00+00:00", "2019-05-01 15:00:00+00:00", "2019-05-01 14:00:00+00:00", "2019-05-01 13:00:00+00:00", "2019-05-01 12:00:00+00:00", "2019-05-01 11:00:00+00:00", "2019-05-01 10:00:00+00:00", "2019-05-01 09:00:00+00:00", "2019-05-01 08:00:00+00:00", "2019-05-01 07:00:00+00:00", "2019-05-01 06:00:00+00:00", "2019-05-01 05:00:00+00:00", "2019-05-01 04:00:00+00:00", "2019-05-01 03:00:00+00:00", "2019-05-01 00:00:00+00:00", "2019-04-30 23:00:00+00:00", "2019-04-30 22:00:00+00:00", "2019-04-30 21:00:00+00:00", "2019-04-30 20:00:00+00:00", "2019-04-30 19:00:00+00:00", "2019-04-30 18:00:00+00:00", "2019-04-30 17:00:00+00:00", "2019-04-30 16:00:00+00:00", "2019-04-30 15:00:00+00:00", "2019-04-30 14:00:00+00:00", "2019-04-30 13:00:00+00:00", "2019-04-30 12:00:00+00:00", "2019-04-30 11:00:00+00:00", "2019-04-30 10:00:00+00:00", "2019-04-30 09:00:00+00:00", "2019-04-30 08:00:00+00:00", "2019-04-30 07:00:00+00:00", "2019-04-30 06:00:00+00:00", "2019-04-30 05:00:00+00:00", "2019-04-30 04:00:00+00:00", "2019-04-30 03:00:00+00:00", "2019-04-30 02:00:00+00:00", "2019-04-30 01:00:00+00:00", "2019-04-30 00:00:00+00:00", "2019-04-29 23:00:00+00:00", "2019-04-29 22:00:00+00:00", "2019-04-29 21:00:00+00:00", "2019-04-29 20:00:00+00:00", "2019-04-29 19:00:00+00:00", "2019-04-29 18:00:00+00:00", "2019-04-29 17:00:00+00:00", "2019-04-29 16:00:00+00:00", "2019-04-29 15:00:00+00:00", "2019-04-29 14:00:00+00:00", "2019-04-29 13:00:00+00:00", "2019-04-29 12:00:00+00:00", "2019-04-29 11:00:00+00:00", "2019-04-29 10:00:00+00:00", "2019-04-29 09:00:00+00:00", "2019-04-29 08:00:00+00:00", "2019-04-29 07:00:00+00:00", "2019-04-29 06:00:00+00:00", "2019-04-29 05:00:00+00:00", "2019-04-29 04:00:00+00:00", "2019-04-29 03:00:00+00:00", "2019-04-29 02:00:00+00:00", "2019-04-29 01:00:00+00:00", "2019-04-29 00:00:00+00:00", "2019-04-28 23:00:00+00:00", "2019-04-28 22:00:00+00:00", "2019-04-28 21:00:00+00:00", "2019-04-28 20:00:00+00:00", "2019-04-28 19:00:00+00:00", "2019-04-28 18:00:00+00:00", "2019-04-28 17:00:00+00:00", "2019-04-28 16:00:00+00:00", "2019-04-28 15:00:00+00:00", "2019-04-28 14:00:00+00:00", "2019-04-28 13:00:00+00:00", "2019-04-28 12:00:00+00:00", "2019-04-28 11:00:00+00:00", "2019-04-28 10:00:00+00:00", "2019-04-28 09:00:00+00:00", "2019-04-27 13:00:00+00:00", "2019-04-27 12:00:00+00:00", "2019-04-27 11:00:00+00:00", "2019-04-27 10:00:00+00:00", "2019-04-27 09:00:00+00:00", "2019-04-27 08:00:00+00:00", "2019-04-27 07:00:00+00:00", "2019-04-27 06:00:00+00:00", "2019-04-27 05:00:00+00:00", "2019-04-27 04:00:00+00:00", "2019-04-27 03:00:00+00:00", "2019-04-27 02:00:00+00:00", "2019-04-27 00:00:00+00:00", "2019-04-26 23:00:00+00:00", "2019-04-26 22:00:00+00:00", "2019-04-26 21:00:00+00:00", "2019-04-26 20:00:00+00:00", "2019-04-26 19:00:00+00:00", "2019-04-26 18:00:00+00:00", "2019-04-26 17:00:00+00:00", "2019-04-26 16:00:00+00:00", "2019-04-26 15:00:00+00:00", "2019-04-26 14:00:00+00:00", "2019-04-26 13:00:00+00:00", "2019-04-26 12:00:00+00:00", "2019-04-26 11:00:00+00:00", "2019-04-26 10:00:00+00:00", "2019-04-26 09:00:00+00:00", "2019-04-26 08:00:00+00:00", "2019-04-26 07:00:00+00:00", "2019-04-26 06:00:00+00:00", "2019-04-26 05:00:00+00:00", "2019-04-26 04:00:00+00:00", "2019-04-26 03:00:00+00:00", "2019-04-26 02:00:00+00:00", "2019-04-26 01:00:00+00:00", "2019-04-26 00:00:00+00:00", "2019-04-25 23:00:00+00:00", "2019-04-25 22:00:00+00:00", "2019-04-25 21:00:00+00:00", "2019-04-25 20:00:00+00:00", "2019-04-25 19:00:00+00:00", "2019-04-25 18:00:00+00:00", "2019-04-25 17:00:00+00:00", "2019-04-25 16:00:00+00:00", "2019-04-25 15:00:00+00:00", "2019-04-25 14:00:00+00:00", "2019-04-25 13:00:00+00:00", "2019-04-25 12:00:00+00:00", "2019-04-25 11:00:00+00:00", "2019-04-25 10:00:00+00:00", "2019-04-25 09:00:00+00:00", "2019-04-25 08:00:00+00:00", "2019-04-25 07:00:00+00:00", "2019-04-25 06:00:00+00:00", "2019-04-25 05:00:00+00:00", "2019-04-25 04:00:00+00:00", "2019-04-25 03:00:00+00:00", "2019-04-25 02:00:00+00:00", "2019-04-25 00:00:00+00:00", "2019-04-24 23:00:00+00:00", "2019-04-24 22:00:00+00:00", "2019-04-24 21:00:00+00:00", "2019-04-24 20:00:00+00:00", "2019-04-24 19:00:00+00:00", "2019-04-24 18:00:00+00:00", "2019-04-24 17:00:00+00:00", "2019-04-24 16:00:00+00:00", "2019-04-24 15:00:00+00:00", "2019-04-24 14:00:00+00:00", "2019-04-24 13:00:00+00:00", "2019-04-24 12:00:00+00:00", "2019-04-24 11:00:00+00:00", "2019-04-24 10:00:00+00:00", "2019-04-24 09:00:00+00:00", "2019-04-24 08:00:00+00:00", "2019-04-24 07:00:00+00:00", "2019-04-24 06:00:00+00:00", "2019-04-24 05:00:00+00:00", "2019-04-24 04:00:00+00:00", "2019-04-24 03:00:00+00:00", "2019-04-24 02:00:00+00:00", "2019-04-24 00:00:00+00:00", "2019-04-23 23:00:00+00:00", "2019-04-23 22:00:00+00:00", "2019-04-23 21:00:00+00:00", "2019-04-23 20:00:00+00:00", "2019-04-23 19:00:00+00:00", "2019-04-23 18:00:00+00:00", "2019-04-23 17:00:00+00:00", "2019-04-23 16:00:00+00:00", "2019-04-23 15:00:00+00:00", "2019-04-23 14:00:00+00:00", "2019-04-23 13:00:00+00:00", "2019-04-23 12:00:00+00:00", "2019-04-23 11:00:00+00:00", "2019-04-23 10:00:00+00:00", "2019-04-23 09:00:00+00:00", "2019-04-23 08:00:00+00:00", "2019-04-23 07:00:00+00:00", "2019-04-23 06:00:00+00:00", "2019-04-23 05:00:00+00:00", "2019-04-23 04:00:00+00:00", "2019-04-23 03:00:00+00:00", "2019-04-23 02:00:00+00:00", "2019-04-23 01:00:00+00:00", "2019-04-23 00:00:00+00:00", "2019-04-22 23:00:00+00:00", "2019-04-22 22:00:00+00:00", "2019-04-22 21:00:00+00:00", "2019-04-22 20:00:00+00:00", "2019-04-22 19:00:00+00:00", "2019-04-22 18:00:00+00:00", "2019-04-22 17:00:00+00:00", "2019-04-22 16:00:00+00:00", "2019-04-22 15:00:00+00:00", "2019-04-22 14:00:00+00:00", "2019-04-22 13:00:00+00:00", "2019-04-22 12:00:00+00:00", "2019-04-22 11:00:00+00:00", "2019-04-22 10:00:00+00:00", "2019-04-22 09:00:00+00:00", "2019-04-22 08:00:00+00:00", "2019-04-22 07:00:00+00:00", "2019-04-22 06:00:00+00:00", "2019-04-22 05:00:00+00:00", "2019-04-22 04:00:00+00:00", "2019-04-22 03:00:00+00:00", "2019-04-22 02:00:00+00:00", "2019-04-22 01:00:00+00:00", "2019-04-22 00:00:00+00:00", "2019-04-21 23:00:00+00:00", "2019-04-21 22:00:00+00:00", "2019-04-21 21:00:00+00:00", "2019-04-21 20:00:00+00:00", "2019-04-21 19:00:00+00:00", "2019-04-21 18:00:00+00:00", "2019-04-21 17:00:00+00:00", "2019-04-21 16:00:00+00:00", "2019-04-21 15:00:00+00:00", "2019-04-21 14:00:00+00:00", "2019-04-21 13:00:00+00:00", "2019-04-21 12:00:00+00:00", "2019-04-21 11:00:00+00:00", "2019-04-21 10:00:00+00:00", "2019-04-21 09:00:00+00:00", "2019-04-21 08:00:00+00:00", "2019-04-21 07:00:00+00:00", "2019-04-21 06:00:00+00:00", "2019-04-21 05:00:00+00:00", "2019-04-21 04:00:00+00:00", "2019-04-21 03:00:00+00:00", "2019-04-21 02:00:00+00:00", "2019-04-21 01:00:00+00:00", "2019-04-21 00:00:00+00:00", "2019-04-20 23:00:00+00:00", "2019-04-20 22:00:00+00:00", "2019-04-20 21:00:00+00:00", "2019-04-20 20:00:00+00:00", "2019-04-20 19:00:00+00:00", "2019-04-20 18:00:00+00:00", "2019-04-20 17:00:00+00:00", "2019-04-20 16:00:00+00:00", "2019-04-20 15:00:00+00:00", "2019-04-20 14:00:00+00:00", "2019-04-20 13:00:00+00:00", "2019-04-20 12:00:00+00:00", "2019-04-20 11:00:00+00:00", "2019-04-20 10:00:00+00:00", "2019-04-20 09:00:00+00:00", "2019-04-20 08:00:00+00:00", "2019-04-20 07:00:00+00:00", "2019-04-20 06:00:00+00:00", "2019-04-20 05:00:00+00:00", "2019-04-20 04:00:00+00:00", "2019-04-20 03:00:00+00:00", "2019-04-20 02:00:00+00:00", "2019-04-20 01:00:00+00:00", "2019-04-20 00:00:00+00:00", "2019-04-19 23:00:00+00:00", "2019-04-19 22:00:00+00:00", "2019-04-19 21:00:00+00:00", "2019-04-19 20:00:00+00:00", "2019-04-19 19:00:00+00:00", "2019-04-19 18:00:00+00:00", "2019-04-19 17:00:00+00:00", "2019-04-19 16:00:00+00:00", "2019-04-19 15:00:00+00:00", "2019-04-19 14:00:00+00:00", "2019-04-19 13:00:00+00:00", "2019-04-19 12:00:00+00:00", "2019-04-19 11:00:00+00:00", "2019-04-19 10:00:00+00:00", "2019-04-19 09:00:00+00:00", "2019-04-19 08:00:00+00:00", "2019-04-19 07:00:00+00:00", "2019-04-19 06:00:00+00:00", "2019-04-19 05:00:00+00:00", "2019-04-19 04:00:00+00:00", "2019-04-19 03:00:00+00:00", "2019-04-19 02:00:00+00:00", "2019-04-19 00:00:00+00:00", "2019-04-18 23:00:00+00:00", "2019-04-18 22:00:00+00:00", "2019-04-18 21:00:00+00:00", "2019-04-18 20:00:00+00:00", "2019-04-18 19:00:00+00:00", "2019-04-18 18:00:00+00:00", "2019-04-18 17:00:00+00:00", "2019-04-18 16:00:00+00:00", "2019-04-18 15:00:00+00:00", "2019-04-18 14:00:00+00:00", "2019-04-18 13:00:00+00:00", "2019-04-18 12:00:00+00:00", "2019-04-18 11:00:00+00:00", "2019-04-18 10:00:00+00:00", "2019-04-18 09:00:00+00:00", "2019-04-18 08:00:00+00:00", "2019-04-18 07:00:00+00:00", "2019-04-18 06:00:00+00:00", "2019-04-18 05:00:00+00:00", "2019-04-18 04:00:00+00:00", "2019-04-18 03:00:00+00:00", "2019-04-18 02:00:00+00:00", "2019-04-18 01:00:00+00:00", "2019-04-18 00:00:00+00:00", "2019-04-17 23:00:00+00:00", "2019-04-17 22:00:00+00:00", "2019-04-17 21:00:00+00:00", "2019-04-17 20:00:00+00:00", "2019-04-17 19:00:00+00:00", "2019-04-17 18:00:00+00:00", "2019-04-17 17:00:00+00:00", "2019-04-17 16:00:00+00:00", "2019-04-17 15:00:00+00:00", "2019-04-17 14:00:00+00:00", "2019-04-17 13:00:00+00:00", "2019-04-17 12:00:00+00:00", "2019-04-17 11:00:00+00:00", "2019-04-17 10:00:00+00:00", "2019-04-17 09:00:00+00:00", "2019-04-17 08:00:00+00:00", "2019-04-17 07:00:00+00:00", "2019-04-17 06:00:00+00:00", "2019-04-17 05:00:00+00:00", "2019-04-17 04:00:00+00:00", "2019-04-17 03:00:00+00:00", "2019-04-17 02:00:00+00:00", "2019-04-17 00:00:00+00:00", "2019-04-16 23:00:00+00:00", "2019-04-16 22:00:00+00:00", "2019-04-16 21:00:00+00:00", "2019-04-16 20:00:00+00:00", "2019-04-16 19:00:00+00:00", "2019-04-16 18:00:00+00:00", "2019-04-16 17:00:00+00:00", "2019-04-16 15:00:00+00:00", "2019-04-16 14:00:00+00:00", "2019-04-16 13:00:00+00:00", "2019-04-16 12:00:00+00:00", "2019-04-16 11:00:00+00:00", "2019-04-16 10:00:00+00:00", "2019-04-16 09:00:00+00:00", "2019-04-16 08:00:00+00:00", "2019-04-16 07:00:00+00:00", "2019-04-16 06:00:00+00:00", "2019-04-16 05:00:00+00:00", "2019-04-16 04:00:00+00:00", "2019-04-16 03:00:00+00:00", "2019-04-16 02:00:00+00:00", "2019-04-16 00:00:00+00:00", "2019-04-15 23:00:00+00:00", "2019-04-15 22:00:00+00:00", "2019-04-15 21:00:00+00:00", "2019-04-15 20:00:00+00:00", "2019-04-15 19:00:00+00:00", "2019-04-15 18:00:00+00:00", "2019-04-15 17:00:00+00:00", "2019-04-15 16:00:00+00:00", "2019-04-15 15:00:00+00:00", "2019-04-15 14:00:00+00:00", "2019-04-15 13:00:00+00:00", "2019-04-15 12:00:00+00:00", "2019-04-15 11:00:00+00:00", "2019-04-15 10:00:00+00:00", "2019-04-15 09:00:00+00:00", "2019-04-15 08:00:00+00:00", "2019-04-15 07:00:00+00:00", "2019-04-15 06:00:00+00:00", "2019-04-15 05:00:00+00:00", "2019-04-15 04:00:00+00:00", "2019-04-15 03:00:00+00:00", "2019-04-15 02:00:00+00:00", "2019-04-15 01:00:00+00:00", "2019-04-15 00:00:00+00:00", "2019-04-14 23:00:00+00:00", "2019-04-14 22:00:00+00:00", "2019-04-14 21:00:00+00:00", "2019-04-14 20:00:00+00:00", "2019-04-14 19:00:00+00:00", "2019-04-14 18:00:00+00:00", "2019-04-14 17:00:00+00:00", "2019-04-14 16:00:00+00:00", "2019-04-14 15:00:00+00:00", "2019-04-14 14:00:00+00:00", "2019-04-14 13:00:00+00:00", "2019-04-14 12:00:00+00:00", "2019-04-14 11:00:00+00:00", "2019-04-14 10:00:00+00:00", "2019-04-14 09:00:00+00:00", "2019-04-14 08:00:00+00:00", "2019-04-14 07:00:00+00:00", "2019-04-14 06:00:00+00:00", "2019-04-14 05:00:00+00:00", "2019-04-14 04:00:00+00:00", "2019-04-14 03:00:00+00:00", "2019-04-14 02:00:00+00:00", "2019-04-14 01:00:00+00:00", "2019-04-14 00:00:00+00:00", "2019-04-13 23:00:00+00:00", "2019-04-13 22:00:00+00:00", "2019-04-13 21:00:00+00:00", "2019-04-13 20:00:00+00:00", "2019-04-13 19:00:00+00:00", "2019-04-13 18:00:00+00:00", "2019-04-13 17:00:00+00:00", "2019-04-13 16:00:00+00:00", "2019-04-13 15:00:00+00:00", "2019-04-13 14:00:00+00:00", "2019-04-13 13:00:00+00:00", "2019-04-13 12:00:00+00:00", "2019-04-13 11:00:00+00:00", "2019-04-13 10:00:00+00:00", "2019-04-13 09:00:00+00:00", "2019-04-13 08:00:00+00:00", "2019-04-13 07:00:00+00:00", "2019-04-13 06:00:00+00:00", "2019-04-13 05:00:00+00:00", "2019-04-13 04:00:00+00:00", "2019-04-13 03:00:00+00:00", "2019-04-13 02:00:00+00:00", "2019-04-13 01:00:00+00:00", "2019-04-13 00:00:00+00:00", "2019-04-12 23:00:00+00:00", "2019-04-12 22:00:00+00:00", "2019-04-12 21:00:00+00:00", "2019-04-12 20:00:00+00:00", "2019-04-12 19:00:00+00:00", "2019-04-12 18:00:00+00:00", "2019-04-12 17:00:00+00:00", "2019-04-12 16:00:00+00:00", "2019-04-12 15:00:00+00:00", "2019-04-12 14:00:00+00:00", "2019-04-12 13:00:00+00:00", "2019-04-12 12:00:00+00:00", "2019-04-12 11:00:00+00:00", "2019-04-12 10:00:00+00:00", "2019-04-12 09:00:00+00:00", "2019-04-12 08:00:00+00:00", "2019-04-12 07:00:00+00:00", "2019-04-12 06:00:00+00:00", "2019-04-12 05:00:00+00:00", "2019-04-12 04:00:00+00:00", "2019-04-12 03:00:00+00:00", "2019-04-12 00:00:00+00:00", "2019-04-11 23:00:00+00:00", "2019-04-11 22:00:00+00:00", "2019-04-11 21:00:00+00:00", "2019-04-11 20:00:00+00:00", "2019-04-11 19:00:00+00:00", "2019-04-11 18:00:00+00:00", "2019-04-11 17:00:00+00:00", "2019-04-11 16:00:00+00:00", "2019-04-11 15:00:00+00:00", "2019-04-11 14:00:00+00:00", "2019-04-11 13:00:00+00:00", "2019-04-11 12:00:00+00:00", "2019-04-11 11:00:00+00:00", "2019-04-11 10:00:00+00:00", "2019-04-11 09:00:00+00:00", "2019-04-11 08:00:00+00:00", "2019-04-11 07:00:00+00:00", "2019-04-11 06:00:00+00:00", "2019-04-11 05:00:00+00:00", "2019-04-11 04:00:00+00:00", "2019-04-11 03:00:00+00:00", "2019-04-11 02:00:00+00:00", "2019-04-11 00:00:00+00:00", "2019-04-10 23:00:00+00:00", "2019-04-10 22:00:00+00:00", "2019-04-10 21:00:00+00:00", "2019-04-10 20:00:00+00:00", "2019-04-10 19:00:00+00:00", "2019-04-10 18:00:00+00:00", "2019-04-10 17:00:00+00:00", "2019-04-10 16:00:00+00:00", "2019-04-10 15:00:00+00:00", "2019-04-10 14:00:00+00:00", "2019-04-10 13:00:00+00:00", "2019-04-10 12:00:00+00:00", "2019-04-10 11:00:00+00:00", "2019-04-10 10:00:00+00:00", "2019-04-10 09:00:00+00:00", "2019-04-10 08:00:00+00:00", "2019-04-10 07:00:00+00:00", "2019-04-10 06:00:00+00:00", "2019-04-10 05:00:00+00:00", "2019-04-10 04:00:00+00:00", "2019-04-10 03:00:00+00:00", "2019-04-10 02:00:00+00:00", "2019-04-10 01:00:00+00:00", "2019-04-10 00:00:00+00:00", "2019-04-09 23:00:00+00:00", "2019-04-09 22:00:00+00:00", "2019-04-09 21:00:00+00:00", "2019-04-09 20:00:00+00:00", "2019-04-09 19:00:00+00:00", "2019-04-09 18:00:00+00:00", "2019-04-09 17:00:00+00:00", "2019-04-09 16:00:00+00:00", "2019-04-09 15:00:00+00:00", "2019-04-09 14:00:00+00:00", "2019-04-09 13:00:00+00:00", "2019-04-09 12:00:00+00:00", "2019-04-09 11:00:00+00:00", "2019-04-09 10:00:00+00:00", "2019-04-09 09:00:00+00:00", "2019-04-09 08:00:00+00:00", "2019-04-09 07:00:00+00:00", "2019-04-09 06:00:00+00:00", "2019-04-09 05:00:00+00:00", "2019-04-09 04:00:00+00:00", "2019-04-09 03:00:00+00:00", "2019-04-09 02:00:00+00:00",
// 					},
// 				},
// 				[][]interface{}{
// 					{"pm25", "BETR801"},
// 					{"pm25", "London Westminster"},
// 					{"no2", "FR04014"},
// 					{"no2", "BETR801"},
// 					{"no2", "London Westminster"},
// 				},
// 			},
// 		},
// 	}
// 	for _, test := range groupByTests {
// 		output, err := test.arg1.GroupBy(test.arg2...)
// 		if !cmp.Equal(output, test.expected, cmp.AllowUnexported(DataFrame{}, Series{}, IndexData{}, Index{}, GroupBy{}), cmpopts.EquateNaNs()) || err != nil {
// 			t.Fatalf("expected:\n%v,\ngot:\n%v,\nerror:\n%v\n", test.expected, output, err)
// 		}
// 	}
// }
