---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- Property reference

<a id="canonical-1212313130220323-2002220212030200-2100013230103102-2233012103211130-3011333131201130-3133123111111231-0103233213320110-0031102300310102"></a>

### Direct properties for `xcsh_tcp_loadbalancer`

- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1011022010101003-1111012300021023-0113013102330111-3011230222012201-0130101131130322-2013322311103222-1122320332303032-2001122100033100): complete subsection reference.

- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000): complete subsection reference.

- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1112201032021131-3221000132200333-3131001223100003-2232313233132313-2111023302121223-0310220211131233-2130013222232323-2223231123010031): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2103301002331000-0011121122110023-2213202121233302-1122011020100033-1201330300100321-3021121333233123-2002013133223103-3102121301033313): complete subsection reference.

<a id="canonical-0033213300222203-3110133001130033-0121100202230300-2230302202321220-0202300222212032-3101212110102021-1033021113021300-2203330101302323"></a>

<a id="canonical-1001110022211303-3030000223223132-1112213103123300-1303010323210222-2123023301101031-1211301030212322-1021121101030233-2312202130331301"></a>

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

- [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1120310112113021-0011220322222102-0130122230032031-0102010320303202-0001311211211213-3321333330113130-2001230311133133-1322321301312102): complete subsection reference.

<a id="canonical-1330332021221011-3122233101332003-3031323100303332-2101000333002222-0121020323102222-2122020121110100-3211313202313031-0321133100012232"></a>

<a id="canonical-0112112132330000-3122230222132233-1222212320330220-1310133332012332-0033103301132323-1312232032102201-0220321000311130-3210213003332302"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the TCPLoadBalancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3012233000221133-3320101201020032-3032000230231012-1323210202201322-3300231012303120-3202131020133300-3210001113311032-3130012331323210"></a>

<a id="canonical-0032032321112112-0112320012203022-1022013313101001-0101301010023310-2233312131133322-1202123221303213-0020300100300332-0022232001331303"></a>

#### `dns_volterra_managed` property

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. This requires the
domain to be delegated to F5XC using the Delegated Domain feature. Defaults to \`false\`. Server
applies default when omitted.

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

- [do_not_advertise](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2301200001232003-2230220223123022-1031331131120332-3030120311113103-0010103000333222-3012312213013120-0313333310330302-1322122222220023): complete subsection reference.

- [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2102221023203330-2101321113012132-3212312223313302-2333003020300013-3212302122333112-1203323112212131-0113122301023121-1123321312122331): complete subsection reference.

<a id="canonical-2000231333311222-2002031100131000-2230002313113222-0113312303300303-1321321023003220-2302230212313230-2311201201321031-3231010312211301"></a>

<a id="canonical-0020133301101303-3011121322013003-2210031023300123-2021213210013013-0213110311231023-3333202223002331-2200030222333010-3211121132232312"></a>

#### `domains` property

Type: `["list", "string"]`. Computed.

A list of Domains (host/authority header) that will be matched to this Load Balancer.

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if SNI is activated on the given TCP Load Balancer. Domains
also indicate the list of names for which DNS resolution will be automatically resolved to IP
addresses by the system.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1233323330321012-0022220130111220-0001311203331000-0032101313230230-0003312301003011-3322313221113031-0000100011331113-2312101310330012): complete subsection reference.

- [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1323103221023221-2313321022023032-0102100310112221-3013030013003012-3002031302202202-0020002001010103-2300123101332113-1301322121013310): complete subsection reference.

- [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2001222213202013-2110122010013020-0222012122321312-2221120003230033-3112000110322320-0030000330021323-3103332101033100-3231323011021122): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2101213003031210-0111210032000102-1202211230323123-2123332011131030-3210312021210331-1330112233100331-2210212303011011-3120212322222111): complete subsection reference.

<a id="canonical-0020333001002103-1122302123333332-2013131230011012-0032212333311122-3210222332133010-3322012300130233-1232132132313130-0123100231113333"></a>

<a id="canonical-3202201231022131-0311322112021211-3301021233011112-2010302102031022-2300311212031103-2100110121021332-0021201212210122-0331202233223301"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2033132233011223-0320031312303221-3311120111100120-3202022023120011-3003000110233213-0312101012300333-0121103033220012-1313133230212303"></a>

<a id="canonical-0200200133230122-2110101002000030-0201003110011120-3033000203331322-2303211021102032-3200010100320003-2323002103111030-1201103231213333"></a>

#### `idle_timeout` property

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
Server applies default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4147200000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4147200000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4147200000"
  }
}
```

<a id="canonical-1020321332302103-2201122000120011-2323122303212201-3010022111331123-2120133002101103-0131132130213220-0022130311101311-2021002232001300"></a>

<a id="canonical-1002330321110021-2330313210020123-0132333222313202-1021203220312312-1133132230323223-3221223303302330-0120221210003310-3332133232000311"></a>

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

<a id="canonical-2222203011331003-0310123211011110-2102011131222123-2003213212232211-2002112321132001-3002100220103221-0230311323201121-0313123210302122"></a>

<a id="canonical-0301223301323302-0320113310120132-0203011113120322-3111331323212001-1130012301032230-0211122203033111-3002221102330013-0231330030133023"></a>

#### `listen_port` property

Type: `"number"`. Computed.

\[OneOf: listen\_port, port\_ranges\] Exclusive with \[port\_ranges\] Listen Port for this load
balancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [listen_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222203011331003-0310123211011110-2102011131222123-2003213212232211-2002112321132001-3002100220103221-0230311323201121-0313123210302122)
- [port_ranges](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3010332213102001-1023232101230303-2030133221032011-3122131332030021-3200103113222130-2101131210332021-3322113020113002-3233031311221002)

Select alternatives according to the provider validators above.

<a id="canonical-2301311022033310-3331211300221332-1300102220212213-2100022212021132-2223210111031103-1313113312110320-3122201210123012-3022103030110031"></a>

<a id="canonical-1111221030013120-1003212332102212-2130201233312323-0213031021312320-0023011033301033-1330003021230133-1031222200110022-0112132132013013"></a>

#### `name` property

Type: `"string"`. Required.

Name of the TCPLoadBalancer.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1200301120201122-2332131010212332-0132320032132212-1012021212111311-0301220220203102-2332031311020330-2123012003010022-2230231012211323"></a>

<a id="canonical-1312230123023300-1023010031103022-1122311222233113-1212030302110221-2211311213113203-3101030023030333-2122110121331023-0122011231300020"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the TCPLoadBalancer exists.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [no_service_policies](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3232303121303303-3310101031211332-3320310013303113-3130011003231312-0221113013330021-0103001323103000-3230002332010321-1123002233310003): complete subsection reference.

- [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3323311322022213-3030031301102132-3222213210012013-2113301030000013-1230313033002102-0010100310021223-2002303301021210-3332001213302330): complete subsection reference.

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2021103311211123-2330202113200202-1332113221211301-1110121021113333-0023121111313031-2113021130032321-3221022112010021-0213201110302101): complete subsection reference.

<a id="canonical-3010332213102001-1023232101230303-2030133221032011-3122131332030021-3200103113222130-2101131210332021-3322113020113002-3233031311221002"></a>

<a id="canonical-1121031312300230-2321113133223020-1333102002331222-1021323211312113-1020103122300200-0122112212303201-0013233222300223-0100332223120033"></a>

#### `port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3030013003230233-1020032322120203-1023112003112123-2102002023033311-2321113323313010-0110100232101332-0123011112033022-0022330033023211): complete subsection reference.

- [service_policies_from_namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0123020233112322-0223121002333303-2100322322223000-3300212133321100-3320102320231010-0321033323023300-3233233123103010-2121331023133111): complete subsection reference.

- [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2133300001211310-3203211112301011-3012211322222112-3322000131211112-0133200313302211-3311012013010002-0330310112023122-1200113222020223): complete subsection reference.

- [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2032030011011133-2021231112123000-0010123000231211-1233113301302232-0220010312011312-1000210021110111-1203111330233211-3303302223213330): complete subsection reference.

- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2220302321031320-2132003000000313-2010131000231032-0103312202200323-3330011103311210-1121000223113223-1203133210113132-1001300030000301): complete subsection reference.

- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113322320020212-1232030003220331-2123323103212303-2201022020230003-0212020022223312-0202133131320033-1021102022300310-1323231323210221): complete subsection reference.

<a id="canonical-3102132300333311-2303311333301203-3330321101312310-0222103133301001-3132321231021200-1002212312332323-0323213232201230-3110212232222331"></a>

### All schema paths for `xcsh_tcp_loadbalancer`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3120023000002123-2201220122222212-0000023001323002-0032102201321122-3201322001220310-3012222002320012-0331211022313201-2312123012011311) |
| `active_service_policies.policies` | [active_service_policies.policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3022103213023223-2031203303201102-0312223311130020-0331000121223103-2022322210132200-2332322112223120-0312332133213333-3310031210322111) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3230231100112011-0310031033230101-2231101031321333-3220002113223022-2201123303312011-0120022203030002-2203311321321211-2110120223333322) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3313200212201121-3301301003300102-0321121121103002-0222003123233201-3021331300210312-2113011030323130-0132031310231313-1111002303221001) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1203021322211230-3321330302211101-2001320231113230-1010131003013111-3032100121331312-0221321232220201-1331001101120201-3001310230022210) |
| `advertise_custom` | [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1013012233212133-3023131131211102-2203001220030233-2131211130201223-1323322001311122-0223103302132103-2032132122323131-2030201131222013) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3310231133113102-1232110011013322-0311030213330120-3103111122222101-2012231121021022-0032231320131032-0312033301101102-3311333210023121) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2231230123023333-1303233212120200-0210210311012300-3313013313132012-3003201213012012-0310130001320333-1010310122120111-0230001010001333) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0221030113121220-1102213003000323-2021322112321222-2000323032321321-3012233203131003-0233310202112013-1323020331223233-2212230030221021) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3200310213312303-2302011230111300-1133010110011100-3331022000202010-1213010102321322-0220021120030320-1322213000203113-1303123203310020) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2020100023101320-1123111211221332-1033313032130223-0212011113130233-0303232310323100-3221312202302110-2100012030112310-0320203333130030) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3112330022132120-0000002110302311-2302311113112103-0202200331112102-1023030110002121-0023023220331010-2031212110111202-1012002222213113) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3002313330003213-0333313000001100-1002121200231321-1102102212030113-1312322022010320-1302103013202203-2031312200100201-0031310332232300) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3303330230303111-3203201002010201-1223221132102301-2131200021210130-0132203222102011-0230113201202311-3033023022221311-1312322323120331) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0123113321121212-1133200130030011-1132322122332021-3211312311333331-0211220320032121-1022301302221311-0113200033321212-2021010030011111) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1310200002013312-0332123320003323-3312212023100100-0012133010222311-1002020020010212-1203211133130113-2222001022301321-1300123102213011) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0100023020332332-3111211211111332-0220012221200023-3221332312223303-2331332212231322-1000230112220121-0201123120332310-2110000232333320) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2331013112211123-3210122222111103-0110211303020330-1230222120220312-0101013212301320-2333001123212200-0322112323231123-0003011131100211) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3103221033000113-3031002313133213-2330003231021330-2102203112311221-0130223013111131-3333122333213330-0000231310230121-3023032232200312) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0021211230030301-0203030030210030-2011300200303230-2110030121312233-2113121201131132-0112222322020333-1233301301313000-0003010221201101) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1020023313211113-1112010333313332-1111122201312111-3231131301031032-1013000132121021-0013123302323120-1123022332311121-0031012222033131) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3233110013333330-1010302023332230-2220222012013210-3103013030012220-0322120132301233-1303231130020303-0031022320112111-1200121001113313) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1120212332121220-2000023032123302-2323322112122211-3130332122121130-3302132312202023-2313322213301201-1002221123230311-3111131021301330) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0110130103100310-2112310323110230-2033122033131201-2021321013202220-1321132213001302-0212123010031321-0320310012213212-1101332223000132) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3120002133123221-1023122202112031-0211321000303220-1301321133200322-0100030102133200-1200201023233302-2311113132230311-3030211111123130) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0212033211312200-3012123331133331-1320302232322023-0112002202333332-3301013120333131-2103320123201023-0122101103123001-2222223321330323) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1132310132123321-2332200002130112-1330023011031022-3221203310003311-0102321022310032-1233030030000333-1112310200332213-1130330303103222) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1123013230331303-0111100002013203-2311002023111001-1211023131203300-3130231122301330-0101012331200301-3333123100311031-2212110022212003) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1220121123303102-3323230312213021-1030120002210112-2331331111102230-1010312323101210-1300021023312231-2012202111020012-0022101233222221) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1323300123212001-2313312220022331-1012230302010111-2022310130131111-0333213312010031-2232333123100302-3203231213133110-1333201001201322) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1122113220301021-2312111120022013-1321012320032302-2130110110211211-0233312310230002-1221032121102311-3310210100300031-3122212031010212) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3000301102201130-2220321212111112-3100312131010120-2310001201030232-3122333332300131-1212100111121112-3121300322212220-3121312130131202) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0221023101223332-2003212210330220-1230031330123113-3200221330223130-0203311001103010-1321111011130222-0212003010211130-2000220121020212) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2211101111223222-1000300201320321-2123121212333231-0300213121300212-1130131110011231-1121300202033002-1021332103102112-2220122323301332) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0012012310122300-3113320200233232-1330102222111230-1111120103333233-3130111130133332-0103120001201133-3112033032100211-3303110330031302) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3133011032003200-0313313230021312-2232021003003131-1021130030333112-3133123303223212-2023101313010112-3003202230122210-0012211022331230) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3222130011021322-3011031312100021-3232233213221313-0303112131223100-1021230332321333-0222123033102130-2200220113302220-1102313333113213) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0011003031001313-1032323211101023-3100100010031132-3012010332230231-2120130312130020-2203023031312023-1112320233212332-3122103033223012) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2033323033031232-0323103233210101-3303000102230312-3310111213330330-2030001001221200-0100000033122233-3330313321030101-2132330222030301) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0300233202133223-0000300210201322-1122132013030330-0032100330013020-2131320110333123-2032110111310203-1201112013013302-1302012203223110) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2100321003032002-3200312200232202-1110110001200103-2332133313223331-0331132122230121-1300002010013201-3012320100233002-0232210331102231) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0310322333022123-1022002101100210-2213302011031032-1210032200203100-0001221333030220-2101333313002102-0103032202132311-1022130022102123) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2112011131311231-3030201200120022-1211300211230212-3123130323301103-2023322000103132-3032321202333030-0030033023231133-2002111120321210) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0000113103303012-3332233323003130-1210032120132333-3330301011121100-2322030311132211-3320002012201213-2022231211023120-2012130003100210) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2102110011023022-1231102131123101-3011023230111312-3220022310310100-2013110020231022-0200330330112330-1201333011121221-2203301010012322) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0221121001323332-1223313301121300-0010303001221332-0320201223131103-2321020130230123-0111320020032011-3321332002221101-0022230323101230) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3101232330312322-0101213102333032-2021311313121302-3331312030003121-2303232110220033-2200033112312120-1103321021001310-1311303233232032) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3203033131312223-1222231300232132-3232010222121313-0220013302122103-1021201003113121-0213333300310111-1121010023232131-0200113230032131) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0322022110131022-1332100303321303-0331221201133010-1320023211313230-3232201100233002-3132133023211310-1221333122113121-1301010332330313) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3213302020033203-2213310121300322-1310111002330112-2223103022120122-2032310203312133-1001323002312211-2211003301300120-1311033002231220) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0313111112302020-0001333312031302-2013021231221332-3210102020233032-1233210001131330-2110133302233322-2133331000021231-3230031201333022) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1103322033230303-3020023223311303-2103130201033110-3322232232121321-1320212302013201-1323130103333033-3023310112313013-3123310330031011) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3000003332203221-0201233032132012-0302100030332111-3302003122230000-2311122301132133-0200210331233121-3012303303222023-1011200111210001) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0023102112332201-1123023323201300-1322203212222332-2001003220030122-1020331102101312-2201332011000032-1330011233032001-1131220320121020) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2311332223031132-0100011021032013-1321220203030310-2300010022212011-0130300201220121-1323321233201101-2121330012301231-0310221121202300) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2330230222331212-2023220202010000-0221322222021103-3131333320123011-1230102022303102-0222113210010122-0200302202103211-1020221311311100) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1100201012023100-1012011000010313-0223331211021111-2001232031220201-3233023333330100-3332110013101011-1122211001321302-0011033221000002) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3133001131012320-1200303131031323-0131310020202300-0011300233321333-0332132131032231-3103003200332133-2322020212322003-3101110003002301) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0120230020323010-1002222102230302-0301030031203033-0222121123023011-2223303031303223-1100212110103003-3110202212121212-2033203030201100) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3111131113300323-1123031310001210-1321030323211210-1233112122310101-3331200221130221-2012310310010113-2212311012033121-0331023201232121) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3201122003211333-3303101211223302-0303021002222103-1110003012102120-2120113022023100-3012100331113222-2112123330201231-0313000331203231) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0012312120212210-0023331020321310-2121022210223303-2300332223321022-0213331131132112-2130122330302233-1210313020310331-0113301101023020) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1202021312220001-3232013103322110-3100112103010323-1121330221122313-3330012013211010-1201031310222032-1113201303033321-3130121021311022) |
| `advertise_on_public` | [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2200031202211021-0313200202312233-1013023211320030-3123000101033002-1113110302231111-2101321210002322-3331003310213213-0312330021311122) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2302330022113120-0132220223113030-1203002312332220-3011132222110320-0100103320332132-3231213110010122-0221301003313120-3312121011200103) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1222030022130330-0331013312310213-0220003120000133-2122222123123030-1223301100031132-0103131211033302-3113031021023121-3022122220220013) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0232110030303000-3323233312210002-0331002220113231-1222032131332110-0320133301131131-0232310110113333-1331330112132330-2303112220222012) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2323212133012203-2231313332303232-0213313201013313-2200010311322211-1221211112100130-0112221313113110-0213322230323013-3310210023331003) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1123330002330211-1310302201203201-1113023310221132-0301203233322302-0032103302032320-2010301030313101-2213001103002101-3232220103002213) |
| `annotations` | [annotations](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0033213300222203-3110133001130033-0121100202230300-2230302202321220-0202300222212032-3101212110102021-1033021113021300-2203330101302323) |
| `default_lb_with_sni` | [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3323001220100213-0332200120333110-0322222321133332-1232031323010000-3212121100010232-0301020212002130-0303222103123032-1320010121121321) |
| `description` | [description](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1330332021221011-3122233101332003-3031323100303332-2101000333002222-0121020323102222-2122020121110100-3211313202313031-0321133100012232) |
| `dns_volterra_managed` | [dns_volterra_managed](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3012233000221133-3320101201020032-3032000230231012-1323210202201322-3300231012303120-3202131020133300-3210001113311032-3130012331323210) |
| `do_not_advertise` | [do_not_advertise](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0001303000210302-2210122211320302-0002200332110310-0310113320302002-1122011331112123-3223311220302121-3230321301031211-3123012131133000) |
| `do_not_retract_cluster` | [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0021013132003110-0210102220132311-3302202220213101-2313122133232322-0211013312030111-2133331021111111-1102303210112013-1030322030021232) |
| `domains` | [domains](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2000231333311222-2002031100131000-2230002313113222-0113312303300303-1321321023003220-2302230212313230-2311201201321031-3231010312211301) |
| `hash_policy_choice_least_active` | [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0213123311230102-3223303113132302-0222311331313011-2233121031121311-1102323230101312-1223301003121202-1122203220021330-1112321202221300) |
| `hash_policy_choice_random` | [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3212033231123032-1011312023323013-1011123300220023-1231033203130221-1201031100111300-1010230221333212-1313301103132100-1122023101232212) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0013003300111331-2221131030212002-0133311311011203-2121231303120023-1320011201221133-2232301003113132-2321033323301211-3331202302033223) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0001313102231321-2102202120232131-3313220313103221-2312113230312002-3312032203302111-0233332302110333-1313210300232111-1311102213312222) |
| `id` | [ID](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0020333001002103-1122302123333332-2013131230011012-0032212333311122-3210222332133010-3322012300130233-1232132132313130-0123100231113333) |
| `idle_timeout` | [idle_timeout](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2033132233011223-0320031312303221-3311120111100120-3202022023120011-3003000110233213-0312101012300333-0121103033220012-1313133230212303) |
| `labels` | [labels](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1020321332302103-2201122000120011-2323122303212201-3010022111331123-2120133002101103-0131132130213220-0022130311101311-2021002232001300) |
| `listen_port` | [listen_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222203011331003-0310123211011110-2102011131222123-2003213212232211-2002112321132001-3002100220103221-0230311323201121-0313123210302122) |
| `name` | [name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2301311022033310-3331211300221332-1300102220212213-2100022212021132-2223210111031103-1313113312110320-3122201210123012-3022103030110031) |
| `namespace` | [namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1200301120201122-2332131010212332-0132320032132212-1012021212111311-0301220220203102-2332031311020330-2123012003010022-2230231012211323) |
| `no_service_policies` | [no_service_policies](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0121112133023001-2133103202133131-1030333112200300-0221311100103201-3330222002023012-0321002320113000-0310301333203033-2200011113112222) |
| `no_sni` | [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2112113331223322-3231232001320233-1302030020311232-3100000332022230-1032022022211232-2010330311022330-0123212123130312-3120230231130212) |
| `origin_pools_weights` | [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3200202332003310-0211023331000103-0121311213201021-2001122321011101-1130023223300221-1121311210313223-3022112123003300-0100230133313000) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1113012100011030-2001322102301030-1131010011033030-1313003012021101-0110011203030111-0032230200000311-3003320330302213-1200101331100111) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2212013331103221-3031332033320031-1210122323123331-1031313113310310-3100223311311321-3120113311131312-3112322132002012-3220331130032303) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3200332330131231-3103203300030030-3031323302132210-3132221121211320-1230030120120212-0003023213012210-1221001331011023-1200032332022033) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0300331231310020-2221010013102111-3022100132132022-3312032231320012-2101033320221132-1123310201312103-1221200033010122-2222310330220323) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3220123132322001-0221230110222113-2301001313220031-2003111011112222-0210003310111311-0231121311123233-3110203332231311-1100133102333121) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1132213133202332-1322012003101323-3330023000200011-0331233133011332-2222300223122200-0331130322222123-2231333130011310-3012311322322233) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1202230130301301-0102231022221021-1132312202023331-0323033331101101-2200020223321123-2133213330222033-3113001133033031-1130322302230120) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3233100010300122-3230221023120200-1113012021131032-0020202101121311-3100220330130000-2333032032122032-2033230320321232-0220233100122310) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3321202331223033-2020301031110310-2212010110322331-2220302023003201-3012111002301110-0111013231221032-1311310333220110-1123230331332102) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0103301112020110-2133210101322130-3120321133113130-1213031012300313-1233031101032300-2023001121123303-0030103230013311-1223211120321312) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1231132003213230-2312112323103213-1303032131103311-1020022230023332-2001001012010113-3201030003313213-2101020123120220-0123312003101312) |
| `port_ranges` | [port_ranges](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3010332213102001-1023232101230303-2030133221032011-3122131332030021-3200103113222130-2101131210332021-3322113020113002-3233031311221002) |
| `retract_cluster` | [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1131023233233002-0323313303102032-2022112113312010-0331130133000123-0003323331110103-0100213211012000-1331100102121003-1302332022231233) |
| `service_policies_from_namespace` | [service_policies_from_namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3301033031210010-2323133031222031-3022302003130302-0332310330103322-1101123130223311-2303021212131110-1301210012202313-1202323210102000) |
| `sni` | [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2312220212030333-0230231220111231-0012033022321102-0311002130002223-2112002002110030-0212103020302010-2021321302012013-2221101100130002) |
| `tcp` | [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1233311120231210-2331220321101311-1220130030202313-0102210322110321-2301130120100010-1311032331121332-3330332030322100-0111113320312013) |
| `tls_tcp` | [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1212102131222122-0010230212231023-2032230231022002-3301321213331102-1011221101033321-2331201103200201-1332121102010012-1210120021232232) |
| `tls_tcp.tls_cert_params` | [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3132223110233100-1320012202202033-1132100030231233-0033102321001200-2232203103300203-1033211111333321-0331200012002110-0312303033331120) |
| `tls_tcp.tls_cert_params.certificates` | [tls_tcp.tls_cert_params.certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2131031203113123-2120023231133210-2031321322203322-3020220130231101-3232323000013013-2130030221033213-2333302213132323-0211123023031222) |
| `tls_tcp.tls_cert_params.certificates.name` | [tls_tcp.tls_cert_params.certificates.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0330213130033302-2311133212003232-1231011202220331-0201310012311002-1302323032110110-2011103103003101-1102100200021002-2331331000332311) |
| `tls_tcp.tls_cert_params.certificates.namespace` | [tls_tcp.tls_cert_params.certificates.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2012221313301311-3121010301222010-1303302113013130-0321122033330101-3323232033310200-3202100230201211-1230222131023332-3303222101121120) |
| `tls_tcp.tls_cert_params.certificates.tenant` | [tls_tcp.tls_cert_params.certificates.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2032333032103301-1310203201210310-2213221300031023-1333112212202311-0011122300022032-1113001303130332-0232330033123311-3123100102033320) |
| `tls_tcp.tls_cert_params.no_mtls` | [tls_tcp.tls_cert_params.no_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1131332332222222-2211033122302333-1310003133220203-1113330202203211-2202111030333333-3331211322133001-3130032120120221-3223110203103302) |
| `tls_tcp.tls_cert_params.tls_config` | [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1121001013333021-1103303101321122-1231003331210113-2031010221010210-1313221103013223-1111313301230202-0232031132211130-3311300223310031) |
| `tls_tcp.tls_cert_params.tls_config.custom_security` | [tls_tcp.tls_cert_params.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1200022311111031-1221101303010133-3200312012113323-1202310320303130-2211212320213322-1310213121233320-2130002103122320-2103013221013303) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2312011321003020-2120011003222131-2230031122022211-1022020022033303-0222131301032332-1312210331323223-3000101212122300-0203013312220020) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.max_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.max_version](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2031121013232310-2212011011310221-1020112310311120-3300333311021221-2333120332310220-3111122010013033-2321113300311023-0323233031102210) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.min_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.min_version](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1013030331321021-2120321112101012-3022010120221302-0111132203322312-2311132031001112-3313111301120021-2033233033101010-2300022301323211) |
| `tls_tcp.tls_cert_params.tls_config.default_security` | [tls_tcp.tls_cert_params.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3000210203333221-0332202020313002-0322130023111302-2130322221231110-2231122122020131-3232223200011221-1130330030223203-2213333332001121) |
| `tls_tcp.tls_cert_params.tls_config.low_security` | [tls_tcp.tls_cert_params.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0030013133102302-3111331100103230-0212131123223001-2130232010030120-1320021111111002-3020012013103223-0313320310121232-2011202330222023) |
| `tls_tcp.tls_cert_params.tls_config.medium_security` | [tls_tcp.tls_cert_params.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1012320000010113-2333030302211220-2313132232130223-2130011220132223-2031110103133103-2201032313000012-1122201103311331-3221213012232132) |
| `tls_tcp.tls_cert_params.use_mtls` | [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3012321301330313-1130000121202210-3132002213121330-2101110133000102-2221110132220130-0130220212232021-2133131310131330-1221300011113332) |
| `tls_tcp.tls_cert_params.use_mtls.client_certificate_optional` | [tls_tcp.tls_cert_params.use_mtls.client_certificate_optional](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3203000003301102-3110322001303011-0132221021312331-3013220331211213-1202220320222030-1010111030021230-1110231202231202-0100301210001233) |
| `tls_tcp.tls_cert_params.use_mtls.crl` | [tls_tcp.tls_cert_params.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1021021310222211-1220302113221211-0213213022332030-2020103012212202-0132311010011202-3003023002121221-3032010100221101-3230133203310033) |
| `tls_tcp.tls_cert_params.use_mtls.crl.name` | [tls_tcp.tls_cert_params.use_mtls.crl.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0013221000300021-1310113131223003-0112323010100101-2200312132322003-2033331121101121-0220013131002030-1101310130133101-0110011112320012) |
| `tls_tcp.tls_cert_params.use_mtls.crl.namespace` | [tls_tcp.tls_cert_params.use_mtls.crl.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0113021201200122-2133120112123120-2111223032133301-3231032221031333-1102132313013120-1013300102301030-1233013331000322-2031121002111330) |
| `tls_tcp.tls_cert_params.use_mtls.crl.tenant` | [tls_tcp.tls_cert_params.use_mtls.crl.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3301003122110202-2120301021221302-2232110131313300-2223321310230010-3113330233022230-3021210113132210-2123000333232012-3022031102230113) |
| `tls_tcp.tls_cert_params.use_mtls.no_crl` | [tls_tcp.tls_cert_params.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1313221332320211-1011200022102330-0322213303201222-0023330132023302-0212301222303112-3101213100332003-1001222201313113-2312020232003130) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3100312211000133-0002100201332030-0312320012202300-1320103211122023-3122332010003131-2332321030001222-1122133120330133-2313011031120223) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.name` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3031030211121023-0330302030001022-3133223320233311-0122020313300322-2031210230323313-2211113133102110-0202223023133120-1321330133012103) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2212121103233230-3211133303132030-3300213010321023-2021211120120203-2111120220001121-2030320203320010-1321330002233010-1023011322120130) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2210122211323312-3012020130321013-0233130102203210-2030013133202311-2321230221121310-1321012003132212-2210132002133330-1300121323132302) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca_url` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca_url](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1232023003032313-1211121000333331-1101023310312120-2222112113233132-1030000122122211-0002200212103033-3103001212020322-0330231220022331) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_disabled` | [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2012330212030210-1130211020301313-1030130013333121-0003003012210021-0023313032110211-3112133000211232-1030103033122100-2031121123003333) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3113333131301233-1012113033202310-1200303321122130-0311033223320133-0300131110102203-0133233000013023-0112032032022021-3132023202202111) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0223223130003202-3320131021110302-2323220000220122-1013112221003111-1100223112320022-1322233300031301-1132133221302111-3002032130322232) |
| `tls_tcp.tls_parameters` | [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2121021311101012-2230010312001123-2312121013311221-3322320113212022-1023322132223233-0233320112310313-2000233212123201-3101323233030001) |
| `tls_tcp.tls_parameters.no_mtls` | [tls_tcp.tls_parameters.no_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3001133123312002-1000320221230212-0222121000310303-1200233122223222-2230332032132333-3313210001031231-1313031322200131-1021111110011000) |
| `tls_tcp.tls_parameters.tls_certificates` | [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1012102230213233-3212321103023301-1031112200203333-0003130113203002-0113233211132130-2100101021100331-2303203302000310-2130011002231231) |
| `tls_tcp.tls_parameters.tls_certificates.certificate_url` | [tls_tcp.tls_parameters.tls_certificates.certificate_url](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1100213020121312-1102013313211022-2021013123233101-1321231012321020-0320132023320322-1303211332030232-2331203210331133-0332202232232121) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3333312302220231-2222020031203330-1203200113121021-2201210323030201-3033223233113213-2302123310131210-1213103311101220-1000102233110122) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0201303311112033-2123130100113320-3331110213002110-1030033210221302-0331221313203021-1322222013121220-3312222013132233-1022000102330003) |
| `tls_tcp.tls_parameters.tls_certificates.description_spec` | [tls_tcp.tls_parameters.tls_certificates.description_spec](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2230333001233120-2021110211021230-1132332320103130-1112110302310001-2322212112121030-3310320120310020-1331013233132303-2210320312222033) |
| `tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling` | [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1011133323230033-0011110010130221-1323003103012012-0103311200301030-1130220220222033-2103111221122030-1320001303310010-2021201132022322) |
| `tls_tcp.tls_parameters.tls_certificates.private_key` | [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1101331003131213-1231312032211322-2233011011231000-1002211131023110-1022021311003332-0223300202331220-3102031301321103-2220022111213132) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3320320232322331-0001011010110300-2033233300111122-0013003310303200-2323210133200202-3030100013323002-0021300202230233-1122010002233112) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3033010312023320-3132320200102133-2303231033212110-0012020102201103-0120213302220303-2211210332001202-1301232302221232-3221221323301100) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0133330130022222-0021302211130321-3023230232220112-0110000232322331-2201321323012013-3103103132211023-1021332230220130-0310013333331303) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2023322332202320-2213301030232110-0302133102330320-3230222011123313-0120333221212013-2003103222201320-1320131031202012-1031323022221231) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1312210322322010-2311302023001232-0032233233302100-1121300131032223-0232313303000000-0132231331200033-2310111021202313-3011303222210213) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3003112000123322-1131112122120313-1113332232201322-2032211033000333-3313120132010033-3130233333012003-1113033130331320-1011201223003101) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1023030133133323-2302200000330013-3220112010000021-1221303222033010-2110120301010220-0131220221211213-0313101120011321-0003322102212211) |
| `tls_tcp.tls_parameters.tls_certificates.use_system_defaults` | [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0133232121100203-3112131312222332-1210211201023203-0213031301113222-1100110303101231-1002321023220200-2130320023203200-2102003202313001) |
| `tls_tcp.tls_parameters.tls_config` | [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2221303331133203-1130203323222012-3330132321032001-2221230001200013-3211320212313313-1330022310200131-2012220011310103-1123313133101101) |
| `tls_tcp.tls_parameters.tls_config.custom_security` | [tls_tcp.tls_parameters.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3013101131313113-0201013023331222-3233331012300022-1301301333232231-3312223021100102-0312131001221132-2103123100122010-0102133012331132) |
| `tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3100112232003232-2300221211320311-2021232303000202-0300221100013232-0311030213023021-2013322320033113-3302320120122101-1310220133233300) |
| `tls_tcp.tls_parameters.tls_config.custom_security.max_version` | [tls_tcp.tls_parameters.tls_config.custom_security.max_version](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0300113000203223-2220033311213120-3310333103003013-0301032103312203-0011301303010303-2001211122103231-0111103021111010-2210022200322330) |
| `tls_tcp.tls_parameters.tls_config.custom_security.min_version` | [tls_tcp.tls_parameters.tls_config.custom_security.min_version](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3331100231131101-2031232201123000-1032022131120020-2331330012113233-2331102112333131-2202232211011330-1202011102333313-1110203132310202) |
| `tls_tcp.tls_parameters.tls_config.default_security` | [tls_tcp.tls_parameters.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1332022222321223-2110230113333020-1021123232000223-3223023231122321-0301011020323100-3232031023123112-2223012332033230-2012203201302001) |
| `tls_tcp.tls_parameters.tls_config.low_security` | [tls_tcp.tls_parameters.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1233020333030001-0023231023111002-2320211210113132-0200321030010330-0222301121001232-2031100331121301-0100303211321012-3311212012121130) |
| `tls_tcp.tls_parameters.tls_config.medium_security` | [tls_tcp.tls_parameters.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3001103131311223-3010020313031202-3200003032201111-0002112323322103-3123310233032000-3131111123113132-3102113330310000-3231122221311101) |
| `tls_tcp.tls_parameters.use_mtls` | [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3223222113230022-2123203033302010-3201002232021200-2112303222332011-1013320230313022-2100203313201023-0320013033033023-2203133121123123) |
| `tls_tcp.tls_parameters.use_mtls.client_certificate_optional` | [tls_tcp.tls_parameters.use_mtls.client_certificate_optional](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0310032321121022-1131300132301122-0101302232113203-1330002133211013-1022010011303221-1310002003321301-2013330120233011-2331230323113310) |
| `tls_tcp.tls_parameters.use_mtls.crl` | [tls_tcp.tls_parameters.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0100200111001303-3031030302320223-0210132001100020-0232033032122220-2023302130333033-3322212220301321-2230201112022322-1103221030212303) |
| `tls_tcp.tls_parameters.use_mtls.crl.name` | [tls_tcp.tls_parameters.use_mtls.crl.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1112302322302231-0333230320023333-1313222210010222-3020100210320032-0212133121233013-1202111100012021-2121001322020311-0033131210003102) |
| `tls_tcp.tls_parameters.use_mtls.crl.namespace` | [tls_tcp.tls_parameters.use_mtls.crl.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0101102323230331-1033111200123212-0313133311121003-3302110203330022-2312030020033123-2013312213303022-1332232331231313-0231323130312312) |
| `tls_tcp.tls_parameters.use_mtls.crl.tenant` | [tls_tcp.tls_parameters.use_mtls.crl.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3121333021032331-0112221131203000-2230102322330110-1212202010203110-3312012231213313-2033322210322331-3323302330310322-1301003203021132) |
| `tls_tcp.tls_parameters.use_mtls.no_crl` | [tls_tcp.tls_parameters.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0332112123301220-2310300121111102-2203101012332220-2313031111331131-0001313301002011-1101331312010313-2130222002312123-1233012313332223) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca` | [tls_tcp.tls_parameters.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1310022002220230-3301211110111003-0120122110301201-2213301110010223-1231112203210021-3301302011233230-3001200111102203-2023313200231030) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.name` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3001302011113030-0232101312023210-1101333020331021-0000310222332101-1030013002322101-3333301122301213-1011203202223312-1101221023101233) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3012232013130300-1211110120202010-1033320132320030-1221001200130331-2212033003210322-0333131030330203-3003113322003101-1301032022120211) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3102121312210012-3230223000230330-0002200132121220-3303012000200202-2231302212012211-3202203102303302-1112112331221013-2332100132203313) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca_url` | [tls_tcp.tls_parameters.use_mtls.trusted_ca_url](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3011301223002313-1302132132023313-0330120020322202-3032302202111112-0010320111013303-3200332330210301-2302112231221303-2112120130322120) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_disabled` | [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3120300001022020-3010321310302003-1230131212331113-2120233033020121-3202321332133232-0122331131222200-3313301220303133-2033330230300020) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options` | [tls_tcp.tls_parameters.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1020130210233233-3331300223301103-0110222032000010-1030001032212101-3020221332323101-1230201030332102-1112303223213101-2100020221011223) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2203003223223132-0103012231300230-2103111213131010-0223322310201003-0300102321200222-1122132202030210-0230003011110302-1112111023231330) |
| `tls_tcp_auto_cert` | [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3122010013013333-0010120031233203-2230121020320021-2303221102121132-3030013013032022-1321333312130213-1313113220130030-1002023233330000) |
| `tls_tcp_auto_cert.no_mtls` | [tls_tcp_auto_cert.no_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1312323322322321-1230033311202000-0021330230212001-2121300030011302-1310011212312321-2321022223213330-2022210331101330-2232332123301233) |
| `tls_tcp_auto_cert.tls_config` | [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2100223333301213-2002332220230023-1111133020100000-1303221120200120-2202022230303211-1102203132031311-1102301231123233-0210203300120303) |
| `tls_tcp_auto_cert.tls_config.custom_security` | [tls_tcp_auto_cert.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3113213220330012-3210103010123110-2322302223233001-1132212132131022-3131220221333211-1113012020101213-1202012113211110-0121032132032023) |
| `tls_tcp_auto_cert.tls_config.custom_security.cipher_suites` | [tls_tcp_auto_cert.tls_config.custom_security.cipher_suites](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0023332303002130-3332030023122220-1331331200202023-2321231133130303-1201010012203210-3332123211201113-1333112202311322-1003113331202113) |
| `tls_tcp_auto_cert.tls_config.custom_security.max_version` | [tls_tcp_auto_cert.tls_config.custom_security.max_version](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1310130000230122-0200213222323231-1211023301310201-1022132332202000-0020112121321330-1012110323020233-0121212220030322-0000310203212233) |
| `tls_tcp_auto_cert.tls_config.custom_security.min_version` | [tls_tcp_auto_cert.tls_config.custom_security.min_version](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2133221010110323-1020302201323201-3310131311130010-0230001321100232-0111101000131201-3221131320303221-1012000032112322-0303123110202021) |
| `tls_tcp_auto_cert.tls_config.default_security` | [tls_tcp_auto_cert.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1231123211121221-3221333310333310-2312011201121020-2110201113100321-1300001133302122-1312233121120032-2301330030010010-3000123120323211) |
| `tls_tcp_auto_cert.tls_config.low_security` | [tls_tcp_auto_cert.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0111030230322333-3323310102022203-3323320111212030-3132222030022222-2333002112200212-3232012221322001-1113202311300031-2333231131011010) |
| `tls_tcp_auto_cert.tls_config.medium_security` | [tls_tcp_auto_cert.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1303003213322311-0201122021101221-1302030000202223-1013001323232310-2300223002320112-1030311322130322-0320202103323000-0132020320332123) |
| `tls_tcp_auto_cert.use_mtls` | [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2000030223230202-2220210002210222-3002220323331312-3303332212110220-2032210011113200-0101313001313133-0222032001300033-0012223200133211) |
| `tls_tcp_auto_cert.use_mtls.client_certificate_optional` | [tls_tcp_auto_cert.use_mtls.client_certificate_optional](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1310222003001333-0201010230220131-3213212222221213-0232301131302131-2213002223313320-3130200000013213-0222123101011210-1102010112200110) |
| `tls_tcp_auto_cert.use_mtls.crl` | [tls_tcp_auto_cert.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2023121130131103-1130230220332001-1103011133232103-3102131201233302-3312020120132220-0000322010023032-2121311233322131-2231233333002110) |
| `tls_tcp_auto_cert.use_mtls.crl.name` | [tls_tcp_auto_cert.use_mtls.crl.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1121101213322130-1221310330020213-2101321311010310-3200322332131121-3311132212223030-0031122102303322-2222331010330031-1120332330003221) |
| `tls_tcp_auto_cert.use_mtls.crl.namespace` | [tls_tcp_auto_cert.use_mtls.crl.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2033021222131122-1230303230120112-3202202133121122-2330102122020101-1231232110303210-0030000310311322-1101103122110011-2201000011321122) |
| `tls_tcp_auto_cert.use_mtls.crl.tenant` | [tls_tcp_auto_cert.use_mtls.crl.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1020322103321221-3220301113003332-0333032132321002-3132112101201021-3312101231121322-1110002012321323-3030110131323203-3233131331032123) |
| `tls_tcp_auto_cert.use_mtls.no_crl` | [tls_tcp_auto_cert.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2310311132122032-0233210101310000-0222220321230000-3113133333012201-0102321222031030-0000100020332131-0323131101032031-2331003112313103) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca` | [tls_tcp_auto_cert.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-0212132103311331-2230222302111013-3331301031021130-0230332012032302-0130203300223022-3220230331220113-3211323102100202-2002312331003211) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.name` | [tls_tcp_auto_cert.use_mtls.trusted_ca.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2302212322000212-3323020331121211-2230020322321131-2133000212233221-3133101011132222-3113021123013021-2332012010300321-0100121033133033) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.namespace` | [tls_tcp_auto_cert.use_mtls.trusted_ca.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2022003210030130-0201311121121312-1030310203231031-1332310201211221-3213131001311212-1231002213110031-1013330311221010-0103033010032120) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.tenant` | [tls_tcp_auto_cert.use_mtls.trusted_ca.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1313330103302210-0333112133213022-1220221203012210-3311311021330130-3233301101022123-3133211211001111-1010211323300311-3001102323013123) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca_url` | [tls_tcp_auto_cert.use_mtls.trusted_ca_url](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3213133123022310-2300013011102123-1103210132220113-3212121330103200-2321021231322121-2323011101302212-0231103102213012-1023210111133332) |
| `tls_tcp_auto_cert.use_mtls.xfcc_disabled` | [tls_tcp_auto_cert.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3232232030103320-3130121211101303-1210132231322022-3013321130230021-0300032021203111-1332012233010021-2033122231021012-1201312103222301) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options` | [tls_tcp_auto_cert.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-2213310110203313-1232313033031131-1332121213033031-0213112212010211-3101012111032012-0312211230202300-2203233212200002-3212230100213220) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-1033233113222132-2323230213321012-2100311332211000-0303232301222200-2303311211222023-1121333223220012-1122311023322032-1232322210203201) |

<a id="canonical-1011022010101003-1111012300021023-0113013102330111-3011230222012201-0130101131130322-2013322311103222-1122320332303032-2001122100033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- active_service_policies

<a id="canonical-3120023000002123-2201220122222212-0000023001323002-0032102201321122-3201322001220310-3012222002320012-0331211022313201-2312123012011311"></a>

Type: `"single"`. Computed.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Additional upstream details:

List of service policies.

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

- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3120023000002123-2201220122222212-0000023001323002-0032102201321122-3201322001220310-3012222002320012-0331211022313201-2312123012011311)
- [no_service_policies](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0121112133023001-2133103202133131-1030333112200300-0221311100103201-3330222002023012-0321002320113000-0310301333203033-2200011113112222)
- [service_policies_from_namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3301033031210010-2323133031222031-3022302003130302-0332310330103322-1101123130223311-2303021212131110-1301210012202313-1202323210102000)

Select alternatives according to the provider validators above.

<a id="canonical-2101222103332201-1323012132211303-2303110203113233-2001000320230311-0021122113220021-2232301001203003-3101003333001000-3023203021312213"></a>

### Direct properties for `active_service_policies`

- [policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0213200311222000-0330011211322200-3302233331211022-2300221212033003-2230103013012113-2330311020220032-0203100103320300-2123131012123113): complete subsection reference.

<a id="canonical-0213200311222000-0330011211322200-3302233331211022-2300221212033003-2230103013012113-2330311020220032-0203100103320300-2123131012123113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies.policies` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1011022010101003-1111012300021023-0113013102330111-3011230222012201-0130101131130322-2013322311103222-1122320332303032-2001122100033100)
- active_service_policies.policies

