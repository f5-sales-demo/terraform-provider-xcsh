---
page_title: "xcsh_data_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type reference."
---

# xcsh_data_type reference

<a id="canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- Property reference

<a id="canonical-2233011032110300-2112130001112330-2230133301302011-0300023011012321-2003112113022020-1333003002213302-1321121100221302-1130213111130032"></a>

### Direct properties for `xcsh_data_type`

<a id="canonical-0130321202313121-0212033322301002-0020322130301011-0212033212230213-0311113112102320-0011233002000333-2213032231013320-0213213312330212"></a>

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

<a id="canonical-3330313003212001-3131000300010011-3103301100301200-3313320230321331-0031200032031312-0330233001300301-3200020230323210-1331303033313311"></a>

<a id="canonical-2023301331013322-0213332201210133-3221300010301333-2211000011111301-1202222132313322-3002103223333000-2321023311000113-0231023232322302"></a>

#### `compliances` property

Type: `["list", "string"]`. Optional.

\[Enum:
GDPR|CCPA|PIPEDA|LGPD|DPA\_UK|PDPA\_SG|APPI|HIPAA|CPRA\_2023|CPA\_CO|SOC2|PCI\_DSS|ISO\_IEC\_27001|ISO\_IEC\_27701|EPRIVACY\_DIRECTIVE|GLBA|SOX\]
Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform
identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance
breach. Possible values are \`GDPR\`, \`CCPA\`, \`PIPEDA\`, \`LGPD\`, \`DPA\_UK\`, \`PDPA\_SG\`,
\`APPI\`, \`HIPAA\`, \`CPRA\_2023\`, \`CPA\_CO\`, \`SOC2\`, \`PCI\_DSS\`, \`ISO\_IEC\_27001\`,
\`ISO\_IEC\_27701\`, \`EPRIVACY\_DIRECTIVE\`, \`GLBA\`, \`SOX\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2222023123333133-3100120201003110-1231122201011212-3130000000222001-1233011201101212-2212330120221112-2133111223032122-2322220122030301"></a>

<a id="canonical-1100113132302000-1020313033033003-3203121210230132-0023030101303203-2131232101000131-0123311323303331-3332202021201222-1101000312120011"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1330032001000013-1110220011003113-0003331002230213-0233330033032010-3212132313320113-0321122132133023-1211232002102132-3032110211302013"></a>

<a id="canonical-3100112201323330-1000033111301201-1132320203122110-2302033320013102-3203123211321320-2132211101333013-3101323233303123-1000122032020032"></a>

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

<a id="canonical-2203202203020321-0323201032331103-3201211131212100-3202033112031223-2201312211023221-1022120021231332-3013123332031201-0110333222301201"></a>

<a id="canonical-2303221000322113-3012110333123332-1102213323011123-3231033032023210-0212323221332021-1132102123110111-1312013300223303-3111100303312013"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0201133122003322-1311230131211102-2031202123320131-1303323011111122-0311030010313013-2201330021303002-2121100211131310-1112201002011231"></a>

<a id="canonical-2112132211313132-0331331130003121-2232003313231013-0332302201123111-0000313020021013-1021011001330302-3020013030023122-0312310220023322"></a>

#### `is_pii` property

Type: `"bool"`. Optional, Computed.

Select this option to classify the custom data type as personally identifiable information (PII).

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

<a id="canonical-1303022200103100-0130120331300002-1103203013023311-1131132033201131-1221022123120023-2223113101312003-3012232231013320-2000021103311011"></a>

<a id="canonical-1101012121102222-2002222232312123-3211122300123231-2303322231020313-1311001212321221-3113121111102203-3311101020201020-2122210220132212"></a>

#### `is_sensitive_data` property

Type: `"bool"`. Optional, Computed.

Select this option to classify the custom data type as sensitive, enabling detection of API
vulnerabilities related to this data type.

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

<a id="canonical-1222213203313121-0210012131012121-3100103333203110-0030321113003322-3013321333032302-2211220010312201-2333301221010122-1220132313031002"></a>

<a id="canonical-0302332202323222-1223023312113330-3033030310321320-2323033322112023-3021311231120010-1331123102303201-0111133300032021-3132013230031301"></a>

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

<a id="canonical-1232001030000331-3102130310311333-2032003123023102-3331330231121032-1300223230131130-0131321210001231-2323121120210133-2232210021300031"></a>

<a id="canonical-1132202133210032-3121223332000002-3212003332310110-0312231231321221-2210221112301312-2022312222130013-2033112232110302-1010333023031001"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Data Type. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0203102222213233-1331132022333030-3031131022133031-0012023021120312-3103213300221331-3011301112033033-1110010301101202-0013103231003213"></a>

<a id="canonical-1233222023030001-2003011220231301-1331310000110312-2002222001233201-2220300103321103-0310301102022101-3333122323202001-3023010000302301"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Data Type is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333): complete subsection reference.

- [timeouts](resources--data_type--reference--group-001.md#canonical-3112201113323131-3312211332202231-1133121032113233-3111120111123311-0203223302123221-2131311310222010-1111201323012203-1203302320022102): complete subsection reference.

<a id="canonical-0010103212203202-2010230203022212-1323030332123200-1010213000203102-3311100301303300-2113203120333130-3111323122113302-1230113213120202"></a>

### All schema paths for `xcsh_data_type`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--data_type--reference--group-001.md#canonical-0130321202313121-0212033322301002-0020322130301011-0212033212230213-0311113112102320-0011233002000333-2213032231013320-0213213312330212) |
| `compliances` | [compliances](resources--data_type--reference--group-001.md#canonical-3330313003212001-3131000300010011-3103301100301200-3313320230321331-0031200032031312-0330233001300301-3200020230323210-1331303033313311) |
| `description` | [description](resources--data_type--reference--group-001.md#canonical-2222023123333133-3100120201003110-1231122201011212-3130000000222001-1233011201101212-2212330120221112-2133111223032122-2322220122030301) |
| `disable` | [disable](resources--data_type--reference--group-001.md#canonical-1330032001000013-1110220011003113-0003331002230213-0233330033032010-3212132313320113-0321122132133023-1211232002102132-3032110211302013) |
| `id` | [ID](resources--data_type--reference--group-001.md#canonical-2203202203020321-0323201032331103-3201211131212100-3202033112031223-2201312211023221-1022120021231332-3013123332031201-0110333222301201) |
| `is_pii` | [is_pii](resources--data_type--reference--group-001.md#canonical-0201133122003322-1311230131211102-2031202123320131-1303323011111122-0311030010313013-2201330021303002-2121100211131310-1112201002011231) |
| `is_sensitive_data` | [is_sensitive_data](resources--data_type--reference--group-001.md#canonical-1303022200103100-0130120331300002-1103203013023311-1131132033201131-1221022123120023-2223113101312003-3012232231013320-2000021103311011) |
| `labels` | [labels](resources--data_type--reference--group-001.md#canonical-1222213203313121-0210012131012121-3100103333203110-0030321113003322-3013321333032302-2211220010312201-2333301221010122-1220132313031002) |
| `name` | [name](resources--data_type--reference--group-001.md#canonical-1232001030000331-3102130310311333-2032003123023102-3331330231121032-1300223230131130-0131321210001231-2323121120210133-2232210021300031) |
| `namespace` | [namespace](resources--data_type--reference--group-001.md#canonical-0203102222213233-1331132022333030-3031131022133031-0012023021120312-3103213300221331-3011301112033033-1110010301101202-0013103231003213) |
| `rules` | [rules](resources--data_type--reference--group-001.md#canonical-0111100130023103-1211221002133031-1103322333132302-2120113133131122-2320302012121031-1023112121131102-0011000230023031-2113332303020200) |
| `rules.key_pattern` | [rules.key_pattern](resources--data_type--reference--group-001.md#canonical-3102222311123112-1120122303022101-1133231123331232-2221220201301323-2130301020132300-0200332102202021-2210103303022320-0000102012120022) |
| `rules.key_pattern.exact_values` | [rules.key_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-1203330311321330-3000120032203223-3232201030231200-0323330321003311-3022123203020302-2011132003211323-2332021302320331-3322300021330020) |
| `rules.key_pattern.exact_values.exact_values` | [rules.key_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-0213202023222103-0310011312302222-1022222131321102-1003223110033321-1231233310322232-0331330121221033-0300301232210101-0233232110311220) |
| `rules.key_pattern.regex_value` | [rules.key_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-2233133033333011-0322011310210231-1032123200013223-0103110203033222-2112100202301223-1120032322123331-2301220123222023-0130231133120100) |
| `rules.key_pattern.substring_value` | [rules.key_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-0222020003033230-1201031313020310-3231030232312323-0110023211101130-2333100013230110-2232323330113313-2013211310131301-3033110022202230) |
| `rules.key_value_pattern` | [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-2213233101221120-1213313213313213-1323330220330101-3001323303323331-1321000311112000-3232020100033201-3211003201300330-1231102333210123) |
| `rules.key_value_pattern.key_pattern` | [rules.key_value_pattern.key_pattern](resources--data_type--reference--group-001.md#canonical-0012301213010211-1313022013133003-0330030302131121-1010123202223331-2123300222220132-1322023200330031-2032020132200233-3102210010031133) |
| `rules.key_value_pattern.key_pattern.exact_values` | [rules.key_value_pattern.key_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-3013210311303311-3202023232303030-3113031302303231-2000133312022023-2111211030021320-2001220000231111-2313003320333223-0133322013022212) |
| `rules.key_value_pattern.key_pattern.exact_values.exact_values` | [rules.key_value_pattern.key_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-0121320331000332-3102322223220130-1012031323111121-3331220133311321-2302303323211231-3313310023223320-3131210000320302-0203131100312031) |
| `rules.key_value_pattern.key_pattern.regex_value` | [rules.key_value_pattern.key_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-1313223331221311-0132112000132230-3211033232230110-3233232022201023-1030102231022221-1232312002103112-1213021030032220-2110313013022132) |
| `rules.key_value_pattern.key_pattern.substring_value` | [rules.key_value_pattern.key_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-0330231032010111-0031213000302121-1332022131101301-1322232211112222-1203330113102200-3003211101321230-1311213231330033-0001012113333302) |
| `rules.key_value_pattern.value_pattern` | [rules.key_value_pattern.value_pattern](resources--data_type--reference--group-001.md#canonical-1302311330302201-0202313312123032-2323213013131330-0031223013311232-0113131301031022-0213203331121030-0321300312103032-3211232301210131) |
| `rules.key_value_pattern.value_pattern.exact_values` | [rules.key_value_pattern.value_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-3102301002222222-2211131203313012-1101303100211232-1012203021032330-0022030013102131-1020202133213330-3022222120033111-0213310000311330) |
| `rules.key_value_pattern.value_pattern.exact_values.exact_values` | [rules.key_value_pattern.value_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-1233101100131211-1113013113320220-2011300010110020-1131202102323322-0133103323233232-1223101233323332-0113330232312213-0303110210132213) |
| `rules.key_value_pattern.value_pattern.regex_value` | [rules.key_value_pattern.value_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-3222312033111123-3232103201113020-1110232300300301-0131211021132113-0213303200223302-3111312023030313-3113123020313002-1010321212022312) |
| `rules.key_value_pattern.value_pattern.substring_value` | [rules.key_value_pattern.value_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-3021021112323202-0122222220123012-1313323230033031-3201321223332023-1300012312213210-1322120302103132-1131330011123031-3203011201332013) |
| `rules.value_pattern` | [rules.value_pattern](resources--data_type--reference--group-001.md#canonical-0203211133311022-1120311101010022-1211120132223131-2021332320012103-3202230211323321-3301022133030201-1330322111231300-1323211230031333) |
| `rules.value_pattern.exact_values` | [rules.value_pattern.exact_values](resources--data_type--reference--group-001.md#canonical-1302013130331120-0331303003213001-0100132030303121-2310331111102120-2232212102013312-1220322013303311-0323123031231323-3001103031232010) |
| `rules.value_pattern.exact_values.exact_values` | [rules.value_pattern.exact_values.exact_values](resources--data_type--reference--group-001.md#canonical-3221102110302302-1300102333200200-0313132301031213-2023113221021022-2111122113300331-3103323121331121-0321130133131331-1200111333330220) |
| `rules.value_pattern.regex_value` | [rules.value_pattern.regex_value](resources--data_type--reference--group-001.md#canonical-1232111100310113-0222210211213321-3020301312020122-1323330012030301-2200332211202330-2211030320110010-1212103213332322-2231231032000103) |
| `rules.value_pattern.substring_value` | [rules.value_pattern.substring_value](resources--data_type--reference--group-001.md#canonical-3123120023100231-2310213213032122-3222021310312032-1011301221332223-0311231221012011-3201133300300231-1231321213200303-1321030203003323) |
| `timeouts` | [timeouts](resources--data_type--reference--group-001.md#canonical-0100021033000000-1113123023031331-1230122112003313-0320332210331300-3230301103230321-0023222203211220-0322312311312103-0120023120231222) |
| `timeouts.create` | [timeouts.create](resources--data_type--reference--group-001.md#canonical-2130210231022221-1302313132312111-0100311003131112-2331321103223302-2130003101310033-2212212220123320-1313102012013021-1200202232130320) |
| `timeouts.delete` | [timeouts.delete](resources--data_type--reference--group-001.md#canonical-0100303220030211-0121301010020211-1033023102001331-2033333021232000-3121221023332033-3322013222303120-2222000121001021-3232030210312232) |
| `timeouts.read` | [timeouts.read](resources--data_type--reference--group-001.md#canonical-1032103101113100-2321010030002213-0233321103032030-1021323032312001-2213011210133120-3100210303103321-0233221333021003-3330301021100311) |
| `timeouts.update` | [timeouts.update](resources--data_type--reference--group-001.md#canonical-0111002333212210-2212221021013233-0232201000130231-2311020031301021-3020321130223103-2130100312021013-0200122130131230-2101020132010212) |

<a id="canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- rules

<a id="canonical-0111100130023103-1211221002133031-1103322333132302-2120113133131122-2320302012121031-1023112121131102-0011000230023031-2113332303020200"></a>

Type: `"object"`. list nested block, Optional.

Configure key-value or regular expression match rules to enable the platform to detect this custom data type in
the API request or response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("key_pattern",
    "key_value_pattern"),
  validators.ConflictingListObjectAttributes("key_pattern",
    "value_pattern"),
  validators.ConflictingListObjectAttributes("key_value_pattern",
    "value_pattern")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020001112010212-2232103121020323-2120133233132021-0212223303012013-1211332223000213-2211013132001203-0020200333000221-0331102122033221"></a>

### Direct properties for `rules`

- [key_pattern](resources--data_type--reference--group-001.md#canonical-1323101103233130-1033223212223212-0332210012332121-3303211001113021-3110222032231302-0111321202233313-0210223320321222-1213222130312332): complete subsection reference.

- [key_value_pattern](resources--data_type--reference--group-001.md#canonical-2203322001222132-2120310231101230-2012030103310313-2333230303021313-1102131230311020-2223212310230203-3303322231301003-0133032033001311): complete subsection reference.

- [value_pattern](resources--data_type--reference--group-001.md#canonical-0210132302033200-2013032022113300-3003000101013321-2012001022212223-1032031301030003-2333213213102000-3013112330122003-0033012210122212): complete subsection reference.

<a id="canonical-1323101103233130-1033223212223212-0332210012332121-3303211001113021-3110222032231302-0111321202233313-0210223320321222-1213222130312332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.key_pattern` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- rules.key_pattern

