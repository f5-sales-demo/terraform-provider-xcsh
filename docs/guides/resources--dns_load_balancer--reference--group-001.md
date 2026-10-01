---
page_title: "xcsh_dns_load_balancer reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_load_balancer reference."
---

# xcsh_dns_load_balancer reference

<a id="canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133022322302211-2023232133033111-2211103011330012-3000220202301121-3013210212122213-3000001110023323-3021212211302101-0012133312330203"></a>

## Property reference — Property reference / 203302123010 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- Property reference

<a id="canonical-0011103302121222-3032130211333130-2030031011101332-0200313001131320-3021001132223233-2000132132130100-1220020302010201-3133303210103302"></a>

## Direct properties — Property reference / 203302123010 / 3

<a id="canonical-0123112110210101-0311130232120312-1130320233013231-1211132300220322-2322023312020303-2112133103312313-3112200102102110-0311130330122011"></a>

<a id="canonical-1132132223023313-2000000301112023-3130222110121031-1211101101022213-3120001322201022-2213132111220012-0132000121013112-1313230302003012"></a>

## annotations property — Property reference / 203302123010 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-3333320210210311-2000303230233001-3021300231112031-0110321032210323-1310103031102132-0131011112301130-3020311332111021-1120212122120010"></a>

<a id="canonical-2023032200333023-3222032301000230-2130220103000121-2122023323010102-3102203030312130-0323333111003100-2223211110203131-2121021023231332"></a>

## description property — Property reference / 203302123010 / 5

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

<a id="canonical-0020213130032223-3323230222201332-3120232320310213-0203003123020023-0013110013030331-2110003222013030-0010300313332323-2010210001211321"></a>

<a id="canonical-3333023232123132-1003222002220021-2201112332230321-1021330222203233-2211113030010232-2211323220323102-2123233133332323-1112032211122112"></a>

## disable property — Property reference / 203302123010 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

- [fallback_pool](resources--dns_load_balancer--reference--group-001.md#canonical-1212210003110211-3001201121231212-1332220221111323-1100201000223220-2301012002013202-2331013101210221-2212310112013322-1223121320211020): complete subsection reference.

<a id="canonical-3230033203122311-3102302310330303-2211122303031120-2322331003123130-2210331010113122-1300331113232213-3333333223012120-0313320301120320"></a>

<a id="canonical-0331003022102011-2011303331012133-3303210030322222-1303331021221130-0021121033111121-1312113001321323-1000303112131300-3322203113132200"></a>

## ID property — Property reference / 203302123010 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0213232102200301-0312111110020112-0232311031100333-3022230033232011-1323200233200000-0110200113023310-0133302313313122-3023132132220020"></a>

<a id="canonical-1213331123321033-3331302103103102-0230032303310011-3031213303031303-2212010012132000-1023230113202111-1310201113013020-2332102210223223"></a>

## labels property — Property reference / 203302123010 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-3330230201022031-0300132003210120-0332232003223023-0302101102002001-2131020033001303-2313122111120332-2210322333010102-1312200000223213"></a>

<a id="canonical-0330233332323222-2201110301230323-1023021131202321-3222123100232111-2301233302202320-1111030330323333-1330031003000023-3130221220133320"></a>

## name property — Property reference / 203302123010 / 9

Type: `"string"`. Required.

Name of the DNS Load Balancer. Must be unique within the namespace.

Upstream description:

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

<a id="canonical-0131030011121112-1132221313303112-1013003033111223-0210203311331120-1202000031031003-3003312300131002-3123122303002201-2011323311030002"></a>

<a id="canonical-1233220331213003-3012120022313301-3321301121200232-1001033130021200-0322213122331310-2102310012300031-2331022322321202-3311020131032122"></a>

## namespace property — Property reference / 203302123010 / 10

Type: `"string"`. Optional, Computed.

Namespace for the DNS Load Balancer. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

<a id="canonical-2312131202103320-1020213310311120-1010202103232013-0021213232232012-1300010202210332-0311111212202112-2022000032031320-2223030033303203"></a>

<a id="canonical-0122033222303210-2300000121130203-0023021303130002-2133123330322003-2100012213323013-1202010231121311-0310332213222133-1122122212222320"></a>

## record_type property — Property reference / 203302123010 / 11

Type: `"string"`. Optional, Computed.

\[Enum: A|AAAA|MX|CNAME|SRV\] Resource Record Type - A: A - AAAA: AAAA - MX: MX - CNAME: CNAME -
SRV: SRV. Possible values are \`A\`, \`AAAA\`, \`MX\`, \`CNAME\`, \`SRV\`. Defaults to \`A\`.

Upstream description:

Resource Record Type

&#8203;- A: A

&#8203;- AAAA: AAAA

&#8203;- MX: MX

&#8203;- CNAME: CNAME

&#8203;- SRV: SRV.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("A",
    "AAAA",
    "MX",
    "CNAME",
    "SRV"),
}
```

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

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033): complete subsection reference.

- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123): complete subsection reference.