<a id="canonical-3022103213023223-2031203303201102-0312223311130020-0331000121223103-2022322210132200-2332322112223120-0312332133213333-3310031210322111"></a>

Type: `"list"`. Computed.

Service Policies is a sequential engine where policies (and rules within the policy) are evaluated
one after the other. It's important to define the correct order (policies evaluated from top to
bottom in the list) for service policies, to GET the intended result. For each request, its
characteristics are evaluated based on the match criteria in each service policy starting at the
top. If there is a match in the current policy, then the policy takes effect, and no more policies
are evaluated. Otherwise, the next policy is evaluated. If all policies are evaluated and none
match, then the request will be denied by default.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0323001300101032-2320301023302303-0231200202123212-2003332130221121-3100331302330100-2030210111103231-1310011312030222-0022222003303010"></a>

### Direct properties for `active_service_policies.policies`

<a id="canonical-3230231100112011-0310031033230101-2231101031321333-3220002113223022-2201123303312011-0120022203030002-2203311321321211-2110120223333322"></a>

#### `active_service_policies.policies.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3313200212201121-3301301003300102-0321121121103002-0222003123233201-3021331300210312-2113011030323130-0132031310231313-1111002303221001"></a>

<a id="canonical-3330120213023001-2123332022221110-2201120212331122-1031102221231120-0210330300110202-2300323002123003-1302021203132232-3303133300200123"></a>

