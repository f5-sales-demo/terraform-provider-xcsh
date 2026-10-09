---
page_title: "xcsh_enhanced_firewall_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_enhanced_firewall_policy reference."
---

# xcsh_enhanced_firewall_policy reference

<a id="canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- Property reference

<a id="canonical-3213323302312002-3211313033323102-0031010220202312-0002001011232311-0302033020320002-0303133310010102-1302310222300100-1202132322030122"></a>

### Direct properties for `xcsh_enhanced_firewall_policy`

- [allow_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1102131223103231-1221212033311132-2212320133110100-2123233220030223-0320001123233230-0220122332332022-1031011132232112-2000130111010012): complete subsection reference.

- [allowed_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1020330031030313-2121223312230100-0130303020203031-0212211201300200-3313233321023213-2031002212123211-2000302131222100-2133000113223233): complete subsection reference.

- [allowed_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0233101011203202-2131101223002310-1303102211012323-0203123013122133-0310123213033203-2230322231320031-3022313211330010-3002230320203231): complete subsection reference.

<a id="canonical-2303121323101112-0310222200213212-0321113310303121-2300332231013321-1010223323231220-2103002220032112-2323103133320021-1132220202103310"></a>

<a id="canonical-0022210003330111-2211220313113022-2310030122030020-2331302303302000-3333303101103112-2202202003103011-2302000011230212-2231032310302220"></a>

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

- [denied_destinations](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2330302201311220-0033000223123322-0002331233031022-1003033033120002-2231321310203322-3303012011303311-0033323313131023-1001131213330232): complete subsection reference.

- [denied_sources](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3030320201232003-3313202120101210-1220013203102130-1220122311333002-2103100102202103-3011023011011210-3322030112202311-0000200001100212): complete subsection reference.

- [deny_all](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2311230322202032-3131231332113201-2233132200232032-2122100001221223-0001203030301022-0101323113220031-3012202300312032-1212221201330030): complete subsection reference.

<a id="canonical-2211001032233033-3222002032012232-0131230022331220-3132333321322210-1000233230222002-0232010232130302-0120312233201201-3330311021020012"></a>

<a id="canonical-1222030033331330-0303320330132222-3002211212311232-0001321133011303-2221131223202312-2330110203002201-0220120100220322-2021123133233313"></a>

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

<a id="canonical-1132032122020211-1310222331113221-0201100313232001-0113330212301323-1221132100132032-1331122013313132-2230002103200333-3121322101233300"></a>

<a id="canonical-1030010300113301-0033302323231101-2100211232033223-0112111202000331-1302033311320103-2223201013303200-0202223311130200-1121130320101311"></a>

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

<a id="canonical-2011031320201011-1322202132211022-2312321302010021-2130000123131003-2033332133102111-2021312103101021-1131302322221111-1302323001020132"></a>

<a id="canonical-1123322001013231-0212030232313332-1111010210033213-2202020330231003-2030333003322210-2131130023213101-3130213230313311-3333122230120311"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1202300011101201-1201121221312133-0203012112010331-3212002211233131-1000302333022203-0322110321110001-0020022333301203-1200221332033301"></a>

<a id="canonical-2132130323322213-1102223032131203-0120130123010233-3010332023303013-1322222233323131-2112113031331213-1003010132121103-2123112300010113"></a>

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

<a id="canonical-1211131100130323-0023213233311033-3332123213212211-2003120000003030-1213232332112313-3333132203210022-2203122123100021-2221300002333330"></a>

<a id="canonical-0313033102130321-0010321320300113-1113233230320231-3231133133310103-2321303013023332-2013000123110313-2022113133201313-3331110310101313"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Enhanced Firewall Policy. Must be unique within the namespace.

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

<a id="canonical-1013212313102133-1131201120313110-0102221000301210-3213203023222233-1300331301311223-1310112023103121-0300211011210223-0113210121332220"></a>

<a id="canonical-1103303020320010-2332001000012110-0120121113230013-1233303132303000-1331212003031301-3012110312222103-1102132333301322-3113113231111001"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Enhanced Firewall Policy is created.

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

- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233): complete subsection reference.

- [timeouts](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1100132312213200-0032121211122030-0200203311120012-2111203021323202-3323203001323231-1123113002211000-3022113223112313-1110213303100000): complete subsection reference.

