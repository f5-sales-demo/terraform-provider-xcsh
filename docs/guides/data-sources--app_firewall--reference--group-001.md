---
page_title: "xcsh_app_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall reference."
---

# xcsh_app_firewall reference

<a id="canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- Property reference

<a id="canonical-2100010313230122-1130210021332210-0010201122330102-2201202112330032-2003213323302131-0203102101100102-0123121100300132-2130213220001103"></a>

### Direct properties for `xcsh_app_firewall`

- [allow_all_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-2313001313310012-0301210321202031-0221010102220023-1012300213220002-2123102221223031-1002013330013010-1113000210233221-2111103222220013): complete subsection reference.

- [allowed_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-2120023301001102-2103000021001323-3230113311223023-3220101220122000-2300110311220000-2121131310320100-2130312102313212-2200031122131102): complete subsection reference.

<a id="canonical-1103120022320213-3321332333211220-3030120220223101-2011201222121213-3211233301001000-2321003330113221-0231011130130001-3231221323133010"></a>

<a id="canonical-1303200322003033-0100320121311212-0200223221231220-2111122113110003-0033100132031303-1122230233000233-1031123100012103-0300020133033331"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [blocking](data-sources--app_firewall--reference--group-001.md#canonical-2330310310111223-0011110013221003-3130133233112122-0132133202311033-3310031301330023-2331222213303332-3101323002031013-1312223120332013): complete subsection reference.

- [blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-2320323133020103-0202323333132023-0231333131131330-3210000030111102-1100133100001232-3201112321100101-3321011100332101-3233030031031031): complete subsection reference.

- [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-2103022221032002-3031232012310300-2201020202303231-0102202313031033-2021133200010103-0330231303130020-0330211202222110-3021203231003220): complete subsection reference.

- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0023020222013102-0220221233101131-3310023020113202-2203132000300231-3310200111300122-1222233233313303-0131012223310232-1211132213211011): complete subsection reference.

- [default_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-1230333210011113-2001223230211300-0100320022133302-3121233203102311-2232332031011332-3022022012030332-0013113012000333-0101213203220220): complete subsection reference.

- [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-1202031113322310-1322320330322332-3311330222223232-3111201232203010-1010320232321303-3212123031302222-3130302330003303-2102013312033211): complete subsection reference.

- [default_detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-1130323331212003-3202232330332222-1311210212001102-3320222310010120-0021301132303300-1031210102102300-0203213003111132-1321211232221121): complete subsection reference.

<a id="canonical-1101012322302212-1223223233313101-0302303131120022-0020030223002220-3112031210302033-0233201130010033-0230132221133301-0233333330313300"></a>

<a id="canonical-1230121020003133-0121202011300312-3220131233333222-2232200321200312-2211120111111303-2212213200210330-3200030210211030-0131230312333023"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AppFirewall.

Additional upstream details:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210): complete subsection reference.

- [disable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-1010320003113321-0320300100231200-3321123033130211-0210221201100101-2302312001221203-3112323102130130-0302223322012101-2113111130031003): complete subsection reference.

- [disable_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-3303321323122022-3010001011230002-0110201311133021-3020110232333222-3002332123012032-1313132030122010-3101221013302031-1232201001022322): complete subsection reference.

- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-3330323012211103-0031131130300122-2011021101121023-3112012123021123-3231300122313331-1101212033212032-3302322233212232-3200321220223211): complete subsection reference.

<a id="canonical-0022021301320033-3221220100001220-3301121103132122-2233000131210301-2302211301123210-2120323231122001-0301130003311012-1332033112131022"></a>

<a id="canonical-1230002100001122-3102221113022232-3021121302212022-3212021122321330-1121320000132031-0330202010200231-1311131210002302-0033112312220301"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1031002310003010-0112310033020003-0023102133021022-3011012210102101-2222032231000220-0222221021302123-2001230331303321-2212301112020212"></a>

<a id="canonical-1032033001033120-3131131122222220-3202120122332033-1201031100011301-2210022210330213-0020213021312211-2213330031203312-1202010232322123"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [monitoring](data-sources--app_firewall--reference--group-001.md#canonical-0132332001110103-0010201300032310-0100330312110132-3301223331032312-2131202331131103-3320010120230110-0301101202322330-0001011321201113): complete subsection reference.

<a id="canonical-1023012033230133-2122330003301100-2020011232101110-1112123231032302-3203201011123003-2011103222233303-1113330022310321-2111001101103022"></a>

<a id="canonical-1222232230003301-3120313223311201-1131210232320030-0131220203312232-1122223012003000-0121222032233310-0002311001102220-1210032033101122"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AppFirewall.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1033221303020330-2323310013203113-0221133323302123-0123031202012013-0221121111223321-1200013111213011-2303213232013013-1220111320110321"></a>

<a id="canonical-3131223320012020-1110123100230310-1023012020133310-3221120020310202-3032322323031302-0033133023210011-2033311203323313-3022323022132222"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AppFirewall exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_default_blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-2133021012333021-2112221112021001-1232130013231113-1033201031303202-2010021000212113-3231101322123331-0123233300232303-3310130211333330): complete subsection reference.

<a id="canonical-0023003211012223-3230303310311113-0220011121110301-0231220112333131-0101012021021002-0302131223202122-3003033031120231-3222232220031331"></a>

