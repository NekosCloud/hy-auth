package main

var cookieStoreKey = []byte{
	68, 234, 15, 58, 124, 32, 161, 44,
	49, 214, 72, 84, 95, 31, 25, 205,
	169, 37, 163, 72, 107, 230, 23, 213,
	177, 41, 237, 17, 112, 246, 31, 252,
	182, 137, 124, 189, 204, 29, 246, 76,
	40, 6, 91, 133, 35, 27, 169, 77,
	116, 120, 129, 152, 198, 113, 23, 116,
	149, 234, 105, 171, 45, 73, 48, 111,
}

var cookieStoreEncryption = []byte{
	136, 97, 53, 14, 208, 7, 176, 205,
	206, 126, 225, 249, 60, 28, 75, 9,
	245, 207, 78, 237, 114, 178, 159, 155,
	34, 139, 76, 188, 65, 222, 142, 209,
}

var sessionStoreKey = []byte{
	121, 103, 117, 190, 175, 13, 112, 89,
	95, 103, 66, 11, 152, 69, 199, 223,
	97, 65, 230, 152, 153, 15, 88, 248,
	175, 24, 109, 25, 118, 207, 212, 28,
	76, 155, 16, 208, 243, 68, 121, 23,
	109, 31, 133, 73, 232, 215, 27, 28,
	152, 42, 216, 253, 46, 24, 240, 242,
	92, 79, 14, 81, 197, 222, 185, 60,
}

var sessionStoreEncryption = []byte{
	44, 18, 25, 100, 179, 150, 31, 221,
	194, 240, 40, 80, 201, 64, 194, 172,
	8, 23, 188, 252, 10, 168, 48, 29,
	189, 95, 199, 109, 174, 191, 128, 92,
}