#### `active_service_policies.policies.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1203021322211230-3321330302211101-2001320231113230-1010131003013111-3032100121331312-0221321232220201-1331001101120201-3001310230022210"></a>

<a id="canonical-2331212110310332-0013202021202232-2130211223010231-0210003021021011-0120201211233022-3300133211033233-2212012010312323-1033210032323330"></a>

#### `active_service_policies.policies.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- advertise_custom

<a id="canonical-1013012233212133-3023131131211102-2203001220030233-2131211130201223-1323322001311122-0223103302132103-2032132122323131-2030201131222013"></a>

Type: `"single"`. Computed.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

Additional upstream details:

This defines a way to advertise a VIP on specific sites.

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

- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1013012233212133-3023131131211102-2203001220030233-2131211130201223-1323322001311122-0223103302132103-2032132122323131-2030201131222013)
- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2200031202211021-0313200202312233-1013023211320030-3123000101033002-1113110302231111-2101321210002322-3331003310213213-0312330021311122)
- [advertise_on_public_default_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1123330002330211-1310302201203201-1113023310221132-0301203233322302-0032103302032320-2010301030313101-2213001103002101-3232220103002213)
- [do_not_advertise](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0001303000210302-2210122211320302-0002200332110310-0310113320302002-1122011331112123-3223311220302121-3230321301031211-3123012131133000)

