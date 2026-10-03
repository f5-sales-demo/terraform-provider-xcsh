---
page_title: "xcsh_dns_load_balancer reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer reference."
---

# xcsh_dns_load_balancer reference

<a id="canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321031023320223-3000003000323210-0033022230330130-2223301132032103-1031321213021112-2011201121223330-3121320121332113-0013212212113210"></a>

## Property reference — Property reference / 322111131311 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- Property reference

<a id="canonical-0022213332201103-3233212210021333-0101011202303212-1303332002311003-3330032222330131-2022302302210101-0303132203321022-1302232030030013"></a>

## Direct properties — Property reference / 322111131311 / 3

<a id="canonical-0113323100322023-0303310101313112-2212021331311223-0013022011023232-2121000313000310-3113333123313111-3312032103323112-0030333212210311"></a>

<a id="canonical-1130230101310020-2111332312000021-2310200132321020-3232331312001130-0130201022201002-3113202211311022-0000113310220233-3231131113231012"></a>

## annotations property — Property reference / 322111131311 / 4

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

<a id="canonical-2121103030310123-2023003003233032-0213120212303231-3031330321032130-1310300030132300-3220003100112333-2032312001000210-3203130133021022"></a>

<a id="canonical-3221322310223131-0310110020033023-3120000222220312-3221223301320232-1123131223313021-0123202100231232-2312002133320311-2011323001333031"></a>

## description property — Property reference / 322111131311 / 5

Type: `"string"`. Computed.

Description of the DNSLoadBalancer.

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

- [fallback_pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-2101122021130201-3101010132313031-3002332123310323-2131103023023221-3121010320123013-2302332030132230-3211120320330301-2212011213122210): complete subsection reference.

<a id="canonical-3331331010211000-1320222030020121-1220123010220332-3011112321330133-2133221101120000-0322230133102021-2101200303110100-1212003031331222"></a>

<a id="canonical-3200221001321310-1110300031003211-0203320313210102-0111320003220022-3012333132231312-0122130202232002-0031320203020213-2011213312313202"></a>

## ID property — Property reference / 322111131311 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0223232130001111-1013302213130233-0022303103021212-2020102233112033-3331113332133302-2023111033322111-3330033032031112-2113003222231021"></a>

<a id="canonical-1313023203102023-3020310102023033-3203033130321022-3101213110012230-1323202002132011-2011001331021200-1031120013222123-2003030310121212"></a>

## labels property — Property reference / 322111131311 / 7

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

<a id="canonical-1331100202102121-0323311333233010-2021100220102132-0211303032213233-0203222111112231-0120301020013112-1121021102200013-0022311313003110"></a>

<a id="canonical-1303302211302121-3221002200323203-0130322100010212-3010211320201201-2220100032123021-3113022030313233-3103023200023333-1103122220332223"></a>

## name property — Property reference / 322111131311 / 8

Type: `"string"`. Required.

Name of the DNSLoadBalancer.

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

<a id="canonical-1310320211112120-0022203303100013-3211003303002331-3230211030131333-3322003223231301-2030223111331212-2232132020331101-1231111001203022"></a>

<a id="canonical-0303111110001231-0332000013133112-1310200103100200-3122222003312211-2230023301121131-3011123322300001-0311110133132320-3130033111311223"></a>

## namespace property — Property reference / 322111131311 / 9

Type: `"string"`. Optional, Computed.

Namespace where the DNSLoadBalancer exists.

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

<a id="canonical-1330233031330000-3132012211223202-3302133300101211-0011200132002331-3130231001313002-3033313330001300-1300202132102301-0022123111332032"></a>

<a id="canonical-3113111223121331-3222022012002123-3023323310020230-3220102030301233-2322331033001322-1030000123203100-1012001313233323-3123322232012222"></a>

## record_type property — Property reference / 322111131311 / 10

Type: `"string"`. Computed.

\[Enum: A|AAAA|MX|CNAME|SRV\] Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME -
SRV: SRV. Possible values are \`A\`, \`AAAA\`, \`MX\`, \`CNAME\`, \`SRV\`. Defaults to \`A\`.

Upstream description:

Resource Record Type

&#8203;- A: A

&#8203;- AAAA: AAAA

&#8203;- MX: MX

&#8203;- CNAME: CNAME

&#8203;- SRV: SRV.

Receipt-pinned upstream constraints:

```json
{
  "default": "A",
  "enum": [
    "A",
    "AAAA",
    "MX",
    "CNAME",
    "SRV"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120): complete subsection reference.

- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223): complete subsection reference.

<a id="canonical-1033111200122202-1332220023013202-3331220313021010-2212020022101210-1030332102012131-3310222312230002-1112322013132331-1121002111230333"></a>

