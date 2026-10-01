---
page_title: "xcsh_udp_loadbalancer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer reference."
---

# xcsh_udp_loadbalancer reference

<a id="canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12140ecc6d6c9b5b51c899a8df8dec6ac62e68493063ae075995c880bbd32733"></a>

## Property reference — Property reference / 565c0f1574cb / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- Property reference

<a id="canonical-3293357c592ee1d28bb487af6d04a3ab6adeb9a5fa93abfe3fde8ccaeeb75864"></a>

## Direct properties — Property reference / 565c0f1574cb / 3

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2eb765de055c7f44fd7c1551b36f0eb5360688b65a6ebd1597bd540b8854e95e): complete subsection reference.

- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c): complete subsection reference.

- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-d9064ae10247a9cba717e2df2b773fae22b197233ca5a3e1704fe32571228c61): complete subsection reference.

- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-4e7cc4e708953afbd0abe22443f62ca3ff474f436045a846be4d4c8abf34e593): complete subsection reference.

<a id="canonical-ea71eb237ccc56336e81b40e1e797442de03aaf6b6c9bb9d96e589e30eaa9946"></a>

<a id="canonical-0cdb3082d706b3e7319bbf29fcc01ced58a30633c6839a14f13f309336cdd7c8"></a>

## annotations property — Property reference / 565c0f1574cb / 4

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

<a id="canonical-0c70194c4dd3931d6f8020c8f711adf5e28f848861116d4343195f090c743628"></a>

<a id="canonical-628a707fa3927d9e148091091434c90041b33429a3a5d1e94f7ff82f160efe81"></a>

## description property — Property reference / 565c0f1574cb / 5

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

<a id="canonical-e8e19876b74e65571f2d3e423515719025464d9536fe8eda779f0bbe1ce86422"></a>

<a id="canonical-efa431ed9a8233586b9de2810d96c6e52ea58bb3d8dc7ddaa960f19a67448ef2"></a>

## disable property — Property reference / 565c0f1574cb / 6

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

<a id="canonical-72d3ee20704266c2bf09e641a1b60cabd6f3482c430e1429869a31ada7fe0da0"></a>

<a id="canonical-4d2edc049a6bae1263c388a0ed4e6a5d5f71a92042137512060770d02003bffb"></a>

## dns_volterra_managed property — Property reference / 565c0f1574cb / 7

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

- [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-8918004147bc82a6741b40ac27a78269f7d85ad32ff9788081e710294eaf0a38): complete subsection reference.

<a id="canonical-478591da172fc563f521858017b6b89e63a9c9a4cde3e6b2e25ad286aba316ec"></a>

<a id="canonical-6c819bf6a11ac7daf0562b18a6f6f0e0550a5cff7cbf4c5ecbc24ce0c46bc32b"></a>

## domains property — Property reference / 565c0f1574cb / 8

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

- [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-9c238b54b4fa70dc93c8d69891fd3b0a930acc5e1179fe65435c9f2b5408d0b7): complete subsection reference.

- [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-0f00a843733878244fa6b6b74122b7e929fc15c072750ee11524745d72f84199): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-6e8bcc7eafb9a3b0f56ec6102957cacd26fdf582b8894df2a649dad99c1500a6): complete subsection reference.

<a id="canonical-538c3a44b2722994d19f2fc78aaf4acc3e30b5a3efb8ea958c9b724b63ba4df1"></a>

<a id="canonical-d1dcd83b2748c1ab29f7c48ed5c3494c6f831c7132b1b38d60bb6ec7b9981084"></a>

## id property — Property reference / 565c0f1574cb / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-03b340da3a10f3b09a274253f9985fe3ba0ac6efb900b772cde781c401cfb41e"></a>

<a id="canonical-b82ca89db78ad4eefe359273eaef3586259e04ea581c74f35ef630c8b46e4986"></a>

## idle_timeout property — Property reference / 565c0f1574cb / 10

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
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-dc27011d955d2cc4d39696797a71f24b622e536d598a63484b70ba5f7d94da2e"></a>

<a id="canonical-783ab211b08191b909b85d6a8489e07069f1f58a83224fdd8d8d7882898f0430"></a>

## labels property — Property reference / 565c0f1574cb / 11

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

<a id="canonical-d89a37a235cc845461ae447fdbb555e2a3789c8678e8e28926a05181a411c4a9"></a>

<a id="canonical-daf27069e6fc8cb29f78870b745d9257710156064e73141fbb164d45ec37fed1"></a>

## listen_port property — Property reference / 565c0f1574cb / 12

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