<a id="canonical-2030032113112001-3000010110030023-2323033100002230-2301131121010210-2333200112102313-1221212323102300-0310130230210122-1031122130321200"></a>

### All schema paths for `xcsh_enhanced_firewall_policy`

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

<a id="canonical-1102131223103231-1221212033311132-2212320133110100-2123233220030223-0320001123233230-0220122332332022-1031011132232112-2000130111010012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allow_all` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- allow_all

<a id="canonical-2133223303012233-1322211203320102-3333133003102201-2032313120001132-1001123310321231-1100102303321221-1221113302131011-3232022011001312"></a>

Type: `["object", {}]`. Optional, Computed.

\[OneOf: allow\_all, allowed\_destinations, allowed\_sources, denied\_destinations, denied\_sources,
deny\_all, rule\_list\] Enable this option. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020330031030313-2121223312230100-0130303020203031-0212211201300200-3313233321023213-2031002212123211-2000302131222100-2133000113223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- allowed_destinations

<a id="canonical-1203212233000123-1212222110033030-0030130000201331-0120300313013213-2023212000302301-2123002113300001-1132331130033312-3031202121032130"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3300310022120122-0312223003330210-2233230202131212-3103100331203133-0223333231013300-1021012332022101-3003303032132213-1112203013012311"></a>

### Direct properties for `allowed_destinations`

<a id="canonical-2001323111123133-1031322002213021-1031132010013300-3230020233023000-1203012022010201-0032310133310112-0302111333023213-1201101320320120"></a>

#### `allowed_destinations.prefix` property

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-0233101011203202-2131101223002310-1303102211012323-0203123013122133-0310123213033203-2230322231320031-3022313211330010-3002230320203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- allowed_sources

<a id="canonical-3110333330110202-3323130322121202-0233302221003111-0131312330101032-2031132122120213-1000231110300131-2132103023031221-2111312321202311"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0021231002021311-0232221322110302-3132312221011330-1233331131023310-2002333030102230-1123103003101322-0301112212030203-1032312132113120"></a>

### Direct properties for `allowed_sources`

<a id="canonical-3030101122232321-3221132312033003-0231330331221213-2110200122101302-1000030022020010-2023332220200101-0331110301322113-1013201310211112"></a>

#### `allowed_sources.prefix` property

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2330302201311220-0033000223123322-0002331233031022-1003033033120002-2231321310203322-3303012011303311-0033323313131023-1001131213330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `denied_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- denied_destinations

<a id="canonical-0211333302313332-2113222200302323-2023112222130331-3220320200300001-2123203203022203-1120033101200103-0303211000110212-1213112012211022"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1232220113211300-2030021301320330-0233231122223112-1013010003331320-2300333000021123-0112230111023200-2323202122221232-0210332301013231"></a>

### Direct properties for `denied_destinations`

<a id="canonical-0010001310002013-1133111323230000-2000323300131013-0012312212132223-0222200321333013-3323331321031121-0330332332333133-0130003031101123"></a>

#### `denied_destinations.prefix` property

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3030320201232003-3313202120101210-1220013203102130-1220122311333002-2103100102202103-3011023011011210-3322030112202311-0000200001100212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `denied_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- denied_sources

<a id="canonical-2101011013200022-2221031222122132-3013110210110103-2311023232112003-2331021012221301-3022330002032310-3012133331301133-1122203311331021"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1233100111103133-1202102010101233-0323100022331100-3032000011202231-0030123301032203-1213223130203220-2013011103222112-1211313113012232"></a>

### Direct properties for `denied_sources`

<a id="canonical-1010013023322010-0320123102100202-1102222330301021-2232022322220133-1201002031301331-0201002002211331-1002223231131331-2131232211201330"></a>

#### `denied_sources.prefix` property

Type: `["list", "string"]`. Optional.

IP Address prefix in string format. String must contain both prefix and prefix-length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2311230322202032-3131231332113201-2233132200232032-2122100001221223-0001203030301022-0101323113220031-3012202300312032-1212221201330030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `deny_all` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- deny_all

<a id="canonical-2000113230323112-1011333330331303-3012121131303002-2003020231023220-2230103012233030-0303111131332102-2212223100003222-2101003033331213"></a>

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
deny_all = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- rule_list

