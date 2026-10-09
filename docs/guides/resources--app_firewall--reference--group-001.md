---
page_title: "xcsh_app_firewall reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall reference."
---

# xcsh_app_firewall reference

<a id="canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- Property reference

<a id="canonical-2311103010020031-2021333120132100-0033223310031111-1221000112210110-1120333320222013-1132330330312223-2133123321221020-2223210001223123"></a>

### Direct properties for `xcsh_app_firewall`

- [allow_all_response_codes](resources--app_firewall--reference--group-001.md#canonical-2121321311000013-1301132310203200-0311030332212320-2120130000301202-1020120212020110-1010311311021120-3203102121231022-2020122312221103): complete subsection reference.

- [allowed_response_codes](resources--app_firewall--reference--group-001.md#canonical-0120311133313213-1113131100301122-3201330201312210-0023331230030310-2032013032030031-0122113221102121-2233331112123232-0130001130213312): complete subsection reference.

<a id="canonical-2313220112100333-1003333033210100-3112122021331001-3122101001133320-1122331203012121-1000003132321021-3033130120301330-1023120033101111"></a>

<a id="canonical-0233321100223212-3213223111013023-0231023301123021-0000111003122223-2110102101133221-0003103103033220-0013120220210020-0022123210311212"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

- [blocking](resources--app_firewall--reference--group-001.md#canonical-1221032300100122-3100233131013023-3021323101202122-2031013300301101-0212103322111331-0232232121030201-0010103312223320-0223001012200312): complete subsection reference.

- [blocking_page](resources--app_firewall--reference--group-001.md#canonical-0013022111203330-2133332311031331-2100130303213113-2012101032003221-0021310233223323-3232333233020001-3331122121132323-0213121111201120): complete subsection reference.

- [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-3311120131133012-0031000313131223-3122333020233031-0133330330300232-0200300331201303-3223220213230013-1111003221331133-0311210320311111): complete subsection reference.

- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-0211102231112321-3303232230133311-0000200300103313-2303321211111030-1013331030331131-0002213130220302-2131101302112300-0121333312111003): complete subsection reference.

- [default_anonymization](resources--app_firewall--reference--group-001.md#canonical-1221101323102312-2232212332122002-2023302223332122-3320323011203133-1031203121011123-3010120030311122-3013223022002003-1321122021131022): complete subsection reference.

- [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-1010333202130203-1311300012320212-2220112222022320-0231332302121133-3213002003102111-1011013332021222-3323022012320133-0320311213330220): complete subsection reference.

- [default_detection_settings](resources--app_firewall--reference--group-001.md#canonical-1123321122320132-0232110113121001-0023231010010011-0220123312312011-2211022310022223-0023120121131300-0031222333313133-2310030122330030): complete subsection reference.

<a id="canonical-0331020132100022-2010332103312313-0021011301302113-2120212113321311-2220131122123103-0101321301000010-0333311210232322-0031132001023311"></a>

<a id="canonical-2102223333133313-3110010031303302-0030001131213102-3023120323202231-3323211211311131-2222130323302323-0032110012200111-0231130302310101"></a>

#### `description` property

Type: `"string"`. Optional.

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

- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211): complete subsection reference.

<a id="canonical-3011203302002001-2113133201200202-1203201110001101-3230013001321310-1023100213002033-1102202310020232-3230312212023213-0232103113130232"></a>

<a id="canonical-0322320012332311-1222222011033011-0020010311332322-1211202013332130-1211311102120333-3132012101233323-2121230220123210-0122333002132231"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

- [disable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-1013223333221021-0021212000210331-3303032012100323-1100320130210001-0102031010220003-2230313233132032-3130100220010103-1322001313122100): complete subsection reference.

- [disable_anonymization](resources--app_firewall--reference--group-001.md#canonical-2231213211031203-1320120022121020-1200223313303012-0213232130001230-2210132331310001-3012103103022100-1320131022330110-3033010232213102): complete subsection reference.

- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-1212321310122210-3103000112213223-1213303113222233-3133233012310201-3202312230230022-1003203010121000-2313120111303003-3030300312002202): complete subsection reference.

<a id="canonical-1110122333333030-1122212031332002-0030323312200133-1321301110322033-2323020002013312-1202132330121001-3031022303311310-1320200013131330"></a>

<a id="canonical-1320123221310001-2100020230323202-2021321302300203-0233101222110122-2111303210112223-1020333320023100-1322003200232333-1113103022231130"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3321323033030210-0130113332023202-0202020322233011-2021003301313012-0320213013120232-3212013200232200-2001130033112100-1330012231233010"></a>

<a id="canonical-3013120211033123-1310213223130011-2303232303011003-2102030030313013-2123010233221112-0132231030002131-3121331130002101-0313003232013031"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

- [monitoring](resources--app_firewall--reference--group-001.md#canonical-0022331023311011-2001021100021301-0332031121120320-3203032321113213-2313221311320312-3011321030233100-1123310232003203-0231031201221123): complete subsection reference.

<a id="canonical-2121020102322311-1022231120101012-1031111133301123-3313303300011203-1031010311001323-2310301123301012-0322221133330331-1332020301300310"></a>

<a id="canonical-2122013013231301-2222023121333132-1120000320330013-0211030122113131-2101120211202100-3122111220020122-2033331303331110-1013103332023320"></a>

#### `name` property

Type: `"string"`. Required.

Name of the App Firewall. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-2211100112023312-1113133010020231-1002203320202011-0131113320213322-0100302301003322-0030310202010200-2232012220021320-0013332303303222"></a>

<a id="canonical-0312211201302223-3131303031102222-0333122211120211-1201322111313211-1322303303210302-0232230330012222-2003020012312322-2020000332330310"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the App Firewall is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--app_firewall--reference--group-001.md#canonical-1213311031022333-2212212121033012-0032131123131010-1323220230310300-0200201302130133-3310033023100103-1332330002130123-3211220013131310): complete subsection reference.

- [use_default_blocking_page](resources--app_firewall--reference--group-001.md#canonical-2331102230223121-1110330110101002-3030032133120032-0221012002221232-2210302310222221-3323120210000331-0233330110021300-1022211131012320): complete subsection reference.

<a id="canonical-2211132101102111-1312323133033323-3313002203232322-1123120333111202-2133223310232310-3223321001032212-2332110020212130-1332122223023213"></a>

### All schema paths for `xcsh_app_firewall`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all_response_codes` | [allow_all_response_codes](resources--app_firewall--reference--group-001.md#canonical-0000210332322030-3012313202012323-2003220333132011-2112223312103110-2123230121311222-2321021012000021-1000130022222323-3011131021120103) |
| `allowed_response_codes` | [allowed_response_codes](resources--app_firewall--reference--group-001.md#canonical-0332213133321323-1232202032222310-3021031232023223-3330322123211120-0231301113102033-0102101013232030-0010133110212310-0300232313200313) |
| `allowed_response_codes.response_code` | [allowed_response_codes.response_code](resources--app_firewall--reference--group-001.md#canonical-1213111333001223-3223333122102312-2203100332103011-3321223023310012-2132021323133320-0210012211113113-3020031030022130-3023202113030033) |
| `annotations` | [annotations](resources--app_firewall--reference--group-001.md#canonical-2313220112100333-1003333033210100-3112122021331001-3122101001133320-1122331203012121-1000003132321021-3033130120301330-1023120033101111) |
| `blocking` | [blocking](resources--app_firewall--reference--group-001.md#canonical-1120322000113112-1122321010311110-2000112003022232-3022211131323131-2330220223012111-1222210233133220-3223201322322303-0210221300200110) |
| `blocking_page` | [blocking_page](resources--app_firewall--reference--group-001.md#canonical-3020300102011131-1030130013211322-0322131321032012-2010221133203113-1101220022311010-2033012201300131-2203233122302322-0023133133202310) |
| `blocking_page.blocking_page` | [blocking_page.blocking_page](resources--app_firewall--reference--group-001.md#canonical-3021233213222333-2301233222030123-2033230220020211-0331221232232233-1301220221123120-0230111002312011-0332323012001200-2132311031233101) |
| `blocking_page.response_code` | [blocking_page.response_code](resources--app_firewall--reference--group-001.md#canonical-3000203022113332-2302210323202202-2222022003333331-3101110121301231-0320200320001110-3223000023010111-1323110230032113-3001223323110213) |
| `bot_protection_setting` | [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-1331202031202210-3133300120133300-3333332320020211-2033110000312103-3130133320303323-0231022003323110-2303110131110003-0020111332311123) |
| `bot_protection_setting.good_bot_action` | [bot_protection_setting.good_bot_action](resources--app_firewall--reference--group-001.md#canonical-3232210331112223-3312021132303103-1133210021012030-0201311302000303-0331130123332232-0002232312302113-0331210131302003-1120103221312233) |
| `bot_protection_setting.malicious_bot_action` | [bot_protection_setting.malicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-1221133203100232-2213201201110323-2131133012312221-2222332120213103-0231013112123030-3231032223133330-2033013103333313-3322210320201302) |
| `bot_protection_setting.suspicious_bot_action` | [bot_protection_setting.suspicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-0220021112103202-0311030212213111-1201231022001020-3333311232232030-2030130011210233-3000120301300011-2320003000101210-1133312111120020) |
| `custom_anonymization` | [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-1310013230312220-0221211202230023-2012031120212000-0010222030233112-0101002111120233-3113200121232032-0011332001001101-1330220213310200) |
| `custom_anonymization.anonymization_config` | [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-1231233032031102-3021120213201113-2200323120222122-1001213312310232-2103120103321003-2211210333130103-0021231132302030-2012121211021332) |
| `custom_anonymization.anonymization_config.cookie` | [custom_anonymization.anonymization_config.cookie](resources--app_firewall--reference--group-001.md#canonical-0131323302023032-0312330312233330-2010323032103133-2022012232132020-1030330033200130-1202102331301202-0310332312231122-2031231000321213) |
| `custom_anonymization.anonymization_config.cookie.cookie_name` | [custom_anonymization.anonymization_config.cookie.cookie_name](resources--app_firewall--reference--group-001.md#canonical-0310031210022120-1211212132120132-2110013020332201-2123031001023030-1201203222100012-2031332322300302-3332232130230321-1300210233131021) |
| `custom_anonymization.anonymization_config.http_header` | [custom_anonymization.anonymization_config.http_header](resources--app_firewall--reference--group-001.md#canonical-2211232112113231-0310131121120033-2213332031321221-3213123223231232-2220133132333231-3212103012233223-0220300212320100-0331022201111322) |
| `custom_anonymization.anonymization_config.http_header.header_name` | [custom_anonymization.anonymization_config.http_header.header_name](resources--app_firewall--reference--group-001.md#canonical-1312210010310033-0301301022202232-0221321300013013-2123203221313220-1133211221000131-0031032121320021-2021122232103311-0210231303323121) |
| `custom_anonymization.anonymization_config.query_parameter` | [custom_anonymization.anonymization_config.query_parameter](resources--app_firewall--reference--group-001.md#canonical-0031220313102010-0003123313003203-3110101030320201-1201313022030031-1330103130211333-3320020110001221-1310103112301113-0133130132203011) |
| `custom_anonymization.anonymization_config.query_parameter.query_param_name` | [custom_anonymization.anonymization_config.query_parameter.query_param_name](resources--app_firewall--reference--group-001.md#canonical-3123300222321301-3223013301203300-0203331310001012-1013010023020131-0213100320232103-0230321101232311-0321102022201201-1031120223220000) |
| `default_anonymization` | [default_anonymization](resources--app_firewall--reference--group-001.md#canonical-2113021323231020-2102121123303111-0010101011032023-0130032121323323-1333013230113021-3113022102110130-1130221322002021-2200011121230312) |
| `default_bot_setting` | [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-2301222202313320-1203300110222121-2113203202002113-1300331111202011-3101130231230011-0001013301311123-3333300221213230-0020103033131303) |
| `default_detection_settings` | [default_detection_settings](resources--app_firewall--reference--group-001.md#canonical-0001010310133223-0112231211112330-1113122133300213-0301331331002301-0213013233033101-2011311131033203-0330121022202230-0212020022331213) |
| `description` | [description](resources--app_firewall--reference--group-001.md#canonical-0331020132100022-2010332103312313-0021011301302113-2120212113321311-2220131122123103-0101321301000010-0333311210232322-0031132001023311) |
| `detection_settings` | [detection_settings](resources--app_firewall--reference--group-001.md#canonical-1300320210110301-2222121021222100-1000212033301021-2230003320132233-3132320213222311-3003022210200013-3123023300013330-1331231301302012) |
| `detection_settings.bot_protection_setting` | [detection_settings.bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-0232110321300133-1131311231010032-0103322331003020-3133322223200113-0112001033113212-2231121300313311-2320003103101220-0311200202212222) |
| `detection_settings.bot_protection_setting.good_bot_action` | [detection_settings.bot_protection_setting.good_bot_action](resources--app_firewall--reference--group-001.md#canonical-2321321012303020-1332302301010211-0132131320122001-2231010213113213-3311113221122013-2003000200301132-3011030020201001-0223320232010021) |
| `detection_settings.bot_protection_setting.malicious_bot_action` | [detection_settings.bot_protection_setting.malicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-3310012011303332-1120203133013202-3003002201202211-0230212112322301-2022011203312330-3123211022000230-2210212322000111-1130120010321210) |
| `detection_settings.bot_protection_setting.suspicious_bot_action` | [detection_settings.bot_protection_setting.suspicious_bot_action](resources--app_firewall--reference--group-001.md#canonical-3000101223311332-0321113001132100-0221332012023311-2202333011103130-3310012032312330-0000110003222302-3100201022113002-1030313231121321) |
| `detection_settings.default_bot_setting` | [detection_settings.default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-0011021211103131-0020011221123330-2321003222303111-0122031030112110-2222233210312031-3021220303221103-0302322230330000-2110212330100320) |
| `detection_settings.default_violation_settings` | [detection_settings.default_violation_settings](resources--app_firewall--reference--group-001.md#canonical-3300323332003203-1121333230001132-1032122232100022-2232120302133133-3202322101233030-0003223213002202-3123021303030111-0033010212012123) |
| `detection_settings.disable_staging` | [detection_settings.disable_staging](resources--app_firewall--reference--group-001.md#canonical-3332323121301121-0221132011011030-1031013003102012-3010212112001121-3101231301300033-3213111120031303-3000011130023331-0101333230310313) |
| `detection_settings.disable_suppression` | [detection_settings.disable_suppression](resources--app_firewall--reference--group-001.md#canonical-1012231301001102-3102213131331110-0122333110300230-0311102223032232-3311033310233202-1013012112230113-3123131002110030-1332233121203210) |
| `detection_settings.disable_threat_campaigns` | [detection_settings.disable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-2002120000312313-1201211211122022-3010303021102232-2011030100232131-1320203130111331-1032321331132312-1223311030302132-0013203210020100) |
| `detection_settings.enable_suppression` | [detection_settings.enable_suppression](resources--app_firewall--reference--group-001.md#canonical-0232112132132331-2102213331013013-3011202032330311-0121121010132132-1103010030030201-3313101000303333-2301233313232100-3121021230030103) |
| `detection_settings.enable_threat_campaigns` | [detection_settings.enable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-3113222221122321-0113322221310131-1113322320120022-1103032112123131-3130032321021300-3232000121121211-1030130203311303-2121011332201333) |
| `detection_settings.signature_selection_setting` | [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-1221131221002000-1011303012301331-2211132103012220-2023130122121101-0102332001103333-2211032000311330-2211110212133023-3323233223202010) |
| `detection_settings.signature_selection_setting.attack_type_settings` | [detection_settings.signature_selection_setting.attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-3021113012223220-2001012300033021-2110211021133331-0031123113021230-0010020203123013-3120120103223033-2201202322202332-2131032020332020) |
| `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` | [detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types](resources--app_firewall--reference--group-001.md#canonical-1300203233010320-0331332232001103-1330103200131233-3101110030003130-1210000311331103-3131232110322000-2011312311323032-2322013332133231) |
| `detection_settings.signature_selection_setting.default_attack_type_settings` | [detection_settings.signature_selection_setting.default_attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-2323233122233120-1001202132223021-3232223202123022-1331123030232321-2200313302122111-3113210030103202-2313030111211023-3010221223223331) |
| `detection_settings.signature_selection_setting.default_signature_setting` | [detection_settings.signature_selection_setting.default_signature_setting](resources--app_firewall--reference--group-001.md#canonical-3120330303231310-3210223330231131-3322113021101103-2332303011133303-1112113213301302-2231231200220201-2300101223110331-0223020032333121) |
| `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-2033230300122203-0100032120331023-0200131322232011-0313323000221203-3022101110202003-1022231130303232-1310122321322213-0112021302002103) |
| `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` | [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-3000203113323030-0010333200311311-3321021023211002-2202311133212023-3223031010030020-0033012233202312-3210310133301230-1032011320320112) |
| `detection_settings.signature_selection_setting.only_high_accuracy_signatures` | [detection_settings.signature_selection_setting.only_high_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-3102312321021222-1233113212221120-0302002033232012-0011330212023133-3231002122000323-1221311202122122-3010220201030221-1021221131200320) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy](resources--app_firewall--reference--group-001.md#canonical-0021213220200132-3231021331122123-3332001022021313-2130323123300012-2201011031020130-1302320030001132-0333222030103122-0011020031333220) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action](resources--app_firewall--reference--group-001.md#canonical-2011320123010012-0133330103021121-2213300322011221-3121100002333300-2300202300200122-0122233010020332-0210033003120130-1030132211123030) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action](resources--app_firewall--reference--group-001.md#canonical-3133202132331003-3210120311213301-2100133031031213-2122202311330103-2222111013313201-3000310120010103-1020311133113011-2331300133110211) |
| `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` | [detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action](resources--app_firewall--reference--group-001.md#canonical-0030111131302023-3123033113223320-3213332131121233-2332001300111012-0333221232033230-1203122101000310-3123200031301102-3332033011231000) |
| `detection_settings.stage_new_and_updated_signatures` | [detection_settings.stage_new_and_updated_signatures](resources--app_firewall--reference--group-001.md#canonical-2323222322123203-1002011131011022-1010331212333232-1320311112233100-0233030221320313-0203320201303131-2031130011003230-1213110000303012) |
| `detection_settings.stage_new_and_updated_signatures.staging_period` | [detection_settings.stage_new_and_updated_signatures.staging_period](resources--app_firewall--reference--group-001.md#canonical-2111022200123333-2331200202001201-2032321300102123-2300133212230003-2211223300312131-0003031012000101-1210230131211212-2010210002101020) |
| `detection_settings.stage_new_signatures` | [detection_settings.stage_new_signatures](resources--app_firewall--reference--group-001.md#canonical-2131200210233320-0231000121122213-3100332223010203-3332303302332000-0122211213231333-0200333110102011-2112321002112000-1320103131220301) |
| `detection_settings.stage_new_signatures.staging_period` | [detection_settings.stage_new_signatures.staging_period](resources--app_firewall--reference--group-001.md#canonical-0200203203031112-3330330232211132-0032111231111211-2201231101020230-2112301131011232-0320230332101311-1311312130310303-0301002133021132) |
| `detection_settings.violation_settings` | [detection_settings.violation_settings](resources--app_firewall--reference--group-001.md#canonical-3000133131001331-0222130300211022-3333020112203222-0220012021101313-2021213012301132-0122230031013330-2132022010010100-2131021131121332) |
| `detection_settings.violation_settings.disabled_violation_types` | [detection_settings.violation_settings.disabled_violation_types](resources--app_firewall--reference--group-001.md#canonical-2102032332012021-2101131333113132-3023020130030131-1112003213221021-1032130330310220-0012130230333013-1122131303312221-3223122013211130) |
| `detection_settings.violations_view` | [detection_settings.violations_view](resources--app_firewall--reference--group-001.md#canonical-1333002213101303-0330222222202010-2112333110021031-3331323313230122-3133231210203310-3333103011330131-1231230210203010-2111020320211102) |
| `detection_settings.violations_view.description_spec` | [detection_settings.violations_view.description_spec](resources--app_firewall--reference--group-001.md#canonical-3032123211201101-2302210200131203-0231321122011322-3123031133322222-1312312110301101-0313113332221120-3010103321202020-1301330012113232) |
| `detection_settings.violations_view.enabled` | [detection_settings.violations_view.enabled](resources--app_firewall--reference--group-001.md#canonical-0101122032030323-0033101000013021-3322201100202210-3331031302132011-2232101313232003-2233212121330211-1221313223322323-3312000200202122) |
| `detection_settings.violations_view.enabled_by_default` | [detection_settings.violations_view.enabled_by_default](resources--app_firewall--reference--group-001.md#canonical-0212320122002231-3033103132012230-1012230012333323-1110202333220210-1013012221122120-1332203103323221-0102211333330310-3310023332331220) |
| `detection_settings.violations_view.name` | [detection_settings.violations_view.name](resources--app_firewall--reference--group-001.md#canonical-3123211010232130-1131211021110300-0001120123313012-0121022231101201-2330112101223100-2231330312200120-0033202302302310-0301300330212003) |
| `detection_settings.violations_view.title` | [detection_settings.violations_view.title](resources--app_firewall--reference--group-001.md#canonical-0330300320132030-3310213330032031-3220231311200103-1320323312322031-0012011001331131-3330231211232311-0032032211231302-1232123110102221) |
| `disable` | [disable](resources--app_firewall--reference--group-001.md#canonical-3011203302002001-2113133201200202-1203201110001101-3230013001321310-1023100213002033-1102202310020232-3230312212023213-0232103113130232) |
| `disable_ai_enhancements` | [disable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-1101232212120232-1123203021231211-1001032332213020-3132321013213002-1312033203021122-3301011203113130-1023020232103320-0332221130101122) |
| `disable_anonymization` | [disable_anonymization](resources--app_firewall--reference--group-001.md#canonical-2133311100320313-1113011301130013-3233100100002110-3100113010203031-3002123112020121-2220313021320001-3303232030100202-3213303213132320) |
| `enable_ai_enhancements` | [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-3201113122130332-3200121211211222-1320120013030312-0231022211002202-0231122112122313-0310211031033300-3010030110012233-1221001021223322) |
| `enable_ai_enhancements.mitigate_high_medium_risk_action` | [enable_ai_enhancements.mitigate_high_medium_risk_action](resources--app_firewall--reference--group-001.md#canonical-0313013321213013-2232212020100312-1333010003320001-0213202200303301-0230000001302010-1000023332021320-3212013010302123-2221330210021211) |
| `enable_ai_enhancements.mitigate_high_risk_action` | [enable_ai_enhancements.mitigate_high_risk_action](resources--app_firewall--reference--group-001.md#canonical-3120012303012223-0223113230311222-3333002131112130-2211302301121102-0213320113032323-1313302032312012-1111100031231012-0232000230001323) |
| `id` | [ID](resources--app_firewall--reference--group-001.md#canonical-1110122333333030-1122212031332002-0030323312200133-1321301110322033-2323020002013312-1202132330121001-3031022303311310-1320200013131330) |
| `labels` | [labels](resources--app_firewall--reference--group-001.md#canonical-3321323033030210-0130113332023202-0202020322233011-2021003301313012-0320213013120232-3212013200232200-2001130033112100-1330012231233010) |
| `monitoring` | [monitoring](resources--app_firewall--reference--group-001.md#canonical-0301011303333011-0313111132031222-3223111332210130-2022313202230323-0303001222000312-1000323113031302-0211121311022302-3000211110221202) |
| `name` | [name](resources--app_firewall--reference--group-001.md#canonical-2121020102322311-1022231120101012-1031111133301123-3313303300011203-1031010311001323-2310301123301012-0322221133330331-1332020301300310) |
| `namespace` | [namespace](resources--app_firewall--reference--group-001.md#canonical-2211100112023312-1113133010020231-1002203320202011-0131113320213322-0100302301003322-0030310202010200-2232012220021320-0013332303303222) |
| `timeouts` | [timeouts](resources--app_firewall--reference--group-001.md#canonical-1300032230220332-1003300221232212-3302000313230300-1133002301201111-3020022101223322-1103100113130322-2131230302310022-2212100130203332) |
| `timeouts.create` | [timeouts.create](resources--app_firewall--reference--group-001.md#canonical-2213112212010212-3232223120301000-0220101120011212-3223331113302130-1133232323122131-3223100301013030-2113013111311312-1321210132021003) |
| `timeouts.delete` | [timeouts.delete](resources--app_firewall--reference--group-001.md#canonical-2013020200221313-3002221130110312-3032131310313310-3303201220211322-1310321233220102-0313122300310122-2022033113001202-2000303212032022) |
| `timeouts.read` | [timeouts.read](resources--app_firewall--reference--group-001.md#canonical-0101322323123130-0000033333021213-2133122102322312-3013233333303103-0300010110031220-3111312222313000-2213300210213133-1013122333023213) |
| `timeouts.update` | [timeouts.update](resources--app_firewall--reference--group-001.md#canonical-0330102011131012-1002130333120011-1133003132100213-3211011322332233-3132212331010220-2011010103121120-0201011230211210-2010030213302113) |
| `use_default_blocking_page` | [use_default_blocking_page](resources--app_firewall--reference--group-001.md#canonical-0220123221320131-1321010320020132-0101212120001103-1320010202120213-3011213210331233-2133133313321011-2230301230122032-0320233300001313) |

<a id="canonical-2121321311000013-1301132310203200-0311030332212320-2120130000301202-1020120212020110-1010311311021120-3203102121231022-2020122312221103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all_response_codes` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- allow_all_response_codes

<a id="canonical-0000210332322030-3012313202012323-2003220333132011-2112223312103110-2123230121311222-2321021012000021-1000130022222323-3011131021120103"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [allow_all_response_codes](resources--app_firewall--reference--group-001.md#canonical-0000210332322030-3012313202012323-2003220333132011-2112223312103110-2123230121311222-2321021012000021-1000130022222323-3011131021120103)
- [allowed_response_codes](resources--app_firewall--reference--group-001.md#canonical-0332213133321323-1232202032222310-3021031232023223-3330322123211120-0231301113102033-0102101013232030-0010133110212310-0300232313200313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all_response_codes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120311133313213-1113131100301122-3201330201312210-0023331230030310-2032013032030031-0122113221102121-2233331112123232-0130001130213312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_response_codes` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- allowed_response_codes

<a id="canonical-0332213133321323-1232202032222310-3021031232023223-3330322123211120-0231301113102033-0102101013232030-0010133110212310-0300232313200313"></a>

Type: `"object"`. single nested block, Optional.

List of HTTP response status codes that are allowed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
```

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

Terraform syntax:

```terraform
allowed_response_codes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101321011311330-1313030301221230-3320120110112121-0201111112021203-1203003020111012-3110303310102330-0012001312032033-2001221132233023"></a>

### Direct properties for `allowed_response_codes`

<a id="canonical-1213111333001223-3223333122102312-2203100332103011-3321223023310012-2132021323133320-0210012211113113-3020031030022130-3023202113030033"></a>

#### `allowed_response_codes.response_code` property

Type: `["list", "number"]`. Optional.

List of HTTP response status codes that are allowed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 48),
}
```

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

<a id="canonical-1221032300100122-3100233131013023-3021323101202122-2031013300301101-0212103322111331-0232232121030201-0010103312223320-0223001012200312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocking` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- blocking

<a id="canonical-1120322000113112-1122321010311110-2000112003022232-3022211131323131-2330220223012111-1222210233133220-3223201322322303-0210221300200110"></a>

Type: `["object", {}]`. Optional.

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

- [blocking](resources--app_firewall--reference--group-001.md#canonical-1120322000113112-1122321010311110-2000112003022232-3022211131323131-2330220223012111-1222210233133220-3223201322322303-0210221300200110)
- [monitoring](resources--app_firewall--reference--group-001.md#canonical-0301011303333011-0313111132031222-3223111332210130-2022313202230323-0303001222000312-1000323113031302-0211121311022302-3000211110221202)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocking = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013022111203330-2133332311031331-2100130303213113-2012101032003221-0021310233223323-3232333233020001-3331122121132323-0213121111201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocking_page` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- blocking_page

<a id="canonical-3020300102011131-1030130013211322-0322131321032012-2010221133203113-1101220022311010-2033012201300131-2203233122302322-0023133133202310"></a>

Type: `"object"`. single nested block, Optional.

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

- [blocking_page](resources--app_firewall--reference--group-001.md#canonical-3020300102011131-1030130013211322-0322131321032012-2010221133203113-1101220022311010-2033012201300131-2203233122302322-0023133133202310)
- [use_default_blocking_page](resources--app_firewall--reference--group-001.md#canonical-0220123221320131-1321010320020132-0101212120001103-1320010202120213-3011213210331233-2133133313321011-2230301230122032-0320233300001313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
blocking_page {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203010321123332-2212103310311123-2030200201223301-1133121031020130-2330210100100120-2220033230031011-2101000033100222-3222003031313112"></a>

### Direct properties for `blocking_page`

<a id="canonical-3021233213222333-2301233222030123-2033230220020211-0331221232232233-1301220221123120-0230111002312011-0332323012001200-2132311031233101"></a>

#### `blocking_page.blocking_page` property

Type: `"string"`. Optional.

Define the content of the response page (e.g., an HTML document or a JSON object), use the
\{\{request\_id\}\} placeholder to provide users with a unique identifier to be able to trace the
blocked request in the logs. The maximum allowed size of response body is 4096 bytes after base64
encoding, which would be about 3070 bytes in plain text.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

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

<a id="canonical-3000203022113332-2302210323202202-2222022003333331-3101110121301231-0320200320001110-3223000023010111-1323110230032113-3001223323110213"></a>

<a id="canonical-0102121331200303-1202331302102023-2133301023220333-1213103220312300-2133202321200133-0332120211132000-1331221101123212-2122222002031102"></a>

#### `blocking_page.response_code` property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Accepted","AlreadyReported","BadGateway","BadRequest","Conflict","Continue","Created","EmptyStatusCode","ExpectationFailed","FailedDependency","Forbidden","Found","GatewayTimeout","Gone","HTTPVersionNotSupported","IMUsed","InsufficientStorage","InternalServerError","LengthRequired","Locked","LoopDetected","MethodNotAllowed","MisdirectedRequest","MovedPermanently","MultiStatus","MultipleChoices","NetworkAuthenticationRequired","NoContent","NonAuthoritativeInformation","NotAcceptable","NotExtended","NotFound","NotImplemented","NotModified","OK","PartialContent","PayloadTooLarge","PaymentRequired","PermanentRedirect","PreconditionFailed","PreconditionRequired","ProxyAuthenticationRequired","RangeNotSatisfiable","RequestHeaderFieldsTooLarge","RequestTimeout","ResetContent","SeeOther","ServiceUnavailable","TemporaryRedirect","TooManyRequests","URITooLong","Unauthorized","UnprocessableEntity","UnsupportedMediaType","UpgradeRequired","UseProxy","VariantAlsoNegotiates"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
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
    "NetworkAuthenticationRequired"),
}
```

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

<a id="canonical-3311120131133012-0031000313131223-3122333020233031-0133330330300232-0200300331201303-3223220213230013-1111003221331133-0311210320311111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_protection_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- bot_protection_setting

<a id="canonical-1331202031202210-3133300120133300-3333332320020211-2033110000312103-3130133320303323-0231022003323110-2303110131110003-0020111332311123"></a>

Type: `"object"`. single nested block, Optional.

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

- [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-1331202031202210-3133300120133300-3333332320020211-2033110000312103-3130133320303323-0231022003323110-2303110131110003-0020111332311123)
- [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-2301222202313320-1203300110222121-2113203202002113-1300331111202011-3101130231230011-0001013301311123-3333300221213230-0020103033131303)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bot_protection_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120230202213210-3233302213230010-3330111312212211-0301112013033121-1020111132120331-2022020223303121-2323313010032233-0331132101112232"></a>

### Direct properties for `bot_protection_setting`

<a id="canonical-3232210331112223-3312021132303103-1133210021012030-0201311302000303-0331130123332232-0002232312302113-0331210131302003-1120103221312233"></a>

#### `bot_protection_setting.good_bot_action` property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BLOCK","IGNORE","REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

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

<a id="canonical-1221133203100232-2213201201110323-2131133012312221-2222332120213103-0231013112123030-3231032223133330-2033013103333313-3322210320201302"></a>

<a id="canonical-3111231020102031-0310312312330002-0130231233122103-3311101320331322-2210000101002221-0211023300331132-1313102300201123-2100112131211201"></a>

#### `bot_protection_setting.malicious_bot_action` property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BLOCK","IGNORE","REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

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

<a id="canonical-0220021112103202-0311030212213111-1201231022001020-3333311232232030-2030130011210233-3000120301300011-2320003000101210-1133312111120020"></a>

<a id="canonical-1131301103010100-1111033300032232-3312101131212220-3121131322001102-3232132231120222-1101323011331103-1311102132032212-1212223331111213"></a>

#### `bot_protection_setting.suspicious_bot_action` property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BLOCK","IGNORE","REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

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

<a id="canonical-0211102231112321-3303232230133311-0000200300103313-2303321211111030-1013331030331131-0002213130220302-2131101302112300-0121333312111003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- custom_anonymization

<a id="canonical-1310013230312220-0221211202230023-2012031120212000-0010222030233112-0101002111120233-3113200121232032-0011332001001101-1330220213310200"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_anonymization, default\_anonymization, disable\_anonymization; Default:
default\_anonymization\] Anonymization settings which is a list of HTTP headers, parameters and
cookies.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("anonymization_config")}
```

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

- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-1310013230312220-0221211202230023-2012031120212000-0010222030233112-0101002111120233-3113200121232032-0011332001001101-1330220213310200)
- [default_anonymization](resources--app_firewall--reference--group-001.md#canonical-2113021323231020-2102121123303111-0010101011032023-0130032121323323-1333013230113021-3113022102110130-1130221322002021-2200011121230312)
- [disable_anonymization](resources--app_firewall--reference--group-001.md#canonical-2133311100320313-1113011301130013-3233100100002110-3100113010203031-3002123112020121-2220313021320001-3303232030100202-3213303213132320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_anonymization {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221133330132223-3322021303211221-1110133110122113-1231033122100103-0030200331011123-1102231221133232-2133203211311202-3010322123132210"></a>

### Direct properties for `custom_anonymization`

- [anonymization_config](resources--app_firewall--reference--group-001.md#canonical-0320102321120003-2323323300213221-0312103003232233-2030321100312120-3333131332132210-2212201110013231-2103103023112031-3323000010002123): complete subsection reference.

<a id="canonical-0320102321120003-2323323300213221-0312103003232233-2030321100312120-3333131332132210-2212201110013231-2103103023112031-3323000010002123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-0211102231112321-3303232230133311-0000200300103313-2303321211111030-1013331030331131-0002213130220302-2131101302112300-0121333312111003)
- custom_anonymization.anonymization_config

<a id="canonical-1231233032031102-3021120213201113-2200323120222122-1001213312310232-2103120103321003-2211210333130103-0021231132302030-2012121211021332"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP headers, cookies and query parameters whose values will be masked.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "http_header"),
  validators.ConflictingListObjectAttributes("cookie",
    "query_parameter"),
  validators.ConflictingListObjectAttributes("http_header",
    "query_parameter")}
```

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

Terraform syntax:

```terraform
anonymization_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020233320112110-1213221222301121-2120011000213010-3123202203230103-3321222022233213-2313131130230120-0230323121100231-1221313202232022"></a>

### Direct properties for `custom_anonymization.anonymization_config`

- [cookie](resources--app_firewall--reference--group-001.md#canonical-2101000230003012-2330321113123112-0322310021020121-2201101001123233-2022132320210033-1201223100123233-0121120101321001-3121013323013030): complete subsection reference.

- [http_header](resources--app_firewall--reference--group-001.md#canonical-3321122222022120-0132202001131102-3002332022300023-3021202332030301-1320003031213312-0031033213302222-3033301111122313-0310322102302033): complete subsection reference.

- [query_parameter](resources--app_firewall--reference--group-001.md#canonical-1200131222023102-3200231020312002-2100020013033130-2000332231021313-3121212320133221-2232313223013220-0212122011322311-3032333102100223): complete subsection reference.

<a id="canonical-2101000230003012-2330321113123112-0322310021020121-2201101001123233-2022132320210033-1201223100123233-0121120101321001-3121013323013030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config.cookie` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-0211102231112321-3303232230133311-0000200300103313-2303321211111030-1013331030331131-0002213130220302-2131101302112300-0121333312111003)
- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-0320102321120003-2323323300213221-0312103003232233-2030321100312120-3333131332132210-2212201110013231-2103103023112031-3323000010002123)
- custom_anonymization.anonymization_config.cookie

<a id="canonical-0131323302023032-0312330312233330-2010323032103133-2022012232132020-1030330033200130-1202102331301202-0310332312231122-2031231000321213"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Cookies.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_name")}
```

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

Terraform syntax:

```terraform
cookie {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102333002233230-1233133130220101-2200021110010130-1303300321033213-2012103313033003-0023103122121133-0013022232000310-1123333210122313"></a>

### Direct properties for `custom_anonymization.anonymization_config.cookie`

<a id="canonical-0310031210022120-1211212132120132-2110013020332201-2123031001023030-1201203222100012-2031332322300302-3332232130230321-1300210233131021"></a>

#### `custom_anonymization.anonymization_config.cookie.cookie_name` property

Type: `"string"`. Optional.

Masks the cookie value. The setting does not mask the cookie name. Wildcard matching can be used by
prefixing or suffixing the cookie name with a wildcard asterisk (\*), or by using only an asterisk
to match any cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-3321122222022120-0132202001131102-3002332022300023-3021202332030301-1320003031213312-0031033213302222-3033301111122313-0310322102302033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config.http_header` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-0211102231112321-3303232230133311-0000200300103313-2303321211111030-1013331030331131-0002213130220302-2131101302112300-0121333312111003)
- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-0320102321120003-2323323300213221-0312103003232233-2030321100312120-3333131332132210-2212201110013231-2103103023112031-3323000010002123)
- custom_anonymization.anonymization_config.http_header

<a id="canonical-2211232112113231-0310131121120033-2213332031321221-3213123223231232-2220133132333231-3212103012233223-0220300212320100-0331022201111322"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("header_name")}
```

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

Terraform syntax:

```terraform
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003330302000233-0103001320011101-2333013021112331-0332030132102323-1312233233131023-2132011032321030-1300212022121031-1222002010131201"></a>

### Direct properties for `custom_anonymization.anonymization_config.http_header`

<a id="canonical-1312210010310033-0301301022202232-0221321300013013-2123203221313220-1133211221000131-0031032121320021-2021122232103311-0210231303323121"></a>

#### `custom_anonymization.anonymization_config.http_header.header_name` property

Type: `"string"`. Optional.

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

<a id="canonical-1200131222023102-3200231020312002-2100020013033130-2000332231021313-3121212320133221-2232313223013220-0212122011322311-3032333102100223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_anonymization.anonymization_config.query_parameter` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [custom_anonymization](resources--app_firewall--reference--group-001.md#canonical-0211102231112321-3303232230133311-0000200300103313-2303321211111030-1013331030331131-0002213130220302-2131101302112300-0121333312111003)
- [custom_anonymization.anonymization_config](resources--app_firewall--reference--group-001.md#canonical-0320102321120003-2323323300213221-0312103003232233-2030321100312120-3333131332132210-2212201110013231-2103103023112031-3323000010002123)
- custom_anonymization.anonymization_config.query_parameter

<a id="canonical-0031220313102010-0003123313003203-3110101030320201-1201313022030031-1330103130211333-3320020110001221-1310103112301113-0133130132203011"></a>

Type: `"object"`. single nested block, Optional.

Configure anonymization for HTTP Parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("query_param_name")}
```

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

Terraform syntax:

```terraform
query_parameter {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210302232330200-2113213213001332-1310011302111332-1323003032322131-1131211323313122-2203322233123113-2013033301000200-2112313003302123"></a>

### Direct properties for `custom_anonymization.anonymization_config.query_parameter`

<a id="canonical-3123300222321301-3223013301203300-0203331310001012-1013010023020131-0213100320232103-0230321101232311-0321102022201201-1031120223220000"></a>

#### `custom_anonymization.anonymization_config.query_parameter.query_param_name` property

Type: `"string"`. Optional.

Masks the query parameter value. The setting does not mask the query parameter name. Wildcard
matching can be used by prefixing or suffixing the query parameter name with a wildcard asterisk
(\*), or by using only an asterisk to match any query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

<a id="canonical-1221101323102312-2232212332122002-2023302223332122-3320323011203133-1031203121011123-3010120030311122-3013223022002003-1321122021131022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_anonymization` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- default_anonymization

<a id="canonical-2113021323231020-2102121123303111-0010101011032023-0130032121323323-1333013230113021-3113022102110130-1130221322002021-2200011121230312"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
default_anonymization = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010333202130203-1311300012320212-2220112222022320-0231332302121133-3213002003102111-1011013332021222-3323022012320133-0320311213330220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_bot_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- default_bot_setting

<a id="canonical-2301222202313320-1203300110222121-2113203202002113-1300331111202011-3101130231230011-0001013301311123-3333300221213230-0020103033131303"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
default_bot_setting = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123321122320132-0232110113121001-0023231010010011-0220123312312011-2211022310022223-0023120121131300-0031222333313133-2310030122330030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_detection_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- default_detection_settings

<a id="canonical-0001010310133223-0112231211112330-1113122133300213-0301331331002301-0213013233033101-2011311131033203-0330121022202230-0212020022331213"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [default_detection_settings](resources--app_firewall--reference--group-001.md#canonical-0001010310133223-0112231211112330-1113122133300213-0301331331002301-0213013233033101-2011311131033203-0330121022202230-0212020022331213)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-1300320210110301-2222121021222100-1000212033301021-2230003320132233-3132320213222311-3003022210200013-3123023300013330-1331231301302012)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_detection_settings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- detection_settings

<a id="canonical-1300320210110301-2222121021222100-1000212033301021-2230003320132233-3132320213222311-3003022210200013-3123023300013330-1331231301302012"></a>

Type: `"object"`. single nested block, Optional.

Specifies detection settings to be used by WAF.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_protection_setting",
    "default_bot_setting"),
  validators.ConflictingObjectAttributes("default_violation_settings",
    "violation_settings"),
  validators.ConflictingObjectAttributes("disable_staging",
    "stage_new_and_updated_signatures"),
  validators.ConflictingObjectAttributes("disable_staging",
    "stage_new_signatures"),
  validators.ConflictingObjectAttributes("disable_suppression",
    "enable_suppression"),
  validators.ConflictingObjectAttributes("disable_threat_campaigns",
    "enable_threat_campaigns"),
  validators.ConflictingObjectAttributes("stage_new_and_updated_signatures",
    "stage_new_signatures")}
```

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

Terraform syntax:

```terraform
detection_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133212111101333-3100302231120102-0021033220022010-2113120013331021-1101100113003221-0313002011031313-1202202132123310-0230120012300231"></a>

### Direct properties for `detection_settings`

- [bot_protection_setting](resources--app_firewall--reference--group-001.md#canonical-0013321000200121-2013120103300322-1220212000321001-1020002012213103-2321111020001221-0030111130313320-1003333213110331-2123202200332110): complete subsection reference.

- [default_bot_setting](resources--app_firewall--reference--group-001.md#canonical-1201100000333133-3320301101211013-0232311311000230-1001210110123121-1010302213303200-1310030100111101-2100022013112130-0023012020033131): complete subsection reference.

- [default_violation_settings](resources--app_firewall--reference--group-001.md#canonical-0303303030120203-0221121121321210-0203203001110020-0023211130100030-1300231123121231-3201221220233003-2023203323003110-2300030003131223): complete subsection reference.

- [disable_staging](resources--app_firewall--reference--group-001.md#canonical-1331012223313101-3233100312323200-3312132121221213-0031013320201031-1130322131021112-3111020010212000-1200312120332023-2323303232131123): complete subsection reference.

- [disable_suppression](resources--app_firewall--reference--group-001.md#canonical-2023020110202212-2002211022013321-2212010222113023-2103302131122123-3223221010232312-0103112011323102-0320232231120113-1123332223003322): complete subsection reference.

- [disable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-2233233213120010-1130023211012332-1321302332322111-3100020121012130-3303131223111002-1312002232132313-0003011011123230-3102020201000132): complete subsection reference.

- [enable_suppression](resources--app_firewall--reference--group-001.md#canonical-3132101302330320-1201321312203212-3000331032100022-1130013303033220-0010033100010222-0303311232112302-2100003201213102-0013212031020030): complete subsection reference.

- [enable_threat_campaigns](resources--app_firewall--reference--group-001.md#canonical-0010331011002321-3311321322202303-2232202203321200-0202202332002310-0303300033003300-1103111001133003-0103102222013000-1333030032302133): complete subsection reference.

- [signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303): complete subsection reference.

- [stage_new_and_updated_signatures](resources--app_firewall--reference--group-001.md#canonical-2230223211123000-1130232201323030-3123133033203012-0331321001020220-2033230203201212-3031302023220131-0031301302010303-1212301300303102): complete subsection reference.

- [stage_new_signatures](resources--app_firewall--reference--group-001.md#canonical-1202022030301321-3210231222102302-0033311203113112-3202001300100302-3301322321002130-1113331033222120-2001233200331100-0210022113121301): complete subsection reference.

- [violation_settings](resources--app_firewall--reference--group-001.md#canonical-0130131000110001-0320333030020333-3322032322201112-3201132311330103-3023323101201303-2123031011330231-3010213220210323-0301311330110100): complete subsection reference.

- [violations_view](resources--app_firewall--reference--group-001.md#canonical-2223210232020100-1101120331032303-2012301231233223-1012101301133313-3303013111120223-3032101020232203-1021012202102300-2132131102313021): complete subsection reference.

<a id="canonical-0013321000200121-2013120103300322-1220212000321001-1020002012213103-2321111020001221-0030111130313320-1003333213110331-2123202200332110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.bot_protection_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.bot_protection_setting

<a id="canonical-0232110321300133-1131311231010032-0103322331003020-3133322223200113-0112001033113212-2231121300313311-2320003103101220-0311200202212222"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
bot_protection_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002332123333313-0320301001113132-0210022311112032-1000212213203302-3133120013102330-0220023333121020-3301013313202330-0022300330020213"></a>

### Direct properties for `detection_settings.bot_protection_setting`

<a id="canonical-2321321012303020-1332302301010211-0132131320122001-2231010213113213-3311113221122013-2003000200301132-3011030020201001-0223320232010021"></a>

#### `detection_settings.bot_protection_setting.good_bot_action` property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BLOCK","IGNORE","REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

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

<a id="canonical-3310012011303332-1120203133013202-3003002201202211-0230212112322301-2022011203312330-3123211022000230-2210212322000111-1130120010321210"></a>

<a id="canonical-0031133110002111-3123020121101130-0322321211031120-1002111331310333-3221310130011011-2021022002131000-0010200112203201-3101133031103333"></a>

#### `detection_settings.bot_protection_setting.malicious_bot_action` property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BLOCK","IGNORE","REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

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

<a id="canonical-3000101223311332-0321113001132100-0221332012023311-2202333011103130-3310012032312330-0000110003222302-3100201022113002-1030313231121321"></a>

<a id="canonical-1000202010330220-2301220312321020-1312103023100100-3123210121023110-0211333200021210-1232022201321021-0323012232031012-3211021113002300"></a>

#### `detection_settings.bot_protection_setting.suspicious_bot_action` property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["BLOCK","IGNORE","REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

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

<a id="canonical-1201100000333133-3320301101211013-0232311311000230-1001210110123121-1010302213303200-1310030100111101-2100022013112130-0023012020033131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.default_bot_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.default_bot_setting

<a id="canonical-0011021211103131-0020011221123330-2321003222303111-0122031030112110-2222233210312031-3021220303221103-0302322230330000-2110212330100320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_bot_setting = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303303030120203-0221121121321210-0203203001110020-0023211130100030-1300231123121231-3201221220233003-2023203323003110-2300030003131223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.default_violation_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.default_violation_settings

<a id="canonical-3300323332003203-1121333230001132-1032122232100022-2232120302133133-3202322101233030-0003223213002202-3123021303030111-0033010212012123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_violation_settings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331012223313101-3233100312323200-3312132121221213-0031013320201031-1130322131021112-3111020010212000-1200312120332023-2323303232131123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.disable_staging` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.disable_staging

<a id="canonical-3332323121301121-0221132011011030-1031013003102012-3010212112001121-3101231301300033-3213111120031303-3000011130023331-0101333230310313"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_staging = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023020110202212-2002211022013321-2212010222113023-2103302131122123-3223221010232312-0103112011323102-0320232231120113-1123332223003322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.disable_suppression` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.disable_suppression

<a id="canonical-1012231301001102-3102213131331110-0122333110300230-0311102223032232-3311033310233202-1013012112230113-3123131002110030-1332233121203210"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_suppression = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233233213120010-1130023211012332-1321302332322111-3100020121012130-3303131223111002-1312002232132313-0003011011123230-3102020201000132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.disable_threat_campaigns` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.disable_threat_campaigns

<a id="canonical-2002120000312313-1201211211122022-3010303021102232-2011030100232131-1320203130111331-1032321331132312-1223311030302132-0013203210020100"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_threat_campaigns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132101302330320-1201321312203212-3000331032100022-1130013303033220-0010033100010222-0303311232112302-2100003201213102-0013212031020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.enable_suppression` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.enable_suppression

<a id="canonical-0232112132132331-2102213331013013-3011202032330311-0121121010132132-1103010030030201-3313101000303333-2301233313232100-3121021230030103"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_suppression = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010331011002321-3311321322202303-2232202203321200-0202202332002310-0303300033003300-1103111001133003-0103102222013000-1333030032302133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.enable_threat_campaigns` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.enable_threat_campaigns

<a id="canonical-3113222221122321-0113322221310131-1113322320120022-1103032112123131-3130032321021300-3232000121121211-1030130203311303-2121011332201333"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_threat_campaigns = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.signature_selection_setting

<a id="canonical-1221131221002000-1011303012301331-2211132103012220-2023130122121101-0102332001103333-2211032000311330-2211110212133023-3323233223202010"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures are patterns that identify attacks on a web application and its components.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("attack_type_settings",
    "default_attack_type_settings"),
  validators.ConflictingObjectAttributes("default_signature_setting",
    "signature_settings_by_accuracy"),
  validators.ConflictingObjectAttributes("high_medium_accuracy_signatures",
    "high_medium_low_accuracy_signatures"),
  validators.ConflictingObjectAttributes("high_medium_accuracy_signatures",
    "only_high_accuracy_signatures"),
  validators.ConflictingObjectAttributes("high_medium_low_accuracy_signatures",
    "only_high_accuracy_signatures")}
```

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

Terraform syntax:

```terraform
signature_selection_setting {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023212232210010-0033133330101102-2011010312221122-0011313030323320-3103011211333330-2212323013320110-3131201003120003-2133231120130132"></a>

### Direct properties for `detection_settings.signature_selection_setting`

- [attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-3120030023022331-2032111230100203-2122332231210113-2321300012311120-3012200002111013-3030130213203002-2133001101133302-2200310122021211): complete subsection reference.

- [default_attack_type_settings](resources--app_firewall--reference--group-001.md#canonical-1032030102202302-1123000200220312-2133231131231010-3101300130131001-1102303011003012-0330233100121001-1303031213010211-0211123123013330): complete subsection reference.

- [default_signature_setting](resources--app_firewall--reference--group-001.md#canonical-1223223103321020-3331001121313112-3113130203102122-1303002212032112-0331330003013313-0032330231203310-1320113033222110-3201213001021200): complete subsection reference.

- [high_medium_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-2323120202123213-2003131120121212-1102100333012033-2001123123313122-1013300233031011-1301332302131220-3322101211123233-1211123232232121): complete subsection reference.

- [high_medium_low_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-3331300001301121-1330011111312030-2301320002032231-3011133113210013-2132302221202122-0132012312121301-0320103221313221-2020212003302300): complete subsection reference.

- [only_high_accuracy_signatures](resources--app_firewall--reference--group-001.md#canonical-1231133023311022-0232031032021132-3303320123322313-1021320220113330-1003003123322101-0100301233033332-1102101020010320-1131112121213120): complete subsection reference.

- [signature_settings_by_accuracy](resources--app_firewall--reference--group-001.md#canonical-2133003013202031-1023113013011331-3213301103221203-2031333103131130-2012312120103120-3320310112023101-0300032210212201-3001100113000332): complete subsection reference.

<a id="canonical-3120030023022331-2032111230100203-2122332231210113-2321300012311120-3012200002111013-3030130213203002-2133001101133302-2200310122021211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.attack_type_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303)
- detection_settings.signature_selection_setting.attack_type_settings

<a id="canonical-3021113012223220-2001012300033021-2110211021133331-0031123113021230-0010020203123013-3120120103223033-2201202322202332-2131032020332020"></a>

Type: `"object"`. single nested block, Optional.

Specifies attack-type settings to be used by WAF.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("disabled_attack_types")}
```

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

Terraform syntax:

```terraform
attack_type_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012133302130011-1123002323303223-0121101100033132-0300001330123322-2312200113001013-0123111230102210-1220002203100202-3313022201230031"></a>

### Direct properties for `detection_settings.signature_selection_setting.attack_type_settings`

<a id="canonical-1300203233010320-0331332232001103-1330103200131233-3101110030003130-1210000311331103-3131232110322000-2011312311323032-2322013332133231"></a>

#### `detection_settings.signature_selection_setting.attack_type_settings.disabled_attack_types` property

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(22),
}
```

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

<a id="canonical-1032030102202302-1123000200220312-2133231131231010-3101300130131001-1102303011003012-0330233100121001-1303031213010211-0211123123013330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.default_attack_type_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303)
- detection_settings.signature_selection_setting.default_attack_type_settings

<a id="canonical-2323233122233120-1001202132223021-3232223202123022-1331123030232321-2200313302122111-3113210030103202-2313030111211023-3010221223223331"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_attack_type_settings = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223223103321020-3331001121313112-3113130203102122-1303002212032112-0331330003013313-0032330231203310-1320113033222110-3201213001021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.default_signature_setting` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303)
- detection_settings.signature_selection_setting.default_signature_setting

<a id="canonical-3120330303231310-3210223330231131-3322113021101103-2332303011133303-1112113213301302-2231231200220201-2300101223110331-0223020032333121"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_signature_setting = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323120202123213-2003131120121212-1102100333012033-2001123123313122-1013300233031011-1301332302131220-3322101211123233-1211123232232121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.high_medium_accuracy_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303)
- detection_settings.signature_selection_setting.high_medium_accuracy_signatures

<a id="canonical-2033230300122203-0100032120331023-0200131322232011-0313323000221203-3022101110202003-1022231130303232-1310122321322213-0112021302002103"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
high_medium_accuracy_signatures = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331300001301121-1330011111312030-2301320002032231-3011133113210013-2132302221202122-0132012312121301-0320103221313221-2020212003302300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303)
- detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures

<a id="canonical-3000203113323030-0010333200311311-3321021023211002-2202311133212023-3223031010030020-0033012233202312-3210310133301230-1032011320320112"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
high_medium_low_accuracy_signatures = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231133023311022-0232031032021132-3303320123322313-1021320220113330-1003003123322101-0100301233033332-1102101020010320-1131112121213120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.only_high_accuracy_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303)
- detection_settings.signature_selection_setting.only_high_accuracy_signatures

<a id="canonical-3102312321021222-1233113212221120-0302002033232012-0011330212023133-3231002122000323-1221311202122122-3010220201030221-1021221131200320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
only_high_accuracy_signatures = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133003013202031-1023113013011331-3213301103221203-2031333103131130-2012312120103120-3320310112023101-0300032210212201-3001100113000332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.signature_selection_setting.signature_settings_by_accuracy` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- [detection_settings.signature_selection_setting](resources--app_firewall--reference--group-001.md#canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303)
- detection_settings.signature_selection_setting.signature_settings_by_accuracy

<a id="canonical-0021213220200132-3231021331122123-3332001022021313-2130323123300012-2201011031020130-1302320030001132-0333222030103122-0011020031333220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
signature_settings_by_accuracy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031100002321323-1100332220022202-3301210011021223-3131213011122203-2211331000321211-3232203312203002-3021012330102021-2133011113131003"></a>

### Direct properties for `detection_settings.signature_selection_setting.signature_settings_by_accuracy`

<a id="canonical-2011320123010012-0133330103021121-2213300322011221-3121100002333300-2300202300200122-0122233010020332-0210033003120130-1030132211123030"></a>

#### `detection_settings.signature_selection_setting.signature_settings_by_accuracy.high_accuracy_action` property

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SIG_BLOCK","SIG_IGNORE","SIG_REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

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

<a id="canonical-3133202132331003-3210120311213301-2100133031031213-2122202311330103-2222111013313201-3000310120010103-1020311133113011-2331300133110211"></a>

<a id="canonical-1023100030012312-2213321323100102-1202110013300331-0213222121010222-0121221110111001-2131300023101212-2133302130211031-3021023230321131"></a>

#### `detection_settings.signature_selection_setting.signature_settings_by_accuracy.low_accuracy_action` property

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SIG_BLOCK","SIG_IGNORE","SIG_REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

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

<a id="canonical-0030111131302023-3123033113223320-3213332131121233-2332001300111012-0333221232033230-1203122101000310-3123200031301102-3332033011231000"></a>

<a id="canonical-2320332033231123-2113131300300300-0301033013111120-0322010202233311-0131321333030210-1002111000313011-1002021222331212-1110213122300302"></a>

#### `detection_settings.signature_selection_setting.signature_settings_by_accuracy.medium_accuracy_action` property

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SIG_BLOCK","SIG_IGNORE","SIG_REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

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

<a id="canonical-2230223211123000-1130232201323030-3123133033203012-0331321001020220-2033230203201212-3031302023220131-0031301302010303-1212301300303102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.stage_new_and_updated_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.stage_new_and_updated_signatures

<a id="canonical-2323222322123203-1002011131011022-1010331212333232-1320311112233100-0233030221320313-0203320201303131-2031130011003230-1213110000303012"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures staging configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("staging_period")}
```

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

Terraform syntax:

```terraform
stage_new_and_updated_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333220231210003-2321110020312203-2311012110330310-1213133302002000-1033110100003221-2100211210020231-2223233200133001-2331022110311023"></a>

### Direct properties for `detection_settings.stage_new_and_updated_signatures`

<a id="canonical-2111022200123333-2331200202001201-2032321300102123-2300133212230003-2211223300312131-0003031012000101-1210230131211212-2010210002101020"></a>

#### `detection_settings.stage_new_and_updated_signatures.staging_period` property

Type: `"number"`. Optional.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

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

<a id="canonical-1202022030301321-3210231222102302-0033311203113112-3202001300100302-3301322321002130-1113331033222120-2001233200331100-0210022113121301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.stage_new_signatures` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.stage_new_signatures

<a id="canonical-2131200210233320-0231000121122213-3100332223010203-3332303302332000-0122211213231333-0200333110102011-2112321002112000-1320103131220301"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures staging configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("staging_period")}
```

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

Terraform syntax:

```terraform
stage_new_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100331000021231-0221313221330303-0230233120222321-2331312223010223-3303021123302312-1212131121023311-0013131322210232-2213111201112230"></a>

### Direct properties for `detection_settings.stage_new_signatures`

<a id="canonical-0200203203031112-3330330232211132-0032111231111211-2201231101020230-2112301131011232-0320230332101311-1311312130310303-0301002133021132"></a>

#### `detection_settings.stage_new_signatures.staging_period` property

Type: `"number"`. Optional.

Define staging period in days. The default staging period is 7 days and the max supported staging
period is 20 days.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

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

<a id="canonical-0130131000110001-0320333030020333-3322032322201112-3201132311330103-3023323101201303-2123031011330231-3010213220210323-0301311330110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.violation_settings` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.violation_settings

<a id="canonical-3000133131001331-0222130300211022-3333020112203222-0220012021101313-2021213012301132-0122230031013330-2132022010010100-2131021131121332"></a>

Type: `"object"`. single nested block, Optional.

Specifies violation settings to be used by WAF.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("disabled_violation_types")}
```

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

Terraform syntax:

```terraform
violation_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132132332010300-0133200313021220-2013021120121113-2231310132203221-2300231122210033-3231310301121110-2230020321011322-2001212021322020"></a>

### Direct properties for `detection_settings.violation_settings`

<a id="canonical-2102032332012021-2101131333113132-3023020130030131-1112003213221021-1032130330310220-0012130230333013-1122131303312221-3223122013211130"></a>

#### `detection_settings.violation_settings.disabled_violation_types` property

Type: `["list", "string"]`. Optional, Deprecated.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(40),
}
```

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

<a id="canonical-2223210232020100-1101120331032303-2012301231233223-1012101301133313-3303013111120223-3032101020232203-1021012202102300-2132131102313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `detection_settings.violations_view` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [detection_settings](resources--app_firewall--reference--group-001.md#canonical-2101100131303122-1303001311323033-1203003312130321-1203330310333200-1311033132212102-2130322113021321-3023231122010200-2310122300220211)
- detection_settings.violations_view

<a id="canonical-1333002213101303-0330222222202010-2112333110021031-3331323313230122-3133231210203310-3333103011330131-1231230210203010-2111020320211102"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
violations_view {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010301112331310-0231123303210123-2030230203311132-3110023330033013-3223010322000213-1231031121013213-1021331120231203-0330212023322200"></a>

### Direct properties for `detection_settings.violations_view`

<a id="canonical-3032123211201101-2302210200131203-0231321122011322-3123031133322222-1312312110301101-0313113332221120-3010103321202020-1301330012113232"></a>

#### `detection_settings.violations_view.description_spec` property

Type: `"string"`. Optional.

Description. Human-readable description text

<a id="canonical-0101122032030323-0033101000013021-3322201100202210-3331031302132011-2232101313232003-2233212121330211-1221313223322323-3312000200202122"></a>

<a id="canonical-3221101110111011-2302133222201020-1020120302312221-0333300311201232-1201111011133131-1023022312100313-2032300230211023-1111313021321323"></a>

#### `detection_settings.violations_view.enabled` property

Type: `"bool"`. Optional.

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

<a id="canonical-0212320122002231-3033103132012230-1012230012333323-1110202333220210-1013012221122120-1332203103323221-0102211333330310-3310023332331220"></a>

<a id="canonical-2032201311020312-0133300211311323-3100011220212101-0113222303122133-2101122031131232-0112102131020213-0222321013233230-0002210033012231"></a>

#### `detection_settings.violations_view.enabled_by_default` property

Type: `"string"`. Optional.

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

<a id="canonical-3123211010232130-1131211021110300-0001120123313012-0121022231101201-2330112101223100-2231330312200120-0033202302302310-0301300330212003"></a>

<a id="canonical-2300033130010303-3121001123022320-1202022202023310-2212333210033123-1301300311102312-1330212211300011-1121230030202123-0020332032033103"></a>

#### `detection_settings.violations_view.name` property

Type: `"string"`. Optional.

Name. Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-0330300320132030-3310213330032031-3220231311200103-1320323312322031-0012011001331131-3330231211232311-0032032211231302-1232123110102221"></a>

<a id="canonical-0220330200233122-2123112223003332-2333213201231033-3101100122102031-2233112011220033-1020101113021223-0012123222223113-3100203121303100"></a>

#### `detection_settings.violations_view.title` property

Type: `"string"`. Optional.

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

<a id="canonical-1013223333221021-0021212000210331-3303032012100323-1100320130210001-0102031010220003-2230313233132032-3130100220010103-1322001313122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_ai_enhancements` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- disable_ai_enhancements

<a id="canonical-1101232212120232-1123203021231211-1001032332213020-3132321013213002-1312033203021122-3301011203113130-1023020232103320-0332221130101122"></a>

Type: `["object", {}]`. Optional, Computed.

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

- [disable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-1101232212120232-1123203021231211-1001032332213020-3132321013213002-1312033203021122-3301011203113130-1023020232103320-0332221130101122)
- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-3201113122130332-3200121211211222-1320120013030312-0231022211002202-0231122112122313-0310211031033300-3010030110012233-1221001021223322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ai_enhancements = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231213211031203-1320120022121020-1200223313303012-0213232130001230-2210132331310001-3012103103022100-1320131022330110-3033010232213102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_anonymization` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- disable_anonymization

<a id="canonical-2133311100320313-1113011301130013-3233100100002110-3100113010203031-3002123112020121-2220313021320001-3303232030100202-3213303213132320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_anonymization = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212321310122210-3103000112213223-1213303113222233-3133233012310201-3202312230230022-1003203010121000-2313120111303003-3030300312002202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ai_enhancements` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- enable_ai_enhancements

<a id="canonical-3201113122130332-3200121211211222-1320120013030312-0231022211002202-0231122112122313-0310211031033300-3010030110012233-1221001021223322"></a>

Type: `"object"`. single nested block, Optional.

Actions complimented by the additional intelligence of the F5 AI Powered Risk-based analysis.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("mitigate_high_medium_risk_action",
    "mitigate_high_risk_action")}
```

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

Terraform syntax:

```terraform
enable_ai_enhancements {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211202021033230-3303231321310332-3330221200321113-1023123133230230-3133021112323100-2330002130320301-1012130312010220-0110111003220311"></a>

### Direct properties for `enable_ai_enhancements`

- [mitigate_high_medium_risk_action](resources--app_firewall--reference--group-001.md#canonical-1332001022202133-2213013110133223-0301032020120312-3221331222313310-1132223212021130-1321313111101211-1001103322133021-3010111321301230): complete subsection reference.

- [mitigate_high_risk_action](resources--app_firewall--reference--group-001.md#canonical-3223233223310233-2330130333121002-1323103311003101-0023011211000332-3300332301001200-2112200302113111-1132330020312310-0113123311023221): complete subsection reference.

<a id="canonical-1332001022202133-2213013110133223-0301032020120312-3221331222313310-1132223212021130-1321313111101211-1001103322133021-3010111321301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ai_enhancements.mitigate_high_medium_risk_action` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-1212321310122210-3103000112213223-1213303113222233-3133233012310201-3202312230230022-1003203010121000-2313120111303003-3030300312002202)
- enable_ai_enhancements.mitigate_high_medium_risk_action

<a id="canonical-0313013321213013-2232212020100312-1333010003320001-0213202200303301-0230000001302010-1000023332021320-3212013010302123-2221330210021211"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
mitigate_high_medium_risk_action = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223233223310233-2330130333121002-1323103311003101-0023011211000332-3300332301001200-2112200302113111-1132330020312310-0113123311023221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ai_enhancements.mitigate_high_risk_action` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- [enable_ai_enhancements](resources--app_firewall--reference--group-001.md#canonical-1212321310122210-3103000112213223-1213303113222233-3133233012310201-3202312230230022-1003203010121000-2313120111303003-3030300312002202)
- enable_ai_enhancements.mitigate_high_risk_action

<a id="canonical-3120012303012223-0223113230311222-3333002131112130-2211302301121102-0213320113032323-1313302032312012-1111100031231012-0232000230001323"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
mitigate_high_risk_action = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022331023311011-2001021100021301-0332031121120320-3203032321113213-2313221311320312-3011321030233100-1123310232003203-0231031201221123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `monitoring` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- monitoring

<a id="canonical-0301011303333011-0313111132031222-3223111332210130-2022313202230323-0303001222000312-1000323113031302-0211121311022302-3000211110221202"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
monitoring = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213311031022333-2212212121033012-0032131123131010-1323220230310300-0200201302130133-3310033023100103-1332330002130123-3211220013131310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- timeouts

<a id="canonical-1300032230220332-1003300221232212-3302000313230300-1133002301201111-3020022101223322-1103100113130322-2131230302310022-2212100130203332"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120230222321332-3030211121301113-1021031323110313-3320320011032310-3333223321220332-2010331111112303-2112121203302203-3102213211323301"></a>

### Direct properties for `timeouts`

<a id="canonical-2213112212010212-3232223120301000-0220101120011212-3223331113302130-1133232323122131-3223100301013030-2113013111311312-1321210132021003"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2013020200221313-3002221130110312-3032131310313310-3303201220211322-1310321233220102-0313122300310122-2022033113001202-2000303212032022"></a>

<a id="canonical-1213131231322123-3213120033321301-3001200312330222-1003133221102000-1120210212030332-1301320121013103-2010333233222310-3112210223001012"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0101322323123130-0000033333021213-2133122102322312-3013233333303103-0300010110031220-3111312222313000-2213300210213133-1013122333023213"></a>

<a id="canonical-1103110312102133-1220111310113220-2331332330001112-2110123111003101-3201221332110032-0132033003202000-0321201011000312-3320201131020331"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0330102011131012-1002130333120011-1133003132100213-3211011322332233-3132212331010220-2011010103121120-0201011230211210-2010030213302113"></a>

<a id="canonical-1023231221121203-2221331300200210-0030102211202223-2201202331131002-2121132131000132-2013322221010322-1233120020030302-1013023020200123"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2331102230223121-1110330110101002-3030032133120032-0221012002221232-2210302310222221-3323120210000331-0233330110021300-1022211131012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_blocking_page` properties

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md#canonical-1302202021100230-1023301331311220-3101230121213001-2212232331001200-2013201222332002-1311323202101000-0112031300303303-2023002001023033)
- [Property reference](resources--app_firewall--reference--group-001.md#canonical-0223030301303211-3210320011100320-2311123111323130-2001221232102223-1331203123203211-2111030310300333-2131113321122022-0103121302000233)
- use_default_blocking_page

<a id="canonical-0220123221320131-1321010320020132-0101212120001103-1320010202120213-3011213210331233-2133133313321011-2230301230122032-0320233300001313"></a>

Type: `["object", {}]`. Optional, Computed.

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

Terraform syntax:

```terraform
use_default_blocking_page = {}
```

This is an empty object or choice marker. It has no direct properties.
