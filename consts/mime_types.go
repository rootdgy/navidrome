package consts

import (
	"mime"
	"sort"
	"strings"
)

type format struct {
	typ      string
	lossless bool
}

var audioFormats = map[string]foirmat{
        ".3gp":        {typ: "audio/3gpp"},
        ".669":        {typ: "audio/x-mod"},
        ".726":        {typ: "audio/32kadpcm"},
        ".aa3":        {typ: "audio/ATRAC3"},
        ".aac":        {typ: "audio/x-aac"},
        ".aal":        {typ: "audio/ATRAC-ADVANCED-LOSSLESS"},
        ".ac3":        {typ: "audio/ac3"},
        ".acn":        {typ: "audio/asc"},
        ".adts":       {typ: "audio/aac"},
        ".aif":        {typ: "audio/x-aiff"},
        ".aifc":       {typ: "audio/x-aiff"},
        ".aiff":       {typ: "audio/x-aiff"},
        ".alac":       {typ: "audio/mp4", lossless: true},
        ".amr":        {typ: "audio/AMR"},
        ".ape":        {typ: "audio/x-monkeys-audio", lossless: true},
        ".ass":        {typ: "audio/aac"},
        ".at3":        {typ: "audio/ATRAC3"},
        ".atx":        {typ: "audio/ATRAC-X"},
        ".au":         {typ: "audio/basic"},
        ".awb":        {typ: "audio/AMR-WB"},
        ".axa":        {typ: "audio/x-annodex"},
        ".dls":        {typ: "audio/dls"},
        ".dsf":        {typ: "audio/dsd", lossless: true},
        ".dts":        {typ: "audio/vnd.dts"},
        ".dtshd":      {typ: "audio/vnd.dts.hd"},
        ".ecelp4800":  {typ: "audio/vnd.nuera.ecelp4800"},
        ".ecelp7470":  {typ: "audio/vnd.nuera.ecelp7470"},
        ".ecelp9600":  {typ: "audio/vnd.nuera.ecelp9600"},
        ".enw":        {typ: "audio/EVRCNW"},
        ".eol":        {typ: "audio/vnd.digital-winds"},
        ".evb":        {typ: "audio/EVRCB"},
        ".evc":        {typ: "audio/EVRC"},
        ".evw":        {typ: "audio/EVRCWB"},
        ".flac":       {typ: "audio/flac", lossless: true},
        ".kar":        {typ: "audio/midi"},
        ".koz":        {typ: "audio/vnd.audiokoz"},
        ".l16":        {typ: "audio/L16"},
        ".lbc":        {typ: "audio/iLBC"},
        ".loas":       {typ: "audio/usac"},
        ".lvp":        {typ: "audio/vnd.lucent.voice"},
        ".m15":        {typ: "audio/x-mod"},
        ".m3u":        {typ: "audio/x-mpegurl"},
        ".m4a":        {typ: "audio/mp4"},
        ".m4b":        {typ: "audio/mp4"},
        ".med":        {typ: "audio/x-mod"},
        ".mhas":       {typ: "audio/mhas"},
        ".mid":        {typ: "audio/midi"},
        ".midi":       {typ: "audio/midi"},
        ".mka":        {typ: "audio/x-matroska"},
        ".mlp":        {typ: "audio/vnd.dolby.mlp"},
        ".mod":        {typ: "audio/x-mod"},
        ".mp1":        {typ: "audio/mpeg"},
        ".mp2":        {typ: "audio/mpeg"},
        ".mp3":        {typ: "audio/mpeg"},
        ".mpc":        {typ: "audio/x-musepack"},
        ".mpga":       {typ: "audio/mpeg"},
        ".mtm":        {typ: "audio/x-mod"},
        ".multitrack": {typ: "audio/vnd.presonus.multitrack"},
        ".mxmf":       {typ: "audio/mobile-xmf"},
        ".oga":        {typ: "audio/ogg"},
        ".ogg":        {typ: "audio/ogg"},
        ".omg":        {typ: "audio/ATRAC3"},
        ".opus":       {typ: "audio/ogg"},
        ".plj":        {typ: "audio/vnd.everad.plj"},
        ".pls":        {typ: "audio/x-scpls"},
        ".psid":       {typ: "audio/prs.sid"},
        ".pya":        {typ: "audio/vnd.ms-playready.media.pya"},
        ".qcp":        {typ: "audio/QCELP"},
        ".ra":         {typ: "audio/x-realaudio"},
        ".ram":        {typ: "audio/x-pn-realaudio"},
        ".rip":        {typ: "audio/vnd.rip"},
        ".rm":         {typ: "audio/x-pn-realaudio"},
        ".s1m":        {typ: "audio/vnd.sealedmedia.softseal.mpeg"},
        ".s3m":        {typ: "audio/x-s3m"},
        ".shn":        {typ: "audio/x-shn", lossless: true},
        ".sid":        {typ: "audio/prs.sid"},
        ".smp":        {typ: "audio/vnd.sealedmedia.softseal.mpeg"},
        ".smp3":       {typ: "audio/vnd.sealedmedia.softseal.mpeg"},
        ".smv":        {typ: "audio/SMV"},
        ".snd":        {typ: "audio/basic"},
        ".sofa":       {typ: "audio/sofa"},
        ".spx":        {typ: "audio/ogg"},
        ".stm":        {typ: "audio/x-stm"},
        ".ult":        {typ: "audio/x-mod"},
        ".uni":        {typ: "audio/x-mod"},
        ".uva":        {typ: "audio/vnd.dece.audio"},
        ".uvva":       {typ: "audio/vnd.dece.audio"},
        ".vbk":        {typ: "audio/vnd.nortel.vbk"},
        ".wav":        {typ: "audio/x-wav", lossless: true},
        ".wax":        {typ: "audio/x-ms-wax"},
        ".wma":        {typ: "audio/x-ms-wma"},
        ".wv":         {typ: "audio/x-wavpack", lossless: true},
        ".wvp":        {typ: "audio/x-wavpack", lossless: true},
        ".xhe":        {typ: "audio/usac"},
}
var imageFormats = map[string]string{
	".gif":  "image/gif",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".webp": "image/webp",
	".png":  "image/png",
	".bmp":  "image/bmp",
}

var LosslessFormats []string

func init() {
	for ext, fmt := range audioFormats {
		_ = mime.AddExtensionType(ext, fmt.typ)
		if fmt.lossless {
			LosslessFormats = append(LosslessFormats, strings.TrimPrefix(ext, "."))
		}
	}
	sort.Strings(LosslessFormats)
	for ext, typ := range imageFormats {
		_ = mime.AddExtensionType(ext, typ)
	}

	// In some circumstances, Windows sets JS mime-type to `text/plain`!
	_ = mime.AddExtensionType(".js", "text/javascript")
	_ = mime.AddExtensionType(".css", "text/css")
}
