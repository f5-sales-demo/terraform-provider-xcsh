---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3130102200011303-0111031030113000-2110033302010100-0122110303103100-1302031032000121-2223322011312132-1131011012312331-2113001031103330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- active_service_policies

<a id="canonical-0212321113002131-1012102032311133-0320031101130112-0303122232123131-0232022323223210-3111003011030131-0213310120230033-2321111010000001"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Additional upstream details:

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

- [active_service_policies](resources--cdn_loadbalancer--reference--group-003.md#canonical-0212321113002131-1012102032311133-0320031101130112-0303122232123131-0232022323223210-3111003011030131-0213310120230033-2321111010000001)
- [no_service_policies](resources--cdn_loadbalancer--reference--group-012.md#canonical-1013310300231013-2321212333203010-0332022031203123-3112002213301121-3010211320310130-3233202333011223-1131111100001030-2310100131120023)
- [service_policies_from_namespace](resources--cdn_loadbalancer--reference--group-015.md#canonical-0212231232132301-3212113303020322-3222132120230030-1123313213201131-1130023023230203-2203303210021230-2103211113213120-2321022020101213)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300301200122322-3010032002010313-3113330332031110-3130311001130001-3133331013031300-1011210111011003-3301302213221332-3122323131112103"></a>

### Direct properties for `active_service_policies`

- [policies](resources--cdn_loadbalancer--reference--group-003.md#canonical-1213002001311210-1231331313121020-3101203021300002-1100032231231102-1303202332130112-3313222330003030-0032210112331110-0223132011212323): complete subsection reference.

<a id="canonical-1213002001311210-1231331313121020-3101203021300002-1100032231231102-1303202332130112-3313222330003030-0032210112331110-0223132011212323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `active_service_policies.policies` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [active_service_policies](resources--cdn_loadbalancer--reference--group-003.md#canonical-3130102200011303-0111031030113000-2110033302010100-0122110303103100-1302031032000121-2223322011312132-1131011012312331-2113001031103330)
- active_service_policies.policies

<a id="canonical-2203102030230021-3223122303010221-2232300022332230-1212103101213120-1312020110233001-1212123122301232-2122102313333101-2312303231322131"></a>

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

<a id="canonical-2103102231003103-0111033131110012-1133122320023132-0002223103121100-2202300203101312-3030231302202002-0322202211012333-2310000113011002"></a>

### Direct properties for `active_service_policies.policies`

<a id="canonical-2010332103023122-0300113323302010-3003033201103323-0032321102023321-2103113121101120-3131102123322020-0121223321033023-3221210312222022"></a>

#### `active_service_policies.policies.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2003302020002333-1131332303123110-1230303112320020-2210331022220211-3313202132320323-0001131333230311-3102133112020203-1200031011023222"></a>

<a id="canonical-1133130330010012-0000122302123032-0310001121223021-1201321102112023-1133212213111303-0022332232213310-3330313123103301-1110110222221220"></a>

#### `active_service_policies.policies.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0313202021320201-3303001221302020-3021121110100101-0030230103223012-2102121111322300-2210211330302023-3200132111320221-3333200103013100"></a>

<a id="canonical-1202112021333132-1220010122221110-2120021003222111-1300301030331112-3123202212321222-3031331111131213-2233112003311002-2313330131001202"></a>

#### `active_service_policies.policies.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- api_rate_limit

<a id="canonical-3032213023202030-1310301123311331-0220121322010230-0311031331000231-1313113230310112-2123021333302120-2001022322130020-2121122322100013"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\]
APIRateLimit.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "custom_ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("bypass_rate_limiting_rules",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-3032213023202030-1310301123311331-0220121322010230-0311031331000231-1313113230310112-2123021333302120-2001022322130020-2121122322100013)
- [disable_rate_limit](resources--cdn_loadbalancer--reference--group-010.md#canonical-3133211010101322-0323301002321331-2233011002232323-1123100210322211-1330323131213103-3011211112311233-3123030230331311-1230310111020221)
- [rate_limit](resources--cdn_loadbalancer--reference--group-014.md#canonical-3323121002313201-3031120322332203-2113131320323133-3022023331331231-1100333301321203-0121201122211101-2223013112203301-0210101213220100)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
api_rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233001203323130-2332002302222223-1300311323032231-3010013202211332-1030011201210330-2100313201103023-0002100232223110-3232032203211013"></a>

### Direct properties for `api_rate_limit`

- [api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031): complete subsection reference.

- [bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110): complete subsection reference.

- [custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-2333012122103001-0111232101323001-2333210301300122-3302011003223222-3002023122303203-1103001221000201-1310203202022230-2010003132313013): complete subsection reference.

- [ip_allowed_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-1131100213132110-3310333232222331-2231310112312012-2311333102233132-3301112001213202-3003223102201221-0101320232303301-1030332212123201): complete subsection reference.

- [no_ip_allowed_list](resources--cdn_loadbalancer--reference--group-005.md#canonical-3313313301303203-1122102000003112-1022231311223221-1302133120330011-1231301301331313-1323332102200203-1110313222013213-0023131100032020): complete subsection reference.

- [server_url_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033): complete subsection reference.

<a id="canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.api_endpoint_rules

<a id="canonical-0202122013031000-2001210303123131-2033301110122020-0031110312301212-2101310111303231-2221331223332320-0221033000133231-0223201222311110"></a>

Type: `"object"`. list nested block, Optional.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("api_endpoint_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_endpoint_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021112001323122-1313230003313113-3031113110012031-2012323232122320-2322110020201100-3302030110213230-2203322022100022-0110130313213230"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-003.md#canonical-3312122323310111-0100303223303133-1323122321133130-1211300120102301-0300331231103213-0031110100203013-0202023132312020-1030003222033012): complete subsection reference.

- [api_endpoint_method](resources--cdn_loadbalancer--reference--group-003.md#canonical-1101312131131202-1133331313023313-0101222211333232-0113023010220332-0131231223032333-3121330010121133-3333222312013111-3212310233111011): complete subsection reference.

<a id="canonical-1322320222332032-3113001302001211-1010300222211110-3211332201003200-3310222323013002-2221130203132003-1012320023113130-1000333200222220"></a>

<a id="canonical-2133021303032313-2022211321020103-2301030302032132-1332202222133131-3011110211300133-1233101211232033-0212130221101231-2011012013013113"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_path` property

Type: `"string"`. Optional.

API Endpoint. The endpoint (path) of the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011): complete subsection reference.

- [inline_rate_limiter](resources--cdn_loadbalancer--reference--group-003.md#canonical-0301122111301232-2013220232300313-3011332300103100-3023223212103132-0232010111232230-0311320233220000-3301331203330121-2320222322032211): complete subsection reference.

- [ref_rate_limiter](resources--cdn_loadbalancer--reference--group-003.md#canonical-3200220130112333-3222120132000131-3123212332333021-3010211231213321-0322000001030333-2000121322111101-2113333323032202-3203102320310321): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332): complete subsection reference.

<a id="canonical-0012022233011232-2331331003010230-0230021322021130-1211113312023310-3313313321203302-0002123010101101-2233112303120111-0031322200330213"></a>

<a id="canonical-3021132121002102-1012100320231310-2122200120000113-2132030331332330-1303100100311002-0321001010010010-3023003032013011-2211223221230033"></a>

#### `api_rate_limit.api_endpoint_rules.specific_domain` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-3312122323310111-0100303223303133-1323122321133130-1211300120102301-0300331231103213-0031110100203013-0202023132312020-1030003222033012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-0030330131101233-3320223110011121-0111012133102213-1331220022321321-2131120112033211-1030103302001200-0233111222112010-1103100001112131"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101312131131202-1133331313023313-0101222211333232-0113023010220332-0131231223032333-3121330010121133-3333222312013111-3212310233111011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.api_endpoint_method` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-3303331012222313-0022302210212001-3001012110123031-1021121023103103-3331321231313332-3110101123313013-3123131102211110-1221331111313121"></a>

Type: `"object"`. single nested block, Optional.

An HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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
api_endpoint_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120220301311112-2131003121103110-1023003311030230-3101230302011312-1131133210210230-0031000211032113-3110302331112201-3123213313311220"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.api_endpoint_method`

<a id="canonical-2023111202213331-2003003230021120-0231230113302132-3331200310312022-1020132231330231-0221122122111213-3202331002001312-1212002313012312"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.invert_matcher` property

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-2123302113320211-0210203323022120-1212322223112131-3303211011131333-3330330232332112-2100321133310230-3020002020323132-2322211133113102"></a>

<a id="canonical-1220213121110020-1321122302132002-3032113033012102-0310222000310222-1220200230222203-2101300313011311-2022013221220100-3200133133012223"></a>

#### `api_rate_limit.api_endpoint_rules.api_endpoint_method.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-0022333322130003-2030003331320223-2300003123201321-1320311221113311-0303122022012110-1201202202130032-0211030021300302-1320210101222312"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233210322213013-2030223002203320-3102133013212121-0133123002212202-2313023230030121-3232120203321022-1022030312312132-2113200300100330"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher`

- [any_client](resources--cdn_loadbalancer--reference--group-003.md#canonical-0011303032100223-0303012110010110-2103000230223132-1130213011003031-2222122100000200-1122100222213321-0011120212310202-2212111301321223): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-003.md#canonical-0223031202323002-3313010003103322-3230123101000101-3101301211030211-2031011101022222-2303002111102311-1010111002112113-0022331100030023): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-003.md#canonical-1331313320132132-0221123101030021-0200123212003033-1321330002111321-3030323310222310-2333310221131022-3223101200232100-0013120003212010): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-0001022001203220-2121001122230012-2333331333123133-0030012311022012-1212113130311223-2223201232120002-3200133030000310-2122201113122301): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-003.md#canonical-1001010033101013-0001332210222232-1213333000202102-0100333211320120-2003031223220121-0230030213112021-3203111303003122-3200212001110110): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-1103011111312020-3023033331120001-3011032300220310-1123001310213023-0323000102202230-3220012202223011-0032310102221020-2311332013100232): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-003.md#canonical-3010121011033022-3222131223021033-3030301130202203-1223133113330322-0121020303220321-2130212332032130-1203333101202213-2333322032322110): complete subsection reference.

- [ip_threat_category_list](resources--cdn_loadbalancer--reference--group-003.md#canonical-3103320212322202-3103001332211102-2331323212212202-2122032300230221-2021033302100111-0212113020113202-2231221013320300-0111102113201022): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-1321200223213022-2122221002022313-3211030313213000-1133332032232321-0311103010000203-1003033002312203-3311210130233103-3110221231021300): complete subsection reference.

<a id="canonical-0011303032100223-0303012110010110-2103000230223132-1130213011003031-2222122100000200-1122100222213321-0011120212310202-2212111301321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-3100200030213031-2301010320212331-3111000102112021-1333033111011121-0022123323332101-1112122121122103-1320111313223232-3303231203231120"></a>

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
any_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223031202323002-3313010003103322-3230123101000101-3101301211030211-2031011101022222-2303002111102311-1010111002112113-0022331100030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-2101011301003123-1121102230113301-1301332133211210-0223100203200202-2233102102101020-0221233021132121-0113023032020221-0002011130032202"></a>

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
any_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331313320132132-0221123101030021-0200123212003033-1321330002111321-3030323310222310-2333310221131022-3223101200232100-0013120003212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-0333100133330201-3123321103001021-2030120002003012-1012300233233222-1203232211001203-0211031030211213-2022302010001102-0103301223110313"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1030210301331200-1300022000123202-0303220202032021-0100212121310123-1132300312120022-2022231121332001-0123132203120030-1323111002212321"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_list`

<a id="canonical-2001012322132023-3020123210210221-2000030201123220-0033033202301231-0211201222221210-1120211323220300-2222322213321032-3132102012312132"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_list.as_numbers` property

Type: `["list", "number"]`. Optional.

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

<a id="canonical-0001022001203220-2121001122230012-2333331333123133-0030012311022012-1212113130311223-2223201232120002-3200133030000310-2122201113122301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-2100120333121013-2322020331022300-1010222213202332-1112021303132033-1100020123331331-3202003032023230-1330131311231112-0031030312323011"></a>

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

<a id="canonical-2330031223013123-0313323022123013-0133223320301110-2133231110120303-3310231111122001-0213033200202333-3233220222311223-1331233102300230"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher`

- [asn_sets](resources--cdn_loadbalancer--reference--group-003.md#canonical-1110011002110102-0001013032233133-2321003010320111-1231120020322113-0202012200301103-2200010230233011-2222111123230231-3122302202231032): complete subsection reference.

<a id="canonical-1110011002110102-0001013032233133-2321003010320111-1231120020322113-0202012200301103-2200010230233011-2222111123230231-3122302202231032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-0001022001203220-2121001122230012-2333331333123133-0030012311022012-1212113130311223-2223201232120002-3200133030000310-2122201113122301)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-3331133120022103-1133033130131323-1301030211331001-3233120112212131-0233201002331330-3001220132033301-0022002231221321-1102320312001122"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231231023221122-2213132022000003-2003122030103231-1013330121000003-0122332330232130-1103213033321332-0112320300313323-2302100100213300"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets`

<a id="canonical-3210330103010300-3131013322120021-0111222012012032-0310312133030233-0130202032012010-0233112302133020-0321021312303233-3113031121133213"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.kind` property

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

<a id="canonical-1301012302331103-1213213111221320-0022233332100131-1012320011222210-1012320032002023-0012100123331221-0022113102211030-2030101231102330"></a>

<a id="canonical-1213033121312032-1032311301322213-2130211000000312-0223212303113001-2222003120302231-0330203011303312-0221121120211223-3203200223211031"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.name` property

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

<a id="canonical-3112211032323333-0130102101002213-1033001111000322-3330201301002211-0030021313233032-0210020233021012-1101030103120123-0112122300323320"></a>

<a id="canonical-0231010321111210-0130222200003032-3221232033221031-0111133323111333-1011011321321021-3313032213023131-2200002301320202-0211302021011302"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2223232200213021-1223110002100023-1222310101233033-0323202333101301-3312112233133122-3013312233223032-0022222333300312-1302331211121301"></a>

<a id="canonical-2301122021133312-0201320130333333-0302001312012310-1132232232301302-1001333032111301-2302033102330201-2332023130322002-2320132203032320"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.tenant` property

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

<a id="canonical-3111001211101133-2200032012033010-2311212302131130-0213013131320320-0320331000131033-3221121133120331-2330001222103133-3101312033321110"></a>

<a id="canonical-1310012323133100-2123220200230223-0033132301012010-0121012032031221-2221031200033220-2030312003100123-0010100111103012-1111210200301120"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets.uid` property

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

<a id="canonical-1001010033101013-0001332210222232-1213333000202102-0100333211320120-2003031223220121-0230030213112021-3203111303003122-3200212001110110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.client_selector` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-0100233121212033-1213000223223030-2232333212033321-3103010122111231-0312313013130023-3030130011211130-0221030131102131-3012013221002332"></a>

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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313002012101321-3132133231020201-0100101321000000-0031222332333021-0011232102020232-3130201223003031-1003110102131111-0102023320122300"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.client_selector`

<a id="canonical-2333131303300311-0112233322032002-3133211110133011-2133130202330221-2022113030122222-3322112310310002-2210222033112001-3133023323201112"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.client_selector.expressions` property

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

<a id="canonical-1103011111312020-3023033331120001-3011032300220310-1123001310213023-0323000102202230-3220012202223011-0032310102221020-2311332013100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-2233100100311311-2301102223233013-1333222330230102-0220011220211211-0030001131021111-0200213222010331-2230022211211103-2313323010323313"></a>

Type: `"object"`. single nested block, Optional.

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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333110100220102-1000112022100113-0211212132231023-2321330001332312-2231313021002311-0021111122210322-1332223322122333-0313013110131112"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher`

<a id="canonical-0301223222010031-3112332322223000-3310001330010001-2222230223210132-1220333022231313-3210202101310300-1103033012203033-0021031210100302"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.invert_matcher` property

Type: `"bool"`. Optional.

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-003.md#canonical-1301303121122010-0000023023222011-0310330032221001-3120012000021222-1311102020322101-1000232323330333-0010211322023313-2012123032003112): complete subsection reference.

<a id="canonical-1301303121122010-0000023023222011-0310330032221001-3120012000021222-1311102020322101-1000232323330333-0010211322023313-2012123032003112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-1103011111312020-3023033331120001-3011032300220310-1123001310213023-0323000102202230-3220012202223011-0032310102221020-2311332013100232)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-1233233332300113-3110102013302122-2313223002130310-1031222223301211-2301331131022213-3201222121122010-0320303030111120-2023031130222110"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100030103202111-2003033033022212-3023113202103110-0032113110301123-1013122100013002-2100230111202231-3310200121300010-0132000002201030"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets`

<a id="canonical-3313332321001010-2130031102213332-2111023210113231-0233101311201023-0111303203331310-1011110322100201-1011220311203221-3012230213111031"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.kind` property

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

<a id="canonical-2033031021231103-1331230300003032-3320212033311123-3031202331003001-3020111200020221-3103101032211221-1303123321201232-0103121330110002"></a>

<a id="canonical-3011213012033221-2000101310330132-2213013022222002-2200111230321322-2100031121311300-2010022332122223-0011333123301213-0230123330103233"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.name` property

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

<a id="canonical-3100003101330121-0320323332121101-0111001321322102-1223200311213031-2102302331312203-0101113333213212-1123021133032220-2322030202331030"></a>

<a id="canonical-2111112021223110-0002222231011231-1000230020301103-0211230222100031-2220300303231030-2013123322202322-3223112333213300-0330331311021133"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3120330321112231-3112220121313103-0321130122210010-2121211011231013-1033202032112300-0031020102131133-0120002013100002-2003120023121023"></a>

<a id="canonical-1331221233102100-2112013102022320-2321103221112312-1120033323110220-0201203332301011-1121120222331033-0103103100021301-2030310121013333"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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

<a id="canonical-1311030012132222-0101130032031200-0333310130211213-0123111000311120-3323103031110001-2102300331223131-0133113202332021-3121111223121320"></a>

<a id="canonical-2031110010020132-1212003213013213-2011232213310031-2010202031321302-2202110333222311-1022220300010012-1030212301122332-1210323013231201"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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

<a id="canonical-3010121011033022-3222131223021033-3030301130202203-1223133113330322-0121020303220321-2130212332032130-1203333101202213-2333322032322110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-1310203010101231-0000120233032032-3131102121211101-0013331012201130-1221320302312100-3022300003211131-3300011302320013-3012010222202300"></a>

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

<a id="canonical-3013111021100011-3202012031311213-2233332113003013-3300203203213111-2333321131301110-0001110023231121-0032203030302131-0203031333330112"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list`

<a id="canonical-3132330101210012-2203100313031100-0131111312211300-1133132313123030-2122311301003001-1321232232131202-1302332231313332-3010010120331312"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.invert_match` property

Type: `"bool"`. Optional.

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

<a id="canonical-3111301322123113-2111212213023300-1223201131033321-1001110310223222-0203211123210210-0331012000212001-2003100002000032-1121203211031103"></a>

<a id="canonical-1020220221203230-1313231323330330-0113330222002221-1220031103011013-1311103320301100-1012312122220211-2103001302130331-2223233012323133"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

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

<a id="canonical-3103320212322202-3103001332211102-2331323212212202-2122032300230221-2021033302100111-0212113020113202-2231221013320300-0111102113201022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-2332022321000221-3121032113010121-1210210310212203-3133332210123202-2230120122130002-1031231133322202-3122033003212311-1132010133322023"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333031233010131-2101132002132121-3321021111212212-0301122102130330-3313300210220113-3223222001310312-3212212201031321-0322210321101200"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list`

<a id="canonical-3230023123303320-1032130213123303-2302102232210220-0300310322001110-0002300000101021-3122113311230011-0022011200000101-2113120233331321"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1321200223213022-2122221002022313-3211030313213000-1133332032232321-0311103010000203-1003033002312203-3311210130233103-3110221231021300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.client_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-2020102032110222-2002300113233132-0033300332023020-3213131320322020-1222120201212333-0303003023321113-3201120301102333-2322000013121011)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1311201101330132-3323202232103120-2030313121313112-3013330212130102-1231013302202021-2002011133120320-3112301223300231-1130020121312102"></a>

Type: `"object"`. single nested block, Optional.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233121320323223-0013100103110312-3311120221330330-2202232103322102-1100012332331121-0121312111130100-0212221111301302-1022212101020220"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-3032310020003122-1232023112113303-2231102203230123-0120203123001030-1213320133130311-3310113113022223-3033112232122102-3013030320223311"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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

<a id="canonical-1331033311003320-3223221300331003-3122120101102232-1021330121100112-3032303112000231-2123313021331222-2233310013313032-1302302131022222"></a>

<a id="canonical-0020022021310122-1130212223000201-3003210203013330-3131232230330300-2131312232132323-3111222031100021-1133003020230010-0201003311030301"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0102223120012332-2011012130221331-3011101321030011-3312113233032312-1112120101023013-3313103232103003-2130303130022221-0111021312303332"></a>

<a id="canonical-3110033330013213-2111323322122211-3311001300132120-0323123220010122-3002033013331320-3010130233230302-1010330230121333-0032220022213302"></a>

#### `api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Optional.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0301122111301232-2013220232300313-3011332300103100-3023223212103132-0232010111232230-0311320233220000-3301331203330121-2320222322032211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-3133023320332202-1221331211011301-3321000231201221-0121021022223000-3100013123312031-3320113100331223-0211010231313031-3313230233100300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inline rate limiter.

Additional upstream details:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
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
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330021022312322-0022111112200100-3003020132212100-1221131103133031-3023300311222013-3131321202213101-2312321000320323-2103222100333012"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter`

- [ref_user_id](resources--cdn_loadbalancer--reference--group-003.md#canonical-1011011011302321-2311311132312202-2123030002000120-1202000333330310-3200120111031101-1133122030132220-0232311301203220-0011321032000333): complete subsection reference.

<a id="canonical-2112103321001122-3331213213211023-0312230210011121-1231133132212121-3123320110110033-2011030212013223-0232212000112213-0222100003330030"></a>

<a id="canonical-2230100323223111-3321333131131110-2102333112321320-2331302302120202-0310202102202010-1310203000023311-3032320300000002-2320303331111130"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.threshold` property

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-2301322112032010-2102000230121022-3001312210301110-3203113222322311-2300132223233113-3020030332001131-2301002301220210-0321320233331300"></a>

<a id="canonical-3232031130301210-3203101012033311-1031311122100230-1213210133031310-1101132033203233-1133220333010223-3133131330002022-1300110110123003"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.unit` property

Type: `"string"`. Optional.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [use_http_lb_user_id](resources--cdn_loadbalancer--reference--group-003.md#canonical-1201130332230100-0321301323113303-2122103120013100-1101020111112013-2221023322110131-2330002310010132-0111121310222110-0121231033331210): complete subsection reference.

<a id="canonical-1011011011302321-2311311132312202-2123030002000120-1202000333330310-3200120111031101-1133122030132220-0232311301203220-0011321032000333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-003.md#canonical-0301122111301232-2013220232300313-3011332300103100-3023223212103132-0232010111232230-0311320233220000-3301331203330121-2320222322032211)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-2211101330301200-2301013100032210-0110311321121313-0222203320010232-2111131313200130-1131202110032131-0122223011001213-2331322101001110"></a>

Type: `"object"`. single nested block, Optional.

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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323231200222232-3121212203313112-3013130101020131-2300233030323331-3110022101202132-0022113030202310-0310133030213203-2023301103320103"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id`

<a id="canonical-3211221322021322-2120011022110201-3031320301333230-3030110033310330-0023132310111030-1221020323133123-2102302032232121-3001210000313312"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2132002321112111-2133221303012002-2332120030000113-0222130221223311-3022233220013223-2232122233110112-1103230020201310-2232302222200220"></a>

<a id="canonical-0103121022102231-1220033200131201-0100131302303212-2130202023320233-2132010031200210-3130230302310212-0313111200013033-2031212311230103"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2030110010030201-2210302100320111-2022030110212112-1331132131110203-1310303231000302-2210203210210020-3130100003222301-1311323333130022"></a>

<a id="canonical-2113332231323020-0311032213110230-2223230322020311-0231130311313302-1203110023300123-1013221001210203-3101031103303200-0000011021103300"></a>

#### `api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-1201130332230100-0321301323113303-2122103120013100-1101020111112013-2221023322110131-2330002310010132-0111121310222110-0121231033331210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-003.md#canonical-0301122111301232-2013220232300313-3011332300103100-3023223212103132-0232010111232230-0311320233220000-3301331203330121-2320222322032211)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-1021300113003013-1003203323002313-3133100321020023-2101030200111021-2032201200030023-0113100302332310-3121300230032203-2300200211200312"></a>

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
use_http_lb_user_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200220130112333-3222120132000131-3123212332333021-3010211231213321-0322000001030333-2000121322111101-2113333323032202-3203102320310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.ref_rate_limiter` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-2001111230132221-3102132332002100-2013133332131030-1200320002200013-2232231213321023-1223023030201100-1233131011233103-0333312201221122"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Additional upstream details:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131132013003331-2222110312303000-2322210201122111-3321211001100211-0223333112220111-2111132103201323-1202330100033003-0231012313201033"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.ref_rate_limiter`

<a id="canonical-1233103131330200-3221221012313110-3210211123330100-3103110122311321-0032100203111133-3020221100232012-3311220101323131-1121012021322011"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1303130032333021-2202133230103313-2203001123131311-1110222321310120-1030302321232001-1101020303033221-3131000020221123-3003323020011033"></a>

<a id="canonical-2011301031023012-2102012200012311-3132112221201303-1221331001011230-3032230200121010-0303000010320100-2111121023120303-2212123100001012"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2020222303331123-2213322221320231-2222333231211013-2131122203230313-2202202210102320-2210210311123233-1030331123203031-0013013030302032"></a>

<a id="canonical-1332030112211210-2202302130121200-1013010200023301-3320033131330332-0302233210132200-0210313120023012-3301020232110231-3110303012120200"></a>

#### `api_rate_limit.api_endpoint_rules.ref_rate_limiter.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-0203032330023212-3312123013031002-0231222013021132-2000200022002023-1112203121330302-3021122103222130-1110220103233322-2131311213230120"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Additional upstream details:

Request conditions for matching a rule.

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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123203332020301-0203102132332223-3213103310220220-3113003333133212-1133202300111132-1320203112200213-1233231230212030-0120011313010222"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher`

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-003.md#canonical-1011301330201232-2020222303113030-0203301033210011-3231000222200012-1303332132111230-0222332131133310-2322221030103013-2031100222131331): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-003.md#canonical-2320220212100031-3230133132110312-3030133003011203-2313310321320300-2011103001003203-1231303013333030-2231213000333120-3220321133102103): complete subsection reference.

- [jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-2022222021122001-3002003010030312-1020011133230112-3013003301012123-0223333311032301-1332111121323103-1331231131301233-0000112001032221): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-2123010330003131-1200001120312112-0103222130123312-2032231130103322-0030222120220210-0022123103310113-0022203231302333-0133112223211113): complete subsection reference.

<a id="canonical-1011301330201232-2020222303113030-0203301033210011-3231000222200012-1303332132111230-0222332131133310-2322221030103013-2031100222131331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-2312023121030331-2110022202312301-3020333031010202-2023223232222210-2030223200312302-3303101331100001-0132213112232120-1020230302233132"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122321022300102-0130121000113023-3323102330332320-3110303121002313-1230003210110102-3102021222020301-3003103203112031-0013013133001031"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-003.md#canonical-2103330200213001-1032030121110013-3323013030000030-2030302212133233-1223123001003112-2022003130012312-2221101121033130-1221133303220111): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-003.md#canonical-1311001033233112-1330302010220102-1303012013313103-0233233222003001-3102221311223231-3200203003312233-0210210113320130-3030302013311031): complete subsection reference.

<a id="canonical-3102100123201001-3301130002131201-2322120033313101-2201312200112301-3312003121223302-0230221222213311-0220112010211001-1003303111232111"></a>

<a id="canonical-1312010200121310-2102131003301003-0012122003222033-1320312320212012-2222032202310320-2200322130233330-0112012230312000-1232000211301222"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

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

- [item](resources--cdn_loadbalancer--reference--group-003.md#canonical-2301232101312033-3133011132123302-0021011220111133-3221020123232310-3112112020100322-0033333133211231-2332311300210210-0030132033112310): complete subsection reference.

<a id="canonical-2033111313120021-0002120033223111-3002223233103010-2013013222321113-1232210012202031-3233203110300100-2313130022232200-3123203310312201"></a>

<a id="canonical-1310331123333122-1332300130133103-1212230020300230-1122113202002213-3123101200101210-0322230221322300-3332310200320032-3231030232011233"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2103330200213001-1032030121110013-3323013030000030-2030302212133233-1223123001003112-2022003130012312-2221101121033130-1221133303220111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-003.md#canonical-1011301330201232-2020222303113030-0203301033210011-3231000222200012-1303332132111230-0222332131133310-2322221030103013-2031100222131331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3333033303121202-1223222103213302-1013010323213131-1230311331212221-3123213121333002-0102321323000031-2131321331312022-0031303330112111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311001033233112-1330302010220102-1303012013313103-0233233222003001-3102221311223231-3200203003312233-0210210113320130-3030302013311031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-003.md#canonical-1011301330201232-2020222303113030-0203301033210011-3231000222200012-1303332132111230-0222332131133310-2322221030103013-2031100222131331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-2131110311002222-3132311001120310-1030123300111132-1111103111210100-2300311102312030-1002221332203001-1011201033131101-3103230332331333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301232101312033-3133011132123302-0021011220111133-3221020123232310-3112112020100322-0033333133211231-2332311300210210-0030132033112310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-003.md#canonical-1011301330201232-2020222303113030-0203301033210011-3231000222200012-1303332132111230-0222332131133310-2322221030103013-2031100222131331)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-3030212000313113-2131201032103020-3331233002001333-2202203323210013-2101203223331012-1313301302213013-1203103323230130-2320101232020221"></a>

Type: `"object"`. single nested block, Optional.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213102300132310-3010221111113233-2330022020010303-2120111203332101-3023032323302012-1031200210301102-1202200313303302-0322103323003212"></a>

### Direct properties for `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item`

<a id="canonical-2123120031300222-2130322113021111-2232231112312331-3103033212310021-0233133133201200-3300200102020223-3332021201111132-3203113313023103"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3311300211302222-0232023112030202-2033220000210021-1231120320223100-0110323313302102-2330323313021110-3120202131132121-1101022100012320"></a>

<a id="canonical-3213332321102132-2300133312131223-3210031232030221-3302033313221030-3021332033203101-2000211321131221-0231130300111220-1233313021331202"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3232103313111102-3300322302113201-3133233300230001-0113111032320121-2031031313113202-2322200112303331-2032033101113121-1113220022030002"></a>

<a id="canonical-1321320100103000-0203022102123203-1220111300221031-3030222021302211-1123002112320113-3212030200300323-2320313213302000-3022211231000023"></a>

#### `api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2320220212100031-3230133132110312-3030133003011203-2313310321320300-2011103001003203-1231303013333030-2231213000333120-3220321133102103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.api_endpoint_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.api_endpoint_rules](resources--cdn_loadbalancer--reference--group-003.md#canonical-2313011122303231-3010000131222220-1312203122131110-1100233231211201-2210010311120200-0300303113023310-1000302010130222-2100023313211031)
- [api_rate_limit.api_endpoint_rules.request_matcher](resources--cdn_loadbalancer--reference--group-003.md#canonical-3220332110030133-2301210013001022-3112300310212323-1301221000233333-2113331132303213-0102323000313201-3000111131220323-3013231111311332)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-0133201313223223-3313320132211210-1012110321232121-2221112020301202-0222330100213103-0322321103030222-2002232103223330-0132102231013312"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```
