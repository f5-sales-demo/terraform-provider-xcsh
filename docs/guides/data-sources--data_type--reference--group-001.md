---
page_title: "xcsh_data_type reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type reference."
---

# xcsh_data_type reference

<a id="canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003203122013133-2212031133311011-3313121101212201-2313122330101123-2231111102322222-1210212100303322-0223100021120201-2003102001200033"></a>

## Property reference — Property reference / 113020023121 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- Property reference

<a id="canonical-0222213311123030-2300320201202003-3231002122303133-0213012203012211-0033100331221211-3301120032120312-2123313011103012-0000333223011101"></a>

## Direct properties — Property reference / 113020023121 / 3

<a id="canonical-2100012210100031-1011022212331001-2132303010022100-3033322122101100-1210123212310301-0121300130333222-3311331313322333-2002310322123333"></a>

<a id="canonical-1330123023321200-3020312132111112-1011321213021100-3110201020101003-0333331213231232-0102132321320113-2322012222210003-2112221322101012"></a>

## annotations property — Property reference / 113020023121 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-2231111232323002-0233331120020233-3213010112333320-2103330200222323-1031302030100023-0012212202231221-2213002232301030-0223102121313023"></a>

<a id="canonical-3130123112310033-3302232021133012-1002013310222321-1022120123112111-2002302212022101-0233333313301233-0303302213131333-1222202113112232"></a>

## compliances property — Property reference / 113020023121 / 5

Type: `["list", "string"]`. Computed.

