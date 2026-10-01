---
page_title: "xcsh_enhanced_firewall_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy reference."
---

# xcsh_enhanced_firewall_policy reference

<a id="canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213323302312002-3211313033323102-0031010220202312-0002001011232311-0302033020320002-0303133310010102-1302310222300100-1202132322030122"></a>

## Property reference — Property reference / 302331202321 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- Property reference

<a id="canonical-0022210003330111-2211220313113022-2310030122030020-2331302303302000-3333303101103112-2202202003103011-2302000011230212-2231032310302220"></a>

## Direct properties — Property reference / 302331202321 / 3

- [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1102131223103231-1221212033311132-2212320133110100-2123233220030223-0320001123233230-0220122332332022-1031011132232112-2000130111010012): complete subsection reference.

- [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1020330031030313-2121223312230100-0130303020203031-0212211201300200-3313233321023213-2031002212123211-2000302131222100-2133000113223233): complete subsection reference.

- [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0233101011203202-2131101223002310-1303102211012323-0203123013122133-0310123213033203-2230322231320031-3022313211330010-3002230320203231): complete subsection reference.

<a id="canonical-2303121323101112-0310222200213212-0321113310303121-2300332231013321-1010223323231220-2103002220032112-2323103133320021-1132220202103310"></a>

<a id="canonical-1222030033331330-0303320330132222-3002211212311232-0001321133011303-2221131223202312-2330110203002201-0220120100220322-2021123133233313"></a>

## annotations property — Property reference / 302331202321 / 4

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

- [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2330302201311220-0033000223123322-0002331233031022-1003033033120002-2231321310203322-3303012011303311-0033323313131023-1001131213330232): complete subsection reference.

- [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3030320201232003-3313202120101210-1220013203102130-1220122311333002-2103100102202103-3011023011011210-3322030112202311-0000200001100212): complete subsection reference.

- [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2311230322202032-3131231332113201-2233132200232032-2122100001221223-0001203030301022-0101323113220031-3012202300312032-1212221201330030): complete subsection reference.

<a id="canonical-2211001032233033-3222002032012232-0131230022331220-3132333321322210-1000233230222002-0232010232130302-0120312233201201-3330311021020012"></a>

<a id="canonical-1030010300113301-0033302323231101-2100211232033223-0112111202000331-1302033311320103-2223201013303200-0202223311130200-1121130320101311"></a>

## description property — Property reference / 302331202321 / 5

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

<a id="canonical-1132032122020211-1310222331113221-0201100313232001-0113330212301323-1221132100132032-1331122013313132-2230002103200333-3121322101233300"></a>

<a id="canonical-1123322001013231-0212030232313332-1111010210033213-2202020330231003-2030333003322210-2131130023213101-3130213230313311-3333122230120311"></a>

## disable property — Property reference / 302331202321 / 6

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

<a id="canonical-2011031320201011-1322202132211022-2312321302010021-2130000123131003-2033332133102111-2021312103101021-1131302322221111-1302323001020132"></a>

<a id="canonical-2132130323322213-1102223032131203-0120130123010233-3010332023303013-1322222233323131-2112113031331213-1003010132121103-2123112300010113"></a>

## ID property — Property reference / 302331202321 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1202300011101201-1201121221312133-0203012112010331-3212002211233131-1000302333022203-0322110321110001-0020022333301203-1200221332033301"></a>

<a id="canonical-0313033102130321-0010321320300113-1113233230320231-3231133133310103-2321303013023332-2013000123110313-2022113133201313-3331110310101313"></a>

## labels property — Property reference / 302331202321 / 8

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

<a id="canonical-1211131100130323-0023213233311033-3332123213212211-2003120000003030-1213232332112313-3333132203210022-2203122123100021-2221300002333330"></a>

<a id="canonical-1103303020320010-2332001000012110-0120121113230013-1233303132303000-1331212003031301-3012110312222103-1102132333301322-3113113231111001"></a>

## name property — Property reference / 302331202321 / 9

Type: `"string"`. Required.

Name of the Enhanced Firewall Policy. Must be unique within the namespace.

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

<a id="canonical-1013212313102133-1131201120313110-0102221000301210-3213203023222233-1300331301311223-1310112023103121-0300211011210223-0113210121332220"></a>

<a id="canonical-2030032113112001-3000010110030023-2323033100002230-2301131121010210-2333200112102313-1221212323102300-0310130230210122-1031122130321200"></a>

## namespace property — Property reference / 302331202321 / 10

Type: `"string"`. Required.

Namespace where the Enhanced Firewall Policy is created.

Upstream description:

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

- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233): complete subsection reference.

- [timeouts](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1100132312213200-0032121211122030-0200203311120012-2111203021323202-3323203001323231-1123113002211000-3022113223112313-1110213303100000): complete subsection reference.

<a id="canonical-3333130313330310-1123130023323121-1020310311021321-3233200010203011-3231132122221121-3013033332223131-3233113333320330-2221302003111121"></a>

