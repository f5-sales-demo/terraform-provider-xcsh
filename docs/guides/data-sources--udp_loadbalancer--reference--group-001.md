---
page_title: "xcsh_udp_loadbalancer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_udp_loadbalancer reference."
---

# xcsh_udp_loadbalancer reference

<a id="canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1eddada5782c93e6d16946e06048d9a91356000abac7f1f5da999090b7634a0"></a>

## Property reference — Property reference / d7a98c1a20ba / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- Property reference

<a id="canonical-64a09306eff3c01b3ce4048bd482e06dd17a011f7e077bd264700ccd48e44dc1"></a>

## Direct properties — Property reference / d7a98c1a20ba / 3

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c67928ccdf37b2be6248c0729166e66a6c9bc509dfb276eed7a585f9c80b4c6d): complete subsection reference.

- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996): complete subsection reference.

- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-d5168f4eb73ed5b138b2244fd19ddf6629ee4860ec7e3750c5cc7580f308ee5d): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b5eb4ceb2d20a15434aee06f351d8acc20cbc0d0bc6f34c5d029c3c0967674d2): complete subsection reference.

<a id="canonical-f4b4fcdcabe558220b0bb646419b0c8accc6299d2db1b9bff31e536a7f8a0cb9"></a>

<a id="canonical-3f6bfc1ea2f4b77fc478ffce81912aa112fc15456dedc5c044fd7d78574f78ce"></a>

## annotations property — Property reference / d7a98c1a20ba / 4

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

<a id="canonical-c0dc9afe681dd12fac2c4b747fffce3ff2b82f91a3818d4ade484bc4d356fd33"></a>

<a id="canonical-32561fe6af60d49fc6ab7795e16a82ec1144176355429788b790327979569d78"></a>

## description property — Property reference / d7a98c1a20ba / 5

Type: `"string"`. Computed.

Description of the UDPLoadBalancer.

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

<a id="canonical-fb3cfa3a8dfcf0a53d0b030e32930c8184ee9b52aaf3f0d9ee413ac6638f3de0"></a>

<a id="canonical-cbe7e55f704bd6ba629853a4f258bf73ba53bd688824628b40657e43db93b51e"></a>

## dns_volterra_managed property — Property reference / d7a98c1a20ba / 6

Type: `"bool"`. Computed.

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

- [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1d03dc3a13c57995a5526df4c20ac937d35490f9430db921da4213bc2adecd0d): complete subsection reference.

<a id="canonical-68786cb93861309e99a837ca62905af0e0a580a8aaa6ca04c946701d7551c224"></a>

<a id="canonical-43dfd9fa236468cb1e3dbcaf167f2b20b313c64699dc9f0860dededa28f55a28"></a>

## domains property — Property reference / d7a98c1a20ba / 7

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to this load balancer.

Upstream description:

A list of domains (host/authority header) that will be matched to this load balancer.

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

- [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c2840de39be22d8c0af335282ccaa658e1b6de712ef708eb0abf925c02f6cec9): complete subsection reference.

- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-6bba225388bef94021eb7f1db16412ca896935d5b8cb4fe6c704d1acbf25f44d): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-368e636030383f2acb10c46c6725dc4705a2c8d9c936ff5c381954f25fcb7478): complete subsection reference.

<a id="canonical-e7d4b4d277687ded6aa1994fff07bb2844394f810e3523ab1eba09346cb4b5b7"></a>

<a id="canonical-161c464fe2f261e0b6a200b5bff97c130a7642796d3ef2e44e1d8d03fbe9c527"></a>

## id property — Property reference / d7a98c1a20ba / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-059c24e994f390ca4d78bf0ad702adce2882a1bc5f1eff6bf52636bc0d7bc110"></a>

<a id="canonical-1fd529f991cf4288cfd8741a2de7a3fd4f175f12996c811c97c27cb0435dcd74"></a>

## idle_timeout property — Property reference / d7a98c1a20ba / 9

Type: `"number"`. Computed.

The amount of time that a session can exist without upstream or downstream activity, in
milliseconds.

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

<a id="canonical-4aca2ab2367900e6bf75f83e2c49a4d82dec7c7375e283d66a7f44f5e5d9c4ae"></a>

<a id="canonical-c930f8b2a8d24838423bf82a63226fe7fb4eb5490eba802838dd05bbf73e87a3"></a>

## labels property — Property reference / d7a98c1a20ba / 10

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

<a id="canonical-21975c9d0f9f1794dcd4647257d4bb9aec2d9d4c7eab0bd8bfccc757eb94bdc4"></a>

<a id="canonical-f766ffc98f02b07f4888d6fbd6fdc3b495eb0f46d9f0fbc7a9c8e534db4562e7"></a>

## listen_port property — Property reference / d7a98c1a20ba / 11

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

- [listen_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-21975c9d0f9f1794dcd4647257d4bb9aec2d9d4c7eab0bd8bfccc757eb94bdc4)
- [port_ranges](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8ec3a344a016946ce7bbd03a4e2b63768a6b1bd9c0c32d69fa3cb29ed3a2b953)

Select alternatives according to the provider validators above.

<a id="canonical-002525a6a956d35b7236aa1d05ad9c572dd174d0814e6cd92e76abc257992e77"></a>

<a id="canonical-b25fc8cd7aa2b42feaf73fc10eeb29c08ece526fad633b3ff59df5b25098662f"></a>

## name property — Property reference / d7a98c1a20ba / 12

Type: `"string"`. Required.

Name of the UDPLoadBalancer.

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

<a id="canonical-ab5f772cd88dac1f9453a39ee948e32c525c68a8865ee6bc2bc7951ceb9b21be"></a>

<a id="canonical-dfe217622948f6531e3c3439e20ec0bb20e936aead9cecb07849f705a95faeda"></a>

## namespace property — Property reference / d7a98c1a20ba / 13

Type: `"string"`. Required.

Namespace where the UDPLoadBalancer exists.

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

- [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-35483bfc197f023ade1e8abae3d8af09cfeb45b6bb50280c18c7f65bc558fdc7): complete subsection reference.

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7): complete subsection reference.

<a id="canonical-8ec3a344a016946ce7bbd03a4e2b63768a6b1bd9c0c32d69fa3cb29ed3a2b953"></a>

<a id="canonical-19c8549516433283a5621e8b71ef0aaebb89ff2722ba56f8662098938c59027f"></a>

## port_ranges property — Property reference / d7a98c1a20ba / 14

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

- [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-dab7f01b3657f3fbbb46b538eb5e89792a2ec50653869106249bf0f8da5d6f30): complete subsection reference.

- [udp](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8029559a5e7fea91d6d135845a81c92c829ceeb7cc054314b011e12dbaa7f04): complete subsection reference.

<a id="canonical-db857935d0ba5f100c4be384dd39288658816cf373b3d3bc94e6963f68407425"></a>