## All schema paths — Property reference / 322111131311 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_load_balancer--reference--group-001.md#canonical-0113323100322023-0303310101313112-2212021331311223-0013022011023232-2121000313000310-3113333123313111-3312032103323112-0030333212210311) |
| `description` | [description](data-sources--dns_load_balancer--reference--group-001.md#canonical-2121103030310123-2023003003233032-0213120212303231-3031330321032130-1310300030132300-3220003100112333-2032312001000210-3203130133021022) |
| `fallback_pool` | [fallback_pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-3220301101022132-3012333212023100-1233312332121133-3112331200223032-3230102313002013-3211321121333032-1102210113013000-0121100232302200) |
| `fallback_pool.name` | [fallback_pool.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-0203300123200230-3123010030012133-1020223231222203-2101022021120213-2301223102113230-1301311323111213-2213210011011233-0321113210121112) |
| `fallback_pool.namespace` | [fallback_pool.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-0221210231003123-1100000322221023-3120313301001123-0311102100103210-0213123013221201-3200220010230023-1113010223133101-1212011121013210) |
| `fallback_pool.tenant` | [fallback_pool.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-1301222322221031-1303003232201003-2022130211132200-0031201111030003-1102222213133303-3303332322212233-1112301101123323-2213323033321323) |
| `id` | [ID](data-sources--dns_load_balancer--reference--group-001.md#canonical-3331331010211000-1320222030020121-1220123010220332-3011112321330133-2133221101120000-0322230133102021-2101200303110100-1212003031331222) |
| `labels` | [labels](data-sources--dns_load_balancer--reference--group-001.md#canonical-0223232130001111-1013302213130233-0022303103021212-2020102233112033-3331113332133302-2023111033322111-3330033032031112-2113003222231021) |
| `name` | [name](data-sources--dns_load_balancer--reference--group-001.md#canonical-1331100202102121-0323311333233010-2021100220102132-0211303032213233-0203222111112231-0120301020013112-1121021102200013-0022311313003110) |
| `namespace` | [namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-1310320211112120-0022203303100013-3211003303002331-3230211030131333-3322003223231301-2030223111331212-2232132020331101-1231111001203022) |
| `record_type` | [record_type](data-sources--dns_load_balancer--reference--group-001.md#canonical-1330233031330000-3132012211223202-3302133300101211-0011200132002331-3130231001313002-3033313330001300-1300202132102301-0022123111332032) |
| `response_cache` | [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1330333221222111-3232030220331332-3222232223133321-3020332313011130-1323001323213231-3021211102132211-0231201111212320-1212030023002021) |
| `response_cache.default_response_cache_parameters` | [response_cache.default_response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-1213033122130021-1221130220103300-0122103213001320-0313211133033121-0230310010030010-3300332000030112-2123213201101213-1221001112101112) |
| `response_cache.disable_spec` | [response_cache.disable_spec](data-sources--dns_load_balancer--reference--group-001.md#canonical-2113110132101110-3122301113313311-3303000310201211-1001221110001211-1222313033011330-2020222200031020-3123303100222003-1011231311012130) |
| `response_cache.response_cache_parameters` | [response_cache.response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-3120131213333022-2123032212131002-0223002212233023-1011220310133102-0321311030132003-0120201320110103-0010103302012120-2020030111202111) |
| `response_cache.response_cache_parameters.cache_cidr_ipv4` | [response_cache.response_cache_parameters.cache_cidr_ipv4](data-sources--dns_load_balancer--reference--group-001.md#canonical-3310013020321022-0020123011121220-1102303103130010-3013200013103233-2321321303201230-0031130332221200-0300303121303300-0132300002211203) |
| `response_cache.response_cache_parameters.cache_cidr_ipv6` | [response_cache.response_cache_parameters.cache_cidr_ipv6](data-sources--dns_load_balancer--reference--group-001.md#canonical-3031122232010011-0122232020201012-1302020321233302-3300000231230203-3301303301031212-0013132232311330-1322330310110302-0321112312122330) |
| `response_cache.response_cache_parameters.cache_ttl` | [response_cache.response_cache_parameters.cache_ttl](data-sources--dns_load_balancer--reference--group-001.md#canonical-3312122223233212-1030101122120200-3000201231122311-3210202320330310-3213332101303002-0010103102220233-3313030312100211-1302301133312210) |
| `rule_list` | [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-0031011301121012-1100012312031122-1013223321101300-1103000002211233-1213022130220323-0101103122310311-0130010122323311-1102033333200312) |
| `rule_list.rules` | [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-1203303211010203-3012113301022131-3021302200231302-3111000321202330-3101311030300213-2303311322133103-2211213002012222-1303301031010012) |
| `rule_list.rules.asn_list` | [rule_list.rules.asn_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-3230313103313102-3032223131233020-3031202232133232-1030230321233002-0130102310010322-0021020213202133-1003101020322331-0320002201113001) |
| `rule_list.rules.asn_list.as_numbers` | [rule_list.rules.asn_list.as_numbers](data-sources--dns_load_balancer--reference--group-001.md#canonical-0003212012202021-3122012113323132-0101302313222030-1200003133033133-1011102021112202-1301312001133312-1132221233112213-0203231132300120) |
| `rule_list.rules.asn_matcher` | [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-2312202013101110-1110101312320131-0113333121101132-1130032130123230-3222202222310003-0000330002122013-1200001021212302-1330333331001310) |
| `rule_list.rules.asn_matcher.asn_sets` | [rule_list.rules.asn_matcher.asn_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-2010022100021130-0033223230031323-2321211200232201-2210321011121303-3122122331111023-3300002331012300-2312312331212200-0110002002202133) |
| `rule_list.rules.asn_matcher.asn_sets.kind` | [rule_list.rules.asn_matcher.asn_sets.kind](data-sources--dns_load_balancer--reference--group-001.md#canonical-3010113311002011-0133323220202121-0131012013331311-3000021233001021-2133123102222023-3232211223301121-1203032230031322-2213223223220302) |
| `rule_list.rules.asn_matcher.asn_sets.name` | [rule_list.rules.asn_matcher.asn_sets.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-1331323002231002-3011101203113302-3332100000010331-1301003030131200-0110020310220213-0333221000233333-0012220123030232-3202121011100202) |
| `rule_list.rules.asn_matcher.asn_sets.namespace` | [rule_list.rules.asn_matcher.asn_sets.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-0330301202201130-2120202111022021-3312133030311302-1221323132122122-3003023230103001-0210110322332303-1112231133031031-0221301030013001) |
| `rule_list.rules.asn_matcher.asn_sets.tenant` | [rule_list.rules.asn_matcher.asn_sets.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-0300302310003110-2330202003112110-2130202030002002-2301113323331300-1003331020022322-0303011001000321-2321021212130230-0232031211202002) |
| `rule_list.rules.asn_matcher.asn_sets.uid` | [rule_list.rules.asn_matcher.asn_sets.uid](data-sources--dns_load_balancer--reference--group-001.md#canonical-2331012220310101-0031233222221230-0330233121130103-0011331011012023-2311102211323213-3330010221131110-0222123200322211-3112002121201131) |
| `rule_list.rules.geo_location_label_selector` | [rule_list.rules.geo_location_label_selector](data-sources--dns_load_balancer--reference--group-001.md#canonical-3330211011013301-1131121002231310-2212010200202332-1022300030001033-0330021310310201-2333130320200113-3210130102303311-3310311120133010) |
| `rule_list.rules.geo_location_label_selector.expressions` | [rule_list.rules.geo_location_label_selector.expressions](data-sources--dns_load_balancer--reference--group-001.md#canonical-0213311101313310-2213111301320203-0232302311213003-0331332033102231-3220322003331121-1003132112120321-2021330123003212-1023212120210112) |
| `rule_list.rules.geo_location_set` | [rule_list.rules.geo_location_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-1201330303301311-3100030231101131-1221313020021121-2202310022011223-3002000101302213-1300032233212111-0123101321303021-0012201332201021) |
| `rule_list.rules.geo_location_set.name` | [rule_list.rules.geo_location_set.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-0213011032233031-3220202231000232-2011300222012030-3101130201122300-3102311111032110-1132332300120002-2022333112332231-1003320222023202) |
| `rule_list.rules.geo_location_set.namespace` | [rule_list.rules.geo_location_set.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-3010300132232120-1030002303301101-0000000213020303-1131123313021021-0021003210201211-2032103123203103-0120021011133313-1131303133233012) |
| `rule_list.rules.geo_location_set.tenant` | [rule_list.rules.geo_location_set.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-2233033233331201-2030213223310230-0020320021011320-1012313222023201-1010301000232300-2010120202123200-3330031121212222-1203120300033223) |
| `rule_list.rules.ip_prefix_list` | [rule_list.rules.ip_prefix_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-2113322000332030-2232332012311333-2030100212321110-0003020211212010-0233011031121122-1100003111221023-3120033211201200-3131100010210220) |
| `rule_list.rules.ip_prefix_list.invert_match` | [rule_list.rules.ip_prefix_list.invert_match](data-sources--dns_load_balancer--reference--group-001.md#canonical-0011100210231222-1033313312211323-2331311030121322-0110201320303111-1232231131101003-1030123301112323-2130333232203212-0132212322310230) |
| `rule_list.rules.ip_prefix_list.ip_prefixes` | [rule_list.rules.ip_prefix_list.ip_prefixes](data-sources--dns_load_balancer--reference--group-001.md#canonical-0031022123000233-0310221002330012-2012021110121103-0223230332331121-1011131302021012-0131100021221032-1123331303002021-0123300302110110) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-3332033131113033-3223220211023121-2323003020131321-3131021310322311-2331111332313123-1231130313232013-3333300120212213-0003202233301301) |
| `rule_list.rules.ip_prefix_set.invert_matcher` | [rule_list.rules.ip_prefix_set.invert_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022203121321033-0323013030113101-0123202121220101-2210030023222333-2221132200013333-3222322033202211-2312233300003313-3332131120213212) |
| `rule_list.rules.ip_prefix_set.prefix_sets` | [rule_list.rules.ip_prefix_set.prefix_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-0111330220322332-2222122313210332-3213233121202112-1202223102220232-2023223033203013-2023030120001000-0113030110312331-3113311311013232) |
| `rule_list.rules.ip_prefix_set.prefix_sets.kind` | [rule_list.rules.ip_prefix_set.prefix_sets.kind](data-sources--dns_load_balancer--reference--group-001.md#canonical-3122321030312102-0311210231300230-3331002231133133-3321312112323022-2110320203322110-2112100102201103-3121221300110012-0113122313332102) |
| `rule_list.rules.ip_prefix_set.prefix_sets.name` | [rule_list.rules.ip_prefix_set.prefix_sets.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-2031022303111023-3231200021210210-1132001222300213-2323013013133223-3212232303302313-2110100010002012-0001310323231132-3310212011112321) |
| `rule_list.rules.ip_prefix_set.prefix_sets.namespace` | [rule_list.rules.ip_prefix_set.prefix_sets.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-0030133031021000-0213210113331200-2221013301032021-1131300301112233-3313313111110311-1330030221211113-3210232300330212-0312132122312213) |
| `rule_list.rules.ip_prefix_set.prefix_sets.tenant` | [rule_list.rules.ip_prefix_set.prefix_sets.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-1032221033221000-0103303021101113-2203311312010131-3223003213312220-3323012323223232-0011111323310322-3111202032021233-0120112011311033) |
| `rule_list.rules.ip_prefix_set.prefix_sets.uid` | [rule_list.rules.ip_prefix_set.prefix_sets.uid](data-sources--dns_load_balancer--reference--group-001.md#canonical-2021132002020222-2131002101030031-1102213311100123-1131331122201021-3231123303101230-1122333201230330-0230031203133333-1322122330310012) |
| `rule_list.rules.pool` | [rule_list.rules.pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-0213110333030222-0112330213022021-0122113000132331-3213123302101221-3001313210133231-3122122001012122-3313210332211003-1001201123030020) |
| `rule_list.rules.pool.name` | [rule_list.rules.pool.name](data-sources--dns_load_balancer--reference--group-001.md#canonical-0210213210123102-1220133333130031-3133220301103221-2032032202321331-3200230321213230-0222021201112223-2013010122003130-2122212033323223) |
| `rule_list.rules.pool.namespace` | [rule_list.rules.pool.namespace](data-sources--dns_load_balancer--reference--group-001.md#canonical-3011000131120303-0331010021022221-3200201331010231-1012111213102000-3022102002120301-3231230032211320-2111020332320233-3020321020011210) |
| `rule_list.rules.pool.tenant` | [rule_list.rules.pool.tenant](data-sources--dns_load_balancer--reference--group-001.md#canonical-1213010322112020-3131302032032023-1223003300121122-3131112130102031-0003010201231322-3202222323033131-1323112030331212-2202220221000011) |
| `rule_list.rules.score` | [rule_list.rules.score](data-sources--dns_load_balancer--reference--group-001.md#canonical-2022020110311001-2202132133211221-0320320212022122-3131230133203210-0320032011023203-2213032322022203-2300311001133013-3223002001232011) |

<a id="canonical-1121311211300323-0101032033020102-2132312330102203-2123212111211030-0003221210203110-0223033231132103-0012113231220113-3220232331022210"></a>

## Next pages — Property reference / 322111131311 / 12

- [fallback_pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-2101122021130201-3101010132313031-3002332123310323-2131103023023221-3121010320123013-2302332030132230-3211120320330301-2212011213122210)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-2101122021130201-3101010132313031-3002332123310323-2131103023023221-3121010320123013-2302332030132230-3211120320330301-2212011213122210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021003220102313-3133131001001133-1123033103230020-1201013230000211-1131301312331211-2121022321013233-1031031000221033-0222203133231113"></a>

## fallback_pool — fallback_pool / 223322331132 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- fallback_pool

<a id="canonical-3220301101022132-3012333212023100-1233312332121133-3112331200223032-3230102313002013-3211321121333032-1102210113013000-0121100232302200"></a>

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

<a id="canonical-1322221230100330-2003213320230133-1032010113110120-0331210000330323-0322313013201021-3110020303222233-1032102032213322-2311031322230211"></a>

## Direct properties — fallback_pool / 223322331132 / 3

<a id="canonical-0203300123200230-3123010030012133-1020223231222203-2101022021120213-2301223102113230-1301311323111213-2213210011011233-0321113210121112"></a>

<a id="canonical-0323011110033233-0302011131330312-3120211202003312-0131123131112010-0113033110002303-3222211211201000-1300101122110122-2321100030210001"></a>

## name property — fallback_pool / 223322331132 / 4

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

<a id="canonical-0221210231003123-1100000322221023-3120313301001123-0311102100103210-0213123013221201-3200220010230023-1113010223133101-1212011121013210"></a>

<a id="canonical-0102011011202001-1001113223101233-3232303120100011-3023130011033033-2022023330113103-2232022202322110-2330130112231231-0001020101203333"></a>

## namespace property — fallback_pool / 223322331132 / 5

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

<a id="canonical-1301222322221031-1303003232201003-2022130211132200-0031201111030003-1102222213133303-3303332322212233-1112301101123323-2213323033321323"></a>

<a id="canonical-3220332013113202-1110022331333110-0013333000121201-2331312113221303-3012131020010001-2121301010211132-1112202211230302-3320302030323112"></a>

## tenant property — fallback_pool / 223322331132 / 6

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

<a id="canonical-0103313011200023-1000202032300101-1100333322211110-2301033012112033-0031200032032130-2210320100303100-2003001021313001-0211022100312020"></a>

## Next pages — fallback_pool / 223322331132 / 7

- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213322200333232-2322031100031332-1323203210003122-3231220302212230-2103303101030031-2232231320232032-0002133012100132-0203301200203210"></a>

## response_cache — response_cache / 232100233322 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- response_cache

<a id="canonical-1330333221222111-3232030220331332-3222232223133321-3020332313011130-1323001323213231-3021211102132211-0231201111212320-1212030023002021"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache.

Upstream description:

Response Cache x-required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

<a id="canonical-1333123032023020-3002231222100200-1323322112201120-1120003121313211-2222001131021312-0202123123020023-3222012213232033-1232022033101301"></a>

## Direct properties — response_cache / 232100233322 / 3

- [default_response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-0033011311220122-3232113303221303-1012003100001323-2223031022220221-3231103130102203-1333022100222212-3213121300313113-3331221303212220): complete subsection reference.

- [disable_spec](data-sources--dns_load_balancer--reference--group-001.md#canonical-1311111011031300-0030203033032300-2323132310331303-1230231021213210-2302203123233310-1232321103130310-2202022221311330-1102130332221133): complete subsection reference.

- [response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-3322233121222221-1232020102211000-0003211203101023-0122200112202212-3100313113023302-2130121120101123-3023221303021123-2101303102300033): complete subsection reference.

<a id="canonical-0220232202012133-1001132302322031-1103121303300003-0203312030123313-2211220210122101-1332212201022113-1201331311332302-1331132202213031"></a>

## Next pages — response_cache / 232100233322 / 4

- [response_cache.default_response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-0033011311220122-3232113303221303-1012003100001323-2223031022220221-3231103130102203-1333022100222212-3213121300313113-3331221303212220)
- [response_cache.disable_spec](data-sources--dns_load_balancer--reference--group-001.md#canonical-1311111011031300-0030203033032300-2323132310331303-1230231021213210-2302203123233310-1232321103130310-2202022221311330-1102130332221133)
- [response_cache.response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-3322233121222221-1232020102211000-0003211203101023-0122200112202212-3100313113023302-2130121120101123-3023221303021123-2101303102300033)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-0033011311220122-3232113303221303-1012003100001323-2223031022220221-3231103130102203-1333022100222212-3213121300313113-3331221303212220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203331201122113-0333211303301031-1231021022223200-3212321010333112-3300101220002011-2231110212110303-2002121221230202-1233220123012311"></a>

## response_cache.default_response_cache_parameters — default_response_cache_parameters / 300122110222 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- response_cache.default_response_cache_parameters

<a id="canonical-1213033122130021-1221130220103300-0122103213001320-0313211133033121-0230310010030010-3300332000030112-2123213201101213-1221001112101112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default response cache parameters.

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

<a id="canonical-1122131300021021-1120033321112232-1232012202131200-3231021032020021-0001012132012322-2100031232311201-1221320003112230-0000133202230111"></a>

## Direct properties — default_response_cache_parameters / 300122110222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332322123023233-3113213331211301-3233213212312003-1130330033330010-2000010221221000-3112310223112221-0033103211330210-0101313201303123"></a>

## Next pages — default_response_cache_parameters / 300122110222 / 4

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-1311111011031300-0030203033032300-2323132310331303-1230231021213210-2302203123233310-1232321103130310-2202022221311330-1102130332221133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002222111103330-3033130313302303-3013013330322312-2133133022021203-3230220203311323-0320012310310320-1133300013011220-3220202202032112"></a>

## response_cache.disable_spec — disable_spec / 031113132030 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- response_cache.disable_spec

<a id="canonical-2113110132101110-3122301113313311-3303000310201211-1001221110001211-1222313033011330-2020222200031020-3123303100222003-1011231311012130"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-1012102102231301-0301130112001100-1121133302133123-2230113001003303-1011111211230033-3110200021210222-2200102031113120-1113003232112001"></a>

## Direct properties — disable_spec / 031113132030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023331121212113-3320021130233303-3032223320030020-1320112330010130-3301013302020031-3231123231100230-2122010333300003-1223013123302113"></a>

## Next pages — disable_spec / 031113132030 / 4

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-3322233121222221-1232020102211000-0003211203101023-0122200112202212-3100313113023302-2130121120101123-3023221303021123-2101303102300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311201213223222-0122011213103213-2221201011130022-3312002113022031-2222332212220022-1013303111132233-3002013021012200-3300221023101023"></a>

## response_cache.response_cache_parameters — response_cache_parameters / 013222233233 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- response_cache.response_cache_parameters

<a id="canonical-3120131213333022-2123032212131002-0223002212233023-1011220310133102-0321311030132003-0120201320110103-0010103302012120-2020030111202111"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache parameters.

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

<a id="canonical-1021222302220320-1030102023330013-2301223003233011-2303113100301202-3132020131223310-0300031312301103-2002023021121321-0133201112111230"></a>

## Direct properties — response_cache_parameters / 013222233233 / 3

<a id="canonical-3310013020321022-0020123011121220-1102303103130010-3013200013103233-2321321303201230-0031130332221200-0300303121303300-0132300002211203"></a>

<a id="canonical-2033002132322013-0121221110010332-0110131232222221-3321102332213223-3322011301000113-3313122013230331-1223121212023323-1333100100331120"></a>

## cache_cidr_ipv4 property — response_cache_parameters / 013222233233 / 4

Type: `"number"`. Computed.

Length of CIDR masks used to group IPv4 clients.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3031122232010011-0122232020201012-1302020321233302-3300000231230203-3301303301031212-0013132232311330-1322330310110302-0321112312122330"></a>

<a id="canonical-2321020211312210-1103113122120322-2211102211233111-2322331001223213-2130001210303032-0321311023313110-0132120013101121-2311131023203201"></a>

## cache_cidr_ipv6 property — response_cache_parameters / 013222233233 / 5

Type: `"number"`. Computed.

Length of CIDR masks used to group IPv6 clients.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-3312122223233212-1030101122120200-3000201231122311-3210202320330310-3213332101303002-0010103102220233-3313030312100211-1302301133312210"></a>

<a id="canonical-2320010230321213-2320020030121101-3310133002122122-0230111331310020-0011323212333111-3103203000120221-3120302211322020-0323113311020001"></a>

## cache_ttl property — response_cache_parameters / 013222233233 / 6

Type: `"number"`. Computed.

TTL. TTL for response cache.

Upstream description:

TTL for response cache.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

<a id="canonical-3103323120212102-0302333210311303-2213111001001200-0021322120202100-3331212211211202-2313013323122103-2330031022233321-2313312133033001"></a>

## Next pages — response_cache_parameters / 013222233233 / 7

- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201323033330310-2213220321032203-3310113011333202-1203303032001230-0312103213331333-2030102310323021-1011310311311302-2311311000101223"></a>

## rule_list — rule_list / 100220132231 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- rule_list

<a id="canonical-0031011301121012-1100012312031122-1013223321101300-1103000002211233-1213022130220323-0101103122310311-0130010122323311-1102033333200312"></a>

Type: `"single"`. Computed.

Load Balancing Rule List. List of the Load Balancing Rules.

Upstream description:

List of the Load Balancing Rules.

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

<a id="canonical-2212333210102100-1012030303200322-1313033220130323-3013122122310131-1001033030123031-2013211321130201-2133330121010123-2031231220213221"></a>

## Direct properties — rule_list / 100220132231 / 3

- [rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103): complete subsection reference.

<a id="canonical-0100330330322332-0210013012330001-2020003112300222-1220021211211010-3100331320312320-3311321231130112-1133120032301121-0010031322120013"></a>

## Next pages — rule_list / 100220132231 / 4

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302122221122012-0030300231320203-1200211220132013-0133313300330123-3330021321023102-2120200330320302-0133210110220131-3132232313232231"></a>

## rule_list.rules — rules / 110332002113 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- rule_list.rules

<a id="canonical-1203303211010203-3012113301022131-3021302200231302-3111000321202330-3101311030300213-2303311322133103-2211213002012222-1303301031010012"></a>

Type: `"list"`. Computed.

Load Balancing Rules. Rules to perform load balancing.

Upstream description:

Rules to perform load balancing.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1032320331102002-3321233122011011-0030121311211301-1110312013211220-0020323011313330-1201323010033323-1100102010231302-2013113312212001"></a>

## Direct properties — rules / 110332002113 / 3

- [asn_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-0312303233022323-1130031221213021-3301322221331113-0311211111010203-0112310101121021-3202033232231232-1121233300123113-0223010110032323): complete subsection reference.

- [asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021): complete subsection reference.

- [geo_location_label_selector](data-sources--dns_load_balancer--reference--group-001.md#canonical-1223222103223132-3310031101003102-0100001031302110-1102220231301200-3000223211112000-0220131220011321-1323310100101033-2121133221233230): complete subsection reference.

- [geo_location_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-2101001211013222-3310330203000121-0322202123032012-2100232331031322-3310013112333132-0002120232000332-3030202231102233-2111133002332002): complete subsection reference.

- [ip_prefix_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1032323121031133-1120303221210112-1202030322120300-0010021123012330-1003020210000013-0102110010100233-1112000203320013-1003331201230022): complete subsection reference.

- [ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031): complete subsection reference.

- [pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-0112111313220331-1330313331010211-0132030000331111-0010130323332331-2312213111123132-0033311200200110-2322101310110300-1302122122212323): complete subsection reference.

<a id="canonical-2022020110311001-2202132133211221-0320320212022122-3131230133203210-0320032011023203-2213032322022203-2300311001133013-3223002001232011"></a>

<a id="canonical-2232111321321112-1111203113030132-1033010333022101-2302210222023203-2202201213332333-3110112333200133-0212322322312222-0120010312232202"></a>

## score property — rules / 110332002113 / 4

Type: `"number"`. Computed.

When multiple load balancing rules match a query, the one with the highest score is chosen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32767,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32767"
  }
}
```

<a id="canonical-1001333103130322-0030101233130322-2301322030021033-2001132120032210-3212100330312223-3010000223110133-2030202020120232-3030211021213122"></a>

## Next pages — rules / 110332002113 / 5

- [rule_list.rules.asn_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-0312303233022323-1130031221213021-3301322221331113-0311211111010203-0112310101121021-3202033232231232-1121233300123113-0223010110032323)
- [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021)
- [rule_list.rules.geo_location_label_selector](data-sources--dns_load_balancer--reference--group-001.md#canonical-1223222103223132-3310031101003102-0100001031302110-1102220231301200-3000223211112000-0220131220011321-1323310100101033-2121133221233230)
- [rule_list.rules.geo_location_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-2101001211013222-3310330203000121-0322202123032012-2100232331031322-3310013112333132-0002120232000332-3030202231102233-2111133002332002)
- [rule_list.rules.ip_prefix_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1032323121031133-1120303221210112-1202030322120300-0010021123012330-1003020210000013-0102110010100233-1112000203320013-1003331201230022)
- [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031)
- [rule_list.rules.pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-0112111313220331-1330313331010211-0132030000331111-0010130323332331-2312213111123132-0033311200200110-2322101310110300-1302122122212323)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-0312303233022323-1130031221213021-3301322221331113-0311211111010203-0112310101121021-3202033232231232-1121233300123113-0223010110032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230331200031302-1102211003121011-1120013131231110-0303131121330102-0332032231001000-2233120103123330-0202023303022032-2032013102111111"></a>

## rule_list.rules.asn_list — asn_list / 001322312331 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.asn_list

<a id="canonical-3230313103313102-3032223131233020-3031202232133232-1030230321233002-0130102310010322-0021020213202133-1003101020322331-0320002201113001"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-3101323032331300-2110113121100233-0100132001111021-3000333130123132-2103330110010023-3303301033310030-1013132001311122-3300001230020210"></a>

## Direct properties — asn_list / 001322312331 / 3

<a id="canonical-0003212012202021-3122012113323132-0101302313222030-1200003133033133-1011102021112202-1301312001133312-1132221233112213-0203231132300120"></a>

<a id="canonical-1322023313132001-1023002231313130-0031323333131211-1222230210002130-0202302302300000-2101132202010030-2200122311000321-0211301231223221"></a>

## as_numbers property — asn_list / 001322312331 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1222001100111013-2333011132302033-0030313010030032-3130012121320210-0122002202321121-3121020213322112-1323010323020321-2220220011112033"></a>

## Next pages — asn_list / 001322312331 / 5

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131320210233310-3100033110301203-1030210320121321-3202222231023213-1223220230110313-0221031031131310-1001010102010013-1332212100311223"></a>

## rule_list.rules.asn_matcher — asn_matcher / 232020332133 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.asn_matcher

<a id="canonical-2312202013101110-1110101312320131-0113333121101132-1130032130123230-3222202222310003-0000330002122013-1200001021212302-1330333331001310"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-3211231313333313-0000020120302012-2010133220200223-3030313021212301-3230011331012120-2132033003333303-0100011311030131-3212122233302320"></a>

## Direct properties — asn_matcher / 232020332133 / 3

- [asn_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-3310030230132020-2301000101023212-1301221101111332-3001023233033012-3013011031220310-2103312302131332-3313003230221310-2030122300001220): complete subsection reference.

<a id="canonical-2023013030122011-3120132302213323-1123301211020032-3212212013202101-0012303230123033-1232101122121200-2310331000032131-0111101322131303"></a>

## Next pages — asn_matcher / 232020332133 / 4

- [rule_list.rules.asn_matcher.asn_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-3310030230132020-2301000101023212-1301221101111332-3001023233033012-3013011031220310-2103312302131332-3313003230221310-2030122300001220)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-3310030230132020-2301000101023212-1301221101111332-3001023233033012-3013011031220310-2103312302131332-3313003230221310-2030122300001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201120112132113-2310202120211221-0210320023303022-1311031212332131-3330301011321031-0021132201331323-2031203232321203-3230102030310220"></a>

## rule_list.rules.asn_matcher.asn_sets — asn_sets / 022232213131 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021)
- rule_list.rules.asn_matcher.asn_sets

<a id="canonical-2010022100021130-0033223230031323-2321211200232201-2210321011121303-3122122331111023-3300002331012300-2312312331212200-0110002002202133"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1323300231013323-3210021003303032-3100230123232003-1022301221200310-1210001033301011-1221023322113230-3301220221300203-0201101002301313"></a>

## Direct properties — asn_sets / 022232213131 / 3

<a id="canonical-3010113311002011-0133323220202121-0131012013331311-3000021233001021-2133123102222023-3232211223301121-1203032230031322-2213223223220302"></a>

<a id="canonical-2032312301001130-0210013030020313-3211130333002303-2232023233022021-1002230321313103-1121021132022231-0230010010323313-1013101320002012"></a>

## kind property — asn_sets / 022232213131 / 4

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

<a id="canonical-1331323002231002-3011101203113302-3332100000010331-1301003030131200-0110020310220213-0333221000233333-0012220123030232-3202121011100202"></a>

<a id="canonical-0321313321022321-1101210302032133-2201311101232010-0303203031002231-1200133120031102-0012212331303022-2201212020211300-1312222120120023"></a>

## name property — asn_sets / 022232213131 / 5

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

<a id="canonical-0330301202201130-2120202111022021-3312133030311302-1221323132122122-3003023230103001-0210110322332303-1112231133031031-0221301030013001"></a>

<a id="canonical-2333010330101030-1020111131323132-2202332213203010-3031102103322311-3013233132200322-3223120031300122-2203012213202332-2303233230330122"></a>

## namespace property — asn_sets / 022232213131 / 6

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

<a id="canonical-0300302310003110-2330202003112110-2130202030002002-2301113323331300-1003331020022322-0303011001000321-2321021212130230-0232031211202002"></a>

<a id="canonical-0033202302113131-1030031311322203-1222120332310323-1302301210133002-3130310002230330-2323110022111032-1212021330202033-0203220010311310"></a>

## tenant property — asn_sets / 022232213131 / 7

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

<a id="canonical-2331012220310101-0031233222221230-0330233121130103-0011331011012023-2311102211323213-3330010221131110-0222123200322211-3112002121201131"></a>

<a id="canonical-3212132333220100-0300101131000312-2320331011011121-2111111333022220-3320301211202011-1011322333132200-1013022212222132-0101220020032223"></a>

## uid property — asn_sets / 022232213131 / 8

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

<a id="canonical-0230101012111230-3022201031002121-3022012013011210-2222132332303300-3003001012321303-3301300302111101-2300012103301323-0212203121323221"></a>

## Next pages — asn_sets / 022232213131 / 9

- [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-1223222103223132-3310031101003102-0100001031302110-1102220231301200-3000223211112000-0220131220011321-1323310100101033-2121133221233230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200023312332101-2131312232001303-0223101222232022-0103001012133331-3103101220131013-3133232033100122-2212313113323030-0220233322112010"></a>

## rule_list.rules.geo_location_label_selector — geo_location_label_selector / 320123002230 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.geo_location_label_selector

<a id="canonical-3330211011013301-1131121002231310-2212010200202332-1022300030001033-0330021310310201-2333130320200113-3210130102303311-3310311120133010"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-1123001022120211-0222202232221320-3001211100211111-0013133312331110-3011120130001231-1330330002302303-3030233303130230-3322000312000120"></a>

## Direct properties — geo_location_label_selector / 320123002230 / 3

<a id="canonical-0213311101313310-2213111301320203-0232302311213003-0331332033102231-3220322003331121-1003132112120321-2021330123003212-1023212120210112"></a>

<a id="canonical-1130120221332210-0021012020231012-2333123222031300-2213333222330200-1131220201130101-0113013123003313-1203333323011122-3113130322020222"></a>

## expressions property — geo_location_label_selector / 320123002230 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1311031130313323-2222213331203001-3120220003312030-3200031331333232-3032313313331313-1200023303012100-3220201312321210-0311131010013121"></a>

## Next pages — geo_location_label_selector / 320123002230 / 5

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-2101001211013222-3310330203000121-0322202123032012-2100232331031322-3310013112333132-0002120232000332-3030202231102233-2111133002332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103121110130222-3331323003231201-2330232101300203-0330131013222031-3023222212110331-3220002333003022-1311123030021023-1003010320200000"></a>

## rule_list.rules.geo_location_set — geo_location_set / 000320211233 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.geo_location_set

<a id="canonical-1201330303301311-3100030231101131-1221313020021121-2202310022011223-3002000101302213-1300032233212111-0123101321303021-0012201332201021"></a>

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

<a id="canonical-0103013222110001-0221030020113333-0123321201302300-1223020102131001-0110023012231031-1310012203022323-3202210222100132-1220312330131131"></a>

## Direct properties — geo_location_set / 000320211233 / 3

<a id="canonical-0213011032233031-3220202231000232-2011300222012030-3101130201122300-3102311111032110-1132332300120002-2022333112332231-1003320222023202"></a>

<a id="canonical-3323211313222213-3130323230233022-1333111212032130-0100200033303322-1203211212103222-3330101310022303-0013022120101020-2133111001332330"></a>

## name property — geo_location_set / 000320211233 / 4

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

<a id="canonical-3010300132232120-1030002303301101-0000000213020303-1131123313021021-0021003210201211-2032103123203103-0120021011133313-1131303133233012"></a>

<a id="canonical-1003320131320313-0002111332232111-0110033000010300-0033303121313312-2303131303133222-2033310223212213-3203310231110103-1213123122311113"></a>

## namespace property — geo_location_set / 000320211233 / 5

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

<a id="canonical-2233033233331201-2030213223310230-0020320021011320-1012313222023201-1010301000232300-2010120202123200-3330031121212222-1203120300033223"></a>

<a id="canonical-3003123310330330-0312030001210110-0333302112322233-1102231032002220-2212311233201102-2231331033213131-3213000003311122-2202012300020013"></a>

## tenant property — geo_location_set / 000320211233 / 6

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

<a id="canonical-0120213113223231-2210101122230120-2022232021023213-3021211302220030-1013112113321201-3122323321332032-1133213020201132-3112000020222223"></a>

## Next pages — geo_location_set / 000320211233 / 7

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-1032323121031133-1120303221210112-1202030322120300-0010021123012330-1003020210000013-0102110010100233-1112000203320013-1003331201230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101000230202101-2020101031301121-1213230132213223-2223112203022132-3312303101302031-3103033001232303-1330203020333101-2010123110212300"></a>

## rule_list.rules.ip_prefix_list — ip_prefix_list / 332212112131 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.ip_prefix_list

<a id="canonical-2113322000332030-2232332012311333-2030100212321110-0003020211212010-0233011031121122-1100003111221023-3120033211201200-3131100010210220"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-0330033313030210-0331213030302203-2221233323303020-0101111020321323-2221011311200220-1310101232312130-3023212122103323-0130122211210302"></a>

## Direct properties — ip_prefix_list / 332212112131 / 3

<a id="canonical-0011100210231222-1033313312211323-2331311030121322-0110201320303111-1232231131101003-1030123301112323-2130333232203212-0132212322310230"></a>

<a id="canonical-3320121222331112-3320303203222220-2331123232310330-3131332212330023-1002102013101001-3310312231101233-3332201132223131-1132222223221123"></a>

## invert_match property — ip_prefix_list / 332212112131 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-0031022123000233-0310221002330012-2012021110121103-0223230332331121-1011131302021012-0131100021221032-1123331303002021-0123300302110110"></a>

<a id="canonical-1222112331111000-0333013332111123-3133122300333212-0031123102012120-3023110130022001-3113112203310212-2313023233021303-1200303133320021"></a>

## ip_prefixes property — ip_prefix_list / 332212112131 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2210320330120233-1132300012200002-3230013233123301-0300032030130031-0033112212012312-2013132023100100-3002213012023230-0221311010102121"></a>

## Next pages — ip_prefix_list / 332212112131 / 6

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233322300311101-1121230331312310-0310330203113130-3322000002130330-3110021112303122-0330233221131332-3022330130221210-0223213010023000"></a>

## rule_list.rules.ip_prefix_set — ip_prefix_set / 212011032212 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.ip_prefix_set

<a id="canonical-3332033131113033-3223220211023121-2323003020131321-3131021310322311-2331111332313123-1231130313232013-3333300120212213-0003202233301301"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-1000033110303330-0012230232123211-2233313011101112-3213211322110002-0003013310330022-1100101001322102-1100120021031112-1003002333113312"></a>

## Direct properties — ip_prefix_set / 212011032212 / 3

<a id="canonical-1022203121321033-0323013030113101-0123202121220101-2210030023222333-2221132200013333-3222322033202211-2312233300003313-3332131120213212"></a>

<a id="canonical-3120230021020001-3312011312210323-1131300333230331-1302211212023232-0131331121023301-2121020003211330-1103233303110310-3210002121032301"></a>

## invert_matcher property — ip_prefix_set / 212011032212 / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-3220312110010010-3303030100022101-1113230110332302-2213013322303121-0113001002230211-0231212231210000-2023320212213100-3130313010122302): complete subsection reference.

<a id="canonical-1321110222002121-1233203222331212-2332323211002122-1102111333101211-2311303110011300-0203311103211202-1013120223031330-3022233122200133"></a>

## Next pages — ip_prefix_set / 212011032212 / 5

- [rule_list.rules.ip_prefix_set.prefix_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-3220312110010010-3303030100022101-1113230110332302-2213013322303121-0113001002230211-0231212231210000-2023320212213100-3130313010122302)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-3220312110010010-3303030100022101-1113230110332302-2213013322303121-0113001002230211-0231212231210000-2023320212213100-3130313010122302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131112000333331-1031333300300020-0213103103331112-1132212012020220-1213222022031330-2010312311200031-2313110000121121-1031013010112131"></a>

## rule_list.rules.ip_prefix_set.prefix_sets — prefix_sets / 030133300220 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031)
- rule_list.rules.ip_prefix_set.prefix_sets

<a id="canonical-0111330220322332-2222122313210332-3213233121202112-1202223102220232-2023223033203013-2023030120001000-0113030110312331-3113311311013232"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-1300233022222310-1000103312223330-3203020123032130-2321301113332003-1002302210323002-2312223130211013-3002101223002331-0003330221130321"></a>

## Direct properties — prefix_sets / 030133300220 / 3

<a id="canonical-3122321030312102-0311210231300230-3331002231133133-3321312112323022-2110320203322110-2112100102201103-3121221300110012-0113122313332102"></a>

<a id="canonical-0120333122300132-2232201331233233-0000131201210332-2330111001320333-1300133210002101-1203203332133210-3010311323030311-3020310222202122"></a>

## kind property — prefix_sets / 030133300220 / 4

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

<a id="canonical-2031022303111023-3231200021210210-1132001222300213-2323013013133223-3212232303302313-2110100010002012-0001310323231132-3310212011112321"></a>

<a id="canonical-1013232012123133-3300202120332013-3313022003023031-2321013220213123-1331113023231231-3023223231221011-2011332012313230-1303310311211312"></a>

## name property — prefix_sets / 030133300220 / 5

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

<a id="canonical-0030133031021000-0213210113331200-2221013301032021-1131300301112233-3313313111110311-1330030221211113-3210232300330212-0312132122312213"></a>

<a id="canonical-0133133323031322-0301120120020212-1300022220131212-0310030220103232-2103302312320010-1231320013320122-2021123202013131-0201120002210311"></a>

## namespace property — prefix_sets / 030133300220 / 6

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

<a id="canonical-1032221033221000-0103303021101113-2203311312010131-3223003213312220-3323012323223232-0011111323310322-3111202032021233-0120112011311033"></a>

<a id="canonical-1100111111120311-0103222102203033-3011302331220323-2011221312013113-2133110122220332-3223202320322011-0022001130322313-3320311102211003"></a>

## tenant property — prefix_sets / 030133300220 / 7

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

<a id="canonical-2021132002020222-2131002101030031-1102213311100123-1131331122201021-3231123303101230-1122333201230330-0230031203133333-1322122330310012"></a>

<a id="canonical-1110102231212020-0102021002323323-3233020102332311-3320222231111021-0031223032032130-1312213232123111-2022312333012321-0132322001332111"></a>

## uid property — prefix_sets / 030133300220 / 8

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

<a id="canonical-0221300011322021-3111022333023000-2133033132133300-0321003221310201-2203122321110103-3120221102112000-2313300221122221-1021233130230000"></a>

## Next pages — prefix_sets / 030133300220 / 9

- [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)

<a id="canonical-0112111313220331-1330313331010211-0132030000331111-0010130323332331-2312213111123132-0033311200200110-2322101310110300-1302122122212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122210031122230-0212001322313300-3300000033313032-0100311212232231-0010030322130101-3021030210311213-2212222311023302-2022102103210112"></a>

## rule_list.rules.pool — pool / 100130223033 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.pool

<a id="canonical-0213110333030222-0112330213022021-0122113000132331-3213123302101221-3001313210133231-3122122001012122-3313210332211003-1001201123030020"></a>

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

<a id="canonical-0110201111232021-0232300131310231-3321201222232123-3021122232013001-2222001211233130-0001000201030123-1123202003331131-0211221000331130"></a>

## Direct properties — pool / 100130223033 / 3

<a id="canonical-0210213210123102-1220133333130031-3133220301103221-2032032202321331-3200230321213230-0222021201112223-2013010122003130-2122212033323223"></a>

<a id="canonical-1212222220332322-2310131020013130-2113331112111221-1222003313313313-3021030200223310-3320023020211222-0320330320200110-1303313102000313"></a>

## name property — pool / 100130223033 / 4

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

<a id="canonical-3011000131120303-0331010021022221-3200201331010231-1012111213102000-3022102002120301-3231230032211320-2111020332320233-3020321020011210"></a>

<a id="canonical-1133032103301330-3100230313222233-1322132303220301-2123330303303000-0000021103102133-0110120312023310-2101012302211112-0232031321011123"></a>

## namespace property — pool / 100130223033 / 5

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

<a id="canonical-1213010322112020-3131302032032023-1223003300121122-3131112130102031-0003010201231322-3202222323033131-1323112030331212-2202220221000011"></a>

<a id="canonical-1330003000110232-3030102310121230-0021133310003211-1323232312031321-2230303301000020-3301302300331321-2001331001221311-2022333231002101"></a>

## tenant property — pool / 100130223033 / 6

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

<a id="canonical-1233121132330322-2220121103300231-3003311013021113-0300133132133201-2323030331331221-0110132333202030-2031022120123030-0300231112323001"></a>

## Next pages — pool / 100130223033 / 7

- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
