totalBlocks = 0
totalPieces = 0
timestampedBlocks = {}
timestampedPieces = {}
timestampedPercentage = {}
timestampedSize = {}


with open("profile") as file:
	for line in file:
		if line[0:3] == '===':
			# print line
			words = line.split(' ')
			words = words[2:]
			print words
			if words[4] == 'Received':
				timestampedBlocks[words[0]] = words[7]
				timestampedPieces[words[0]] = words[8]
				block = words[7]
				# print block
				if words[6] != totalPieces:
					totalBlocks += int(block)
					totalPieces = words[6]
			elif words[4] == 'Completed':
				timestampedPercentage[words[0]] = words[6]
			elif words[4] == 'Downloaded':
				timestampedSize[words[0]] = words[6]
				
print "Total blocks: ", totalBlocks
print "Total pieces: ", totalPieces