## All schema paths — Property reference / d7a98c1a20ba / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-caf9b4a4473b8bae6d7eb2c41db6bd8b2c98b1e6413b775aaf2bbb4698c43efe) |
| `active_service_policies.policies` | [active_service_policies.policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9583b5000524c6044d85f04941f9cf65af633f0e769bc301be7675eee20b6c58) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-f8550174813f6d9066c8fa3ed4aa08762e8b79576d7dfbf859d10d10f382785c) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8e56224506d0caeee521abb806a79b4727aa401b0c4b609a3739eddcde0b630d) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e316021f9c9a13978bd3f57560aa899f75c656c8474b34a592f25987a1120590) |
| `advertise_custom` | [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-506f8eaf16db1ae8b3ead1f9d75f4aa09f14bb68b1f48424f7c36cf989be5f06) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-faf3fbdf719b040f4b10d009c973fd49a1d30a8b8391ace2252d062d255cfd05) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4ce67cca439be75a7300f5d68bddbad1f5c93064d64837676989fa10fd0af013) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-5c224400e84610431aaea5582e6930699c9b5bc8411561baf562fe636c34e8e6) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b989129c613a395a81a91bb35c9feb80501ff0d0c752eb4295e0699d5960f9f6) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-482ef59869d83580c3b2563b56f3e8bf0a5e5beb0e27b636d780cb01dabd927b) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a7afc0f0b5fb6d043c34b1b76f019b58fda86ddfa65a60944a6dd1d9a81ee7d0) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0fe312d8c0486274d5f5fcf50ab649e0e150ed289e394f37e41f6fae1639a207) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9455b0257df2587ca2cd3d300e042fda4de80f1145a41b4fd39be29dcf63bbf8) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3b28e0ee9bf9b85896bda9d66e1ff4bf107dc3e37568e70658f0570ae5f04fc4) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-dd2cc4011ca57810329a81eca22cac46ede669c7a8e25ecc3f1ac11b6ca2e756) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-39c3d0e238fb36452414908d84cb1c63d0344e434672deec375f964374b21256) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a85d0c7aab8101a2d44faa6944a407c23f7946cfd92c94b3b23c4af583cb41fe) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c8b5c66303bf4ab2081cd2d9a14169c4121bcd290d3083d372332ff3bb91cce1) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-716c7b953b2965fd882f2e39505804a147054dc4f0eb1433b246e5aef66c4b9d) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-5391fe3fc312740930ae527b66932f2ce09c2f6730c683be0f8f034c51631005) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c7afb070e861f24e24fde2314ef7e688c1dad4b435e4ccef3fa9da0240442220) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-968b47a2c29dfd7debb7c4fefe78230db99dad50e2ba05f2595e6b577a717d44) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](data-sources--udp_loadbalancer--reference--group-001.md#canonical-fb797370e81dd6315c30ec71d739ababeb1368cd824261b5e7fab2f0f96a4ad3) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a6a767152e50372944f6de21a0e75602d8009b909870479faa17f3c389939a75) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-151c7295000af1ecf54150885354127eaf45b63102584a33f99f5939fcd8a5d5) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-00670fcdd99e2b5e504a42ab3cc72007d1ea822d0222acda96d8e26f54476190) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-5fa4c80c00d47473caf0eb226338b27a7fcff2fa991af3b199b551f491321b58) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-0d891dc7f9ad0f501ecb44470295aa7bb42e5b47685adbb8ae88b43bb661d23f) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-d755fcae0515f1bdc0b4a2707677d11ea7888ff7194fa54977f0ef0aa221a3a8) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e76c979073aef8ffb99e1b6b73170bfae135ebaa61f7aaa4aa9323b99c761e85) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b9f35391cf0784b50ced0479868e70948548efe842f8bc2f78a2bfc933c57929) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-31836349c924f89e675716f34d5f42ead3ebed533cdab0e33f54cd1fcd7a374b) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-835fed5803e45014eb4c8c1bb8801a466a9b572cc1fbb60801d577b06716bfe5) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-04a279ca07672ed414935c889bef6e242c57715ccc78822e53d5ae56a9976a60) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-130c37c8c4f7eaacbc31c1e814c07dd49d5ec0938b3f2aa1d02b0dd905922840) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e68a6a3bf18bea8e0ad7f368f18df7d804d2120f6ff4d09a40400d9c6b832c17) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9061ea967f620202fb25cbd7213955ee0b3295650a7ba67b99169d8dc941245a) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9d025d1bcbf7ccdc3ccf22825d6dac10f6747b87f39c17c8663e9eb374e664eb) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1a6aae53cf1311614a143fc99966f4ee9a3165b3e638aff1fcd2dcef1a290bba) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-481171f25461d22d7164155eebc8ed6940f86052c272a71fe3c037d7c962b6f3) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-5a8ca5e15664763e038f0f526e2b56e7e20a52960b200c0719a31ef5b2354d35) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-676d70f085372b0fdb3265dc8496fef569cb3fac4968c723de8e4d95aa45e094) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9a8de5046def0ec912287ffd1ec237bc657f7ba14c1ee78399fee822f73121a3) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a659eb84d48f65479b1e2d503eb43e901ae8b65335e1e002d3b98aded41b8dc3) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-ceb8ac6867b0ef3c55c7da0e338de562e22741327898cfa3c514da683852de70) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9e277b29b5c5be2cae0cd1ceded003d6ab9107c70351ab225a6f22e96f42b5d7) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3c53f59d598c79a2bcb2f33cd5a1793bf88760f624f7cf4d46902de7aeaf85c2) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-feff7f533ada708a19a126f0a5e62dfd61c2f5beb504ed93baa862d088f9fdb3) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-6e57f8554d7880d622ad53f725628ecd270bbacd5c2d88396bd089dade0c0b40) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-6f2b01ab333f3a00f812f3f802d0f282dc3b549063f6ec18ab7a75522e354c91) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8addf50e48c39d140b12682f4be0910e3c783cad7290c6c2252a6322e65a4d25) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-ac82e971d56147a27516d9a6709c88c720a415979507165e90f965590e9f51bb) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-556126d348fb84ff868bb66195e9798b16a0c77b8e743b5e6eb77882f19da675) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-921d43a2c1ba075ff09cd88ac3ed763ddc7ad1a1ac40d2c19c5e5177f5b15f01) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-5d2eef8b53e72630869469a965b36323f7a2d97df2a11a444b091d86bba08348) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b30da8e784cf85e36fecb72a0d997df77b7bf0dd3e5cc21da9ba55f1649acb19) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7048066f756fd3b10897c5ca0ab3b7ce41955cd2df61ed4bc98f700a41b336f5) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-84ae926b200245d21db2ed4df47764716d5e234d68534e974d8dabcd6ed0c55e) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7dbef697841fa637133361fa8df3a355ed2a30b04b0c8aa773c0d75587c10a1c) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-074aaffd2c44f9cf4a7299bc825ffe5214c74f54f9fdca2f835571938fee6b62) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-40fc8df6244644653bfd24ef4785affd6ad758b72fd66707fb0f64e92afda191) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e96ed6787985651542d3fc82df91006f7fd4012e6310a0b75b24fee7908b54f7) |
| `advertise_on_public` | [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c2496f692c3248fafa7fe8560a189583638c1dd790eedbe6b21d479d733fc33f) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8bd7453d768e7c56d262312aa8577e43e29ecbf919b826b095d9df14d470c3d6) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c200698e083393d32b4cf3517a74e8ebb4ff4a33650c3d2c7f990b1a2b427f99) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-02d8bf12de40b94df80dc596fb188e05c7709a710166d16d39259336f88654b3) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1cdbb02b44c067dc65fe3cea8e7cb70184fbf29ab8d91926704fe9126f13e885) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b4e5a59b1bb9de65f78677156c820e8ccb2d0f5bf1ae81bab4d58c8cc8b61530) |
| `annotations` | [annotations](data-sources--udp_loadbalancer--reference--group-001.md#canonical-f4b4fcdcabe558220b0bb646419b0c8accc6299d2db1b9bff31e536a7f8a0cb9) |
| `description` | [description](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c0dc9afe681dd12fac2c4b747fffce3ff2b82f91a3818d4ade484bc4d356fd33) |
| `dns_volterra_managed` | [dns_volterra_managed](data-sources--udp_loadbalancer--reference--group-001.md#canonical-fb3cfa3a8dfcf0a53d0b030e32930c8184ee9b52aaf3f0d9ee413ac6638f3de0) |
| `do_not_advertise` | [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-116708bf49f53a7a81817599cb6d0dc09f3b35a1d6e478e3707e2114ef59b978) |
| `domains` | [domains](data-sources--udp_loadbalancer--reference--group-001.md#canonical-68786cb93861309e99a837ca62905af0e0a580a8aaa6ca04c946701d7551c224) |
| `hash_policy_choice_random` | [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-5aad3fcb6f7df3486afc79a36520af1935ac44953101b27ae00bb6225f15952d) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-f1b328288612454845b907f088f36b22dfb6dfe6149b7d87cb47eff5cc86c29c) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-aa986a91e9d0c055ff991dc2e46dc133575d35241eb58f59a88dcb021a3152dc) |
| `id` | [id](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e7d4b4d277687ded6aa1994fff07bb2844394f810e3523ab1eba09346cb4b5b7) |
| `idle_timeout` | [idle_timeout](data-sources--udp_loadbalancer--reference--group-001.md#canonical-059c24e994f390ca4d78bf0ad702adce2882a1bc5f1eff6bf52636bc0d7bc110) |
| `labels` | [labels](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4aca2ab2367900e6bf75f83e2c49a4d82dec7c7375e283d66a7f44f5e5d9c4ae) |
| `listen_port` | [listen_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-21975c9d0f9f1794dcd4647257d4bb9aec2d9d4c7eab0bd8bfccc757eb94bdc4) |
| `name` | [name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-002525a6a956d35b7236aa1d05ad9c572dd174d0814e6cd92e76abc257992e77) |
| `namespace` | [namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-ab5f772cd88dac1f9453a39ee948e32c525c68a8865ee6bc2bc7951ceb9b21be) |
| `no_service_policies` | [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4e9d8ef93f9f96f932e3cff70e2315a8c75729e4edf984b897414ce8685b5bcc) |
| `origin_pools_weights` | [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1f1421826b7784862b3833f2d8009928e5328c84734716d083c36a26aab3f4ac) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4d965e320d766842274f7fa134e5eb11ac95c939f5792aeaa52b8d8e23b09624) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-ed8f03d1111e1a46c84c43fe60eb65c4c4deb07d4fd7749636607933e55dfa4b) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-ecaad0099c4e8333eb9f997a8179621919a974e075dd7750526fae76cda75aa5) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-cce9166994a1fca32a9b2ebc4065181aa0ced8f7afbd2c8a166ca31f587be4b8) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1478454b614689f916ca494f161c6d8dcd0c48afcbb85b703fb05f74a7f8cb13) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7b7d183a67cb28564e5559d3f1a480927bf9bfe5739a0e54bc8395daacbf6dd3) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](data-sources--udp_loadbalancer--reference--group-001.md#canonical-6c4a7229e6883f3812c972007736773a359ccfd88853e74ec622b49738233980) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b9828aeafe43a2dc4d37e1b2fca2dc4ff24eb14f66644f904fed51d0b80c6067) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4df4327d1247875528a3a25566644d6495033a90ce15633788b566bdd71d351c) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](data-sources--udp_loadbalancer--reference--group-001.md#canonical-307e125d127cade86eedfd22e471f5dd543a64565ae18fb9772e4f69e00d2542) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c9e02595c0a17cfe44ecbad89769058109eff9eb96510c45150979339ad68f6b) |
| `port_ranges` | [port_ranges](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8ec3a344a016946ce7bbd03a4e2b63768a6b1bd9c0c32d69fa3cb29ed3a2b953) |
| `service_policies_from_namespace` | [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4dbc966db9d1d33246273ef3fba34f087674f9f47eaa6b07c12b74440c4ee878) |
| `udp` | [udp](data-sources--udp_loadbalancer--reference--group-001.md#canonical-3abd8ed3ddf2ea410ad301d36d0a459dcb11f8639b54fb7c7c534d5cd5850a84) |

<a id="canonical-82261e92772810b3959f0b2f7ed03f5f938adc42e876cc8c8babac224b43c3a4"></a>

## Next pages — Property reference / d7a98c1a20ba / 16

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c67928ccdf37b2be6248c0729166e66a6c9bc509dfb276eed7a585f9c80b4c6d)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-d5168f4eb73ed5b138b2244fd19ddf6629ee4860ec7e3750c5cc7580f308ee5d)
- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b5eb4ceb2d20a15434aee06f351d8acc20cbc0d0bc6f34c5d029c3c0967674d2)
- [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1d03dc3a13c57995a5526df4c20ac937d35490f9430db921da4213bc2adecd0d)
- [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c2840de39be22d8c0af335282ccaa658e1b6de712ef708eb0abf925c02f6cec9)
- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-6bba225388bef94021eb7f1db16412ca896935d5b8cb4fe6c704d1acbf25f44d)
- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-368e636030383f2acb10c46c6725dc4705a2c8d9c936ff5c381954f25fcb7478)
- [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-35483bfc197f023ade1e8abae3d8af09cfeb45b6bb50280c18c7f65bc558fdc7)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7)
- [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-dab7f01b3657f3fbbb46b538eb5e89792a2ec50653869106249bf0f8da5d6f30)
- [udp](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8029559a5e7fea91d6d135845a81c92c829ceeb7cc054314b011e12dbaa7f04)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-c67928ccdf37b2be6248c0729166e66a6c9bc509dfb276eed7a585f9c80b4c6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbf94109b58f91553979d497bbfb371b83fb16586fd7976c5e5f3a5a092fc208"></a>

## active_service_policies — active_service_policies / cae54c717955 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- active_service_policies