Select alternatives according to the provider validators above.

<a id="canonical-1121021002130000-3113302220030221-1233221203031212-0212213211131202-2003102033102112-1121221000313300-2323211131213213-3330202300130020"></a>

### Direct properties for `advertise_custom`

- [advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310): complete subsection reference.

<a id="canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- advertise_custom.advertise_where

<a id="canonical-3310231133113102-1232110011013322-0311030213330120-3103111122222101-2012231121021022-0032231320131032-0312033301101102-3311333210023121"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2102232130200103-2200202021012003-0021130331132233-2013111130322300-2203322220123003-2021100021012321-1202022101101132-2212223303220023"></a>

### Direct properties for `advertise_custom.advertise_where`

- [advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0331101202330300-3013302221110011-0223120300222212-2201311202100110-1200210333031121-3113033021103211-1222032331013330-3121013002120201): complete subsection reference.

- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3133222223020333-2012330001010231-0113121311102003-3022122103102003-1112322121313133-3300023132223312-1021121221030110-3133333310303301): complete subsection reference.

- [advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1012303003200121-3010330321203201-0223111121101321-1302123210213020-0202133103201322-2001320210021000-0311023312133103-3311212213023213): complete subsection reference.

<a id="canonical-1120212332121220-2000023032123302-2323322112122211-3130332122121130-3302132312202023-2313322213301201-1002221123230311-3111131021301330"></a>