- [listen_port](resources--udp_loadbalancer--reference--group-001.md#canonical-d89a37a235cc845461ae447fdbb555e2a3789c8678e8e28926a05181a411c4a9)
- [port_ranges](resources--udp_loadbalancer--reference--group-001.md#canonical-a320355abf18bd9a6d489aeae7d29a80dd6150a9676d10dffc267ef90a5af52d)

Select alternatives according to the provider validators above.

<a id="canonical-3df203ffee0b46352410c9b6d6b5056cd08726ffd42a531e0a711c02bda353ee"></a>

<a id="canonical-1fe9dbf1ae4137af46350e29025863a05d75c0d646af000383b2b49d80008b2d"></a>

## name property — Property reference / 565c0f1574cb / 13

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

<a id="canonical-49a03420ce2802acb6ca63ff9f7ebb4f52d921aea0fa55cbf69191c18d3fa87b"></a>

<a id="canonical-ad9178e48ee243fc2d05b2ef5055eec65dc63048f27628ea9f617d88b051bb09"></a>

## namespace property — Property reference / 565c0f1574cb / 14

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

- [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-aad9f2a76c2fc2765eedd5ecbff8af5a9c083c5f8e7df90fd0dc82310dbeeb06): complete subsection reference.

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3): complete subsection reference.

<a id="canonical-a320355abf18bd9a6d489aeae7d29a80dd6150a9676d10dffc267ef90a5af52d"></a>

<a id="canonical-a8a3ac89ba14d8a621dcc0c45127157c1c4ad5a8f3117d6a6d0724fb803e15b6"></a>

## port_ranges property — Property reference / 565c0f1574cb / 15

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

- [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-a943aedfea42ed379f11cd0ded88b44ef75a05109c7e9807a51bbfe5039c7cec): complete subsection reference.

- [timeouts](resources--udp_loadbalancer--reference--group-001.md#canonical-6a1cd6fb1eca53775784d20ffcccb4f95e864f18ca39406d848757a84e4f7c95): complete subsection reference.

- [udp](resources--udp_loadbalancer--reference--group-002.md#canonical-d95efc1645c481a1dff9b350771ac01bbb10588a5a3ef681c1dce0e98f8af44c): complete subsection reference.

<a id="canonical-36ee1742870ac46d4e9c2582186829c3db46b68a1ab55c3e8b1684a834d9079a"></a>

## All schema paths — Property reference / 565c0f1574cb / 16

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-5743947cbc6976b0fc4c91c9295ed922456ce3899018f763cbced92c4801ee28) |
| `active_service_policies.policies` | [active_service_policies.policies](resources--udp_loadbalancer--reference--group-001.md#canonical-c9e3cbe92a1c700768374106e28e05279cf65146712aedd3675fbee3157d178e) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](resources--udp_loadbalancer--reference--group-001.md#canonical-753fc391b0f497a669f9a3fdcf5db1af274b60f8305e7fcab4d6095a441ffeb6) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-c07edc1c487bf98b2b2a2f3846f1683cb59125b1836d5df4f052a9a39133c431) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-13fa4911711e890f0ddbffdcf7c9549b0d3c195791acca28a9df258e0ded0e56) |
| `advertise_custom` | [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-a85a6369eb922defa74228bc89a5d3930713bdfb290f9ff30d81a5d4d534ba89) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-0492260194c789cfaf6e326ad809bca3af236246fdc68faa8f343c718591a98f) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-bc44aa6f3b23dd52b7a33820c94f8a88db93921549accbd9fa143c065fd59b30) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-d69c9c9702b270c6ec58c29c66dc3de94d58647827f9073c1d155e06b9309499) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-47a60c9bccbdfd556e7cedbaa6c4b1a8a6168b0e536d4e27d5c904c28085be40) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-ebf10bd3b1a9ad8e372d8bd949368a2426c110e6a8ad9fe6ee93baf26166b70b) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-2c52e5807fc57b287d25531022d223f1dda71f34fcdc4b42ad3403e1e4ae1be4) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-9d8b2fd06792979d3938a9fc54edf1b717690a57af7ebcfa7def4b17e97d6bad) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-d1ee059fb92e782d3a520ec2253bfbfdc6cdb1946d988cb58a1237b909786fcd) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-a7b43af290b022987fd527928ac3d3739ab69763af1db9a655f12aea57b2c4d4) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-452c1b32091c3bf58e8fc5d70b927f3eefd81b4d073385cd1c88296a27f71bb3) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-714699e3442fc83b930a620b0a271b13949c9903de2ca9d6487af18bd6ffadf8) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-1af50843d119f55a9cd9d32e2bcec8ac8ed64877fb636d79cff75de6965eeb67) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-f13f02f77f01392ca7892a74d7fa26c9b5499cf724f45b89dc297cd4af792f8c) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-cb808c2ed5e5a072f4a23b97bf8f97f3c646f86dc51b7469d8fa6b26e4a682fc) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-e40b301a3d30d88a86001d728961fdb032eb0c0f4505f9e970b1a6c71b119062) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-fbb86b04b4bf2c638187c531325723a2d54a61db7b6e4f5969bf67d4aee529b8) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](resources--udp_loadbalancer--reference--group-001.md#canonical-20f8459eb48522800a2a60d675388208e0ee01d2b3cf40a96d68dae02e8cc046) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](resources--udp_loadbalancer--reference--group-001.md#canonical-9b5d7efdc7ccf939d925688f07e1db08170dbbdeb3234516cff249e961a96bda) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-087933082ce9f9fdb8367f31c0036b17194eb612c096469c1cecc3510ae48f26) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](resources--udp_loadbalancer--reference--group-001.md#canonical-01da621462d0e361850ac4b81aec078360e70adbc97d25665e935619aa2c3b13) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](resources--udp_loadbalancer--reference--group-001.md#canonical-b031d8aeef0895ebdbcbae822e2cc485b51df3f86c0caf1108f0355653e22a40) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](resources--udp_loadbalancer--reference--group-001.md#canonical-5cd2fafe560f75e305a9d689c68674aea7f737d011b74ee0ba8a3754e19cb0f6) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-9d1b0615cdf758bbd16c799490f02250c8f9c8fb63008ca5a8715c954f7c1c02) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-8ccd4a8ebae674bdd5c347b5bfea0e9b3ffde761025562d116e4da13394e8141) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-385246fafc94dac23c7e9550d8d95b48f89d16143aa6481445a5a37632a83d16) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](resources--udp_loadbalancer--reference--group-001.md#canonical-30ee0b30d377ad49fc04e36995d701d738ee9531d8a7a6b68f45854619843f93) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-b197266511d523fff7766c382c9f74c0b1628e671bcb1a1ea88e6a7de802de1e) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-8c96ad2057d0d236ae55c2b912ed433934a1e72939b45921db51318d3aaa114d) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-2b2c6738d31f2ffa2c155bec602a44ed4567bda95f962ba26271d36a79c4b9db) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-1f4aa6d96c2589c808a8fe087c88a28550953d12a8b5803b0fbc7a35fa072651) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-3e8ad8fc338247453a9610b9358d1c6108ac59182835ef1f9da590a818387bb3) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-9a8d30efe1c256543cf154fa889f8ef881e8ea541a694f84803396fb38b6d9c7) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--udp_loadbalancer--reference--group-001.md#canonical-9868c37ffce5eef41edd5b8a601a1cd8bf1cffee05b55e7a49c332d7e40f1a19) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-86036fad9b815d40bd0e4f27bdaa038ffb587d587ce1da179c2edcec1b0f312e) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-aad25099bdac93d88dda5f4825df3b52850d15a727659cd964ad986a7c9a3ec7) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-ee20549408fd07bfee157d4ca81f5e135ecdfe71b360d450f0aaf4bbc99396cd) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](resources--udp_loadbalancer--reference--group-001.md#canonical-686a723f88835140b30d2e874e9ff49bd915431519a34bb78037d4c43a8f7bc3) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-7d6006293ce5dc9c34fb9656cba0d089c7d860221b644e5fe8ed97672458f4f5) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-36fb4da37969c4cb4dfb470fa3d63e61c58fbe52145d3a857675a3789045fa8f) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-bf2d459d240333c9ad2695218184589fb622521d996d3b6f681b8fb8d256da69) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-bf5a481af3163c5c5398719f65cfa480433c6ce614ed650976340b848c14102c) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-731f931eeead98ad065e78c4a0072ec1fae7d25ea9e6464ff0a4890a079d4ba4) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--udp_loadbalancer--reference--group-001.md#canonical-7c668402686f4d83785ac447e8a2429458a0f676939f42426d52f4c2414dd89e) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](resources--udp_loadbalancer--reference--group-001.md#canonical-15556c2a5e8c5b5a811d80cdc5a469058e2dbc9e34bebdff7b55b9461813502e) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-63f2f38529cd0cbb4776f2dc6875307a4180ee8d1a06e989d94ebef04bfb695e) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-e31a7893fc339175a2e93c1746ff1061c9c5ada39028b8b7c2be500ec34670c9) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-4505e5f1ff10700f9c65dc2bbc7def2bdd83e3a51dd0ae43b429ded4885c7af5) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-28d9b9af6e8ac5b9b3b66110ff31345d4f5d70b0cfa974c6c6b37e18253898d5) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-1e8d43cf6b357f82089e8d3a29ba02baa628ee4254ea292b128afa3645bcdb3b) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](resources--udp_loadbalancer--reference--group-001.md#canonical-54045ef5b28bcbd40c65169e8eb27d704c13af87080dcdb579427935c3f1ac80) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-b3870faa062c1a4c978f9bb9edb2d9851a8e7c8e75ea02e9f468cda7c822ba26) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-05e7c6a7620c511837ff89c172aebb075f433c778ab82bf28541c926417f099b) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-5d616a4e147216f755e7f836445e57d80e8c6f59c31e91db4687adce42cb3784) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-f132e12661fffa7ff4ac51d3ad46819b209e91c5e7fb37e1a5f15cb7861c7580) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--udp_loadbalancer--reference--group-001.md#canonical-22269fa3639ab7b1235d71b7d7990b9bae5151f0db509525fb7ad2197ce60167) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-574469de3e2d2e8824e14ae54c07d6ced972a3e471aa99332523da55e331de4f) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-1d006e3d5f6880bc2fe5f1ba618e4df1a41446d04159b942fbeec07d006a28c8) |
| `advertise_on_public` | [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-5df43900bd5c0407ae764c67495ace806a9f36c132e913c1047635999ba07957) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-304dae4bc07ba465aca88a5f06379964b033d4f8a1342908bcc69b34e66e27cd) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](resources--udp_loadbalancer--reference--group-001.md#canonical-155193a3be4db577501d5764072ecb4ade5bb6967e92507909e61d958b025dbe) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-a3a7dc7a1fc425670f065ab0f8a7a6fdff4cb2cbc6ff22d7fc5f0300798a8b9c) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-fa7f95dd43823742da5a3f66cba715a310b98a770f2a7a1d5292ce9e59b3727a) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-b2edb84e8196bc951ea0fdc2c9c63576398749833bf9a7b65df2ae1fb47f806f) |
| `annotations` | [annotations](resources--udp_loadbalancer--reference--group-001.md#canonical-ea71eb237ccc56336e81b40e1e797442de03aaf6b6c9bb9d96e589e30eaa9946) |
| `description` | [description](resources--udp_loadbalancer--reference--group-001.md#canonical-0c70194c4dd3931d6f8020c8f711adf5e28f848861116d4343195f090c743628) |
| `disable` | [disable](resources--udp_loadbalancer--reference--group-001.md#canonical-e8e19876b74e65571f2d3e423515719025464d9536fe8eda779f0bbe1ce86422) |
| `dns_volterra_managed` | [dns_volterra_managed](resources--udp_loadbalancer--reference--group-001.md#canonical-72d3ee20704266c2bf09e641a1b60cabd6f3482c430e1429869a31ada7fe0da0) |
| `do_not_advertise` | [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-299da560f186457acdaa0655c3e4327ae8c6b1c124ac88ca3bc4d486defa93a0) |
| `domains` | [domains](resources--udp_loadbalancer--reference--group-001.md#canonical-478591da172fc563f521858017b6b89e63a9c9a4cde3e6b2e25ad286aba316ec) |
| `hash_policy_choice_random` | [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-2593b25f47b2d6b0ca6e599ea8124322a0ccee145c7be102a199d9af68f8ab58) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-e30839a403148e9818b5f0e75b5b07813437b93c330ffd670030f0ecf904a448) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-f6e867f0958edb0b92159ec6a85b755c0ffd272b2fcdb4f18faeda97ccac0e43) |
| `id` | [id](resources--udp_loadbalancer--reference--group-001.md#canonical-538c3a44b2722994d19f2fc78aaf4acc3e30b5a3efb8ea958c9b724b63ba4df1) |
| `idle_timeout` | [idle_timeout](resources--udp_loadbalancer--reference--group-001.md#canonical-03b340da3a10f3b09a274253f9985fe3ba0ac6efb900b772cde781c401cfb41e) |
| `labels` | [labels](resources--udp_loadbalancer--reference--group-001.md#canonical-dc27011d955d2cc4d39696797a71f24b622e536d598a63484b70ba5f7d94da2e) |
| `listen_port` | [listen_port](resources--udp_loadbalancer--reference--group-001.md#canonical-d89a37a235cc845461ae447fdbb555e2a3789c8678e8e28926a05181a411c4a9) |
| `name` | [name](resources--udp_loadbalancer--reference--group-001.md#canonical-3df203ffee0b46352410c9b6d6b5056cd08726ffd42a531e0a711c02bda353ee) |
| `namespace` | [namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-49a03420ce2802acb6ca63ff9f7ebb4f52d921aea0fa55cbf69191c18d3fa87b) |
| `no_service_policies` | [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-8566e3ab18cd76803d80747be5b443cbf6c666d10ba37595089ac4457560b3c8) |
| `origin_pools_weights` | [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-b14d6230706c3a14278b27880febcca55d99e25c7ffa7a54ef7360cc73f054da) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](resources--udp_loadbalancer--reference--group-001.md#canonical-8ef1005996684b05ccca852b9fc42884c6ec032b7df08150052ccf912ac80896) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](resources--udp_loadbalancer--reference--group-001.md#canonical-80fc60542d098491f30d99047ee570d290f9f88337c6760473f5363ad87af240) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-85070bf77b9bf460549faf438832d453d36ed360d715c3df3735353432378cc9) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-dac5fabfa657efc4c1544b44228c914698820db4f48feb064f63b5ef47442187) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](resources--udp_loadbalancer--reference--group-001.md#canonical-9ed9b555c3725b19f079d2a7ff5df90ea183612b3901b3348b4913fedb2e7a61) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](resources--udp_loadbalancer--reference--group-001.md#canonical-e637954e7b12d3ce4ca460cfcc383d949e4f65cf24e16e5f354a7b16bb0f7bc3) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](resources--udp_loadbalancer--reference--group-001.md#canonical-126c664776d70d2aa1466cac801afa3646248b2beb89922bf54e02635ba80b80) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-a7105bbd4e782fd8369614c8a913b550345a04a19b8a1d09415dd216e616738c) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](resources--udp_loadbalancer--reference--group-001.md#canonical-e46becef0ad629da3719fa60187cb22d5c26c35c3bc6fcd895b2933204d627b9) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](resources--udp_loadbalancer--reference--group-001.md#canonical-a84bfe8ea1ac276c0c9bc4c13b2eff6648aea542c8bb8834bfd780f031d59ca8) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](resources--udp_loadbalancer--reference--group-001.md#canonical-23ffdac21e5043473e49b28ae73a5efd26df2373025739857daaf7025102b30b) |
| `port_ranges` | [port_ranges](resources--udp_loadbalancer--reference--group-001.md#canonical-a320355abf18bd9a6d489aeae7d29a80dd6150a9676d10dffc267ef90a5af52d) |
| `service_policies_from_namespace` | [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-af52f868a3f0414fa7c4768e484c8f217966c71e5fc2ff84c9059cfd65031785) |
| `timeouts` | [timeouts](resources--udp_loadbalancer--reference--group-001.md#canonical-8929daabbf9810669cfa00dd08c809be7cedd7783505382a2a8f23673cf25421) |
| `timeouts.create` | [timeouts.create](resources--udp_loadbalancer--reference--group-001.md#canonical-c04b66d9aa1cd13aef64e8af5415ee96e7c864d2ec3c69d27e5c074ad680c161) |
| `timeouts.delete` | [timeouts.delete](resources--udp_loadbalancer--reference--group-002.md#canonical-aad5010b11f3777368a49591eae0ffe43b7cae41a9afe936f245bb40f8e53127) |
| `timeouts.read` | [timeouts.read](resources--udp_loadbalancer--reference--group-002.md#canonical-a008e7da5d3bc52306de89977122267bae3338412452280fae0f0b63e9eec139) |
| `timeouts.update` | [timeouts.update](resources--udp_loadbalancer--reference--group-002.md#canonical-717d775e3bf523c3de93b50b3301a6afe937251ea478f596a8c52b54f4a674f4) |
| `udp` | [udp](resources--udp_loadbalancer--reference--group-002.md#canonical-679f3148a643ae01c68df2c6228142bc4fbde22a45250bdfa3beb537fbb80270) |

<a id="canonical-6e3d3e9db39bc22bad7604babfbd6f8d361d229bb25bba5f638852eee8422819"></a>

## Next pages — Property reference / 565c0f1574cb / 17

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2eb765de055c7f44fd7c1551b36f0eb5360688b65a6ebd1597bd540b8854e95e)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-d9064ae10247a9cba717e2df2b773fae22b197233ca5a3e1704fe32571228c61)
- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-4e7cc4e708953afbd0abe22443f62ca3ff474f436045a846be4d4c8abf34e593)
- [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-8918004147bc82a6741b40ac27a78269f7d85ad32ff9788081e710294eaf0a38)
- [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-9c238b54b4fa70dc93c8d69891fd3b0a930acc5e1179fe65435c9f2b5408d0b7)
- [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-0f00a843733878244fa6b6b74122b7e929fc15c072750ee11524745d72f84199)
- [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-6e8bcc7eafb9a3b0f56ec6102957cacd26fdf582b8894df2a649dad99c1500a6)
- [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-aad9f2a76c2fc2765eedd5ecbff8af5a9c083c5f8e7df90fd0dc82310dbeeb06)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3)
- [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-a943aedfea42ed379f11cd0ded88b44ef75a05109c7e9807a51bbfe5039c7cec)
- [timeouts](resources--udp_loadbalancer--reference--group-001.md#canonical-6a1cd6fb1eca53775784d20ffcccb4f95e864f18ca39406d848757a84e4f7c95)
- [udp](resources--udp_loadbalancer--reference--group-002.md#canonical-d95efc1645c481a1dff9b350771ac01bbb10588a5a3ef681c1dce0e98f8af44c)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-2eb765de055c7f44fd7c1551b36f0eb5360688b65a6ebd1597bd540b8854e95e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0bad1b3dbd60abcb16d905f928a5cf5a857ea212c93433cc1d28f5aec111561c"></a>

## active_service_policies — active_service_policies / 781a6ce53b4a / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- active_service_policies

<a id="canonical-5743947cbc6976b0fc4c91c9295ed922456ce3899018f763cbced92c4801ee28"></a>

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

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-5743947cbc6976b0fc4c91c9295ed922456ce3899018f763cbced92c4801ee28)
- [no_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-8566e3ab18cd76803d80747be5b443cbf6c666d10ba37595089ac4457560b3c8)
- [service_policies_from_namespace](resources--udp_loadbalancer--reference--group-001.md#canonical-af52f868a3f0414fa7c4768e484c8f217966c71e5fc2ff84c9059cfd65031785)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
active_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3b6b906329f641b038b7e74a63a3d8bb173a2a0d4a948ba53df3dbe5b90d1ed2"></a>

## Direct properties — active_service_policies / 781a6ce53b4a / 3

- [policies](resources--udp_loadbalancer--reference--group-001.md#canonical-a714924c68e2e38d98cb34e545635c59dcc93febf54237c956b1b09ca0da8e20): complete subsection reference.

<a id="canonical-d7d5dc05bd60ac1d7d58fda73a5e9591e5e2e54474dfc0af478471c45928dee7"></a>

## Next pages — active_service_policies / 781a6ce53b4a / 4

- [active_service_policies.policies](resources--udp_loadbalancer--reference--group-001.md#canonical-a714924c68e2e38d98cb34e545635c59dcc93febf54237c956b1b09ca0da8e20)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-a714924c68e2e38d98cb34e545635c59dcc93febf54237c956b1b09ca0da8e20"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f38b3cbef96628aaf9513fcad89147034de846a9db2928263bc9a6265c381c5f"></a>

## active_service_policies.policies — active_service_policies.policies / 2f7111ded0d3 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2eb765de055c7f44fd7c1551b36f0eb5360688b65a6ebd1597bd540b8854e95e)
- active_service_policies.policies

<a id="canonical-c9e3cbe92a1c700768374106e28e05279cf65146712aedd3675fbee3157d178e"></a>

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

<a id="canonical-c4d418025287c47d7ac931d362793d9a725ae8cb8ece4b8aa90f6f0df42be396"></a>

## Direct properties — active_service_policies.policies / 2f7111ded0d3 / 3

<a id="canonical-753fc391b0f497a669f9a3fdcf5db1af274b60f8305e7fcab4d6095a441ffeb6"></a>

<a id="canonical-a1a8a65120b220e97d1250a23bb37ff6f992bb8ab1a775ebd9f11e9d04e02f2b"></a>

## name property — active_service_policies.policies / 2f7111ded0d3 / 4

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

<a id="canonical-c07edc1c487bf98b2b2a2f3846f1683cb59125b1836d5df4f052a9a39133c431"></a>

<a id="canonical-733f71a71ff41a219b583722ce88a2feadee146d963fce08cfefc6868f7ea071"></a>

## namespace property — active_service_policies.policies / 2f7111ded0d3 / 5

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

<a id="canonical-13fa4911711e890f0ddbffdcf7c9549b0d3c195791acca28a9df258e0ded0e56"></a>

<a id="canonical-9fc4701b93f753dec40dd8270e7363216376b56352564ac7d082958e25c84ff4"></a>

## tenant property — active_service_policies.policies / 2f7111ded0d3 / 6

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

<a id="canonical-0f0be46a3229cb289cb4981f1f0b78135486cfcee260ca5da89171ce271d6b84"></a>

## Next pages — active_service_policies.policies / 2f7111ded0d3 / 7

- [active_service_policies](resources--udp_loadbalancer--reference--group-001.md#canonical-2eb765de055c7f44fd7c1551b36f0eb5360688b65a6ebd1597bd540b8854e95e)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fc633a7fb1a2cc9d335ef9d848a02f79cad24f5e411a2d97ee67962a7c4a30d"></a>

## advertise_custom — advertise_custom / 027702772fe1 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- advertise_custom

<a id="canonical-a85a6369eb922defa74228bc89a5d3930713bdfb290f9ff30d81a5d4d534ba89"></a>

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

- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-a85a6369eb922defa74228bc89a5d3930713bdfb290f9ff30d81a5d4d534ba89)
- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-5df43900bd5c0407ae764c67495ace806a9f36c132e913c1047635999ba07957)
- [advertise_on_public_default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-b2edb84e8196bc951ea0fdc2c9c63576398749833bf9a7b65df2ae1fb47f806f)
- [do_not_advertise](resources--udp_loadbalancer--reference--group-001.md#canonical-299da560f186457acdaa0655c3e4327ae8c6b1c124ac88ca3bc4d486defa93a0)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-dd2cdc8b31d013aebeaed2e05a34c1d2e8ae9f0df5dda1d60e79bdefb795a8a2"></a>

## Direct properties — advertise_custom / 027702772fe1 / 3

- [advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911): complete subsection reference.

<a id="canonical-33ee42c28fe0c059e06f59edade4c42ef8ca2114744f42f944f9e6668e44245e"></a>

## Next pages — advertise_custom / 027702772fe1 / 4

- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58dccf4392280fb2733313074b66a53286547abbc1913d194ffa86609366262a"></a>

## advertise_custom.advertise_where — advertise_custom.advertise_where / 4453f49eea54 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- advertise_custom.advertise_where

<a id="canonical-0492260194c789cfaf6e326ad809bca3af236246fdc68faa8f343c718591a98f"></a>

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

<a id="canonical-b68b4fbcc5575103c87dd1217dc13a54c3f9cf5cfa0ea0341aca20acd89a8479"></a>

## Direct properties — advertise_custom.advertise_where / 4453f49eea54 / 3

- [advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-ef5df0371c7ad3ca8c46d55ebad67737d80fb496da4f360658a6a192928cee9f): complete subsection reference.

- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-4f3ab94e8105fc20a214b4b2b3e713e5272ddad0732a95b2c0d1884c421e4401): complete subsection reference.

- [advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0fa479eb77e8015e4314425133e9de52900c9651371dd972a4d6d80921d2325f): complete subsection reference.

<a id="canonical-20f8459eb48522800a2a60d675388208e0ee01d2b3cf40a96d68dae02e8cc046"></a>

<a id="canonical-4eb4bdd6dfac15c60c02cb4cd9f71f33dc20e96c735672d196f41932cc4768aa"></a>

## port property — advertise_custom.advertise_where / 4453f49eea54 / 4

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

<a id="canonical-9b5d7efdc7ccf939d925688f07e1db08170dbbdeb3234516cff249e961a96bda"></a>

<a id="canonical-0bc94abdfdeeb26a03e204f499382d3f70db24b6568663e85c682fd9daaeb3f9"></a>

## port_ranges property — advertise_custom.advertise_where / 4453f49eea54 / 5

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

- [site](resources--udp_loadbalancer--reference--group-001.md#canonical-9d0dd20a65b48538b919d73be80414f37a7dd471f2f9f734bccf5c9f8e1bd6c3): complete subsection reference.

- [use_default_port](resources--udp_loadbalancer--reference--group-001.md#canonical-eb3f098e7fd6914708208eaa0785e5cd86821259629b73ea12bd40abc691fbde): complete subsection reference.

- [virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304): complete subsection reference.

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-ddbcc812edfb98f39df6d4addb365f523cd76e2b2577a5c2710ea9b81f3d6e3c): complete subsection reference.

- [virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-7cf1e5410ce18045e6b57379d4afc1000d647abffe63f2e242400e3ba2d37672): complete subsection reference.

- [vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-ae9ac0183af821b0ff3f4415a04dce95210f82748604d82c9e8d5969b57e656a): complete subsection reference.

<a id="canonical-0cafe1f8e1824a1ea496ce4fb28de53d09fabc86b6a978fc3bb3b6c8354764ab"></a>

## Next pages — advertise_custom.advertise_where / 4453f49eea54 / 6

- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-ef5df0371c7ad3ca8c46d55ebad67737d80fb496da4f360658a6a192928cee9f)
- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-4f3ab94e8105fc20a214b4b2b3e713e5272ddad0732a95b2c0d1884c421e4401)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0fa479eb77e8015e4314425133e9de52900c9651371dd972a4d6d80921d2325f)
- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-9d0dd20a65b48538b919d73be80414f37a7dd471f2f9f734bccf5c9f8e1bd6c3)
- [advertise_custom.advertise_where.use_default_port](resources--udp_loadbalancer--reference--group-001.md#canonical-eb3f098e7fd6914708208eaa0785e5cd86821259629b73ea12bd40abc691fbde)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304)
- [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-ddbcc812edfb98f39df6d4addb365f523cd76e2b2577a5c2710ea9b81f3d6e3c)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-7cf1e5410ce18045e6b57379d4afc1000d647abffe63f2e242400e3ba2d37672)
- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-ae9ac0183af821b0ff3f4415a04dce95210f82748604d82c9e8d5969b57e656a)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-ef5df0371c7ad3ca8c46d55ebad67737d80fb496da4f360658a6a192928cee9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b5bc613b847c66dcafe5ff4b69cf2ee827bbd09920a37f34e4c23cca57d57c60"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_custom.advertise_where.advertise_dualstack_on_public / 0086b084662f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-bc44aa6f3b23dd52b7a33820c94f8a88db93921549accbd9fa143c065fd59b30"></a>

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

<a id="canonical-4cbad62db61f2926ce8fed57687091e0a458afe1970add797e04230d4c68f70d"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public / 0086b084662f / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-762ee1ab22a718fd4616ce85eec46c009e4848ca67e0d561618ebb62ec6a5195): complete subsection reference.

<a id="canonical-bda30d9e60ef55683e2bc0e344d680f6165b4da7206355ee6a7d3b4ab74cda7a"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public / 0086b084662f / 4

- [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-762ee1ab22a718fd4616ce85eec46c009e4848ca67e0d561618ebb62ec6a5195)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-762ee1ab22a718fd4616ce85eec46c009e4848ca67e0d561618ebb62ec6a5195"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f6abb6739d63309752662927d3f1470c810b7b8e2736426aa1c1ae14148b0ab"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / f9e3e199e51a / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-ef5df0371c7ad3ca8c46d55ebad67737d80fb496da4f360658a6a192928cee9f)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-d69c9c9702b270c6ec58c29c66dc3de94d58647827f9073c1d155e06b9309499"></a>

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

<a id="canonical-a06285d428f3b81f0586c6fcd336920a2867129be07727bf9d386dfff01540c5"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / f9e3e199e51a / 3

<a id="canonical-47a60c9bccbdfd556e7cedbaa6c4b1a8a6168b0e536d4e27d5c904c28085be40"></a>

<a id="canonical-859111dc2d09d9d2143cecbd0fbc5ccf2bdeef9314d5600153c167752a5b7151"></a>

## name property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / f9e3e199e51a / 4

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

<a id="canonical-ebf10bd3b1a9ad8e372d8bd949368a2426c110e6a8ad9fe6ee93baf26166b70b"></a>

<a id="canonical-07072da45f28fce68457c6e1e0f30e223da5ff991c585f9104e4d15c1cec60c0"></a>

## namespace property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / f9e3e199e51a / 5

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

<a id="canonical-2c52e5807fc57b287d25531022d223f1dda71f34fcdc4b42ad3403e1e4ae1be4"></a>

<a id="canonical-20b001a4b480d44fd55793de2090bcfc3eb2075ba73e928d2309c273e3bc9397"></a>

## tenant property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / f9e3e199e51a / 6

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

<a id="canonical-8bbd1da179fdb3b40992917ae52136eb6e5fb9e721015a980bf90065887f7c82"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / f9e3e199e51a / 7

- [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-ef5df0371c7ad3ca8c46d55ebad67737d80fb496da4f360658a6a192928cee9f)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-4f3ab94e8105fc20a214b4b2b3e713e5272ddad0732a95b2c0d1884c421e4401"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002965e029f8e53b34a31c2dbc160a6fe8df570905e74a6e46f0eefdf2911fe"></a>

## advertise_custom.advertise_where.advertise_on_public — advertise_custom.advertise_where.advertise_on_public / 6becffe619b9 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-9d8b2fd06792979d3938a9fc54edf1b717690a57af7ebcfa7def4b17e97d6bad"></a>

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

<a id="canonical-1fc728a9b1bbab59b7ab2cc1a5b9c0dc7543adb21e935512cb4ccdbe8765d542"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public / 6becffe619b9 / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-21add006f08af2863a4f460922397a20c9f8992527e9fda117e667d23f51c843): complete subsection reference.

<a id="canonical-b736b13d0bb726ec35682bc510b7c91e7b94afe28f7f976a91da9d72a5048425"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public / 6becffe619b9 / 4

- [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-21add006f08af2863a4f460922397a20c9f8992527e9fda117e667d23f51c843)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-21add006f08af2863a4f460922397a20c9f8992527e9fda117e667d23f51c843"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a3621619e62f1f4c32bd2dfed6400bd233a79899b1500e2dfde34c3beeb6a80a"></a>

## advertise_custom.advertise_where.advertise_on_public.public_ip — advertise_custom.advertise_where.advertise_on_public.public_ip / c25692a7f0ef / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-4f3ab94e8105fc20a214b4b2b3e713e5272ddad0732a95b2c0d1884c421e4401)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-d1ee059fb92e782d3a520ec2253bfbfdc6cdb1946d988cb58a1237b909786fcd"></a>

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

<a id="canonical-950416260b6a674efb20be8bc00774b0ed8269c3f32843ed68c17cf0caadb637"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public.public_ip / c25692a7f0ef / 3

<a id="canonical-a7b43af290b022987fd527928ac3d3739ab69763af1db9a655f12aea57b2c4d4"></a>

<a id="canonical-d4edf8291b47a90d0b442e8b6b6b984fff77199fe55dd9763679c49695512124"></a>

## name property — advertise_custom.advertise_where.advertise_on_public.public_ip / c25692a7f0ef / 4

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

<a id="canonical-452c1b32091c3bf58e8fc5d70b927f3eefd81b4d073385cd1c88296a27f71bb3"></a>

<a id="canonical-ee5289ec5ea1b7c6364a270cc596c9f937a82d23cdb864396202da2707e98686"></a>

## namespace property — advertise_custom.advertise_where.advertise_on_public.public_ip / c25692a7f0ef / 5

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

<a id="canonical-714699e3442fc83b930a620b0a271b13949c9903de2ca9d6487af18bd6ffadf8"></a>

<a id="canonical-db7186be5d36a900cf80b7a01eaaa6f830e50e974f8c7de8d40fadbf36e9fe33"></a>

## tenant property — advertise_custom.advertise_where.advertise_on_public.public_ip / c25692a7f0ef / 6

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

<a id="canonical-19a2965e0bcf873bad8270cf518dfea3a7a685c6e488690ba0c02af3ea69df35"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public.public_ip / c25692a7f0ef / 7

- [advertise_custom.advertise_where.advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-4f3ab94e8105fc20a214b4b2b3e713e5272ddad0732a95b2c0d1884c421e4401)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-0fa479eb77e8015e4314425133e9de52900c9651371dd972a4d6d80921d2325f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a09f499058884e36f0ae8c1b57d22ef6f728ee6368d353c6ca736164a14d188a"></a>

## advertise_custom.advertise_where.advertise_v6_on_public — advertise_custom.advertise_where.advertise_v6_on_public / fc5dc4654826 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-1af50843d119f55a9cd9d32e2bcec8ac8ed64877fb636d79cff75de6965eeb67"></a>

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

<a id="canonical-b8efbec06d5a1dc97aeaa32dad7b82d763612a8a03fd6623cc1c12b54ce6d0ab"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public / fc5dc4654826 / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-b7619e1bb8631e2823926fda4faa7161e437de36104e8667ccb7abce2e87e430): complete subsection reference.

<a id="canonical-4221e1e65770a35b645ee149efaac80aea0748ca3d22cddc297320b03e515a9e"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public / fc5dc4654826 / 4

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-b7619e1bb8631e2823926fda4faa7161e437de36104e8667ccb7abce2e87e430)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-b7619e1bb8631e2823926fda4faa7161e437de36104e8667ccb7abce2e87e430"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a97af9b2da4648ea351727bd5b37a0a3a2edc6a53492e778faa3fd2a7cf99ea2"></a>

## advertise_custom.advertise_where.advertise_v6_on_public.public_ip — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 6c734c46f03e / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0fa479eb77e8015e4314425133e9de52900c9651371dd972a4d6d80921d2325f)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-f13f02f77f01392ca7892a74d7fa26c9b5499cf724f45b89dc297cd4af792f8c"></a>

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

<a id="canonical-55bf8ea4212292140b5d94c14407d8efba402868977183552ce00c8fc401e656"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 6c734c46f03e / 3

<a id="canonical-cb808c2ed5e5a072f4a23b97bf8f97f3c646f86dc51b7469d8fa6b26e4a682fc"></a>

<a id="canonical-1ba1033400aa41d08f4e70e05befdb0866c5fa7abe7360b802417ed3224af629"></a>

## name property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 6c734c46f03e / 4

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

<a id="canonical-e40b301a3d30d88a86001d728961fdb032eb0c0f4505f9e970b1a6c71b119062"></a>

<a id="canonical-d2beed9f7d24718890e75c48b3561352d5f6b2d812dd0e3e3df6eb3f17153c0b"></a>

## namespace property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 6c734c46f03e / 5

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

<a id="canonical-fbb86b04b4bf2c638187c531325723a2d54a61db7b6e4f5969bf67d4aee529b8"></a>

<a id="canonical-46e48aa48463edc5524db6e1db92e60a2109420c7b9d8542b5433248b8eaf034"></a>

## tenant property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 6c734c46f03e / 6

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

<a id="canonical-79a7aafb5158944c8ebcfe1d33cd3a43618467414bf290d30d3a40e5b437b904"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / 6c734c46f03e / 7

- [advertise_custom.advertise_where.advertise_v6_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-0fa479eb77e8015e4314425133e9de52900c9651371dd972a4d6d80921d2325f)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-9d0dd20a65b48538b919d73be80414f37a7dd471f2f9f734bccf5c9f8e1bd6c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28f4e31103516cda42c48f38256619d3ceff7f6000f1d3625530b1171e07cf51"></a>

## advertise_custom.advertise_where.site — advertise_custom.advertise_where.site / 1cea7b83347f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.site

<a id="canonical-087933082ce9f9fdb8367f31c0036b17194eb612c096469c1cecc3510ae48f26"></a>

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

<a id="canonical-25d8bd88e4e37ec83aeccf372143317206f803981ed741f53f1dcea3e64024a6"></a>

## Direct properties — advertise_custom.advertise_where.site / 1cea7b83347f / 3

<a id="canonical-01da621462d0e361850ac4b81aec078360e70adbc97d25665e935619aa2c3b13"></a>

<a id="canonical-51e157d33e76853e44716f5566ac17cc90425ca79a731cfdde8a1d2071a28b6d"></a>

## ip property — advertise_custom.advertise_where.site / 1cea7b83347f / 4

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

<a id="canonical-b031d8aeef0895ebdbcbae822e2cc485b51df3f86c0caf1108f0355653e22a40"></a>

<a id="canonical-5aeb8fd63b7f0ab70d63ffec979dd3a152b3a75384f3ff480ea1bf0534853a5e"></a>

## network property — advertise_custom.advertise_where.site / 1cea7b83347f / 5

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

- [site](resources--udp_loadbalancer--reference--group-001.md#canonical-a6a76d34118c52cffe026a2144877fd2f3ca3bc1e6464c95111a85778227f582): complete subsection reference.

<a id="canonical-333f9f7fc844dc88fa9fdcebff9ad0a8517f7226ccfa662e29cf136831fedd77"></a>

## Next pages — advertise_custom.advertise_where.site / 1cea7b83347f / 6

- [advertise_custom.advertise_where.site.site](resources--udp_loadbalancer--reference--group-001.md#canonical-a6a76d34118c52cffe026a2144877fd2f3ca3bc1e6464c95111a85778227f582)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-a6a76d34118c52cffe026a2144877fd2f3ca3bc1e6464c95111a85778227f582"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3817fcc06c1f3d4dd66457fdcfe7a43c73e1a787a56ae1eb6c29b542db69dd7d"></a>

## advertise_custom.advertise_where.site.site — advertise_custom.advertise_where.site.site / 914a5b4f7c96 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-9d0dd20a65b48538b919d73be80414f37a7dd471f2f9f734bccf5c9f8e1bd6c3)
- advertise_custom.advertise_where.site.site

<a id="canonical-5cd2fafe560f75e305a9d689c68674aea7f737d011b74ee0ba8a3754e19cb0f6"></a>

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

<a id="canonical-02ec5e2e3db830adedcf721b1128e48cd29b561eb2d5b4b0ecfc4e617df2f55b"></a>

## Direct properties — advertise_custom.advertise_where.site.site / 914a5b4f7c96 / 3

<a id="canonical-9d1b0615cdf758bbd16c799490f02250c8f9c8fb63008ca5a8715c954f7c1c02"></a>

<a id="canonical-c0153927d610ab5537c2678ee7028c75ea99e716bccf115a9fcf6eaece13dc7f"></a>

## name property — advertise_custom.advertise_where.site.site / 914a5b4f7c96 / 4

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

<a id="canonical-8ccd4a8ebae674bdd5c347b5bfea0e9b3ffde761025562d116e4da13394e8141"></a>

<a id="canonical-a6e99e5c022c903a99a4d446997609f261dc5d0910e64e1a854c8877b4a29d26"></a>

## namespace property — advertise_custom.advertise_where.site.site / 914a5b4f7c96 / 5

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

<a id="canonical-385246fafc94dac23c7e9550d8d95b48f89d16143aa6481445a5a37632a83d16"></a>

<a id="canonical-f9019d000cd863686c26222ac22ccb4349b8ffb5052e4414258ead83a0ee45cb"></a>

## tenant property — advertise_custom.advertise_where.site.site / 914a5b4f7c96 / 6

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

<a id="canonical-2819389fa68c6a42d93ff6b21571b1616c099f6dd41282e19a2e1c5b9feda2ee"></a>

## Next pages — advertise_custom.advertise_where.site.site / 914a5b4f7c96 / 7

- [advertise_custom.advertise_where.site](resources--udp_loadbalancer--reference--group-001.md#canonical-9d0dd20a65b48538b919d73be80414f37a7dd471f2f9f734bccf5c9f8e1bd6c3)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-eb3f098e7fd6914708208eaa0785e5cd86821259629b73ea12bd40abc691fbde"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6565b5369bea977d7a4a2b97bd998da1993f4a418788196142c9912f214d517"></a>

## advertise_custom.advertise_where.use_default_port — advertise_custom.advertise_where.use_default_port / 45b22959118f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-30ee0b30d377ad49fc04e36995d701d738ee9531d8a7a6b68f45854619843f93"></a>

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

<a id="canonical-8a60123094692fefcfa197487d67c4688a18fda9e4cc041c28b6d3cc04701485"></a>

## Direct properties — advertise_custom.advertise_where.use_default_port / 45b22959118f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ecf60bb96e0fd5e832623712904ecdb5e029b303df480e60177304d6cd3acef2"></a>

## Next pages — advertise_custom.advertise_where.use_default_port / 45b22959118f / 4

- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-210b39967f2000ca7cfe8997e19c7ffcf7bdf6cc4fdd588693da86dab1e64f6b"></a>

## advertise_custom.advertise_where.virtual_network — advertise_custom.advertise_where.virtual_network / be65d21a13b8 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-b197266511d523fff7766c382c9f74c0b1628e671bcb1a1ea88e6a7de802de1e"></a>

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

<a id="canonical-0425caf2c4586827a8b5cc94aa6e130530d864f92ec0e7635b29054383f08dba"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network / be65d21a13b8 / 3

- [default_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-494c0bebf3f58ef77b5ab7e9c7b6f193a02b44b4cd2fc6fc3d77df034e1bb161): complete subsection reference.

- [default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-220b8b5536202c54580cb99a447274413606328a70c268361ec8c758410db2cb): complete subsection reference.

<a id="canonical-1f4aa6d96c2589c808a8fe087c88a28550953d12a8b5803b0fbc7a35fa072651"></a>

<a id="canonical-6afa78e29ab5e98a859784c5af8cbfd1db37ce7881e0ad72bd534e800e5fa567"></a>

## specific_v6_vip property — advertise_custom.advertise_where.virtual_network / be65d21a13b8 / 4

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

<a id="canonical-3e8ad8fc338247453a9610b9358d1c6108ac59182835ef1f9da590a818387bb3"></a>

<a id="canonical-0bf3b9073d94141948213bfd4be1f9abf7a1dba891da0c60e5dba0918ebebf99"></a>

## specific_vip property — advertise_custom.advertise_where.virtual_network / be65d21a13b8 / 5

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

- [virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-715fb0e47677f21516ec1f3729140196d6bb0e2bf2fb887ee4d67b2918b1f373): complete subsection reference.

<a id="canonical-1bc3e721fb7b334d1956ec65f645ca828a6b82c6178d30e71692ea99413e523f"></a>

## Next pages — advertise_custom.advertise_where.virtual_network / be65d21a13b8 / 6

- [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-494c0bebf3f58ef77b5ab7e9c7b6f193a02b44b4cd2fc6fc3d77df034e1bb161)
- [advertise_custom.advertise_where.virtual_network.default_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-220b8b5536202c54580cb99a447274413606328a70c268361ec8c758410db2cb)
- [advertise_custom.advertise_where.virtual_network.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-715fb0e47677f21516ec1f3729140196d6bb0e2bf2fb887ee4d67b2918b1f373)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-494c0bebf3f58ef77b5ab7e9c7b6f193a02b44b4cd2fc6fc3d77df034e1bb161"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36adf4785c56ceac0d137d9db2306841439e601b07c403daf3ff653f50247ebb"></a>

## advertise_custom.advertise_where.virtual_network.default_v6_vip — advertise_custom.advertise_where.virtual_network.default_v6_vip / 295fb5e45cf2 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-8c96ad2057d0d236ae55c2b912ed433934a1e72939b45921db51318d3aaa114d"></a>

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

<a id="canonical-b7f1b8ba1d93a3e62f0982505e1f1338f515786515f0661a78b412b231b85c7c"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_v6_vip / 295fb5e45cf2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14a90608b87b71c7395f74e0110b2f1a71a2782f960a716a932a04ad81033e67"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_v6_vip / 295fb5e45cf2 / 4

- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-220b8b5536202c54580cb99a447274413606328a70c268361ec8c758410db2cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95b3d996c5fa3eeeb819b20ec08c9a7382fb44a703dce033fba7fdb8b042a441"></a>

## advertise_custom.advertise_where.virtual_network.default_vip — advertise_custom.advertise_where.virtual_network.default_vip / 5f1e67bd8e87 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-2b2c6738d31f2ffa2c155bec602a44ed4567bda95f962ba26271d36a79c4b9db"></a>

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

<a id="canonical-f8fdf77713bc5520e5086715f787780a2f4bc0150309ba1af8b47ec8d271d652"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_vip / 5f1e67bd8e87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ea5761533b77f823453e1e929eba48e6800e89b8bcfd50d14accb3e53aaed5ec"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_vip / 5f1e67bd8e87 / 4

- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-715fb0e47677f21516ec1f3729140196d6bb0e2bf2fb887ee4d67b2918b1f373"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94216bd850b19804468ef874a168504de7f4297448d941126369827586b43d59"></a>

## advertise_custom.advertise_where.virtual_network.virtual_network — advertise_custom.advertise_where.virtual_network.virtual_network / 9b7ff8743e39 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-9a8d30efe1c256543cf154fa889f8ef881e8ea541a694f84803396fb38b6d9c7"></a>

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

<a id="canonical-e6f0dba3decadc9a4e5c7cf7a2518cf38b13e58c7557956d0e604f6f5eddf923"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.virtual_network / 9b7ff8743e39 / 3

<a id="canonical-9868c37ffce5eef41edd5b8a601a1cd8bf1cffee05b55e7a49c332d7e40f1a19"></a>

<a id="canonical-4f5e2dac970bb710ff831c630cb4ee13da93769ce36f46dde136a7c53227d403"></a>

## name property — advertise_custom.advertise_where.virtual_network.virtual_network / 9b7ff8743e39 / 4

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

<a id="canonical-86036fad9b815d40bd0e4f27bdaa038ffb587d587ce1da179c2edcec1b0f312e"></a>

<a id="canonical-09a4a14d948813946314c3fb97429a5399a9adc29f22b464a901fe30ec358846"></a>

## namespace property — advertise_custom.advertise_where.virtual_network.virtual_network / 9b7ff8743e39 / 5

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

<a id="canonical-aad25099bdac93d88dda5f4825df3b52850d15a727659cd964ad986a7c9a3ec7"></a>

<a id="canonical-b4878d3095edb87b45650f9a0123f98cc0f5e8fc158769c3c2363c60301a5edf"></a>

## tenant property — advertise_custom.advertise_where.virtual_network.virtual_network / 9b7ff8743e39 / 6

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

<a id="canonical-460264395a152dafbbdd1774422a526de9368930ddb6b511580786b318e34d3d"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.virtual_network / 9b7ff8743e39 / 7

- [advertise_custom.advertise_where.virtual_network](resources--udp_loadbalancer--reference--group-001.md#canonical-abc33554665382bbeb704dac9067d249b1a829626e910034b76491a4548c3304)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-ddbcc812edfb98f39df6d4addb365f523cd76e2b2577a5c2710ea9b81f3d6e3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-716afdcfff7ef9854b58ed61ee8650904b72a9275d031f1287cff1f036e1ebff"></a>

## advertise_custom.advertise_where.virtual_site — advertise_custom.advertise_where.virtual_site / 49e7413e4346 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-ee20549408fd07bfee157d4ca81f5e135ecdfe71b360d450f0aaf4bbc99396cd"></a>

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

<a id="canonical-3a347f28e8540f8f0b3f6437ec40f3c61d21c2a2f9b2aba27c0ec0a365c3177e"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site / 49e7413e4346 / 3

<a id="canonical-686a723f88835140b30d2e874e9ff49bd915431519a34bb78037d4c43a8f7bc3"></a>

<a id="canonical-49432b245ea65903dd3a13672827402f24b09627cc6f751c8a5c26cf9c7a6ac0"></a>

## network property — advertise_custom.advertise_where.virtual_site / 49e7413e4346 / 4

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

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-157235765ae69f0b218a23667f9fc2e68e19c6803a770b66e17b4b45c71b7d46): complete subsection reference.

<a id="canonical-1c6b8e5188a2dc5aae168d5f25b14406ff87ed8131bb255055215b192c60f229"></a>

## Next pages — advertise_custom.advertise_where.virtual_site / 49e7413e4346 / 5

- [advertise_custom.advertise_where.virtual_site.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-157235765ae69f0b218a23667f9fc2e68e19c6803a770b66e17b4b45c71b7d46)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-157235765ae69f0b218a23667f9fc2e68e19c6803a770b66e17b4b45c71b7d46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5a3e2786df95926cf4088393c7e8b8cf328a7ecec2cfe865c49446b14b6069b"></a>

## advertise_custom.advertise_where.virtual_site.virtual_site — advertise_custom.advertise_where.virtual_site.virtual_site / 12f4fc7209e6 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-ddbcc812edfb98f39df6d4addb365f523cd76e2b2577a5c2710ea9b81f3d6e3c)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-7d6006293ce5dc9c34fb9656cba0d089c7d860221b644e5fe8ed97672458f4f5"></a>

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

<a id="canonical-11abf3d852143e7116d28b5ddca9a53ff1d6a9bf20bfad3b585a93e49833d124"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site.virtual_site / 12f4fc7209e6 / 3

<a id="canonical-36fb4da37969c4cb4dfb470fa3d63e61c58fbe52145d3a857675a3789045fa8f"></a>

<a id="canonical-b86e5e8897d8cdb797f5b739ff4d9db5ab5825e09818c55d5094103c86b57a31"></a>

## name property — advertise_custom.advertise_where.virtual_site.virtual_site / 12f4fc7209e6 / 4

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

<a id="canonical-bf2d459d240333c9ad2695218184589fb622521d996d3b6f681b8fb8d256da69"></a>

<a id="canonical-a25bfd8bdf83a8b51335e977a5e91eec7d9c77636ff07d562cd01e9f1cd22a85"></a>

## namespace property — advertise_custom.advertise_where.virtual_site.virtual_site / 12f4fc7209e6 / 5

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

<a id="canonical-bf5a481af3163c5c5398719f65cfa480433c6ce614ed650976340b848c14102c"></a>

<a id="canonical-42103b46a9ebe5a8a94a03b5dece40e92fbe5c62b6239c1c37af76439831a7ad"></a>

## tenant property — advertise_custom.advertise_where.virtual_site.virtual_site / 12f4fc7209e6 / 6

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

<a id="canonical-c48d8e94f9d1fa017218a45fc80104a327807e7f8b566e84660e64ca7bb4232c"></a>

## Next pages — advertise_custom.advertise_where.virtual_site.virtual_site / 12f4fc7209e6 / 7

- [advertise_custom.advertise_where.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-ddbcc812edfb98f39df6d4addb365f523cd76e2b2577a5c2710ea9b81f3d6e3c)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-7cf1e5410ce18045e6b57379d4afc1000d647abffe63f2e242400e3ba2d37672"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aab2a6887e1713ca2e4937107a2bcfb6d28be38761a48d6a9e6604f491530049"></a>

## advertise_custom.advertise_where.virtual_site_with_vip — advertise_custom.advertise_where.virtual_site_with_vip / b778007e5cc6 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-731f931eeead98ad065e78c4a0072ec1fae7d25ea9e6464ff0a4890a079d4ba4"></a>

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

<a id="canonical-650f69b7ad27c73a89fb72a68e9f6f43f0d12bc8119d2c8de0d44b60546f69ab"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip / b778007e5cc6 / 3

<a id="canonical-7c668402686f4d83785ac447e8a2429458a0f676939f42426d52f4c2414dd89e"></a>

<a id="canonical-3449ba488960ca030ac7f05a0d509cbd35f8855f06e6bb2c9fb67a352893f48a"></a>

## ip property — advertise_custom.advertise_where.virtual_site_with_vip / b778007e5cc6 / 4

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

<a id="canonical-15556c2a5e8c5b5a811d80cdc5a469058e2dbc9e34bebdff7b55b9461813502e"></a>

<a id="canonical-961183e5092fe07233b265e883f47b72112855e70153bfa62a420c5cd3a9c55d"></a>

## network property — advertise_custom.advertise_where.virtual_site_with_vip / b778007e5cc6 / 5

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

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-dbcaf92dc4a2fb5f6982f7b73ff4a2dd6d26506fe75b7bcf360881e5d4276167): complete subsection reference.

<a id="canonical-540d876b5ea8684df339c6375ba25fa68e7deaf672b170096d099f6a4ecd59b6"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip / b778007e5cc6 / 6

- [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-dbcaf92dc4a2fb5f6982f7b73ff4a2dd6d26506fe75b7bcf360881e5d4276167)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-dbcaf92dc4a2fb5f6982f7b73ff4a2dd6d26506fe75b7bcf360881e5d4276167"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b2d1a642ba0e84a399e9464b78797085e3fca8350ef67565e4d870af2e62533"></a>

## advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 287692539b10 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-7cf1e5410ce18045e6b57379d4afc1000d647abffe63f2e242400e3ba2d37672)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-63f2f38529cd0cbb4776f2dc6875307a4180ee8d1a06e989d94ebef04bfb695e"></a>

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

<a id="canonical-694113637772d0c38890f2866cf6c7538a94586d1d62e27465d9663b47f9bb89"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 287692539b10 / 3

<a id="canonical-e31a7893fc339175a2e93c1746ff1061c9c5ada39028b8b7c2be500ec34670c9"></a>

<a id="canonical-bffa05324c301dac42da4c74e368334316f094f5c3f52ef171966c4462558896"></a>

## name property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 287692539b10 / 4

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

<a id="canonical-4505e5f1ff10700f9c65dc2bbc7def2bdd83e3a51dd0ae43b429ded4885c7af5"></a>

<a id="canonical-14cad4f19a03b508f571df52017263efc3440be336838581c4d85eb5f6c29749"></a>

## namespace property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 287692539b10 / 5

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

<a id="canonical-28d9b9af6e8ac5b9b3b66110ff31345d4f5d70b0cfa974c6c6b37e18253898d5"></a>

<a id="canonical-5d35388151cb14f9140600213bfba20710d331382e1be85115174e9ba0461d58"></a>

## tenant property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 287692539b10 / 6

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

<a id="canonical-68bcb5e7097a273b740213da7c374a2980d5863701256a16922040ac87f0f5d5"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 287692539b10 / 7

- [advertise_custom.advertise_where.virtual_site_with_vip](resources--udp_loadbalancer--reference--group-001.md#canonical-7cf1e5410ce18045e6b57379d4afc1000d647abffe63f2e242400e3ba2d37672)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-ae9ac0183af821b0ff3f4415a04dce95210f82748604d82c9e8d5969b57e656a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9a9b271b67e78bd113f2b7eb0ba8881542392127b46c17dd78c5ff4f25193b8"></a>

## advertise_custom.advertise_where.vk8s_service — advertise_custom.advertise_where.vk8s_service / 4b4de8458c68 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-1e8d43cf6b357f82089e8d3a29ba02baa628ee4254ea292b128afa3645bcdb3b"></a>

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

<a id="canonical-7c06271716fdd3518050f29f6d8d9c5c386a9a0c4f05d56f31d931fd7b4eeb7c"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service / 4b4de8458c68 / 3

- [site](resources--udp_loadbalancer--reference--group-001.md#canonical-565f14a26c5777b43a7a6cac5c259f8b2a3ae38b52f3192eec993f480cfc44b9): complete subsection reference.

- [virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-ac029de327d1ac8812bce8f74e9d707a6e1eda3066289c6c67155b0fa21c6f5f): complete subsection reference.

<a id="canonical-76d0dadf19cd766d56dbbef939255cb0317f2ca00719749401c9b93514ec1ad1"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service / 4b4de8458c68 / 4

- [advertise_custom.advertise_where.vk8s_service.site](resources--udp_loadbalancer--reference--group-001.md#canonical-565f14a26c5777b43a7a6cac5c259f8b2a3ae38b52f3192eec993f480cfc44b9)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--udp_loadbalancer--reference--group-001.md#canonical-ac029de327d1ac8812bce8f74e9d707a6e1eda3066289c6c67155b0fa21c6f5f)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-565f14a26c5777b43a7a6cac5c259f8b2a3ae38b52f3192eec993f480cfc44b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-21369f2b95ba1d54e8995934a54f64d761c23b8423f3eea3c6506edc3231ebb6"></a>

## advertise_custom.advertise_where.vk8s_service.site — advertise_custom.advertise_where.vk8s_service.site / 22801582a885 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-ae9ac0183af821b0ff3f4415a04dce95210f82748604d82c9e8d5969b57e656a)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-54045ef5b28bcbd40c65169e8eb27d704c13af87080dcdb579427935c3f1ac80"></a>

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

<a id="canonical-4406cbc7460fd15c427bdeff73e1c44935f8f97d8d089b441b43063fba5fe4f8"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.site / 22801582a885 / 3

<a id="canonical-b3870faa062c1a4c978f9bb9edb2d9851a8e7c8e75ea02e9f468cda7c822ba26"></a>

<a id="canonical-85b1b7484aa85ffb5c138ff00681369cedadfa25980f69203455ffd23eb1511c"></a>

## name property — advertise_custom.advertise_where.vk8s_service.site / 22801582a885 / 4

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

<a id="canonical-05e7c6a7620c511837ff89c172aebb075f433c778ab82bf28541c926417f099b"></a>

<a id="canonical-ee0ae4d38d50e61a6d6487daa120e9da8ac6a7a3324c4bde8be251c860fe3dab"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.site / 22801582a885 / 5

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

<a id="canonical-5d616a4e147216f755e7f836445e57d80e8c6f59c31e91db4687adce42cb3784"></a>

<a id="canonical-fa7d8d1771f792407056c0568930d677fc8189acca49c4ccd94ca6090ada70db"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.site / 22801582a885 / 6

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

<a id="canonical-a8d8c1c88bd624df47421d1a2604a5fa63f5eeb4bfc5651a80da6ee131d2be35"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.site / 22801582a885 / 7

- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-ae9ac0183af821b0ff3f4415a04dce95210f82748604d82c9e8d5969b57e656a)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-ac029de327d1ac8812bce8f74e9d707a6e1eda3066289c6c67155b0fa21c6f5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1577a38b0275da37d36485ee84aca027906ba2b221912e9f41d60f45f7f8fb3"></a>

## advertise_custom.advertise_where.vk8s_service.virtual_site — advertise_custom.advertise_where.vk8s_service.virtual_site / f4791e192cfa / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_custom](resources--udp_loadbalancer--reference--group-001.md#canonical-fdadf0137afa30b95af93ba1a52a0698274a36b12ea7605d8691d0f026fbe90c)
- [advertise_custom.advertise_where](resources--udp_loadbalancer--reference--group-001.md#canonical-22066adbf841c7c6b5e723de8b3033692f89f464924ab02e71c86559b354e911)
- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-ae9ac0183af821b0ff3f4415a04dce95210f82748604d82c9e8d5969b57e656a)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-f132e12661fffa7ff4ac51d3ad46819b209e91c5e7fb37e1a5f15cb7861c7580"></a>

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

<a id="canonical-2a7f14b35f56f764b18ed24e683b3a94329bff87ee9bd242bf401e55aa23b001"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.virtual_site / f4791e192cfa / 3

<a id="canonical-22269fa3639ab7b1235d71b7d7990b9bae5151f0db509525fb7ad2197ce60167"></a>

<a id="canonical-991002c46a4d7e46a3c37e343652b611336ccfcc4e59fd6a21c8ed479c67a1bb"></a>

## name property — advertise_custom.advertise_where.vk8s_service.virtual_site / f4791e192cfa / 4

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

<a id="canonical-574469de3e2d2e8824e14ae54c07d6ced972a3e471aa99332523da55e331de4f"></a>

<a id="canonical-fb6e43dabf1e7a555d954cde036ff6f6ecbeb70927357361379b9e4b75cafbfd"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.virtual_site / f4791e192cfa / 5

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

<a id="canonical-1d006e3d5f6880bc2fe5f1ba618e4df1a41446d04159b942fbeec07d006a28c8"></a>

<a id="canonical-53870915fc199f07c629e08bf26623c3de06ed9469e4f5625bb74c2e313aa824"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.virtual_site / f4791e192cfa / 6

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

<a id="canonical-3e92f8621316a1412c08ae50adf7e125ebb21c20deaf77554f911d7ad3684891"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.virtual_site / f4791e192cfa / 7

- [advertise_custom.advertise_where.vk8s_service](resources--udp_loadbalancer--reference--group-001.md#canonical-ae9ac0183af821b0ff3f4415a04dce95210f82748604d82c9e8d5969b57e656a)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-d9064ae10247a9cba717e2df2b773fae22b197233ca5a3e1704fe32571228c61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00ca102ec8f1f5e8607d5275db5ae0477e054d882853feaf74c4874f66bb51ec"></a>

## advertise_on_public — advertise_on_public / cd2e0cfdaf9c / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- advertise_on_public

<a id="canonical-5df43900bd5c0407ae764c67495ace806a9f36c132e913c1047635999ba07957"></a>

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

<a id="canonical-8efb5f57bd07082d541425fc0f96af7d989ec00e4f8d02661e810e79f6110e62"></a>

## Direct properties — advertise_on_public / cd2e0cfdaf9c / 3

- [public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-10967a62dcf2ab4f4aeb8370e6de8d7e06f68db619e2e56f3d3fbf5a7798a813): complete subsection reference.

<a id="canonical-f85f2c27f18240f7392739d11b1fe3e0c043b45325c5d5f0351b1decf9905e2a"></a>

## Next pages — advertise_on_public / cd2e0cfdaf9c / 4

- [advertise_on_public.public_ip](resources--udp_loadbalancer--reference--group-001.md#canonical-10967a62dcf2ab4f4aeb8370e6de8d7e06f68db619e2e56f3d3fbf5a7798a813)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-10967a62dcf2ab4f4aeb8370e6de8d7e06f68db619e2e56f3d3fbf5a7798a813"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d724e6add60db3b12b38aad4381cc3d4ffe5495c599f3266045bdae3223e66f3"></a>

## advertise_on_public.public_ip — advertise_on_public.public_ip / a23e1ab084c6 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-d9064ae10247a9cba717e2df2b773fae22b197233ca5a3e1704fe32571228c61)
- advertise_on_public.public_ip

<a id="canonical-304dae4bc07ba465aca88a5f06379964b033d4f8a1342908bcc69b34e66e27cd"></a>

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

<a id="canonical-d144c2884fcef49a1a3d5c4b77c773a0b525a2fcab38528073bcbf38ef68578a"></a>

## Direct properties — advertise_on_public.public_ip / a23e1ab084c6 / 3

<a id="canonical-155193a3be4db577501d5764072ecb4ade5bb6967e92507909e61d958b025dbe"></a>

<a id="canonical-908808133143c700911d69f3c552158cc108851f2c7d793649cb648ca6dcd28e"></a>

## name property — advertise_on_public.public_ip / a23e1ab084c6 / 4

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

<a id="canonical-a3a7dc7a1fc425670f065ab0f8a7a6fdff4cb2cbc6ff22d7fc5f0300798a8b9c"></a>

<a id="canonical-3446929817b4ce40db3d7da8f91bd190e8c775ecccc0b12dae66291f94734ac4"></a>

## namespace property — advertise_on_public.public_ip / a23e1ab084c6 / 5

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

<a id="canonical-fa7f95dd43823742da5a3f66cba715a310b98a770f2a7a1d5292ce9e59b3727a"></a>

<a id="canonical-64d0d3268e1fbdaac378a4ff48dab9782031535c9b37b43b3654afddaa8e2ce8"></a>

## tenant property — advertise_on_public.public_ip / a23e1ab084c6 / 6

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

<a id="canonical-ba05d4046d00fcd4b111a35be58ec6357d037a4fc0964f2a8a728f9ff5ae6402"></a>

## Next pages — advertise_on_public.public_ip / a23e1ab084c6 / 7

- [advertise_on_public](resources--udp_loadbalancer--reference--group-001.md#canonical-d9064ae10247a9cba717e2df2b773fae22b197233ca5a3e1704fe32571228c61)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-4e7cc4e708953afbd0abe22443f62ca3ff474f436045a846be4d4c8abf34e593"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01f5cb4151969c20ccc67493ba0058ab0fff37617a09daf261935a30b436de8f"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / 1c4db6a9a8e9 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- advertise_on_public_default_vip

<a id="canonical-b2edb84e8196bc951ea0fdc2c9c63576398749833bf9a7b65df2ae1fb47f806f"></a>

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

<a id="canonical-5e329c5c9df7709035f67f5847ebe5a9c3b8308814247b2ea4768980aced065c"></a>

## Direct properties — advertise_on_public_default_vip / 1c4db6a9a8e9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cbc058e9b49b2c7747b905b4750b442114ce7bb2f3effd6c8cebb6471d778622"></a>

## Next pages — advertise_on_public_default_vip / 1c4db6a9a8e9 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-8918004147bc82a6741b40ac27a78269f7d85ad32ff9788081e710294eaf0a38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8ebd6b5a7852a890c64253cb3efea174034597d447fc1fb40513eafac6abd22"></a>

## do_not_advertise — do_not_advertise / 07d756529414 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- do_not_advertise

<a id="canonical-299da560f186457acdaa0655c3e4327ae8c6b1c124ac88ca3bc4d486defa93a0"></a>

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

<a id="canonical-7bb063895cf1d4340809d50ce339743a2f8f255d8ea80d9c325db528d16a46b9"></a>

## Direct properties — do_not_advertise / 07d756529414 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c8832808cdf9cc829ef7636c44ea6e5653d56a65000dffdb119ae58db4fb5922"></a>

## Next pages — do_not_advertise / 07d756529414 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-9c238b54b4fa70dc93c8d69891fd3b0a930acc5e1179fe65435c9f2b5408d0b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4294ff2f71fee45cde8c3c6b2ba8f798deb151cf4143842b5ed70473cf4be3bf"></a>

## hash_policy_choice_random — hash_policy_choice_random / 85998b88440f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- hash_policy_choice_random

<a id="canonical-2593b25f47b2d6b0ca6e599ea8124322a0ccee145c7be102a199d9af68f8ab58"></a>

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

- [hash_policy_choice_random](resources--udp_loadbalancer--reference--group-001.md#canonical-2593b25f47b2d6b0ca6e599ea8124322a0ccee145c7be102a199d9af68f8ab58)
- [hash_policy_choice_round_robin](resources--udp_loadbalancer--reference--group-001.md#canonical-e30839a403148e9818b5f0e75b5b07813437b93c330ffd670030f0ecf904a448)
- [hash_policy_choice_source_ip_stickiness](resources--udp_loadbalancer--reference--group-001.md#canonical-f6e867f0958edb0b92159ec6a85b755c0ffd272b2fcdb4f18faeda97ccac0e43)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hash_policy_choice_random = {}
```

<a id="canonical-0d91d9c31bb589792c1e28291062da8176b6ddda4b1d5109c4e4b0fcb0e7529f"></a>

## Direct properties — hash_policy_choice_random / 85998b88440f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6c1aba16b99124dc185c3dfe805ff519e458c4201c15961901ce69c4281d4d5e"></a>

## Next pages — hash_policy_choice_random / 85998b88440f / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-0f00a843733878244fa6b6b74122b7e929fc15c072750ee11524745d72f84199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a59c8ae6158173e7016de071908f7db633721ce6c363bcf3bb4bdfe11a6a0ff9"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / a9dc3df671e2 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- hash_policy_choice_round_robin

<a id="canonical-e30839a403148e9818b5f0e75b5b07813437b93c330ffd670030f0ecf904a448"></a>

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

<a id="canonical-b7b89eb4fafe4fd6d22f3bb7c29514f8210d96acc114d96c25a0828d83f869d7"></a>

## Direct properties — hash_policy_choice_round_robin / a9dc3df671e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b90202ac5e31e98a470befe8045bc78f03629ef2344c202a3284ea655649ead2"></a>

## Next pages — hash_policy_choice_round_robin / a9dc3df671e2 / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-6e8bcc7eafb9a3b0f56ec6102957cacd26fdf582b8894df2a649dad99c1500a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22a1b0e014959898c120279348d09b1c9046992e83d42214eb83afd1bcafd803"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / bdfbef7bc31f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-f6e867f0958edb0b92159ec6a85b755c0ffd272b2fcdb4f18faeda97ccac0e43"></a>

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

<a id="canonical-602ae7f89e89c53a9ddc8fe19bea26da9638bc8083cc71fab20f4e00b5a6ba99"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / bdfbef7bc31f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb58f11ba1c3914085dcdad86e354669372b312dd293760e8e01b32478b7c9fd"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / bdfbef7bc31f / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-aad9f2a76c2fc2765eedd5ecbff8af5a9c083c5f8e7df90fd0dc82310dbeeb06"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-adf1cdf49a84187c8589c76c08aacaffda52bda9012cced82733f071fb68d9ed"></a>

## no_service_policies — no_service_policies / f7ac9e5b091d / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- no_service_policies

<a id="canonical-8566e3ab18cd76803d80747be5b443cbf6c666d10ba37595089ac4457560b3c8"></a>

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

<a id="canonical-ecec62c8e313fa783572a6515911e9b0def7bceacfb77a0fd5434e6601d65ba9"></a>

## Direct properties — no_service_policies / f7ac9e5b091d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33acb890d6273bae79748dad4367bb7abfa1ef1c0baec6c661743ce552eb1643"></a>

## Next pages — no_service_policies / f7ac9e5b091d / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afe9e7c0a246baa1bbe5be1caa9d64d0c85f156dc3e5a1669241123443fcadb8"></a>

## origin_pools_weights — origin_pools_weights / ad9688a34f1f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- origin_pools_weights

<a id="canonical-b14d6230706c3a14278b27880febcca55d99e25c7ffa7a54ef7360cc73f054da"></a>

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

<a id="canonical-a59b9b144326db05462099b5f4c3849ad25fe4368132ac322d61307b0651d026"></a>

## Direct properties — origin_pools_weights / ad9688a34f1f / 3

- [cluster](resources--udp_loadbalancer--reference--group-001.md#canonical-908e6f891b8253f335b593047358f870185290cf9a15588ca3297629740f9db1): complete subsection reference.

- [endpoint_subsets](resources--udp_loadbalancer--reference--group-001.md#canonical-a6f7b0d54cdd820ee619b65ff0bd0e82dc3d9782454688b32069f96e8909aa3d): complete subsection reference.

- [pool](resources--udp_loadbalancer--reference--group-001.md#canonical-3a9f99e2e252d51fb905d871be4f605f5b1e51abb2e5cbc038ffaf2707d00b8a): complete subsection reference.

<a id="canonical-a84bfe8ea1ac276c0c9bc4c13b2eff6648aea542c8bb8834bfd780f031d59ca8"></a>

<a id="canonical-75ea9a34c15b61810f4a1dd769d1e5e5b20a53cb63e0993d4c1a7f704916e501"></a>

## priority property — origin_pools_weights / ad9688a34f1f / 4

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-23ffdac21e5043473e49b28ae73a5efd26df2373025739857daaf7025102b30b"></a>

<a id="canonical-21d040b99049cc8e6f3e979c7fd2e274f15468848cd67eb241970f33240f196a"></a>

## weight property — origin_pools_weights / ad9688a34f1f / 5

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
  }
}
```

<a id="canonical-a01416b53617fbe9b3e3e63b9fd880ea35fad86114ddec7b8eff3ff97b48a199"></a>

## Next pages — origin_pools_weights / ad9688a34f1f / 6

- [origin_pools_weights.cluster](resources--udp_loadbalancer--reference--group-001.md#canonical-908e6f891b8253f335b593047358f870185290cf9a15588ca3297629740f9db1)
- [origin_pools_weights.endpoint_subsets](resources--udp_loadbalancer--reference--group-001.md#canonical-a6f7b0d54cdd820ee619b65ff0bd0e82dc3d9782454688b32069f96e8909aa3d)
- [origin_pools_weights.pool](resources--udp_loadbalancer--reference--group-001.md#canonical-3a9f99e2e252d51fb905d871be4f605f5b1e51abb2e5cbc038ffaf2707d00b8a)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-908e6f891b8253f335b593047358f870185290cf9a15588ca3297629740f9db1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c38d7d85d03641943d25c25e43c6053dfbea68c5008cae9a7fe4d5bb93cb6602"></a>

## origin_pools_weights.cluster — origin_pools_weights.cluster / 755e55bb4af4 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3)
- origin_pools_weights.cluster

<a id="canonical-8ef1005996684b05ccca852b9fc42884c6ec032b7df08150052ccf912ac80896"></a>

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

<a id="canonical-cf74c717f7b545a63e640e54f12707d62e32acdeca35707aed7c9e250b439710"></a>

## Direct properties — origin_pools_weights.cluster / 755e55bb4af4 / 3

<a id="canonical-80fc60542d098491f30d99047ee570d290f9f88337c6760473f5363ad87af240"></a>

<a id="canonical-b748307cd7adbb8d762c60b8f7be105960950f8ac12b3f44e43fec0bc872d5fe"></a>

## name property — origin_pools_weights.cluster / 755e55bb4af4 / 4

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

<a id="canonical-85070bf77b9bf460549faf438832d453d36ed360d715c3df3735353432378cc9"></a>

<a id="canonical-723a5bd32717e8d371d12ecf1aec4cfc602ce58eb2f66ad30b3a8d208c9e008f"></a>

## namespace property — origin_pools_weights.cluster / 755e55bb4af4 / 5

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

<a id="canonical-dac5fabfa657efc4c1544b44228c914698820db4f48feb064f63b5ef47442187"></a>

<a id="canonical-37671199dd9bbad6d7d1b15f49970cc6e8aabce6d950ec811aeeb1e0ed087fdc"></a>

## tenant property — origin_pools_weights.cluster / 755e55bb4af4 / 6

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

<a id="canonical-4d941d80dbed99702a41727ee3c3a23687de532c3af65a1a51ff775b251ebc5c"></a>

## Next pages — origin_pools_weights.cluster / 755e55bb4af4 / 7

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-a6f7b0d54cdd820ee619b65ff0bd0e82dc3d9782454688b32069f96e8909aa3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3715fa2002c9229e882f3a86fb07012640b098522383e30f73c710a041bd44fa"></a>

## origin_pools_weights.endpoint_subsets — origin_pools_weights.endpoint_subsets / ecc08bc4227a / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3)
- origin_pools_weights.endpoint_subsets

<a id="canonical-9ed9b555c3725b19f079d2a7ff5df90ea183612b3901b3348b4913fedb2e7a61"></a>

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

<a id="canonical-f71648643e0565871682488b179476cbae9c6b41b206b62e56e8f26eac321d81"></a>

## Direct properties — origin_pools_weights.endpoint_subsets / ecc08bc4227a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7debfb4660f6f101a96ca3add7b7ce7892487250b567facad2e661d2bd1b44b9"></a>

## Next pages — origin_pools_weights.endpoint_subsets / ecc08bc4227a / 4

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-3a9f99e2e252d51fb905d871be4f605f5b1e51abb2e5cbc038ffaf2707d00b8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bc7bcead971c94d3695d77e729d9b39c0b2adce2feb0ec75bc2f37d3bae0e35"></a>

## origin_pools_weights.pool — origin_pools_weights.pool / de7bb8587e41 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3)
- origin_pools_weights.pool

<a id="canonical-e637954e7b12d3ce4ca460cfcc383d949e4f65cf24e16e5f354a7b16bb0f7bc3"></a>

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
pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1e51a8ed7c8855a42d50ccfa441ee2c11d70f7b03149b3229e7446041342103e"></a>

## Direct properties — origin_pools_weights.pool / de7bb8587e41 / 3

<a id="canonical-126c664776d70d2aa1466cac801afa3646248b2beb89922bf54e02635ba80b80"></a>

<a id="canonical-517b68b660d740a6b741a360790ac7fe448819440997e632511ed74b78b0b592"></a>

## name property — origin_pools_weights.pool / de7bb8587e41 / 4

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

<a id="canonical-a7105bbd4e782fd8369614c8a913b550345a04a19b8a1d09415dd216e616738c"></a>

<a id="canonical-83757877e584f773d0302ff14400bd45ee05d18d01febeceb0a4980a4c821159"></a>

## namespace property — origin_pools_weights.pool / de7bb8587e41 / 5

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

<a id="canonical-e46becef0ad629da3719fa60187cb22d5c26c35c3bc6fcd895b2933204d627b9"></a>

<a id="canonical-f58c4aa005287663ae2a1a35f3a950cb26b1ae12e5cab79dad162585534456cf"></a>

## tenant property — origin_pools_weights.pool / de7bb8587e41 / 6

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

<a id="canonical-f2ae236cdaead5e9a216dd92c10b37e53a73e8c2b440539afb4bc607026ac6e7"></a>

## Next pages — origin_pools_weights.pool / de7bb8587e41 / 7

- [origin_pools_weights](resources--udp_loadbalancer--reference--group-001.md#canonical-0e9f7431e615647c45ff9e40bf3f4fcdb94af95c814aeaa54930bde98b378cf3)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-a943aedfea42ed379f11cd0ded88b44ef75a05109c7e9807a51bbfe5039c7cec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dca1be072ae16116211b7f12fe8f8a25bf0b6d8db88c6124592af1ff7904f55c"></a>

## service_policies_from_namespace — service_policies_from_namespace / d1883e0bed5d / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- service_policies_from_namespace

<a id="canonical-af52f868a3f0414fa7c4768e484c8f217966c71e5fc2ff84c9059cfd65031785"></a>

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
service_policies_from_namespace = {}
```

<a id="canonical-ff385428baef949ecc472ed47b45d252497411cefb43bc92d855e8dd962a8bfb"></a>

## Direct properties — service_policies_from_namespace / d1883e0bed5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-206fe9f7a8f0d09fcf7130af5a2f8c3b1924059b3d6ef00c8b0d470cb1e33b1f"></a>

## Next pages — service_policies_from_namespace / d1883e0bed5d / 4

- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)

<a id="canonical-6a1cd6fb1eca53775784d20ffcccb4f95e864f18ca39406d848757a84e4f7c95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f09bd847bf928bb0ec9d66452c0e3d431dfa38e1fa51aede18712b65a5a9b54"></a>

## timeouts — timeouts / 6a0f045370d9 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md#canonical-8d0941bd2b64991cda54db4f129d03cdf063ebe72f0f5cee355179e865ff8ae6)
- [Property reference](resources--udp_loadbalancer--reference--group-001.md#canonical-822e7540266823efd2a77ae1ad30d26e70b6ac855b61834d36929910865f2513)
- timeouts

<a id="canonical-8929daabbf9810669cfa00dd08c809be7cedd7783505382a2a8f23673cf25421"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-da825c07ba0bc5ded3584f896952ecee86feddce9370216934e1052039e9b37d"></a>

## Direct properties — timeouts / 6a0f045370d9 / 3

<a id="canonical-c04b66d9aa1cd13aef64e8af5415ee96e7c864d2ec3c69d27e5c074ad680c161"></a>