<a id="canonical-caf9b4a4473b8bae6d7eb2c41db6bd8b2c98b1e6413b775aaf2bbb4698c43efe"></a>

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

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-caf9b4a4473b8bae6d7eb2c41db6bd8b2c98b1e6413b775aaf2bbb4698c43efe)
- [no_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4e9d8ef93f9f96f932e3cff70e2315a8c75729e4edf984b897414ce8685b5bcc)
- [service_policies_from_namespace](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4dbc966db9d1d33246273ef3fba34f087674f9f47eaa6b07c12b74440c4ee878)

Select alternatives according to the provider validators above.

<a id="canonical-8123fb1965de4d90019222763a9cf753aa7b86d66367be00c078cc23ffd82232"></a>

## Direct properties — active_service_policies / cae54c717955 / 3

- [policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-494a0fbd48bb1aa147883076c56f7b310b471b0a0ee703b59ad5046786173263): complete subsection reference.

<a id="canonical-481368094f15d483e273d47f5987386bcc4c377c260e798a4b3e937929bcd290"></a>

## Next pages — active_service_policies / cae54c717955 / 4

- [active_service_policies.policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-494a0fbd48bb1aa147883076c56f7b310b471b0a0ee703b59ad5046786173263)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-494a0fbd48bb1aa147883076c56f7b310b471b0a0ee703b59ad5046786173263"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d317d0e9d91bb2d0f08765032981553c48c08efa187b5f878a1d4ebbf507c9ff"></a>

## active_service_policies.policies — active_service_policies.policies / 2d04f80bf098 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c67928ccdf37b2be6248c0729166e66a6c9bc509dfb276eed7a585f9c80b4c6d)
- active_service_policies.policies

<a id="canonical-9583b5000524c6044d85f04941f9cf65af633f0e769bc301be7675eee20b6c58"></a>

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

<a id="canonical-03832caa9157fc754f58a53cd393211dd0ab1a6d3ef3241bc09f03906c819ea8"></a>

## Direct properties — active_service_policies.policies / 2d04f80bf098 / 3

<a id="canonical-f8550174813f6d9066c8fa3ed4aa08762e8b79576d7dfbf859d10d10f382785c"></a>

<a id="canonical-24d1eb7e8e881df3e96feb6a6433f4f54ceba68f0671ae178f688217403c2d7c"></a>

## name property — active_service_policies.policies / 2d04f80bf098 / 4

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

<a id="canonical-8e56224506d0caeee521abb806a79b4727aa401b0c4b609a3739eddcde0b630d"></a>

<a id="canonical-3725cd920ec01affcf1b14aa59e0fc766cc3db7541d2ea3bcdf5fb70ec1d23bf"></a>

## namespace property — active_service_policies.policies / 2d04f80bf098 / 5

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

<a id="canonical-e316021f9c9a13978bd3f57560aa899f75c656c8474b34a592f25987a1120590"></a>

<a id="canonical-d46ebec65d63f8e761b1db94ffb23ad6ea12fd1a323d4615e5a1f582eb7f370b"></a>

## tenant property — active_service_policies.policies / 2d04f80bf098 / 6

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

<a id="canonical-943940b725c6007ab0a23e55e793dca7eafb6a19235463da534c2ebab9c442f2"></a>

## Next pages — active_service_policies.policies / 2d04f80bf098 / 7

- [active_service_policies](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c67928ccdf37b2be6248c0729166e66a6c9bc509dfb276eed7a585f9c80b4c6d)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63e86dbd543708cbb62544e3763cf18cca741c8bd4e4ca6cc9ecfd233b86cd71"></a>

## advertise_custom — advertise_custom / 345d5cfb3b2e / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- advertise_custom

<a id="canonical-506f8eaf16db1ae8b3ead1f9d75f4aa09f14bb68b1f48424f7c36cf989be5f06"></a>

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

- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-506f8eaf16db1ae8b3ead1f9d75f4aa09f14bb68b1f48424f7c36cf989be5f06)
- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c2496f692c3248fafa7fe8560a189583638c1dd790eedbe6b21d479d733fc33f)
- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b4e5a59b1bb9de65f78677156c820e8ccb2d0f5bf1ae81bab4d58c8cc8b61530)
- [do_not_advertise](data-sources--udp_loadbalancer--reference--group-001.md#canonical-116708bf49f53a7a81817599cb6d0dc09f3b35a1d6e478e3707e2114ef59b978)

Select alternatives according to the provider validators above.

<a id="canonical-8f6630cfd1e7a08fd09768986fe19d89b70d2f6d92b88927c0c167a4ce5329a7"></a>

## Direct properties — advertise_custom / 345d5cfb3b2e / 3

- [advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617): complete subsection reference.

<a id="canonical-672256c90ef8dbd36404e654f508841bbb9492dc8edb1e01213eabc0bc6f809e"></a>

## Next pages — advertise_custom / 345d5cfb3b2e / 4

- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afbbb192b9b21b0700e89390fc1fdd9f28980f305d0fadadd4688328b9ab937c"></a>

## advertise_custom.advertise_where — advertise_custom.advertise_where / 8708237a822f / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- advertise_custom.advertise_where

<a id="canonical-faf3fbdf719b040f4b10d009c973fd49a1d30a8b8391ace2252d062d255cfd05"></a>

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

<a id="canonical-58d8bf9f4bbcecdddcc3c41a20dafbdea848263cbea18e28fefd591bb0ba91e1"></a>

## Direct properties — advertise_custom.advertise_where / 8708237a822f / 3

- [advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-911e7b8d81049f4992640fdac1710660600f575a509cb398783829033c94666c): complete subsection reference.

- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-68726beb9d074b213b2866d481cf633a97dd4662305fa7e9932187ec4ea126ae): complete subsection reference.

- [advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a267550051395f0f1fa132938a6bd0320f2f3134ed5093dc08c772e25054a15d): complete subsection reference.

<a id="canonical-968b47a2c29dfd7debb7c4fefe78230db99dad50e2ba05f2595e6b577a717d44"></a>

<a id="canonical-c705c17eb63ab044f5587a028bf17e663c776906b89ec338ab145feb05e6580a"></a>

## port property — advertise_custom.advertise_where / 8708237a822f / 4

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

<a id="canonical-fb797370e81dd6315c30ec71d739ababeb1368cd824261b5e7fab2f0f96a4ad3"></a>

<a id="canonical-891f48b641d3df05019cc913cd7a3eac7f1697cf486a879b94faf517d53285e6"></a>

## port_ranges property — advertise_custom.advertise_where / 8708237a822f / 5

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

- [site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-01e4422723d7bed42f62d682fd4d27da3df14f1fc4c83321c61b4c303499373d): complete subsection reference.

- [use_default_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9068df4283088d54603361b25c9c8a15536a1a25009996e52ae1f5f220a39190): complete subsection reference.

- [virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187): complete subsection reference.

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e3744b7a41f67683e48a0745dc20d3accdeaead0c59b15491f85286673c468b5): complete subsection reference.

- [virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a6693a30231263f66d37d6765afb02b246cd9bcff2fed3757e51efc5b07eed34): complete subsection reference.

- [vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7a6d49a64fa8bd2949bc29e099a1eec058e153c2860b56e8087604ed3b7a2cf8): complete subsection reference.

<a id="canonical-dfd915f3cfc47208ef92171ff0e6b1606971cfabaabdd56a3ef680f510c834cf"></a>

## Next pages — advertise_custom.advertise_where / 8708237a822f / 6

- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-911e7b8d81049f4992640fdac1710660600f575a509cb398783829033c94666c)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-68726beb9d074b213b2866d481cf633a97dd4662305fa7e9932187ec4ea126ae)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a267550051395f0f1fa132938a6bd0320f2f3134ed5093dc08c772e25054a15d)
- [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-01e4422723d7bed42f62d682fd4d27da3df14f1fc4c83321c61b4c303499373d)
- [advertise_custom.advertise_where.use_default_port](data-sources--udp_loadbalancer--reference--group-001.md#canonical-9068df4283088d54603361b25c9c8a15536a1a25009996e52ae1f5f220a39190)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187)
- [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e3744b7a41f67683e48a0745dc20d3accdeaead0c59b15491f85286673c468b5)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a6693a30231263f66d37d6765afb02b246cd9bcff2fed3757e51efc5b07eed34)
- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7a6d49a64fa8bd2949bc29e099a1eec058e153c2860b56e8087604ed3b7a2cf8)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-911e7b8d81049f4992640fdac1710660600f575a509cb398783829033c94666c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-588b9da43a13bf276a8b67e58f9864b70544049919101cf8b054dbff11be4b36"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_custom.advertise_where.advertise_dualstack_on_public / 3344684efecc / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-4ce67cca439be75a7300f5d68bddbad1f5c93064d64837676989fa10fd0af013"></a>

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

<a id="canonical-b4e991b92199a078c454d7380dfaa0d1af29955524f852a23816e574176238ae"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public / 3344684efecc / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-152a17b69f0fadcf2ddc1a46fdc398d9fa81e6373e730887c7a3eebd4fb6ac2a): complete subsection reference.

<a id="canonical-a0332ff6fcf1a0b828350b5a8c15d35c0c95bc49e90f3752d01a9c08666057ee"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public / 3344684efecc / 4

- [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-152a17b69f0fadcf2ddc1a46fdc398d9fa81e6373e730887c7a3eebd4fb6ac2a)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-152a17b69f0fadcf2ddc1a46fdc398d9fa81e6373e730887c7a3eebd4fb6ac2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-437fb18e1c4f2652e030c307f618fe2d0eaf52d295d2247943bd0c66f2b04876"></a>

## advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / b01ba667c599 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-911e7b8d81049f4992640fdac1710660600f575a509cb398783829033c94666c)
- advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-5c224400e84610431aaea5582e6930699c9b5bc8411561baf562fe636c34e8e6"></a>

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

<a id="canonical-d77da9f37108b9ac3cbc611573bbc45d8e2e57d5a5cc8b933114012fc77e2010"></a>

## Direct properties — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / b01ba667c599 / 3

<a id="canonical-b989129c613a395a81a91bb35c9feb80501ff0d0c752eb4295e0699d5960f9f6"></a>

<a id="canonical-1191a731842e7ba18cd8bacc73aee48a4964bb77141edc933a3354d9db52551c"></a>

## name property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / b01ba667c599 / 4

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

<a id="canonical-482ef59869d83580c3b2563b56f3e8bf0a5e5beb0e27b636d780cb01dabd927b"></a>

<a id="canonical-c6e4279a84975b12ee43bad7acc3958d9452c1feb749ef594c33b30fab29bd27"></a>

