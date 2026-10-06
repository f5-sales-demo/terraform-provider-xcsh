---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1030030122011102-1200203100110303-0222320330313000-1212333211001012-3011132203133320-3203212303220322-2213321232120023-2033013323323230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- active_service_policies

<a id="canonical-2103221110303112-0033133102200020-1121233332111101-3123220121032033-0201023223331112-2100113020011131-0332232302102333-2201210210123120"></a>

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

- [active_service_policies](data-sources--http_loadbalancer--reference--group-004.md#canonical-2103221110303112-0033133102200020-1121233332111101-3123220121032033-0201023223331112-2100113020011131-0332232302102333-2201210210123120)
- [no_service_policies](data-sources--http_loadbalancer--reference--group-021.md#canonical-2210312233210311-0002123133331020-1220122121222001-0132233321112131-1121331112303320-2113021232203331-1202011022320022-1211301213031113)
- [service_policies_from_namespace](data-sources--http_loadbalancer--reference--group-026.md#canonical-2210322011133032-1130200230121211-3102212123010321-3312101211130233-3323101010013002-1002011302333221-3302300230113112-1120220333133013)

Select alternatives according to the provider validators above.

<a id="canonical-2202330122022222-3031320302003030-3202132302100102-2013221122212113-0203310012100021-2300202223311110-2100133120132002-0230210233012211"></a>

### Direct properties for `active_service_policies`

- [policies](data-sources--http_loadbalancer--reference--group-004.md#canonical-3220300211201313-3112031121321203-3332020000200002-2231330100013210-1100233003020313-0030022132020223-2111002320310312-2130232003330020): complete subsection reference.

<a id="canonical-3220300211201313-3112031121321203-3332020000200002-2231330100013210-1100233003020313-0030022132020223-2111002320310312-2130232003330020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies.policies` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [active_service_policies](data-sources--http_loadbalancer--reference--group-004.md#canonical-1030030122011102-1200203100110303-0222320330313000-1212333211001012-3011132203133320-3203212303220322-2213321232120023-2033013323323230)
- active_service_policies.policies

<a id="canonical-3322212233321013-3033301320201210-0301113012003003-3123123220003010-0003020313110021-3022002003132230-1110110130323233-2130122231020203"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1211020330021300-2201223033130332-0131021013320112-2132101320221330-1303213302302013-3330332130323013-1300001131330123-1111030003233300"></a>

### Direct properties for `active_service_policies.policies`

<a id="canonical-2011122102111310-2023012332313023-1313313130130211-3103312233000123-2132323321003103-3101023222033321-1101222132300331-1213013213010310"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1222212322121022-2001102131220321-2212300231331330-3131030212010131-0332212030210021-3303123302000220-0222022123000312-1300031202020311"></a>

<a id="canonical-2202231202301311-2033033222333203-0220200222021303-0121333322110110-3012100003120321-0211300310230230-3122222323002232-3132303032102201"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0302321102333311-2010132031003213-1110320100203203-1330220203311021-2332331001002331-2032232103203033-3113202003200233-0300322113122023"></a>

<a id="canonical-2113221332233112-3203330123312100-3313231332322213-1123203212101033-0321231023003333-1310133231230232-1310021120323221-2123320011200200"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- advertise_custom

<a id="canonical-1211322103210032-1313121133121221-0013022312002112-2300231200130220-1202312213222212-0122302321302011-2110030132000203-3233120030201103"></a>

Type: `"single"`. Computed.

\[OneOf: advertise\_custom, advertise\_dualstack\_on\_public, advertise\_on\_public,
advertise\_on\_public\_default\_vip, advertise\_v6\_on\_public, do\_not\_advertise; Default:
advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP on specific sites.

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

- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-1211322103210032-1313121133121221-0013022312002112-2300231200130220-1202312213222212-0122302321302011-2110030132000203-3233120030201103)
- [advertise_dualstack_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-0303330131303232-0032101200321202-2333020312033323-1133203011110300-2013002032133102-1331231230202003-3211301301303202-0022333320122123)
- [advertise_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-0131001033312203-1233133313001030-3313122220112000-3012100022321011-2213000212113331-1311321221112331-3220312211312213-2212231101232121)
- [advertise_on_public_default_vip](data-sources--http_loadbalancer--reference--group-004.md#canonical-3110212023223203-2310311120102130-1211123102111330-3233113030202022-3131033223003130-0200221320210332-1021232103112310-1333131001101322)
- [advertise_v6_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-3321100211232200-0201302220112112-2312220001200002-3310313211332132-3121131101201001-2203323330032033-0220330303001130-1213112220232121)
- [do_not_advertise](data-sources--http_loadbalancer--reference--group-017.md#canonical-0021302010203112-0300133233023002-1233200231022231-0002221120200231-1313312000123101-1111103332203311-2120010322013122-2021313130023120)

Select alternatives according to the provider validators above.

<a id="canonical-0031022001203133-3000021032130220-0320331021131113-2122302022122330-3103112301002330-3122130211301310-2232223031330120-1130210003200132"></a>

### Direct properties for `advertise_custom`

- [advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213): complete subsection reference.

<a id="canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- advertise_custom.advertise_where

<a id="canonical-0320221211312232-3002222312311023-2201022023112331-0321201222003221-0331333321330320-2321112002231221-2103133200122103-0200211130102010"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1003200002010010-2111220012222112-2021333120201222-2033302331310231-3131130003011222-1302300100032023-1030122330131321-3013311011033132"></a>

### Direct properties for `advertise_custom.advertise_where`

- [advertise_dualstack_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-3023323030132021-2200331203230303-2010201020131010-2200331321101100-1123123200121010-2211312202113213-1022201201212030-3332232333012332): complete subsection reference.

- [advertise_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-3001312121330320-1020103101122103-2002330223203013-1023233132022232-1221212231333210-2121233231313311-2320013113200121-2220102203121101): complete subsection reference.

- [advertise_v6_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-2132221203022111-1030012012210111-1130330003113321-3310012011010232-0110330101023202-3030123223022333-1232122123123023-2321311123310102): complete subsection reference.

<a id="canonical-1221113312100012-0331312223303310-3111011203030312-2201122312322230-3230203222212302-3010111101103202-3121213300101033-2210033030300002"></a>

<a id="canonical-0021132221103123-2230233002231033-0131301232122312-1311300203110321-3030033311301102-2000220123221122-2011200001012113-1221310100303300"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1312323232032011-1200022202103221-2100320022000120-0030323103013011-0121022130131311-0220021132203230-0323132312202102-3002131101203131"></a>

<a id="canonical-0121222313022103-3032030130013003-2220113303120332-2031211303233132-1101202303303312-3310010020200101-2302103120320032-2001001233202002"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [site](data-sources--http_loadbalancer--reference--group-004.md#canonical-2011322212233300-1203201313200103-3132222113222221-2230313112130221-3300330010111130-0123222333222131-0123211221030200-3203133030100232): complete subsection reference.

- [use_default_port](data-sources--http_loadbalancer--reference--group-004.md#canonical-2130333312313203-1331000112332100-2201112322203201-2232330300023103-2101020333202332-3320220200213023-2303222323221002-2301310032011302): complete subsection reference.

- [virtual_network](data-sources--http_loadbalancer--reference--group-004.md#canonical-1332223010333312-0131102221313003-0300220233112013-0122111312302200-2033133230020123-2130313110323123-1023313033002113-0111200011131110): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-004.md#canonical-3212332101232031-3022230013003113-3033210301033111-3010230321333330-0110223000001120-3001131233320121-1021331013203213-0230023001133223): complete subsection reference.

- [virtual_site_with_vip](data-sources--http_loadbalancer--reference--group-004.md#canonical-3320310222302220-3110120131200231-0223312331202232-1333031112221111-1123222320303310-2102130023123102-2232312221331333-1330002023030230): complete subsection reference.

- [vk8s_service](data-sources--http_loadbalancer--reference--group-004.md#canonical-1202331113211312-0330011103010312-0102220113223110-1003221323210230-0012000232123110-0100030330103213-0113023001033001-1222121212322012): complete subsection reference.

<a id="canonical-3023323030132021-2200331203230303-2010201020131010-2200331321101100-1123123200121010-2211312202113213-1022201201212030-3332232333012332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2122101221031000-1121313201221031-2000122111310201-1202222100230232-3300313021122131-3322001130031201-0300033232332220-2123000020233112"></a>

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

<a id="canonical-1102031321032212-3313320330110213-1133320021200002-0210111302222112-0020013323310211-1130120201202030-1312032020331202-1120201102130201"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](data-sources--http_loadbalancer--reference--group-004.md#canonical-3133312003302202-0132213310323132-3031230123033121-2320230233101320-2022323020301310-3133021012300231-0232113210211010-2033020300312310): complete subsection reference.

<a id="canonical-3133312003302202-0132213310323132-3031230123033121-2320230233101320-2022323020301310-3133021012300231-0232113210211010-2033020300312310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-3023323030132021-2200331203230303-2010201020131010-2200331321101100-1123123200121010-2211312202113213-1022201201212030-3332232333012332)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-0021021333120000-2032330000231203-2113201000013112-3101130110233000-0110301333021022-1211321332202312-0213221131303320-0002320202323230"></a>

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

<a id="canonical-2131003223003201-0102103332230333-3223110123231210-1212223213132221-1032320112312121-3022331222330011-2121332002031323-3303121021021310"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-1201211231231231-3003231031000130-2311013023221111-0110111320331323-1212220030211012-2200032100211002-2320220223113122-0213021323302332"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0112232121123221-2300011132113222-2100332201120013-1221113333102011-1110200002320020-0302000122220201-1221010132232023-1023231303011232"></a>

<a id="canonical-2211123110222313-2322201023200023-1231210002333332-2121201013311331-3010203003312233-3211133001313230-1203222331222010-0232321023321021"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0330220212131100-3310203121023000-1201130331302312-3333030332200211-0033110003011033-2330312123231000-3033202111230112-0031022333002121"></a>

<a id="canonical-3101212112223330-0120112313213330-1220032322203231-0233202110130231-3110133102133002-0313111222021323-2232303120210301-0212210311010333"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3001312121330320-1020103101122103-2002330223203013-1023233132022232-1221212231333210-2121233231313311-2320013113200121-2220102203121101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-3012133002001121-1210301110231101-1203123002103332-3023103312231300-2311002302000121-3302112012011210-2202321212012332-3230301123230302"></a>

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

<a id="canonical-1200022122313223-1222203213022110-1313300020231332-1202220211223132-1203102121010000-2013213130111322-1002203020113012-1000200330021133"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public`

- [public_ip](data-sources--http_loadbalancer--reference--group-004.md#canonical-2212101232303320-1002103021022201-3102330333010212-2132321112132320-0203111230133022-2222302120311013-1122310030120122-1122132300020330): complete subsection reference.

<a id="canonical-2212101232303320-1002103021022201-3102330333010212-2132321112132320-0203111230133022-2222302120311013-1122310030120122-1122132300020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-3001312121330320-1020103101122103-2002330223203013-1023233132022232-1221212231333210-2121233231313311-2320013113200121-2220102203121101)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-3100310202331101-2230012131133021-2011100323103022-3120130212303213-1312303331133000-1231033301130323-3113022303210022-0033222311131320"></a>

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

<a id="canonical-2103122311301132-3200103333311033-1023001312201313-0123330303002222-1111001131130331-3320202331010130-0022001202101111-2203221000003102"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-0112322010223300-0033133003332302-1013100231212220-2231000112010112-3101030111123031-2330323320010333-0011233323211133-0031300101021232"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3011002010333102-2200212001232110-3331031023232021-2203322203302012-0221002033101032-0211301023031312-2133233223132013-0300003330202021"></a>

<a id="canonical-1102133110200103-0010200120301011-0301020320101320-1220020133002030-1231233110031330-3320300330023230-1311102210000132-0131313203010232"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1312213333001300-1001112220132132-2312320231112333-2133031011312321-0212222203222122-3203203233311323-0100323220010001-0311221020230311"></a>

<a id="canonical-2220000231301200-0203202133020323-0102023001202303-1200100130010303-1113031231232003-2210233323321220-3323201010033311-3230122132031312"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2132221203022111-1030012012210111-1130330003113321-3310012011010232-0110330101023202-3030123223022333-1232122123123023-2321311123310102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-1333333302202310-1303320001330230-1131232131123212-0302010333212133-0010223111001132-1313310311220101-1130323032110331-0001222131313023"></a>

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

<a id="canonical-1032213113310213-1101232322102133-0323303211201121-1103021233023210-2020303222122231-0231103123011102-2121113113112010-0232200213302123"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](data-sources--http_loadbalancer--reference--group-004.md#canonical-3023001002322302-2322122303130003-1211130001130312-2232313213112013-1332013031003223-2311201230331223-3120210211032020-0121030111221310): complete subsection reference.

<a id="canonical-3023001002322302-2322122303130003-1211130001130312-2232313213112013-1332013031003223-2311201230331223-3120210211032020-0121030111221310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-2132221203022111-1030012012210111-1130330003113321-3310012011010232-0110330101023202-3030123223022333-1232122123123023-2321311123310102)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-1323223210002230-2100121211121320-0121100301202213-3002233110330321-3311001230223202-1003012232030322-0031223031332130-1120313003112012"></a>

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

<a id="canonical-2312300011233013-3200120112011213-1300233302031100-2112200331110013-3201030023312010-2103200122302103-2033111312010113-2113200112220323"></a>

### Direct properties for `advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-2110021023002233-3223031313212201-3220312131300032-1332130331203110-3121022123302122-0203222203202101-2232111203221302-2222333230002323"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1233320010320303-2131102110002112-3113312230313332-3230301200230122-0032323033003031-0133103010210211-2101203110100333-3321020132331100"></a>

<a id="canonical-1222003132113330-3132112103230321-2201110200133101-0111331312003021-0110321130132113-0303332231112011-0230010030123333-0132333033002231"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3103312333222310-1010313002130220-0302003131033222-0021100031003131-3112211203311313-3123021122330110-1322010302010132-0120322201131203"></a>

<a id="canonical-0102203120032331-2031211331123313-2213120122300022-0222310320110230-1002002322102023-0221223212322120-1310111312233003-0221200123300223"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2011322212233300-1203201313200103-3132222113222221-2230313112130221-3300330010111130-0123222333222131-0123211221030200-3203133030100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.site

<a id="canonical-1331111122102312-2033013201223121-0311013123002212-2200100010210330-0320311122010000-2111203331130302-1022003210313300-1101031322111000"></a>

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

<a id="canonical-2220111233301320-3333130312000020-3331212111112220-3001003130320301-0321103201023333-2302002302331302-3120003202112331-0233121313003112"></a>

### Direct properties for `advertise_custom.advertise_where.site`

<a id="canonical-3202100210203323-3323031210020010-3301013211301002-1100120210121110-3311211103310313-1003112200220323-1033101303320021-1013203012331133"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1203031122133020-0121132102321303-2321233120221223-3230323233010332-1332330330003003-0033231203302311-3332002002210210-1203322002011023"></a>

<a id="canonical-2111002332100203-2121102332122213-3110200011113210-2211200233030002-2021033111323303-3013003212302331-1211212312011000-2122031223211020"></a>

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

- [site](data-sources--http_loadbalancer--reference--group-004.md#canonical-1313113101322320-3121303302320322-2110320030031031-2022100013220311-2231312003321100-0203012121333311-1121220210303010-2220012222103312): complete subsection reference.

<a id="canonical-1313113101322320-3121303302320322-2110320030031031-2022100013220311-2231312003321100-0203012121333311-1121220210303010-2220012222103312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.site](data-sources--http_loadbalancer--reference--group-004.md#canonical-2011322212233300-1203201313200103-3132222113222221-2230313112130221-3300330010111130-0123222333222131-0123211221030200-3203133030100232)
- advertise_custom.advertise_where.site.site

<a id="canonical-3032133112212002-2012332331301201-3002203301222000-0230103321221212-1303001201132020-3011230012233221-3122001120322010-2300003303201101"></a>

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

<a id="canonical-2132032201311212-1332220010212303-3201202122332133-1133002203310230-2323100101101303-0310332232323202-0023222220030022-0033231220201311"></a>

### Direct properties for `advertise_custom.advertise_where.site.site`

<a id="canonical-2131112030233231-0113011220100331-3233111310102100-2321311330310323-0200200001211332-1021300202230201-1302102031232203-1020302212011310"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2330121101001211-2203133003123030-0330021201202003-2020131310131310-1102112220212231-0303011033122023-3000321303311221-1231002003000102"></a>

<a id="canonical-1212110230300333-1122311023233223-2201201112013131-3213302333301101-3300001232200220-3312010020220312-2220112031213331-3102200001101130"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0200322032003323-0323013120022110-0232202320000033-3121022101033023-0211112110321032-2110002323232200-1303300330201301-1311033021320301"></a>

<a id="canonical-0311200211300021-2332201023023232-3210002110111012-0211300000130310-2321301101003212-2000013320300111-1331200000130211-2003203221233301"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2130333312313203-1331000112332100-2201112322203201-2232330300023103-2101020333202332-3320220200213023-2303222323221002-2301310032011302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-3010303011213033-1033211222121203-1002010012133201-2301100122132003-1133200212213203-1002013001110210-1322003102003132-0231302213301211"></a>

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

<a id="canonical-1332223010333312-0131102221313003-0300220233112013-0122111312302200-2033133230020123-2130313110323123-1023313033002113-0111200011131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-3103111221101312-2210303303112121-3223100312022231-3031222300233012-3220111030101102-1023012132102002-0203013112321302-3133112200033011"></a>

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

<a id="canonical-0103001230123110-1301320230210222-3230331013302321-3230033323130011-0313312110102232-2220131133320232-0021203301130103-2312112020012001"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](data-sources--http_loadbalancer--reference--group-004.md#canonical-3112323103012002-2223120312320121-0102322120313332-3323103123212331-3321311332100000-0213330313100121-1302310222021323-1101201113300102): complete subsection reference.

- [default_vip](data-sources--http_loadbalancer--reference--group-004.md#canonical-1103210001212332-3130123231212023-2013032320320210-3021131310131122-3102021032100300-0101102002320033-0213330121113033-3020130100100022): complete subsection reference.

<a id="canonical-2223103312332230-3002333102330101-1101000113031232-2212113102122200-0322201023220310-0110311123122212-1202202023220101-2222300101133023"></a>

<a id="canonical-3300223230023322-3112010121320112-3300301223210123-2023203300131033-1011001122200111-1130321102220100-1223110102020012-1132302032330001"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2223230223202221-0113103110033102-0320312020121301-2321130033202233-3303101231122131-3213011030102003-3220100333000013-0103321303201110"></a>

<a id="canonical-1111202103100020-0321103021130002-3332031022212311-3223002223330201-3121012031312212-1112201323011330-2120003112220012-3302033001032202"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [virtual_network](data-sources--http_loadbalancer--reference--group-004.md#canonical-0332313110211122-3000212232312301-3220303321021210-3213313020201233-0103301312131022-3123233233203032-1031330332030302-3233311131213120): complete subsection reference.

<a id="canonical-3112323103012002-2223120312320121-0102322120313332-3323103123212331-3321311332100000-0213330313100121-1302310222021323-1101201113300102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.virtual_network](data-sources--http_loadbalancer--reference--group-004.md#canonical-1332223010333312-0131102221313003-0300220233112013-0122111312302200-2033133230020123-2130313110323123-1023313033002113-0111200011131110)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-3302300120211120-1130220312131030-3030031210232301-1203220310232332-3211302332201220-1223321120122001-1022120221201122-2103212011102331"></a>

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

<a id="canonical-1103210001212332-3130123231212023-2013032320320210-3021131310131122-3102021032100300-0101102002320033-0213330121113033-3020130100100022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.default_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.virtual_network](data-sources--http_loadbalancer--reference--group-004.md#canonical-1332223010333312-0131102221313003-0300220233112013-0122111312302200-2033133230020123-2130313110323123-1023313033002113-0111200011131110)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-0332302131312323-1232200000120312-3300133310033211-2100311322002332-1121331322323001-2330322230120111-3213331201202331-3010033030212220"></a>

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

<a id="canonical-0332313110211122-3000212232312301-3220303321021210-3213313020201233-0103301312131022-3123233233203032-1031330332030302-3233311131213120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_network.virtual_network` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.virtual_network](data-sources--http_loadbalancer--reference--group-004.md#canonical-1332223010333312-0131102221313003-0300220233112013-0122111312302200-2033133230020123-2130313110323123-1023313033002113-0111200011131110)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-0002022230020313-3112330200130200-1031031013100223-1121030120101333-3222322132112301-1332321113100122-2301310011002302-0232011333202130"></a>

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

<a id="canonical-2133223020203231-3102010032021233-3021200231003031-1302211013331200-2320300023202323-1032001311302302-3220101022223011-0030022132021321"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_network.virtual_network`

<a id="canonical-0333221331332301-3200120201200021-1010113213111102-3000113120211102-3103112201310022-1121020103020003-1131333120230112-0231020012112022"></a>

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0130332212301010-1120213002300202-3322131300302002-0221201321123123-3322211222103110-3023123120111331-2102311333132211-3331232020111032"></a>

<a id="canonical-0323210320213311-0312331312131120-2123310133321323-0121302233030032-3311203232231102-3233011303100230-3023120033223332-2333332312300220"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2332221120000130-2220032020013113-1320332030130120-0011133321232023-2112130232231031-1220231122221231-0101303113030221-2020030132131301"></a>

<a id="canonical-0322002021333221-2313300212301233-1311033020131311-0020113330222311-3130320203133102-2112311013322022-0320122223312010-0322233200213220"></a>

#### `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3212332101232031-3022230013003113-3033210301033111-3010230321333330-0110223000001120-3001131233320121-1021331013203213-0230023001133223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-1301112212122110-3200022121023211-0003213221320121-3001000223103000-1011033202303112-0020221113030220-1101321032100021-0021233101010233"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2310120320211300-1230311010022333-1011032320212010-3213133130200101-0133303320301211-2302202323103133-0102211333102101-2322221021112233"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site`

<a id="canonical-2032122210030001-0100223100320002-0010213000030022-1333032302302022-0020201021113333-3333010320131133-0121212301220033-1223133301020120"></a>

#### `advertise_custom.advertise_where.virtual_site.network` property

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

- [virtual_site](data-sources--http_loadbalancer--reference--group-004.md#canonical-0232133201222203-1201122200223310-0032212021112110-2101331131312131-1011110310012222-1012312300303333-1200000003021302-0213133321103020): complete subsection reference.

<a id="canonical-0232133201222203-1201122200223310-0032212021112110-2101331131312131-1011110310012222-1012312300303333-1200000003021302-0213133321103020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.virtual_site](data-sources--http_loadbalancer--reference--group-004.md#canonical-3212332101232031-3022230013003113-3033210301033111-3010230321333330-0110223000001120-3001131233320121-1021331013203213-0230023001133223)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-3001200013132310-3021030132021311-2231031303333320-0203323231222023-3010331300021310-0200120113300321-1323002123112311-1010103313030002"></a>

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

<a id="canonical-2300130201112212-3303121012210102-1232101302300120-3211222100101100-2001312213220010-3022232220301002-3211232012023021-0213011023000302"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site.virtual_site`

<a id="canonical-2312021313220103-0010211120132322-3103310212231010-3212212001221211-2023032331012020-0212010012203112-0302200021011330-0322212331030123"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1030032200033232-2333020213021313-3313222011330031-3330103322123332-2031001303123300-2222030131033102-1021212011312032-2023220030211313"></a>

<a id="canonical-0322201232333031-0320120331333223-0233333012301020-2321012331002100-1323020111123233-2230303233300011-0030201300211122-1220202110320320"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0113012203211220-0111213000110001-2001222221111203-2131231033120321-3333010112133000-1111231230123222-1223001320131222-2020020111101002"></a>

<a id="canonical-1023303200031010-3033000121112031-0012022131131122-1022203110102213-1003211131123132-0102002101020301-0010332030311101-1031333003132301"></a>

#### `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3320310222302220-3110120131200231-0223312331202232-1333031112221111-1123222320303310-2102130023123102-2232312221331333-1330002023030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-0202323310120231-1003023001202132-0112032220232123-2120302313023112-1131302013123110-2233303130133332-2122110200332300-2312201333303102"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2201213100213312-3102310232303313-0312133020023121-1302030032313200-3010030332130300-3310330200031122-1112100221120202-1201213002213310"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip`

<a id="canonical-2000022033213033-2132312023032211-3312330230130230-1110112001212312-1103222011300333-0303231211010331-3132121333030011-0031330032131011"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.ip` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0122013130032011-3100302013320332-2303133103213032-0032010213312020-1321100303103003-2110302012323003-3311001322113210-1322230230230301"></a>

<a id="canonical-0313233032020013-2002223132233230-0313221101210333-1302212110031003-1120230030000021-0303200101231323-3223322232001003-0222331223303221"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.network` property

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on virtual-site with specified VIP

All outside networks.

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

- [virtual_site](data-sources--http_loadbalancer--reference--group-004.md#canonical-2221020301023100-1221320301210332-1021032203231113-1310232123120200-1321013313213030-2031202312130323-2201120002313133-0111102231002000): complete subsection reference.

<a id="canonical-2221020301023100-1221320301210332-1021032203231113-1310232123120200-1321013313213030-2031202312130323-2201120002313133-0111102231002000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--http_loadbalancer--reference--group-004.md#canonical-3320310222302220-3110120131200231-0223312331202232-1333031112221111-1123222320303310-2102130023123102-2232312221331333-1330002023030230)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-3102310312302233-1011002301032211-0112332332023203-2033112211020323-1232232020111331-0230100331100010-0122302130313023-3223133122323101"></a>

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

<a id="canonical-0130011110133221-3322123310210301-0103313323032111-1312033100102231-0010222112121200-0201320133113223-0123113002100022-2332121001000303"></a>

### Direct properties for `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site`

<a id="canonical-1323102333220123-2012020300310233-1222102312113230-0200101131102122-3032331110123211-0333210331323121-2132201303121223-2100023010110101"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1020223002310112-0032322320223230-0222011232323302-3022113003210110-3033123120323200-2100230322223331-1100021103101220-0203321210133101"></a>

<a id="canonical-1232012123201313-3001311023102120-3231112131002100-1102002332301023-2323013320020300-1221030023110313-1030301022101111-3123313223333123"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0233031320021303-0003100033002030-3312133022012000-0132203132332220-1220011313112103-1112332121210133-2213331032222203-2001323001301320"></a>

<a id="canonical-1221300300312211-3101013210331000-3231213320212212-3111321131101312-2231303033330303-3312032223100120-3312210010210030-1002212313011032"></a>

#### `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1202331113211312-0330011103010312-0102220113223110-1003221323210230-0012000232123110-0100030330103213-0113023001033001-1222121212322012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-1312303000201031-0131222100302021-2031322202323313-2212121312101031-1002031331220230-2113300002200101-1300001303332110-0310100112333032"></a>

Type: `"single"`. Computed.

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

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

<a id="canonical-1223013231201012-3113022221303212-3030021101322022-3032033100121013-0233320323321311-0220233031310033-0021221023200003-2001203310031331"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service`

- [site](data-sources--http_loadbalancer--reference--group-004.md#canonical-3133331031333003-0203312232321210-0313212110202203-1332111233110200-2333120123223312-2210300223201201-3203033310201021-1001110021000100): complete subsection reference.

- [virtual_site](data-sources--http_loadbalancer--reference--group-004.md#canonical-2030110332200223-3322120111101001-2110202121233323-2122210021121133-1112132222323111-0123222320200011-2221130110121331-3100232301123312): complete subsection reference.

<a id="canonical-3133331031333003-0203312232321210-0313212110202203-1332111233110200-2333120123223312-2210300223201201-3203033310201021-1001110021000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.vk8s_service](data-sources--http_loadbalancer--reference--group-004.md#canonical-1202331113211312-0330011103010312-0102220113223110-1003221323210230-0012000232123110-0100030330103213-0113023001033001-1222121212322012)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-1010010010201201-2023222002031123-3323321132100213-3321222322130210-1033103220103112-2122212322223010-3233312031032311-2230022031233001"></a>

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

<a id="canonical-2101131230223130-0000032222110011-2113203332320033-0213101031101030-2011230231323203-1320012010022200-3300110113122023-3220200223121233"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.site`

<a id="canonical-1021103220312302-2003201313221021-3003110212213230-1223331121310313-1121321102212312-3121111231011230-3100332202101213-0033200010103222"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1023130121112113-3031110312231021-1212102323313233-1222030220132220-0321132112023112-3010101322302001-1220322011033012-1212330211112131"></a>

<a id="canonical-3031232000112000-2033100232002212-0220221200102320-2311102021100130-3000003220311331-1200131112300020-1122120031231302-0312303233213003"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1011223102202020-3100313331120302-1012012223101013-0123100221212202-2201100320223312-3003321113101021-3120321112003122-2332011112113223"></a>

<a id="canonical-1211132101320321-2311101111231332-1013333000322021-3101303211021123-2200322333202200-2113321302211122-1003202230001213-2120211310131300"></a>

#### `advertise_custom.advertise_where.vk8s_service.site.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2030110332200223-3322120111101001-2110202121233323-2122210021121133-1112132222323111-0123222320200011-2221130110121331-3100232301123312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_custom.advertise_where.vk8s_service.virtual_site` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_custom](data-sources--http_loadbalancer--reference--group-004.md#canonical-2321010232222113-2212023011131230-0322020011213020-2030233113312303-2221333102012201-1001013033321102-3013221120322303-1300122003302000)
- [advertise_custom.advertise_where](data-sources--http_loadbalancer--reference--group-004.md#canonical-1201102113133312-0232033231222232-3031102111031223-2301032032110210-0031020220003021-3133320211333312-3200101201100323-0233201011221213)
- [advertise_custom.advertise_where.vk8s_service](data-sources--http_loadbalancer--reference--group-004.md#canonical-1202331113211312-0330011103010312-0102220113223110-1003221323210230-0012000232123110-0100030330103213-0113023001033001-1222121212322012)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-0101301222103101-3133221310021323-0330323232221200-2201021121103300-3132330320222011-0332013201312013-1031131322020031-0323230332210231"></a>

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

<a id="canonical-1310202003132311-3131122322301300-1233021202223312-2312200021031000-3311103233102323-2131022020333030-2213101201211301-3213000113012110"></a>

### Direct properties for `advertise_custom.advertise_where.vk8s_service.virtual_site`

<a id="canonical-1202031011203213-2320301313223210-2032130031002312-1112120110000212-2322010022202311-1302233223222213-3111321121120310-3133133222223223"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1003032213021031-2102122103333310-1223232311323333-3121320221123100-0322003323331022-2031132123200210-2232223121013020-0112030231033332"></a>

<a id="canonical-1103222030311321-3311300231131122-0100310333220132-3321020030301111-1330013101123321-1023212220311122-1303032131022310-0212132110133132"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2032033130123200-1321202311003030-3323332230131123-0000321001323201-0110121000231120-3331022011202120-0122231032101000-3213130231113101"></a>

<a id="canonical-2331230221333210-2223001303323123-0332033202303332-3131313011321332-3213202300223131-0132033031003003-2100002011331131-1221121213223302"></a>

#### `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2230111322213330-1133313211131321-0000311132313220-1123102333032011-1231222323323010-0102210111000010-0332002211011211-1310131231110211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- advertise_dualstack_on_public

<a id="canonical-0303330131303232-0032101200321202-2333020312033323-1133203011110300-2013002032133102-1331231230202003-3211301301303202-0022333320122123"></a>

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

<a id="canonical-3100303220131221-3123221012010302-3033202333303213-3123123302233002-0132330200212123-2200111103111113-0310200312001110-2131301121312130"></a>

### Direct properties for `advertise_dualstack_on_public`

- [public_ip](data-sources--http_loadbalancer--reference--group-004.md#canonical-1030213110031030-0301320301021023-2213230212333310-2230012031131101-1000332122213300-2200202120220200-2301202001032012-1110201031232320): complete subsection reference.

<a id="canonical-1030213110031030-0301320301021023-2213230212333310-2230012031131101-1000332122213300-2200202120220200-2301202001032012-1110201031232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_dualstack_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-2230111322213330-1133313211131321-0000311132313220-1123102333032011-1231222323323010-0102210111000010-0332002211011211-1310131231110211)
- advertise_dualstack_on_public.public_ip

<a id="canonical-1122123001303033-2200110231231022-0130021311303110-1330331033002101-2223232300313121-1230100220211200-1323011103313012-1331320121313232"></a>

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

<a id="canonical-3032013013202111-0212221320032002-3102032101123023-0232231220211120-2333200232202022-2211013100112100-2002000003310222-3302111223233220"></a>

### Direct properties for `advertise_dualstack_on_public.public_ip`

<a id="canonical-1102212010012213-3130013210101311-0312001002010033-0203101122312212-0013132313031212-1222333220111223-2220312312131220-3302203030221011"></a>

#### `advertise_dualstack_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1111011001122010-0220300203021331-3000123002110111-3021222121102020-1021211203310123-0010010300112012-0002221102211123-1121032002302123"></a>

<a id="canonical-3333332233233210-3013102220010223-2133312012121301-2023103212022013-3231022201032010-0101132100001100-3020212031102312-0203023213331233"></a>

#### `advertise_dualstack_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2211003200323132-2202213031300101-0111002033232021-0031013222111312-2331233011020303-0033323110220022-0011231030213101-2323130103320310"></a>

<a id="canonical-2312103313000222-0033003020112221-1112133202023000-2221121222100100-1322231221333332-1131220312100130-3103100203300120-3331322013111013"></a>

#### `advertise_dualstack_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2110110202303002-3212310231101322-0113133010033132-1231001030020312-3101021020302221-3211111113213332-0202113200011002-0213233111330230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- advertise_on_public

<a id="canonical-0131001033312203-1233133313001030-3313122220112000-3012100022321011-2213000212113331-1311321221112331-3220312211312213-2212231101232121"></a>

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

<a id="canonical-3311132000232222-3021323310313130-0302230032322302-3211230201331011-1312333212222101-3120221010131302-1113333203322001-2211031022303003"></a>

### Direct properties for `advertise_on_public`

- [public_ip](data-sources--http_loadbalancer--reference--group-004.md#canonical-3110100330010033-2202121011222213-3113132332133103-0102133220033203-0033022033320011-3110220221221312-3023233210231111-3210102000103233): complete subsection reference.

<a id="canonical-3110100330010033-2202121011222213-3113132332133103-0102133220033203-0033022033320011-3110220221221312-3023233210231111-3210102000103233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [advertise_on_public](data-sources--http_loadbalancer--reference--group-004.md#canonical-2110110202303002-3212310231101322-0113133010033132-1231001030020312-3101021020302221-3211111113213332-0202113200011002-0213233111330230)
- advertise_on_public.public_ip

<a id="canonical-3032221020303003-0323122001121123-0212031203321202-2103130321302102-1023023103033310-3031310013102302-3110303012003320-1132202223020202"></a>

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

<a id="canonical-0321311232213311-1201230001320131-1031333012113310-2322001033110122-3133212101021300-3000330102120133-1003300102100212-2213000013011022"></a>

### Direct properties for `advertise_on_public.public_ip`

<a id="canonical-1023303303303231-3200010030123222-0212003010030003-2202300310310012-2100022223322233-0223211313033133-0303101111222201-2020332232323213"></a>

#### `advertise_on_public.public_ip.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1201203311010023-3232302331212210-3022311022102130-1212200030312230-2232020030311022-2311130103330003-0210133322033010-0111203131000022"></a>

<a id="canonical-0220020200100223-0221102030130130-1331321221222021-1032322320003221-2201013022120313-0023010232221032-1321031210031110-0313300033200101"></a>

#### `advertise_on_public.public_ip.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0103201011332031-2332020201020011-1030113310201012-0323203031102212-1220130101213223-0221013121223222-0303312132013000-0212132032211303"></a>

<a id="canonical-1010213202001023-2332020313121311-2032102212023120-2212033100010102-1011332130323300-0302023202000100-3123213332021301-2031000322303100"></a>

#### `advertise_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0023021022222102-2330002011133123-3330210113010220-1322222203110001-2013001032123010-2211332013223321-1122030302003010-3011102323002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_on_public_default_vip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- advertise_on_public_default_vip

<a id="canonical-3110212023223203-2310311120102130-1211123102111330-3233113030202022-3131033223003130-0200221320210332-1021232103112310-1333131001101322"></a>

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

<a id="canonical-3233101001322201-3010101330010210-1230122302103210-1020231313123311-1123022210033012-0001211313213020-2123001023213012-0132232031212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- advertise_v6_on_public

<a id="canonical-3321100211232200-0201302220112112-2312220001200002-3310313211332132-3121131101201001-2203323330032033-0220330303001130-1213112220232121"></a>

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

<a id="canonical-3133031322201021-2313031202223112-2100020112011121-2320320332000302-1132321232211121-2333010302223321-3111231101122111-1130131022122133"></a>

### Direct properties for `advertise_v6_on_public`

- [public_ip](data-sources--http_loadbalancer--reference--group-005.md#canonical-3131233300321233-0212301102322121-1202111033301033-2113200310232222-0003112222313301-0222301020302212-3020100103323101-2212320023210010): complete subsection reference.