## All schema paths — Property reference / 302331202321 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allow_all` | [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2133223303012233-1322211203320102-3333133003102201-2032313120001132-1001123310321231-1100102303321221-1221113302131011-3232022011001312) |
| `allowed_destinations` | [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1203212233000123-1212222110033030-0030130000201331-0120300313013213-2023212000302301-2123002113300001-1132331130033312-3031202121032130) |
| `allowed_destinations.prefix` | [allowed_destinations.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2001323111123133-1031322002213021-1031132010013300-3230020233023000-1203012022010201-0032310133310112-0302111333023213-1201101320320120) |
| `allowed_sources` | [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3110333330110202-3323130322121202-0233302221003111-0131312330101032-2031132122120213-1000231110300131-2132103023031221-2111312321202311) |
| `allowed_sources.prefix` | [allowed_sources.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3030101122232321-3221132312033003-0231330331221213-2110200122101302-1000030022020010-2023332220200101-0331110301322113-1013201310211112) |
| `annotations` | [annotations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2303121323101112-0310222200213212-0321113310303121-2300332231013321-1010223323231220-2103002220032112-2323103133320021-1132220202103310) |
| `denied_destinations` | [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0211333302313332-2113222200302323-2023112222130331-3220320200300001-2123203203022203-1120033101200103-0303211000110212-1213112012211022) |
| `denied_destinations.prefix` | [denied_destinations.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0010001310002013-1133111323230000-2000323300131013-0012312212132223-0222200321333013-3323331321031121-0330332332333133-0130003031101123) |
| `denied_sources` | [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2101011013200022-2221031222122132-3013110210110103-2311023232112003-2331021012221301-3022330002032310-3012133331301133-1122203311331021) |
| `denied_sources.prefix` | [denied_sources.prefix](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1010013023322010-0320123102100202-1102222330301021-2232022322220133-1201002031301331-0201002002211331-1002223231131331-2131232211201330) |
| `deny_all` | [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2000113230323112-1011333330331303-3012121131303002-2003020231023220-2230103012233030-0303111131332102-2212223100003222-2101003033331213) |
| `description` | [description](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2211001032233033-3222002032012232-0131230022331220-3132333321322210-1000233230222002-0232010232130302-0120312233201201-3330311021020012) |
| `disable` | [disable](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1132032122020211-1310222331113221-0201100313232001-0113330212301323-1221132100132032-1331122013313132-2230002103200333-3121322101233300) |
| `id` | [ID](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2011031320201011-1322202132211022-2312321302010021-2130000123131003-2033332133102111-2021312103101021-1131302322221111-1302323001020132) |
| `labels` | [labels](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1202300011101201-1201121221312133-0203012112010331-3212002211233131-1000302333022203-0322110321110001-0020022333301203-1200221332033301) |
| `name` | [name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1211131100130323-0023213233311033-3332123213212211-2003120000003030-1213232332112313-3333132203210022-2203122123100021-2221300002333330) |
| `namespace` | [namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1013212313102133-1131201120313110-0102221000301210-3213203023222233-1300331301311223-1310112023103121-0300211011210223-0113210121332220) |
| `rule_list` | [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2202011303310301-1022220123013310-2111000100022230-3210133001102222-1230300111121233-3110212333200313-2202132223011021-3321221100031111) |
| `rule_list.rules` | [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3301100221132102-3103020103003132-0322322333123121-0223003111133303-2130303320203322-0222022000000210-0220312133212132-0231130333021023) |
| `rule_list.rules.advanced_action` | [rule_list.rules.advanced_action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3201112220310003-2012033200331331-0300010133200220-2213032222033203-1010123001110223-3301131212123301-3221202123031322-3023100100310332) |
| `rule_list.rules.advanced_action.action` | [rule_list.rules.advanced_action.action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1221323031312322-3120111223123312-2020020331320120-1321210222332330-0000230232123130-0133301230323031-0021130320203020-2301020332312131) |
| `rule_list.rules.all_destinations` | [rule_list.rules.all_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2123003221131311-1032122200103233-1231003131200012-0311212303111011-2021333200111110-1201112233221103-2011231203231132-0223303320202120) |
| `rule_list.rules.all_sli_vips` | [rule_list.rules.all_sli_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2133123132102300-2131000301031023-1132011201323102-3200312232313331-3330010300133200-0013313303221323-3020201013112330-3120131221020212) |
| `rule_list.rules.all_slo_vips` | [rule_list.rules.all_slo_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1123213030032131-0222020133003300-2202220132133323-1210001032312013-1231010211122030-1111301323123321-3120013111323032-2220321011320331) |
| `rule_list.rules.all_sources` | [rule_list.rules.all_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2310011032123110-3003202032303212-0210233201112220-3122201113332202-0113223131103220-0323212130323030-2221130331302121-2232003130113100) |
| `rule_list.rules.all_tcp_traffic` | [rule_list.rules.all_tcp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1201211210311032-2013132033001000-3212121113221103-3133102010202100-0132200203310000-1220220301001003-2111320133102231-0332231333211211) |
| `rule_list.rules.all_traffic` | [rule_list.rules.all_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0332030211123200-1022030033101211-0031323300320232-0312203332233233-1102210322131111-2222123322330011-1201212212221130-2232013013211023) |
| `rule_list.rules.all_udp_traffic` | [rule_list.rules.all_udp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2102230322013231-0321003201230301-0132113030121222-1310001221002012-0301011100103211-0113233311032022-2003311312030313-3232231032101000) |
| `rule_list.rules.allow` | [rule_list.rules.allow](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1211003102300000-2122203110031310-3202323033232200-1320300202103102-0120230111101220-0211111120330322-2023310222113100-2000330210212320) |
| `rule_list.rules.applications` | [rule_list.rules.applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0111330200023022-0020112202113210-0332022220003121-2332022200032111-2002313133030002-2200213320200220-1103332312010102-3332223322021132) |
| `rule_list.rules.applications.applications` | [rule_list.rules.applications.applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0311021022221301-2200202010232201-1201011301223213-3312233011330011-0101320220133230-2312312323130210-3101101103311330-3323231031120211) |
| `rule_list.rules.deny` | [rule_list.rules.deny](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0322003203330100-1231313032302110-1101303223001322-0120202230100232-0121001131220023-1200133022302230-1030103001012231-1132101132231201) |
| `rule_list.rules.destination_aws_vpc_ids` | [rule_list.rules.destination_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0213223032210131-1212320013013200-0313012232131001-3313221330332113-1203233010301310-2131310112033032-3212220010211130-2000213323132200) |
| `rule_list.rules.destination_aws_vpc_ids.vpc_id` | [rule_list.rules.destination_aws_vpc_ids.vpc_id](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0211031103223322-0233111102011112-0302302122030203-0300333332320101-1110210202210232-3013102211231231-3311223113002321-2110212132012232) |
| `rule_list.rules.destination_ip_prefix_set` | [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0231001002013222-0313002303030213-2313122302120322-2333000023003202-0300213031223012-2133100023200012-3021030310121213-3320301032132321) |
| `rule_list.rules.destination_ip_prefix_set.ref` | [rule_list.rules.destination_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3233211122121213-3032130020011312-1001320222202133-2323223020301223-1011031331330122-2330331200101320-0233301030011012-2222333122122000) |
| `rule_list.rules.destination_ip_prefix_set.ref.kind` | [rule_list.rules.destination_ip_prefix_set.ref.kind](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2030210122233102-3001033001212321-0320211002123333-0312012222231331-0212011203021201-2032102332021002-3233222002323230-0332210300221210) |
| `rule_list.rules.destination_ip_prefix_set.ref.name` | [rule_list.rules.destination_ip_prefix_set.ref.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3233022321133102-3222300032312223-2200221011101020-1331202111333322-2132133333201203-0222300101011300-2210212232030010-1002123001021322) |
| `rule_list.rules.destination_ip_prefix_set.ref.namespace` | [rule_list.rules.destination_ip_prefix_set.ref.namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0022210031222111-2021000223012102-3100111002101330-0303131223033302-2103122220232032-1031220211300202-1001123302133200-1032001212231223) |
| `rule_list.rules.destination_ip_prefix_set.ref.tenant` | [rule_list.rules.destination_ip_prefix_set.ref.tenant](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0132201200111233-1301103220010320-3312110201320012-3303003323022300-3000031231033221-3113301201121113-0223132310223030-3011000300123300) |
| `rule_list.rules.destination_ip_prefix_set.ref.uid` | [rule_list.rules.destination_ip_prefix_set.ref.uid](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3001202010230223-1031231013102223-2021331322331010-0112233230003113-1032320010322112-2003300121123002-1103031233102023-0101010100101133) |
| `rule_list.rules.destination_label_selector` | [rule_list.rules.destination_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1000003213001003-3212130311130220-1230000130000111-0122033020233101-0013131311311323-0233212033320121-3232203113223113-1011102122021310) |
| `rule_list.rules.destination_label_selector.expressions` | [rule_list.rules.destination_label_selector.expressions](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2232201230201211-3230021101020001-0110010132230220-0110232231021022-3220111203330133-2122011330303200-0103322233202202-0230203213121030) |
| `rule_list.rules.destination_prefix_list` | [rule_list.rules.destination_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0102101103220223-3123222220220101-2300023002202103-2003020312001320-1300123120323032-0011201202310020-2132323322132332-0200300223202033) |
| `rule_list.rules.destination_prefix_list.prefixes` | [rule_list.rules.destination_prefix_list.prefixes](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0233213202100012-1321013220210321-0013102120102032-2120312201202120-3111202320011122-3102303011220223-0110220323303100-0332103112312123) |
| `rule_list.rules.insert_service` | [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0330232121321220-2230033233312333-3030021201123032-2201221321323003-2021100022111003-0130030110331202-3202210311110211-0321012003031232) |
| `rule_list.rules.insert_service.nfv_service` | [rule_list.rules.insert_service.nfv_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1221013211011302-2100321323003233-2211311201112230-3121002303202023-2331333313333310-2322010300230322-1200233120012023-3303103101332233) |
| `rule_list.rules.insert_service.nfv_service.name` | [rule_list.rules.insert_service.nfv_service.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1132011222310102-1032331023231302-3020213332021100-3332131112030212-1023103333011222-0222031001212101-0030231002310113-1302121320212131) |
| `rule_list.rules.insert_service.nfv_service.namespace` | [rule_list.rules.insert_service.nfv_service.namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2232101230023231-2032331123210112-0103123133110202-1222122203320310-2131033103223103-3002001001003020-3233332200300302-2002200130212233) |
| `rule_list.rules.insert_service.nfv_service.tenant` | [rule_list.rules.insert_service.nfv_service.tenant](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3032223111010003-2220100131201320-3321232113200031-2303332121322122-3210020211223103-0323102023333010-1010113213331000-3300220222321131) |
| `rule_list.rules.inside_destinations` | [rule_list.rules.inside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1113212031011031-0323210212230020-2323210313201002-0311301330330303-0110200121011213-2020211321200123-0233131303130200-1121033222133011) |
| `rule_list.rules.inside_sources` | [rule_list.rules.inside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3033302001103222-0223010223210112-1220132222311200-2333112332031133-1123122020210010-3021103113030330-2321303110312021-1202310203013203) |
| `rule_list.rules.label_matcher` | [rule_list.rules.label_matcher](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2202330130132032-3100121001123100-1100323122313203-0203322223130222-1023320020032023-2123300020311322-2301020300230321-0202121121330322) |
| `rule_list.rules.label_matcher.keys` | [rule_list.rules.label_matcher.keys](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2103221321301300-1020131231020022-2310020001300321-3000032302233302-1300302303030330-2313211100032012-2330113003103133-1232211001110113) |
| `rule_list.rules.metadata` | [rule_list.rules.metadata](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1222300233003030-1231232022223031-3232323102303012-2311230313231013-3223230112202101-1113323203302113-2310201310122033-0231332110323022) |
| `rule_list.rules.metadata.description_spec` | [rule_list.rules.metadata.description_spec](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2110032023312203-2301302202113301-1122222120030132-1012030211233323-0113112221010021-1230001110003201-1133120302102020-1020032131322122) |
| `rule_list.rules.metadata.name` | [rule_list.rules.metadata.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2021001021313331-3032203330101210-1233113333202023-0221323011000201-2312102223100302-2220303202111010-0211012121302113-2302212322330032) |
| `rule_list.rules.outside_destinations` | [rule_list.rules.outside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3001121111111320-3130333320210103-0011032301121212-2312133000030310-0023333311322231-1112330320033131-3300333132120303-0210112302022102) |
| `rule_list.rules.outside_sources` | [rule_list.rules.outside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2310111011310202-3010121200230221-3012201022011032-0100301232013322-2112303313122000-2132331201321211-3121022232121120-3332312102311200) |
| `rule_list.rules.protocol_port_range` | [rule_list.rules.protocol_port_range](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0213211003321110-1021012310200103-0002301012300313-2123202331002102-3222112200130233-1010131023310333-2203032332330200-3312221121133110) |
| `rule_list.rules.protocol_port_range.port_ranges` | [rule_list.rules.protocol_port_range.port_ranges](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3231202300220023-3213213213122201-3103133303132001-3231100123211211-2110332331221123-2223231332002033-0011112210323013-2201311013122032) |
| `rule_list.rules.protocol_port_range.protocol` | [rule_list.rules.protocol_port_range.protocol](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3121303300313220-3133122332331112-1232213333101322-3110321233320223-3100201202222123-3111122310221122-2132230220102030-1303033301211111) |
| `rule_list.rules.source_aws_vpc_ids` | [rule_list.rules.source_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3302010130321123-2002133022312321-3110023232323332-0013101203233233-2021030211123001-0313333220213033-0220202113230133-0021103012300122) |
| `rule_list.rules.source_aws_vpc_ids.vpc_id` | [rule_list.rules.source_aws_vpc_ids.vpc_id](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3233132012111112-2330303121101013-0233322221022023-0221033001113030-2012101113120221-3000001003111301-0332212102011223-0000323102321320) |
| `rule_list.rules.source_ip_prefix_set` | [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0013330032312112-2220233311203032-3300012232013332-3012323323333113-3332023223201131-1131313003112230-1213213210031322-3132200002213113) |
| `rule_list.rules.source_ip_prefix_set.ref` | [rule_list.rules.source_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3032022013331321-3313121311231100-2221033010111012-0203221223000002-3000123022002223-2323300113320201-3231032130010012-1230312310102031) |
| `rule_list.rules.source_ip_prefix_set.ref.kind` | [rule_list.rules.source_ip_prefix_set.ref.kind](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3310033003300111-1333203022331230-0002110200213012-1331312202223233-3122221010212133-1331211222310322-3031213233100331-3203130011000301) |
| `rule_list.rules.source_ip_prefix_set.ref.name` | [rule_list.rules.source_ip_prefix_set.ref.name](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0310022112123113-0322022121130221-2303302133121333-1113021101320013-0333013130301122-3321011030333000-0310000321210322-0222322331300202) |
| `rule_list.rules.source_ip_prefix_set.ref.namespace` | [rule_list.rules.source_ip_prefix_set.ref.namespace](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3103003002113303-0112113300231002-0321311111222101-2123100020011321-1302003132032202-1021231202123313-2203133203010221-2132023031132122) |
| `rule_list.rules.source_ip_prefix_set.ref.tenant` | [rule_list.rules.source_ip_prefix_set.ref.tenant](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2233232320111013-1200321010312211-2231033231302110-2100200022003132-2332301103031313-0000112132210200-3130313302310320-0123130021313222) |
| `rule_list.rules.source_ip_prefix_set.ref.uid` | [rule_list.rules.source_ip_prefix_set.ref.uid](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0321122200133011-0300103113203000-1110210230113031-0301021313113303-3133230200232122-2221112212221033-0330232232110010-1331022302011202) |
| `rule_list.rules.source_label_selector` | [rule_list.rules.source_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1011100133213022-1220303013330231-2310123003010003-2030030121321201-3311213032330030-3022223311312120-2013223002323311-3331121313130311) |
| `rule_list.rules.source_label_selector.expressions` | [rule_list.rules.source_label_selector.expressions](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0311003332100323-1223222330210020-0303131000112230-1220312022311222-2012212210101201-1232110233032121-3121003103201110-0113312311112232) |
| `rule_list.rules.source_prefix_list` | [rule_list.rules.source_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3122222301310230-2332123023123000-3031012223203122-2021330303121033-3321320322312313-0121231010011223-3020331233211120-2121011011032111) |
| `rule_list.rules.source_prefix_list.prefixes` | [rule_list.rules.source_prefix_list.prefixes](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2033113011000112-1022010303313111-3131113313303101-3232232200012312-1322223111231201-3321110321302313-2120003110301323-1212211311120312) |
| `timeouts` | [timeouts](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1302112313333020-1301021223132010-1001112211111132-3001222032223310-3230311122020201-3112012000230021-3130011210033221-1303300322030111) |
| `timeouts.create` | [timeouts.create](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0201301120300033-2003112021320030-0030022120331021-3100103103103033-2011003030123322-3122033100010122-1030302002222233-3102002311313011) |
| `timeouts.delete` | [timeouts.delete](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0321223300322132-1113323221331020-3111230001222230-0013330033332210-2212213122232133-3013313022322232-3310010013223313-0020211221122012) |
| `timeouts.read` | [timeouts.read](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3212203211030210-1101020212231302-0210212233332022-3212301212333031-0230220013010212-1330201012331112-0301002023300111-3123312102203001) |
| `timeouts.update` | [timeouts.update](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2223202031102230-3031311130301212-3211320303131131-0120223030220121-2012302010321100-0323110231233310-3322031200231002-3222332211101130) |