## namespace property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / b01ba667c599 / 5

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

<a id="canonical-a7afc0f0b5fb6d043c34b1b76f019b58fda86ddfa65a60944a6dd1d9a81ee7d0"></a>

<a id="canonical-55c2372598c89b8e7c2b95f533a914c14dea5277ec182c98e7f096b354102852"></a>

## tenant property — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / b01ba667c599 / 6

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

<a id="canonical-31f23b65d6d8edaeef3181573c0199036067cfe9062895269cb6a2ae41c3ec9e"></a>

## Next pages — advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip / b01ba667c599 / 7

- [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-911e7b8d81049f4992640fdac1710660600f575a509cb398783829033c94666c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-68726beb9d074b213b2866d481cf633a97dd4662305fa7e9932187ec4ea126ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ee822f5e1bba2b368772b1879e674f30aeb69ed9b42ee23ed22b39c2a801269"></a>

## advertise_custom.advertise_where.advertise_on_public — advertise_custom.advertise_where.advertise_on_public / 2f27736239b6 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.advertise_on_public

<a id="canonical-0fe312d8c0486274d5f5fcf50ab649e0e150ed289e394f37e41f6fae1639a207"></a>

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

<a id="canonical-f4f3af6b855cdffdcb05b9bdc1db929d942ed53a4575d9a6338566f7e88c4f16"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public / 2f27736239b6 / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8592f1e68b1b474b9fb3654d9e8cd3a7f522810450d03037fe9d584cf6b5d5cf): complete subsection reference.

<a id="canonical-8b34369557446d3b8de37790a40c53fabee569742fe83ba620de2668e9a028e6"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public / 2f27736239b6 / 4

- [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8592f1e68b1b474b9fb3654d9e8cd3a7f522810450d03037fe9d584cf6b5d5cf)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-8592f1e68b1b474b9fb3654d9e8cd3a7f522810450d03037fe9d584cf6b5d5cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ebaf4980d22a4426f4016ae12ff7ef515f0b1b4b69cf6b897d43ae4889a4af6f"></a>

## advertise_custom.advertise_where.advertise_on_public.public_ip — advertise_custom.advertise_where.advertise_on_public.public_ip / 790e710e3e00 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-68726beb9d074b213b2866d481cf633a97dd4662305fa7e9932187ec4ea126ae)
- advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-9455b0257df2587ca2cd3d300e042fda4de80f1145a41b4fd39be29dcf63bbf8"></a>

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

<a id="canonical-bf577ac03621b5ad5671cf476c23f215b30cd048384ae57b126b50735fe7b0ab"></a>

## Direct properties — advertise_custom.advertise_where.advertise_on_public.public_ip / 790e710e3e00 / 3

<a id="canonical-3b28e0ee9bf9b85896bda9d66e1ff4bf107dc3e37568e70658f0570ae5f04fc4"></a>

<a id="canonical-07377896418595421eb6b748dea27277074d6535a0c33262ca7db6fcd0d243cd"></a>

## name property — advertise_custom.advertise_where.advertise_on_public.public_ip / 790e710e3e00 / 4

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

<a id="canonical-dd2cc4011ca57810329a81eca22cac46ede669c7a8e25ecc3f1ac11b6ca2e756"></a>

<a id="canonical-bf74db8419c0a1c4b7d6f047345bf02625619c396c55470877e737c8f2c9baa2"></a>

## namespace property — advertise_custom.advertise_where.advertise_on_public.public_ip / 790e710e3e00 / 5

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

<a id="canonical-39c3d0e238fb36452414908d84cb1c63d0344e434672deec375f964374b21256"></a>

<a id="canonical-d8de529ddd6384406059e825822cba2be4c123e9a191518adeb62d294070a061"></a>

## tenant property — advertise_custom.advertise_where.advertise_on_public.public_ip / 790e710e3e00 / 6

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

<a id="canonical-3cb559c90b0e94a3673e6273f7f72d740fafe76b0e295cffdf6c6220e36e7c3a"></a>

## Next pages — advertise_custom.advertise_where.advertise_on_public.public_ip / 790e710e3e00 / 7

- [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-68726beb9d074b213b2866d481cf633a97dd4662305fa7e9932187ec4ea126ae)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-a267550051395f0f1fa132938a6bd0320f2f3134ed5093dc08c772e25054a15d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50d3b8dd26b9ef8babac906ad05409a9dc4e8b1975250a8a6bf20cf9bd23b02c"></a>

## advertise_custom.advertise_where.advertise_v6_on_public — advertise_custom.advertise_where.advertise_v6_on_public / 1beae6c15be7 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-a85d0c7aab8101a2d44faa6944a407c23f7946cfd92c94b3b23c4af583cb41fe"></a>

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

<a id="canonical-4d02591cac1e6d73691cbe7a48f4e951899b0b7b88d15680bb505653c9dd5fca"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public / 1beae6c15be7 / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-64671f1d3a27ca89ff5e71672bb38985f7ae1b11718eab02cbcd88b68b287cc6): complete subsection reference.

<a id="canonical-3350f60c2a5e7ecddf34543f12a5a07e35a1c62cbe9563eebe2d191e70236310"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public / 1beae6c15be7 / 4

- [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-64671f1d3a27ca89ff5e71672bb38985f7ae1b11718eab02cbcd88b68b287cc6)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-64671f1d3a27ca89ff5e71672bb38985f7ae1b11718eab02cbcd88b68b287cc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29eb1595115f7f4b410e00787292e1e0f332f82096836aa1f29f137d10e2775d"></a>

## advertise_custom.advertise_where.advertise_v6_on_public.public_ip — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / bf04fece3c0a / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a267550051395f0f1fa132938a6bd0320f2f3134ed5093dc08c772e25054a15d)
- advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-c8b5c66303bf4ab2081cd2d9a14169c4121bcd290d3083d372332ff3bb91cce1"></a>

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

<a id="canonical-44175455b863ca49d35b428e09f837d99bf4dc667291dd551b40e072a92b03e2"></a>

## Direct properties — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / bf04fece3c0a / 3

<a id="canonical-716c7b953b2965fd882f2e39505804a147054dc4f0eb1433b246e5aef66c4b9d"></a>

<a id="canonical-a3f86b0d522474b621c6c077c7cb4970b401715ca392a922de019d53f7ad8a0d"></a>

## name property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / bf04fece3c0a / 4

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

<a id="canonical-5391fe3fc312740930ae527b66932f2ce09c2f6730c683be0f8f034c51631005"></a>

<a id="canonical-d42c78fd010c1b5aee8528e17a2554bf2f6fcb36155a4f150af58243200ea268"></a>

## namespace property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / bf04fece3c0a / 5

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

<a id="canonical-c7afb070e861f24e24fde2314ef7e688c1dad4b435e4ccef3fa9da0240442220"></a>

<a id="canonical-5fce745c5a725ca09ad2db9ba0e924c4e449b74337f5efd3e987f653b2108950"></a>

## tenant property — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / bf04fece3c0a / 6

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

<a id="canonical-0754a95deb3a0b7ee53ae8b8674e0fc2c784ef323d771cc616290684d4400b55"></a>

## Next pages — advertise_custom.advertise_where.advertise_v6_on_public.public_ip / bf04fece3c0a / 7

- [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a267550051395f0f1fa132938a6bd0320f2f3134ed5093dc08c772e25054a15d)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-01e4422723d7bed42f62d682fd4d27da3df14f1fc4c83321c61b4c303499373d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ab5a346f5545e6ad283a1caa98714493e6ce2860c11846eff35bf5cc5897768"></a>

## advertise_custom.advertise_where.site — advertise_custom.advertise_where.site / 3910e131a951 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.site

<a id="canonical-a6a767152e50372944f6de21a0e75602d8009b909870479faa17f3c389939a75"></a>

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

<a id="canonical-4d898b30f86f1a8e30e6de47f2b549b67b55270cec3cb3cbbecc96fc0e6b8e02"></a>

## Direct properties — advertise_custom.advertise_where.site / 3910e131a951 / 3

<a id="canonical-151c7295000af1ecf54150885354127eaf45b63102584a33f99f5939fcd8a5d5"></a>

<a id="canonical-6f2cdea3d06adda9f7f5886bce0378422bbc80a7e44c113a2a208fe76c3c5b93"></a>

## ip property — advertise_custom.advertise_where.site / 3910e131a951 / 4

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

<a id="canonical-00670fcdd99e2b5e504a42ab3cc72007d1ea822d0222acda96d8e26f54476190"></a>

<a id="canonical-f7fb34c6922fdcc8a6ae70458b00b870400e7aadc577983dbae909e47cb19625"></a>

## network property — advertise_custom.advertise_where.site / 3910e131a951 / 5

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

- [site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1e9a9a111f9037a02a75aa4216241abcacb7d9d56c20043a53dcc21dc731c8e9): complete subsection reference.

<a id="canonical-99f2f4cf8fc1a1d21cb45109f139f609ead0b6734faeee4b05c771ea224c845e"></a>

## Next pages — advertise_custom.advertise_where.site / 3910e131a951 / 6

- [advertise_custom.advertise_where.site.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1e9a9a111f9037a02a75aa4216241abcacb7d9d56c20043a53dcc21dc731c8e9)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-1e9a9a111f9037a02a75aa4216241abcacb7d9d56c20043a53dcc21dc731c8e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef8c08ae204a184d989d43fe7f7f85bda6ddf8a63763cadd4c2c366762710e06"></a>

## advertise_custom.advertise_where.site.site — advertise_custom.advertise_where.site.site / cdf6a6cf16c5 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-01e4422723d7bed42f62d682fd4d27da3df14f1fc4c83321c61b4c303499373d)
- advertise_custom.advertise_where.site.site

<a id="canonical-5fa4c80c00d47473caf0eb226338b27a7fcff2fa991af3b199b551f491321b58"></a>

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

<a id="canonical-6519d065d6a4838950212060fd8e114fe859fb4761695ab2ae3feaa84d11eb29"></a>

## Direct properties — advertise_custom.advertise_where.site.site / cdf6a6cf16c5 / 3

<a id="canonical-0d891dc7f9ad0f501ecb44470295aa7bb42e5b47685adbb8ae88b43bb661d23f"></a>

<a id="canonical-455aa7e84b060694dec487f757ebc9224f405cce58cdf2a2969dcd8096bb6f63"></a>

