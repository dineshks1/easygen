package easygenapi

import "regexp"

var (
	re  *regexp.Regexp
	rec *regexp.Regexp
)

// TFStringsInit does initialization for strings related template functions
func TFStringsInit() {
	re = regexp.MustCompile(`(?i)` + Opts.StrFrom)
	// case sensitive string replace
	rec = regexp.MustCompile(Opts.StrFrom)

}

//==========================================================================
// template function

func replace(replStr string) string {
	return re.ReplaceAllString(replStr, Opts.StrTo)
}

// replacec does a case sensitive string replace
func replacec(replStr string) string {
	return rec.ReplaceAllString(replStr, Opts.StrTo)
}