<a id="canonical-2202011303310301-1022220123013310-2111000100022230-3210133001102222-1230300111121233-3110212333200313-2202132223011021-3321221100031111"></a>

Type: `"object"`. single nested block, Optional.

Custom Enhanced Firewall Policy Rules. Custom Enhanced Firewall Policy Rules.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2120313303001300-3231321300033333-3313001131220323-2112303331300220-1302300212031031-2301221002313313-2331130321102102-0030023320001132"></a>

### Direct properties for `rule_list`

- [rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012): complete subsection reference.

<a id="canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules` properties

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
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2000030130100223-0023203232133333-2320312102030133-1320003321033003-2330201103202120-2003221100301023-1320111131023202-1221222213231112"></a>

### Direct properties for `rule_list.rules`

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

<a id="canonical-1321010123302002-2320200102133330-3210131231113103-2212020310013210-3311131332320311-2000010132133233-1232322130322220-1000212000333110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.advanced_action` properties

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

<a id="canonical-1221122223313300-0333321222031000-0201301211112110-1000000232201331-0203200333122222-1200111233233212-0323000322322232-3220123322031130"></a>

### Direct properties for `rule_list.rules.advanced_action`

<a id="canonical-1221323031312322-3120111223123312-2020020331320120-1321210222332330-0000230232123130-0133301230323031-0021130320203020-2301020332312131"></a>

