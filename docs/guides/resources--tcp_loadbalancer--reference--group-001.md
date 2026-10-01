---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d48046f8801d9ba6ecda4f0442f892844be82f97ed0c0f63c6ec7fa3ec2aa11"></a>

## Property reference — Property reference / 89d73e3380a2 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- Property reference

<a id="canonical-b54528bb9f1610b4b33473eba6f30f5deffbd7c1f7db690c79c0e1bdddfa9b54"></a>

## Direct properties — Property reference / 89d73e3380a2 / 3

- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-9bf84c084bb7532e7035a4454e2e392860fc4eb33b7dba834871fadf5674c8d2): complete subsection reference.

- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c): complete subsection reference.

- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-5145a903bc57dcc49c9709e298e99c068fe602cad3de3f2ab20093ee69049729): complete subsection reference.

- [advertise_on_public_default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-ff42e920406b0375c6cbaac8ce29eedff4194fc6bf1fc0585caa1787584c2174): complete subsection reference.

<a id="canonical-9d02c3572755948c8325d313dba332b91760c996e14aca0f2cc07cbb826eba1f"></a>

<a id="canonical-de9ab56e71e2052625b01191e2fe0370aa946911ffa1a6a513eb68b19ae0f6b5"></a>

## annotations property — Property reference / 89d73e3380a2 / 4

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

- [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-6fc17a77d48f9cf85e426e65ef541a24a263cb7378eb08dbdebeaddb4c6c5404): complete subsection reference.

<a id="canonical-24c0ec8f441a6551b0dcb4174fa9b2be7817d1695177b12ef4bcc6991a3099cb"></a>

<a id="canonical-a7f6449a1ee59f6ca790d337b3bb8afe923e31e06053fa996ec49a2f71fef339"></a>

## description property — Property reference / 89d73e3380a2 / 5

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

<a id="canonical-e5fe0854d330b6b5e4f43a8c29b75aa474d59fb6cdd7d85e298e029b3c607736"></a>

<a id="canonical-6d0f3aeb55b635d94e90c40ade8d67bd92d6d934dd442cc949c5061b6f7e66e3"></a>

## disable property — Property reference / 89d73e3380a2 / 6

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

<a id="canonical-711ccd0e79b67ba4dd098ffb67ad6d5699f978a573bd18491bfe4ec1c7117d62"></a>

<a id="canonical-583c403e8e10a8c673da612c09a773cad7a5c4024cc29ba4963b4a4c089c52be"></a>

## dns_volterra_managed property — Property reference / 89d73e3380a2 / 7

Type: `"bool"`. Optional, Computed.

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

