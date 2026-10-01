---
page_title: "xcsh_dns_lb_pool reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool reference."
---

# xcsh_dns_lb_pool reference

<a id="canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6877ed62bfac033e0ae7cd8bebc909ad7492751c11cbed982ec1784461ab8916"></a>

## Property reference — Property reference / c0d759dd51a2 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- Property reference

<a id="canonical-41f16983a87c9f0bdb7f7ec0fb8c81100d186a13917f81b2fbf39c15545aab46"></a>

## Direct properties — Property reference / c0d759dd51a2 / 3

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b): complete subsection reference.

- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a698a730e71b5de3a466aa7f87d6826a0a53788b7b54c1b5cd88cce717a9eecf): complete subsection reference.

<a id="canonical-6455b7e4a5564038b4393249afe0f9ea17fc64b7ce6f407005c304fd9c0152fe"></a>

<a id="canonical-b3cdd3887cb70d6de1cef4573fce07a775bbe2e9239b45d8762a83bda5306701"></a>

## annotations property — Property reference / c0d759dd51a2 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
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

- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28): complete subsection reference.

<a id="canonical-6965ddd6caa527792992e0db2403e08c95bb61db806f7076590532f84a127554"></a>

<a id="canonical-2895a5d59f42c3cca65d31c39a8c22c27d9fa1bb60b6ada78c269b19e304dcc9"></a>

## description property — Property reference / c0d759dd51a2 / 5

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

<a id="canonical-9ad4fbc4a6bd8413f9c894426f7ddfef1af2c84dec301a2c4fc7bd850f49deb1"></a>

<a id="canonical-8fa61e94d979a3de28b7467e7bd32b7d3146406ec414f68328f000c9878dea7a"></a>

## disable property — Property reference / c0d759dd51a2 / 6

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

<a id="canonical-fd730427341e3c514d11eb93ad075b142812aa5b6ce360a1ac06726f1417f079"></a>

<a id="canonical-87c1e2e67f3c5896623b860291dda1395ea2fa60453e318c7b28043c13cc92be"></a>

## id property — Property reference / c0d759dd51a2 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-91ee0dc9b397fce62ebe6831716eec29d508a9cccbb99b0a819dd41bba1b8d26"></a>

<a id="canonical-4293bde10ba126a9882d57744b30ee76afc517dc73d627acb7e0692bb2268a7e"></a>

## labels property — Property reference / c0d759dd51a2 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
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

<a id="canonical-1905a79fdae9ca7dd78cbcae80dd295ae5737c5455e26df7496a9091e3bcfdc8"></a>

<a id="canonical-560246373802c6dd643df37951d119964f611185d8881ab56cf1d7202a7aa8d1"></a>

## load_balancing_mode property — Property reference / c0d759dd51a2 / 9

Type: `"string"`. Optional, Computed.

