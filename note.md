Commands:
1E (hex) -> 30
10 (bin) -> 2
N go (up) -> GO
N SHOUTING (low) -> shouting 
N fOOLISH (cap) -> Foolish
punctuation handling
group punctuation 
quotues
article 


Order of operations. Apply the rules in this order so the result is well defined: first resolve the tag rules ((hex), (bin), (up), (low), (cap)) and remove the tags, then fix punctuation spacing, then apply the a → an article rule. The provided examples assume this order.


while processing if tag is invalid then I have to left it as it and also give notification that that was not changed




Language detector 


( up , 8 )
(up, 8)



	/* LOGIC OF THIS FUNCTION
	range loop for token in tokens
	if token is tag -> apply tag function where tag executes based on tag itself+
	if it starts with (a, e, i, o, u ,h) && A or a then article function
	last fixing with punctuation func 
	returns ready []string
	*/


for input.txt



TO DO LIST 08.08.2026(Saturday):
// fix error ( said :' I  )
// write func that convert from []string to string
// writing everything to output file

// ai
// Makefile
// README