- [do_not_advertise](resources--tcp_loadbalancer--reference--group-002.md#canonical-ba9e95942d760250bb35617997afdb3d548b484a3b5ac545e364fd505c61a698): complete subsection reference.

- [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-ba5e7e3c123f000e5c21dcf85c3f1ab41075becc703e6fcc8f87971d9c0fce07): complete subsection reference.

<a id="canonical-b52ac0c5fba86f8dbcbd034b2c05669ce13a6912380f6133ce201f7095b9d971"></a>

<a id="canonical-cd3e47a3ddf48488674f75bd0a9963d51f02aea1db98c06e5ea5a2ca4e6a2bba"></a>

## domains property — Property reference / 89d73e3380a2 / 8

Type: `["list", "string"]`. Optional.

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

- [hash_policy_choice_least_active](resources--tcp_loadbalancer--reference--group-002.md#canonical-edd915446ceae49d7455932f6f248f30860d874aec8e428a7dd297607454e0bc): complete subsection reference.

- [hash_policy_choice_random](resources--tcp_loadbalancer--reference--group-002.md#canonical-9f8669f3e8de1af30d714b546776d46ef9bc8683f210f0b322894e386a825f0e): complete subsection reference.

- [hash_policy_choice_round_robin](resources--tcp_loadbalancer--reference--group-002.md#canonical-85cdf1528965a6285ff2c75a0267cb449348c1ca88921d060042419de75f1ace): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--reference--group-002.md#canonical-dbedf978ab25250f8b6bdd8d76481f4d8b6a7e6f331b982a6fc699ddc46b8ac6): complete subsection reference.

<a id="canonical-ad9d4e0662a636c307a9befacdcbff8dd81740594f1d0485f97fb10f87e0c8f6"></a>

<a id="canonical-a69dfc1e15d6ecb1a08e728a2604df0b00ff58ab9f217c77f56e24d8b749d7a6"></a>

## id property — Property reference / 89d73e3380a2 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1d5a7b535ec4b033e7a05967b300c1069364f692c09e84eb94a17d30e7d31cf6"></a>

<a id="canonical-f18918ccae398fb141721ba7e957c76a4ce36b5d7163e1ea7e0a40673f4b01eb"></a>

## idle_timeout property — Property reference / 89d73e3380a2 / 10

Type: `"number"`. Optional, Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
Server applies default when omitted.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(4147200000),
}
```

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

<a id="canonical-3e56b1d1c7a072840fbe2c104be165dad8f4834ebfa870ffda216e1b64d36926"></a>

<a id="canonical-6678957dfb5dff0c81d052d0e93bf8108f174ccdd43c364daa3d34e67c028ab8"></a>

## labels property — Property reference / 89d73e3380a2 / 11

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

<a id="canonical-134bf295899c5393ce4c482ce2f145a4dbe1f9dca34fa807ba0f0b65d77c4477"></a>

<a id="canonical-6646951ed5cb2e75c9991040120d727bdca9840f8202de360369e92ae2a2e4dc"></a>

## listen_port property — Property reference / 89d73e3380a2 / 12

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

- [listen_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-134bf295899c5393ce4c482ce2f145a4dbe1f9dca34fa807ba0f0b65d77c4477)
- [port_ranges](resources--tcp_loadbalancer--reference--group-001.md#canonical-a44c3b980811a5a4d972ce9f74b0f6f5c339b9af307beeb15a5cbe6123e2d1e6)

Select alternatives according to the provider validators above.

<a id="canonical-e839c845549573ff2bc9a72d758f2bdef34f6c4a15c79cc525d497fb4f3ce3c0"></a>

<a id="canonical-2d27bb021aa7260f75c24386ed5e69b2ccb7e01e6cd2a5866e7b91beb66fb053"></a>

## name property — Property reference / 89d73e3380a2 / 13

Type: `"string"`. Required.

Name of the TCP Load Balancer. Must be unique within the namespace.

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

<a id="canonical-200156ad2e9e49e69320b386ae2bf49dff70227e464125f1d4f07eab03c8af7c"></a>

<a id="canonical-7589dd6cbbc5b2b34039e340b996da9574ea07f4eecdc5dab992a00c70bf8563"></a>

## namespace property — Property reference / 89d73e3380a2 / 14

Type: `"string"`. Required.

Namespace where the TCP Load Balancer is created.

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

- [no_service_policies](resources--tcp_loadbalancer--reference--group-002.md#canonical-25ccd69fc2452f16f34ad58d3b9af414cbf6a2ca1af6864921afc4c6286aca01): complete subsection reference.

- [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-3e6782861d51d0a9e5a0f8e8ffa45db8737335ce850aafcfc19a13424444586b): complete subsection reference.

- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda): complete subsection reference.

<a id="canonical-a44c3b980811a5a4d972ce9f74b0f6f5c339b9af307beeb15a5cbe6123e2d1e6"></a>

<a id="canonical-27793667955f43c2ed71d456dcc1bdb31d3914d461d417b6709ef641ef21c450"></a>

## port_ranges property — Property reference / 89d73e3380a2 / 15

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

- [retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-a3f2eb692d42a4518e81b4d2ca346c903a6e9aedbc28daf1e229ac85d2217fbf): complete subsection reference.

- [service_policies_from_namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-7202f6813ce28cd3211969f35f4a551cbf7c6921c78f64f4a67273125a5d0ceb): complete subsection reference.

- [sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-65d837fc6b970ea897c23eaec474f1cabe159cddf3ae011ef00b692480d12e3a): complete subsection reference.

- [tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-a8b8d21b3fa7b0073cd75420b779305e70da290853df37c205c279e25e1912e5): complete subsection reference.

- [timeouts](resources--tcp_loadbalancer--reference--group-002.md#canonical-6b17fb1ab324a8664c2799ba750302145ca7f5afebaf304cad2896a517809a45): complete subsection reference.

- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8): complete subsection reference.

- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd): complete subsection reference.

<a id="canonical-a9401fc400ac2240c03e2f829920c2b9ac61c8b606b187af5fde1dc9f4004565"></a>

## All schema paths — Property reference / 89d73e3380a2 / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-659f5423c72f84f3eb0fa8f83d9c428401ff14240ac8f0c2b27e6d2cd6003907) |
| `active_service_policies.policies` | [active_service_policies.policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-49a0456a0637522726e8d341b2ea2ac30b4769477b1f735a494809c8cebdfdca) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-f24ba9ab28187514d519f5fbe2bd9557f74e6ad7b407b6e1f533551e5d233dc8) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-43bca3cd4cea136acf7aaf1bafae271c7118ac4433e97c3a605a4104e6feb094) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-95d3b8da8ca3707890a9bae5f84ac49276ddb0bacf93ac4206a67c9dc01b18b0) |
| `advertise_custom` | [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-71eca36c430c121b5e87e6eeb5af43b65e2442766d24f4415d1e766593b09da3) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-dd0cd5cb2812c4b73052a58cf05ee7affd4e34c5380a1e9a662f9240d5bfaa2b) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-291dd98f0763be52fba0d7672a720aa937b4a80a71762bd88e9d317164976490) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-fe06bff3ef0202b3b9edb4b60acc930049d36b3a3a5c65b77c4ce0807f7d4655) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-29ea2ae15b43af37ed7f5d762925ab2e5ba722220517e90422f75ce2cfd629b4) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-fb4bf958119068434eee1f2369f5fbaefa4d03f3ce4f2ccccf80df6803d00822) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-3789988360e283d95e87a2128b44d2ee1b4bb15a19bcc22a40d3ca67b92eebc2) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-6180b2610a3212c313e0f77d00c5282dc128e230fa9a07ce79581bc43bb755c5) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-3ebcf6f612643a46eaa5921a0a224d6fee882cc49c51f0c3cc313803ad6557b6) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-f0de4b313ead309a358812482220f88a448837ca438d969494457ed3bbe78b43) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-ad9c0888c2a3076ec575b1f16248694a7edb5beb6feacbd4bb2cf218dba511ce) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-63091ffd354cf54f2bb2e9a3bf183496e053f6439ebabe1d37ab9092e22a8f59) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-b3fd43f92eb4aca095d337a192b89f0ba077a13d1120a4425fb9d6046bf05d74) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-0619eefe4a5ad7da60383f6c575c6a9ca58fd23b00e6eff543a2634809658db4) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-777d6cf9f7c3a887ccb66126144705f29fcc9435bb34a3d58b5b75710828ae25) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-bf1a3749adfe6183a84f2fb24b6d5e259ff9c7f944d212d6642ba7b83d2d5761) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-14e1c8ddd3f34add0917ab506e7471c62b40156cd9e39838d03dc7f975e729dd) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](resources--tcp_loadbalancer--reference--group-001.md#canonical-827eb5557ec6f24b8be165381e45de306aeaadd0ea24d583da5da9b831d090a0) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](resources--tcp_loadbalancer--reference--group-001.md#canonical-88bc97fa6bb7a6f707966fd5da3912eb1a59279445011821d1728290daa3168b) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-2f0a65afd5f3978dd9abf53ac9f0dc2e73a2ee1f05ef35b2ce459d6a54afb3c2) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-57b308427e9ad8c7c4ccf3a0ef586c9cce17a6366a42df1a246a6c36b4338f4a) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](resources--tcp_loadbalancer--reference--group-001.md#canonical-c55415264b949c80a59ab84353415f5580d3e2c8d95347e08be19a05acb76ad1) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-3fde104ab7367e6d1e251083ec2cd5c0d8a3b82c2d32233b575717497f3785e2) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-64bb6c84727a1f847696a0f8d0d55468b57b8e62b7be1b258a526ccd3ac2781f) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-8f68df4810419f2ce4094c99876ad6afadaddd3c43dd63f6709c80806b12ef98) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-d75db1b737736af546a474d15a9783fec0ce84f041c5c162fb40162a92c9acac) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-9862aa1afb18062aa24aff6f912aca4d21e723d2f02b7f2f6075c9b0b5d139a4) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-a3d2da93488bbcbaf017009990bbf728b2947fbeb9d1543ac7d1c3498b77f5a1) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-b0b3cbf1e837c88e351f808375ec9f754877c74c1e2497abce8db784b6dfb131) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-a7fe050bf3d2e0de06147621d44334a04d185a14594803bf6e7d8d9fa58ceb41) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-09aae16a4b6b41f55898693ae609b2f340a30f3f0a43cc7ebcc448ffca38dbac) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-2afad266e7a0c6777028134c266a106031eac21b564ef7765823ee859f7492bf) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-87ce824dc486a9f4e16571eb286176f0fc12eb27ca10e40455ca7ff7a0c176c9) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-95fa2b429fc440fe289fa1df9717a6624d5e0d6e26711a357851a68827577aa9) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-a4da37b53213739c24d36ad0008da4e94eea22352d59684796d7cea8a676b902) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-f0da9baa497531c47b96394cd4ea9e3485a3ea84ebff6eb7f997c8ab95f93b31) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-91308e4c08b0cf943266116f2e880aca3f7faa10b525aecf3f34e82e58cf8e10) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](resources--tcp_loadbalancer--reference--group-001.md#canonical-b1ccad1b9e60ade621dc7df6a19fbc807781d71d69d900d69d6b8a2c09cb7ba5) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-a5e9ab30ffca5358d2f840a3d99fd4630c4cc5c91fcd91e547fba8c79c6d7868) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-f55c8f66c97f640e2d8f9dba3828617327c1baf54f94485d76aba8fd3af2c64a) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-e199e622ef4394de26c5b9b4afb20377af3b734fcda743d3cb1e0a776872d64a) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-0821e390c7016706e7ab150b035165079dd96a2a4505bd2cb9065ef5994abb2c) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-3d162dad8162c94fba8949e9d165c3abd17e3f0bd7c8665a786913cc24f879db) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-75001a100eae707748ff933fc4fd5cb85c0ad3aad348054b75c7c856cbd5140a) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](resources--tcp_loadbalancer--reference--group-001.md#canonical-608829ecf8f91e9783a38f280ddc944f62346907aef9961d784c5a0d1c03fb81) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-41899cb0ad5ce13bea67726b5dd461f1d384c5a144dd448c5f443905b5c32ac2) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-f06b11ed8a7101a252a0e64f776ef6903a5c85588a98c0a76f6501f7df0a81ac) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-030fc97b5d7a18e735e09349908012e25ad49e2784af454e9730a734d19f712c) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-b502cb4dd2f99323240961f076c78c7e027ff9d210f3170c88604ceba7ab77fe) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-8ece6b0320aec582deba74d48114afb23e82007e8f5751c3c6097b9cd2461ddd) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-cfb0afb92d9af0de5e0d87a0a536da056d4776f9b4120604572e8bf5272fc371) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](resources--tcp_loadbalancer--reference--group-001.md#canonical-dc3ba9ccce2921b7f790c770e51c8ccc716a8e0aa14ecfb7617918b8d0a73195) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-6e7e538357ec90a5be6ec3a498bbb29efb835928f249bbcc0c0ff4ea5e22c7d0) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](resources--tcp_loadbalancer--reference--group-001.md#canonical-00dac523dbf767ff4a55e9fc800640aed85a29033f0c443186b4dbe6ddd0202f) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--tcp_loadbalancer--reference--group-002.md#canonical-966d0364d83a9ed98fcbfe5cfb83e0eb0f634809b46f91110b61efa4768e9d46) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-15c3669a45b81e82bf8d58c62a26d9cd769d96656d19273e08cc933da1220350) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-56cc8c0463b677f58f80b081e530848dbfe3e338fc8477c70368b179ed87da2a) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-a06692555aa0a9222ae73ee852b7d52571008ef5eec10162f1bee11889732db2) |
| `advertise_on_public` | [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-32ddddd3b88cd727dfc738e4a881d56c2f176fb82c812cd6b5af1a6bceba73a5) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](resources--tcp_loadbalancer--reference--group-002.md#canonical-baa551955ff9311107bba0f7822fdcdde5a5a0f1f1992bbee147f7b3a5931e6f) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-0283f86159db777fea17a1ab85b4fc403294497eaea5705e4b96b62c30830f9d) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-7687b3dadb92306fb8af8a2faa861c662cc63fcef498a00fc08d9e9bbffe7b40) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-7356de531c6b5f5d247a1d1a340e199641612607aa53c74080d1f87c1d29117f) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-398ab0b68068225a40ff55c122883245ccb581e7f66f1534fbb67519359d5b79) |
| `annotations` | [annotations](resources--tcp_loadbalancer--reference--group-001.md#canonical-9d02c3572755948c8325d313dba332b91760c996e14aca0f2cc07cbb826eba1f) |
| `default_lb_with_sni` | [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-3788fb217091bb364117acb918953f23d8b90c1c58ced8a152d22577b6555028) |
| `description` | [description](resources--tcp_loadbalancer--reference--group-001.md#canonical-24c0ec8f441a6551b0dcb4174fa9b2be7817d1695177b12ef4bcc6991a3099cb) |
| `disable` | [disable](resources--tcp_loadbalancer--reference--group-001.md#canonical-e5fe0854d330b6b5e4f43a8c29b75aa474d59fb6cdd7d85e298e029b3c607736) |
| `dns_volterra_managed` | [dns_volterra_managed](resources--tcp_loadbalancer--reference--group-001.md#canonical-711ccd0e79b67ba4dd098ffb67ad6d5699f978a573bd18491bfe4ec1c7117d62) |
| `do_not_advertise` | [do_not_advertise](resources--tcp_loadbalancer--reference--group-002.md#canonical-de7e64a189c21e9f3f46ce0121947d119ee5bee16f16d6a10e89d0504f2a673d) |
| `do_not_retract_cluster` | [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-9e6628b6f34d48dd1b66dc924ac1d5ec91d3c3774d7fbcafb402b3f8dc162bef) |
| `domains` | [domains](resources--tcp_loadbalancer--reference--group-001.md#canonical-b52ac0c5fba86f8dbcbd034b2c05669ce13a6912380f6133ce201f7095b9d971) |
| `hash_policy_choice_least_active` | [hash_policy_choice_least_active](resources--tcp_loadbalancer--reference--group-002.md#canonical-982847a6e067eab3f677405e68402501c7e29da80d0971dc6b854489d213b593) |
| `hash_policy_choice_random` | [hash_policy_choice_random](resources--tcp_loadbalancer--reference--group-002.md#canonical-38521a8224ad394cfe8cd652d3035a77824b9a41bc3a4efbf14ab6614d1dac6d) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](resources--tcp_loadbalancer--reference--group-002.md#canonical-fae47bac70c13a3fa8e46c01d2322e558fc2fde428219f768ed6920169f060a6) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--reference--group-002.md#canonical-d12c6f4812985b415deec0717b71a02ce26d8da7e057db994ed37e035a4594ea) |
| `id` | [id](resources--tcp_loadbalancer--reference--group-001.md#canonical-ad9d4e0662a636c307a9befacdcbff8dd81740594f1d0485f97fb10f87e0c8f6) |
| `idle_timeout` | [idle_timeout](resources--tcp_loadbalancer--reference--group-001.md#canonical-1d5a7b535ec4b033e7a05967b300c1069364f692c09e84eb94a17d30e7d31cf6) |
| `labels` | [labels](resources--tcp_loadbalancer--reference--group-001.md#canonical-3e56b1d1c7a072840fbe2c104be165dad8f4834ebfa870ffda216e1b64d36926) |
| `listen_port` | [listen_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-134bf295899c5393ce4c482ce2f145a4dbe1f9dca34fa807ba0f0b65d77c4477) |
| `name` | [name](resources--tcp_loadbalancer--reference--group-001.md#canonical-e839c845549573ff2bc9a72d758f2bdef34f6c4a15c79cc525d497fb4f3ce3c0) |
| `namespace` | [namespace](resources--tcp_loadbalancer--reference--group-001.md#canonical-200156ad2e9e49e69320b386ae2bf49dff70227e464125f1d4f07eab03c8af7c) |
| `no_service_policies` | [no_service_policies](resources--tcp_loadbalancer--reference--group-002.md#canonical-12c893c9b154a54bbd2643e61d7c9e8fc00d6788f4f4fe74c14f0612bacefd95) |
| `no_sni` | [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-c7ca021715eb3eac0052ca372e3b473fbb49cd1d446d5f30a8b5d6c0a5e2d789) |
| `origin_pools_weights` | [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-f04ad50b5ad516b9df6f8c3cd2a82dfb4162bccfeb44f6910e644751ef3e536e) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-942bd378bb92030aa1cd9915b20be97862c55f3bdad9b1ddea6c1e596280437c) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-a984aaa851ba57bd65eb5d462deac346d8bbdc71a1690bf84a086c466256a46f) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-76e5f22b8b233e3257e5fa65f43eaed9a817d9486b26f64a8273e8b351506c82) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-3370414cdcc3b8c966737e9eaec3a22cfcd4707dd555b616ff248e2c3586b869) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](resources--tcp_loadbalancer--reference--group-002.md#canonical-d8bd47f63991028ac0e577f4dc4512696bc7ef1f53ad4b2cac2031492f71f2ac) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](resources--tcp_loadbalancer--reference--group-002.md#canonical-a5504463dea2b40665148d07dfbf6b838cc82d0a218f03b6a2e90550953f2900) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-ed7e56172412822cce07fc1efca6093483967a488ae7c8c1e1e83ffdac939ff1) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-75c72e7953856558b310d1dfcddf8e71a7a170d1b867517ca4f3ee13fbe3b83b) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-2437b6cacfa240a19aeb58c09d060fefc8e7f0e29014b1ff769019879e4f5422) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](resources--tcp_loadbalancer--reference--group-002.md#canonical-8b2f0ac94c1faca03b213d64d99636f61e6b4c55b4fc302cb8d0e77aed694cb0) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](resources--tcp_loadbalancer--reference--group-002.md#canonical-702ed77400d2d458554f69b652a904a7576f5cc7599fa01cbc01ed2e50e70cdc) |
| `port_ranges` | [port_ranges](resources--tcp_loadbalancer--reference--group-001.md#canonical-a44c3b980811a5a4d972ce9f74b0f6f5c339b9af307beeb15a5cbe6123e2d1e6) |
| `retract_cluster` | [retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-2b9a9a7b43fc5b3bb54f0086ca58dfff2563b0e53b5fac4e073e87d03972460f) |
| `service_policies_from_namespace` | [service_policies_from_namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-c1ad3fff48340b5a358cdb53dfdbe34e72962062e51d663e4320d1daef574a2c) |
| `sni` | [sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-99f5e75dcf3506323c7b2293c0a12580c5d75d61a0d59b6bc5a40868e1d1663a) |
| `tcp` | [tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-f8a3300666d29fc38bb80075835da84e389d0f20e01b4bfd63d8b3b38f6c0fa1) |
| `timeouts` | [timeouts](resources--tcp_loadbalancer--reference--group-002.md#canonical-2ddf1f8d291bf93ba2a0d30af0129e2812f85e6039aafd9c5fb7bd067fff75aa) |
| `timeouts.create` | [timeouts.create](resources--tcp_loadbalancer--reference--group-002.md#canonical-6bbc654538118f3b9442b4420925d63ae60e555308b8d1e8f2999236feedd49e) |
| `timeouts.delete` | [timeouts.delete](resources--tcp_loadbalancer--reference--group-002.md#canonical-b1e05b0dd7f397839eda1921cd879fcc361bcba32112bcacaff2ad9b9e5ac993) |
| `timeouts.read` | [timeouts.read](resources--tcp_loadbalancer--reference--group-002.md#canonical-b63277a70989155692a5117d3eccf29a336d8a8141a54381253d73b84469202d) |
| `timeouts.update` | [timeouts.update](resources--tcp_loadbalancer--reference--group-002.md#canonical-fbf72386fc53c991b3b9fe9c2b0fa1deb2422471a3db9f77855ac737743bdcc4) |
| `tls_tcp` | [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-f5d11b5136db56e396813bcb2d1d10476e78227b0523aa6575c4c8c4a64cc8fb) |
| `tls_tcp.tls_cert_params` | [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--reference--group-002.md#canonical-839adb6c9acef80e1b4a4ae9b1fd31890a98e94682061014c982ad34abdca7ab) |
| `tls_tcp.tls_cert_params.certificates` | [tls_tcp.tls_cert_params.certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-4bf7c44f4821368c9bf9db15ac4591e64bc6b15685ccf63d895a097417ad77bf) |
| `tls_tcp.tls_cert_params.certificates.name` | [tls_tcp.tls_cert_params.certificates.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-065d444cda11e11777d8b0c74b0bfce39060f6056cc12a6a091ce8048de4c4af) |
| `tls_tcp.tls_cert_params.certificates.namespace` | [tls_tcp.tls_cert_params.certificates.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-e119b8645677106a5f8706c4278fe3f3f331cf1b38c667246ff70793776a8877) |
| `tls_tcp.tls_cert_params.certificates.tenant` | [tls_tcp.tls_cert_params.certificates.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-e954d9bc1a092d2cb105bce4ffb42ad8c46f08fcb84a188ca0f782623d2225c4) |
| `tls_tcp.tls_cert_params.no_mtls` | [tls_tcp.tls_cert_params.no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-2d3bdbf3e98bb9d69ac82de23ada90ed2c82656feab7063d25cda4b6817fe5be) |
| `tls_tcp.tls_cert_params.tls_config` | [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-30ba8f8e6931cac9d549246cf322a7be2950e02978374eb7816a600afeb22f09) |
| `tls_tcp.tls_cert_params.tls_config.custom_security` | [tls_tcp.tls_cert_params.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-97efd0877528b98ff62a6dce1a9c9d997394c4587c0b3e2c0efd85c2bbda6469) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--reference--group-002.md#canonical-d958ecdd97ffc497b20aa15b791b99dab14168ca7203dc9109cb4ab14b9b8e1e) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.max_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.max_version](resources--tcp_loadbalancer--reference--group-002.md#canonical-a73e9cce9fbaf9a80461468d5dfac06d4e0c63006188843392d5deae76112830) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.min_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.min_version](resources--tcp_loadbalancer--reference--group-002.md#canonical-59af86cf863b8a767395402ffeb09e4ca231596c37f04d12297f363b1256a33e) |
| `tls_tcp.tls_cert_params.tls_config.default_security` | [tls_tcp.tls_cert_params.tls_config.default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-76445d70996e5f723a036b8c79095576a550eb4e0c9c85068ae5809cf7302c06) |
| `tls_tcp.tls_cert_params.tls_config.low_security` | [tls_tcp.tls_cert_params.tls_config.low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-7c537e85348f88cdd9d2368711b4dfd2e1a44e59b592414380af4d0add268ef8) |
| `tls_tcp.tls_cert_params.tls_config.medium_security` | [tls_tcp.tls_cert_params.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-a62bf1ba751fa8a13eb16dcd085d47c9e04fc8ca4dd711b1843267fed018f980) |
| `tls_tcp.tls_cert_params.use_mtls` | [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-d564850dc8103f3a5a2e16b1fa8495a26b5e284453498560ca6c4c4167472ace) |
| `tls_tcp.tls_cert_params.use_mtls.client_certificate_optional` | [tls_tcp.tls_cert_params.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--reference--group-002.md#canonical-c7e25aae57d84d4287f7acceda456ba799fdcf08b7b9df696d84d89480f20df4) |
| `tls_tcp.tls_cert_params.use_mtls.crl` | [tls_tcp.tls_cert_params.use_mtls.crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-2102e226a2d5162573054dda7b5c5ff20bc00f5f91ade2ac81e4278e5578a47b) |
| `tls_tcp.tls_cert_params.use_mtls.crl.name` | [tls_tcp.tls_cert_params.use_mtls.crl.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-1aa399ffce5889d600757cb51cde34dd0dd8a5b1b78ba73b95fc0fce9faefe15) |
| `tls_tcp.tls_cert_params.use_mtls.crl.namespace` | [tls_tcp.tls_cert_params.use_mtls.crl.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-ea960781cda343ce74d0e3f8f0a23a2153e5a12f4b0623428e0f9362abab779a) |
| `tls_tcp.tls_cert_params.use_mtls.crl.tenant` | [tls_tcp.tls_cert_params.use_mtls.crl.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-e241de93108b158dafd74cb1fbad1fa5c8d30c0fa55fddd3b5db058ff2ccfc02) |
| `tls_tcp.tls_cert_params.use_mtls.no_crl` | [tls_tcp.tls_cert_params.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-f91a2cbab02af12b020b84d0a19b0f669dbb2bccb5e755d650ad7de0a35b487f) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-81662e01a2b658548318786c28df285fbc043aa0467f119ebebfbf3272370026) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.name` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-0c0fe0b1cf9d367ea2439cc8ce05f41378df1e4353bd7757989481353da56c4f) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-cfceeff7da90d114521d0bffe90ca9ee5fbc422a5169f49b38800149776be32a) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-1f771d732064a50e77c8c6cf80ff8c7b8b302e2e43c913fe96cb0546007c9aa7) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca_url` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--reference--group-002.md#canonical-7d50d6041f589aaff46e34535832fc3c926e94278e4e9d7bc9861839c20d1427) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_disabled` | [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-520e0626013ad608c41aedf492771ce3949f9e4213407b8e461bc798c6546b2c) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-817f51e33de2a08d48b000736ce01f7a0f77bfbe94c4b1f33f2d858ffff02d0a) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--reference--group-002.md#canonical-d427cd4bbc4532c4b4824041fc4b26196d877455265ba7eed62e559f2b8f6b35) |
| `tls_tcp.tls_parameters` | [tls_tcp.tls_parameters](resources--tcp_loadbalancer--reference--group-002.md#canonical-d4769c860887e9b2ae5d731f715a34ff5b5839a189b45decccd16f21cb5ef71d) |
| `tls_tcp.tls_parameters.no_mtls` | [tls_tcp.tls_parameters.no_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-51e63514831425ccde074d2f430f29f8315ee7a2512e9ab1ea172009876d9efb) |
| `tls_tcp.tls_parameters.tls_certificates` | [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--reference--group-002.md#canonical-62b666102cdee1a411a3fb10df10554fd099d0dc8e559389e7607e206bae9ec6) |
| `tls_tcp.tls_parameters.tls_certificates.certificate_url` | [tls_tcp.tls_parameters.tls_certificates.certificate_url](resources--tcp_loadbalancer--reference--group-002.md#canonical-fbdb0eb311cd17f81736c112164a63cb509fb2272fb32b2d8d90eb6cc181c5a6) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](resources--tcp_loadbalancer--reference--group-002.md#canonical-cb3848221be67cb4201786c8b13fa885e91479c4c0b43ee0d1aadd44100cfbaf) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--tcp_loadbalancer--reference--group-002.md#canonical-a6a186f99f370f3d15afeafeed3dded1928e5689408af507fdcd786a104f4025) |
| `tls_tcp.tls_parameters.tls_certificates.description_spec` | [tls_tcp.tls_parameters.tls_certificates.description_spec](resources--tcp_loadbalancer--reference--group-002.md#canonical-a7c113dcf28f29d73afbd90621e603a4c8ff43401f1722776f0a4e16d7ecf6bc) |
| `tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling` | [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--tcp_loadbalancer--reference--group-002.md#canonical-cbd45497c9ae17d67b95902e3393b833bc7073c4c238b40126ef7b6e99b9c6bf) |
| `tls_tcp.tls_parameters.tls_certificates.private_key` | [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--reference--group-002.md#canonical-5b1748d7c91697be4f4f1d574040d2e87f587b86005a97a7fa3f34edabb5450e) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-fec46dcc08b92908c0d55fea9f584a98ea1b138234b91b82fbb1d687bdbb139e) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--tcp_loadbalancer--reference--group-002.md#canonical-8692c4721e8eb2ae682cad958e59eb3cb1a8fe153c237539200c3837059c05fd) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](resources--tcp_loadbalancer--reference--group-002.md#canonical-f3c6f480b6ecaa2e2bb1fb47918457f5d7ab318a07ed3b59d9b3765dc5c33a1e) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--tcp_loadbalancer--reference--group-002.md#canonical-4bf6be6b098efb4fd44eb8c833a644bb6e8dda1c9702230b481c75d08d2b92b9) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--tcp_loadbalancer--reference--group-002.md#canonical-59f990737c34c0febe2e09685755198517f82d62b68c2112c02c044b5856f203) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](resources--tcp_loadbalancer--reference--group-002.md#canonical-12d71bbe2f6e433ca300a70f8e8a48bc3d488b93f77d536123334f99e41f9aee) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url](resources--tcp_loadbalancer--reference--group-002.md#canonical-1c891faae28b14d8261107515821358dcf87043983c3147852720dc5ddaef4b4) |
| `tls_tcp.tls_parameters.tls_certificates.use_system_defaults` | [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](resources--tcp_loadbalancer--reference--group-002.md#canonical-f9f968bd3ea088f4333ea94dd5381109caf22f242cca206d63129879f8a3add9) |
| `tls_tcp.tls_parameters.tls_config` | [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--reference--group-002.md#canonical-26c0ca004c48c60eec18c10c0026fb93281961932e9c86faae2490720608f07a) |
| `tls_tcp.tls_parameters.tls_config.custom_security` | [tls_tcp.tls_parameters.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-918ec4c30acce8ad1daeb546ef1b64e7412c93c19028d16963481ef971401a19) |
| `tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--reference--group-002.md#canonical-9856c3a6f10498084b7ef3a654ce1505892aca615d22031fabb99a3793d68697) |
| `tls_tcp.tls_parameters.tls_config.custom_security.max_version` | [tls_tcp.tls_parameters.tls_config.custom_security.max_version](resources--tcp_loadbalancer--reference--group-002.md#canonical-4dfdec101f6e26c20047907f8d69d0f38aa41a202fc2f039beefa79477264ff8) |
| `tls_tcp.tls_parameters.tls_config.custom_security.min_version` | [tls_tcp.tls_parameters.tls_config.custom_security.min_version](resources--tcp_loadbalancer--reference--group-002.md#canonical-3e1c44b62f035948bd02a5e118dc280e676b036e43a8090bdabf7b8934ced109) |
| `tls_tcp.tls_parameters.tls_config.default_security` | [tls_tcp.tls_parameters.tls_config.default_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-90405f4e788751e268de7fc91e700f3dc84b6ea8a1acde55d737e6e3f9c6d93e) |
| `tls_tcp.tls_parameters.tls_config.low_security` | [tls_tcp.tls_parameters.tls_config.low_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-69fa96b5dc15cc29622e1bacc77a628f8af0ccf2acd65f48c7eec38299b8fd71) |
| `tls_tcp.tls_parameters.tls_config.medium_security` | [tls_tcp.tls_parameters.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-002.md#canonical-4834a793215722a3dfb4701b5a78a353ac1e18034db90ce141998c9839f80b82) |
| `tls_tcp.tls_parameters.use_mtls` | [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--reference--group-002.md#canonical-6ea72b670d9342d72553586b580b69d885e70506e030706a904fba9ef26b618c) |
| `tls_tcp.tls_parameters.use_mtls.client_certificate_optional` | [tls_tcp.tls_parameters.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--reference--group-002.md#canonical-e9547855e72e4233dd339b81e348b0d312044ca4864ed6bd63868b50e995e61b) |
| `tls_tcp.tls_parameters.use_mtls.crl` | [tls_tcp.tls_parameters.use_mtls.crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-ad0b02b934e081f9c7b41cbc2a2a12037fd89a881b0a378de92d2b13bb936b30) |
| `tls_tcp.tls_parameters.use_mtls.crl.name` | [tls_tcp.tls_parameters.use_mtls.crl.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-2c1b79330f12b7c730816f5e396a9acfdec46325f28c7477ede3abecc03839eb) |
| `tls_tcp.tls_parameters.use_mtls.crl.namespace` | [tls_tcp.tls_parameters.use_mtls.crl.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-17a7128c69ab97b1046d94e23026dce626e298196b543634f9d61580b789a6ab) |
| `tls_tcp.tls_parameters.use_mtls.crl.tenant` | [tls_tcp.tls_parameters.use_mtls.crl.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-f94cbcba5cc93cb0eec999ade7e2134fee9d053391ccadbf935b36ae68e26680) |
| `tls_tcp.tls_parameters.use_mtls.no_crl` | [tls_tcp.tls_parameters.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-002.md#canonical-dc67fa54b332672f119c4aeee658ee0ef43fff943cefbe0b5a02a0a3ac5aba81) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca` | [tls_tcp.tls_parameters.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-002.md#canonical-2b89dc2634daa717de787e5f57118572562eb480d78e5036703ee740b3472935) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.name` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--reference--group-002.md#canonical-a698b662b72fce792de59525dbc4b07adca20b49e6dfad29f78a44516e366aed) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-ac731bacc11592d2b21532cb83666dcb5448c48c8337baaf180aa1b2683fc7d7) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--reference--group-002.md#canonical-1c6cd973263adb0ded51346501d07a09ceaa1c0181ddec1ba00cde00e9af5e3f) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca_url` | [tls_tcp.tls_parameters.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--reference--group-002.md#canonical-8396a709121ca368355627f16b026c3c183d967492e49dd2704b6ab9b992273b) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_disabled` | [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-002.md#canonical-71c57554afe54c60e15deff1db2e063f9c33ce6cdc467771f83e4050b22ddf2b) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options` | [tls_tcp.tls_parameters.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-002.md#canonical-6e143add6dd848c238c2a03f88fe0226071145f139a1bde7fdf2a07519836a12) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--reference--group-002.md#canonical-6d113ed51f2aa65dcde5738e97f78ef10445efca76b5c0907d830ef4c1f8f7c9) |
| `tls_tcp_auto_cert` | [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-490a6432751fa2d70f483cf952ae0e4fd5499dee0876f285d95600aa388da7e3) |
| `tls_tcp_auto_cert.no_mtls` | [tls_tcp_auto_cert.no_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-b76060dd354205d41b3cf6e0c381ec28490050e9e0011f219419dc62de7e8b07) |
| `tls_tcp_auto_cert.tls_config` | [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--reference--group-003.md#canonical-f9d219f6524fac96a79a280ef2b5143c27a7f5654bd547db08dd8dd670722b12) |
| `tls_tcp_auto_cert.tls_config.custom_security` | [tls_tcp_auto_cert.tls_config.custom_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-dae249d8ded308f90b8df7a98679aa4e5e244aacd4ffa94360524be0e25d493f) |
| `tls_tcp_auto_cert.tls_config.custom_security.cipher_suites` | [tls_tcp_auto_cert.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--reference--group-003.md#canonical-d4b24d2945ef03ba1701cd848ab5b5cb4ac8b6a9d5c1f3e9d1e5fe90efa64960) |
| `tls_tcp_auto_cert.tls_config.custom_security.max_version` | [tls_tcp_auto_cert.tls_config.custom_security.max_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-f2b238884819c50c885f0f18c697c98ea64ee7ed6ca247bff12ee0bb4b09eec6) |
| `tls_tcp_auto_cert.tls_config.custom_security.min_version` | [tls_tcp_auto_cert.tls_config.custom_security.min_version](resources--tcp_loadbalancer--reference--group-003.md#canonical-bd1bd4bcaad7d6a1d24f63113817fa0e1379c208644c3b4c307365d7da5ba6b0) |
| `tls_tcp_auto_cert.tls_config.default_security` | [tls_tcp_auto_cert.tls_config.default_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-841321beacee51a236b218d2f005ccba7cd5d24217ca05579a9eee54a560dbc6) |
| `tls_tcp_auto_cert.tls_config.low_security` | [tls_tcp_auto_cert.tls_config.low_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-06f71cd71a1cd624d57fc08a6bb9c128274990f6ffcf5e2dbce0c27b5e18396e) |
| `tls_tcp_auto_cert.tls_config.medium_security` | [tls_tcp_auto_cert.tls_config.medium_security](resources--tcp_loadbalancer--reference--group-003.md#canonical-a3f5b47de0a00b09163da0873fba92dd4319be8641ea9a2180017413953fb436) |
| `tls_tcp_auto_cert.use_mtls` | [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--reference--group-003.md#canonical-9a6f4d817ec8e7d816a0a0573e25a30a3db14016fc4bd4564b3b034a1f3ff74e) |
| `tls_tcp_auto_cert.use_mtls.client_certificate_optional` | [tls_tcp_auto_cert.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--reference--group-003.md#canonical-49fa2a4d05e8f50300ea151eb68e1a69744b9455bba66f79c09811cbf38d3f3e) |
| `tls_tcp_auto_cert.use_mtls.crl` | [tls_tcp_auto_cert.use_mtls.crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-484eab373ada77fa19c7bb4d3231256a300a3618044f45a2dc46807ac0a27b37) |
| `tls_tcp_auto_cert.use_mtls.crl.name` | [tls_tcp_auto_cert.use_mtls.crl.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-8a551dc437eaf6477d062c186b53c7164c7674c69e513261c7647e210413ea33) |
| `tls_tcp_auto_cert.use_mtls.crl.namespace` | [tls_tcp_auto_cert.use_mtls.crl.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-9b9dd3d711ceb02556160599d4faae50e7281bad706998ac06d6c90fffe7aad5) |
| `tls_tcp_auto_cert.use_mtls.crl.tenant` | [tls_tcp_auto_cert.use_mtls.crl.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-cbc06ec5d6997216e3888f99e2432b8dc0ec6f776508db4f7b2f5ac0c1a7a630) |
| `tls_tcp_auto_cert.use_mtls.no_crl` | [tls_tcp_auto_cert.use_mtls.no_crl](resources--tcp_loadbalancer--reference--group-003.md#canonical-7d8e5434b81927930c100af8f0fd518584d1dc0044edb3d6449108d1b4e2baf2) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca` | [tls_tcp_auto_cert.use_mtls.trusted_ca](resources--tcp_loadbalancer--reference--group-003.md#canonical-34c229689c4c84321291ae7c8cd294a4e5575d5ef34e50ef7b8b247b53186745) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.name` | [tls_tcp_auto_cert.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--reference--group-003.md#canonical-1159fd3b503a0d5f85558686183d2cc853cbcb8d3b47d7db24eb2c7d175c7caa) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.namespace` | [tls_tcp_auto_cert.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--reference--group-003.md#canonical-d0fd23e36c6885b1c42e7cc058d0332cdd196469ff2f174711a14f9c1cd838f8) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.tenant` | [tls_tcp_auto_cert.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--reference--group-003.md#canonical-de0afcf1d217c9a6617b28a1273ed47e1511808d3c0cd109e750547345351231) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca_url` | [tls_tcp_auto_cert.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--reference--group-003.md#canonical-71bf92f1e9c5be87d8560415e77597a87a742704497c877daa92c73577cac3d8) |
| `tls_tcp_auto_cert.use_mtls.xfcc_disabled` | [tls_tcp_auto_cert.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--reference--group-003.md#canonical-074ec61f8e0c8fd8574c8335f022a9ff62b3f538641f246b7bea55c4f9e13271) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options` | [tls_tcp_auto_cert.use_mtls.xfcc_options](resources--tcp_loadbalancer--reference--group-003.md#canonical-a739c6c05f47696d1c43ab48dffb9464a25d74679fa7bbf0db8cd85f7ae33512) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--reference--group-003.md#canonical-5acb5161ef09d5e61528baf5c2050ff49e5741b721512e6cbe4ef24d6d96e5de) |

<a id="canonical-781d582b74d7b183fb4d5f09f8591dcdac543fceebef4eb288c5283ddb6510a9"></a>

## Next pages — Property reference / 89d73e3380a2 / 17

- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-9bf84c084bb7532e7035a4454e2e392860fc4eb33b7dba834871fadf5674c8d2)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-5145a903bc57dcc49c9709e298e99c068fe602cad3de3f2ab20093ee69049729)
- [advertise_on_public_default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-ff42e920406b0375c6cbaac8ce29eedff4194fc6bf1fc0585caa1787584c2174)
- [default_lb_with_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-6fc17a77d48f9cf85e426e65ef541a24a263cb7378eb08dbdebeaddb4c6c5404)
- [do_not_advertise](resources--tcp_loadbalancer--reference--group-002.md#canonical-ba9e95942d760250bb35617997afdb3d548b484a3b5ac545e364fd505c61a698)
- [do_not_retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-ba5e7e3c123f000e5c21dcf85c3f1ab41075becc703e6fcc8f87971d9c0fce07)
- [hash_policy_choice_least_active](resources--tcp_loadbalancer--reference--group-002.md#canonical-edd915446ceae49d7455932f6f248f30860d874aec8e428a7dd297607454e0bc)
- [hash_policy_choice_random](resources--tcp_loadbalancer--reference--group-002.md#canonical-9f8669f3e8de1af30d714b546776d46ef9bc8683f210f0b322894e386a825f0e)
- [hash_policy_choice_round_robin](resources--tcp_loadbalancer--reference--group-002.md#canonical-85cdf1528965a6285ff2c75a0267cb449348c1ca88921d060042419de75f1ace)
- [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--reference--group-002.md#canonical-dbedf978ab25250f8b6bdd8d76481f4d8b6a7e6f331b982a6fc699ddc46b8ac6)
- [no_service_policies](resources--tcp_loadbalancer--reference--group-002.md#canonical-25ccd69fc2452f16f34ad58d3b9af414cbf6a2ca1af6864921afc4c6286aca01)
- [no_sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-3e6782861d51d0a9e5a0f8e8ffa45db8737335ce850aafcfc19a13424444586b)
- [origin_pools_weights](resources--tcp_loadbalancer--reference--group-002.md#canonical-9010378f7d918df4260eb3f27869e01272844da0ffa7dc25fe2fa5c5121b4bda)
- [retract_cluster](resources--tcp_loadbalancer--reference--group-002.md#canonical-a3f2eb692d42a4518e81b4d2ca346c903a6e9aedbc28daf1e229ac85d2217fbf)
- [service_policies_from_namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-7202f6813ce28cd3211969f35f4a551cbf7c6921c78f64f4a67273125a5d0ceb)
- [sni](resources--tcp_loadbalancer--reference--group-002.md#canonical-65d837fc6b970ea897c23eaec474f1cabe159cddf3ae011ef00b692480d12e3a)
- [tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-a8b8d21b3fa7b0073cd75420b779305e70da290853df37c205c279e25e1912e5)
- [timeouts](resources--tcp_loadbalancer--reference--group-002.md#canonical-6b17fb1ab324a8664c2799ba750302145ca7f5afebaf304cad2896a517809a45)
- [tls_tcp](resources--tcp_loadbalancer--reference--group-002.md#canonical-e23c296f716a96184f1b94a4b91de1cf35da41a48b56dc6f5e4a941256b9a3d8)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--reference--group-002.md#canonical-4803fa7626ede63c16c722f64429587626623e72c537a6380a35e33ec1b289bd)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-9bf84c084bb7532e7035a4454e2e392860fc4eb33b7dba834871fadf5674c8d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d67470b89c9b7f211a491473a241019961a985120cc581092d5e75ef4427d21"></a>

## active_service_policies — active_service_policies / d22cadcb9c6b / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- active_service_policies

<a id="canonical-659f5423c72f84f3eb0fa8f83d9c428401ff14240ac8f0c2b27e6d2cd6003907"></a>

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

- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-659f5423c72f84f3eb0fa8f83d9c428401ff14240ac8f0c2b27e6d2cd6003907)
- [no_service_policies](resources--tcp_loadbalancer--reference--group-002.md#canonical-12c893c9b154a54bbd2643e61d7c9e8fc00d6788f4f4fe74c14f0612bacefd95)
- [service_policies_from_namespace](resources--tcp_loadbalancer--reference--group-002.md#canonical-c1ad3fff48340b5a358cdb53dfdbe34e72962062e51d663e4320d1daef574a2c)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-52f6a246c6ea319aac49105ef95bcc4c3ef6d712505c77cf274b9c3a8479cf39"></a>

## Direct properties — active_service_policies / d22cadcb9c6b / 3

- [policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-dd069de925f4b281140e3b057605a9df954ee3f618fd127ba3448baec46a8041): complete subsection reference.

<a id="canonical-d6ffc79f5428de20da9151f914aedeb5b59ce61cdc862a491daed6f262fdbd0c"></a>

## Next pages — active_service_policies / d22cadcb9c6b / 4

- [active_service_policies.policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-dd069de925f4b281140e3b057605a9df954ee3f618fd127ba3448baec46a8041)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-dd069de925f4b281140e3b057605a9df954ee3f618fd127ba3448baec46a8041"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46847bcec370c241c1f15bd7681c9411ad8524857b6b1b16ac96173e86961237"></a>

## active_service_policies.policies — active_service_policies.policies / 431e80d38cd1 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-9bf84c084bb7532e7035a4454e2e392860fc4eb33b7dba834871fadf5674c8d2)
- active_service_policies.policies

<a id="canonical-49a0456a0637522726e8d341b2ea2ac30b4769477b1f735a494809c8cebdfdca"></a>

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

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5559e219300253c354c890bb89c04b3eec83160395452bf75dd1ea38f2cf330"></a>

## Direct properties — active_service_policies.policies / 431e80d38cd1 / 3

<a id="canonical-f24ba9ab28187514d519f5fbe2bd9557f74e6ad7b407b6e1f533551e5d233dc8"></a>

<a id="canonical-12d9da6d2952730c31aa5357a5fca204ed40ca7a88af94e2e34edea685a2d7dd"></a>

## name property — active_service_policies.policies / 431e80d38cd1 / 4

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

<a id="canonical-43bca3cd4cea136acf7aaf1bafae271c7118ac4433e97c3a605a4104e6feb094"></a>

<a id="canonical-cdf4b1e84c941364474bb5c825d0867fc132ae5f96745ac54b57f022134bb33d"></a>

## namespace property — active_service_policies.policies / 431e80d38cd1 / 5

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

<a id="canonical-95d3b8da8ca3707890a9bae5f84ac49276ddb0bacf93ac4206a67c9dc01b18b0"></a>

<a id="canonical-ca003877e7b2a2f517d7f1fb3ed84e441960e435b1bc442192abbd6041f54942"></a>

## tenant property — active_service_policies.policies / 431e80d38cd1 / 6

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

<a id="canonical-ff103af532f4de3c5cf7829071232c6f412ef01dc8095ba6011f05fd238226b5"></a>

## Next pages — active_service_policies.policies / 431e80d38cd1 / 7

- [active_service_policies](resources--tcp_loadbalancer--reference--group-001.md#canonical-9bf84c084bb7532e7035a4454e2e392860fc4eb33b7dba834871fadf5674c8d2)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52fb3001ea23ba6fcb8f497086307582cbfb99a8ae432903d4d1de430b4d39df"></a>

## advertise_custom — advertise_custom / 654ccd78c575 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- advertise_custom

<a id="canonical-71eca36c430c121b5e87e6eeb5af43b65e2442766d24f4415d1e766593b09da3"></a>

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

- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-71eca36c430c121b5e87e6eeb5af43b65e2442766d24f4415d1e766593b09da3)
- [advertise_on_public](resources--tcp_loadbalancer--reference--group-002.md#canonical-32ddddd3b88cd727dfc738e4a881d56c2f176fb82c812cd6b5af1a6bceba73a5)
- [advertise_on_public_default_vip](resources--tcp_loadbalancer--reference--group-002.md#canonical-398ab0b68068225a40ff55c122883245ccb581e7f66f1534fbb67519359d5b79)
- [do_not_advertise](resources--tcp_loadbalancer--reference--group-002.md#canonical-de7e64a189c21e9f3f46ce0121947d119ee5bee16f16d6a10e89d0504f2a673d)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2a3019f3dcbd4ee995519a57548d37988f6d6aa72840dd12da57a67caa1dc8e2"></a>

## Direct properties — advertise_custom / 654ccd78c575 / 3

- [advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe): complete subsection reference.

<a id="canonical-04217ecf42058348223606780641f01f3ddea239e65850d653a603cdf0fd211e"></a>

## Next pages — advertise_custom / 654ccd78c575 / 4

- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa116c995e1e5abae6d76f7e205c965c1dbf87015643e17f911db4710e758ccb"></a>

## advertise_custom.advertise_where — advertise_custom.advertise_where / ccd91b185acd / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- advertise_custom.advertise_where

<a id="canonical-dd0cd5cb2812c4b73052a58cf05ee7affd4e34c5380a1e9a662f9240d5bfaa2b"></a>

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
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-c92f31843d69fea066b3ca4263c6fac4d696195d8ad944c8cd744d8a5efdab7f"></a>

## Direct properties — advertise_custom.advertise_where / ccd91b185acd / 3

- [advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-c987eff1694d7b31aa71ac1b738f7376971911e8ff891f96d65d2a1bc34e3d01): complete subsection reference.

- [advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-f73da15706e40629ccc6b9923ad4ec41768cef2e6f385b93b218c6a608ee8a1c): complete subsection reference.

- [advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-081384be8ca54f88320cd5e565b94cbb9fd3089eddd6e6e5d864de646103026c): complete subsection reference.

<a id="canonical-827eb5557ec6f24b8be165381e45de306aeaadd0ea24d583da5da9b831d090a0"></a>

<a id="canonical-138991688356ddf3e86fdfc2582a3a6a9216f96cdcf1c58629f264a7ea2d6481"></a>

## port property — advertise_custom.advertise_where / ccd91b185acd / 4

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

<a id="canonical-88bc97fa6bb7a6f707966fd5da3912eb1a59279445011821d1728290daa3168b"></a>

<a id="canonical-2c5c818e372a4c6f832c6b5319d6fb9f77e2b4f24e3df0c7c67c0d57f28cffb0"></a>

## port_ranges property — advertise_custom.advertise_where / ccd91b185acd / 5

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

- [site](resources--tcp_loadbalancer--reference--group-001.md#canonical-26a2dc79721bfc784f8d9bb8b6355e3945dab028b1943d2b3aae9c375500f2f9): complete subsection reference.

- [use_default_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-518c5cf7e99022e4323c937d996b2facfb322e40c1f6eeb8fca092974cd64596): complete subsection reference.

- [virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348): complete subsection reference.

- [virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-5ab4c9d572a5cb0d3dd7c74517ee809160f708db54f37990efdfe466b4e42447): complete subsection reference.

- [virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-5987f0c1e54c132e58207bc26199657ab73a1bb0b77bf939c8683022a3a10c71): complete subsection reference.

- [vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-b4392d7e78138a30c0757783cdc3c073c5da6748af1bf6264da06f58d293c38a): complete subsection reference.

<a id="canonical-a876343001bb5ef14a1ed02e23cc0c586cef10771a092fd1e4b6fb89a57412db"></a>

## Next pages — advertise_custom.advertise_where / ccd91b185acd / 6

- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-c987eff1694d7b31aa71ac1b738f7376971911e8ff891f96d65d2a1bc34e3d01)
- [advertise_custom.advertise_where.advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-f73da15706e40629ccc6b9923ad4ec41768cef2e6f385b93b218c6a608ee8a1c)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-081384be8ca54f88320cd5e565b94cbb9fd3089eddd6e6e5d864de646103026c)
- [advertise_custom.advertise_where.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-26a2dc79721bfc784f8d9bb8b6355e3945dab028b1943d2b3aae9c375500f2f9)
- [advertise_custom.advertise_where.use_default_port](resources--tcp_loadbalancer--reference--group-001.md#canonical-518c5cf7e99022e4323c937d996b2facfb322e40c1f6eeb8fca092974cd64596)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348)
- [advertise_custom.advertise_where.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-5ab4c9d572a5cb0d3dd7c74517ee809160f708db54f37990efdfe466b4e42447)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-5987f0c1e54c132e58207bc26199657ab73a1bb0b77bf939c8683022a3a10c71)
- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-b4392d7e78138a30c0757783cdc3c073c5da6748af1bf6264da06f58d293c38a)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-c987eff1694d7b31aa71ac1b738f7376971911e8ff891f96d65d2a1bc34e3d01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e5b20f607cbe7d4cc0998b90dd84e2e0375508c2701045e240ee92fce22d6a"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_custom.advertise_where.advertise_dualstack_on_public / 60c36b52c379 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-291dd98f0763be52fba0d7672a720aa937b4a80a71762bd88e9d317164976490"></a>

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

<a id="canonical-5c8ac850ce27b48cac0af374c5ce91ad71921ecccd3a27cac6d922f474cb7e33"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public / 60c36b52c379 / 3

- [public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-b2db6ce6dfc70527a69692c2a6583108baf265b80e1f3e95524d93b300614877): complete subsection reference.

<a id="canonical-1e80cc1e8215c906086fdb2ada433f62d7cef2171227c491512e0f67b5203450"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public / 60c36b52c379 / 4

- [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-b2db6ce6dfc70527a69692c2a6583108baf265b80e1f3e95524d93b300614877)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-b2db6ce6dfc70527a69692c2a6583108baf265b80e1f3e95524d93b300614877"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-621014ebb24eea5dc5883abe2838f807352b02fd4226d604f37e5b1030624adc"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / e70d80066337 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-c987eff1694d7b31aa71ac1b738f7376971911e8ff891f96d65d2a1bc34e3d01)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-fe06bff3ef0202b3b9edb4b60acc930049d36b3a3a5c65b77c4ce0807f7d4655"></a>

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

<a id="canonical-80947fde7be670aaf656f09c8860b3277e051a98a150ed4f8b0cdfbf1f5c6f2c"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / e70d80066337 / 3

<a id="canonical-29ea2ae15b43af37ed7f5d762925ab2e5ba722220517e90422f75ce2cfd629b4"></a>

<a id="canonical-da4b903ad633280aec442e8506ecef9616293c3b32e8a83b6702c7f79381c665"></a>

## name property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / e70d80066337 / 4

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

<a id="canonical-fb4bf958119068434eee1f2369f5fbaefa4d03f3ce4f2ccccf80df6803d00822"></a>

<a id="canonical-19b3017be1e8172143a963bfd00a1af35feb2772e31737d94495c5770a4b2ed7"></a>

## namespace property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / e70d80066337 / 5

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

<a id="canonical-3789988360e283d95e87a2128b44d2ee1b4bb15a19bcc22a40d3ca67b92eebc2"></a>

<a id="canonical-4dfcaabeb0607c73659b27a1f44501637f15558e086286419dba6c556451661c"></a>

## tenant property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / e70d80066337 / 6

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

<a id="canonical-844409a402a8d6f8702636507ae0f444e54e56cc4ef7eea23790f1a2b2b38ccb"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / e70d80066337 / 7

- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-c987eff1694d7b31aa71ac1b738f7376971911e8ff891f96d65d2a1bc34e3d01)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-f73da15706e40629ccc6b9923ad4ec41768cef2e6f385b93b218c6a608ee8a1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-debb5e0c2c29bf93a9f0d573df86b192d11bf31fe2398e27b92e0fffe0a11b40"></a>

## advertise_custom.advertise_where.advertise_on_public — advertise_custom.advertise_where.advertise_on_public / 02e61fb113ea / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-6180b2610a3212c313e0f77d00c5282dc128e230fa9a07ce79581bc43bb755c5"></a>

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

<a id="canonical-0978f7eb116876c7460bfc88f1421aafdf91ebb217129db2ae2fe75b87b64c4a"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public / 02e61fb113ea / 3

- [public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-197949fbd9f214aeb0735468d939893d9ad8e3aefd96ea6f674a999fe329c427): complete subsection reference.

<a id="canonical-ee684cb4b6163b62a75d81b197af82fd3e8e98098432321ab8f9507d801341ea"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public / 02e61fb113ea / 4

- [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-197949fbd9f214aeb0735468d939893d9ad8e3aefd96ea6f674a999fe329c427)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-197949fbd9f214aeb0735468d939893d9ad8e3aefd96ea6f674a999fe329c427"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98a9a3881d0a163be3e845a42b92a29bac055016fae02da8485677a18a14bef2"></a>

## advertise_custom.advertise_where.advertise_on_public.public_ip — advertise_custom.advertise_where.advertise_on_public.public_ip / 81ece78c8186 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-f73da15706e40629ccc6b9923ad4ec41768cef2e6f385b93b218c6a608ee8a1c)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-3ebcf6f612643a46eaa5921a0a224d6fee882cc49c51f0c3cc313803ad6557b6"></a>

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

<a id="canonical-5afde28e45587dca97be20b0ca47bbc72f8cf8d588c7ceab394a896c5b73f6b0"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public.public_ip / 81ece78c8186 / 3

<a id="canonical-f0de4b313ead309a358812482220f88a448837ca438d969494457ed3bbe78b43"></a>

<a id="canonical-7e24c36832602b3f62b91c4d37ff73e8558c2da5995d6589ea2111e9db7ef610"></a>

## name property — advertise_custom.advertise_where.advertise_on_public.public_ip / 81ece78c8186 / 4

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

<a id="canonical-ad9c0888c2a3076ec575b1f16248694a7edb5beb6feacbd4bb2cf218dba511ce"></a>

<a id="canonical-5db7a7cf6cf8781c0ce4d80972fa687c0886fd5fd95362aa0f59b28f68cb7b26"></a>

## namespace property — advertise_custom.advertise_where.advertise_on_public.public_ip / 81ece78c8186 / 5

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

<a id="canonical-63091ffd354cf54f2bb2e9a3bf183496e053f6439ebabe1d37ab9092e22a8f59"></a>

<a id="canonical-ec4c6bb88a6a4add77f1fd1f1eb62c7a995b268d8ec9c018c62db8700232eac4"></a>

## tenant property — advertise_custom.advertise_where.advertise_on_public.public_ip / 81ece78c8186 / 6

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

<a id="canonical-d851ec8842138e6a24a972149be575dcc33ffd31b83e9c06d8bb2e22f778c33d"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public.public_ip / 81ece78c8186 / 7

- [advertise_custom.advertise_where.advertise_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-f73da15706e40629ccc6b9923ad4ec41768cef2e6f385b93b218c6a608ee8a1c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-081384be8ca54f88320cd5e565b94cbb9fd3089eddd6e6e5d864de646103026c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0df29683657b3505becc863e493cce2c2db5ff4635987fe86fd65f3e5657b7c1"></a>

## advertise_custom.advertise_where.advertise_v6_on_public — advertise_custom.advertise_where.advertise_v6_on_public / 65a116d481d3 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-b3fd43f92eb4aca095d337a192b89f0ba077a13d1120a4425fb9d6046bf05d74"></a>

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

<a id="canonical-ecd022646c7fcaf4a6c7dbf3bfeaa903fec27c9cfa408ddbb7fb3909ada99db8"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public / 65a116d481d3 / 3

- [public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-076ea005928295f9a675ce6d539fabe866651a44b0267705405c071e455b4caa): complete subsection reference.

<a id="canonical-4724389b945450d2202df7b8081f5cfe0d86d41ab0dbf24af3e2455d7a8d7214"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public / 65a116d481d3 / 4

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--tcp_loadbalancer--reference--group-001.md#canonical-076ea005928295f9a675ce6d539fabe866651a44b0267705405c071e455b4caa)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-076ea005928295f9a675ce6d539fabe866651a44b0267705405c071e455b4caa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dacac801565de73c4aa0e075806a35844d95c984f32fc0c9e1b84a2be7dd9e04"></a>

## advertise_custom.advertise_where.advertise_v6_on_public.public_ip — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / e25f3158f561 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-081384be8ca54f88320cd5e565b94cbb9fd3089eddd6e6e5d864de646103026c)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-0619eefe4a5ad7da60383f6c575c6a9ca58fd23b00e6eff543a2634809658db4"></a>

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

<a id="canonical-83e263e9b279a6168dcf7756c6d78eb336e099a993be27e8d879008c6b6f3e79"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / e25f3158f561 / 3

<a id="canonical-777d6cf9f7c3a887ccb66126144705f29fcc9435bb34a3d58b5b75710828ae25"></a>

<a id="canonical-2e140a749bcc0b07c05479099feeb8e8eeb2e3ce128e52ccf84759d025faa390"></a>

## name property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / e25f3158f561 / 4

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

<a id="canonical-bf1a3749adfe6183a84f2fb24b6d5e259ff9c7f944d212d6642ba7b83d2d5761"></a>

<a id="canonical-6d8230ff402838c59cd9b370be026bc89a92c3ac305fdb1cfa979a7e88c29949"></a>

## namespace property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / e25f3158f561 / 5

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

<a id="canonical-14e1c8ddd3f34add0917ab506e7471c62b40156cd9e39838d03dc7f975e729dd"></a>

<a id="canonical-66ec2892b8caa9bfeda19c315419d6793e2a1e53b9df30d430336a635835469b"></a>

## tenant property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / e25f3158f561 / 6

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

<a id="canonical-284a858559ee5c99bcda89817beb6ed6d091a8b286da64663f69619466089356"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / e25f3158f561 / 7

- [advertise_custom.advertise_where.advertise_v6_on_public](resources--tcp_loadbalancer--reference--group-001.md#canonical-081384be8ca54f88320cd5e565b94cbb9fd3089eddd6e6e5d864de646103026c)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-26a2dc79721bfc784f8d9bb8b6355e3945dab028b1943d2b3aae9c375500f2f9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72cd38323e97cbe48dc6bee745bd11c5f874d400ab7937099982c4ddf23caca6"></a>

## advertise_custom.advertise_where.site — advertise_custom.advertise_where.site / 8e53aa1cd850 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.site

<a id="canonical-2f0a65afd5f3978dd9abf53ac9f0dc2e73a2ee1f05ef35b2ce459d6a54afb3c2"></a>

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

<a id="canonical-0dcc5da828b7f489d2e89fd818e59565ac3d107480e22586d5b2843ff12289ce"></a>

## Direct properties — advertise_custom.advertise_where.site / 8e53aa1cd850 / 3

<a id="canonical-57b308427e9ad8c7c4ccf3a0ef586c9cce17a6366a42df1a246a6c36b4338f4a"></a>

<a id="canonical-d4e7a1fc7468c836d60571ee63ff6b51e55b852d7e320639834ef03881b23748"></a>

## ip property — advertise_custom.advertise_where.site / 8e53aa1cd850 / 4

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

<a id="canonical-c55415264b949c80a59ab84353415f5580d3e2c8d95347e08be19a05acb76ad1"></a>

<a id="canonical-7bd9de79cd4bccd095940173813bd93a2dac40547c02af3d34c39f6809b4d3f3"></a>

## network property — advertise_custom.advertise_where.site / 8e53aa1cd850 / 5

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

- [site](resources--tcp_loadbalancer--reference--group-001.md#canonical-8098e76ccdacde34f92a0e7b16320a5794be07eb91b2fa3d95039a924dde3e5a): complete subsection reference.

<a id="canonical-7fb544c2c31d5b38cfe66ed2464990d865a31cd88bb6d77b4f13c72c36d92c3b"></a>

## Next pages — advertise_custom.advertise_where.site / 8e53aa1cd850 / 6

- [advertise_custom.advertise_where.site.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-8098e76ccdacde34f92a0e7b16320a5794be07eb91b2fa3d95039a924dde3e5a)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-8098e76ccdacde34f92a0e7b16320a5794be07eb91b2fa3d95039a924dde3e5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a44cbd0c48cfe493f6caabebd3ba49d47684a6444b2cfa5236c2b1fe14314e29"></a>

## advertise_custom.advertise_where.site.site — advertise_custom.advertise_where.site.site / df875fd49b3b / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-26a2dc79721bfc784f8d9bb8b6355e3945dab028b1943d2b3aae9c375500f2f9)
- advertise_custom.advertise_where.site.site

<a id="canonical-3fde104ab7367e6d1e251083ec2cd5c0d8a3b82c2d32233b575717497f3785e2"></a>

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

<a id="canonical-aebc81a9d57ff1370e39222d15b4f09494df882983eba0413a7728e9b3e90b5b"></a>

## Direct properties — advertise_custom.advertise_where.site.site / df875fd49b3b / 3

<a id="canonical-64bb6c84727a1f847696a0f8d0d55468b57b8e62b7be1b258a526ccd3ac2781f"></a>

<a id="canonical-0e5259327e8f5ef47160f7316075777ca7b142ec48c16594a8aae63f60233887"></a>

## name property — advertise_custom.advertise_where.site.site / df875fd49b3b / 4

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

<a id="canonical-8f68df4810419f2ce4094c99876ad6afadaddd3c43dd63f6709c80806b12ef98"></a>

<a id="canonical-ade656377c6da39ff0d8f65d09e92e90b08a46697c9dc2ce906a340304c072c6"></a>

## namespace property — advertise_custom.advertise_where.site.site / df875fd49b3b / 5

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

<a id="canonical-d75db1b737736af546a474d15a9783fec0ce84f041c5c162fb40162a92c9acac"></a>

<a id="canonical-fd1ed91addce47044c75d3189e30a3651902da506e32abfb3f877b220d6f3d22"></a>

## tenant property — advertise_custom.advertise_where.site.site / df875fd49b3b / 6

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

<a id="canonical-5b6646f36c58b29c1fb524b6f5973377be626e645ed45e8bc962f7c90acafc83"></a>

## Next pages — advertise_custom.advertise_where.site.site / df875fd49b3b / 7

- [advertise_custom.advertise_where.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-26a2dc79721bfc784f8d9bb8b6355e3945dab028b1943d2b3aae9c375500f2f9)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-518c5cf7e99022e4323c937d996b2facfb322e40c1f6eeb8fca092974cd64596"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c3f109c61d80876c3f792be13b988bc7fc559efefcdc464980007400a6e91ec"></a>

## advertise_custom.advertise_where.use_default_port — advertise_custom.advertise_where.use_default_port / 90f25c33368f / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-9862aa1afb18062aa24aff6f912aca4d21e723d2f02b7f2f6075c9b0b5d139a4"></a>

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

<a id="canonical-5b62297a02b02246c11c35ea11d82ec27f5dd05c50246002091576702dd76f62"></a>

## Direct properties — advertise_custom.advertise_where.use_default_port / 90f25c33368f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e75336707b18aced17ae9a367dfd32830e04d3aa799329de4fc5736c86e9323"></a>

## Next pages — advertise_custom.advertise_where.use_default_port / 90f25c33368f / 4

- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6edefafc4b5f4617ccbbffc10ceb76720e6e67cb40d4d3b41fa0647a94f0cf8e"></a>

## advertise_custom.advertise_where.virtual_network — advertise_custom.advertise_where.virtual_network / 7e078582e462 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-a3d2da93488bbcbaf017009990bbf728b2947fbeb9d1543ac7d1c3498b77f5a1"></a>

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

<a id="canonical-099063909e24fa9e454121643e01d168baf38c00a90f22ed7727d3ad0e3c4333"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network / 7e078582e462 / 3

- [default_v6_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-caa15f0b8fc5764513445fb39e1852e2873a404ecc1588314fc3a2a78ff56cf7): complete subsection reference.

- [default_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-8e72c9ecf35c2222cc9f3d5071609a75a1c1bccdbe94e03240ed57c6a6a581aa): complete subsection reference.

<a id="canonical-09aae16a4b6b41f55898693ae609b2f340a30f3f0a43cc7ebcc448ffca38dbac"></a>

<a id="canonical-049c31ddcfbeb02819079ae635adb6c3bf8582334fb1f0d16c00291829498bcb"></a>

## specific_v6_vip property — advertise_custom.advertise_where.virtual_network / 7e078582e462 / 4

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

<a id="canonical-2afad266e7a0c6777028134c266a106031eac21b564ef7765823ee859f7492bf"></a>

<a id="canonical-c3baa09197a8ef3aa7d6839dbf6615c0957c7388aaa660d564c67494e8a26b92"></a>

## specific_vip property — advertise_custom.advertise_where.virtual_network / 7e078582e462 / 5

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

- [virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-380042b1f00d2a6ee3824da007b11196063d2fc1804315f6e84cc66792b95475): complete subsection reference.

<a id="canonical-197e6919ea827df49b1858973a8b27f66a9255418956b73921973e004d379e98"></a>

## Next pages — advertise_custom.advertise_where.virtual_network / 7e078582e462 / 6

- [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-caa15f0b8fc5764513445fb39e1852e2873a404ecc1588314fc3a2a78ff56cf7)
- [advertise_custom.advertise_where.virtual_network.default_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-8e72c9ecf35c2222cc9f3d5071609a75a1c1bccdbe94e03240ed57c6a6a581aa)
- [advertise_custom.advertise_where.virtual_network.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-380042b1f00d2a6ee3824da007b11196063d2fc1804315f6e84cc66792b95475)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-caa15f0b8fc5764513445fb39e1852e2873a404ecc1588314fc3a2a78ff56cf7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c98b276752e905237f7fac1e0a7ab4345008a7ebdf99ec2b23d4dbacca9ea860"></a>

## advertise_custom.advertise_where.virtual_network.default_v6_vip — advertise_custom.advertise_where.virtual_network.default_v6_vip / f7a0b7e6c5b5 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-b0b3cbf1e837c88e351f808375ec9f754877c74c1e2497abce8db784b6dfb131"></a>

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

<a id="canonical-9505383d89daa88dd7e0d7d0ea0305a47aa085efa5f84a36928f414a785734b5"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_v6_vip / f7a0b7e6c5b5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f52b75766b2e4006f8fd359ff709df94d5c50a536acc855376e5c807908a4d31"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_v6_vip / f7a0b7e6c5b5 / 4

- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-8e72c9ecf35c2222cc9f3d5071609a75a1c1bccdbe94e03240ed57c6a6a581aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b073e2d4b4e17bcddebd6cffef24bb4df223ff5d042988b7c69184766d030c86"></a>

## advertise_custom.advertise_where.virtual_network.default_vip — advertise_custom.advertise_where.virtual_network.default_vip / 1815023da7bd / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-a7fe050bf3d2e0de06147621d44334a04d185a14594803bf6e7d8d9fa58ceb41"></a>

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

<a id="canonical-83c958f906a2c49bcb3dfe7713f4d2692f16c84cbda1e1051f0530268667fb85"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_vip / 1815023da7bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2f03fc3255066fe888687f4590b84f6e8944e931c0737d3178472e026233bbb"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_vip / 1815023da7bd / 4

- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-380042b1f00d2a6ee3824da007b11196063d2fc1804315f6e84cc66792b95475"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8d220d6fa8a91115f113420d6f7fc3e7af5847e73d7bef647a380e1f780653b"></a>

## advertise_custom.advertise_where.virtual_network.virtual_network — advertise_custom.advertise_where.virtual_network.virtual_network / 1a97c4e40dfe / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-87ce824dc486a9f4e16571eb286176f0fc12eb27ca10e40455ca7ff7a0c176c9"></a>

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

<a id="canonical-036f305af248062a022d8472006cc568b2feb51ab524a55449f296b5d583c789"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.virtual_network / 1a97c4e40dfe / 3

<a id="canonical-95fa2b429fc440fe289fa1df9717a6624d5e0d6e26711a357851a68827577aa9"></a>

<a id="canonical-811ad22b166a746b7c764acb0e7d562ebac3de066ea3319c278467741aea7581"></a>

## name property — advertise_custom.advertise_where.virtual_network.virtual_network / 1a97c4e40dfe / 4

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

<a id="canonical-a4da37b53213739c24d36ad0008da4e94eea22352d59684796d7cea8a676b902"></a>

<a id="canonical-f7808374d24e95644abc049a7bb48658715351041d0dd498b3eb34b80642e2c7"></a>

## namespace property — advertise_custom.advertise_where.virtual_network.virtual_network / 1a97c4e40dfe / 5

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

<a id="canonical-f0da9baa497531c47b96394cd4ea9e3485a3ea84ebff6eb7f997c8ab95f93b31"></a>

<a id="canonical-405e020e08cb4b9271ae4cb17cc3fcfc41bf6bad86f0e9fc4735ed6c42048dd7"></a>

## tenant property — advertise_custom.advertise_where.virtual_network.virtual_network / 1a97c4e40dfe / 6

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

<a id="canonical-63349b6ac3fd3685f1fa916392abe493f34f92d46f126e7e71f11e18bba507ff"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.virtual_network / 1a97c4e40dfe / 7

- [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--reference--group-001.md#canonical-435b717e887751410cafa4d262c1ee11552a66894cc581963d26b33f3c23c348)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-5ab4c9d572a5cb0d3dd7c74517ee809160f708db54f37990efdfe466b4e42447"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f4680760ca4b84be4d4e632a7dd1c8dc3d8c1fb208c0d144931f89b716a5c5b"></a>

## advertise_custom.advertise_where.virtual_site — advertise_custom.advertise_where.virtual_site / 8a3b4df61557 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-91308e4c08b0cf943266116f2e880aca3f7faa10b525aecf3f34e82e58cf8e10"></a>

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

<a id="canonical-304292fe725c71d9c8e79049cfd0fb3828d24dc6259c872fb381b631ead0e4a8"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site / 8a3b4df61557 / 3

<a id="canonical-b1ccad1b9e60ade621dc7df6a19fbc807781d71d69d900d69d6b8a2c09cb7ba5"></a>

<a id="canonical-2e6d35329ea7e26e8c61111c820f1a678d39fc6c4ea6d84ed3118c7fd521f49e"></a>

## network property — advertise_custom.advertise_where.virtual_site / 8a3b4df61557 / 4

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

- [virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-0b6b47b14eb354aa1f959d66724d451d2280d4d376a619460646df22710a854b): complete subsection reference.

<a id="canonical-5cdd6ff9de4888e83602519b37ec8eb539a41b4523578edc9dc5f14962e52176"></a>

## Next pages — advertise_custom.advertise_where.virtual_site / 8a3b4df61557 / 5

- [advertise_custom.advertise_where.virtual_site.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-0b6b47b14eb354aa1f959d66724d451d2280d4d376a619460646df22710a854b)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-0b6b47b14eb354aa1f959d66724d451d2280d4d376a619460646df22710a854b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1cad5112e516ca4ba7d97a76256e90a01bbb9040434472b0a71bf586938ba5b"></a>

## advertise_custom.advertise_where.virtual_site.virtual_site — advertise_custom.advertise_where.virtual_site.virtual_site / 8512c7ec2b73 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-5ab4c9d572a5cb0d3dd7c74517ee809160f708db54f37990efdfe466b4e42447)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-a5e9ab30ffca5358d2f840a3d99fd4630c4cc5c91fcd91e547fba8c79c6d7868"></a>

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

<a id="canonical-9535eaa05c815debb64257195a7502793adc2d25d57baed3b854bdfdeab79ed0"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site.virtual_site / 8512c7ec2b73 / 3

<a id="canonical-f55c8f66c97f640e2d8f9dba3828617327c1baf54f94485d76aba8fd3af2c64a"></a>

<a id="canonical-7c76626ace1a1c362e602ae2df0c61e300868f8311085ce9a3b63668fca32ee9"></a>

## name property — advertise_custom.advertise_where.virtual_site.virtual_site / 8512c7ec2b73 / 4

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

<a id="canonical-e199e622ef4394de26c5b9b4afb20377af3b734fcda743d3cb1e0a776872d64a"></a>

<a id="canonical-2db9f3f73e0aa001c9971c36dfe3a9d35a9d366fbc609a7b2f887df1ef00ec32"></a>

## namespace property — advertise_custom.advertise_where.virtual_site.virtual_site / 8512c7ec2b73 / 5

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

<a id="canonical-0821e390c7016706e7ab150b035165079dd96a2a4505bd2cb9065ef5994abb2c"></a>

<a id="canonical-1843bf354c770b78c5ea0522e9c18ce7934cf23f7abc18c7a7c990e7c2006a9d"></a>

## tenant property — advertise_custom.advertise_where.virtual_site.virtual_site / 8512c7ec2b73 / 6

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

<a id="canonical-93d2c94a4e40df182b7e347547f05abde66e63c25ed3bdb7aa20ea1a40a8898c"></a>

## Next pages — advertise_custom.advertise_where.virtual_site.virtual_site / 8512c7ec2b73 / 7

- [advertise_custom.advertise_where.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-5ab4c9d572a5cb0d3dd7c74517ee809160f708db54f37990efdfe466b4e42447)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-5987f0c1e54c132e58207bc26199657ab73a1bb0b77bf939c8683022a3a10c71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b2b46791ceb9b5aa115516f41e1ffb089bc6f0cd06f90b609f2a0efbe777a7e"></a>

## advertise_custom.advertise_where.virtual_site_with_vip — advertise_custom.advertise_where.virtual_site_with_vip / 01253609c7ed / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-3d162dad8162c94fba8949e9d165c3abd17e3f0bd7c8665a786913cc24f879db"></a>

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

<a id="canonical-f5dc4de28b9febf80b24d97b8f2886c6e7d23094eebfad791a685cc63ebbe1bd"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip / 01253609c7ed / 3

<a id="canonical-75001a100eae707748ff933fc4fd5cb85c0ad3aad348054b75c7c856cbd5140a"></a>

<a id="canonical-9194eee2a487e4073a014d336a3e22ac2ff3710751b000ce66b6900adda2c9f5"></a>

## ip property — advertise_custom.advertise_where.virtual_site_with_vip / 01253609c7ed / 4

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

<a id="canonical-608829ecf8f91e9783a38f280ddc944f62346907aef9961d784c5a0d1c03fb81"></a>

<a id="canonical-8cf8181b3e61a78562538a4108d260e4bf16d7b0319bacc46323208b2295b44f"></a>

## network property — advertise_custom.advertise_where.virtual_site_with_vip / 01253609c7ed / 5

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

- [virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-85488c99a77d3be6cca951ff4bdfd1c15b37d25b239d0d01911428ff8f57a1db): complete subsection reference.

<a id="canonical-9993c824e9b2f18434f1ae54b76a3c5d21a3148970e16c7d83a748aa834c5d9f"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip / 01253609c7ed / 6

- [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-85488c99a77d3be6cca951ff4bdfd1c15b37d25b239d0d01911428ff8f57a1db)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-85488c99a77d3be6cca951ff4bdfd1c15b37d25b239d0d01911428ff8f57a1db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-839be6cd7273837e75c483acc2c4a88c48396015a22cec3d160aa785e1d25963"></a>

## advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 5de5eab9aaf2 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-5987f0c1e54c132e58207bc26199657ab73a1bb0b77bf939c8683022a3a10c71)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-41899cb0ad5ce13bea67726b5dd461f1d384c5a144dd448c5f443905b5c32ac2"></a>

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

<a id="canonical-ac07b31794e403fe5ba35c57d57a21a5e6ee2026484af194dadadd2bafdd9106"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 5de5eab9aaf2 / 3

<a id="canonical-f06b11ed8a7101a252a0e64f776ef6903a5c85588a98c0a76f6501f7df0a81ac"></a>

<a id="canonical-b4bd0f37892074dd40357361fdb05045c5fc983382ed5183628d66b0efa283f8"></a>

## name property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 5de5eab9aaf2 / 4

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

<a id="canonical-030fc97b5d7a18e735e09349908012e25ad49e2784af454e9730a734d19f712c"></a>

<a id="canonical-5ed4aaadf250974518b8eb2ed69483c5af602a22b64ecc31e9c8af4f26550371"></a>

## namespace property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 5de5eab9aaf2 / 5

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

<a id="canonical-b502cb4dd2f99323240961f076c78c7e027ff9d210f3170c88604ceba7ab77fe"></a>

<a id="canonical-4377369cf31a247ad664bc7324be25efb4ab07190dd3fcdbaa1d1a725112f655"></a>

## tenant property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 5de5eab9aaf2 / 6

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

<a id="canonical-356de57d59975511d9374056e40d380547e4024e05b37cf49b03672b4a54f61d"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 5de5eab9aaf2 / 7

- [advertise_custom.advertise_where.virtual_site_with_vip](resources--tcp_loadbalancer--reference--group-001.md#canonical-5987f0c1e54c132e58207bc26199657ab73a1bb0b77bf939c8683022a3a10c71)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-b4392d7e78138a30c0757783cdc3c073c5da6748af1bf6264da06f58d293c38a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9aed7e763e19a68e1653bdfb90d6939f0363d2a0f1984d6f4a99d4b63249cdab"></a>

## advertise_custom.advertise_where.vk8s_service — advertise_custom.advertise_where.vk8s_service / be3bb460928a / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-8ece6b0320aec582deba74d48114afb23e82007e8f5751c3c6097b9cd2461ddd"></a>

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

<a id="canonical-a6cdaf3efe5b539921bdd410af0669e9b9e99a7f408481ccc3f3ea4e759f7c87"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service / be3bb460928a / 3

- [site](resources--tcp_loadbalancer--reference--group-001.md#canonical-056b4313fab2ed6aae7cbce4b1c8e018c103cb1c2ed447cc4946e7368dd17246): complete subsection reference.

- [virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-12091f35e3d77af04cc9a6b5bdfa332921715db2745aad9cdfc417c7bf704bfc): complete subsection reference.

<a id="canonical-f5bd18d2b4694bfea2c58d32f091a5bfd2a3cd262aa4350cbda4b79c4b371f4c"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service / be3bb460928a / 4

- [advertise_custom.advertise_where.vk8s_service.site](resources--tcp_loadbalancer--reference--group-001.md#canonical-056b4313fab2ed6aae7cbce4b1c8e018c103cb1c2ed447cc4946e7368dd17246)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--tcp_loadbalancer--reference--group-001.md#canonical-12091f35e3d77af04cc9a6b5bdfa332921715db2745aad9cdfc417c7bf704bfc)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-056b4313fab2ed6aae7cbce4b1c8e018c103cb1c2ed447cc4946e7368dd17246"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df7e533d9bc9734f725b4b7fb941f5a95cdf20a7f2573456c66f6ba71652dd34"></a>

## advertise_custom.advertise_where.vk8s_service.site — advertise_custom.advertise_where.vk8s_service.site / 91dbceba7633 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Property reference](resources--tcp_loadbalancer--reference--group-001.md#canonical-869f7b6ed7a9bd4c762d24a87f33328ade5a5db7347f7a3df79a7dfc9e69d541)
- [advertise_custom](resources--tcp_loadbalancer--reference--group-001.md#canonical-111dced2152011ee6b0e4bc9d68b85087428bd7443017630b688a8872ad6645c)
- [advertise_custom.advertise_where](resources--tcp_loadbalancer--reference--group-001.md#canonical-a835c8d2a41431c2f5981ca03ca12a27363f2aa3f1d784ba8283550b072298fe)
- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-b4392d7e78138a30c0757783cdc3c073c5da6748af1bf6264da06f58d293c38a)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-cfb0afb92d9af0de5e0d87a0a536da056d4776f9b4120604572e8bf5272fc371"></a>

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

<a id="canonical-879b344f921e11e4f6cd84e9bd4cefe959338452fb2f6b976cea1e8fc67001b2"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.site / 91dbceba7633 / 3

<a id="canonical-dc3ba9ccce2921b7f790c770e51c8ccc716a8e0aa14ecfb7617918b8d0a73195"></a>

<a id="canonical-d889c913ccd965d93fbc253aef15c924f9c0cc901dcca662c59c58d6fdfddd87"></a>

## name property — advertise_custom.advertise_where.vk8s_service.site / 91dbceba7633 / 4

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

<a id="canonical-6e7e538357ec90a5be6ec3a498bbb29efb835928f249bbcc0c0ff4ea5e22c7d0"></a>

<a id="canonical-e594b3901c22c2c78e057e6ba1befee56fd63134e5049e2b96bd13e4c6a9d0e8"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.site / 91dbceba7633 / 5

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

<a id="canonical-00dac523dbf767ff4a55e9fc800640aed85a29033f0c443186b4dbe6ddd0202f"></a>

<a id="canonical-3ecf8a408ba008b9e1525800fa8fd1f543021de56927f7d138f4753d7ce2daed"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.site / 91dbceba7633 / 6

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

<a id="canonical-e91567401dd09f84f993ceeaf40bfc3d1c053d956444e52dd4d4623aecff5d58"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.site / 91dbceba7633 / 7

- [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--reference--group-001.md#canonical-b4392d7e78138a30c0757783cdc3c073c5da6748af1bf6264da06f58d293c38a)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-12091f35e3d77af04cc9a6b5bdfa332921715db2745aad9cdfc417c7bf704bfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
