// This file is available under the Apache 2.0 License
// This file was copied from: https://github.com/briandowns/spinner
//
// Please see the LICENSE file for the copy of the Apache 2.0 License.
//
// Modifications:
//
// - removed runtime generation of CharSets 37 and 38; made them literals

package yacspin

// CharSets contains the default character sets from
// https://github.com/briandowns/spinner.
var CharSets = map[int][]string{
	0:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	1:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
	2:  {"▖", "▘", "▝", "▗"},
	3:  {"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
	4:  {"◢", "◣", "◤", "◥"},
	5:  {"◰", "◳", "◲", "◱"},
	6:  {"◴", "◷", "◶", "◵"},
	7:  {"◐", "◓", "◑", "◒"},
	8:  {".", "o", "O", "@", "*"},
	9:  {"|", "/", "-", "\\"},
	10: {"◡◡", "⊙⊙", "◠◠"},
	11: {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	12: {">))'>", " >))'>", "  >))'>", "   >))'>", "    >))'>", "   <'((<", "  <'((<", " <'((<"},
	13: {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	14: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	15: {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	16: {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	17: {"■", "□", "▪", "▫"},
	18: {"←", "↑", "→", "↓"},
	19: {"╫", "╪"},
	20: {"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
	21: {"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
	22: {"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
	23: {"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
	24: {"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
	25: {"ｦ", "ｧ", "ｨ", "ｩ", "ｪ", "ｫ", "ｬ", "ｭ", "ｮ", "ｯ", "ｱ", "ｲ", "ｳ", "ｴ", "ｵ", "ｶ", "ｷ", "ｸ", "ｹ", "ｺ", "ｻ", "ｼ", "ｽ", "ｾ", "ｿ", "ﾀ", "ﾁ", "ﾂ", "ﾃ", "ﾄ", "ﾅ", "ﾆ", "ﾇ", "ﾈ", "ﾉ", "ﾊ", "ﾋ", "ﾌ", "ﾍ", "ﾎ", "ﾏ", "ﾐ", "ﾑ", "ﾒ", "ﾓ", "ﾔ", "ﾕ", "ﾖ", "ﾗ", "ﾘ", "ﾙ", "ﾚ", "ﾛ", "ﾜ", "ﾝ"},
	26: {".", "..", "..."},
	27: {"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▉", "▊", "▋", "▌", "▍", "▎", "▏", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁"},
	28: {".", "o", "O", "°", "O", "o", "."},
	29: {"+", "x"},
	30: {"v", "<", "^", ">"},
	31: {">>--->", " >>--->", "  >>--->", "   >>--->", "    >>--->", "    <---<<", "   <---<<", "  <---<<", " <---<<", "<---<<"},
	32: {"|", "||", "|||", "||||", "|||||", "||||||", "|||||||", "||||||||", "|||||||", "||||||", "|||||", "||||", "|||", "||", "|"},
	33: {"[          ]", "[=         ]", "[==        ]", "[===       ]", "[====      ]", "[=====     ]", "[======    ]", "[=======   ]", "[========  ]", "[========= ]", "[==========]"},
	34: {"(*---------)", "(-*--------)", "(--*-------)", "(---*------)", "(----*-----)", "(-----*----)", "(------*---)", "(-------*--)", "(--------*-)", "(---------*)"},
	35: {"█▒▒▒▒▒▒▒▒▒", "███▒▒▒▒▒▒▒", "█████▒▒▒▒▒", "███████▒▒▒", "██████████"},
	36: {"[                    ]", "[=>                  ]", "[===>                ]", "[=====>              ]", "[======>             ]", "[========>           ]", "[==========>         ]", "[============>       ]", "[==============>     ]", "[================>   ]", "[==================> ]", "[===================>]"},
	37: {"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"},                                                             // clock emoji: one per hour for hours 1~12
	38: {"🕐", "🕜", "🕑", "🕝", "🕒", "🕞", "🕓", "🕟", "🕔", "🕠", "🕕", "🕡", "🕖", "🕢", "🕗", "🕣", "🕘", "🕤", "🕙", "🕥", "🕚", "🕦", "🕛", "🕧"}, // clock emoji: one per half hour for hours 1~12
	39: {"🌍", "🌎", "🌏"},
	40: {"◜", "◝", "◞", "◟"},
	41: {"⬒", "⬔", "⬓", "⬕"},
	42: {"⬖", "⬘", "⬗", "⬙"},
	43: {"[>>>          >]", "[]>>>>        []", "[]  >>>>      []", "[]    >>>>    []", "[]      >>>>  []", "[]        >>>>[]", "[>>          >>]"},
	44: {"♠", "♣", "♥", "♦"},
	45: {"➞", "➟", "➠", "➡", "➠", "➟"},
	46: {"  |  ", ` \   `, "_    ", ` \   `, "  |  ", "   / ", "    _", "   / "},
	47: {"  . . . .", ".   . . .", ". .   . .", ". . .   .", ". . . .  ", ". . . . ."},
	48: {" |     ", "  /    ", "   _   ", `    \  `, "     | ", `    \  `, "   _   ", "  /    "},
	49: {"⎺", "⎻", "⎼", "⎽", "⎼", "⎻"},
	50: {"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"},
	51: {"[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"},
	52: {"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )"},
	53: {"✶", "✸", "✹", "✺", "✹", "✷"},
	54: {"▐|\\____________▌", "▐_|\\___________▌", "▐__|\\__________▌", "▐___|\\_________▌", "▐____|\\________▌", "▐_____|\\_______▌", "▐______|\\______▌", "▐_______|\\_____▌", "▐________|\\____▌", "▐_________|\\___▌", "▐__________|\\__▌", "▐___________|\\_▌", "▐____________|\\▌", "▐____________/|▌", "▐___________/|_▌", "▐__________/|__▌", "▐_________/|___▌", "▐________/|____▌", "▐_______/|_____▌", "▐______/|______▌", "▐_____/|_______▌", "▐____/|________▌", "▐___/|_________▌", "▐__/|__________▌", "▐_/|___________▌", "▐/|____________▌"},
	55: {"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌", "▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌", "▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌", "▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌", "▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌", "▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌", "▐⠠       ▌"},
	56: {"¿", "?"},
	57: {"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"},
	58: {"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"},
	59: {".  ", ".. ", "...", " ..", "  .", "   "},
	60: {".", "o", "O", "°", "O", "o", "."},
	61: {"▓", "▒", "░"},
	62: {"▌", "▀", "▐", "▄"},
	63: {"⊶", "⊷"},
	64: {"▪", "▫"},
	65: {"□", "■"},
	66: {"▮", "▯"},
	67: {"-", "=", "≡"},
	68: {"d", "q", "p", "b"},
	69: {"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"},
	70: {"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "},
	71: {"☗", "☖"},
	72: {"⧇", "⧆"},
	73: {"◉", "◎"},
	74: {"㊂", "㊀", "㊁"},
	75: {"⦾", "⦿"},
	76: {"ဝ", "၀"},
	77: {"▌", "▀", "▐▄"},
}