## name property — advertise_custom.advertise_where.site.site / cdf6a6cf16c5 / 4

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

<a id="canonical-d755fcae0515f1bdc0b4a2707677d11ea7888ff7194fa54977f0ef0aa221a3a8"></a>

<a id="canonical-cea93b9a99ea35707ba043e87f56e9d6cf878d6d5f2de7532bf45aa6167c0bb4"></a>

## namespace property — advertise_custom.advertise_where.site.site / cdf6a6cf16c5 / 5

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

<a id="canonical-e76c979073aef8ffb99e1b6b73170bfae135ebaa61f7aaa4aa9323b99c761e85"></a>

<a id="canonical-f9243d0e75bfb0a3408bb14c2fc0c2b2e5debeb5d2ad038cedcb810897a116c1"></a>

## tenant property — advertise_custom.advertise_where.site.site / cdf6a6cf16c5 / 6

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

<a id="canonical-a80662bcbe48f31517a50c109a124f0f5821bb59e95037128607c40bfe6ee154"></a>

## Next pages — advertise_custom.advertise_where.site.site / cdf6a6cf16c5 / 7

- [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-01e4422723d7bed42f62d682fd4d27da3df14f1fc4c83321c61b4c303499373d)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-9068df4283088d54603361b25c9c8a15536a1a25009996e52ae1f5f220a39190"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66c4563c96a6ccc7304c1917b6851f92572d745d40fbc40f16ebd24bdce738ad"></a>

## advertise_custom.advertise_where.use_default_port — advertise_custom.advertise_where.use_default_port / f14e778a5f87 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.use_default_port

<a id="canonical-b9f35391cf0784b50ced0479868e70948548efe842f8bc2f78a2bfc933c57929"></a>

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

<a id="canonical-2642395d3659246cae06358435a9dad2053d93c32f1aad9eeef5ef074298db18"></a>

## Direct properties — advertise_custom.advertise_where.use_default_port / f14e778a5f87 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1d679131e6250e166e6e0a999480cf6b0e254b327819af856fc01d972992099b"></a>

## Next pages — advertise_custom.advertise_where.use_default_port / f14e778a5f87 / 4

- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b89429158bb8404b5a907995fcf7236edf419d9648bca8b5ccb55ed8008ec3e8"></a>

## advertise_custom.advertise_where.virtual_network — advertise_custom.advertise_where.virtual_network / e93071c63bec / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.virtual_network

<a id="canonical-31836349c924f89e675716f34d5f42ead3ebed533cdab0e33f54cd1fcd7a374b"></a>

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

<a id="canonical-f4c05c2e68dc81c3fa67f378558d66eae519ccf3a094810d884cfc77a5f11809"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network / e93071c63bec / 3

- [default_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4b7005b4a7940eb3353bcf175c72b10b4a3ce04f478d8acc4404adcfbc48f892): complete subsection reference.

- [default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e13bcc137167da98d68bb6286f5dc101df58e47460188b1d66f51364b4a2e19d): complete subsection reference.

<a id="canonical-130c37c8c4f7eaacbc31c1e814c07dd49d5ec0938b3f2aa1d02b0dd905922840"></a>

<a id="canonical-e9923489115164d654bbc336d8814b79efa0c5ecbae8d599a2a1465c5d03a408"></a>

## specific_v6_vip property — advertise_custom.advertise_where.virtual_network / e93071c63bec / 4

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

<a id="canonical-e68a6a3bf18bea8e0ad7f368f18df7d804d2120f6ff4d09a40400d9c6b832c17"></a>

<a id="canonical-ed438c8b1c485d8caf4b2e30c250cab0a7dccefe9ec9cca5801fdad89e3421ba"></a>

## specific_vip property — advertise_custom.advertise_where.virtual_network / e93071c63bec / 5

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

- [virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c8c8f299efa8cdd977363e95694663d179855273c591a2f860dfe989bb87a1a7): complete subsection reference.

<a id="canonical-3887c9bec30dc8d2702450e7f1c1cfe9c6f3397e1d7c63a5c9f42669c9f893bf"></a>

## Next pages — advertise_custom.advertise_where.virtual_network / e93071c63bec / 6

- [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-4b7005b4a7940eb3353bcf175c72b10b4a3ce04f478d8acc4404adcfbc48f892)
- [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e13bcc137167da98d68bb6286f5dc101df58e47460188b1d66f51364b4a2e19d)
- [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-c8c8f299efa8cdd977363e95694663d179855273c591a2f860dfe989bb87a1a7)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-4b7005b4a7940eb3353bcf175c72b10b4a3ce04f478d8acc4404adcfbc48f892"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55e863367c8be310d777ca485fe8ef3e6472e5fd287193b53594a6e4d512efb1"></a>

## advertise_custom.advertise_where.virtual_network.default_v6_vip — advertise_custom.advertise_where.virtual_network.default_v6_vip / 533bc1b487a3 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187)
- advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-835fed5803e45014eb4c8c1bb8801a466a9b572cc1fbb60801d577b06716bfe5"></a>

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

<a id="canonical-6b0d9ee862568bfd8901b130dde883e6fa8e6367a67548cb5f7e187f96d209f7"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_v6_vip / 533bc1b487a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-95b0cc4044ab67568fb5cd23472ad901fd5e7bf1cb04cbdcf2e8fe6799a95917"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_v6_vip / 533bc1b487a3 / 4

- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-e13bcc137167da98d68bb6286f5dc101df58e47460188b1d66f51364b4a2e19d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-993ede9e16409a5343184d1f6c1bce05bf8a7d8b753c58c864ae1bcd70498318"></a>

## advertise_custom.advertise_where.virtual_network.default_vip — advertise_custom.advertise_where.virtual_network.default_vip / 4d24c9330b13 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187)
- advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-04a279ca07672ed414935c889bef6e242c57715ccc78822e53d5ae56a9976a60"></a>

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

<a id="canonical-9cce4e01c59a5a56f0170583cdb0d4b2709e42c0d28091ec82365ba7d4984e86"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.default_vip / 4d24c9330b13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-682d5d8c6025ec3d9452a61724fa370ac913bb0b293c894144581d61850a6fa9"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.default_vip / 4d24c9330b13 / 4

- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-c8c8f299efa8cdd977363e95694663d179855273c591a2f860dfe989bb87a1a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bb70bbd344ae9ff373c6070239cd005d7c11bb93d2320ac41f3368ca09a3ef6"></a>

## advertise_custom.advertise_where.virtual_network.virtual_network — advertise_custom.advertise_where.virtual_network.virtual_network / e7209e55db75 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187)
- advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-9061ea967f620202fb25cbd7213955ee0b3295650a7ba67b99169d8dc941245a"></a>

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

<a id="canonical-12b06c628dace3d15d1826433c0bd8baa4efd3065d45e75ba8995287d0013206"></a>

## Direct properties — advertise_custom.advertise_where.virtual_network.virtual_network / e7209e55db75 / 3

<a id="canonical-9d025d1bcbf7ccdc3ccf22825d6dac10f6747b87f39c17c8663e9eb374e664eb"></a>

<a id="canonical-f1c901560b2971f87c8179f19d4067268d2d6159378a1c8bfff850dc3f94655d"></a>

## name property — advertise_custom.advertise_where.virtual_network.virtual_network / e7209e55db75 / 4

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

<a id="canonical-1a6aae53cf1311614a143fc99966f4ee9a3165b3e638aff1fcd2dcef1a290bba"></a>

<a id="canonical-01c0d92f7923c94d8f776641ca0eebac5218369133e3b98fc02fec43b2de48bb"></a>

## namespace property — advertise_custom.advertise_where.virtual_network.virtual_network / e7209e55db75 / 5

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

<a id="canonical-481171f25461d22d7164155eebc8ed6940f86052c272a71fe3c037d7c962b6f3"></a>

<a id="canonical-fc4a541f2d55d5a44ad41af7ca8f7079ce42b163990b4a081547dc956e54fdf0"></a>

## tenant property — advertise_custom.advertise_where.virtual_network.virtual_network / e7209e55db75 / 6

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

<a id="canonical-3dd190c3146ec1d4f86162f25b53e31a3fcf187eb9ce3466e7dfbc39c5c65ba0"></a>

## Next pages — advertise_custom.advertise_where.virtual_network.virtual_network / e7209e55db75 / 7

- [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e843a49c3ba658d5299e58fb48a3a7202ce397fbe84de3a816b97c9dc70da187)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-e3744b7a41f67683e48a0745dc20d3accdeaead0c59b15491f85286673c468b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b1f3fe043ba9bf2c605a671c51c8c3e4341d72d720429a31d1878fe21277503"></a>

## advertise_custom.advertise_where.virtual_site — advertise_custom.advertise_where.virtual_site / c434a8fed511 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.virtual_site

<a id="canonical-5a8ca5e15664763e038f0f526e2b56e7e20a52960b200c0719a31ef5b2354d35"></a>

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

<a id="canonical-c1fd09ed34a43119829b203d447f28a1edee9ce8558443e519f9935d7f31ef12"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site / c434a8fed511 / 3

<a id="canonical-676d70f085372b0fdb3265dc8496fef569cb3fac4968c723de8e4d95aa45e094"></a>

<a id="canonical-b40f2c6c4cdd0e13b184b67e6d54bed1ee7614ff3b9ba5cca88c8f14abba36a1"></a>

## network property — advertise_custom.advertise_where.virtual_site / c434a8fed511 / 4

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

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-d791a09112a91b7dbf7e17bbc0b6d179279bf1709ea49ce247bd55db2173744b): complete subsection reference.

<a id="canonical-35eb44243b4f2fb0a908a8b3029bf4d1c134de1bc520c69364bafcb417b403ee"></a>

## Next pages — advertise_custom.advertise_where.virtual_site / c434a8fed511 / 5

- [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-d791a09112a91b7dbf7e17bbc0b6d179279bf1709ea49ce247bd55db2173744b)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-d791a09112a91b7dbf7e17bbc0b6d179279bf1709ea49ce247bd55db2173744b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcd2ef456d0ca75e9b19424d9692a8e039f774cb16793aa6b9d443fe4d907879"></a>

## advertise_custom.advertise_where.virtual_site.virtual_site — advertise_custom.advertise_where.virtual_site.virtual_site / 8d0dfd64c331 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e3744b7a41f67683e48a0745dc20d3accdeaead0c59b15491f85286673c468b5)
- advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-9a8de5046def0ec912287ffd1ec237bc657f7ba14c1ee78399fee822f73121a3"></a>

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

