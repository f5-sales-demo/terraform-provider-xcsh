---
page_title: "xcsh_udp_loadbalancer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer reference."
---

# xcsh_udp_loadbalancer reference

<a id="canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102011000323030-1231123021231123-1101302021212220-3133203132301222-3012023212201021-0300120322320013-1121211130202000-2323310302130303"></a>

## Property reference — Property reference / 011113103023 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- Property reference

<a id="canonical-0302210303111330-1121023232013102-2023231020132233-1231001022032223-1222313223212211-3322210322233332-0333313220303022-3232231311201210"></a>

## Direct properties — Property reference / 011113103023 / 3

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132): complete subsection reference.

- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030): complete subsection reference.

- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3121001210223201-0002101322213023-2213011332023133-0223131303332232-0202230121130203-0330221122033201-1300103332030211-1301020220301201): complete subsection reference.

- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1032133030103213-0020211103223323-3100222332020210-1003331202302203-3333101310331003-1200101122201012-2332103110302022-2333031032112103): complete subsection reference.

<a id="canonical-3222130132230203-1330303011120303-1232200123100032-0132132113101002-3132000322223312-2312302123232131-2112321120213203-0032222221211012"></a>

<a id="canonical-0030312303002002-3113001223033213-0301212323330221-3330300001303231-1120220300120303-3012200321220110-3301033303002103-0312303131133020"></a>

## annotations property — Property reference / 011113103023 / 4

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

<a id="canonical-1202202213001333-2203210213312132-0110200021010021-0110031030210000-1001230303100221-2203221131013221-1033133333200233-0112003233322001"></a>

## description property — Property reference / 011113103023 / 5

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

<a id="canonical-3220320121201312-2313103212111113-0133023103321002-0311011113012100-0211101210312111-0312333220323122-1313213300232332-0130322012100202"></a>

<a id="canonical-3233221003013231-2122200203031120-1223213132022001-0031211230123211-0232221120232303-3120313013313122-2221120033012122-1213101020323302"></a>

## disable property — Property reference / 011113103023 / 6

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

<a id="canonical-1302310332320200-1300100212123002-2333002132121001-2201231200302223-3112330310200230-1003003201100221-2012212203012231-2213333200312200"></a>

<a id="canonical-1031023231300010-2122122322320102-1203300320202200-3231103212221131-1133130122210200-1002010313110102-0012001313003100-0200000323333323"></a>

## dns_volterra_managed property — Property reference / 011113103023 / 7

Type: `"bool"`. Optional, Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain to be delegated to F5 Distributed Cloud using the Delegated Domain feature or a DNS CNAME
record must be created in your DNS provider's portal.

Upstream description:

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

- [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-2021012000001001-1013233020022212-1310012310002230-0213221320021221-3313312011223103-0233332113202000-2001321301000221-1032223300220320): complete subsection reference.

<a id="canonical-1013201121013122-0113023330111203-3311020120112000-0113231223202132-1203222130212210-3031320332122302-3202112231022012-2223220301123230"></a>

<a id="canonical-1230200121233312-2201012230133122-3300111202230120-2212331233003200-1111002211303333-1330233310301132-3023300210303200-3010122330030223"></a>

## domains property — Property reference / 011113103023 / 8

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to this load balancer.

Upstream description:

A list of domains (host/authority header) that will be matched to this load balancer.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-2130020320231110-2310332213003130-2103302031122120-2101333103230022-2103002230301132-0101132133321211-1003113021330223-1110002031002313): complete subsection reference.

- [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-0033000022201003-1303032013200210-1033221223122313-1001020223133221-0221333001113000-1302131100323201-0111021013101131-1302332010012121): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-1232202330301332-2233232122032300-3311123230120100-0221111330223031-0212333133112002-2320202110313302-2212102131223121-2130011100002212): complete subsection reference.

<a id="canonical-1103203003221010-2302130202212110-3101213302333013-2022223310223030-0332030023112203-3233232032222111-2030212313021023-1203232210313301"></a>

<a id="canonical-3101313031200323-0213102030012223-0221331330102032-3111300310211030-1233200301301301-0302230123032031-1200232312323013-2321212001002010"></a>

## ID property — Property reference / 011113103023 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0003230310003122-0322010033032300-2122021310021103-3321212011333203-2322002230123233-2321000023131302-3031321320013010-0001303323100132"></a>

<a id="canonical-2320023022202131-2313202231103232-3332031121021303-3222323303112012-0211213200103222-1120013013103303-1132331203003020-2310123210212012"></a>

## idle_timeout property — Property reference / 011113103023 / 10

Type: `"number"`. Optional, Computed.

The amount of time that a session can exist without upstream or downstream activity, in
milliseconds.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-3130021300010131-2111113102303010-3103211221121321-1322130133021023-1202023211031231-1121202212031020-1023130023221133-1331211031220232"></a>

<a id="canonical-1320032223020101-2300200121012321-0021232011311222-2010202132001300-1221330133112022-2003020210333131-2031203113202002-2021203300100300"></a>

## labels property — Property reference / 011113103023 / 11

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

<a id="canonical-3120212203132202-0311303020101110-1201223210101333-3123231111113202-2203132021302012-1320322032022021-0212220011012001-2210010130102221"></a>

<a id="canonical-3122330213001221-3212333020302302-2133132020130023-1310113121021113-1301000111120012-1032130301100133-2323011210311011-3230031333323101"></a>

## listen_port property — Property reference / 011113103023 / 12

Type: `"number"`. Optional, Computed.

\[OneOf: listen\_port, port\_ranges\] Exclusive with \[port\_ranges\] Listen Port for this load
balancer.

Upstream description:

Exclusive with \[port\_ranges\] Listen Port for this load balancer.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0133322131233301-2232100103132233-1012031100320221-0002112012032200-1131131130003112-1012223300000003-2003230223102131-2000000020230231"></a>

## name property — Property reference / 011113103023 / 13

Type: `"string"`. Required.

Name of the UDP Load Balancer. Must be unique within the namespace.

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

<a id="canonical-1021220003100200-3032022000022230-2312302212033333-2133133223231033-1102312102012232-2200332211113023-3312210121013001-2031033322201323"></a>

<a id="canonical-2231210113203210-2032320210033330-0231001123023233-1100111132323012-1131301203001020-3302131202203222-2133120113312020-2300110123230021"></a>

## namespace property — Property reference / 011113103023 / 14

Type: `"string"`. Required.

Namespace where the UDP Load Balancer is created.

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

- [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2222312133022213-1230023330021312-1132323131113230-2333332022331122-2130002003301133-2032133133210033-3100313020020301-0031233232230012): complete subsection reference.

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303): complete subsection reference.

<a id="canonical-2203020003111122-2333012023312122-1231102021223222-3213310221222000-3131120111002221-1213123101003133-3330021213323321-0022112233110231"></a>

<a id="canonical-2220220322302021-2322011031202212-0201313030003010-1101021301111330-0130102231112220-3303010113311222-1231001302103323-2000033201112312"></a>

## port_ranges property — Property reference / 011113103023 / 15

Type: `"string"`. Optional, Computed.

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0312323201131002-2013002230101231-1032213002112002-0120122002213003-3123101223122022-0122231111300332-2023011220102220-0310312100132122"></a>