\[Enum: ROUND\_ROBIN|RATIO\_MEMBER|STATIC\_PERSIST|PRIORITY\] - ROUND\_ROBIN: Round-Robin Round
Robin will ensure random equal distribution of requests among all pool members in a pool. -
RATIO\_MEMBER: Ratio-Member Ratio-Member performs load balancing of requests across the pool members
based on the ratio assigned to each pool member - STATIC\_PERSIST.. Possible values are
\`ROUND\_ROBIN\`, \`RATIO\_MEMBER\`, \`STATIC\_PERSIST\`, \`PRIORITY\`. Defaults to
\`ROUND\_ROBIN\`.

Upstream description:

&#8203;- ROUND\_ROBIN: Round-Robin

Round Robin will ensure random equal distribution of requests among all pool members in a pool.
&#8203;- RATIO\_MEMBER: Ratio-Member

Ratio-Member performs load balancing of requests across the pool members based on the ratio assigned
to each pool member &#8203;- STATIC\_PERSIST: Static-Persist

The Static Persist load balancing method uses the persist mask, with the source IP address of the
Local Domain Name Server (LDNS), in a deterministic algorithm to send requests to a specific pool
member. If the DNS resolver passes ECS (EDNS-Client-Subnet) information, then a hash of it will be
used, to send the client to the same pool member &#8203;- PRIORITY: Priority

The Priority load balancing method returns all available endpoints in a pool with the highest
priority. Pool Members have a priority value, starting from zero, where a lower value means a higher
priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a638e71fb1dc5aac1b7456c9208f6cb86eaf14b84667e5425b68df2490d56db8): complete subsection reference.

<a id="canonical-7760509eac38147a3d077aeaeec6549b39c8d8058b92e12f9de839159f87b2a2"></a>

<a id="canonical-9da2219eaf7b6de12ebfb9caa24d803e7bdc1638aa2aac0e39a69e6a00ad4ae0"></a>

## name property — Property reference / c0d759dd51a2 / 10

Type: `"string"`. Required.

Name of the DNS LB Pool. Must be unique within the namespace.

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

<a id="canonical-287203a39970efb5a5dd3c444ba93647f58e91a0e8f2a8038095f7477385d3a2"></a>

<a id="canonical-cf411c68c8cfd2014913c549ff3c3541e4778e9a4bf067bdfed9165b341d4c59"></a>

## namespace property — Property reference / c0d759dd51a2 / 11

Type: `"string"`. Optional, Computed.

Namespace for the DNS LB Pool. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

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

- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-c0d7a3808b846937daabc8d2ff91a5bb4f57beb6dab6e0faf0549f3ca8f74d7e): complete subsection reference.

- [timeouts](resources--dns_lb_pool--reference--group-001.md#canonical-31e8545911e0acc8af7f1a9adbfbe27fc9e8d7767f88361f35f13f9f6af750c0): complete subsection reference.

<a id="canonical-7dc3162cd1d188fafa5fb16809f714bf22e13a0e665a7737d05bf2d23bb66ba6"></a>

<a id="canonical-afa47e0851ab0181689f265e8c941912c8dd448989351744cbca2e561d6aaa0f"></a>

## ttl property — Property reference / c0d759dd51a2 / 12

Type: `"number"`. Optional, Computed.

\[OneOf: ttl, use\_rrset\_ttl\] Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default
&#8203;30) for responses from this pool.

Upstream description:

Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default 30) for responses from this pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

OneOf alternatives in this subsection:

- [ttl](resources--dns_lb_pool--reference--group-001.md#canonical-7dc3162cd1d188fafa5fb16809f714bf22e13a0e665a7737d05bf2d23bb66ba6)
- [use_rrset_ttl](resources--dns_lb_pool--reference--group-001.md#canonical-54fe72d5b9eabe62b5778a239546a5f426562fbd32efbabd3d1c5f955f3d58d7)

Select alternatives according to the provider validators above.

- [use_rrset_ttl](resources--dns_lb_pool--reference--group-001.md#canonical-61787dc50e3037e058b416ce7241c0eef43ae674606b147ed012be99217e997d): complete subsection reference.

<a id="canonical-50de60756b3ddc9c790a77f565af43914fd15a9bd96aede2d882a47dd71d762f"></a>

## All schema paths — Property reference / c0d759dd51a2 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `a_pool` | [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-ef43b2af1b10e1661a98322cd52b2e32c4a25dae7bb4a9a313170cfb6e71c341) |
| `a_pool.disable_health_check` | [a_pool.disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-a51ca0217c575e37d1a0c875f9d7d4514e4a5b3ca8aab0b672445de3e07f7fda) |
| `a_pool.health_check` | [a_pool.health_check](resources--dns_lb_pool--reference--group-001.md#canonical-7f6b498bebc130bc552d835158b17e2ecd228f54fe9b755198646630452bac6e) |
| `a_pool.health_check.name` | [a_pool.health_check.name](resources--dns_lb_pool--reference--group-001.md#canonical-2300710c80c6fa8239cc2ebeea273d5610f73dbc394db837fbc213d2e145bae6) |
| `a_pool.health_check.namespace` | [a_pool.health_check.namespace](resources--dns_lb_pool--reference--group-001.md#canonical-090de96b104bda57b64fc340c430b378181edbfa80d62053fd23b0cf2672e31c) |
| `a_pool.health_check.tenant` | [a_pool.health_check.tenant](resources--dns_lb_pool--reference--group-001.md#canonical-8478971634702b101e42a720afcd839545a217db75befbf1e3296da6d28fbd7f) |
| `a_pool.max_answers` | [a_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-1af97733c8b7101f0656264b79a086d4ba11df777d4307262cf28d31d471f781) |
| `a_pool.members` | [a_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-542e8d0f04d57321eeaa23b07087ed17a2b8d8e78359566cdd0cf2d1bce77cd1) |
| `a_pool.members.disable_spec` | [a_pool.members.disable_spec](resources--dns_lb_pool--reference--group-001.md#canonical-d30b9721ebf26cef0f53118d9088784a206bebdbd9eed0229fa951a03f04071b) |
| `a_pool.members.ip_endpoint` | [a_pool.members.ip_endpoint](resources--dns_lb_pool--reference--group-001.md#canonical-33822e54e15097816474022ea59808f5189a8b9884d021d475ca9c6e460969ec) |
| `a_pool.members.name` | [a_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-744dbe2ebee7f8896a95169b082fb5d1ada475a9583e6f4c2d20bcf83572d082) |
| `a_pool.members.priority` | [a_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-fce8f9d153d734da26c369607fcbbbd43935e0c83dfc463018ae482143feabc5) |
| `a_pool.members.ratio` | [a_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-543cc26ff351b997a3c041655e8146374b81ef9298d2214e4c1e2ada8512821a) |
| `aaaa_pool` | [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-334dd7fb48ce39facf7f614e84d98eebc2278b298b1f93bc238e44fd7b6555f2) |
| `aaaa_pool.max_answers` | [aaaa_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-ca575e84b44c180a7aa2a99f6bfaddd1d61d625135e4cfaf8d3b8e285bb17cbe) |
| `aaaa_pool.members` | [aaaa_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-2bebbe6fc46932eccc0c0ff43997a10773568d7c238bec33fbefdbe3232c50c0) |
| `aaaa_pool.members.disable_spec` | [aaaa_pool.members.disable_spec](resources--dns_lb_pool--reference--group-001.md#canonical-2b1d5027b712822e0b5abc2861203b2cff8721b04b30e42cadc49173fb2f62c0) |
| `aaaa_pool.members.ip_endpoint` | [aaaa_pool.members.ip_endpoint](resources--dns_lb_pool--reference--group-001.md#canonical-6db872f9cff7420f122c4d295a04b864efd3638eb3c53979c4fecfc9af8aa760) |
| `aaaa_pool.members.name` | [aaaa_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-1505ed0b82ecebe359a6dc1682f3f319116640308dd465db869d4deea309ad2a) |
| `aaaa_pool.members.priority` | [aaaa_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-00feedc3d45bd49c5866bb16aa28f509288dc62a19bd408c0636625aabade1e5) |
| `aaaa_pool.members.ratio` | [aaaa_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-4076d893967f5c1208bf6db6cb96456f40ce479cb5160008cdda551edd304987) |
| `annotations` | [annotations](resources--dns_lb_pool--reference--group-001.md#canonical-6455b7e4a5564038b4393249afe0f9ea17fc64b7ce6f407005c304fd9c0152fe) |
| `cname_pool` | [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-521d052a43b461025f590b553379cde8aa84d11cf75747f59eea1a61dc1b944e) |
| `cname_pool.disable_health_check` | [cname_pool.disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-85491cf026bee091a0a71650364fb1a96b917a5d3cc5e55dfbcd20262dd88aed) |
| `cname_pool.health_check` | [cname_pool.health_check](resources--dns_lb_pool--reference--group-001.md#canonical-7915a629e342517331c1f5a88bce86e2486485c56497e72616df64e345ceb249) |
| `cname_pool.health_check.name` | [cname_pool.health_check.name](resources--dns_lb_pool--reference--group-001.md#canonical-2860a13268af499d3ce664696ab30453ce6f34249c354a083d5bf3f2ebb46a8c) |
| `cname_pool.health_check.namespace` | [cname_pool.health_check.namespace](resources--dns_lb_pool--reference--group-001.md#canonical-71a0dddd9f7ff75e3873f4716dc269cb2ba4b5506e5041e20948b5df12c0e044) |
| `cname_pool.health_check.tenant` | [cname_pool.health_check.tenant](resources--dns_lb_pool--reference--group-001.md#canonical-d4120d946936e90223f6f733ef1cf8c4c80c0dd5e8888d610e8dbe56e67ee759) |
| `cname_pool.members` | [cname_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-dc39302a454e3815fa7ded418c7ab7845deaa5557453d3f189983d14e0742a05) |
| `cname_pool.members.domain` | [cname_pool.members.domain](resources--dns_lb_pool--reference--group-001.md#canonical-0be7a547ce039cff9aac16aff00209eacf97ac1ec768c23f587612dab13d8c85) |
| `cname_pool.members.final_translation` | [cname_pool.members.final_translation](resources--dns_lb_pool--reference--group-001.md#canonical-259a3c3cd499fc1194586387f499b85906967579af68b0277324790327cf3be2) |
| `cname_pool.members.name` | [cname_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-fd6c398e386751289eccf4c5dd83caf4de1c3fa514d2a7380521ed2dc8c973f8) |
| `cname_pool.members.priority` | [cname_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-31134b7fe7fa159a5dcf8be2f23d3d2c47bfccc7b5ea8db9e79423a307e2e9bc) |
| `cname_pool.members.ratio` | [cname_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-8b0badc46a2648cd61fa7c81ba081950865ca8e3fefd2fea0b29f3be4b2783a4) |
| `description` | [description](resources--dns_lb_pool--reference--group-001.md#canonical-6965ddd6caa527792992e0db2403e08c95bb61db806f7076590532f84a127554) |
| `disable` | [disable](resources--dns_lb_pool--reference--group-001.md#canonical-9ad4fbc4a6bd8413f9c894426f7ddfef1af2c84dec301a2c4fc7bd850f49deb1) |
| `id` | [id](resources--dns_lb_pool--reference--group-001.md#canonical-fd730427341e3c514d11eb93ad075b142812aa5b6ce360a1ac06726f1417f079) |
| `labels` | [labels](resources--dns_lb_pool--reference--group-001.md#canonical-91ee0dc9b397fce62ebe6831716eec29d508a9cccbb99b0a819dd41bba1b8d26) |
| `load_balancing_mode` | [load_balancing_mode](resources--dns_lb_pool--reference--group-001.md#canonical-1905a79fdae9ca7dd78cbcae80dd295ae5737c5455e26df7496a9091e3bcfdc8) |
| `mx_pool` | [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-906f918bf77ffc479ddec803f7fbdea3b4250ba18337ab2278c7bdaf78953e3e) |
| `mx_pool.max_answers` | [mx_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-34517024bb29d5515cf397a07779493daf31de066b3b5b293539f35c9d2daf5a) |
| `mx_pool.members` | [mx_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-890aa25e3211a9a02c4c524a44d3625c3af13e07beadbb9acf9d9e6e79bc8f61) |
| `mx_pool.members.domain` | [mx_pool.members.domain](resources--dns_lb_pool--reference--group-001.md#canonical-82b03dbe3972d7ede1b73b986ede2d5d27c0428d111949e02841d6e31be3d101) |
| `mx_pool.members.name` | [mx_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-3498ff24504acfed8961ba103891a6b8a4b5b982524f355b510b55cfb097b65f) |
| `mx_pool.members.priority` | [mx_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-8df2685c9942c914fb580e3390b95b1a47477426d5c34e52f7196bc52943b682) |
| `mx_pool.members.ratio` | [mx_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-636213f4df31a1023b2377c558713727b178e04862e319b82bc96629b4c36fa0) |
| `name` | [name](resources--dns_lb_pool--reference--group-001.md#canonical-7760509eac38147a3d077aeaeec6549b39c8d8058b92e12f9de839159f87b2a2) |
| `namespace` | [namespace](resources--dns_lb_pool--reference--group-001.md#canonical-287203a39970efb5a5dd3c444ba93647f58e91a0e8f2a8038095f7477385d3a2) |
| `srv_pool` | [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1e55f944d877db53b756428b7ba788cf987c4175c70219a7332f28cdb478628b) |
| `srv_pool.max_answers` | [srv_pool.max_answers](resources--dns_lb_pool--reference--group-001.md#canonical-89235113e728a4742f443f9e30a19d06be0d24955f2b744ff498d42f36a27834) |
| `srv_pool.members` | [srv_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-2bf08d77a1d606979e158dfbf07abc18ca17f8c4e0f12a3018f6809e1ca987d0) |
| `srv_pool.members.final_translation` | [srv_pool.members.final_translation](resources--dns_lb_pool--reference--group-001.md#canonical-22824b918f5a918ad3c7034dc666108f2bb5153e87c22be243e1c2c2dea0427b) |
| `srv_pool.members.name` | [srv_pool.members.name](resources--dns_lb_pool--reference--group-001.md#canonical-0652f1137a1ef1b111ef98ae49eff5a513975c5fee34921368eaa81ed233c248) |
| `srv_pool.members.port` | [srv_pool.members.port](resources--dns_lb_pool--reference--group-001.md#canonical-07f998d779c59a10be8ec65ac1b4867e3dda4df191fe248d8cb32160fe73052b) |
| `srv_pool.members.priority` | [srv_pool.members.priority](resources--dns_lb_pool--reference--group-001.md#canonical-2743ac9e875ec036e195268bed319b777ea0a9e6bb09d2b0e9d0f32550bbbf8f) |
| `srv_pool.members.ratio` | [srv_pool.members.ratio](resources--dns_lb_pool--reference--group-001.md#canonical-12befb6bd7292ae6db3b84fb56b7954ca37d85d2d8d796c11d6541915fd47b3e) |
| `srv_pool.members.target` | [srv_pool.members.target](resources--dns_lb_pool--reference--group-001.md#canonical-09f053765b89b4c795dbde5961488485fa76944fb13aba42e272eef5ff5a1c1c) |
| `srv_pool.members.weight` | [srv_pool.members.weight](resources--dns_lb_pool--reference--group-001.md#canonical-76e42f27e971d346a6d552f0b4c9ea4025323d05765467b186dce8ebfbae7579) |
| `timeouts` | [timeouts](resources--dns_lb_pool--reference--group-001.md#canonical-6f4246cb70c5ccb0295ed7b4836143651c5e1f2d116e65e3fcf74ff0d2067f10) |
| `timeouts.create` | [timeouts.create](resources--dns_lb_pool--reference--group-001.md#canonical-016b101a9b29b309b5700703390daeee9daf36b57d746e4a1962ce663d0d8cb9) |
| `timeouts.delete` | [timeouts.delete](resources--dns_lb_pool--reference--group-001.md#canonical-65d2328072579843187b02c888de65f8d186453b6ba9f0ff9e7a476593f2aebf) |
| `timeouts.read` | [timeouts.read](resources--dns_lb_pool--reference--group-001.md#canonical-e267be24310cb4518951651293ecf4a1ff1106a7e9e40db1803b4b80136213b2) |
| `timeouts.update` | [timeouts.update](resources--dns_lb_pool--reference--group-001.md#canonical-9239dbc7517403801cf792acaa1447ef6960f5dae4cb5951c6add569dae36539) |
| `ttl` | [ttl](resources--dns_lb_pool--reference--group-001.md#canonical-7dc3162cd1d188fafa5fb16809f714bf22e13a0e665a7737d05bf2d23bb66ba6) |
| `use_rrset_ttl` | [use_rrset_ttl](resources--dns_lb_pool--reference--group-001.md#canonical-54fe72d5b9eabe62b5778a239546a5f426562fbd32efbabd3d1c5f955f3d58d7) |

<a id="canonical-087a0288516466148f3200239b0c6211bf5c1ec3e6bd5c274fc9525cf2f6df8a"></a>

## Next pages — Property reference / c0d759dd51a2 / 14

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b)
- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a698a730e71b5de3a466aa7f87d6826a0a53788b7b54c1b5cd88cce717a9eecf)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28)
- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a638e71fb1dc5aac1b7456c9208f6cb86eaf14b84667e5425b68df2490d56db8)
- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-c0d7a3808b846937daabc8d2ff91a5bb4f57beb6dab6e0faf0549f3ca8f74d7e)
- [timeouts](resources--dns_lb_pool--reference--group-001.md#canonical-31e8545911e0acc8af7f1a9adbfbe27fc9e8d7767f88361f35f13f9f6af750c0)
- [use_rrset_ttl](resources--dns_lb_pool--reference--group-001.md#canonical-61787dc50e3037e058b416ce7241c0eef43ae674606b147ed012be99217e997d)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddb649f1886265ce905c70afb2c42a7fb030539c0b5571f44ea3c0b2941c803e"></a>

## a_pool — a_pool / fd550c2fd0d0 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- a_pool

<a id="canonical-ef43b2af1b10e1661a98322cd52b2e32c4a25dae7bb4a9a313170cfb6e71c341"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: a\_pool, aaaa\_pool, cname\_pool, mx\_pool, srv\_pool\] Pool for A Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members"),
  validators.ConflictingObjectAttributes("disable_health_check",
    "health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

OneOf alternatives in this subsection:

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-ef43b2af1b10e1661a98322cd52b2e32c4a25dae7bb4a9a313170cfb6e71c341)
- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-334dd7fb48ce39facf7f614e84d98eebc2278b298b1f93bc238e44fd7b6555f2)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-521d052a43b461025f590b553379cde8aa84d11cf75747f59eea1a61dc1b944e)
- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-906f918bf77ffc479ddec803f7fbdea3b4250ba18337ab2278c7bdaf78953e3e)
- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1e55f944d877db53b756428b7ba788cf987c4175c70219a7332f28cdb478628b)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
a_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad4a146896d4b1d5c95daeba6892579d85e4d739d22719e51f68131a11388889"></a>

## Direct properties — a_pool / fd550c2fd0d0 / 3

- [disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-249c29c0a2b904d314eee449c3806dfa74271c00778a4a447e59a40ccdfaf895): complete subsection reference.

- [health_check](resources--dns_lb_pool--reference--group-001.md#canonical-3bd373e88b4c65067e4635cabcaccc819128ede97c867dfcbd49319edec885e9): complete subsection reference.

<a id="canonical-1af97733c8b7101f0656264b79a086d4ba11df777d4307262cf28d31d471f781"></a>

<a id="canonical-4e70e195898b81a0f678b3bbaca2bb35556f0a36ccf17881fb5533dbbc074660"></a>

## max_answers property — a_pool / fd550c2fd0d0 / 4

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-9a2957380107b511e4832835bb43cc9407fa6a40fec70b8aa59ea52a529c0e1c): complete subsection reference.

<a id="canonical-418c4d084c120c606509e91efb620a9f95815932562106f585d6e82c07a48361"></a>

## Next pages — a_pool / fd550c2fd0d0 / 5

- [a_pool.disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-249c29c0a2b904d314eee449c3806dfa74271c00778a4a447e59a40ccdfaf895)
- [a_pool.health_check](resources--dns_lb_pool--reference--group-001.md#canonical-3bd373e88b4c65067e4635cabcaccc819128ede97c867dfcbd49319edec885e9)
- [a_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-9a2957380107b511e4832835bb43cc9407fa6a40fec70b8aa59ea52a529c0e1c)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-249c29c0a2b904d314eee449c3806dfa74271c00778a4a447e59a40ccdfaf895"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1229d6c4a917fd1b5977d900057f119e5aa79c05723fa70cac0e765919b0b8d0"></a>

## a_pool.disable_health_check — a_pool.disable_health_check / 763fe625a4e5 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b)
- a_pool.disable_health_check

<a id="canonical-a51ca0217c575e37d1a0c875f9d7d4514e4a5b3ca8aab0b672445de3e07f7fda"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable health check.

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
disable_health_check = {}
```