<a id="canonical-3102222311123112-1120122303022101-1133231123331232-2221220201301323-2130301020132300-0200332102202021-2210103303022320-0000102012120022"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for key pattern.

Additional upstream details:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
key_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223020301003031-2110010133101132-0332323021311320-0232123212221221-3333312313000131-1232033202023100-0221332030001222-0103003220311313"></a>

### Direct properties for `rules.key_pattern`

- [exact_values](resources--data_type--reference--group-001.md#canonical-3222020332200113-3130221131312203-2110003310110121-3131013230112320-1023231202120012-3020301101123231-0032130333223312-2323032300121023): complete subsection reference.

<a id="canonical-2233133033333011-0322011310210231-1032123200013223-0103110203033222-2112100202301223-1120032322123331-2301220123222023-0130231133120100"></a>

<a id="canonical-1201000113303210-2000220020300131-3011233123212003-2002101021001022-1113030211130210-2012120331301031-3312221030030212-1021002113312112"></a>

#### `rules.key_pattern.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0222020003033230-1201031313020310-3231030232312323-0110023211101130-2333100013230110-2232323330113313-2013211310131301-3033110022202230"></a>

<a id="canonical-2230212302330103-3132331300212130-3333012321213033-2310103213210221-1132200130033213-1111133230001231-0032013201301233-0323120311312330"></a>

#### `rules.key_pattern.substring_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-3222020332200113-3130221131312203-2110003310110121-3131013230112320-1023231202120012-3020301101123231-0032130333223312-2323032300121023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.key_pattern.exact_values` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- [rules.key_pattern](resources--data_type--reference--group-001.md#canonical-1323101103233130-1033223212223212-0332210012332121-3303211001113021-3110222032231302-0111321202233313-0210223320321222-1213222130312332)
- rules.key_pattern.exact_values

<a id="canonical-1203330311321330-3000120032203223-3232201030231200-0323330321003311-3022123203020302-2011132003211323-2332021302320331-3322300021330020"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Additional upstream details:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322001202232131-1322032032200120-3103303201123233-0012211001200013-3030322331300022-3322123331222023-0020233331300232-3030012211201022"></a>

### Direct properties for `rules.key_pattern.exact_values`

<a id="canonical-0213202023222103-0310011312302222-1022222131321102-1003223110033321-1231233310322232-0331330121221033-0300301232210101-0233232110311220"></a>

#### `rules.key_pattern.exact_values.exact_values` property

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2203322001222132-2120310231101230-2012030103310313-2333230303021313-1102131230311020-2223212310230203-3303322231301003-0133032033001311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.key_value_pattern` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- rules.key_value_pattern

<a id="canonical-2213233101221120-1213313213313213-1323330220330101-3001323303323331-1321000311112000-3232020100033201-3211003201300330-1231102333210123"></a>

Type: `"object"`. single nested block, Optional.

Search for specific key &amp; value patterns in the specified sections.

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
key_value_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223233030021200-3313320221021113-0301330110033111-1210102111112322-0102001323333232-2123202221223021-0310311021332231-2013320210202012"></a>

### Direct properties for `rules.key_value_pattern`

- [key_pattern](resources--data_type--reference--group-001.md#canonical-2203231202320212-3131310130323332-0211000021112131-0210002321021002-1133200023101321-1123023102003323-3020012022301122-0300010011021213): complete subsection reference.

- [value_pattern](resources--data_type--reference--group-001.md#canonical-2131001001031100-0201133233332332-0233303120233132-1233011300212301-0311220031301113-0110122312110133-3030231331120303-0232211111230102): complete subsection reference.

<a id="canonical-2203231202320212-3131310130323332-0211000021112131-0210002321021002-1133200023101321-1123023102003323-3020012022301122-0300010011021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.key_value_pattern.key_pattern` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-2203322001222132-2120310231101230-2012030103310313-2333230303021313-1102131230311020-2223212310230203-3303322231301003-0133032033001311)
- rules.key_value_pattern.key_pattern