<a id="canonical-fb45a10ca9a1e247f26e57733d5b12f236d471fea32ca6fb925ac5bad7e39b85"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site.virtual_site / 8d0dfd64c331 / 3

<a id="canonical-a659eb84d48f65479b1e2d503eb43e901ae8b65335e1e002d3b98aded41b8dc3"></a>

<a id="canonical-aa999cb6dd6a92c34cb72a1ba0a23a484b4598d84a82c88d222802feddd15fe1"></a>

## name property — advertise_custom.advertise_where.virtual_site.virtual_site / 8d0dfd64c331 / 4

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

<a id="canonical-ceb8ac6867b0ef3c55c7da0e338de562e22741327898cfa3c514da683852de70"></a>

<a id="canonical-b3929a385f83c77c43ca449f4ce1dda823a8fd3108c35e456edb3acdc66c6a1f"></a>

## namespace property — advertise_custom.advertise_where.virtual_site.virtual_site / 8d0dfd64c331 / 5

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

<a id="canonical-9e277b29b5c5be2cae0cd1ceded003d6ab9107c70351ab225a6f22e96f42b5d7"></a>

<a id="canonical-6a1662c5276873247b0299353e05bd9b251fc899bb75b51a55add2b9e6883023"></a>

## tenant property — advertise_custom.advertise_where.virtual_site.virtual_site / 8d0dfd64c331 / 6

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

<a id="canonical-cbef140738a1327c8209ae3573e734cab46640b60c078fc4770c7b346fcceee6"></a>

## Next pages — advertise_custom.advertise_where.virtual_site.virtual_site / 8d0dfd64c331 / 7

- [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-e3744b7a41f67683e48a0745dc20d3accdeaead0c59b15491f85286673c468b5)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-a6693a30231263f66d37d6765afb02b246cd9bcff2fed3757e51efc5b07eed34"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6b074fdfffabff40a70401cdcce6a815f171f634378560a9fdcde0cebcd8f5b"></a>

## advertise_custom.advertise_where.virtual_site_with_vip — advertise_custom.advertise_where.virtual_site_with_vip / 34c8b9ad0a8a / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-3c53f59d598c79a2bcb2f33cd5a1793bf88760f624f7cf4d46902de7aeaf85c2"></a>

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

<a id="canonical-46ba820d0381a13a20fc7bba1a8a50e8ff96bd59f93247e75f193c0114991ea4"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip / 34c8b9ad0a8a / 3

<a id="canonical-feff7f533ada708a19a126f0a5e62dfd61c2f5beb504ed93baa862d088f9fdb3"></a>

<a id="canonical-69a5ad285011762108149389cbfe78e3b7984e7522c9c14108c64c9f6baa4d46"></a>

## ip property — advertise_custom.advertise_where.virtual_site_with_vip / 34c8b9ad0a8a / 4

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

<a id="canonical-6e57f8554d7880d622ad53f725628ecd270bbacd5c2d88396bd089dade0c0b40"></a>

<a id="canonical-687c20eff46c3a2daac0bc992b3739c27d8437b19a081587c7f60e75be2d417c"></a>

## network property — advertise_custom.advertise_where.virtual_site_with_vip / 34c8b9ad0a8a / 5

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

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-dccb76ff925483742be4596ace5e3fe4bd4dc61734a104e5db1eefb9a213cb7e): complete subsection reference.

<a id="canonical-646d3ea25632d3c25c9c84d949d43bca05f2f5b116d72e5923b37848c153b79a"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip / 34c8b9ad0a8a / 6

- [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-dccb76ff925483742be4596ace5e3fe4bd4dc61734a104e5db1eefb9a213cb7e)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-dccb76ff925483742be4596ace5e3fe4bd4dc61734a104e5db1eefb9a213cb7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-132c715456d020edc883a536472d5aca5c66f9b698c0afe89dd2870d83cfdc33"></a>

## advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 686e02180dfa / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a6693a30231263f66d37d6765afb02b246cd9bcff2fed3757e51efc5b07eed34)
- advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-6f2b01ab333f3a00f812f3f802d0f282dc3b549063f6ec18ab7a75522e354c91"></a>

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

<a id="canonical-19ae75d8a0ca5e5dc404fe1ec6cd4f3fb33109bfd13cdf864b21e91708621a93"></a>

## Direct properties — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 686e02180dfa / 3

<a id="canonical-8addf50e48c39d140b12682f4be0910e3c783cad7290c6c2252a6322e65a4d25"></a>

<a id="canonical-2d2e189002335571e6e0b399ebe9e7f2124a2537f955b933812beadfd6f61613"></a>

## name property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 686e02180dfa / 4

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

<a id="canonical-ac82e971d56147a27516d9a6709c88c720a415979507165e90f965590e9f51bb"></a>

<a id="canonical-6368f8c9d497942b8153b39df016bf1726c4cf482a450afd5a11074a1cb61763"></a>

## namespace property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 686e02180dfa / 5

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

<a id="canonical-556126d348fb84ff868bb66195e9798b16a0c77b8e743b5e6eb77882f19da675"></a>

<a id="canonical-69cea4883225cb9dd4a9f95ec4db6f944603fb4a3da2319d157021d83a9b2404"></a>

## tenant property — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 686e02180dfa / 6

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

<a id="canonical-6a92034e9b6cfdf3602c43ed24c58c9cfb192a40fb0086e705ff8182e65464ba"></a>

## Next pages — advertise_custom.advertise_where.virtual_site_with_vip.virtual_site / 686e02180dfa / 7

- [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-a6693a30231263f66d37d6765afb02b246cd9bcff2fed3757e51efc5b07eed34)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-7a6d49a64fa8bd2949bc29e099a1eec058e153c2860b56e8087604ed3b7a2cf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-020a71b8b5c593064e341471f3bf83876855f3c177fd4cf2158725abbbb8f3a5"></a>

## advertise_custom.advertise_where.vk8s_service — advertise_custom.advertise_where.vk8s_service / 3124dfbdd9ac / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- advertise_custom.advertise_where.vk8s_service

<a id="canonical-921d43a2c1ba075ff09cd88ac3ed763ddc7ad1a1ac40d2c19c5e5177f5b15f01"></a>

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

<a id="canonical-5b630e0eac75c26cfb472ad92996e1a6e1ed1c9e4273d0de673c09a0f3fc5150"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service / 3124dfbdd9ac / 3

- [site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-db80ae020cd93bd2965a6079780072039a595ba624318dd0147ea004f255eafd): complete subsection reference.

- [virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8d2e1f005eb13c9d86fe580663e72fbcda3ecd2a079e3b70e619697dc70be43b): complete subsection reference.

<a id="canonical-1f52f349de0d4634fb5d5f455c40e41717fd59d1971a64da26b831312e278485"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service / 3124dfbdd9ac / 4

- [advertise_custom.advertise_where.vk8s_service.site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-db80ae020cd93bd2965a6079780072039a595ba624318dd0147ea004f255eafd)
- [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--udp_loadbalancer--reference--group-001.md#canonical-8d2e1f005eb13c9d86fe580663e72fbcda3ecd2a079e3b70e619697dc70be43b)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-db80ae020cd93bd2965a6079780072039a595ba624318dd0147ea004f255eafd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c2383c44b9ba8dfbf87b82ef425f22b199f0e7c30d55863dc3bfc12a9ead9dc"></a>

## advertise_custom.advertise_where.vk8s_service.site — advertise_custom.advertise_where.vk8s_service.site / 6b7d7e8fc6a1 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7a6d49a64fa8bd2949bc29e099a1eec058e153c2860b56e8087604ed3b7a2cf8)
- advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-5d2eef8b53e72630869469a965b36323f7a2d97df2a11a444b091d86bba08348"></a>

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

<a id="canonical-f82eeb8a54e520c0d605157ab7bb234dc58a6d950195ec1f858de71345a8021e"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.site / 6b7d7e8fc6a1 / 3

<a id="canonical-b30da8e784cf85e36fecb72a0d997df77b7bf0dd3e5cc21da9ba55f1649acb19"></a>

<a id="canonical-955b25d8b00e778b219574189430d4ba46b0eafaf4f4e504d86b2da4fa6578cf"></a>

## name property — advertise_custom.advertise_where.vk8s_service.site / 6b7d7e8fc6a1 / 4

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

<a id="canonical-7048066f756fd3b10897c5ca0ab3b7ce41955cd2df61ed4bc98f700a41b336f5"></a>

<a id="canonical-f279211e23365cd340a13183ef1ee29a7c6301524370a14bbcf7763081481a2f"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.site / 6b7d7e8fc6a1 / 5

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

<a id="canonical-84ae926b200245d21db2ed4df47764716d5e234d68534e974d8dabcd6ed0c55e"></a>

<a id="canonical-92c6ddae22c445b7bb4cfe605a8b3110d1b544ccdb629bf5a461c6f06ae6b385"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.site / 6b7d7e8fc6a1 / 6

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

<a id="canonical-26c3c2b41961fe9b2e223a8625b553988d88c71549b2309817364cb3a79b2d81"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.site / 6b7d7e8fc6a1 / 7

- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7a6d49a64fa8bd2949bc29e099a1eec058e153c2860b56e8087604ed3b7a2cf8)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-8d2e1f005eb13c9d86fe580663e72fbcda3ecd2a079e3b70e619697dc70be43b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aafae01c4cbc743fd87f5b9012772efdc9cabbd2bcbba3293441a14c5c02d407"></a>

## advertise_custom.advertise_where.vk8s_service.virtual_site — advertise_custom.advertise_where.vk8s_service.virtual_site / 1c0d59977b56 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_custom](data-sources--udp_loadbalancer--reference--group-001.md#canonical-1217abbb8eb8ca4cc71a70cf5d21dc5b98101af311436b6ae4e9aceab94e8996)
- [advertise_custom.advertise_where](data-sources--udp_loadbalancer--reference--group-001.md#canonical-bcfd8cb11a5cec5a54a80ddd371323141209fa3f6bed7c6a4712950d81211617)
- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7a6d49a64fa8bd2949bc29e099a1eec058e153c2860b56e8087604ed3b7a2cf8)
- advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-7dbef697841fa637133361fa8df3a355ed2a30b04b0c8aa773c0d75587c10a1c"></a>

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

<a id="canonical-58a4f14d68b79cb3c0f372962bbe8b41c1f5d01115c9dc3906efc9dfbc6207e7"></a>