<a id="canonical-0203302313102032-1113313313011101-0012301213221300-2331302301330001-2222333133332120-2310212300302011-3330013001132211-0233210012232330"></a>

#### `advertise_custom.advertise_where.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0110130103100310-2112310323110230-2033122033131201-2021321013202220-1321132213001302-0212123010031321-0320310012213212-1101332223000132"></a>

<a id="canonical-1130211312021000-3211133322333333-2313321032323301-1032210030102122-1121200010232223-1302320301200221-3203221302130023-2201003200130331"></a>

#### `advertise_custom.advertise_where.port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1201111030300103-0132333121123310-2133301110203021-3233002311120021-1232002300111300-0102111011200322-2303020001310101-2131101312301130): complete subsection reference.

- [use_default_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3023311201032023-2113010112030331-3112210232133220-1333113123311133-2330321201130222-1122322012201230-3302100102301001-0003201111323231): complete subsection reference.

- [virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3212203202033003-0230033003330111-2203020032311021-0010323231331320-3020200213200112-2123312323032030-3203020021133103-0120320132001200): complete subsection reference.

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0122113233033122-1212031131202012-1003210101121201-1330102202322232-2130220221202212-3003100212222003-3213102310020231-1101311333121000): complete subsection reference.

- [virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0310233013030220-3111100223222230-0312202122200212-3113110332121233-0221220033233110-1300030222302200-1321121120220012-0030120222320100): complete subsection reference.