<a id="canonical-e266decad4244f28fb26f4a618f9b0e931c53b7189f376aeb1f5965439232397"></a>

## Direct properties — a_pool.disable_health_check / 763fe625a4e5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4994237306da8b24d329c386fd51d16bb16df4cb84660818694edb58719ccf27"></a>

## Next pages — a_pool.disable_health_check / 763fe625a4e5 / 4

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-3bd373e88b4c65067e4635cabcaccc819128ede97c867dfcbd49319edec885e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d596b85c02737f9a273c01f4bcb2c110fb296e78aaa099233109c93ebaa20ca8"></a>

## a_pool.health_check — a_pool.health_check / 924a8b83bcbb / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b)
- a_pool.health_check

<a id="canonical-7f6b498bebc130bc552d835158b17e2ecd228f54fe9b755198646630452bac6e"></a>

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
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3d101f6004e2e1f291aa31fd9d0048ecb6b8d17b1e71b79d72472f9896e3b357"></a>

## Direct properties — a_pool.health_check / 924a8b83bcbb / 3

<a id="canonical-2300710c80c6fa8239cc2ebeea273d5610f73dbc394db837fbc213d2e145bae6"></a>

<a id="canonical-bac12aefb48e8edc6cd6f36c7c5d5ced520d89a904ba6211ce866720c95e8349"></a>