## Direct properties — advertise_custom.advertise_where.vk8s_service.virtual_site / 1c0d59977b56 / 3

<a id="canonical-074aaffd2c44f9cf4a7299bc825ffe5214c74f54f9fdca2f835571938fee6b62"></a>

<a id="canonical-7575191fd237b8e0b5fe24a36a1aea8ca0cc1f2766e2720339a11f6182bc3e9e"></a>

## name property — advertise_custom.advertise_where.vk8s_service.virtual_site / 1c0d59977b56 / 4

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

<a id="canonical-40fc8df6244644653bfd24ef4785affd6ad758b72fd66707fb0f64e92afda191"></a>

<a id="canonical-31965c4259bef0c35cea224c51acc9aa28db3ad7d9583b1e2466eab119b46704"></a>

## namespace property — advertise_custom.advertise_where.vk8s_service.virtual_site / 1c0d59977b56 / 5

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

<a id="canonical-e96ed6787985651542d3fc82df91006f7fd4012e6310a0b75b24fee7908b54f7"></a>

<a id="canonical-84dd7f26a85000ebfb75de3d764e9c3a2b17613dd6cffe364c88bc6573de8714"></a>

## tenant property — advertise_custom.advertise_where.vk8s_service.virtual_site / 1c0d59977b56 / 6

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

<a id="canonical-58ab33d7dd066b07127f792e0649c6db1509725b411a0113e4bb9608a52c80f1"></a>

## Next pages — advertise_custom.advertise_where.vk8s_service.virtual_site / 1c0d59977b56 / 7

