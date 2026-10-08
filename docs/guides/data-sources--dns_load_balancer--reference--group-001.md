---
page_title: "xcsh_dns_load_balancer reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer reference."
---

# xcsh_dns_load_balancer reference

<a id="canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- Property reference

<a id="canonical-2321031023320223-3000003000323210-0033022230330130-2223301132032103-1031321213021112-2011201121223330-3121320121332113-0013212212113210"></a>

### Direct properties for `xcsh_dns_load_balancer`

<a id="canonical-0113323100322023-0303310101313112-2212021331311223-0013022011023232-2121000313000310-3113333123313111-3312032103323112-0030333212210311"></a>

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

<a id="canonical-2121103030310123-2023003003233032-0213120212303231-3031330321032130-1310300030132300-3220003100112333-2032312001000210-3203130133021022"></a>

<a id="canonical-0022213332201103-3233212210021333-0101011202303212-1303332002311003-3330032222330131-2022302302210101-0303132203321022-1302232030030013"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the DNSLoadBalancer.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1130230101310020-2111332312000021-2310200132321020-3232331312001130-0130201022201002-3113202211311022-0000113310220233-3231131113231012"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0223232130001111-1013302213130233-0022303103021212-2020102233112033-3331113332133302-2023111033322111-3330033032031112-2113003222231021"></a>

<a id="canonical-3221322310223131-0310110020033023-3120000222220312-3221223301320232-1123131223313021-0123202100231232-2312002133320311-2011323001333031"></a>

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

<a id="canonical-1331100202102121-0323311333233010-2021100220102132-0211303032213233-0203222111112231-0120301020013112-1121021102200013-0022311313003110"></a>

<a id="canonical-3200221001321310-1110300031003211-0203320313210102-0111320003220022-3012333132231312-0122130202232002-0031320203020213-2011213312313202"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DNSLoadBalancer.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1313023203102023-3020310102023033-3203033130321022-3101213110012230-1323202002132011-2011001331021200-1031120013222123-2003030310121212"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the DNSLoadBalancer exists.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1303302211302121-3221002200323203-0130322100010212-3010211320201201-2220100032123021-3113022030313233-3103023200023333-1103122220332223"></a>

#### `record_type` property

Type: `"string"`. Computed.

\[Enum: A|AAAA|MX|CNAME|SRV\] Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME -
SRV: SRV. Possible values are \`A\`, \`AAAA\`, \`MX\`, \`CNAME\`, \`SRV\`. Defaults to \`A\`.

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

<a id="canonical-0303111110001231-0332000013133112-1310200103100200-3122222003312211-2230023301121131-3011123322300001-0311110133132320-3130033111311223"></a>

### All schema paths for `xcsh_dns_load_balancer`

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

<a id="canonical-2101122021130201-3101010132313031-3002332123310323-2131103023023221-3121010320123013-2302332030132230-3211120320330301-2212011213122210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `fallback_pool` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- fallback_pool

<a id="canonical-3220301101022132-3012333212023100-1233312332121133-3112331200223032-3230102313002013-3211321121333032-1102210113013000-0121100232302200"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0021003220102313-3133131001001133-1123033103230020-1201013230000211-1131301312331211-2121022321013233-1031031000221033-0222203133231113"></a>

### Direct properties for `fallback_pool`

<a id="canonical-0203300123200230-3123010030012133-1020223231222203-2101022021120213-2301223102113230-1301311323111213-2213210011011233-0321113210121112"></a>

#### `fallback_pool.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1322221230100330-2003213320230133-1032010113110120-0331210000330323-0322313013201021-3110020303222233-1032102032213322-2311031322230211"></a>

#### `fallback_pool.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0323011110033233-0302011131330312-3120211202003312-0131123131112010-0113033110002303-3222211211201000-1300101122110122-2321100030210001"></a>

#### `fallback_pool.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cache` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- response_cache

<a id="canonical-1330333221222111-3232030220331332-3222232223133321-3020332313011130-1323001323213231-3021211102132211-0231201111212320-1212030023002021"></a>

