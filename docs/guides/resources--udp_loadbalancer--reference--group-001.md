---
page_title: "xcsh_udp_loadbalancer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer reference."
---

# xcsh_udp_loadbalancer reference

<a id="canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- Property reference

<a id="canonical-0102011000323030-1231123021231123-1101302021212220-3133203132301222-3012023212201021-0300120322320013-1121211130202000-2323310302130303"></a>

### Direct properties for `xcsh_udp_loadbalancer`

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132): complete subsection reference.

- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030): complete subsection reference.

- [advertise_on_public](resources--udp_loadbalancer--reference--group-002.md#canonical-3121001210223201-0002101322213023-2213011332023133-0223131303332232-0202230121130203-0330221122033201-1300103332030211-1301020220301201): complete subsection reference.

- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-002.md#canonical-1032133030103213-0020211103223323-3100222332020210-1003331202302203-3333101310331003-1200101122201012-2332103110302022-2333031032112103): complete subsection reference.

<a id="canonical-3222130132230203-1330303011120303-1232200123100032-0132132113101002-3132000322223312-2312302123232131-2112321120213203-0032222221211012"></a>

<a id="canonical-0302210303111330-1121023232013102-2023231020132233-1231001022032223-1222313223212211-3322210322233332-0333313220303022-3232231311201210"></a>

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

<a id="canonical-0030130001211030-1031310321030131-1233200002003020-3313010122313311-3202203320102020-1201010112311003-1003012111330021-0030131003120220"></a>

<a id="canonical-0030312303002002-3113001223033213-0301212323330221-3330300001303231-1120220300120303-3012200321220110-3301033303002103-0312303131133020"></a>

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

<a id="canonical-3220320121201312-2313103212111113-0133023103321002-0311011113012100-0211101210312111-0312333220323122-1313213300232332-0130322012100202"></a>

<a id="canonical-1202202213001333-2203210213312132-0110200021010021-0110031030210000-1001230303100221-2203221131013221-1033133333200233-0112003233322001"></a>

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

<a id="canonical-1302310332320200-1300100212123002-2333002132121001-2201231200302223-3112330310200230-1003003201100221-2012212203012231-2213333200312200"></a>

<a id="canonical-3233221003013231-2122200203031120-1223213132022001-0031211230123211-0232221120232303-3120313013313122-2221120033012122-1213101020323302"></a>

#### `dns_volterra_managed` property

Type: `"bool"`. Optional, Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain to be delegated to F5 Distributed Cloud using the Delegated Domain feature or a DNS CNAME
record must be created in your DNS provider's portal.

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

- [do_not_advertise](resources--udp_loadbalancer--reference--group-002.md#canonical-2021012000001001-1013233020022212-1310012310002230-0213221320021221-3313312011223103-0233332113202000-2001321301000221-1032223300220320): complete subsection reference.

<a id="canonical-1013201121013122-0113023330111203-3311020120112000-0113231223202132-1203222130212210-3031320332122302-3202112231022012-2223220301123230"></a>

<a id="canonical-1031023231300010-2122122322320102-1203300320202200-3231103212221131-1133130122210200-1002010313110102-0012001313003100-0200000323333323"></a>

#### `domains` property

Type: `["list", "string"]`. Optional.

A list of domains (host/authority header) that will be matched to this load balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-002.md#canonical-2130020320231110-2310332213003130-2103302031122120-2101333103230022-2103002230301132-0101132133321211-1003113021330223-1110002031002313): complete subsection reference.

- [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-002.md#canonical-0033000022201003-1303032013200210-1033221223122313-1001020223133221-0221333001113000-1302131100323201-0111021013101131-1302332010012121): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-002.md#canonical-1232202330301332-2233232122032300-3311123230120100-0221111330223031-0212333133112002-2320202110313302-2212102131223121-2130011100002212): complete subsection reference.

<a id="canonical-1103203003221010-2302130202212110-3101213302333013-2022223310223030-0332030023112203-3233232032222111-2030212313021023-1203232210313301"></a>

<a id="canonical-1230200121233312-2201012230133122-3300111202230120-2212331233003200-1111002211303333-1330233310301132-3023300210303200-3010122330030223"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0003230310003122-0322010033032300-2122021310021103-3321212011333203-2322002230123233-2321000023131302-3031321320013010-0001303323100132"></a>

<a id="canonical-3101313031200323-0213102030012223-0221331330102032-3111300310211030-1233200301301301-0302230123032031-1200232312323013-2321212001002010"></a>

#### `idle_timeout` property

Type: `"number"`. Optional, Computed.

The amount of time that a session can exist without upstream or downstream activity, in
milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(30000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
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
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-3130021300010131-2111113102303010-3103211221121321-1322130133021023-1202023211031231-1121202212031020-1023130023221133-1331211031220232"></a>

<a id="canonical-2320023022202131-2313202231103232-3332031121021303-3222323303112012-0211213200103222-1120013013103303-1132331203003020-2310123210212012"></a>

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

<a id="canonical-3120212203132202-0311303020101110-1201223210101333-3123231111113202-2203132021302012-1320322032022021-0212220011012001-2210010130102221"></a>

<a id="canonical-1320032223020101-2300200121012321-0021232011311222-2010202132001300-1221330133112022-2003020210333131-2031203113202002-2021203300100300"></a>

#### `listen_port` property

Type: `"number"`. Optional, Computed.

\[OneOf: listen\_port, port\_ranges\] Exclusive with \[port\_ranges\] Listen Port for this load
balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(65535),
}
```

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [listen_port](resources--udp_loadbalancer--reference--group-001.md#canonical-3120212203132202-0311303020101110-1201223210101333-3123231111113202-2203132021302012-1320322032022021-0212220011012001-2210010130102221)
- [port_ranges](resources--udp_loadbalancer--reference--group-001.md#canonical-2203020003111122-2333012023312122-1231102021223222-3213310221222000-3131120111002221-1213123101003133-3330021213323321-0022112233110231)

Select alternatives according to the provider validators above.

<a id="canonical-0331330200033333-3232002310120311-0210010030212312-3112231100111230-3100201302123333-3110022211030132-0022130101300002-2331220311033232"></a>

<a id="canonical-3122330213001221-3212333020302302-2133132020130023-1310113121021113-1301000111120012-1032130301100133-2323011210311011-3230031333323101"></a>

#### `name` property

Type: `"string"`. Required.

Name of the UDP Load Balancer. Must be unique within the namespace.

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

<a id="canonical-1021220003100200-3032022000022230-2312302212033333-2133133223231033-1102312102012232-2200332211113023-3312210121013001-2031033322201323"></a>

<a id="canonical-0133322131233301-2232100103132233-1012031100320221-0002112012032200-1131131130003112-1012223300000003-2003230223102131-2000000020230231"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the UDP Load Balancer is created.

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

- [no_service_policies](resources--udp_loadbalancer--reference--group-002.md#canonical-2222312133022213-1230023330021312-1132323131113230-2333332022331122-2130002003301133-2032133133210033-3100313020020301-0031233232230012): complete subsection reference.

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-002.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303): complete subsection reference.

<a id="canonical-2203020003111122-2333012023312122-1231102021223222-3213310221222000-3131120111002221-1213123101003133-3330021213323321-0022112233110231"></a>

<a id="canonical-2231210113203210-2032320210033330-0231001123023233-1100111132323012-1131301203001020-3302131202203222-2133120113312020-2300110123230021"></a>

#### `port_ranges` property

Type: `"string"`. Optional, Computed.

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2221100322323133-3222100232310313-2133010130310031-3231202023101032-3313112200110100-2130133221200013-2211012323333211-0003213013303230): complete subsection reference.

- [timeouts](resources--udp_loadbalancer--reference--group-002.md#canonical-1222013031123323-0132302211031313-1113201031020033-3330303023103321-1132201210330120-3022032110001231-2010201311132220-1032103313302111): complete subsection reference.

- [udp](resources--udp_loadbalancer--reference--group-002.md#canonical-3121113233300112-1011301020012201-3133332123031100-1313012230000123-2323010011202022-1122033233122001-3001313032003221-2033202233101030): complete subsection reference.

<a id="canonical-2220220322302021-2322011031202212-0201313030003010-1101021301111330-0130102231112220-3303010113311222-1231001302103323-2000033201112312"></a>

### All schema paths for `xcsh_udp_loadbalancer`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-1113100321101330-2330122113122300-3330103021013021-0221113231210202-1011123032032021-2100012033131203-3023303231210230-1020000132320220) |
| `active_service_policies.policies` | [active_service_policies.policies](resources--udp_loadbalancer--reference--group-001.md#canonical-3021320330233221-0222013013000013-1220031310010012-3202203200110213-2130331211011012-1301022232313103-1213113323323203-0111133101132032) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](resources--udp_loadbalancer--reference--group-001.md#canonical-1311033330032101-2300331021132212-1221332122033331-3033113123012233-0213102312003320-0300113213333022-2310311200211122-1010013333322312) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-3000133231300130-1020132333212023-0223022202330320-1012330112200330-2311210102112301-2003123111313310-3300110222212203-2101030330100301) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-0103332210210101-1301013220210033-0031312333333130-3313302111102123-0031033001211113-2101223030220220-2221313302112032-0031323100321112) |
| `advertise_custom` | [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-2220112212031221-3223210202313233-2213100202202330-2021221131032103-0013010323313323-0221003321333303-0031200122113110-3111031023222021) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0010210202120001-2110301320213033-2233123203021222-3120002123302203-2233020312021012-3331301220332222-2033031003301301-2011210122212033) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-2330101022221233-0323020331311102-2313220303200200-3021103320222020-3123210321020111-1021223030233121-3322011003300012-1133311121230300) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-3112213021302113-0002230213003012-3230112030022130-1212313003313221-1031112012101320-0213332100130330-0131011111320012-2321030021102121) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-1013221200302123-3030233133311111-1232133032312322-2212301023012220-2212011220230032-1103123110320213-3111302100103002-2000201123321000) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-3223330100233103-2301222122312032-0313023120233121-1021031220220210-0212300101003212-2220223121333212-3232210323223302-1201121223130023) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-0230110232112000-1333301113230220-1331021111030100-0202310202033301-3131221301330310-3330313010231002-2231031000033201-3210223201233210) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-2131202302333100-1213210221132131-0321032022213330-1110323133012313-0113122100221113-2233133223303322-1331323310230113-3221133112232231) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-3101323200112133-2321023213200231-0322110200323002-0211032333233331-3012303123012110-1231212020302311-2022010203132321-0021132012333031) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-2213231003223302-2100230002022120-1333311102132102-2022300331031303-2122231221131203-2233013123212212-1111330102223222-1113230230103110) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-1011023001230302-0021013003233311-2032203330113113-0023210213330332-3233312001231031-0013030320113031-0130202002211222-0213331301232303) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-1301101221213203-1010023330200323-2103002212020023-0022021301230103-2110213021210003-3132023022213112-1020132233012023-3112333322313320) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0122331100201003-3101012133111122-2130312131030232-0223303230202230-2032311210201313-3323120312311321-3033331311313212-2112113232231213) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-3301033300023313-1333000103210230-2213202102221310-3113332202123021-2311102121303313-0210331011232021-3130022113303110-2233132102332030) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-3023200020300232-3111321122001302-3310220203232113-2333203321133303-3012101233201231-3011012313101221-3120332212230212-3210221220023330) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-3210002303000122-0331030031202022-2012000001311302-2021120133312300-0302322300300033-1011001133213221-1300230122123013-0123010121001202) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-3323232012230010-2310233302301203-2001201330110301-0302111302032202-3111102212013123-1323123210331121-1221233312133110-2232321102212320) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](resources--udp_loadbalancer--reference--group-001.md#canonical-0200332010112132-2310201102022000-0022022212003112-1311032020020020-3200323200013102-2303303310002221-1231122031223200-0232203030001012) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](resources--udp_loadbalancer--reference--group-001.md#canonical-2123113113323331-3013303033210321-3121021112202033-0013320131230020-0113003123233132-2303020310110112-3033330210213221-1201222112233122) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-0020132103030020-0230322133213331-2320031213330301-3000000312230113-0121103223120102-3000211210122130-0130323030031101-0022321020330212) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](resources--udp_loadbalancer--reference--group-001.md#canonical-0001312212020110-1202310032031201-2011002230102320-0122323000132003-1200321300223123-3021133102111212-1132210311120121-2222023003230103) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](resources--udp_loadbalancer--reference--group-001.md#canonical-2300030131202232-3233002021113223-3123302322322002-0232023030102011-2311013133033320-1230003022330101-0020330003111112-1103320202221000) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](resources--udp_loadbalancer--reference--group-001.md#canonical-1130310233223332-1112003313113203-0011222131122021-3012201213102232-2213331303133100-0101231310323200-2322202203131110-3201213023003312) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-2131012300120111-3031331311202323-3101123013212110-2100330002021100-3020332130203323-1203000020302211-2220130111302111-1033133001300002) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-2030303110222032-2322321213102331-3111300310132311-2333322200322123-0333333132131201-0002111112023101-0112321031220103-0321103220011001) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-0320110210123322-3330211031223002-0330133221111100-3120312111231020-3320213101120110-0322221210200110-1011221122031312-0302222003310112) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](resources--udp_loadbalancer--reference--group-001.md#canonical-0300323200230300-3103131322311021-3330001032031221-2111311300013113-0320323221110301-3120221322122312-2033101120111012-0121201003332103) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2301211302121211-0101311102033333-3313131212300320-0230213313103000-2301120220321213-0123302301220132-2220203212221331-3220000231320132) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-2030211222310200-1113310031020312-2232111130022321-0102323110030321-0310220132130221-0321231011210201-3123110103012031-0322222201011031) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-0223023012130320-3103013302333322-0230011111233230-1200022210103231-1011121323312221-1133211202232202-1202130131031222-1321301023213123) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-0133102222123121-1230021120213020-0020222033320020-1330202022022011-1100211103310102-2220231120000323-0033233013220311-3322001302121101) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-0332202231203330-0303200210131011-0322211201002321-0311203101301201-0020223011210120-0220031132330133-2131221121002220-0120032013232303) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](resources--udp_loadbalancer--reference--group-002.md#canonical-2122203103003233-3201300211121110-0330330111103322-2020213320323320-2001322032221110-0122122110332010-2000030321123323-0320231231213013) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--udp_loadbalancer--reference--group-002.md#canonical-2120122030031333-3330321132323310-0132313111232022-1200012201303120-2333013033333232-0011231111321322-1021300303023113-3210003301220121) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2012000312332231-2123200111311000-2331003210330213-2331222200032033-3323112013311120-1330320131220113-2130023231303230-0123003303010232) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-2222310211002121-2331223021033120-2031312211331020-0211313303231102-2011003101112213-0213121121303121-1210223121201222-1330212203323013) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-002.md#canonical-3232020011102110-0020333100132333-3232011113311030-2220013311320103-1132303133321301-2303120031101100-3300222233102323-3021210321123031) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](resources--udp_loadbalancer--reference--group-002.md#canonical-1220122213020333-2020200311011000-2303003102322013-1032213333102123-3121011110030111-0121220310232313-2000031331103010-0322203313233003) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](resources--udp_loadbalancer--reference--group-002.md#canonical-1331120000120221-0330321131302130-0310332321121112-3023220031002021-3013312012000202-0123121010321133-3220323121131213-0210112033103311) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--udp_loadbalancer--reference--group-002.md#canonical-0312332310312203-1321122130103023-1031332310130033-2203311203321201-3011203323321102-0110113103222011-1312131122031320-2100101133222033) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2333023110112131-0210000303033021-2231021221110201-2001201011202133-2312020211020131-2121123103231233-1220012320332320-3102111231221221) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-2333112210200122-3303011203301130-1103212013012133-1211303322102000-1003033012303212-0110323112110021-1312031000232010-2030011001000230) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-002.md#canonical-1303013321030132-3232223121202231-0012113213203010-2200001302323001-3322321331021132-2221321210121033-3300221020210022-0013213110232210) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--udp_loadbalancer--reference--group-002.md#canonical-1330121220100002-1220123310312003-1320112230101013-3220220210022110-1120220033121312-2103213310021002-1231110233103002-1001103131202132) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](resources--udp_loadbalancer--reference--group-002.md#canonical-0111111112300222-1132203011231122-2001013120003031-3011221012210011-2032023123302132-0310233223313333-1323111123211012-0120010311000232) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--udp_loadbalancer--reference--group-002.md#canonical-1203330233032011-0221303100302323-1013131233023130-1220131103001322-1001200032322031-0122001232212021-3121103223323300-1023332312211132) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--udp_loadbalancer--reference--group-002.md#canonical-3203012213202103-3330030321011311-2202322103300113-1012333301001201-3021301122312203-2100022023202313-3002233211000032-3003101213003021) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-1011001132113301-3333010013000033-2130121131300223-2330133132330223-3131200332032211-0131310022321003-2310022131323110-2020113013223311) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-0220312123212233-1232202230112321-2303231212010100-3333030103101131-1033113113002300-3033222113103012-3012230313320120-0211032021203111) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-002.md#canonical-0132203110033033-1223031113332002-0020213220310322-0221232200022322-2212022032321002-1110322202210223-0102202233220312-1011233031230323) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](resources--udp_loadbalancer--reference--group-002.md#canonical-1110001011323311-2302202330233110-0030121101122132-2032230213311300-1030010322332013-0020003130312311-1321100213210311-3003330122302000) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](resources--udp_loadbalancer--reference--group-002.md#canonical-2303201300332222-0012023001221030-2113203321232321-3231230231212011-0122203213302032-1311322200023221-3310122030312213-3020020223220212) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-0011321330122213-1202003011010120-0313333320213001-1302223223230013-1133100303301313-2022232002233302-2011100130210212-1001133300212123) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-1131120112221032-0110130201123313-1111321333200312-1010113211133120-0032203012331121-3003013221013123-1012201322313032-1002302303132010) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--udp_loadbalancer--reference--group-002.md#canonical-3301030232010212-1201333333221333-3310223011013103-2231101220012123-0200213221013011-3213332303133201-2211330111302313-2012013013112000) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--udp_loadbalancer--reference--group-002.md#canonical-0202021221332203-1203212223132301-0203113113012313-3113212100232123-2232110111013300-3123110021110211-3323132231020121-1330321200011213) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-1113101012213132-0332023102322020-0210320110223211-1030001331123032-3121130222033210-1301222221210303-0211020331221111-3203030131321033) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-0131000012320331-1133122020002330-0233321133012322-1201203210313301-2210011010123100-1001112123211002-3323323230001331-0000122202203020) |
| `advertise_on_public` | [advertise_on_public](resources--udp_loadbalancer--reference--group-002.md#canonical-1131331003210000-2331113000100013-2232131210301213-1021112230322000-1222213303123001-0302322101033001-0010131203112121-2123220013211113) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-002.md#canonical-0300103122321023-3000132322101211-2230222020221133-0012031321211210-2300030331103320-2201031002210020-2330301221230310-3212123202133031) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-002.md#canonical-0111110121032203-2332103123111313-1100013111131210-0013023230231022-3132112323122112-1332210211001321-0021321201312111-2023000211312332) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2203221331301322-0133301002111213-0033001211222300-3320221322123331-3333103023023023-3012333302023113-3330113300030000-1321202220232130) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-3322133321113131-1003200203131002-3122112203331212-3023221301112203-0100232120221313-0033022213220131-1102210230322132-1121230313021322) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-002.md#canonical-2302323123201032-2001211223302111-0132220033313002-3021301203111312-0321201310212003-0323332122132312-1131330222320133-2310133320001233) |
| `annotations` | [annotations](resources--udp_loadbalancer--reference--group-001.md#canonical-3222130132230203-1330303011120303-1232200123100032-0132132113101002-3132000322223312-2312302123232131-2112321120213203-0032222221211012) |
| `description` | [description](resources--udp_loadbalancer--reference--group-001.md#canonical-0030130001211030-1031310321030131-1233200002003020-3313010122313311-3202203320102020-1201010112311003-1003012111330021-0030131003120220) |
| `disable` | [disable](resources--udp_loadbalancer--reference--group-001.md#canonical-3220320121201312-2313103212111113-0133023103321002-0311011113012100-0211101210312111-0312333220323122-1313213300232332-0130322012100202) |
| `dns_volterra_managed` | [dns_volterra_managed](resources--udp_loadbalancer--reference--group-001.md#canonical-1302310332320200-1300100212123002-2333002132121001-2201231200302223-3112330310200230-1003003201100221-2012212203012231-2213333200312200) |
| `do_not_advertise` | [do_not_advertise](resources--udp_loadbalancer--reference--group-002.md#canonical-0221213122111200-3301201210111322-3031222200121111-3003321003021322-3220301223013001-0210223020203022-0323301031102012-3132332221032200) |
| `domains` | [domains](resources--udp_loadbalancer--reference--group-001.md#canonical-1013201121013122-0113023330111203-3311020120112000-0113231223202132-1203222130212210-3031320332122302-3202112231022012-2223220301123230) |
| `hash_policy_choice_random` | [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-002.md#canonical-0211210323021133-1013230231122300-3022123211212132-2220010210030202-2200303032320110-1130132332010002-2201212131212233-1220332022231120) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-002.md#canonical-3203002003212210-0003011020322120-0120231133003213-1123112300132001-0310031323210330-0303003333311213-0000030033003230-3321001022101020) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-002.md#canonical-3312322012133300-2111203231230023-2102011121323012-2220112313111130-0033333102130223-0233303123103301-2033223231222113-3030223000321003) |
| `id` | [ID](resources--udp_loadbalancer--reference--group-001.md#canonical-1103203003221010-2302130202212110-3101213302333013-2022223310223030-0332030023112203-3233232032222111-2030212313021023-1203232210313301) |
| `idle_timeout` | [idle_timeout](resources--udp_loadbalancer--reference--group-001.md#canonical-0003230310003122-0322010033032300-2122021310021103-3321212011333203-2322002230123233-2321000023131302-3031321320013010-0001303323100132) |
| `labels` | [labels](resources--udp_loadbalancer--reference--group-001.md#canonical-3130021300010131-2111113102303010-3103211221121321-1322130133021023-1202023211031231-1121202212031020-1023130023221133-1331211031220232) |
| `listen_port` | [listen_port](resources--udp_loadbalancer--reference--group-001.md#canonical-3120212203132202-0311303020101110-1201223210101333-3123231111113202-2203132021302012-1320322032022021-0212220011012001-2210010130102221) |
| `name` | [name](resources--udp_loadbalancer--reference--group-001.md#canonical-0331330200033333-3232002310120311-0210010030212312-3112231100111230-3100201302123333-3110022211030132-0022130101300002-2331220311033232) |
| `namespace` | [namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-1021220003100200-3032022000022230-2312302212033333-2133133223231033-1102312102012232-2200332211113023-3312210121013001-2031033322201323) |
| `no_service_policies` | [no_service_policies](resources--udp_loadbalancer--reference--group-002.md#canonical-2011121232032223-0120303113122000-0331200013101323-3211231010033023-3312301212123101-0023220313112111-0020212230101011-1311120023033020) |
| `origin_pools_weights` | [origin_pools_weights](resources--udp_loadbalancer--reference--group-002.md#canonical-2301103112020300-1300123003220110-0213202302132020-0033322330302211-1131212132021130-1333332213221110-3233130312003030-1303330011103122) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](resources--udp_loadbalancer--reference--group-002.md#canonical-2032330100001121-2112122010230011-3030302220110223-2133301002202010-3012323000030223-1331330020011100-0011023030332101-0222302000202112) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](resources--udp_loadbalancer--reference--group-002.md#canonical-2000333012001110-0231002120102101-3303003121210010-1332321113003102-2100332133202003-0313301213120010-1303331103120322-3120132233021000) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2011001300233313-1323212333101200-1110213322331003-2020030231101103-3103123231031200-3113011130033133-0313031103110310-0302031320303021) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-3122301133222333-2212111332333010-3001111010231010-0202203021011012-2120200200312310-3310203332230012-1033120323113233-1013101002012013) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](resources--udp_loadbalancer--reference--group-002.md#canonical-2132312123111111-3003130211230121-3300132131022213-3333113133210032-2201200312010223-0321000123030310-2023102101033332-3123023213221201) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](resources--udp_loadbalancer--reference--group-002.md#canonical-3212031321111032-1323010231033032-1030221012003033-3030032003312110-2132103312113033-0210320112321133-0311102213230112-2323003313233003) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](resources--udp_loadbalancer--reference--group-002.md#canonical-0102123012121013-1312311300310222-2201101212302230-2000012233220312-1012021020230223-3223202121020223-3311103200021203-1123222000232000) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2213010011232331-1032132002333120-0312211201103020-2221010323111100-0310112200102201-2123202201310021-1001113131020112-3212011213032030) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-3210122332303233-0022311202213122-0313012133221200-0120133023020231-1130021230031130-0323301233303120-2111230221030302-0010311202132321) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](resources--udp_loadbalancer--reference--group-002.md#canonical-2220102333322032-2201223002131230-0030212330103001-0323023233331212-1020223222111002-3020232320200310-2333311320003300-0301311121302220) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](resources--udp_loadbalancer--reference--group-002.md#canonical-0203333331223002-0132110010031013-0332102123022022-3213032211323331-0212313302031303-0002111303212011-1331222233130002-1101000223030023) |
| `port_ranges` | [port_ranges](resources--udp_loadbalancer--reference--group-001.md#canonical-2203020003111122-2333012023312122-1231102021223222-3213310221222000-3131120111002221-1213123101003133-3330021213323321-0022112233110231) |
| `service_policies_from_namespace` | [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2233110233201220-2203330010011033-2213301013122032-1020103020330201-1321121230130132-1133300233332010-3021001121303331-1211000301132011) |
| `timeouts` | [timeouts](resources--udp_loadbalancer--reference--group-002.md#canonical-2021022131222223-2333212001001212-2130332200003131-0020302000212332-1330323131131320-0311001103200222-0222203302031213-0330330211100201) |
| `timeouts.create` | [timeouts.create](resources--udp_loadbalancer--reference--group-002.md#canonical-3000102312123121-2222013031010322-3233121032202233-1110011132322112-3213302012103102-3230033012213102-1332113000131022-3112200030011201) |
| `timeouts.delete` | [timeouts.delete](resources--udp_loadbalancer--reference--group-002.md#canonical-2222311100010023-0101330313131303-1220221021112101-3222320033333210-0323133022321001-2221223332210312-3302101123231000-3320321103010213) |
| `timeouts.read` | [timeouts.read](resources--udp_loadbalancer--reference--group-002.md#canonical-2200002032133122-1131032330110203-0012313220212113-1301020202121323-2232030303201001-0210110202200033-2232003300231203-3221323230010321) |
| `timeouts.update` | [timeouts.update](resources--udp_loadbalancer--reference--group-002.md#canonical-1301133113131132-0323331102033003-3132210323110023-0303000122122233-3221031302110132-2210132033112112-2220301102231110-3310221213103310) |
| `udp` | [udp](resources--udp_loadbalancer--reference--group-002.md#canonical-1213213303011020-2212100322320001-3012203133023012-0202200110022330-1033233132020222-1011021100233133-2203233223110313-3323232000021300) |

<a id="canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- active_service_policies

<a id="canonical-1113100321101330-2330122113122300-3330103021013021-0221113231210202-1011123032032021-2100012033131203-3023303231210230-1020000132320220"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Additional upstream details:

List of service policies.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
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

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-1113100321101330-2330122113122300-3330103021013021-0221113231210202-1011123032032021-2100012033131203-3023303231210230-1020000132320220)
- [no_service_policies](resources--udp_loadbalancer--reference--group-002.md#canonical-2011121232032223-0120303113122000-0331200013101323-3211231010033023-3312301212123101-0023220313112111-0020212230101011-1311120023033020)
- [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2233110233201220-2203330010011033-2213301013122032-1020103020330201-1321121230130132-1133300233332010-3021001121303331-1211000301132011)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023223101230331-2331120022233023-0112312100113321-0220221130331122-2011133222020102-3021031003033030-0131022033112232-3001010111120130"></a>

### Direct properties for `active_service_policies`

- [policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2213011021021030-1220320232032031-2120302303103211-1011120311301121-3130302103333223-3311100203133021-1112230123002130-2200312220320200): complete subsection reference.

<a id="canonical-2213011021021030-1220320232032031-2120302303103211-1011120311301121-3130302103333223-3311100203133021-1112230123002130-2200312220320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies.policies` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132)
- active_service_policies.policies

<a id="canonical-3021320330233221-0222013013000013-1220031310010012-3202203200110213-2130331211011012-1301022232313103-1213113323323203-0111133101132032"></a>

Type: `"object"`. list nested block, Optional.

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its
characteristics are evaluated based on the match criteria in each service policy starting at the
top. If there is a match in the current policy, then the policy takes effect, and no more policies
are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none
match, then the request will be denied by default.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303202303302332-3321121202202222-3321110103333022-3120210110130003-1031322010122221-3123022102200212-0323302122120212-1130032001301133"></a>

### Direct properties for `active_service_policies.policies`

<a id="canonical-1311033330032101-2300331021132212-1221332122033331-3033113123012233-0213102312003320-0300113213333022-2310311200211122-1010013333322312"></a>

#### `active_service_policies.policies.name` property

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

<a id="canonical-3000133231300130-1020132333212023-0223022202330320-1012330112200330-2311210102112301-2003123111313310-3300110222212203-2101030330100301"></a>

<a id="canonical-3010311001200002-1102201330101331-1322302103013103-1202132103312122-1302112232203023-2032303210232022-2221003312330031-3310022332032112"></a>

#### `active_service_policies.policies.namespace` property

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

<a id="canonical-0103332210210101-1301013220210033-0031312333333130-3313302111102123-0031033001211113-2101223030220220-2221313302112032-0031323100321112"></a>

<a id="canonical-2201222022121101-0200230202003221-1331010211002202-0323230313333312-3321210223232022-2301221313113223-3121330101322131-0010320002330223"></a>

#### `active_service_policies.policies.tenant` property

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

<a id="canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- advertise_custom

<a id="canonical-2220112212031221-3223210202313233-2213100202202330-2021221131032103-0013010323313323-0221003321333303-0031200122113110-3111031023222021"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

Additional upstream details:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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

- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-2220112212031221-3223210202313233-2213100202202330-2021221131032103-0013010323313323-0221003321333303-0031200122113110-3111031023222021)
- [advertise_on_public](resources--udp_loadbalancer--reference--group-002.md#canonical-1131331003210000-2331113000100013-2232131210301213-1021112230322000-1222213303123001-0302322101033001-0010131203112121-2123220013211113)
- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-002.md#canonical-2302323123201032-2001211223302111-0132220033313002-3021301203111312-0321201310212003-0323332122132312-1131330222320133-2310133320001233)
- [do_not_advertise](resources--udp_loadbalancer--reference--group-002.md#canonical-0221213122111200-3301201210111322-3031222200121111-3003321003021322-3220301223013001-0210223020203022-0323301031102012-3132332221032200)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033301203032213-3323012202303021-3103031132332131-2010202200023313-2130223102103311-3210010122023121-1332321213211202-2213301022030031"></a>

### Direct properties for `advertise_custom`

- [advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101): complete subsection reference.

<a id="canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- advertise_custom.advertise_where

<a id="canonical-0010210202120001-2110301320213033-2233123203021222-3120002123302203-2233020312021012-3331301220332222-2033031003301301-2011210122212033"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-1120313030331003-2102022000332302-1303030301030013-1023121222110302-2012111013222323-3001210103310121-1033332220121200-2103121202120222"></a>

### Direct properties for `advertise_custom.advertise_where`

- [advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133): complete subsection reference.

- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001): complete subsection reference.

- [advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133): complete subsection reference.

<a id="canonical-0200332010112132-2310201102022000-0022022212003112-1311032020020020-3200323200013102-2303303310002221-1231122031223200-0232203030001012"></a>

<a id="canonical-2312202310332330-3011111311010003-3020133131010201-1331300103221110-3003332130331130-3322003222000310-0122302202002230-3120212220101321"></a>

#### `advertise_custom.advertise_where.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2123113113323331-3013303033210321-3121021112202033-0013320131230020-0113003123233132-2303020310110112-3033330210213221-1201222112233122"></a>

<a id="canonical-1032231023313112-3133223001113012-0030000230231030-3121331301330303-3130020032211230-1303111213023101-2112331001210302-3030101312202222"></a>

#### `advertise_custom.advertise_where.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

- [site](resources--udp_loadbalancer--reference--group-001.md#canonical-2131003131020022-1211231020110320-2321012131130323-3220001001103303-1322133131101301-3302332133130310-2330303311302133-2032012331123003): complete subsection reference.

- [use_default_port](resources--udp_loadbalancer--reference--group-001.md#canonical-3223033300212032-1333311221011013-0020020020322222-0013201132113031-2012200201021121-1202212313033222-0102233110002223-3012210133233132): complete subsection reference.

- [virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010): complete subsection reference.

- [virtual_site](resources--udp_loadbalancer--reference--group-002.md#canonical-3131233030200102-3231332321203303-2131331231102231-3123031211331102-0330311312320223-0211131322113002-1301003222212320-0133033112320330): complete subsection reference.

- [virtual_site_with_vip](resources--udp_loadbalancer--reference--group-002.md#canonical-1330330132111001-0030320120001011-3212231113031321-3110223330010000-0031121013222333-3332120333023202-1002100000320323-2202310313121302): complete subsection reference.

- [vk8s_service](resources--udp_loadbalancer--reference--group-002.md#canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222): complete subsection reference.

<a id="canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2330101022221233-0323020331311102-2313220303200200-3021103320222020-3123210321020111-1021223030233121-3322011003300012-1133311121230300"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311233012010323-2010133012123130-2233321133331023-1221303302323220-0213232331002121-0200220313330310-3210300203303022-1113311113301200"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-1312023232012223-0202221301203331-1012011230322011-3232301012300000-2132102010203022-1213320031111201-1201203223231202-3230122211012111): complete subsection reference.

<a id="canonical-1312023232012223-0202221301203331-1012011230322011-3232301012300000-2132102010203022-1213320031111201-1201203223231202-3230122211012111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-3112213021302113-0002230213003012-3230112030022130-1212313003313221-1031112012101320-0213332100130330-0131011111320012-2321030021102121"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233122223231213-0321311203030021-1311021212022102-1331033301101300-3020010023132320-3202130312100212-2222013001223201-1001102023002223"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-1013221200302123-3030233133311111-1232133032312322-2212301023012220-2212011220230032-1103123110320213-3111302100103002-2000201123321000"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` property

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

<a id="canonical-3223330100233103-2301222122312032-0313023120233121-1021031220220210-0212300101003212-2220223121333212-3232210323223302-1201121223130023"></a>

<a id="canonical-2200120220113110-0220330323200133-0011201230123330-3103031221020022-0220121301022123-3200131302132333-2131032012313333-3300011110003011"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` property

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

<a id="canonical-0230110232112000-1333301113230220-1331021111030100-0202310202033301-3131221301330310-3330313010231002-2231031000033201-3210223201233210"></a>

<a id="canonical-2011210101013130-0231002131213102-0110033032302331-0033233011303033-0223313232332103-0110311112000001-1103300112131311-0222112313011101"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` property

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

<a id="canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-2131202302333100-1213210221132131-0321032022213330-1110323133012313-0113122100221113-2233133223303322-1331323310230113-3221133112232231"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000000221121132-0002213320321103-2303102203013002-3123300112002212-3332203133111300-2100113213102212-3210123300323233-3133022101013332"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public`

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-0201223131000012-3300202233022012-0322103310120021-0202032113220200-3021332021210211-0213322133312201-0113321212133102-0333110130201003): complete subsection reference.

<a id="canonical-0201223131000012-3300202233022012-0322103310120021-0202032113220200-3021332021210211-0213322133312201-0113321212133102-0333110130201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-3101323200112133-2321023213200231-0322110200323002-0211032333233331-3012303123012110-1231212020302311-2022010203132321-0021132012333031"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203120201120121-3212023301331030-0302233102313332-3112100000233102-0303221321202121-2301110000320231-3331320310300323-3232231222200022"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-2213231003223302-2100230002022120-1333311102132102-2022300331031303-2122231221131203-2233013123212212-1111330102223222-1113230230103110"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.name` property

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

<a id="canonical-1011023001230302-0021013003233311-2032203330113113-0023210213330332-3233312001231031-0013030320113031-0130202002211222-0213331301232303"></a>

<a id="canonical-2111001001120212-0023122212131032-3323020023322023-3000001313102300-3231200212213003-3303022010033231-1220300113303300-3022223123120313"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` property

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

<a id="canonical-1301101221213203-1010023330200323-2103002212020023-0022021301230103-2110213021210003-3132023022213112-1020132233012023-3112333322313320"></a>

<a id="canonical-3110323133200221-0123101322210031-0023101002322023-1223122321201033-3333131301212133-3211113131211312-0312132130102112-2111110102010210"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` property

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

<a id="canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-0122331100201003-3101012133111122-2130312131030232-0223303230202230-2032311210201313-3323120312311321-3033331311313212-2112113232231213"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200213310212100-1120202010320312-3300223220300123-1113310202323312-3313022032321203-1220310311033012-3022130312011210-2201103101202022"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-2313120121320123-2320120301320220-0203210212333122-1033222213011201-3210031331320312-0100103220121213-3030231322233032-0232201332100300): complete subsection reference.

<a id="canonical-2313120121320123-2320120301320220-0203210212333122-1033222213011201-3210031331320312-0100103220121213-3030231322233032-0232201332100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-3301033300023313-1333000103210230-2213202102221310-3113332202123021-2311102121303313-0210331011232021-3130022113303110-2233132102332030"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221132233212302-3122101210203222-0311011302132331-1123031322002203-2202323130122211-0310210232131320-3322220333310222-1330332121322202"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-3023200020300232-3111321122001302-3310220203232113-2333203321133303-3012101233201231-3011012313101221-3120332212230212-3210221220023330"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` property

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

<a id="canonical-3210002303000122-0331030031202022-2012000001311302-2021120133312300-0302322300300033-1011001133213221-1300230122123013-0123010121001202"></a>

<a id="canonical-1111233320322210-0201020221020110-0023113121103001-1010001331203233-2322100002201220-2113130120031111-0230320000302033-3010000132121112"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` property

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

<a id="canonical-3323232012230010-2310233302301203-2001201330110301-0302111302032202-3111102212013123-1323123210331121-1221233312133110-2232321102212320"></a>

<a id="canonical-0123220100030310-0000222210013100-2033103213003200-1123323331230020-1212301133221322-2332130312002320-0002100113323103-0202102233120221"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` property

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

<a id="canonical-2131003131020022-1211231020110320-2321012131130323-3220001001103303-1322133131101301-3302332133130310-2330303311302133-2032012331123003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.site

<a id="canonical-0020132103030020-0230322133213331-2320031213330301-3000000312230113-0121103223120102-3000211210122130-0130323030031101-0022321020330212"></a>

Type: `"object"`. single nested block, Optional.

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220331032030101-0003110112303122-1002301020330320-0211121201213103-3032333313331200-0000330131031202-1111030023010113-0132001330331101"></a>

### Direct properties for `advertise_custom.advertise_where.site`

<a id="canonical-0001312212020110-1202310032031201-2011002230102320-0122323000132003-1200321300223123-3021133102111212-1132210311120121-2222023003230103"></a>

#### `advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2300030131202232-3233002021113223-3123302322322002-0232023030102011-2311013133033320-1230003022330101-0020330003111112-1103320202221000"></a>

<a id="canonical-0211312023312020-3210320313323020-0322323030330313-0201100303011302-0012332000032120-0132311310013311-0333013130322203-3212100002102212"></a>

#### `advertise_custom.advertise_where.site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--udp_loadbalancer--reference--group-001.md#canonical-2212221312310310-0101203011023033-3332000212220201-1010201313333102-3303302203233001-3212101210302111-0101012220111313-2002021333112002): complete subsection reference.

<a id="canonical-2212221312310310-0101203011023033-3332000212220201-1010201313333102-3303302203233001-3212101210302111-0101012220111313-2002021333112002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-2131003131020022-1211231020110320-2321012131130323-3220001001103303-1322133131101301-3302332133130310-2330303311302133-2032012331123003)
- advertise_custom.advertise_where.site.site

<a id="canonical-1130310233223332-1112003313113203-0011222131122021-3012201213102232-2213331303133100-0101231310323200-2322202203131110-3201213023003312"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320011333303000-1230013303311031-3112121011133331-3033321322100330-1303320122132013-2211122232013223-1230022123111002-3123122131311331"></a>

### Direct properties for `advertise_custom.advertise_where.site.site`

<a id="canonical-2131012300120111-3031331311202323-3101123013212110-2100330002021100-3020332130203323-1203000020302211-2220130111302111-1033133001300002"></a>

#### `advertise_custom.advertise_where.site.site.name` property

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

<a id="canonical-2030303110222032-2322321213102331-3111300310132311-2333322200322123-0333333132131201-0002111112023101-0112321031220103-0321103220011001"></a>

<a id="canonical-0002323011320232-0331232003002231-3231303313020123-0101022032102030-3102212311120132-2302311123102300-3230333010321201-1331330233111123"></a>

#### `advertise_custom.advertise_where.site.site.namespace` property

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

<a id="canonical-0320110210123322-3330211031223002-0330133221111100-3120312111231020-3320213101120110-0322221210200110-1011221122031312-0302222003310112"></a>

<a id="canonical-3000011103210213-3112010022231111-0313300212132032-3213000220301311-3222212132130112-2330303301011122-2133303312322232-3032010331301333"></a>

#### `advertise_custom.advertise_where.site.site.tenant` property

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

<a id="canonical-3223033300212032-1333311221011013-0020020020322222-0013201132113031-2012200201021121-1202212313033222-0102233110002223-3012210133233132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-0300323200230300-3103131322311021-3330001032031221-2111311300013113-0320323221110301-3120221322122312-2033101120111012-0121201003332103"></a>

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
use_default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-2301211302121211-0101311102033333-3313131212300320-0230213313103000-2301120220321213-0123302301220132-2220203212221331-3220000231320132"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_v6_vip",
    "specific_v6_vip"),
  validators.ConflictingObjectAttributes("default_vip",
    "specific_vip")}
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
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201002303212112-1333020000003022-1330333220212113-3201213013333330-3313233133123030-1033313111202012-2103312220123122-2301321210331223"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1021103000233223-3303331120323313-1323112223133221-3013231233012103-2200022310102310-3031023330123330-0331131331330003-1032012323011201): complete subsection reference.

- [default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-0202002320231111-0312020002301110-1120003023212122-1010130213101001-0312001203022022-1300300212200312-0132302030131120-1001003123023023): complete subsection reference.

<a id="canonical-0133102222123121-1230021120213020-0020222033320020-1330202022022011-1100211103310102-2220231120000323-0033233013220311-3322001302121101"></a>

<a id="canonical-0010021130223302-3010112012200213-2220231130302110-2222123201030011-0300312012103321-0232300032131203-1123022100111003-2003330020312322"></a>

#### `advertise_custom.advertise_where.virtual_network.specific_v6_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0332202231203330-0303200210131011-0322211201002321-0311203101301201-0020223011210120-0220031132330133-2131221121002220-0120032013232303"></a>

<a id="canonical-1222332213203202-2122231132212022-2011211320103011-2233203023333101-3123031330321320-2001320022311302-2331110310322000-0032113322111213"></a>

#### `advertise_custom.advertise_where.virtual_network.specific_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--udp_loadbalancer--reference--group-002.md#canonical-1301113323003210-1312131333020111-0112323001330313-0221011000012112-3112232300320223-3302332320201332-3210311213230221-0120230133031303): complete subsection reference.

<a id="canonical-1021103000233223-3303331120323313-1323112223133221-3013231233012103-2200022310102310-3031023330123330-0331131331330003-1032012323011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-2030211222310200-1113310031020312-2232111130022321-0102323110030321-0310220132130221-0321231011210201-3123110103012031-0322222201011031"></a>

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
default_v6_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202002320231111-0312020002301110-1120003023212122-1010130213101001-0312001203022022-1300300212200312-0132302030131120-1001003123023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-0223023012130320-3103013302333322-0230011111233230-1200022210103231-1011121323312221-1133211202232202-1202130131031222-1321301023213123"></a>

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
default_vip = {}
```

This is an empty object or choice marker. It has no direct properties.