- [vk8s_service](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0132122101013332-2002300333102030-3301212123122331-3230111111010301-2012311030311102-1231333202212000-0001103133211311-0101301312120333): complete subsection reference.

<a id="canonical-0331101202330300-3013302221110011-0223120300222212-2201311202100110-1200210333031121-3113033021103211-1222032331013330-3121013002120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2231230123023333-1303233212120200-0210210311012300-3313013313132012-3003201213012012-0310130001320333-1010310122120111-0230001010001333"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3213302033211122-2312312302003312-0223320311321320-0112333133230233-3232001332101112-2110123011123132-3330231233210201-1220022031201010"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0203113320000023-0331123121133130-3231022310020211-0231012212333002-2321113002210112-2202011011322210-1210230130113311-2110321311011323): complete subsection reference.

<a id="canonical-0203113320000023-0331123121133130-3231022310020211-0231012212333002-2321113002210112-2202011011322210-1210230130113311-2110321311011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0331101202330300-3013302221110011-0223120300222212-2201311202100110-1200210333031121-3113033021103211-1222032331013330-3121013002120201)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-0221030113121220-1102213003000323-2021322112321222-2000323032321321-3012233203131003-0233310202112013-1323020331223233-2212230030221021"></a>

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

<a id="canonical-0032300131033131-1021203123302000-1332001023211220-3213023122332230-1222022022323330-1331123031020223-2012330313333033-1323130203300020"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-3200310213312303-2302011230111300-1133010110011100-3331022000202010-1213010102321322-0220021120030320-1322213000203113-1303123203310020"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2020100023101320-1123111211221332-1033313032130223-0212011113130233-0303232310323100-3221312202302110-2100012030112310-0320203333130030"></a>

<a id="canonical-0312332303321311-2123120110202300-0302123012120311-0010301300110130-2002130300222113-1011111023103010-2010113201101211-1231130203103030"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3112330022132120-0000002110302311-2302311113112103-0202200331112102-1023030110002121-0023023220331010-2031212110111202-1012002222213113"></a>

<a id="canonical-3333201202220231-0312133201322222-0121030011132212-0331003211113211-0300013223200211-2002330033310003-0333000102333220-1123220321100020"></a>

#### `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3133222223020333-2012330001010231-0113121311102003-3022122103102003-1112322121313133-3300023132223312-1021121221030110-3133333310303301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-3002313330003213-0333313000001100-1002121200231321-1102102212030113-1312322022010320-1302103013202203-2031312200100201-0031310332232300"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3311130133002203-3121012231130301-2000232133011203-2031210213203032-2211011212120332-3320320223201333-0100122320230001-3233213022222230"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public`

- [public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2221332030202103-1210310322012230-1323232202310111-0110000111213100-2212113010032012-2033213313201223-1230033030110031-1313310102321212): complete subsection reference.

<a id="canonical-2221332030202103-1210310322012230-1323232202310111-0110000111213100-2212113010032012-2033213313201223-1230033030110031-1313310102321212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3133222223020333-2012330001010231-0113121311102003-3022122103102003-1112322121313133-3300023132223312-1021121221030110-3133333310303301)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-3303330230303111-3203201002010201-1223221132102301-2131200021210130-0132203222102011-0230113201202311-3033023022221311-1312322323120331"></a>

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

<a id="canonical-2112221033032112-0310300120103020-3331132223322232-3032030100101313-3213131013331102-1211323333103032-1301220123030022-0232132031323112"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-0123113321121212-1133200130030011-1132322122332021-3211312311333331-0211220320032121-1022301302221311-0113200033321212-2021010030011111"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1310200002013312-0332123320003323-3312212023100100-0012133010222311-1002020020010212-1203211133130113-2222001022301321-1300123102213011"></a>

<a id="canonical-2131020113012313-0101003230221200-0021200121222332-2331131100011001-0030320203200313-2013021311023333-1032202112011223-0213131212310003"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0100023020332332-3111211211111332-0220012221200023-3221332312223303-2331332212231322-1000230112220121-0201123120332310-2110000232333320"></a>

<a id="canonical-2000321203023200-2113103230223120-0310203120321302-2311232120003301-0220320010003232-1321203032021000-2220320323102002-1222001211332212"></a>

#### `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1012303003200121-3010330321203201-0223111121101321-1302123210213020-0202133103201322-2001320210021000-0311023312133103-3311212213023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-2331013112211123-3210122222111103-0110211303020330-1230222120220312-0101013212301320-2333001123212200-0322112323231123-0003011131100211"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2310002121001132-0012001310223323-2221130001222311-3320302300130311-2210000200101131-2132210301030022-3102203313200202-1213103203011022"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3101102310000321-2200302000130112-2112322032201013-3230133000023231-1220323221221201-0203021120030012-3100032003102200-3213030203020310): complete subsection reference.

<a id="canonical-3101102310000321-2200302000130112-2112322032201013-3230133000023231-1220323221221201-0203021120030012-3100032003102200-3213030203020310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1012303003200121-3010330321203201-0223111121101321-1302123210213020-0202133103201322-2001320210021000-0311023312133103-3311212213023213)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-3103221033000113-3031002313133213-2330003231021330-2102203112311221-0130223013111131-3333122333213330-0000231310230121-3023032232200312"></a>

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

