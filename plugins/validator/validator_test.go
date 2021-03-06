package validator

import (
	"fmt"

	"github.com/iotaledger/iota.go/address"
	. "github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/signing"
	. "github.com/iotaledger/iota.go/trinary"
)

const (
	exampleHash  = "999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	exampleSeed  = exampleHash
	exmapleIndex = 0
	exampleSec   = SecurityLevelLow
)

// Creates bundle signature fragments for the given address index and bundle hash.
// Each signature fragment after the first must go into its own meta transaction with value = 0.
func signature(seed Trytes, index uint64, sec SecurityLevel, bundleHash Hash) []Trytes {
	// compute seed based on address index
	subseed, _ := signing.Subseed(seed, index)
	// generate the private key
	prvKey, _ := signing.Key(subseed, sec)

	normalizedBundleHash := signing.NormalizedBundleHash(bundleHash)

	signatureFragments := make([]Trytes, sec)
	for i := 0; i < int(sec); i++ {
		// each security level signs one third of the (normalized) bundle hash
		signedFragTrits, _ := signing.SignatureFragment(
			normalizedBundleHash[i*HashTrytesSize/3:(i+1)*HashTrytesSize/3],
			prvKey[i*KeyFragmentLength:(i+1)*KeyFragmentLength],
		)
		signatureFragments[i] = MustTritsToTrytes(signedFragTrits)
	}

	return signatureFragments
}

func ExamplePLUGIN() {
	// corresponding address to validate against.
	addr, _ := address.GenerateAddress(exampleSeed, exmapleIndex, exampleSec)
	fmt.Println(addr)

	// compute the signature fragments which would be added to the (meta) transactions
	signatureFragments := signature(exampleSeed, exmapleIndex, exampleSec, exampleHash)
	fmt.Println(signatureFragments[0])

	// Output:
	// BSIXFJENGVJSOWPVHVALMPOPO9PUKHXDQI9VDELCBJXN9TCNQPTFEDMPQCVBOJSZUHEOABYYYAT9IAHHY
	// GHHKPBXOOBOEHGGEEKYPH9MANWEKSQTQJFJ9KUTMJQAVITYRZMNLUESQARNHAWUJAPPZSQ9A9RUKABCE9KZPJDUEHVZEOSCQMTCC9AWBGWZLZEXMJ9YOQUVIBGMXSINCOLUATYDDUBAALHCBIONNRQIVIPUFPOIFHYRBFBGXXNVYXFZUSTTA9LYGGITTAJCVDE9GCFRGIOTXLQ9ZJDLONDLZ9OPS9TNYVKLTCGFBH9QPJWLIGADWMTJVCLAUCOZFDSRRCAMVWYFXRPGPMIOPIW9GBWANVSMPONQOTNLLYYHXAMZMMNRHMRXHEIXPVNORNGZZ9ZAU9RAWASOZNIBKDWYZWKCMLEUE9UVDHZ9XXGPXZABB9FGTNDTDFTYCKLKRRC9GZFKHKDGAWPBWEUPPWISYBBNZCIBERPXTMZPZHPKKUQUPBIJBIKZAGFHDDNAGCRQMWOMLUMAYKRBMHPMDWZK9JRBDWCJCBJQYMDUBNKOIRSJSVTCNKROZ9KLFBZLOXQOASLCFETCNZRPZULOABOFCUO9WKNQILLLTQ9GWVDBASBGSKUHFHRXOKQIBRCLUYZBZMTXTIG9BJNYHTJQQOECXOWLIDOYKMFJWKRCYW99VZILSPU9I9ZSTTBZVGISUHPCWLGKCFNLIHJNCL9OWQDNAKJAGRKTGCTDRHXVAYXOHNFVJYBMZLMXV9VINNIAWONYDYOKHHMOFFEOOVBMVMYABWRWLZTWJECKKAGPCIMUDZZIEJCFBXFIYKDRMWZIOEUZNLOXZJRDHVVKOTJWMLTIXVIRJSXUBLFGOCCLEIZVCDYD9FEMCRUOERPRDFGUJSALRSOBN9J9XDTUAJZFLHUGQI9MCXZCYWTTIHNQUPUYPDRJLRZG9HAXHYQDSSCQNPTBYKNQUWZDE9QUESZJASRXHNW9OKAVUKLLMVGOJJRZCPRXSRYUECLNQEFIHI9S9NNEN9KACVIKCZYDEKCDNUASUJWMTVLSPBOBQMQEMZJXJVQAMUGBTMNWEWVJSXNZKIAADSQCCLISYSUZICSIVXZUG9MTICGWXKXKJDW9TOUBS9BTOUFUKWEBVIIJTGD9IBLRHBCPICWSZQNJQERTBOZGLJFCXKGQTAHIWKOSGHRMMWXABQYHVHOPG9XDIXMIRBXHOSYBCHSFWORNLUD9JAB9ICBIPXYVLIXYNRHJVEDMIRSAGXKZKSFZADJ9GA9DGJZAJTXZGIKRXVBCCBGJPJWJJZXZRQNWLEUZEFTWOXUBTAGDPPKKPKRYPGXVSRWLRNEDAXHZYT9DRN9L9ZWXPTTOSKMGTPQQXHACAKESRQXVXXNOLIATRKDGGJNIDWWYKQSLTC9ERTPMNXQHZNVNSBGIRRQHMOCOGDWPQAU9WPRSGZMPXZWQADUFUAWVGESLIWZNV9WNANDMZAOLXIHAOSFBADWVVAHMJVFNX9BGMMYGMJCUOYCSKJWIUMYHQFQXCFQXQNB9VTBLAYGKUZLFH9UVWIQJVLMLOZDLLIPJZSNXBPWAKKZWKCVSWUSBSQLBIAX9SQGMNPCJWTQDQEASSWWCSTVJRFDBPBLNYU9CNFUYINVMQPJZGKKUH9QBMUVWFSLPXWKBBWKNLMHGCEMJWCTNXZYWCFXYU9XLTWDSROJDTCRARMBNYDDD99HCFMXMUCO9NJSRA9G9HGWRTWNDBDQLBTCNYIVRMWRWPDJDDYCDODGEBNFTNINPNMZYMJJHVZSNEIJOAPGHAIVCZHQIULTRIZ9ML9LCWTQVGLBKKBGJYZTOZZIYUBCBKHKYUHCFGZKDERTWYHNYWSWLGPUGRB9WNQTHOMBFPKUQZREUQCNXL9MFSZCNBN9PTAVCERMWTTFDZL9BJQMC9OUBWGDTURAEYTYRDNFUBATOWFSVNXJC9JUPARMU9MINY9RWRHIXBPNIUADFAEP9F9FWNJNRPNGLWHRYYCV9ZIWBOUZPFZTWDLOCNOYZQLWFJHZ99ZBLUDSIQBJOJXMQJBUCYYMROBCJJJNCETVUYRXKHAWGUBIWOKQXOIOYBQKNDXZCKXQZLWEMXYLJPODRMOQUYOAATZZQ9JZDR9KPIHRQKIEAQNO9OVXNHDFCUUIZRQDWYGKUAYIGHGIIJIOIERLVNDUEBZUAQGDZMWNGXQPYSNWUEGF9BQDFJEQRPEGFGJTQFWO9PWECFGNDH9LW
}
