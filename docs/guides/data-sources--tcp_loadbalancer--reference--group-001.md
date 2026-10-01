---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66ddca3b82a26320901ec4d2af19395cc5fdd85cdf6d556d13be7e140d4b0d12"></a>

## Property reference — Property reference / 7eba654d3f60 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- Property reference

<a id="canonical-4150a973cc02bade569d36f07313b92a9b2f144d65c4c9ba4965132fb689cf71"></a>

## Direct properties — Property reference / 7eba654d3f60 / 3

- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-45284443551b024b171d2f15c5b2a1a11c45d73a87eb54ea5ae3ecce816903d0): complete subsection reference.

- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800): complete subsection reference.

- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5684e25de901e83fdd06b403aedef7b7952f266b34a2576f9c1eabbbabb5b10d): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-93c42f400565a50ba7899bf25a14840f61f30439c967fbdb821dfad3d26713f7): complete subsection reference.

<a id="canonical-0f9f0aa3d47c170f19422b30acca2e6822c2a98ed19944894f257270a3f11cbb"></a>

<a id="canonical-1659ef00dab2a7af6a9b8f28747fe1be0f4f17bb76b8e4a128e40d5ce49c3fb2"></a>

## annotations property — Property reference / 7eba654d3f60 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-58d165c905a3aa921c6ac38d12138ce201d65967f9ffc5dc81b357df7ae71d92): complete subsection reference.

<a id="canonical-7cf89a45dabd1f83cded0cfe9103f0aa1923b4aa9a219510e5de2dcd397d01ae"></a>

<a id="canonical-0e3b959616e068ca4a1f744111c442f4afd9d7fa626e9ce708c10c3e0ab81f73"></a>

## description property — Property reference / 7eba654d3f60 / 5

Type: `"string"`. Computed.

Description of the TCPLoadBalancer.

Upstream description:

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

<a id="canonical-c6bc0a5ff846120ece02cb467b92287af0b46cd8e27487f0e4057d4edc1bdee4"></a>

<a id="canonical-087f1473c567a1c3a434bc1b899e41c727535b4bff8ab0bda032afc4e565ebb6"></a>

## dns_volterra_managed property — Property reference / 7eba654d3f60 / 6

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. This requires the
domain to be delegated to F5XC using the Delegated Domain feature. Defaults to \`false\`. Server
applies default when omitted.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. This requires the
domain to be delegated to F5XC using the Delegated Domain feature.

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

- [do_not_advertise](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-b1801b83aca2b6ca4df5d63ecc6355d3044c0feac6da71d837ff4f327a6aaa0b): complete subsection reference.

- [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-92a4b8fc91e5719ee6dabdf2bf0c8c07e6c9afd663ed699d176b12d95be766bd): complete subsection reference.

<a id="canonical-80b7fd6a82350740ac0b75ea17db3c3379e4b0e8b2b26decb5861e4ded136971"></a>

<a id="canonical-e286d29d35e96265f126f15684c9234ab0d663539051927e0986691a3d8afaf1"></a>

## domains property — Property reference / 7eba654d3f60 / 7

Type: `["list", "string"]`. Computed.

List of Domains (host/authority header) that will be matched to this Load Balancer. Supported
Domains and search order: 1. Exact Domain names: www&#46;example.com. 2.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6fefce460aa1c56801d63f400e477b2c03db10c5fade95cd00405f57b6474f06): complete subsection reference.

- [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-7b4e92e9b7e4a2ce124345a9c73070c6c23728a208081113b06d1f9771e991f4): complete subsection reference.

- [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-81aa7887946841c82a19ae76a9603b0fd6014eb80c03c27bd3f913d0edec525a): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-919c33641590e0126296cedb9bf8574ce4d8993d7c5af43da49b3145d89baa95): complete subsection reference.

<a id="canonical-08fc10935ac9bffe8776c1460e9bfd5ae4abe7c4fa1b072f6e79eddc1b42d5ff"></a>

<a id="canonical-2081fb1a9444200c210d4158cf023f7ab394948ee0110e03bb09354c614ed9ff"></a>

## id property — Property reference / 7eba654d3f60 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-8f7af16b38376ce9f5615418e228b605c3014be736446c3f194cfa06777ec9b3"></a>

<a id="canonical-42f39509bcde421b1efeade2498e8db65f7aceebe9af3cbc18a640f4fe7ee035"></a>

## idle_timeout property — Property reference / 7eba654d3f60 / 9

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
Server applies default when omitted.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.

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
    "ves.io.schema.rules.uint32.lte": "4147200000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4147200000"
  }
}
```

<a id="canonical-48e7ec93a1680605bb6b39a1c4295f5b987c24531d79c9e80a735475890ae070"></a>

<a id="canonical-31af1ef2385f461e2315763ad5f7b9815c1b13ac256a33d5c2a52f072df0c7cb"></a>

## labels property — Property reference / 7eba654d3f60 / 10

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-aa8c5f43346e51549215da9b839e6ba5825b9781c24284e92cd7b859376e4c9a"></a>

<a id="canonical-55a4c1d8439be4a69c86fdbb27349db80b14fc4f7c0c9b1f4daa050a1679e1c7"></a>

## listen_port property — Property reference / 7eba654d3f60 / 11

Type: `"number"`. Computed.

\[OneOf: listen\_port, port\_ranges\] Exclusive with \[port\_ranges\] Listen Port for this load
balancer.

Upstream description:

Exclusive with \[port\_ranges\] Listen Port for this load balancer.

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

- [listen_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa8c5f43346e51549215da9b839e6ba5825b9781c24284e92cd7b859376e4c9a)
- [port_ranges](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c4fa74814bb91b338c7e9385da77e309e04d7a9c91764f89fa5c85c2ef375a42)

Select alternatives according to the provider validators above.

<a id="canonical-b1d4a3f4fd970a7e704a89a7902a625eab915353775f6538da8646c6ca4cc50d"></a>

<a id="canonical-76b1b2f04b10d4ca5ad6abd766332529a5d675e3d130b33f9a519f4b1a16dc08"></a>

## name property — Property reference / 7eba654d3f60 / 12

Type: `"string"`. Required.

Name of the TCPLoadBalancer.

Upstream description:

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

<a id="canonical-60c5885abe7449be1ee0e7a64626657531a288d2be37523c9b18310aacb4697b"></a>

<a id="canonical-59376c2cb95dfac87f482f6a49ee5d97484dac201a5a6ce107beac2b10fab60f"></a>

## namespace property — Property reference / 7eba654d3f60 / 13

Type: `"string"`. Required.

Namespace where the TCPLoadBalancer exists.

Upstream description:

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

- [no_service_policies](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-eecd9cf3f444d97ef8d07cd7dc143b76295c7f091307b4c0ec0be1395b0afd03): complete subsection reference.

- [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-fbd7a2a7cc37149eea9e418797c4c0076cdcf0920443426b82cf1264fe067cbc): complete subsection reference.

- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91): complete subsection reference.

<a id="canonical-c4fa74814bb91b338c7e9385da77e309e04d7a9c91764f89fa5c85c2ef375a42"></a>

<a id="canonical-d27b0ff5b3d7fc63fce51db42a4dfc41dee6d260429b6fbb3b9ee86cd49aeabd"></a>

## port_ranges property — Property reference / 7eba654d3f60 / 14

Type: `"string"`. Computed.

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[listen\_port\] A string containing a comma separated list of port ranges. Each port
range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-cc1c3b2f483ba6234b58359b9208b3f5b95fbdc41442e47e1b1563ca0af0f2e5): complete subsection reference.

- [service_policies_from_namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1b22f5ba2b642ff390ebaac0f099fe50f84b8b44393fb2f0efbdb4c499f4b7d5): complete subsection reference.

- [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-9fc01974e3956c45c697aa96fa01d9561f837ca5f51871023cd162da605ea22b): complete subsection reference.

- [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8e30515f89b566c0046c0b656f5f1cae28136176409095156357cbe5f3cab9fc): complete subsection reference.

- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031): complete subsection reference.

- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929): complete subsection reference.

<a id="canonical-ec4008ce2e285d3bc7b513fd03b1105f9f16e839ef8fb11697affb6f59240e3c"></a>