## name property — a_pool.health_check / 924a8b83bcbb / 4

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

<a id="canonical-090de96b104bda57b64fc340c430b378181edbfa80d62053fd23b0cf2672e31c"></a>

<a id="canonical-20e8cd20280fa87fd5e0b750306380cd4e26e3df8221e2d3644195c9a604ca39"></a>

## namespace property — a_pool.health_check / 924a8b83bcbb / 5

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

<a id="canonical-8478971634702b101e42a720afcd839545a217db75befbf1e3296da6d28fbd7f"></a>

<a id="canonical-061b7f99d30985aa52aed145a1d2f846d581700a6727f5a62d555408d044fa19"></a>

## tenant property — a_pool.health_check / 924a8b83bcbb / 6

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

<a id="canonical-8a339074e75f3930becca2aa2e31a970262dc5f1d65e91da3845c06a3f5fb0dd"></a>

## Next pages — a_pool.health_check / 924a8b83bcbb / 7

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-9a2957380107b511e4832835bb43cc9407fa6a40fec70b8aa59ea52a529c0e1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f0289945676cd09147e1d054af4bbf3077f4c8aa70bb59f45a88d5e82bb69bed"></a>

## a_pool.members — a_pool.members / 04c473480944 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b)
- a_pool.members

<a id="canonical-542e8d0f04d57321eeaa23b07087ed17a2b8d8e78359566cdd0cf2d1bce77cd1"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_endpoint")}
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
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-b548158e340beaaf84f816645c46814252b32a4dd4c7289340d2b128fc079217"></a>

## Direct properties — a_pool.members / 04c473480944 / 3

<a id="canonical-d30b9721ebf26cef0f53118d9088784a206bebdbd9eed0229fa951a03f04071b"></a>

<a id="canonical-5f5e2bce529c5ab4dba4241a918f2ef138bfdf6345475c416d1d3256f946fd8c"></a>

## disable_spec property — a_pool.members / 04c473480944 / 4

Type: `"bool"`. Optional.

Value of true will disable the pool-member.

<a id="canonical-33822e54e15097816474022ea59808f5189a8b9884d021d475ca9c6e460969ec"></a>

<a id="canonical-74b0e945bac99f4f77c34b2a1b910f473a19024179aaeebc6852247386401d23"></a>

## ip_endpoint property — a_pool.members / 04c473480944 / 5

Type: `"string"`. Optional.

Public IP. Public IP address.

Upstream description:

