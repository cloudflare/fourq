package fourq

import (
	"math/big"
)

var (
	Order, _ = new(big.Int).SetString("73846995687063900142583536357581573884798075859800097461294096333596429543", 10)
	Gx, Gy   *big.Int

	aMask uint64 = 0x7fffffffffffffff
	bMask uint64 = 0xffffffffffffffff

	g = &point{
		x: gfP2{
			x: baseFieldElem{0x286592ad7b3833aa, 0x1a3472237c2fb305},
			y: baseFieldElem{0x96869fb360ac77f6, 0x1e1f553f2878aa9c},
		},
		y: gfP2{
			x: baseFieldElem{0xb924a2462bcbb287, 0xe3fee9ba120785a},
			y: baseFieldElem{0x49a7c344844c8b5c, 0x6e1c4af8630e0242},
		},
		z: *newGFp2().SetOne(),
	}

	d = &gfP2{
		x: baseFieldElem{0x142, 0xe4},
		y: baseFieldElem{0xb3821488f1fc0c8d, 0x5e472f846657e0fc},
	}

	// generatorBase is pre-computed multiples of the curve's generator.
	generatorBase = [...]*point{
		&point{
			x: gfP2{baseFieldElem{0x0000000000000000, 0x0000000000000000}, baseFieldElem{0x0000000000000000, 0x0000000000000000}},
			y: gfP2{baseFieldElem{0x0000000000000001, 0x0000000000000000}, baseFieldElem{0x0000000000000000, 0x0000000000000000}},
			t: gfP2{baseFieldElem{0xffffffffffffffff, 0x7fffffffffffffff}, baseFieldElem{0x0000000000000000, 0x0000000000000000}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x286592ad7b3833aa, 0x1a3472237c2fb305}, baseFieldElem{0x96869fb360ac77f6, 0x1e1f553f2878aa9c}},
			y: gfP2{baseFieldElem{0xb924a2462bcbb287, 0x0e3fee9ba120785a}, baseFieldElem{0x49a7c344844c8b5c, 0x6e1c4af8630e0242}},
			t: gfP2{baseFieldElem{0x894ba36ee8cee416, 0x35bfa1947fb0913e}, baseFieldElem{0x673c574d296cd8d0, 0x7bfb41a38e7076ac}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xdffd6556d311ce43, 0x210a7d9f9782a38c}, baseFieldElem{0x023c5e59afc61df4, 0x58d4179cfc261e7b}},
			y: gfP2{baseFieldElem{0x35a2323d01cb626c, 0x2db3fc78c3d93dfe}, baseFieldElem{0xee7c9525e2919bf8, 0x44c04cb98a015452}},
			t: gfP2{baseFieldElem{0xd9bd090b4a993f99, 0x69dafd67125a8e46}, baseFieldElem{0xd3fe212ab9612f83, 0x0ddaa61be71d1f99}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x821ff2e80dc5e252, 0x6a9819b5c0f0f512}, baseFieldElem{0x7f29641b85d56f5c, 0x1dd2c4814e7439e7}},
			y: gfP2{baseFieldElem{0x070763c94e098671, 0x6caaddc6d7b431a8}, baseFieldElem{0xb4e0f6026423303e, 0x771ca389a001970f}},
			t: gfP2{baseFieldElem{0xc33adbdde93e4249, 0x5a76cd697c181ae1}, baseFieldElem{0xe2c01d7a9bcddf88, 0x2f44f58091142d38}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x66161959b80546a7, 0x5ed74234051a217b}, baseFieldElem{0xc64448cc446529b2, 0x25e463a9625bea0b}},
			y: gfP2{baseFieldElem{0x9f9097db745193e8, 0x75ee43643940c19d}, baseFieldElem{0x7eb80ac86a7981ef, 0x6497521029d87701}},
			t: gfP2{baseFieldElem{0x13228459ee8f996b, 0x4bf482f6b439200a}, baseFieldElem{0x45a48ef604943ef8, 0x436b7c13a0fcf62e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1ffea44f74fb99b9, 0x7b245d0b4cc0c060}, baseFieldElem{0xc4e6eddd7cf59e81, 0x7c199861991b6f1c}},
			y: gfP2{baseFieldElem{0xff4b55c525692082, 0x647419c9cca643f0}, baseFieldElem{0x3e1a1727f25fb9e9, 0x05f581f3e4746809}},
			t: gfP2{baseFieldElem{0x01113fcee82a6534, 0x11c36a4b8f6970fc}, baseFieldElem{0x2548f9cd107c15d7, 0x3219befb139fd127}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2239ece38a075309, 0x0667fc4bd09c18d9}, baseFieldElem{0x6377686a6b43513e, 0x59ea75d0445abd9d}},
			y: gfP2{baseFieldElem{0x6c3079477fa0e2da, 0x3b67adc4c938fc5a}, baseFieldElem{0xf770768173f59469, 0x1a9e014c10eca4f2}},
			t: gfP2{baseFieldElem{0xd6ebc0f505cab344, 0x3b4347d0d3e569cf}, baseFieldElem{0x8d95879b2c45f676, 0x679461bfbc8f036e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x533062953f8e4cb9, 0x2041b5b9f098b1b3}, baseFieldElem{0x907f105792a9bd99, 0x7a3e836ccd3ee8de}},
			y: gfP2{baseFieldElem{0x9ca9d389059fcabc, 0x59c85712bc2f8748}, baseFieldElem{0x1fff68c4793c1ec5, 0x2a9b33474b8ce2b5}},
			t: gfP2{baseFieldElem{0xc407a53633b21097, 0x4fb795f4d3c70b13}, baseFieldElem{0xec104047ba5cc49b, 0x747ef881dfabcf91}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x431f41fb2252cc0d, 0x46a99df6da5bd7b9}, baseFieldElem{0xc5bd84fbb84a3dbe, 0x7132fe9df342729a}},
			y: gfP2{baseFieldElem{0x82489442c0baf355, 0x45caf3b446c89bd4}, baseFieldElem{0xc9520db6e2e4f2d1, 0x3be71fb0a5d3de6d}},
			t: gfP2{baseFieldElem{0x9cfe820cf582f4c1, 0x783a37e29952bc41}, baseFieldElem{0x3f3b5cbc068fa30a, 0x42e85455bf52a7b2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xafc5a6722f20ac6d, 0x6732bf6f801ef46b}, baseFieldElem{0x90f0cb09a515f89c, 0x69ef409b9ff40f06}},
			y: gfP2{baseFieldElem{0xbfbb62ed47edd7ab, 0x6c8918fbf377c79d}, baseFieldElem{0x9a81ef68c6154973, 0x786e9981890db884}},
			t: gfP2{baseFieldElem{0x6fcbbaf2235db183, 0x5a91f3333d070130}, baseFieldElem{0x28aa2eaaa13ee4c6, 0x0b65ff3b73009b54}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3b1643ec2e277c21, 0x7a9012b08cd75939}, baseFieldElem{0xe3f5877929fbbecc, 0x31c16e959a4834b2}},
			y: gfP2{baseFieldElem{0x272e9f93326a113d, 0x30b3b2b7c1af7c8f}, baseFieldElem{0xbd8108e9ffd1ae09, 0x30c366c885e08a29}},
			t: gfP2{baseFieldElem{0xe6aa5e4022f39f0b, 0x518e41d19abb46e3}, baseFieldElem{0x5a8c9e04fa584154, 0x2b6b2eba76219a47}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe53263a7bb1e9952, 0x0a5e06a44689d29e}, baseFieldElem{0x38b33fcb1dde7580, 0x5e5d8aa07cef0853}},
			y: gfP2{baseFieldElem{0x040b8285af60fb45, 0x16cb3ef9406a769e}, baseFieldElem{0x0cb054fca6865a64, 0x02c6aa5e47059960}},
			t: gfP2{baseFieldElem{0x600e4ebaffca7b7d, 0x6620505b95c52f3d}, baseFieldElem{0xc68f8a6f806f79f7, 0x11371573eff1851c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb92b573d2c4b06ff, 0x6b62d585800a9f6a}, baseFieldElem{0xecb6dfb3fa1acb7c, 0x0d9d9f54a8335e2b}},
			y: gfP2{baseFieldElem{0xdf3bd744d9bb783d, 0x2b827eeda23988a6}, baseFieldElem{0x947c187247366cdd, 0x3b7e00ba2f9525b3}},
			t: gfP2{baseFieldElem{0x8ed1dec866469842, 0x26f73a30be35882e}, baseFieldElem{0xb739e78fdbd67f5d, 0x0c1d77d0df4ce382}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xf255f8d3cb2257a2, 0x5147de82dc70d94f}, baseFieldElem{0x0753c3b4cf5e3e5e, 0x003b8d95055fdcb2}},
			y: gfP2{baseFieldElem{0x407acf7453c26c17, 0x1aeea88dc06c3f76}, baseFieldElem{0xddd3535f6061577c, 0x6784168aafcd0d14}},
			t: gfP2{baseFieldElem{0x5cfb2af3e9da4ca3, 0x64b9f0a9a5332684}, baseFieldElem{0x4ff99faf5ecf1036, 0x3f67644ed65cf859}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2f9f29622b93b8c3, 0x764cd83f65b58163}, baseFieldElem{0xd0652084cefd19cf, 0x4233bf5c7885c810}},
			y: gfP2{baseFieldElem{0xea2c746fa21a90fa, 0x03ab2058c136a9a3}, baseFieldElem{0x789d5f6bf45c01b0, 0x0136c9823cd8dbf4}},
			t: gfP2{baseFieldElem{0xf17cf2eb6d5670a0, 0x5baf8b3a82bb28ec}, baseFieldElem{0x57eb5c979f3915f0, 0x0c1ae18069902148}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xff4bdab28d3a7750, 0x61330ad8675bbff8}, baseFieldElem{0x508c3357eca15a5f, 0x2cd8db72eb7c074c}},
			y: gfP2{baseFieldElem{0xf948934df0e92d18, 0x416e8c22547271a9}, baseFieldElem{0x027dbac3ef70bc5a, 0x77370e224cab9bd8}},
			t: gfP2{baseFieldElem{0x1a344ed24655e2ca, 0x1529c386f4ab0f61}, baseFieldElem{0x26bb146305513e95, 0x2ad30f3f2e38305d}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd3e4ebbf372235b4, 0x7c112e3c41407291}, baseFieldElem{0x4a17119d649f9d40, 0x2ac8b7c7b4376705}},
			y: gfP2{baseFieldElem{0xe18e5696ac93eb5e, 0x3e0f67dad044a500}, baseFieldElem{0x2b1840f29c624ee7, 0x30f9588cfd532740}},
			t: gfP2{baseFieldElem{0xf64711fc6a224816, 0x6833afdf948f8384}, baseFieldElem{0x09338cd4043addbb, 0x457a71d9dfbb8aa5}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x6ca1e6dcd0204bec, 0x7a4edff7eb6dcf8a}, baseFieldElem{0xcc43992e1bfbd17d, 0x25348f38117139b0}},
			y: gfP2{baseFieldElem{0x78549240a6657808, 0x51d2a529d8068b11}, baseFieldElem{0x3fc188817c598e19, 0x6f2e11f14259a0ca}},
			t: gfP2{baseFieldElem{0x4ab8f09dd90c907d, 0x5b000a03e1c28c5d}, baseFieldElem{0xc491faf77744dbe2, 0x2499ac543a5b9c98}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x6ec321a8fe862fad, 0x49061263c9afa20e}, baseFieldElem{0x43a5bde2588f2019, 0x53d3d30b303e759e}},
			y: gfP2{baseFieldElem{0xbe710371e852ed3b, 0x436aed6dddb0eea2}, baseFieldElem{0xc21b7f37cd302edd, 0x5dfe3931d431afb4}},
			t: gfP2{baseFieldElem{0x83808be234957e84, 0x1d12149f0c446c20}, baseFieldElem{0x5095513a91f20a8c, 0x386a5da33a6620cf}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x5121482386938889, 0x0bf392b410798f7b}, baseFieldElem{0xe7a840297bfbe4ec, 0x6b99cee4de8496f5}},
			y: gfP2{baseFieldElem{0x0200f8baf1601969, 0x0bc9966a9a21416f}, baseFieldElem{0xf8fa97c646fcbbb3, 0x4eda4320419ce098}},
			t: gfP2{baseFieldElem{0x2918b66efb09f2c0, 0x41d5a51a9b227859}, baseFieldElem{0x6d1d0e6895589653, 0x2d9a2aa326475159}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x87209762bfa3b3da, 0x167f432e3d2541ac}, baseFieldElem{0x08d484fbbab44477, 0x5dcd5b7e375f2542}},
			y: gfP2{baseFieldElem{0xe6b8209909543b61, 0x339f437273651783}, baseFieldElem{0xd121ad82ed5b3095, 0x3e444a6f8c7c99d4}},
			t: gfP2{baseFieldElem{0x0a4ab2744964672b, 0x31f3f881fd5ef09e}, baseFieldElem{0x358412700e8b6b5b, 0x357be4331ca7ef95}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4ec208bf0d66d9b7, 0x246954f2c4de5b22}, baseFieldElem{0x0232be2cbfbfb162, 0x729110cad9138819}},
			y: gfP2{baseFieldElem{0x563484616d1f4f26, 0x2c68dbc253b58d1f}, baseFieldElem{0xf18ab86d8bf55353, 0x63064b3b6055a115}},
			t: gfP2{baseFieldElem{0x8af1dc3058036565, 0x6fe0f6c97cf69e31}, baseFieldElem{0x1e1ca6df9b2bc11d, 0x75eef0ddeb0c4bb3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x57511ca3ea131553, 0x20cbd7378e15abb5}, baseFieldElem{0x29e1a20d1335aa5b, 0x5855cc4b796448b9}},
			y: gfP2{baseFieldElem{0x28cc5f7ef709e7f3, 0x355732f9089ccefc}, baseFieldElem{0xd837aabaad2ca71e, 0x16d7c92c3c1973d7}},
			t: gfP2{baseFieldElem{0x4eca77ce65f3ac67, 0x271c9b280b048f44}, baseFieldElem{0x645919bb6f5eec9d, 0x3efadcc675302d66}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8b88a7b523e0a8f1, 0x4d2712bf08fb277a}, baseFieldElem{0x28bb73bc78506138, 0x2b3429fba35bd0be}},
			y: gfP2{baseFieldElem{0x891911699d96325a, 0x08d207a4cd9f86f2}, baseFieldElem{0xcb680e7602cb0bee, 0x7f9af54bbc50da3f}},
			t: gfP2{baseFieldElem{0xce0b04be2ddae364, 0x332d07d83c6cbb14}, baseFieldElem{0xb91539e271d5b1d7, 0x4423494026352ee7}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x549048ebdbd33cd8, 0x3183a8e2d4791ba6}, baseFieldElem{0x00d3b6b8c78e6d65, 0x41fc757663243eb2}},
			y: gfP2{baseFieldElem{0x0a8f3352e0ec0761, 0x387d83e71a722f20}, baseFieldElem{0x3bac82f9e613a444, 0x0d855620f4899f91}},
			t: gfP2{baseFieldElem{0x0b894f512b8c524f, 0x2b5ac76371854b5b}, baseFieldElem{0x85a93a99d4e88aa8, 0x30d7d7dc2cc99a77}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x5f7306a4b466dc57, 0x2904f5bc605a0b36}, baseFieldElem{0xb9ebe1f4b6e39127, 0x7fa41f600779b7a9}},
			y: gfP2{baseFieldElem{0x70ad812d52984f68, 0x029799111f51dc00}, baseFieldElem{0xf8c98b438be719e5, 0x04ca81576efd48fd}},
			t: gfP2{baseFieldElem{0x7a3ab710798d4393, 0x4133dff6742f7e00}, baseFieldElem{0x12ce0a1394817165, 0x0082e4b631ff77cb}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x598898dc14be16d0, 0x55eeefa8312fd7de}, baseFieldElem{0xee55a8dd1fa6c097, 0x1f2b5ceb9a6805e5}},
			y: gfP2{baseFieldElem{0x30476d974da8f2d6, 0x2afe3fb138ee6346}, baseFieldElem{0xc9e71da9afd4b3eb, 0x214431c47d94bf32}},
			t: gfP2{baseFieldElem{0xebdcfb433d2257b8, 0x32cc2e73f50d320c}, baseFieldElem{0xc1dda2f2c7e11151, 0x1809e6d690ebf2cf}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x13e325ae03964686, 0x5e447f054d2fc925}, baseFieldElem{0x1d4e1abcd0d992b8, 0x041fd31bd8d7bc58}},
			y: gfP2{baseFieldElem{0x8f91ccd4b7ea8666, 0x1a59e1967a4b17f7}, baseFieldElem{0xf3843cc088c1b33a, 0x1834b83fac9f6558}},
			t: gfP2{baseFieldElem{0xf2915639b46f1e19, 0x1eaf1231030e5fad}, baseFieldElem{0xc80b3e6799924b18, 0x6afed44cdc74b2ac}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0b5e439cb9e85bf6, 0x1f3ebe6c87ea0eff}, baseFieldElem{0x52618d5f344409f9, 0x753e9e8f88641db2}},
			y: gfP2{baseFieldElem{0xb3a8f97726bc2878, 0x5678aaadbe0c247f}, baseFieldElem{0x74ebeec150b912f9, 0x777f4beba4a261b2}},
			t: gfP2{baseFieldElem{0x255c222e26df6908, 0x7cdbef1d35b1dd96}, baseFieldElem{0xd1b7b4d9e885fa65, 0x4fc406673da4b45f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd62b0271a12e75d9, 0x4834608e16e5dc89}, baseFieldElem{0xd78cc2058d5f922f, 0x65c8c8f75f23672f}},
			y: gfP2{baseFieldElem{0xd5a7bf6b536e3fca, 0x77204c6e0ee05785}, baseFieldElem{0x48a22df52e718a8e, 0x0158a2855445f75c}},
			t: gfP2{baseFieldElem{0x964588fd62c29def, 0x0f04d9444fe43299}, baseFieldElem{0xc9edf41953e6771e, 0x12a0ba54c28584a0}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4606dc91edb838e7, 0x269182d754b0e2a5}, baseFieldElem{0x16c002fce112e56d, 0x6874611159e3a3f7}},
			y: gfP2{baseFieldElem{0xd887c044282feb0d, 0x043184cdec85b205}, baseFieldElem{0x2503fa98530ff320, 0x76c96bd470d52f8a}},
			t: gfP2{baseFieldElem{0x06579dea53128596, 0x1df97ef964255a85}, baseFieldElem{0x43259ece22915158, 0x7a97734c83b1118f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x870b5b8f65068e98, 0x49c26117656a448c}, baseFieldElem{0x3b9d49987636b94c, 0x2ec814cd687539dd}},
			y: gfP2{baseFieldElem{0xaa5d2177dcdbbf03, 0x39192e13eb303835}, baseFieldElem{0x56a6aec0adfb4bdb, 0x39fb4c22b4f9334e}},
			t: gfP2{baseFieldElem{0xe1b1c5a37b7485ae, 0x50c042cb1a84ee0d}, baseFieldElem{0x044d93b84c10557c, 0x64840ef56ff6cc24}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3573739fca037c44, 0x5586fb6dbd2a8294}, baseFieldElem{0xe647aa21ee392052, 0x4c53e4dfd63cf77b}},
			y: gfP2{baseFieldElem{0xeb255170a17584d1, 0x048ec7c597200739}, baseFieldElem{0x057fcbdd6afb0962, 0x38969192b10c0a72}},
			t: gfP2{baseFieldElem{0x2ed21d0140b36e2d, 0x2ab1dc2488b5c1c4}, baseFieldElem{0xa665ae5469a375e2, 0x4ad597acb173e83e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x146c309674d3e80c, 0x56bbfc5f066f5226}, baseFieldElem{0x4b0b00b4e5c10bcc, 0x61428b493c021df3}},
			y: gfP2{baseFieldElem{0x106b61defaf57e34, 0x0ad3de40e9173a38}, baseFieldElem{0x6cf45a5d753c87a9, 0x164b4e34c83ea465}},
			t: gfP2{baseFieldElem{0x44620d86432b79a0, 0x1d019905a720efe9}, baseFieldElem{0x7ea2b73910ae2929, 0x6ea28a0ee99c8e45}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xff258bb8e89ec101, 0x1585c1d10a1fa63e}, baseFieldElem{0xe83bddf07bd366f1, 0x435ddc2155c6ca14}},
			y: gfP2{baseFieldElem{0xec0219631260d70f, 0x21ecd416fce695fb}, baseFieldElem{0x480ae4d71407a53c, 0x74ce6e7070354263}},
			t: gfP2{baseFieldElem{0x95c9416f8f2725d8, 0x7f3b2403d4e236b7}, baseFieldElem{0xe17b49f848f6c66b, 0x4d1ab2d46cc4fe7b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd97e40fae1d59a16, 0x7661ae4e73555046}, baseFieldElem{0xdc2afa9cb912e774, 0x153e24507924d652}},
			y: gfP2{baseFieldElem{0x53440ab04f60f9b5, 0x29f26dc4460ef433}, baseFieldElem{0x5b1e7b1a426edc58, 0x3d523aaabb2941a4}},
			t: gfP2{baseFieldElem{0x0a83598d45bdb6d2, 0x002e9196d708f7db}, baseFieldElem{0xd6b00275d5b32176, 0x42dbf55f2dee8a5c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xf6bee2c6d1d57c46, 0x4b89669008ca2d36}, baseFieldElem{0xeefaccdf4ff793bd, 0x76a99a43a9fb23a2}},
			y: gfP2{baseFieldElem{0x6195b0304327bdb7, 0x59a2c17502114162}, baseFieldElem{0x811389ee57d227ab, 0x251b956a406733c4}},
			t: gfP2{baseFieldElem{0xa000145a7379cd48, 0x5df865ef5aae36ec}, baseFieldElem{0x64c9fd9c12a832a8, 0x38836a0cf13be81c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9e58fad0232424ea, 0x1f6d9105660b1c20}, baseFieldElem{0xc857c3d9d01a8efb, 0x71f6cb61d87ab0df}},
			y: gfP2{baseFieldElem{0x1f3f1b678f1a5479, 0x49afea77527ef2d5}, baseFieldElem{0x48626dc111eb9a18, 0x150448edaa1af77b}},
			t: gfP2{baseFieldElem{0x60e3c531bc0fcb88, 0x1a9a54e4d63cdcf5}, baseFieldElem{0x665f774d37c7869b, 0x48596a1aa2e3ef06}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd1e5193da0ebf0bb, 0x0172c391cab4e5bf}, baseFieldElem{0x7be6694323315fba, 0x50fd7a4173070142}},
			y: gfP2{baseFieldElem{0xbb6558b0268666ea, 0x355cd7a79d6afc0b}, baseFieldElem{0x7da3b34016f00a84, 0x7e5c400aa7bb2914}},
			t: gfP2{baseFieldElem{0x1cf508c0a272abb4, 0x70ab58f6309a7de8}, baseFieldElem{0x42dff7919febf9b6, 0x6b3578391800f417}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0424874df1f3d9d7, 0x239184bb27e62a87}, baseFieldElem{0x3992682dae9cd9f2, 0x3a2865b0b9498114}},
			y: gfP2{baseFieldElem{0x70311d136431c394, 0x062b3ba84ce6b72e}, baseFieldElem{0xb968a6a5fbead50a, 0x66f1a0fe550b5c44}},
			t: gfP2{baseFieldElem{0xbcbb60645862f7de, 0x7b682025da2a4d18}, baseFieldElem{0x4ff7ad1524ab50bc, 0x0fa0b2856cb26de2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xf5e2a454fed1ed28, 0x3cb3ca236c00c7d5}, baseFieldElem{0xf03fb674eb72ce8b, 0x3da295be4acb13b9}},
			y: gfP2{baseFieldElem{0xf8cc335e0804d114, 0x7b1313e510ed05fb}, baseFieldElem{0xe401dadee412b4bd, 0x07637bb3817c7b3e}},
			t: gfP2{baseFieldElem{0x35b3a2eba3b1e086, 0x4772cbd85f210ee0}, baseFieldElem{0xbdae5ffd7853e3e4, 0x7e9d7d0115766e1a}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x074d6f531651a3cc, 0x037ecc86d72fa4cc}, baseFieldElem{0xbef7984223383e2a, 0x065929fc404c5c9e}},
			y: gfP2{baseFieldElem{0xfea66882e04ef0ff, 0x26b8a54db7036f38}, baseFieldElem{0x49f7a1a7b95eb1de, 0x7ac91e3af9e9ecf3}},
			t: gfP2{baseFieldElem{0xda89908b9aaf4228, 0x7e0005adf35a6ba9}, baseFieldElem{0x56ce0c8e31491e66, 0x4e0f0f9495cb385f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x35e804878256125b, 0x1a00c69858b6d95f}, baseFieldElem{0x0a7c3067f109df1b, 0x5edff2c038c8503d}},
			y: gfP2{baseFieldElem{0x8044391d9a70d503, 0x1918558062957118}, baseFieldElem{0x1dc85e485495e16b, 0x5aab6af50539ceb1}},
			t: gfP2{baseFieldElem{0xdfb1f391aaf94fce, 0x6e03f7259b2edc65}, baseFieldElem{0x73a1fa969087a66e, 0x7d22fcbec8268cc2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x6455ee8f388b871b, 0x143d44655e64878c}, baseFieldElem{0x8443cd250213b013, 0x22d0758e6f8cc790}},
			y: gfP2{baseFieldElem{0x573b34866939814e, 0x368b894b9c1a3423}, baseFieldElem{0x80fe0a255e838ecc, 0x4f641a74c525fe46}},
			t: gfP2{baseFieldElem{0x8ff11479c367a0ff, 0x2d05706a9f9edd5f}, baseFieldElem{0x694bab066a169114, 0x51d7f34fd7100e33}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd09eaeb44cd355ba, 0x4d0f30d13efffcaa}, baseFieldElem{0x3ea0d3ff9a2cb1c0, 0x46a2e282b078e5af}},
			y: gfP2{baseFieldElem{0x2f014ba43d13bafd, 0x18b61523651e285e}, baseFieldElem{0x8783de9a333dfb18, 0x7b19571d5127a0f7}},
			t: gfP2{baseFieldElem{0x7b19336681be9947, 0x3d3ca0b95d17af93}, baseFieldElem{0xfc36072ed3d6ec4c, 0x00a66cf27644e6b9}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xddaaa5f2c1a52208, 0x50e77a5d76407bac}, baseFieldElem{0xdd182944db59d623, 0x69c3555d63d2beb5}},
			y: gfP2{baseFieldElem{0x8d3dea544faa9118, 0x6c78e43cf8865c60}, baseFieldElem{0x3bd0ec824c4e9eb4, 0x7fcd335885fdd279}},
			t: gfP2{baseFieldElem{0xb2960b4e3b3b5a31, 0x548c209f2bad3ea9}, baseFieldElem{0x2fdaba4a22f28e84, 0x5a7f0c4fd57d272d}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x66c680aaee921840, 0x0f603481e45c8b8d}, baseFieldElem{0x6bf77035c25df67e, 0x278e6246df552207}},
			y: gfP2{baseFieldElem{0x9bbe4264467101ae, 0x07601729a5bd10e2}, baseFieldElem{0xee2bdd86db14177f, 0x11d58856a6e71b47}},
			t: gfP2{baseFieldElem{0x8f46d3d3f412e867, 0x2de63311786eb62f}, baseFieldElem{0x9a2087a988e0148c, 0x42a3caee08e14f08}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x17a8f1dc505208d5, 0x0a23f94011ad6b64}, baseFieldElem{0x81e9dfdea292e19d, 0x37896a9e3f10ea7a}},
			y: gfP2{baseFieldElem{0xc58ec3f014838f96, 0x74afd897c66d4a78}, baseFieldElem{0x2a69688081045b00, 0x4f7bfcbef42aa75d}},
			t: gfP2{baseFieldElem{0x4bc12afa9e9a2352, 0x7c6063b83aac9bba}, baseFieldElem{0xb49d8276fbefaeab, 0x090b8cc43704e6f4}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb9f205e77d22fbac, 0x4b95b80f3ceeb591}, baseFieldElem{0xff531a3b1dc90d51, 0x289d9029ab47d43d}},
			y: gfP2{baseFieldElem{0x882bac15e2709b8f, 0x7d40c2a38a1160aa}, baseFieldElem{0xbd9c74d5ff46671c, 0x3bae2013543c4ba8}},
			t: gfP2{baseFieldElem{0xa1cb4f90ec81cf62, 0x6e32f6da8985f246}, baseFieldElem{0xfc9dca3e98a9195b, 0x66b9bd3ecccbac8b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7671f0e49c4dfb66, 0x4dfbe1241786b29c}, baseFieldElem{0xae0e0e2ac9407101, 0x36723098ebe4aaa0}},
			y: gfP2{baseFieldElem{0x0e6c899d5dd15574, 0x01a7b0b24116cfa8}, baseFieldElem{0x0eb54b443ef44184, 0x16d89b3d1e403f2e}},
			t: gfP2{baseFieldElem{0x1a1ae360b9418708, 0x607d4737fad56fcb}, baseFieldElem{0x1cabd500bfbf5a06, 0x08ee8c483a6a4bcf}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7f0edc4acc229728, 0x56af9e0e005ca5a6}, baseFieldElem{0xe0e22b724b790b2a, 0x347950b3bfb917d6}},
			y: gfP2{baseFieldElem{0xde6df45923e2c41a, 0x78031acc78107123}, baseFieldElem{0x1ff7c3fd7952c081, 0x29a614f437feb8d3}},
			t: gfP2{baseFieldElem{0x8eed080aab8ceb7a, 0x1c6bb356c9699f4b}, baseFieldElem{0x9c0a3954ea39fcfe, 0x4f8591ecded90c37}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xeb99dafd8e7ac3de, 0x6a319f9d5f6bcabe}, baseFieldElem{0x0ad7f5349ec4132f, 0x6bac573a745e1325}},
			y: gfP2{baseFieldElem{0x6edaebbdb50d0f36, 0x176289235245bcd6}, baseFieldElem{0x51e95eadd1f79d25, 0x3f9054a48c7e373c}},
			t: gfP2{baseFieldElem{0x724bb52e2d8662d0, 0x23cd813cd1b9b4ca}, baseFieldElem{0x60bf97305ddb7c25, 0x0111b20e99d1f9ca}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x6ce752a542c98e78, 0x65b58ed1b7a371f0}, baseFieldElem{0x2bef6a3221411895, 0x5e14b3451f2a9d29}},
			y: gfP2{baseFieldElem{0x1be5fc484353ed62, 0x224a08043b333d9f}, baseFieldElem{0xd337fd1a999ced8d, 0x6e5b56d9a9d02272}},
			t: gfP2{baseFieldElem{0x055c04d957b1d4fc, 0x04f0cc18f923c218}, baseFieldElem{0xa45f95e52ab1ecf2, 0x421e09d9a40de945}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x83f325004684714c, 0x34e6ff8f947c1738}, baseFieldElem{0x125c338d419db2fa, 0x470fc69bdb73b928}},
			y: gfP2{baseFieldElem{0xf0eeedeebad4a7b4, 0x7d9091ea6f1b9a80}, baseFieldElem{0x69608ae0fbcac249, 0x447990d44195f68e}},
			t: gfP2{baseFieldElem{0xe462c1621aa34ad4, 0x408e84fbc83d1665}, baseFieldElem{0x4ab1bc52969a0091, 0x6fe649bf580efc27}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb28bb69e43175b52, 0x490d06ab793ed74b}, baseFieldElem{0x79aed252dbf870bf, 0x753a45af126c5573}},
			y: gfP2{baseFieldElem{0xd3007d44af2956a1, 0x76b90c26e1c4ca33}, baseFieldElem{0x953526a768fe2f91, 0x53d08fb6eb41b147}},
			t: gfP2{baseFieldElem{0xec397e6345a80cf5, 0x4816bea15894f79c}, baseFieldElem{0xf114c6d29ff48b68, 0x775b8df7b7062a0c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3dec2cd4e0bd7b0a, 0x2ec8e184feb52545}, baseFieldElem{0x1e2e825a0d865e45, 0x6aefcf336220638c}},
			y: gfP2{baseFieldElem{0xd5c50c1c23ccdcc2, 0x3630a9387fdc3549}, baseFieldElem{0x5cc31391091a68ec, 0x26ceb2742fb5d146}},
			t: gfP2{baseFieldElem{0x5cf7e710eb837dd4, 0x17ba01930bc8aeee}, baseFieldElem{0x72b8b21abb5e0e61, 0x29b55c0f9c2c9339}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xfc91fc9eb33c1f2c, 0x582c4a0f70463b2c}, baseFieldElem{0x3c9866ae1c9c6b30, 0x31aa2f3c7bf06a73}},
			y: gfP2{baseFieldElem{0x1d8c2d65dee6ce05, 0x6b9c8fac55d42a30}, baseFieldElem{0x7e303ea9e9ee54e4, 0x0661038b2a83c113}},
			t: gfP2{baseFieldElem{0xe6ee04bc95ed9d69, 0x6420aad834d2386e}, baseFieldElem{0x2174e9e8f6049b61, 0x68633124fe91c162}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd7d93ada3b1a9b83, 0x14a465d54be9ce3d}, baseFieldElem{0x88c6533a98ea66c0, 0x0af0319ee722bab5}},
			y: gfP2{baseFieldElem{0x07437a681eeca95c, 0x2896a2d2734c0218}, baseFieldElem{0x3d6d0b4894fa0adc, 0x63c8a7df21f2928d}},
			t: gfP2{baseFieldElem{0xd0cfa66fef0bf32e, 0x6c842c900562eb5b}, baseFieldElem{0x5ffcc726847c1959, 0x7f9b39765e8593fb}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe39fe495c4413102, 0x33bbd95370375c88}, baseFieldElem{0x9c4c80ad3ed62a95, 0x750580b05fcb5903}},
			y: gfP2{baseFieldElem{0x7a3548f2bc872e58, 0x6c3e8b311908712d}, baseFieldElem{0x21777fb2f6e8779f, 0x02f8bdaa52a94667}},
			t: gfP2{baseFieldElem{0x16d5bb273f12ceff, 0x79b28736a62a9a19}, baseFieldElem{0x38515e311e2c078e, 0x6a85c9015ad627c2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xfbf09a9aebb5baf5, 0x0be0580715ca7030}, baseFieldElem{0x8f6f3fccf9779e54, 0x461ff294353bed69}},
			y: gfP2{baseFieldElem{0x008e0430a5107ce5, 0x2e23dc58ab40aa4c}, baseFieldElem{0x30b5a9fe698260cb, 0x262fa9a66f6e466f}},
			t: gfP2{baseFieldElem{0xb2209da44f541ef6, 0x0068df46c4764917}, baseFieldElem{0xa98efe86d801db3f, 0x49c27212c353024d}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0f73cb90b78214c6, 0x2ee816269da811da}, baseFieldElem{0xdeb10c098d5fcbdc, 0x3398dfbd8124b4fd}},
			y: gfP2{baseFieldElem{0x6d5718095a0f64a2, 0x671fb2216f84a2f6}, baseFieldElem{0x7596e5a4d3862789, 0x01bc5ad8789171e0}},
			t: gfP2{baseFieldElem{0xabef77aa83c6119a, 0x0195e3b969456bff}, baseFieldElem{0x4d6087cb3c1d7929, 0x78be8b2017ad52a1}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xc1f8e15f669968bb, 0x6e451c3a420732b2}, baseFieldElem{0xf039eb068011bec5, 0x2827334e695d42be}},
			y: gfP2{baseFieldElem{0x7c291fa8914e92c9, 0x00c186017fb127db}, baseFieldElem{0xc3dfa21376d94f83, 0x3599e41371771d1b}},
			t: gfP2{baseFieldElem{0x90ef0a28531fec70, 0x46a895bc968fc276}, baseFieldElem{0xac9a0115f31b51bf, 0x47e2fb4e03abc1ce}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2676053bf6f4041d, 0x788fc22392bf19ea}, baseFieldElem{0x97c3783f2b017892, 0x70c9fe667a6d2acb}},
			y: gfP2{baseFieldElem{0xb181db293c53875b, 0x5edd9afe53b67d56}, baseFieldElem{0x396c875c08cb422f, 0x09ea479140b3616c}},
			t: gfP2{baseFieldElem{0xa2a1294c60f2ba47, 0x2efcafe38be9f316}, baseFieldElem{0x8ac5dd989b17e889, 0x35e9f3c3be131277}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xbe0e3cc81aeb118a, 0x7504833f0f1bfd32}, baseFieldElem{0xec0a28fb3f6a67b8, 0x750258f4794f1391}},
			y: gfP2{baseFieldElem{0x88eee436451dba1d, 0x0554095231f9d862}, baseFieldElem{0xa3a7aadfad89f7bf, 0x134f81cc1b989ca4}},
			t: gfP2{baseFieldElem{0x9b6ccd7e4fa83b54, 0x0ca1eaf29c6b18df}, baseFieldElem{0xb23241d8879ba5f0, 0x3c369bd1a65aff41}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe8aa9a6a6ef14bd8, 0x023836806025a5d1}, baseFieldElem{0x2ef9ce3239667207, 0x7f94a2964faab283}},
			y: gfP2{baseFieldElem{0x12657cbbdf07290e, 0x5e64f1874ee0cd35}, baseFieldElem{0x7a1882480004d492, 0x23be4e8dfd199e8a}},
			t: gfP2{baseFieldElem{0x807fbb18f6aa605e, 0x77f9bc016692e80f}, baseFieldElem{0x042bbde12ea852d6, 0x362670b631625370}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3cafa9e52667a296, 0x2dd52b79b66ffbcc}, baseFieldElem{0xcbb73229144f63b5, 0x1dcfcbfd3e7cb3ce}},
			y: gfP2{baseFieldElem{0x29e13c5e94d13fac, 0x6cac992e12a89f49}, baseFieldElem{0x449f9f3d89f7e5f8, 0x636826aa0312653d}},
			t: gfP2{baseFieldElem{0x1a37d4f5d0f5eeee, 0x250636046a600628}, baseFieldElem{0x7e7f93824adc43a1, 0x7a000fb589e7ac6e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9cab0bca223613b5, 0x2d2e8fd57a2dad43}, baseFieldElem{0xde8f330b8057c5ad, 0x5d831d634614e315}},
			y: gfP2{baseFieldElem{0x7755eb9b153aabc7, 0x1ed741d4ab02532e}, baseFieldElem{0x33cec259a6319478, 0x315daa4ebc36bbc9}},
			t: gfP2{baseFieldElem{0x397d1e708ae2b732, 0x56cf8a1fa402caac}, baseFieldElem{0x6246f38d3ab30868, 0x51b14fca334df50b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xaba1b04579d13c2f, 0x4a40f87d79119e84}, baseFieldElem{0xd5504a4bbbadd424, 0x3bcedb08b08c6c78}},
			y: gfP2{baseFieldElem{0x0e6dbd6e283aec46, 0x0da0e08cda1273cc}, baseFieldElem{0x2dc4f8f4224b1219, 0x77656c4832f14122}},
			t: gfP2{baseFieldElem{0x32cde83040ae7444, 0x0c4ba0d3d4f925bd}, baseFieldElem{0x01ac352d3c24ec65, 0x7ddbf230ae33c960}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x44b81b7f070054c4, 0x3b8bc63e7e351a0c}, baseFieldElem{0xd68a167a0a58931a, 0x585b7e8e735e6456}},
			y: gfP2{baseFieldElem{0xa2d570a32dff6aea, 0x5167a118d4cebaf7}, baseFieldElem{0xecddcf95338585cc, 0x58c3fb48ff816558}},
			t: gfP2{baseFieldElem{0x359936c40a2aa569, 0x7d0e46e10fc5a215}, baseFieldElem{0x425e4a2930eabdf2, 0x78d6f90a7df51367}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd67135f16c0d8685, 0x0cdfb783cbb496c5}, baseFieldElem{0xe605ddfd12a39021, 0x2c8bab04eba22576}},
			y: gfP2{baseFieldElem{0x00c01d5794f037d6, 0x416fe613e42c867e}, baseFieldElem{0x9bae34172277e3b5, 0x70bd25fbc16f8037}},
			t: gfP2{baseFieldElem{0x0678d5057aa37ee8, 0x5a50f44756d6d8e3}, baseFieldElem{0xdd626eb1cc3e1f69, 0x24cbc3941e4f5812}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd8300311ed6a59e5, 0x38d1277698408c76}, baseFieldElem{0xc5b1916588d50296, 0x09b3cf5e45841d4a}},
			y: gfP2{baseFieldElem{0x9715e9239a6e35c8, 0x01ff0f85fc186103}, baseFieldElem{0x6cd4ef7ccda7ef68, 0x4cbd4994ddc3955e}},
			t: gfP2{baseFieldElem{0xd092eb9143165905, 0x4e78cf4dac79c8ad}, baseFieldElem{0x2258c1f47a8de9bf, 0x055d2b75d457dc2e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xbb8a35d5ae23e9de, 0x178250e1bcc3e364}, baseFieldElem{0x3342457deb64da2c, 0x70a8798edaa2ad7e}},
			y: gfP2{baseFieldElem{0x9a3f4fa005c43c1f, 0x1200e17d191356b6}, baseFieldElem{0x22fa06d1c88f8cbb, 0x027296153890de32}},
			t: gfP2{baseFieldElem{0xa645f2e99442a3dc, 0x1799d6fda2eaf489}, baseFieldElem{0xf13eff11455359d2, 0x2d0dcfe0c51709db}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3d876227eba40a01, 0x675249cb479a18dd}, baseFieldElem{0xc212d8bc4907f50f, 0x3a7de38da167b7a0}},
			y: gfP2{baseFieldElem{0x4dede52099b678db, 0x513abc4acd491cd9}, baseFieldElem{0xb1dd156ae907d4a4, 0x1cdcd82c8413ebec}},
			t: gfP2{baseFieldElem{0x18ef999a2a7c22b2, 0x503f480c7aafcaf3}, baseFieldElem{0x8a561d40e6e4a7b6, 0x74adc69ad15a016e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3bdeff9fb192fd3b, 0x7ac491e40b5a2b5e}, baseFieldElem{0x9ce78334fc44ac70, 0x4d944e9861f86044}},
			y: gfP2{baseFieldElem{0xd592f10b9b3d22dd, 0x28034e3873f33969}, baseFieldElem{0xfa473dba42e98263, 0x14ab3529ff21e727}},
			t: gfP2{baseFieldElem{0x3f1f1cb11fc56f54, 0x30f868cfe49d68a0}, baseFieldElem{0x8a962f11bc800637, 0x6353f4652ffb515d}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x363e9e8fe345c8ec, 0x04a70482d6f92ff1}, baseFieldElem{0x833dd8ff095fb237, 0x07dfeb52bc72ab2f}},
			y: gfP2{baseFieldElem{0xf6f4c9a12bcc687e, 0x7d81473d1653d7b4}, baseFieldElem{0x72c475f90217b93d, 0x4a11b3c466f7e8b8}},
			t: gfP2{baseFieldElem{0x24a61051f613bbd4, 0x731a745fac37dcf1}, baseFieldElem{0xe6ce19e92c21090c, 0x592f6656e45644f5}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4f8539b6907e3950, 0x38c88d8471db2f46}, baseFieldElem{0xa58a458b2ddab33f, 0x48156e983e31a95f}},
			y: gfP2{baseFieldElem{0x4e01070327be6f96, 0x4a73cd4dd28b8ef7}, baseFieldElem{0xc9e2910285334a1a, 0x0a94fc83d15f4a96}},
			t: gfP2{baseFieldElem{0xfb3fa26fb503aa72, 0x7e73b142f56c3ad1}, baseFieldElem{0xce822075d7004a07, 0x1df0cadb00572b32}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd4d00572117954e8, 0x4c45284a5d4d01ee}, baseFieldElem{0x58bc731ab80555f6, 0x4adfde3a0bb85a33}},
			y: gfP2{baseFieldElem{0x1c7c8788f2ce274e, 0x5f45428be820c028}, baseFieldElem{0x73f42acd789d0369, 0x171771f9be882040}},
			t: gfP2{baseFieldElem{0x02369abdcfec4d2f, 0x14210c44d6b6abde}, baseFieldElem{0x805a6b8dbe852058, 0x1ca789eb23598820}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3ca1a05ee59979f8, 0x52abbbcc2f6f3d75}, baseFieldElem{0x2db0347d0af91f6d, 0x756d9cb8969d86bb}},
			y: gfP2{baseFieldElem{0x5eed7a9bd95524a6, 0x32540c230507c219}, baseFieldElem{0xbc1f1f518e3e91d8, 0x7e90a5ef6ab4bd8d}},
			t: gfP2{baseFieldElem{0x266e358f98489f6c, 0x23c873d4f84f603d}, baseFieldElem{0x195fdf6a1ee66e78, 0x6f8ecff2dc351dc3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1ecedc66c0572065, 0x472fb23d4ceaf313}, baseFieldElem{0x31653bbb9487f359, 0x6ba494b78ba4b963}},
			y: gfP2{baseFieldElem{0xd4a9c03ebbf86155, 0x2893f613ed718dad}, baseFieldElem{0x6989d5190357bba6, 0x7c0a69f50acfa52a}},
			t: gfP2{baseFieldElem{0xd03a7eaec0cf6c98, 0x38b5f557d5d43734}, baseFieldElem{0x8ff57387316d3f0a, 0x5505c8dfe1b2f1a4}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd1237b657c47a3aa, 0x16be50318109c68c}, baseFieldElem{0xfbb06dcad6a8a372, 0x2e15de49be74162c}},
			y: gfP2{baseFieldElem{0x0803bf619f48cb6a, 0x412fb5ca3c371829}, baseFieldElem{0xbbc8204ee3f5be79, 0x4b51d8925de9806c}},
			t: gfP2{baseFieldElem{0xb5d80e53b6598a19, 0x549af43d93ffa125}, baseFieldElem{0xcfb83acd3acde4da, 0x17e4b0ba3eeabc58}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x390f28516ca605dd, 0x03067cd24db62de1}, baseFieldElem{0x946e73a656630dda, 0x71b60aae2e3cf27d}},
			y: gfP2{baseFieldElem{0x711a16286194a405, 0x611090b84a34f5c0}, baseFieldElem{0x6fe099400b16c447, 0x4be677ef990e59a7}},
			t: gfP2{baseFieldElem{0x74595e9c9a33f927, 0x6401f63a14d79df6}, baseFieldElem{0x3ee0e64f84d0e559, 0x6b496d779a31862e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x477d27c4a717002a, 0x6b97e1e62f1ecca8}, baseFieldElem{0x04d5a828b1e52934, 0x3b49743d954f8926}},
			y: gfP2{baseFieldElem{0x4735586fb8bab4bc, 0x6a20f308285430d1}, baseFieldElem{0xc510924fcc47b9ad, 0x2d0dcd702859af5f}},
			t: gfP2{baseFieldElem{0xab39c36fa267fc4f, 0x50f3fd0f49ded7c3}, baseFieldElem{0x550dd420cc08caa8, 0x0ff3091a46561ded}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4707ccc042c3ef44, 0x36dee49450a033ee}, baseFieldElem{0xf2e6a8c9916651a9, 0x56a9cfbdc7d39e55}},
			y: gfP2{baseFieldElem{0x6fc649e11546eb57, 0x36f325fba1bb2aaa}, baseFieldElem{0x8a630bf34d667e8c, 0x6b7c0202379156b7}},
			t: gfP2{baseFieldElem{0x1ff5e011844a9ceb, 0x1ab0fd2ab7e2ae31}, baseFieldElem{0xffb180ebb6c2a514, 0x556b19ec505813f9}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x16eaf3d9caf21cef, 0x6afbf9852c85265e}, baseFieldElem{0x60166cd85a64796c, 0x652d08ae27da325b}},
			y: gfP2{baseFieldElem{0xa08051b2a41a16b7, 0x3d6a720296dd92fc}, baseFieldElem{0xd644789d4a5e022e, 0x51c22c62dd887ba1}},
			t: gfP2{baseFieldElem{0xf688fa81abc3f145, 0x6b3c9b577a8321bf}, baseFieldElem{0x52c66237c7df7a88, 0x4a47a2bedf6df8a4}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4856cfc2e6be911b, 0x381606837764f780}, baseFieldElem{0x022ae26d523212c4, 0x75c5dc9239017980}},
			y: gfP2{baseFieldElem{0x222855d0ecae3179, 0x6f4f98408d713675}, baseFieldElem{0x1af8a7bf9104d34e, 0x7062307227b8cbdc}},
			t: gfP2{baseFieldElem{0xa77814fc21389d76, 0x0078888574bc3951}, baseFieldElem{0x5465447d8df906ac, 0x4f5c7a7047214494}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3c93d7b35dafffbf, 0x7d41d493559eeb79}, baseFieldElem{0xdf0dac90f4ed7ddf, 0x6a4b408fe3e1eac7}},
			y: gfP2{baseFieldElem{0x393fde38b66bea05, 0x4e80621b0db050ae}, baseFieldElem{0x32854e0a96436ab4, 0x56de93ae392741f7}},
			t: gfP2{baseFieldElem{0x1ffdebd25d626288, 0x4da4cdf815e7ad34}, baseFieldElem{0x3015feb7e8e55372, 0x52448f62f75f4687}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0589155d7d5d78c1, 0x408e76c37965686c}, baseFieldElem{0xd76ec79cbfbbce92, 0x2f33674a7246e084}},
			y: gfP2{baseFieldElem{0x075cdeee686bf2bc, 0x069657aa6cdd0eb3}, baseFieldElem{0x6d4287f0cdae322d, 0x1b0ef7174981ed3b}},
			t: gfP2{baseFieldElem{0xd9a24f053a4cb61a, 0x3391f8104b773a9b}, baseFieldElem{0x5ae84339cf9adc69, 0x27a018d448a252d6}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xacca1cd1e52aeabb, 0x3c0f78708dcb26d5}, baseFieldElem{0x9849ee50d34f5173, 0x15053c5f59a0d2ff}},
			y: gfP2{baseFieldElem{0xc49a9b033fc0bbf4, 0x14c40ea5b21f25d2}, baseFieldElem{0x50c3f32c8f5f6d05, 0x55ae2d5acf1f6e0e}},
			t: gfP2{baseFieldElem{0x23bdf205ac884d3c, 0x64e5168db58d6a22}, baseFieldElem{0x90538a9bd7b591cf, 0x7f64696bd8db0973}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x07e2854c84186ffa, 0x66a4a01858cd437a}, baseFieldElem{0x49427f9d6989f800, 0x007e79b255f4941e}},
			y: gfP2{baseFieldElem{0x5a538d1c99ea4a5a, 0x251da7f7569a0e08}, baseFieldElem{0xe6531e450706dd9f, 0x3b68cf48e3cf8ba9}},
			t: gfP2{baseFieldElem{0xd90dbf3eaac43082, 0x723c6a87f9854135}, baseFieldElem{0x2419cbf3bf578cbd, 0x2a0cc7067638782f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xad6b99baeaf18007, 0x45bd1bc455acf4bb}, baseFieldElem{0x99e6231a3c7bd311, 0x68ba83f95e84f40c}},
			y: gfP2{baseFieldElem{0xdad194086c4ff20b, 0x68d7497ba1d93188}, baseFieldElem{0xb9bd57db0bc96adf, 0x1b8cd9ba647bc52e}},
			t: gfP2{baseFieldElem{0x2f258155c3ad8db8, 0x2dbe14e65ede3ef5}, baseFieldElem{0x77116e12ea323578, 0x344d55e3ae663ac1}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe5c0fbdc1213b5da, 0x064ca0dd08339e9b}, baseFieldElem{0xf552033ba9e4d521, 0x113b1c96c6bf9af6}},
			y: gfP2{baseFieldElem{0x9112145aa5a37257, 0x70187f4abe99b39a}, baseFieldElem{0xccb899f559ad8c21, 0x61b1d1688076e453}},
			t: gfP2{baseFieldElem{0xe9a237c844328e6b, 0x2344f75a3651b6fc}, baseFieldElem{0x074dd704f067c988, 0x67e1ebbb3020ccbb}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb8a688ef592b16d0, 0x421649daac52452f}, baseFieldElem{0x29e0116a532b6020, 0x5db46b7a5446fe8a}},
			y: gfP2{baseFieldElem{0xaad438b9c1b85dfa, 0x5a6fd8e8ae8a3cf4}, baseFieldElem{0xc1985ae653201d12, 0x5873b82ccee35a09}},
			t: gfP2{baseFieldElem{0x9ce07d980e4eb853, 0x3acf01feefbbf823}, baseFieldElem{0xf6d8fc88ffd2ae42, 0x20b450fd08695d85}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xbe7a3cebdfab58ed, 0x0acc1508ef6fc4e9}, baseFieldElem{0xf67e5d9aa1d186e9, 0x708be3525fa5d789}},
			y: gfP2{baseFieldElem{0xbc927ea1778e62b8, 0x7fc08bfb223ae827}, baseFieldElem{0xb5e4d1e46efd7d98, 0x156ccf0f33932c6a}},
			t: gfP2{baseFieldElem{0x75533dd83150e41e, 0x37db49399e77df3a}, baseFieldElem{0xe2c06cc7f0255eb9, 0x18504c87b0b1ca70}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x62b6feb3b4b2f027, 0x518f40ae9176f65b}, baseFieldElem{0x6985161757607329, 0x1dce6874f4145d54}},
			y: gfP2{baseFieldElem{0x5998013966101058, 0x2da26498b3118650}, baseFieldElem{0x7cd085383471de70, 0x7ad92650bcade487}},
			t: gfP2{baseFieldElem{0x4adb09f3d0558f56, 0x17a95c79a1380bbd}, baseFieldElem{0x1f44937b5815c738, 0x2833a5017e9d3be1}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb9ce529870e57b5d, 0x5efe3b34f45fae9c}, baseFieldElem{0xe38fda1ffc05e23d, 0x2434fc1ed990d252}},
			y: gfP2{baseFieldElem{0xefbd2b0e8769cd59, 0x3b12e6380b9bb87d}, baseFieldElem{0x4fe2d1200797f3e6, 0x50ed3d49556fba96}},
			t: gfP2{baseFieldElem{0x2771f87139da4f19, 0x3b3f54e9e4f3f505}, baseFieldElem{0xc6c01b2d1260bb7e, 0x4691633d8d6490fa}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa49bd8866509c375, 0x79b4f7ab85b8dbdf}, baseFieldElem{0x5cb0d3668e0b68ec, 0x7c6a06b1df0cc95a}},
			y: gfP2{baseFieldElem{0xfc72dd204283fac3, 0x2d1df0fd328ca1c4}, baseFieldElem{0xa5bbf87c870c6f19, 0x2dcb1099a2228cd1}},
			t: gfP2{baseFieldElem{0x3319214e190c79c8, 0x098b5310c039c8b1}, baseFieldElem{0x1605fe0b17ee0066, 0x398b2fa0f8723912}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x47955e50e1f28659, 0x2a9e0fd9c6e4b0c2}, baseFieldElem{0x598e4323917b4490, 0x5d6eaf2d702ec8e1}},
			y: gfP2{baseFieldElem{0xb77f26370324e468, 0x74486c7b11760f2d}, baseFieldElem{0x69c6d2232b33922f, 0x671a5d4d7df7c0ba}},
			t: gfP2{baseFieldElem{0xd744befcedb08bce, 0x77f5b198cb8b78cd}, baseFieldElem{0x39a0833988f38b3c, 0x396251cb973b9395}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x37544555edcd0458, 0x011312bef123016c}, baseFieldElem{0xfbd285b3a4454e0b, 0x3b6885951b9a40cf}},
			y: gfP2{baseFieldElem{0x33da202f3ee67853, 0x749d876f7db3083d}, baseFieldElem{0x7ef5c57cde1ad64a, 0x3b284620794e0ac4}},
			t: gfP2{baseFieldElem{0x0dc1d261e0c1760d, 0x6fc62b2cb5c73cab}, baseFieldElem{0xa11ec1b2b340815d, 0x387f19a613431baf}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2486e18912445d28, 0x598381cd79860bd3}, baseFieldElem{0x6d6127a16a644226, 0x45c7dd8c80ae000a}},
			y: gfP2{baseFieldElem{0x44032e5671542a33, 0x125418dfaa6298a9}, baseFieldElem{0x98f79bec4395ae3b, 0x100efcc4f003b0bf}},
			t: gfP2{baseFieldElem{0xa12aaafcea224037, 0x6b1019c51b18db2a}, baseFieldElem{0xa63616e535e2731b, 0x45c385e58a8dc5cf}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9f86ee93123f2bbe, 0x1d5a24a41ffb8052}, baseFieldElem{0x80087ee76d8e5d73, 0x23fc99d5a7a2bc96}},
			y: gfP2{baseFieldElem{0x0ceba56f2572c5bc, 0x25f3dee609d8c65a}, baseFieldElem{0x4a618a26934a5352, 0x7b1e134e587b8aa8}},
			t: gfP2{baseFieldElem{0x4e6781bf225b92fc, 0x23b8653d303668af}, baseFieldElem{0x431fdbeb166e4203, 0x2a207fba28946655}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xde5384e0debd1908, 0x36ed9b8471bbc1a8}, baseFieldElem{0xd1ce4810ef60184d, 0x5db9336c26b4c357}},
			y: gfP2{baseFieldElem{0xa5bb18a0f38166a0, 0x15c76ac6f4faf700}, baseFieldElem{0x62bb2d61d2544144, 0x0f251f8f26ca1d88}},
			t: gfP2{baseFieldElem{0xe6bc1ecff196e8c1, 0x1756a32aaa9fc77d}, baseFieldElem{0x8c9910c4dbe654c2, 0x26cbf44ac578fc13}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xeb4077d4407c62de, 0x07e722044f57a009}, baseFieldElem{0x91b042ae59891a3d, 0x2e4076e58988e533}},
			y: gfP2{baseFieldElem{0xf67e6d6a516066c9, 0x791947d8956faa88}, baseFieldElem{0x5dd353db1f216b1f, 0x3e980d9bfa2c5840}},
			t: gfP2{baseFieldElem{0xfe6635c475e8a572, 0x7add8eabc80782b0}, baseFieldElem{0xbefb5ec69e032a3c, 0x3d29bf7396f77312}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x445d44aa2f7f8134, 0x1b236a0a087cb27c}, baseFieldElem{0x890b26ced9d1b385, 0x5c1939e79c12e543}},
			y: gfP2{baseFieldElem{0x3b2c5e28247316d3, 0x4d37745bd5bcf1be}, baseFieldElem{0x70dd1aa64c646d68, 0x7527c5dcf5d1d7eb}},
			t: gfP2{baseFieldElem{0x4258361c15636541, 0x1e547b67cb2c181b}, baseFieldElem{0x1eaee787a272bd5c, 0x21a5c112c633c5c2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x03ffa6994e85a1fb, 0x1e0b338064b12438}, baseFieldElem{0x0a1abc5dd99b8358, 0x78c20fab3d6d2768}},
			y: gfP2{baseFieldElem{0x43eba9128e40f7fc, 0x303aa7e2d015c95e}, baseFieldElem{0x2bebb2cdd1d6d116, 0x0c3270ba1bd1a684}},
			t: gfP2{baseFieldElem{0x06b28df83c21b4ea, 0x7cf55afed5b5eeb8}, baseFieldElem{0x4602c25990872f30, 0x314d4dd85c91a9ff}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd3d1cb4f417bdd84, 0x41c70790324ae9e6}, baseFieldElem{0xfb248cf87cbb3dae, 0x0d4b4c4864b289a3}},
			y: gfP2{baseFieldElem{0x30537232cd58533a, 0x79be422953e61bcf}, baseFieldElem{0xdb8ab0ef8b0be930, 0x656e6946dff5b49c}},
			t: gfP2{baseFieldElem{0x7133387c19ad9017, 0x7930d7422c31fd4c}, baseFieldElem{0xf4645e7b92e5cb42, 0x1cdd10ee7f04ba29}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1b857b17aab6318e, 0x2e337e89286a7a9a}, baseFieldElem{0xe2578ae2277d1419, 0x227394026248392a}},
			y: gfP2{baseFieldElem{0x9e7af81bcdaf9806, 0x350b62c528525b7b}, baseFieldElem{0x15e897ac9ca7fc7c, 0x2748027ec850b763}},
			t: gfP2{baseFieldElem{0xf33b12ec6102a95f, 0x6cac5c1d4f8f92c5}, baseFieldElem{0x35c25cbd7aa8f01f, 0x1c51c01b4d5304ca}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x820df01bb0a43a7a, 0x0db54e68c9804359}, baseFieldElem{0x693a580ce398a200, 0x79a5619a5a5cd60e}},
			y: gfP2{baseFieldElem{0xc1f2df48ff7b11a6, 0x2a97da5284961ebb}, baseFieldElem{0xa6ad597cf812e0b5, 0x116c81f6d5fa8d50}},
			t: gfP2{baseFieldElem{0x221ca7f08232b8fa, 0x2d284fa778c38b90}, baseFieldElem{0x9d190a22ded1a443, 0x428c04b1f80b6772}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb424fd7b817a0815, 0x350cdc3e4ead138c}, baseFieldElem{0xcba78e63cef3d876, 0x0bf11cf3c107ec8f}},
			y: gfP2{baseFieldElem{0x9d6ca2d0c941c8ae, 0x596957229e1b8e10}, baseFieldElem{0x334df55db5b29b4a, 0x69d1d616bb6c4758}},
			t: gfP2{baseFieldElem{0x1a18ed76ebbd82cc, 0x71f15a37b7a07c78}, baseFieldElem{0xf6a2fa01073bfd17, 0x58cdd4ed63c10615}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2688e1c7f142b87e, 0x2570f676097c0344}, baseFieldElem{0xef223afa23c96bcb, 0x4e4b7afa9dcf7003}},
			y: gfP2{baseFieldElem{0x82645c613f1cfd1c, 0x5b94095d88794c18}, baseFieldElem{0x0557b21a893e0673, 0x74a56db02ae3c1f8}},
			t: gfP2{baseFieldElem{0x368831b20daa5dbd, 0x5ba77fc9e855b5aa}, baseFieldElem{0x97f7d7ce69b742e6, 0x6894877721e73860}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2e0d2a3289de5f05, 0x1c8fb1603f9c2739}, baseFieldElem{0xcf3a551470e35e39, 0x07696f50761d006c}},
			y: gfP2{baseFieldElem{0x84ab9c71e715abfb, 0x688db45d789a0929}, baseFieldElem{0x7daeaf0035f2f8db, 0x16e79c01e6fcde9f}},
			t: gfP2{baseFieldElem{0x96bee43e8c894fbe, 0x2d20fe4cadc58aed}, baseFieldElem{0xda0aaedbb4ff0758, 0x58335630976c485e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x117a51101ba6644e, 0x0861d509303b9618}, baseFieldElem{0xa8fff040241bbe30, 0x6732a7018e7c1b6e}},
			y: gfP2{baseFieldElem{0x898a36c24ef61531, 0x579b44f911380225}, baseFieldElem{0xf3523a635851488a, 0x58fa86da29bb3056}},
			t: gfP2{baseFieldElem{0x576370ae9a72795e, 0x14b836173d201c0e}, baseFieldElem{0x656c1b9b5abfe350, 0x33e5b767436c8b83}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xae2b402d94bfd3f7, 0x7a1192c69c37e4c6}, baseFieldElem{0xc8ea7d1041bc99fd, 0x61924f2907256216}},
			y: gfP2{baseFieldElem{0xa9a7d38643bc59a0, 0x553199561e112e78}, baseFieldElem{0x9879540f6302620e, 0x3926fee928b88f17}},
			t: gfP2{baseFieldElem{0xb339f9000817b96f, 0x41cb2ff335f2662d}, baseFieldElem{0xf3e86c5c9ded9bdf, 0x40644956ee8cde70}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x700291a21b6e5a10, 0x33d0dc0be56b508b}, baseFieldElem{0x1357c0db988f8933, 0x74841ef999ecd923}},
			y: gfP2{baseFieldElem{0xcf7294b38119fe0c, 0x52fec3a8dc35938f}, baseFieldElem{0x70ba08b551fe7eb6, 0x11603c5363711433}},
			t: gfP2{baseFieldElem{0xcaddf9069927f55f, 0x37010be5ba6291ff}, baseFieldElem{0x0036551c42939cf9, 0x6e5025a44a61eb76}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xc3fb6399605c45ac, 0x1a6ed5a6e118f669}, baseFieldElem{0xa05dbd739dfcbfa9, 0x1128ab682d7983fc}},
			y: gfP2{baseFieldElem{0x90f0da83052b4748, 0x65b93bcf5a8929a4}, baseFieldElem{0x97cfd3fcb3eabbc8, 0x38b76154d3c11a82}},
			t: gfP2{baseFieldElem{0xc6abc1a6c0f5dae7, 0x4c53b492be6bc3a2}, baseFieldElem{0x18748755382b1f81, 0x05a5be9b41fc9f95}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb28b182838c92e77, 0x7cc2434889cdcd7e}, baseFieldElem{0xb3586f72db2ccdd3, 0x1b0c20515a182172}},
			y: gfP2{baseFieldElem{0x04c4cc91300a6c2d, 0x3c6ba02c2ec6158b}, baseFieldElem{0x13e6f63c024eae4a, 0x50d3c7bb7c6d51e7}},
			t: gfP2{baseFieldElem{0x2c24b9c0bcd56578, 0x656bc7b824680152}, baseFieldElem{0xe04b4be479098208, 0x0606af82466f29c8}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7f721467e5c39c0c, 0x5b87a58df44e6854}, baseFieldElem{0x735de27d520de60f, 0x6680dae7bc9ee4c9}},
			y: gfP2{baseFieldElem{0xa7753d96b9f7f8f5, 0x4dfab5e3bc14c640}, baseFieldElem{0xae08ad199d7dcfb5, 0x48a9aeb6972abb3a}},
			t: gfP2{baseFieldElem{0x0e36f15672ee02a8, 0x42e4d4c66ec81a01}, baseFieldElem{0x82b42c91298bd18e, 0x36ac3da97e43a738}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1f5d883dad117063, 0x6fa08012003ee43f}, baseFieldElem{0xa9f695b403155424, 0x6e9fe8c1e1649606}},
			y: gfP2{baseFieldElem{0xcc160e1dc470b889, 0x0c1374da6df3ade9}, baseFieldElem{0x61fe023110018555, 0x6c527978bf9bbcb8}},
			t: gfP2{baseFieldElem{0xaca863ed10da9aee, 0x116ffcd49f6bbe10}, baseFieldElem{0x587398e101500ccd, 0x37840fc8c9348257}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x85f10c49d0f5f1a8, 0x5188f1df6850e706}, baseFieldElem{0xec0fc7fd56af6b48, 0x2fef1a2faee6f528}},
			y: gfP2{baseFieldElem{0xc896c3c03af372d7, 0x38349b37a358f4a4}, baseFieldElem{0xda65a797281de309, 0x7caf268c0de61a32}},
			t: gfP2{baseFieldElem{0x7ed5a173bd2419c5, 0x715296f3714ba22d}, baseFieldElem{0x15b28d6ed519d849, 0x18e46eaa2ec77105}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xbac4d67c4264b934, 0x1e667b60c930ebaa}, baseFieldElem{0x72476bb4a5ae332e, 0x42c82345f8a9ca0b}},
			y: gfP2{baseFieldElem{0x86c9a215eed344a8, 0x393689b59e6bbfea}, baseFieldElem{0xdd94bc3e2f23e50e, 0x115792abddaddc1e}},
			t: gfP2{baseFieldElem{0xd7e3067dd89fa396, 0x35cb6999737eb0b2}, baseFieldElem{0x86b197d1b91e8796, 0x140686f88c0132b4}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9735cf210ba5fe01, 0x4e4f9c3ca6f05fb2}, baseFieldElem{0x1e3d166b703dd3bb, 0x4701e6100a571beb}},
			y: gfP2{baseFieldElem{0xdec35e519cf560f0, 0x574dfd4ff9094c5b}, baseFieldElem{0x6bdfe50736866005, 0x5cf3d10d3f7f5a19}},
			t: gfP2{baseFieldElem{0x17c574e83472dd8f, 0x16e781086e51604a}, baseFieldElem{0x5a97c01056d08fa4, 0x672a8a1450fedb1f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xc3a80332c7ac6959, 0x742fa286a319b01b}, baseFieldElem{0x6d549b288c3ef93f, 0x5696a0aedc043438}},
			y: gfP2{baseFieldElem{0x77439c43ca3b3b79, 0x4a6bc4f962fa09ca}, baseFieldElem{0x4aa4ec522219e972, 0x1295db20cf469e99}},
			t: gfP2{baseFieldElem{0x811d4a33a9bc624a, 0x1cc2ad04415bc752}, baseFieldElem{0xe89f6155a846c2df, 0x75d74bba69c72a75}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa07ade0ca50284ee, 0x394b147150868403}, baseFieldElem{0x2a57ea0e98a37225, 0x6446cef3e5ca8d99}},
			y: gfP2{baseFieldElem{0x169a3a9b89bd3af5, 0x6b24f20073e839d9}, baseFieldElem{0x449d434a817ead4f, 0x634d9259d0dbeff9}},
			t: gfP2{baseFieldElem{0xd54682ea90306cb1, 0x7b1e57b1af2fd18e}, baseFieldElem{0xd2379f29daf4a2ad, 0x31b4899f71475b81}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x57767528d61cd681, 0x2f114a77e837b09b}, baseFieldElem{0x10f79b3b194328cc, 0x348ec927bb269ed5}},
			y: gfP2{baseFieldElem{0x148119ef54c2bb4b, 0x0dc7d415aaa6649c}, baseFieldElem{0x2a9fbf88e227a276, 0x6192abfd39656e91}},
			t: gfP2{baseFieldElem{0x1664da9a5996a304, 0x6165cc9b3356bb26}, baseFieldElem{0xcec3f1a57e96e1d7, 0x7fb2274e8e138dc4}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd3b80bffbdf54a7f, 0x607560679ddbce83}, baseFieldElem{0xc7704f33b4d17612, 0x48da924dc142fabe}},
			y: gfP2{baseFieldElem{0x9b154ec297ea6d17, 0x2c48b473bfbeba57}, baseFieldElem{0xfa581d2b45afdf15, 0x6388cee5f916ddc8}},
			t: gfP2{baseFieldElem{0x5f1607eea094ec1b, 0x5fbb4e82a43cc7e7}, baseFieldElem{0x207dc9a5fc58e6a3, 0x4709529903357769}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x80aca2e0893e7085, 0x37e21bb18acd43df}, baseFieldElem{0x594aaa6e79180b74, 0x669194365e77a39d}},
			y: gfP2{baseFieldElem{0x5c273557de119f30, 0x7c3541fbf1ce84d8}, baseFieldElem{0xadf9ee962dfcb7d8, 0x34d78aeba73abdad}},
			t: gfP2{baseFieldElem{0xc7e1d58b0e188deb, 0x301c50020dfaa089}, baseFieldElem{0xbffc5e0b7bda5cce, 0x41cb69810c0edeb3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x976ebff43cd9a9ad, 0x53fbb14bb28a7a81}, baseFieldElem{0x51d611e874b6d507, 0x7ea0ed86f0d862b0}},
			y: gfP2{baseFieldElem{0x14c1183aedebc76c, 0x2b9051236a8db451}, baseFieldElem{0x98f5fba73e6d4109, 0x5ac7eddee1fef1aa}},
			t: gfP2{baseFieldElem{0x3161d7b2cce019c6, 0x66fb1ef257416587}, baseFieldElem{0x250a42604b9a0f35, 0x4bc14b5df87b662f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7086ce2ab48c0ed5, 0x5f9f6b4fe83118da}, baseFieldElem{0xa3580a2845f234be, 0x5cba5344dd5790e3}},
			y: gfP2{baseFieldElem{0xcab39b5cf8b228e4, 0x419e63c28570171a}, baseFieldElem{0x4ba5b61a789e4dad, 0x59628d3ffa365766}},
			t: gfP2{baseFieldElem{0x40419f7c631ae04b, 0x1159dee5586f977d}, baseFieldElem{0xf615849b502ec81f, 0x7bb4cafe4585c388}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb1ddb3b9c80c45d4, 0x6394a36a2dda0876}, baseFieldElem{0x0dd12ea0a0a18b2e, 0x75d6d75be4e4a224}},
			y: gfP2{baseFieldElem{0x67db16dca43df252, 0x374f9843cf47ddc7}, baseFieldElem{0x102bd18af505213b, 0x51320b2388b44773}},
			t: gfP2{baseFieldElem{0x2b9666de53ca9b26, 0x52384ecc2eb1c865}, baseFieldElem{0x200983c310b527b2, 0x7d70e4762aa454e9}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb57cc002f2796eec, 0x3eda63ed1e8ae356}, baseFieldElem{0xc89939810d379e0b, 0x06f27e5d9c1a8fb8}},
			y: gfP2{baseFieldElem{0x1594f0e3294e8a7f, 0x2e7b5050c4a68999}, baseFieldElem{0xe665f7331614ca86, 0x392de160b6b20614}},
			t: gfP2{baseFieldElem{0xc689526b1cabac50, 0x3fa4177f56a788cc}, baseFieldElem{0x557c50101d564c93, 0x401a96c696a423a6}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd7a07cf298f7fa81, 0x29e6c11d1501dd16}, baseFieldElem{0x65d5a25de02af55a, 0x3310157014d1c1bb}},
			y: gfP2{baseFieldElem{0xa642b4f85b0bdd45, 0x03d8798282586df2}, baseFieldElem{0xb184e867f0dd2cf2, 0x3387989ec0a9efce}},
			t: gfP2{baseFieldElem{0x4b4639ae7bd83aaa, 0x660afb6fd795f002}, baseFieldElem{0xe8eee2d1e7be6bb5, 0x6cb92d2a54167604}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x017aea40ac386f1d, 0x1167b185a53e593d}, baseFieldElem{0x5e24455bd986d72e, 0x150b4f744836669f}},
			y: gfP2{baseFieldElem{0x8761bc7a3b6cd75f, 0x7afe4b7b821db151}, baseFieldElem{0xa17b05b8ee0e25db, 0x3e98a0aef837a4cf}},
			t: gfP2{baseFieldElem{0xd9f671ce12eaefdd, 0x1a7b1fa7bbd09792}, baseFieldElem{0x20204ed00ec25171, 0x23c09368eebc01f6}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd135f1250909995f, 0x1299ebea1682111a}, baseFieldElem{0x3c9f4f7eb1e9a1af, 0x447e6749e9a6700f}},
			y: gfP2{baseFieldElem{0xf3e29ab18d0f4efd, 0x16d5a64384520ea8}, baseFieldElem{0xf106500d4d714be0, 0x61e4650277480ee2}},
			t: gfP2{baseFieldElem{0x1204d59e20a37d0f, 0x292e83f1a04ef442}, baseFieldElem{0x0330d1c14de56fe8, 0x0531294e9c74a685}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x48a4a70f6a4b0ca4, 0x6a7d76504edbd12e}, baseFieldElem{0xf6b4d11d9eca6074, 0x3b7cc39c80898eb8}},
			y: gfP2{baseFieldElem{0x869acef5b4f2ed9d, 0x4793b4b36b865f5d}, baseFieldElem{0xd00b0dda0b609557, 0x2352c700e8d4ceb9}},
			t: gfP2{baseFieldElem{0x35671da8541fe967, 0x6c9fa431d8f376f9}, baseFieldElem{0x81d9375abd4e9a06, 0x11454208102ce0cf}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8454c377239e827e, 0x50b911111a80bacf}, baseFieldElem{0x53f04fedc85a2810, 0x1bc35de6a216a551}},
			y: gfP2{baseFieldElem{0xcf565e6d16976adc, 0x63668bc5cbcdb1b4}, baseFieldElem{0x1b0897f93f67193f, 0x179da22b2fe29d92}},
			t: gfP2{baseFieldElem{0xb28e62ef55cbd1cc, 0x53a4b2fd74815ae4}, baseFieldElem{0x5e60e453755129a1, 0x369eb5368c22f10f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x80b928499ec2b80e, 0x11b18d95f59a7b2e}, baseFieldElem{0xea3ce970df4d521a, 0x4c61009362a237d6}},
			y: gfP2{baseFieldElem{0x724f6b3cc45ea9f9, 0x764a7031c7b65290}, baseFieldElem{0xba5cd524032acd47, 0x41ef7466e42682b8}},
			t: gfP2{baseFieldElem{0xb7ab72cc1808f702, 0x56a754d870301523}, baseFieldElem{0x95d095a66a59cd39, 0x79a0621d9ac7d1c4}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb82fe55f6bdd1fa1, 0x7c311c7ab97c7e65}, baseFieldElem{0x7c6940f281216fd6, 0x34db3166c3e07ef4}},
			y: gfP2{baseFieldElem{0xf7738cc4cdb856be, 0x0ababe76226ef72d}, baseFieldElem{0x6218188c87a23eaa, 0x7ede5f14f7b655d0}},
			t: gfP2{baseFieldElem{0x7db41fc2ad32c8ed, 0x67377d072832d06d}, baseFieldElem{0x12c471fa02dad066, 0x29b05a85452b4387}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x17b2fd30387d681c, 0x3a159f65ea85330c}, baseFieldElem{0x942938829cf80490, 0x464773bf203d5d03}},
			y: gfP2{baseFieldElem{0xd0ec85768ad235ce, 0x48e9bf7f8fefd4bf}, baseFieldElem{0x89ae149d1c81c4c8, 0x6b0d989b042b2c14}},
			t: gfP2{baseFieldElem{0x52ae212cfce340cd, 0x7acdabba70368c80}, baseFieldElem{0xa8ed910ef147529a, 0x615e8fe3d23981a0}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xaeeb4e74d548087a, 0x33aae9c7de91acb7}, baseFieldElem{0x4bdcde6bcfa080c7, 0x618b9eb507735a8b}},
			y: gfP2{baseFieldElem{0x20b5818b67693eab, 0x41ff9b7eab8c883c}, baseFieldElem{0x1d18dd65dc53d22b, 0x430ef176772cd002}},
			t: gfP2{baseFieldElem{0xcb404eb23274aaab, 0x595c7f517a6196b9}, baseFieldElem{0x7f5f2d546fe5d904, 0x1e21ec5cf1c7ba16}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xefd23ac8e6bc4c49, 0x31f0fcdba775237e}, baseFieldElem{0xb0f332c25d18ee50, 0x177927d04fc9c2ed}},
			y: gfP2{baseFieldElem{0x2576308e55e78724, 0x3c187355541769df}, baseFieldElem{0x00778d73b8010306, 0x1e7f2654c3e9d9e6}},
			t: gfP2{baseFieldElem{0x47c65ce9f768756b, 0x78ca406cf5fa7290}, baseFieldElem{0xf3d8450b358897c2, 0x742ada64b6f7d713}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8d591890efb417d7, 0x683e3ce426ef7d76}, baseFieldElem{0x8258979b471c0f15, 0x57a2baa019783339}},
			y: gfP2{baseFieldElem{0x93751cd523cf9171, 0x66890daee042bfaa}, baseFieldElem{0xfd6240688fab7537, 0x5381190539eb8347}},
			t: gfP2{baseFieldElem{0x68f642a4b585d005, 0x62b04485a0f13a11}, baseFieldElem{0x0276542feea37d56, 0x35a9aa1780fb7d2c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1f7b398796acd5f2, 0x65ff2d4b3d976b41}, baseFieldElem{0xa47ddf2d812a0ba8, 0x67b4f102f6d338e1}},
			y: gfP2{baseFieldElem{0x526bce1ef919f79a, 0x5aab1399603b86ed}, baseFieldElem{0xd77f4416ed7091c1, 0x11c7419d2db7f7c9}},
			t: gfP2{baseFieldElem{0xcc35564ce26c9e91, 0x1f9ee51ca3836c88}, baseFieldElem{0xe8fa18bb94893a40, 0x0fb6db0b39e6f660}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe139a2ba7dd07a8b, 0x6677a0389cd3fd13}, baseFieldElem{0x7fc17dd1a2c896b5, 0x3fe403203d1de8ea}},
			y: gfP2{baseFieldElem{0xf79ebbaea33fffd4, 0x3e6db493f5d3530d}, baseFieldElem{0xeb9e682b19fcce2e, 0x4b79e5baef558aeb}},
			t: gfP2{baseFieldElem{0x1b40d13243603506, 0x1b2c7dd405546f61}, baseFieldElem{0xee0ace7552bd5030, 0x7af0b98bf8b13a80}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xabd64a18cd33c04d, 0x5826e2e5ddf93e54}, baseFieldElem{0xf8d4843dc726dbee, 0x6316dd14d50f3aa6}},
			y: gfP2{baseFieldElem{0x505d654c5fae859d, 0x73459e6b3187faaa}, baseFieldElem{0x9ac25ad719d7fb4f, 0x1e91b7dffadbbcf9}},
			t: gfP2{baseFieldElem{0xad8066dacc3b8d38, 0x036aefefb7778291}, baseFieldElem{0xc937b790f724016e, 0x6e259e4f909a234d}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xed142b1ee04e3d38, 0x20e59fb25e46ff1c}, baseFieldElem{0xa892673172f76697, 0x0a0efb5b96432a93}},
			y: gfP2{baseFieldElem{0xfe415ebef78b6c58, 0x3a2441f3fd5c5f93}, baseFieldElem{0xb9646275d5208205, 0x4413fb4cb5dfb512}},
			t: gfP2{baseFieldElem{0x4fac76136aeb73f9, 0x74eed0031b09eed5}, baseFieldElem{0x60350dda7d7c82a4, 0x5beb464c27567ef2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb77acdefdafa93d4, 0x72bb43bead744a66}, baseFieldElem{0xe046492192770114, 0x7586c5ba3fe84482}},
			y: gfP2{baseFieldElem{0x09175ef74d1597c1, 0x0d39160d451f9ec2}, baseFieldElem{0x0c1f0ac2bed80401, 0x2fbb52f3e74b42fa}},
			t: gfP2{baseFieldElem{0xa1af22106f0c236b, 0x52635e2e19c02781}, baseFieldElem{0x25b719a8e033ab2f, 0x6a85040b8bac72c0}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xbe3645919e4ae891, 0x58b04b1ef446aa76}, baseFieldElem{0x74e2eafdf8139112, 0x2b9dfcc28399ec5c}},
			y: gfP2{baseFieldElem{0xfc4473fcedb03bf4, 0x0548084529d0558d}, baseFieldElem{0xea65313cdf9fb45a, 0x1c87b8af1505b912}},
			t: gfP2{baseFieldElem{0xcd59d522c29171ee, 0x07f3fade02d0a444}, baseFieldElem{0x237bbcc26b93c68c, 0x22b31d285908c500}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xf8dbf86ca0ac4635, 0x529e141060f705f3}, baseFieldElem{0xc9af353a2a99605e, 0x4a304385054c8939}},
			y: gfP2{baseFieldElem{0x64fe650883185856, 0x2d9acb6a1d57b5fe}, baseFieldElem{0xf76cae0ca9eb1f2f, 0x3720c378de291f20}},
			t: gfP2{baseFieldElem{0x95b8a0e3f1754ede, 0x4ac15f869a86a276}, baseFieldElem{0x3ad01e1b4785c6ec, 0x7d226394b55326b2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2333e963fa84b9b6, 0x6645a808bbecc536}, baseFieldElem{0x9d6cf2bcba5f1f43, 0x25a789487f1c26ba}},
			y: gfP2{baseFieldElem{0x7e686b44e16aeb9d, 0x0e387eb7cbc2b806}, baseFieldElem{0xbd609dc1e9503587, 0x278a6644761ddad3}},
			t: gfP2{baseFieldElem{0x3923b155e7de59cc, 0x06adff10ac0f3e5a}, baseFieldElem{0x89f391fe3f82cdd1, 0x79d502936f6ad6ce}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9db2355ddaaf009a, 0x299c26ccf48f286e}, baseFieldElem{0x5c16bdbbc151b314, 0x5c3d565fcc30c61a}},
			y: gfP2{baseFieldElem{0x059a7493a27a5b22, 0x071b2ccefac9b10f}, baseFieldElem{0xfcc5231bed1a0d5e, 0x612f9cebfa36b890}},
			t: gfP2{baseFieldElem{0x6d2b7a24e48eb267, 0x4ef4f3ca647c2b50}, baseFieldElem{0x699aabd5dfb19288, 0x4ff71f091bfa987b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x10fa36057631e928, 0x6945eef25e381344}, baseFieldElem{0xe6af1d3c3bf8aa55, 0x5c88f32324b15cd8}},
			y: gfP2{baseFieldElem{0xedd874fd6bdec0d6, 0x556e611bf4fd6b8d}, baseFieldElem{0x8a3d5a755586fa1f, 0x07f09c264835d8b9}},
			t: gfP2{baseFieldElem{0x9e56427d843e0a37, 0x6f002b4b27834bdf}, baseFieldElem{0xe694ab0992ddb17d, 0x0a01e4f0a94eba7e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb847daf99aa0be99, 0x5a743afa5ef8be50}, baseFieldElem{0x30cabc936303647b, 0x514879ecd3a31981}},
			y: gfP2{baseFieldElem{0x4b4cd7e128ae67e5, 0x36eb52554b1b0145}, baseFieldElem{0xadc81fa0274bd623, 0x3ba0bb8c36e77171}},
			t: gfP2{baseFieldElem{0xbb826d9fddf6734f, 0x59abeb7d9cfe05d6}, baseFieldElem{0xc9f1d0f938506a41, 0x709ad70f8aed9f86}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa6b81b9e6bf5975d, 0x77df67f7c7b640ac}, baseFieldElem{0xfd7e1704f8f1344e, 0x38f1c8f7f3253ae5}},
			y: gfP2{baseFieldElem{0x7134ddd42dfb8a1f, 0x51f87d4190b06cc5}, baseFieldElem{0x73a8fcee790ecb0b, 0x74b5cdebd30d053b}},
			t: gfP2{baseFieldElem{0x2632dd72c910481c, 0x61b681d582820cc8}, baseFieldElem{0xcf05eead49006d10, 0x3bbdace10f9ec6f5}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x13eb037a52c98df4, 0x53a91bc09cef7140}, baseFieldElem{0x19d3440ae7e1a747, 0x37447c99549334a4}},
			y: gfP2{baseFieldElem{0x04eade8db3c4b52e, 0x488c12f5a4a138a8}, baseFieldElem{0x72f586da660a9b85, 0x2d57565fcb02b47c}},
			t: gfP2{baseFieldElem{0x6b24d5f43b1af706, 0x54ec590249a9c0c9}, baseFieldElem{0x3a6355d27b8b3440, 0x58e0d2380531952a}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x83d2686c9a7c02c7, 0x23e0d09626c4c54f}, baseFieldElem{0xd7205996c02201f7, 0x00e79ec234ed8cf9}},
			y: gfP2{baseFieldElem{0x8519ea988d5d559d, 0x7d5e79d8c576955c}, baseFieldElem{0x2ad1e05ba4a499b4, 0x78ebf284d4781101}},
			t: gfP2{baseFieldElem{0x70203f68bd4b602c, 0x5900b7d1f07fc6a7}, baseFieldElem{0x5087e4a434cd12bc, 0x667b0f9e3dcbe7cd}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x119614412d778cf3, 0x555dd71e0a33d0d4}, baseFieldElem{0xb8b0ad64bd3bb454, 0x0b250980fe28038e}},
			y: gfP2{baseFieldElem{0x389a3f174772ec92, 0x786493a6c2cde789}, baseFieldElem{0x8a06173d9ed5fc81, 0x64b50d86050a36f2}},
			t: gfP2{baseFieldElem{0xf98b89fe6248249a, 0x0fa90d46391b7c19}, baseFieldElem{0x5a1b4c2f7fd0b363, 0x1e5cacc81f0995f1}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x23092f2afcd995ce, 0x38a752085270b864}, baseFieldElem{0x560388912562730f, 0x4c4c1420bf76b086}},
			y: gfP2{baseFieldElem{0x6bcefd719199f47b, 0x336308deb1747fed}, baseFieldElem{0x8f7c82815fdaf206, 0x783953d89f9699e8}},
			t: gfP2{baseFieldElem{0x04291c2322740b45, 0x41a96042c216a06f}, baseFieldElem{0x038f0c1a5d5a1978, 0x0a54cb25424f9669}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9cbe258c80fa4a0a, 0x2afe08c086466146}, baseFieldElem{0xa1fd56b31ec2856f, 0x2c68b10ac8a8bc55}},
			y: gfP2{baseFieldElem{0xb23c89d2662b5bf3, 0x4aa84b14d8edca61}, baseFieldElem{0x9d395f764ea40bef, 0x6520f59c7a344e3e}},
			t: gfP2{baseFieldElem{0xf3419a510bc98606, 0x1662d14c70298946}, baseFieldElem{0x2abb05b7a183f329, 0x0c76f499c96dd121}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x698f14ff25f1f634, 0x415c59529da9921d}, baseFieldElem{0xaae02fac36d09f19, 0x08d8587cc99ec30c}},
			y: gfP2{baseFieldElem{0xc16200cede168461, 0x525fcf5c352470b1}, baseFieldElem{0xd5281f277fe01ddd, 0x676b77fbe2d00119}},
			t: gfP2{baseFieldElem{0x6e46161c1361d4c5, 0x70f3bc155e3028d1}, baseFieldElem{0xfb1ec56079dacdc4, 0x7efe385a209ad566}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xdefb722b7a5492ec, 0x63978ce4df353f80}, baseFieldElem{0xa88a806e1b0b037f, 0x2c62bfd473a6eaad}},
			y: gfP2{baseFieldElem{0xf705e13de17256d3, 0x715bec7cae3b5497}, baseFieldElem{0x63c44cca0a8e5ced, 0x725dc649479cdfc1}},
			t: gfP2{baseFieldElem{0xa323475ec2a77616, 0x2c2efde56b310dfe}, baseFieldElem{0xc104db425443215f, 0x5604ccf37bc32ace}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x04dd960949f669db, 0x525dea761e5d2036}, baseFieldElem{0x280b328cce0d8a75, 0x55b8a89b29ba4989}},
			y: gfP2{baseFieldElem{0xce7f50182b4fb06d, 0x228e876a2f23635d}, baseFieldElem{0x572738e52790611c, 0x540d2a849d926bf3}},
			t: gfP2{baseFieldElem{0xa425d085b0d19384, 0x2adf1638ba0bac54}, baseFieldElem{0x108bc1b98697200c, 0x543ed3cafdacb66c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xd11730aee413ea11, 0x36ba8ae1d0b11d71}, baseFieldElem{0x30a0adaa051b6871, 0x12ef26791b13a441}},
			y: gfP2{baseFieldElem{0xd37b32195bc6b5b9, 0x7891fbe026ec6eb1}, baseFieldElem{0x9a3a75fa8365d5b3, 0x45ee0bd0257e9363}},
			t: gfP2{baseFieldElem{0x1decca9e2257c81a, 0x02787e85628b1c69}, baseFieldElem{0x0294f24633fadc15, 0x56eb7f9d6ebd0ee3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe5f095027904729c, 0x0dc26ab499925b91}, baseFieldElem{0x7f9f95004142a73d, 0x501e84f0e13f24a5}},
			y: gfP2{baseFieldElem{0xfb1a6e9399b4d269, 0x1fad4ea253f16826}, baseFieldElem{0xd04b4345a15174bb, 0x718b2983dfc57dd4}},
			t: gfP2{baseFieldElem{0x5e17c6efc0229606, 0x54d576d51d20dcba}, baseFieldElem{0x9f910a20fab32071, 0x0d3b534d5e7dce01}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1340b7d88dc2fa04, 0x5d19f989d4eb95a2}, baseFieldElem{0x49316672dc923607, 0x675966b88a5c2411}},
			y: gfP2{baseFieldElem{0x6b7fcf396eb6305b, 0x22fe8ed6c7f43f14}, baseFieldElem{0x75c5b677ef904179, 0x06624f8dcd44be88}},
			t: gfP2{baseFieldElem{0x9a11ee2c2aac9e0c, 0x0264ad6a31993ac3}, baseFieldElem{0x8691f8c216398d5f, 0x17e23191ddd2542b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8cf270a98cf8cda4, 0x50673c733fde5e5f}, baseFieldElem{0x05836d896a3c6ec8, 0x5f209097bba2193c}},
			y: gfP2{baseFieldElem{0x967166fa497ffc2c, 0x1be666cf87956191}, baseFieldElem{0x743e59aaeeb341a0, 0x2ecff24f58d805d7}},
			t: gfP2{baseFieldElem{0x12b98c8d8e3eae98, 0x4342fc0e8f62298c}, baseFieldElem{0x4892120e7e88c5d8, 0x1fffa02a9ed0745f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2766365cefd19914, 0x4e55a0ad9eb1aa47}, baseFieldElem{0x21f997b877c7fa6b, 0x542e199362ee2cbd}},
			y: gfP2{baseFieldElem{0x066ca7b27886e791, 0x4bf76ac787f2af56}, baseFieldElem{0xcb2e3c3ed73b547b, 0x56684131e2b6974c}},
			t: gfP2{baseFieldElem{0x4474ca04070e1160, 0x26e7e3e4a43d7bd0}, baseFieldElem{0x9c779bc9fe035dde, 0x1cd83b0b09201656}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8efbebfd41fb445f, 0x43ae324edc56c313}, baseFieldElem{0x39c23534d6a952ba, 0x46546748330d239e}},
			y: gfP2{baseFieldElem{0x26f2c5bfc0e11f30, 0x3711fd281b52ca82}, baseFieldElem{0x2f2c901d6fcddcff, 0x2641bbf37a58c07a}},
			t: gfP2{baseFieldElem{0xa01a19219fcb15ba, 0x43eef32db1d15e34}, baseFieldElem{0x0abe93ec318f52f7, 0x6b634a35b563d29c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1c86991159d67e5e, 0x66c8fe29dcf9881d}, baseFieldElem{0x7d086759fbb4e042, 0x6d2cbf9d65e32571}},
			y: gfP2{baseFieldElem{0x39fc530a399c53e8, 0x7fa65e1306b3bf46}, baseFieldElem{0x9bb7f81f118dc895, 0x29ce58575a4c2090}},
			t: gfP2{baseFieldElem{0xf588e0f75c2b3b8d, 0x232feaef78c7086d}, baseFieldElem{0x0f126e97cdba9574, 0x1c13887f41e9fc6e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa89401330533afc1, 0x234d01fad73af8a5}, baseFieldElem{0x894dcb40a04261ab, 0x758dd368af726b98}},
			y: gfP2{baseFieldElem{0x9e9f0c699794bcdc, 0x65d5b44d76281d2d}, baseFieldElem{0xa344aacf12c77a8d, 0x7465015b7a0f99a6}},
			t: gfP2{baseFieldElem{0x599f67c71424eafc, 0x39c8a3ff29030186}, baseFieldElem{0x5f42913659fef505, 0x10c9c47a0e6e8bc0}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xc54034d48ae2043a, 0x08238a258f9a34e1}, baseFieldElem{0x2ab3ab6bc8510d24, 0x41990ac7bf52d279}},
			y: gfP2{baseFieldElem{0xbfbf22598ff64860, 0x7233803e7cede9c6}, baseFieldElem{0x40b0f6f07f4d0339, 0x53653c2bfe7af56e}},
			t: gfP2{baseFieldElem{0x7a5f24dad8a0a3a6, 0x021b6550fbe3faee}, baseFieldElem{0x699b2c0997fe042d, 0x7c4d3245b07da8e4}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4242953dca55cc0b, 0x36c0f42ed028ce24}, baseFieldElem{0xb7e57fe72f0bf086, 0x0fc0f2672475ecb0}},
			y: gfP2{baseFieldElem{0x6b1860244dbb75d4, 0x0c3b8669a518de13}, baseFieldElem{0x90c3d8ee269116fe, 0x4628234511a57c94}},
			t: gfP2{baseFieldElem{0x03dd1bc2ac3a454a, 0x37b9fdaf6418593b}, baseFieldElem{0x9d7e34a2b6493a88, 0x4f2deb4239c190d6}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa73b1d90d7974dc3, 0x5ffa53e44f0db1ba}, baseFieldElem{0x4e7f6f90277d7a00, 0x3414fdd58e42c887}},
			y: gfP2{baseFieldElem{0x6492891d6be21aeb, 0x1609336d6f71caa8}, baseFieldElem{0x997726f7acbf7156, 0x0936dfc6dc960db2}},
			t: gfP2{baseFieldElem{0xcd41671bca94c9ea, 0x0768c0993767f22c}, baseFieldElem{0x17c1bbe32285633e, 0x6d4b6278a77cd36b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3ad4f18358ff1719, 0x1e6a8e9498fb54d6}, baseFieldElem{0x3a1f135c367c6a9f, 0x2ecb9e8761a90c10}},
			y: gfP2{baseFieldElem{0xe3f9b1ca8bed3abf, 0x2420d78d223f3c1d}, baseFieldElem{0xa3401f8de73bca8e, 0x6393efb1de09763c}},
			t: gfP2{baseFieldElem{0xca4d96d099fbd043, 0x6b7858ebc0021054}, baseFieldElem{0x5a8ffe40ff9cb1c0, 0x3efbf5df2a7157aa}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x041558a4380c80ef, 0x579f494637e8a47d}, baseFieldElem{0xe833b35f6e343742, 0x5fd930802f367738}},
			y: gfP2{baseFieldElem{0x5e2b410cef8b5b10, 0x63a41e61ae5bc4e4}, baseFieldElem{0x85605baa518b4a3f, 0x34d73ec1bf59aee0}},
			t: gfP2{baseFieldElem{0xe1b9687e175ce7bb, 0x702c3d1776519f24}, baseFieldElem{0x667df117b30fd2d9, 0x0309dc001ece109c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3bdb1f36d0988bb6, 0x430cd0fc957e892c}, baseFieldElem{0x8dd7a8c66b0c300c, 0x3608757f21051d42}},
			y: gfP2{baseFieldElem{0x6e2de491722908f0, 0x3f958ca8ecdea789}, baseFieldElem{0x052313cc9915ea78, 0x2da5f6409e8f5141}},
			t: gfP2{baseFieldElem{0x932bf74cb2b2d8ef, 0x0a30499ac314eebd}, baseFieldElem{0xc8490319607c437a, 0x347154d6a67416a2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa3959808672b866d, 0x13e518db3c8db6f2}, baseFieldElem{0x64b6422970bd6b1f, 0x639d97f3ebeefe4d}},
			y: gfP2{baseFieldElem{0x9e1d13f9ca558050, 0x1e16823fcf141eb8}, baseFieldElem{0xaa8ee4b237879a40, 0x68ad9aa6b41fb1d3}},
			t: gfP2{baseFieldElem{0xb28e86f97a3f0af4, 0x655cda39230f316b}, baseFieldElem{0x8b9175dea68d6c35, 0x336a3ac225916d71}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1047c03274ad1719, 0x264f8d2c6f8b3887}, baseFieldElem{0x55d783c4abed9b41, 0x6031ae03de3a6b92}},
			y: gfP2{baseFieldElem{0xf090fbca01f01fc6, 0x3d4c0ab06b3ca9bd}, baseFieldElem{0x1856bec1ed89db69, 0x3afa4772ba25b56e}},
			t: gfP2{baseFieldElem{0x384e8906a716b128, 0x046f330711931491}, baseFieldElem{0xb69b0f50465024f2, 0x51c6bf423199f6cb}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x938be8d51be06524, 0x5714c8d613086d2d}, baseFieldElem{0x6e084315217ea203, 0x4a8b7bcfb8d1d668}},
			y: gfP2{baseFieldElem{0x52c965c052ce199a, 0x3297e7ab93ceee6d}, baseFieldElem{0x6e353748c5ef4fc0, 0x36933b47ba9346cf}},
			t: gfP2{baseFieldElem{0xadb43af287ac1661, 0x7bd5c96c95be5ac2}, baseFieldElem{0x13b5a3f8bb509cea, 0x0b2f782d7bbef4d3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x407ed3d03a368d38, 0x62165f82a58396aa}, baseFieldElem{0x80b36e4d7d2812a2, 0x34a55e299a8ba46b}},
			y: gfP2{baseFieldElem{0x38876a44a7dbb8c0, 0x0fe02e9fafc8092a}, baseFieldElem{0x102b61eeed4ed8b5, 0x4ab1255aafb8fbe9}},
			t: gfP2{baseFieldElem{0x2cbd02f9352d8103, 0x1df9a6a41482b5e9}, baseFieldElem{0x7e69d30d579692ba, 0x73a93fa4aae6a236}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb0aef0efaeed6175, 0x2b7ed1e876d63149}, baseFieldElem{0xee3b620691e6f158, 0x0a745887942d01ec}},
			y: gfP2{baseFieldElem{0x6281c4958316bc0c, 0x0ab69ee23963e2a8}, baseFieldElem{0x82486b3a7f206848, 0x01d1e369f99b0826}},
			t: gfP2{baseFieldElem{0x1bf9e39981144202, 0x5f7ca87c41589a01}, baseFieldElem{0x508eaad07280dbed, 0x4c384e770165b7e5}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x972871d541c765c5, 0x4de4f3ddb61791b2}, baseFieldElem{0x7442c2334a4f1078, 0x118c4e12ea299ffa}},
			y: gfP2{baseFieldElem{0xa38cae9333be8876, 0x1c07dccdfe2fd95e}, baseFieldElem{0xcbceab2e1298ba2e, 0x6c1f173056a3e634}},
			t: gfP2{baseFieldElem{0x3872aa54b6d74378, 0x30b575838ab3958e}, baseFieldElem{0x257e2116188f684a, 0x0fc795bed2dfedd3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x5ecec47ee065d484, 0x7f0b0011488cb4eb}, baseFieldElem{0x7666274829eb12e7, 0x501ee81657a18c3a}},
			y: gfP2{baseFieldElem{0x84a4065833ee84c9, 0x55110387ca60d011}, baseFieldElem{0x6c8f89855dcd73af, 0x3b14310d09d51ec5}},
			t: gfP2{baseFieldElem{0x2e94422769463db7, 0x6fe122bb9155c23e}, baseFieldElem{0x4b331b100e7cbc51, 0x41a17d13b211526a}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0afbe417f5f1084e, 0x1dd4af49f3a03af9}, baseFieldElem{0xf2cef8e3dc605461, 0x6df5e02e8f99e04c}},
			y: gfP2{baseFieldElem{0x16a8b34d587f30a5, 0x7a01bf4b149ed0a1}, baseFieldElem{0xfd16a85cd212810b, 0x658dbe02caf2dfda}},
			t: gfP2{baseFieldElem{0x2ecbce07a6765797, 0x656aabc259e91d72}, baseFieldElem{0x7f3759a223a9c1e7, 0x5db261289d5c13e6}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x57fa9429ad1fd04a, 0x1337beb7a503b5c7}, baseFieldElem{0x443b0a4f8c353dbb, 0x58bf6db00d79bfa3}},
			y: gfP2{baseFieldElem{0x8d238d0902aa881a, 0x41626335a374e82f}, baseFieldElem{0x648d71a5cc1b654a, 0x5662ee08d67e8fb7}},
			t: gfP2{baseFieldElem{0xb71c20e7ad2a9e8e, 0x14721f11d781a3c1}, baseFieldElem{0xa4b3c4f2c24393f6, 0x15d90f1a0d317737}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x49fcd22ff915800a, 0x41f6bc9d3aee27f1}, baseFieldElem{0xcbb277688253a855, 0x05d8265d2d591bfe}},
			y: gfP2{baseFieldElem{0xf3b2890bab4cfab6, 0x2aa242b54ac6fc2e}, baseFieldElem{0xe8d5e428951853cd, 0x33235db85e332310}},
			t: gfP2{baseFieldElem{0x7d179d60508ff6a5, 0x1cb46c8f1661c852}, baseFieldElem{0x93bbb01fcaec5b00, 0x356cce4782bf4852}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9b33621b25fe83e4, 0x0918c49fb4062991}, baseFieldElem{0x737a8ca960ccc545, 0x0ecfea853f944c24}},
			y: gfP2{baseFieldElem{0x499d1ba2d806c99c, 0x4eb6d87ef7dc52f9}, baseFieldElem{0xf01731740e44e907, 0x696d74025df07273}},
			t: gfP2{baseFieldElem{0x03edcd975c2b9b61, 0x640da047e5d91c51}, baseFieldElem{0x34dbb46501093f42, 0x46f2546b3e4d5230}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4e1fdd4e2beef90a, 0x66c7ea7880211454}, baseFieldElem{0xa3608d3f2efd88fd, 0x32bc1eef38a2f588}},
			y: gfP2{baseFieldElem{0x1d5187fd805f1d73, 0x2f1e55f917a62f19}, baseFieldElem{0x77ea93e4b88b302d, 0x336b6cf3f32a7bb5}},
			t: gfP2{baseFieldElem{0xa2e64d8e8004b1f5, 0x237cc2a6b01a58b8}, baseFieldElem{0xad1c24a2d58a2968, 0x593b5131ace05f29}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8243ae99fc17abc5, 0x5abd66f7834b8ee3}, baseFieldElem{0x3a148120e3057b97, 0x36df593fe7aacb30}},
			y: gfP2{baseFieldElem{0xcccbb71d2c4731c6, 0x0920573c69348fe7}, baseFieldElem{0x5c9157af2ec73aa7, 0x1a89661b5aa71372}},
			t: gfP2{baseFieldElem{0xbd99efbd2c84a290, 0x206aeed4a55f1f8c}, baseFieldElem{0x7fc0fe1d2740c327, 0x131b47cbf356132c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb8461b219662c50f, 0x37facb7096a2902d}, baseFieldElem{0x8963e324768d03fb, 0x4fa604fbe77e1bbc}},
			y: gfP2{baseFieldElem{0x17132fc5aed8fad1, 0x18316acaf41b02a4}, baseFieldElem{0x9976836a00daf020, 0x454ba4901067c9d6}},
			t: gfP2{baseFieldElem{0xd2021cfb1b2eb3b4, 0x6b7e694b57b6af7c}, baseFieldElem{0x95906fd4ebe7bd59, 0x37f38ac5fd4ae71b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x86311fa1acfbae21, 0x5bf9826cf543bdbf}, baseFieldElem{0xc4d3e18fab253be6, 0x49e78f1115091bac}},
			y: gfP2{baseFieldElem{0xb8999af4641d2e35, 0x1fbf7d7c821f5fde}, baseFieldElem{0x586ad6dbcce8efea, 0x0eea618b50721f36}},
			t: gfP2{baseFieldElem{0x580144465188540c, 0x212c707af361b925}, baseFieldElem{0x12e9426f9f4f93e2, 0x5b9a216691dec5f8}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x92ad13745664ce8a, 0x1ab11837c2cb97fa}, baseFieldElem{0x3a8e7925d24bc918, 0x3fa413ccc8e23086}},
			y: gfP2{baseFieldElem{0x47c98ea31632ce1b, 0x29ed73ae58e0fdc7}, baseFieldElem{0xe59042d5ecdb6e54, 0x69d4f8f4916c4c46}},
			t: gfP2{baseFieldElem{0x8b1edffeb15afd0d, 0x595d39ea2284233c}, baseFieldElem{0xa724983e80178d37, 0x4cfe5391f6c71228}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa6768c7b9cb94202, 0x6bf12b7004fdcd77}, baseFieldElem{0xb3de344c53083aeb, 0x406b36f911d60cc7}},
			y: gfP2{baseFieldElem{0xc99f544c50f3e417, 0x14ee215f3c1ecb69}, baseFieldElem{0x8a156a0a52278ed1, 0x0b41d1680cb10d9c}},
			t: gfP2{baseFieldElem{0x4a074b9a4064d88a, 0x7dda6a1e8008672c}, baseFieldElem{0x72fdf412b0a136ec, 0x50c9dd81da504d54}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0ea2e538fc296086, 0x3ce914af89926f1d}, baseFieldElem{0xa604672df9d64533, 0x1402cc13f6b9911b}},
			y: gfP2{baseFieldElem{0xb91386a71d3801af, 0x40fc0c9637d83cc9}, baseFieldElem{0xc9c9b276c83fbc00, 0x5d74ec43e71bbe62}},
			t: gfP2{baseFieldElem{0x97c1c00ab60e2e23, 0x6e2048b27860671a}, baseFieldElem{0xc17a16dd4fcdf366, 0x18494eadf68c7664}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x43b7f7bcf6df80fe, 0x23e3a3ecdaff454e}, baseFieldElem{0x103310d65feb01a9, 0x0e4569110830842a}},
			y: gfP2{baseFieldElem{0x306bbc81a9cdd58d, 0x1e042ad467dcb9a1}, baseFieldElem{0x4850b12f87c9713d, 0x55c9f6bd22df9180}},
			t: gfP2{baseFieldElem{0xfadd20368a4153c4, 0x3289b7e131ad59a2}, baseFieldElem{0xfb6994c6ee0dce59, 0x3d486fc6c8f4fa6d}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x548aedcc53000ce4, 0x59f9f42ba0c26db9}, baseFieldElem{0xe62a7c6a3b2be4e9, 0x2c987a904fe7e467}},
			y: gfP2{baseFieldElem{0x8c65f5d6c8ef4503, 0x31305ff339e9fbc8}, baseFieldElem{0x8329234120bd7073, 0x11c9c8b051eb6838}},
			t: gfP2{baseFieldElem{0x410dc00f9a83543b, 0x1dd9b616691bf894}, baseFieldElem{0x1b809096267f571d, 0x70ed51c11cfc0d97}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xc99f667314c78ab8, 0x7644f12b92611547}, baseFieldElem{0x3ae2500eaa4ce824, 0x197c019b6b519d0e}},
			y: gfP2{baseFieldElem{0x570b9a4ff8d98662, 0x57cf91fe7da16a28}, baseFieldElem{0x3da578a2190d455c, 0x011245665b17ebf6}},
			t: gfP2{baseFieldElem{0xf1b078eb5f1f6b38, 0x4170d9b9f306be3d}, baseFieldElem{0x6ec705bcda59aa8d, 0x1d5a95cf7d034073}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7a3fec7283433737, 0x4fd4dc8443ab45cf}, baseFieldElem{0x182830413261e18d, 0x7fc2e92cc31ebaa6}},
			y: gfP2{baseFieldElem{0xfe683a6337f30bf2, 0x0b5f3aff3ca28f23}, baseFieldElem{0xfe8b61d0119c990e, 0x21567e70dcceb45c}},
			t: gfP2{baseFieldElem{0xf771d14aba0dfd5c, 0x54b8fe94ba51d370}, baseFieldElem{0xc93869e0c0d7b178, 0x0b32b36ef01abab1}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0a2a69f1b4a59315, 0x37ab859e638cbceb}, baseFieldElem{0x73c4343f438dd6c5, 0x789c2056bae273bf}},
			y: gfP2{baseFieldElem{0x9c0a497225510d71, 0x07b4313685c7e6c5}, baseFieldElem{0x3ee952728c59915a, 0x453e982c9689882b}},
			t: gfP2{baseFieldElem{0x8d76d08b45c5da03, 0x4ce773d8c2d39aa3}, baseFieldElem{0xe2e0d14437022d0c, 0x7a89df23f5c63ca1}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x63827515993fb5cb, 0x0e4a07f9f8fe42dd}, baseFieldElem{0x215ae08dd8c8fd6a, 0x37cd898bb39b1832}},
			y: gfP2{baseFieldElem{0xb31dc817641a8769, 0x40ec9b7d4ebe7fb3}, baseFieldElem{0x4fa32a30e947add9, 0x31fe0faa62901675}},
			t: gfP2{baseFieldElem{0xe253c3d5932d2cf0, 0x160258b200beb9b7}, baseFieldElem{0xed6b982a689cdcc6, 0x33c0d705a7979158}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x635a43f48a432ae7, 0x44f060020da330d8}, baseFieldElem{0x809a8a216e916c7f, 0x036b604cf4a70d69}},
			y: gfP2{baseFieldElem{0xc8d83baf4b63d5d7, 0x6312cbffa8632927}, baseFieldElem{0x31570b55ed407ae7, 0x164d04beee70aabe}},
			t: gfP2{baseFieldElem{0xc6ba7d4efe28e79f, 0x02ba1321c44ee49d}, baseFieldElem{0x7b37208f05d6f27d, 0x556dc56af51b4140}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8588e89b891f72d6, 0x74b93321f66d3e25}, baseFieldElem{0x1ee6867ca5025b5c, 0x36b1e8c44f2af265}},
			y: gfP2{baseFieldElem{0xc20903a582f4396b, 0x48518c65defdd3ab}, baseFieldElem{0x2b7fe80859994879, 0x188c76959c09bc94}},
			t: gfP2{baseFieldElem{0xe8c5e590c38f8e98, 0x22963c0609323249}, baseFieldElem{0x10d89892b2c2d7b4, 0x0d5190dc55066de3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x67c0c63d2351b813, 0x40a3bf87e10e2031}, baseFieldElem{0xb0bbf02a9b655085, 0x2136d0d5ade31d0b}},
			y: gfP2{baseFieldElem{0x13ebbbf0c9a46c24, 0x3213d038340fce79}, baseFieldElem{0x4c14ac49cd95055a, 0x598179dd2f66555c}},
			t: gfP2{baseFieldElem{0x826228cd329b8d2a, 0x2ebe1e1d8c29e164}, baseFieldElem{0x05fc349d391402b5, 0x5303a57f94bb85ad}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xf8ece4401f5151c6, 0x741277b0456d4e5d}, baseFieldElem{0xdbbbcc0993932df1, 0x1ef1b2f06bd3c461}},
			y: gfP2{baseFieldElem{0x6962dd6df598bcf5, 0x79e20a98fa30cd14}, baseFieldElem{0xb5b1ea6173caefef, 0x4481151faafc944f}},
			t: gfP2{baseFieldElem{0x1af09c41dd788cca, 0x071639d56ad85a4c}, baseFieldElem{0x9b6a3b3b252e334b, 0x38094a8db3fb7bca}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x602e3fdd1cbe037f, 0x259bda839b02ab8c}, baseFieldElem{0x1266734904b68e40, 0x6873d3d8e08539c4}},
			y: gfP2{baseFieldElem{0xd22eacb17fd25628, 0x5498234dc2bd60bf}, baseFieldElem{0x750c97a57af3e990, 0x3b5d079fb763f9d8}},
			t: gfP2{baseFieldElem{0x4e92289400cdfdd2, 0x47b266bc0158a2c7}, baseFieldElem{0xb8fbde33f052fa41, 0x25050027f0b5b944}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x46cc7376d8a6795e, 0x335ee315d0ac6a70}, baseFieldElem{0xaffa2b91086b8e2e, 0x5fb0ff5b967339b5}},
			y: gfP2{baseFieldElem{0x7fb0f744ca26f719, 0x3d849b1a3cc43ecc}, baseFieldElem{0x1cebb959d8402e26, 0x0bf7f9f772940339}},
			t: gfP2{baseFieldElem{0x1ab01aa02542bc1b, 0x59472b741edaad5d}, baseFieldElem{0xc3c79977d403d51f, 0x456c029afbc88796}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x51f4ebc06a402341, 0x294dc46fd702d702}, baseFieldElem{0x30a2564e81b009b7, 0x4b9636cfb2a9a467}},
			y: gfP2{baseFieldElem{0xa02632601006253c, 0x6b2f99c22f09cd33}, baseFieldElem{0xeb23cbdf11122e24, 0x0d43dff8a44786df}},
			t: gfP2{baseFieldElem{0x3850db13b1d0783f, 0x38d87f11699df4fe}, baseFieldElem{0x046162910dcd874f, 0x0af0e52629b2a145}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x4bd251bfdf1c9369, 0x2e24fb6f2580ccf4}, baseFieldElem{0x898edecfb6bb44e0, 0x3a0a041892329424}},
			y: gfP2{baseFieldElem{0xd42c00cf63ec6b8b, 0x2a3e42e4c7e33b50}, baseFieldElem{0xc5b06e08d4d434b1, 0x613e5ac8e0e7a254}},
			t: gfP2{baseFieldElem{0x5c5edd85105a3805, 0x2bdeb26af5734384}, baseFieldElem{0x8888f917f37a0b2b, 0x1679ee5381b10c43}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7923e6c95e8361b1, 0x1024e7b28165ff58}, baseFieldElem{0xab4e5a371d56f361, 0x3bf8e970c016e171}},
			y: gfP2{baseFieldElem{0x4bc82d7e943152ab, 0x6f99b747e0a47c85}, baseFieldElem{0xed082e67f2226f21, 0x385c5e8a3cb5d3d2}},
			t: gfP2{baseFieldElem{0x6a8ef993ee048711, 0x1c1c569ffd9206f2}, baseFieldElem{0x61b2fd95adb0cf65, 0x62fbcfcb8ac47f49}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x5e50f8f769f69a53, 0x31290f73ea58121d}, baseFieldElem{0x35f841bec72b9cf3, 0x78646b6006a392d7}},
			y: gfP2{baseFieldElem{0xff8017baad47661a, 0x29a822aa04d59fde}, baseFieldElem{0xb942b910e782f368, 0x31f3febded7f5c27}},
			t: gfP2{baseFieldElem{0x4f176889a6bd9ba8, 0x3bd5442a2fd27ced}, baseFieldElem{0x67679d4e8a411f19, 0x7ab7f13a163c72f1}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8e8eeb181495b07e, 0x396c791e08106bc8}, baseFieldElem{0xa06fd1fd61253014, 0x2620808382ee4335}},
			y: gfP2{baseFieldElem{0x4bf2a6d21eb50bcf, 0x6dde0b856098deda}, baseFieldElem{0x0ff95de88d711d96, 0x66c263c4adc9bbdd}},
			t: gfP2{baseFieldElem{0x6f463a3644a8d09b, 0x264b25036718297f}, baseFieldElem{0xb0f8cc28f44ce648, 0x3e511f8678bd1a54}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x252a5e51f66b01d0, 0x433c4ce7aaa1d050}, baseFieldElem{0x7655000da6783d46, 0x103c9943cfa2acfb}},
			y: gfP2{baseFieldElem{0xf37eb2f608a132a4, 0x65f750df8c541d56}, baseFieldElem{0x1de22add70b14349, 0x7390b2239322386f}},
			t: gfP2{baseFieldElem{0x20954d1ddb35aeb0, 0x5ea8d59d59be898b}, baseFieldElem{0xa814786ee136b81b, 0x05fa72c5782db2d2}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2306ba00c06c5e96, 0x47acca4569141082}, baseFieldElem{0xc153788bd8356fca, 0x6e49f0d49ac09e32}},
			y: gfP2{baseFieldElem{0x9bac3e9d221be988, 0x3b1ecef507b6e404}, baseFieldElem{0xa0dd5d045936f75d, 0x24b84ed884ef118a}},
			t: gfP2{baseFieldElem{0x59cd15c8e69139ce, 0x392bbe2d16982788}, baseFieldElem{0x861bc81a709f83b7, 0x3aee93699d4e3bba}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xca14ad2c225d4dcc, 0x40a659e06d7fb4c8}, baseFieldElem{0x25d1e0b4041b06da, 0x603f47e6b5419479}},
			y: gfP2{baseFieldElem{0xd7f1163c12711a6e, 0x09e12902edd871d5}, baseFieldElem{0x24c34cd634ee14cd, 0x07c5434672fd7962}},
			t: gfP2{baseFieldElem{0x6da0cc42ac3071be, 0x6595d882190610ca}, baseFieldElem{0xcf50bb67348bb728, 0x513b85cb31408d2b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xc454929b04e79c84, 0x3f161d9151fc5136}, baseFieldElem{0x565970ccb180aa02, 0x4b4ebf6113096144}},
			y: gfP2{baseFieldElem{0x5e2d45dfa1da701c, 0x41d0cb6fb05b2935}, baseFieldElem{0x7175cb606eede1a0, 0x3d7d457e90d02519}},
			t: gfP2{baseFieldElem{0x570fce68c0d5617e, 0x40ca60dae80f7c1b}, baseFieldElem{0xda1d8abbe6ee776c, 0x29cad91cffc8301c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xdc38ee5e44804c41, 0x7929896fd88ff11d}, baseFieldElem{0x1d34482bea94e035, 0x47feea217b3ea1a5}},
			y: gfP2{baseFieldElem{0x449add2340e43f57, 0x6101f9c3555c8091}, baseFieldElem{0x96ec06537cb630f3, 0x48c19fa788cee08f}},
			t: gfP2{baseFieldElem{0xad200040e5a75d3f, 0x7aa712782eea227f}, baseFieldElem{0x7d16366225901bee, 0x25bb8e6513f1bc59}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x1830d91766e4ff6b, 0x0e8c19d7c9a5f479}, baseFieldElem{0xde9d38d436b41d56, 0x7d58743a47ee361c}},
			y: gfP2{baseFieldElem{0x5fa7f5aacb6dcb16, 0x2ee3f795dd5a679c}, baseFieldElem{0xd67559580741079a, 0x47b1f80983d7f6a9}},
			t: gfP2{baseFieldElem{0x214a6a08c5be08d1, 0x1560537979a993d6}, baseFieldElem{0x00ba88278223e537, 0x3787d9b881454131}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x9caf1df1a32f1286, 0x1be2818db1f5c37b}, baseFieldElem{0xb37e2954ec15648e, 0x2e72cc61051e8fdd}},
			y: gfP2{baseFieldElem{0xe74d3224e02a92e6, 0x15c3acda5866ae01}, baseFieldElem{0x7e29f2047bb63692, 0x43861ab3938a5650}},
			t: gfP2{baseFieldElem{0xce3c90bd420dfa2b, 0x4367e0b39a156789}, baseFieldElem{0x9bfaed026861f9cf, 0x723d8106c0afdb4c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe7af4b580f9cbdee, 0x394042048fd2e6d8}, baseFieldElem{0x9db208874ab4c070, 0x50e4dcf5fd2a04f5}},
			y: gfP2{baseFieldElem{0xe4ee1c8eca7f5960, 0x69bd2c80f048b782}, baseFieldElem{0xc7552f0d3313bdb0, 0x171a1cbe38ab9438}},
			t: gfP2{baseFieldElem{0x6ecdc4944cfff8d5, 0x0e3822c946baa273}, baseFieldElem{0x968e67bed8fd86d2, 0x3981b18354f9aa7e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2cf958f0c059b554, 0x7666db742162745b}, baseFieldElem{0x8c754a1d28e6fd60, 0x323e925455f3d204}},
			y: gfP2{baseFieldElem{0x3d7f0fb60a75eb0b, 0x7c0e64c4d089ab76}, baseFieldElem{0x7c1ca8f3b68d6219, 0x5a57c5ac4586d3f7}},
			t: gfP2{baseFieldElem{0xe497316bf73f2561, 0x640e9f2b37b2b9f5}, baseFieldElem{0x6a997de4fcaf8a0b, 0x22770503c8ea00cb}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x0c4ff023f6e16fd4, 0x3520806107328b72}, baseFieldElem{0xbb9a33c8d2479c47, 0x56060ac579559355}},
			y: gfP2{baseFieldElem{0x4d2fbdc264ab68df, 0x4f61367caf487687}, baseFieldElem{0x9ded9f17fed6869d, 0x42b6d650ab6aa52c}},
			t: gfP2{baseFieldElem{0xd04a91f7970cd84d, 0x13859159d0c125ea}, baseFieldElem{0x1e683086b6fa7ab5, 0x0e5f1a6dd199d266}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb4bff532a192977a, 0x12de9726c423b1aa}, baseFieldElem{0x63193141e9f87788, 0x2122423458bf91b7}},
			y: gfP2{baseFieldElem{0xdb549045e85cda9f, 0x55f74f4073c72d42}, baseFieldElem{0xc767a560c126b419, 0x11d76a8588eafabb}},
			t: gfP2{baseFieldElem{0xcefa807940bff41f, 0x11886837a6450a88}, baseFieldElem{0x2794c03749985bb6, 0x35e560adf0177cde}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xf074c42327e7dda9, 0x53995a8862869406}, baseFieldElem{0x5a09fa93202fa9eb, 0x2f14051011e81bb5}},
			y: gfP2{baseFieldElem{0x69a49f73f9dba0e2, 0x65d52b4fc42ae0c3}, baseFieldElem{0xce6cfb73231197b8, 0x681dde290e837d2b}},
			t: gfP2{baseFieldElem{0xff75001b17408764, 0x4876c54cbfdf54e1}, baseFieldElem{0x487571c0f566a0c4, 0x3fab7842e36f9ec9}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb12a4df4953ed9a4, 0x4501f0c6e0c414c1}, baseFieldElem{0x185f2ebdc3914583, 0x6b7b098f44e93a62}},
			y: gfP2{baseFieldElem{0xc6b4eaa6ed0338e9, 0x7c3a46d731a49b61}, baseFieldElem{0x1b632562db157a07, 0x053e0bf58d7363a9}},
			t: gfP2{baseFieldElem{0x11499ae7ee6d3637, 0x6c39edce40794a9a}, baseFieldElem{0x0ff93c5fbad6c49b, 0x0695182d30ca3dc9}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x3ed03d17d8f8e11c, 0x667a17d618f771b0}, baseFieldElem{0xda70d3245a38a0e7, 0x162641f3d46699db}},
			y: gfP2{baseFieldElem{0xfe707fca56efec2a, 0x502947a424c0181d}, baseFieldElem{0xe005aa3cb272dc0b, 0x309cf667354df85f}},
			t: gfP2{baseFieldElem{0xdf37eb470ec78fd6, 0x75c22b8f9541c583}, baseFieldElem{0xb9cf349be3c71b3c, 0x6fe24abb3d20a063}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xba7206670c1be500, 0x537c714159412d6c}, baseFieldElem{0x94dc95f5b108bc40, 0x7f13a244b33f7f4d}},
			y: gfP2{baseFieldElem{0xd3592af00e85ef30, 0x22880f1f19f69ef8}, baseFieldElem{0x4322ce3916a9d4a1, 0x62e5d7522f5d780b}},
			t: gfP2{baseFieldElem{0xb07d1397e71ee1f1, 0x73ad4afc259f4c8c}, baseFieldElem{0xd7498b792f935e47, 0x483c3c7820adb379}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2cd9e0cfa7074781, 0x434670217f0fb990}, baseFieldElem{0x7308fffa168abec7, 0x1646b21cce6fa41d}},
			y: gfP2{baseFieldElem{0x95cda62f554094ff, 0x6bd5e22e4f94fb51}, baseFieldElem{0x2b71c5a8fbc02c97, 0x707fb120a1aa040a}},
			t: gfP2{baseFieldElem{0x7071516944fbc745, 0x095d4b202b5af6b2}, baseFieldElem{0x05b6ba3ab29087d6, 0x56e3c7be74fcaf58}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x5795657de4a03a9f, 0x2b14dbafd5a0cff8}, baseFieldElem{0xca4323c0ecec3bec, 0x4c489a95a906e31a}},
			y: gfP2{baseFieldElem{0xa0049dd4c5deb0c5, 0x7bf30d06091b8933}, baseFieldElem{0xe627fab0eabd60e4, 0x19a0dad21ff72a58}},
			t: gfP2{baseFieldElem{0xcaf96dbb115aad81, 0x5dc23a0d7643bf4b}, baseFieldElem{0x7be921435142d26e, 0x03349ff057cf1c5f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x8530344eec59ff5a, 0x4244bdd74f2da84b}, baseFieldElem{0xdb1be401e790bfe5, 0x13e84bc1624b29b6}},
			y: gfP2{baseFieldElem{0x6187739b8040c037, 0x05275b535d1a12d5}, baseFieldElem{0x5a9c70929368ac48, 0x59c2ce911655ec31}},
			t: gfP2{baseFieldElem{0xa5dad8f32241d89d, 0x5150e1d20ba67a6b}, baseFieldElem{0xdd71fb3f08072fb1, 0x37b8fe98082a2205}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x895dd68e4336706e, 0x67ec11b7342528d6}, baseFieldElem{0xba5c966206931199, 0x585b83c138583817}},
			y: gfP2{baseFieldElem{0x7a5a199164502fbd, 0x4d98f448801a9d20}, baseFieldElem{0xf8d9468b79af5cd8, 0x24c355fcc73af746}},
			t: gfP2{baseFieldElem{0x261a39e7b10d37dc, 0x7acbbee0b7317ed0}, baseFieldElem{0x23572031cab8b70f, 0x36bde89942bfc493}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe2c198f97f523996, 0x25078045532aaa09}, baseFieldElem{0x5fd02ea67008a234, 0x6af5cb5dfb454d74}},
			y: gfP2{baseFieldElem{0xbef31c78676ccd6c, 0x4578bccd2776364b}, baseFieldElem{0x9d29314d9dde1781, 0x05f30da4eff4eea0}},
			t: gfP2{baseFieldElem{0xa3fcdfbf0136fba7, 0x142f979583d8c0ac}, baseFieldElem{0x7d90d7b3a57784ba, 0x53f1a9b0421a792c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x82acc0ed16357d1c, 0x0771177ad5e1849d}, baseFieldElem{0x00fcf52f472db50d, 0x3f0ea2eb2d220c9e}},
			y: gfP2{baseFieldElem{0xe2e82a976c184812, 0x5937a170bb8185b7}, baseFieldElem{0xdcf4d79921833e04, 0x06f916e214b4e20e}},
			t: gfP2{baseFieldElem{0x79a08dfc9113d008, 0x0f9148bad64519a0}, baseFieldElem{0x8780cec3233cce75, 0x13104b4cfb7c0768}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa30824aca2d803ab, 0x1ab313987b47b6a3}, baseFieldElem{0x3fc719fe437d389e, 0x13ad174130e63cb8}},
			y: gfP2{baseFieldElem{0xcf98d1bd49d36da8, 0x1b0fe8da7c7d5d67}, baseFieldElem{0x561806205d9fad7f, 0x6a0bee6901b82ebd}},
			t: gfP2{baseFieldElem{0xb4268eaccab319af, 0x7fefec4ca4c4f4c7}, baseFieldElem{0x1bc42dd8bb512d02, 0x4d2a339ba971a05f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x66c23fc45abe0254, 0x0c5c901c2c1dbc98}, baseFieldElem{0x6ce53e61954aeaca, 0x37db209e8fcd5f56}},
			y: gfP2{baseFieldElem{0x9ba989e9f52f6a3d, 0x15973c81b9b7d786}, baseFieldElem{0x9066de195c4cd24e, 0x5c7ea1909a9a9611}},
			t: gfP2{baseFieldElem{0x40cb642fac4bab0c, 0x1b772ca290433264}, baseFieldElem{0x11a6940df6592037, 0x74f6d124ffb92d28}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xfb6ce26f14d2a79f, 0x4171afcccb9089e1}, baseFieldElem{0xbffb68ef1e6c99bf, 0x5f6aca5fc186485c}},
			y: gfP2{baseFieldElem{0xa53c75d0803431a2, 0x5c693070b5e3e5e5}, baseFieldElem{0xae4ec2bdac73c41f, 0x53eba1eb0645ed79}},
			t: gfP2{baseFieldElem{0x4bd0dc7ef2a66bf8, 0x519e841d2840aff4}, baseFieldElem{0xd7e23a70485dc696, 0x774d9d38ac4874bc}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x46c8c1541ae4e083, 0x1888ce37ea0dc27d}, baseFieldElem{0x4677e9eb67124585, 0x270fdcd25ba37642}},
			y: gfP2{baseFieldElem{0x2c0af95f1919f743, 0x557594766a67bd7d}, baseFieldElem{0x4c5128e57c0a9273, 0x33e0b4080e5fc802}},
			t: gfP2{baseFieldElem{0x00ca593cd183338d, 0x111b931c88b33a45}, baseFieldElem{0x9e94049885ca3ad9, 0x4c0b135fa75ad6ae}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x83191b0d46422ee1, 0x512c3d4a8d520f6d}, baseFieldElem{0x6d1e484a99ffc082, 0x1d4b095a28a471a3}},
			y: gfP2{baseFieldElem{0xbfed41c7b955f85a, 0x4af8a413ae6e67fe}, baseFieldElem{0x8cac406f24dff350, 0x520138de0e3577b5}},
			t: gfP2{baseFieldElem{0xa6d09f2f012e5928, 0x7a99b2dadf197a93}, baseFieldElem{0xfd2052a9bb13491e, 0x2f5fe4537f1b3920}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xfbe39ef0c0066647, 0x0b018ca8184b12c7}, baseFieldElem{0x4e135a5e1fe58cdf, 0x55fbaa238ab6ee35}},
			y: gfP2{baseFieldElem{0x22d727338875441a, 0x75308b6d73b8a48c}, baseFieldElem{0x8194eb94a42fcaab, 0x4d0bb098b4c4bf8e}},
			t: gfP2{baseFieldElem{0x9afb4f32612e86fe, 0x2c44617adeea2448}, baseFieldElem{0xede3dc0c35e5ed43, 0x6d622c8f95246ce5}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7030a4a76f15734c, 0x05f28d213864acea}, baseFieldElem{0x99e41adcc0bd8a6e, 0x64b6ed7787abbb42}},
			y: gfP2{baseFieldElem{0x690bed2639967ad7, 0x212c59764a88cdf5}, baseFieldElem{0x12ba3961898c87d4, 0x0db097a0cac38c3a}},
			t: gfP2{baseFieldElem{0xd322c5c2f955b4e5, 0x675bcbe99617d751}, baseFieldElem{0xebbeff89433afd54, 0x4b3d6d029fd742be}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe5be7c3031366b0c, 0x0b8d427e7d80a209}, baseFieldElem{0xba2d4fa0b51fc68f, 0x5df7a796ebbd1caf}},
			y: gfP2{baseFieldElem{0xabf33c5cc8fd9eca, 0x6f1064664b4183b3}, baseFieldElem{0x11ec26cc4793d529, 0x7d946e133a352ab9}},
			t: gfP2{baseFieldElem{0xb5bf50fa55b02b70, 0x642ac3efee9ad7df}, baseFieldElem{0xf3c2c34aa4135351, 0x2ccddbf94fc6876e}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x6945ff5243c777e7, 0x5d330157be260000}, baseFieldElem{0xfb5100bc426aed9d, 0x6d1660148bcc8624}},
			y: gfP2{baseFieldElem{0xfd260633f3fa6429, 0x57c241fd2bcecab2}, baseFieldElem{0x1c4bebc62ff57904, 0x3085c4fc415497f3}},
			t: gfP2{baseFieldElem{0x3f41a89cfba69922, 0x2ffce2b8ecad4910}, baseFieldElem{0xeaee45a467bba0bb, 0x0a087321e6f94719}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x40f8692b47db6e79, 0x11996e6601fdebfb}, baseFieldElem{0x2a5285bb637d565e, 0x7aafa8bf37c8ace1}},
			y: gfP2{baseFieldElem{0x758b620c38b105a9, 0x3747094d92f4906f}, baseFieldElem{0x804bd12ba71cd65c, 0x5bab6d3557297e28}},
			t: gfP2{baseFieldElem{0x0b6f7f9609756d45, 0x0886d3c5c7f03f21}, baseFieldElem{0x1ad82e2f41662907, 0x281acfa04cd9b273}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xcbd6a203af74ff3f, 0x3dc1c3cceeabce5e}, baseFieldElem{0xe5163a587f205aef, 0x0819a92bb599c082}},
			y: gfP2{baseFieldElem{0x06aef1cdd29d362a, 0x14852b7fa1621956}, baseFieldElem{0xf2a0a8bf4c940fdf, 0x09784909abe05b9e}},
			t: gfP2{baseFieldElem{0xe99462780408ff10, 0x1e7aa452faf78b37}, baseFieldElem{0x07d671425bb855cb, 0x4dc5a40e9b4eb25c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7292a882910c705e, 0x3e2a92c4f454452b}, baseFieldElem{0x76ac0ee75b14fcb4, 0x43271080e47d0ea6}},
			y: gfP2{baseFieldElem{0xf343283090b290cf, 0x5f579dfb37836aab}, baseFieldElem{0x37b1df4b520ad532, 0x06c43e297300fb66}},
			t: gfP2{baseFieldElem{0x7475db5d1e8c291c, 0x669660e9e5e6a7e6}, baseFieldElem{0x82bef004bbc9dd70, 0x7c06e8b1f1b9717c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x96ab41f1de05c954, 0x3ebc89f7ec80ad32}, baseFieldElem{0x71321e545a13046c, 0x305855dd596dcdf3}},
			y: gfP2{baseFieldElem{0x130e79d7bbf5c1d7, 0x3245717697000f51}, baseFieldElem{0xa1dc2f773e938189, 0x54c6d4f8a8c1fd7e}},
			t: gfP2{baseFieldElem{0x48d0798b95cb5722, 0x188e51eb53019596}, baseFieldElem{0x1421f7f1355f52a6, 0x69351892dfee4867}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xf377e9e852184162, 0x07aafe817d5a74d9}, baseFieldElem{0x95fcfd28e8d930fe, 0x0ae924ca03e1329b}},
			y: gfP2{baseFieldElem{0x55d27d1a99bb8928, 0x6f5bb2f67e6c8b9e}, baseFieldElem{0x0adc26ce5a94e6c5, 0x4869add4eb44b7ad}},
			t: gfP2{baseFieldElem{0x05b42ba543e9d5c7, 0x1b389a41d060fefa}, baseFieldElem{0x4ba9f8071e68bc19, 0x2c448fea51fab7e3}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x58c4cf7ee549e155, 0x235fc76939f8e429}, baseFieldElem{0x5380a36f86f0ed8a, 0x2ad2fd2328f12442}},
			y: gfP2{baseFieldElem{0xce0b41c27456ab57, 0x7e759eb7afc2b6e2}, baseFieldElem{0xbf028ce2cd631098, 0x512ccc51a10cd7ad}},
			t: gfP2{baseFieldElem{0xc74e41f0ec7f84ac, 0x1b46f24ce590f451}, baseFieldElem{0x6e6d5aa79d92fe0f, 0x72c4ccce8d7dcfe9}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb8319a38290d8f17, 0x1248f13f8668732e}, baseFieldElem{0xd7b39cfcc9561784, 0x09b321c9ff49c433}},
			y: gfP2{baseFieldElem{0x41a0ef5686eb445e, 0x6fedb42acf59a99a}, baseFieldElem{0xd796738883419a57, 0x0619f136b23910eb}},
			t: gfP2{baseFieldElem{0xafe8ccd32c190edd, 0x0b00f95ed4391e75}, baseFieldElem{0xaea44853f4e17c8e, 0x24aad910b9e8be42}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xe89d4efc74b0d454, 0x1107014986ed295c}, baseFieldElem{0xe0089f3d48f0d2b3, 0x151be0d161cb3c73}},
			y: gfP2{baseFieldElem{0xe729d2506afb4598, 0x328876a934517c1a}, baseFieldElem{0x9cdb38051db3e595, 0x750dc9acfe873e3b}},
			t: gfP2{baseFieldElem{0x82709f549d5de926, 0x3b8789beb9f6fd96}, baseFieldElem{0x5c04da6dd3ff186d, 0x46f94eb034263c65}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb2aca0c08789c0dd, 0x6e2cd8ba1acff374}, baseFieldElem{0x4252487fe939e0df, 0x77c41f450339ce45}},
			y: gfP2{baseFieldElem{0x5f4977af83adb29b, 0x35542e0bdd9192f9}, baseFieldElem{0xefe0c455d7a696ee, 0x6b28d74ba028022b}},
			t: gfP2{baseFieldElem{0x3ff49cfc450b5c28, 0x660f143352ac1b51}, baseFieldElem{0xfd2214ef2916b59d, 0x1e0e58417e09327b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x345533220836a1bb, 0x0d1888aae856b11b}, baseFieldElem{0x91266d7d2dcde645, 0x57b5d2b40ff74c6a}},
			y: gfP2{baseFieldElem{0x39c6b183ca531807, 0x1c948531dff7b09b}, baseFieldElem{0x0847bc81ef664c46, 0x331ac09a19a2d9ef}},
			t: gfP2{baseFieldElem{0x7a9b2b1414847b8c, 0x5044eb27ac0a7476}, baseFieldElem{0x0006beeebc80f1a3, 0x41621de9fda91983}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x6b333576d021c047, 0x347a2b2f48b645e7}, baseFieldElem{0x29048c901ca3c09e, 0x62f9797504a9fc95}},
			y: gfP2{baseFieldElem{0x2fa0b471cd63eb02, 0x0e205e2f61ea0450}, baseFieldElem{0x58ec8153a4e6a956, 0x03053789c7e266c3}},
			t: gfP2{baseFieldElem{0x9880b22d9f0d699f, 0x2a4cdae5fffd456e}, baseFieldElem{0x91f17ad44e856ff2, 0x0a95f377772be7fa}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x071bf58ef248d79d, 0x7403f4eb28b270df}, baseFieldElem{0x40fd47755fefde8d, 0x575f0e47b55d98d3}},
			y: gfP2{baseFieldElem{0xe52807fcec02ac87, 0x037db214ddb7fdac}, baseFieldElem{0x8180dd34f11b1fe0, 0x52851f4ae73609a2}},
			t: gfP2{baseFieldElem{0x5be1d9189a38d837, 0x65ad8253633d2517}, baseFieldElem{0x33074f8cc993eacf, 0x26cf691966e2ae2b}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x7fe847940e95e0d6, 0x58f3631f5ea473d4}, baseFieldElem{0x5778c99bc943e6c3, 0x25736aa76539f315}},
			y: gfP2{baseFieldElem{0x799514e1d6c61bf5, 0x5594c266d10dbcbe}, baseFieldElem{0xa130e8ecc8325b5d, 0x1bf51f9ff826a1e9}},
			t: gfP2{baseFieldElem{0x8ed695c35d39119b, 0x1611f9fe7d3de7f6}, baseFieldElem{0xbea6bf7f9c80e185, 0x5fc6c3e1e369b4c0}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0x2a45b0e0c298fdf2, 0x54b7ffc67feeafcc}, baseFieldElem{0xf631bfb6c32146ba, 0x286a04280499daf3}},
			y: gfP2{baseFieldElem{0x0bf8bb685c7695a4, 0x5a20241139dd55b0}, baseFieldElem{0xf0eed7a619a26b0d, 0x445dc8427e90d2a4}},
			t: gfP2{baseFieldElem{0xe590169db23d2808, 0x40d6d88e415635e2}, baseFieldElem{0x0890f3713982f2de, 0x2ce52bcf964aab7f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb6305541497dfef8, 0x41d8d61f6086bcbb}, baseFieldElem{0xdca308a459e8dbb2, 0x23c9b3adb0a5f7d6}},
			y: gfP2{baseFieldElem{0x235d10a644fda06e, 0x021c2dbb97c7a742}, baseFieldElem{0x1e54d443d0c8ebdf, 0x1634b4c26010c6af}},
			t: gfP2{baseFieldElem{0xb3982a9c7020c13a, 0x09d552f6a569e1c8}, baseFieldElem{0x514d4bddd0d2c463, 0x072a93f4bb18ce8f}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xb1a02c6f3cd7c191, 0x62232cae008f0b9f}, baseFieldElem{0x22ba8e66138eb448, 0x4c2c74e489fc5e32}},
			y: gfP2{baseFieldElem{0x930bed9a92420d5f, 0x3194eec351cc9805}, baseFieldElem{0xa9ce443a76af14ca, 0x33b6d278fb0dc303}},
			t: gfP2{baseFieldElem{0x13cf3e45a913f203, 0x21743b1ff177e206}, baseFieldElem{0x410a4bdd8ad323db, 0x4ab7dbc2dc2f6e3a}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		}, &point{
			x: gfP2{baseFieldElem{0xa02f390a51250c21, 0x5400391b4812b953}, baseFieldElem{0x05e58dd8a5d62d0c, 0x1f9619fc97f31562}},
			y: gfP2{baseFieldElem{0x326428f33d68994d, 0x76c4c99533163fc4}, baseFieldElem{0x1909d9fbab355924, 0x53e6186e7964194f}},
			t: gfP2{baseFieldElem{0x13ac5d1193aa1830, 0x16d15e40fca2a64f}, baseFieldElem{0x703197eca778e512, 0x59dca51723704f2c}},
			z: gfP2{baseFieldElem{0x1, 0x0}, baseFieldElem{0x0, 0x0}},
		},
	}

	// Constants exclusively for tests.
	p, _ = new(big.Int).SetString("170141183460469231731687303715884105727", 10)
)

func init() {
	feMul(&g.t, &g.x, &g.y)
	Gx, Gy = g.Int()
}