## All schema paths — Property reference / 7eba654d3f60 / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d82c009ba1a1aaa6002c1ec20e4a1e5ae1e81a34c6a82e063d94ade1b66c6175) |
| `active_service_policies.policies` | [active_service_policies.policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-ca4e72eb8d8f385236af57083d019ad38aea47a0bee96ad836f9f9fff4364e95) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-ecb505853434fb11ad44de7fe8097acaa16f3d85182a3302a3d79e659462bffa) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-f7826859f1c43c12396594c22a0dbbe1c9f709369714cedc1e374b77550b3a41) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6327a96cf9f3295181e2d5ec447431d5ce419f7629e6ea217d051621c1d2c2a4) |
| `advertise_custom` | [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-471af99fcb75d952a306832f9d95c86b7be81d5a2b4f27938e79aedd8c85da87) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-f4b5f5d26e5051fa35327f18d355aa9186b5924a0eb7874e363f1452f5fe42d9) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-adb1b2ff73be6620249351b0f71f7786c386718634701e3f44d1a6152c04407f) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-29317668529c303b89e96e6a80ecee79c6be37432fd225877b23daefa6b0ca49) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e0d27db3b216c5705f114150fd28088467112e7a282583387a9c08d7736e3d08) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-8840b4785b565a7e4fdce72b2615772f33bb4ed0e9da2c949018c5b4388ff70c) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d6f0a79800094cb5b2d575932283d5924b3140990b2e8f448d994562460aa9d7) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c2dfc0e73fdc005042660b79524a631776e8a138724c78a38dda04210dd3ebb0) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-f3f2ccd5e38421216ba5e4b19d80991c1e8ea4852c5e18b5cf2caa7576ebb63d) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1b5f96665f81c3055ee9af89e5db5ffd25a383994ac72a751780fe668910c155) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-748021f63e6f80fbf698b410067c4ab5422081266395f717aa04ac79706d29c5) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-102c8fbed596557e281a980be9fb6af3bdfa6b7a40b16a19216d8fb49402eff8) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-bd1d695be46aa5531497323c6ca98a36111e6c78bf05b9a03a5bbb5b0315d425) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d3a4f017cd0b77e7bc0ed27c928d6d691cac755dff6bf9fc00b74b19cb3ae836) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0996c3312330c90c85c20cec94319daf9766175e16aba23f6fc71dc003129851) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-482f79575613fdfe556a1d95ed77134e4701e649076f2ed85b2bed590d1aa3dd) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-ef507ffc44c8bfaca8a861e4d31cc1a83a61ec6f73b5c2330d2b8595606415f7) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-589be668802ce6f2bbe966a5dcf9a65cf27b688bb7ea7c6142a5bb35d5749c7c) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1471343496d3b52c8f68f76189e478a8797a7072266c437938d069e651fab01e) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d809f6e94b6a258d25e40ce871e5f83a103127e06084bbf2b55deb35cc9556dc) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-263e5da0c66fd7fd78caee8b160a2ffef11d8fdd93e1b84b1a4536c1aaaf9f3b) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5ed1e6f9be8027167c2c534ae98f40f512e4ad0e6f30c03f56d20fa75cf334ea) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5b1ecf73154021e3b508b541652dd8f0dcb5ac7c111bd831ff6d0d4da650a983) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6865bcd2fbb369c94c602916bdf554ac44dbb4647024bdad868952060a46faa9) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-7bc1b981b7da82bd46b321158ad1c7553f9f610daefdb432e3b677d47f84187a) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5a5e8c49b6558287791b83b29c5149652fdb4b02693994b5f4910c0dda98d126) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c0c5285ca8e66556d0d9d118b406132edaffec1d66415656d9c3a9a8d9d9c762) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-292d1afe839a4f286c37c6d7e0a7cadc23d414c47954572a260c495c80a19226) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a5455aea40c21e399b666fed309d9c265c75416d59c223c249f93496a86bbc7e) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-061b46b0d7e20bee7c4aa56c55613fefdc55c7fe1360185fd63ce425f353c372) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-df14e0e037dec276ae2430dd4970cfd6df6f3ae68b477116c38ac6a40694af6c) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-ea70527ac5376409eebe7a773359dad049b3ee7f2a6cf49ca0a17ca852dff5e7) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-050cd0774eee544bd040435ec613eb2d98736708a32cdd8b56e2f9beda4cfac6) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-8fecf36e3b4ef911f3012b36f4567f3c8c041a601000f6affcdf93119ef2a331) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-30be27eb00c2487a5a78733c0e43c1c89de14fdb8e515d23615871f2721a3ad4) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-90e43382e0da0ba254501813be7f7afd3d79ab19700841e1c6e10bc22e93d4ad) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-34ebf29b4a091424a7c8534e643a08d001a7f32891ff7092133a27b54a70a49b) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-9615dd6dcc86060a65c25b26db73bc538be804decee62fcc0c3cbb5f82558e64) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-005d3cc6febfb0dc643987bffcc45650ba3357a5f80868678ab652d886703424) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-925052ca6d49d6d1c52ec576e82b4d1087508b4a20f3c5bc61fc5669a3c441ba) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-29641efe6bdf167004cc1a7e3886b753b921cb1b15e08385f9f82a510ab3b46c) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d1bbcdba119d2fce89d77672fdd8c0d9b3b94a0fa03d6d9853e4907475cefb8e) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e33dddab6ab70b9eee12a677281f2693498435d927ff0d155910bb9d205ec39d) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3a29474a7e433e733da617c4782e5decee850bc2de7cb97469fda5d97113ef37) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e7c883e3a7d19c3a74542f16ab4ca61a8ed23d9f41ec2da5a50f1c18753c2b68) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-37556c8801ff63728726da7ee4488bce6f90177c947f2bfa9ff4026dec361fca) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-53e8fb33c82ebd73937213d4fabae679789b21e17b713fcfcbd16dc7dbd3c345) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c00fe8e921bce7863240cf95f20dab00b56b179f2093dbd9c6cf3a8b45815901) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0b496fa15b2fb8707a8e6abe810e831a48f52476a1f8500e7c16f3815da38648) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-b5fab35e1014938779a23334b010a9851cc21a197be6f85199f06c6d34a598b0) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-bcb2af668ba2210029eaa253ddff86c56c48acd22a5e411a20ca24e548a75d50) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-508462d0461401372bf6525581b8da21ef2fff10fe5074455a941e72053e9002) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-df05d1b860cdd37b1dd088b005c2fe7f3e79d3add30e0f9fba226e83d15030b1) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-18b08ec442a92b323130d8cf2a65b2c5abccdceb509944c3d48a66668f8cc850) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d5757c3b5b3740647933b9646f59ad11fd82972986d34117a6d463d93d2e1b99) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e168397ff3465af233242a93540c6498985ca2d0c643d5ea966fc86d3703d8ed) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-06d989a40bf48e74992a4af3b0fabe4a27f5d7969c6bccaf64dc8d3d17c512c8) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-62276a01ee1d3e94d059313b59f296b7fc18794461374a8e578733f9dc649d4a) |
| `advertise_on_public` | [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a036294937822daf472e5e0cdb0113c257532b5591e640bafd0f49e736f09d5a) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-b2f0a5d81ea2b5cc630b6fa8c57aa538104f8f9eed9d411a29c43dd8f6645813) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6a30a73c3d1f6d27280d801f9aa9b6cc6bc5035e137653f2d73492d9ca6a8a07) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-2e50ccc0fbbf69023d0a85ed6a39df94387f175d2ed145ff7df167bcb35a8a86) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-bb99f1a3addfecee27de11f7a0135ea56995641c16a775d427eacec7f490bf43) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-5bf02f2574ca18e1572f4a5e318efeb20e4f23b884c4cdd1a7053091eea130a7) |
| `annotations` | [annotations](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-0f9f0aa3d47c170f19422b30acca2e6822c2a98ed19944894f257270a3f11cbb) |
| `default_lb_with_sni` | [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-fb0684273e818fd43aab97fe6e37b100e665012e3122609c33a936ce78119679) |
| `description` | [description](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-7cf89a45dabd1f83cded0cfe9103f0aa1923b4aa9a219510e5de2dcd397d01ae) |
| `dns_volterra_managed` | [dns_volterra_managed](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c6bc0a5ff846120ece02cb467b92287af0b46cd8e27487f0e4057d4edc1bdee4) |
| `do_not_advertise` | [do_not_advertise](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-01cc0932a46a5e320283e534345f8c825a17d59bebd68c99ece71365db19d7c0) |
| `do_not_retract_cluster` | [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-091de0d4244a87b5f28a89d1b769fbba251f63159ff4955552ce45874ce8c26e) |
| `domains` | [domains](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-80b7fd6a82350740ac0b75ea17db3c3379e4b0e8b2b26decb5861e4ded136971) |
| `hash_policy_choice_least_active` | [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-276f5b12ebcd77b22ad7ddc5af64d67552eec4766bc436625a8e827c56e62a70) |
| `hash_policy_choice_random` | [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e63ed6ce45d8bec7456f0a0b6d3e37296135057044b29fe677c537905a2d1ba6) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-070f057da974c9821fd7516399b7360b78161a5faec435deb93fbc65fd8b23eb) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-01dd2b7992898b9df7a374e9b65ecd82f63a3c952ffb253f77930b95754a7daa) |
| `id` | [id](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-08fc10935ac9bffe8776c1460e9bfd5ae4abe7c4fa1b072f6e79eddc1b42d5ff) |
| `idle_timeout` | [idle_timeout](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-8f7af16b38376ce9f5615418e228b605c3014be736446c3f194cfa06777ec9b3) |
| `labels` | [labels](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-48e7ec93a1680605bb6b39a1c4295f5b987c24531d79c9e80a735475890ae070) |
| `listen_port` | [listen_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa8c5f43346e51549215da9b839e6ba5825b9781c24284e92cd7b859376e4c9a) |
| `name` | [name](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-b1d4a3f4fd970a7e704a89a7902a625eab915353775f6538da8646c6ca4cc50d) |
| `namespace` | [namespace](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-60c5885abe7449be1ee0e7a64626657531a288d2be37523c9b18310aacb4697b) |
| `no_service_policies` | [no_service_policies](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1959f2c19f4e27dd4cfd683029d504e1fca822c6390b85c034c7f8cfa01575aa) |
| `no_sni` | [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-965fdafaedb81e2f72308d6ed003e2ac4e28a96e84f352bc1b99b736d8b2d726) |
| `origin_pools_weights` | [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e08be0f4252fd01319d67849816b91515c2ebc2959d64debca59b0f010b1fdc0) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-5719014c81e92c4c5d1053cc770c6251141633150eb20035c3e3cca76047d415) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a61fd4e9cdf8fe0d646bb6fd4ddd7d34d0af5d79d85f5776d6e9e086e8f5c3b3) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e0fbc76dd38f030ccdef27a4dea599786c318626032e71a46907d14b603be28f) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-30f6dd08a9107495ca41e78af63ade06913f8a5e5bd21d936980f11aaad3ca3b) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e86dee8129b14a97b1077a0d835455aa240f45752d6756efd48feb75507d2fd9) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-5e9df8be7a18347bfc2c08053dbdf17eaac2b6a03d73aa9badfdc174c6d7aeaf) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-62b1cc7112b4aa495eda22fd3b3fd451a022be5b9f9fca8fd705f3cd5ceb2b18) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-ef404c1aeca4b6205718974e08891675d0a3c700bf38e68e8fb38e6e28bd06b4) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f98bdacf88c4d534a6114ebda8c8b0e1c6542c54151eda4e75d3fa145bb3df92) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-13c562149f911e9cd8e5f5dc67346c376f3513b08b0596f30c4ec1f56b958e76) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6d7839ecb65bb4e77339d4f5482ac2fe81046117e1303de79121b6281bd83476) |
| `port_ranges` | [port_ranges](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c4fa74814bb91b338c7e9385da77e309e04d7a9c91764f89fa5c85c2ef375a42) |
| `retract_cluster` | [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-5d2efbc23bdf348e8a597d843d71f01b03efd513109e51807d41264372f8ab6f) |
| `service_policies_from_namespace` | [service_policies_from_namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f13cd904bb7cda8dcac837323ed3c4fa516dcaf5b3266754719068b762ee4480) |
| `sni` | [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-b6a2633f2cb6856d063cae523509c0ab9608250c264c8c8489e72187a9450702) |
| `tcp` | [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6fd58b64bda394756870c8b71293a539b1718404753bd67efcf8ce90155f8d87) |
| `tls_tcp` | [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6649da9a04b26b4b8eb2d282f1e67f5245a513f9bd8538217e65210664609bae) |
| `tls_tcp.tls_cert_params` | [tls_tcp.tls_cert_params](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-dead4bd0781a288f5e40cb6f0f4b9060ae8d3c234f955ff93d80609436ccff58) |
| `tls_tcp.tls_cert_params.certificates` | [tls_tcp.tls_cert_params.certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-9d3635db982ed7e48de7a8fac8a1cb51eeec01c79c3293e7bfca77bb256cb36a) |
| `tls_tcp.tls_cert_params.certificates.name` | [tls_tcp.tls_cert_params.certificates.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3c9dc3f2b57e60ee6d162a3d21d06d4272ece514854d30d152420242bdf40fb5) |
| `tls_tcp.tls_cert_params.certificates.namespace` | [tls_tcp.tls_cert_params.certificates.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-86a77c75d9131a8473c971dc3968ff11fbb8fd20e242c8656ca9d2fef3a91658) |
| `tls_tcp.tls_cert_params.certificates.tenant` | [tls_tcp.tls_cert_params.certificates.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8efce4f1748e1934a7a7034b7f5a68b5056b028e5707373e2ef0f6f5db4123f8) |
| `tls_tcp.tls_cert_params.no_mtls` | [tls_tcp.tls_cert_params.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-5dfbeaaaa53dacbf740dfa2357f228e5a254cffffd97a7c1dc398629eb5234f2) |
| `tls_tcp.tls_cert_params.tls_config` | [tls_tcp.tls_cert_params.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-59047fc953cd1e5a6d0fd9178d12912477a531eb55df1b222e35e95cf5c2bd0d) |
| `tls_tcp.tls_cert_params.tls_config.custom_security` | [tls_tcp.tls_cert_params.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-602b554d6947311fe0d865fb62d38cdca59b89fa749d9bf89c0936b8931e91f3) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-b61790c898143a9dac35a2a54a20a3f32a7713be7693deebc04666b0231f6a08) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.max_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.max_version](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8d647bb4a6145d29485b4d58f0ff5269bf63ed28d56841cfb95f0d4b3bbcd4a4) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.min_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.min_version](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4733de4998e56446ca118a72157a3eb6b578d056f75716098fbcf444b02b1ee5) |
| `tls_tcp.tls_cert_params.tls_config.default_security` | [tls_tcp.tls_cert_params.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c0923fe93e888dc23a70b5729cea9b54ad69a21deeae01695cf0cae3a7ffe059) |
| `tls_tcp.tls_cert_params.tls_config.low_security` | [tls_tcp.tls_cert_params.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0c1df4b2d5f504ec2675bac19cb8431878255542c81874eb37e3466e858bca8b) |
| `tls_tcp.tls_cert_params.tls_config.medium_security` | [tls_tcp.tls_cert_params.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-46e00117bf332968b77ae72b9c1687ab8d5137d3a13b70065a853d7de99c6b9e) |
| `tls_tcp.tls_cert_params.use_mtls` | [tls_tcp.tls_cert_params.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c6e71f375c0198a4de0a767c9151f012a951ea1c1ca26b899f77477c69c055fe) |
| `tls_tcp.tls_cert_params.use_mtls.client_certificate_optional` | [tls_tcp.tls_cert_params.use_mtls.client_certificate_optional](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e3003c52d4e81cc51ea49dbdc7a3d96762a38a8c4454c26c54b62b6210c6406f) |
| `tls_tcp.tls_cert_params.use_mtls.crl` | [tls_tcp.tls_cert_params.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-49274aa568c97a65279caf8c884c69a21ed44162c32c2669ce110a51ec7e3d0f) |
| `tls_tcp.tls_cert_params.use_mtls.crl.name` | [tls_tcp.tls_cert_params.use_mtls.crl.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-07a40c09745ddac316ec4411a0d9ee838ff59459281dd08c51d1c7d114156e06) |
| `tls_tcp.tls_cert_params.use_mtls.crl.namespace` | [tls_tcp.tls_cert_params.use_mtls.crl.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1726181a9f6166d895ace7f1ed3a937f527b71d847c12c4c6f1fd03a8d64257c) |
| `tls_tcp.tls_cert_params.use_mtls.crl.tenant` | [tls_tcp.tls_cert_params.use_mtls.crl.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f10da52298c49a72ae51ddf0abe74b04d7f2f2acc99177a49b03fb86ca352b17) |
| `tls_tcp.tls_cert_params.use_mtls.no_crl` | [tls_tcp.tls_cert_params.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-77a7ee254580a4bc3a9f386a0bf1e2f226c6acd6d19d0f8341aa1dd7b622e0dc) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d0da501f02421f8c36e068b0784e568bdaf840ddbee4c06a5a7d8f1fb714d62b) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.name` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-cd32564b3cc8c04adfaf8bf51a237c3a8d92cef7a55df49422acb7d879f1f193) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a6653bece57f378cf09c4e4b89958623956280598ce23e0479f02bc44b17a61c) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a46a5ef6c621ce472f7128e48c1df8b5b9b29674791837a6a47827fc7067b7b2) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca_url` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca_url](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6e2c33b765640ffd512f4d98aa597bde4c01a6a5028264cfd306623a3cb682bd) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_disabled` | [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-86f263245c948c774c707fd9030c69090bdce525d67c096e4c4cf6908d65b0ff) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d7fddc6f465cf8b460cf969c353ebe1f307544a31fbc01cb1638e289de2e2895) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-2badc0e2f8749532bba00a1a475a90d550ad6e0a7abf03715e7e9c95c239ceae) |
| `tls_tcp.tls_parameters` | [tls_tcp.tls_parameters](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-99275446ac13605bb6647d69fae1798a4be9eaef2fe16d3780be66e1d1eef301) |
| `tls_tcp.tls_parameters.no_mtls` | [tls_tcp.tls_parameters.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c17dbd8240e29b262a640d3360bdaaeaacf8e7bff790136d7737a81d49554140) |
| `tls_tcp.tls_parameters.tls_certificates` | [tls_tcp.tls_parameters.tls_certificates](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-464ac9efe6e532f14d5a08ff037178c217be579c9044943db38f20349c142b6d) |
| `tls_tcp.tls_parameters.tls_certificates.certificate_url` | [tls_tcp.tls_parameters.tls_certificates.certificate_url](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-509c8676521f794a891dbbd179b46e483878be3a7397e32ebd8e4f5f3e8aeb99) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-ffdb2a2daa20d8fc63817649a193b321cfaef5e7b26f4764674f5468404af51a) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-21cf558f9b7105f8fd5270944c3e4a723da778c97aa87668f6a877af4a012f03) |
| `tls_tcp.tls_parameters.tls_certificates.description_spec` | [tls_tcp.tls_parameters.tls_certificates.description_spec](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-acfc1bd88952526c5efb84dc56532d01ba99664cf4e18d087d1ef7b3a4e36a8f) |
| `tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling` | [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-457fbb0f055047297b0d318613d60c4c5ca28a8f9356968c78073d048985e2ba) |
| `tls_tcp.tls_parameters.tls_certificates.private_key` | [tls_tcp.tls_parameters.tls_certificates.private_key](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-51f437676dd8e97aaf145b404295d2d44a2750fe2bc22f68d2371e53a82959de) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f8e2eebd011445308fbf055a070f4ce0bb91f822cc407ec209c22b2f5a102bd6) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-cf1362f8dee2049fb3b4f99406212853189f2a33a593e06271bb2a6ee9a7bc50) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1ff1c2aa09ca5739cbb2ea161402eebda1e7b187d34de94b49faca1c341fff73) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8bebe8b8a7c4cb94327d2f38eca856f718fe9987834ea8787874d8864decaa6d) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-7693ae84b5c8b06e0ebefc9059c1d3ab2edf30001eb7d80fb45498b7c5cea927) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c35806fa5d59a63757fae87a8e94f03ff761e10fdcbff183573dcf784586b0d1) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4b31f7fbb2800f07e858400969cea3c4946311281da299673745817903e929a5) |
| `tls_tcp.tls_parameters.tls_certificates.use_system_defaults` | [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1fb99423d6776abe649612e3273715ea5053346d42e4ba209ce0b8e0920e2dc1) |
| `tls_tcp.tls_parameters.tls_config` | [tls_tcp.tls_parameters.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a9cfd7e35c8fba86fc7b9381a9b01807e5e26df77c2b481d86a05d135bddf451) |
| `tls_tcp.tls_parameters.tls_config.custom_security` | [tls_tcp.tls_parameters.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c745ddd7211cbf6aeff46c0a71c7fbadf6ac941236741a5e936d0684127c6f5e) |
| `tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d05ae0eeb0a65e3589bb302230a501ee353272c987eb83d7f2e1869174a1fbf0) |
| `tls_tcp.tls_parameters.tls_config.custom_security.max_version` | [tls_tcp.tls_parameters.tls_config.custom_security.max_version](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-305c08eba83f59d8f4fd30c731393da305c731338195a4ed154c9544a42a0ebc) |
| `tls_tcp.tls_parameters.tls_config.custom_security.min_version` | [tls_tcp.tls_parameters.tls_config.custom_security.min_version](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-fd42d7518dba16c04e29d608bdf065efbd496fdda2ba517c62152ff7548ded22) |
| `tls_tcp.tls_parameters.tls_config.default_security` | [tls_tcp.tls_parameters.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-7e2aae6b94b17fc8496ee02beb2ed6b931148ed0ee34b6d6ab1be3ec868e1c81) |
| `tls_tcp.tls_parameters.tls_config.low_security` | [tls_tcp.tls_parameters.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6f23f3010bb4b542b89645de20e4c13c2ac5906e8d43d67110ce5e46f598665c) |
| `tls_tcp.tls_parameters.tls_config.medium_security` | [tls_tcp.tls_parameters.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c14ddd6bc4237362e00ce855025bbe93dbd2f380dd55b5ded25fcd00ed6a9d51) |
| `tls_tcp.tls_parameters.use_mtls` | [tls_tcp.tls_parameters.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-eba97b0a9b8cfc84e10ae26096ceaf8547e2cdca908f784b381cf3cba37d96db) |
| `tls_tcp.tls_parameters.use_mtls.client_certificate_optional` | [tls_tcp.tls_parameters.use_mtls.client_certificate_optional](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-343b964a5dc1ec5a11cae5e37c09f9474a105ce974083e7187f18bc5bdb3b5f4) |
| `tls_tcp.tls_parameters.use_mtls.crl` | [tls_tcp.tls_parameters.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-10815073cd332e2b247814082e3ce6a88bc9cfcffa9a8c79ac8562ba53a4c9b3) |
| `tls_tcp.tls_parameters.use_mtls.crl.name` | [tls_tcp.tls_parameters.use_mtls.crl.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-56cbacad3fb382ff77aa412ac8424e0e267d9bc7625501899907a2350f7640d2) |
| `tls_tcp.tls_parameters.use_mtls.crl.namespace` | [tls_tcp.tls_parameters.use_mtls.crl.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-114bbb3d4f5606e6377f5643f2523f0ab63083db87da7cca7ebbdb772dedcdb6) |
| `tls_tcp.tls_parameters.use_mtls.crl.tenant` | [tls_tcp.tls_parameters.use_mtls.crl.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d9fc93bd16a5d8c0ac4baf14668848d4f61ad9f78fea4ebdfbcbcd3a710e325e) |
| `tls_tcp.tls_parameters.use_mtls.no_crl` | [tls_tcp.tls_parameters.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-3e59bc68b4c19552a3446fa8b7355f5d01df108551f761379ca82d9b6f1b7fab) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca` | [tls_tcp.tls_parameters.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-74282a2cf195454318694c61a7c5412b6d5a3909f1c85becc18154a38bde0b4c) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.name` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.name](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c1c855cc2e4762e451fc8f4900d2af914c1c2e91ffc5ac67458e2af651a4b46f) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c6b87730655188844fe1ee0c6906073da63c393a3f74cf23c35fa0d17138a625) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d2676906ecac0b3c0281e668f3180822adca61a5e28d2cf2565bda47be41e8f7) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca_url` | [tls_tcp.tls_parameters.use_mtls.trusted_ca_url](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-c5c6b0b77279e2f73c608ea2ceca255604e151f3e0fbc931b25ada739661ce98) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_disabled` | [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d8c01288c4e74c836c766f5798bcf219e2e7e7ee1af5daa0f7c68cdf8ff2cc08) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options` | [tls_tcp.tls_parameters.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-48724beffdc2bc5314a8e0044c04e991c8a7eed16c84cf9256ceb9d19022916b) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a30ebade131adc2c935677442beb4843304b982a5a7a23242c0c55325654bb7c) |
| `tls_tcp_auto_cert` | [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-da1071ff0460dbe3ac648e09b3a5265ecc1c738a79ff6727775e870c422eff00) |
| `tls_tcp_auto_cert.no_mtls` | [tls_tcp_auto_cert.no_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-76efaeb96c3f588009f2c98199c0c17274166db9b92ab9fc8a93d47caef9bc6f) |
| `tls_tcp_auto_cert.tls_config` | [tls_tcp_auto_cert.tls_config](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-90affc6782fa8b0b557c840073a58818a22acce5528de37552c6d6ef248f0633) |
| `tls_tcp_auto_cert.tls_config.custom_security` | [tls_tcp_auto_cert.tls_config.custom_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-d79e8f06e44c46d4bacabbc15e99e74adda29fe557188467621979541939e38b) |
| `tls_tcp_auto_cert.tls_config.custom_security.cipher_suites` | [tls_tcp_auto_cert.tls_config.custom_security.cipher_suites](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-0bfb309cfe30b6a87df6088bb9b5f733611068e4fe6e58577f5a2d7a435fd897) |
| `tls_tcp_auto_cert.tls_config.custom_security.max_version` | [tls_tcp_auto_cert.tls_config.custom_security.max_version](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-74700b1a209eaeed652f1d214a7be88008599e7c4653b22f199a833a00d239af) |
| `tls_tcp_auto_cert.tls_config.custom_security.min_version` | [tls_tcp_auto_cert.tls_config.custom_security.min_version](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-9fa4453b48ca1ee1f47757042c07942e15440761e9778ce94600e5ba336d4889) |
| `tls_tcp_auto_cert.tls_config.default_security` | [tls_tcp_auto_cert.tls_config.default_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6d6e5669e9ff4ff4b6161648948574397005fc9a76bd960eb1f0c104c06d8ee5) |
| `tls_tcp_auto_cert.tls_config.low_security` | [tls_tcp_auto_cert.tls_config.low_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1532cebffbd122a3fbe1598cdea8c2aabf096826ee1a9e81578b5c0dbfb5d144) |
| `tls_tcp_auto_cert.tls_config.medium_security` | [tls_tcp_auto_cert.tls_config.medium_security](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-730e7eb521689469723008ab4707bbb4b0ac2e164cd7a73a38893ec01e238f9b) |
| `tls_tcp_auto_cert.use_mtls` | [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8032bb22a890292ac2a3bf76f3fa65288e9055e011dc1ddf2a381c0f06ae07e5) |
| `tls_tcp_auto_cert.use_mtls.client_certificate_optional` | [tls_tcp_auto_cert.use_mtls.client_certificate_optional](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-74a8307f2112ca1de79aaa672ec5dc9da70abdf8dc8001e72a6d116452116814) |
| `tls_tcp_auto_cert.use_mtls.crl` | [tls_tcp_auto_cert.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-8b65c7535cb28f815315fb93d2761bf2f62187a800e842ce99d6fe9dadbff094) |
| `tls_tcp_auto_cert.use_mtls.crl.name` | [tls_tcp_auto_cert.use_mtls.crl.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-59467e9c69d3c22791e75134e0ebe759f57a6acc0d692cfaaaf44f0d58fbc0e9) |
| `tls_tcp_auto_cert.use_mtls.crl.namespace` | [tls_tcp_auto_cert.use_mtls.crl.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-8f26a75a6ccec616e289f65abc49a2116db94ce40c034d7a514da505a1005e5a) |
| `tls_tcp_auto_cert.use_mtls.crl.tenant` | [tls_tcp_auto_cert.use_mtls.crl.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-48e93e69e8c570fe3f39ee42de591849f646d67a54086e7bcc51dee3ef77d39b) |
| `tls_tcp_auto_cert.use_mtls.no_crl` | [tls_tcp_auto_cert.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-b4d5e68e2f911d002aa39b00d77ff1a112e6a34c00408f9d3b75138dbd0d6dd3) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca` | [tls_tcp_auto_cert.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-26793d7dacab2547fdc4d25c2cf863b21c8f0acae8b3da17e5ed242282dbd0e5) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.name` | [tls_tcp_auto_cert.use_mtls.trusted_ca.name](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-b29ba026fb23d665ac23ae5d9f026be9df4457aad725b1c9be184c391064f7cf) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.namespace` | [tls_tcp_auto_cert.use_mtls.trusted_ca.namespace](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-8a0e431c21d596764cd23b4d7ed21969e7741d666d0a750d47f35a44133c4398) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.tenant` | [tls_tcp_auto_cert.use_mtls.trusted_ca.tenant](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-77f13ca43f59f9ca68a631a4f5d49f1cefc5129bdf9650554497bc35c14bb1db) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca_url` | [tls_tcp_auto_cert.use_mtls.trusted_ca_url](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-e77db2b4b01c549b5391ea17e667c4e0b926de99bb151ca62d4d29c64b9157fe) |
| `tls_tcp_auto_cert.use_mtls.xfcc_disabled` | [tls_tcp_auto_cert.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-eeb8c4f8dc665473647ade8ac7e5cb09303898d57e1af1098f6ad24661d93ab1) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options` | [tls_tcp_auto_cert.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-a7d148f76edcf35d7e6673cd275a6125d11953863696c8b0a3be6802e6b109e8) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-4fbd7a9ebbb27e4690d7e94033bb1aa0b3d65a8b59feba065ad4be8e6eea48e1) |

<a id="canonical-04407d75f3e001686a13bd08b888f72cf31866b903f6a0db96389cfe0efed8df"></a>

## Next pages — Property reference / 7eba654d3f60 / 16

- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-45284443551b024b171d2f15c5b2a1a11c45d73a87eb54ea5ae3ecce816903d0)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5684e25de901e83fdd06b403aedef7b7952f266b34a2576f9c1eabbbabb5b10d)
- [advertise_on_public_default_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-93c42f400565a50ba7899bf25a14840f61f30439c967fbdb821dfad3d26713f7)
- [default_lb_with_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-58d165c905a3aa921c6ac38d12138ce201d65967f9ffc5dc81b357df7ae71d92)
- [do_not_advertise](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-b1801b83aca2b6ca4df5d63ecc6355d3044c0feac6da71d837ff4f327a6aaa0b)
- [do_not_retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-92a4b8fc91e5719ee6dabdf2bf0c8c07e6c9afd663ed699d176b12d95be766bd)
- [hash_policy_choice_least_active](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-6fefce460aa1c56801d63f400e477b2c03db10c5fade95cd00405f57b6474f06)
- [hash_policy_choice_random](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-7b4e92e9b7e4a2ce124345a9c73070c6c23728a208081113b06d1f9771e991f4)
- [hash_policy_choice_round_robin](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-81aa7887946841c82a19ae76a9603b0fd6014eb80c03c27bd3f913d0edec525a)
- [hash_policy_choice_source_ip_stickiness](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-919c33641590e0126296cedb9bf8574ce4d8993d7c5af43da49b3145d89baa95)
- [no_service_policies](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-eecd9cf3f444d97ef8d07cd7dc143b76295c7f091307b4c0ec0be1395b0afd03)
- [no_sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-fbd7a2a7cc37149eea9e418797c4c0076cdcf0920443426b82cf1264fe067cbc)
- [origin_pools_weights](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-894f595bbc8978227e5e9971546495ff0b655dcd9725c3b9e929610927854c91)
- [retract_cluster](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-cc1c3b2f483ba6234b58359b9208b3f5b95fbdc41442e47e1b1563ca0af0f2e5)
- [service_policies_from_namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1b22f5ba2b642ff390ebaac0f099fe50f84b8b44393fb2f0efbdb4c499f4b7d5)
- [sni](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-9fc01974e3956c45c697aa96fa01d9561f837ca5f51871023cd162da605ea22b)
- [tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-8e30515f89b566c0046c0b656f5f1cae28136176409095156357cbe5f3cab9fc)
- [tls_tcp](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-a8cb93789e0c003784740b4e13da283bfc153d645902b5eb637e45de41c0c031)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-45284443551b024b171d2f15c5b2a1a11c45d73a87eb54ea5ae3ecce816903d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91a93fa17b19e973b35235ef81038b3509697a09aec418c3d10ff040cb8c9da7"></a>

## active_service_policies — active_service_policies / 99d71e54f243 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- active_service_policies

<a id="canonical-d82c009ba1a1aaa6002c1ec20e4a1e5ae1e81a34c6a82e063d94ade1b66c6175"></a>

Type: `"single"`. Computed.

\[OneOf: active\_service\_policies, no\_service\_policies, service\_policies\_from\_namespace;
Default: no\_service\_policies\] Configuration parameter for active service policies.

Upstream description:

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

- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d82c009ba1a1aaa6002c1ec20e4a1e5ae1e81a34c6a82e063d94ade1b66c6175)
- [no_service_policies](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-1959f2c19f4e27dd4cfd683029d504e1fca822c6390b85c034c7f8cfa01575aa)
- [service_policies_from_namespace](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-f13cd904bb7cda8dcac837323ed3c4fa516dcaf5b3266754719068b762ee4480)

Select alternatives according to the provider validators above.

<a id="canonical-3fbda0f1c9952ec31a09a4002a72a86bbb1e982835c73f4489d494a0fa468ea3"></a>

## Direct properties — active_service_policies / 99d71e54f243 / 3

- [policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-27835a803c165ea0f2bfd94ab0a663c3ac4c7197bcd48a0e23413e309b7466d7): complete subsection reference.

<a id="canonical-3240d2fee9a0786c0c907df4fd0433f6561d25d78707c0ee825f3c5fb076cd38"></a>

## Next pages — active_service_policies / 99d71e54f243 / 4

- [active_service_policies.policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-27835a803c165ea0f2bfd94ab0a663c3ac4c7197bcd48a0e23413e309b7466d7)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-27835a803c165ea0f2bfd94ab0a663c3ac4c7197bcd48a0e23413e309b7466d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b07044eb8c4bcb32d8226e683f9ca59d0f72f108c9154ed7417632a0aa83cc4"></a>

## active_service_policies.policies — active_service_policies.policies / 01af85b87a26 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-45284443551b024b171d2f15c5b2a1a11c45d73a87eb54ea5ae3ecce816903d0)
- active_service_policies.policies

<a id="canonical-ca4e72eb8d8f385236af57083d019ad38aea47a0bee96ad836f9f9fff4364e95"></a>

Type: `"list"`. Computed.

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

<a id="canonical-fc6272c19bf8aa54a1626f5a4d4a9b5824f30522b0ec26c3722637aef37f081b"></a>

## Direct properties — active_service_policies.policies / 01af85b87a26 / 3

<a id="canonical-ecb505853434fb11ad44de7fe8097acaa16f3d85182a3302a3d79e659462bffa"></a>

<a id="canonical-bd994d3e078898ae9c96b12d240c924518865bcaf07e53efa6184dbb4f90eefc"></a>

## name property — active_service_policies.policies / 01af85b87a26 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-f7826859f1c43c12396594c22a0dbbe1c9f709369714cedc1e374b77550b3a41"></a>

<a id="canonical-2b0f09a1a0f020aade29b8a29e9cd48c603e7ffff6dd9ed9ee0e85b4b88a37ef"></a>

## namespace property — active_service_policies.policies / 01af85b87a26 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-6327a96cf9f3295181e2d5ec447431d5ce419f7629e6ea217d051621c1d2c2a4"></a>

<a id="canonical-871e11b717b5a76713ff66adc2d8eeb7fe491fd36f2efca8038ec3fb515e81e6"></a>

## tenant property — active_service_policies.policies / 01af85b87a26 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3562b3ea58e9e8ee6a00af5f8e55cb7808300d18c012dc49c83be9f6084a808a"></a>

## Next pages — active_service_policies.policies / 01af85b87a26 / 7

- [active_service_policies](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-45284443551b024b171d2f15c5b2a1a11c45d73a87eb54ea5ae3ecce816903d0)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59242700d7ca83296fa63366269e57628348f49659a40df0bb95d9e7fc8b0708"></a>

## advertise_custom — advertise_custom / 7e6261a9bbc2 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- advertise_custom

<a id="canonical-471af99fcb75d952a306832f9d95c86b7be81d5a2b4f27938e79aedd8c85da87"></a>

Type: `"single"`. Computed.

\[OneOf: advertise\_custom, advertise\_on\_public, advertise\_on\_public\_default\_vip,
do\_not\_advertise; Default: advertise\_on\_public\_default\_vip\] Defines a way to advertise a VIP
on specific sites.

Upstream description:

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

- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-471af99fcb75d952a306832f9d95c86b7be81d5a2b4f27938e79aedd8c85da87)
- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a036294937822daf472e5e0cdb0113c257532b5591e640bafd0f49e736f09d5a)
- [advertise_on_public_default_vip](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-5bf02f2574ca18e1572f4a5e318efeb20e4f23b884c4cdd1a7053091eea130a7)
- [do_not_advertise](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-01cc0932a46a5e320283e534345f8c825a17d59bebd68c99ece71365db19d7c0)

Select alternatives according to the provider validators above.

<a id="canonical-45cfadd0dd712a3b4948b4532f609188cb7cc45a7363992e107102cc375399fa"></a>

## Direct properties — advertise_custom / 7e6261a9bbc2 / 3

- [advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4): complete subsection reference.

<a id="canonical-3b94df0f55742e3a1d8a017950b71b8c01d08e74eb7cc881b3daa75f2c672e59"></a>

## Next pages — advertise_custom / 7e6261a9bbc2 / 4

- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92b9c813a08891830973d7af8755ceb0a3ea86c3894091b96229145ea6af3a0b"></a>

## advertise_custom.advertise_where — advertise_custom.advertise_where / 0a3a2b5ec751 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- advertise_custom.advertise_where

<a id="canonical-f4b5f5d26e5051fa35327f18d355aa9186b5924a0eb7874e363f1452f5fe42d9"></a>

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

<a id="canonical-23cb748e57df715106c67a70bdcb1f01aafdff98b49b0c85fc1c17a52f906bbc"></a>

## Direct properties — advertise_custom.advertise_where / 0a3a2b5ec751 / 3

- [advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3d462f30c7ca95052b630aa6a1d624146093f359d73c94e56a3bd1fcd91c2621): complete subsection reference.

- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-dfaab23f86f0112d17675483ca69348356e99ddff02deaf649669314dfff4cf1): complete subsection reference.

- [advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-46cc3819c4f398e12b559479726e49c8227d387a81e24240352f67d3f59a72e7): complete subsection reference.

<a id="canonical-589be668802ce6f2bbe966a5dcf9a65cf27b688bb7ea7c6142a5bb35d5749c7c"></a>

<a id="canonical-5c976240e57fafffb7e4eef14e90c49a59804bab72e31829e3a7270ba10e073d"></a>

## port property — advertise_custom.advertise_where / 0a3a2b5ec751 / 4

Type: `"number"`. Computed.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1471343496d3b52c8f68f76189e478a8797a7072266c437938d069e651fab01e"></a>

<a id="canonical-206b6c40d9b7cb68f095976aeda728333b2d45d7e082d798c1f560ec3735c70a"></a>

## port_ranges property — advertise_custom.advertise_where / 0a3a2b5ec751 / 5

Type: `"string"`. Computed.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6154cc131efd96f49fc548c9ef0b56096e0b05701254583ab3201d119d476c5c): complete subsection reference.

- [use_default_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-cbd6138b9711633dd692e7e87f5dbd5fbce6172a5ae8686cf2412c4103855eed): complete subsection reference.

- [virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060): complete subsection reference.

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1a5ef3da6635d886439116617c4a2eae9ca298a6c3426a83e74b422d51d7f640): complete subsection reference.

- [virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-34bc7328d542baac3689a826d753e66f29a0fbd47032aca079658a060c62ae10): complete subsection reference.

- [vk8s_service](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1e6911fe82c3f48cf199b6bdec55513186d4cd526dfe2980014df97511c7663f): complete subsection reference.

<a id="canonical-679905370da826fe80cf76f1a96798e9b7457125988648b2a591a7b20f7cf0f6"></a>

## Next pages — advertise_custom.advertise_where / 0a3a2b5ec751 / 6

- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3d462f30c7ca95052b630aa6a1d624146093f359d73c94e56a3bd1fcd91c2621)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-dfaab23f86f0112d17675483ca69348356e99ddff02deaf649669314dfff4cf1)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-46cc3819c4f398e12b559479726e49c8227d387a81e24240352f67d3f59a72e7)
- [advertise_custom.advertise_where.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6154cc131efd96f49fc548c9ef0b56096e0b05701254583ab3201d119d476c5c)
- [advertise_custom.advertise_where.use_default_port](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-cbd6138b9711633dd692e7e87f5dbd5fbce6172a5ae8686cf2412c4103855eed)
- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060)
- [advertise_custom.advertise_where.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1a5ef3da6635d886439116617c4a2eae9ca298a6c3426a83e74b422d51d7f640)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-34bc7328d542baac3689a826d753e66f29a0fbd47032aca079658a060c62ae10)
- [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1e6911fe82c3f48cf199b6bdec55513186d4cd526dfe2980014df97511c7663f)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-3d462f30c7ca95052b630aa6a1d624146093f359d73c94e56a3bd1fcd91c2621"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7c8f95ab6db20f62be35e7816fdfb2fee07e456946c56defcb6f9216828d844"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_custom.advertise_where.advertise_dualstack_on_public / 843103e6b277 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-adb1b2ff73be6620249351b0f71f7786c386718634701e3f44d1a6152c04407f"></a>

Type: `"single"`. Computed.

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

<a id="canonical-5fa1c05486ac42ab50f44d0908beb0d41f88b3043d957b689d8f9ae3e6eef253"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public / 843103e6b277 / 3

- [public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-235f800b3d6d97dced2b42252d1a6fc2b95c2916a2145ea464b1c5f594e7517b): complete subsection reference.

<a id="canonical-c412f960bd4ab31e5ffe9e071aeab1732eb8fab8ac049b429055558b5ed9f7da"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public / 843103e6b277 / 4

- [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-235f800b3d6d97dced2b42252d1a6fc2b95c2916a2145ea464b1c5f594e7517b)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-235f800b3d6d97dced2b42252d1a6fc2b95c2916a2145ea464b1c5f594e7517b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ec1d3dd498dbc807e04b968e72dafac6a28aefc7d6cd22b86f37fcf7b723c08"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / 0756a98d6e0c / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3d462f30c7ca95052b630aa6a1d624146093f359d73c94e56a3bd1fcd91c2621)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-29317668529c303b89e96e6a80ecee79c6be37432fd225877b23daefa6b0ca49"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-36fb3e759b6148b0326c663504c7051c82730a974554b4c4845e14656d7234cc"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / 0756a98d6e0c / 3

<a id="canonical-e0d27db3b216c5705f114150fd28088467112e7a282583387a9c08d7736e3d08"></a>

<a id="canonical-ff862a2d367e1eaa193057a63d0e55e5301eb82582f0fd033f012fe85ba39408"></a>

## name property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / 0756a98d6e0c / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-8840b4785b565a7e4fdce72b2615772f33bb4ed0e9da2c949018c5b4388ff70c"></a>

<a id="canonical-1dda66540f6cdb7171fbe57a63161138956cd57065865e41fd2dc2c3941475af"></a>

## namespace property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / 0756a98d6e0c / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-d6f0a79800094cb5b2d575932283d5924b3140990b2e8f448d994562460aa9d7"></a>

<a id="canonical-1996508517b04bb7c31c5aae5e6fb0c485a177df3a069e0e817414532c399357"></a>

## tenant property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / 0756a98d6e0c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-10d614843ff610a65f62c0e8396de7dabbcdcc7f5262220b6cbc39fd5e4c0ede"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / 0756a98d6e0c / 7

- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-3d462f30c7ca95052b630aa6a1d624146093f359d73c94e56a3bd1fcd91c2621)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-dfaab23f86f0112d17675483ca69348356e99ddff02deaf649669314dfff4cf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f571f0a3d91ad73180b9f1638d9278cea516663ef8e2b87f106b8b01ef9caaac"></a>

## advertise_custom.advertise_where.advertise_on_public — advertise_custom.advertise_where.advertise_on_public / af8873c9d11d / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-c2dfc0e73fdc005042660b79524a631776e8a138724c78a38dda04210dd3ebb0"></a>

Type: `"single"`. Computed.

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

<a id="canonical-21f1833bb2df66a34b182a4034c996fc76cf807e057dab4a22f183f42996f5ba"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public / af8873c9d11d / 3

- [public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a9f8c89364d3a1ac7bba2d15140159d0a65c43868f9f786b6c3cc50d77d12e66): complete subsection reference.

<a id="canonical-992935a1e83103a4fcc7815c78b9f2f64d2b7a27641f16645dd7764e375703a4"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public / af8873c9d11d / 4

- [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a9f8c89364d3a1ac7bba2d15140159d0a65c43868f9f786b6c3cc50d77d12e66)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-a9f8c89364d3a1ac7bba2d15140159d0a65c43868f9f786b6c3cc50d77d12e66"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96a4f39634c184c8fd7abeaece310477e7747f5265eff4ce71a1b30a2e78ded6"></a>

## advertise_custom.advertise_where.advertise_on_public.public_ip — advertise_custom.advertise_where.advertise_on_public.public_ip / 20072d8bd55e / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-dfaab23f86f0112d17675483ca69348356e99ddff02deaf649669314dfff4cf1)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-f3f2ccd5e38421216ba5e4b19d80991c1e8ea4852c5e18b5cf2caa7576ebb63d"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-9d2171b7110eca6009819abebd7501410ce23837872752ff4e89616b27766d03"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public.public_ip / 20072d8bd55e / 3

<a id="canonical-1b5f96665f81c3055ee9af89e5db5ffd25a383994ac72a751780fe668910c155"></a>

<a id="canonical-80e632e0974ecad8348d8e72b5b980f128e040ee798ce240a8e3b4826a065fa6"></a>

## name property — advertise_custom.advertise_where.advertise_on_public.public_ip / 20072d8bd55e / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-748021f63e6f80fbf698b410067c4ab5422081266395f717aa04ac79706d29c5"></a>

<a id="canonical-fa9cfe8f1fed4f5343d48cc4f6ab6b2bf72bfa8cf915bc627ce37ed964b9fcac"></a>

## namespace property — advertise_custom.advertise_where.advertise_on_public.public_ip / 20072d8bd55e / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-102c8fbed596557e281a980be9fb6af3bdfa6b7a40b16a19216d8fb49402eff8"></a>

<a id="canonical-03d4371a150d6248566d4d006e9a72bda1196b25266bdfe8954944d694de264a"></a>

## tenant property — advertise_custom.advertise_where.advertise_on_public.public_ip / 20072d8bd55e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-52f2f0e3eb6e74dc3b767938fc45146de69d433340a5533fada1c935f5487e9a"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public.public_ip / 20072d8bd55e / 7

- [advertise_custom.advertise_where.advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-dfaab23f86f0112d17675483ca69348356e99ddff02deaf649669314dfff4cf1)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-46cc3819c4f398e12b559479726e49c8227d387a81e24240352f67d3f59a72e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b409905e06074afba9701ab5f8cb0735a402045d9e93130ad28f7822674e314a"></a>

## advertise_custom.advertise_where.advertise_v6_on_public — advertise_custom.advertise_where.advertise_v6_on_public / 6b24d4f0e93b / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-bd1d695be46aa5531497323c6ca98a36111e6c78bf05b9a03a5bbb5b0315d425"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2fac2273eee7e0c0fed6fc1e4b0d7ff4bd9f44a4c926423e1a7ad0e9cfd7709e"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public / 6b24d4f0e93b / 3

- [public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d14b4039a0c8071696e8e847ec7c02ed68ee9a6123258306d03834a0e7323234): complete subsection reference.

<a id="canonical-9803ee02d0d0fc0967fc460d4f40ea1ebcc8c71560daef1bc1e432bb77e1e104"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public / 6b24d4f0e93b / 4

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-d14b4039a0c8071696e8e847ec7c02ed68ee9a6123258306d03834a0e7323234)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-d14b4039a0c8071696e8e847ec7c02ed68ee9a6123258306d03834a0e7323234"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48499d2e1bd44784a6c37e9decb77b60657983fda3b03e8ae7bb7c97ff0e2b22"></a>

## advertise_custom.advertise_where.advertise_v6_on_public.public_ip — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 18919d2bd253 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-46cc3819c4f398e12b559479726e49c8227d387a81e24240352f67d3f59a72e7)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-d3a4f017cd0b77e7bc0ed27c928d6d691cac755dff6bf9fc00b74b19cb3ae836"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-77cec09a10c1756bb88d7210da123d9f920d3ed93b327855351a07973137786a"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 18919d2bd253 / 3

<a id="canonical-0996c3312330c90c85c20cec94319daf9766175e16aba23f6fc71dc003129851"></a>

<a id="canonical-82c2448b4c5f81496814ec739db208ef25be714b2376081583a6e96faf10176c"></a>

## name property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 18919d2bd253 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-482f79575613fdfe556a1d95ed77134e4701e649076f2ed85b2bed590d1aa3dd"></a>

<a id="canonical-97777d39e20f1cb5c5e811e6172ef971d8cad28089391c0dfc153e123b7add57"></a>

## namespace property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 18919d2bd253 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-ef507ffc44c8bfaca8a861e4d31cc1a83a61ec6f73b5c2330d2b8595606415f7"></a>

<a id="canonical-b1b921e87c2d8c1fd35e19bbae39fe680acd6f57cd73733d72fe61b4ca5caa39"></a>

## tenant property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 18919d2bd253 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-6b5ccd5e6878b2e584b3fb8719f9052608959f491aa1f8d3d62088e71c8e2c2a"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 18919d2bd253 / 7

- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-46cc3819c4f398e12b559479726e49c8227d387a81e24240352f67d3f59a72e7)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-6154cc131efd96f49fc548c9ef0b56096e0b05701254583ab3201d119d476c5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff92e127ae0316caccac3fab77acc8629fccca3e32981fa6c511c59f7bc476ff"></a>

## advertise_custom.advertise_where.site — advertise_custom.advertise_where.site / dc70a223910f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.site

<a id="canonical-d809f6e94b6a258d25e40ce871e5f83a103127e06084bbf2b55deb35cc9556dc"></a>

Type: `"single"`. Computed.

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

<a id="canonical-34dc3231d47fb2e7697394f9367d6c1715fe38a807b748d21f4695b820986e91"></a>

## Direct properties — advertise_custom.advertise_where.site / dc70a223910f / 3

<a id="canonical-263e5da0c66fd7fd78caee8b160a2ffef11d8fdd93e1b84b1a4536c1aaaf9f3b"></a>

<a id="canonical-21e72691db975c4f98d66a2f8c8d0ab45ede232116f8a238d6d338e5a043e695"></a>

## ip property — advertise_custom.advertise_where.site / dc70a223910f / 4

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-5ed1e6f9be8027167c2c534ae98f40f512e4ad0e6f30c03f56d20fa75cf334ea"></a>

<a id="canonical-b67481ef0f83a2cfc6b4fe3a38a5d15cb9d90b3a011c0f0a275a67500f7e6f86"></a>

## network property — advertise_custom.advertise_where.site / dc70a223910f / 5

Type: `"string"`. Computed.

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

- [site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-7b35937a184c1bf4aa39f6c1428d17e974778d7a6b143251b7edfbec84eac4c6): complete subsection reference.

<a id="canonical-c485f312b4028495a82cf1425f0924d820225ecdb19c7b72009538976880d3ad"></a>

## Next pages — advertise_custom.advertise_where.site / dc70a223910f / 6

- [advertise_custom.advertise_where.site.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-7b35937a184c1bf4aa39f6c1428d17e974778d7a6b143251b7edfbec84eac4c6)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-7b35937a184c1bf4aa39f6c1428d17e974778d7a6b143251b7edfbec84eac4c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8df5388bed397d364164c8e2ce62bc57d3d300b7843762818676ce7114d1139"></a>

## advertise_custom.advertise_where.site.site — advertise_custom.advertise_where.site.site / a3cac9b989fa / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6154cc131efd96f49fc548c9ef0b56096e0b05701254583ab3201d119d476c5c)
- advertise_custom.advertise_where.site.site

<a id="canonical-5b1ecf73154021e3b508b541652dd8f0dcb5ac7c111bd831ff6d0d4da650a983"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-ab38b4ed73e0e2671625ea58a461f3f8291718d46fd4290332c6fa00b4a4bd82"></a>

## Direct properties — advertise_custom.advertise_where.site.site / a3cac9b989fa / 3

<a id="canonical-6865bcd2fbb369c94c602916bdf554ac44dbb4647024bdad868952060a46faa9"></a>

<a id="canonical-153d28af10cd61dc8400d959c671a9f31e25789772ac98a458f833d9d242be44"></a>

## name property — advertise_custom.advertise_where.site.site / a3cac9b989fa / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-7bc1b981b7da82bd46b321158ad1c7553f9f610daefdb432e3b677d47f84187a"></a>

<a id="canonical-9b9095235dd4fb5484ce94407860b71a0db7af0ba05a3efb75fd061e2cf7e878"></a>

## namespace property — advertise_custom.advertise_where.site.site / a3cac9b989fa / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-5a5e8c49b6558287791b83b29c5149652fdb4b02693994b5f4910c0dda98d126"></a>

<a id="canonical-9cce3149e3c4ed4cdd7dd3b4ce615a6a9a9ed3a967f3a0123ecf7f4eaee2e469"></a>

## tenant property — advertise_custom.advertise_where.site.site / a3cac9b989fa / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-a4537a672a873c113b9e07c5f3b14b6a123f5db8689c2db21772057cb98f112c"></a>

## Next pages — advertise_custom.advertise_where.site.site / a3cac9b989fa / 7

- [advertise_custom.advertise_where.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6154cc131efd96f49fc548c9ef0b56096e0b05701254583ab3201d119d476c5c)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-cbd6138b9711633dd692e7e87f5dbd5fbce6172a5ae8686cf2412c4103855eed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ca5c502966ab07bc0e75b1cfd3bf20aae0872ed7aa275cb3dad191f6d78b86b"></a>

## advertise_custom.advertise_where.use_default_port — advertise_custom.advertise_where.use_default_port / d37c1f7a1a00 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-c0c5285ca8e66556d0d9d118b406132edaffec1d66415656d9c3a9a8d9d9c762"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-297ec33c64b2f49e65f6b9384cfffded2e437c867dc022309faf1d3b27ed4803"></a>

## Direct properties — advertise_custom.advertise_where.use_default_port / d37c1f7a1a00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e37727027c6a249e3acdd83a81dbd97818cc60cd5124277259f722a0605a52a5"></a>

## Next pages — advertise_custom.advertise_where.use_default_port / d37c1f7a1a00 / 4

- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f509422d83c5e51f832817b0490e51299acfa2d9d0564acdf18c5ea2af55b8b4"></a>

## advertise_custom.advertise_where.virtual_network — advertise_custom.advertise_where.virtual_network / db73cdff8373 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-292d1afe839a4f286c37c6d7e0a7cadc23d414c47954572a260c495c80a19226"></a>

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

<a id="canonical-39400d71c9c2bd91ca9f73a39ed5f0045c2b3ef53c0e692cf5ffecd9a7127027"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network / db73cdff8373 / 3

- [default_v6_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-7259436d32018fc54e78f0c45316491778fa4471799b56ca8f0f0e6fa38f9007): complete subsection reference.

- [default_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a554924f6e6eb721b41a66411722ccb19c6b5f3bc801827c5e65beab15a7e2c4): complete subsection reference.

<a id="canonical-df14e0e037dec276ae2430dd4970cfd6df6f3ae68b477116c38ac6a40694af6c"></a>

<a id="canonical-553f50daaf62886f8f10c79e7219f27de357b54d3d28dc9fc45646b3b7055e46"></a>

## specific_v6_vip property — advertise_custom.advertise_where.virtual_network / db73cdff8373 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-ea70527ac5376409eebe7a773359dad049b3ee7f2a6cf49ca0a17ca852dff5e7"></a>

<a id="canonical-e636c7a5353385f4fba021b00b616c2e8600f910d1fc7a75c2c2b08f84d19c44"></a>

## specific_vip property — advertise_custom.advertise_where.virtual_network / db73cdff8373 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c71463730981f3c04ffd2bfa6066bac72c8e8a1202ddf589b55c69b2b68801ca): complete subsection reference.

<a id="canonical-ebdf3e05d631baff5a7bb96a88545022fcae34e03482d79df89880155e445dfe"></a>

## Next pages — advertise_custom.advertise_where.virtual_network / db73cdff8373 / 6

- [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-7259436d32018fc54e78f0c45316491778fa4471799b56ca8f0f0e6fa38f9007)
- [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a554924f6e6eb721b41a66411722ccb19c6b5f3bc801827c5e65beab15a7e2c4)
- [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-c71463730981f3c04ffd2bfa6066bac72c8e8a1202ddf589b55c69b2b68801ca)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-7259436d32018fc54e78f0c45316491778fa4471799b56ca8f0f0e6fa38f9007"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-877d3bce8290bad01285bd7476116cc08fe00d0d14a6af402cf3f989e19296c7"></a>

## advertise_custom.advertise_where.virtual_network.default_v6_vip — advertise_custom.advertise_where.virtual_network.default_v6_vip / dc420c1698f2 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-a5455aea40c21e399b666fed309d9c265c75416d59c223c249f93496a86bbc7e"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-d6566cbabd1c6415f221ec096ef283e904a97538a96b0668dcd2de7dcd59d435"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_v6_vip / dc420c1698f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f897464a7d012fe699ba0859f01703d43691b48aec0358c5414149efcb984202"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_v6_vip / dc420c1698f2 / 4

- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-a554924f6e6eb721b41a66411722ccb19c6b5f3bc801827c5e65beab15a7e2c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e92f14f95e66feb8b6cc7acaa4f08d24ca83a0de86e0d13d88ef5bfde14cb1b7"></a>

## advertise_custom.advertise_where.virtual_network.default_vip — advertise_custom.advertise_where.virtual_network.default_vip / 956ffbfe4732 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-061b46b0d7e20bee7c4aa56c55613fefdc55c7fe1360185fd63ce425f353c372"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5a82ca185d30be8659700a54865d7f69cdcea20e1590bbe8b1440a6e8dc3cf2b"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_vip / 956ffbfe4732 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-258dfb65231137f7852040704f4922b40cc623f6d5ca86471b48e0ed19256f32"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_vip / 956ffbfe4732 / 4

- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-c71463730981f3c04ffd2bfa6066bac72c8e8a1202ddf589b55c69b2b68801ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-573a5a7920b11d6c4d3ceb6a74a0cf5a29d808c0a86ea5bb818d176368e10465"></a>

## advertise_custom.advertise_where.virtual_network.virtual_network — advertise_custom.advertise_where.virtual_network.virtual_network / a5633f1d5b98 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-050cd0774eee544bd040435ec613eb2d98736708a32cdd8b56e2f9beda4cfac6"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-50d48771c4fdd7bd0bf2c57e3401aed22829501332be7c689cb3b46913e34693"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.virtual_network / a5633f1d5b98 / 3

<a id="canonical-8fecf36e3b4ef911f3012b36f4567f3c8c041a601000f6affcdf93119ef2a331"></a>

<a id="canonical-2bdb3a6f6e9501d81ce612d935f1d4bca9e24ef5f0105e97ba9e70e4d942d4e0"></a>

## name property — advertise_custom.advertise_where.virtual_network.virtual_network / a5633f1d5b98 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-30be27eb00c2487a5a78733c0e43c1c89de14fdb8e515d23615871f2721a3ad4"></a>

<a id="canonical-fc0fedd81ad4cbcd6db65214e95d9c80b603805ecf7137d8e93e463a45a12887"></a>

## namespace property — advertise_custom.advertise_where.virtual_network.virtual_network / a5633f1d5b98 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-90e43382e0da0ba254501813be7f7afd3d79ab19700841e1c6e10bc22e93d4ad"></a>

<a id="canonical-99a94d4c8c082183a8bbd8ad9010b7b61d6ffc2d0c0d4e24e0bc9d49aa73e8dc"></a>

## tenant property — advertise_custom.advertise_where.virtual_network.virtual_network / a5633f1d5b98 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-ab434e1395bdc2ce464891a8abcf57470ac92c103ee67e4650b6eafa66382e41"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.virtual_network / a5633f1d5b98 / 7

- [advertise_custom.advertise_where.virtual_network](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-e68e23c32c3c3f15a320ed4904eedf78c88278169bdbb38ce32097d318e1e060)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-1a5ef3da6635d886439116617c4a2eae9ca298a6c3426a83e74b422d51d7f640"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06c92b8679d5572cb4ce5d67e65aac06e8289aa973e886584f0dadf8c6a0ed02"></a>

## advertise_custom.advertise_where.virtual_site — advertise_custom.advertise_where.virtual_site / 4292b7e18889 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-34ebf29b4a091424a7c8534e643a08d001a7f32891ff7092133a27b54a70a49b"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2496c38f06aa1452493448580684b45f059d67098b4bd688fcff589c897df3cc"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site / 4292b7e18889 / 3

<a id="canonical-9615dd6dcc86060a65c25b26db73bc538be804decee62fcc0c3cbb5f82558e64"></a>

<a id="canonical-12e40d9cad3389d69a16c494514b6b4c7816edd1d32d66254892d21c9e6a6bc9"></a>

## network property — advertise_custom.advertise_where.virtual_site / 4292b7e18889 / 4

Type: `"string"`. Computed.

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

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-4b807a091c7b032f998f8e04e907c97f8f88ddd1bb298d6d1b08a9bd0b6b4d30): complete subsection reference.

<a id="canonical-049e87aeb4db904225edee1945d9a7cd56618efefd937f1542eb9a16dcae7404"></a>

## Next pages — advertise_custom.advertise_where.virtual_site / 4292b7e18889 / 5

- [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-4b807a091c7b032f998f8e04e907c97f8f88ddd1bb298d6d1b08a9bd0b6b4d30)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-4b807a091c7b032f998f8e04e907c97f8f88ddd1bb298d6d1b08a9bd0b6b4d30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d672dc58001cc702233073022b27273bc434147ba93dbb2391be8a6f2e52ec69"></a>

## advertise_custom.advertise_where.virtual_site.virtual_site — advertise_custom.advertise_where.virtual_site.virtual_site / b2efd0130465 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1a5ef3da6635d886439116617c4a2eae9ca298a6c3426a83e74b422d51d7f640)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-005d3cc6febfb0dc643987bffcc45650ba3357a5f80868678ab652d886703424"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-64f45774664145b9cf128c471a16b362dfc937c61ad96f50afad9ba0aa43f7ff"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site.virtual_site / b2efd0130465 / 3

<a id="canonical-925052ca6d49d6d1c52ec576e82b4d1087508b4a20f3c5bc61fc5669a3c441ba"></a>

<a id="canonical-4098891b0bab28b73bcf230583ed4ceeea3a9665c9c105c725b58061410321a9"></a>

## name property — advertise_custom.advertise_where.virtual_site.virtual_site / b2efd0130465 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-29641efe6bdf167004cc1a7e3886b753b921cb1b15e08385f9f82a510ab3b46c"></a>

<a id="canonical-adbf344b2408bd38f56e2dff35c541d4ce99c5350c17b8a30cdb240b67f50faf"></a>

## namespace property — advertise_custom.advertise_where.virtual_site.virtual_site / b2efd0130465 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-d1bbcdba119d2fce89d77672fdd8c0d9b3b94a0fa03d6d9853e4907475cefb8e"></a>

<a id="canonical-1656c15dad5e25d74438672d8997225d9d4ec03be5e3e4fbd34041c4a9d17152"></a>

## tenant property — advertise_custom.advertise_where.virtual_site.virtual_site / b2efd0130465 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-69d7d554226d6e9c3093bc426f0c2d132f942a058e19850630062800d924a25d"></a>

## Next pages — advertise_custom.advertise_where.virtual_site.virtual_site / b2efd0130465 / 7

- [advertise_custom.advertise_where.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1a5ef3da6635d886439116617c4a2eae9ca298a6c3426a83e74b422d51d7f640)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-34bc7328d542baac3689a826d753e66f29a0fbd47032aca079658a060c62ae10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff27816ed34dcc7489b7952cd8b38cb5fa97fb931b3341b3cb27317a33b0cf47"></a>

## advertise_custom.advertise_where.virtual_site_with_vip — advertise_custom.advertise_where.virtual_site_with_vip / 508ffa05e3b6 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-e33dddab6ab70b9eee12a677281f2693498435d927ff0d155910bb9d205ec39d"></a>

Type: `"single"`. Computed.

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

<a id="canonical-354ad7445bc6ffe3a3b70cce004b542ca36e0808b8d4446f1e75877d73c34ad0"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip / 508ffa05e3b6 / 3

<a id="canonical-3a29474a7e433e733da617c4782e5decee850bc2de7cb97469fda5d97113ef37"></a>

<a id="canonical-d03de9a2113411200870ba6f5618c8cd1c401511fdcca3172f85bce13840cc55"></a>

## ip property — advertise_custom.advertise_where.virtual_site_with_vip / 508ffa05e3b6 / 4

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-e7c883e3a7d19c3a74542f16ab4ca61a8ed23d9f41ec2da5a50f1c18753c2b68"></a>

<a id="canonical-4272768c17145be083a86b81571c395b70d3d3cd57dfa4f261d4479ca188184a"></a>

## network property — advertise_custom.advertise_where.virtual_site_with_vip / 508ffa05e3b6 / 5

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

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

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a82797b69349dd76f7412f4c86417f35ad269252d5f96f040a8b985ff68ebb54): complete subsection reference.

<a id="canonical-1d435b74ac142ca9840c1b16fbf8600d300fe2796791cc366021c40e013fbea4"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip / 508ffa05e3b6 / 6

- [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-a82797b69349dd76f7412f4c86417f35ad269252d5f96f040a8b985ff68ebb54)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-a82797b69349dd76f7412f4c86417f35ad269252d5f96f040a8b985ff68ebb54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93325b5226e4c9ef12f3f02118105eff33d764b0c0111c24c4dee853a88b2b95"></a>

## advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / abd6938182e2 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-34bc7328d542baac3689a826d753e66f29a0fbd47032aca079658a060c62ae10)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-37556c8801ff63728726da7ee4488bce6f90177c947f2bfa9ff4026dec361fca"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-16bab7f63c4654ce0a4124a91bd2231cb7fec411f4e2fad737075830a5827a51"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / abd6938182e2 / 3

<a id="canonical-53e8fb33c82ebd73937213d4fabae679789b21e17b713fcfcbd16dc7dbd3c345"></a>

<a id="canonical-aa9c485a52c530203d54736f048474a6c5292cb73e799fb5ae831d1c887baa2d"></a>

## name property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / abd6938182e2 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-c00fe8e921bce7863240cf95f20dab00b56b179f2093dbd9c6cf3a8b45815901"></a>

<a id="canonical-e20d8db44730aab68993769867b15335889e42b1d5cfa36b1534d2553c8823c5"></a>

## namespace property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / abd6938182e2 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0b496fa15b2fb8707a8e6abe810e831a48f52476a1f8500e7c16f3815da38648"></a>

<a id="canonical-3b743818776a4238369e5f0f2bae5e0ea4c6a3d360efa369a44d4134dc7da32d"></a>

## tenant property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / abd6938182e2 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-172cdd381377e77285d4379241221473ad85df80e64f9c4c5677559ffd9f737b"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / abd6938182e2 / 7

- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-34bc7328d542baac3689a826d753e66f29a0fbd47032aca079658a060c62ae10)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-1e6911fe82c3f48cf199b6bdec55513186d4cd526dfe2980014df97511c7663f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e6a530e1019f77388398a9487a7afd4740548231a0cfe18bf136fa84329ed2e"></a>

## advertise_custom.advertise_where.vk8s_service — advertise_custom.advertise_where.vk8s_service / 93a298df5a59 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-b5fab35e1014938779a23334b010a9851cc21a197be6f85199f06c6d34a598b0"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

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

<a id="canonical-846325eab51fd3ea66badebad1d74ba1dc2b9e51ad8e5d3104491edcb485fd65"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service / 93a298df5a59 / 3

- [site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6eaaa669607edcbc7f817a84706e0b18929673c5d9cfc608bc7ce4aab632dca9): complete subsection reference.

- [virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-9e3c8a5cecc0ca05a5cbc138c1227e2547fae623ec7c1d3293a99fd7ecafd2ef): complete subsection reference.

<a id="canonical-b39dc89e57fd0d39dc09372c8601cba424aa7fc4a4ec0f4cf83ca7d3a72796ee"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service / 93a298df5a59 / 4

- [advertise_custom.advertise_where.vk8s_service.site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-6eaaa669607edcbc7f817a84706e0b18929673c5d9cfc608bc7ce4aab632dca9)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-9e3c8a5cecc0ca05a5cbc138c1227e2547fae623ec7c1d3293a99fd7ecafd2ef)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-6eaaa669607edcbc7f817a84706e0b18929673c5d9cfc608bc7ce4aab632dca9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb78ba2790c96f915b7611a17b7bfa5b76a2af062aecda6ef37327545575ed95"></a>

## advertise_custom.advertise_where.vk8s_service.site — advertise_custom.advertise_where.vk8s_service.site / 7c392c1fc606 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1e6911fe82c3f48cf199b6bdec55513186d4cd526dfe2980014df97511c7663f)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-bcb2af668ba2210029eaa253ddff86c56c48acd22a5e411a20ca24e548a75d50"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-3add91e65e663d125489668974ef72038c71c26aa0839d7adef101962cef0865"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.site / 7c392c1fc606 / 3

<a id="canonical-508462d0461401372bf6525581b8da21ef2fff10fe5074455a941e72053e9002"></a>

<a id="canonical-dc58d0a2469750e6160c832674f7961cc1023947e6dea71a03ea33cc841acb09"></a>

## name property — advertise_custom.advertise_where.vk8s_service.site / 7c392c1fc606 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-df05d1b860cdd37b1dd088b005c2fe7f3e79d3add30e0f9fba226e83d15030b1"></a>

<a id="canonical-6738d714dd77ae50228b0e78c2a3e0dc73ced5f2f62c73cbfd6ec83265047f3a"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.site / 7c392c1fc606 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-18b08ec442a92b323130d8cf2a65b2c5abccdceb509944c3d48a66668f8cc850"></a>

<a id="canonical-737482e9178f181a3d8c7b962a53e3820fce3ffc643b109f882dcd5347bc151b"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.site / 7c392c1fc606 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-b633315c158642923f0f2a6313921180292e1c9302d38af51edce8f91e3233a9"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.site / 7c392c1fc606 / 7

- [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1e6911fe82c3f48cf199b6bdec55513186d4cd526dfe2980014df97511c7663f)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-9e3c8a5cecc0ca05a5cbc138c1227e2547fae623ec7c1d3293a99fd7ecafd2ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c72b8a8dbd412dbaf5fcaf88b516472ff1d950eaff596c3bd8dc3690694251d4"></a>

## advertise_custom.advertise_where.vk8s_service.virtual_site — advertise_custom.advertise_where.vk8s_service.virtual_site / c9785342ceb1 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_custom](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5369a8063b3ad84afa4b51f6a2a231a762e2a2207952933f53643da7b6f8d800)
- [advertise_custom.advertise_where](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-aa5fc2cd403f4b5c7ff3aad1dc7ad13537c90b8950820f79acd82c594d64ceb4)
- [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1e6911fe82c3f48cf199b6bdec55513186d4cd526dfe2980014df97511c7663f)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-d5757c3b5b3740647933b9646f59ad11fd82972986d34117a6d463d93d2e1b99"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-f4207e716563a82a5c66a8d2d1fe9d5e5e3d956dffe1c374a506096ff92b81bd"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.virtual_site / c9785342ceb1 / 3

<a id="canonical-e168397ff3465af233242a93540c6498985ca2d0c643d5ea966fc86d3703d8ed"></a>

<a id="canonical-da570cc8ee52c23beea364cdf37fbd1cb216c6ec4285cdd728e29b899a0e44d0"></a>

## name property — advertise_custom.advertise_where.vk8s_service.virtual_site / c9785342ceb1 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-06d989a40bf48e74992a4af3b0fabe4a27f5d7969c6bccaf64dc8d3d17c512c8"></a>

<a id="canonical-32381c15519cb1fc4a098056059cf50448c51bb04cf8bd71649e34b6fa088367"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.virtual_site / c9785342ceb1 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-62276a01ee1d3e94d059313b59f296b7fc18794461374a8e578733f9dc649d4a"></a>

<a id="canonical-33433c43e4f58fe99a96e7c1af2439cbb58fcdd97b333e5f0dc7ecd0b7c2873b"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.virtual_site / c9785342ceb1 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-a90fdf22ae9bd9c24e91d44e739a0d94b645b713f62ec1650bb91526a3eb3eb5"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.virtual_site / c9785342ceb1 / 7

- [advertise_custom.advertise_where.vk8s_service](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-1e6911fe82c3f48cf199b6bdec55513186d4cd526dfe2980014df97511c7663f)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-5684e25de901e83fdd06b403aedef7b7952f266b34a2576f9c1eabbbabb5b10d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78772db94809d0e8b4233df6858f58a272fe30642b04620311c3879bfed8a9e3"></a>

## advertise_on_public — advertise_on_public / 51b2c2ec729d / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- advertise_on_public

<a id="canonical-a036294937822daf472e5e0cdb0113c257532b5591e640bafd0f49e736f09d5a"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1b261bd061f9325bad1261e69bc3b6957a97a2a66d5291edea37e1977d1aaf8d"></a>

## Direct properties — advertise_on_public / 51b2c2ec729d / 3

- [public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-bb521c582e424d4fe702c1ae19c4b81fc3909f1d3d22f97b088f634d0bb8f1e2): complete subsection reference.

<a id="canonical-a2c543b651b70cb2db9a3c8bc03af51c6ec503cb36d416796de77edf156f9809"></a>

## Next pages — advertise_on_public / 51b2c2ec729d / 4

- [advertise_on_public.public_ip](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-bb521c582e424d4fe702c1ae19c4b81fc3909f1d3d22f97b088f634d0bb8f1e2)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-bb521c582e424d4fe702c1ae19c4b81fc3909f1d3d22f97b088f634d0bb8f1e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef9af24a675909cabb56c8d3eb3fec66697a0a31489a8d96d2ea17519139bd56"></a>

## advertise_on_public.public_ip — advertise_on_public.public_ip / 13db4c8fec6b / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [advertise_on_public](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-5684e25de901e83fdd06b403aedef7b7952f266b34a2576f9c1eabbbabb5b10d)
- advertise_on_public.public_ip

<a id="canonical-b2f0a5d81ea2b5cc630b6fa8c57aa538104f8f9eed9d411a29c43dd8f6645813"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-63215f325b4f80fc6fffe7b8a951883a272b656a25b21aae0acd7f23a078ffbb"></a>

## Direct properties — advertise_on_public.public_ip / 13db4c8fec6b / 3

<a id="canonical-6a30a73c3d1f6d27280d801f9aa9b6cc6bc5035e137653f2d73492d9ca6a8a07"></a>

<a id="canonical-3ce4b9371724b2fcaebf04d17b4d464a38fac407a8815aaa9be5c351b226494a"></a>

## name property — advertise_on_public.public_ip / 13db4c8fec6b / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2e50ccc0fbbf69023d0a85ed6a39df94387f175d2ed145ff7df167bcb35a8a86"></a>

<a id="canonical-1eede8c5114992c7944773594470d93be17294f741919cf930b9cdc5d5ec5f4d"></a>

## namespace property — advertise_on_public.public_ip / 13db4c8fec6b / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-bb99f1a3addfecee27de11f7a0135ea56995641c16a775d427eacec7f490bf43"></a>