<a id="canonical-3223320112112230-2212303021303220-0032003330102130-1132113032131230-2223131221013023-3312310331210030-0000300002231203-0331000231110033"></a>

## Next pages — Property reference / 302331202321 / 12

- [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1102131223103231-1221212033311132-2212320133110100-2123233220030223-0320001123233230-0220122332332022-1031011132232112-2000130111010012)
- [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1020330031030313-2121223312230100-0130303020203031-0212211201300200-3313233321023213-2031002212123211-2000302131222100-2133000113223233)
- [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0233101011203202-2131101223002310-1303102211012323-0203123013122133-0310123213033203-2230322231320031-3022313211330010-3002230320203231)
- [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2330302201311220-0033000223123322-0002331233031022-1003033033120002-2231321310203322-3303012011303311-0033323313131023-1001131213330232)
- [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3030320201232003-3313202120101210-1220013203102130-1220122311333002-2103100102202103-3011023011011210-3322030112202311-0000200001100212)
- [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2311230322202032-3131231332113201-2233132200232032-2122100001221223-0001203030301022-0101323113220031-3012202300312032-1212221201330030)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [timeouts](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1100132312213200-0032121211122030-0200203311120012-2111203021323202-3323203001323231-1123113002211000-3022113223112313-1110213303100000)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-1102131223103231-1221212033311132-2212320133110100-2123233220030223-0320001123233230-0220122332332022-1031011132232112-2000130111010012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121202232322301-3033110102133322-0133230011323001-0021223132333132-1202120023023323-2121310220010223-2332022322310110-2023310302012310"></a>

## allow_all — allow_all / 313103333003 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- allow_all

<a id="canonical-2133223303012233-1322211203320102-3333133003102201-2032313120001132-1001123310321231-1100102303321221-1221113302131011-3232022011001312"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: allow\_all, allowed\_destinations, allowed\_sources, denied\_destinations, denied\_sources,
deny\_all, rule\_list\] Enable this option. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

- [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2133223303012233-1322211203320102-3333133003102201-2032313120001132-1001123310321231-1100102303321221-1221113302131011-3232022011001312)
- [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1203212233000123-1212222110033030-0030130000201331-0120300313013213-2023212000302301-2123002113300001-1132331130033312-3031202121032130)
- [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3110333330110202-3323130322121202-0233302221003111-0131312330101032-2031132122120213-1000231110300131-2132103023031221-2111312321202311)
- [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0211333302313332-2113222200302323-2023112222130331-3220320200300001-2123203203022203-1120033101200103-0303211000110212-1213112012211022)
- [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2101011013200022-2221031222122132-3013110210110103-2311023232112003-2331021012221301-3022330002032310-3012133331301133-1122203311331021)
- [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2000113230323112-1011333330331303-3012121131303002-2003020231023220-2230103012233030-0303111131332102-2212223100003222-2101003033331213)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2202011303310301-1022220123013310-2111000100022230-3210133001102222-1230300111121233-3110212333200313-2202132223011021-3321221100031111)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
allow_all = {}
```

<a id="canonical-1301003122133131-3202323233033100-0313131013102022-1201223113231300-1213223012102302-2001113230312023-0203030221031313-2323222332231122"></a>

## Direct properties — allow_all / 313103333003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301302010002102-3303210203232213-1213100300332202-0031313222132221-0012330222011123-2021010033221312-3200110130112100-2130111302303300"></a>

## Next pages — allow_all / 313103333003 / 4

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-1020330031030313-2121223312230100-0130303020203031-0212211201300200-3313233321023213-2031002212123211-2000302131222100-2133000113223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300310022120122-0312223003330210-2233230202131212-3103100331203133-0223333231013300-1021012332022101-3003303032132213-1112203013012311"></a>

## allowed_destinations — allowed_destinations / 201132033221 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- allowed_destinations

<a id="canonical-1203212233000123-1212222110033030-0030130000201331-0120300313013213-2023212000302301-2123002113300001-1132331130033312-3031202121032130"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
allowed_destinations {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320022203222202-3333032123033010-0130212310231213-1330303203123222-0012222332133302-2331232301110031-3123232201123103-2233032200133223"></a>

## Direct properties — allowed_destinations / 201132033221 / 3

<a id="canonical-2001323111123133-1031322002213021-1031132010013300-3230020233023000-1203012022010201-0032310133310112-0302111333023213-1201101320320120"></a>

<a id="canonical-1022312022321323-0022230221303113-0333302331330003-0323302301032231-3133102030300213-2333200212010011-2220330111020112-1113102232012230"></a>

## prefix property — allowed_destinations / 201132033221 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-0323111223220100-0233000213320032-2330220320101001-0322120203312231-3030021313132100-2332030222000230-1332201032122011-3230132233032212"></a>

## Next pages — allowed_destinations / 201132033221 / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0233101011203202-2131101223002310-1303102211012323-0203123013122133-0310123213033203-2230322231320031-3022313211330010-3002230320203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021231002021311-0232221322110302-3132312221011330-1233331131023310-2002333030102230-1123103003101322-0301112212030203-1032312132113120"></a>

## allowed_sources — allowed_sources / 011103022033 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- allowed_sources

<a id="canonical-3110333330110202-3323130322121202-0233302221003111-0131312330101032-2031132122120213-1000231110300131-2132103023031221-2111312321202311"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
allowed_sources {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130102330320100-1020030320011320-1121212202210103-0220023313132130-2312100003120033-0310310321222321-1321232313203123-0221033120331332"></a>

## Direct properties — allowed_sources / 011103022033 / 3

<a id="canonical-3030101122232321-3221132312033003-0231330331221213-2110200122101302-1000030022020010-2023332220200101-0331110301322113-1013201310211112"></a>

<a id="canonical-2033233022013220-2030103212212322-3213023121202110-2221322120213003-3012100332333003-1122023031101311-3311131211033100-0103323011223201"></a>

## prefix property — allowed_sources / 011103022033 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-0330210220013020-1212032013001313-1222000103333232-2322213100130331-1210001113023213-0033320003113023-3010033133123110-2330120230113311"></a>

## Next pages — allowed_sources / 011103022033 / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2330302201311220-0033000223123322-0002331233031022-1003033033120002-2231321310203322-3303012011303311-0033323313131023-1001131213330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232220113211300-2030021301320330-0233231122223112-1013010003331320-2300333000021123-0112230111023200-2323202122221232-0210332301013231"></a>

## denied_destinations — denied_destinations / 303300320222 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- denied_destinations

<a id="canonical-0211333302313332-2113222200302323-2023112222130331-3220320200300001-2123203203022203-1120033101200103-0303211000110212-1213112012211022"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
denied_destinations {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102101010112313-0322121000123130-1322120313231021-3022122211331301-2201130013313130-2112230222230023-3233310332013230-1030032320300233"></a>

## Direct properties — denied_destinations / 303300320222 / 3

<a id="canonical-0010001310002013-1133111323230000-2000323300131013-0012312212132223-0222200321333013-3323331321031121-0330332332333133-0130003031101123"></a>

<a id="canonical-0111320113320123-3331131201332211-2220210313300033-1311311222012312-0000031313101122-0012122131300011-3130120331100103-0233201213231213"></a>

## prefix property — denied_destinations / 303300320222 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2132020333203313-2012330322210333-0331230112103221-0320103000311113-2203102130011220-0332202121011012-2031232031203230-2231323322100033"></a>

## Next pages — denied_destinations / 303300320222 / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3030320201232003-3313202120101210-1220013203102130-1220122311333002-2103100102202103-3011023011011210-3322030112202311-0000200001100212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233100111103133-1202102010101233-0323100022331100-3032000011202231-0030123301032203-1213223130203220-2013011103222112-1211313113012232"></a>

## denied_sources — denied_sources / 231101112202 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- denied_sources

<a id="canonical-2101011013200022-2221031222122132-3013110210110103-2311023232112003-2331021012221301-3022330002032310-3012133331301133-1122203311331021"></a>

Type: `"object"`. single nested block, Optional.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

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
denied_sources {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102222113220320-1101001300213021-1302030230012221-3223333100111023-0002330102111003-3223020000033221-1011113102222220-0323003312032300"></a>

## Direct properties — denied_sources / 231101112202 / 3

<a id="canonical-1010013023322010-0320123102100202-1102222330301021-2232022322220133-1201002031301331-0201002002211331-1002223231131331-2131232211201330"></a>

<a id="canonical-1311020101220033-3232213133302022-3011010221131223-1001332311303321-1032001302000222-0211020220233221-3323101131102020-0200232113321122"></a>

## prefix property — denied_sources / 231101112202 / 4

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-1001102023022012-3031331310111303-1323323213021003-0021131331000212-0033113120000311-2212332230011213-0131221233022012-0112031231212100"></a>

## Next pages — denied_sources / 231101112202 / 5

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2311230322202032-3131231332113201-2233132200232032-2122100001221223-0001203030301022-0101323113220031-3012202300312032-1212221201330030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032132110012103-1032103320332330-3101033330122221-2331220100202213-2022212232100102-2203330323232033-1102223233210020-0021300330231202"></a>

## deny_all — deny_all / 301021000201 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- deny_all

<a id="canonical-2000113230323112-1011333330331303-3012121131303002-2003020231023220-2230103012233030-0303111131332102-2212223100003222-2101003033331213"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
deny_all = {}
```

<a id="canonical-0331112023220111-2133133333322331-0120303232303122-3301302332231113-3113310202231212-3331000200101311-3023213101012030-2200332002111311"></a>

## Direct properties — deny_all / 301021000201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101200211131310-2312220112120302-2203213232001030-1122112023232110-0221002010131320-3231001213200200-3120333213300012-2201120222133311"></a>

## Next pages — deny_all / 301021000201 / 4

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120313303001300-3231321300033333-3313001131220323-2112303331300220-1302300212031031-2301221002313313-2331130321102102-0030023320001132"></a>

## rule_list — rule_list / 023222230021 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- rule_list

<a id="canonical-2202011303310301-1022220123013310-2111000100022230-3210133001102222-1230300111121233-3110212333200313-2202132223011021-3321221100031111"></a>

Type: `"object"`. single nested block, Optional.

Custom Enhanced Firewall Policy Rules. Custom Enhanced Firewall Policy Rules.

Upstream description:

Custom Enhanced Firewall Policy Rules.

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