- [timeouts](resources--dns_load_balancer--reference--group-001.md#canonical-3020101101220100-1331123312200221-0133330233120332-1010123102022230-0221113111221210-1030110202230033-2301033013012302-2010221031233011): complete subsection reference.

<a id="canonical-2021212112332011-3300100222313102-2210312120030333-2101311233000113-0310103021223223-3203110020230202-3132220201132102-1131132131211130"></a>

## All schema paths — Property reference / 203302123010 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_load_balancer--reference--group-001.md#canonical-0123112110210101-0311130232120312-1130320233013231-1211132300220322-2322023312020303-2112133103312313-3112200102102110-0311130330122011) |
| `description` | [description](resources--dns_load_balancer--reference--group-001.md#canonical-3333320210210311-2000303230233001-3021300231112031-0110321032210323-1310103031102132-0131011112301130-3020311332111021-1120212122120010) |
| `disable` | [disable](resources--dns_load_balancer--reference--group-001.md#canonical-0020213130032223-3323230222201332-3120232320310213-0203003123020023-0013110013030331-2110003222013030-0010300313332323-2010210001211321) |
| `fallback_pool` | [fallback_pool](resources--dns_load_balancer--reference--group-001.md#canonical-2121111302210113-0313112310213022-2331130333201223-1111301111113333-1102120003013113-2002212230032123-1122012010121022-1310021230101303) |
| `fallback_pool.name` | [fallback_pool.name](resources--dns_load_balancer--reference--group-001.md#canonical-2330132021200100-1013021123303300-0332203022303320-0323100130231112-0213233300302203-0223323131000033-1221233312021313-1002322321003101) |
| `fallback_pool.namespace` | [fallback_pool.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-2303303121302002-0230003312332323-1311132220002001-0222002011213200-2122102031133222-1300202021303220-3020100102311201-1012112221010300) |
| `fallback_pool.tenant` | [fallback_pool.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-1233013202002220-0103202213211310-2222203110032332-0211033231230122-2300320023230001-1211102302112200-3233202031312130-2300110203330322) |
| `id` | [id](resources--dns_load_balancer--reference--group-001.md#canonical-3230033203122311-3102302310330303-2211122303031120-2322331003123130-2210331010113122-1300331113232213-3333333223012120-0313320301120320) |
| `labels` | [labels](resources--dns_load_balancer--reference--group-001.md#canonical-0213232102200301-0312111110020112-0232311031100333-3022230033232011-1323200233200000-0110200113023310-0133302313313122-3023132132220020) |
| `name` | [name](resources--dns_load_balancer--reference--group-001.md#canonical-3330230201022031-0300132003210120-0332232003223023-0302101102002001-2131020033001303-2313122111120332-2210322333010102-1312200000223213) |
| `namespace` | [namespace](resources--dns_load_balancer--reference--group-001.md#canonical-0131030011121112-1132221313303112-1013003033111223-0210203311331120-1202000031031003-3003312300131002-3123122303002201-2011323311030002) |
| `record_type` | [record_type](resources--dns_load_balancer--reference--group-001.md#canonical-2312131202103320-1020213310311120-1010202103232013-0021213232232012-1300010202210332-0311111212202112-2022000032031320-2223030033303203) |
| `response_cache` | [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-1011311312311033-2123301323200022-0232303221333113-1332230110202230-2130122221301303-1131200232012300-3031111303033220-1031001103322312) |
| `response_cache.default_response_cache_parameters` | [response_cache.default_response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-2112120233033220-2221322023023101-3112111310331101-0100102222100120-3020330010313321-1331003333002302-0133303000332003-0331113310223200) |
| `response_cache.disable_spec` | [response_cache.disable_spec](resources--dns_load_balancer--reference--group-001.md#canonical-2211301201022123-2312012131300003-2030013232022022-2220331222100230-1302230233133222-1200300323202003-2331032130333333-0011330021032001) |
| `response_cache.response_cache_parameters` | [response_cache.response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-2233312312221021-0313321001033022-1123130330313113-1010102312220030-0110332303303302-2230011023020000-3323220023200000-3212033001030232) |
| `response_cache.response_cache_parameters.cache_cidr_ipv4` | [response_cache.response_cache_parameters.cache_cidr_ipv4](resources--dns_load_balancer--reference--group-001.md#canonical-0202210202022213-2103003032303231-1122000011131320-3311212013121310-2200212202120223-0002212121002233-1012323033123120-1302201102211101) |
| `response_cache.response_cache_parameters.cache_cidr_ipv6` | [response_cache.response_cache_parameters.cache_cidr_ipv6](resources--dns_load_balancer--reference--group-001.md#canonical-0323110302202133-1101332221211010-2220101120303203-0023100001323103-0101102013011300-3313301312232021-0330003201101332-1130133101113303) |
| `response_cache.response_cache_parameters.cache_ttl` | [response_cache.response_cache_parameters.cache_ttl](resources--dns_load_balancer--reference--group-001.md#canonical-3223213003233022-3103032223221100-3001202223322122-1302233123011330-0302110232232120-1221122333132013-2232210020103330-1303310300203332) |
| `rule_list` | [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-1011030300133330-1232233331000203-0001220301010012-2022333222201231-3120200231121310-3013230200232231-2331121301011003-2302030232030212) |
| `rule_list.rules` | [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-1232202012132220-2130223201022300-3300032011330232-0131322223001322-3231013201023211-3110301303321132-3323003320222200-0020020011100120) |
| `rule_list.rules.asn_list` | [rule_list.rules.asn_list](resources--dns_load_balancer--reference--group-001.md#canonical-0303111100131233-2122100220230301-0310111132022321-1223201023322220-3330310202122313-0002120310233102-3132313013000202-1333000212211330) |
| `rule_list.rules.asn_list.as_numbers` | [rule_list.rules.asn_list.as_numbers](resources--dns_load_balancer--reference--group-001.md#canonical-0032202121213011-2011012031103101-1231031013102002-1132001010011012-3112131101023231-1333310002121002-3122101001103311-1010300002100001) |
| `rule_list.rules.asn_matcher` | [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-2101022121002302-1311113101123303-0020233213202023-3132012011112233-1130013220313033-2300031000232312-0100202331111230-0203333311013302) |
| `rule_list.rules.asn_matcher.asn_sets` | [rule_list.rules.asn_matcher.asn_sets](resources--dns_load_balancer--reference--group-001.md#canonical-3301100110310203-2302122112312121-0000313331002223-1113022321310213-1220130201320103-3320320302113131-3012023023133232-2103032032002030) |
| `rule_list.rules.asn_matcher.asn_sets.kind` | [rule_list.rules.asn_matcher.asn_sets.kind](resources--dns_load_balancer--reference--group-001.md#canonical-1100212222231111-0023003002312132-2203110031003312-1200221102310321-3013030220322330-2120311133010113-0023033222001103-1101311333023211) |
| `rule_list.rules.asn_matcher.asn_sets.name` | [rule_list.rules.asn_matcher.asn_sets.name](resources--dns_load_balancer--reference--group-001.md#canonical-0132111231212322-0320011113301001-0313321103333021-3313032013021110-0332133202001023-0023222303130320-2200320021203222-1022302331020220) |
| `rule_list.rules.asn_matcher.asn_sets.namespace` | [rule_list.rules.asn_matcher.asn_sets.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-2133202323022100-1201101332102232-0102313223333103-0133130111022030-1032320320321333-0123201033001311-3200132313211130-3010101030200231) |
| `rule_list.rules.asn_matcher.asn_sets.tenant` | [rule_list.rules.asn_matcher.asn_sets.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-3303321103013203-3222222333201233-3031331223231130-1330333123311123-1300003222223310-2221031122033100-1313333103233133-3023323003013323) |
| `rule_list.rules.asn_matcher.asn_sets.uid` | [rule_list.rules.asn_matcher.asn_sets.uid](resources--dns_load_balancer--reference--group-001.md#canonical-3122313020313311-2133302233031212-1020101233103201-1010130200231102-2122223202212023-2122302002311021-0102300031333102-1202303333322003) |
| `rule_list.rules.geo_location_label_selector` | [rule_list.rules.geo_location_label_selector](resources--dns_load_balancer--reference--group-001.md#canonical-3123111103031123-0221320330232113-0131100233131223-1111100000302020-3011331030102313-2302222222110012-0331121022123200-1013001213321201) |
| `rule_list.rules.geo_location_label_selector.expressions` | [rule_list.rules.geo_location_label_selector.expressions](resources--dns_load_balancer--reference--group-001.md#canonical-1332000101300231-1331121130200010-3313233120320301-2301302031230311-3330222130212200-3312333233132301-1310322002332000-2030320230110123) |
| `rule_list.rules.geo_location_set` | [rule_list.rules.geo_location_set](resources--dns_load_balancer--reference--group-001.md#canonical-3123330103223022-2321232202022103-1222012011122113-1330321303102313-0330031332130121-1332203012012231-1332310202113223-0121032010310102) |
| `rule_list.rules.geo_location_set.name` | [rule_list.rules.geo_location_set.name](resources--dns_load_balancer--reference--group-001.md#canonical-2313113032112230-2010203000030113-1131123330120132-1021001101103221-2320102233213131-2233012201201300-2020021112121320-1102110111330001) |
| `rule_list.rules.geo_location_set.namespace` | [rule_list.rules.geo_location_set.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-2222121213133231-3122123333311332-1112322132311111-1210202302213110-0102000002223330-2130211330211130-1103200311211101-2211201013113023) |
| `rule_list.rules.geo_location_set.tenant` | [rule_list.rules.geo_location_set.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-2302233213000123-1313033100323103-3010133011031120-0010031132322102-1003313131003100-1301133212213001-3331002010213012-0020002300210132) |
| `rule_list.rules.ip_prefix_list` | [rule_list.rules.ip_prefix_list](resources--dns_load_balancer--reference--group-001.md#canonical-1001031001310010-2120022232102133-0231311313310222-1102331201131022-2223031331302113-0200001220113132-1130333121231230-2332032200113030) |
| `rule_list.rules.ip_prefix_list.invert_match` | [rule_list.rules.ip_prefix_list.invert_match](resources--dns_load_balancer--reference--group-001.md#canonical-0222023130020230-3321233233300013-1213121123301323-3202233223233131-3232110203203013-0123312230302332-1012003001130103-3201132002232322) |
| `rule_list.rules.ip_prefix_list.ip_prefixes` | [rule_list.rules.ip_prefix_list.ip_prefixes](resources--dns_load_balancer--reference--group-001.md#canonical-3220120310211223-3013222332010101-0123213102022013-0313232102010312-3122031230322322-0330323331321331-2321112122012000-1322113003332230) |
| `rule_list.rules.ip_prefix_set` | [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-0202201303013120-3233121221112022-0011013021113003-2203203132032322-2210112213221211-2122033213322303-0231012231123223-3001332230222103) |
| `rule_list.rules.ip_prefix_set.invert_matcher` | [rule_list.rules.ip_prefix_set.invert_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-3132211120001130-1022202301020123-0313112312203031-1212203003323112-2230200223201121-2123030203131313-1122203031223133-2231000230202122) |
| `rule_list.rules.ip_prefix_set.prefix_sets` | [rule_list.rules.ip_prefix_set.prefix_sets](resources--dns_load_balancer--reference--group-001.md#canonical-1212210332313011-0123032000322223-2003321301100110-3210002302123301-3031002131320331-1000023121132311-2311303110130021-0201030303123100) |
| `rule_list.rules.ip_prefix_set.prefix_sets.kind` | [rule_list.rules.ip_prefix_set.prefix_sets.kind](resources--dns_load_balancer--reference--group-001.md#canonical-3000333110122100-0123003003031231-1300203110222212-2102103303131032-0022320033231033-0030120100222313-3130032201313111-0032032331113022) |
| `rule_list.rules.ip_prefix_set.prefix_sets.name` | [rule_list.rules.ip_prefix_set.prefix_sets.name](resources--dns_load_balancer--reference--group-001.md#canonical-2233221130313231-2211302030013012-0311233133030320-3223220000323212-1300033121333330-2023000001022110-3223020012310210-1131130130321111) |
| `rule_list.rules.ip_prefix_set.prefix_sets.namespace` | [rule_list.rules.ip_prefix_set.prefix_sets.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-3112323112003120-3332301022130301-3210302121201321-1113011211303021-3233220223003130-1002010333203200-2103112231010232-2101113231133202) |
| `rule_list.rules.ip_prefix_set.prefix_sets.tenant` | [rule_list.rules.ip_prefix_set.prefix_sets.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-0102212213013202-3322032003111333-3111013002300013-1221322230211222-3322200022213200-3110312331212020-2101111211220320-3230113100023013) |
| `rule_list.rules.ip_prefix_set.prefix_sets.uid` | [rule_list.rules.ip_prefix_set.prefix_sets.uid](resources--dns_load_balancer--reference--group-001.md#canonical-1331130302001233-0112011332123103-0333013013212012-2110100133031013-3112011331030101-3221311131203102-1320330031102023-2103131121200230) |
| `rule_list.rules.pool` | [rule_list.rules.pool](resources--dns_load_balancer--reference--group-001.md#canonical-0311311200123322-2031130100003101-0210121203003311-1223033133301232-2111313131203223-0032133012021310-2131030323023333-2300012203001022) |
| `rule_list.rules.pool.name` | [rule_list.rules.pool.name](resources--dns_load_balancer--reference--group-001.md#canonical-3110133132013010-2112232222330100-0203232212112032-3013211223223333-3100201120001120-0300113311320112-0202330321222300-1011321013213302) |
| `rule_list.rules.pool.namespace` | [rule_list.rules.pool.namespace](resources--dns_load_balancer--reference--group-001.md#canonical-1022223310210010-3100303200132012-0331210012110120-3312021313002202-3033320103000212-2223303123332023-0013322000131021-1002012200133120) |
| `rule_list.rules.pool.tenant` | [rule_list.rules.pool.tenant](resources--dns_load_balancer--reference--group-001.md#canonical-0320113212331333-2221120320100323-2313121030031022-0213200233231300-2313121001113211-2230132201010310-2202130231230120-3100112232313213) |
| `rule_list.rules.score` | [rule_list.rules.score](resources--dns_load_balancer--reference--group-001.md#canonical-3013123131003322-0101020012331332-2233131131013033-3010110300120002-1320203122122322-3022211232201030-0022320203232133-0000033030020300) |
| `timeouts` | [timeouts](resources--dns_load_balancer--reference--group-001.md#canonical-3132030131303210-3312033122111002-1003030222230032-3321023112103313-3013201022023211-2302123103202221-2100130001101131-0101123000330111) |
| `timeouts.create` | [timeouts.create](resources--dns_load_balancer--reference--group-001.md#canonical-0011231111201302-3210101312233030-0230123111330201-2003121113112333-1011011100131221-0023020101022203-2002103201333232-3332111311112100) |
| `timeouts.delete` | [timeouts.delete](resources--dns_load_balancer--reference--group-001.md#canonical-0220032032131323-0022003201011123-1320121111032003-2000211322010133-1003300210330011-1013111221233310-0102320031221230-2023111102301301) |
| `timeouts.read` | [timeouts.read](resources--dns_load_balancer--reference--group-001.md#canonical-0310310332303000-3333103323002033-1311130022210202-2201030320111233-0220222202100033-0330100020111323-1102203002101313-3301021330111221) |
| `timeouts.update` | [timeouts.update](resources--dns_load_balancer--reference--group-001.md#canonical-0202332303001212-0102232210022310-1323032320330232-0101312222313003-3220130320013100-3020232131323031-3102111231030133-3331031132311221) |

<a id="canonical-1331333111332213-1223110122222221-0233032001020032-3111112321230213-2031021332213010-0103112212112311-0312213310111033-3233111331202301"></a>

## Next pages — Property reference / 203302123010 / 13

- [fallback_pool](resources--dns_load_balancer--reference--group-001.md#canonical-1212210003110211-3001201121231212-1332220221111323-1100201000223220-2301012002013202-2331013101210221-2212310112013322-1223121320211020)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [timeouts](resources--dns_load_balancer--reference--group-001.md#canonical-3020101101220100-1331123312200221-0133330233120332-1010123102022230-0221113111221210-1030110202230033-2301033013012302-2010221031233011)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-1212210003110211-3001201121231212-1332220221111323-1100201000223220-2301012002013202-2331013101210221-2212310112013322-1223121320211020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031103203230221-1310230231001220-1331032113230322-3221233203103003-3323300123333313-1020033211212213-0301233011333310-0213000233213133"></a>

## fallback_pool — fallback_pool / 301110112011 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- fallback_pool

<a id="canonical-2121111302210113-0313112310213022-2331130333201223-1111301111113333-1102120003013113-2002212230032123-1122012010121022-1310021230101303"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
fallback_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301322030032213-3200103023131022-0010330130131310-1101111011113312-3302210222101000-0121112323131113-3000100022022233-1301103322332101"></a>

## Direct properties — fallback_pool / 301110112011 / 3

<a id="canonical-2330132021200100-1013021123303300-0332203022303320-0323100130231112-0213233300302203-0223323131000033-1221233312021313-1002322321003101"></a>

<a id="canonical-0121123313022131-3211322003301231-3313101313003012-0221132133030302-0030310100312333-3202102330221110-3120310022200321-2100212001220312"></a>

## name property — fallback_pool / 301110112011 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2303303121302002-0230003312332323-1311132220002001-0222002011213200-2122102031133222-1300202021303220-3020100102311201-1012112221010300"></a>

<a id="canonical-1101021300122220-2130201120303020-1323333230221002-3313312223212011-1223321131013010-2203223300021010-3221301320002333-1320333202100100"></a>

## namespace property — fallback_pool / 301110112011 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1233013202002220-0103202213211310-2222203110032332-0211033231230122-2300320023230001-1211102302112200-3233202031312130-2300110203330322"></a>

<a id="canonical-2012322010322120-0112131013213001-1332100002101122-3101312312333310-2220330010001031-0001310203003000-1212231211233322-2130213231230200"></a>

## tenant property — fallback_pool / 301110112011 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1121020111110103-2232122111323201-1022110331030330-3031000001231103-3331311300210210-0201330201102103-1300332120213030-2010202301333311"></a>

## Next pages — fallback_pool / 301110112011 / 7

- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203012030201023-2300322122032011-2121022201232033-3223120202132313-0123123212232331-0213202011333023-2313203030302130-0020201102002032"></a>

## response_cache — response_cache / 121203031113 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- response_cache

<a id="canonical-1011311312311033-2123301323200022-0232303221333113-1332230110202230-2130122221301303-1131200232012300-3031111303033220-1031001103322312"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for response cache.

Upstream description:

Response Cache x-required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_response_cache_parameters",
    "response_cache_parameters"),
  validators.ConflictingObjectAttributes("disable_spec",
    "response_cache_parameters")}
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
  "x-ves-oneof-field-response_cache_parameters_choice": "[\"default_response_cache_parameters\",\"disable\",\"response_cache_parameters\"]"
}
```

Terraform syntax:

```terraform
response_cache {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200311323002031-0132122201300212-3330303110303331-1112112321131030-3230302100333020-0223212321331030-0313201123132010-3331123033013201"></a>

## Direct properties — response_cache / 121203031113 / 3

- [default_response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-2032203223131302-1320313200321200-1233201312003033-3331221203011330-0211113110211120-0010000031210201-1303120232222020-1030012203022232): complete subsection reference.

- [disable_spec](resources--dns_load_balancer--reference--group-001.md#canonical-1223303111031130-1303220113211001-0001103012223332-1322021003002032-0330331220231112-2121222020300232-3121212112232300-3101123133000323): complete subsection reference.

- [response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-3331112003223121-2003203323102330-3202003210020000-1330302321221011-2010113030232121-1220011210121321-1300100132020112-3010132200032130): complete subsection reference.

<a id="canonical-3002003033033223-1331323301122110-1133203100223202-3332311311202033-3231202002133022-1000211220010133-1210012131222133-2232122011001111"></a>

## Next pages — response_cache / 121203031113 / 4

- [response_cache.default_response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-2032203223131302-1320313200321200-1233201312003033-3331221203011330-0211113110211120-0010000031210201-1303120232222020-1030012203022232)
- [response_cache.disable_spec](resources--dns_load_balancer--reference--group-001.md#canonical-1223303111031130-1303220113211001-0001103012223332-1322021003002032-0330331220231112-2121222020300232-3121212112232300-3101123133000323)
- [response_cache.response_cache_parameters](resources--dns_load_balancer--reference--group-001.md#canonical-3331112003223121-2003203323102330-3202003210020000-1330302321221011-2010113030232121-1220011210121321-1300100132020112-3010132200032130)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-2032203223131302-1320313200321200-1233201312003033-3331221203011330-0211113110211120-0010000031210201-1303120232222020-1030012203022232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300002203032113-3103103311111000-3322032032030303-2012113323100302-3111330223333322-1120203210023210-0213210211032333-1213131312010021"></a>

## response_cache.default_response_cache_parameters — default_response_cache_parameters / 101010332301 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033)
- response_cache.default_response_cache_parameters

<a id="canonical-2112120233033220-2221322023023101-3112111310331101-0100102222100120-3020330010313321-1331003333002302-0133303000332003-0331113310223200"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_response_cache_parameters = {}
```

<a id="canonical-1322323312011031-1133312100203120-3313112111132022-3231220300232231-0213011001111203-1323103021011012-0100121303133311-3323031022133222"></a>

## Direct properties — default_response_cache_parameters / 101010332301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212202200231002-3133301313301211-3230010013313330-0032113230211312-3303202223201103-2333000010122012-3021002302033213-1113333333231330"></a>

## Next pages — default_response_cache_parameters / 101010332301 / 4

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-1223303111031130-1303220113211001-0001103012223332-1322021003002032-0330331220231112-2121222020300232-3121212112232300-3101123133000323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003211311020332-3003110010233002-2332233130101023-0333020031330122-1233103020211313-0112330013320101-3023102200100322-1032301332122222"></a>

## response_cache.disable_spec — disable_spec / 303212101312 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033)
- response_cache.disable_spec

<a id="canonical-2211301201022123-2312012131300003-2030013232022022-2220331222100230-1302230233133222-1200300323202003-2331032130333333-0011330021032001"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-3031333000010021-1011202100333320-3102222123332030-0210033210200313-1323330313021210-0122330300003321-1033030220110213-3123012210320202"></a>

## Direct properties — disable_spec / 303212101312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013222123213301-3003320023003310-2310111002010223-1211230300102300-1020302130033313-0212011220113332-2103221332121212-3012023021101332"></a>

## Next pages — disable_spec / 303212101312 / 4

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-3331112003223121-2003203323102330-3202003210020000-1330302321221011-2010113030232121-1220011210121321-1300100132020112-3010132200032130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023132113202322-2213323122202221-1332332200103303-2223010021220303-3112302213012120-2001230002102333-1133230230233332-3122012012103110"></a>

## response_cache.response_cache_parameters — response_cache_parameters / 010021210001 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033)
- response_cache.response_cache_parameters

<a id="canonical-2233312312221021-0313321001033022-1123130330313113-1010102312220030-0110332303303302-2230011023020000-3323220023200000-3212033001030232"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for response cache parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cache_cidr_ipv6")}
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
response_cache_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202223021300311-0231311333123320-1300132302101330-2120022030313232-1000111113203330-0200113033102000-2120233203102333-2103023132011202"></a>

## Direct properties — response_cache_parameters / 010021210001 / 3

<a id="canonical-0202210202022213-2103003032303231-1122000011131320-3311212013121310-2200212202120223-0002212121002233-1012323033123120-1302201102211101"></a>

<a id="canonical-3020003110021123-1230301222331300-0223112223111131-3022302301212132-0111112331321122-1230031103113222-2120231323200213-1013200202001133"></a>

## cache_cidr_ipv4 property — response_cache_parameters / 010021210001 / 4

Type: `"number"`. Optional.

Length of CIDR masks used to group IPv4 clients.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0323110302202133-1101332221211010-2220101120303203-0023100001323103-0101102013011300-3313301312232021-0330003201101332-1130133101113303"></a>

<a id="canonical-0131220130120301-2110333320122322-0231121131302032-2213312012022221-3100230202200232-0112123222031120-1013313010332123-1130200112023013"></a>

## cache_cidr_ipv6 property — response_cache_parameters / 010021210001 / 5

Type: `"number"`. Optional.

Length of CIDR masks used to group IPv6 clients.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 128),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3223213003233022-3103032223221100-3001202223322122-1302233123011330-0302110232232120-1221122333132013-2232210020103330-1303310300203332"></a>

<a id="canonical-0031201012102111-3333103022230300-2131031320000210-2330110203310233-1001232101212323-0220030322232222-1210231131313132-0032032102322133"></a>

## cache_ttl property — response_cache_parameters / 010021210001 / 6

Type: `"number"`. Optional.

TTL. TTL for response cache.

Upstream description:

TTL for response cache.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(0),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1001301201132333-0233221231100200-0220223030112012-2322012102011101-2012221233023230-1131222333213312-1223000131131233-1011203020313212"></a>

## Next pages — response_cache_parameters / 010021210001 / 7

- [response_cache](resources--dns_load_balancer--reference--group-001.md#canonical-2222120002121230-3110032110303221-1121221300210031-2320302310233001-2230301012221321-3211131131121010-1203210022321033-2021103202112033)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213012002300022-3000221313311221-0102212300220300-0222320032302223-3323231212010233-1301123131100100-1023023121210002-2132100303211322"></a>

## rule_list — rule_list / 222002113132 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- rule_list

<a id="canonical-1011030300133330-1232233331000203-0001220301010012-2022333222201231-3120200231121310-3013230200232231-2331121301011003-2302030232030212"></a>

Type: `"object"`. single nested block, Optional.

Load Balancing Rule List. List of the Load Balancing Rules.

Upstream description:

List of the Load Balancing Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222133301112333-3023031203133033-2313010002212232-1102322031231012-3333200033330233-3213121100300213-2222303113203301-1322003001131313"></a>

## Direct properties — rule_list / 222002113132 / 3

- [rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121): complete subsection reference.

<a id="canonical-3002122210231123-0133110311211323-0221033220113132-2121122221023222-0302100010330233-2203121220323211-0112002031021121-3233123312210031"></a>

## Next pages — rule_list / 222002113132 / 4

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100331030131203-0013123120212122-0120032220311111-1311020131232232-0210101301101330-3103110121231000-0111201030000202-3123313113212113"></a>

## rule_list.rules — rules / 221222000130 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- rule_list.rules

<a id="canonical-1232202012132220-2130223201022300-3300032011330232-0131322223001322-3231013201023211-3110301303321132-3323003320222200-0020020011100120"></a>

Type: `"object"`. list nested block, Optional.

Load Balancing Rules. Rules to perform load balancing.

Upstream description:

Rules to perform load balancing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("score"),
  validators.ConflictingListObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("asn_list",
    "geo_location_label_selector"),
  validators.ConflictingListObjectAttributes("asn_list",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_list",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "geo_location_label_selector"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_matcher",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "geo_location_set"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("geo_location_label_selector",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("geo_location_set",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("geo_location_set",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("ip_prefix_list",
    "ip_prefix_set")}
```

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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031002032330231-1202233301000231-0330330122021120-2330100120312120-2211102012301130-2231202221003011-3102230101201121-0330321000022333"></a>

## Direct properties — rules / 221222000130 / 3

- [asn_list](resources--dns_load_balancer--reference--group-001.md#canonical-3321223222330000-2012300200211201-0211021022002223-1323100222100113-2130330102122120-2210202021030101-0231212100212101-1013103112330113): complete subsection reference.

- [asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-2123003121333211-2302013212313302-1123210233230220-3220010200311110-3122322222313332-2021201201110220-2000103323332121-1112232302000201): complete subsection reference.

- [geo_location_label_selector](resources--dns_load_balancer--reference--group-001.md#canonical-0332233333311132-0033231202303020-1202301012223023-3121100330111001-3022332013333230-1320300121230223-3133011133030303-3030003002200210): complete subsection reference.

- [geo_location_set](resources--dns_load_balancer--reference--group-001.md#canonical-3332303023312132-1112123121210330-1320200331132112-1020010023231010-2023130131311001-1313230010111112-1020201132121303-0002311310232110): complete subsection reference.

- [ip_prefix_list](resources--dns_load_balancer--reference--group-001.md#canonical-2233311212333031-2302030022200111-2001221012113113-2222313300031001-0220201102123230-1012113000212222-3111103321100130-1011110022233301): complete subsection reference.

- [ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-1021103121300020-3113310130101200-2121330213301130-1123000102230021-2310311211332002-0032132003201311-1312223122232100-2322120232033331): complete subsection reference.

- [pool](resources--dns_load_balancer--reference--group-001.md#canonical-1010331032221303-1133022320120002-3302300021123021-1222330200033032-2302112320201120-1301102220312223-3020300022311330-2220010221020021): complete subsection reference.

<a id="canonical-3013123131003322-0101020012331332-2233131131013033-3010110300120002-1320203122122322-3022211232201030-0022320203232133-0000033030020300"></a>

<a id="canonical-3023032121312232-2322133103212313-1333031030132202-0231232223232123-0122223110321132-3021132300112030-0202023130122002-1001313300233201"></a>

## score property — rules / 221222000130 / 4

Type: `"number"`. Optional.

When multiple load balancing rules match a query, the one with the highest score is chosen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32767),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2212112033132003-1131310333021220-1323203300121120-1100223302232023-1303121312100130-1332220330311120-2232122231302100-1020030101321032"></a>

## Next pages — rules / 221222000130 / 5

- [rule_list.rules.asn_list](resources--dns_load_balancer--reference--group-001.md#canonical-3321223222330000-2012300200211201-0211021022002223-1323100222100113-2130330102122120-2210202021030101-0231212100212101-1013103112330113)
- [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-2123003121333211-2302013212313302-1123210233230220-3220010200311110-3122322222313332-2021201201110220-2000103323332121-1112232302000201)
- [rule_list.rules.geo_location_label_selector](resources--dns_load_balancer--reference--group-001.md#canonical-0332233333311132-0033231202303020-1202301012223023-3121100330111001-3022332013333230-1320300121230223-3133011133030303-3030003002200210)
- [rule_list.rules.geo_location_set](resources--dns_load_balancer--reference--group-001.md#canonical-3332303023312132-1112123121210330-1320200331132112-1020010023231010-2023130131311001-1313230010111112-1020201132121303-0002311310232110)
- [rule_list.rules.ip_prefix_list](resources--dns_load_balancer--reference--group-001.md#canonical-2233311212333031-2302030022200111-2001221012113113-2222313300031001-0220201102123230-1012113000212222-3111103321100130-1011110022233301)
- [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-1021103121300020-3113310130101200-2121330213301130-1123000102230021-2310311211332002-0032132003201311-1312223122232100-2322120232033331)
- [rule_list.rules.pool](resources--dns_load_balancer--reference--group-001.md#canonical-1010331032221303-1133022320120002-3302300021123021-1222330200033032-2302112320201120-1301102220312223-3020300022311330-2220010221020021)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-3321223222330000-2012300200211201-0211021022002223-1323100222100113-2130330102122120-2210202021030101-0231212100212101-1013103112330113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321300210121221-1030331203200122-1332030020120222-3121201012033022-1022231233230212-0002302101311003-3321013030302113-2120320130233110"></a>

## rule_list.rules.asn_list — asn_list / 301113223100 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- rule_list.rules.asn_list

<a id="canonical-0303111100131233-2122100220230301-0310111132022321-1223201023322220-3330310202122313-0002120310233102-3132313013000202-1333000212211330"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222032033132003-3113300013003120-0033013021222221-1313033320013301-1121113011111100-2231032010022221-3013331022103000-0310033131333031"></a>

## Direct properties — asn_list / 301113223100 / 3

<a id="canonical-0032202121213011-2011012031103101-1231031013102002-1132001010011012-3112131101023231-1333310002121002-3122101001103311-1010300002100001"></a>

<a id="canonical-3011103230331323-2211233201333322-3010232022031100-2001002313231200-2220130203031231-3002131203031303-2123131102033320-1122131112001002"></a>

## as_numbers property — asn_list / 301113223100 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-2030103121013203-0302120203121310-3130011112112102-3102010333233121-2220032311332021-2130211223002230-1223301311113323-0310233200323233"></a>

## Next pages — asn_list / 301113223100 / 5

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-2123003121333211-2302013212313302-1123210233230220-3220010200311110-3122322222313332-2021201201110220-2000103323332121-1112232302000201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102320332121023-0030000300233220-2131033012333320-0003323033312201-0200202303030230-3022201220033222-2011223322310233-3303210113120303"></a>

## rule_list.rules.asn_matcher — asn_matcher / 303320230321 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- rule_list.rules.asn_matcher

<a id="canonical-2101022121002302-1311113101123303-0020233213202023-3132012011112233-1130013220313033-2300031000232312-0100202331111230-0203333311013302"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212301331022111-0131010321200232-2110311201130232-3120031002102123-0023021301113013-1300122101101111-0111222200230103-3222212103330031"></a>

## Direct properties — asn_matcher / 303320230321 / 3

- [asn_sets](resources--dns_load_balancer--reference--group-001.md#canonical-0323211003323103-0210122232312120-0100001033202233-1131103311011332-3110113300132001-1120133022220202-0332030332233031-2021013003320122): complete subsection reference.

<a id="canonical-0211220332223022-0012323000211100-3112102200010210-2112220200221102-2122210011111030-3032212012132011-3210012101230330-3303322110301111"></a>

## Next pages — asn_matcher / 303320230321 / 4

- [rule_list.rules.asn_matcher.asn_sets](resources--dns_load_balancer--reference--group-001.md#canonical-0323211003323103-0210122232312120-0100001033202233-1131103311011332-3110113300132001-1120133022220202-0332030332233031-2021013003320122)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-0323211003323103-0210122232312120-0100001033202233-1131103311011332-3110113300132001-1120133022220202-0332030332233031-2021013003320122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101213202023302-3030211322120122-1033300331101133-2113202010311302-3202322031020021-1110133123113020-1012232303013230-3230122123201033"></a>

## rule_list.rules.asn_matcher.asn_sets — asn_sets / 132103330121 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-2123003121333211-2302013212313302-1123210233230220-3220010200311110-3122322222313332-2021201201110220-2000103323332121-1112232302000201)
- rule_list.rules.asn_matcher.asn_sets

<a id="canonical-3301100110310203-2302122112312121-0000313331002223-1113022321310213-1220130201320103-3320320302113131-3012023023133232-2103032032002030"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021013313131110-2012232210232232-3121022202200112-1010131032301132-1012230003202212-0123320031000102-2002120132131003-0100333101212032"></a>

## Direct properties — asn_sets / 132103330121 / 3

<a id="canonical-1100212222231111-0023003002312132-2203110031003312-1200221102310321-3013030220322330-2120311133010113-0023033222001103-1101311333023211"></a>

<a id="canonical-3310023223110223-0131013120101321-1331313323331133-0313123221213110-2133032331202230-0222210012310223-1223210222233013-1220131320213012"></a>

## kind property — asn_sets / 132103330121 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0132111231212322-0320011113301001-0313321103333021-3313032013021110-0332133202001023-0023222303130320-2200320021203222-1022302331020220"></a>

<a id="canonical-1033231321223100-3103012233232220-2033023031300031-2312133311222131-1100310213121223-1030103011102031-3123300120200011-0103311212322011"></a>

## name property — asn_sets / 132103330121 / 5

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2133202323022100-1201101332102232-0102313223333103-0133130111022030-1032320320321333-0123201033001311-3200132313211130-3010101030200231"></a>

<a id="canonical-3002211203130021-2300030020030101-0230122230321301-1311110120202230-2332213122111201-0132330200311300-0233223230000323-2023001212322222"></a>

## namespace property — asn_sets / 132103330121 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3303321103013203-3222222333201233-3031331223231130-1330333123311123-1300003222223310-2221031122033100-1313333103233133-3023323003013323"></a>

<a id="canonical-1220322212002103-2101000121333210-1102111120023231-3320201111003123-0220120133201100-0322101103010230-0232122223333322-3012212331231231"></a>

## tenant property — asn_sets / 132103330121 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3122313020313311-2133302233031212-1020101233103201-1010130200231102-2122223202212023-2122302002311021-0102300031333102-1202303333322003"></a>

<a id="canonical-2002031010032212-2111021120223202-0120032301331310-3002333232210002-1032110001120003-0121121302223221-3111120322003310-0101202002030213"></a>

## uid property — asn_sets / 132103330121 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2122320111132023-0212022001112223-3132112222130200-1011303302111100-0313322131020113-1102201231320330-2220002130030130-1203012030002123"></a>

## Next pages — asn_sets / 132103330121 / 9

- [rule_list.rules.asn_matcher](resources--dns_load_balancer--reference--group-001.md#canonical-2123003121333211-2302013212313302-1123210233230220-3220010200311110-3122322222313332-2021201201110220-2000103323332121-1112232302000201)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-0332233333311132-0033231202303020-1202301012223023-3121100330111001-3022332013333230-1320300121230223-3133011133030303-3030003002200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122200233330101-2210222321233322-2133023121132220-0013120100111201-1223211023210202-2231320200330202-2230010032123233-1233321313023230"></a>

## rule_list.rules.geo_location_label_selector — geo_location_label_selector / 023202112322 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- rule_list.rules.geo_location_label_selector

<a id="canonical-3123111103031123-0221320330232113-0131100233131223-1111100000302020-3011331030102313-2302222222110012-0331121022123200-1013001213321201"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
geo_location_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102323110122303-0011320322311122-0201300033130010-1311102130100122-0111121222320210-0133133203122333-3113311303110330-1131130333023223"></a>

## Direct properties — geo_location_label_selector / 023202112322 / 3

<a id="canonical-1332000101300231-1331121130200010-3313233120320301-2301302031230311-3330222130212200-3312333233132301-1310322002332000-2030320230110123"></a>

<a id="canonical-3331003210300130-1120221013113313-0221102120330032-1220331221323002-0031301002320103-1303001301220203-2311022311020232-1022323233113031"></a>

## expressions property — geo_location_label_selector / 023202112322 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0210322300012110-3030310000233030-1301031313031113-2002320201123130-2020303031220000-3201313213322330-3010230210021320-1110123322000011"></a>

## Next pages — geo_location_label_selector / 023202112322 / 5

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-3332303023312132-1112123121210330-1320200331132112-1020010023231010-2023130131311001-1313230010111112-1020201132121303-0002311310232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122130311122013-2210033123201010-2001330101321212-3320212203130232-2203121000320120-0031122130311301-2220231212011332-1012012233330100"></a>

## rule_list.rules.geo_location_set — geo_location_set / 022132321033 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- rule_list.rules.geo_location_set

<a id="canonical-3123330103223022-2321232202022103-1222012011122113-1330321303102313-0330031332130121-1332203012012231-1332310202113223-0121032010310102"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
geo_location_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021010233331223-1303130113030333-3230030332003121-2002302221312112-2213020003211012-0100101002110020-1133032123030222-3130100321003020"></a>

## Direct properties — geo_location_set / 022132321033 / 3

<a id="canonical-2313113032112230-2010203000030113-1131123330120132-1021001101103221-2320102233213131-2233012201201300-2020021112121320-1102110111330001"></a>

<a id="canonical-3201223122002222-1222203112313323-0001122002321120-1013010210333002-2002033313112100-0000022030213000-2011102001330022-3213022023113113"></a>

## name property — geo_location_set / 022132321033 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2222121213133231-3122123333311332-1112322132311111-1210202302213110-0102000002223330-2130211330211130-1103200311211101-2211201013113023"></a>

<a id="canonical-2211123322113123-1133132303021233-0201102131003103-3203210013202213-2333102212301220-2332031020310312-0002112100023220-2201133322211200"></a>

## namespace property — geo_location_set / 022132321033 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2302233213000123-1313033100323103-3010133011031120-0010031132322102-1003313131003100-1301133212213001-3331002010213012-0020002300210132"></a>

<a id="canonical-0003121131303223-1113103323310320-3123123212130130-2312210322212110-2333323111103011-1131133230212112-3010221131122200-0203120313312020"></a>

## tenant property — geo_location_set / 022132321033 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1211033213120013-3133033321013031-2111122332322023-0023323200231232-2230012012200111-1301033000323132-3032230132123012-0312210203022002"></a>

## Next pages — geo_location_set / 022132321033 / 7

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-2233311212333031-2302030022200111-2001221012113113-2222313300031001-0220201102123230-1012113000212222-3111103321100130-1011110022233301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012211312023110-2310001122323132-2321003110332122-2332031100330313-2323000013331020-0230312333302331-0302100112103000-0320311130031213"></a>

## rule_list.rules.ip_prefix_list — ip_prefix_list / 131010113302 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- rule_list.rules.ip_prefix_list

<a id="canonical-1001031001310010-2120022232102133-0231311313310222-1102331201131022-2223031331302113-0200001220113132-1130333121231230-2332032200113030"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311212212021111-1110300212321323-0002212012131113-2300102230131331-3223311211002010-3022003312112000-2012002023130310-1021012103120131"></a>

## Direct properties — ip_prefix_list / 131010113302 / 3

<a id="canonical-0222023130020230-3321233233300013-1213121123301323-3202233223233131-3232110203203013-0123312230302332-1012003001130103-3201132002232322"></a>

<a id="canonical-3331110213323301-1310331011120230-1332122232302333-1033222221310111-0101010132212133-2013130023012003-2321330130231032-3011303022101133"></a>

## invert_match property — ip_prefix_list / 131010113302 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-3220120310211223-3013222332010101-0123213102022013-0313232102010312-3122031230322322-0330323331321331-2321112122012000-1322113003332230"></a>

<a id="canonical-3132210001012300-0001330001333101-1110322200132103-3011310301021310-3310012320002103-0101222003231233-2332332211233310-1221210331013320"></a>

## ip_prefixes property — ip_prefix_list / 131010113302 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-1201212231130301-1101103220102110-2122303212300102-3020003130101123-1111031001010301-2221000212011220-0313312010010300-1130320303022110"></a>

## Next pages — ip_prefix_list / 131010113302 / 6

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-1021103121300020-3113310130101200-2121330213301130-1123000102230021-2310311211332002-0032132003201311-1312223122232100-2322120232033331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201102212020113-1110121011021213-2011102210022332-3122001212000320-2022233322202210-3122100303313231-3113012230312333-0320300201101022"></a>

## rule_list.rules.ip_prefix_set — ip_prefix_set / 223031203230 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- rule_list.rules.ip_prefix_set

<a id="canonical-0202201303013120-3233121221112022-0011013021113003-2203203132032322-2210112213221211-2122033213322303-0231012231123223-3001332230222103"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130012323201020-1313231232003200-2003232112021113-0333231230110300-1023333210110211-1320130313222110-3332121300313120-2320132331102303"></a>

## Direct properties — ip_prefix_set / 223031203230 / 3

<a id="canonical-3132211120001130-1022202301020123-0313112312203031-1212203003323112-2230200223201121-2123030203131313-1122203031223133-2231000230202122"></a>

<a id="canonical-3031103220022300-0310230033301222-3312313020123123-0132222303303110-0023011032100210-3103021132021113-3311300022021113-0321033323332221"></a>

## invert_matcher property — ip_prefix_set / 223031203230 / 4

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--dns_load_balancer--reference--group-001.md#canonical-0211331202330210-3112232312031201-1310232303221210-3132022212222010-3320200310231113-2011001323333202-0331111201332310-0021012113021232): complete subsection reference.

<a id="canonical-0330200313133023-3213323113101130-1023132101103003-0331121230230210-3001223001231200-2122033022233102-1031123132103200-3310030220230221"></a>

## Next pages — ip_prefix_set / 223031203230 / 5

- [rule_list.rules.ip_prefix_set.prefix_sets](resources--dns_load_balancer--reference--group-001.md#canonical-0211331202330210-3112232312031201-1310232303221210-3132022212222010-3320200310231113-2011001323333202-0331111201332310-0021012113021232)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-0211331202330210-3112232312031201-1310232303221210-3132022212222010-3320200310231113-2011001323333202-0331111201332310-0021012113021232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203130112233122-2312210311022010-1233213130301030-3022133213203030-0112110330200032-2322001213002013-3301221202020100-2103120011031210"></a>

## rule_list.rules.ip_prefix_set.prefix_sets — prefix_sets / 203200103201 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-1021103121300020-3113310130101200-2121330213301130-1123000102230021-2310311211332002-0032132003201311-1312223122232100-2322120232033331)
- rule_list.rules.ip_prefix_set.prefix_sets

<a id="canonical-1212210332313011-0123032000322223-2003321301100110-3210002302123301-3031002131320331-1000023121132311-2311303110130021-0201030303123100"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013111322011211-1022101311130133-3032331123021222-2110122311332130-3320102320323233-3122320101100022-2011103121233023-3030332220322331"></a>

## Direct properties — prefix_sets / 203200103201 / 3

<a id="canonical-3000333110122100-0123003003031231-1300203110222212-2102103303131032-0022320033231033-0030120100222313-3130032201313111-0032032331113022"></a>

<a id="canonical-0302211212213202-2101332002320302-1222101311310101-2320002011131331-2121103333102222-2321303133301222-0112122233302023-1031231223023222"></a>

## kind property — prefix_sets / 203200103201 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2233221130313231-2211302030013012-0311233133030320-3223220000323212-1300033121333330-2023000001022110-3223020012310210-1131130130321111"></a>

<a id="canonical-3321323201323123-1112333330120301-0003013231202313-0200030000123323-2213213010123030-0010230123032021-0332213003320011-0233330231213312"></a>

## name property — prefix_sets / 203200103201 / 5

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3112323112003120-3332301022130301-3210302121201321-1113011211303021-3233220223003130-1002010333203200-2103112231010232-2101113231133202"></a>

<a id="canonical-3012203121123330-0102212233103000-1123123221110031-3202220000222031-3030033220301320-1102133222300033-3030213021302032-0311202002101131"></a>

## namespace property — prefix_sets / 203200103201 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0102212213013202-3322032003111333-3111013002300013-1221322230211222-3322200022213200-3110312331212020-2101111211220320-3230113100023013"></a>

<a id="canonical-0133132010213213-0002310332230001-1223233020033313-1030001033120103-3322323122032213-3003300122113022-2201313201033210-0032223221012201"></a>

## tenant property — prefix_sets / 203200103201 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1331130302001233-0112011332123103-0333013013212012-2110100133031013-3112011331030101-3221311131203102-1320330031102023-2103131121200230"></a>

<a id="canonical-2003331200121302-0001310132022023-2010032232200322-2202111012231120-1113111323210311-0033323111132333-2332113022002032-3013003120122313"></a>

## uid property — prefix_sets / 203200103201 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1320112121001110-3332121113332031-1223201321233303-1130201302300312-1301221210022200-0310332020230103-1221112123001131-0222121133022233"></a>

## Next pages — prefix_sets / 203200103201 / 9

- [rule_list.rules.ip_prefix_set](resources--dns_load_balancer--reference--group-001.md#canonical-1021103121300020-3113310130101200-2121330213301130-1123000102230021-2310311211332002-0032132003201311-1312223122232100-2322120232033331)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-1010331032221303-1133022320120002-3302300021123021-1222330200033032-2302112320201120-1301102220312223-3020300022311330-2220010221020021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330130022011033-3302313213212332-0313232212120100-2120000002130101-2211003312201020-2321012031133133-2002103323202132-1203321300033123"></a>

## rule_list.rules.pool — pool / 332101312112 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [rule_list](resources--dns_load_balancer--reference--group-001.md#canonical-3123001011030103-3030021310200030-2113331102010132-0303130311130011-0310303333022133-2121102212302031-2010301021020111-3112210032132123)
- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- rule_list.rules.pool

<a id="canonical-0311311200123322-2031130100003101-0210121203003311-1223033133301232-2111313131203223-0032133012021310-2131030323023333-2300012203001022"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222033320102031-3022101300313103-1330222102110312-3013332230210002-0232200213333303-1212012202002320-0010201000010021-2230032303130332"></a>

## Direct properties — pool / 332101312112 / 3

<a id="canonical-3110133132013010-2112232222330100-0203232212112032-3013211223223333-3100201120001120-0300113311320112-0202330321222300-1011321013213302"></a>

<a id="canonical-0030110031102132-1232330222333302-3231002021313230-1120310021021123-1112310021201022-0010223231331321-3133232012000003-0211233321231112"></a>

## name property — pool / 332101312112 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1022223310210010-3100303200132012-0331210012110120-3312021313002202-3033320103000212-2223303123332023-0013322000131021-1002012200133120"></a>

<a id="canonical-3010123111323222-1300123233132100-3310113220101002-1113332203321213-1123030302320122-3103321110011301-3210123030000030-0030133102302132"></a>

## namespace property — pool / 332101312112 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0320113212331333-2221120320100323-2313121030031022-0213200233231300-2313121001113211-2230132201010310-2202130231230120-3100112232313213"></a>

<a id="canonical-0020102232223301-0020230303012022-1002303210312322-2331002000120312-1233010303113021-0121033203103320-3331200302103132-0013310302202122"></a>

## tenant property — pool / 332101312112 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0202310320023023-3010020020321222-0023011211020322-2010232103002133-3202300330110220-3300210023232120-2330101331333100-0031200133323222"></a>

## Next pages — pool / 332101312112 / 7

- [rule_list.rules](resources--dns_load_balancer--reference--group-001.md#canonical-2321132022133030-1332131221201231-2113100122130112-2110030202103300-2121322011013302-0211112102122231-1122101301213330-2223002010100121)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)

<a id="canonical-3020101101220100-1331123312200221-0133330233120332-1010123102022230-0221113111221210-1030110202230033-2301033013012302-2010221031233011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333231220302212-0112003301320121-3133310223312223-0301320122211330-2232101000223032-3202202102001333-1300220122230210-2332010232330200"></a>

## timeouts — timeouts / 033101013300 / 2

Breadcrumbs:

- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- timeouts

<a id="canonical-3132030131303210-3312033122111002-1003030222230032-3321023112103313-3013201022023211-2302123103202221-2100130001101131-0101123000330111"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333112030102330-0200301233203302-2022001320102130-2223222130301013-1321021103203031-2121120220332310-0000032132330311-0031011010332210"></a>

## Direct properties — timeouts / 033101013300 / 3

<a id="canonical-0011231111201302-3210101312233030-0230123111330201-2003121113112333-1011011100131221-0023020101022203-2002103201333232-3332111311112100"></a>

<a id="canonical-1202200313111101-2122000131332132-0312233322332031-1210332010020102-0122132123312012-1131020022000202-1002010001101120-2102133230200003"></a>

## create property — timeouts / 033101013300 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0220032032131323-0022003201011123-1320121111032003-2000211322010133-1003300210330011-1013111221233310-0102320031221230-2023111102301301"></a>

<a id="canonical-3100113220022020-1110112320131201-1113102223022123-0130210130023321-2331031221200213-0232222102001303-2233123022233120-1220020020332300"></a>

## delete property — timeouts / 033101013300 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0310310332303000-3333103323002033-1311130022210202-2201030320111233-0220222202100033-0330100020111323-1102203002101313-3301021330111221"></a>

<a id="canonical-2022031101023112-3002332111100123-1202113211222022-3103111121233123-2230202000331112-1133011302301033-2113000000202013-0132131213111312"></a>

## read property — timeouts / 033101013300 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0202332303001212-0102232210022310-1323032320330232-0101312222313003-3220130320013100-3020232131323031-3102111231030133-3331031132311221"></a>

<a id="canonical-0321210130202001-2122131331101010-0313301131102320-1010133331103003-1002330002112000-3322332102201033-3100112233123310-2332030020010130"></a>

## update property — timeouts / 033101013300 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3032201332130103-3131230303320233-0302303121032230-3011001023030111-3332301100333223-1021323011231113-3221122020133110-1212133310201311"></a>

## Next pages — timeouts / 033101013300 / 8

- [Property reference](resources--dns_load_balancer--reference--group-001.md#canonical-3321030232020002-0011103311203300-2121000022123101-3221311213012322-1033303302120031-1011000133302122-3312210311123110-2223313220310031)
- [xcsh_dns_load_balancer](../resources/dns_load_balancer.md#canonical-0033303122332020-0032121023333003-1311312102200303-3203210001332002-1223332123021131-3313033203131000-0232222020323020-2121001221232233)