<a id="canonical-0012301213010211-1313022013133003-0330030302131121-1010123202223331-2123300222220132-1322023200330031-2032020132200233-3102210010031133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for key pattern.

Additional upstream details:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
key_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332330322022330-0003023020000021-1310332000222322-3002230031200020-3331023032101202-3213200113311011-1100131322001122-0202220213101312"></a>

### Direct properties for `rules.key_value_pattern.key_pattern`

- [exact_values](resources--data_type--reference--group-001.md#canonical-0302333232221013-0011212100001221-0331122032202233-3103232223100202-0033213022221110-1010002333011320-2110330202033321-1022320220100233): complete subsection reference.

<a id="canonical-1313223331221311-0132112000132230-3211033232230110-3233232022201023-1030102231022221-1232312002103112-1213021030032220-2110313013022132"></a>

<a id="canonical-1020223211010323-2302120331021022-3311322201031310-0112013232332302-1223023322000023-0231301030021131-1133010330313002-3231001130331310"></a>

#### `rules.key_value_pattern.key_pattern.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0330231032010111-0031213000302121-1332022131101301-1322232211112222-1203330113102200-3003211101321230-1311213231330033-0001012113333302"></a>

<a id="canonical-3122003210210021-1102133221120322-2301301011330201-0012222203222011-1120131123121231-3210313100313211-3232221322012003-3130330022201011"></a>

#### `rules.key_value_pattern.key_pattern.substring_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-0302333232221013-0011212100001221-0331122032202233-3103232223100202-0033213022221110-1010002333011320-2110330202033321-1022320220100233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.key_value_pattern.key_pattern.exact_values` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-2203322001222132-2120310231101230-2012030103310313-2333230303021313-1102131230311020-2223212310230203-3303322231301003-0133032033001311)
- [rules.key_value_pattern.key_pattern](resources--data_type--reference--group-001.md#canonical-2203231202320212-3131310130323332-0211000021112131-0210002321021002-1133200023101321-1123023102003323-3020012022301122-0300010011021213)
- rules.key_value_pattern.key_pattern.exact_values

<a id="canonical-3013210311303311-3202023232303030-3113031302303231-2000133312022023-2111211030021320-2001220000231111-2313003320333223-0133322013022212"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Additional upstream details:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113221120312322-2222022223330323-2330120120313122-3332123101130223-3030302131200003-1222112013122320-1331002030331111-0132110213111013"></a>

### Direct properties for `rules.key_value_pattern.key_pattern.exact_values`

<a id="canonical-0121320331000332-3102322223220130-1012031323111121-3331220133311321-2302303323211231-3313310023223320-3131210000320302-0203131100312031"></a>

#### `rules.key_value_pattern.key_pattern.exact_values.exact_values` property

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131001001031100-0201133233332332-0233303120233132-1233011300212301-0311220031301113-0110122312110133-3030231331120303-0232211111230102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.key_value_pattern.value_pattern` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-2203322001222132-2120310231101230-2012030103310313-2333230303021313-1102131230311020-2223212310230203-3303322231301003-0133032033001311)
- rules.key_value_pattern.value_pattern

<a id="canonical-1302311330302201-0202313312123032-2323213013131330-0031223013311232-0113131301031022-0213203331121030-0321300312103032-3211232301210131"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for value pattern.

Additional upstream details:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
value_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013112120303323-0300033310322120-1013111002321010-2023301213231021-2322003330103223-0332300303103113-2121313103020112-1201300001233032"></a>

### Direct properties for `rules.key_value_pattern.value_pattern`

- [exact_values](resources--data_type--reference--group-001.md#canonical-2102300102311332-0011000011222313-2313223022000202-3003210011003022-1020323323302331-2101230130212220-1133231231131021-1301002320310200): complete subsection reference.

<a id="canonical-3222312033111123-3232103201113020-1110232300300301-0131211021132113-0213303200223302-3111312023030313-3113123020313002-1010321212022312"></a>

<a id="canonical-3213331133221030-0211220103101300-2330333301000210-1111202012101213-1320030002113102-1311331310001020-1103220220013311-3012122101103111"></a>

#### `rules.key_value_pattern.value_pattern.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3021021112323202-0122222220123012-1313323230033031-3201321223332023-1300012312213210-1322120302103132-1131330011123031-3203011201332013"></a>

<a id="canonical-0133232333100123-2310332200111001-1012021211212013-0130211012200131-2230110010310221-3212021123330201-1203131121000210-3330303310113023"></a>

#### `rules.key_value_pattern.value_pattern.substring_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-2102300102311332-0011000011222313-2313223022000202-3003210011003022-1020323323302331-2101230130212220-1133231231131021-1301002320310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.key_value_pattern.value_pattern.exact_values` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- [rules.key_value_pattern](resources--data_type--reference--group-001.md#canonical-2203322001222132-2120310231101230-2012030103310313-2333230303021313-1102131230311020-2223212310230203-3303322231301003-0133032033001311)
- [rules.key_value_pattern.value_pattern](resources--data_type--reference--group-001.md#canonical-2131001001031100-0201133233332332-0233303120233132-1233011300212301-0311220031301113-0110122312110133-3030231331120303-0232211111230102)
- rules.key_value_pattern.value_pattern.exact_values

<a id="canonical-3102301002222222-2211131203313012-1101303100211232-1012203021032330-0022030013102131-1020202133213330-3022222120033111-0213310000311330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Additional upstream details:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0133011003022021-1020222010000320-0221313030022111-2103133012231231-1220222222331131-3302103320232130-2331301321000202-3302023301133203"></a>

### Direct properties for `rules.key_value_pattern.value_pattern.exact_values`

<a id="canonical-1233101100131211-1113013113320220-2011300010110020-1131202102323322-0133103323233232-1223101233323332-0113330232312213-0303110210132213"></a>

#### `rules.key_value_pattern.value_pattern.exact_values.exact_values` property

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0210132302033200-2013032022113300-3003000101013321-2012001022212223-1032031301030003-2333213213102000-3013112330122003-0033012210122212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.value_pattern` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- rules.value_pattern

<a id="canonical-0203211133311022-1120311101010022-1211120132223131-2021332320012103-3202230211323321-3301022133030201-1330322111231300-1323211230031333"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for value pattern.

Additional upstream details:

Test

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_values",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_values",
    "substring_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "substring_value")}
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
  "x-ves-oneof-field-type_choice": "[\"exact_values\",\"regex_value\",\"substring_value\"]"
}
```

Terraform syntax:

```terraform
value_pattern {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221011320033221-1322121320223133-1210230220211231-2101130023222102-2103122331113133-0011232130212101-1102333001013121-1102310131120032"></a>

### Direct properties for `rules.value_pattern`

- [exact_values](resources--data_type--reference--group-001.md#canonical-0311113201132310-0330321303103120-0113100013323112-3011031130033321-0100113103330302-3101120312111212-1322011200200300-2223021233033121): complete subsection reference.

<a id="canonical-1232111100310113-0222210211213321-3020301312020122-1323330012030301-2200332211202330-2211030320110010-1212103213332322-2231231032000103"></a>

<a id="canonical-0200113301101230-0122203203023010-1203231021222232-3330031221033030-3211212120312133-1033101130020230-1112031301213311-3213322020020032"></a>

#### `rules.value_pattern.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3123120023100231-2310213213032122-3222021310312032-1011301221332223-0311231221012011-3201133300300231-1231321213200303-1321030203003323"></a>

<a id="canonical-2230200000320331-2022331233300221-0203131200333301-2101220132022211-1023001200012232-3203322213222022-3220222110003032-3333001212212323"></a>

#### `rules.value_pattern.substring_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-0311113201132310-0330321303103120-0113100013323112-3011031130033321-0100113103330302-3101120312111212-1322011200200300-2223021233033121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.value_pattern.exact_values` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- [rules](resources--data_type--reference--group-001.md#canonical-0013102232230332-1230112101113132-0002130131130213-1120123010331313-1223112222221313-1022100132010130-3003101200231123-0010230022231333)
- [rules.value_pattern](resources--data_type--reference--group-001.md#canonical-0210132302033200-2013032022113300-3003000101013321-2012001022212223-1032031301030003-2333213213102000-3013112330122003-0033012210122212)
- rules.value_pattern.exact_values

<a id="canonical-1302013130331120-0331303003213001-0100132030303121-2310331111102120-2232212102013312-1220322013303311-0323123031231323-3001103031232010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for exact values.

Additional upstream details:

List of exact values to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("exact_values")}
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
exact_values {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332212022013003-2033212111112333-3320331130111323-3311211113311033-3021233312113133-1330131112010223-0033002322312122-1222012032032020"></a>

### Direct properties for `rules.value_pattern.exact_values`

<a id="canonical-3221102110302302-1300102333200200-0313132301031213-2023113221021022-2111122113300331-3103323121331121-0321130133131331-1200111333330220"></a>

#### `rules.value_pattern.exact_values.exact_values` property

Type: `["list", "string"]`. Optional.

Exact Values. List of exact values to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3112201113323131-3312211332202231-1133121032113233-3111120111123311-0203223302123221-2131311310222010-1111201323012203-1203302320022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_data_type](../resources/data_type.md#canonical-0130333301110020-0132111313313303-3030102230323221-2030031322131321-0220013133320111-2333311320030320-2201101033200331-2212312313233130)
- [Property reference](resources--data_type--reference--group-001.md#canonical-0023102330323310-3033213000323320-2320222030130123-2213110132221320-3102131103230320-3032011020130102-3220121232123130-1021121011232020)
- timeouts

<a id="canonical-0100021033000000-1113123023031331-1230122112003313-0320332210331300-3230301103230321-0023222203211220-0322312311312103-0120023120231222"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101203023312103-1022202210330323-1302333111323022-3022111223312001-1200223233220030-1010220123212113-3002333233012120-1021003012013031"></a>

### Direct properties for `timeouts`

<a id="canonical-2130210231022221-1302313132312111-0100311003131112-2331321103223302-2130003101310033-2212212220123320-1313102012013021-1200202232130320"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0100303220030211-0121301010020211-1033023102001331-2033333021232000-3121221023332033-3322013222303120-2222000121001021-3232030210312232"></a>

<a id="canonical-2203313310123310-1030131031030013-2233320330112302-3301021201120211-2111021121133211-1230031333011223-2103023202311223-2330021312222032"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1032103101113100-2321010030002213-0233321103032030-1021323032312001-2213011210133120-3100210303103321-0233221333021003-3330301021100311"></a>

<a id="canonical-2030013310100113-3231120132323111-1222132320113300-2300233203001010-0322302110000022-1001301100112002-0230031211110313-1333200110202130"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0111002333212210-2212221021013233-0232201000130231-2311020031301021-3020321130223103-2130100312021013-0200122130131230-2101020132010212"></a>

<a id="canonical-0113230320121121-1131012102333322-3101123012212333-0023300210022212-0323302121113302-1330033120023200-2231320012122023-1213130010121011"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
