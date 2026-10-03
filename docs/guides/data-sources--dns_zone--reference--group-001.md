---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200120133223003-1320023021013013-0122102132103020-1223133312200130-2032230022210312-0001321322002012-3212323203200212-3113021201232000"></a>

## Property reference — Property reference / 121200323121 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- Property reference

<a id="canonical-2102100130111323-1230123120130122-0333232300100001-0223120021103230-0030022133030200-1201333101202210-0200301001330130-3330213322321233"></a>

## Direct properties — Property reference / 121200323121 / 3

<a id="canonical-3221313102233023-3020301321310200-0111300210103000-2202211230033103-0321320110120003-1003232312202033-3032022323011201-0321232133213221"></a>

<a id="canonical-1233203333212331-1201011013101331-1203012030020123-1133313010322212-0222230332013310-0102111210111123-1231220222213321-0232213011030033"></a>

## annotations property — Property reference / 121200323121 / 4

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

<a id="canonical-3331313120113002-1202330323333013-1113000020020211-3000103210033330-3300213221310233-2130120021202010-1002233000203300-3322000333302332"></a>

<a id="canonical-1303323101002012-2313002321303222-0113212302001012-1302210233320301-3203303312230022-0210102201110310-2320011312101332-0023121123231232"></a>

## description property — Property reference / 121200323121 / 5

Type: `"string"`. Computed.

Description of the DNSZone.

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

<a id="canonical-1023101021203311-3023130231010130-1322330000132330-2023230311123100-2012123002112310-2031031201102312-3231010100321200-0102022130022031"></a>

<a id="canonical-2133330332332332-0020233232200331-3002100033331002-3100233011013312-3203211302033311-3133223221000123-0130012202102212-3102010232000223"></a>

## ID property — Property reference / 121200323121 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0110100020320101-0130021310202230-2013020332220002-3000101022211122-0311203211020322-3000321020210233-3303130032213211-1221300131331333"></a>

<a id="canonical-3022222102100213-1203012333332133-2303021332133230-3132003133202220-2301030012313322-1332032102200220-2332211003302020-1003230013133102"></a>

## labels property — Property reference / 121200323121 / 7

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

<a id="canonical-3201221312031130-2310120132001323-3331220033233321-3220123113311320-3130003023320121-1103211323231200-2002020013313233-0303233223320230"></a>

<a id="canonical-3232213022201221-0122102122210222-3013210222231331-2310102130203231-3113023232033220-2222322322113230-2202310322211110-3321312000301220"></a>

## name property — Property reference / 121200323121 / 8

Type: `"string"`. Required.

Name of the DNSZone.

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

<a id="canonical-2220103201320323-2303320122033130-3310232030330201-0313001220012332-0213013110301102-1022302003301323-0231101321331310-0011300030103113"></a>

<a id="canonical-2130013200210202-0102021300200323-2122331001101110-1233223301202011-0100013113202131-3323200032310213-3023000330110001-0131021323213301"></a>

## namespace property — Property reference / 121200323121 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSZone exists.

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

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331): complete subsection reference.

- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023): complete subsection reference.

<a id="canonical-3021312002033220-1203102111012331-0220230100131011-3301313312210100-2000313133313021-3230313030202233-1002230202311200-2030310102101332"></a>