#### `rule_list.rules.advanced_action.action` property

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Additional upstream details:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["LOG","NOLOG"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-3332223003121103-0100232322211100-3233232313010000-2111222012221103-3302201203232012-3201033330313102-0323311033130212-0113320030223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_destinations

<a id="canonical-2123003221131311-1032122200103233-1231003131200012-0311212303111011-2021333200111110-1201112233221103-2011231203231132-0223303320202120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all destinations.

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
all_destinations = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320210302122011-1122312300223112-1112322010203001-0212231203012100-0330023300321111-3231033010233002-1303013130311210-0210100233013311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_sli_vips` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_sli_vips

<a id="canonical-2133123132102300-2131000301031023-1132011201323102-3200312232313331-3330010300133200-0013313303221323-3020201013112330-3120131221020212"></a>

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
all_sli_vips = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210010332032211-1132103201011133-0213113111322302-1210232313222102-3331111112121232-2231133003310122-2031222322122332-1001112012233101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_slo_vips` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_slo_vips

<a id="canonical-1123213030032131-0222020133003300-2202220132133323-1210001032312013-1231010211122030-1111301323123321-3120013111323032-2220321011320331"></a>

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
all_slo_vips = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110331230112330-3102332302302200-2012112312202303-1300223133010212-3331230120330221-0312001000023112-1220302111130221-0102210110320222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_sources

<a id="canonical-2310011032123110-3003202032303212-0210233201112220-3122201113332202-0113223131103220-0323212130323030-2221130331302121-2232003130113100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all sources.

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
all_sources = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233310331200213-3330302312302201-3012302020132221-0010133011202110-0031321002021022-0010330012102032-3333231320022312-3103321102222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_tcp_traffic

<a id="canonical-1201211210311032-2013132033001000-3212121113221103-3133102010202100-0132200203310000-1220220301001003-2111320133102231-0332231333211211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221321120210323-0202302011101332-2330003112113310-3032202011230313-0233212031132210-0120132123110013-0023030311203121-3123120233103002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_traffic` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_traffic

<a id="canonical-0332030211123200-1022030033101211-0031323300320232-0312203332233233-1102210322131111-2222123322330011-1201212212221130-2232013013211023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313102110112012-2300312113203123-1222330300130300-0200110113020333-2102300200121323-0311302331220130-1020020232133232-0102013013201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.all_udp_traffic

<a id="canonical-2102230322013231-0321003201230301-0132113030121222-1310001221002012-0301011100103211-0113233311032022-2003311312030313-3232231032101000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122001303211011-0322112010232320-3031110311233031-1033011011331302-2333332030321022-2113021010112032-1031113103312312-1222000233230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.allow` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.allow

<a id="canonical-1211003102300000-2122203110031310-3202323033232200-1320300202103102-0120230111101220-0211111120330322-2023310222113100-2000330210212320"></a>

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
allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330031202013312-0323002013321210-1313331003023100-1330110103110111-0220123123000332-3133202231132231-3301330023321212-3331132132003003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.applications` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.applications

<a id="canonical-0111330200023022-0020112202113210-0332022220003121-2332022200032111-2002313133030002-2200213320200220-1103332312010102-3332223322021132"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Additional upstream details:

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

<a id="canonical-0223303000021022-3000101300103231-2212113131120331-1011330031323111-0332302313031101-0022030320131203-1001322030301331-3231103312230012"></a>

### Direct properties for `rule_list.rules.applications`

<a id="canonical-0311021022221301-2200202010232201-1201011301223213-3312233011330011-0101320220133230-2312312323130210-3101101103311330-3323231031120211"></a>

#### `rule_list.rules.applications.applications` property

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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

<a id="canonical-1211021022123210-0232100032323000-2203312122200231-3312031112011222-1032121332132330-0131120033320111-3020200233210221-2323000113010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.deny` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.deny

<a id="canonical-0322003203330100-1231313032302110-1101303223001322-0120202230100232-0121001131220023-1200133022302230-1030103001012231-1132101132231201"></a>

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
deny = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203111133203023-1330113002232331-0110303203000213-1122030333032222-1003021330103020-1330113012310301-0203313022022110-1203121113122201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_aws_vpc_ids` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.destination_aws_vpc_ids

<a id="canonical-0213223032210131-1212320013013200-0313012232131001-3313221330332113-1203233010301310-2131310112033032-3212220010211130-2000213323132200"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for destination aws vpc IDs.

Additional upstream details:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3113011212212131-3133322203002010-1101013102012130-0102233322032023-3021320312303302-1131100322003332-2300331301112110-3033221032123011"></a>

### Direct properties for `rule_list.rules.destination_aws_vpc_ids`

<a id="canonical-0211031103223322-0233111102011112-0302302122030203-0300333332320101-1110210202210232-3013102211231231-3311223113002321-2110212132012232"></a>

#### `rule_list.rules.destination_aws_vpc_ids.vpc_id` property

Type: `["list", "string"]`. Optional.

AWS VPC List. List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1331310301100100-2111111311110322-1021212123330203-0302011101203200-3120013012301321-1003123232301313-3202303101011213-2323223111111330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_ip_prefix_set` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.destination_ip_prefix_set

<a id="canonical-0231001002013222-0313002303030213-2313122302120322-2333000023003202-0300213031223012-2133100023200012-3021030310121213-3320301032132321"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1001132020132102-0122223113003313-2220210112123032-0311331222223011-1232302113211101-3230330010232212-3222211221303312-3023213000120331"></a>

### Direct properties for `rule_list.rules.destination_ip_prefix_set`

- [ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3330211030203212-2103032230030131-2231011310221003-1231212023303032-3210101023130100-0020111133102022-2332122322101100-1001313013233030): complete subsection reference.

<a id="canonical-3330211030203212-2103032230030131-2231011310221003-1231212023303032-3210101023130100-0020111133102022-2332122322101100-1001313013233030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [rule_list.rules.destination_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-1331310301100100-2111111311110322-1021212123330203-0302011101203200-3120013012301321-1003123232301313-3202303101011213-2323223111111330)
- rule_list.rules.destination_ip_prefix_set.ref

<a id="canonical-3233211122121213-3032130020011312-1001320222202133-2323223020301223-1011031331330122-2330331200101320-0233301030011012-2222333122122000"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-0021220012133200-2211300303301120-3121011302033333-2023333130232303-2200010110323110-3111330231023332-0312202130101213-1300121103311130"></a>

### Direct properties for `rule_list.rules.destination_ip_prefix_set.ref`

<a id="canonical-2030210122233102-3001033001212321-0320211002123333-0312012222231331-0212011203021201-2032102332021002-3233222002323230-0332210300221210"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.kind` property

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

<a id="canonical-3233022321133102-3222300032312223-2200221011101020-1331202111333322-2132133333201203-0222300101011300-2210212232030010-1002123001021322"></a>

<a id="canonical-1310101110303212-2003233101132103-2022123200110113-0212331301113222-1220003133033002-3111012003100321-3302130212133010-1300000333000030"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0022210031222111-2021000223012102-3100111002101330-0303131223033302-2103122220232032-1031220211300202-1001123302133200-1032001212231223"></a>

<a id="canonical-3012100312301000-1231231100100302-2110231232101201-1200002221222220-1112333000311211-0331233033221010-2312301320011122-0133033302122010"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0132201200111233-1301103220010320-3312110201320012-3303003323022300-3000031231033221-3113301201121113-0223132310223030-3011000300123300"></a>

<a id="canonical-0313033221122031-0202021000312322-3101111000132222-3011111212113031-1031330320220201-1133030111133231-1321130112002222-2332133103002021"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.tenant` property

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

<a id="canonical-3001202010230223-1031231013102223-2021331322331010-0112233230003113-1032320010322112-2003300121123002-1103031233102023-0101010100101133"></a>

<a id="canonical-2132110000021003-2010231020212102-1022030022322221-2310221032113011-2113110030211110-2332200320032320-3123310010030301-2112321022322021"></a>

#### `rule_list.rules.destination_ip_prefix_set.ref.uid` property

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

<a id="canonical-2101213123123320-2132022333123021-3032321233212023-1212213222010002-0012321132202313-3003230002003002-2202233320003221-1003210330000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_label_selector` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.destination_label_selector

<a id="canonical-1000003213001003-3212130311130220-1230000130000111-0122033020233101-0013131311311323-0233212033320121-3232203113223113-1011102122021310"></a>

Type: `"object"`. single nested block, Optional.

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
EnumExtractionComplete: false
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

<a id="canonical-0122302330010120-2223211330022211-1010331203310330-1321312012312200-2202013201132213-2002132002231222-0023230230220113-1123202223132122"></a>

### Direct properties for `rule_list.rules.destination_label_selector`

<a id="canonical-2232201230201211-3230021101020001-0110010132230220-0110232231021022-3220111203330133-2122011330303200-0103322233202202-0230203213121030"></a>

#### `rule_list.rules.destination_label_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2313000201120200-3020021003012101-1022033123301103-1202100121322203-1221333021001201-2021301220322310-2111121332221323-1310303103332033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.destination_prefix_list` properties

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

<a id="canonical-1011030121021111-1013201021301110-1231330200222022-0111031003131023-1022020133021331-3000231100233312-1133323212001012-2011110033112013"></a>

### Direct properties for `rule_list.rules.destination_prefix_list`

<a id="canonical-0233213202100012-1321013220210321-0013102120102032-2120312201202120-3111202320011122-3102303011220223-0110220323303100-0332103112312123"></a>

#### `rule_list.rules.destination_prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2020300331210132-2230130332021010-0000221103022320-2233032323132332-0032102332020013-1103313022132121-2133133120310222-2132221222210330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.insert_service` properties

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

<a id="canonical-1221123103020102-2200111303000311-3020012122100123-1131102233013011-0330313331331003-0211310031212120-3310320223200123-1033313101013102"></a>

### Direct properties for `rule_list.rules.insert_service`

- [nfv_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3221033101233120-2332300221310232-2203011213223000-2022132332303022-3322003022031011-0311110332303122-0223002033030222-3120212101133123): complete subsection reference.

<a id="canonical-3221033101233120-2332300221310232-2203011213223000-2022132332303022-3322003022031011-0311110332303122-0223002033030222-3120212101133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.insert_service.nfv_service` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [rule_list.rules.insert_service](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2020300331210132-2230130332021010-0000221103022320-2233032323132332-0032102332020013-1103313022132121-2133133120310222-2132221222210330)
- rule_list.rules.insert_service.nfv_service

<a id="canonical-1221013211011302-2100321323003233-2211311201112230-3121002303202023-2331333313333310-2322010300230322-1200233120012023-3303103101332233"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2210110200313021-2212110200101302-2333023121012332-1333130102120332-3201033330323211-3012311203231300-3300232132012100-1013230332003111"></a>

### Direct properties for `rule_list.rules.insert_service.nfv_service`

<a id="canonical-1132011222310102-1032331023231302-3020213332021100-3332131112030212-1023103333011222-0222031001212101-0030231002310113-1302121320212131"></a>

#### `rule_list.rules.insert_service.nfv_service.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2012112211003302-1312111303332213-0133132222000331-2123022020210121-0132111121132333-2012123230023302-1121111231003030-1230112221033003"></a>

#### `rule_list.rules.insert_service.nfv_service.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1310013231120302-2310303013001112-1333032201020232-3301310302300323-3013211002313330-3301231021033232-2330032131002230-1212320010222211"></a>

#### `rule_list.rules.insert_service.nfv_service.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0210123303133133-0121121100000030-3023232322013332-2120310101100302-2013012010102303-2331300111330020-1020310333000322-1322222123200332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.inside_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.inside_destinations

<a id="canonical-1113212031011031-0323210212230020-2323210313201002-0311301330330303-0110200121011213-2020211321200123-0233131303130200-1121033222133011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside destinations.

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
inside_destinations = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202202312023033-1031023030121320-1221012313301320-3200233222223002-1221300322200302-2123123032312001-3123222131330133-1000211213213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.inside_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.inside_sources

<a id="canonical-3033302001103222-0223010223210112-1220132222311200-2333112332031133-1123122020210010-3021103113030330-2321303110312021-1202310203013203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside sources.

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
inside_sources = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303001303312103-0312032001221303-0132103311122021-2300032320020113-1013123313121012-2113002221232130-1223331023323231-0211322130133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.label_matcher` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.label_matcher

<a id="canonical-2202330130132032-3100121001123100-1100323122313203-0203322223130222-1023320020032023-2123300020311322-2301020300230321-0202121121330322"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1110100312131233-0123112000331020-1231312323003323-0100310111212121-3212022003221323-0231231131332131-1033031322321213-0112312002000100"></a>

### Direct properties for `rule_list.rules.label_matcher`

<a id="canonical-2103221321301300-1020131231020022-2310020001300321-3000032302233302-1300302303030330-2313211100032012-2330113003103133-1232211001110113"></a>

#### `rule_list.rules.label_matcher.keys` property

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0001102333021033-3310131112313200-3320011333221213-3302322222313113-0220110101212213-0120330022110003-1000030320323323-2100020012323003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.metadata` properties

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
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0302331210031230-0110310022220302-1220121230310030-3200232003200100-2131322233211313-2122130311010301-0120321310213022-1300323021323212"></a>

### Direct properties for `rule_list.rules.metadata`

<a id="canonical-2110032023312203-2301302202113301-1122222120030132-1012030211233323-0113112221010021-1230001110003201-1133120302102020-1020032131322122"></a>

#### `rule_list.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2021001021313331-3032203330101210-1233113333202023-0221323011000201-2312102223100302-2220303202111010-0211012121302113-2302212322330032"></a>

<a id="canonical-2312001032231133-2232311121012312-0022213231102220-1323231113100321-2223121031112210-2032213200113131-2031332221120100-3133223020200323"></a>

#### `rule_list.rules.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3122030321331013-3001330001303333-2301200212231230-2113231333332210-1231321221030122-0102111023122121-3310110331212113-1301323121110312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.outside_destinations` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.outside_destinations

<a id="canonical-3001121111111320-3130333320210103-0011032301121212-2312133000030310-0023333311322231-1112330320033131-3300333132120303-0210112302022102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside destinations.

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
outside_destinations = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001020022231030-1302203121112110-0330100103120321-1211022213100232-1030000013220010-3213213123323002-0101120220032133-2031003122331320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.outside_sources` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.outside_sources

<a id="canonical-2310111011310202-3010121200230221-3012201022011032-0100301232013322-2112303313122000-2132331201321211-3121022232121120-3332312102311200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside sources.

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
outside_sources = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002103111113300-0332102033002033-2132302000221023-3221111201100303-0110202033011102-0333312301123312-1122301212333330-3331110232302232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.protocol_port_range

<a id="canonical-0213211003321110-1021012310200103-0002301012300313-2123202331002102-3222112200130233-1010131023310333-2203032332330200-3312221121133110"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

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

<a id="canonical-1233323111323003-1232331222011020-1233200332100033-3112013021123120-3321133121213202-1333212312123120-3312002301300323-1020210312012132"></a>

### Direct properties for `rule_list.rules.protocol_port_range`

<a id="canonical-3231202300220023-3213213213122201-3103133303132001-3231100123211211-2110332331221123-2223231332002033-0011112210323013-2201311013122032"></a>

#### `rule_list.rules.protocol_port_range.port_ranges` property

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2002012310131300-1333130321331120-3300233110021313-3210330023231301-3231300210321011-1320212211032020-0132013321010211-2310320031023003"></a>

#### `rule_list.rules.protocol_port_range.protocol` property

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ALL","ICMP","TCP","UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-0320000103011022-2102323210130022-1320113322021020-1100200102121211-3113320000301032-3222121021012200-2212111133210211-3011201122110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_aws_vpc_ids` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.source_aws_vpc_ids

<a id="canonical-3302010130321123-2002133022312321-3110023232323332-0013101203233233-2021030211123001-0313333220213033-0220202113230133-0021103012300122"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for source aws vpc IDs.

Additional upstream details:

List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1022213232123331-2330131103220212-3033301010032230-3213100202323110-1200231120320222-3033121330030212-2003101302232101-1312000302213212"></a>

### Direct properties for `rule_list.rules.source_aws_vpc_ids`

<a id="canonical-3233132012111112-2330303121101013-0233322221022023-0221033001113030-2012101113120221-3000001003111301-0332212102011223-0000323102321320"></a>

#### `rule_list.rules.source_aws_vpc_ids.vpc_id` property

Type: `["list", "string"]`. Optional.

AWS VPC List. List of VPC Identifiers in AWS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_ip_prefix_set` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.source_ip_prefix_set

<a id="canonical-0013330032312112-2220233311203032-3300012232013332-3012323323333113-3332023223201131-1131313003112230-1213213210031322-3132200002213113"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0201103012213323-3022213111232101-2003303123213102-2121011102000003-1012231100223222-0232203010301102-0101213131021111-3230000111131320"></a>

### Direct properties for `rule_list.rules.source_ip_prefix_set`

- [ref](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0022332111103323-3221221120230001-0303233200102203-3203113011110333-0121323003133213-3313233233231013-3322023323002212-1213122211002103): complete subsection reference.

<a id="canonical-0022332111103323-3221221120230001-0303233200102203-3203113011110333-0121323003133213-3313233233231013-3322023323002212-1213122211002103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- [rule_list.rules.source_ip_prefix_set](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3012003330101000-0131122022122220-1121200200131330-2121103312111201-1332320121321321-2222100102313123-1311212121200110-0010201111031221)
- rule_list.rules.source_ip_prefix_set.ref

<a id="canonical-3032022013331321-3313121311231100-2221033010111012-0203221223000002-3000123022002223-2323300113320201-3231032130010012-1230312310102031"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-1302110022101312-1111312131222002-2303121322033103-0222233321323001-1000012011303021-0303302122103330-3222012010301002-3322113321223320"></a>

### Direct properties for `rule_list.rules.source_ip_prefix_set.ref`

<a id="canonical-3310033003300111-1333203022331230-0002110200213012-1331312202223233-3122221010212133-1331211222310322-3031213233100331-3203130011000301"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.kind` property

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

<a id="canonical-0310022112123113-0322022121130221-2303302133121333-1113021101320013-0333013130301122-3321011030333000-0310000321210322-0222322331300202"></a>

<a id="canonical-2013323123131120-2111202113311020-1302022323133003-2200012220303313-3330012333232002-0123011332201320-3033210130023332-3111103110201123"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3103003002113303-0112113300231002-0321311111222101-2123100020011321-1302003132032202-1021231202123313-2203133203010221-2132023031132122"></a>

<a id="canonical-0300232101132120-2220100200113211-3012202320102321-1100333120222300-1311130300102013-0213010330112010-2312122331203201-0121221132021002"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2233232320111013-1200321010312211-2231033231302110-2100200022003132-2332301103031313-0000112132210200-3130313302310320-0123130021313222"></a>

<a id="canonical-0301320023113232-1122002201302011-1310132213203031-0123032100310333-1202002301223210-2230122212033002-0010210313230220-1323332332231011"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.tenant` property

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

<a id="canonical-0321122200133011-0300103113203000-1110210230113031-0301021313113303-3133230200232122-2221112212221033-0330232232110010-1331022302011202"></a>

<a id="canonical-2222211111113211-0121200112123123-3300103122130232-3322300132323231-3123022102100213-2310133310203210-0021330111220221-0123002230023021"></a>

#### `rule_list.rules.source_ip_prefix_set.ref.uid` property

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

<a id="canonical-2311323100110022-0010031233102023-0223111213203020-3113301302322232-3320133310032212-1021310132231130-3321200101220231-3212313311011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_label_selector` properties

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../resources/enhanced_firewall_policy.md#canonical-0230122003121113-2121322323013210-0332210003313031-3130130121001020-0320301310111013-2131011113313110-1232313200113310-3023330120001120)
- [Property reference](resources--enhanced_firewall_policy--reference--group-001.md#canonical-3120211302012031-2120323122122020-0210103121120212-1321311022000233-2100211332133033-0032332233212111-3203333123232212-3220222311000331)
- [rule_list](resources--enhanced_firewall_policy--reference--group-001.md#canonical-2023223000011102-0132201101331330-1113221301031013-3010123023223223-3021032333033122-2112013320331320-3013110101001212-1111220202322233)
- [rule_list.rules](resources--enhanced_firewall_policy--reference--group-001.md#canonical-0202310021233130-3232223030000020-2211131230200110-3101110213101132-3300322230203330-2322112220231103-0032111321020220-2311302211030012)
- rule_list.rules.source_label_selector

<a id="canonical-1011100133213022-1220303013330231-2310123003010003-2030030121321201-3311213032330030-3022223311312120-2013223002323311-3331121313130311"></a>

Type: `"object"`. single nested block, Optional.

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
EnumExtractionComplete: false
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

<a id="canonical-0131312020123211-1331210022120220-0213320012121311-0110323031330232-3133210202113120-1003011100112203-1031010020220101-0021212312221231"></a>

### Direct properties for `rule_list.rules.source_label_selector`

<a id="canonical-0311003332100323-1223222330210020-0303131000112230-1220312022311222-2012212210101201-1232110233032121-3121003103201110-0113312311112232"></a>

#### `rule_list.rules.source_label_selector.expressions` property

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3023232001313222-3033313230302113-3233020302311303-2132130120132131-3232210330000101-2013313212220311-3033112311101010-1210330101321311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rule_list.rules.source_prefix_list` properties

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

<a id="canonical-2221013101202033-1200112002302031-0330212132303311-0323323123033110-3110221003132311-1330001033002300-2321101030310303-0320002220302113"></a>

### Direct properties for `rule_list.rules.source_prefix_list`

<a id="canonical-2033113011000112-1022010303313111-3131113313303101-3232232200012312-1322223111231201-3321110321302313-2120003110301323-1212211311120312"></a>

#### `rule_list.rules.source_prefix_list.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1100132312213200-0032121211122030-0200203311120012-2111203021323202-3323203001323231-1123113002211000-3022113223112313-1110213303100000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

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

<a id="canonical-0000323010133031-1111033332013202-1001010310311012-3330123323132201-2223023320320001-1002130031232123-0221012013330312-0000133202312231"></a>

### Direct properties for `timeouts`

<a id="canonical-0201301120300033-2003112021320030-0030022120331021-3100103103103033-2011003030123322-3122033100010122-1030302002222233-3102002311313011"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0321223300322132-1113323221331020-3111230001222230-0013330033332210-2212213122232133-3013313022322232-3310010013223313-0020211221122012"></a>

<a id="canonical-1112212032321321-3110222312003211-1012322112130122-2003010012331232-1000013332300211-2012003212303333-3231233031001302-0100120132120233"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3212203211030210-1101020212231302-0210212233332022-3212301212333031-0230220013010212-1330201012331112-0301002023300111-3123312102203001"></a>

<a id="canonical-1323322132200332-3033022220311231-3012221121000323-0212330120022223-3130101302133212-0202223023322121-1120000310312001-1122031322122210"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2223202031102230-3031311130301212-3211320303131131-0120223030220121-2012302010321100-0323110231233310-3322031200231002-3222332211101130"></a>

<a id="canonical-2313322321330313-0103133121112233-2001011130013033-3032032320303312-1002101213022033-0233311211123111-0112002212313123-0301333103303001"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