<a id="canonical-1020102121310232-0123311010132010-2212300313322131-3230231313231200-1211132120033331-2203230003322022-3213232313302113-3333003202230202"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-0021211230030301-0203030030210030-2011300200303230-2110030121312233-2113121201131132-0112222322020333-1233301301313000-0003010221201101"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1020023313211113-1112010333313332-1111122201312111-3231131301031032-1013000132121021-0013123302323120-1123022332311121-0031012222033131"></a>

<a id="canonical-1313303230002122-0100300113111223-2320203113020100-3122010203312133-2102003103323121-0323030213201111-0311012200132113-0301031313201222"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3233110013333330-1010302023332230-2220222012013210-3103013030012220-0322120132301233-1303231130020303-0031022320112111-1200121001113313"></a>

<a id="canonical-2002300210102023-1030113320011021-1220011032301303-2131230200203233-0211233213011023-0203131200200111-2003221232211233-2233010001131230"></a>

#### `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1201111030300103-0132333121123310-2133301110203021-3233002311120021-1232002300111300-0102111011200322-2303020001310101-2131101312301130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.site

<a id="canonical-3120002133123221-1023122202112031-0211321000303220-1301321133200322-0100030102133200-1200201023233302-2311113132230311-3030211111123130"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3333210232010213-2232000301123022-3030223003332223-1313223030201202-2133303030220332-0302212001332212-3011010130112133-1323301013123333"></a>

### Direct properties for `advertise_custom.advertise_where.site`

<a id="canonical-0212033211312200-3012123331133331-1320302232322023-0112002202333332-3301013120333131-2103320123201023-0122101103123001-2222223321330323"></a>

#### `advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1132310132123321-2332200002130112-1330023011031022-3221203310003311-0102321022310032-1233030030000333-1112310200332213-1130330303103222"></a>

<a id="canonical-0310313003020301-3110133323023213-1221130321103321-0312133112300113-0111333203202220-0013231310203102-0133101221112320-0200212012322101"></a>

#### `advertise_custom.advertise_where.site.network` property

Type: `"string"`. Computed.

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

- [site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1323031121031322-0120103001233310-2222032133123001-1002203101133221-1310131320311322-1223011003021101-2313323133233230-2010322230103012): complete subsection reference.

<a id="canonical-1323031121031322-0120103001233310-2222032133123001-1002203101133221-1310131320311322-1223011003021101-2313323133233230-2010322230103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1201111030300103-0132333121123310-2133301110203021-3233002311120021-1232002300111300-0102111011200322-2303020001310101-2131101312301130)
- advertise_custom.advertise_where.site.site

<a id="canonical-1123013230331303-0111100002013203-2311002023111001-1211023131203300-3130231122301330-0101012331200301-3333123100311031-2212110022212003"></a>

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

<a id="canonical-3120313311032020-2332310321133103-1210011210302032-0230321202233011-1331033103000023-1320100313120220-0120121312303213-0101103101010321"></a>

### Direct properties for `advertise_custom.advertise_where.site.site`

<a id="canonical-1220121123303102-3323230312213021-1030120002210112-2331331111102230-1010312323101210-1300021023312231-2012202111020012-0022101233222221"></a>

#### `advertise_custom.advertise_where.site.site.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1323300123212001-2313312220022331-1012230302010111-2022310130131111-0333213312010031-2232333123100302-3203231213133110-1333201001201322"></a>

<a id="canonical-2223032023103231-1303320032021213-0112021132221120-2210120133033320-0221011301203110-1233311002210003-0302301233220000-2310221023312002"></a>

#### `advertise_custom.advertise_where.site.site.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1122113220301021-2312111120022013-1321012320032302-2130110110211211-0233312310230002-1221032121102311-3310210100300031-3122212031010212"></a>

<a id="canonical-0111033102202233-0100303112013130-2010000031211121-3012130122213303-0132021113202113-1302223021202210-1120332003033121-3102100223321010"></a>

#### `advertise_custom.advertise_where.site.site.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3023311201032023-2113010112030331-3112210232133220-1333113123311133-2330321201130222-1122322012201230-3302100102301001-0003201111323231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-3000301102201130-2220321212111112-3100312131010120-2310001201030232-3122333332300131-1212100111121112-3121300322212220-3121312130131202"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212203202033003-0230033003330111-2203020032311021-0010323231331320-3020200213200112-2123312323032030-3203020021133103-0120320132001200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-0221023101223332-2003212210330220-1230031330123113-3200221330223130-0203311001103010-1321111011130222-0212003010211130-2000220121020212"></a>

Type: `"single"`. Computed.

Parameters to advertise on a given virtual network.

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

<a id="canonical-3311002110020231-2003301132110133-2003022001132300-1021003211010221-2122303322023121-3100111210223031-3301203011322202-2233111123202310"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1302112110031231-0302000120333011-1032132033003010-1103011210210113-1320332210101301-1321212311123022-2033003300321233-2203203321000013): complete subsection reference.

- [default_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2211111021021033-1232123223130201-2310012212121001-0113020230302301-2130122311330323-3020000120021330-1132121123322223-0111221332023010): complete subsection reference.

<a id="canonical-3133011032003200-0313313230021312-2232021003003131-1021130030333112-3133123303223212-2023101313010112-3003202230122210-0012211022331230"></a>

<a id="canonical-0321100000311301-3021300223312101-3022213313032203-2132311133000010-1130022303323311-0330003212210230-3311333332303121-2213010213000213"></a>

#### `advertise_custom.advertise_where.virtual_network.specific_v6_vip` property

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3222130011021322-3011031312100021-3232233213221313-0303112131223100-1021230332321333-0222123033102130-2200220113302220-1102313333113213"></a>

<a id="canonical-1111033311003122-2233120220201233-2033010030132132-1302012133021331-3203111323111031-0331022031302133-3010111210122303-2313001111321012"></a>

#### `advertise_custom.advertise_where.virtual_network.specific_vip` property

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3013011012031303-0021200133033000-1033333102233322-1200121223223013-0230203220220102-0002313133112021-2311113012212302-2312202000013022): complete subsection reference.

<a id="canonical-1302112110031231-0302000120333011-1032132033003010-1103011210210113-1320332210101301-1321212311123022-2033003300321233-2203203321000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3212203202033003-0230033003330111-2203020032311021-0010323231331320-3020200213200112-2123312323032030-3203020021133103-0120320132001200)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-2211101111223222-1000300201320321-2123121212333231-0300213121300212-1130131110011231-1121300202033002-1021332103102112-2220122323301332"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211111021021033-1232123223130201-2310012212121001-0113020230302301-2130122311330323-3020000120021330-1132121123322223-0111221332023010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3212203202033003-0230033003330111-2203020032311021-0010323231331320-3020200213200112-2123312323032030-3203020021133103-0120320132001200)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-0012012310122300-3113320200233232-1330102222111230-1111120103333233-3130111130133332-0103120001201133-3112033032100211-3303110330031302"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013011012031303-0021200133033000-1033333102233322-1200121223223013-0230203220220102-0002313133112021-2311113012212302-2312202000013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.virtual_network` properties

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-3333120121311030-1222210233221001-3312122111101220-3230032211310111-0320230000212310-0112221113320201-0312111221101202-2321212213311233)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0103323311003132-3301302203223221-0203231010322201-0313110133112213-0110221020311100-2200213031313001-2012020103110231-2011020211202023)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1103122122200012-0323032231201022-3322102311013312-2202220203012213-1202320222020200-1321110221030333-1103121003312213-2312332031200000)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2222113330023031-1000033310231130-1333330322223101-3130132231010311-0313302100232021-1100200200331321-2230312002301121-1031121030322310)
- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3212203202033003-0230033003330111-2203020032311021-0010323231331320-3020200213200112-2123312323032030-3203020021133103-0120320132001200)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-0011003031001313-1032323211101023-3100100010031132-3012010332230231-2120130312130020-2203023031312023-1112320233212332-3122103033223012"></a>

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

<a id="canonical-1113032211221321-0200230101311230-1031033032231222-1310220030331122-0221312000203000-2220123222112323-2001203101131203-1220320100101211"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network.virtual_network`

<a id="canonical-2033323033031232-0323103233210101-3303000102230312-3310111213330330-2030001001221200-0100000033122233-3330313321030101-2132330222030301"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0300233202133223-0000300210201322-1122132013030330-0032100330013020-2131320110333123-2032110111310203-1201112013013302-1302012203223110"></a>