### All schema paths for `xcsh_app_firewall`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_response_codes` | [allow_all_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-2030120232132200-0132203113121330-2320200323121013-2131002201103102-2021302233020212-2011033322213113-0020223131013022-3123103202001323) |
| `allowed_response_codes` | [allowed_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-2322002103222111-1012031013010111-3011232330032003-2003333132321322-1032123201021120-1321312313001000-1000200232131313-0212312011210101) |
| `allowed_response_codes.response_code` | [allowed_response_codes.response_code](data-sources--app_firewall--reference--group-001.md#canonical-0320022303322203-2233030003301233-3202311333213102-2212021121102313-1221002213313320-2012002302230001-2000003233000110-1230033120033202) |
| `annotations` | [annotations](data-sources--app_firewall--reference--group-001.md#canonical-1103120022320213-3321332333211220-3030120220223101-2011201222121213-3211233301001000-2321003330113221-0231011130130001-3231221323133010) |
| `blocking` | [blocking](data-sources--app_firewall--reference--group-001.md#canonical-0300002133113213-2203210210200113-2310222132003003-2312303203213212-2230321122133022-0331302012120031-2231320012103012-3222312013303302) |
| `blocking_page` | [blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-0223133111001102-3313132310303000-3333312020123332-3231101000003330-2223021110110231-2113320202021112-3101020233300323-2122333003020021) |
| `blocking_page.blocking_page` | [blocking_page.blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-0003323302321032-2330122303201122-1233103120022121-2030022222030113-1112113001030210-3000021311321310-3332202102200320-2332222312011032) |
| `blocking_page.response_code` | [blocking_page.response_code](data-sources--app_firewall--reference--group-001.md#canonical-0331031200210323-2310311003323212-0322020231210020-2312313003203132-0210303122033010-0120030120023203-2121322302232111-0332100000121022) |
| `bot_protection_setting` | [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3002231223212103-0011012011120012-0322120001221313-0211103031210213-2310300220230302-0010300230222001-3223230220110011-3010330320112233) |
| `bot_protection_setting.good_bot_action` | [bot_protection_setting.good_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-1301120112333110-0200213211300221-2120232220012010-0120032200001102-2111212323221220-0221333132322200-1301301123221130-3131030120201213) |
| `bot_protection_setting.malicious_bot_action` | [bot_protection_setting.malicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-1010220011202202-1330211212021102-2112002313332011-3220002202000223-0302212331303331-0310331023323033-2130221212103313-1222203121112113) |
| `bot_protection_setting.suspicious_bot_action` | [bot_protection_setting.suspicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-0020102231332320-0101200112013003-0211011032112232-1120122131220213-3031322113103230-3213201130131202-0203110202021002-0121113020100010) |
| `custom_anonymization` | [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-2012101332111332-3233220203102320-0220330113310320-2321100332212313-3002200311130022-0210001120121312-2323023230312131-1311232202332231) |
| `custom_anonymization.anonymization_config` | [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-0020201313133032-3213003220030023-3331022032320132-2230323203331333-1120333101323222-1112333221021322-0001231020032112-1200130131113031) |
| `custom_anonymization.anonymization_config.cookie` | [custom_anonymization.anonymization_config.cookie](data-sources--app_firewall--reference--group-001.md#canonical-0300322321001231-3310130300220101-1323120212111210-0030212231321102-0010003112121303-0130232123102221-1112130210132321-2031133202032121) |
| `custom_anonymization.anonymization_config.cookie.cookie_name` | [custom_anonymization.anonymization_config.cookie.cookie_name](data-sources--app_firewall--reference--group-001.md#canonical-1230312123301131-0220102033323302-2031222300121031-3223003031211231-1231220301301021-1213030012133132-0111032020220311-3113012130010121) |
| `custom_anonymization.anonymization_config.http_header` | [custom_anonymization.anonymization_config.http_header](data-sources--app_firewall--reference--group-001.md#canonical-0100213322023322-0322311112033333-3213120120000333-0122003032231100-3021033320301330-2120122112333200-1210333323231011-2022323200011120) |
| `custom_anonymization.anonymization_config.http_header.header_name` | [custom_anonymization.anonymization_config.http_header.header_name](data-sources--app_firewall--reference--group-001.md#canonical-0112330032302213-2030101203002303-2231010302122232-0320011020122111-2200013233320330-3223220021122031-1020233331100012-2121202112322112) |
| `custom_anonymization.anonymization_config.query_parameter` | [custom_anonymization.anonymization_config.query_parameter](data-sources--app_firewall--reference--group-001.md#canonical-3020202223103303-2202202113313132-1211202100212312-1321023313133233-3031030230233123-2033231010021231-2030022120012221-3301202223012113) |
| `custom_anonymization.anonymization_config.query_parameter.query_param_name` | [custom_anonymization.anonymization_config.query_parameter.query_param_name](data-sources--app_firewall--reference--group-001.md#canonical-1322223200012200-1203002202223311-0233131000300203-3113032030120112-3021201302310221-0113012122031220-1122013222300131-0101111013001211) |
| `default_anonymization` | [default_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-3033312130321301-2011101021301322-1202031001023100-3100200110303123-2120322321020030-0210313213001200-2303120101133223-2323233222030223) |
| `default_bot_setting` | [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-1032213023131131-2233121312222130-1120233323213131-2013222202113103-3332102200322100-1332031020222110-2132130300031231-1112121020002130) |
| `default_detection_settings` | [default_detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-0000113131131211-1032200030131323-2130321012123211-2012022313302202-2313202232332103-0120303201132300-3012221313303200-3303212221130301) |
| `description` | [description](data-sources--app_firewall--reference--group-001.md#canonical-1101012322302212-1223223233313101-0302303131120022-0020030223002220-3112031210302033-0233201130010033-0230132221133301-0233333330313300) |
| `detection_settings` | [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-0021023012003132-1331020220200333-2223201122312223-2222023303313320-2230221233233330-2031012111301103-2131103032231203-2023033212112012) |
| `detection_settings.bot_protection_setting` | [detection_settings.bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3110220100211021-0213120200001131-2002022002123201-1213133220012113-1021022312022322-1003211221213322-2131222323110233-2210202210113101) |
| `detection_settings.bot_protection_setting.good_bot_action` | [detection_settings.bot_protection_setting.good_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-2121023332010203-2021132001311312-2122212023322303-0021133213101121-1022123012033031-0220332001123132-2332011003300010-2003313023031030) |
| `detection_settings.bot_protection_setting.malicious_bot_action` | [detection_settings.bot_protection_setting.malicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-2011123021113231-3023333000233311-1002200133100313-2000321001122100-3313321023131231-2212221332033313-3300113310012303-2322133103132130) |
| `detection_settings.bot_protection_setting.suspicious_bot_action` | [detection_settings.bot_protection_setting.suspicious_bot_action](data-sources--app_firewall--reference--group-001.md#canonical-3300113322102011-1133111233120312-0003303033122103-3020320013123203-2331213303320131-3303122313021132-1331010123110220-0222333103302020) |
| `detection_settings.default_bot_setting` | [detection_settings.default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-1201221103023030-1131011011121220-2020300032011202-3303003013301200-1120023311223032-1012200302331121-1320021211210310-0313333300123321) |
| `detection_settings.default_violation_settings` | [detection_settings.default_violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-2212222301011001-2022120321331102-3312101023200230-2002200301113032-0313101120311313-1121210312230121-2100103011210032-1111113303333201) |
| `detection_settings.disable_staging` | [detection_settings.disable_staging](data-sources--app_firewall--reference--group-001.md#canonical-0110321030123220-3012102103010223-3201321232330110-3033333111112030-0200210301000213-0332001211103012-3210011310102030-2133001310322010) |
| `detection_settings.disable_suppression` | [detection_settings.disable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-2133001213000231-3302213013002323-1301013123230012-2210322110230021-0101233102121002-1303020333231010-2230120033110303-2323133310012120) |
| `detection_settings.disable_threat_campaigns` | [detection_settings.disable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-2302030221232311-3100123021110102-2130211020302220-2022210132300023-1022300313111100-1330022133020230-3231122003001010-1121221210100012) |
| `detection_settings.enable_suppression` | [detection_settings.enable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-3313033320130201-3223012101323001-0221132013301102-2121222132332013-0102102300330102-3111203031120131-3331122301001103-1313212031213230) |
| `detection_settings.enable_threat_campaigns` | [detection_settings.enable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-3010313021102222-3113203303203011-1022131130123223-3100331321023210-1001322000312203-2021111123302312-1323303013023200-0212323122311222) |
| `detection_settings.signature_selection_setting` | [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-1110121301000301-2022101211330021-3230310023223111-1303230332202002-3313203132231023-0032010102101310-3110222331123000-2202323213002211) |
| `detection_settings.signature_selection_setting.attack_type_settings` | [detection_settings.signature_selection_setting.attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-3330312130103121-0113011132223030-1011111301320113-3332230213132031-3300311132230213-0200301120111112-1233011223301232-2212313233123102) |
| `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` | [detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types](data-sources--app_firewall--reference--group-001.md#canonical-2021031200301220-3300112332330330-3300310331230203-3201033113012132-1222010230001010-3033232323113102-1223011202132310-1303030021210213) |
| `detection_settings.signature_selection_setting.default_attack_type_settings` | [detection_settings.signature_selection_setting.default_attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-2213103100022023-2223220212033013-1221222100133013-1013013212131023-2230022221202022-3021332223222213-3022103110302331-0201303001112131) |
| `detection_settings.signature_selection_setting.default_signature_setting` | [detection_settings.signature_selection_setting.default_signature_setting](data-sources--app_firewall--reference--group-001.md#canonical-3301122032303132-1222212031010000-3210213110010030-3322102232032330-0112220113220313-0213113032302013-1223233003002130-1030020303013223) |
| `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-2332133310112020-2131113200012221-3323020331232220-2221031300010311-0221332001211131-2001203231030120-1231232200233110-0320322121011321) |
| `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-3320122120012101-3320323331132202-1033220000113023-3023000122013223-2033031003000301-3023331211311002-1210303202102232-0112021302300331) |
| `detection_settings.signature_selection_setting.only_high_accuracy_signatures` | [detection_settings.signature_selection_setting.only_high_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-0211221033311230-1003012022230001-0110000123203300-0213200012101033-2130031302022332-3230210011331302-3220111121302310-2230323323113320) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy](data-sources--app_firewall--reference--group-001.md#canonical-3000233222232323-2020200211122111-2103312200200320-1130330212132210-3001322310213001-2001301202022200-1330201323312003-1002222330202120) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action](data-sources--app_firewall--reference--group-001.md#canonical-1111233023211301-0210102322322012-1032331322020131-3200300022133331-2103003302102032-3010220202211230-2013023130302213-1232033311130013) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action](data-sources--app_firewall--reference--group-001.md#canonical-2002302102032300-0332333032103201-1133122003312122-1330123222131323-3201320102323320-0130332001030023-1031032313121021-1201112101302030) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action](data-sources--app_firewall--reference--group-001.md#canonical-1111221210233300-2311102131331013-1223320133300132-3331323231130223-2202121112231201-0203223133131031-1332012013233310-1310033223233131) |
| `detection_settings.stage_new_and_updated_signatures` | [detection_settings.stage_new_and_updated_signatures](data-sources--app_firewall--reference--group-001.md#canonical-2201301213222212-2231213131333210-1312323123123201-2133300021013131-3010011101321210-3333022002010101-3222310121330331-1310233122301003) |
| `detection_settings.stage_new_and_updated_signatures.staging_period` | [detection_settings.stage_new_and_updated_signatures.staging_period](data-sources--app_firewall--reference--group-001.md#canonical-2331322200201221-0022322221232311-2003011021303213-3222133213233232-1332100202321003-0112020132133202-1033113121131330-0311232320321321) |
| `detection_settings.stage_new_signatures` | [detection_settings.stage_new_signatures](data-sources--app_firewall--reference--group-001.md#canonical-0230303201022123-1222030230130302-3032303320023001-0022331112333023-1230131032210202-0222112010103320-3302232323231120-3323213200203221) |
| `detection_settings.stage_new_signatures.staging_period` | [detection_settings.stage_new_signatures.staging_period](data-sources--app_firewall--reference--group-001.md#canonical-1231203222031221-2030031232201100-3100121330213021-3130131222130121-1312021222013322-3333100002222223-0011030222323110-0123223302330003) |
| `detection_settings.violation_settings` | [detection_settings.violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-1213010100321222-3300210220300200-3101021102222212-0331303202020330-1123221012210020-1130130222231232-3021200130323313-2200213103032030) |
| `detection_settings.violation_settings.disabled_violation_types` | [detection_settings.violation_settings.disabled_violation_types](data-sources--app_firewall--reference--group-001.md#canonical-0002010230321123-0310302130113321-2221213201211122-2122123100032211-3321333302221131-2203300112013201-3002333201123320-2001002023231131) |
| `detection_settings.violations_view` | [detection_settings.violations_view](data-sources--app_firewall--reference--group-001.md#canonical-1203332321221313-2111302023313301-1121111011313203-3031212123321323-3002213213332021-1311022221011221-3321133133132322-3103300003033221) |
| `detection_settings.violations_view.description_spec` | [detection_settings.violations_view.description_spec](data-sources--app_firewall--reference--group-001.md#canonical-3233030023230210-3331030103331111-2313011330122001-0130210220102303-2200020130201023-2031321231202131-1221020110200130-0022321231110301) |
| `detection_settings.violations_view.enabled` | [detection_settings.violations_view.enabled](data-sources--app_firewall--reference--group-001.md#canonical-2201202233333302-1021002011131322-2123010332211011-1130012311013213-3212213233312010-1021031130031213-0300133233300201-1311231222130201) |
| `detection_settings.violations_view.enabled_by_default` | [detection_settings.violations_view.enabled_by_default](data-sources--app_firewall--reference--group-001.md#canonical-3122231200123022-0330333021011023-1123022120200213-3213311013122112-0211220220232213-2202211030200301-3321311330132023-1230331332211202) |
| `detection_settings.violations_view.name` | [detection_settings.violations_view.name](data-sources--app_firewall--reference--group-001.md#canonical-3222313310102331-2301203323202211-1200033003221221-1000312311110003-2230111032102321-0323100301011003-2330120301302131-2221103212220003) |
| `detection_settings.violations_view.title` | [detection_settings.violations_view.title](data-sources--app_firewall--reference--group-001.md#canonical-3312131311013220-1220220111102232-1232223110031023-2330301012232202-0101011210100220-3000333130202330-0023133133222033-0223100123022203) |
| `disable_ai_enhancements` | [disable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-2300010332013322-0303123313112103-1220300113120332-0232130130131333-1320320202211332-2221222010100331-1322112110303201-3322203032332022) |
| `disable_anonymization` | [disable_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0013200302132212-3020321020200210-1121010002320121-0123030211233331-1021003213223133-2302013301233300-2331211233122322-3303212130002220) |
| `enable_ai_enhancements` | [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-2220132121302230-1300333113210112-0202002203103301-2200203022213223-3220012223020220-2100303013011130-0303222001011302-1302022130003213) |
| `enable_ai_enhancements.mitigate_high_medium_risk_action` | [enable_ai_enhancements.mitigate_high_medium_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-1233012332022201-0313112131211300-3021033321001100-1012222301311321-0103022012312130-1220132020331202-3303213032221230-3111232133330101) |
| `enable_ai_enhancements.mitigate_high_risk_action` | [enable_ai_enhancements.mitigate_high_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-3103332330113132-3232223333332022-1313222002302200-1211031103032112-0102302312111132-0011322100000102-1131322302122022-3002320301331032) |
| `id` | [ID](data-sources--app_firewall--reference--group-001.md#canonical-0022021301320033-3221220100001220-3301121103132122-2233000131210301-2302211301123210-2120323231122001-0301130003311012-1332033112131022) |
| `labels` | [labels](data-sources--app_firewall--reference--group-001.md#canonical-1031002310003010-0112310033020003-0023102133021022-3011012210102101-2222032231000220-0222221021302123-2001230331303321-2212301112020212) |
| `monitoring` | [monitoring](data-sources--app_firewall--reference--group-001.md#canonical-3300221033303003-1022023232101001-1333330003133123-1102332101201023-2021111201013202-2220203021331330-2323002302132000-3231201000003313) |
| `name` | [name](data-sources--app_firewall--reference--group-001.md#canonical-1023012033230133-2122330003301100-2020011232101110-1112123231032302-3203201011123003-2011103222233303-1113330022310321-2111001101103022) |
| `namespace` | [namespace](data-sources--app_firewall--reference--group-001.md#canonical-1033221303020330-2323310013203113-0221133323302123-0123031202012013-0221121111223321-1200013111213011-2303213232013013-1220111320110321) |
| `use_default_blocking_page` | [use_default_blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-0321123131001122-3332213331330211-1211010211223121-2331023011102220-1333113121111231-0212312003023000-3120130323030001-0130201232331003) |

<a id="canonical-2313001313310012-0301210321202031-0221010102220023-1012300213220002-2123102221223031-1002013330013010-1113000210233221-2111103222220013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all_response_codes` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- allow_all_response_codes

<a id="canonical-2030120232132200-0132203113121330-2320200323121013-2131002201103102-2021302233020212-2011033322213113-0020223131013022-3123103202001323"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all\_response\_codes, allowed\_response\_codes\] Configuration parameter for allow
all response codes. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [allow_all_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-2030120232132200-0132203113121330-2320200323121013-2131002201103102-2021302233020212-2011033322213113-0020223131013022-3123103202001323)
- [allowed_response_codes](data-sources--app_firewall--reference--group-001.md#canonical-2322002103222111-1012031013010111-3011232330032003-2003333132321322-1032123201021120-1321312313001000-1000200232131313-0212312011210101)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120023301001102-2103000021001323-3230113311223023-3220101220122000-2300110311220000-2121131310320100-2130312102313212-2200031122131102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_response_codes` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- allowed_response_codes

<a id="canonical-2322002103222111-1012031013010111-3011232330032003-2003333132321322-1032123201021120-1321312313001000-1000200232131313-0212312011210101"></a>

Type: `"single"`. Computed.

List of HTTP response status codes that are allowed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2202030103123030-1110123023312120-2011002132201132-3130021010220303-1002102131133100-1000001333323202-1022333331221100-0122013102210013"></a>

### Direct properties for `allowed_response_codes`

<a id="canonical-0320022303322203-2233030003301233-3202311333213102-2212021121102313-1221002213313320-2012002302230001-2000003233000110-1230033120033202"></a>

#### `allowed_response_codes.response_code` property

Type: `["list", "number"]`. Computed.

List of HTTP response status codes that are allowed.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 48,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.uint32.gte": "100",
    "ves.io.schema.rules.repeated.items.uint32.lte": "999",
    "ves.io.schema.rules.repeated.max_items": "48",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2330310310111223-0011110013221003-3130133233112122-0132133202311033-3310031301330023-2331222213303332-3101323002031013-1312223120332013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocking` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- blocking

<a id="canonical-0300002133113213-2203210210200113-2310222132003003-2312303203213212-2230321122133022-0331302012120031-2231320012103012-3222312013303302"></a>

Type: `["object", {}]`. Computed.

\[OneOf: blocking, monitoring\] Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [blocking](data-sources--app_firewall--reference--group-001.md#canonical-0300002133113213-2203210210200113-2310222132003003-2312303203213212-2230321122133022-0331302012120031-2231320012103012-3222312013303302)
- [monitoring](data-sources--app_firewall--reference--group-001.md#canonical-3300221033303003-1022023232101001-1333330003133123-1102332101201023-2021111201013202-2220203021331330-2323002302132000-3231201000003313)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320323133020103-0202323333132023-0231333131131330-3210000030111102-1100133100001232-3201112321100101-3321011100332101-3233030031031031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocking_page` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- blocking_page

<a id="canonical-0223133111001102-3313132310303000-3333312020123332-3231101000003330-2223021110110231-2113320202021112-3101020233300323-2122333003020021"></a>

Type: `"single"`. Computed.

\[OneOf: blocking\_page, use\_default\_blocking\_page; Default: use\_default\_blocking\_page\]
Custom Blocking Response Page. Custom blocking response page body.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-0223133111001102-3313132310303000-3333312020123332-3231101000003330-2223021110110231-2113320202021112-3101020233300323-2122333003020021)
- [use_default_blocking_page](data-sources--app_firewall--reference--group-001.md#canonical-0321123131001122-3332213331330211-1211010211223121-2331023011102220-1333113121111231-0212312003023000-3120130323030001-0130201232331003)

Select alternatives according to the provider validators above.

<a id="canonical-0101212113123131-0211212011112222-3303012032202020-3320003210132020-1231223131333323-2101213331303333-0121303201101023-2231103000203023"></a>

### Direct properties for `blocking_page`

<a id="canonical-0003323302321032-2330122303201122-1233103120022121-2030022222030113-1112113001030210-3000021311321310-3332202102200320-2332222312011032"></a>

#### `blocking_page.blocking_page` property

Type: `"string"`. Computed.

Define the content of the response page (e.g., an HTML document or a JSON object), use the
\{\{request\_id\}\} placeholder to provide users with a unique identifier to be able to trace the
blocked request in the logs. The maximum allowed size of response body is 4096 bytes after base64
encoding, which would be about 3070 bytes in plain text.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 4096,
    "metadata": {
      "category": "content",
      "confidence": 0.99,
      "note": "Must be a valid URI (uri_ref). API rejects inline HTML despite the description example.",
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0331031200210323-2310311003323212-0322020231210020-2312313003203132-0210303122033010-0120030120023203-2121322302232111-0332100000121022"></a>

<a id="canonical-2020231010000022-1333202131101201-3320233333211323-2303101312023123-1200211133202110-1002212221110020-3213132332013302-1221100131202012"></a>

#### `blocking_page.response_code` property

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Additional upstream details:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103022221032002-3031232012310300-2201020202303231-0102202313031033-2021133200010103-0330231303130020-0330211202222110-3021203231003220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_protection_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- bot_protection_setting

<a id="canonical-3002231223212103-0011012011120012-0322120001221313-0211103031210213-2310300220230302-0010300230222001-3223230220110011-3010330320112233"></a>

Type: `"single"`. Computed.

\[OneOf: bot\_protection\_setting, default\_bot\_setting; Default: default\_bot\_setting\]
Configuration parameter for bot protection setting.

Additional upstream details:

Configuration of WAF Bot Protection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3002231223212103-0011012011120012-0322120001221313-0211103031210213-2310300220230302-0010300230222001-3223230220110011-3010330320112233)
- [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-1032213023131131-2233121312222130-1120233323213131-2013222202113103-3332102200322100-1332031020222110-2132130300031231-1112121020002130)

Select alternatives according to the provider validators above.

<a id="canonical-0231122231100213-0211310131131020-1110230332110302-1131200220132223-2112112112301300-1320220322200332-0223311223212113-0201222301011122"></a>

### Direct properties for `bot_protection_setting`

<a id="canonical-1301120112333110-0200213211300221-2120232220012010-0120032200001102-2111212323221220-0221333132322200-1301301123221130-3131030120201213"></a>

#### `bot_protection_setting.good_bot_action` property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1010220011202202-1330211212021102-2112002313332011-3220002202000223-0302212331303331-0310331023323033-2130221212103313-1222203121112113"></a>

<a id="canonical-2033220011100312-2211231212303313-3111103102020331-2300200210310231-1023020323311023-1122201322111111-2313322310313201-3331003321002312"></a>

#### `bot_protection_setting.malicious_bot_action` property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020102231332320-0101200112013003-0211011032112232-1120122131220213-3031322113103230-3213201130131202-0203110202021002-0121113020100010"></a>

<a id="canonical-2232312332001000-0033202121031333-0211121230112002-1332103102011213-3211321331231221-2110010102001321-0232100202002122-0203201022313011"></a>

#### `bot_protection_setting.suspicious_bot_action` property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0023020222013102-0220221233101131-3310023020113202-2203132000300231-3310200111300122-1222233233313303-0131012223310232-1211132213211011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- custom_anonymization

<a id="canonical-2012101332111332-3233220203102320-0220330113310320-2321100332212313-3002200311130022-0210001120121312-2323023230312131-1311232202332231"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_anonymization, default\_anonymization, disable\_anonymization; Default:
default\_anonymization\] Anonymization settings which is a list of HTTP headers, parameters and
cookies.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-2012101332111332-3233220203102320-0220330113310320-2321100332212313-3002200311130022-0210001120121312-2323023230312131-1311232202332231)
- [default_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-3033312130321301-2011101021301322-1202031001023100-3100200110303123-2120322321020030-0210313213001200-2303120101133223-2323233222030223)
- [disable_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0013200302132212-3020321020200210-1121010002320121-0123030211233331-1021003213223133-2302013301233300-2331211233122322-3303212130002220)

Select alternatives according to the provider validators above.

<a id="canonical-1001211330122232-0221132112210321-3322011031230321-1233010233102302-3030132130112231-3231333233031331-2013233211331233-1101100002300001"></a>

### Direct properties for `custom_anonymization`

- [anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-3313103023110012-3112123312210312-2333322032100303-0002031333003321-0330010220033330-2232312011032230-0002312223300213-3010110200220233): complete subsection reference.

<a id="canonical-3313103023110012-3112123312210312-2333322032100303-0002031333003321-0330010220033330-2232312011032230-0002312223300213-3010110200220233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0023020222013102-0220221233101131-3310023020113202-2203132000300231-3310200111300122-1222233233313303-0131012223310232-1211132213211011)
- custom_anonymization.anonymization_config

<a id="canonical-0020201313133032-3213003220030023-3331022032320132-2230323203331333-1120333101323222-1112333221021322-0001231020032112-1200130131113031"></a>

Type: `"list"`. Computed.

List of HTTP headers, cookies and query parameters whose values will be masked.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2021331330312001-1021303320302223-0113322133211311-3013203001032011-2130230202232123-0323132203301221-2023010011201212-0332133202223012"></a>

### Direct properties for `custom_anonymization.anonymization_config`

- [cookie](data-sources--app_firewall--reference--group-001.md#canonical-0002031121233033-3331213132001322-0131001022112023-2131333200211310-3110013130001123-3123110231001221-3113221002232311-2202223232011211): complete subsection reference.

- [http_header](data-sources--app_firewall--reference--group-001.md#canonical-0100321222011122-2331012001321130-2201200113002333-1232110223010102-3232230233001000-3100023323011100-2232001213202312-1030032031132020): complete subsection reference.

- [query_parameter](data-sources--app_firewall--reference--group-001.md#canonical-3223321311230011-3302030021021133-1210012223323311-2112000322330230-1331223300031322-1120130002012321-3203200323123110-3222120002113302): complete subsection reference.

<a id="canonical-0002031121233033-3331213132001322-0131001022112023-2131333200211310-3110013130001123-3123110231001221-3113221002232311-2202223232011211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config.cookie` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0023020222013102-0220221233101131-3310023020113202-2203132000300231-3310200111300122-1222233233313303-0131012223310232-1211132213211011)
- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-3313103023110012-3112123312210312-2333322032100303-0002031333003321-0330010220033330-2232312011032230-0002312223300213-3010110200220233)
- custom_anonymization.anonymization_config.cookie

<a id="canonical-0300322321001231-3310130300220101-1323120212111210-0030212231321102-0010003112121303-0130232123102221-1112130210132321-2031133202032121"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Cookies.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2122331112201311-3232211201320232-2020002132332320-2230112221000122-0323112230233221-2203201121303312-3323013211132230-1112032310230001"></a>

### Direct properties for `custom_anonymization.anonymization_config.cookie`

<a id="canonical-1230312123301131-0220102033323302-2031222300121031-3223003031211231-1231220301301021-1213030012133132-0111032020220311-3113012130010121"></a>

#### `custom_anonymization.anonymization_config.cookie.cookie_name` property

Type: `"string"`. Computed.

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0100321222011122-2331012001321130-2201200113002333-1232110223010102-3232230233001000-3100023323011100-2232001213202312-1030032031132020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config.http_header` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0023020222013102-0220221233101131-3310023020113202-2203132000300231-3310200111300122-1222233233313303-0131012223310232-1211132213211011)
- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-3313103023110012-3112123312210312-2333322032100303-0002031333003321-0330010220033330-2232312011032230-0002312223300213-3010110200220233)
- custom_anonymization.anonymization_config.http_header

<a id="canonical-0100213322023322-0322311112033333-3213120120000333-0122003032231100-3021033320301330-2120122112333200-1210333323231011-2022323200011120"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120232221131323-2210122130310102-0122103321131221-1332000101203011-2302133133032201-1331002113110313-0302002113023033-3131031103133323"></a>

### Direct properties for `custom_anonymization.anonymization_config.http_header`

<a id="canonical-0112330032302213-2030101203002303-2231010302122232-0320011020122111-2200013233320330-3223220021122031-1020233331100012-2121202112322112"></a>

#### `custom_anonymization.anonymization_config.http_header.header_name` property

Type: `"string"`. Computed.

Masks the HTTP header value. The setting does not mask the HTTP header name. Wildcard matching can
be used by prefixing or suffixing the HTTP header name with a wildcard asterisk (\*), or by using
only an asterisk to match any HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true"
  }
}
```

<a id="canonical-3223321311230011-3302030021021133-1210012223323311-2112000322330230-1331223300031322-1120130002012321-3203200323123110-3222120002113302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config.query_parameter` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [custom_anonymization](data-sources--app_firewall--reference--group-001.md#canonical-0023020222013102-0220221233101131-3310023020113202-2203132000300231-3310200111300122-1222233233313303-0131012223310232-1211132213211011)
- [custom_anonymization.anonymization_config](data-sources--app_firewall--reference--group-001.md#canonical-3313103023110012-3112123312210312-2333322032100303-0002031333003321-0330010220033330-2232312011032230-0002312223300213-3010110200220233)
- custom_anonymization.anonymization_config.query_parameter

<a id="canonical-3020202223103303-2202202113313132-1211202100212312-1321023313133233-3031030230233123-2033231010021231-2030022120012221-3301202223012113"></a>

Type: `"single"`. Computed.

Configure anonymization for HTTP Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0303311313133220-0200300201223221-2210010123002101-3311002103333233-0113322132031023-0211103030322231-2210002111022210-1112233131211102"></a>

### Direct properties for `custom_anonymization.anonymization_config.query_parameter`

<a id="canonical-1322223200012200-1203002202223311-0233131000300203-3113032030120112-3021201302310221-0113012122031220-1122013222300131-0101111013001211"></a>

#### `custom_anonymization.anonymization_config.query_parameter.query_param_name` property

Type: `"string"`. Computed.

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1230333210011113-2001223230211300-0100320022133302-3121233203102311-2232332031011332-3022022012030332-0013113012000333-0101213203220220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_anonymization` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- default_anonymization

<a id="canonical-3033312130321301-2011101021301322-1202031001023100-3100200110303123-2120322321020030-0210313213001200-2303120101133223-2323233222030223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default anonymization. Defaults to \`map\[\]\`. Server applies default
when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202031113322310-1322320330322332-3311330222223232-3111201232203010-1010320232321303-3212123031302222-3130302330003303-2102013312033211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_bot_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- default_bot_setting

<a id="canonical-1032213023131131-2233121312222130-1120233323213131-2013222202113103-3332102200322100-1332031020222110-2132130300031231-1112121020002130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default bot setting. Defaults to \`map\[\]\`. Server applies default
when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130323331212003-3202232330332222-1311210212001102-3320222310010120-0021301132303300-1031210102102300-0203213003111132-1321211232221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_detection_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- default_detection_settings

<a id="canonical-0000113131131211-1032200030131323-2130321012123211-2012022313302202-2313202232332103-0120303201132300-3012221313303200-3303212221130301"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_detection\_settings, detection\_settings; Default: default\_detection\_settings\]
Configuration parameter for default detection settings. Defaults to \`map\[\]\`. Server applies
default when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [default_detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-0000113131131211-1032200030131323-2130321012123211-2012022313302202-2313202232332103-0120303201132300-3012221313303200-3303212221130301)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-0021023012003132-1331020220200333-2223201122312223-2222023303313320-2230221233233330-2031012111301103-2131103032231203-2023033212112012)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- detection_settings

<a id="canonical-0021023012003132-1331020220200333-2223201122312223-2222023303313320-2230221233233330-2031012111301103-2131103032231203-2023033212112012"></a>

Type: `"single"`. Computed.

Specifies detection settings to be used by WAF.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-bot_protection_choice": "[\"bot_protection_setting\",\"default_bot_setting\"]",
  "x-ves-oneof-field-false_positive_suppression": "[\"disable_suppression\",\"enable_suppression\"]",
  "x-ves-oneof-field-signatures_staging_settings": "[\"disable_staging\",\"stage_new_and_updated_signatures\",\"stage_new_signatures\"]",
  "x-ves-oneof-field-threat_campaign_choice": "[\"disable_threat_campaigns\",\"enable_threat_campaigns\"]",
  "x-ves-oneof-field-violation_detection_setting": "[\"default_violation_settings\",\"violation_settings\"]"
}
```

<a id="canonical-2223001121313203-0003232020123101-2232312320122311-0302222213222302-3010112122020001-1123120012333321-3100211111331020-0231122020002303"></a>

### Direct properties for `detection_settings`

- [bot_protection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3120233020110133-0002320311033313-1300011030333303-2323113212300203-1223031312021000-3112223031220011-0122021023321331-0131222123323023): complete subsection reference.

- [default_bot_setting](data-sources--app_firewall--reference--group-001.md#canonical-0132333013020130-1102021300122310-3321131010200120-3303223000000300-0001103003213031-1113221110220231-1323300133111203-1312002032223021): complete subsection reference.

- [default_violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-0021222013330233-0210302130131223-3032110201201200-3012302111321300-3000333330212111-0300102113110121-0131210001120131-3001111321113302): complete subsection reference.

- [disable_staging](data-sources--app_firewall--reference--group-001.md#canonical-3213122303123011-2113012332332233-3203221222322202-3301001112220100-2333033310013222-3033022010102020-2210211213001133-0210103232222120): complete subsection reference.

- [disable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-0212311203201312-1022102020331320-1301222221010331-2312011013223331-2112200133201102-3212223322202211-0031121332112022-2010010203201200): complete subsection reference.

- [disable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-2020320002100003-1213122201330131-3133312031210112-2302103301000130-1320010202001222-1112322113230233-0010133030111322-3331130223101332): complete subsection reference.

- [enable_suppression](data-sources--app_firewall--reference--group-001.md#canonical-0123200210002121-3201100320232232-1223021233303011-0111221131211232-0010222100201220-3212233210223331-2330210231200123-1332320313302221): complete subsection reference.

- [enable_threat_campaigns](data-sources--app_firewall--reference--group-001.md#canonical-1301302000221331-1030030030233212-3100211113233202-3312300213032130-1013003122103132-1103210213231012-1012213212210132-0211212232333110): complete subsection reference.

- [signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230): complete subsection reference.

- [stage_new_and_updated_signatures](data-sources--app_firewall--reference--group-001.md#canonical-3232200212223300-3320023110332300-2331320211020031-2030102130000220-0223112011310321-1201312331132011-3211130013001233-0110330201332011): complete subsection reference.

- [stage_new_signatures](data-sources--app_firewall--reference--group-001.md#canonical-0101303121302322-2223101111022103-0223111322103201-2331300003032212-0010323021122300-0030013133200230-2223201020120133-0121121332212103): complete subsection reference.

- [violation_settings](data-sources--app_firewall--reference--group-001.md#canonical-1121231313133202-1000332220232313-2120221311012113-0021001032211022-0120102000133031-2233210131133020-2201101111231221-3332020012132031): complete subsection reference.

- [violations_view](data-sources--app_firewall--reference--group-001.md#canonical-3100112002213131-1121100000222220-1222121201303032-3033012303001222-2102330311310130-1021103301002112-0221013203310100-2322122000201311): complete subsection reference.

<a id="canonical-3120233020110133-0002320311033313-1300011030333303-2323113212300203-1223031312021000-3112223031220011-0122021023321331-0131222123323023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.bot_protection_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.bot_protection_setting

<a id="canonical-3110220100211021-0213120200001131-2002022002123201-1213133220012113-1021022312022322-1003211221213322-2131222323110233-2210202210113101"></a>

Type: `"single"`. Computed.

Configuration parameter for bot protection setting.

Additional upstream details:

Configuration of WAF Bot Protection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0013220033110033-0000023032002202-3332320303221211-2103331221302332-0202011100103333-3131031010032201-3300211300210331-3131211123332102"></a>

### Direct properties for `detection_settings.bot_protection_setting`

<a id="canonical-2121023332010203-2021132001311312-2122212023322303-0021133213101121-1022123012033031-0220332001123132-2332011003300010-2003313023031030"></a>

#### `detection_settings.bot_protection_setting.good_bot_action` property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2011123021113231-3023333000233311-1002200133100313-2000321001122100-3313321023131231-2212221332033313-3300113310012303-2322133103132130"></a>

<a id="canonical-3033012203313021-0030223122111103-3210311233303012-0011130201123230-2131121313003120-3213203023312010-1230003122102110-0320202300201301"></a>

#### `detection_settings.bot_protection_setting.malicious_bot_action` property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3300113322102011-1133111233120312-0003303033122103-3020320013123203-2331213303320131-3303122313021132-1331010123110220-0222333103302020"></a>

<a id="canonical-3010313300302011-0302331111231012-2202311222111100-3300303223312312-0300033210121013-1302011333330103-1022110013330131-0101312203232111"></a>

#### `detection_settings.bot_protection_setting.suspicious_bot_action` property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0132333013020130-1102021300122310-3321131010200120-3303223000000300-0001103003213031-1113221110220231-1323300133111203-1312002032223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.default_bot_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.default_bot_setting

<a id="canonical-1201221103023030-1131011011121220-2020300032011202-3303003013301200-1120023311223032-1012200302331121-1320021211210310-0313333300123321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default bot setting.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021222013330233-0210302130131223-3032110201201200-3012302111321300-3000333330212111-0300102113110121-0131210001120131-3001111321113302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.default_violation_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.default_violation_settings

<a id="canonical-2212222301011001-2022120321331102-3312101023200230-2002200301113032-0313101120311313-1121210312230121-2100103011210032-1111113303333201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default violation settings.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213122303123011-2113012332332233-3203221222322202-3301001112220100-2333033310013222-3033022010102020-2210211213001133-0210103232222120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.disable_staging` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.disable_staging

<a id="canonical-0110321030123220-3012102103010223-3201321232330110-3033333111112030-0200210301000213-0332001211103012-3210011310102030-2133001310322010"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212311203201312-1022102020331320-1301222221010331-2312011013223331-2112200133201102-3212223322202211-0031121332112022-2010010203201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.disable_suppression` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.disable_suppression

<a id="canonical-2133001213000231-3302213013002323-1301013123230012-2210322110230021-0101233102121002-1303020333231010-2230120033110303-2323133310012120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable suppression.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020320002100003-1213122201330131-3133312031210112-2302103301000130-1320010202001222-1112322113230233-0010133030111322-3331130223101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.disable_threat_campaigns` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.disable_threat_campaigns

<a id="canonical-2302030221232311-3100123021110102-2130211020302220-2022210132300023-1022300313111100-1330022133020230-3231122003001010-1121221210100012"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123200210002121-3201100320232232-1223021233303011-0111221131211232-0010222100201220-3212233210223331-2330210231200123-1332320313302221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.enable_suppression` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.enable_suppression

<a id="canonical-3313033320130201-3223012101323001-0221132013301102-2121222132332013-0102102300330102-3111203031120131-3331122301001103-1313212031213230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable suppression.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301302000221331-1030030030233212-3100211113233202-3312300213032130-1013003122103132-1103210213231012-1012213212210132-0211212232333110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.enable_threat_campaigns` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.enable_threat_campaigns

<a id="canonical-3010313021102222-3113203303203011-1022131130123223-3100331321023210-1001322000312203-2021111123302312-1323303013023200-0212323122311222"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.signature_selection_setting

<a id="canonical-1110121301000301-2022101211330021-3230310023223111-1303230332202002-3313203132231023-0032010102101310-3110222331123000-2202323213002211"></a>

Type: `"single"`. Computed.

Attack Signatures are patterns that identify attacks on a web application and its components.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-attack_type_setting": "[\"attack_type_settings\",\"default_attack_type_settings\"]",
  "x-ves-oneof-field-signature_protection_choice": "[\"default_signature_setting\",\"signature_settings_by_accuracy\"]",
  "x-ves-oneof-field-signature_selection_by_accuracy": "[\"high_medium_accuracy_signatures\",\"high_medium_low_accuracy_signatures\",\"only_high_accuracy_signatures\"]"
}
```

<a id="canonical-1131231301202011-0301010200132310-2322313232012210-2312010101233003-2122123111313012-1330123230203222-3200100121131310-2102122020222321"></a>

### Direct properties for `detection_settings.signature_selection_setting`

- [attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-3221203122131311-3100111230001311-1210213123303233-0310222213330333-1311232002332113-0223211313203110-3113303022120031-3303310001132023): complete subsection reference.

- [default_attack_type_settings](data-sources--app_firewall--reference--group-001.md#canonical-0330120222322222-0311213211003033-3013131113302220-1303132030100230-1210122110020233-1013331323102122-2211200211111100-3233310323103122): complete subsection reference.

- [default_signature_setting](data-sources--app_firewall--reference--group-001.md#canonical-0230113311030022-1101331313222002-1013302130232132-3032103232012100-1213100102002001-2111200121002020-3201301133203202-1103130032102301): complete subsection reference.

- [high_medium_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-0201013103122001-3301002020033222-3011223121032320-1202202222212310-0033303120023320-0000023232321002-1300323021303320-1211022100103102): complete subsection reference.

- [high_medium_low_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-3330321130231200-0323132213300203-3221203311133032-1122202210312033-2320202022333232-3303102302000011-3221322012111000-1212320230330233): complete subsection reference.

- [only_high_accuracy_signatures](data-sources--app_firewall--reference--group-001.md#canonical-0200303013011003-1012011220100202-2002301213013230-2302022330301100-1333212110223030-0221112031203303-1130233331121112-0102113230023022): complete subsection reference.

- [signature_settings_by_accuracy](data-sources--app_firewall--reference--group-001.md#canonical-0223310013322210-0201120110122302-3210110002210201-3320130230112032-3011301330100222-2110113001223200-0133333012023221-2021202010331032): complete subsection reference.

<a id="canonical-3221203122131311-3100111230001311-1210213123303233-0310222213330333-1311232002332113-0223211313203110-3113303022120031-3303310001132023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.attack_type_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230)
- detection_settings.signature_selection_setting.attack_type_settings

<a id="canonical-3330312130103121-0113011132223030-1011111301320113-3332230213132031-3300311132230213-0200301120111112-1233011223301232-2212313233123102"></a>

Type: `"single"`. Computed.

Specifies attack-type settings to be used by WAF.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2223000030212013-1113100020232111-3101130322213313-3023333323302003-1223330303122300-1231102121100310-0210010201011122-0101020123101023"></a>

### Direct properties for `detection_settings.signature_selection_setting.attack_type_settings`

<a id="canonical-2021031200301220-3300112332330330-3300310331230203-3201033113012132-1222010230001010-3033232323113102-1223011202132310-1303030021210213"></a>

#### `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` property

Type: `["list", "string"]`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of Attack Types that will be ignored and not trigger a detection. Possible values are
\`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 22,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "22",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "22",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0330120222322222-0311213211003033-3013131113302220-1303132030100230-1210122110020233-1013331323102122-2211200211111100-3233310323103122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.default_attack_type_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230)
- detection_settings.signature_selection_setting.default_attack_type_settings

<a id="canonical-2213103100022023-2223220212033013-1221222100133013-1013013212131023-2230022221202022-3021332223222213-3022103110302331-0201303001112131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default attack type settings.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230113311030022-1101331313222002-1013302130232132-3032103232012100-1213100102002001-2111200121002020-3201301133203202-1103130032102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.default_signature_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230)
- detection_settings.signature_selection_setting.default_signature_setting

<a id="canonical-3301122032303132-1222212031010000-3210213110010030-3322102232032330-0112220113220313-0213113032302013-1223233003002130-1030020303013223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default signature setting.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201013103122001-3301002020033222-3011223121032320-1202202222212310-0033303120023320-0000023232321002-1300323021303320-1211022100103102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230)
- detection_settings.signature_selection_setting.high_medium_accuracy_signatures

<a id="canonical-2332133310112020-2131113200012221-3323020331232220-2221031300010311-0221332001211131-2001203231030120-1231232200233110-0320322121011321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for high medium accuracy signatures.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330321130231200-0323132213300203-3221203311133032-1122202210312033-2320202022333232-3303102302000011-3221322012111000-1212320230330233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230)
- detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures

<a id="canonical-3320122120012101-3320323331132202-1033220000113023-3023000122013223-2033031003000301-3023331211311002-1210303202102232-0112021302300331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for high medium low accuracy signatures.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200303013011003-1012011220100202-2002301213013230-2302022330301100-1333212110223030-0221112031203303-1130233331121112-0102113230023022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.only_high_accuracy_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230)
- detection_settings.signature_selection_setting.only_high_accuracy_signatures

<a id="canonical-0211221033311230-1003012022230001-0110000123203300-0213200012101033-2130031302022332-3230210011331302-3220111121302310-2230323323113320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for only high accuracy signatures.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223310013322210-0201120110122302-3210110002210201-3320130230112032-3011301330100222-2110113001223200-0133333012023221-2021202010331032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.signature_settings_by_accuracy` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- [detection_settings.signature_selection_setting](data-sources--app_firewall--reference--group-001.md#canonical-3211300213032331-3031011223300102-3100232032323213-3022013011333101-2332320031111110-3301322300002010-1313233031021020-2013300130023230)
- detection_settings.signature_selection_setting.signature_settings_by_accuracy

<a id="canonical-3000233222232323-2020200211122111-2103312200200320-1130330212132210-3001322310213001-2001301202022200-1330201323312003-1002222330202120"></a>

Type: `"single"`. Computed.

Configuration of WAF Signature Protection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0213230302302000-2333023303222221-3011132033333330-1213202320312302-2012112003223130-2301311002233130-0021110130131120-3101232332331022"></a>

### Direct properties for `detection_settings.signature_selection_setting.signature_settings_by_accuracy`

<a id="canonical-1111233023211301-0210102322322012-1032331322020131-3200300022133331-2103003302102032-3010220202211230-2013023130302213-1232033311130013"></a>

#### `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` property

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2002302102032300-0332333032103201-1133122003312122-1330123222131323-3201320102323320-0130332001030023-1031032313121021-1201112101302030"></a>

<a id="canonical-1011002230302230-2200333113021213-3110320030132022-0302203213002013-0320033013220122-0203013103201202-2113100001233330-1311203310120312"></a>

#### `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` property

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1111221210233300-2311102131331013-1223320133300132-3331323231130223-2202121112231201-0203223133131031-1332012013233310-1310033223233131"></a>

<a id="canonical-3021332212103312-2030322002122130-1113030232201223-2301221102321220-2010032003101102-0100323203301302-3233230311231213-3233011133203320"></a>

#### `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` property

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3232200212223300-3320023110332300-2331320211020031-2030102130000220-0223112011310321-1201312331132011-3211130013001233-0110330201332011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.stage_new_and_updated_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.stage_new_and_updated_signatures

<a id="canonical-2201301213222212-2231213131333210-1312323123123201-2133300021013131-3010011101321210-3333022002010101-3222310121330331-1310233122301003"></a>

Type: `"single"`. Computed.

Attack Signatures staging configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1121233213201122-1001120230203333-2330312213010300-1001002030130102-3323022322302322-3000100013032103-2011221221231322-3323011101022333"></a>

### Direct properties for `detection_settings.stage_new_and_updated_signatures`

<a id="canonical-2331322200201221-0022322221232311-2003011021303213-3222133213233232-1332100202321003-0112020132133202-1033113121131330-0311232320321321"></a>

#### `detection_settings.stage_new_and_updated_signatures.staging_period` property

Type: `"number"`. Computed.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Staging period in days. Default 7, max 20. Applies to both stage_new_and_updated_signatures and stage_new_signatures.",
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-0101303121302322-2223101111022103-0223111322103201-2331300003032212-0010323021122300-0030013133200230-2223201020120133-0121121332212103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.stage_new_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.stage_new_signatures

<a id="canonical-0230303201022123-1222030230130302-3032303320023001-0022331112333023-1230131032210202-0222112010103320-3302232323231120-3323213200203221"></a>

Type: `"single"`. Computed.

Attack Signatures staging configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3123200121332013-3013103123012000-2000012300231223-0103323313010223-0313301330311200-1212300221013333-2331132120122130-3101011201113220"></a>

### Direct properties for `detection_settings.stage_new_signatures`

<a id="canonical-1231203222031221-2030031232201100-3100121330213021-3130131222130121-1312021222013322-3333100002222223-0011030222323110-0123223302330003"></a>

#### `detection_settings.stage_new_signatures.staging_period` property

Type: `"number"`. Computed.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Staging period in days. Default 7, max 20. Applies to both stage_new_and_updated_signatures and stage_new_signatures.",
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-1121231313133202-1000332220232313-2120221311012113-0021001032211022-0120102000133031-2233210131133020-2201101111231221-3332020012132031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.violation_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.violation_settings

<a id="canonical-1213010100321222-3300210220300200-3101021102222212-0331303202020330-1123221012210020-1130130222231232-3021200130323313-2200213103032030"></a>

Type: `"single"`. Computed.

Specifies violation settings to be used by WAF.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1021230133233230-3002201000322330-0022310320122013-3133201201322130-3102221232322011-3033211322122000-1222023131013202-0033102111112013"></a>

### Direct properties for `detection_settings.violation_settings`

<a id="canonical-0002010230321123-0310302130113321-2221213201211122-2122123100032211-3321333302221131-2203300112013201-3002333201123320-2001002023231131"></a>

#### `detection_settings.violation_settings.disabled_violation_types` property

Type: `["list", "string"]`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
Disabled Violations. List of violations to be excluded. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "40",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "40",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3100112002213131-1121100000222220-1222121201303032-3033012303001222-2102330311310130-1021103301002112-0221013203310100-2322122000201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.violations_view` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [detection_settings](data-sources--app_firewall--reference--group-001.md#canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210)
- detection_settings.violations_view

<a id="canonical-1203332321221313-2111302023313301-1121111011313203-3031212123321323-3002213213332021-1311022221011221-3321133133132322-3103300003033221"></a>

Type: `"list"`. Computed.

List of violation checks that are performed on HTTP request to ensure the requests are properly
formatted, detection of evasion techniques and other violations.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2222033223010223-0203031111211323-2310032100130000-3212203222213031-2200013120130331-0203020022233232-0211300331111101-1110020133121000"></a>

### Direct properties for `detection_settings.violations_view`

<a id="canonical-3233030023230210-3331030103331111-2313011330122001-0130210220102303-2200020130201023-2031321231202131-1221020110200130-0022321231110301"></a>

#### `detection_settings.violations_view.description_spec` property

Type: `"string"`. Computed.

Description. Human-readable description text

<a id="canonical-2201202233333302-1021002011131322-2123010332211011-1130012311013213-3212213233312010-1021031130031213-0300133233300201-1311231222130201"></a>

<a id="canonical-0103210203301303-2030112310232013-3000131022223312-0132133201021011-2211122330232110-0100000122010231-3233332321200110-1213323013300313"></a>

#### `detection_settings.violations_view.enabled` property

Type: `"bool"`. Computed.

State. Enable or disable the feature

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122231200123022-0330333021011023-1123022120200213-3213311013122112-0211220220232213-2202211030200301-3321311330132023-1230331332211202"></a>

<a id="canonical-2122333313303121-0211133102130202-3001312110112230-0212223330012013-0010030133120231-0213211010211233-1200312331320313-2220232222001313"></a>

#### `detection_settings.violations_view.enabled_by_default` property

Type: `"string"`. Computed.

Violations that are enabled by default by F5 are advisable to leave enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3222313310102331-2301203323202211-1200033003221221-1000312311110003-2230111032102321-0323100301011003-2330120301302131-2221103212220003"></a>

<a id="canonical-3230220332211201-1030113313111112-3332323311201120-1130230103130013-3310222301210012-1002010221112000-3302031131032322-0320021112230123"></a>

#### `detection_settings.violations_view.name` property

Type: `"string"`. Computed.

Name. Human-readable name for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3312131311013220-1220220111102232-1232223110031023-2330301012232202-0101011210100220-3000333130202330-0023133133222033-0223100123022203"></a>

<a id="canonical-3211100203003210-2103013333321021-3111032201122031-1322032033130333-1113133211010333-3310112221102030-3222220212030302-3210223210203212"></a>

#### `detection_settings.violations_view.title` property

Type: `"string"`. Computed.

Title. Human-readable title for the resource

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1010320003113321-0320300100231200-3321123033130211-0210221201100101-2302312001221203-3112323102130130-0302223322012101-2113111130031003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ai_enhancements` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- disable_ai_enhancements

<a id="canonical-2300010332013322-0303123313112103-1220300113120332-0232130130131333-1320320202211332-2221222010100331-1322112110303201-3322203032332022"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_ai\_enhancements, enable\_ai\_enhancements; Default: disable\_ai\_enhancements\]
Configuration parameter for disable ai enhancements. Defaults to \`map\[\]\`. Server applies default
when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

OneOf alternatives in this subsection:

- [disable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-2300010332013322-0303123313112103-1220300113120332-0232130130131333-1320320202211332-2221222010100331-1322112110303201-3322203032332022)
- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-2220132121302230-1300333113210112-0202002203103301-2200203022213223-3220012223020220-2100303013011130-0303222001011302-1302022130003213)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303321323122022-3010001011230002-0110201311133021-3020110232333222-3002332123012032-1313132030122010-3101221013302031-1232201001022322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_anonymization` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- disable_anonymization

<a id="canonical-0013200302132212-3020321020200210-1121010002320121-0123030211233331-1021003213223133-2302013301233300-2331211233122322-3303212130002220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable anonymization.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330323012211103-0031131130300122-2011021101121023-3112012123021123-3231300122313331-1101212033212032-3302322233212232-3200321220223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ai_enhancements` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- enable_ai_enhancements

<a id="canonical-2220132121302230-1300333113210112-0202002203103301-2200203022213223-3220012223020220-2100303013011130-0303222001011302-1302022130003213"></a>

Type: `"single"`. Computed.

Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-risk_score_action_choice": "[\"mitigate_high_medium_risk_action\",\"mitigate_high_risk_action\"]"
}
```

<a id="canonical-0203310123033002-3331103322221200-0300220003101103-1232113330110302-2001303100233132-2023331301201130-1031103211122013-3011000111210100"></a>

### Direct properties for `enable_ai_enhancements`

- [mitigate_high_medium_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-2303110021023333-3102213213101221-3110302022133211-0033321120331222-3121003120122011-1132131201212123-2301201201113133-1112020331201123): complete subsection reference.

- [mitigate_high_risk_action](data-sources--app_firewall--reference--group-001.md#canonical-1131031322200111-2303332232031231-1002321033322021-3122330102212310-2201203300102010-0130130210201133-2320321333023311-2233110102320233): complete subsection reference.

<a id="canonical-2303110021023333-3102213213101221-3110302022133211-0033321120331222-3121003120122011-1132131201212123-2301201201113133-1112020331201123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ai_enhancements.mitigate_high_medium_risk_action` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-3330323012211103-0031131130300122-2011021101121023-3112012123021123-3231300122313331-1101212033212032-3302322233212232-3200321220223211)
- enable_ai_enhancements.mitigate_high_medium_risk_action

<a id="canonical-1233012332022201-0313112131211300-3021033321001100-1012222301311321-0103022012312130-1220132020331202-3303213032221230-3111232133330101"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131031322200111-2303332232031231-1002321033322021-3122330102212310-2201203300102010-0130130210201133-2320321333023311-2233110102320233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ai_enhancements.mitigate_high_risk_action` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- [enable_ai_enhancements](data-sources--app_firewall--reference--group-001.md#canonical-3330323012211103-0031131130300122-2011021101121023-3112012123021123-3231300122313331-1101212033212032-3302322233212232-3200321220223211)
- enable_ai_enhancements.mitigate_high_risk_action

<a id="canonical-3103332330113132-3232223333332022-1313222002302200-1211031103032112-0102302312111132-0011322100000102-1131322302122022-3002320301331032"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132332001110103-0010201300032310-0100330312110132-3301223331032312-2131202331131103-3320010120230110-0301101202322330-0001011321201113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `monitoring` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- monitoring

<a id="canonical-3300221033303003-1022023232101001-1333330003133123-1102332101201023-2021111201013202-2220203021331330-2323002302132000-3231201000003313"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133021012333021-2112221112021001-1232130013231113-1033201031303202-2010021000212113-3231101322123331-0123233300232303-3310130211333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_blocking_page` properties

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-0320023201111121-3000130210130132-3330020030101023-0011302221323021-1102320231200030-1023333300332200-0122301003313000-0301011232103312)
- [Property reference](data-sources--app_firewall--reference--group-001.md#canonical-3230302312020213-0303301210111010-3121313101331331-1312303103033320-1333333032122033-3200221010203202-3031313132223223-3102110120032111)
- use_default_blocking_page

<a id="canonical-0321123131001122-3332213331330211-1211010211223121-2331023011102220-1333113121111231-0212312003023000-3120130323030001-0130201232331003"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.