- [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--reference--group-001.md#canonical-7a6d49a64fa8bd2949bc29e099a1eec058e153c2860b56e8087604ed3b7a2cf8)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-d5168f4eb73ed5b138b2244fd19ddf6629ee4860ec7e3750c5cc7580f308ee5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023cb4fa5258127d6034cb62c039807b32a26303a5e51977cb7d695ea9b2a4e"></a>

## advertise_on_public — advertise_on_public / 53af23e85c09 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- advertise_on_public

<a id="canonical-c2496f692c3248fafa7fe8560a189583638c1dd790eedbe6b21d479d733fc33f"></a>

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

<a id="canonical-3441388b705481bb11c415d1cad3f137faf5a636ec398c3b5007a744da4a0910"></a>

## Direct properties — advertise_on_public / 53af23e85c09 / 3

- [public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-757cb6d8d698e73c880f7e17d63e9bdb7ee4263f8562a1929adf55cd0a4a64c4): complete subsection reference.

<a id="canonical-330de529a1ade743d736d5f62eb0e4734d7f4db58025797905d3564bcf77eb6a"></a>

## Next pages — advertise_on_public / 53af23e85c09 / 4

- [advertise_on_public.public_ip](data-sources--udp_loadbalancer--reference--group-001.md#canonical-757cb6d8d698e73c880f7e17d63e9bdb7ee4263f8562a1929adf55cd0a4a64c4)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-757cb6d8d698e73c880f7e17d63e9bdb7ee4263f8562a1929adf55cd0a4a64c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a0c06e7856da3134c71fbb87cc4d047b3fe1ac333017c90d41a640f2931022f"></a>

## advertise_on_public.public_ip — advertise_on_public.public_ip / e0198f63c243 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-d5168f4eb73ed5b138b2244fd19ddf6629ee4860ec7e3750c5cc7580f308ee5d)
- advertise_on_public.public_ip

<a id="canonical-8bd7453d768e7c56d262312aa8577e43e29ecbf919b826b095d9df14d470c3d6"></a>

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

<a id="canonical-1d697c4bc746387df3d81ae060d3e860d7e1649f07a741e85537914096d71acb"></a>

## Direct properties — advertise_on_public.public_ip / e0198f63c243 / 3

<a id="canonical-c200698e083393d32b4cf3517a74e8ebb4ff4a33650c3d2c7f990b1a2b427f99"></a>

<a id="canonical-b3fb8be8994670c512f0afdece1afc8a24629753ec6324ab6d58a2cd235cf218"></a>

## name property — advertise_on_public.public_ip / e0198f63c243 / 4

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

<a id="canonical-02d8bf12de40b94df80dc596fb188e05c7709a710166d16d39259336f88654b3"></a>

<a id="canonical-a609a96acf84da6fd418b125463703ee03f8d3019a3059b0f814a2e576186f2b"></a>

## namespace property — advertise_on_public.public_ip / e0198f63c243 / 5

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

<a id="canonical-1cdbb02b44c067dc65fe3cea8e7cb70184fbf29ab8d91926704fe9126f13e885"></a>

<a id="canonical-a39b08771a7c0f45ec6ab96d7290e2f4d32acf42983ba43cb34fa3369e7f6a75"></a>

## tenant property — advertise_on_public.public_ip / e0198f63c243 / 6

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

<a id="canonical-cdebae46968a20d72fe468fe746184f7e43d8a504816f16c64fe44da932bae02"></a>

## Next pages — advertise_on_public.public_ip / e0198f63c243 / 7

- [advertise_on_public](data-sources--udp_loadbalancer--reference--group-001.md#canonical-d5168f4eb73ed5b138b2244fd19ddf6629ee4860ec7e3750c5cc7580f308ee5d)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-b5eb4ceb2d20a15434aee06f351d8acc20cbc0d0bc6f34c5d029c3c0967674d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08890e8d164aa32aae667708d8f432d3f8f17f8164a956237b11aae3dba61c81"></a>

## advertise_on_public_default_vip — advertise_on_public_default_vip / 7ee962975f90 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- advertise_on_public_default_vip

<a id="canonical-b4e5a59b1bb9de65f78677156c820e8ccb2d0f5bf1ae81bab4d58c8cc8b61530"></a>

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

<a id="canonical-7de9cf711481da0e0b169afa0606aa73bf3c20503a5d75417612caeef498d074"></a>

## Direct properties — advertise_on_public_default_vip / 7ee962975f90 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b64f11805ba4cb312ecbd3425a75f36b986028457f66a2731f9f7666448b72c"></a>

## Next pages — advertise_on_public_default_vip / 7ee962975f90 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-1d03dc3a13c57995a5526df4c20ac937d35490f9430db921da4213bc2adecd0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a144c74184429000613104d7864b3acaec6138474273791322e183a2eb6f983"></a>

## do_not_advertise — do_not_advertise / a876d49f3d57 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- do_not_advertise

<a id="canonical-116708bf49f53a7a81817599cb6d0dc09f3b35a1d6e478e3707e2114ef59b978"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-b836961834d3bfd01b4d0a258050d43b1961a0e18d691a286ab6d0ae3694cac9"></a>

## Direct properties — do_not_advertise / a876d49f3d57 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f8f7eb63946330d7e2d72e195a27cf6a65af62de27b778a245bad9243a446c7c"></a>

## Next pages — do_not_advertise / a876d49f3d57 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-c2840de39be22d8c0af335282ccaa658e1b6de712ef708eb0abf925c02f6cec9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233e9dcb53c11647430cc18fb89b367b8e143718bcf5e36d64201224202a77c"></a>

## hash_policy_choice_random — hash_policy_choice_random / 9f651695ee13 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- hash_policy_choice_random

<a id="canonical-5aad3fcb6f7df3486afc79a36520af1935ac44953101b27ae00bb6225f15952d"></a>

Type: `["object", {}]`. Computed.

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

- [hash_policy_choice_random](data-sources--udp_loadbalancer--reference--group-001.md#canonical-5aad3fcb6f7df3486afc79a36520af1935ac44953101b27ae00bb6225f15952d)
- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--reference--group-001.md#canonical-f1b328288612454845b907f088f36b22dfb6dfe6149b7d87cb47eff5cc86c29c)
- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--reference--group-001.md#canonical-aa986a91e9d0c055ff991dc2e46dc133575d35241eb58f59a88dcb021a3152dc)

Select alternatives according to the provider validators above.

<a id="canonical-06a5a34a456e0adedfb12973f91e16f31d527e7ad55a0e2d46e2dd6b4fcd1ca6"></a>

## Direct properties — hash_policy_choice_random / 9f651695ee13 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8908d6f306a11e0b82afa7e6c30a58fee3cae4dc613bbc4b6cd2644c896077ea"></a>

## Next pages — hash_policy_choice_random / 9f651695ee13 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-6bba225388bef94021eb7f1db16412ca896935d5b8cb4fe6c704d1acbf25f44d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d418b3195ccf6116d51307f39de2a3f1631c839c701438f177cade5a4405d552"></a>

## hash_policy_choice_round_robin — hash_policy_choice_round_robin / 3585f65bd4d6 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- hash_policy_choice_round_robin

<a id="canonical-f1b328288612454845b907f088f36b22dfb6dfe6149b7d87cb47eff5cc86c29c"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-5c5470c09c779c8b05a9d5445d58dbd199a7416930462f6b394ada9a5afbabf2"></a>

## Direct properties — hash_policy_choice_round_robin / 3585f65bd4d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-696a3dd60dfd80163afe789f6b3b07433c2d62e23f7b1f9487f5686005ae2f79"></a>

## Next pages — hash_policy_choice_round_robin / 3585f65bd4d6 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-368e636030383f2acb10c46c6725dc4705a2c8d9c936ff5c381954f25fcb7478"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69277377591e64f56f82ca84aada99462b86bdad5977b2de6780ba8951614d50"></a>

## hash_policy_choice_source_ip_stickiness — hash_policy_choice_source_ip_stickiness / dfe4004d7e43 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- hash_policy_choice_source_ip_stickiness

<a id="canonical-aa986a91e9d0c055ff991dc2e46dc133575d35241eb58f59a88dcb021a3152dc"></a>

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

<a id="canonical-7ed8f8ae40ec484a9b3b1f6e0337287173212501d356ed1dfe27fdc79c693786"></a>

## Direct properties — hash_policy_choice_source_ip_stickiness / dfe4004d7e43 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-170133b12e8b439aa61125d408d79418c55799ca49891be7a08889aab8d6d764"></a>

## Next pages — hash_policy_choice_source_ip_stickiness / dfe4004d7e43 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-35483bfc197f023ade1e8abae3d8af09cfeb45b6bb50280c18c7f65bc558fdc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43c99170977bd6819364ae82265b62f62f0d708cd2214ee3ae88da3ea0b6737b"></a>

## no_service_policies — no_service_policies / 0a6839242be6 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- no_service_policies

<a id="canonical-4e9d8ef93f9f96f932e3cff70e2315a8c75729e4edf984b897414ce8685b5bcc"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-aba076d541470b2eb657dbd658b83877a35acc0ac7b30d4381649dd1253e7370"></a>

## Direct properties — no_service_policies / 0a6839242be6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ceea9fdc9e2d0262f8332e850ac344edb5721c945e1260c4a3663037d776a456"></a>

## Next pages — no_service_policies / 0a6839242be6 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9021e36e7276e0fadec4ccb3b29d253ac095f76e7945aee69b1c869efd17f2c"></a>

## origin_pools_weights — origin_pools_weights / 5ce601410191 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- origin_pools_weights

<a id="canonical-1f1421826b7784862b3833f2d8009928e5328c84734716d083c36a26aab3f4ac"></a>

Type: `"list"`. Computed.

Origin pools with weights and priorities used for this load balancer.

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

<a id="canonical-b381e2cc28d849f9a053e273ef361298a6a01de6514aa56b36a5b5859e834f83"></a>

## Direct properties — origin_pools_weights / 5ce601410191 / 3

- [cluster](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2c71f8bf417e40520eee99e8c506835e55bb8e8337254dbf27b6d5a86496ed92): complete subsection reference.

- [endpoint_subsets](data-sources--udp_loadbalancer--reference--group-001.md#canonical-72c7f9715ba513dd9ecca68d7e6ece7bc6c62354844b0a4f139fccf28d4069a2): complete subsection reference.

- [pool](data-sources--udp_loadbalancer--reference--group-001.md#canonical-fefb843799b216408a1d6f75e1434465b919f350ace3caecd0a44a9e1ed2dd38): complete subsection reference.

<a id="canonical-307e125d127cade86eedfd22e471f5dd543a64565ae18fb9772e4f69e00d2542"></a>

<a id="canonical-2c0b9df27c1c2ed230aebd1bb5cdb540e829beb40a4046feadf6ec01b9e5d203"></a>

## priority property — origin_pools_weights / 5ce601410191 / 4

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

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

<a id="canonical-c9e02595c0a17cfe44ecbad89769058109eff9eb96510c45150979339ad68f6b"></a>

<a id="canonical-79c012890ce02cceab8eef886af9d0a745e0f2a3e46fa89e2f98f6ba127c6fb2"></a>

## weight property — origin_pools_weights / 5ce601410191 / 5

Type: `"number"`. Computed.

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

<a id="canonical-138156bbf62373d1da15d4c629e6542e1bcc0ab74056741159c002d85e50ad2a"></a>

## Next pages — origin_pools_weights / 5ce601410191 / 6

- [origin_pools_weights.cluster](data-sources--udp_loadbalancer--reference--group-001.md#canonical-2c71f8bf417e40520eee99e8c506835e55bb8e8337254dbf27b6d5a86496ed92)
- [origin_pools_weights.endpoint_subsets](data-sources--udp_loadbalancer--reference--group-001.md#canonical-72c7f9715ba513dd9ecca68d7e6ece7bc6c62354844b0a4f139fccf28d4069a2)
- [origin_pools_weights.pool](data-sources--udp_loadbalancer--reference--group-001.md#canonical-fefb843799b216408a1d6f75e1434465b919f350ace3caecd0a44a9e1ed2dd38)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-2c71f8bf417e40520eee99e8c506835e55bb8e8337254dbf27b6d5a86496ed92"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64c21bde2bc1a2c213941a6e3d677b4a0864abcca7e18bef75967eb7301bba39"></a>

## origin_pools_weights.cluster — origin_pools_weights.cluster / 2877ea74bf2d / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7)
- origin_pools_weights.cluster

<a id="canonical-4d965e320d766842274f7fa134e5eb11ac95c939f5792aeaa52b8d8e23b09624"></a>

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

<a id="canonical-329c1ff92b5b926ae9d3454a50c3fe69e09ab843e5ce302a3b3b080d4723822f"></a>

## Direct properties — origin_pools_weights.cluster / 2877ea74bf2d / 3

<a id="canonical-ed8f03d1111e1a46c84c43fe60eb65c4c4deb07d4fd7749636607933e55dfa4b"></a>

<a id="canonical-e08750ecd2393d0036ec6ed08aa8eceb9bf4b6f058328cdf8c3ca68fb0e6f0aa"></a>

## name property — origin_pools_weights.cluster / 2877ea74bf2d / 4

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

<a id="canonical-ecaad0099c4e8333eb9f997a8179621919a974e075dd7750526fae76cda75aa5"></a>

<a id="canonical-587908fd1b89fdfb45997c53ad3ffa9be9c29957a185d6a84336ae9d9d0db885"></a>

## namespace property — origin_pools_weights.cluster / 2877ea74bf2d / 5

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

<a id="canonical-cce9166994a1fca32a9b2ebc4065181aa0ced8f7afbd2c8a166ca31f587be4b8"></a>

<a id="canonical-c43a962b7666ce5687a8a383a2043fc8d6ec80c762a14a981a880cc62f15a357"></a>

## tenant property — origin_pools_weights.cluster / 2877ea74bf2d / 6

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

<a id="canonical-684d0e9c767a0c8fe6b9881ad44d08b0c06a467903d25dfed0d2089deb81e763"></a>

## Next pages — origin_pools_weights.cluster / 2877ea74bf2d / 7

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-72c7f9715ba513dd9ecca68d7e6ece7bc6c62354844b0a4f139fccf28d4069a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-667b70c7ebcfadfcf28a406b38fd47268b5d6abb8b87721cebf7ad795d767a08"></a>

## origin_pools_weights.endpoint_subsets — origin_pools_weights.endpoint_subsets / 334c30510e6e / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7)
- origin_pools_weights.endpoint_subsets

<a id="canonical-1478454b614689f916ca494f161c6d8dcd0c48afcbb85b703fb05f74a7f8cb13"></a>

Type: `"single"`. Computed.

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

<a id="canonical-9933bd74d200192e75b44155a90a4aef60044df318d270e9843caa2d15852458"></a>

## Direct properties — origin_pools_weights.endpoint_subsets / 334c30510e6e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-242ade5e9bfba089c177891e316e5afc5595a2d6dcd9af7c5b3935777b817dad"></a>

## Next pages — origin_pools_weights.endpoint_subsets / 334c30510e6e / 4

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-fefb843799b216408a1d6f75e1434465b919f350ace3caecd0a44a9e1ed2dd38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19403c910473000b3c0b4eda64132d1c513a1376098342217e1756e61ae7f7c5"></a>

## origin_pools_weights.pool — origin_pools_weights.pool / 97a4d01f6778 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7)
- origin_pools_weights.pool

<a id="canonical-7b7d183a67cb28564e5559d3f1a480927bf9bfe5739a0e54bc8395daacbf6dd3"></a>

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

<a id="canonical-e5e71d259dba664a914d91277f892a8f1a016cb94619ff61600e221f8d67ef9d"></a>

## Direct properties — origin_pools_weights.pool / 97a4d01f6778 / 3

<a id="canonical-6c4a7229e6883f3812c972007736773a359ccfd88853e74ec622b49738233980"></a>

<a id="canonical-80f04c5586f423c758eef5a09d61bb605a96eb2810001e19a836177f738dd9a1"></a>

## name property — origin_pools_weights.pool / 97a4d01f6778 / 4

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

<a id="canonical-b9828aeafe43a2dc4d37e1b2fca2dc4ff24eb14f66644f904fed51d0b80c6067"></a>

<a id="canonical-11a89b525ab243bb1196a528647ee8ee4077e4455ebcea4aeac48beaaf06da35"></a>

## namespace property — origin_pools_weights.pool / 97a4d01f6778 / 5

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

<a id="canonical-4df4327d1247875528a3a25566644d6495033a90ce15633788b566bdd71d351c"></a>

<a id="canonical-35adcadfdff84f523b366a6d75808a517e5742c6d0dfa0899359b7f5f8c2c961"></a>

## tenant property — origin_pools_weights.pool / 97a4d01f6778 / 6

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

<a id="canonical-f24f09026a903a8f46a94fe9ea4a49195c87366c1638d7e6492bbe49b5a76614"></a>

## Next pages — origin_pools_weights.pool / 97a4d01f6778 / 7

- [origin_pools_weights](data-sources--udp_loadbalancer--reference--group-001.md#canonical-b8d4ff425fe199a738fd751d984278a9af2053317a1cc9995d6293e48ef106a7)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-dab7f01b3657f3fbbb46b538eb5e89792a2ec50653869106249bf0f8da5d6f30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd6f810b0d5e514dd1b8d9e32f5c40ee435b13724901f5245c28c5362f6a2680"></a>

## service_policies_from_namespace — service_policies_from_namespace / 6b03e1402356 / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- service_policies_from_namespace

<a id="canonical-4dbc966db9d1d33246273ef3fba34f087674f9f47eaa6b07c12b74440c4ee878"></a>

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

<a id="canonical-f6fb5c254edbf9e8f296a82041183936cd2b792c850c8e42802d51fd93309b81"></a>

## Direct properties — service_policies_from_namespace / 6b03e1402356 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a284678375993f3620057234456ea9978d03261e9fdfd28a6114ec6a3a4f066"></a>

## Next pages — service_policies_from_namespace / 6b03e1402356 / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)

<a id="canonical-b8029559a5e7fea91d6d135845a81c92c829ceeb7cc054314b011e12dbaa7f04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-901270537f262259e94df1e2b205f04038142d2c19aeac923e240dd15ca0ecae"></a>

## udp — udp / 90d615f39cba / 2

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- udp

<a id="canonical-3abd8ed3ddf2ea410ad301d36d0a459dcb11f8639b54fb7c7c534d5cd5850a84"></a>

Type: `"single"`. Computed.

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

<a id="canonical-8b6863a3ed64825d329bad9a74c159f62c8808aaf046f5358181fa0d4b567751"></a>

## Direct properties — udp / 90d615f39cba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36bbfd67224a8e5ece7eee84a6fc43f1080df3b69bfeed42e3cbda3ed9ea12f3"></a>

## Next pages — udp / 90d615f39cba / 4

- [Property reference](data-sources--udp_loadbalancer--reference--group-001.md#canonical-90bfc5c745211e4757205cd8fe220a2b5c2c1103253488671914bb68c847c39c)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md#canonical-156fae62fee7cb9e0e712d30277f9f65b8ba96c3336177b95c39c819598c5ba3)