## All schema paths — Property reference / 121200323121 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_zone--reference--group-001.md#canonical-3221313102233023-3020301321310200-0111300210103000-2202211230033103-0321320110120003-1003232312202033-3032022323011201-0321232133213221) |
| `description` | [description](data-sources--dns_zone--reference--group-001.md#canonical-3331313120113002-1202330323333013-1113000020020211-3000103210033330-3300213221310233-2130120021202010-1002233000203300-3322000333302332) |
| `id` | [ID](data-sources--dns_zone--reference--group-001.md#canonical-1023101021203311-3023130231010130-1322330000132330-2023230311123100-2012123002112310-2031031201102312-3231010100321200-0102022130022031) |
| `labels` | [labels](data-sources--dns_zone--reference--group-001.md#canonical-0110100020320101-0130021310202230-2013020332220002-3000101022211122-0311203211020322-3000321020210233-3303130032213211-1221300131331333) |
| `name` | [name](data-sources--dns_zone--reference--group-001.md#canonical-3201221312031130-2310120132001323-3331220033233321-3220123113311320-3130003023320121-1103211323231200-2002020013313233-0303233223320230) |
| `namespace` | [namespace](data-sources--dns_zone--reference--group-001.md#canonical-2220103201320323-2303320122033130-3310232030330201-0313001220012332-0213013110301102-1022302003301323-0231101321331310-0011300030103113) |
| `primary` | [primary](data-sources--dns_zone--reference--group-001.md#canonical-2121010121102223-3310323133310112-1113111110203223-3000201032032200-1001233231232110-0102013302313230-3223110102210011-1002120320220331) |
| `primary.allow_http_lb_managed_records` | [primary.allow_http_lb_managed_records](data-sources--dns_zone--reference--group-001.md#canonical-0011201311331332-0000001012122003-1122202131303312-2201102031231123-1020212100002012-3212210231021302-3313210321012000-2333130112103312) |
| `primary.default_rr_set_group` | [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-1113102033121113-1000301100233101-1323300223303133-1331120113002132-3232122222030111-1202002013022232-2332320011300313-0020230311202010) |
| `primary.default_rr_set_group.a_record` | [primary.default_rr_set_group.a_record](data-sources--dns_zone--reference--group-001.md#canonical-2202030000303320-2203033023002101-2121302113212320-1131112300223133-1122103223031130-0003001132133122-0032130303301233-0203212132221021) |
| `primary.default_rr_set_group.a_record.name` | [primary.default_rr_set_group.a_record.name](data-sources--dns_zone--reference--group-001.md#canonical-2213233021320010-3133010020313213-2313001313101301-2102002013301311-2202113133323321-1011203112301310-3320200211203133-3003333032001011) |
| `primary.default_rr_set_group.a_record.values` | [primary.default_rr_set_group.a_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1210311210102033-3132330103023322-1313233121311111-2222002211303113-1131020022311300-0032131201212002-3011111312133232-0102300230123221) |
| `primary.default_rr_set_group.aaaa_record` | [primary.default_rr_set_group.aaaa_record](data-sources--dns_zone--reference--group-001.md#canonical-0313033121023202-0312110313020213-1232110322312331-3133000312113310-1011221300011230-2320303010232220-2220220333322322-1320130222113131) |
| `primary.default_rr_set_group.aaaa_record.name` | [primary.default_rr_set_group.aaaa_record.name](data-sources--dns_zone--reference--group-001.md#canonical-3102033001011313-3312213133110031-1221232032231231-1231030232132110-3323101230131120-2032210102310203-0320231313323302-3022120111322120) |
| `primary.default_rr_set_group.aaaa_record.values` | [primary.default_rr_set_group.aaaa_record.values](data-sources--dns_zone--reference--group-001.md#canonical-2303002320003221-3110210003001112-2112001002022133-2002203131120320-2101302123032200-2312303000022032-3212323203320221-1211023011122323) |
| `primary.default_rr_set_group.afsdb_record` | [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-3132130120303123-3111331113133222-0003220322132231-1221120021333122-0133313012220113-2012021312102302-1112000213233011-1313022312013111) |
| `primary.default_rr_set_group.afsdb_record.name` | [primary.default_rr_set_group.afsdb_record.name](data-sources--dns_zone--reference--group-001.md#canonical-0110300331130123-2132312123101333-3313303003231130-3210110311213333-1222000333101001-2331120331301321-3210032333021002-0223203023231010) |
| `primary.default_rr_set_group.afsdb_record.values` | [primary.default_rr_set_group.afsdb_record.values](data-sources--dns_zone--reference--group-001.md#canonical-3130002333022131-0230003023021203-3030303302332332-2010201122323313-1311131232021021-3023030200203011-3122322002332311-3213020132110310) |
| `primary.default_rr_set_group.afsdb_record.values.hostname` | [primary.default_rr_set_group.afsdb_record.values.hostname](data-sources--dns_zone--reference--group-001.md#canonical-1002033312212101-2331120003311212-1112012311020212-3131132231320032-2131001021113100-3013333100122233-2332101002131012-1330211130110103) |
| `primary.default_rr_set_group.afsdb_record.values.subtype` | [primary.default_rr_set_group.afsdb_record.values.subtype](data-sources--dns_zone--reference--group-001.md#canonical-1321311233222302-0032022013222201-1222322121032202-3331031020302221-2230133312111312-2100330331332211-2030101201212121-2022100230311220) |
| `primary.default_rr_set_group.alias_record` | [primary.default_rr_set_group.alias_record](data-sources--dns_zone--reference--group-001.md#canonical-3123303123312231-1120201130031210-3000200201321020-1023312200012301-2022110200001130-0010132031321001-3231322013233110-2022311301033311) |
| `primary.default_rr_set_group.alias_record.value` | [primary.default_rr_set_group.alias_record.value](data-sources--dns_zone--reference--group-001.md#canonical-0003333310003211-0232300123200202-3310313300103201-2001033302312032-3012303310010222-3100122323030303-3202303122022202-1230220121111033) |
| `primary.default_rr_set_group.caa_record` | [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-2111230222101021-1320210000333020-1020303331213122-1001003201300100-0303330020020022-3231031130123220-1012230321300322-2302033111230121) |
| `primary.default_rr_set_group.caa_record.name` | [primary.default_rr_set_group.caa_record.name](data-sources--dns_zone--reference--group-001.md#canonical-3000003332023313-1123310023302312-3132100123333110-2330122120332011-2103222122330321-3211022322100131-0122033110321331-2012101200211130) |
| `primary.default_rr_set_group.caa_record.values` | [primary.default_rr_set_group.caa_record.values](data-sources--dns_zone--reference--group-001.md#canonical-0132232112330122-3020303022121233-2123111101133121-2302010323110230-0030000122103020-1120001103323123-3010320311321223-2011213300121003) |
| `primary.default_rr_set_group.caa_record.values.flags` | [primary.default_rr_set_group.caa_record.values.flags](data-sources--dns_zone--reference--group-001.md#canonical-2323311312222001-0013120300311322-3331332101313022-3010113010130212-3203101022000313-3332211000313010-1211133333223132-1200133111131111) |
| `primary.default_rr_set_group.caa_record.values.tag` | [primary.default_rr_set_group.caa_record.values.tag](data-sources--dns_zone--reference--group-001.md#canonical-3310212311323301-2103133002233110-0311220022301230-3012331211003302-1120010102201103-1213323032331123-1131130222211113-3130023313233130) |
| `primary.default_rr_set_group.caa_record.values.value` | [primary.default_rr_set_group.caa_record.values.value](data-sources--dns_zone--reference--group-001.md#canonical-3300221002010112-1222221230022020-0133301132310033-0322101220101113-1013000300312121-0202011010101102-1212100222220323-1301102033110132) |
| `primary.default_rr_set_group.cds_record` | [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-2210130232002212-1221333233322332-3031222320301131-2210031202320332-1303221321010000-2021332020120011-3123023213220111-0223010032133231) |
| `primary.default_rr_set_group.cds_record.name` | [primary.default_rr_set_group.cds_record.name](data-sources--dns_zone--reference--group-001.md#canonical-1201233312111231-1312223013101111-3000233222220130-0203330022131200-2111132003323200-2331110300033223-2031011031110210-2203031110220001) |
| `primary.default_rr_set_group.cds_record.values` | [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-3103311021310220-2021121122010131-1301303121113303-1101000131030330-2222221321211310-0132203020332033-0302233303312213-3113321102110313) |
| `primary.default_rr_set_group.cds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.cds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-001.md#canonical-3233213123322010-1211232312021001-1321021030302312-2130010311131123-1230122033211101-2032310133202201-3121332323122032-0100202220221311) |
| `primary.default_rr_set_group.cds_record.values.key_tag` | [primary.default_rr_set_group.cds_record.values.key_tag](data-sources--dns_zone--reference--group-001.md#canonical-1013003310003321-0233033113131202-2030122110213212-2032012121113032-2220021001032001-1302233023132021-0023230203033212-0120111303021020) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-2103023032101032-1210032133333022-0203320020311231-2220032330311202-3002111302123031-0122002213123323-0102101321233310-3210232031201230) |
| `primary.default_rr_set_group.cds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-001.md#canonical-1120131122112210-0130233133221023-1202331310003101-3200323303211222-2232202233101230-1012013122122111-2112101103333222-2001111220211112) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-1032321100131320-0221111023312313-0232120133303312-0230001023123021-2001012000313121-1313233003020322-1112222101013232-3232100131020112) |
| `primary.default_rr_set_group.cds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-001.md#canonical-0330110211220103-0112233012302321-0320000301011010-0301013311202131-3102130303001111-0000023002230313-1032013222121111-3201032233202022) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-001.md#canonical-1111012111012013-0202302223320200-0201300033122233-0333233300102011-0023220200002302-0222201302312131-3001131323213233-3010033312010232) |
| `primary.default_rr_set_group.cds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.cds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-001.md#canonical-0302321003303001-0122012011100300-1121323322030303-1301010313100320-2101220130023133-3132221310321100-2011333230130233-2202031120302103) |
| `primary.default_rr_set_group.cert_record` | [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-2330230012123132-1221230312310033-2133111133112033-3131023230131220-0231210301103033-2211311331101331-3332221321202331-3003110220113301) |
| `primary.default_rr_set_group.cert_record.name` | [primary.default_rr_set_group.cert_record.name](data-sources--dns_zone--reference--group-001.md#canonical-3312220333221031-3220130201031212-2002110032020233-1201230221201312-2000101303322310-3311330031230311-0111131011310213-3312210201113022) |
| `primary.default_rr_set_group.cert_record.values` | [primary.default_rr_set_group.cert_record.values](data-sources--dns_zone--reference--group-001.md#canonical-2300123012220120-2033212302211023-0112223130101001-0313202113112320-0123022011111033-0311122233330000-1111031311102232-3200212322130333) |
| `primary.default_rr_set_group.cert_record.values.algorithm` | [primary.default_rr_set_group.cert_record.values.algorithm](data-sources--dns_zone--reference--group-001.md#canonical-2110103123313222-3133011311023130-1330323122011122-0100032000232133-2122211311121031-1231120200000132-3031230301321010-2323313000212112) |
| `primary.default_rr_set_group.cert_record.values.cert_key_tag` | [primary.default_rr_set_group.cert_record.values.cert_key_tag](data-sources--dns_zone--reference--group-001.md#canonical-2200300230131231-1333023020230233-0133020022321030-3033103233010112-1000022211020301-0101212100002200-1010100100013130-1100013000013031) |
| `primary.default_rr_set_group.cert_record.values.cert_type` | [primary.default_rr_set_group.cert_record.values.cert_type](data-sources--dns_zone--reference--group-001.md#canonical-1002200101302031-2132212233210101-1311100333232013-2333200133211232-3223303003302033-2303111213110322-2122330221223223-1222220212311011) |
| `primary.default_rr_set_group.cert_record.values.certificate` | [primary.default_rr_set_group.cert_record.values.certificate](data-sources--dns_zone--reference--group-001.md#canonical-3122321133333333-1313311023321310-2110012323101232-1320010313330122-3203211011023331-1101001120332021-0110112311122320-3311310102211311) |
| `primary.default_rr_set_group.cname_record` | [primary.default_rr_set_group.cname_record](data-sources--dns_zone--reference--group-001.md#canonical-2121301111310012-2001113013020011-2301022103002203-0100021133213031-2030220010101133-1201321223200231-2001321103013223-2031131302020121) |
| `primary.default_rr_set_group.cname_record.name` | [primary.default_rr_set_group.cname_record.name](data-sources--dns_zone--reference--group-001.md#canonical-3020113003202222-3332121000120002-0233003201323021-3031130003001032-0313022212232211-1133110132212202-0123303011210313-3131311203223010) |
| `primary.default_rr_set_group.cname_record.value` | [primary.default_rr_set_group.cname_record.value](data-sources--dns_zone--reference--group-001.md#canonical-1310203013003032-2033100230131202-1013032123032013-0223122123333120-1303232300121302-1111013312202312-3021302322110300-3212003110312100) |
| `primary.default_rr_set_group.description_spec` | [primary.default_rr_set_group.description_spec](data-sources--dns_zone--reference--group-001.md#canonical-0113312013210200-2113030131021323-3132103200313313-1213300120001232-0113212330132300-1123320311201132-3000210010001030-1130233112210000) |
| `primary.default_rr_set_group.ds_record` | [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-1211203130131120-0121001321313311-2231310210220030-0130122123010221-3022333331222000-0331300031001001-2220030322321320-0320011122133131) |
| `primary.default_rr_set_group.ds_record.name` | [primary.default_rr_set_group.ds_record.name](data-sources--dns_zone--reference--group-001.md#canonical-2210113100211302-3023013321210022-3333000201102313-0132020301300332-3213330302032230-0032313022031123-2333123021331230-0220332230231321) |
| `primary.default_rr_set_group.ds_record.values` | [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-2000102011223330-2213001100022300-1100221211120111-0313312230222303-0203031212312221-0010011322110212-2131011000032101-3103301212212110) |
| `primary.default_rr_set_group.ds_record.values.ds_key_algorithm` | [primary.default_rr_set_group.ds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-001.md#canonical-0122212333010013-2111301021003303-2103230323033300-3333323220330332-2121323332011232-0311313230312000-2032212021031322-0212212220312213) |
| `primary.default_rr_set_group.ds_record.values.key_tag` | [primary.default_rr_set_group.ds_record.values.key_tag](data-sources--dns_zone--reference--group-001.md#canonical-2302203310103211-2223100220132232-1032020122123322-3311133002203121-0312003313110320-2223111030000202-1303300013223313-0230131332222113) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-1223112030310123-0221210302013131-1302031020223003-0232102210013120-0312301030233230-3011132002111230-2011112321221331-2031230100020003) |
| `primary.default_rr_set_group.ds_record.values.sha1_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-2023121103322110-3103132303130122-2323231213010203-0220231130100102-3101121223031323-3230022000013212-3230211311003023-0133203133102013) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-1211103222131221-0032322120303010-2101302001321032-3231020200110231-3301212111210120-1102020121203303-3220031031021223-3022213213001031) |
| `primary.default_rr_set_group.ds_record.values.sha256_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-1031223323002303-2111302203323302-0012223231121122-1012003012203033-1120321223130001-0321321302003312-2210122333202000-2000010001121130) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-1203222223031333-2213313232123133-0132023212213320-2122210011111231-3131223111200131-1123001213031032-3110211201230233-3010301332013201) |
| `primary.default_rr_set_group.ds_record.values.sha384_digest.digest` | [primary.default_rr_set_group.ds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-2100200132011100-3021232201132311-0110031023111230-0230311130112321-1022312111330101-1000311330010302-0011213030023103-0120313223313312) |
| `primary.default_rr_set_group.eui48_record` | [primary.default_rr_set_group.eui48_record](data-sources--dns_zone--reference--group-002.md#canonical-1101003211112233-0020221210212232-2301212122120221-2300232221312301-1030131112022130-3022331321302121-3330012020221302-1003130221133220) |
| `primary.default_rr_set_group.eui48_record.name` | [primary.default_rr_set_group.eui48_record.name](data-sources--dns_zone--reference--group-002.md#canonical-3013102300200103-0202323022102033-1200330200122232-1210000233003231-3231021101211002-1232001131313323-2031200010211030-2213112022200230) |
| `primary.default_rr_set_group.eui48_record.value` | [primary.default_rr_set_group.eui48_record.value](data-sources--dns_zone--reference--group-002.md#canonical-0001211302020011-2101312333310022-1210110221303312-1000202000302121-2130300202121322-3310030321033312-2100100100120331-3101332333002001) |
| `primary.default_rr_set_group.eui64_record` | [primary.default_rr_set_group.eui64_record](data-sources--dns_zone--reference--group-002.md#canonical-3000330130313332-3332310121201212-2130221001122010-1300301102003012-1013322223111132-2331120122013201-2312123333332202-1012112130103120) |
| `primary.default_rr_set_group.eui64_record.name` | [primary.default_rr_set_group.eui64_record.name](data-sources--dns_zone--reference--group-002.md#canonical-1110310131032031-2121201121311221-3220103312311032-2023310333211031-3221112200122201-2212302330023321-1332231122301120-0312310130220002) |
| `primary.default_rr_set_group.eui64_record.value` | [primary.default_rr_set_group.eui64_record.value](data-sources--dns_zone--reference--group-002.md#canonical-2222210322121023-2223222212201303-3012320021301102-2131013223033011-3131023001200320-1120233032223220-1021303201121302-3223121031211301) |
| `primary.default_rr_set_group.lb_record` | [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-3321013212110302-0023102310032001-3322010023033222-2230130001031320-3001233100201112-1020110302111000-1220132112322010-3203032000220202) |
| `primary.default_rr_set_group.lb_record.name` | [primary.default_rr_set_group.lb_record.name](data-sources--dns_zone--reference--group-002.md#canonical-0213222223323021-1232220023221221-2113002203131303-3031102030213313-0231132333033122-2200210100232031-2003322233303023-2010000022313133) |
| `primary.default_rr_set_group.lb_record.value` | [primary.default_rr_set_group.lb_record.value](data-sources--dns_zone--reference--group-002.md#canonical-3132203013302103-1312022333333202-3130320101313310-3303130301333122-2113200331320023-0321031213112232-0322210330021013-0121332320212031) |
| `primary.default_rr_set_group.lb_record.value.name` | [primary.default_rr_set_group.lb_record.value.name](data-sources--dns_zone--reference--group-002.md#canonical-1131222303231032-1113322200021231-3133023102131231-1200300223133320-3311022322023333-2020231221020003-0011023000010012-3221322012121310) |
| `primary.default_rr_set_group.lb_record.value.namespace` | [primary.default_rr_set_group.lb_record.value.namespace](data-sources--dns_zone--reference--group-002.md#canonical-1212022211001321-1200111110200222-2303103333002230-1013031323000102-0122222231132301-1333211130211121-0131132321223001-0101122332120212) |
| `primary.default_rr_set_group.lb_record.value.tenant` | [primary.default_rr_set_group.lb_record.value.tenant](data-sources--dns_zone--reference--group-002.md#canonical-3230102333320033-2110000030300102-0300320022222021-3012013123221123-2300002232010302-1112122220233213-0102120012112220-0321103120121233) |
| `primary.default_rr_set_group.loc_record` | [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-3302122211311222-3103311102023100-3220103323100201-2123010111112323-3123113301023332-0230002213201013-0330231020002311-2321300011222312) |
| `primary.default_rr_set_group.loc_record.name` | [primary.default_rr_set_group.loc_record.name](data-sources--dns_zone--reference--group-002.md#canonical-1103230001020130-3331023302200223-0112220003002223-1000103012121303-0222220111130123-2020333213331303-0120213102211032-0313311132021233) |
| `primary.default_rr_set_group.loc_record.values` | [primary.default_rr_set_group.loc_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3113200210231122-0221222203132301-0212030311303310-0330133133130233-2203302103302001-2132123322012331-2103102010130012-1013320022212120) |
| `primary.default_rr_set_group.loc_record.values.altitude` | [primary.default_rr_set_group.loc_record.values.altitude](data-sources--dns_zone--reference--group-002.md#canonical-3311200200213133-2021210303021311-2133231022013130-0331120213113330-2022203010312111-1301032302320021-1313021230230100-1121013032231210) |
| `primary.default_rr_set_group.loc_record.values.horizontal_precision` | [primary.default_rr_set_group.loc_record.values.horizontal_precision](data-sources--dns_zone--reference--group-002.md#canonical-2300120201211002-2310321321030112-2130230300003332-0231021010202202-3122130320132231-3000013013122001-0222001100232331-3001110311320213) |
| `primary.default_rr_set_group.loc_record.values.latitude_degree` | [primary.default_rr_set_group.loc_record.values.latitude_degree](data-sources--dns_zone--reference--group-002.md#canonical-3110303113003102-1031302331030321-0323110313132003-2131313102230020-0221013012300111-3111333123101320-0010001012103212-0230030110121310) |
| `primary.default_rr_set_group.loc_record.values.latitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.latitude_hemisphere](data-sources--dns_zone--reference--group-002.md#canonical-0031230002231000-1210022311012023-0100113101103331-1101303222012222-3232112130033001-0320130022210312-1103201301000220-0102333100100110) |
| `primary.default_rr_set_group.loc_record.values.latitude_minute` | [primary.default_rr_set_group.loc_record.values.latitude_minute](data-sources--dns_zone--reference--group-002.md#canonical-2302110130023131-3121330000331301-1003332213113101-3130231322033301-0111122301311023-1201001100203033-1023313100311211-3131233302311013) |
| `primary.default_rr_set_group.loc_record.values.latitude_second` | [primary.default_rr_set_group.loc_record.values.latitude_second](data-sources--dns_zone--reference--group-002.md#canonical-1311022203303033-3112312100332301-3310013101031230-3202311111222113-0202330221113131-3320312123311020-2233312002213221-3100303212010301) |
| `primary.default_rr_set_group.loc_record.values.location_diameter` | [primary.default_rr_set_group.loc_record.values.location_diameter](data-sources--dns_zone--reference--group-002.md#canonical-1332020110030220-0110323332033122-0211221103213133-2132023210320332-3200130321302230-3113133201030012-0201213332203223-1203223113202330) |
| `primary.default_rr_set_group.loc_record.values.longitude_degree` | [primary.default_rr_set_group.loc_record.values.longitude_degree](data-sources--dns_zone--reference--group-002.md#canonical-3120232123033331-0300133111133213-3111211323203213-3310210223031100-1131111302113223-0211120310003212-3032312321123012-0002301303031332) |
| `primary.default_rr_set_group.loc_record.values.longitude_hemisphere` | [primary.default_rr_set_group.loc_record.values.longitude_hemisphere](data-sources--dns_zone--reference--group-002.md#canonical-2100132230332122-0112003022133101-1111223133113332-1113333203031030-0011231300022013-1210110123230003-3321212000222003-3220101322211323) |
| `primary.default_rr_set_group.loc_record.values.longitude_minute` | [primary.default_rr_set_group.loc_record.values.longitude_minute](data-sources--dns_zone--reference--group-002.md#canonical-2023123133322101-3220201332030302-2223300002022200-1211012232222333-3213230220333023-0101130013312211-3132312322023222-1133300303332100) |
| `primary.default_rr_set_group.loc_record.values.longitude_second` | [primary.default_rr_set_group.loc_record.values.longitude_second](data-sources--dns_zone--reference--group-002.md#canonical-0130330210323301-2302302203032002-3013230121123223-2012002223231113-3003312111101111-1133012123001123-1001222022213200-3233112113322123) |
| `primary.default_rr_set_group.loc_record.values.vertical_precision` | [primary.default_rr_set_group.loc_record.values.vertical_precision](data-sources--dns_zone--reference--group-002.md#canonical-2111033231212233-2011201101302003-2221003230122121-1013131311032202-0001230011101203-3113101020101303-1213112220312030-3331012100003310) |
| `primary.default_rr_set_group.mx_record` | [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-1032311100202302-3103323103112013-1333010002323213-2002133011303012-0230220002213110-1003213032321220-1130322103013231-3303203111301121) |
| `primary.default_rr_set_group.mx_record.name` | [primary.default_rr_set_group.mx_record.name](data-sources--dns_zone--reference--group-002.md#canonical-0333112301302200-3132212312222130-0022323331310010-2303001231202322-0313332131233111-2202322212320001-0012222330120330-2020010020321221) |
| `primary.default_rr_set_group.mx_record.values` | [primary.default_rr_set_group.mx_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3233232323213020-1102033211231011-1321202112322120-0333320230101121-0320302103020112-1331300012031212-0022010320131022-3001202130130233) |
| `primary.default_rr_set_group.mx_record.values.domain` | [primary.default_rr_set_group.mx_record.values.domain](data-sources--dns_zone--reference--group-002.md#canonical-3212233103010021-2112330330101021-3311113123112320-3210112120133231-1123310220120112-1301333133002022-2113311003233020-1203122111002323) |
| `primary.default_rr_set_group.mx_record.values.priority` | [primary.default_rr_set_group.mx_record.values.priority](data-sources--dns_zone--reference--group-002.md#canonical-3322212023020301-2102010023101223-1120101300130022-3000311103023230-3220132000213212-3210333213332211-3112113323201011-3332223233032202) |
| `primary.default_rr_set_group.naptr_record` | [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-1131003222002322-0331310113232231-0120321102333300-2000203102120231-0302103021322220-3232023223211033-2330023121211213-2302003331022033) |
| `primary.default_rr_set_group.naptr_record.name` | [primary.default_rr_set_group.naptr_record.name](data-sources--dns_zone--reference--group-002.md#canonical-1310020030321333-0101103311331302-2303012331310110-3003213333031333-1013020301321101-1322312211323103-3013310130323002-3013132232123002) |
| `primary.default_rr_set_group.naptr_record.values` | [primary.default_rr_set_group.naptr_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2320200022012033-1033033032110021-0102202303021210-3030031212130021-0300133300102312-2203100103330133-2300012122032212-1100023312311020) |
| `primary.default_rr_set_group.naptr_record.values.flags` | [primary.default_rr_set_group.naptr_record.values.flags](data-sources--dns_zone--reference--group-002.md#canonical-3220320022000021-1233032000023232-0021330031311101-1003210020112310-0111322332300202-2330101220201211-0131110023033311-2302310101333232) |
| `primary.default_rr_set_group.naptr_record.values.order` | [primary.default_rr_set_group.naptr_record.values.order](data-sources--dns_zone--reference--group-002.md#canonical-0300031121323202-3023112232120133-0130222230303332-0302000303203010-0010113133012022-2012120103012321-2323100221100333-2202023212021101) |
| `primary.default_rr_set_group.naptr_record.values.preference` | [primary.default_rr_set_group.naptr_record.values.preference](data-sources--dns_zone--reference--group-002.md#canonical-2011021203322230-2231013332301131-3132101230023310-3023000312210330-1333030302112031-2210123203222133-0111202001321110-2302302111210113) |
| `primary.default_rr_set_group.naptr_record.values.regexp` | [primary.default_rr_set_group.naptr_record.values.regexp](data-sources--dns_zone--reference--group-002.md#canonical-3330121203020203-0221132100100003-2002333210221223-1311220321133311-3323323311333033-2122031012210132-1223102033211211-3032111300330202) |
| `primary.default_rr_set_group.naptr_record.values.replacement` | [primary.default_rr_set_group.naptr_record.values.replacement](data-sources--dns_zone--reference--group-002.md#canonical-3033113003322110-2210300231010000-0003313223021212-0212010220101222-1000113022332113-2210013200131230-1312001131132203-0331103032022122) |
| `primary.default_rr_set_group.naptr_record.values.service` | [primary.default_rr_set_group.naptr_record.values.service](data-sources--dns_zone--reference--group-002.md#canonical-1103001310312002-2033033120313313-3301110213102322-1200103113111121-0111201223100233-1210011003112100-3331010212033110-2111113103012012) |
| `primary.default_rr_set_group.ns_record` | [primary.default_rr_set_group.ns_record](data-sources--dns_zone--reference--group-002.md#canonical-3120011311313020-0111201012211122-2213331000100130-1232101330013230-0030203001332200-2000311000020320-3311002011130312-1322011010013123) |
| `primary.default_rr_set_group.ns_record.name` | [primary.default_rr_set_group.ns_record.name](data-sources--dns_zone--reference--group-002.md#canonical-2213320012022313-1021103303322033-0111031323202112-3312112010303203-3011210120120320-2230203102200111-2211022130202111-3010320210120211) |
| `primary.default_rr_set_group.ns_record.values` | [primary.default_rr_set_group.ns_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0300120203320102-3303321332223332-0323100011233013-3230312112133013-3101333233332231-1033102102210322-3132203032100032-3220031211032023) |
| `primary.default_rr_set_group.ptr_record` | [primary.default_rr_set_group.ptr_record](data-sources--dns_zone--reference--group-002.md#canonical-3303011132132001-3122103213121003-3311211220312210-1111031013000031-3212320011003331-3323012020013020-0120110300202132-1012322330233103) |
| `primary.default_rr_set_group.ptr_record.name` | [primary.default_rr_set_group.ptr_record.name](data-sources--dns_zone--reference--group-002.md#canonical-1023002231030203-1111200200202010-1320211303302021-0033222220111300-1011203222322021-1122221223213310-0210233000330313-0013022333303210) |
| `primary.default_rr_set_group.ptr_record.values` | [primary.default_rr_set_group.ptr_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2313101211020213-0201123002020023-1113010113221221-2301313211102232-3200202311133321-3332021221031222-0321110113201202-0200112033033110) |
| `primary.default_rr_set_group.srv_record` | [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-2001102313121111-1031012001223021-1320032301003300-1221012201112033-0210231131333011-1223133113110100-3122001233211300-1233200332020001) |
| `primary.default_rr_set_group.srv_record.name` | [primary.default_rr_set_group.srv_record.name](data-sources--dns_zone--reference--group-002.md#canonical-0323130100321031-0322103030031330-0332231023332230-2002133130023113-3111213100333233-3133102030110313-3120321101233323-1223230200201300) |
| `primary.default_rr_set_group.srv_record.values` | [primary.default_rr_set_group.srv_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3011201211313333-0132121313130231-1223101003211111-3213221203211211-2320300231031112-1033212012000030-3311101111020022-3113122301120032) |
| `primary.default_rr_set_group.srv_record.values.port` | [primary.default_rr_set_group.srv_record.values.port](data-sources--dns_zone--reference--group-002.md#canonical-0001332202013201-2121012121222022-1020110233101232-1311321312220111-0023022022232021-0130000111100131-3221031321133220-0211320210103302) |
| `primary.default_rr_set_group.srv_record.values.priority` | [primary.default_rr_set_group.srv_record.values.priority](data-sources--dns_zone--reference--group-002.md#canonical-0203121331131200-1221202321111221-2331131303003330-3232111021110213-3010231102103002-0003330311113320-1330012010132020-2320212112122011) |
| `primary.default_rr_set_group.srv_record.values.target` | [primary.default_rr_set_group.srv_record.values.target](data-sources--dns_zone--reference--group-002.md#canonical-0130212300322010-0021212210213323-3301110210322211-2131222211120020-2303212202032322-0322020110111220-2311332113310212-1303202022020203) |
| `primary.default_rr_set_group.srv_record.values.weight` | [primary.default_rr_set_group.srv_record.values.weight](data-sources--dns_zone--reference--group-002.md#canonical-1231123310210022-3110320103030330-1033213211212230-1102120031210331-2123020301223321-2202023200133222-2023030101310323-3302223332123313) |
| `primary.default_rr_set_group.sshfp_record` | [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2013012331332022-3033213302020033-1223002311330010-0102020002302130-2100313302123103-2003323030131310-2232212330323011-3212033020230212) |
| `primary.default_rr_set_group.sshfp_record.name` | [primary.default_rr_set_group.sshfp_record.name](data-sources--dns_zone--reference--group-002.md#canonical-1023122132223330-2001321301031200-1030312232001233-0201023130202100-2122322002301000-3102310331020333-3320013132232223-3030021323120222) |
| `primary.default_rr_set_group.sshfp_record.values` | [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0002031333331122-1032020130110201-3102202233202031-1002322002200203-2030133203023222-3020310021213033-2031110213101210-2130013011222231) |
| `primary.default_rr_set_group.sshfp_record.values.algorithm` | [primary.default_rr_set_group.sshfp_record.values.algorithm](data-sources--dns_zone--reference--group-002.md#canonical-0303030022003011-3020231322301030-0330110112122301-0330022000303231-2203332223112112-3201230320313201-1023202102110220-1001333003113030) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-2301022221023201-1013221230312100-0313330033313231-3210020032122202-2321203113211310-2021331300331312-0010123010032000-2212302103211002) |
| `primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint.fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-2202012233001101-3131012212320120-2212030130212331-2331310220220331-0001230232101201-1221122111131023-0221201300130200-3003311132232222) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-1201210333203213-2223231311222211-0201320320001023-3110022121030130-2302312101110210-0301333010330033-0230022230112212-2210011303031321) |
| `primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint.fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-1113020303311333-0132332222011212-0232212100333233-3110110303222112-2023213310022210-2300321230330113-2222032010303302-0030102310221231) |
| `primary.default_rr_set_group.tlsa_record` | [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-0032202110310221-3032213010113022-0333121311132030-3203010222233313-0303021332131032-3222131113020221-0011020030321203-0301021200311302) |
| `primary.default_rr_set_group.tlsa_record.name` | [primary.default_rr_set_group.tlsa_record.name](data-sources--dns_zone--reference--group-002.md#canonical-1230213033331222-0330232100013210-3213221211033210-3211223220332020-1101010120033211-0233030231000310-0212212300203223-2210211111022110) |
| `primary.default_rr_set_group.tlsa_record.values` | [primary.default_rr_set_group.tlsa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3002001033031230-0223021112020031-2031230013011312-2212321010023312-2130233001011032-3333311220331311-2220322001301203-3031032102121020) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_association_data` | [primary.default_rr_set_group.tlsa_record.values.certificate_association_data](data-sources--dns_zone--reference--group-002.md#canonical-0101102321003212-3113212202302201-2110303021101311-3303012000332311-0311220231312102-0132122221022212-0211013111023301-1111100322111200) |
| `primary.default_rr_set_group.tlsa_record.values.certificate_usage` | [primary.default_rr_set_group.tlsa_record.values.certificate_usage](data-sources--dns_zone--reference--group-002.md#canonical-0123210230200102-3213021213222203-2002331032101201-1123001303221233-3323100202113322-3130212113311100-3033330032301210-1231012020213333) |
| `primary.default_rr_set_group.tlsa_record.values.matching_type` | [primary.default_rr_set_group.tlsa_record.values.matching_type](data-sources--dns_zone--reference--group-002.md#canonical-1123030022332000-1322302232022012-3221001233221210-3233322123131031-2113030131111212-1200000322300232-2221233112232322-1131020233000320) |
| `primary.default_rr_set_group.tlsa_record.values.selector` | [primary.default_rr_set_group.tlsa_record.values.selector](data-sources--dns_zone--reference--group-002.md#canonical-1212311331323330-3121101311123002-3322032023130310-1333000010211313-0211023201030232-2111022020011002-3320313123203303-1302133303211202) |
| `primary.default_rr_set_group.ttl` | [primary.default_rr_set_group.ttl](data-sources--dns_zone--reference--group-001.md#canonical-1120233323223031-2202211123332021-0131113311112233-3133300000310331-2210010332013333-0012130212232203-3223202223302113-2300001101113012) |
| `primary.default_rr_set_group.txt_record` | [primary.default_rr_set_group.txt_record](data-sources--dns_zone--reference--group-002.md#canonical-0113213300213103-1012321032121210-0320323211322311-2113133002212131-2200322100201330-2001332233103233-3303333311332233-0121112232302110) |
| `primary.default_rr_set_group.txt_record.name` | [primary.default_rr_set_group.txt_record.name](data-sources--dns_zone--reference--group-002.md#canonical-3132322012012100-2213013323323121-2113100213332123-3223330323222110-1300133212023222-0012333100203032-3211111032012332-2333213122021123) |
| `primary.default_rr_set_group.txt_record.values` | [primary.default_rr_set_group.txt_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0232301130321032-1123011101323311-1103103322303022-1212000212321333-0213003222303132-3211223133203202-1111111310212333-3023300210011323) |
| `primary.default_soa_parameters` | [primary.default_soa_parameters](data-sources--dns_zone--reference--group-002.md#canonical-1200102321322311-3223202022210202-3010323322310100-2013003011231200-1032331233311132-0021012320100102-3113332223310121-0101330133331122) |
| `primary.dnssec_mode` | [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-1003323000301330-1322020032333013-1121022131303132-2311101012002202-0130133003222100-0133113113112003-2121213120121032-2021301010121112) |
| `primary.dnssec_mode.disable_spec` | [primary.dnssec_mode.disable_spec](data-sources--dns_zone--reference--group-002.md#canonical-0133101022110312-1012230221200332-0112223131330022-3122233333232332-0303220302221003-1230223103002022-3112210203012202-3333110013323211) |
| `primary.dnssec_mode.enable` | [primary.dnssec_mode.enable](data-sources--dns_zone--reference--group-002.md#canonical-1021202202023231-1203202332111212-0003030323332333-2100202023101323-3223302330203000-3231022021132313-1010101012320311-0033321331223221) |
| `primary.rr_set_group` | [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-2231202303211010-2021101001111030-2013023101223233-3321213130020213-3332320300011023-2121211132011013-2312031132210200-1131031002133031) |
| `primary.rr_set_group.metadata` | [primary.rr_set_group.metadata](data-sources--dns_zone--reference--group-002.md#canonical-1330003303230112-2023000311003232-0313021100232031-2130100211300330-1322102200310132-0111211310022101-3331032000001001-1003333223010012) |
| `primary.rr_set_group.metadata.description_spec` | [primary.rr_set_group.metadata.description_spec](data-sources--dns_zone--reference--group-002.md#canonical-1322330123203331-2320210002222202-3221301311020012-2330303222321212-1031011322100311-2110232000022102-0032000113233121-0303020233122311) |
| `primary.rr_set_group.metadata.name` | [primary.rr_set_group.metadata.name](data-sources--dns_zone--reference--group-002.md#canonical-2202023131100022-1103302232221201-3011011233003132-1222131313102331-0013123113203032-1301132213112102-3002233120002323-1130303100213221) |
| `primary.rr_set_group.rr_set` | [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3313212012322332-1132311103230202-1103220203333130-3310221013011303-1033220103302101-1021000000330033-1210313021213300-0031330331100211) |
| `primary.rr_set_group.rr_set.a_record` | [primary.rr_set_group.rr_set.a_record](data-sources--dns_zone--reference--group-002.md#canonical-3011212320111000-3202333013120221-0021212111332023-2032123232032111-0201112310120223-3310200123333112-0212223220022001-2001113312010031) |
| `primary.rr_set_group.rr_set.a_record.name` | [primary.rr_set_group.rr_set.a_record.name](data-sources--dns_zone--reference--group-002.md#canonical-3213102231313033-0303331322121002-3300312033112130-0201111312112012-1003032023110211-2223300333323120-1211322302312013-2021322230330230) |
| `primary.rr_set_group.rr_set.a_record.values` | [primary.rr_set_group.rr_set.a_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0112323113010023-3023131013113031-3102131201322031-2233113002200212-3130333332313022-3002220000221331-1010312303011011-1023021030113300) |
| `primary.rr_set_group.rr_set.aaaa_record` | [primary.rr_set_group.rr_set.aaaa_record](data-sources--dns_zone--reference--group-002.md#canonical-2023331012311212-2202230303010330-2220121231320033-0011233220203122-1222030112301320-1131232001323310-2311111111133130-0132120210120321) |
| `primary.rr_set_group.rr_set.aaaa_record.name` | [primary.rr_set_group.rr_set.aaaa_record.name](data-sources--dns_zone--reference--group-002.md#canonical-0132130000323000-0120210230321113-0332300303303101-2131331032001110-0103203311103130-3131232113100322-3103033011222301-3203031100330310) |
| `primary.rr_set_group.rr_set.aaaa_record.values` | [primary.rr_set_group.rr_set.aaaa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-1201210031130232-1121021132332032-0100200001133112-2013220012213033-0321033333113332-2002303321201102-3103030013230103-3032011201322021) |
| `primary.rr_set_group.rr_set.afsdb_record` | [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-3033302222102131-3033210012122221-1223220200233121-2000203301300123-2311103123330212-0213212012130122-0230201320011032-0231132102321330) |
| `primary.rr_set_group.rr_set.afsdb_record.name` | [primary.rr_set_group.rr_set.afsdb_record.name](data-sources--dns_zone--reference--group-002.md#canonical-0032301233120202-0033231113100112-3111330023301212-2200010130112022-0331221002213211-0033300110200210-2322302222122332-1300113202211222) |
| `primary.rr_set_group.rr_set.afsdb_record.values` | [primary.rr_set_group.rr_set.afsdb_record.values](data-sources--dns_zone--reference--group-002.md#canonical-1023001200301202-0110001232003000-2021300103333001-0123111002000032-1320201020302133-1233211310230010-2233331130101130-1032330312322202) |
| `primary.rr_set_group.rr_set.afsdb_record.values.hostname` | [primary.rr_set_group.rr_set.afsdb_record.values.hostname](data-sources--dns_zone--reference--group-002.md#canonical-1110231102320131-3211323223302120-0301301330103020-3102030311203001-1310333221022233-2201120001221211-1031312223130302-2000200323131022) |
| `primary.rr_set_group.rr_set.afsdb_record.values.subtype` | [primary.rr_set_group.rr_set.afsdb_record.values.subtype](data-sources--dns_zone--reference--group-002.md#canonical-3310302212033311-1100003023013330-3132210322002310-0121211123201130-2020120003011121-1210013320110013-2231203210201011-3323222313231230) |
| `primary.rr_set_group.rr_set.alias_record` | [primary.rr_set_group.rr_set.alias_record](data-sources--dns_zone--reference--group-002.md#canonical-3330101333113112-1201010022313322-3013003311010121-2021213321311301-1320223332310312-3313023020020102-0333222132313023-1313022000331133) |
| `primary.rr_set_group.rr_set.alias_record.value` | [primary.rr_set_group.rr_set.alias_record.value](data-sources--dns_zone--reference--group-002.md#canonical-0121102012223122-3013201022220312-1033133222301232-0231202321131130-3233312331331022-3330023120303030-0332131132301222-3301101221303330) |
| `primary.rr_set_group.rr_set.caa_record` | [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-3330112101203321-2003331021202232-1010102323102020-3030102201031133-2032203333030330-1012301022310203-3102020010230303-2121311003313311) |
| `primary.rr_set_group.rr_set.caa_record.name` | [primary.rr_set_group.rr_set.caa_record.name](data-sources--dns_zone--reference--group-002.md#canonical-3203023010220211-2130302031120031-0012331023301333-3310303222120232-2113103010133130-2311001321201020-0132023102310033-1033300320133310) |
| `primary.rr_set_group.rr_set.caa_record.values` | [primary.rr_set_group.rr_set.caa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2233002323122031-0323113231103233-3332300130230020-2301012111103101-0202213230312031-0022032223213103-2220221031331222-2222213023101033) |
| `primary.rr_set_group.rr_set.caa_record.values.flags` | [primary.rr_set_group.rr_set.caa_record.values.flags](data-sources--dns_zone--reference--group-002.md#canonical-2211321133212110-1031313232033132-3103331013220112-0303210120233121-0301021131013032-3201001113121002-3021012212122312-1232103220103322) |
| `primary.rr_set_group.rr_set.caa_record.values.tag` | [primary.rr_set_group.rr_set.caa_record.values.tag](data-sources--dns_zone--reference--group-002.md#canonical-1101210323223113-1331223213132323-1322200301010110-2000012101132313-2122322123301211-2023001023323010-2332013021031110-3221032200302220) |
| `primary.rr_set_group.rr_set.caa_record.values.value` | [primary.rr_set_group.rr_set.caa_record.values.value](data-sources--dns_zone--reference--group-002.md#canonical-0230313221211321-3323301030103232-1121102112023002-3201100201022111-2333133220013211-3120020021013301-2311131201123333-3220101303303333) |
| `primary.rr_set_group.rr_set.cds_record` | [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3112102222133222-3003220331112000-1320021313102202-0011300212321131-2113010310310211-3210200013001230-0010310233323110-3012303012123200) |
| `primary.rr_set_group.rr_set.cds_record.name` | [primary.rr_set_group.rr_set.cds_record.name](data-sources--dns_zone--reference--group-002.md#canonical-3302113021000030-3032201310101020-1212330002102131-3202003013323323-3123322320212003-0232103121011002-2203201100131123-1031000300221300) |
| `primary.rr_set_group.rr_set.cds_record.values` | [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0323130310222111-2002023011012232-0210113302102211-1233200322033330-0113111301001120-2203020131211220-1322030002033312-3132100113001313) |
| `primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.cds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-002.md#canonical-1311020133210033-0320312100301033-0200030313011222-2333111231113311-1311030322303311-1101313022202023-2111222020331103-2030022111321310) |
| `primary.rr_set_group.rr_set.cds_record.values.key_tag` | [primary.rr_set_group.rr_set.cds_record.values.key_tag](data-sources--dns_zone--reference--group-002.md#canonical-3113103231121323-3231030031230133-1300200122020033-3330323030313330-3011112122111203-2003302222012021-2322112322131103-1301213223313032) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-2321022202123020-0222121220103123-2222131300102200-0103002133131132-0210100012112032-1102002023303223-1012030202000002-1221002330103130) |
| `primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-2133322000331020-0030100322011210-0300332323232303-0333110303131021-1210111223031103-0302010333221131-0213212032011220-3201232230330122) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-2000202202001212-1210220312331330-0110131133032212-1300311323302110-3101211203203101-2321001303100120-3312232210121021-0332220100000133) |
| `primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-2202033002002021-1000111021220031-1331033210122233-3313031000000211-3000300233113102-2111230330123332-2230020201302231-0023002103223131) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-2012033320333210-1320312201333310-0331302222232311-1003331302102232-1032101231112102-1231021021230101-3333201132021111-0232113010211320) |
| `primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.cds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-002.md#canonical-2211030303020013-0331200320121123-3212311320021110-1031231332012103-2131122222313110-2022313331101021-0200232030320220-1201101232112112) |
| `primary.rr_set_group.rr_set.cert_record` | [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-1333323223212332-2013302100133331-1221312031023331-0321203302021123-3333310322103003-1213310001113303-0211312030311301-3023103012323033) |
| `primary.rr_set_group.rr_set.cert_record.name` | [primary.rr_set_group.rr_set.cert_record.name](data-sources--dns_zone--reference--group-002.md#canonical-0012022202310032-3221110201121102-1101331203011021-2020213332011022-1330102310002331-0000003100331121-2100120221011233-0303223210232310) |
| `primary.rr_set_group.rr_set.cert_record.values` | [primary.rr_set_group.rr_set.cert_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2130123321101122-3033233202331100-1111013012031301-1313123121032113-3023223120331121-2000332312300303-3000022210322301-2100102202121122) |
| `primary.rr_set_group.rr_set.cert_record.values.algorithm` | [primary.rr_set_group.rr_set.cert_record.values.algorithm](data-sources--dns_zone--reference--group-002.md#canonical-3333032312101212-2202131123321222-1020131320202200-3323220001231020-3333330012103133-2031012001301320-2003021211022132-3322000303031121) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_key_tag` | [primary.rr_set_group.rr_set.cert_record.values.cert_key_tag](data-sources--dns_zone--reference--group-002.md#canonical-3100313120323021-3110303122033131-1131312211110202-1133031331222312-3233300031121000-0113332033231221-1003232210002310-2001111102013021) |
| `primary.rr_set_group.rr_set.cert_record.values.cert_type` | [primary.rr_set_group.rr_set.cert_record.values.cert_type](data-sources--dns_zone--reference--group-002.md#canonical-1001133220032031-1213300121231022-1001202300323021-1232011122201033-2233112130131312-3300203202310010-2121310221121232-2033001023012003) |
| `primary.rr_set_group.rr_set.cert_record.values.certificate` | [primary.rr_set_group.rr_set.cert_record.values.certificate](data-sources--dns_zone--reference--group-002.md#canonical-1022303100312121-2133320033312012-3220332213330112-0202030130002133-3112131122023121-2101032110003213-2222023222321232-3003020310021133) |
| `primary.rr_set_group.rr_set.cname_record` | [primary.rr_set_group.rr_set.cname_record](data-sources--dns_zone--reference--group-003.md#canonical-1030012213303033-0002303312011021-3111130311232311-0020011033100021-0332202101323301-3133002320230131-2222332200333102-0102302112003002) |
| `primary.rr_set_group.rr_set.cname_record.name` | [primary.rr_set_group.rr_set.cname_record.name](data-sources--dns_zone--reference--group-003.md#canonical-2003100302220033-3122311233212001-0113022310332003-3200210312121032-0000313000310133-0001200132032130-1301203111221020-0103323320203321) |
| `primary.rr_set_group.rr_set.cname_record.value` | [primary.rr_set_group.rr_set.cname_record.value](data-sources--dns_zone--reference--group-003.md#canonical-3330130313321200-3212011323333300-3023021223312031-3210212311310232-2232213012131223-3301030123122211-2032112121113302-3021112323212230) |
| `primary.rr_set_group.rr_set.description_spec` | [primary.rr_set_group.rr_set.description_spec](data-sources--dns_zone--reference--group-002.md#canonical-2000320222210232-1300311103033231-3231232123032330-3001223121001231-2331332212302100-1203320111201203-0121133303013112-1110222233321033) |
| `primary.rr_set_group.rr_set.ds_record` | [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-003.md#canonical-3323123231232232-3121122132113033-2010032032232020-0223223101013113-1320002203003203-0021121122213032-0110113000320310-3131320002200321) |
| `primary.rr_set_group.rr_set.ds_record.name` | [primary.rr_set_group.rr_set.ds_record.name](data-sources--dns_zone--reference--group-003.md#canonical-2120010001021333-2330312330231221-3013231302312000-1012330233333101-2112330010102121-3230301012220311-2211313001301132-3201002122332000) |
| `primary.rr_set_group.rr_set.ds_record.values` | [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2002210212313213-0103320021303300-3103231001322333-2111012132002320-3231020112223333-1320100103000123-0230020232113323-1030130320013033) |
| `primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm` | [primary.rr_set_group.rr_set.ds_record.values.ds_key_algorithm](data-sources--dns_zone--reference--group-003.md#canonical-2000112311312320-0220020111313121-0200111002020332-1212301331001213-1313232133301232-1332113323012120-1122200133112113-2122201122320213) |
| `primary.rr_set_group.rr_set.ds_record.values.key_tag` | [primary.rr_set_group.rr_set.ds_record.values.key_tag](data-sources--dns_zone--reference--group-003.md#canonical-2230223230000300-1232232312112300-2032233211113100-0022102120111130-1120222113201100-2330003332012331-1323220230211111-0301333031331021) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-003.md#canonical-1311001023013010-2120332201020300-0330303232223012-0003321111113220-3233132311300323-0331100222123203-0010100312122300-0202111211311011) |
| `primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha1_digest.digest](data-sources--dns_zone--reference--group-003.md#canonical-0211022311323301-2333203131012220-1031110301003033-1132331011133230-3030133201011313-2100123333112000-1122031102003013-1032211121233101) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-003.md#canonical-3120333302320123-1021333022223330-3302111213102300-2111313030123132-0303031303033232-0210010310002012-0212000111133301-2111312021021231) |
| `primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha256_digest.digest](data-sources--dns_zone--reference--group-003.md#canonical-0000202333011311-2211122320113121-2311200101111101-3200233123201201-2222202310123300-3133121112121013-1213110333131131-3032030033221030) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-003.md#canonical-3003231321331023-2130302302202211-1301303101130013-1003130313231020-1111332313312312-1033021122012133-3101033031223101-2333020010213003) |
| `primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest` | [primary.rr_set_group.rr_set.ds_record.values.sha384_digest.digest](data-sources--dns_zone--reference--group-003.md#canonical-1101221320221223-1231332032211112-3231113020303321-0300321103023000-1133300001133003-3132032231213321-0001011320020302-2303232101212231) |
| `primary.rr_set_group.rr_set.eui48_record` | [primary.rr_set_group.rr_set.eui48_record](data-sources--dns_zone--reference--group-003.md#canonical-2001000003012201-0103211033302211-3210331303310200-0132323302313312-2131011003203221-3101321110100030-1212121121002232-2010002212211130) |
| `primary.rr_set_group.rr_set.eui48_record.name` | [primary.rr_set_group.rr_set.eui48_record.name](data-sources--dns_zone--reference--group-003.md#canonical-3130212232120323-1122323212113010-0001012120200023-0201130300120321-2230111012303020-1110303000001301-0000030020001230-3011110120212102) |
| `primary.rr_set_group.rr_set.eui48_record.value` | [primary.rr_set_group.rr_set.eui48_record.value](data-sources--dns_zone--reference--group-003.md#canonical-1121220222210001-2202001311302222-3222323232333211-2020031120112200-2111201311232301-1330201121301331-3302310300220021-1123133321323333) |
| `primary.rr_set_group.rr_set.eui64_record` | [primary.rr_set_group.rr_set.eui64_record](data-sources--dns_zone--reference--group-003.md#canonical-1212201231123200-0110310002232221-1000113123120321-1200121211201322-1123100112113002-1230103230003231-3122120313201201-3133223022301123) |
| `primary.rr_set_group.rr_set.eui64_record.name` | [primary.rr_set_group.rr_set.eui64_record.name](data-sources--dns_zone--reference--group-003.md#canonical-3223031003000022-1301323031312210-1302322221021013-1133310210310132-1220023121201013-0220301331002032-0330103110030000-1211221101132120) |
| `primary.rr_set_group.rr_set.eui64_record.value` | [primary.rr_set_group.rr_set.eui64_record.value](data-sources--dns_zone--reference--group-003.md#canonical-2331200101303131-2201302132103211-1310001230130200-3321132230320223-3133322100001223-1011002231031201-3112211203302113-1311103100303112) |
| `primary.rr_set_group.rr_set.lb_record` | [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-1010002011111230-0030000301202311-2013002322300101-3033112001200130-3300230103231010-2032221112102132-3233310201023023-1110101302322332) |
| `primary.rr_set_group.rr_set.lb_record.name` | [primary.rr_set_group.rr_set.lb_record.name](data-sources--dns_zone--reference--group-003.md#canonical-0000321020211220-0303010000002220-0212332120131033-2232210021012112-3213233323302222-2022012000131232-0211310122022312-0213101122333323) |
| `primary.rr_set_group.rr_set.lb_record.value` | [primary.rr_set_group.rr_set.lb_record.value](data-sources--dns_zone--reference--group-003.md#canonical-3121320223110223-1332123103002222-2313010003220230-2012231231333302-1010223011211013-2321221233000212-0331033121313101-0023233010031111) |
| `primary.rr_set_group.rr_set.lb_record.value.name` | [primary.rr_set_group.rr_set.lb_record.value.name](data-sources--dns_zone--reference--group-003.md#canonical-0000202220212313-3111321311123302-2131021122203101-3000230310120101-1112231002303110-1303023323321232-0013211232021121-2031223113331022) |
| `primary.rr_set_group.rr_set.lb_record.value.namespace` | [primary.rr_set_group.rr_set.lb_record.value.namespace](data-sources--dns_zone--reference--group-003.md#canonical-3313322000321000-3113213123012211-2032012223101033-2231033201320333-1310023100033020-0100030313121201-0210333130312303-2221022213230133) |
| `primary.rr_set_group.rr_set.lb_record.value.tenant` | [primary.rr_set_group.rr_set.lb_record.value.tenant](data-sources--dns_zone--reference--group-003.md#canonical-2330303013222230-0011223112102330-0022222023310020-2032231332020121-1232102130201023-1202203103200000-0203222211233301-3332202201032101) |
| `primary.rr_set_group.rr_set.loc_record` | [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-2310233030001321-2302022120200120-2011323023012232-3330132222012231-3121102131100201-0332131311020001-2001331033212200-1023211110112202) |
| `primary.rr_set_group.rr_set.loc_record.name` | [primary.rr_set_group.rr_set.loc_record.name](data-sources--dns_zone--reference--group-003.md#canonical-3132112120333211-2233232033231211-1000331213201310-1101220230213201-0011032113331230-1030330122200111-1031020013333310-2333122001003212) |
| `primary.rr_set_group.rr_set.loc_record.values` | [primary.rr_set_group.rr_set.loc_record.values](data-sources--dns_zone--reference--group-003.md#canonical-1211112130033313-0001333210220013-1012303312003132-1302203221002201-1011022323200210-2103131232211132-1112322032113322-2101333131202112) |
| `primary.rr_set_group.rr_set.loc_record.values.altitude` | [primary.rr_set_group.rr_set.loc_record.values.altitude](data-sources--dns_zone--reference--group-003.md#canonical-2222132112000030-0310330123330132-0321123120201130-0111123130111021-2322113203022132-3221101203323322-3122223211202021-1220121231133330) |
| `primary.rr_set_group.rr_set.loc_record.values.horizontal_precision` | [primary.rr_set_group.rr_set.loc_record.values.horizontal_precision](data-sources--dns_zone--reference--group-003.md#canonical-3022330112300003-2101102102123232-2033210031113203-3002322120120220-2213030031112110-2311010213002330-3012311300031232-3021121331102103) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.latitude_degree](data-sources--dns_zone--reference--group-003.md#canonical-3110222123202312-3023302020220102-3303211313030132-2333010102201203-3221232320101031-0213101110130202-2121323110220213-1300013223030123) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.latitude_hemisphere](data-sources--dns_zone--reference--group-003.md#canonical-0111330302313111-2332211212300211-3101302322000200-1233233202312011-0210232101131002-1303121021013232-3223312011021132-3201300323021120) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.latitude_minute](data-sources--dns_zone--reference--group-003.md#canonical-3200123113211231-3332120110000101-0110320132310200-0001222030032322-0310122331020230-2311203101032201-2301033033103122-2321133131203232) |
| `primary.rr_set_group.rr_set.loc_record.values.latitude_second` | [primary.rr_set_group.rr_set.loc_record.values.latitude_second](data-sources--dns_zone--reference--group-003.md#canonical-2210102333110102-0031330021011331-3020033131333202-1033111233002220-2110211233321231-0032032001331021-3103002220111222-0202321222020020) |
| `primary.rr_set_group.rr_set.loc_record.values.location_diameter` | [primary.rr_set_group.rr_set.loc_record.values.location_diameter](data-sources--dns_zone--reference--group-003.md#canonical-1212030312111130-1312210022310221-1010201220330032-0220302120333312-2222332330230100-0032301211312102-1011212222221011-0112212201020311) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_degree` | [primary.rr_set_group.rr_set.loc_record.values.longitude_degree](data-sources--dns_zone--reference--group-003.md#canonical-2111001200013213-0110120330110322-0331221330203303-3102132333302121-3110020033013120-2310032301031003-2312013123323222-0010323001321223) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere` | [primary.rr_set_group.rr_set.loc_record.values.longitude_hemisphere](data-sources--dns_zone--reference--group-003.md#canonical-3111201232032021-1220021122201120-2033010222130000-0022210021203103-3101002332121313-1311331012330212-3000003302112111-2123010131000321) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_minute` | [primary.rr_set_group.rr_set.loc_record.values.longitude_minute](data-sources--dns_zone--reference--group-003.md#canonical-1030211331000333-2301033202230211-1103101233121032-1032131212102230-1332223222132232-0312010111001312-3211332012212022-3020023001300003) |
| `primary.rr_set_group.rr_set.loc_record.values.longitude_second` | [primary.rr_set_group.rr_set.loc_record.values.longitude_second](data-sources--dns_zone--reference--group-003.md#canonical-2331223312002332-2222001221322123-0102102132203110-3321230121322300-1032102220010030-0211012221133102-1011212332200333-1031221332101102) |
| `primary.rr_set_group.rr_set.loc_record.values.vertical_precision` | [primary.rr_set_group.rr_set.loc_record.values.vertical_precision](data-sources--dns_zone--reference--group-003.md#canonical-2200331331113020-0321300301110030-1002113022231312-0123033233113033-2101130132120012-3012022111201103-2230222003012022-0332210312312211) |
| `primary.rr_set_group.rr_set.mx_record` | [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-1032301230121302-2230303233101300-3200131020003011-3223231131120111-3211000013130303-3113013032123030-2313121333332032-0030212331102010) |
| `primary.rr_set_group.rr_set.mx_record.name` | [primary.rr_set_group.rr_set.mx_record.name](data-sources--dns_zone--reference--group-003.md#canonical-3311033132001020-0102022220030111-1232031110121310-2231022232031233-1133110222101113-1132313012310030-2313023133102322-1333311331220330) |
| `primary.rr_set_group.rr_set.mx_record.values` | [primary.rr_set_group.rr_set.mx_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2010113303312310-1321103120212023-2203002233330103-3133021300012001-0313201120310203-2332220233012323-3310013003130201-2321220322121321) |
| `primary.rr_set_group.rr_set.mx_record.values.domain` | [primary.rr_set_group.rr_set.mx_record.values.domain](data-sources--dns_zone--reference--group-003.md#canonical-2212112331311011-1102003102311101-3303301010021132-1132200332102233-1303330021322112-2102312121323021-3020103332131030-3211011300301002) |
| `primary.rr_set_group.rr_set.mx_record.values.priority` | [primary.rr_set_group.rr_set.mx_record.values.priority](data-sources--dns_zone--reference--group-003.md#canonical-1102132003010311-2221313113001020-3201012012000233-0212222330213313-0203330131320113-1211222100232300-3202231313221013-3313300300310020) |
| `primary.rr_set_group.rr_set.naptr_record` | [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-0133232032210101-0001200232000323-1130202303131203-3331333233303332-0012131002202331-0320323131233332-0000212013233331-1011012221020233) |
| `primary.rr_set_group.rr_set.naptr_record.name` | [primary.rr_set_group.rr_set.naptr_record.name](data-sources--dns_zone--reference--group-003.md#canonical-1212011312322322-0312203331301013-3322121111210101-2200211220000111-3032112033200333-0300322110002031-3202232010012220-1112222133303213) |
| `primary.rr_set_group.rr_set.naptr_record.values` | [primary.rr_set_group.rr_set.naptr_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2200231101020302-0013202003222102-3000130322230201-1032213200122031-0022233212211123-0021232220010203-3020201223101211-0110020030232033) |
| `primary.rr_set_group.rr_set.naptr_record.values.flags` | [primary.rr_set_group.rr_set.naptr_record.values.flags](data-sources--dns_zone--reference--group-003.md#canonical-1031012120012113-1313312322333103-2032103222123012-2121333312111012-3031130302031121-2202230212033331-0203301113002231-2312323013113013) |
| `primary.rr_set_group.rr_set.naptr_record.values.order` | [primary.rr_set_group.rr_set.naptr_record.values.order](data-sources--dns_zone--reference--group-003.md#canonical-2311202231301323-3121333022301101-2223030001121022-0022120312113201-0303102103310032-2220312031222320-0312013323332111-2000122301101021) |
| `primary.rr_set_group.rr_set.naptr_record.values.preference` | [primary.rr_set_group.rr_set.naptr_record.values.preference](data-sources--dns_zone--reference--group-003.md#canonical-0032220230321232-0210110100300333-0111130031311110-3013012130120210-0211233310211013-1131033010331121-2223031133003030-2013111232122032) |
| `primary.rr_set_group.rr_set.naptr_record.values.regexp` | [primary.rr_set_group.rr_set.naptr_record.values.regexp](data-sources--dns_zone--reference--group-003.md#canonical-0101230230333232-1332200302210321-3130112222202020-0102100120101330-0121133023313321-3112011001110013-2323021301022023-0003332310132202) |
| `primary.rr_set_group.rr_set.naptr_record.values.replacement` | [primary.rr_set_group.rr_set.naptr_record.values.replacement](data-sources--dns_zone--reference--group-003.md#canonical-2200200231122010-1201013102211212-2301011122201231-2223033312113200-2311302021012310-3030120312101310-0121001121111000-0222110200010111) |
| `primary.rr_set_group.rr_set.naptr_record.values.service` | [primary.rr_set_group.rr_set.naptr_record.values.service](data-sources--dns_zone--reference--group-003.md#canonical-0222322101131033-2330013302030033-2322120201123121-2313311332010213-2210100232003131-3301000021132131-1001003200202201-0313032032003033) |
| `primary.rr_set_group.rr_set.ns_record` | [primary.rr_set_group.rr_set.ns_record](data-sources--dns_zone--reference--group-003.md#canonical-0330312201100120-2112312301203220-3230003000130221-0023232000030332-3220220102110023-1132123131100021-3103230321221112-3322223320131103) |
| `primary.rr_set_group.rr_set.ns_record.name` | [primary.rr_set_group.rr_set.ns_record.name](data-sources--dns_zone--reference--group-003.md#canonical-0200211233030201-3122121200101333-3220202313311001-1213312221301101-2333121220001202-3100112100122322-0301211011133220-3031313202011301) |
| `primary.rr_set_group.rr_set.ns_record.values` | [primary.rr_set_group.rr_set.ns_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2121333000030022-1111112311202130-1321222312321100-0231012012113003-0022310110111233-3320111233223231-1230011103023030-3130102330101301) |
| `primary.rr_set_group.rr_set.ptr_record` | [primary.rr_set_group.rr_set.ptr_record](data-sources--dns_zone--reference--group-003.md#canonical-3030302130012103-1031200320000313-3213320132112032-3011133211333100-2211102023313110-0000310203120132-2233122200311201-0012110010030112) |
| `primary.rr_set_group.rr_set.ptr_record.name` | [primary.rr_set_group.rr_set.ptr_record.name](data-sources--dns_zone--reference--group-003.md#canonical-2320112321012331-2211032133102013-2012112213232302-2200312233332133-2030122012311323-3000122321100123-2301210012312211-0232130132131110) |
| `primary.rr_set_group.rr_set.ptr_record.values` | [primary.rr_set_group.rr_set.ptr_record.values](data-sources--dns_zone--reference--group-003.md#canonical-3022000221212222-1223011132033230-1001032003202132-3013221033030122-0030010331211121-3002101102030231-2032222322202102-0122110132221302) |
| `primary.rr_set_group.rr_set.srv_record` | [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-1000221220030112-2010100232132001-3022002120021130-1103323000130332-1102212231213330-0122231210232320-2020130132322232-3121222113332300) |
| `primary.rr_set_group.rr_set.srv_record.name` | [primary.rr_set_group.rr_set.srv_record.name](data-sources--dns_zone--reference--group-003.md#canonical-2302322211300310-3222300011022310-0210330303222112-1321232112133201-1223132210322111-1213310002320301-2223302100310113-0202212113031003) |
| `primary.rr_set_group.rr_set.srv_record.values` | [primary.rr_set_group.rr_set.srv_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2212233103113113-0311022000300322-0310231201320310-1100021111230311-2303130003123100-2223332202103301-0111000113330012-1221313232200032) |
| `primary.rr_set_group.rr_set.srv_record.values.port` | [primary.rr_set_group.rr_set.srv_record.values.port](data-sources--dns_zone--reference--group-003.md#canonical-3003201220200130-1021330232001031-2321011021133032-0002310022312232-2102213113331110-0132012022323033-1223031023100212-0321202210233230) |
| `primary.rr_set_group.rr_set.srv_record.values.priority` | [primary.rr_set_group.rr_set.srv_record.values.priority](data-sources--dns_zone--reference--group-003.md#canonical-3210220321011103-0221123113220123-3101321201123011-3132321110103301-2232212102010012-3320012113021320-3201132330203101-1003103012001021) |
| `primary.rr_set_group.rr_set.srv_record.values.target` | [primary.rr_set_group.rr_set.srv_record.values.target](data-sources--dns_zone--reference--group-003.md#canonical-2222300302002330-0331121232303121-1323233320131120-2101102203212132-2213100223130100-1131220022202331-2100132311031101-0210020030000301) |
| `primary.rr_set_group.rr_set.srv_record.values.weight` | [primary.rr_set_group.rr_set.srv_record.values.weight](data-sources--dns_zone--reference--group-003.md#canonical-0313220012311033-3131021312031102-2010003303132000-2313020302003312-3301221231001320-0300103203001322-1232303020123102-3030021033312210) |
| `primary.rr_set_group.rr_set.sshfp_record` | [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-2121012320102322-0110331111102013-2132110133203320-3100031022001321-0032102321220212-3221111323222212-3100312000000230-0120322113132122) |
| `primary.rr_set_group.rr_set.sshfp_record.name` | [primary.rr_set_group.rr_set.sshfp_record.name](data-sources--dns_zone--reference--group-003.md#canonical-1231031000113201-3321323320010331-3323322201120313-0222231133212021-1013111310313302-1232032001323032-1233021231002103-3113110111201021) |
| `primary.rr_set_group.rr_set.sshfp_record.values` | [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2331300330310133-0031203323201200-0020002201201200-0030033113010311-3330221213323321-1321210003322111-3011230300321203-0332322021001321) |
| `primary.rr_set_group.rr_set.sshfp_record.values.algorithm` | [primary.rr_set_group.rr_set.sshfp_record.values.algorithm](data-sources--dns_zone--reference--group-003.md#canonical-1300023031211311-3313232011221303-0122212033232122-0212102021130121-1002320332020233-2310132330313320-1231033230020221-1213001232222300) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-1310013002022232-3123100310320020-3302132323103132-3010312231002020-0322211133101202-0330130210120012-0023132001322001-0221131011332221) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint.fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-0232000100001222-2300203103302233-3121333113330031-3221212301303111-2210323011131110-1130103021323111-0102312333130231-1103232131000330) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-2021231231131011-2002010300220312-2013311332002322-2323103231333301-3032102030233320-0321333100133023-2223003113133021-3222311112110112) |
| `primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint` | [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint.fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-3201203001133212-0131133232032213-2011311202003222-2021102032301211-1201033223000230-2302323122302213-2013030113111010-0021122211020113) |
| `primary.rr_set_group.rr_set.tlsa_record` | [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-1101223212221102-3322031333001130-1103322111122133-1033323303213231-3332201223232100-1321231100221231-2130122203112013-0222012110223001) |
| `primary.rr_set_group.rr_set.tlsa_record.name` | [primary.rr_set_group.rr_set.tlsa_record.name](data-sources--dns_zone--reference--group-003.md#canonical-3013011123320300-1330002211301302-0010222000332022-0331110332323222-3013010011022032-0310312331031133-3232332231113010-0011100130233301) |
| `primary.rr_set_group.rr_set.tlsa_record.values` | [primary.rr_set_group.rr_set.tlsa_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2312332012210332-2122213103112133-1132301230330200-0223000200211331-1121010310320232-3020220333130203-3313313323312102-2330321021132333) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_association_data](data-sources--dns_zone--reference--group-003.md#canonical-2022331111110123-2100103133330003-2011211220002120-2200222331011202-0233133231121211-2100130131020002-0020133032101201-1120100312130101) |
| `primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage` | [primary.rr_set_group.rr_set.tlsa_record.values.certificate_usage](data-sources--dns_zone--reference--group-003.md#canonical-1032033121302201-0313000230112110-1201323010101122-2203310201303130-2302113130330111-3201013302201122-1023203000012300-0202130302300122) |
| `primary.rr_set_group.rr_set.tlsa_record.values.matching_type` | [primary.rr_set_group.rr_set.tlsa_record.values.matching_type](data-sources--dns_zone--reference--group-003.md#canonical-2231300020001110-2202032131312021-3013103200231130-1122101031311230-3002123203312100-1023230201233010-3121330111100032-2221122330021322) |
| `primary.rr_set_group.rr_set.tlsa_record.values.selector` | [primary.rr_set_group.rr_set.tlsa_record.values.selector](data-sources--dns_zone--reference--group-003.md#canonical-0301100322000222-2210310111030100-0120200130323031-0313300323302321-2020023000003023-2230202232131031-2302031133313230-0222231030103131) |
| `primary.rr_set_group.rr_set.ttl` | [primary.rr_set_group.rr_set.ttl](data-sources--dns_zone--reference--group-002.md#canonical-0113323032231011-0223333222302211-3322300003130030-0323113011131110-0313123301200332-3032113230120122-0323110023212100-0112221033132131) |
| `primary.rr_set_group.rr_set.txt_record` | [primary.rr_set_group.rr_set.txt_record](data-sources--dns_zone--reference--group-003.md#canonical-3120013321330311-3132233230321201-0111310300333022-1333330321012022-3022221032332233-0201113210031120-0121111121022320-2322222212030010) |
| `primary.rr_set_group.rr_set.txt_record.name` | [primary.rr_set_group.rr_set.txt_record.name](data-sources--dns_zone--reference--group-003.md#canonical-1230123002302333-0320100331032333-0332032213020213-2313010123113222-3130032100232102-0001031233212030-3111013300000210-2232023222320232) |
| `primary.rr_set_group.rr_set.txt_record.values` | [primary.rr_set_group.rr_set.txt_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2313020101101202-1330213132331132-0000110133301131-0012232110110003-3303020000303202-3130122103003331-0120203331232122-3001322231031230) |
| `primary.soa_parameters` | [primary.soa_parameters](data-sources--dns_zone--reference--group-003.md#canonical-3332200032033121-0332301033131232-0222202113000111-3302133332331313-1302101233232222-2003033200120003-1030130033320323-1003113213310120) |
| `primary.soa_parameters.expire` | [primary.soa_parameters.expire](data-sources--dns_zone--reference--group-003.md#canonical-1201033000303111-2223303120002331-0021033000232021-3123122112002012-1123232021130110-3023211303113021-3311312133110210-0002231111232302) |
| `primary.soa_parameters.negative_ttl` | [primary.soa_parameters.negative_ttl](data-sources--dns_zone--reference--group-003.md#canonical-3100223313012320-2031021332301022-2331201032203202-2002023020101122-2110133031121231-3201313011302023-0120221133100100-3220113013302122) |
| `primary.soa_parameters.refresh` | [primary.soa_parameters.refresh](data-sources--dns_zone--reference--group-003.md#canonical-3120102023112030-3232121120322003-0021302112212101-2201012221110322-0220003032222203-2310310313302020-1202332331112213-0131133131130230) |
| `primary.soa_parameters.retry` | [primary.soa_parameters.retry](data-sources--dns_zone--reference--group-003.md#canonical-0132113021222221-2013001303010022-3221122332130011-0033113000132220-1303321123221002-0302320220223033-2122331011201321-2323110332310111) |
| `primary.soa_parameters.ttl` | [primary.soa_parameters.ttl](data-sources--dns_zone--reference--group-003.md#canonical-3213120213122202-0213310131301121-3202133221121323-0102313213023020-0023231121222111-1321320022330200-1010321303222311-0122010110212101) |
| `secondary` | [secondary](data-sources--dns_zone--reference--group-003.md#canonical-3313001100221131-0333113010030320-1300031010002223-1231331333010010-0122201331100023-3022321102020300-0200013201131111-3320301233010323) |
| `secondary.primary_servers` | [secondary.primary_servers](data-sources--dns_zone--reference--group-003.md#canonical-2231231030010330-2312101322001132-3330301120301212-2132112131013123-3301303301023113-1220032023023021-1321133312132220-0110301231200022) |
| `secondary.tsig_key_algorithm` | [secondary.tsig_key_algorithm](data-sources--dns_zone--reference--group-003.md#canonical-3330223310100001-0110311112003331-3013103213122000-2103321322103210-2012323002310000-1321203300233211-2200023201100210-2022222311131132) |
| `secondary.tsig_key_name` | [secondary.tsig_key_name](data-sources--dns_zone--reference--group-003.md#canonical-2032021232222221-3331232103113031-0222131211003030-0111213310012111-0120232201320103-3220323322013310-1223033222120332-1100022012211321) |
| `secondary.tsig_key_value` | [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-3101233323332301-0312030330010020-0112200003233111-3030302033213132-1231033132112202-1221323011331233-3132101020311220-0322033320313121) |
| `secondary.tsig_key_value.blindfold_secret_info` | [secondary.tsig_key_value.blindfold_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-3322021301003221-1011000011013122-3221210313313322-1010211101010230-0021322012320102-3301222202013311-0330313211203121-3310320323300313) |
| `secondary.tsig_key_value.blindfold_secret_info.decryption_provider` | [secondary.tsig_key_value.blindfold_secret_info.decryption_provider](data-sources--dns_zone--reference--group-003.md#canonical-2210020032131012-1313301020210112-2223100111000300-1320110301123013-0121113322102313-0110122232020112-2301213112313002-3122000002222200) |
| `secondary.tsig_key_value.blindfold_secret_info.location` | [secondary.tsig_key_value.blindfold_secret_info.location](data-sources--dns_zone--reference--group-003.md#canonical-1300313031010331-1211322202132302-0323221133333301-2120210322212120-2320303010232122-3302202122310030-0301313211010021-1100333111010221) |
| `secondary.tsig_key_value.blindfold_secret_info.store_provider` | [secondary.tsig_key_value.blindfold_secret_info.store_provider](data-sources--dns_zone--reference--group-003.md#canonical-1332321201022032-0211212221300332-2022310122013230-2211233030113203-0220013000023200-1133221300112003-2122103012323023-0322031022312012) |
| `secondary.tsig_key_value.clear_secret_info` | [secondary.tsig_key_value.clear_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-0002213001331323-1200100121010123-2301120030233222-2203330130212332-1331302213010022-1330213231323202-2010321231203310-0110223012100222) |
| `secondary.tsig_key_value.clear_secret_info.provider_ref` | [secondary.tsig_key_value.clear_secret_info.provider_ref](data-sources--dns_zone--reference--group-003.md#canonical-3012021300301121-2020233020332311-3000200220320123-2000003303023023-3202331302310133-2231133310330133-3302321211300300-3300030030031111) |
| `secondary.tsig_key_value.clear_secret_info.url` | [secondary.tsig_key_value.clear_secret_info.url](data-sources--dns_zone--reference--group-003.md#canonical-1330123220130302-2101200121133331-3010022021113122-3331023211020110-2230132031312100-0101030030100311-3213331132110111-3103103113303113) |

<a id="canonical-0203331121332210-0000023210001001-2200220012012331-2022032001000120-0323122103220120-0212120213033133-0133302022333010-1133123133232122"></a>

## Next pages — Property reference / 121200323121 / 11

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301220220111313-2203102101003321-1221012203131121-3123002321131032-1321330032103313-3302010213113210-1221201022213033-0131110130002023"></a>

## primary — primary / 332231132222 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- primary

<a id="canonical-2121010121102223-3310323133310112-1113111110203223-3000201032032200-1001233231232110-0102013302313230-3223110102210011-1002120320220331"></a>

Type: `"single"`. Computed.

\[OneOf: primary, secondary\] PrimaryDNSCreateSpecType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-soa_record_parameters_choice": "[\"default_soa_parameters\",\"soa_parameters\"]"
}
```

OneOf alternatives in this subsection:

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-2121010121102223-3310323133310112-1113111110203223-3000201032032200-1001233231232110-0102013302313230-3223110102210011-1002120320220331)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-3313001100221131-0333113010030320-1300031010002223-1231331333010010-0122201331100023-3022321102020300-0200013201131111-3320301233010323)

Select alternatives according to the provider validators above.

<a id="canonical-3132231021310111-0212111121031210-2223311112313012-3022221112111231-3032222032101122-2131122201003001-2303213211222233-0220321313331203"></a>

## Direct properties — primary / 332231132222 / 3

<a id="canonical-0011201311331332-0000001012122003-1122202131303312-2201102031231123-1020212100002012-3212210231021302-3313210321012000-2333130112103312"></a>

<a id="canonical-1303221103302302-1031120233331001-1101221332101032-1320101020221300-1313000230020211-1110232032020011-2123220312011030-2120321033123113"></a>

## allow_http_lb_managed_records property — primary / 332231132222 / 4

Type: `"bool"`. Computed.

Option to allow user-created HTTP, TCP, and CDN load balancer related resource records to be
automatically managed in a protected RRset.

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

- [default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111): complete subsection reference.

- [default_soa_parameters](data-sources--dns_zone--reference--group-002.md#canonical-1013032121121211-0200310230111302-2302123231122131-1132011211233131-3233303200303110-1001330123313000-0122110323001111-2131223301100220): complete subsection reference.

- [dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102): complete subsection reference.

- [rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202): complete subsection reference.

- [soa_parameters](data-sources--dns_zone--reference--group-003.md#canonical-3111301200310120-2012211111103233-2001110322021102-1132112033133212-1100110333112313-0210200023120120-3233033130322010-1032311322222221): complete subsection reference.

<a id="canonical-0003312130202111-1210212030211211-2002321303220013-1130213201020332-1222311020231301-2132111011312132-2221221002333323-3000221012110012"></a>

## Next pages — primary / 332231132222 / 5

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_soa_parameters](data-sources--dns_zone--reference--group-002.md#canonical-1013032121121211-0200310230111302-2302123231122131-1132011211233131-3233303200303110-1001330123313000-0122110323001111-2131223301100220)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.soa_parameters](data-sources--dns_zone--reference--group-003.md#canonical-3111301200310120-2012211111103233-2001110322021102-1132112033133212-1100110333112313-0210200023120120-3233033130322010-1032311322222221)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320121001101331-1300110132013031-3121323222022203-0122321012230031-3010112333212023-3210102002130110-3201002030121023-2123001011312300"></a>

## primary.default_rr_set_group — default_rr_set_group / 131131132311 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.default_rr_set_group

<a id="canonical-1113102033121113-1000301100233101-1323300223303133-1331120113002132-3232122222030111-1202002013022232-2332320011300313-0020230311202010"></a>

Type: `"list"`. Computed.

Add and manage DNS resource record sets part of Default set group.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
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
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

<a id="canonical-0311010301221331-0113013131132013-0103133320213313-0022000012311233-0222333100111002-1023011122033121-3302320231130031-0310312201320233"></a>

## Direct properties — default_rr_set_group / 131131132311 / 3

- [a_record](data-sources--dns_zone--reference--group-001.md#canonical-2312200030111200-1212321102221122-1303233301122023-2113100103230120-3230111301023210-1321113113303300-0311102130231112-3322113202113010): complete subsection reference.

- [aaaa_record](data-sources--dns_zone--reference--group-001.md#canonical-2122110231021031-2311002310121211-0030101131322312-2002213023010323-1003310222101102-3121232322232230-0032201201331221-2022310230031212): complete subsection reference.

- [afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-2212321030130323-0331003210301222-1002023233203223-3001001333203330-2310200301223030-2230033302212033-1133303131313312-0102220132310223): complete subsection reference.

- [alias_record](data-sources--dns_zone--reference--group-001.md#canonical-2201132013312310-0031120122103231-3200313312112001-2022323120201202-3021322302213003-2300121002001030-2300132201121310-0303202132032030): complete subsection reference.

- [caa_record](data-sources--dns_zone--reference--group-001.md#canonical-2231130311212021-2231121000310210-0333021133202030-2212032233303300-2233330122101130-3300333102031111-0312302332003200-2201113110032012): complete subsection reference.

- [cds_record](data-sources--dns_zone--reference--group-001.md#canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220): complete subsection reference.

- [cert_record](data-sources--dns_zone--reference--group-001.md#canonical-2103113313321033-0201132202120131-1011321103103120-0212131230022211-1011221201210111-1203211223103012-3213112020013312-0122002221223221): complete subsection reference.

- [cname_record](data-sources--dns_zone--reference--group-001.md#canonical-1210222313232023-3302132113200010-1102000331232222-1301112000130123-0332301130133312-0112100212202202-3031221112323220-3020332031122200): complete subsection reference.

<a id="canonical-0113312013210200-2113030131021323-3132103200313313-1213300120001232-0113212330132300-1123320311201132-3000210010001030-1130233112210000"></a>

<a id="canonical-2313003101102232-3122030102002120-2021023221102213-0033312203011303-3002232311311111-0000022323303112-0222121300212111-2220003231121001"></a>

## description_spec property — default_rr_set_group / 131131132311 / 4

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](data-sources--dns_zone--reference--group-001.md#canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100): complete subsection reference.

- [eui48_record](data-sources--dns_zone--reference--group-002.md#canonical-3131300332020230-3331033131000131-1121003101002330-1233102210013032-0213323303121120-0011202223031101-3113322221020122-1031012321320303): complete subsection reference.

- [eui64_record](data-sources--dns_zone--reference--group-002.md#canonical-1001031330110020-1133122011223212-1221020310212103-2121030102103310-1031332030211201-3010332112131120-0002121331230022-3122212233332113): complete subsection reference.

- [lb_record](data-sources--dns_zone--reference--group-002.md#canonical-1321110233323023-3033012222210110-1100112210100301-1331110010313330-0202131111030201-0031121030023110-3303012321221102-1202212120312023): complete subsection reference.

- [loc_record](data-sources--dns_zone--reference--group-002.md#canonical-3310122013002202-2201310111100013-3320132322330303-3003311222123331-1221002132100302-2122021012023311-0330301233113311-0002130121010103): complete subsection reference.

- [mx_record](data-sources--dns_zone--reference--group-002.md#canonical-1210121201311120-2131010222033122-0330222031322111-3200022000122302-0333021032123112-2220203201001212-2120030320310132-2322032202111212): complete subsection reference.

- [naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2000221213012233-1222212300001003-2201121122210020-3133301203202102-1332032202320301-1003102301122213-2203320000301213-0012113303201103): complete subsection reference.

- [ns_record](data-sources--dns_zone--reference--group-002.md#canonical-2313231112212011-0201330312333230-1201033232031120-1123102230132211-2311130313033231-0220202113331330-3030032011212201-2022201023220103): complete subsection reference.

- [ptr_record](data-sources--dns_zone--reference--group-002.md#canonical-0032230300312201-2330203102230021-2122122330231113-1232330231302310-3001020022321003-0100000031302303-0320002132012223-3332123310323113): complete subsection reference.

- [srv_record](data-sources--dns_zone--reference--group-002.md#canonical-3022002303322101-2133310220222102-1221303323002002-2021320313022100-3022301033201311-3100301120122232-0121032103302112-2011211020313032): complete subsection reference.

- [sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021): complete subsection reference.

- [tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-2101121000001100-2301010023202222-0201001332133023-3100101303202310-3001221233210201-2132313122010110-3100111130011132-3022131111200230): complete subsection reference.

<a id="canonical-1120233323223031-2202211123332021-0131113311112233-3133300000310331-2210010332013333-0012130212232203-3223202223302113-2300001101113012"></a>

<a id="canonical-3012333303312321-3220301113221011-0130321210100113-0321223201002302-2331002303101030-3233111230332231-0323213212321113-3122023333100111"></a>

## TTL property — default_rr_set_group / 131131132311 / 5

Type: `"number"`. Computed.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](data-sources--dns_zone--reference--group-002.md#canonical-2023003321332312-0023132123210230-1102330021223031-3112003211020100-3031102022010133-0321133333232212-0313320223232033-0232003031030313): complete subsection reference.

<a id="canonical-2200203022221312-0000223213212111-2021203322012301-3202131333212210-2102232023302300-0130130023032331-0011000000223321-2203213213000331"></a>

## Next pages — default_rr_set_group / 131131132311 / 6

- [primary.default_rr_set_group.a_record](data-sources--dns_zone--reference--group-001.md#canonical-2312200030111200-1212321102221122-1303233301122023-2113100103230120-3230111301023210-1321113113303300-0311102130231112-3322113202113010)
- [primary.default_rr_set_group.aaaa_record](data-sources--dns_zone--reference--group-001.md#canonical-2122110231021031-2311002310121211-0030101131322312-2002213023010323-1003310222101102-3121232322232230-0032201201331221-2022310230031212)
- [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-2212321030130323-0331003210301222-1002023233203223-3001001333203330-2310200301223030-2230033302212033-1133303131313312-0102220132310223)
- [primary.default_rr_set_group.alias_record](data-sources--dns_zone--reference--group-001.md#canonical-2201132013312310-0031120122103231-3200313312112001-2022323120201202-3021322302213003-2300121002001030-2300132201121310-0303202132032030)
- [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-2231130311212021-2231121000310210-0333021133202030-2212032233303300-2233330122101130-3300333102031111-0312302332003200-2201113110032012)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220)
- [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-2103113313321033-0201132202120131-1011321103103120-0212131230022211-1011221201210111-1203211223103012-3213112020013312-0122002221223221)
- [primary.default_rr_set_group.cname_record](data-sources--dns_zone--reference--group-001.md#canonical-1210222313232023-3302132113200010-1102000331232222-1301112000130123-0332301130133312-0112100212202202-3031221112323220-3020332031122200)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100)
- [primary.default_rr_set_group.eui48_record](data-sources--dns_zone--reference--group-002.md#canonical-3131300332020230-3331033131000131-1121003101002330-1233102210013032-0213323303121120-0011202223031101-3113322221020122-1031012321320303)
- [primary.default_rr_set_group.eui64_record](data-sources--dns_zone--reference--group-002.md#canonical-1001031330110020-1133122011223212-1221020310212103-2121030102103310-1031332030211201-3010332112131120-0002121331230022-3122212233332113)
- [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-1321110233323023-3033012222210110-1100112210100301-1331110010313330-0202131111030201-0031121030023110-3303012321221102-1202212120312023)
- [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-3310122013002202-2201310111100013-3320132322330303-3003311222123331-1221002132100302-2122021012023311-0330301233113311-0002130121010103)
- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-1210121201311120-2131010222033122-0330222031322111-3200022000122302-0333021032123112-2220203201001212-2120030320310132-2322032202111212)
- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2000221213012233-1222212300001003-2201121122210020-3133301203202102-1332032202320301-1003102301122213-2203320000301213-0012113303201103)
- [primary.default_rr_set_group.ns_record](data-sources--dns_zone--reference--group-002.md#canonical-2313231112212011-0201330312333230-1201033232031120-1123102230132211-2311130313033231-0220202113331330-3030032011212201-2022201023220103)
- [primary.default_rr_set_group.ptr_record](data-sources--dns_zone--reference--group-002.md#canonical-0032230300312201-2330203102230021-2122122330231113-1232330231302310-3001020022321003-0100000031302303-0320002132012223-3332123310323113)
- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-3022002303322101-2133310220222102-1221303323002002-2021320313022100-3022301033201311-3100301120122232-0121032103302112-2011211020313032)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-2101121000001100-2301010023202222-0201001332133023-3100101303202310-3001221233210201-2132313122010110-3100111130011132-3022131111200230)
- [primary.default_rr_set_group.txt_record](data-sources--dns_zone--reference--group-002.md#canonical-2023003321332312-0023132123210230-1102330021223031-3112003211020100-3031102022010133-0321133333232212-0313320223232033-0232003031030313)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2312200030111200-1212321102221122-1303233301122023-2113100103230120-3230111301023210-1321113113303300-0311102130231112-3322113202113010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222001220032131-1331210100122311-0020100300323022-0130311321110331-1313031030121232-2332301303200331-0133132210333023-1013121220222300"></a>

## primary.default_rr_set_group.a_record — a_record / 323001023202 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.a_record

<a id="canonical-2202030000303320-2203033023002101-2121302113212320-1131112300223133-1122103223031130-0003001132133122-0032130303301233-0203212132221021"></a>

Type: `"single"`. Computed.

DNSAResourceRecord. A Records

Upstream description:

A Records

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

<a id="canonical-2311200232200113-1223212213110313-1002123002302322-3002332301223121-3333221010313030-0200211230203002-2220223232330101-0122322213001110"></a>

## Direct properties — a_record / 323001023202 / 3

<a id="canonical-2213233021320010-3133010020313213-2313001313101301-2102002013301311-2202113133323321-1011203112301310-3320200211203133-3003333032001011"></a>

<a id="canonical-2210330212001130-1123321020100312-1303322302033233-3310313203212232-3300033001233012-2003232330101233-0301033213323302-3130002012032210"></a>

## name property — a_record / 323001023202 / 4

Type: `"string"`. Computed.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-1210311210102033-3132330103023322-1313233121311111-2222002211303113-1131020022311300-0032131201212002-3011111312133232-0102300230123221"></a>

<a id="canonical-0130022003000031-2201033120213003-2210223231300221-2320201033033021-2021020210222003-2000000311011301-0201102120021112-1120320122323220"></a>

## values property — a_record / 323001023202 / 5

Type: `["list", "string"]`. Computed.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2002133031131220-1010123022230100-2131320102002211-0222023220203123-2130002322001130-3122021033102120-0133111311213303-2012102313222103"></a>

## Next pages — a_record / 323001023202 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2122110231021031-2311002310121211-0030101131322312-2002213023010323-1003310222101102-3121232322232230-0032201201331221-2022310230031212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332023010000213-3220121310112321-3203222023122002-1231021032212022-3203301011211321-0000112331222231-2332330122123330-0331013210110001"></a>

## primary.default_rr_set_group.aaaa_record — aaaa_record / 210022300122 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.aaaa_record

<a id="canonical-0313033121023202-0312110313020213-1232110322312331-3133000312113310-1011221300011230-2320303010232220-2220220333322322-1320130222113131"></a>

Type: `"single"`. Computed.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

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

<a id="canonical-3200122131332031-3222013012220310-3033332212303013-2330220211223001-1112032321302210-0133111032231132-0203001003210212-1303232221332332"></a>

## Direct properties — aaaa_record / 210022300122 / 3

<a id="canonical-3102033001011313-3312213133110031-1221232032231231-1231030232132110-3323101230131120-2032210102310203-0320231313323302-3022120111322120"></a>

<a id="canonical-0120011100201212-0302010120301133-1101320020222103-0020000331030103-0323320303212331-1122230102202212-0012330201333323-2330102200121122"></a>

## name property — aaaa_record / 210022300122 / 4

Type: `"string"`. Computed.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2303002320003221-3110210003001112-2112001002022133-2002203131120320-2101302123032200-2312303000022032-3212323203320221-1211023011122323"></a>

<a id="canonical-3321112102103223-1032221113311220-2212010203320000-0013301110132132-3202220210021231-2103323112331212-0003331031123331-1111333320303122"></a>

## values property — aaaa_record / 210022300122 / 5

Type: `["list", "string"]`. Computed.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2232102023211011-3002303222012332-2220223032021220-0001120202113122-3320221131333310-1002121003322201-2213213133201203-0110000101023210"></a>

## Next pages — aaaa_record / 210022300122 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2212321030130323-0331003210301222-1002023233203223-3001001333203330-2310200301223030-2230033302212033-1133303131313312-0102220132310223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101131223322330-0113313003100021-3321230012310233-1311022010321023-1133320101321000-0120312330022221-0002213232101331-0102013113112102"></a>

## primary.default_rr_set_group.afsdb_record — afsdb_record / 012330133102 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.afsdb_record

<a id="canonical-3132130120303123-3111331113133222-0003220322132231-1221120021333122-0133313012220113-2012021312102302-1112000213233011-1313022312013111"></a>

Type: `"single"`. Computed.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

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

<a id="canonical-1112132022012120-3002222133120120-1300330121100122-2012001302320001-2320133321111002-3001112102230311-0030130202001103-3123301301003112"></a>

## Direct properties — afsdb_record / 012330133102 / 3

<a id="canonical-0110300331130123-2132312123101333-3313303003231130-3210110311213333-1222000333101001-2331120331301321-3210032333021002-0223203023231010"></a>

<a id="canonical-1003120103133332-0020112202120102-2113213230211200-2031001223033331-1000120302113331-3102303130003321-2121301231211201-1101300111310222"></a>

## name property — afsdb_record / 012330133102 / 4

Type: `"string"`. Computed.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-001.md#canonical-3323233331320022-1000000300301231-2103221320301002-1003123123131031-3002312031121300-0122330002032202-1000311132321011-2013131113033310): complete subsection reference.

<a id="canonical-0110123212220130-0313232111100013-2020003112311202-0133000123102200-0122122222203322-0320132013222330-3101202310023211-1203123320100320"></a>

## Next pages — afsdb_record / 012330133102 / 5

- [primary.default_rr_set_group.afsdb_record.values](data-sources--dns_zone--reference--group-001.md#canonical-3323233331320022-1000000300301231-2103221320301002-1003123123131031-3002312031121300-0122330002032202-1000311132321011-2013131113033310)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3323233331320022-1000000300301231-2103221320301002-1003123123131031-3002312031121300-0122330002032202-1000311132321011-2013131113033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220013201300130-0323333302023120-3202033331010023-1210031003101032-0312100212031130-1210100120011011-3122012103221222-3201200320030020"></a>

## primary.default_rr_set_group.afsdb_record.values — values / 221133320202 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-2212321030130323-0331003210301222-1002023233203223-3001001333203330-2310200301223030-2230033302212033-1133303131313312-0102220132310223)
- primary.default_rr_set_group.afsdb_record.values

<a id="canonical-3130002333022131-0230003023021203-3030303302332332-2010201122323313-1311131232021021-3023030200203011-3122322002332311-3213020132110310"></a>

Type: `"list"`. Computed.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0022123213201213-2010000101113111-0303323320111313-2132221302320121-3002131223322111-3330220130031320-1321020322033100-2123110103322130"></a>

## Direct properties — values / 221133320202 / 3

<a id="canonical-1002033312212101-2331120003311212-1112012311020212-3131132231320032-2131001021113100-3013333100122233-2332101002131012-1330211130110103"></a>

<a id="canonical-0113111312023022-3202132320203333-2331312102033101-2312210030311320-1220200212330201-1312030333030321-3100223323020212-1332212032203332"></a>

## hostname property — values / 221133320202 / 4

Type: `"string"`. Computed.

Server name of the AFS cell database server or the DCE name server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-1321311233222302-0032022013222201-1222322121032202-3331031020302221-2230133312111312-2100330331332211-2030101201212121-2022100230311220"></a>

<a id="canonical-3320023113112102-0033200203222120-3302313230011321-1033012233113202-3312223322112030-2200103103130131-3101022030300310-0233232033020303"></a>

## subtype property — values / 221133320202 / 5

Type: `"string"`. Computed.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0312213332331102-0033230123220011-2332220121023101-3103000230010332-3201133011013202-1022302100011323-0322023030011303-1113211020202230"></a>

## Next pages — values / 221133320202 / 6

- [primary.default_rr_set_group.afsdb_record](data-sources--dns_zone--reference--group-001.md#canonical-2212321030130323-0331003210301222-1002023233203223-3001001333203330-2310200301223030-2230033302212033-1133303131313312-0102220132310223)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2201132013312310-0031120122103231-3200313312112001-2022323120201202-3021322302213003-2300121002001030-2300132201121310-0303202132032030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133203021022221-2112000211022210-3321023121200112-3002103311233021-1123302300131201-3023312233211221-2033201022311302-2130000001020221"></a>

## primary.default_rr_set_group.alias_record — alias_record / 222120021131 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.alias_record

<a id="canonical-3123303123312231-1120201130031210-3000200201321020-1023312200012301-2022110200001130-0010132031321001-3231322013233110-2022311301033311"></a>

Type: `"single"`. Computed.

Configuration parameter for alias record.

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

<a id="canonical-3130200232032300-3002323332302301-2030213030203032-2020330033012320-0332313333031000-0303303132033033-1230020300111031-3230003312303332"></a>

## Direct properties — alias_record / 222120021131 / 3

<a id="canonical-0003333310003211-0232300123200202-3310313300103201-2001033302312032-3012303310010222-3100122323030303-3202303122022202-1230220121111033"></a>

<a id="canonical-0330301113323030-3102301113300132-3232211010031121-0330303331200212-2203321201313123-0223221201123332-3002120101003030-1222222212000111"></a>

## value property — alias_record / 222120021131 / 4

Type: `"string"`. Computed.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-2113320212202010-3313031332120332-1312232032121332-2112122122332302-2233133330320010-1033212231302211-2022223301323310-2331110222231211"></a>

## Next pages — alias_record / 222120021131 / 5

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2231130311212021-2231121000310210-0333021133202030-2212032233303300-2233330122101130-3300333102031111-0312302332003200-2201113110032012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203211313130213-3112101220301323-2202132111012132-0131331023323031-3310332031302023-2333120123322220-2320200221112201-3112011010220112"></a>

## primary.default_rr_set_group.caa_record — caa_record / 102102312022 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.caa_record

<a id="canonical-2111230222101021-1320210000333020-1020303331213122-1001003201300100-0303330020020022-3231031130123220-1012230321300322-2302033111230121"></a>

Type: `"single"`. Computed.

DNSCAAResourceRecord.

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

<a id="canonical-2302023012013033-3212132231123201-0120312321202220-0032121223100201-0232202030222033-0031213000032223-1201020232010002-2000330231300302"></a>

## Direct properties — caa_record / 102102312022 / 3

<a id="canonical-3000003332023313-1123310023302312-3132100123333110-2330122120332011-2103222122330321-3211022322100131-0122033110321331-2012101200211130"></a>

<a id="canonical-3211200021203313-0330320233323123-3110331122202120-1121032321302333-3131200110103300-2320013132131110-2220013203332332-1332212132013130"></a>

## name property — caa_record / 102102312022 / 4

Type: `"string"`. Computed.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-001.md#canonical-1003010210202332-3333202023033313-3202310332302023-0113323111021000-0323333203102112-2302123311023320-0110332132200102-3322102013001203): complete subsection reference.

<a id="canonical-0122110003000322-2033330300020000-1121213032332221-1222311323102020-2220332322313223-1032320320310010-3313120000002331-1133233022120233"></a>

## Next pages — caa_record / 102102312022 / 5

- [primary.default_rr_set_group.caa_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1003010210202332-3333202023033313-3202310332302023-0113323111021000-0323333203102112-2302123311023320-0110332132200102-3322102013001203)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1003010210202332-3333202023033313-3202310332302023-0113323111021000-0323333203102112-2302123311023320-0110332132200102-3322102013001203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212102112111123-2321331110312200-1100131013322030-1323310321132101-1120110221102032-1203130130131220-1132320000110031-0212302002121100"></a>

## primary.default_rr_set_group.caa_record.values — values / 301121213203 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-2231130311212021-2231121000310210-0333021133202030-2212032233303300-2233330122101130-3300333102031111-0312302332003200-2201113110032012)
- primary.default_rr_set_group.caa_record.values

<a id="canonical-0132232112330122-3020303022121233-2123111101133121-2302010323110230-0030000122103020-1120001103323123-3010320311321223-2011213300121003"></a>

Type: `"list"`. Computed.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-3310311010331001-3220303301232230-1111330320103301-2031130331001030-3133203223310113-1113121312020002-0132031210122300-0003120130211321"></a>

## Direct properties — values / 301121213203 / 3

<a id="canonical-2323311312222001-0013120300311322-3331332101313022-3010113010130212-3203101022000313-3332211000313010-1211133333223132-1200133111131111"></a>

<a id="canonical-0231233011120100-1001111101122311-3230020001010110-0200023213130132-2323312020011131-2223323321310103-3322010023000201-2200220232222130"></a>

## flags property — values / 301121213203 / 4

Type: `"number"`. Computed.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-3310212311323301-2103133002233110-0311220022301230-3012331211003302-1120010102201103-1213323032331123-1131130222211113-3130023313233130"></a>

<a id="canonical-3000001310222022-1122012010120201-3021130311120331-3101100020300033-2230321302002132-1201320320311323-2103323310233131-0223132012201022"></a>

## tag property — values / 301121213203 / 5

Type: `"string"`. Computed.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-3300221002010112-1222221230022020-0133301132310033-0322101220101113-1013000300312121-0202011010101102-1212100222220323-1301102033110132"></a>

<a id="canonical-2210232320213333-3103201213313303-2212022322032212-0023312321102120-3003002330221313-1012311313332020-1133330032122213-0202130321222133"></a>

## value property — values / 301121213203 / 6

Type: `"string"`. Computed.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2312201320011101-3300213213203000-3330112220002202-2103233031031310-2121130202030211-1012012310020303-3033023321023010-2332302132022020"></a>

## Next pages — values / 301121213203 / 7

- [primary.default_rr_set_group.caa_record](data-sources--dns_zone--reference--group-001.md#canonical-2231130311212021-2231121000310210-0333021133202030-2212032233303300-2233330122101130-3300333102031111-0312302332003200-2201113110032012)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210223023303133-2112230302132202-3121210030020121-0100033231230101-1332123212200003-1323201010103000-3302011033021210-0032203133233301"></a>

## primary.default_rr_set_group.cds_record — cds_record / 002202301001 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.cds_record

<a id="canonical-2210130232002212-1221333233322332-3031222320301131-2210031202320332-1303221321010000-2021332020120011-3123023213220111-0223010032133231"></a>

Type: `"single"`. Computed.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

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

<a id="canonical-3331132330223233-0002022123231233-0200023030230120-1332031220111230-0223212011332131-3121013333210213-0112203112023212-3230300210321120"></a>

## Direct properties — cds_record / 002202301001 / 3

<a id="canonical-1201233312111231-1312223013101111-3000233222220130-0203330022131200-2111132003323200-2331110300033223-2031011031110210-2203031110220001"></a>

<a id="canonical-0233233230322033-2330003210121121-3031133022101200-2303122010203212-3223331233332003-3333023300033200-0132030211132031-0301310230200113"></a>

## name property — cds_record / 002202301001 / 4

Type: `"string"`. Computed.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213): complete subsection reference.

<a id="canonical-1012102003033222-1223303320032030-0031201110320122-1301303211003130-0203113011331021-1220213220331030-3203323311131303-0122132030301031"></a>

## Next pages — cds_record / 002202301001 / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020200231211201-3010220303123002-2302300310233313-0323312210131011-0302303133331210-2012213210112323-3222111001010210-1111133312120031"></a>

## primary.default_rr_set_group.cds_record.values — values / 220022032200 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220)
- primary.default_rr_set_group.cds_record.values

<a id="canonical-3103311021310220-2021121122010131-1301303121113303-1101000131030330-2222221321211310-0132203020332033-0302233303312213-3113321102110313"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3033123222211322-2201211100333320-0132032130131302-3203132200020200-0212023202013322-0100102313323301-2002030030012311-3110233212133320"></a>

## Direct properties — values / 220022032200 / 3

<a id="canonical-3233213123322010-1211232312021001-1321021030302312-2130010311131123-1230122033211101-2032310133202201-3121332323122032-0100202220221311"></a>

<a id="canonical-1313111100223012-1301100311002123-2221333011022331-2010202222103131-1022202232333201-2023332132131032-1323300331311202-3132123220023301"></a>

## ds_key_algorithm property — values / 220022032200 / 4

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1013003310003321-0233033113131202-2030122110213212-2032012121113032-2220021001032001-1302233023132021-0023230203033212-0120111303021020"></a>

<a id="canonical-0233031002311200-3320321232030300-3021000211033223-2110320203102032-2120311332030132-2003230223122110-3301320332220222-2233100303223102"></a>

## key_tag property — values / 220022032200 / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-2133022302012123-0220323023301322-3323231013023302-3212011312223012-1102220100330322-3330030001112321-2113203332012033-3023101320100111): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-2231130131103101-0130122012121130-3211003230012003-0130223100003331-3223211002130320-1332012313033213-0203122220232300-3301303321003311): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-001.md#canonical-2101320332232333-1300033130101121-1223311220002300-1021130021021331-3013020122210230-1000211022231110-2001013303211320-3121322232302331): complete subsection reference.

<a id="canonical-1030020232103303-2221000011331221-2200101131323211-0332321213331232-2213112230210133-2132133101302133-2033230101021330-0313321120333312"></a>

## Next pages — values / 220022032200 / 6

- [primary.default_rr_set_group.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-001.md#canonical-2133022302012123-0220323023301322-3323231013023302-3212011312223012-1102220100330322-3330030001112321-2113203332012033-3023101320100111)
- [primary.default_rr_set_group.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-001.md#canonical-2231130131103101-0130122012121130-3211003230012003-0130223100003331-3223211002130320-1332012313033213-0203122220232300-3301303321003311)
- [primary.default_rr_set_group.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-001.md#canonical-2101320332232333-1300033130101121-1223311220002300-1021130021021331-3013020122210230-1000211022231110-2001013303211320-3121322232302331)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2133022302012123-0220323023301322-3323231013023302-3212011312223012-1102220100330322-3330030001112321-2113203332012033-3023101320100111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101100200130301-1021310332102033-3033201100112111-3230220220021012-1320013003103111-1302331331302202-1233300033332233-1030133001122213"></a>

## primary.default_rr_set_group.cds_record.values.sha1_digest — sha1_digest / 321231230211 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220)
- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213)
- primary.default_rr_set_group.cds_record.values.sha1_digest

<a id="canonical-2103023032101032-1210032133333022-0203320020311231-2220032330311202-3002111302123031-0122002213123323-0102101321233310-3210232031201230"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-0130322201222303-0001300330100123-1331011122001030-1131113210303012-1230321102211200-0003033111233231-3303031002112332-2221331013320301"></a>

## Direct properties — sha1_digest / 321231230211 / 3

<a id="canonical-1120131122112210-0130233133221023-1202331310003101-3200323303211222-2232202233101230-1012013122122111-2112101103333222-2001111220211112"></a>

<a id="canonical-0121031031233232-3221320103303131-1203131203211020-1101011012331021-3033112103232331-0213322012213320-2320122201031100-2111003231213010"></a>

## digest property — sha1_digest / 321231230211 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-0022310122000302-3312200132323030-2130123213020332-3300011203031311-0320111301122301-2120121121222112-3230202331010310-2311233331012110"></a>

## Next pages — sha1_digest / 321231230211 / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2231130131103101-0130122012121130-3211003230012003-0130223100003331-3223211002130320-1332012313033213-0203122220232300-3301303321003311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003203212012131-3002000012311123-1120213002213131-2003112103130323-2020332302123011-1220301312131222-2202100123121103-3220221010331233"></a>

## primary.default_rr_set_group.cds_record.values.sha256_digest — sha256_digest / 020031331133 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220)
- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213)
- primary.default_rr_set_group.cds_record.values.sha256_digest

<a id="canonical-1032321100131320-0221111023312313-0232120133303312-0230001023123021-2001012000313121-1313233003020322-1112222101013232-3232100131020112"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-3312220030010111-1020203130211321-2333320130230122-3221011213031200-0202131220203322-1221133123022223-3021203102032321-3113310300232100"></a>

## Direct properties — sha256_digest / 020031331133 / 3

<a id="canonical-0330110211220103-0112233012302321-0320000301011010-0301013311202131-3102130303001111-0000023002230313-1032013222121111-3201032233202022"></a>

<a id="canonical-1111131123322101-3031233230331202-1220033113223003-1231032031330321-1100020103030310-1003310300231131-2022203112123332-2021133013001313"></a>

## digest property — sha256_digest / 020031331133 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0010020310133120-1133132313321003-3002001010300010-0301222202011312-2103300302031103-3323212323320203-1230020322212232-2130201123223110"></a>

## Next pages — sha256_digest / 020031331133 / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2101320332232333-1300033130101121-1223311220002300-1021130021021331-3013020122210230-1000211022231110-2001013303211320-3121322232302331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133010002210230-1112302011302300-0233332210212221-3221231321021320-3132332113103012-1112001013033332-2211113312031232-3103321332213100"></a>

## primary.default_rr_set_group.cds_record.values.sha384_digest — sha384_digest / 200313001313 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.cds_record](data-sources--dns_zone--reference--group-001.md#canonical-0110033330222032-0123033013030301-1302023330330311-1330232122221312-3133301300302331-0131033023022301-0230230201212031-1221311203320220)
- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213)
- primary.default_rr_set_group.cds_record.values.sha384_digest

<a id="canonical-1111012111012013-0202302223320200-0201300033122233-0333233300102011-0023220200002302-0222201302312131-3001131323213233-3010033312010232"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-1223113302101112-2312002333021011-2303131103032321-3322131101230222-3111101233200221-3101213113020322-1001120003302032-0000222302131020"></a>

## Direct properties — sha384_digest / 200313001313 / 3

<a id="canonical-0302321003303001-0122012011100300-1121323322030303-1301010313100320-2101220130023133-3132221310321100-2011333230130233-2202031120302103"></a>

<a id="canonical-2023101331223001-1132333022303201-3220323322033222-0301233233102220-3331331231110033-0233110113230211-0030021013223001-3312133210202130"></a>

## digest property — sha384_digest / 200313001313 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-3301331320111123-2230001202201101-3223132333022132-1310010122301312-2100331002222332-0120100100013110-1322000121331201-2321020333102323"></a>

## Next pages — sha384_digest / 200313001313 / 5

- [primary.default_rr_set_group.cds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1320030231302213-2013220000302201-2213201212323011-1001112012013200-2201010011202112-3030303221322101-2001002102213022-0231200030313213)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2103113313321033-0201132202120131-1011321103103120-0212131230022211-1011221201210111-1203211223103012-3213112020013312-0122002221223221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320220203010133-2101103030133130-0033111030302133-1131331113122202-1103313201212030-2223012300203331-3033102233220011-2311020023011103"></a>

## primary.default_rr_set_group.cert_record — cert_record / 300222232320 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.cert_record

<a id="canonical-2330230012123132-1221230312310033-2133111133112033-3131023230131220-0231210301103033-2211311331101331-3332221321202331-3003110220113301"></a>

Type: `"single"`. Computed.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

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

<a id="canonical-2020231111300122-2332003222112330-1031220333301001-0233032231301213-2130000110323111-2220003333101230-1003110330033301-3211221313011032"></a>

## Direct properties — cert_record / 300222232320 / 3

<a id="canonical-3312220333221031-3220130201031212-2002110032020233-1201230221201312-2000101303322310-3311330031230311-0111131011310213-3312210201113022"></a>

<a id="canonical-3213311033210112-0121132231130323-1321102320120323-0012021231122023-0222232310030011-0011212320003332-3222003111101031-3132231203322330"></a>

## name property — cert_record / 300222232320 / 4

Type: `"string"`. Computed.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-001.md#canonical-1120031120230330-2222022101310233-1110032021010231-3321311203222233-1320112010302133-3232100311322323-1103102112030310-2120132302312320): complete subsection reference.

<a id="canonical-1023301113210130-0132110111003301-0123220031103332-3200230233230123-0330333000033220-3120120231231022-0123312023201002-1220023031122311"></a>

## Next pages — cert_record / 300222232320 / 5

- [primary.default_rr_set_group.cert_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1120031120230330-2222022101310233-1110032021010231-3321311203222233-1320112010302133-3232100311322323-1103102112030310-2120132302312320)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1120031120230330-2222022101310233-1110032021010231-3321311203222233-1320112010302133-3232100311322323-1103102112030310-2120132302312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133122311230213-0121333131111320-0011203303120313-2222203311321132-0300231333131231-3200320221331330-0010022321320220-1200223211332020"></a>

## primary.default_rr_set_group.cert_record.values — values / 101323202103 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-2103113313321033-0201132202120131-1011321103103120-0212131230022211-1011221201210111-1203211223103012-3213112020013312-0122002221223221)
- primary.default_rr_set_group.cert_record.values

<a id="canonical-2300123012220120-2033212302211023-0112223130101001-0313202113112320-0123022011111033-0311122233330000-1111031311102232-3200212322130333"></a>

Type: `"list"`. Computed.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2112223032331210-2300130220233222-0011232112003003-3231211032011030-2302003102111221-3303331023211232-0012213000131322-3231322223031200"></a>

## Direct properties — values / 101323202103 / 3

<a id="canonical-2110103123313222-3133011311023130-1330323122011122-0100032000232133-2122211311121031-1231120200000132-3031230301321010-2323313000212112"></a>

<a id="canonical-3323132311112301-1302202031322230-1210003213222332-2131302230232020-2102232332210210-3030121230130201-0311030010331322-0102230001110003"></a>

## algorithm property — values / 101323202103 / 4

Type: `"string"`. Computed.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Upstream description:

CERT algorithm value must be compatible with the specified algorithm.

&#8203;- RESERVEDALGORITHM: RESERVEDALGORITHM

&#8203;- RSAMD5: RSAMD5

&#8203;- DH: DH

&#8203;- DSASHA1: DSASHA1

&#8203;- ECC: ECC

&#8203;- RSASHA1ALGORITHM: RSA-SHA1

&#8203;- INDIRECT: INDIRECT

&#8203;- PRIVATEDNS: PRIVATEDNS

&#8203;- PRIVATEOID: PRIVATEOID.

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2200300230131231-1333023020230233-0133020022321030-3033103233010112-1000022211020301-0101212100002200-1010100100013130-1100013000013031"></a>

<a id="canonical-0331323131303130-0202232203132022-2321301133113131-3001201232101030-1100202132030120-2103130111232031-1321133023103011-3133121213231301"></a>

## cert_key_tag property — values / 101323202103 / 5

Type: `"number"`. Computed.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1002200101302031-2132212233210101-1311100333232013-2333200133211232-3223303003302033-2303111213110322-2122330221223223-1222220212311011"></a>

<a id="canonical-3033313000110213-2131011002010333-3000100032113311-2010113003300302-2220003302002223-3033110321313313-3003122003310010-1023123021302301"></a>

## cert_type property — values / 101323202103 / 6

Type: `"string"`. Computed.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Upstream description:

CERT type value must be compatible with the specified types.

&#8203;- INVALIDCERTTYPE: INVALIDCERTTYPE

&#8203;- PKIX: PKIX

&#8203;- SPKI: SPKI

&#8203;- PGP: PGP

&#8203;- IPKIX: IPKIX

&#8203;- ISPKI: ISPKI

&#8203;- IPGP: IPGP

&#8203;- ACPKIX: ACPKIX

&#8203;- IACPKIX: IACPKIX

&#8203;- URI\_: URI

&#8203;- OID: OID.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122321133333333-1313311023321310-2110012323101232-1320010313330122-3203211011023331-1101001120332021-0110112311122320-3311310102211311"></a>

<a id="canonical-1321221332011323-1001120023220101-2302102221211323-0330221332303110-2210303031333130-2132101122011233-3113212313011130-0321213033303002"></a>

## certificate property — values / 101323202103 / 7

Type: `"string"`. Computed.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3332230012101230-0330022330330230-2331323323010302-3231031100013311-2331303321101322-0033100000212020-1211133023321110-2232010303132123"></a>

## Next pages — values / 101323202103 / 8

- [primary.default_rr_set_group.cert_record](data-sources--dns_zone--reference--group-001.md#canonical-2103113313321033-0201132202120131-1011321103103120-0212131230022211-1011221201210111-1203211223103012-3213112020013312-0122002221223221)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1210222313232023-3302132113200010-1102000331232222-1301112000130123-0332301130133312-0112100212202202-3031221112323220-3020332031122200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233133211212200-2131001032221103-0302213031112200-3313303321320021-2330233322201321-2333221311330032-3011022003123301-0132100001223033"></a>

## primary.default_rr_set_group.cname_record — cname_record / 220200132223 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.cname_record

<a id="canonical-2121301111310012-2001113013020011-2301022103002203-0100021133213031-2030220010101133-1201321223200231-2001321103013223-2031131302020121"></a>

Type: `"single"`. Computed.

DNSCNAMEResourceRecord.

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

<a id="canonical-2321310211122231-2112300220000021-0103111202032031-1323110101310220-1112212103320121-3020300021112230-3201200111120022-3103103330032023"></a>

## Direct properties — cname_record / 220200132223 / 3

<a id="canonical-3020113003202222-3332121000120002-0233003201323021-3031130003001032-0313022212232211-1133110132212202-0123303011210313-3131311203223010"></a>

<a id="canonical-1211211031023103-3221002300333111-3121123100220200-2221103323231332-3301001133330122-0033133012013310-0323002212221300-1210332330001312"></a>

## name property — cname_record / 220200132223 / 4

Type: `"string"`. Computed.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-1310203013003032-2033100230131202-1013032123032013-0223122123333120-1303232300121302-1111013312202312-3021302322110300-3212003110312100"></a>

<a id="canonical-2003222321233201-1322101122320232-3220000021130112-3031020203013210-0230110233303100-2211133132212312-2033020132203301-2331003113023020"></a>

## value property — cname_record / 220200132223 / 5

Type: `"string"`. Computed.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-1213213303000200-1321233130021120-2012303200112333-2223211220303130-0113001100120030-3023023202302302-2102202102131331-1203010112122211"></a>

## Next pages — cname_record / 220200132223 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133122031230201-3001301222020030-2322111231231300-0301023313300101-2103023323313100-1131003131033003-1101321221112122-2233311010313121"></a>

## primary.default_rr_set_group.ds_record — ds_record / 113310202233 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.ds_record

<a id="canonical-1211203130131120-0121001321313311-2231310210220030-0130122123010221-3022333331222000-0331300031001001-2220030322321320-0320011122133131"></a>

Type: `"single"`. Computed.

DNS DS Record. DNS DS Record.

Upstream description:

DNS DS Record.

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

<a id="canonical-1120221223113002-1331030103022333-3012220330312310-2112102022210221-3003103320010122-3122123120001001-1332333123223132-3020132023333020"></a>

## Direct properties — ds_record / 113310202233 / 3

<a id="canonical-2210113100211302-3023013321210022-3333000201102313-0132020301300332-3213330302032230-0032313022031123-2333123021331230-0220332230231321"></a>

<a id="canonical-2303002110222030-1133022203210322-3231133103123201-0023030301331231-1023012320110233-2132330012203023-3202321203110222-0303223122211231"></a>

## name property — ds_record / 113310202233 / 4

Type: `"string"`. Computed.

DS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100): complete subsection reference.

<a id="canonical-2221200010212012-3223230012113012-3221300303003310-1333220322122101-3022102020102022-0322001313010113-0023030202030321-2302122003013102"></a>

## Next pages — ds_record / 113310202233 / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331210213020210-2131231303132312-0031201003012023-3023130110213223-2223312133201120-2220010210220112-0032200302111220-1302230021001000"></a>

## primary.default_rr_set_group.ds_record.values — values / 320132011321 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100)
- primary.default_rr_set_group.ds_record.values

<a id="canonical-2000102011223330-2213001100022300-1100221211120111-0313312230222303-0203031212312221-0010011322110212-2131011000032101-3103301212212110"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2323130320231121-3103320330330332-0311021120310232-2112332012030120-1213133122020111-1210300312223011-3202103020222110-1200012021000202"></a>

## Direct properties — values / 320132011321 / 3

<a id="canonical-0122212333010013-2111301021003303-2103230323033300-3333323220330332-2121323332011232-0311313230312000-2032212021031322-0212212220312213"></a>

<a id="canonical-1302223033013132-0332221113120123-0200332302101121-1303103100023312-2331120030311123-1010213202000120-2201323321032003-0302201333323231"></a>

## ds_key_algorithm property — values / 320132011321 / 4

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2302203310103211-2223100220132232-1032020122123322-3311133002203121-0312003313110320-2223111030000202-1303300013223313-0230131332222113"></a>

<a id="canonical-0122212312301220-1112002133000303-0210203121012223-2300100300210210-3231230331022323-2201303120212122-0231213020000021-0012230200330213"></a>

## key_tag property — values / 320132011321 / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-2330211211113323-0133320000331330-3212111101203123-1220330011232233-1300320223102222-3211311200130210-1032213010030230-1222030202212313): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-1020002003030032-1130201032112031-3010331300321220-3320103121300300-0331321313303330-1310231231013013-1320221110333110-2202021210002033): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-2022232333211100-0121202021103121-2302000333012022-1232200321013202-0201321121130123-3310311330202123-2022000220113232-0111002333132313): complete subsection reference.