## All schema paths — Property reference / 011113103023 / 16

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
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2122203103003233-3201300211121110-0330330111103322-2020213320323320-2001322032221110-0122122110332010-2000030321123323-0320231231213013) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--udp_loadbalancer--reference--group-001.md#canonical-2120122030031333-3330321132323310-0132313111232022-1200012201303120-2333013033333232-0011231111321322-1021300303023113-3210003301220121) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-2012000312332231-2123200111311000-2331003210330213-2331222200032033-3323112013311120-1330320131220113-2130023231303230-0123003303010232) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-2222310211002121-2331223021033120-2031312211331020-0211313303231102-2011003101112213-0213121121303121-1210223121201222-1330212203323013) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3232020011102110-0020333100132333-3232011113311030-2220013311320103-1132303133321301-2303120031101100-3300222233102323-3021210321123031) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](resources--udp_loadbalancer--reference--group-001.md#canonical-1220122213020333-2020200311011000-2303003102322013-1032213333102123-3121011110030111-0121220310232313-2000031331103010-0322203313233003) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-1331120000120221-0330321131302130-0310332321121112-3023220031002021-3013312012000202-0123121010321133-3220323121131213-0210112033103311) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-0312332310312203-1321122130103023-1031332310130033-2203311203321201-3011203323321102-0110113103222011-1312131122031320-2100101133222033) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-2333023110112131-0210000303033021-2231021221110201-2001201011202133-2312020211020131-2121123103231233-1220012320332320-3102111231221221) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-2333112210200122-3303011203301130-1103212013012133-1211303322102000-1003033012303212-0110323112110021-1312031000232010-2030011001000230) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1303013321030132-3232223121202231-0012113213203010-2200001302323001-3322321331021132-2221321210121033-3300221020210022-0013213110232210) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--udp_loadbalancer--reference--group-001.md#canonical-1330121220100002-1220123310312003-1320112230101013-3220220210022110-1120220033121312-2103213310021002-1231110233103002-1001103131202132) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](resources--udp_loadbalancer--reference--group-001.md#canonical-0111111112300222-1132203011231122-2001013120003031-3011221012210011-2032023123302132-0310233223313333-1323111123211012-0120010311000232) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-1203330233032011-0221303100302323-1013131233023130-1220131103001322-1001200032322031-0122001232212021-3121103223323300-1023332312211132) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-3203012213202103-3330030321011311-2202322103300113-1012333301001201-3021301122312203-2100022023202313-3002233211000032-3003101213003021) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-1011001132113301-3333010013000033-2130121131300223-2330133132330223-3131200332032211-0131310022321003-2310022131323110-2020113013223311) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-0220312123212233-1232202230112321-2303231212010100-3333030103101131-1033113113002300-3033222113103012-3012230313320120-0211032021203111) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-0132203110033033-1223031113332002-0020213220310322-0221232200022322-2212022032321002-1110322202210223-0102202233220312-1011233031230323) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](resources--udp_loadbalancer--reference--group-001.md#canonical-1110001011323311-2302202330233110-0030121101122132-2032230213311300-1030010322332013-0020003130312311-1321100213210311-3003330122302000) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-2303201300332222-0012023001221030-2113203321232321-3231230231212011-0122203213302032-1311322200023221-3310122030312213-3020020223220212) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-0011321330122213-1202003011010120-0313333320213001-1302223223230013-1133100303301313-2022232002233302-2011100130210212-1001133300212123) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-1131120112221032-0110130201123313-1111321333200312-1010113211133120-0032203012331121-3003013221013123-1012201322313032-1002302303132010) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3301030232010212-1201333333221333-3310223011013103-2231101220012123-0200213221013011-3213332303133201-2211330111302313-2012013013112000) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-0202021221332203-1203212223132301-0203113113012313-3113212100232123-2232110111013300-3123110021110211-3323132231020121-1330321200011213) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-1113101012213132-0332023102322020-0210320110223211-1030001331123032-3121130222033210-1301222221210303-0211020331221111-3203030131321033) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-0131000012320331-1133122020002330-0233321133012322-1201203210313301-2210011010123100-1001112123211002-3323323230001331-0000122202203020) |
| `advertise_on_public` | [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1131331003210000-2331113000100013-2232131210301213-1021112230322000-1222213303123001-0302322101033001-0010131203112121-2123220013211113) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-0300103122321023-3000132322101211-2230222020221133-0012031321211210-2300030331103320-2201031002210020-2330301221230310-3212123202133031) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-0111110121032203-2332103123111313-1100013111131210-0013023230231022-3132112323122112-1332210211001321-0021321201312111-2023000211312332) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-2203221331301322-0133301002111213-0033001211222300-3320221322123331-3333103023023023-3012333302023113-3330113300030000-1321202220232130) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-3322133321113131-1003200203131002-3122112203331212-3023221301112203-0100232120221313-0033022213220131-1102210230322132-1121230313021322) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-2302323123201032-2001211223302111-0132220033313002-3021301203111312-0321201310212003-0323332122132312-1131330222320133-2310133320001233) |
| `annotations` | [annotations](resources--udp_loadbalancer--reference--group-001.md#canonical-3222130132230203-1330303011120303-1232200123100032-0132132113101002-3132000322223312-2312302123232131-2112321120213203-0032222221211012) |
| `description` | [description](resources--udp_loadbalancer--reference--group-001.md#canonical-0030130001211030-1031310321030131-1233200002003020-3313010122313311-3202203320102020-1201010112311003-1003012111330021-0030131003120220) |
| `disable` | [disable](resources--udp_loadbalancer--reference--group-001.md#canonical-3220320121201312-2313103212111113-0133023103321002-0311011113012100-0211101210312111-0312333220323122-1313213300232332-0130322012100202) |
| `dns_volterra_managed` | [dns_volterra_managed](resources--udp_loadbalancer--reference--group-001.md#canonical-1302310332320200-1300100212123002-2333002132121001-2201231200302223-3112330310200230-1003003201100221-2012212203012231-2213333200312200) |
| `do_not_advertise` | [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-0221213122111200-3301201210111322-3031222200121111-3003321003021322-3220301223013001-0210223020203022-0323301031102012-3132332221032200) |
| `domains` | [domains](resources--udp_loadbalancer--reference--group-001.md#canonical-1013201121013122-0113023330111203-3311020120112000-0113231223202132-1203222130212210-3031320332122302-3202112231022012-2223220301123230) |
| `hash_policy_choice_random` | [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-0211210323021133-1013230231122300-3022123211212132-2220010210030202-2200303032320110-1130132332010002-2201212131212233-1220332022231120) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-3203002003212210-0003011020322120-0120231133003213-1123112300132001-0310031323210330-0303003333311213-0000030033003230-3321001022101020) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-3312322012133300-2111203231230023-2102011121323012-2220112313111130-0033333102130223-0233303123103301-2033223231222113-3030223000321003) |
| `id` | [ID](resources--udp_loadbalancer--reference--group-001.md#canonical-1103203003221010-2302130202212110-3101213302333013-2022223310223030-0332030023112203-3233232032222111-2030212313021023-1203232210313301) |
| `idle_timeout` | [idle_timeout](resources--udp_loadbalancer--reference--group-001.md#canonical-0003230310003122-0322010033032300-2122021310021103-3321212011333203-2322002230123233-2321000023131302-3031321320013010-0001303323100132) |
| `labels` | [labels](resources--udp_loadbalancer--reference--group-001.md#canonical-3130021300010131-2111113102303010-3103211221121321-1322130133021023-1202023211031231-1121202212031020-1023130023221133-1331211031220232) |
| `listen_port` | [listen_port](resources--udp_loadbalancer--reference--group-001.md#canonical-3120212203132202-0311303020101110-1201223210101333-3123231111113202-2203132021302012-1320322032022021-0212220011012001-2210010130102221) |
| `name` | [name](resources--udp_loadbalancer--reference--group-001.md#canonical-0331330200033333-3232002310120311-0210010030212312-3112231100111230-3100201302123333-3110022211030132-0022130101300002-2331220311033232) |
| `namespace` | [namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-1021220003100200-3032022000022230-2312302212033333-2133133223231033-1102312102012232-2200332211113023-3312210121013001-2031033322201323) |
| `no_service_policies` | [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2011121232032223-0120303113122000-0331200013101323-3211231010033023-3312301212123101-0023220313112111-0020212230101011-1311120023033020) |
| `origin_pools_weights` | [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-2301103112020300-1300123003220110-0213202302132020-0033322330302211-1131212132021130-1333332213221110-3233130312003030-1303330011103122) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](resources--udp_loadbalancer--reference--group-001.md#canonical-2032330100001121-2112122010230011-3030302220110223-2133301002202010-3012323000030223-1331330020011100-0011023030332101-0222302000202112) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](resources--udp_loadbalancer--reference--group-001.md#canonical-2000333012001110-0231002120102101-3303003121210010-1332321113003102-2100332133202003-0313301213120010-1303331103120322-3120132233021000) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-2011001300233313-1323212333101200-1110213322331003-2020030231101103-3103123231031200-3113011130033133-0313031103110310-0302031320303021) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-3122301133222333-2212111332333010-3001111010231010-0202203021011012-2120200200312310-3310203332230012-1033120323113233-1013101002012013) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](resources--udp_loadbalancer--reference--group-001.md#canonical-2132312123111111-3003130211230121-3300132131022213-3333113133210032-2201200312010223-0321000123030310-2023102101033332-3123023213221201) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](resources--udp_loadbalancer--reference--group-002.md#canonical-3212031321111032-1323010231033032-1030221012003033-3030032003312110-2132103312113033-0210320112321133-0311102213230112-2323003313233003) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](resources--udp_loadbalancer--reference--group-002.md#canonical-0102123012121013-1312311300310222-2201101212302230-2000012233220312-1012021020230223-3223202121020223-3311103200021203-1123222000232000) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2213010011232331-1032132002333120-0312211201103020-2221010323111100-0310112200102201-2123202201310021-1001113131020112-3212011213032030) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](resources--udp_loadbalancer--reference--group-002.md#canonical-3210122332303233-0022311202213122-0313012133221200-0120133023020231-1130021230031130-0323301233303120-2111230221030302-0010311202132321) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](resources--udp_loadbalancer--reference--group-001.md#canonical-2220102333322032-2201223002131230-0030212330103001-0323023233331212-1020223222111002-3020232320200310-2333311320003300-0301311121302220) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](resources--udp_loadbalancer--reference--group-001.md#canonical-0203333331223002-0132110010031013-0332102123022022-3213032211323331-0212313302031303-0002111303212011-1331222233130002-1101000223030023) |
| `port_ranges` | [port_ranges](resources--udp_loadbalancer--reference--group-001.md#canonical-2203020003111122-2333012023312122-1231102021223222-3213310221222000-3131120111002221-1213123101003133-3330021213323321-0022112233110231) |
| `service_policies_from_namespace` | [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2233110233201220-2203330010011033-2213301013122032-1020103020330201-1321121230130132-1133300233332010-3021001121303331-1211000301132011) |
| `timeouts` | [timeouts](resources--udp_loadbalancer--reference--group-002.md#canonical-2021022131222223-2333212001001212-2130332200003131-0020302000212332-1330323131131320-0311001103200222-0222203302031213-0330330211100201) |
| `timeouts.create` | [timeouts.create](resources--udp_loadbalancer--reference--group-002.md#canonical-3000102312123121-2222013031010322-3233121032202233-1110011132322112-3213302012103102-3230033012213102-1332113000131022-3112200030011201) |
| `timeouts.delete` | [timeouts.delete](resources--udp_loadbalancer--reference--group-002.md#canonical-2222311100010023-0101330313131303-1220221021112101-3222320033333210-0323133022321001-2221223332210312-3302101123231000-3320321103010213) |
| `timeouts.read` | [timeouts.read](resources--udp_loadbalancer--reference--group-002.md#canonical-2200002032133122-1131032330110203-0012313220212113-1301020202121323-2232030303201001-0210110202200033-2232003300231203-3221323230010321) |
| `timeouts.update` | [timeouts.update](resources--udp_loadbalancer--reference--group-002.md#canonical-1301133113131132-0323331102033003-3132210323110023-0303000122122233-3221031302110132-2210132033112112-2220301102231110-3310221213103310) |
| `udp` | [udp](resources--udp_loadbalancer--reference--group-002.md#canonical-1213213303011020-2212100322320001-3012203133023012-0202200110022330-1033233132020222-1011021100233133-2203233223110313-3323232000021300) |

<a id="canonical-1232033103322131-2303212330020223-2231131200102322-2333233112332031-0312013102022123-2302112323221133-1203202011023232-3220100202200121"></a>

## Next pages — Property reference / 011113103023 / 17

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3121001210223201-0002101322213023-2213011332023133-0223131303332232-0202230121130203-0330221122033201-1300103332030211-1301020220301201)
- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1032133030103213-0020211103223323-3100222332020210-1003331202302203-3333101310331003-1200101122201012-2332103110302022-2333031032112103)
- [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-2021012000001001-1013233020022212-1310012310002230-0213221320021221-3313312011223103-0233332113202000-2001321301000221-1032223300220320)
- [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-2130020320231110-2310332213003130-2103302031122120-2101333103230022-2103002230301132-0101132133321211-1003113021330223-1110002031002313)
- [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-0033000022201003-1303032013200210-1033221223122313-1001020223133221-0221333001113000-1302131100323201-0111021013101131-1302332010012121)
- [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-1232202330301332-2233232122032300-3311123230120100-0221111330223031-0212333133112002-2320202110313302-2212102131223121-2130011100002212)
- [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2222312133022213-1230023330021312-1132323131113230-2333332022331122-2130002003301133-2032133133210033-3100313020020301-0031233232230012)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303)
- [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2221100322323133-3222100232310313-2133010130310031-3231202023101032-3313112200110100-2130133221200013-2211012323333211-0003213013303230)
- [timeouts](resources--udp_loadbalancer--reference--group-002.md#canonical-1222013031123323-0132302211031313-1113201031020033-3330303023103321-1132201210330120-3022032110001231-2010201311132220-1032103313302111)
- [udp](resources--udp_loadbalancer--reference--group-002.md#canonical-3121113233300112-1011301020012201-3133332123031100-1313012230000123-2323010011202022-1122033233122001-3001313032003221-2033202233101030)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023223101230331-2331120022233023-0112312100113321-0220221130331122-2011133222020102-3021031003033030-0131022033112232-3001010111120130"></a>

## active_service_policies — active_service_policies / 321103231022 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- active_service_policies

<a id="canonical-1113100321101330-2330122113122300-3330103021013021-0221113231210202-1011123032032021-2100012033131203-3023303231210230-1020000132320220"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Upstream description:

List of service policies.

Provider validators and defaults (from schema source):

```go
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
- [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2011121232032223-0120303113122000-0331200013101323-3211231010033023-3312301212123101-0023220313112111-0020212230101011-1311120023033020)
- [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-002.md#canonical-2233110233201220-2203330010011033-2213301013122032-1020103020330201-1321121230130132-1133300233332010-3021001121303331-1211000301132011)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323122321001203-0221331210012300-0320231332131022-1203220331202323-0113032202220031-1022211020232211-0331330331233211-2321003101323102"></a>

## Direct properties — active_service_policies / 321103231022 / 3

- [policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2213011021021030-1220320232032031-2120302303103211-1011120311301121-3130302103333223-3311100203133021-1112230123002130-2200312220320200): complete subsection reference.

<a id="canonical-3113311131300011-2331120022300131-1331112033312213-0322113221112101-3211320232111010-1310313330002233-1013201013013010-1121022031323213"></a>

## Next pages — active_service_policies / 321103231022 / 4

- [active_service_policies.policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2213011021021030-1220320232032031-2120302303103211-1011120311301121-3130302103333223-3311100203133021-1112230123002130-2200312220320200)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2213011021021030-1220320232032031-2120302303103211-1011120311301121-3130302103333223-3311100203133021-1112230123002130-2200312220320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303202303302332-3321121202202222-3321110103333022-3120210110130003-1031322010122221-3123022102200212-0323302122120212-1130032001301133"></a>

## active_service_policies.policies — policies / 313231003103 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132)
- active_service_policies.policies

<a id="canonical-3021320330233221-0222013013000013-1220031310010012-3202203200110213-2130331211011012-1301022232313103-1213113323323203-0111133101132032"></a>

Type: `"object"`. list nested block, Optional.

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its..

Upstream description:

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its
characteristics are evaluated based on the match criteria in each service policy starting at the
top. If there is a match in the current policy, then the policy takes effect, and no more policies
are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none
match, then the request will be denied by default.

Provider validators and defaults (from schema source):

```go
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

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010311001200002-1102201330101331-1322302103013103-1202132103312122-1302112232203023-2032303210232022-2221003312330031-3310022332032112"></a>

## Direct properties — policies / 313231003103 / 3

<a id="canonical-1311033330032101-2300331021132212-1221332122033331-3033113123012233-0213102312003320-0300113213333022-2310311200211122-1010013333322312"></a>

<a id="canonical-2201222022121101-0200230202003221-1331010211002202-0323230313333312-3321210223232022-2301221313113223-3121330101322131-0010320002330223"></a>

## name property — policies / 313231003103 / 4

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

<a id="canonical-3000133231300130-1020132333212023-0223022202330320-1012330112200330-2311210102112301-2003123111313310-3300110222212203-2101030330100301"></a>

<a id="canonical-1303033313012213-0133331001220201-2123112003130202-3032202022023332-2231323201101231-2112033330320020-3033323330122012-2033133222001301"></a>

## namespace property — policies / 313231003103 / 5

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

<a id="canonical-0103332210210101-1301013220210033-0031312333333130-3313302111102123-0031033001211113-2101223030220220-2221313302112032-0031323100321112"></a>

<a id="canonical-2133301013000123-2103331311033132-3010003131200213-0032130312030201-1203131223111203-1102111210223013-3100200221112032-0211302010333310"></a>

## tenant property — policies / 313231003103 / 6

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

<a id="canonical-0033002332101222-0302022130230220-2130231021200133-0133002313200103-1110201230333032-3202120030221131-2220210113013032-0213013112232010"></a>

## Next pages — policies / 313231003103 / 7

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-0232231312113132-0011113013331010-3331133001111101-2303123300322311-0312001220202312-1122123223310111-2113233111100023-2020111032211132)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033301203032213-3323012202303021-3103031132332131-2010202200023313-2130223102103311-3210010122023121-1332321213211202-2213301022030031"></a>

## advertise_custom — advertise_custom / 131302333201 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- advertise_custom

<a id="canonical-2220112212031221-3223210202313233-2213100202202330-2021221131032103-0013010323313323-0221003321333303-0031200122113110-3111031023222021"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
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
- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1131331003210000-2331113000100013-2232131210301213-1021112230322000-1222213303123001-0302322101033001-0010131203112121-2123220013211113)
- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-2302323123201032-2001211223302111-0132220033313002-3021301203111312-0321201310212003-0323332122132312-1131330222320133-2310133320001233)
- [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-0221213122111200-3301201210111322-3031222200121111-3003321003021322-3220301223013001-0210223020203022-0323301031102012-3132332221032200)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131023031302023-0301310001032232-2332223231023200-1122031030013102-3220223221330031-3311313122013112-0032132123313233-2313211122202202"></a>

## Direct properties — advertise_custom / 131302333201 / 3

- [advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101): complete subsection reference.

<a id="canonical-0303323210023002-2033320030001121-3200123311213231-2231321030100232-3320302202010110-1310103310023321-1010332132121212-2032101002101132"></a>

## Next pages — advertise_custom / 131302333201 / 4

- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120313030331003-2102022000332302-1303030301030013-1023121222110302-2012111013222323-3001210103310121-1033332220121200-2103121202120222"></a>

## advertise_custom.advertise_where — advertise_where / 213232221110 / 2

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

<a id="canonical-2312202310332330-3011111311010003-3020133131010201-1331300103221110-3003332130331130-3322003222000310-0122302202002230-3120212220101321"></a>

## Direct properties — advertise_where / 213232221110 / 3

- [advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133): complete subsection reference.

- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001): complete subsection reference.

- [advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133): complete subsection reference.

<a id="canonical-0200332010112132-2310201102022000-0022022212003112-1311032020020020-3200323200013102-2303303310002221-1231122031223200-0232203030001012"></a>

<a id="canonical-1032231023313112-3133223001113012-0030000230231030-3121331301330303-3130020032211230-1303111213023101-2112331001210302-3030101312202222"></a>

## port property — advertise_where / 213232221110 / 4

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0023302110222331-3331323223021222-0003320200103310-2121032002310333-1300312302102312-1112201212033220-1130122002333121-3122223223033321"></a>

## port_ranges property — advertise_where / 213232221110 / 5

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
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

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3131233030200102-3231332321203303-2131331231102231-3123031211331102-0330311312320223-0211131322113002-1301003222212320-0133033112320330): complete subsection reference.

- [virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1330330132111001-0030320120001011-3212231113031321-3110223330010000-0031121013222333-3332120333023202-1002100000320323-2202310313121302): complete subsection reference.

- [vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222): complete subsection reference.

<a id="canonical-0030223332013320-3201200210220132-2210211230321033-2302203132110331-0021332223302012-2312222113203330-0323230323123020-0311101312102223"></a>

## Next pages — advertise_where / 213232221110 / 6

- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133)
- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133)
- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-2131003131020022-1211231020110320-2321012131130323-3220001001103303-1322133131101301-3302332133130310-2330303311302133-2032012331123003)
- [advertise_custom.advertise_where.use_default_port](resources--udp_loadbalancer--reference--group-001.md#canonical-3223033300212032-1333311221011013-0020020020322222-0013201132113031-2012200201021121-1202212313033222-0102233110002223-3012210133233132)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010)
- [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3131233030200102-3231332321203303-2131331231102231-3123031211331102-0330311312320223-0211131322113002-1301003222212320-0133033112320330)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1330330132111001-0030320120001011-3212231113031321-3110223330010000-0031121013222333-3332120333023202-1002100000320323-2202310313121302)
- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311233012010323-2010133012123130-2233321133331023-1221303302323220-0213232331002121-0200220313330310-3210300203303022-1113311113301200"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_dualstack_on_public / 201012120233 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2330101022221233-0323020331311102-2313220303200200-3021103320222020-3123210321020111-1021223030233121-3322011003300012-1133311121230300"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

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

<a id="canonical-1030232231120231-2312013302210212-3032203332311113-1220130021013200-2210112022333201-2113002231311321-1332001002030031-1030122033130031"></a>

## Direct properties — advertise_dualstack_on_public / 201012120233 / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-1312023232012223-0202221301203331-1012011230322011-3232301012300000-2132102010203022-1213320031111201-1201203223231202-3230122211012111): complete subsection reference.

<a id="canonical-2331220300312132-1200323311111220-0332022330003203-1010311220003312-0112112310312213-0200120311113232-1222133103231022-2313103031221322"></a>

## Next pages — advertise_dualstack_on_public / 201012120233 / 4

- [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-1312023232012223-0202221301203331-1012011230322011-3232301012300000-2132102010203022-1213320031111201-1201203223231202-3230122211012111)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1312023232012223-0202221301203331-1012011230322011-3232301012300000-2132102010203022-1213320031111201-1201203223231202-3230122211012111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233122223231213-0321311203030021-1311021212022102-1331033301101300-3020010023132320-3202130312100212-2222013001223201-1001102023002223"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — public_ip / 212132110122 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-3112213021302113-0002230213003012-3230112030022130-1212313003313221-1031112012101320-0213332100130330-0131011111320012-2321030021102121"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200120220113110-0220330323200133-0011201230123330-3103031221020022-0220121301022123-3200131302132333-2131032012313333-3300011110003011"></a>

## Direct properties — public_ip / 212132110122 / 3

<a id="canonical-1013221200302123-3030233133311111-1232133032312322-2212301023012220-2212011220230032-1103123110320213-3111302100103002-2000201123321000"></a>

<a id="canonical-2011210101013130-0231002131213102-0110033032302331-0033233011303033-0223313232332103-0110311112000001-1103300112131311-0222112313011101"></a>

## name property — public_ip / 212132110122 / 4

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

<a id="canonical-3223330100233103-2301222122312032-0313023120233121-1021031220220210-0212300101003212-2220223121333212-3232210323223302-1201121223130023"></a>

<a id="canonical-0013001302312210-1133022033303212-2010111330123201-3200330300320202-0331221133332121-0130112011332101-0010321031011130-0130323012003000"></a>

## namespace property — public_ip / 212132110122 / 5

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

<a id="canonical-0230110232112000-1333301113230220-1331021111030100-0202310202033301-3131221301330310-3330313010231002-2231031000033201-3210223201233210"></a>

<a id="canonical-0200230000012210-2310200031101033-3111111321033132-0200210023303330-0332230200131123-2213033221022031-0203002130021303-3203233021032113"></a>

## tenant property — public_ip / 212132110122 / 6

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

<a id="canonical-2023233101312201-1321333123032310-0021210221011322-3211020103123223-1232113323213213-0201000111222120-0023332100001211-2020133313302002"></a>

## Next pages — public_ip / 212132110122 / 7

- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3233113133000313-0130132231033022-2030101231111132-2322311213130313-3120003323102112-3122103303120012-1120221222012102-2102203032322133)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000000221121132-0002213320321103-2303102203013002-3123300112002212-3332203133111300-2100113213102212-3210123300323233-3133022101013332"></a>

## advertise_custom.advertise_where.advertise_on_public — advertise_on_public / 321201212321 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-2131202302333100-1213210221132131-0321032022213330-1110323133012313-0113122100221113-2233133223303322-1331323310230113-3221133112232231"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

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

<a id="canonical-0133301302202221-2301232322231121-2313222302303001-2211232130003130-1311100322312302-0132210311110102-3023103030312332-2013121131111002"></a>

## Direct properties — advertise_on_public / 321201212321 / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-0201223131000012-3300202233022012-0322103310120021-0202032113220200-3021332021210211-0213322133312201-0113321212133102-0333110130201003): complete subsection reference.

<a id="canonical-2313031223010331-0023231302123230-0311122002233011-0100231330210132-1323211022333202-2033133321131222-2101312221311302-2211001020100211"></a>

## Next pages — advertise_on_public / 321201212321 / 4

- [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-0201223131000012-3300202233022012-0322103310120021-0202032113220200-3021332021210211-0213322133312201-0113321212133102-0333110130201003)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0201223131000012-3300202233022012-0322103310120021-0202032113220200-3021332021210211-0213322133312201-0113321212133102-0333110130201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203120201120121-3212023301331030-0302233102313332-3112100000233102-0303221321202121-2301110000320231-3331320310300323-3232231222200022"></a>

## advertise_custom.advertise_where.advertise_on_public.public_ip — public_ip / 221333003233 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-3101323200112133-2321023213200231-0322110200323002-0211032333233331-3012303123012110-1231212020302311-2022010203132321-0021132012333031"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111001001120212-0023122212131032-3323020023322023-3000001313102300-3231200212213003-3303022010033231-1220300113303300-3022223123120313"></a>

## Direct properties — public_ip / 221333003233 / 3

<a id="canonical-2213231003223302-2100230002022120-1333311102132102-2022300331031303-2122231221131203-2233013123212212-1111330102223222-1113230230103110"></a>

<a id="canonical-3110323133200221-0123101322210031-0023101002322023-1223122321201033-3333131301212133-3211113131211312-0312132130102112-2111110102010210"></a>

## name property — public_ip / 221333003233 / 4

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

<a id="canonical-1011023001230302-0021013003233311-2032203330113113-0023210213330332-3233312001231031-0013030320113031-0130202002211222-0213331301232303"></a>

<a id="canonical-3232110220213230-1132220123133012-0312102202130030-3011211230213321-0313222002310203-3031232012100321-1202000231220213-0013322120122012"></a>

## namespace property — public_ip / 221333003233 / 5

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

<a id="canonical-1301101221213203-1010023330200323-2103002212020023-0022021301230103-2110213021210003-3132023022213112-1020132233012023-3112333322313320"></a>

<a id="canonical-3123130120122332-1131031222210000-3033200023132200-0132222222123320-0300321100322113-1033203013313220-3110003322312333-0312322133320303"></a>

## tenant property — public_ip / 221333003233 / 6

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

<a id="canonical-0121220221121132-0023303320130323-2231200213003033-1101203133322203-2213221220113012-3210202012210023-2200300002223303-3222122131330311"></a>

## Next pages — public_ip / 221333003233 / 7

- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1033032223211032-2001001133300200-2202011023102302-2303321301033211-0213023131223100-1303022221112302-3000310120201030-1002013210100001)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200213310212100-1120202010320312-3300223220300123-1113310202323312-3313022032321203-1220310311033012-3022130312011210-2201103101202022"></a>

## advertise_custom.advertise_where.advertise_v6_on_public — advertise_v6_on_public / 121110200212 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-0122331100201003-3101012133111122-2130312131030232-0223303230202230-2032311210201313-3323120312311321-3033331311313212-2112113232231213"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

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

<a id="canonical-2320323323323000-1231112201313021-1322322222030231-2231132320023113-1203120102222022-0003333112120203-3030013001022311-1030321231002223"></a>

## Direct properties — advertise_v6_on_public / 121110200212 / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-2313120121320123-2320120301320220-0203210212333122-1033222213011201-3210031331320312-0100103220121213-3030231322233032-0232201332100300): complete subsection reference.

<a id="canonical-1002020132013212-1113130022031123-1210113232011021-3233222230200022-3222001310203022-0331020230313130-0221130302002300-0332110111222132"></a>

## Next pages — advertise_v6_on_public / 121110200212 / 4

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-2313120121320123-2320120301320220-0203210212333122-1033222213011201-3210031331320312-0100103220121213-3030231322233032-0232201332100300)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2313120121320123-2320120301320220-0203210212333122-1033222213011201-3210031331320312-0100103220121213-3030231322233032-0232201332100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221132233212302-3122101210203222-0311011302132331-1123031322002203-2202323130122211-0310210232131320-3322220333310222-1330332121322202"></a>

## advertise_custom.advertise_where.advertise_v6_on_public.public_ip — public_ip / 101233000332 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-3301033300023313-1333000103210230-2213202102221310-3113332202123021-2311102121303313-0210331011232021-3130022113303110-2233132102332030"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111233320322210-0201020221020110-0023113121103001-1010001331203233-2322100002201220-2113130120031111-0230320000302033-3010000132121112"></a>

## Direct properties — public_ip / 101233000332 / 3

<a id="canonical-3023200020300232-3111321122001302-3310220203232113-2333203321133303-3012101233201231-3011012313101221-3120332212230212-3210221220023330"></a>

<a id="canonical-0123220100030310-0000222210013100-2033103213003200-1123323331230020-1212301133221322-2332130312002320-0002100113323103-0202102233120221"></a>

## name property — public_ip / 101233000332 / 4

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

<a id="canonical-3210002303000122-0331030031202022-2012000001311302-2021120133312300-0302322300300033-1011001133213221-1300230122123013-0123010121001202"></a>

<a id="canonical-3102233232312133-1331021013012020-2100321311301020-2303111201031102-3111331223023120-0102313100320332-0331331232230333-0113011103300023"></a>

## namespace property — public_ip / 101233000332 / 5

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

<a id="canonical-3323232012230010-2310233302301203-2001201330110301-0302111302032202-3111102212013123-1323123210331121-1221233312133110-2232321102212320"></a>

<a id="canonical-1012321020222210-2010120332313011-1102103123123201-3123210232120022-0201002110020030-1323213120111002-2311100303021020-2320322233000310"></a>

## tenant property — public_ip / 101233000332 / 6

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

<a id="canonical-1321221322223323-1101112021101030-2032233033320131-0303303103221003-1201201012131001-1023330221003103-0031032210003211-2310031323210010"></a>

## Next pages — public_ip / 101233000332 / 7

- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0033221013213223-1313322000011132-1003011010021101-0303322131321102-2100003021121101-0313013131211302-2210311231200021-0201310203021133)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2131003131020022-1211231020110320-2321012131130323-3220001001103303-1322133131101301-3302332133130310-2330303311302133-2032012331123003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220331032030101-0003110112303122-1002301020330320-0211121201213103-3032333313331200-0000330131031202-1111030023010113-0132001330331101"></a>

## advertise_custom.advertise_where.site — site / 200303101333 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.site

<a id="canonical-0020132103030020-0230322133213331-2320031213330301-3000000312230113-0121103223120102-3000211210122130-0130323030031101-0022321020330212"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

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

<a id="canonical-0211312023312020-3210320313323020-0322323030330313-0201100303011302-0012332000032120-0132311310013311-0333013130322203-3212100002102212"></a>

## Direct properties — site / 200303101333 / 3

<a id="canonical-0001312212020110-1202310032031201-2011002230102320-0122323000132003-1200321300223123-3021133102111212-1132210311120121-2222023003230103"></a>

<a id="canonical-1101320111133103-0332131220110332-1010130112331111-1212223001133030-2100100211302213-2122130301303331-3132202201310200-1301220220231231"></a>

## ip property — site / 200303101333 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2300030131202232-3233002021113223-3123302322322002-0232023030102011-2311013133033320-1230003022330101-0020330003111112-1103320202221000"></a>

<a id="canonical-1122322320333112-0323133300222313-0031120333333230-2113213131032201-1102230322131103-2010330333331020-0032220123330011-0310201103221132"></a>

## network property — site / 200303101333 / 5

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0303033321331333-3020101031302020-3322213331303223-3333212231002220-1101133313020212-3030332212120232-0221303301031220-0301333231311313"></a>

## Next pages — site / 200303101333 / 6

- [advertise_custom.advertise_where.site.site](resources--udp_loadbalancer--reference--group-001.md#canonical-2212221312310310-0101203011023033-3332000212220201-1010201313333102-3303302203233001-3212101210302111-0101012220111313-2002021333112002)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2212221312310310-0101203011023033-3332000212220201-1010201313333102-3303302203233001-3212101210302111-0101012220111313-2002021333112002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320011333303000-1230013303311031-3112121011133331-3033321322100330-1303320122132013-2211122232013223-1230022123111002-3123122131311331"></a>

## advertise_custom.advertise_where.site.site — site / 103313302112 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-2131003131020022-1211231020110320-2321012131130323-3220001001103303-1322133131101301-3302332133130310-2330303311302133-2032012331123003)
- advertise_custom.advertise_where.site.site

<a id="canonical-1130310233223332-1112003313113203-0011222131122021-3012201213102232-2213331303133100-0101231310323200-2322202203131110-3201213023003312"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0002323011320232-0331232003002231-3231303313020123-0101022032102030-3102212311120132-2302311123102300-3230333010321201-1331330233111123"></a>

## Direct properties — site / 103313302112 / 3

<a id="canonical-2131012300120111-3031331311202323-3101123013212110-2100330002021100-3020332130203323-1203000020302211-2220130111302111-1033133001300002"></a>

<a id="canonical-3000011103210213-3112010022231111-0313300212132032-3213000220301311-3222212132130112-2330303301011122-2133303312322232-3032010331301333"></a>

## name property — site / 103313302112 / 4

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

<a id="canonical-2030303110222032-2322321213102331-3111300310132311-2333322200322123-0333333132131201-0002111112023101-0112321031220103-0321103220011001"></a>

<a id="canonical-2212322121321130-0002023021000322-2121221031101012-2121131200213302-1201313011310021-0100321210320122-2011103020201313-2310220221310212"></a>

## namespace property — site / 103313302112 / 5

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

<a id="canonical-0320110210123322-3330211031223002-0330133221111100-3120312111231020-3320213101120110-0322221210200110-1011221122031312-0302222003310112"></a>

<a id="canonical-3321000121310000-0030312012031220-1230021202020222-3002023030231003-1021232033332311-0011023210100110-0211203222312003-2200323210113023"></a>

## tenant property — site / 103313302112 / 6

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

<a id="canonical-0220012103202133-2212203012221002-3121033333122302-0111130123011201-1230002121331231-3110010220023201-2122023201301123-2133323122023232"></a>

## Next pages — site / 103313302112 / 7

- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-2131003131020022-1211231020110320-2321012131130323-3220001001103303-1322133131101301-3302332133130310-2330303311302133-2032012331123003)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-3223033300212032-1333311221011013-0020020020322222-0013201132113031-2012200201021121-1202212313033222-0102233110002223-3012210133233132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012111211231103-1221233222211313-3113221022022321-1323312121203122-0121210333102210-0120132020012112-0110023021210102-3302011031110113"></a>

## advertise_custom.advertise_where.use_default_port — use_default_port / 112101012033 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-0300323200230300-3103131322311021-3330001032031221-2111311300013113-0320323221110301-3120221322122312-2033101120111012-0121201003332103"></a>

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
use_default_port = {}
```

<a id="canonical-2022120001020300-2110122102333233-3033220121131020-1331121330101220-2022012033312221-3210303000100130-0220231231033030-0010130001102011"></a>

## Direct properties — use_default_port / 112101012033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230331200232321-1232003331113220-0302120203130102-2100103230312311-3200022123030003-3133102000321200-0113130300103112-3031032230323302"></a>

## Next pages — use_default_port / 112101012033 / 4

- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201002303212112-1333020000003022-1330333220212113-3201213013333330-3313233133123030-1033313111202012-2103312220123122-2301321210331223"></a>

## advertise_custom.advertise_where.virtual_network — virtual_network / 012201032320 / 2

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

<a id="canonical-0010021130223302-3010112012200213-2220231130302110-2222123201030011-0300312012103321-0232300032131203-1123022100111003-2003330020312322"></a>

## Direct properties — virtual_network / 012201032320 / 3

- [default_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1021103000233223-3303331120323313-1323112223133221-3013231233012103-2200022310102310-3031023330123330-0331131331330003-1032012323011201): complete subsection reference.

- [default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-0202002320231111-0312020002301110-1120003023212122-1010130213101001-0312001203022022-1300300212200312-0132302030131120-1001003123023023): complete subsection reference.

<a id="canonical-0133102222123121-1230021120213020-0020222033320020-1330202022022011-1100211103310102-2220231120000323-0033233013220311-3322001302121101"></a>

<a id="canonical-1222332213203202-2122231132212022-2011211320103011-2233203023333101-3123031330321320-2001320022311302-2331110310322000-0032113322111213"></a>

## specific_v6_vip property — virtual_network / 012201032320 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0332202231203330-0303200210131011-0322211201002321-0311203101301201-0020223011210120-0220031132330133-2131221121002220-0120032013232303"></a>

<a id="canonical-0023330323210013-0331211001100121-1020020103233331-1023320133212223-3313220131232220-2101312200301200-3211312322002101-2032233223332121"></a>

## specific_vip property — virtual_network / 012201032320 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-1301113323003210-1312131333020111-0112323001330313-0221011000012112-3112232300320223-3302332320201332-3210311213230221-0120230133031303): complete subsection reference.

<a id="canonical-0123300332130201-3323132303031031-0121111232301211-3312101130222002-2022122320023012-0113203103003213-0112210232222121-1001033211020333"></a>

## Next pages — virtual_network / 012201032320 / 6

- [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1021103000233223-3303331120323313-1323112223133221-3013231233012103-2200022310102310-3031023330123330-0331131331330003-1032012323011201)
- [advertise_custom.advertise_where.virtual_network.default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-0202002320231111-0312020002301110-1120003023212122-1010130213101001-0312001203022022-1300300212200312-0132302030131120-1001003123023023)
- [advertise_custom.advertise_where.virtual_network.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-1301113323003210-1312131333020111-0112323001330313-0221011000012112-3112232300320223-3302332320201332-3210311213230221-0120230133031303)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1021103000233223-3303331120323313-1323112223133221-3013231233012103-2200022310102310-3031023330123330-0331131331330003-1032012323011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312223133101320-1130111230322230-0031010313312131-2302030012201001-1003213212000123-0013301000033122-3303333312110333-1100021013322323"></a>

## advertise_custom.advertise_where.virtual_network.default_v6_vip — default_v6_vip / 321011303302 / 2

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
default_v6_vip = {}
```

<a id="canonical-2313330123202322-0131210322033212-0233002120021100-1132013301030320-3311011113201211-0111330012120122-1320231001022302-0301232011301330"></a>

## Direct properties — default_v6_vip / 321011303302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110222100120020-2320132313013013-0321113313103200-0101002302330122-1301220213200233-2112002213011222-2103022200102231-2001000303321213"></a>

## Next pages — default_v6_vip / 321011303302 / 4

- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0202002320231111-0312020002301110-1120003023212122-1010130213101001-0312001203022022-1300300212200312-0132302030131120-1001003123023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111230331212112-3011332203323232-2320012123020032-3000203021221303-2002332310102213-0003313032000303-3323221333312320-2300100222101001"></a>

## advertise_custom.advertise_where.virtual_network.default_vip — default_vip / 233120322013 / 2

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
default_vip = {}
```

<a id="canonical-3320333133131313-0103233011110200-3211002012130111-3313201313200022-0233102330000111-0003002123220122-3320231013323020-3102130131121102"></a>

## Direct properties — default_vip / 233120322013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222111312011103-0323131333200203-1011033201322102-2132232210203212-2000003220212320-2330333111003101-1022303023033211-0322223231113230"></a>

## Next pages — default_vip / 233120322013 / 4

- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1301113323003210-1312131333020111-0112323001330313-0221011000012112-3112232300320223-3302332320201332-3210311213230221-0120230133031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110020112233120-1100230121200010-1012203233201310-2201122011001031-3213331002211310-1020312110010102-1203122120021311-2012231003311121"></a>

## advertise_custom.advertise_where.virtual_network.virtual_network — virtual_network / 131003320321 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-2122203103003233-3201300211121110-0330330111103322-2020213320323320-2001322032221110-0122122110332010-2000030321123323-0320231231213013"></a>

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212330031232203-3132302231302122-1032113013303313-2202110120303303-2023010332112030-1311111321111231-0032120010331233-1132313133210203"></a>

## Direct properties — virtual_network / 131003320321 / 3

<a id="canonical-2120122030031333-3330321132323310-0132313111232022-1200012201303120-2333013033333232-0011231111321322-1021300303023113-3210003301220121"></a>

<a id="canonical-1033113202312230-2113002323130100-3333200301301203-0030231032320103-3122210313122130-3203123310123131-3201031222133011-0302021331100003"></a>

## name property — virtual_network / 131003320321 / 4

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

<a id="canonical-2012000312332231-2123200111311000-2331003210330213-2331222200032033-3323112013311120-1330320131220113-2130023231303230-0123003303010232"></a>

<a id="canonical-0021221022011031-2110202001032110-1203011030033323-2113100221221103-2121222122313002-2133020223101210-2221000133320300-3230031120201012"></a>

## namespace property — virtual_network / 131003320321 / 5

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

<a id="canonical-2222310211002121-2331223021033120-2031312211331020-0211313303231102-2011003101112213-0213121121303121-1210223121201222-1330212203323013"></a>

<a id="canonical-2310201320310300-2111323123201323-1011121100332122-0001020333212030-3000331132203330-0111201312213003-3002031203301200-0300012211323133"></a>

## tenant property — virtual_network / 131003320321 / 6

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

<a id="canonical-1012000212100321-1122011102312233-2323313101131310-1002022211021231-3221031220210300-3131231223110101-1120001320122303-0120320310310331"></a>

## Next pages — virtual_network / 131003320321 / 7

- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-2223300303111110-1212110320022323-3223130010312230-2100121331021021-2301222002211202-1232210100000310-2313121021012210-1110203003030010)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-3131233030200102-3231332321203303-2131331231102231-3123031211331102-0330311312320223-0211131322113002-1301003222212320-0133033112320330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301122233313033-3333133233212011-1023112032311201-3232201211002100-1023130222210213-1131000301330102-2013303333013300-0312320132233333"></a>

## advertise_custom.advertise_where.virtual_site — virtual_site / 033210031012 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-3232020011102110-0020333100132333-3232011113311030-2220013311320103-1132303133321301-2303120031101100-3300222233102323-3021210321123031"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322031013330220-3220111000332033-0023033312100313-3230100033033012-0131020130022202-3321230222232202-1330003230002203-1211300301131332"></a>

## Direct properties — virtual_site / 033210031012 / 3

<a id="canonical-1220122213020333-2020200311011000-2303003102322013-1032213333102123-3121011110030111-0121220310232313-2000031331103010-0322203313233003"></a>

<a id="canonical-1021100302230210-1132221211210003-3131032201031213-0220021310000233-0210230021120213-3030123313110130-2022113002123033-2130132212223000"></a>

## network property — virtual_site / 033210031012 / 4

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
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

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-0111130203111312-1122321221330023-0201202202031212-1333213330023212-2032012130122000-0322131300231212-3201132310231011-3013012313311012): complete subsection reference.

<a id="canonical-0130122320321101-2020220231301122-2232011220311133-0211230110100012-3333201332312001-0301232302111100-1111020111230121-0230120033020221"></a>

## Next pages — virtual_site / 033210031012 / 5

- [advertise_custom.advertise_where.virtual_site.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-0111130203111312-1122321221330023-0201202202031212-1333213330023212-2032012130122000-0322131300231212-3201132310231011-3013012313311012)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0111130203111312-1122321221330023-0201202202031212-1333213330023212-2032012130122000-0322131300231212-3201132310231011-3013012313311012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011220332021320-1231332111210212-3033100020200321-0330133220232030-3303022022133230-3230023033322012-1130102110101223-0110231200122123"></a>

## advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 130200213212 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3131233030200102-3231332321203303-2131331231102231-3123031211331102-0330311312320223-0211131322113002-1301003222212320-0133033112320330)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-1331120000120221-0330321131302130-0310332321121112-3023220031002021-3013312012000202-0123121010321133-3220323121131213-0210112033103311"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101222333033120-1102011003321301-0112310220231131-3130222122110333-3301311222212333-0200233322310323-1120112221033210-2120030331010210"></a>

## Direct properties — virtual_site / 130200213212 / 3

<a id="canonical-0312332310312203-1321122130103023-1031332310130033-2203311203321201-3011203323321102-0110113103222011-1312131122031320-2100101133222033"></a>

<a id="canonical-2320123211322020-2113312030312313-2113331123130321-3333103121312311-2223112002113200-2120012030111131-1100211001000330-2012231113220301"></a>

## name property — virtual_site / 130200213212 / 4

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

<a id="canonical-2333023110112131-0210000303033021-2231021221110201-2001201011202133-2312020211020131-2121123103231233-1220012320332320-3102111231221221"></a>

<a id="canonical-2202112333312023-3133200322202311-0103031132211313-2211322101323230-1331213013131203-1233330013311112-0230310001322133-0130310202222011"></a>

## namespace property — virtual_site / 130200213212 / 5

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

<a id="canonical-2333112210200122-3303011203301130-1103212013012133-1211303322102000-1003033012303212-0110323112110021-1312031000232010-2030011001000230"></a>

<a id="canonical-1002010003231012-2221322332112220-2221102200032311-3132303210003221-0233233211301202-2312020321300130-0313223313121003-2120030122132231"></a>

## tenant property — virtual_site / 130200213212 / 6

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

<a id="canonical-3010203120322110-3321310133220001-1302012022101133-3020000100102203-0213200013321333-2023111212322010-1212003212103022-1323231002030230"></a>

## Next pages — virtual_site / 130200213212 / 7

- [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3131233030200102-3231332321203303-2131331231102231-3123031211331102-0330311312320223-0211131322113002-1301003222212320-0133033112320330)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1330330132111001-0030320120001011-3212231113031321-3110223330010000-0031121013222333-3332120333023202-1002100000320323-2202310313121302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222230222122020-1332011301033022-0232102103130100-1322022330332312-3102202332032013-1201221020311222-2132121200103310-2101110300001021"></a>

## advertise_custom.advertise_where.virtual_site_with_vip — virtual_site_with_vip / 133211303012 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-1303013321030132-3232223121202231-0012113213203010-2200001302323001-3322321331021132-2221321210121033-3300221020210022-0013213110232210"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211003312212313-2231021330130322-2021332313022212-2032213312331003-3300310102233020-0101213102302031-3200311010231200-1110123312212223"></a>

## Direct properties — virtual_site_with_vip / 133211303012 / 3

<a id="canonical-1330121220100002-1220123310312003-1320112230101013-3220220210022110-1120220033121312-2103213310021002-1231110233103002-1001103131202132"></a>

<a id="canonical-0310102123221020-2021120030220003-0022301333001122-0031110021302331-0311332020111133-0012321223230230-2133231213220311-0220210333102022"></a>

## ip property — virtual_site_with_vip / 133211303012 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0111111112300222-1132203011231122-2001013120003031-3011221012210011-2032023123302132-0310233223313333-1323111123211012-0120010311000232"></a>

<a id="canonical-2112010120033211-0021023332001302-0303230212113220-2003331013231302-0101022011113213-0001110323332212-0222100200301130-3103222130111131"></a>

## network property — virtual_site_with_vip / 133211303012 / 5

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3123302233210231-3010220233231133-1221200233132313-0333331022023131-1231021211001233-3213112313233033-0312002020013211-3110021312011213): complete subsection reference.

<a id="canonical-1110003120131223-1132222012201031-3303032130120313-1123220211332212-2032133132223312-1302230113000021-1231002121331222-1032303111212312"></a>

## Next pages — virtual_site_with_vip / 133211303012 / 6

- [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-3123302233210231-3010220233231133-1221200233132313-0333331022023131-1231021211001233-3213112313233033-0312002020013211-3110021312011213)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-3123302233210231-3010220233231133-1221200233132313-0333331022023131-1231021211001233-3213112313233033-0312002020013211-3110021312011213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223023101221210-0223220032201022-0321213221101210-2313201321130020-1132033330222003-1100323312131112-1132103120130022-3302321202110303"></a>

## advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — virtual_site / 110321230100 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1330330132111001-0030320120001011-3212231113031321-3110223330010000-0031121013222333-3332120333023202-1002100000320323-2202310313121302)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-1203330233032011-0221303100302323-1013131233023130-1220131103001322-1001200032322031-0122001232212021-3121103223323300-1023332312211132"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221100101031203-1313130231003003-2020210033022012-1230331230131103-2022211011201231-0131120232021310-1211312112120323-1013332123232021"></a>

## Direct properties — virtual_site / 110321230100 / 3

<a id="canonical-3203012213202103-3330030321011311-2202322103300113-1012333301001201-3021301122312203-2100022023202313-3002233211000032-3003101213003021"></a>

<a id="canonical-2333332200110302-1030030001312230-1002312210301310-3203122003031003-0112330021103311-3003331102323301-1301211212301010-1202111120202112"></a>

## name property — virtual_site / 110321230100 / 4

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

<a id="canonical-1011001132113301-3333010013000033-2130121131300223-2330133132330223-3131200332032211-0131310022321003-2310022131323110-2020113013223311"></a>

<a id="canonical-0110302231103301-2122000323110020-3311130131331102-0001130212033233-3003101000233203-0312200320112001-3010312011322311-3312300221131021"></a>

## namespace property — virtual_site / 110321230100 / 5

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

<a id="canonical-0220312123212233-1232202230112321-2303231212010100-3333030103101131-1033113113002300-3033222113103012-3012230313320120-0211032021203111"></a>

<a id="canonical-1131031103202001-1101302301103321-0110001200000201-0323332322020013-0100310303010320-0232012332201101-0111011310322123-2200101201311120"></a>

## tenant property — virtual_site / 110321230100 / 6

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

<a id="canonical-1220233023113213-0021132202130323-1310000201033122-1330031310220221-2000311120120313-0001021112220112-2102020010002230-2013330033113111"></a>

## Next pages — virtual_site / 110321230100 / 7

- [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1330330132111001-0030320120001011-3212231113031321-3110223330010000-0031121013222333-3332120333023202-1002100000320323-2202310313121302)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321222123021301-2312133213202331-0101033302231332-2300232220202001-1110020321020102-1323101230011331-3113203011333310-3302110121032320"></a>

## advertise_custom.advertise_where.vk8s_service — vk8s_service / 101120301220 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-0132203110033033-1223031113332002-0020213220310322-0221232200022322-2212022032321002-1110322202210223-0102202233220312-1011233031230323"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330001202130113-0112333131031101-2000110033022133-1231203121301130-0320122221220030-1033001131111233-0301312103013331-1323103232231330"></a>

## Direct properties — vk8s_service / 101120301220 / 3

- [site](resources--udp_loadbalancer--reference--group-001.md#canonical-1112113301102202-1230111313132310-0322132212302230-1130021121332023-0222032232032023-1102330301210232-3230212103331020-0030333010102321): complete subsection reference.

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-2230000221313203-0213310122302020-0102233032203313-1032213113001322-1232013231220300-1212022021301230-1213011111230033-2202013012331133): complete subsection reference.

<a id="canonical-1312310031223133-0121303113121231-1112312323323321-0321021111302300-0301133302302200-0013012113102110-0001302123210311-0110323001223101"></a>

## Next pages — vk8s_service / 101120301220 / 4

- [advertise_custom.advertise_where.vk8s_service.site](resources--udp_loadbalancer--reference--group-001.md#canonical-1112113301102202-1230111313132310-0322132212302230-1130021121332023-0222032232032023-1102330301210232-3230212103331020-0030333010102321)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-2230000221313203-0213310122302020-0102233032203313-1032213113001322-1232013231220300-1212022021301230-1213011111230033-2202013012331133)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1112113301102202-1230111313132310-0322132212302230-1130021121332023-0222032232032023-1102330301210232-3230212103331020-0030333010102321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201031221330223-2111232201311110-3220212111210310-2211103312103113-1201300203232010-0203330332322203-3012110012323130-0302030132232312"></a>

## advertise_custom.advertise_where.vk8s_service.site — site / 200222202011 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-1110001011323311-2302202330233110-0030121101122132-2032230213311300-1030010322332013-0020003130312311-1321100213210311-3003330122302000"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010001230233013-1012003331011130-1002132331323333-1303320130101021-0311332033211331-2031002021231010-0123100300120333-2322113332103320"></a>

## Direct properties — site / 200222202011 / 3

<a id="canonical-2303201300332222-0012023001221030-2113203321232321-3231230231212011-0122203213302032-1311322200023221-3310122030312213-3020020223220212"></a>

<a id="canonical-2011230123131020-1022222011333323-1130010320333300-0012200103122130-3231223133220211-2120003312210200-0310111133333102-0332230111010130"></a>

## name property — site / 200222202011 / 4

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

<a id="canonical-0011321330122213-1202003011010120-0313333320213001-1302223223230013-1133100303301313-2022232002233302-2011100130210212-1001133300212123"></a>

<a id="canonical-3232002232103103-2031110032120122-1231121020133122-2201020032213122-2022301222132203-0302103010233132-2023320211013020-1200333203312223"></a>

## namespace property — site / 200222202011 / 5

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

<a id="canonical-1131120112221032-0110130201123313-1111321333200312-1010113211133120-0032203012331121-3003013221013123-1012201322313032-1002302303132010"></a>

<a id="canonical-3322133120310113-1301331321021000-1300111230001112-2021030031121313-3330200120212230-3022102130103030-3121103022120021-0022312213003123"></a>

## tenant property — site / 200222202011 / 6

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

<a id="canonical-2220312030013020-2023311202103133-1013100201310122-0212001022113322-1203331132322310-2333301112110122-2000312212323201-0301310223320311"></a>

## Next pages — site / 200222202011 / 7

- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2230000221313203-0213310122302020-0102233032203313-1032213113001322-1232013231220300-1212022021301230-1213011111230033-2202013012331133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001111313220320-2300021311312203-1331031210201132-3220102230220002-1321001223220223-0202012101023221-3310013112003310-1133133320332303"></a>

## advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 012102303322 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-3331223133000103-1322332203002321-1122332103232201-2211022200122120-0213102203122301-0232221312001131-2012210131003300-0212332332210030)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0202001212223123-3320100130133012-2311321302033132-2023030003031221-0233202133101210-2102102223000232-1301302012111121-2303111032210101)
- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-3301030232010212-1201333333221333-3310223011013103-2231101220012123-0200213221013011-3213332303133201-2211330111302313-2012013013112000"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222133301102303-1133111233131210-2301203231021032-1220032303222110-0302212333332013-3232212331021002-2333100001321111-2222020323000001"></a>

## Direct properties — virtual_site / 012102303322 / 3

<a id="canonical-0202021221332203-1203212223132301-0203113113012313-3113212100232123-2232110111013300-3123110021110211-3323132231020121-1330321200011213"></a>

<a id="canonical-2121010000023010-1222103113321012-2203300313320310-0312110223120101-0303123030333030-1032112133311222-0201302032311013-2130121322012323"></a>

## name property — virtual_site / 012102303322 / 4

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

<a id="canonical-1113101012213132-0332023102322020-0210320110223211-1030001331123032-3121130222033210-1301222221210303-0211020331221111-3203030131321033"></a>

<a id="canonical-3323123210033122-2333013213221111-1131211110303132-0003123333123312-3230233223130021-0213031113031201-0313212321321023-1311302233233331"></a>

## namespace property — virtual_site / 012102303322 / 5

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

<a id="canonical-0131000012320331-1133122020002330-0233321133012322-1201203210313301-2210011010123100-1001112123211002-3323323230001331-0000122202203020"></a>

<a id="canonical-1103201300210111-3330012121330013-3012022132002023-3302121202033003-3132001232312110-1221321033111202-1123231310300232-0301032222200210"></a>

## tenant property — virtual_site / 012102303322 / 6

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

<a id="canonical-0332210233201202-0103011222011001-0230002022321100-2231331332010211-3223230201300200-3132223313131111-1033210101311322-3103122010202101"></a>

## Next pages — virtual_site / 012102303322 / 7

- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-2232212230000120-0322332002012300-3333033310100111-2200103130322111-0201003320021310-2012001031200230-2132203111211221-2311133212111222)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-3121001210223201-0002101322213023-2213011332023133-0223131303332232-0202230121130203-0330221122033201-1300103332030211-1301020220301201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000302201000232-3020330133113220-1200133111021311-3123112232001013-1332001110312020-0220110333322233-1310301020131033-1212232311013230"></a>

## advertise_on_public — advertise_on_public / 333122332130 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- advertise_on_public

<a id="canonical-1131331003210000-2331113000100013-2232131210301213-1021112230322000-1222213303123001-0302322101033001-0010131203112121-2123220013211113"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

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

<a id="canonical-2032332311331113-2331001300200231-1110011002113330-0033211222331331-2120213230000032-1033203100021212-0132200100321321-3312010100321202"></a>

## Direct properties — advertise_on_public / 333122332130 / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-0100211213221202-3130330222231033-1022322320031300-3212313220311332-0012331220312312-0121320232111233-0331033323331122-1313212022200103): complete subsection reference.

<a id="canonical-3320113302300213-3301200210003313-0321021303213101-0123013332033200-3000100323101103-0211301131113300-0311012301313230-3321210011320222"></a>

## Next pages — advertise_on_public / 333122332130 / 4

- [advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-0100211213221202-3130330222231033-1022322320031300-3212313220311332-0012331220312312-0121320232111233-0331033323331122-1313212022200103)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0100211213221202-3130330222231033-1022322320031300-3212313220311332-0012331220312312-0121320232111233-0331033323331122-1313212022200103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113021032122231-3112003123032301-0223032022223110-0320013030033110-3333321110211130-1121213303021212-0010112331223203-0202033212123303"></a>

## advertise_on_public.public_ip — public_ip / 230020103012 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3121001210223201-0002101322213023-2213011332023133-0223131303332232-0202230121130203-0330221122033201-1300103332030211-1301020220301201)
- advertise_on_public.public_ip

<a id="canonical-0300103122321023-3000132322101211-2230222020221133-0012031321211210-2300030331103320-2201031002210020-2330301221230310-3212123202133031"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101101030022020-1033303233102122-0122033111301023-1313301313032200-2311021122023330-2223032011022000-1303233023330320-3233122011132022"></a>

## Direct properties — public_ip / 230020103012 / 3

<a id="canonical-0111110121032203-2332103123111313-1100013111131210-0013023230231022-3132112323122112-1332210211001321-0021321201312111-2023000211312332"></a>

<a id="canonical-2100202000200103-0301100330130000-2101013112213303-3011110201112030-3001002020110133-0230133113210312-1021302312102030-2212313031022032"></a>

## name property — public_ip / 230020103012 / 4

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

<a id="canonical-2203221331301322-0133301002111213-0033001211222300-3320221322123331-3333103023023023-3012333302023113-3330113300030000-1321202220232130"></a>

<a id="canonical-0310101221022120-0113231030321000-3123033113312220-3321012331012100-3220301313113230-3030300023010231-2232121202210133-2110130310223010"></a>

## namespace property — public_ip / 230020103012 / 5

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

<a id="canonical-3322133321113131-1003200203131002-3122112203331212-3023221301112203-0100232120221313-0033022213220131-1102210230322132-1121230313021322"></a>

<a id="canonical-1210310031030212-2032013323312222-3003132022103333-1020312223211320-0200030111031130-2123031323100323-0312111022333131-2222203202303220"></a>

## tenant property — public_ip / 230020103012 / 6

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

<a id="canonical-2322001131100010-1231000033303110-2301010122031123-3211203230120311-1331000313221033-3000211210330222-2022130220332133-3311223212100002"></a>

## Next pages — public_ip / 230020103012 / 7

- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-3121001210223201-0002101322213023-2213011332023133-0223131303332232-0202230121130203-0330221122033201-1300103332030211-1301020220301201)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1032133030103213-0020211103223323-3100222332020210-1003331202302203-3333101310331003-1200101122201012-2332103110302022-2333031032112103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001331130231001-1101211221300200-3030301213102103-2322000011202223-0033333303131201-1322002131223302-1201210311220300-2310031231322033"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / 222122203221 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- advertise_on_public_default_vip

<a id="canonical-2302323123201032-2001211223302111-0132220033313002-3021301203111312-0321201310212003-0323332122132312-1131330222320133-2310133320001233"></a>

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
advertise_on_public_default_vip = {}
```

<a id="canonical-1132030221301130-2131331313002100-0311331213331120-1013322332112221-3003232003002020-0110021013230232-2210131220212000-2230323100121130"></a>

## Direct properties — advertise_on_public_default_vip / 222122203221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023300011203221-2310212302301313-1013232100112310-1311002310100201-0110303213232302-3303323333311230-2030322323121013-0131131320120202"></a>

## Next pages — advertise_on_public_default_vip / 222122203221 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2021012000001001-1013233020022212-1310012310002230-0213221320021221-3313312011223103-0233332113202000-2001321301000221-1032223300220320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320322331122311-2213201102222021-0030121002110330-2303323332220113-1000031011211331-1010133330013323-1000110103322233-2230122223310202"></a>

## do_not_advertise — do_not_advertise / 110221100110 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- do_not_advertise

<a id="canonical-0221213122111200-3301201210111322-3031222200121111-3003321003021322-3220301223013001-0210223020203022-0323301031102012-3132332221032200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

<a id="canonical-1323230012032021-1130330131100310-0020002131110030-3203032113100322-0233203302111131-2032222000312130-0302113123110220-3101122210122321"></a>

## Direct properties — do_not_advertise / 110221100110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020200302200020-3031332130302002-2132331312031230-1010322212321112-1103311112221211-0000003133333123-0101212232112031-2310332311210202"></a>

## Next pages — do_not_advertise / 110221100110 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2130020320231110-2310332213003130-2103302031122120-2101333103230022-2103002230301132-0101132133321211-1003113021330223-1110002031002313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002211033330233-1301333232101130-3132203003301223-0223222033132120-3132230111013033-1001100320100223-1132311300101303-3033102332032333"></a>

## hash_policy_choice_random — hash_policy_choice_random / 202010100033 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- hash_policy_choice_random

<a id="canonical-0211210323021133-1013230231122300-3022123211212132-2220010210030202-2200303032320110-1130132332010002-2201212131212233-1220332022231120"></a>

Type: `["object", {}]`. Optional.

\[OneOf: hash\_policy\_choice\_random, hash\_policy\_choice\_round\_robin,
hash\_policy\_choice\_source\_ip\_stickiness\] Configuration parameter for hash policy choice
random.

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

- [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-0211210323021133-1013230231122300-3022123211212132-2220010210030202-2200303032320110-1130132332010002-2201212131212233-1220332022231120)
- [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-3203002003212210-0003011020322120-0120231133003213-1123112300132001-0310031323210330-0303003333311213-0000030033003230-3321001022101020)
- [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-3312322012133300-2111203231230023-2102011121323012-2220112313111130-0033333102130223-0233303123103301-2033223231222113-3030223000321003)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hash_policy_choice_random = {}
```

<a id="canonical-0031210131213003-0123231120211321-0230013202200221-0100120231222001-1312231231313122-1023013111010021-3010321023003330-2300321311022133"></a>

## Direct properties — hash_policy_choice_random / 202010100033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230012223220112-2321210102103130-0120113003313332-2000113333110121-3210112030100200-0130011121120121-0001303212213010-0220013110311132"></a>

## Next pages — hash_policy_choice_random / 202010100033 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0033000022201003-1303032013200210-1033221223122313-1001020223133221-0221333001113000-1302131100323201-0111021013101131-1302332010012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211213020223212-0111200113033213-0001123132001301-2100203313312312-0303130201303212-3003120323303303-2323102331333201-0122122200333321"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / 331213013202 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- hash_policy_choice_round_robin

<a id="canonical-3203002003212210-0003011020322120-0120231133003213-1123112300132001-0310031323210330-0303003333311213-0000030033003230-3321001022101020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for hash policy choice round robin.

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
hash_policy_choice_round_robin = {}
```

<a id="canonical-2313232021322310-3322333210333112-3102023303232313-3002211101103320-0201003121122230-3001011031211230-0211220020022031-2003332012213113"></a>

## Direct properties — hash_policy_choice_round_robin / 331213013202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321000200022230-1132030132212022-1013002332333220-0010112330132033-0003120221323302-0310103002000222-0302201032221211-1112102132223102"></a>

## Next pages — hash_policy_choice_round_robin / 331213013202 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-1232202330301332-2233232122032300-3311123230120100-0221111330223031-0212333133112002-2320202110313302-2212102131223121-2130011100002212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202220123003200-0110211121202120-3001020002132103-1020310021230130-2100101221210232-2003311002020110-3223200322333101-2330223331200003"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / 132330030133 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-3312322012133300-2111203231230023-2102011121323012-2220112313111130-0033333102130223-0233303123103301-2033223231222113-3030223000321003"></a>

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
hash_policy_choice_source_ip_stickiness = {}
```

<a id="canonical-1200022232133320-2132202130110322-2131313020333201-2123322202123122-2112032023302000-2003303013013322-2302003310320000-2311221223222121"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / 132330030133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023112033010123-2201300321011000-2011313031223120-1232031110121221-0313022303010231-3102210313120032-2032000123030210-1320231330213331"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / 132330030133 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2222312133022213-1230023330021312-1132323131113230-2333332022331122-2130002003301133-2032133133210033-3100313020020301-0031233232230012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231330130313310-2122201001201330-2011202130131230-0020222230223333-3122110223312221-0001023030323120-0213030333001301-3323122031213231"></a>

## no_service_policies — no_service_policies / 112300210131 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- no_service_policies

<a id="canonical-2011121232032223-0120303113122000-0331200013101323-3211231010033023-3312301212123101-0023220313112111-0020212230101011-1311120023033020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

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
no_service_policies = {}
```

<a id="canonical-3230323012023020-3203010333221320-0311130222121101-1121010132212300-3132331323303222-3033231313220033-3111100310321212-0001311211232221"></a>

## Direct properties — no_service_policies / 112300210131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303223023202100-3112021303232232-1321131020312231-1003121323231322-2333220132330130-0023223230123012-1201131003303211-1102322301121003"></a>

## Next pages — no_service_policies / 112300210131 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233322132133000-2202101223222201-2323321123320130-2222213112103100-3020113301111231-3003321122011212-2102100101020310-1003333022312320"></a>

## origin_pools_weights — origin_pools_weights / 220310330133 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- origin_pools_weights

<a id="canonical-2301103112020300-1300123003220110-0213202302132020-0033322330302211-1131212132021130-1333332213221110-3233130312003030-1303330011103122"></a>

Type: `"object"`. list nested block, Optional.

Origin pools with weights and priorities used for this load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_pools_weights {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211212321230110-1003021231230011-1012020021212311-3310300320102122-3102113332100312-2001030222300302-0231120103001323-0012110131000212"></a>

## Direct properties — origin_pools_weights / 220310330133 / 3

- [cluster](resources--udp_loadbalancer--reference--group-001.md#canonical-2100203212332021-0123200211033303-0311231121030010-1303112033201300-0120110221003033-2122011111202030-2203022113120221-1310003321312301): complete subsection reference.

- [endpoint_subsets](resources--udp_loadbalancer--reference--group-001.md#canonical-2212331323003111-1030313120020032-3212012123121133-3300233100322002-3130033121132002-1011101220202303-0200122133211232-2021002122220331): complete subsection reference.

- [pool](resources--udp_loadbalancer--reference--group-002.md#canonical-0322213321213202-3202110231110133-2321001131201301-2332103312001133-1123013211012223-2302321130233000-0320333322330213-0013310000232022): complete subsection reference.

<a id="canonical-2220102333322032-2201223002131230-0030212330103001-0323023233331212-1020223222111002-3020232320200310-2333311320003300-0301311121302220"></a>

<a id="canonical-1311322221220310-3001112312012001-0033102201313113-1221310132113211-2302002211033023-1203320021210331-1030012213331300-1021011232110001"></a>

## priority property — origin_pools_weights / 220310330133 / 4

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0203333331223002-0132110010031013-0332102123022022-3213032211323331-0212313302031303-0002111303212011-1331222233130002-1101000223030023"></a>

<a id="canonical-0201310010002321-2100102130302032-1233033221132130-1333310232021310-3301111012202010-2030311213322302-1001211300330303-0210003301211222"></a>

## weight property — origin_pools_weights / 220310330133 / 5

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2200011001122311-0312011333233221-2303320332120323-2133312020003222-0311332231201201-0110313132301323-2032333303333321-1323102022012121"></a>

## Next pages — origin_pools_weights / 220310330133 / 6

- [origin_pools_weights.cluster](resources--udp_loadbalancer--reference--group-001.md#canonical-2100203212332021-0123200211033303-0311231121030010-1303112033201300-0120110221003033-2122011111202030-2203022113120221-1310003321312301)
- [origin_pools_weights.endpoint_subsets](resources--udp_loadbalancer--reference--group-001.md#canonical-2212331323003111-1030313120020032-3212012123121133-3300233100322002-3130033121132002-1011101220202303-0200122133211232-2021002122220331)
- [origin_pools_weights.pool](resources--udp_loadbalancer--reference--group-002.md#canonical-0322213321213202-3202110231110133-2321001131201301-2332103312001133-1123013211012223-2302321130233000-0320333322330213-0013310000232022)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2100203212332021-0123200211033303-0311231121030010-1303112033201300-0120110221003033-2122011111202030-2203022113120221-1310003321312301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003203113312011-3100031210012110-0331021130021132-1003301200110331-3323322212203011-0000203022322122-1333321031112323-2103302312120002"></a>

## origin_pools_weights.cluster — cluster / 232310223310 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303)
- origin_pools_weights.cluster

<a id="canonical-2032330100001121-2112122010230011-3030302220110223-2133301002202010-3012323000030223-1331330020011100-0011023030332101-0222302000202112"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033131030130113-3313231110112212-0332121000321110-3301021300133112-0232030222303132-3022031113001322-3231133021320211-0023100321130100"></a>

## Direct properties — cluster / 232310223310 / 3

<a id="canonical-2000333012001110-0231002120102101-3303003121210010-1332321113003102-2100332133202003-0313301213120010-1303331103120322-3120132233021000"></a>

<a id="canonical-2313102003001330-3113223123232031-1312023012002320-3313233201001121-1200211100332022-3001022303331010-3210033332300023-3020130231113332"></a>

## name property — cluster / 232310223310 / 4

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

<a id="canonical-2011001300233313-1323212333101200-1110213322331003-2020030231101103-3103123231031200-3113011130033133-0313031103110310-0302031320303021"></a>

<a id="canonical-1302032211233103-0213011332203103-1301310102323033-0122323010303330-1200023032112032-2302331212223103-0023032220310200-2030213200002033"></a>

## namespace property — cluster / 232310223310 / 5

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

<a id="canonical-3122301133222333-2212111332333010-3001111010231010-0202203021011012-2120200200312310-3310203332230012-1033120323113233-1013101002012013"></a>

<a id="canonical-0313121301012121-3131212323223112-3113310123011133-1021211300303012-3220222223303212-3121110032302001-0122323223013200-3231002013333130"></a>

## tenant property — cluster / 232310223310 / 6

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

<a id="canonical-1031211001312000-3123323121211300-0222100113021332-3203300322020312-2013313211030230-0322331211220122-1101333313131123-0211013223301130"></a>

## Next pages — cluster / 232310223310 / 7

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)

<a id="canonical-2212331323003111-1030313120020032-3212012123121133-3300233100322002-3130033121132002-1011101220202303-0200122133211232-2021002122220331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313011133220200-0002302102022132-2020023303222012-3323001300010212-1000230021201102-0203200332030033-1303301301002200-1001233110103322"></a>

## origin_pools_weights.endpoint_subsets — endpoint_subsets / 301002021322 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-2002023213111000-0212122002033233-3102221313223201-2231030031021232-1300231222302011-1123120120031031-0312210221210100-2012113302110103)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0032213313100301-3212011112101330-1011333321321000-2333033310333031-2321102233211130-2001102232222211-1021030023313221-2023031320303303)
- origin_pools_weights.endpoint_subsets

<a id="canonical-2132312123111111-3003130211230121-3300132131022213-3333113133210032-2201200312010223-0321000123030310-2023102101033332-3123023213221201"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

<a id="canonical-3313011210201210-0332001112112013-0112200210202023-0113211013123023-2232213012231001-2302001223120232-1112322033021232-2230030201312001"></a>

## Direct properties — endpoint_subsets / 301002021322 / 3

This is an empty object or choice marker. It has no direct properties.