Public IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-744dbe2ebee7f8896a95169b082fb5d1ada475a9583e6f4c2d20bcf83572d082"></a>

<a id="canonical-1ca33598aa3be3c2e9b1f01330ec91dbb57edaa78382e1cca23adcf894a4cbe3"></a>

## name property — a_pool.members / 04c473480944 / 6

Type: `"string"`. Optional.

Name. Pool member name.

Upstream description:

Pool member name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-fce8f9d153d734da26c369607fcbbbd43935e0c83dfc463018ae482143feabc5"></a>

<a id="canonical-4db33a4d89cb10c4a44796c6c7ae820b8cb8e935eea4fbefb2f39a195e7b70c4"></a>

## priority property — a_pool.members / 04c473480944 / 7

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-543cc26ff351b997a3c041655e8146374b81ef9298d2214e4c1e2ada8512821a"></a>

<a id="canonical-e244c56ca552d267ce8a8588fdf6e57fe6ee103324df1895da257edc3729c2c6"></a>

## ratio property — a_pool.members / 04c473480944 / 8

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Ratio-Member.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-5ba9ac6b61cf5c25e05875d6449078d4fc976eb48e171a8e73c18fc5c35cf566"></a>

## Next pages — a_pool.members / 04c473480944 / 9

- [a_pool](resources--dns_lb_pool--reference--group-001.md#canonical-1bd7d2cc9b92b92e886552a70fc317dcc76db6b3a89aef111ec878ce008f430b)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-a698a730e71b5de3a466aa7f87d6826a0a53788b7b54c1b5cd88cce717a9eecf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce3a162b89a585b19b1cff7e44670298a3c59f2e5dece3987b295ec6c314afbb"></a>

## aaaa_pool — aaaa_pool / b6fd0fc2bdcf / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- aaaa_pool

<a id="canonical-334dd7fb48ce39facf7f614e84d98eebc2278b298b1f93bc238e44fd7b6555f2"></a>

Type: `"object"`. single nested block, Optional.

Pool for AAAA Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members")}
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
aaaa_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-be6a379fa2a16e9a1fcd6d2d2d9c9e31565ff7f7592ff15afdccd20d3aa3e5ce"></a>

## Direct properties — aaaa_pool / b6fd0fc2bdcf / 3

<a id="canonical-ca575e84b44c180a7aa2a99f6bfaddd1d61d625135e4cfaf8d3b8e285bb17cbe"></a>

<a id="canonical-64a75c21fa8bacd3735b315010e10889acbf8fb3fda520d39e3a876d45c4c35e"></a>

## max_answers property — aaaa_pool / b6fd0fc2bdcf / 4

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-a5363ba2a4e437a531cff8fbb6628a481501a17e84180de01ee56db673340f59): complete subsection reference.

<a id="canonical-1a0eb2be66cc6de218f13e08ffb56e45ab9d488aecbef3ec351026c29ffd64a0"></a>

## Next pages — aaaa_pool / b6fd0fc2bdcf / 5

- [aaaa_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-a5363ba2a4e437a531cff8fbb6628a481501a17e84180de01ee56db673340f59)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-a5363ba2a4e437a531cff8fbb6628a481501a17e84180de01ee56db673340f59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c859afb8d1a11e8132a6870d22b20a7c258d732d7acaf75500816158913dcdb8"></a>

## aaaa_pool.members — aaaa_pool.members / 8199fa3b455d / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a698a730e71b5de3a466aa7f87d6826a0a53788b7b54c1b5cd88cce717a9eecf)
- aaaa_pool.members

<a id="canonical-2bebbe6fc46932eccc0c0ff43997a10773568d7c238bec33fbefdbe3232c50c0"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_endpoint")}
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
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-03adb9049c031291e5ec2d3009fd6d57811d3332e57d5957d56f5efdf91035c5"></a>

## Direct properties — aaaa_pool.members / 8199fa3b455d / 3

<a id="canonical-2b1d5027b712822e0b5abc2861203b2cff8721b04b30e42cadc49173fb2f62c0"></a>

<a id="canonical-8d89a8b81433af9892a1c503bedbbc831a589f5ab743c8322866d9b04e9af197"></a>

## disable_spec property — aaaa_pool.members / 8199fa3b455d / 4

Type: `"bool"`. Optional.

Value of true will disable the pool-member.

<a id="canonical-6db872f9cff7420f122c4d295a04b864efd3638eb3c53979c4fecfc9af8aa760"></a>

<a id="canonical-018118c3d69d88ce6d0876db1cbf7c4dc9edbe192cfd663fb94b195ec802cbc7"></a>

## ip_endpoint property — aaaa_pool.members / 8199fa3b455d / 5

Type: `"string"`. Optional.

Public IP. Public IP address.

Upstream description:

Public IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1505ed0b82ecebe359a6dc1682f3f319116640308dd465db869d4deea309ad2a"></a>

<a id="canonical-88c10f798c96aa0338dea4b23ccc4aa8b38011228366dc2bc45acd7a4304d29d"></a>

## name property — aaaa_pool.members / 8199fa3b455d / 6

Type: `"string"`. Optional.

Name. Pool member name.

Upstream description:

Pool member name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-00feedc3d45bd49c5866bb16aa28f509288dc62a19bd408c0636625aabade1e5"></a>

<a id="canonical-a77114e59815a4db461606ef89c91d2883581a86f9af95dfc67ee2b096e0a6d3"></a>

## priority property — aaaa_pool.members / 8199fa3b455d / 7

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-4076d893967f5c1208bf6db6cb96456f40ce479cb5160008cdda551edd304987"></a>

<a id="canonical-ee58cc8486071abdb397f626c046e46d15e683c43e8544accd52c11b75e13896"></a>

## ratio property — aaaa_pool.members / 8199fa3b455d / 8

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Ratio-Member.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-7afb9fb1aa49084bfa03914c848e00e3b7a850c31b2be32234fb5ee733574dad"></a>

## Next pages — aaaa_pool.members / 8199fa3b455d / 9

- [aaaa_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a698a730e71b5de3a466aa7f87d6826a0a53788b7b54c1b5cd88cce717a9eecf)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a900fb27f70c55f70cd29bfee7f4c7486c599dd5390c730e8e5c386595870cc"></a>

## cname_pool — cname_pool / 90a23991313a / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- cname_pool

<a id="canonical-521d052a43b461025f590b553379cde8aa84d11cf75747f59eea1a61dc1b944e"></a>

Type: `"object"`. single nested block, Optional.