Type: `"single"`. Computed.

Configuration parameter for response cache.

Additional upstream details:

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

<a id="canonical-3213322200333232-2322031100031332-1323203210003122-3231220302212230-2103303101030031-2232231320232032-0002133012100132-0203301200203210"></a>

### Direct properties for `response_cache`

- [default_response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-0033011311220122-3232113303221303-1012003100001323-2223031022220221-3231103130102203-1333022100222212-3213121300313113-3331221303212220): complete subsection reference.

- [disable_spec](data-sources--dns_load_balancer--reference--group-001.md#canonical-1311111011031300-0030203033032300-2323132310331303-1230231021213210-2302203123233310-1232321103130310-2202022221311330-1102130332221133): complete subsection reference.

- [response_cache_parameters](data-sources--dns_load_balancer--reference--group-001.md#canonical-3322233121222221-1232020102211000-0003211203101023-0122200112202212-3100313113023302-2130121120101123-3023221303021123-2101303102300033): complete subsection reference.

<a id="canonical-0033011311220122-3232113303221303-1012003100001323-2223031022220221-3231103130102203-1333022100222212-3213121300313113-3331221303212220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cache.default_response_cache_parameters` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- response_cache.default_response_cache_parameters

<a id="canonical-1213033122130021-1221130220103300-0122103213001320-0313211133033121-0230310010030010-3300332000030112-2123213201101213-1221001112101112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default response cache parameters.

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

<a id="canonical-1311111011031300-0030203033032300-2323132310331303-1230231021213210-2302203123233310-1232321103130310-2202022221311330-1102130332221133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cache.disable_spec` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [response_cache](data-sources--dns_load_balancer--reference--group-001.md#canonical-1022220100313001-1303313330111232-0300132223021231-1123321233211231-1103123031102201-2032002203302332-1030111321331021-3000112323222120)
- response_cache.disable_spec

<a id="canonical-2113110132101110-3122301113313311-3303000310201211-1001221110001211-1222313033011330-2020222200031020-3123303100222003-1011231311012130"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322233121222221-1232020102211000-0003211203101023-0122200112202212-3100313113023302-2130121120101123-3023221303021123-2101303102300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `response_cache.response_cache_parameters` properties

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

<a id="canonical-1311201213223222-0122011213103213-2221201011130022-3312002113022031-2222332212220022-1013303111132233-3002013021012200-3300221023101023"></a>

### Direct properties for `response_cache.response_cache_parameters`

<a id="canonical-3310013020321022-0020123011121220-1102303103130010-3013200013103233-2321321303201230-0031130332221200-0300303121303300-0132300002211203"></a>

#### `response_cache.response_cache_parameters.cache_cidr_ipv4` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1021222302220320-1030102023330013-2301223003233011-2303113100301202-3132020131223310-0300031312301103-2002023021121321-0133201112111230"></a>

#### `response_cache.response_cache_parameters.cache_cidr_ipv6` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2033002132322013-0121221110010332-0110131232222221-3321102332213223-3322011301000113-3313122013230331-1223121212023323-1333100100331120"></a>

#### `response_cache.response_cache_parameters.cache_ttl` property

Type: `"number"`. Computed.

TTL. TTL for response cache.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- rule_list

<a id="canonical-0031011301121012-1100012312031122-1013223321101300-1103000002211233-1213022130220323-0101103122310311-0130010122323311-1102033333200312"></a>

Type: `"single"`. Computed.

Load Balancing Rule List. List of the Load Balancing Rules.

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

<a id="canonical-1201323033330310-2213220321032203-3310113011333202-1203303032001230-0312103213331333-2030102310323021-1011310311311302-2311311000101223"></a>

### Direct properties for `rule_list`

- [rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103): complete subsection reference.

<a id="canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- rule_list.rules

<a id="canonical-1203303211010203-3012113301022131-3021302200231302-3111000321202330-3101311030300213-2303311322133103-2211213002012222-1303301031010012"></a>

Type: `"list"`. Computed.

Load Balancing Rules. Rules to perform load balancing.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3302122221122012-0030300231320203-1200211220132013-0133313300330123-3330021321023102-2120200330320302-0133210110220131-3132232313232231"></a>

### Direct properties for `rule_list.rules`

- [asn_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-0312303233022323-1130031221213021-3301322221331113-0311211111010203-0112310101121021-3202033232231232-1121233300123113-0223010110032323): complete subsection reference.

- [asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021): complete subsection reference.

- [geo_location_label_selector](data-sources--dns_load_balancer--reference--group-001.md#canonical-1223222103223132-3310031101003102-0100001031302110-1102220231301200-3000223211112000-0220131220011321-1323310100101033-2121133221233230): complete subsection reference.

- [geo_location_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-2101001211013222-3310330203000121-0322202123032012-2100232331031322-3310013112333132-0002120232000332-3030202231102233-2111133002332002): complete subsection reference.

- [ip_prefix_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1032323121031133-1120303221210112-1202030322120300-0010021123012330-1003020210000013-0102110010100233-1112000203320013-1003331201230022): complete subsection reference.

- [ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031): complete subsection reference.

- [pool](data-sources--dns_load_balancer--reference--group-001.md#canonical-0112111313220331-1330313331010211-0132030000331111-0010130323332331-2312213111123132-0033311200200110-2322101310110300-1302122122212323): complete subsection reference.

<a id="canonical-2022020110311001-2202132133211221-0320320212022122-3131230133203210-0320032011023203-2213032322022203-2300311001133013-3223002001232011"></a>

<a id="canonical-1032320331102002-3321233122011011-0030121311211301-1110312013211220-0020323011313330-1201323010033323-1100102010231302-2013113312212001"></a>

#### `rule_list.rules.score` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0312303233022323-1130031221213021-3301322221331113-0311211111010203-0112310101121021-3202033232231232-1121233300123113-0223010110032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.asn_list` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.asn_list

<a id="canonical-3230313103313102-3032223131233020-3031202232133232-1030230321233002-0130102310010322-0021020213202133-1003101020322331-0320002201113001"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2230331200031302-1102211003121011-1120013131231110-0303131121330102-0332032231001000-2233120103123330-0202023303022032-2032013102111111"></a>

### Direct properties for `rule_list.rules.asn_list`

<a id="canonical-0003212012202021-3122012113323132-0101302313222030-1200003133033133-1011102021112202-1301312001133312-1132221233112213-0203231132300120"></a>

#### `rule_list.rules.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.asn_matcher` properties

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

<a id="canonical-0131320210233310-3100033110301203-1030210320121321-3202222231023213-1223220230110313-0221031031131310-1001010102010013-1332212100311223"></a>

### Direct properties for `rule_list.rules.asn_matcher`

- [asn_sets](data-sources--dns_load_balancer--reference--group-001.md#canonical-3310030230132020-2301000101023212-1301221101111332-3001023233033012-3013011031220310-2103312302131332-3313003230221310-2030122300001220): complete subsection reference.

<a id="canonical-3310030230132020-2301000101023212-1301221101111332-3001023233033012-3013011031220310-2103312302131332-3313003230221310-2030122300001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [rule_list.rules.asn_matcher](data-sources--dns_load_balancer--reference--group-001.md#canonical-2312111032120320-3220222130131200-2001132220223132-1301132133132002-0120322212131231-1120333133331221-0233221033202233-1131122300232021)
- rule_list.rules.asn_matcher.asn_sets

<a id="canonical-2010022100021130-0033223230031323-2321211200232201-2210321011121303-3122122331111023-3300002331012300-2312312331212200-0110002002202133"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2201120112132113-2310202120211221-0210320023303022-1311031212332131-3330301011321031-0021132201331323-2031203232321203-3230102030310220"></a>

### Direct properties for `rule_list.rules.asn_matcher.asn_sets`

<a id="canonical-3010113311002011-0133323220202121-0131012013331311-3000021233001021-2133123102222023-3232211223301121-1203032230031322-2213223223220302"></a>

#### `rule_list.rules.asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1323300231013323-3210021003303032-3100230123232003-1022301221200310-1210001033301011-1221023322113230-3301220221300203-0201101002301313"></a>

#### `rule_list.rules.asn_matcher.asn_sets.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2032312301001130-0210013030020313-3211130333002303-2232023233022021-1002230321313103-1121021132022231-0230010010323313-1013101320002012"></a>

#### `rule_list.rules.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0321313321022321-1101210302032133-2201311101232010-0303203031002231-1200133120031102-0012212331303022-2201212020211300-1312222120120023"></a>

#### `rule_list.rules.asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2333010330101030-1020111131323132-2202332213203010-3031102103322311-3013233132200322-3223120031300122-2203012213202332-2303233230330122"></a>

#### `rule_list.rules.asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1223222103223132-3310031101003102-0100001031302110-1102220231301200-3000223211112000-0220131220011321-1323310100101033-2121133221233230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.geo_location_label_selector` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.geo_location_label_selector

<a id="canonical-3330211011013301-1131121002231310-2212010200202332-1022300030001033-0330021310310201-2333130320200113-3210130102303311-3310311120133010"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2200023312332101-2131312232001303-0223101222232022-0103001012133331-3103101220131013-3133232033100122-2212313113323030-0220233322112010"></a>

### Direct properties for `rule_list.rules.geo_location_label_selector`

<a id="canonical-0213311101313310-2213111301320203-0232302311213003-0331332033102231-3220322003331121-1003132112120321-2021330123003212-1023212120210112"></a>

#### `rule_list.rules.geo_location_label_selector.expressions` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2101001211013222-3310330203000121-0322202123032012-2100232331031322-3310013112333132-0002120232000332-3030202231102233-2111133002332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.geo_location_set` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.geo_location_set

<a id="canonical-1201330303301311-3100030231101131-1221313020021121-2202310022011223-3002000101302213-1300032233212111-0123101321303021-0012201332201021"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1103121110130222-3331323003231201-2330232101300203-0330131013222031-3023222212110331-3220002333003022-1311123030021023-1003010320200000"></a>

### Direct properties for `rule_list.rules.geo_location_set`

<a id="canonical-0213011032233031-3220202231000232-2011300222012030-3101130201122300-3102311111032110-1132332300120002-2022333112332231-1003320222023202"></a>

#### `rule_list.rules.geo_location_set.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0103013222110001-0221030020113333-0123321201302300-1223020102131001-0110023012231031-1310012203022323-3202210222100132-1220312330131131"></a>

#### `rule_list.rules.geo_location_set.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3323211313222213-3130323230233022-1333111212032130-0100200033303322-1203211212103222-3330101310022303-0013022120101020-2133111001332330"></a>

#### `rule_list.rules.geo_location_set.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1032323121031133-1120303221210112-1202030322120300-0010021123012330-1003020210000013-0102110010100233-1112000203320013-1003331201230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.ip_prefix_list` properties

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

<a id="canonical-3101000230202101-2020101031301121-1213230132213223-2223112203022132-3312303101302031-3103033001232303-1330203020333101-2010123110212300"></a>

### Direct properties for `rule_list.rules.ip_prefix_list`

<a id="canonical-0011100210231222-1033313312211323-2331311030121322-0110201320303111-1232231131101003-1030123301112323-2130333232203212-0132212322310230"></a>

#### `rule_list.rules.ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-0330033313030210-0331213030302203-2221233323303020-0101111020321323-2221011311200220-1310101232312130-3023212122103323-0130122211210302"></a>

#### `rule_list.rules.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.ip_prefix_set` properties

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

<a id="canonical-1233322300311101-1121230331312310-0310330203113130-3322000002130330-3110021112303122-0330233221131332-3022330130221210-0223213010023000"></a>

### Direct properties for `rule_list.rules.ip_prefix_set`

<a id="canonical-1022203121321033-0323013030113101-0123202121220101-2210030023222333-2221132200013333-3222322033202211-2312233300003313-3332131120213212"></a>

#### `rule_list.rules.ip_prefix_set.invert_matcher` property

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

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

<a id="canonical-3220312110010010-3303030100022101-1113230110332302-2213013322303121-0113001002230211-0231212231210000-2023320212213100-3130313010122302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.ip_prefix_set.prefix_sets` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- [rule_list.rules.ip_prefix_set](data-sources--dns_load_balancer--reference--group-001.md#canonical-3103121010111300-1213220310221000-3302231030011231-0201211121312112-2313302320301132-1233021113113323-3312203020011213-3203213321121031)
- rule_list.rules.ip_prefix_set.prefix_sets

<a id="canonical-0111330220322332-2222122313210332-3213233121202112-1202223102220232-2023223033203013-2023030120001000-0113030110312331-3113311311013232"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0131112000333331-1031333300300020-0213103103331112-1132212012020220-1213222022031330-2010312311200031-2313110000121121-1031013010112131"></a>

### Direct properties for `rule_list.rules.ip_prefix_set.prefix_sets`

<a id="canonical-3122321030312102-0311210231300230-3331002231133133-3321312112323022-2110320203322110-2112100102201103-3121221300110012-0113122313332102"></a>

#### `rule_list.rules.ip_prefix_set.prefix_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1300233022222310-1000103312223330-3203020123032130-2321301113332003-1002302210323002-2312223130211013-3002101223002331-0003330221130321"></a>

#### `rule_list.rules.ip_prefix_set.prefix_sets.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0120333122300132-2232201331233233-0000131201210332-2330111001320333-1300133210002101-1203203332133210-3010311323030311-3020310222202122"></a>

#### `rule_list.rules.ip_prefix_set.prefix_sets.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1013232012123133-3300202120332013-3313022003023031-2321013220213123-1331113023231231-3023223231221011-2011332012313230-1303310311211312"></a>

#### `rule_list.rules.ip_prefix_set.prefix_sets.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0133133323031322-0301120120020212-1300022220131212-0310030220103232-2103302312320010-1231320013320122-2021123202013131-0201120002210311"></a>

#### `rule_list.rules.ip_prefix_set.prefix_sets.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0112111313220331-1330313331010211-0132030000331111-0010130323332331-2312213111123132-0033311200200110-2322101310110300-1302122122212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.pool` properties

Breadcrumbs:

- [xcsh_dns_load_balancer](../data-sources/dns_load_balancer.md#canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103)
- [Property reference](data-sources--dns_load_balancer--reference--group-001.md#canonical-3132200331022103-0202102030232123-3221123330030310-3132333003123202-1220120012203031-2132331133031102-0103322031301112-1330312100102222)
- [rule_list](data-sources--dns_load_balancer--reference--group-001.md#canonical-1200302330200113-2111201232232222-0320032111123202-0022023201211103-0331313123013213-0212322233213130-3130223133120312-3323032210313223)
- [rule_list.rules](data-sources--dns_load_balancer--reference--group-001.md#canonical-0132020012322023-0131332323132323-1302002121000322-3310030103202013-1311021202311033-3313120012222101-1332013123123000-0333301120331103)
- rule_list.rules.pool

<a id="canonical-0213110333030222-0112330213022021-0122113000132331-3213123302101221-3001313210133231-3122122001012122-3313210332211003-1001201123030020"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2122210031122230-0212001322313300-3300000033313032-0100311212232231-0010030322130101-3021030210311213-2212222311023302-2022102103210112"></a>

### Direct properties for `rule_list.rules.pool`

<a id="canonical-0210213210123102-1220133333130031-3133220301103221-2032032202321331-3200230321213230-0222021201112223-2013010122003130-2122212033323223"></a>

#### `rule_list.rules.pool.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0110201111232021-0232300131310231-3321201222232123-3021122232013001-2222001211233130-0001000201030123-1123202003331131-0211221000331130"></a>

#### `rule_list.rules.pool.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1212222220332322-2310131020013130-2113331112111221-1222003313313313-3021030200223310-3320023020211222-0320330320200110-1303313102000313"></a>

#### `rule_list.rules.pool.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