<a id="canonical-1313130031201033-3332322211112111-2102213113332210-3111032010232331-3230113022010212-2300201001013312-3123203100103020-0021112200330322"></a>

## Direct properties — rule_list / 023222230021 / 3

- [rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012): complete subsection reference.

<a id="canonical-1002100201030212-2121001111321210-0321202221120301-1300202201101322-2021230001002122-2021121121331030-1302200330231300-2121032011132301"></a>

## Next pages — rule_list / 023222230021 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000030130100223-0023203232133333-2320312102030133-1320003321033003-2330201103202120-2003221100301023-1320111131023202-1221222213231112"></a>

## rule_list.rules — rules / 101212201213 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- rule_list.rules

<a id="canonical-3301100221132102-3103020103003132-0322322333123121-0223003111133303-2130303320203322-0222022000000210-0220312133212132-0231130333021023"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policy Rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_destinations",
    "all_sli_vips"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "all_slo_vips"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_destinations",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "all_slo_vips"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_sli_vips",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("all_slo_vips",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("all_sources",
    "inside_sources"),
  validators.ConflictingListObjectAttributes("all_sources",
    "outside_sources"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("all_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("allow",
    "deny"),
  validators.ConflictingListObjectAttributes("allow",
    "insert_service"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("deny",
    "insert_service"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_aws_vpc_ids",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "destination_label_selector"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_ip_prefix_set",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "destination_prefix_list"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_label_selector",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("destination_prefix_list",
    "inside_destinations"),
  validators.ConflictingListObjectAttributes("destination_prefix_list",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("inside_destinations",
    "outside_destinations"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "outside_sources"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("inside_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_aws_vpc_ids"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("outside_sources",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_ip_prefix_set"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("source_aws_vpc_ids",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_ip_prefix_set",
    "source_label_selector"),
  validators.ConflictingListObjectAttributes("source_ip_prefix_set",
    "source_prefix_list"),
  validators.ConflictingListObjectAttributes("source_label_selector",
    "source_prefix_list")}
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300230202113102-2201223202010300-3202123010003022-1302020300000003-1100323123032233-1220023032213013-1110123122020313-3210000100110110"></a>

## Direct properties — rules / 101212201213 / 3

- [advanced_action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1321010123302002-2320200102133330-3210131231113103-2212020310013210-3311131332320311-2000010132133233-1232322130322220-1000212000333110): complete subsection reference.

- [all_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3332223003121103-0100232322211100-3233232313010000-2111222012221103-3302201203232012-3201033330313102-0323311033130212-0113320030223211): complete subsection reference.

- [all_sli_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0320210302122011-1122312300223112-1112322010203001-0212231203012100-0330023300321111-3231033010233002-1303013130311210-0210100233013311): complete subsection reference.

- [all_slo_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2210010332032211-1132103201011133-0213113111322302-1210232313222102-3331111112121232-2231133003310122-2031222322122332-1001112012233101): complete subsection reference.

- [all_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0110331230112330-3102332302302200-2012112312202303-1300223133010212-3331230120330221-0312001000023112-1220302111130221-0102210110320222): complete subsection reference.

- [all_tcp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0233310331200213-3330302312302201-3012302020132221-0010133011202110-0031321002021022-0010330012102032-3333231320022312-3103321102222110): complete subsection reference.

- [all_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2221321120210323-0202302011101332-2330003112113310-3032202011230313-0233212031132210-0120132123110013-0023030311203121-3123120233103002): complete subsection reference.

- [all_udp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2313102110112012-2300312113203123-1222330300130300-0200110113020333-2102300200121323-0311302331220130-1020020232133232-0102013013201231): complete subsection reference.

- [allow](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0122001303211011-0322112010232320-3031110311233031-1033011011331302-2333332030321022-2113021010112032-1031113103312312-1222000233230131): complete subsection reference.

- [applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0330031202013312-0323002013321210-1313331003023100-1330110103110111-0220123123000332-3133202231132231-3301330023321212-3331132132003003): complete subsection reference.

- [deny](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1211021022123210-0232100032323000-2203312122200231-3312031112011222-1032121332132330-0131120033320111-3020200233210221-2323000113010210): complete subsection reference.

- [destination_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0203111133203023-1330113002232331-0110303203000213-1122030333032222-1003021330103020-1330113012310301-0203313022022110-1203121113122201): complete subsection reference.

- [destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1331310301100100-2111111311110322-1021212123330203-0302011101203200-3120013012301321-1003123232301313-3202303101011213-2323223111111330): complete subsection reference.

- [destination_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2101213123123320-2132022333123021-3032321233212023-1212213222010002-0012321132202313-3003230002003002-2202233320003221-1003210330000033): complete subsection reference.

- [destination_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2313000201120200-3020021003012101-1022033123301103-1202100121322203-1221333021001201-2021301220322310-2111121332221323-1310303103332033): complete subsection reference.

- [insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2020300331210132-2230130332021010-0000221103022320-2233032323132332-0032102332020013-1103313022132121-2133133120310222-2132221222210330): complete subsection reference.

- [inside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0210123303133133-0121121100000030-3023232322013332-2120310101100302-2013012010102303-2331300111330020-1020310333000322-1322222123200332): complete subsection reference.

- [inside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3202202312023033-1031023030121320-1221012313301320-3200233222223002-1221300322200302-2123123032312001-3123222131330133-1000211213213003): complete subsection reference.

- [label_matcher](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0303001303312103-0312032001221303-0132103311122021-2300032320020113-1013123313121012-2113002221232130-1223331023323231-0211322130133020): complete subsection reference.

- [metadata](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0001102333021033-3310131112313200-3320011333221213-3302322222313113-0220110101212213-0120330022110003-1000030320323323-2100020012323003): complete subsection reference.

- [outside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3122030321331013-3001330001303333-2301200212231230-2113231333332210-1231321221030122-0102111023122121-3310110331212113-1301323121110312): complete subsection reference.

- [outside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3001020022231030-1302203121112110-0330100103120321-1211022213100232-1030000013220010-3213213123323002-0101120220032133-2031003122331320): complete subsection reference.

- [protocol_port_range](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2002103111113300-0332102033002033-2132302000221023-3221111201100303-0110202033011102-0333312301123312-1122301212333330-3331110232302232): complete subsection reference.

- [source_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0320000103011022-2102323210130022-1320113322021020-1100200102121211-3113320000301032-3222121021012200-2212111133210211-3011201122110221): complete subsection reference.

- [source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221): complete subsection reference.

- [source_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2311323100110022-0010031233102023-0223111213203020-3113301302322232-3320133310032212-1021310132231130-3321200101220231-3212313311011231): complete subsection reference.

- [source_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3023232001313222-3033313230302113-3233020302311303-2132130120132131-3232210330000101-2013313212220311-3033112311101010-1210330101321311): complete subsection reference.

<a id="canonical-2231021130333223-0130013201100122-3313032230113032-1310210000100122-3320013111033300-2110233120210011-2010013203321030-1302202102323221"></a>

## Next pages — rules / 101212201213 / 4

- [rule_list.rules.advanced_action](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1321010123302002-2320200102133330-3210131231113103-2212020310013210-3311131332320311-2000010132133233-1232322130322220-1000212000333110)
- [rule_list.rules.all_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3332223003121103-0100232322211100-3233232313010000-2111222012221103-3302201203232012-3201033330313102-0323311033130212-0113320030223211)
- [rule_list.rules.all_sli_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0320210302122011-1122312300223112-1112322010203001-0212231203012100-0330023300321111-3231033010233002-1303013130311210-0210100233013311)
- [rule_list.rules.all_slo_vips](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2210010332032211-1132103201011133-0213113111322302-1210232313222102-3331111112121232-2231133003310122-2031222322122332-1001112012233101)
- [rule_list.rules.all_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0110331230112330-3102332302302200-2012112312202303-1300223133010212-3331230120330221-0312001000023112-1220302111130221-0102210110320222)
- [rule_list.rules.all_tcp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0233310331200213-3330302312302201-3012302020132221-0010133011202110-0031321002021022-0010330012102032-3333231320022312-3103321102222110)
- [rule_list.rules.all_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2221321120210323-0202302011101332-2330003112113310-3032202011230313-0233212031132210-0120132123110013-0023030311203121-3123120233103002)
- [rule_list.rules.all_udp_traffic](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2313102110112012-2300312113203123-1222330300130300-0200110113020333-2102300200121323-0311302331220130-1020020232133232-0102013013201231)
- [rule_list.rules.allow](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0122001303211011-0322112010232320-3031110311233031-1033011011331302-2333332030321022-2113021010112032-1031113103312312-1222000233230131)
- [rule_list.rules.applications](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0330031202013312-0323002013321210-1313331003023100-1330110103110111-0220123123000332-3133202231132231-3301330023321212-3331132132003003)
- [rule_list.rules.deny](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1211021022123210-0232100032323000-2203312122200231-3312031112011222-1032121332132330-0131120033320111-3020200233210221-2323000113010210)
- [rule_list.rules.destination_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0203111133203023-1330113002232331-0110303203000213-1122030333032222-1003021330103020-1330113012310301-0203313022022110-1203121113122201)
- [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1331310301100100-2111111311110322-1021212123330203-0302011101203200-3120013012301321-1003123232301313-3202303101011213-2323223111111330)
- [rule_list.rules.destination_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2101213123123320-2132022333123021-3032321233212023-1212213222010002-0012321132202313-3003230002003002-2202233320003221-1003210330000033)
- [rule_list.rules.destination_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2313000201120200-3020021003012101-1022033123301103-1202100121322203-1221333021001201-2021301220322310-2111121332221323-1310303103332033)
- [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2020300331210132-2230130332021010-0000221103022320-2233032323132332-0032102332020013-1103313022132121-2133133120310222-2132221222210330)
- [rule_list.rules.inside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0210123303133133-0121121100000030-3023232322013332-2120310101100302-2013012010102303-2331300111330020-1020310333000322-1322222123200332)
- [rule_list.rules.inside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3202202312023033-1031023030121320-1221012313301320-3200233222223002-1221300322200302-2123123032312001-3123222131330133-1000211213213003)
- [rule_list.rules.label_matcher](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0303001303312103-0312032001221303-0132103311122021-2300032320020113-1013123313121012-2113002221232130-1223331023323231-0211322130133020)
- [rule_list.rules.metadata](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0001102333021033-3310131112313200-3320011333221213-3302322222313113-0220110101212213-0120330022110003-1000030320323323-2100020012323003)
- [rule_list.rules.outside_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3122030321331013-3001330001303333-2301200212231230-2113231333332210-1231321221030122-0102111023122121-3310110331212113-1301323121110312)
- [rule_list.rules.outside_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3001020022231030-1302203121112110-0330100103120321-1211022213100232-1030000013220010-3213213123323002-0101120220032133-2031003122331320)
- [rule_list.rules.protocol_port_range](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2002103111113300-0332102033002033-2132302000221023-3221111201100303-0110202033011102-0333312301123312-1122301212333330-3331110232302232)
- [rule_list.rules.source_aws_vpc_ids](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0320000103011022-2102323210130022-1320113322021020-1100200102121211-3113320000301032-3222121021012200-2212111133210211-3011201122110221)
- [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221)
- [rule_list.rules.source_label_selector](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2311323100110022-0010031233102023-0223111213203020-3113301302322232-3320133310032212-1021310132231130-3321200101220231-3212313311011231)
- [rule_list.rules.source_prefix_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3023232001313222-3033313230302113-3233020302311303-2132130120132131-3232210330000101-2013313212220311-3033112311101010-1210330101321311)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-1321010123302002-2320200102133330-3210131231113103-2212020310013210-3311131332320311-2000010132133233-1232322130322220-1000212000333110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221122223313300-0333321222031000-0201301211112110-1000000232201331-0203200333122222-1200111233233212-0323000322322232-3220123322031130"></a>

## rule_list.rules.advanced_action — advanced_action / 011001310201 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.advanced_action

<a id="canonical-3201112220310003-2012033200331331-0300010133200220-2213032222033203-1010123001110223-3301131212123301-3221202123031322-3023100100310332"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
advanced_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323120313130020-2210120321112220-0322120023121003-3003121330033012-3122111301301313-0022012023232212-3222010203331303-3212311111100103"></a>

## Direct properties — advanced_action / 011001310201 / 3

<a id="canonical-1221323031312322-3120111223123312-2020020331320120-1321210222332330-0000230232123130-0133301230323031-0021130320203020-2301020332312131"></a>

<a id="canonical-2321233333131123-2210310110000133-2100333133011022-2303210002213132-0033002033303022-3001300330120033-2203300312330123-1130113313212132"></a>

## action property — advanced_action / 011001310201 / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2221113023113123-0313311223103213-0312133133013203-1002011011312013-0002202330021203-1311213023213303-1331302230120123-1210231231023101"></a>

## Next pages — advanced_action / 011001310201 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3332223003121103-0100232322211100-3233232313010000-2111222012221103-3302201203232012-3201033330313102-0323311033130212-0113320030223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323121313033230-2313302101210003-1220231010113313-0103113031302101-2311330101320133-0011203021221330-1322211100131113-1032010330023002"></a>

## rule_list.rules.all_destinations — all_destinations / 002330200331 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_destinations

<a id="canonical-2123003221131311-1032122200103233-1231003131200012-0311212303111011-2021333200111110-1201112233221103-2011231203231132-0223303320202120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

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
all_destinations = {}
```

<a id="canonical-3230213120221200-3310020132011323-3302200310000313-0002332323130103-0020131322121022-0020002333133333-2332203010313113-0200302222111123"></a>

## Direct properties — all_destinations / 002330200331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232121323020231-3021032210332203-3311103231120203-0301122031032111-2030313331221003-3212200210000122-1221030020310101-3212100010003300"></a>

## Next pages — all_destinations / 002330200331 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0320210302122011-1122312300223112-1112322010203001-0212231203012100-0330023300321111-3231033010233002-1303013130311210-0210100233013311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133023322110302-3302301212032331-0303121232010022-2203323122112203-0332120201333033-3231220130302233-0310111311032320-3032300333020021"></a>

## rule_list.rules.all_sli_vips — all_sli_vips / 203313332230 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_sli_vips

<a id="canonical-2133123132102300-2131000301031023-1132011201323102-3200312232313331-3330010300133200-0013313303221323-3020201013112330-3120131221020212"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_sli_vips = {}
```

<a id="canonical-0301023201012032-1023123112010300-3001222000232330-2110310120222230-2101000231233213-0123221321230333-1110230132320310-3123321033211033"></a>

## Direct properties — all_sli_vips / 203313332230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232102031213233-0302221112020233-3321320200031310-0032221011123001-3023121231231202-1030313233002132-0003101012310131-1103332233223112"></a>

## Next pages — all_sli_vips / 203313332230 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2210010332032211-1132103201011133-0213113111322302-1210232313222102-3331111112121232-2231133003310122-2031222322122332-1001112012233101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331203202122102-2030321020303333-2300030310213001-3131233213210021-3112302111301032-3300303321213221-3313000031301312-0210322221320110"></a>

## rule_list.rules.all_slo_vips — all_slo_vips / 100012212111 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_slo_vips

<a id="canonical-1123213030032131-0222020133003300-2202220132133323-1210001032312013-1231010211122030-1111301323123321-3120013111323032-2220321011320331"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_slo_vips = {}
```

<a id="canonical-1003012133022023-2321200313311301-0321301100200133-3132203032210022-1100021230331332-0320230103122113-0220310332120123-2302212101120130"></a>

## Direct properties — all_slo_vips / 100012212111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232221302011313-2031200231132031-1021310110132010-3023223300100212-0302230131211302-3322023313323100-1300321010321000-0130021120302330"></a>

## Next pages — all_slo_vips / 100012212111 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0110331230112330-3102332302302200-2012112312202303-1300223133010212-3331230120330221-0312001000023112-1220302111130221-0102210110320222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210120103233121-3003001223130102-3313211202032130-1203130130323030-0203100332102233-3110223310132230-0301131122310021-0111012222120002"></a>

## rule_list.rules.all_sources — all_sources / 102110231303 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_sources

<a id="canonical-2310011032123110-3003202032303212-0210233201112220-3122201113332202-0113223131103220-0323212130323030-2221130331302121-2232003130113100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all sources.

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
all_sources = {}
```

<a id="canonical-3033313110200303-2121202012000212-0232032030302331-1200122331001322-3120233110233332-2013210312123131-0311012023033102-0000020131131133"></a>

## Direct properties — all_sources / 102110231303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113220233021131-3001000123312310-2113113230312100-1322110323232102-2232102221300323-2100223203000301-0123303021203322-2031032132230000"></a>

## Next pages — all_sources / 102110231303 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0233310331200213-3330302312302201-3012302020132221-0010133011202110-0031321002021022-0010330012102032-3333231320022312-3103321102222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202332202200111-0123332212132003-3220203232222322-2032112210202232-0023211033121223-3103102212101333-3101112210021012-0103312220112233"></a>

## rule_list.rules.all_tcp_traffic — all_tcp_traffic / 333221111333 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_tcp_traffic

<a id="canonical-1201211210311032-2013132033001000-3212121113221103-3133102010202100-0132200203310000-1220220301001003-2111320133102231-0332231333211211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

<a id="canonical-3221002002120330-3123101001312212-1331303103312220-1122112100300112-1202121332000113-3122032132022111-2212103110123202-1033203021213021"></a>

## Direct properties — all_tcp_traffic / 333221111333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300121120321101-3120020310020113-2333232300203231-2312312301023201-2223110211023321-2311111210002100-3100132102320312-0331321102331202"></a>

## Next pages — all_tcp_traffic / 333221111333 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2221321120210323-0202302011101332-2330003112113310-3032202011230313-0233212031132210-0120132123110013-0023030311203121-3123120233103002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212200233133103-1220320211032223-3032021011232000-3021110001300223-3300000131123332-1121011003220123-2030011310002131-1122232031130321"></a>

## rule_list.rules.all_traffic — all_traffic / 022311213103 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_traffic

<a id="canonical-0332030211123200-1022030033101211-0031323300320232-0312203332233233-1102210322131111-2222123322330011-1201212212221130-2232013013211023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

<a id="canonical-1130120332033303-0223221012133132-1121021010323102-3213031000012222-2322203333221233-1223023012033011-2332001003300320-1120110001300302"></a>

## Direct properties — all_traffic / 022311213103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013000231122211-0220203203222010-1302033033030322-3212020101012021-3312023221230021-1031101031101013-0300022310030130-0201310321330222"></a>

## Next pages — all_traffic / 022311213103 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2313102110112012-2300312113203123-1222330300130300-0200110113020333-2102300200121323-0311302331220130-1020020232133232-0102013013201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100011101013131-1322022031012233-1301132303011001-0330210022011223-2110233123311332-0001223131102300-1002113330332111-0122022231330302"></a>

## rule_list.rules.all_udp_traffic — all_udp_traffic / 323120010330 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_udp_traffic

<a id="canonical-2102230322013231-0321003201230301-0132113030121222-1310001221002012-0301011100103211-0113233311032022-2003311312030313-3232231032101000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

<a id="canonical-1010220102102301-0121000230333321-0131033233113303-2210022120121002-0011132030303303-0230221303203112-1230101113131221-0103313010333120"></a>

## Direct properties — all_udp_traffic / 323120010330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132310330130133-1221320101120331-0011103130232301-0320101001122201-0221103232013131-3223123023013221-0113323100233232-2233122023200301"></a>

## Next pages — all_udp_traffic / 323120010330 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0122001303211011-0322112010232320-3031110311233031-1033011011331302-2333332030321022-2113021010112032-1031113103312312-1222000233230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032313311022102-1233230032333002-3330020210321111-1112032303003310-2330210222120000-1123322101313203-3212320233322223-3320012203330212"></a>

## rule_list.rules.allow — allow / 131123303313 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.allow

<a id="canonical-1211003102300000-2122203110031310-3202323033232200-1320300202103102-0120230111101220-0211111120330322-2023310222113100-2000330210212320"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
allow = {}
```

<a id="canonical-2230110220330301-1221012110310113-3201133323120230-2332311200201020-3021032200232223-0233123030112003-0130101031233133-0322303233120013"></a>

## Direct properties — allow / 131123303313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330012212121032-2021103300102100-2301322201222120-1331000003123110-2113100103101200-1302100203221001-0132130121231200-0201313103310212"></a>

## Next pages — allow / 131123303313 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0330031202013312-0323002013321210-1313331003023100-1330110103110111-0220123123000332-3133202231132231-3301330023321212-3331132132003003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223303000021022-3000101300103231-2212113131120331-1011330031323111-0332302313031101-0022030320131203-1001322030301331-3231103312230012"></a>

## rule_list.rules.applications — applications / 333103331211 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.applications

<a id="canonical-0111330200023022-0020112202113210-0332022220003121-2332022200032111-2002313133030002-2200213320200220-1103332312010102-3332223322021132"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-3000222301100001-2310130102202031-2210231010232031-0211132020002231-2000232302223112-3212221113331023-3132000312333113-1013103311003301"></a>

## Direct properties — applications / 333103331211 / 3

<a id="canonical-0311021022221301-2200202010232201-1201011301223213-3312233011330011-0101320220133230-2312312323130210-3101101103311330-3323231031120211"></a>

<a id="canonical-1211320021212002-0120211321113312-2212031223110030-0223020333310013-0200203110332112-0230210213132310-2223332212231211-0121123202023320"></a>

## applications property — applications / 333103331211 / 4

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-3013000123222110-2331310331023032-2021132111021323-1333233102301022-2300301033102021-3332312232030111-3331001101332113-3220200110123312"></a>

## Next pages — applications / 333103331211 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-1211021022123210-0232100032323000-2203312122200231-3312031112011222-1032121332132330-0131120033320111-3020200233210221-2323000113010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330222132112230-1131010001230231-0001203302003111-3031223202013120-1212102000322101-1213113110311023-0112313122333113-2100003221123203"></a>

## rule_list.rules.deny — deny / 011031231000 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.deny

<a id="canonical-0322003203330100-1231313032302110-1101303223001322-0120202230100232-0121001131220023-1200133022302230-1030103001012231-1132101132231201"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
deny = {}
```

<a id="canonical-2231023321213002-3313131300032103-3213001300301231-0012122220112012-2222201321300321-0223301232100213-1332020110210310-2001333122003030"></a>

## Direct properties — deny / 011031231000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123021020012002-0302231211333121-2030210232222221-0223303110313310-0130333130121002-3133000333320032-2133332302012300-3121302333131012"></a>

## Next pages — deny / 011031231000 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0203111133203023-1330113002232331-0110303203000213-1122030333032222-1003021330103020-1330113012310301-0203313022022110-1203121113122201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113011212212131-3133322203002010-1101013102012130-0102233322032023-3021320312303302-1131100322003332-2300331301112110-3033221032123011"></a>

## rule_list.rules.destination_aws_vpc_ids — destination_aws_vpc_ids / 002303323000 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.destination_aws_vpc_ids

<a id="canonical-0213223032210131-1212320013013200-0313012232131001-3313221330332113-1203233010301310-2131310112033032-3212220010211130-2000213323132200"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for destination aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("vpc_id")}
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
destination_aws_vpc_ids {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303031001221003-1120323303011031-3211133221222300-3021102220113322-2220221222211123-3222131100030000-1323030113130013-3113131311220232"></a>

## Direct properties — destination_aws_vpc_ids / 002303323000 / 3

<a id="canonical-0211031103223322-0233111102011112-0302302122030203-0300333332320101-1110210202210232-3013102211231231-3311223113002321-2110212132012232"></a>

<a id="canonical-3133300322211301-1032302002222113-1010133123002022-3320231111330133-1201010113303212-2213021110200010-2221211102112202-3111132020231232"></a>

## vpc_id property — destination_aws_vpc_ids / 002303323000 / 4

Type: `["list", "string"]`. Optional.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0032311323331332-2003301033220002-3210032103033103-3313100310112330-1122222013223202-3100021220231330-2110011010212210-0232132000213231"></a>

## Next pages — destination_aws_vpc_ids / 002303323000 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-1331310301100100-2111111311110322-1021212123330203-0302011101203200-3120013012301321-1003123232301313-3202303101011213-2323223111111330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001132020132102-0122223113003313-2220210112123032-0311331222223011-1232302113211101-3230330010232212-3222211221303312-3023213000120331"></a>

## rule_list.rules.destination_ip_prefix_set — destination_ip_prefix_set / 320210232131 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.destination_ip_prefix_set

<a id="canonical-0231001002013222-0313002303030213-2313122302120322-2333000023003202-0300213031223012-2133100023200012-3021030310121213-3320301032132321"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
destination_ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022022001132223-2103133031233313-1110223311002022-1310012113022021-3223103203010201-3120200233303032-3121320033220221-2220230110011200"></a>

## Direct properties — destination_ip_prefix_set / 320210232131 / 3

- [ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3330211030203212-2103032230030131-2231011310221003-1231212023303032-3210101023130100-0020111133102022-2332122322101100-1001313013233030): complete subsection reference.

<a id="canonical-2232210331132331-1000323013232032-1010211023210101-2322332011120222-3221003303032001-2020120031032001-2201112100111303-2321020333220121"></a>

## Next pages — destination_ip_prefix_set / 320210232131 / 4

- [rule_list.rules.destination_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3330211030203212-2103032230030131-2231011310221003-1231212023303032-3210101023130100-0020111133102022-2332122322101100-1001313013233030)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3330211030203212-2103032230030131-2231011310221003-1231212023303032-3210101023130100-0020111133102022-2332122322101100-1001313013233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021220012133200-2211300303301120-3121011302033333-2023333130232303-2200010110323110-3111330231023332-0312202130101213-1300121103311130"></a>

## rule_list.rules.destination_ip_prefix_set.ref — ref / 011033121020 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1331310301100100-2111111311110322-1021212123330203-0302011101203200-3120013012301321-1003123232301313-3202303101011213-2323223111111330)
- rule_list.rules.destination_ip_prefix_set.ref

<a id="canonical-3233211122121213-3032130020011312-1001320222202133-2323223020301223-1011031331330122-2330331200101320-0233301030011012-2222333122122000"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1310101110303212-2003233101132103-2022123200110113-0212331301113222-1220003133033002-3111012003100321-3302130212133010-1300000333000030"></a>

## Direct properties — ref / 011033121020 / 3

<a id="canonical-2030210122233102-3001033001212321-0320211002123333-0312012222231331-0212011203021201-2032102332021002-3233222002323230-0332210300221210"></a>

<a id="canonical-3012100312301000-1231231100100302-2110231232101201-1200002221222220-1112333000311211-0331233033221010-2312301320011122-0133033302122010"></a>

## kind property — ref / 011033121020 / 4

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

<a id="canonical-3233022321133102-3222300032312223-2200221011101020-1331202111333322-2132133333201203-0222300101011300-2210212232030010-1002123001021322"></a>

<a id="canonical-0313033221122031-0202021000312322-3101111000132222-3011111212113031-1031330320220201-1133030111133231-1321130112002222-2332133103002021"></a>

## name property — ref / 011033121020 / 5

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

<a id="canonical-0022210031222111-2021000223012102-3100111002101330-0303131223033302-2103122220232032-1031220211300202-1001123302133200-1032001212231223"></a>

<a id="canonical-2132110000021003-2010231020212102-1022030022322221-2310221032113011-2113110030211110-2332200320032320-3123310010030301-2112321022322021"></a>

## namespace property — ref / 011033121020 / 6

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

<a id="canonical-0132201200111233-1301103220010320-3312110201320012-3303003323022300-3000031231033221-3113301201121113-0223132310223030-3011000300123300"></a>

<a id="canonical-0330032033221300-1202103123323033-3100301133230300-3203232031311010-1111233202212001-1133120221301211-1022121010221213-2112130123112120"></a>

## tenant property — ref / 011033121020 / 7

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

<a id="canonical-3001202010230223-1031231013102223-2021331322331010-0112233230003113-1032320010322112-2003300121123002-1103031233102023-0101010100101133"></a>

<a id="canonical-0021232310023133-1222122202031233-3013010211111123-3233020022122021-0123010301013010-0013230133132211-2133310113131330-3331110310223101"></a>

## uid property — ref / 011033121020 / 8

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

<a id="canonical-2033212222003313-2213203101120033-1131312310120200-2111103121000033-3001112223120113-3333333033331123-0332222300200100-0123230200322032"></a>

## Next pages — ref / 011033121020 / 9

- [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1331310301100100-2111111311110322-1021212123330203-0302011101203200-3120013012301321-1003123232301313-3202303101011213-2323223111111330)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2101213123123320-2132022333123021-3032321233212023-1212213222010002-0012321132202313-3003230002003002-2202233320003221-1003210330000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122302330010120-2223211330022211-1010331203310330-1321312012312200-2202013201132213-2002132002231222-0023230230220113-1123202223132122"></a>

## rule_list.rules.destination_label_selector — destination_label_selector / 110032131000 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.destination_label_selector

<a id="canonical-1000003213001003-3212130311130220-1230000130000111-0122033020233101-0013131311311323-0233212033320121-3232203113223113-1011102122021310"></a>

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
destination_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102220332002300-0031223221222130-0123020012211320-0321221002130302-3030312133010032-3121212210323033-0211311113302020-1323113130120032"></a>

## Direct properties — destination_label_selector / 110032131000 / 3

<a id="canonical-2232201230201211-3230021101020001-0110010132230220-0110232231021022-3220111203330133-2122011330303200-0103322233202202-0230203213121030"></a>

<a id="canonical-3030213321001011-1101222221233101-2302231230223211-3311121331121112-1003303002311333-0021001301030230-0210013333210032-2330200232230112"></a>

## expressions property — destination_label_selector / 110032131000 / 4

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

<a id="canonical-0021330332212332-0123003123111222-1220222303302232-3203031213231312-0003000332120230-3310312333013213-2121302011023132-3032022320320033"></a>

## Next pages — destination_label_selector / 110032131000 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2313000201120200-3020021003012101-1022033123301103-1202100121322203-1221333021001201-2021301220322310-2111121332221323-1310303103332033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011030121021111-1013201021301110-1231330200222022-0111031003131023-1022020133021331-3000231100233312-1133323212001012-2011110033112013"></a>

## rule_list.rules.destination_prefix_list — destination_prefix_list / 232010002322 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.destination_prefix_list

<a id="canonical-0102101103220223-3123222220220101-2300023002202103-2003020312001320-1300123120323032-0011201202310020-2132323322132332-0200300223202033"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
destination_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202121210123012-1013202202201132-3113130203300312-0200332320333323-1332302222330223-2301100010212000-3322033033102201-1010220323031013"></a>

## Direct properties — destination_prefix_list / 232010002322 / 3

<a id="canonical-0233213202100012-1321013220210321-0013102120102032-2120312201202120-3111202320011122-3102303011220223-0110220323303100-0332103112312123"></a>

<a id="canonical-3212102121132322-0011023023221332-2221131223203112-1133300002322311-2133113123010030-3000023002230322-1010111322331030-2000013121031021"></a>

## prefixes property — destination_prefix_list / 232010002322 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103123133122100-3131302103322101-0132100330223313-3001022022201101-1222322223213200-0002232212311300-3003112211111120-2302210222302001"></a>

## Next pages — destination_prefix_list / 232010002322 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2020300331210132-2230130332021010-0000221103022320-2233032323132332-0032102332020013-1103313022132121-2133133120310222-2132221222210330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221123103020102-2200111303000311-3020012122100123-1131102233013011-0330313331331003-0211310031212120-3310320223200123-1033313101013102"></a>

## rule_list.rules.insert_service — insert_service / 310203222013 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.insert_service

<a id="canonical-0330232121321220-2230033233312333-3030021201123032-2201221321323003-2021100022111003-0130030110331202-3202210311110211-0321012003031232"></a>

Type: `"object"`. single nested block, Optional.

Action to forward traffic to external service.

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
insert_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132133210222203-2322023031032012-2113210122221202-0313131100031103-0020213320331210-3133120300231013-3111022011301130-2321321232023110"></a>

## Direct properties — insert_service / 310203222013 / 3

- [nfv_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3221033101233120-2332300221310232-2203011213223000-2022132332303022-3322003022031011-0311110332303122-0223002033030222-3120212101133123): complete subsection reference.

<a id="canonical-1112032300000303-2103210000210001-1202323231330221-1111013232020011-3112321113033322-1000103212121312-2233301300330101-3312212130223201"></a>

## Next pages — insert_service / 310203222013 / 4

- [rule_list.rules.insert_service.nfv_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3221033101233120-2332300221310232-2203011213223000-2022132332303022-3322003022031011-0311110332303122-0223002033030222-3120212101133123)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3221033101233120-2332300221310232-2203011213223000-2022132332303022-3322003022031011-0311110332303122-0223002033030222-3120212101133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210110200313021-2212110200101302-2333023121012332-1333130102120332-3201033330323211-3012311203231300-3300232132012100-1013230332003111"></a>

## rule_list.rules.insert_service.nfv_service — nfv_service / 122330130131 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2020300331210132-2230130332021010-0000221103022320-2233032323132332-0032102332020013-1103313022132121-2133133120310222-2132221222210330)
- rule_list.rules.insert_service.nfv_service

<a id="canonical-1221013211011302-2100321323003233-2211311201112230-3121002303202023-2331333313333310-2322010300230322-1200233120012023-3303103101332233"></a>

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
nfv_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012112211003302-1312111303332213-0133132222000331-2123022020210121-0132111121132333-2012123230023302-1121111231003030-1230112221033003"></a>

## Direct properties — nfv_service / 122330130131 / 3

<a id="canonical-1132011222310102-1032331023231302-3020213332021100-3332131112030212-1023103333011222-0222031001212101-0030231002310113-1302121320212131"></a>

<a id="canonical-1310013231120302-2310303013001112-1333032201020232-3301310302300323-3013211002313330-3301231021033232-2330032131002230-1212320010222211"></a>

## name property — nfv_service / 122330130131 / 4

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

<a id="canonical-2232101230023231-2032331123210112-0103123133110202-1222122203320310-2131033103223103-3002001001003020-3233332200300302-2002200130212233"></a>

<a id="canonical-0233213103002323-3023033030332322-0103331201001122-0031300203310030-3000031221200111-0310211203133100-1112111122322233-1301001103200233"></a>

## namespace property — nfv_service / 122330130131 / 5

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

<a id="canonical-3032223111010003-2220100131201320-3321232113200031-2303332121322122-3210020211223103-0323102023333010-1010113213331000-3300220222321131"></a>

<a id="canonical-0222030130300023-3320130111330032-3000031213333103-2030313023221311-1311301330210020-2212211310001013-2001312233333020-3021013303100203"></a>

## tenant property — nfv_service / 122330130131 / 6

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

<a id="canonical-2332221023130231-3020010133112220-1203312003230203-3010030022023333-0310220002010120-0233201320012213-2210030013232313-2210302231113111"></a>

## Next pages — nfv_service / 122330130131 / 7

- [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2020300331210132-2230130332021010-0000221103022320-2233032323132332-0032102332020013-1103313022132121-2133133120310222-2132221222210330)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0210123303133133-0121121100000030-3023232322013332-2120310101100302-2013012010102303-2331300111330020-1020310333000322-1322222123200332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213132201123103-0023211211230102-1113033122202032-1322210300332203-3232200301103113-3102101231011031-1030212000020210-0020000331030323"></a>

## rule_list.rules.inside_destinations — inside_destinations / 023010301102 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.inside_destinations

<a id="canonical-1113212031011031-0323210212230020-2323210313201002-0311301330330303-0110200121011213-2020211321200123-0233131303130200-1121033222133011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside destinations.

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
inside_destinations = {}
```

<a id="canonical-2221103232321200-2100101213130103-0031332020102123-2023131323220233-1102221301213002-1312212320200233-2130200213220210-2201300232233003"></a>

## Direct properties — inside_destinations / 023010301102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320113132311031-1201323022211213-1230000023122223-2101013312330020-0232102001003033-3300323303001113-3233321113231010-3200230300032201"></a>

## Next pages — inside_destinations / 023010301102 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3202202312023033-1031023030121320-1221012313301320-3200233222223002-1221300322200302-2123123032312001-3123222131330133-1000211213213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221333101031101-1003131113211232-0211331030101202-3022231220103101-1332123103000110-2132311132333230-0202313113203202-2120311222122102"></a>

## rule_list.rules.inside_sources — inside_sources / 323111120210 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.inside_sources

<a id="canonical-3033302001103222-0223010223210112-1220132222311200-2333112332031133-1123122020210010-3021103113030330-2321303110312021-1202310203013203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside sources.

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
inside_sources = {}
```

<a id="canonical-0320212230033210-1100322212022201-2331211302303030-0011030302031200-0323103012023030-1202120200333230-1020330022302230-0023202110032023"></a>

## Direct properties — inside_sources / 323111120210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230202300323010-2302003130301303-0330103011131000-0022033122111311-0033012232332112-0012012133221022-0131230303200122-2213023030231313"></a>

## Next pages — inside_sources / 323111120210 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0303001303312103-0312032001221303-0132103311122021-2300032320020113-1013123313121012-2113002221232130-1223331023323231-0211322130133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110100312131233-0123112000331020-1231312323003323-0100310111212121-3212022003221323-0231231131332131-1033031322321213-0112312002000100"></a>

## rule_list.rules.label_matcher — label_matcher / 322202023001 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.label_matcher

<a id="canonical-2202330130132032-3100121001123100-1100323122313203-0203322223130222-1023320020032023-2123300020311322-2301020300230321-0202121121330322"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222300022302001-0033021300322300-0301301113002322-0233132020100130-2003032203113303-1331100023102033-1001210012032302-1332232100110031"></a>

## Direct properties — label_matcher / 322202023001 / 3

<a id="canonical-2103221321301300-1020131231020022-2310020001300321-3000032302233302-1300302303030330-2313211100032012-2330113003103133-1232211001110113"></a>

<a id="canonical-2322212120111233-0233001323002203-3133312132312030-0002002300320112-1122223121323021-2032110302001131-3132120203300210-2303022032113112"></a>

## keys property — label_matcher / 322202023001 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0201113023102000-3032130111311110-1021033030110102-0020233332220332-3122122313101212-3033203303112121-2201231123312130-3301133132223001"></a>

## Next pages — label_matcher / 322202023001 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0001102333021033-3310131112313200-3320011333221213-3302322222313113-0220110101212213-0120330022110003-1000030320323323-2100020012323003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302331210031230-0110310022220302-1220121230310030-3200232003200100-2131322233211313-2122130311010301-0120321310213022-1300323021323212"></a>

## rule_list.rules.metadata — metadata / 031101300301 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.metadata

<a id="canonical-1222300233003030-1231232022223031-3232323102303012-2311230313231013-3223230112202101-1113323203302113-2310201310122033-0231332110323022"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312001032231133-2232311121012312-0022213231102220-1323231113100321-2223121031112210-2032213200113131-2031332221120100-3133223020200323"></a>

## Direct properties — metadata / 031101300301 / 3

<a id="canonical-2110032023312203-2301302202113301-1122222120030132-1012030211233323-0113112221010021-1230001110003201-1133120302102020-1020032131322122"></a>

<a id="canonical-3021202121112321-2032133230310101-1011003003322033-2010331330113322-2013202022131013-1102013200030133-3300221230220030-1101330022320310"></a>

## description_spec property — metadata / 031101300301 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2021001021313331-3032203330101210-1233113333202023-0221323011000201-2312102223100302-2220303202111010-0211012121302113-2302212322330032"></a>

<a id="canonical-1131022000202130-1101033333323121-2212301000303320-0112203211312203-3230112333111033-0232332111222000-3231203131232112-0213322322230230"></a>

## name property — metadata / 031101300301 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0112213102212310-2311032233010332-3300003133203310-0331220132303101-1131031113132123-3213223123321233-1312313012012310-3012313020312313"></a>

## Next pages — metadata / 031101300301 / 6

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3122030321331013-3001330001303333-2301200212231230-2113231333332210-1231321221030122-0102111023122121-3310110331212113-1301323121110312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101033031232322-2023321013213022-3203313032130232-0311303030203120-2030021302312023-2212332302322023-1121101111203101-2132332202113223"></a>

## rule_list.rules.outside_destinations — outside_destinations / 311132203123 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.outside_destinations

<a id="canonical-3001121111111320-3130333320210103-0011032301121212-2312133000030310-0023333311322231-1112330320033131-3300333132120303-0210112302022102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside destinations.

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
outside_destinations = {}
```

<a id="canonical-3012123301323333-1203112110102331-0121200121011312-3221002201101311-3211201001002101-2002010301023301-1020330001231030-2301202323233011"></a>

## Direct properties — outside_destinations / 311132203123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020102222002202-3323233203213323-3113322012120022-3112022103202110-2223223320320303-1101201022313310-0311200203210232-0033231111123220"></a>

## Next pages — outside_destinations / 311132203123 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3001020022231030-1302203121112110-0330100103120321-1211022213100232-1030000013220010-3213213123323002-0101120220032133-2031003122331320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203101301230330-1022011223223100-3003230110021221-2032031122332030-2012121113103111-3030220310032122-2111010023030232-3203310220032020"></a>

## rule_list.rules.outside_sources — outside_sources / 330322130323 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.outside_sources

<a id="canonical-2310111011310202-3010121200230221-3012201022011032-0100301232013322-2112303313122000-2132331201321211-3121022232121120-3332312102311200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside sources.

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
outside_sources = {}
```

<a id="canonical-2113320322310333-0022031010221032-1003221032200310-1013011131123300-2322102130211313-1013231321303312-0013212003120112-3100131233211202"></a>

## Direct properties — outside_sources / 330322130323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333323203031123-2110131113102220-0122233003222203-3031112023133301-0013332103111100-0310331321311210-3112001332211113-1131200100322100"></a>

## Next pages — outside_sources / 330322130323 / 4

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2002103111113300-0332102033002033-2132302000221023-3221111201100303-0110202033011102-0333312301123312-1122301212333330-3331110232302232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233323111323003-1232331222011020-1233200332100033-3112013021123120-3321133121213202-1333212312123120-3312002301300323-1020210312012132"></a>

## rule_list.rules.protocol_port_range — protocol_port_range / 123303330232 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.protocol_port_range

<a id="canonical-0213211003321110-1021012310200103-0002301012300313-2123202331002102-3222112200130233-1010131023310333-2203032332330200-3312221121133110"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002012310131300-1333130321331120-3300233110021313-3210330023231301-3231300210321011-1320212211032020-0132013321010211-2310320031023003"></a>

## Direct properties — protocol_port_range / 123303330232 / 3

<a id="canonical-3231202300220023-3213213213122201-3103133303132001-3231100123211211-2110332331221123-2223231332002033-0011112210323013-2201311013122032"></a>

<a id="canonical-2011033011331030-3031303012310011-0103302112233020-1122310110200200-2120012011310202-3300230012002310-1011222111130312-2300211033001020"></a>

## port_ranges property — protocol_port_range / 123303330232 / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-3121303300313220-3133122332331112-1232213333101322-3110321233320223-3100201202222123-3111122310221122-2132230220102030-1303033301211111"></a>

<a id="canonical-0123331222030031-2232021011313202-0021330233213001-0100213020032202-0131010102322032-0320103121121303-2031220002031300-3121223030320311"></a>

## protocol property — protocol_port_range / 123303330232 / 5

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-0330122222210320-1321000202010202-3110112102032123-2021200320111223-1110113123202313-0100313231313302-3022101120012102-0322113202310112"></a>

## Next pages — protocol_port_range / 123303330232 / 6

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0320000103011022-2102323210130022-1320113322021020-1100200102121211-3113320000301032-3222121021012200-2212111133210211-3011201122110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022213232123331-2330131103220212-3033301010032230-3213100202323110-1200231120320222-3033121330030212-2003101302232101-1312000302213212"></a>

## rule_list.rules.source_aws_vpc_ids — source_aws_vpc_ids / 232203201222 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.source_aws_vpc_ids

<a id="canonical-3302010130321123-2002133022312321-3110023232323332-0013101203233233-2021030211123001-0313333220213033-0220202113230133-0021103012300122"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for source aws vpc ids.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("vpc_id")}
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
source_aws_vpc_ids {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033322120321112-0031100320212022-1231331122033113-3013213010323010-2231032021020301-1323003122120122-3010102131131203-2021332133021121"></a>

## Direct properties — source_aws_vpc_ids / 232203201222 / 3

<a id="canonical-3233132012111112-2330303121101013-0233322221022023-0221033001113030-2012101113120221-3000001003111301-0332212102011223-0000323102321320"></a>

<a id="canonical-2110331021323013-1100111133133113-1121111122100130-1113322022132301-0002023223112102-2130110011030013-0013322121311002-0023213200330322"></a>

## vpc_id property — source_aws_vpc_ids / 232203201222 / 4

Type: `["list", "string"]`. Optional.

AWS VPC List. List of VPC Identifiers in AWS.

Upstream description:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3311013310333012-3100221212212330-1320101122301132-2200203022031131-3212200123131213-1330111203322323-0120323200122210-0003323121211202"></a>

## Next pages — source_aws_vpc_ids / 232203201222 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201103012213323-3022213111232101-2003303123213102-2121011102000003-1012231100223222-0232203010301102-0101213131021111-3230000111131320"></a>

## rule_list.rules.source_ip_prefix_set — source_ip_prefix_set / 330010000111 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.source_ip_prefix_set

<a id="canonical-0013330032312112-2220233311203032-3300012232013332-3012323323333113-3332023223201131-1131313003112230-1213213210031322-3132200002213113"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
source_ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113110112333113-3321231130233322-0101133023032021-2023022022322032-0232321312012102-0001322103333100-0331303103213311-3313313312211333"></a>

## Direct properties — source_ip_prefix_set / 330010000111 / 3

- [ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0022332111103323-3221221120230001-0303233200102203-3203113011110333-0121323003133213-3313233233231013-3322023323002212-1213122211002103): complete subsection reference.

<a id="canonical-1202123202031210-0120303201311113-3111300220313213-0111221313103011-3321003210112120-0022113312320320-1331100130022211-3030310300302132"></a>

## Next pages — source_ip_prefix_set / 330010000111 / 4

- [rule_list.rules.source_ip_prefix_set.ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0022332111103323-3221221120230001-0303233200102203-3203113011110333-0121323003133213-3313233233231013-3322023323002212-1213122211002103)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-0022332111103323-3221221120230001-0303233200102203-3203113011110333-0121323003133213-3313233233231013-3322023323002212-1213122211002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302110022101312-1111312131222002-2303121322033103-0222233321323001-1000012011303021-0303302122103330-3222012010301002-3322113321223320"></a>

## rule_list.rules.source_ip_prefix_set.ref — ref / 202011210231 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221)
- rule_list.rules.source_ip_prefix_set.ref

<a id="canonical-3032022013331321-3313121311231100-2221033010111012-0203221223000002-3000123022002223-2323300113320201-3231032130010012-1230312310102031"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013323123131120-2111202113311020-1302022323133003-2200012220303313-3330012333232002-0123011332201320-3033210130023332-3111103110201123"></a>

## Direct properties — ref / 202011210231 / 3

<a id="canonical-3310033003300111-1333203022331230-0002110200213012-1331312202223233-3122221010212133-1331211222310322-3031213233100331-3203130011000301"></a>

<a id="canonical-0300232101132120-2220100200113211-3012202320102321-1100333120222300-1311130300102013-0213010330112010-2312122331203201-0121221132021002"></a>

## kind property — ref / 202011210231 / 4

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

<a id="canonical-0310022112123113-0322022121130221-2303302133121333-1113021101320013-0333013130301122-3321011030333000-0310000321210322-0222322331300202"></a>

<a id="canonical-0301320023113232-1122002201302011-1310132213203031-0123032100310333-1202002301223210-2230122212033002-0010210313230220-1323332332231011"></a>

## name property — ref / 202011210231 / 5

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

<a id="canonical-3103003002113303-0112113300231002-0321311111222101-2123100020011321-1302003132032202-1021231202123313-2203133203010221-2132023031132122"></a>

<a id="canonical-2222211111113211-0121200112123123-3300103122130232-3322300132323231-3123022102100213-2310133310203210-0021330111220221-0123002230023021"></a>

## namespace property — ref / 202011210231 / 6

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

<a id="canonical-2233232320111013-1200321010312211-2231033231302110-2100200022003132-2332301103031313-0000112132210200-3130313302310320-0123130021313222"></a>

<a id="canonical-3003120320001023-3202333331320113-3121222110221203-2030303121013100-2212322221100211-3331103233200020-1003231013023311-0323012223131203"></a>

## tenant property — ref / 202011210231 / 7

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

<a id="canonical-0321122200133011-0300103113203000-1110210230113031-0301021313113303-3133230200232122-2221112212221033-0330232232110010-1331022302011202"></a>

<a id="canonical-1120031030120321-2032130112220130-2030300323121100-0213030110312013-2323102300231332-3312023223032323-0311012311113310-1132223203121020"></a>

## uid property — ref / 202011210231 / 8

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

<a id="canonical-0210332113030213-2100313312212332-3210221223330130-1123031331002111-2113313200000112-3011133121233013-0111322220332123-3201312133233110"></a>

## Next pages — ref / 202011210231 / 9

- [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-2311323100110022-0010031233102023-0223111213203020-3113301302322232-3320133310032212-1021310132231130-3321200101220231-3212313311011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131312020123211-1331210022120220-0213320012121311-0110323031330232-3133210202113120-1003011100112203-1031010020220101-0021212312221231"></a>

## rule_list.rules.source_label_selector — source_label_selector / 103103031001 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.source_label_selector

<a id="canonical-1011100133213022-1220303013330231-2310123003010003-2030030121321201-3311213032330030-3022223311312120-2013223002323311-3331121313130311"></a>

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
source_label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201330121222112-0333112103023201-3321012220032203-0030223130120011-3003310333232202-3223213121101022-2322121203110020-0213303031230203"></a>

## Direct properties — source_label_selector / 103103031001 / 3

<a id="canonical-0311003332100323-1223222330210020-0303131000112230-1220312022311222-2012212210101201-1232110233032121-3121003103201110-0113312311112232"></a>

<a id="canonical-1301003222231131-3000131302022202-1031320332322131-1123233232221032-2001213320002100-3332213230211332-2132232022221010-0322312001031213"></a>

## expressions property — source_label_selector / 103103031001 / 4

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

<a id="canonical-3102132030100201-3323013320103002-1130211202212213-3313131122131213-3323101131313311-0031232331200000-3131101322211213-2022130232322212"></a>

## Next pages — source_label_selector / 103103031001 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-3023232001313222-3033313230302113-3233020302311303-2132130120132131-3232210330000101-2013313212220311-3033112311101010-1210330101321311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221013101202033-1200112002302031-0330212132303311-0323323123033110-3110221003132311-1330001033002300-2321101030310303-0320002220302113"></a>

## rule_list.rules.source_prefix_list — source_prefix_list / 312323321112 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.source_prefix_list

<a id="canonical-3122222301310230-2332123023123000-3031012223203122-2021330303121033-3321320322312313-0121231010011223-3020331233211120-2121011011032111"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
source_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213310320032312-0003011000331303-2322020231032111-1123330032110002-1212111012312103-3323113321001022-1033023303320310-1131300221330230"></a>

## Direct properties — source_prefix_list / 312323321112 / 3

<a id="canonical-2033113011000112-1022010303313111-3131113313303101-3232232200012312-1322223111231201-3321110321302313-2120003110301323-1212211311120312"></a>

<a id="canonical-2112202212010203-0032020111212033-0200320221302300-3202301223201110-0333310221030100-0301313012330312-1221323100212312-1212221211021102"></a>

## prefixes property — source_prefix_list / 312323321112 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2221231233233102-2002011013131330-2130110102001330-2021323320022112-3233331200020332-0303122010131020-3031032133132201-0013132033221210"></a>

## Next pages — source_prefix_list / 312323321112 / 5

- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)

<a id="canonical-1100132312213200-0032121211122030-0200203311120012-2111203021323202-3323203001323231-1123113002211000-3022113223112313-1110213303100000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000323010133031-1111033332013202-1001010310311012-3330123323132201-2223023320320001-1002130031232123-0221012013330312-0000133202312231"></a>

## timeouts — timeouts / 211331020332 / 2

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- timeouts

<a id="canonical-1302112313333020-1301021223132010-1001112211111132-3001222032223310-3230311122020201-3112012000230021-3130011210033221-1303300322030111"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112212032321321-3110222312003211-1012322112130122-2003010012331232-1000013332300211-2012003212303333-3231233031001302-0100120132120233"></a>

## Direct properties — timeouts / 211331020332 / 3

<a id="canonical-0201301120300033-2003112021320030-0030022120331021-3100103103103033-2011003030123322-3122033100010122-1030302002222233-3102002311313011"></a>

<a id="canonical-1323322132200332-3033022220311231-3012221121000323-0212330120022223-3130101302133212-0202223023322121-1120000310312001-1122031322122210"></a>

## create property — timeouts / 211331020332 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0321223300322132-1113323221331020-3111230001222230-0013330033332210-2212213122232133-3013313022322232-3310010013223313-0020211221122012"></a>

<a id="canonical-2313322321330313-0103133121112233-2001011130013033-3032032320303312-1002101213022033-0233311211123111-0112002212313123-0301333103303001"></a>

## delete property — timeouts / 211331020332 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3212203211030210-1101020212231302-0210212233332022-3212301212333031-0230220013010212-1330201012331112-0301002023300111-3123312102203001"></a>

<a id="canonical-1320312002130232-3201223023023021-1131031122002331-2202300210332130-1202123122212222-0003130121121001-2303110310032102-0032301210300312"></a>

## read property — timeouts / 211331020332 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2223202031102230-3031311130301212-3211320303131131-0120223030220121-2012302010321100-0323110231233310-3322031200231002-3222332211101130"></a>

<a id="canonical-1212131321020012-3030211331321020-1020121010030203-2122000332201230-3222200301130332-2331310022300331-1230032213110313-2302100130032332"></a>

## update property — timeouts / 211331020332 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2213222322323010-1120113120221321-0231131202320232-0300310100310120-0223231122322203-2010130323213111-2000112022300231-2311002003230101"></a>

## Next pages — timeouts / 211331020332 / 8

- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
