---
page_title: "xcsh_virtual_k8s reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s reference."
---

# xcsh_virtual_k8s reference

<a id="canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012231110121201-2011132330003102-1132210211101320-2210100001312130-0301232223303301-0202200311323112-2000230231203220-3112112130331210"></a>

## Property reference — Property reference / 223103333132 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
- Property reference

<a id="canonical-2123311202233231-1310033010301023-3231110010212022-2313113330313203-1022232103300220-0103023301300330-2013213220223221-0203121103022122"></a>

## Direct properties — Property reference / 223103333132 / 3

<a id="canonical-1230202123300021-1200311010231103-3131030003001002-3323203103333102-1201212210220223-1310133302131310-3331032310320321-3122312032201213"></a>

<a id="canonical-3012023323233023-3012302121130002-0211113002123023-2321122303130100-3011122110110131-0001300103320110-2333331022031220-3201122333212223"></a>

## annotations property — Property reference / 223103333132 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

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

- [default_flavor_ref](data-sources--virtual_k8s--reference--group-001.md#canonical-0310331022123312-0232011223312112-3223020002200102-2303230030120131-2223233233201011-0201011211112022-3321001322131101-0112102111212331): complete subsection reference.

<a id="canonical-2303100013321121-1201323122110202-3230213130123010-3312211111111103-0101330330033312-2312330313302222-3333210312320030-3301033312003002"></a>

<a id="canonical-3331232233011133-2302232121013103-1300002323210122-3130233010120211-1301311123223000-3013000213122212-3022100223203100-1103000103031023"></a>

## description property — Property reference / 223103333132 / 5

Type: `"string"`. Computed.

Description of the VirtualK8S.

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

- [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-0103231212031130-1111030212223030-0201311232122103-3011222101321133-1002031202300101-1011013121010321-3130231201312333-0202121011021113): complete subsection reference.

<a id="canonical-3002000030101021-3013011323020323-1203122212011201-0032120000003003-2011103232313321-1000111133122000-2102002033130330-2101300302131331"></a>

<a id="canonical-3211312300113222-2131220213031223-0323103133311220-1331322123213023-0000312331022032-3020312222310303-2331132132233323-3330300333330221"></a>

## ID property — Property reference / 223103333132 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-3301022120013122-3000112301302332-3310330331333013-3312023310123112-0311120310320123-0022110311001122-1003021030311000-3131200122100330): complete subsection reference.

<a id="canonical-3030223312210330-0002221211100021-0022303212133001-3203233331200330-3002003222213200-0233300111103321-1123313023231202-1331332112103331"></a>

<a id="canonical-1110000102102200-1203321021032032-1333212032010221-1310023303123311-3031131121330212-0003003311230103-3333122201113110-1003301331000023"></a>

## labels property — Property reference / 223103333132 / 7

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

<a id="canonical-1121331122333222-2102002122013020-0100331301120123-3101101110012130-1232130012200120-0132012120031022-1203330002303211-3213032123131312"></a>

<a id="canonical-1332301232120302-2300313122211322-2212113331023330-0230102200112201-1010333112020302-3130212200210102-2333022032302002-1302230112223310"></a>

## name property — Property reference / 223103333132 / 8

Type: `"string"`. Required.

Name of the VirtualK8S.

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

<a id="canonical-2121133333103123-3123201001011200-0032211032312332-3212011011012100-1320221232302001-2011223331212113-0032223322120003-2100123103031032"></a>

<a id="canonical-2121100302103032-3122121033022232-2332211113233132-1332333101112300-0123303110122311-3033002121321310-3323023313323030-3111331001313132"></a>

## namespace property — Property reference / 223103333132 / 9

Type: `"string"`. Required.

Namespace where the VirtualK8S exists.

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

- [vsite_refs](data-sources--virtual_k8s--reference--group-001.md#canonical-2220313111103300-3321111021211031-3032231310121133-3130001012030101-1022230330211132-0302212212212220-3013321022121013-3003321003103022): complete subsection reference.

<a id="canonical-0311233211311201-1223100011130013-2002313003202202-3113131231112111-2100100100132112-3330220002001231-2203102032331112-1100123210011330"></a>

## All schema paths — Property reference / 223103333132 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--virtual_k8s--reference--group-001.md#canonical-1230202123300021-1200311010231103-3131030003001002-3323203103333102-1201212210220223-1310133302131310-3331032310320321-3122312032201213) |
| `default_flavor_ref` | [default_flavor_ref](data-sources--virtual_k8s--reference--group-001.md#canonical-0302012310333232-0331210022001332-3323321330312022-2221030012200323-0022112101102212-1120132130303103-0031222333031200-1320331203130031) |
| `default_flavor_ref.name` | [default_flavor_ref.name](data-sources--virtual_k8s--reference--group-001.md#canonical-3030223030110303-1001000003303221-0230221202000033-3011222003233131-1020210302001200-0311233133012100-3013222313023110-0222303122330223) |
| `default_flavor_ref.namespace` | [default_flavor_ref.namespace](data-sources--virtual_k8s--reference--group-001.md#canonical-2121030323021101-2002102333111031-2031121013033111-2023001033310021-1020032101012203-0000203320231031-1311313232122013-3002221020233222) |
| `default_flavor_ref.tenant` | [default_flavor_ref.tenant](data-sources--virtual_k8s--reference--group-001.md#canonical-2303230232210132-0202022122031200-2030031122112100-1301230300031332-2220000100131311-3312013102102320-2111300223021033-1111101333121230) |
| `description` | [description](data-sources--virtual_k8s--reference--group-001.md#canonical-2303100013321121-1201323122110202-3230213130123010-3312211111111103-0101330330033312-2312330313302222-3333210312320030-3301033312003002) |
| `disabled` | [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-0232201221232113-1212111311020203-3302202002300301-2101133322130231-3300323023001023-1131030231221103-1112012123010203-0301223222021000) |
| `id` | [ID](data-sources--virtual_k8s--reference--group-001.md#canonical-3002000030101021-3013011323020323-1203122212011201-0032120000003003-2011103232313321-1000111133122000-2102002033130330-2101300302131331) |
| `isolated` | [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-3312020301313032-3223121312002011-1112221130233102-2121322112022110-0101201023323222-2011312000220031-3012130330221301-2012333210002111) |
| `labels` | [labels](data-sources--virtual_k8s--reference--group-001.md#canonical-3030223312210330-0002221211100021-0022303212133001-3203233331200330-3002003222213200-0233300111103321-1123313023231202-1331332112103331) |
| `name` | [name](data-sources--virtual_k8s--reference--group-001.md#canonical-1121331122333222-2102002122013020-0100331301120123-3101101110012130-1232130012200120-0132012120031022-1203330002303211-3213032123131312) |
| `namespace` | [namespace](data-sources--virtual_k8s--reference--group-001.md#canonical-2121133333103123-3123201001011200-0032211032312332-3212011011012100-1320221232302001-2011223331212113-0032223322120003-2100123103031032) |
| `vsite_refs` | [vsite_refs](data-sources--virtual_k8s--reference--group-001.md#canonical-2201332003332100-3323012011012001-0330123101332001-0133221332010232-2121303331022320-1331333011033013-1030032121332131-0310202031311311) |
| `vsite_refs.kind` | [vsite_refs.kind](data-sources--virtual_k8s--reference--group-001.md#canonical-3123012220231033-2101032332202310-0022100313300122-1220013010121320-3212111013032211-2123310202003202-2301121213203301-1110331131030000) |
| `vsite_refs.name` | [vsite_refs.name](data-sources--virtual_k8s--reference--group-001.md#canonical-0303220131230030-1212023003103101-2121021201003132-3021032121322000-2332002110011230-1121221320202211-1323203001333323-2112031010303300) |
| `vsite_refs.namespace` | [vsite_refs.namespace](data-sources--virtual_k8s--reference--group-001.md#canonical-3301111202323001-3201032331322321-3003330112201301-3212301322230112-0202310211001122-0312132121103331-1032002322013013-2103020202230102) |
| `vsite_refs.tenant` | [vsite_refs.tenant](data-sources--virtual_k8s--reference--group-001.md#canonical-0300132121113100-2123033212200220-3002332303220321-1123201233103100-1321312320002010-2031030000301301-2231210132013113-1232331100120332) |
| `vsite_refs.uid` | [vsite_refs.uid](data-sources--virtual_k8s--reference--group-001.md#canonical-2030031021021021-1110101302123122-0300220021101301-3210033210003311-1023221310312322-1230233301023221-3131002310023320-1331013131312311) |

<a id="canonical-3222122101002211-0032011113300230-1212202003323002-1123123132220131-3132013302130210-0110122210022311-2121101113122111-0113321322120200"></a>

## Next pages — Property reference / 223103333132 / 11

- [default_flavor_ref](data-sources--virtual_k8s--reference--group-001.md#canonical-0310331022123312-0232011223312112-3223020002200102-2303230030120131-2223233233201011-0201011211112022-3321001322131101-0112102111212331)
- [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-0103231212031130-1111030212223030-0201311232122103-3011222101321133-1002031202300101-1011013121010321-3130231201312333-0202121011021113)
- [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-3301022120013122-3000112301302332-3310330331333013-3312023310123112-0311120310320123-0022110311001122-1003021030311000-3131200122100330)
- [vsite_refs](data-sources--virtual_k8s--reference--group-001.md#canonical-2220313111103300-3321111021211031-3032231310121133-3130001012030101-1022230330211132-0302212212212220-3013321022121013-3003321003103022)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)

<a id="canonical-0310331022123312-0232011223312112-3223020002200102-2303230030120131-2223233233201011-0201011211112022-3321001322131101-0112102111212331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310111213313002-3130011212222112-3202110121201131-2310231211133132-0321003320223122-1121100112110100-3023100112012103-0300301113320222"></a>

## default_flavor_ref — default_flavor_ref / 010201020020 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- default_flavor_ref

<a id="canonical-0302012310333232-0331210022001332-3323321330312022-2221030012200323-0022112101102212-1120132130303103-0031222333031200-1320331203130031"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2022100101021131-1020131101032103-0113113222212002-2131131302032131-3010201123112210-2102111023212020-0202013110003331-1323000101320020"></a>

## Direct properties — default_flavor_ref / 010201020020 / 3

<a id="canonical-3030223030110303-1001000003303221-0230221202000033-3011222003233131-1020210302001200-0311233133012100-3013222313023110-0222303122330223"></a>

<a id="canonical-2210212023103032-0020130012032033-3302330311021121-3223233330012313-1023222122032030-1231313122311111-3213313111300330-0131123032123301"></a>

## name property — default_flavor_ref / 010201020020 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2121030323021101-2002102333111031-2031121013033111-2023001033310021-1020032101012203-0000203320231031-1311313232122013-3002221020233222"></a>

<a id="canonical-3310203102030221-3001123130230311-0313001031303032-0122133331102120-3103103020322133-3311002130302301-1302020101003010-0113301000121210"></a>

## namespace property — default_flavor_ref / 010201020020 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2303230232210132-0202022122031200-2030031122112100-1301230300031332-2220000100131311-3312013102102320-2111300223021033-1111101333121230"></a>

<a id="canonical-2313032332112113-2312110330320212-0121311310232021-2333301132013331-2233023100132012-1120222120100030-2013033223031223-2310233133000333"></a>

## tenant property — default_flavor_ref / 010201020020 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2321310031101033-3320212031210333-0030121320020102-0303303232310030-1022203211330320-1003101310121020-1223213011331203-0131333011331003"></a>

## Next pages — default_flavor_ref / 010201020020 / 7

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)

<a id="canonical-0103231212031130-1111030212223030-0201311232122103-3011222101321133-1002031202300101-1011013121010321-3130231201312333-0202121011021113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120211013023322-1233231223021113-3221000020213133-1003102312302231-0233211000013202-0313223001311221-3300233211010002-2100100221122330"></a>

## disabled — disabled / 100131220102 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- disabled

<a id="canonical-0232201221232113-1212111311020203-3302202002300301-2101133322130231-3300323023001023-1131030231221103-1112012123010203-0301223222021000"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disabled, isolated\] Enable this option

Upstream description:

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

- [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-0232201221232113-1212111311020203-3302202002300301-2101133322130231-3300323023001023-1131030231221103-1112012123010203-0301223222021000)
- [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-3312020301313032-3223121312002011-1112221130233102-2121322112022110-0101201023323222-2011312000220031-3012130330221301-2012333210002111)

Select alternatives according to the provider validators above.

<a id="canonical-3303231221312102-2030321032233100-0332300313332322-0213011222111323-3302203032002101-3303212300120133-2333303310022020-0003011132210230"></a>

## Direct properties — disabled / 100131220102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132200132203130-3022301311233320-3212022110223011-3312013320103030-1220211031203332-1210021132310233-0122210331310010-3333012100000012"></a>

## Next pages — disabled / 100131220102 / 4

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)

<a id="canonical-3301022120013122-3000112301302332-3310330331333013-3312023310123112-0311120310320123-0022110311001122-1003021030311000-3131200122100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320232133201232-1103111302031312-3210023202110021-1013000013231132-0220303312213113-1230322200133122-2121300030030032-1132132112033000"></a>

## isolated — isolated / 132200021330 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- isolated

<a id="canonical-3312020301313032-3223121312002011-1112221130233102-2121322112022110-0101201023323222-2011312000220031-3012130330221301-2012333210002111"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2003111203230022-3101211323133031-1311030331332012-1120223130203030-1310123311012031-1003212312333022-2031312013030320-1231132302222030"></a>

## Direct properties — isolated / 132200021330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303322300003331-3012231033323332-1230232211102013-3323112002222313-0102100320220111-3212331120030013-1200100232102103-1102100021123201"></a>

## Next pages — isolated / 132200021330 / 4

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)

<a id="canonical-2220313111103300-3321111021211031-3032231310121133-3130001012030101-1022230330211132-0302212212212220-3013321022121013-3003321003103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030133213223113-3323220233210333-3321211221300033-0220011113103331-1321011132012331-3302331022020131-0220101110100010-1022312313303132"></a>

## vsite_refs — vsite_refs / 233213022113 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- vsite_refs

<a id="canonical-2201332003332100-3323012011012001-0330123101332001-0133221332010232-2121303331022320-1331333011033013-1030032121332131-0310202031311311"></a>

Type: `"list"`. Computed.

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

Upstream description:

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-0332130212313320-0120113111300221-0201210333330013-2021113110033320-2000322110111030-3201123001100302-3202212000133220-0032103021213312"></a>

## Direct properties — vsite_refs / 233213022113 / 3

<a id="canonical-3123012220231033-2101032332202310-0022100313300122-1220013010121320-3212111013032211-2123310202003202-2301121213203301-1110331131030000"></a>

<a id="canonical-2021302103211113-1301100131233110-0220022323132011-2321233230012101-2111100022132021-0313101310303302-0323020103331000-3300123202212130"></a>

## kind property — vsite_refs / 233213022113 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0303220131230030-1212023003103101-2121021201003132-3021032121322000-2332002110011230-1121221320202211-1323203001333323-2112031010303300"></a>

<a id="canonical-1303103330033300-3023121001221033-2110301303001232-0003001223201222-3320311333010320-1001200101013303-1000213220310231-2131103333010220"></a>

## name property — vsite_refs / 233213022113 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3301111202323001-3201032331322321-3003330112201301-3212301322230112-0202310211001122-0312132121103331-1032002322013013-2103020202230102"></a>

<a id="canonical-1310111332233203-2111302311000223-0210220332222032-2311230302100122-2331121033331233-2221031023203013-0321322033003112-2020212321312232"></a>

## namespace property — vsite_refs / 233213022113 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0300132121113100-2123033212200220-3002332303220321-1123201233103100-1321312320002010-2031030000301301-2231210132013113-1232331100120332"></a>

<a id="canonical-3211032202032111-1101111102330010-2122312221323122-1333012003310113-3332120303112113-0111111200201221-3212013123301232-1123020303333132"></a>

## tenant property — vsite_refs / 233213022113 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2030031021021021-1110101302123122-0300220021101301-3210033210003311-1023221310312322-1230233301023221-3131002310023320-1331013131312311"></a>

<a id="canonical-1110103021321203-1102101013112201-0103230233223010-0130231013223122-2223201210201223-2310123110322122-1100031002020331-3032331200233333"></a>

## uid property — vsite_refs / 233213022113 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0010031020012322-0100032203011300-0003123000213302-1132133110031010-0032110003220111-2212322202121003-1330031102010301-2232212031322003"></a>

## Next pages — vsite_refs / 233213022113 / 9

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-1210030023202200-3120102322223222-0332300202003322-3311013210232010-1021131222011021-3110010202333001-0300033310121021-3203223232220130)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0001102031231122-3320312101003201-1113300011213000-3123002303020333-1313210131123221-3100013021110200-0131100111322221-1300103320031332)