\[Enum:
GDPR|CCPA|PIPEDA|LGPD|DPA\_UK|PDPA\_SG|APPI|HIPAA|CPRA\_2023|CPA\_CO|SOC2|PCI\_DSS|ISO\_IEC\_27001|ISO\_IEC\_27701|EPRIVACY\_DIRECTIVE|GLBA|SOX\]
Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform
identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance
breach. Possible values are \`GDPR\`, \`CCPA\`, \`PIPEDA\`, \`LGPD\`, \`DPA\_UK\`, \`PDPA\_SG\`,
\`APPI\`, \`HIPAA\`, \`CPRA\_2023\`, \`CPA\_CO\`, \`SOC2\`, \`PCI\_DSS\`, \`ISO\_IEC\_27001\`,
\`ISO\_IEC\_27701\`, \`EPRIVACY\_DIRECTIVE\`, \`GLBA\`, \`SOX\`.

Upstream description:

Choose applicable compliance frameworks such as GDPR, PCI/DSS, or CCPA to ensure the platform
identifies whether vulnerabilities in API endpoints handling this data type may cause a compliance
breach.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1230310202011333-3003332102331003-0232231330102130-2012220322100220-3331010300312222-3301032221220322-1333202230231311-2011003222123323"></a>

<a id="canonical-2132210021021030-0013213332232102-0101302322020211-2132313300110013-3313112012022312-0013211320110302-3003103011331233-3013123201031330"></a>

## description property — Property reference / 113020023121 / 6

Type: `"string"`. Computed.

Description of the DataType.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0021113012101231-0321121223311232-3200302202232110-0013320032101200-1112102123231123-2321013210213020-1211001202212001-0011330010331331"></a>

<a id="canonical-0311003133130211-3012021231010112-3121122012313313-1031300301320133-0230110110331333-1022211022123302-1333020323310301-1003313003212021"></a>

## ID property — Property reference / 113020023121 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1002130312112131-2322230101331232-0213310233130211-0023232221131223-2231133222022312-2101123310022020-1303022222123213-2011213123101123"></a>

<a id="canonical-3133222312021232-1321312211031302-3130113002033320-0231030313201012-3330023131231112-2211102230101133-3333223100100313-3313332113121000"></a>

## is_pii property — Property reference / 113020023121 / 8

Type: `"bool"`. Computed.

Select this option to classify the custom data type as personally identifiable information (PII).

Upstream description:

Select this option to classify the custom data type as personally identifiable information (PII)

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

<a id="canonical-1122113023211013-0310223221313301-3303132331301330-0222320001010210-2111233222030122-1111033033011113-3002033321320030-0130301330222233"></a>

<a id="canonical-1123113112211300-3201222212222112-1221302131220201-1323112112211132-2322231011012121-2221233001232122-2022312032021120-3321213302132222"></a>

## is_sensitive_data property — Property reference / 113020023121 / 9

Type: `"bool"`. Computed.

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

<a id="canonical-3311103100231013-1311202002112322-1023010111312121-0233122302033312-1030310021003212-0022000200123111-0211210000130012-0102130222202203"></a>

<a id="canonical-0301023213130321-2012320003221212-0203031302031220-1133111332103332-2103011313023310-2223222301001331-0220030013321210-1133101012031013"></a>

## labels property — Property reference / 113020023121 / 10

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

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

<a id="canonical-3212220223233323-3311112220313303-2322200302113302-2233302012212320-0001233323223021-0031131310031210-0112223333201010-1302223202133300"></a>

<a id="canonical-1111202232131112-1210300032230000-0312001031201301-0010211032032131-3213313023213212-0033232011110220-2032213330220002-3221021332000220"></a>

## name property — Property reference / 113020023121 / 11

Type: `"string"`. Required.

Name of the DataType.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0300300100130202-2212200030021202-3100312023302331-3310030232220330-2323300213010121-2231122210202110-1023223311303302-3131323022111333"></a>

<a id="canonical-1302030332203012-2123223012300121-3203031220101200-3121122210011320-2213230211110211-0133010100330023-3013113320222020-3002300301313122"></a>

## namespace property — Property reference / 113020023121 / 12

Type: `"string"`. Required.

Namespace where the DataType exists.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311): complete subsection reference.

<a id="canonical-3232323023331213-3013132022231133-2330320311003012-2211122132223202-1231300200300322-1102312230122011-1223203130212223-3222210023133302"></a>

## All schema paths — Property reference / 113020023121 / 13

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--data_type--reference--group-001.md#canonical-2100012210100031-1011022212331001-2132303010022100-3033322122101100-1210123212310301-0121300130333222-3311331313322333-2002310322123333) |
| `compliances` | [compliances](data-sources--data_type--reference--group-001.md#canonical-2231111232323002-0233331120020233-3213010112333320-2103330200222323-1031302030100023-0012212202231221-2213002232301030-0223102121313023) |
| `description` | [description](data-sources--data_type--reference--group-001.md#canonical-1230310202011333-3003332102331003-0232231330102130-2012220322100220-3331010300312222-3301032221220322-1333202230231311-2011003222123323) |
| `id` | [ID](data-sources--data_type--reference--group-001.md#canonical-0021113012101231-0321121223311232-3200302202232110-0013320032101200-1112102123231123-2321013210213020-1211001202212001-0011330010331331) |
| `is_pii` | [is_pii](data-sources--data_type--reference--group-001.md#canonical-1002130312112131-2322230101331232-0213310233130211-0023232221131223-2231133222022312-2101123310022020-1303022222123213-2011213123101123) |
| `is_sensitive_data` | [is_sensitive_data](data-sources--data_type--reference--group-001.md#canonical-1122113023211013-0310223221313301-3303132331301330-0222320001010210-2111233222030122-1111033033011113-3002033321320030-0130301330222233) |
| `labels` | [labels](data-sources--data_type--reference--group-001.md#canonical-3311103100231013-1311202002112322-1023010111312121-0233122302033312-1030310021003212-0022000200123111-0211210000130012-0102130222202203) |
| `name` | [name](data-sources--data_type--reference--group-001.md#canonical-3212220223233323-3311112220313303-2322200302113302-2233302012212320-0001233323223021-0031131310031210-0112223333201010-1302223202133300) |
| `namespace` | [namespace](data-sources--data_type--reference--group-001.md#canonical-0300300100130202-2212200030021202-3100312023302331-3310030232220330-2323300213010121-2231122210202110-1023223311303302-3131323022111333) |
| `rules` | [rules](data-sources--data_type--reference--group-001.md#canonical-1102023132003220-1322221133122330-0320100031121322-1220103102033320-0323231101000113-2013033323022102-2311232230122320-1301222030133102) |
| `rules.key_pattern` | [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-2113200022233213-1233322300302331-1313321320131202-2313120032203021-3201302023300123-1033313103200331-2030223031023202-3031303032102031) |
| `rules.key_pattern.exact_values` | [rules.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-1121103103221303-2000120203200300-2302202322203310-3021223310220122-3312332302023101-3310300133000103-3000001332120133-3212010102222302) |
| `rules.key_pattern.exact_values.exact_values` | [rules.key_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-0033232313102323-3320222132032010-2332321001101231-3103001133121330-0213313010300222-3321113313010130-2213311322021230-2022312113302212) |
| `rules.key_pattern.regex_value` | [rules.key_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-2300322132322303-3203132300003312-1011300031300030-3113302333202020-0033002200212333-0221111030332323-2212230102231321-0120313101113120) |
| `rules.key_pattern.substring_value` | [rules.key_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-2310301312230022-3312020100200320-3210001212201201-2030103122223302-1221223313102120-3133331201321203-2213012212300031-3022301301113330) |
| `rules.key_value_pattern` | [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-1111302203322120-3112320233223210-0333010211220210-0212023333311112-3011102202013201-3023112023202010-2322123333313103-1333222322000132) |
| `rules.key_value_pattern.key_pattern` | [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-2203121020232001-1221313122000221-1012302111103132-1101332133022312-2310321202303130-0022323312012120-3002032323223333-0022312301313123) |
| `rules.key_value_pattern.key_pattern.exact_values` | [rules.key_value_pattern.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-0203013022210000-2011320200031320-1300103033333201-2302333013033322-2110001332333320-3331221030112103-1132321123133212-3030013300312212) |
| `rules.key_value_pattern.key_pattern.exact_values.exact_values` | [rules.key_value_pattern.key_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-3222102022230220-3001300203302330-2103302021031011-1001020130000202-2130031320213020-2112232121032320-3130102020321220-2233233100213222) |
| `rules.key_value_pattern.key_pattern.regex_value` | [rules.key_value_pattern.key_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-0312130012233321-3213331310313022-2113223201113032-2322130302210232-3301010011222102-1113320023020321-1331303301332323-1113311123311030) |
| `rules.key_value_pattern.key_pattern.substring_value` | [rules.key_value_pattern.key_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-0222132211032333-2220111220112102-2322303232031231-1010032111113323-2020302211133212-2200132103110113-0002311121102312-1320210001000001) |
| `rules.key_value_pattern.value_pattern` | [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-2033233111320001-0002030320012322-1032231103311010-3013013330303223-2202233301031102-2300322111331122-0022022133120301-2122000002110321) |
| `rules.key_value_pattern.value_pattern.exact_values` | [rules.key_value_pattern.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-1131111020133101-3130220131111020-0222211200300000-0001101300121032-1231232023220202-1313002230031013-0121131131110003-0020120130103302) |
| `rules.key_value_pattern.value_pattern.exact_values.exact_values` | [rules.key_value_pattern.value_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-2112303223232302-1120032022030130-2231013011111222-1231223123133323-1033120202202310-3101023022203302-3101200010101211-0100301323330312) |
| `rules.key_value_pattern.value_pattern.regex_value` | [rules.key_value_pattern.value_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-0001130233222223-0312330112132233-2002323303220232-0012032032103131-3001322111311320-1300211213022300-1023113323000010-1101332113232332) |
| `rules.key_value_pattern.value_pattern.substring_value` | [rules.key_value_pattern.value_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-3201031233031333-2200131312021230-2201101301211330-2032222133303110-3113120100323220-1120010203022031-3003310320030231-2032320213232012) |
| `rules.value_pattern` | [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-3101203223002211-2301330001200112-2012211213211020-2002313201201220-2002230111003221-2121230302323112-2133213013032031-1032333321222011) |
| `rules.value_pattern.exact_values` | [rules.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-2131033023100001-1010002032003231-1032220220031112-2303000322212023-3030033111221123-0300321322111113-3013300332020300-2030211203001332) |
| `rules.value_pattern.exact_values.exact_values` | [rules.value_pattern.exact_values.exact_values](data-sources--data_type--reference--group-001.md#canonical-2030012232201211-2132203012003021-3211113202130003-3232333012123013-3033003313213330-1111211320230103-2323130133310330-3320102023322102) |
| `rules.value_pattern.regex_value` | [rules.value_pattern.regex_value](data-sources--data_type--reference--group-001.md#canonical-1102111103311021-3111203003312322-0330301112100130-0310220132203233-3221001321130332-3203123111122321-1330033022322103-1211013220220010) |
| `rules.value_pattern.substring_value` | [rules.value_pattern.substring_value](data-sources--data_type--reference--group-001.md#canonical-3010000303130303-0203223310323333-0201113003212120-2322302231000311-1331300231123130-1201012310123320-3113123233220011-1233032303323200) |

<a id="canonical-0310121000002323-2211021231020210-1120022320130230-2213322013132023-0312001223002203-1222232011101212-1012110030031331-2322312302010001"></a>

## Next pages — Property reference / 113020023121 / 14

- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133133201220300-3112003133222300-1311131122302312-2232323323032211-1233132033310010-3032101102013020-2133103013123321-2113312111223033"></a>

## rules — rules / 320010010303 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- rules

<a id="canonical-1102023132003220-1322221133122330-0320100031121322-1220103102033320-0323231101000113-2013033323022102-2311232230122320-1301222030133102"></a>

Type: `"list"`. Computed.

Configure key-value or regular expression match rules to enable the platform to detect this custom data type in
the API request or response.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0300030132133221-0332222230102312-3131301313210101-1213202020032223-1332031203102113-1323033122113321-0022112322132110-2113130122220202"></a>

## Direct properties — rules / 320010010303 / 3

- [key_pattern](data-sources--data_type--reference--group-001.md#canonical-0121320013020002-2121012313103223-3122202112022112-1131200013020113-3120022101203010-1120211322330200-2230300213212102-1301120233130213): complete subsection reference.

- [key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312): complete subsection reference.

- [value_pattern](data-sources--data_type--reference--group-001.md#canonical-2311201012331032-3110222023000021-1022013100310301-2103103233333220-3003111113333102-2222233001122101-2321203300011132-2301112113332222): complete subsection reference.

<a id="canonical-0323213210021112-1203033111020000-3122233320102230-0120213203021032-1322023301301323-1113011232301303-0301223131320301-0230220103311022"></a>

## Next pages — rules / 320010010303 / 4

- [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-0121320013020002-2121012313103223-3122202112022112-1131200013020113-3120022101203010-1120211322330200-2230300213212102-1301120233130213)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312)
- [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-2311201012331032-3110222023000021-1022013100310301-2103103233333220-3003111113333102-2222233001122101-2321203300011132-2301112113332222)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-0121320013020002-2121012313103223-3122202112022112-1131200013020113-3120022101203010-1120211322330200-2230300213212102-1301120233130213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320233332323132-3300302210013033-1000331231012103-1312213200100320-2301330002333123-0111303001021201-2122132120200030-1300132213303231"></a>

## rules.key_pattern — key_pattern / 303000102111 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- rules.key_pattern

<a id="canonical-2113200022233213-1233322300302331-1313321320131202-2313120032203021-3201302023300123-1033313103200331-2030223031023202-3031303032102031"></a>

Type: `"single"`. Computed.

Configuration parameter for key pattern.

Upstream description:

Test

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

<a id="canonical-1022111300130210-3202010011031000-1131322221030113-0211231112101013-0100020301233332-1021032222211202-2211230013331201-3020330022213133"></a>

## Direct properties — key_pattern / 303000102111 / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-1133021210322100-0122031232222220-1100020121100033-0023123200122333-2013102012100333-1200131103322133-1030303233231221-2103023002301132): complete subsection reference.

<a id="canonical-2300322132322303-3203132300003312-1011300031300030-3113302333202020-0033002200212333-0221111030332323-2212230102231321-0120313101113120"></a>

<a id="canonical-2123331121300022-3222012003320223-3101111132232132-3323330333201310-0301101303021100-0220101220102012-0010332331131200-3331221131231002"></a>

## regex_value property — key_pattern / 303000102111 / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2310301312230022-3312020100200320-3210001212201201-2030103122223302-1221223313102120-3133331201321203-2213012212300031-3022301301113330"></a>

<a id="canonical-1030002112300130-0333223201321313-2212031331121211-0020112200310320-0133122000110233-1222201100002102-1233201330103302-0211033020003013"></a>

## substring_value property — key_pattern / 303000102111 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2230020333031333-0122001230020023-0213023112333103-3000102012202230-0312203222111221-3121130301013330-3111210322220220-3220223323210113"></a>

## Next pages — key_pattern / 303000102111 / 6

- [rules.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-1133021210322100-0122031232222220-1100020121100033-0023123200122333-2013102012100333-1200131103322133-1030303233231221-2103023002301132)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-1133021210322100-0122031232222220-1100020121100033-0023123200122333-2013102012100333-1200131103322133-1030303233231221-2103023002301132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322321023201133-3023003101230330-2330000020122021-0121200202220012-1011131100221023-1302030122001121-2030031103302002-3331003103021200"></a>

## rules.key_pattern.exact_values — exact_values / 101212231133 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-0121320013020002-2121012313103223-3122202112022112-1131200013020113-3120022101203010-1120211322330200-2230300213212102-1301120233130213)
- rules.key_pattern.exact_values

<a id="canonical-1121103103221303-2000120203200300-2302202322203310-3021223310220122-3312332302023101-3310300133000103-3000001332120133-3212010102222302"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-2310022213121100-3222221010310223-1020211322012001-1301002031300313-2033101321033003-1120320333122330-2121333020322132-1122133111000103"></a>

## Direct properties — exact_values / 101212231133 / 3

<a id="canonical-0033232313102323-3320222132032010-2332321001101231-3103001133121330-0213313010300222-3321113313010130-2213311322021230-2022312113302212"></a>

<a id="canonical-3123000111222232-0313101110202102-0132032323131003-2231322203020310-3131123202322222-1311011211000032-0320200030223211-3222122303023013"></a>

## exact_values property — exact_values / 101212231133 / 4

Type: `["list", "string"]`. Computed.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2201000303133031-2232022102010213-3201222102030330-3321233223321011-0330232223213321-3332302010100123-0303033000322220-2101200231231203"></a>

## Next pages — exact_values / 101212231133 / 5

- [rules.key_pattern](data-sources--data_type--reference--group-001.md#canonical-0121320013020002-2121012313103223-3122202112022112-1131200013020113-3120022101203010-1120211322330200-2230300213212102-1301120233130213)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012323103233013-3103110121113221-3011303120112130-3203301001003110-2331312203013120-1211312230321111-0132331001003123-2012200021000201"></a>

## rules.key_value_pattern — key_value_pattern / 232003330202 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- rules.key_value_pattern

<a id="canonical-1111302203322120-3112320233223210-0333010211220210-0212023333311112-3011102202013201-3023112023202010-2322123333313103-1333222322000132"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0321023131120113-2300020313312131-1121032301031220-2301130121122032-0203313331122310-1213023232301203-3302322133110132-0032313130310233"></a>

## Direct properties — key_value_pattern / 232003330202 / 3

- [key_pattern](data-sources--data_type--reference--group-001.md#canonical-1111131211022103-1022013203103023-3022022323101233-3112002102001101-3001202222010203-3132000221200102-3013013011200230-1000211123332000): complete subsection reference.

- [value_pattern](data-sources--data_type--reference--group-001.md#canonical-2232113200010323-2211310113320110-3031230312333122-3222010310313310-3300302310312023-0322330201223302-2213310010223032-1221210100122211): complete subsection reference.

<a id="canonical-3030332032112302-0232002210020100-1100123131132012-2321233322210313-3111233200321112-2233130012103333-2222022221010333-0131011322120012"></a>

## Next pages — key_value_pattern / 232003330202 / 4

- [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-1111131211022103-1022013203103023-3022022323101233-3112002102001101-3001202222010203-3132000221200102-3013013011200230-1000211123332000)
- [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-2232113200010323-2211310113320110-3031230312333122-3222010310313310-3300302310312023-0322330201223302-2213310010223032-1221210100122211)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-1111131211022103-1022013203103023-3022022323101233-3112002102001101-3001202222010203-3132000221200102-3013013011200230-1000211123332000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031233330012202-1012320233123210-0311020331121111-2113311012013020-2201331213331002-3220210121020121-1112221122232122-3210002120302321"></a>

## rules.key_value_pattern.key_pattern — key_pattern / 231101023311 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312)
- rules.key_value_pattern.key_pattern

<a id="canonical-2203121020232001-1221313122000221-1012302111103132-1101332133022312-2310321202303130-0022323312012120-3002032323223333-0022312301313123"></a>

Type: `"single"`. Computed.

Configuration parameter for key pattern.

Upstream description:

Test

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

<a id="canonical-3231323232130010-1211003121230113-2200120010203103-1222110120122310-1210011331132312-3133111232033123-1212123022203033-2031211302301332"></a>

## Direct properties — key_pattern / 231101023311 / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-1010323223303123-0112133331023033-3210322221311312-3121011310222322-0332233111312320-1033032002001200-3113002003322213-0130333222110011): complete subsection reference.

<a id="canonical-0312130012233321-3213331310313022-2113223201113032-2322130302210232-3301010011222102-1113320023020321-1331303301332323-1113311123311030"></a>

<a id="canonical-0332013111020101-0023323220222220-1203003001131230-0222020022230112-1000213030322032-2002202012213202-3201100323112230-2220301112320120"></a>

## regex_value property — key_pattern / 231101023311 / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0222132211032333-2220111220112102-2322303232031231-1010032111113323-2020302211133212-2200132103110113-0002311121102312-1320210001000001"></a>

<a id="canonical-3231011333121213-3330300120332330-0303120101131013-1130112222303010-2220321223000020-1213310030221102-2310022303331323-1000220121211013"></a>

## substring_value property — key_pattern / 231101023311 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0211303110301002-0102110202132121-3033331100111123-3203102313113233-3123211330131212-1030001330302301-1301013130201301-2322312312320300"></a>

## Next pages — key_pattern / 231101023311 / 6

- [rules.key_value_pattern.key_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-1010323223303123-0112133331023033-3210322221311312-3121011310222322-0332233111312320-1033032002001200-3113002003322213-0130333222110011)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-1010323223303123-0112133331023033-3210322221311312-3121011310222322-0332233111312320-1033032002001200-3113002003322213-0130333222110011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213310133013012-3332302232331321-0021103033100002-3332210130000133-3122103122312322-1030021102023111-3301333211220211-2320200232331100"></a>

## rules.key_value_pattern.key_pattern.exact_values — exact_values / 230132321003 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312)
- [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-1111131211022103-1022013203103023-3022022323101233-3112002102001101-3001202222010203-3132000221200102-3013013011200230-1000211123332000)
- rules.key_value_pattern.key_pattern.exact_values

<a id="canonical-0203013022210000-2011320200031320-1300103033333201-2302333013033322-2110001332333320-3331221030112103-1132321123133212-3030013300312212"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-2001033103001222-1020001020112322-1103212323103110-3203120103032032-2320202123121020-1331302100031312-2130030223011120-2002220023321333"></a>

## Direct properties — exact_values / 230132321003 / 3

<a id="canonical-3222102022230220-3001300203302330-2103302021031011-1001020130000202-2130031320213020-2112232121032320-3130102020321220-2233233100213222"></a>

<a id="canonical-1230022220302310-2031303013002031-2211101222332310-1300100111001113-3111003132122233-2103032013300101-0032321021103003-0313002301303333"></a>

## exact_values property — exact_values / 230132321003 / 4

Type: `["list", "string"]`. Computed.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3301033102121000-3320331030031002-1333133131331121-0332231123232320-3122002112222211-0021001222110112-3231303222000313-0330003333030321"></a>

## Next pages — exact_values / 230132321003 / 5

- [rules.key_value_pattern.key_pattern](data-sources--data_type--reference--group-001.md#canonical-1111131211022103-1022013203103023-3022022323101233-3112002102001101-3001202222010203-3132000221200102-3013013011200230-1000211123332000)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-2232113200010323-2211310113320110-3031230312333122-3222010310313310-3300302310312023-0322330201223302-2213310010223032-1221210100122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230213212301203-1311113130013323-3213321023333003-1112100221301130-1201201323023130-0330233300330112-0310133330300222-0132033010320120"></a>

## rules.key_value_pattern.value_pattern — value_pattern / 222113122322 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312)
- rules.key_value_pattern.value_pattern

<a id="canonical-2033233111320001-0002030320012322-1032231103311010-3013013330303223-2202233301031102-2300322111331122-0022022133120301-2122000002110321"></a>

Type: `"single"`. Computed.

Configuration parameter for value pattern.

Upstream description:

Test

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

<a id="canonical-1301131220331302-3110231312013110-2111312103122123-3033301021133133-0323202311311103-0202310100030210-3310312023032200-0012221113321023"></a>

## Direct properties — value_pattern / 222113122322 / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-1303221331013302-1230301312031222-3003223303000031-2110222020203030-3332201130232200-2203133213011302-2301332203232200-0333230231221012): complete subsection reference.

<a id="canonical-0001130233222223-0312330112132233-2002323303220232-0012032032103131-3001322111311320-1300211213022300-1023113323000010-1101332113232332"></a>

<a id="canonical-0011021133113003-3310200202002030-2002333310232313-3012023313120102-0221011100221202-2303203320313222-1313232321121110-2012232222232311"></a>

## regex_value property — value_pattern / 222113122322 / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3201031233031333-2200131312021230-2201101301211330-2032222133303110-3113120100323220-1120010203022031-3003310320030231-2032320213232012"></a>

<a id="canonical-2133002223012012-3313113123013303-3101201210322230-2210020322230302-2020030320101231-3201102023103120-2312031330131233-2132230331201200"></a>

## substring_value property — value_pattern / 222113122322 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1011100202223332-3210222310333203-2210110121222003-2232020111100220-0030120203000011-3110030132203303-3032130031030113-0230032302330313"></a>

## Next pages — value_pattern / 222113122322 / 6

- [rules.key_value_pattern.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-1303221331013302-1230301312031222-3003223303000031-2110222020203030-3332201130232200-2203133213011302-2301332203232200-0333230231221012)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-1303221331013302-1230301312031222-3003223303000031-2110222020203030-3332201130232200-2203133213011302-2301332203232200-0333230231221012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031200210120331-1131010133122331-3312301322113033-2212202233000030-3202200320321300-0211031121220003-1211121311332210-2202112310033030"></a>

## rules.key_value_pattern.value_pattern.exact_values — exact_values / 132313100030 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [rules.key_value_pattern](data-sources--data_type--reference--group-001.md#canonical-0102131130013322-1332131032332000-2113231210301231-3023100201331020-2311113103333121-2100310221333033-1332032333211112-1222223132121312)
- [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-2232113200010323-2211310113320110-3031230312333122-3222010310313310-3300302310312023-0322330201223302-2213310010223032-1221210100122211)
- rules.key_value_pattern.value_pattern.exact_values

<a id="canonical-1131111020133101-3130220131111020-0222211200300000-0001101300121032-1231232023220202-1313002230031013-0121131131110003-0020120130103302"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-0320120330303122-0203330013300230-2012123011323102-3013001020011110-1321120220300000-0003102010332032-3010123030033102-0011311123232130"></a>

## Direct properties — exact_values / 132313100030 / 3

<a id="canonical-2112303223232302-1120032022030130-2231013011111222-1231223123133323-1033120202202310-3101023022203302-3101200010101211-0100301323330312"></a>

<a id="canonical-1111011211113312-1113000023112303-1230032030030323-2020010311110001-1020300203011130-3122331311300010-2330100221002222-2031003211302321"></a>

## exact_values property — exact_values / 132313100030 / 4

Type: `["list", "string"]`. Computed.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3321110231113102-0222300320233320-0022002010033330-2011203011221021-0012022320101220-1122103321133321-0100303211001010-2312111120123323"></a>

## Next pages — exact_values / 132313100030 / 5

- [rules.key_value_pattern.value_pattern](data-sources--data_type--reference--group-001.md#canonical-2232113200010323-2211310113320110-3031230312333122-3222010310313310-3300302310312023-0322330201223302-2213310010223032-1221210100122211)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-2311201012331032-3110222023000021-1022013100310301-2103103233333220-3003111113333102-2222233001122101-2321203300011132-2301112113332222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331233321100123-1201210301323113-0303323223300231-2111102013201211-1330202013313320-1220022202130311-1111032331323032-2031203201333231"></a>

## rules.value_pattern — value_pattern / 302013331332 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- rules.value_pattern

<a id="canonical-3101203223002211-2301330001200112-2012211213211020-2002313201201220-2002230111003221-2121230302323112-2133213013032031-1032333321222011"></a>

Type: `"single"`. Computed.

Configuration parameter for value pattern.

Upstream description:

Test

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

<a id="canonical-3313030013022103-3120112230303331-3331300111311100-2312223122233133-2122302133133132-3332002210130223-3010320022331210-3231233323211131"></a>

## Direct properties — value_pattern / 302013331332 / 3

- [exact_values](data-sources--data_type--reference--group-001.md#canonical-0222222011132111-0330311021301101-0203311003212001-3001023201000103-1122133311320200-1130112111122130-0303100011202321-3332221311100330): complete subsection reference.

<a id="canonical-1102111103311021-3111203003312322-0330301112100130-0310220132203233-3221001321130332-3203123111122321-1330033022322103-1211013220220010"></a>

<a id="canonical-2123322121130132-3100320111313331-0033311120101233-0111011300333133-1213030102323103-3022323122202313-3233132103130010-0331021031111111"></a>

## regex_value property — value_pattern / 302013331332 / 4

Type: `"string"`. Computed.

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

Upstream description:

Exclusive with \[exact\_values substring\_value\] Search for values matching this regular
expression.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3010000303130303-0203223310323333-0201113003212120-2322302231000311-1331300231123130-1201012310123320-3113123233220011-1233032303323200"></a>

<a id="canonical-0113313102333032-1121130023200021-0212020012301222-0230031222111220-2313312021101112-1311213333222030-2032133100302331-3320332012020111"></a>

## substring_value property — value_pattern / 302013331332 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

Upstream description:

Exclusive with \[exact\_values regular expression\_value\] Search for values that include this substring.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2233311111133013-0021203023012231-3011031312203331-3200110203211100-1312000132012230-1103232302302330-2132011020022111-2133131022321021"></a>

## Next pages — value_pattern / 302013331332 / 6

- [rules.value_pattern.exact_values](data-sources--data_type--reference--group-001.md#canonical-0222222011132111-0330311021301101-0203311003212001-3001023201000103-1122133311320200-1130112111122130-0303100011202321-3332221311100330)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)

<a id="canonical-0222222011132111-0330311021301101-0203311003212001-3001023201000103-1122133311320200-1130112111122130-0303100011202321-3332221311100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321322011130232-3020200101001222-2332202112323020-3230233122312332-1011121121220003-1212030113312032-2020031213303230-2120330301331333"></a>

## rules.value_pattern.exact_values — exact_values / 203011023312 / 2

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
- [Property reference](data-sources--data_type--reference--group-001.md#canonical-0101121221321220-0300032012130112-3013110233001110-1201101201213122-1232011210133132-2033200300332030-1002231313212311-0232103231121012)
- [rules](data-sources--data_type--reference--group-001.md#canonical-0333333102113332-1130203111213230-0322211311000122-3122201320013033-1032002310111120-3213132130302321-1303212303202121-1000201022213311)
- [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-2311201012331032-3110222023000021-1022013100310301-2103103233333220-3003111113333102-2222233001122101-2321203300011132-2301112113332222)
- rules.value_pattern.exact_values

<a id="canonical-2131033023100001-1010002032003231-1032220220031112-2303000322212023-3030033111221123-0300321322111113-3013300332020300-2030211203001332"></a>

Type: `"single"`. Computed.

Configuration parameter for exact values.

Upstream description:

List of exact values to match.

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

<a id="canonical-1110110113210232-2222131110330323-1210130120313223-3312020033223200-3103023100120310-3212231221212110-2012002100302312-1310113320130201"></a>

## Direct properties — exact_values / 203011023312 / 3

<a id="canonical-2030012232201211-2132203012003021-3211113202130003-3232333012123013-3033003313213330-1111211320230103-2323130133310330-3320102023322102"></a>

<a id="canonical-0220321220113113-3101111020222003-1010131010011121-2221003133311010-0112132112130100-0130001322313200-2012130020131310-0030012122011120"></a>

## exact_values property — exact_values / 203011023312 / 4

Type: `["list", "string"]`. Computed.

Exact Values. List of exact values to match.

Upstream description:

List of exact values to match.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2000331302012003-1010222000333012-2022310210112332-0232200021000131-1313301323032012-3031103103133210-1320021232001330-0203031122022001"></a>

## Next pages — exact_values / 203011023312 / 5

- [rules.value_pattern](data-sources--data_type--reference--group-001.md#canonical-2311201012331032-3110222023000021-1022013100310301-2103103233333220-3003111113333102-2222233001122101-2321203300011132-2301112113332222)
- [xcsh_data_type](../data-sources/data_type.md#canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133)