Pool for CNAME Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("members"),
  validators.ConflictingObjectAttributes("disable_health_check",
    "health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

Terraform syntax:

```terraform
cname_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-39a5b0c04bb5c6c73e7c4796eb7c9efaeb402203c863f1c199ca295734c51b57"></a>

## Direct properties — cname_pool / 90a23991313a / 3

- [disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-a7d9901eb9c8067aed21be8f85924b90eae364ee817130476bc676418b7ff5b2): complete subsection reference.

- [health_check](resources--dns_lb_pool--reference--group-001.md#canonical-f80d4aa5aed1234ec854299681f56f287396177915474a49a22630d9f8fb4e83): complete subsection reference.

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-611e675a7fd968416f3de93c5a85f2ea60bffd0f0be2492fa59f7d42f2472cd8): complete subsection reference.

<a id="canonical-92f1b1480023132b9b8d4a06e612c641ae2ca56bb73a7bf549da06c562315def"></a>

## Next pages — cname_pool / 90a23991313a / 4

- [cname_pool.disable_health_check](resources--dns_lb_pool--reference--group-001.md#canonical-a7d9901eb9c8067aed21be8f85924b90eae364ee817130476bc676418b7ff5b2)
- [cname_pool.health_check](resources--dns_lb_pool--reference--group-001.md#canonical-f80d4aa5aed1234ec854299681f56f287396177915474a49a22630d9f8fb4e83)
- [cname_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-611e675a7fd968416f3de93c5a85f2ea60bffd0f0be2492fa59f7d42f2472cd8)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-a7d9901eb9c8067aed21be8f85924b90eae364ee817130476bc676418b7ff5b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0077178a1b75eb286acc96c31d529eedb37938b37b6fd8c46c270220594b21ee"></a>

## cname_pool.disable_health_check — cname_pool.disable_health_check / 1dab7b57c634 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28)
- cname_pool.disable_health_check

<a id="canonical-85491cf026bee091a0a71650364fb1a96b917a5d3cc5e55dfbcd20262dd88aed"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable health check.

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
disable_health_check = {}
```

<a id="canonical-8f42b4ab1b0c7f78eb4be3e5441d5437b7e26fc661bfcfd6a77320b3ce5e0137"></a>

## Direct properties — cname_pool.disable_health_check / 1dab7b57c634 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-998a2be69afcd229221450a91305a6c490c8ecbb29af1db97bdc469a8d1b465d"></a>

## Next pages — cname_pool.disable_health_check / 1dab7b57c634 / 4

- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-f80d4aa5aed1234ec854299681f56f287396177915474a49a22630d9f8fb4e83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-638764a333bd7f96fceadb8eb75c13349ec4c2293723ea02b85d48e030bf8d8d"></a>

## cname_pool.health_check — cname_pool.health_check / 4868a7a94b77 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28)
- cname_pool.health_check

<a id="canonical-7915a629e342517331c1f5a88bce86e2486485c56497e72616df64e345ceb249"></a>

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
health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-caa3f82bf17e71a65834a7a6439a4b062d812daebf68a791f8105ba439749cd8"></a>

## Direct properties — cname_pool.health_check / 4868a7a94b77 / 3

<a id="canonical-2860a13268af499d3ce664696ab30453ce6f34249c354a083d5bf3f2ebb46a8c"></a>

<a id="canonical-40a9ecd5448f86ccf134cb38157836243cd2725a74ec1f4bff7d9f852bc530f8"></a>

## name property — cname_pool.health_check / 4868a7a94b77 / 4

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

<a id="canonical-71a0dddd9f7ff75e3873f4716dc269cb2ba4b5506e5041e20948b5df12c0e044"></a>

<a id="canonical-631f01d16d9a5d0dd6f5c70713150ecef767992752935859587f9cc29f0ac747"></a>

## namespace property — cname_pool.health_check / 4868a7a94b77 / 5

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

<a id="canonical-d4120d946936e90223f6f733ef1cf8c4c80c0dd5e8888d610e8dbe56e67ee759"></a>

<a id="canonical-4f2c866a9d20181182d4fd87659b9e0aa6d3af7863d5ed24e7a86cf25a6c2393"></a>

## tenant property — cname_pool.health_check / 4868a7a94b77 / 6

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

<a id="canonical-37fa6d42ab55783d2640fe2e2cf01f66cba19e0368e33c22804fbdf809b7b6e5"></a>

## Next pages — cname_pool.health_check / 4868a7a94b77 / 7

- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-611e675a7fd968416f3de93c5a85f2ea60bffd0f0be2492fa59f7d42f2472cd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e630f3d250af48dbc896578c1d51b2e3f7774d696b2e764309f4517928c2ccf"></a>

## cname_pool.members — cname_pool.members / fc632c33ba59 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28)
- cname_pool.members

<a id="canonical-dc39302a454e3815fa7ded418c7ab7845deaa5557453d3f189983d14e0742a05"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-c97c8783077235d69b7664cee7264fd96eba5111910697948bbf88c4decefc81"></a>

## Direct properties — cname_pool.members / fc632c33ba59 / 3

<a id="canonical-0be7a547ce039cff9aac16aff00209eacf97ac1ec768c23f587612dab13d8c85"></a>

<a id="canonical-33bbb201ced63d9b121b35b498ee78dc3c4f6a79516387b7dd397d1c1259f1d7"></a>

## domain property — cname_pool.members / fc632c33ba59 / 4

Type: `"string"`. Optional.

Specifies the fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-259a3c3cd499fc1194586387f499b85906967579af68b0277324790327cf3be2"></a>

<a id="canonical-c26e3dd57260cd8968ade11920fcea23a16af1b0dc86f4525eddbdf7efca5cbc"></a>

## final_translation property — cname_pool.members / fc632c33ba59 / 5

Type: `"bool"`. Optional.

If this flag is true, the CNAME record will not be translated further.

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

<a id="canonical-fd6c398e386751289eccf4c5dd83caf4de1c3fa514d2a7380521ed2dc8c973f8"></a>

<a id="canonical-f4ba12a4299f3039f5edb0a7c0d48e08195e010d7a3ebcdf39ed65b628571f94"></a>

## name property — cname_pool.members / fc632c33ba59 / 6

Type: `"string"`. Optional.

Name. Pool member name.

Upstream description:

Pool member name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-31134b7fe7fa159a5dcf8be2f23d3d2c47bfccc7b5ea8db9e79423a307e2e9bc"></a>

<a id="canonical-b716feb7140556011ca1ce4b8b22532602d8a56d76a79eb03f719009cace375f"></a>

## priority property — cname_pool.members / fc632c33ba59 / 7

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is
routed to pool members. The lower the number, the higher the priority, making those members active
while higher-numbered members act as backups.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-8b0badc46a2648cd61fa7c81ba081950865ca8e3fefd2fea0b29f3be4b2783a4"></a>

<a id="canonical-ae8e726cef28403e764483471fc6b7c0fc29bde7ae9c764c70b2cdef9a97d8bb"></a>

## ratio property — cname_pool.members / fc632c33ba59 / 8

Type: `"number"`. Optional.

Used if the pool’s load balancing mode is set to Ratio-Member.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-5231daa250eed56856d15ed5126f7490b51c51b3571bfac76f9b3f91fff17a75"></a>

## Next pages — cname_pool.members / fc632c33ba59 / 9

- [cname_pool](resources--dns_lb_pool--reference--group-001.md#canonical-9270d8e01c1b433763e199a17cec6350868832fe33f9e1f6131bdb737b6c9b28)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-a638e71fb1dc5aac1b7456c9208f6cb86eaf14b84667e5425b68df2490d56db8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42b0bcd6aaa5cc05516037891e9b5d102c49e9bf6f1d42a97ad3c696920b7c23"></a>

## mx_pool — mx_pool / d35144f9505e / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- mx_pool

<a id="canonical-906f918bf77ffc479ddec803f7fbdea3b4250ba18337ab2278c7bdaf78953e3e"></a>

Type: `"object"`. single nested block, Optional.

Pool for MX Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members")}
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
mx_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-12b9d499d937d31f4a0b1c9b0d9c9afd5e99d8d047d209b8d9cbde2797257d92"></a>

## Direct properties — mx_pool / d35144f9505e / 3

<a id="canonical-34517024bb29d5515cf397a07779493daf31de066b3b5b293539f35c9d2daf5a"></a>

<a id="canonical-1b47eef1128b2c071f579687803aa56ee2e990fd7a5a67b2c9c34dd41b721eb8"></a>

## max_answers property — mx_pool / d35144f9505e / 4

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-d2faa2fca6a5f3cf44df98a3ac2fcc445f66ad96d1225c9afbc2d0d92a4fcccb): complete subsection reference.

<a id="canonical-dde05f7bc253ae1a66ad6b9a54ca3d27cc026416292907f68f7be70b972bbd0c"></a>

## Next pages — mx_pool / d35144f9505e / 5

- [mx_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-d2faa2fca6a5f3cf44df98a3ac2fcc445f66ad96d1225c9afbc2d0d92a4fcccb)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-d2faa2fca6a5f3cf44df98a3ac2fcc445f66ad96d1225c9afbc2d0d92a4fcccb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e9175b6fb714b79fc3d925fad6ea87cd122e02b27e743987d4334eee3f2609e"></a>

## mx_pool.members — mx_pool.members / a3a7bf2d75b7 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a638e71fb1dc5aac1b7456c9208f6cb86eaf14b84667e5425b68df2490d56db8)
- mx_pool.members

<a id="canonical-890aa25e3211a9a02c4c524a44d3625c3af13e07beadbb9acf9d9e6e79bc8f61"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-87d0ee9b8043ff44c03477331e908827ac46547d80c4b5f986844d86eb990c03"></a>

## Direct properties — mx_pool.members / a3a7bf2d75b7 / 3

<a id="canonical-82b03dbe3972d7ede1b73b986ede2d5d27c0428d111949e02841d6e31be3d101"></a>

<a id="canonical-9e3da04b12badef71d2da1d7a4df5ed5f869cf3149a629df2dfe578571d71e6d"></a>

## domain property — mx_pool.members / a3a7bf2d75b7 / 4

Type: `"string"`. Optional.

Domain name for routing and identification.

Upstream description:

Domain name for routing and identification

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3498ff24504acfed8961ba103891a6b8a4b5b982524f355b510b55cfb097b65f"></a>

<a id="canonical-ded0fae3801d8d7d803d0cc9d86da76626c7ce5a355c19c22127b572abb0b6f2"></a>

## name property — mx_pool.members / a3a7bf2d75b7 / 5

Type: `"string"`. Optional.

Name. Pool member name.

Upstream description:

Pool member name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-8df2685c9942c914fb580e3390b95b1a47477426d5c34e52f7196bc52943b682"></a>

<a id="canonical-66091418a8389bb117b9cd9d0ee4f287592063f605911e883e7246b28f53a766"></a>

## priority property — mx_pool.members / a3a7bf2d75b7 / 6

Type: `"number"`. Optional.

MX Record Priority. MX Record priority.

Upstream description:

MX Record priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-636213f4df31a1023b2377c558713727b178e04862e319b82bc96629b4c36fa0"></a>

<a id="canonical-2df66da97082eaff7c746bfffff573b9e4fd5a3243f6da249ca973cd8972428f"></a>

## ratio property — mx_pool.members / a3a7bf2d75b7 / 7

Type: `"number"`. Optional.

Load Balancing Ratio. Load Balancing Ratio.

Upstream description:

Load Balancing Ratio.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-b299f16c53d548fe41c91898efa91820f20cbb30bc9690cde74648f33255c830"></a>

## Next pages — mx_pool.members / a3a7bf2d75b7 / 8

- [mx_pool](resources--dns_lb_pool--reference--group-001.md#canonical-a638e71fb1dc5aac1b7456c9208f6cb86eaf14b84667e5425b68df2490d56db8)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-c0d7a3808b846937daabc8d2ff91a5bb4f57beb6dab6e0faf0549f3ca8f74d7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18ea71554fd27c332e7af45cb77cdf0afb9bad610e74e3700bd9ada7c7c0b431"></a>

## srv_pool — srv_pool / 340f0b4d2a4a / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- srv_pool

<a id="canonical-1e55f944d877db53b756428b7ba788cf987c4175c70219a7332f28cdb478628b"></a>

Type: `"object"`. single nested block, Optional.

Pool for SRV Record.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("max_answers",
    "members")}
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
srv_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-aefbaaed5aa722136e075b908bff496e4fe9c20e8ea70f41004fbd74b9714205"></a>

## Direct properties — srv_pool / 340f0b4d2a4a / 3

<a id="canonical-89235113e728a4742f443f9e30a19d06be0d24955f2b744ff498d42f36a27834"></a>

<a id="canonical-8853f94acaa8527c18711f5b7ee849ca7dea91945d86c521757f7ff9142815e3"></a>

## max_answers property — srv_pool / 340f0b4d2a4a / 4

Type: `"number"`. Optional.

Limit on number of Resource Records to be included in the response to query.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 32),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](resources--dns_lb_pool--reference--group-001.md#canonical-aa405c4b87f51cdd102624afae3840ed55b8bbb1c72eebd794c430ef7fa24eae): complete subsection reference.

<a id="canonical-3f0b9807ba10addc528df14d397a09a95747ce476bfaa28b6a868e1b96f0ea46"></a>

## Next pages — srv_pool / 340f0b4d2a4a / 5

- [srv_pool.members](resources--dns_lb_pool--reference--group-001.md#canonical-aa405c4b87f51cdd102624afae3840ed55b8bbb1c72eebd794c430ef7fa24eae)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-aa405c4b87f51cdd102624afae3840ed55b8bbb1c72eebd794c430ef7fa24eae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48efcb6698e4dd6f60ba83f80db72c2c7c691027e9f903b2fdcdb0101615ddea"></a>

## srv_pool.members — srv_pool.members / 655a2e3fc7ec / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-c0d7a3808b846937daabc8d2ff91a5bb4f57beb6dab6e0faf0549f3ca8f74d7e)
- srv_pool.members

<a id="canonical-2bf08d77a1d606979e158dfbf07abc18ca17f8c4e0f12a3018f6809e1ca987d0"></a>

Type: `"object"`. list nested block, Optional.

Pool Members. Configuration parameter for members

Upstream description:

Configuration parameter for members

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("port",
    "priority",
    "target",
    "weight")}
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
members {
  # Configure direct properties listed below.
}
```

<a id="canonical-d70d48df7e7de0f16103dbd2e9e27b9787a70015fdf84207b0513b6efb068828"></a>

## Direct properties — srv_pool.members / 655a2e3fc7ec / 3

<a id="canonical-22824b918f5a918ad3c7034dc666108f2bb5153e87c22be243e1c2c2dea0427b"></a>

<a id="canonical-6b6ee12283133f5d12929d9ca4459de70ff24fa9922fe7e43721294ed88f1492"></a>

## final_translation property — srv_pool.members / 655a2e3fc7ec / 4

Type: `"bool"`. Optional.

If this flag is true, the SRV record will not be translated further.

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

<a id="canonical-0652f1137a1ef1b111ef98ae49eff5a513975c5fee34921368eaa81ed233c248"></a>

<a id="canonical-c5915e37f3c4c2f8ef70b2a596b4c25cdff2cee73363e32ac1125f3705f83ead"></a>

## name property — srv_pool.members / 655a2e3fc7ec / 5

Type: `"string"`. Optional.

Name. Pool member name.

Upstream description:

Pool member name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-07f998d779c59a10be8ec65ac1b4867e3dda4df191fe248d8cb32160fe73052b"></a>

<a id="canonical-55da3e378fe0364225f12ba4bcbc81cbf15becc16e85408cf98c5183653656be"></a>

## port property — srv_pool.members / 655a2e3fc7ec / 6

Type: `"number"`. Optional.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2743ac9e875ec036e195268bed319b777ea0a9e6bb09d2b0e9d0f32550bbbf8f"></a>

<a id="canonical-bde462fda9d90130f53d6811d7103d55a2272f425e16ec760c2848106c12882e"></a>

## priority property — srv_pool.members / 655a2e3fc7ec / 7

Type: `"number"`. Optional.

Priority of the target. A lower number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-12befb6bd7292ae6db3b84fb56b7954ca37d85d2d8d796c11d6541915fd47b3e"></a>

<a id="canonical-03826a0ea3671dd003b38231d819e5ef85f642f4f9817baa213459942a65b916"></a>

## ratio property — srv_pool.members / 655a2e3fc7ec / 8

Type: `"number"`. Optional.

Load Balancing Ratio. Configuration parameter for ratio

Upstream description:

Configuration parameter for ratio

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-09f053765b89b4c795dbde5961488485fa76944fb13aba42e272eef5ff5a1c1c"></a>

<a id="canonical-7a51c62683eb0d9606b0a6f27b84345dbbdb094219756b87208613343c269a90"></a>

## target property — srv_pool.members / 655a2e3fc7ec / 9

Type: `"string"`. Optional.

Domain name of the machine providing the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-76e42f27e971d346a6d552f0b4c9ea4025323d05765467b186dce8ebfbae7579"></a>

<a id="canonical-2c941dd244d1f2b965e228d65cce9f6c5f7f167dd7a321e1baf823f0d941514b"></a>

## weight property — srv_pool.members / 655a2e3fc7ec / 10

Type: `"number"`. Optional.

Weight of the target. A higher number indicates a higher preference.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-dad398cd56b5b391024684a504007c43ecdc341e45f1a35cc2e785a071a3fd26"></a>

## Next pages — srv_pool.members / 655a2e3fc7ec / 11

- [srv_pool](resources--dns_lb_pool--reference--group-001.md#canonical-c0d7a3808b846937daabc8d2ff91a5bb4f57beb6dab6e0faf0549f3ca8f74d7e)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-31e8545911e0acc8af7f1a9adbfbe27fc9e8d7767f88361f35f13f9f6af750c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9bf0c43f8777d85d97758d2869556b491ceadcd9891ca785a76f9b8170cfdf7"></a>

## timeouts — timeouts / bef3a8858bcc / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- timeouts

<a id="canonical-6f4246cb70c5ccb0295ed7b4836143651c5e1f2d116e65e3fcf74ff0d2067f10"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3bb4c10d48f0d90d0cffa8fcc3acdfd73015fa690f22864ffba1a81bb7aa3acb"></a>

## Direct properties — timeouts / bef3a8858bcc / 3

<a id="canonical-016b101a9b29b309b5700703390daeee9daf36b57d746e4a1962ce663d0d8cb9"></a>

<a id="canonical-5c87f0580aa1cac1760a5416488dcce593cadd3d37e09503acaeafde53cfb810"></a>

## create property — timeouts / bef3a8858bcc / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-65d2328072579843187b02c888de65f8d186453b6ba9f0ff9e7a476593f2aebf"></a>

<a id="canonical-016a80aed8fb01fe8aa1de4fb57983443691b874e3dd6ba3da14759e11a7013e"></a>

## delete property — timeouts / bef3a8858bcc / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-e267be24310cb4518951651293ecf4a1ff1106a7e9e40db1803b4b80136213b2"></a>

<a id="canonical-d2d987153a9b2dfa25fcfafdf9a66afa1334e7b4b21d0764423bc9a0f2ee6c26"></a>

## read property — timeouts / bef3a8858bcc / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-9239dbc7517403801cf792acaa1447ef6960f5dae4cb5951c6add569dae36539"></a>

<a id="canonical-446ec5f0ba269f0cd7f9457d21757002857c3cda9e0d12d75daa107c8ffc2841"></a>

## update property — timeouts / bef3a8858bcc / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d0f7bfa324439e54704426ce082cc6dd64e2e993b54651056757c2a75105f434"></a>

## Next pages — timeouts / bef3a8858bcc / 8

- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)

<a id="canonical-61787dc50e3037e058b416ce7241c0eef43ae674606b147ed012be99217e997d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a55b8d11402521c8499a028b3284ed7d4c4966bc0edeb3020238d4196d911a5"></a>

## use_rrset_ttl — use_rrset_ttl / f820e80b5971 / 2

Breadcrumbs:

- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- use_rrset_ttl

<a id="canonical-54fe72d5b9eabe62b5778a239546a5f426562fbd32efbabd3d1c5f955f3d58d7"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use rrset ttl.

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
use_rrset_ttl = {}
```

<a id="canonical-48eff2d0143adcbd3ff8d9f543dbd85bf14ca219e622d0dbc64e6955a4a4150b"></a>

## Direct properties — use_rrset_ttl / f820e80b5971 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-424c337643a00e8873f772f5721dd7a5f5328243d6f12deec934da7e1c35c80f"></a>

## Next pages — use_rrset_ttl / f820e80b5971 / 4

- [Property reference](resources--dns_lb_pool--reference--group-001.md#canonical-e051d8e0ad04bf7bfa8ec8011b77167797a43be3ca460d8599a26eb6434b48cd)
- [xcsh_dns_lb_pool](../resources/dns_lb_pool.md#canonical-cbd8be9979f5969a414be1fc8eba160acc05faa80445e78a41e1ae2fed10b6be)
