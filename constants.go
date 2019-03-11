package main

var excuses = []string{
	"excuse me waiter, there's a 404 in my link",
	"404 outlook not so good",
	"do you like pina coladas, and getting FILE NOT FOUND?",
	"always, I know, you'll be, 404",
	"it's a bittersweet symphony, this 404",
	"hey now, you're a rockstar, get the show on, 404",
	"sorry girl, but you missed out, well tough luck that 404's now",
	"music hits me, so hard, makes me say, 404",
	"i'm blue, 404044, 404, 404, 404044 (4 4 4..)",
	"if you wanna be my lover, you gotta 404 my friends",
	"if I go 404 will you still call me superman?",
	"oops 4 0 4 again",
	"you better lose yourself in the music, the moment, you hold it, you better never 404",
	"i feel stupid, and contagious, 404 now, imitate us",
}

var fragments = []string{
	"abc",
	"act",
	"age",
	"air",
	"ale",
	"all",
	"and",
	"ant",
	"any",
	"ape",
	"are",
	"art",
	"ash",
	"ate",
	"bad",
	"bag",
	"bah",
	"ban",
	"bar",
	"bat",
	"bed",
	"bee",
	"beg",
	"bet",
	"bin",
	"bit",
	"boa",
	"bra",
	"bud",
	"bun",
	"bus",
	"buy",
	"cab",
	"cad",
	"cam",
	"can",
	"cap",
	"car",
	"cat",
	"ceo",
	"cod",
	"con",
	"cot",
	"cow",
	"cue",
	"cuz",
	"dab",
	"dad",
	"dam",
	"dan",
	"day",
	"deb",
	"dee",
	"del",
	"den",
	"der",
	"des",
	"dew",
	"die",
	"dis",
	"dna",
	"doc",
	"doe",
	"don",
	"dot",
	"dub",
	"dye",
	"ear",
	"eat",
	"eel",
	"ego",
	"elm",
	"end",
	"eon",
	"era",
	"ere",
	"err",
	"ese",
	"esp",
	"eta",
	"eve",
	"eww",
	"fad",
	"fan",
	"far",
	"fat",
	"fbi",
	"fed",
	"fee",
	"fen",
	"few",
	"fib",
	"fir",
	"fix",
	"flu",
	"foe",
	"foo",
	"for",
	"fox",
	"fro",
	"fry",
	"fur",
	"gab",
	"gad",
	"gag",
	"gal",
	"gap",
	"gas",
	"gee",
	"gel",
	"gem",
	"geo",
	"get",
	"gin",
	"god",
	"goo",
	"got",
	"got",
	"grr",
	"had",
	"hag",
	"hah",
	"hal",
	"ham",
	"has",
	"hat",
	"hay",
	"hem",
	"hen",
	"hep",
	"her",
	"hey",
	"hip",
	"his",
	"hit",
	"hit",
	"hoe",
	"hop",
	"hot",
	"hue",
	"ian",
	"ice",
	"ike",
	"ikr",
	"ill",
	"imp",
	"ink",
	"inn",
	"ion",
	"ire",
	"its",
	"ive",
	"jag",
	"jam",
	"jan",
	"jar",
	"jaw",
	"jay",
	"jet",
	"jib",
	"jig",
	"job",
	"joe",
	"jog",
	"jon",
	"jot",
	"joy",
	"joy",
	"kal",
	"kay",
	"kea",
	"keg",
	"ken",
	"key",
	"key",
	"kgb",
	"kid",
	"kim",
	"kin",
	"kip",
	"kit",
	"lab",
	"lad",
	"lam",
	"lap",
	"las",
	"law",
	"led",
	"lee",
	"leg",
	"len",
	"leo",
	"let",
	"lid",
	"lie",
	"lol",
	"los",
	"lot",
	"luv",
	"mac",
	"mad",
	"mal",
	"man",
	"map",
	"map",
	"mar",
	"mat",
	"mba",
	"med",
	"mel",
	"men",
	"met",
	"moo",
	"msg",
	"mug",
	"nab",
	"nag",
	"nan",
	"nap",
	"nay",
	"ned",
	"neo",
	"net",
	"new",
	"nil",
	"nod",
	"nod",
	"nor",
	"not",
	"not",
	"oak",
	"oar",
	"oat",
	"odd",
	"ode",
	"off",
	"oil",
	"old",
	"one",
	"opt",
	"orb",
	"ore",
	"owe",
	"pad",
	"pal",
	"pan",
	"par",
	"pat",
	"pay",
	"pea",
	"peg",
	"peg",
	"pen",
	"pep",
	"pet",
	"pho",
	"pie",
	"poe",
	"pot",
	"pre",
	"pro",
	"pun",
	"qed",
	"qui",
	"quo",
	"rag",
	"rah",
	"ram",
	"ran",
	"rap",
	"rat",
	"raw",
	"red",
	"red",
	"rep",
	"rev",
	"rio",
	"roe",
	"rue",
	"rye",
	"sad",
	"sag",
	"sal",
	"sam",
	"sap",
	"sat",
	"say",
	"set",
	"she",
	"sin",
	"sir",
	"sly",
	"sod",
	"soy",
	"spa",
	"sry",
	"sty",
	"sun",
	"tab",
	"tad",
	"tam",
	"tan",
	"tap",
	"tar",
	"tat",
	"tea",
	"ted",
	"tee",
	"ten",
	"the",
	"the",
	"tie",
	"tnt",
	"toe",
	"ton",
	"top",
	"tot",
	"try",
	"two",
	"txt",
	"ufo",
	"uni",
	"uni",
	"uno",
	"ups",
	"urn",
	"usa",
	"use",
	"ute",
	"vac",
	"van",
	"van",
	"vat",
	"vet",
	"vex",
	"via",
	"vie",
	"vim",
	"vim",
	"vip",
	"vol",
	"von",
	"vow",
	"wad",
	"wag",
	"wan",
	"war",
	"was",
	"was",
	"way",
	"web",
	"wed",
	"wee",
	"wet",
	"who",
	"who",
	"wig",
	"win",
	"wit",
	"woe",
	"won",
	"woo",
	"wow",
	"wry",
	"wut",
	"xyz",
	"xyz",
	"yak",
	"yam",
	"yap",
	"yaw",
	"yea",
	"yen",
	"yep",
	"yes",
	"yes",
	"yet",
	"yew",
	"yin",
	"yip",
	"yon",
	"you",
	"zag",
	"zap",
	"zed",
	"zee",
	"zen",
	"zig",
	"zip",
	"zit",
	"zoe",
	"zoo",
}