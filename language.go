package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

// Lang is a string representation of a two character language string.
type Lang string

const (
	// LangDefault is set to 'en'.
	LangDefault Lang = LanguageEN
	// LanguageAA is a string language key.
	LanguageAA Lang = "aa"
	// LanguageAB is a string language key.
	LanguageAB Lang = "ab"
	// LanguageAE is a string language key.
	LanguageAE Lang = "ae"
	// LanguageAF is a string language key.
	LanguageAF Lang = "af"
	// LanguageAK is a string language key.
	LanguageAK Lang = "ak"
	// LanguageAM is a string language key.
	LanguageAM Lang = "am"
	// LanguageAN is a string language key.
	LanguageAN Lang = "an"
	// LanguageAR is a string language key.
	LanguageAR Lang = "ar"
	// LanguageAS is a string language key.
	LanguageAS Lang = "as"
	// LanguageAV is a string language key.
	LanguageAV Lang = "av"
	// LanguageAY is a string language key.
	LanguageAY Lang = "ay"
	// LanguageAZ is a string language key.
	LanguageAZ Lang = "az"
	// LanguageBA is a string language key.
	LanguageBA Lang = "ba"
	// LanguageBE is a string language key.
	LanguageBE Lang = "be"
	// LanguageBG is a string language key.
	LanguageBG Lang = "bg"
	// LanguageBH is a string language key.
	LanguageBH Lang = "bh"
	// LanguageBI is a string language key.
	LanguageBI Lang = "bi"
	// LanguageBM is a string language key.
	LanguageBM Lang = "bm"
	// LanguageBN is a string language key.
	LanguageBN Lang = "bn"
	// LanguageBO is a string language key.
	LanguageBO Lang = "bo"
	// LanguageBR is a string language key.
	LanguageBR Lang = "br"
	// LanguageBS is a string language key.
	LanguageBS Lang = "bs"
	// LanguageCA is a string language key.
	LanguageCA Lang = "ca"
	// LanguageCE is a string language key.
	LanguageCE Lang = "ce"
	// LanguageCH is a string language key.
	LanguageCH Lang = "ch"
	// LanguageCO is a string language key.
	LanguageCO Lang = "co"
	// LanguageCR is a string language key.
	LanguageCR Lang = "cr"
	// LanguageCS is a string language key.
	LanguageCS Lang = "cs"
	// LanguageCU is a string language key.
	LanguageCU Lang = "cu"
	// LanguageCV is a string language key.
	LanguageCV Lang = "cv"
	// LanguageCY is a string language key.
	LanguageCY Lang = "cy"
	// LanguageDA is a string language key.
	LanguageDA Lang = "da"
	// LanguageDE is a string language key.
	LanguageDE Lang = "de"
	// LanguageDV is a string language key.
	LanguageDV Lang = "dv"
	// LanguageDZ is a string language key.
	LanguageDZ Lang = "dz"
	// LanguageEE is a string language key.
	LanguageEE Lang = "ee"
	// LanguageEL is a string language key.
	LanguageEL Lang = "el"
	// LanguageEN is a string language key.
	LanguageEN Lang = "en"
	// LanguageEO is a string language key.
	LanguageEO Lang = "eo"
	// LanguageES is a string language key.
	LanguageES Lang = "es"
	// LanguageET is a string language key.
	LanguageET Lang = "et"
	// LanguageEU is a string language key.
	LanguageEU Lang = "eu"
	// LanguageFA is a string language key.
	LanguageFA Lang = "fa"
	// LanguageFF is a string language key.
	LanguageFF Lang = "ff"
	// LanguageFI is a string language key.
	LanguageFI Lang = "fi"
	// LanguageFJ is a string language key.
	LanguageFJ Lang = "fj"
	// LanguageFO is a string language key.
	LanguageFO Lang = "fo"
	// LanguageFR is a string language key.
	LanguageFR Lang = "fr"
	// LanguageFY is a string language key.
	LanguageFY Lang = "fy"
	// LanguageGA is a string language key.
	LanguageGA Lang = "ga"
	// LanguageGD is a string language key.
	LanguageGD Lang = "gd"
	// LanguageGL is a string language key.
	LanguageGL Lang = "gl"
	// LanguageGN is a string language key.
	LanguageGN Lang = "gn"
	// LanguageGU is a string language key.
	LanguageGU Lang = "gu"
	// LanguageGV is a string language key.
	LanguageGV Lang = "gv"
	// LanguageHA is a string language key.
	LanguageHA Lang = "ha"
	// LanguageHE is a string language key.
	LanguageHE Lang = "he"
	// LanguageHI is a string language key.
	LanguageHI Lang = "hi"
	// LanguageHO is a string language key.
	LanguageHO Lang = "ho"
	// LanguageHR is a string language key.
	LanguageHR Lang = "hr"
	// LanguageHT is a string language key.
	LanguageHT Lang = "ht"
	// LanguageHU is a string language key.
	LanguageHU Lang = "hu"
	// LanguageHY is a string language key.
	LanguageHY Lang = "hy"
	// LanguageHZ is a string language key.
	LanguageHZ Lang = "hz"
	// LanguageIA is a string language key.
	LanguageIA Lang = "ia"
	// LanguageID is a string language key.
	LanguageID Lang = "id"
	// LanguageIE is a string language key.
	LanguageIE Lang = "ie"
	// LanguageIG is a string language key.
	LanguageIG Lang = "ig"
	// LanguageII is a string language key.
	LanguageII Lang = "ii"
	// LanguageIK is a string language key.
	LanguageIK Lang = "ik"
	// LanguageIO is a string language key.
	LanguageIO Lang = "io"
	// LanguageIS is a string language key.
	LanguageIS Lang = "is"
	// LanguageIT is a string language key.
	LanguageIT Lang = "it"
	// LanguageIU is a string language key.
	LanguageIU Lang = "iu"
	// LanguageJA is a string language key.
	LanguageJA Lang = "ja"
	// LanguageJV is a string language key.
	LanguageJV Lang = "jv"
	// LanguageKA is a string language key.
	LanguageKA Lang = "ka"
	// LanguageKG is a string language key.
	LanguageKG Lang = "kg"
	// LanguageKI is a string language key.
	LanguageKI Lang = "ki"
	// LanguageKJ is a string language key.
	LanguageKJ Lang = "kj"
	// LanguageKK is a string language key.
	LanguageKK Lang = "kk"
	// LanguageKL is a string language key.
	LanguageKL Lang = "kl"
	// LanguageKM is a string language key.
	LanguageKM Lang = "km"
	// LanguageKN is a string language key.
	LanguageKN Lang = "kn"
	// LanguageKO is a string language key.
	LanguageKO Lang = "ko"
	// LanguageKR is a string language key.
	LanguageKR Lang = "kr"
	// LanguageKS is a string language key.
	LanguageKS Lang = "ks"
	// LanguageKU is a string language key.
	LanguageKU Lang = "ku"
	// LanguageKV is a string language key.
	LanguageKV Lang = "kv"
	// LanguageKW is a string language key.
	LanguageKW Lang = "kw"
	// LanguageKY is a string language key.
	LanguageKY Lang = "ky"
	// LanguageLA is a string language key.
	LanguageLA Lang = "la"
	// LanguageLB is a string language key.
	LanguageLB Lang = "lb"
	// LanguageLG is a string language key.
	LanguageLG Lang = "lg"
	// LanguageLI is a string language key.
	LanguageLI Lang = "li"
	// LanguageLN is a string language key.
	LanguageLN Lang = "ln"
	// LanguageLO is a string language key.
	LanguageLO Lang = "lo"
	// LanguageLT is a string language key.
	LanguageLT Lang = "lt"
	// LanguageLU is a string language key.
	LanguageLU Lang = "lu"
	// LanguageLV is a string language key.
	LanguageLV Lang = "lv"
	// LanguageMG is a string language key.
	LanguageMG Lang = "mg"
	// LanguageMH is a string language key.
	LanguageMH Lang = "mh"
	// LanguageMI is a string language key.
	LanguageMI Lang = "mi"
	// LanguageMK is a string language key.
	LanguageMK Lang = "mk"
	// LanguageML is a string language key.
	LanguageML Lang = "ml"
	// LanguageMN is a string language key.
	LanguageMN Lang = "mn"
	// LanguageMR is a string language key.
	LanguageMR Lang = "mr"
	// LanguageMS is a string language key.
	LanguageMS Lang = "ms"
	// LanguageMT is a string language key.
	LanguageMT Lang = "mt"
	// LanguageMY is a string language key.
	LanguageMY Lang = "my"
	// LanguageNA is a string language key.
	LanguageNA Lang = "na"
	// LanguageNB is a string language key.
	LanguageNB Lang = "nb"
	// LanguageND is a string language key.
	LanguageND Lang = "nd"
	// LanguageNE is a string language key.
	LanguageNE Lang = "ne"
	// LanguageNG is a string language key.
	LanguageNG Lang = "ng"
	// LanguageNL is a string language key.
	LanguageNL Lang = "nl"
	// LanguageNN is a string language key.
	LanguageNN Lang = "nn"
	// LanguageNO is a string language key.
	LanguageNO Lang = "no"
	// LanguageNR is a string language key.
	LanguageNR Lang = "nr"
	// LanguageNV is a string language key.
	LanguageNV Lang = "nv"
	// LanguageNY is a string language key.
	LanguageNY Lang = "ny"
	// LanguageOC is a string language key.
	LanguageOC Lang = "oc"
	// LanguageOJ is a string language key.
	LanguageOJ Lang = "oj"
	// LanguageOM is a string language key.
	LanguageOM Lang = "om"
	// LanguageOR is a string language key.
	LanguageOR Lang = "or"
	// LanguageOS is a string language key.
	LanguageOS Lang = "os"
	// LanguagePA is a string language key.
	LanguagePA Lang = "pa"
	// LanguagePI is a string language key.
	LanguagePI Lang = "pi"
	// LanguagePL is a string language key.
	LanguagePL Lang = "pl"
	// LanguagePS is a string language key.
	LanguagePS Lang = "ps"
	// LanguagePT is a string language key.
	LanguagePT Lang = "pt"
	// LanguageQU is a string language key.
	LanguageQU Lang = "qu"
	// LanguageRM is a string language key.
	LanguageRM Lang = "rm"
	// LanguageRN is a string language key.
	LanguageRN Lang = "rn"
	// LanguageRO is a string language key.
	LanguageRO Lang = "ro"
	// LanguageRU is a string language key.
	LanguageRU Lang = "ru"
	// LanguageRW is a string language key.
	LanguageRW Lang = "rw"
	// LanguageSA is a string language key.
	LanguageSA Lang = "sa"
	// LanguageSC is a string language key.
	LanguageSC Lang = "sc"
	// LanguageSD is a string language key.
	LanguageSD Lang = "sd"
	// LanguageSE is a string language key.
	LanguageSE Lang = "se"
	// LanguageSG is a string language key.
	LanguageSG Lang = "sg"
	// LanguageSI is a string language key.
	LanguageSI Lang = "si"
	// LanguageSK is a string language key.
	LanguageSK Lang = "sk"
	// LanguageSL is a string language key.
	LanguageSL Lang = "sl"
	// LanguageSM is a string language key.
	LanguageSM Lang = "sm"
	// LanguageSN is a string language key.
	LanguageSN Lang = "sn"
	// LanguageSO is a string language key.
	LanguageSO Lang = "so"
	// LanguageSQ is a string language key.
	LanguageSQ Lang = "sq"
	// LanguageSR is a string language key.
	LanguageSR Lang = "sr"
	// LanguageSS is a string language key.
	LanguageSS Lang = "ss"
	// LanguageST is a string language key.
	LanguageST Lang = "st"
	// LanguageSU is a string language key.
	LanguageSU Lang = "su"
	// LanguageSV is a string language key.
	LanguageSV Lang = "sv"
	// LanguageSW is a string language key.
	LanguageSW Lang = "sw"
	// LanguageTA is a string language key.
	LanguageTA Lang = "ta"
	// LanguageTE is a string language key.
	LanguageTE Lang = "te"
	// LanguageTG is a string language key.
	LanguageTG Lang = "tg"
	// LanguageTH is a string language key.
	LanguageTH Lang = "th"
	// LanguageTI is a string language key.
	LanguageTI Lang = "ti"
	// LanguageTK is a string language key.
	LanguageTK Lang = "tk"
	// LanguageTL is a string language key.
	LanguageTL Lang = "tl"
	// LanguageTN is a string language key.
	LanguageTN Lang = "tn"
	// LanguageTO is a string language key.
	LanguageTO Lang = "to"
	// LanguageTR is a string language key.
	LanguageTR Lang = "tr"
	// LanguageTS is a string language key.
	LanguageTS Lang = "ts"
	// LanguageTT is a string language key.
	LanguageTT Lang = "tt"
	// LanguageTW is a string language key.
	LanguageTW Lang = "tw"
	// LanguageTY is a string language key.
	LanguageTY Lang = "ty"
	// LanguageUG is a string language key.
	LanguageUG Lang = "ug"
	// LanguageUK is a string language key.
	LanguageUK Lang = "uk"
	// LanguageUR is a string language key.
	LanguageUR Lang = "ur"
	// LanguageUZ is a string language key.
	LanguageUZ Lang = "uz"
	// LanguageVE is a string language key.
	LanguageVE Lang = "ve"
	// LanguageVI is a string language key.
	LanguageVI Lang = "vi"
	// LanguageVO is a string language key.
	LanguageVO Lang = "vo"
	// LanguageWA is a string language key.
	LanguageWA Lang = "wa"
	// LanguageWO is a string language key.
	LanguageWO Lang = "wo"
	// LanguageXH is a string language key.
	LanguageXH Lang = "xh"
	// LanguageYI is a string language key.
	LanguageYI Lang = "yi"
	// LanguageYO is a string language key.
	LanguageYO Lang = "yo"
	// LanguageZA is a string language key.
	LanguageZA Lang = "za"
	// LanguageZH is a string language key.
	LanguageZH Lang = "zh"
	// LanguageZU is a string language key.
	LanguageZU Lang = "zu"
)
