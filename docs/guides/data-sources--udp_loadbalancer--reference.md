---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 36823, "body_sha256": "sha256:e8f441721f28fa1d5930c74b8e381e0b03aa7c2f5ec43662a55224869e1eb289", "canonical_id": "xcsh-docs:data-sources:udp_loadbalancer:reference", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:properties:active_service_policies", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_custom", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_on_public", "xcsh-docs:data-sources:udp_loadbalancer:properties:advertise_on_public_default_vip", "xcsh-docs:data-sources:udp_loadbalancer:properties:do_not_advertise", "xcsh-docs:data-sources:udp_loadbalancer:properties:hash_policy_choice_random", "xcsh-docs:data-sources:udp_loadbalancer:properties:hash_policy_choice_round_robin", "xcsh-docs:data-sources:udp_loadbalancer:properties:hash_policy_choice_source_ip_stickiness", "xcsh-docs:data-sources:udp_loadbalancer:properties:no_service_policies", "xcsh-docs:data-sources:udp_loadbalancer:properties:origin_pools_weights", "xcsh-docs:data-sources:udp_loadbalancer:properties:service_policies_from_namespace", "xcsh-docs:data-sources:udp_loadbalancer:properties:udp"], "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:reference", "parent_id": "xcsh-docs:data-sources:udp_loadbalancer:fundamentals", "path": "docs/guides/data-sources--udp_loadbalancer--reference.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
- Property reference

## Direct properties

- [active_service_policies](data-sources--udp_loadbalancer--properties--active_service_policies.md): complete subsection reference.

- [advertise_custom](data-sources--udp_loadbalancer--properties--advertise_custom.md): complete subsection reference.

- [advertise_on_public](data-sources--udp_loadbalancer--properties--advertise_on_public.md): complete subsection reference.

- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--properties--advertise_on_public_default_vip.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

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

<a id="schema-description"></a>

### description property

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

<a id="schema-dns_volterra_managed"></a>

### dns_volterra_managed property

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

- [do_not_advertise](data-sources--udp_loadbalancer--properties--do_not_advertise.md): complete subsection reference.

<a id="schema-domains"></a>

### domains property

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

- [hash_policy_choice_random](data-sources--udp_loadbalancer--properties--hash_policy_choice_random.md): complete subsection reference.

- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--properties--hash_policy_choice_round_robin.md): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--properties--hash_policy_choice_source_ip_stickiness.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-idle_timeout"></a>

### idle_timeout property

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

<a id="schema-labels"></a>

### labels property

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

<a id="schema-listen_port"></a>

### listen_port property

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

- [listen_port](data-sources--udp_loadbalancer--reference.md#schema-listen_port)
- [port_ranges](data-sources--udp_loadbalancer--reference.md#schema-port_ranges)

Select alternatives according to the provider validators above.

<a id="schema-name"></a>

### name property

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

<a id="schema-namespace"></a>

### namespace property

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

- [no_service_policies](data-sources--udp_loadbalancer--properties--no_service_policies.md): complete subsection reference.

- [origin_pools_weights](data-sources--udp_loadbalancer--properties--origin_pools_weights.md): complete subsection reference.

<a id="schema-port_ranges"></a>

### port_ranges property

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

- [service_policies_from_namespace](data-sources--udp_loadbalancer--properties--service_policies_from_namespace.md): complete subsection reference.

- [udp](data-sources--udp_loadbalancer--properties--udp.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](data-sources--udp_loadbalancer--properties--active_service_policies.md#section) |
| `active_service_policies.policies` | [active_service_policies.policies](data-sources--udp_loadbalancer--properties--active_service_policies--policies.md#section) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](data-sources--udp_loadbalancer--properties--active_service_policies--policies.md#schema-active_service_policies--policies--name) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](data-sources--udp_loadbalancer--properties--active_service_policies--policies.md#schema-active_service_policies--policies--namespace) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](data-sources--udp_loadbalancer--properties--active_service_policies--policies.md#schema-active_service_policies--policies--tenant) |
| `advertise_custom` | [advertise_custom](data-sources--udp_loadbalancer--properties--advertise_custom.md#section) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where.md#section) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public.md#section) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#section) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--namespace) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--tenant) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public.md#section) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#section) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_on_public--public_ip--name) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_on_public--public_ip--namespace) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_on_public--public_ip--tenant) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md#section) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#section) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where.md#schema-advertise_custom--advertise_where--port) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where.md#schema-advertise_custom--advertise_where--port_ranges) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site.md#section) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site.md#schema-advertise_custom--advertise_where--site--ip) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site.md#schema-advertise_custom--advertise_where--site--network) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#section) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#schema-advertise_custom--advertise_where--site--site--name) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#schema-advertise_custom--advertise_where--site--site--namespace) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#schema-advertise_custom--advertise_where--site--site--tenant) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md#section) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md#section) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_v6_vip.md#section) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_vip.md#section) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--specific_v6_vip) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--specific_vip) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#section) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--virtual_network--name) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--virtual_network--namespace) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--virtual_network--tenant) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md#section) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--network) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#section) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--virtual_site--name) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--virtual_site--namespace) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--virtual_site--tenant) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md#section) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--ip) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--network) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#section) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--namespace) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--tenant) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service.md#section) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#section) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#schema-advertise_custom--advertise_where--vk8s_service--site--name) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#schema-advertise_custom--advertise_where--vk8s_service--site--namespace) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#schema-advertise_custom--advertise_where--vk8s_service--site--tenant) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#section) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-advertise_custom--advertise_where--vk8s_service--virtual_site--name) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-advertise_custom--advertise_where--vk8s_service--virtual_site--namespace) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--udp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-advertise_custom--advertise_where--vk8s_service--virtual_site--tenant) |
| `advertise_on_public` | [advertise_on_public](data-sources--udp_loadbalancer--properties--advertise_on_public.md#section) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](data-sources--udp_loadbalancer--properties--advertise_on_public--public_ip.md#section) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](data-sources--udp_loadbalancer--properties--advertise_on_public--public_ip.md#schema-advertise_on_public--public_ip--name) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](data-sources--udp_loadbalancer--properties--advertise_on_public--public_ip.md#schema-advertise_on_public--public_ip--namespace) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](data-sources--udp_loadbalancer--properties--advertise_on_public--public_ip.md#schema-advertise_on_public--public_ip--tenant) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](data-sources--udp_loadbalancer--properties--advertise_on_public_default_vip.md#section) |
| `annotations` | [annotations](data-sources--udp_loadbalancer--reference.md#schema-annotations) |
| `description` | [description](data-sources--udp_loadbalancer--reference.md#schema-description) |
| `dns_volterra_managed` | [dns_volterra_managed](data-sources--udp_loadbalancer--reference.md#schema-dns_volterra_managed) |
| `do_not_advertise` | [do_not_advertise](data-sources--udp_loadbalancer--properties--do_not_advertise.md#section) |
| `domains` | [domains](data-sources--udp_loadbalancer--reference.md#schema-domains) |
| `hash_policy_choice_random` | [hash_policy_choice_random](data-sources--udp_loadbalancer--properties--hash_policy_choice_random.md#section) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--properties--hash_policy_choice_round_robin.md#section) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--properties--hash_policy_choice_source_ip_stickiness.md#section) |
| `id` | [id](data-sources--udp_loadbalancer--reference.md#schema-id) |
| `idle_timeout` | [idle_timeout](data-sources--udp_loadbalancer--reference.md#schema-idle_timeout) |
| `labels` | [labels](data-sources--udp_loadbalancer--reference.md#schema-labels) |
| `listen_port` | [listen_port](data-sources--udp_loadbalancer--reference.md#schema-listen_port) |
| `name` | [name](data-sources--udp_loadbalancer--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--udp_loadbalancer--reference.md#schema-namespace) |
| `no_service_policies` | [no_service_policies](data-sources--udp_loadbalancer--properties--no_service_policies.md#section) |
| `origin_pools_weights` | [origin_pools_weights](data-sources--udp_loadbalancer--properties--origin_pools_weights.md#section) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](data-sources--udp_loadbalancer--properties--origin_pools_weights--cluster.md#section) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](data-sources--udp_loadbalancer--properties--origin_pools_weights--cluster.md#schema-origin_pools_weights--cluster--name) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](data-sources--udp_loadbalancer--properties--origin_pools_weights--cluster.md#schema-origin_pools_weights--cluster--namespace) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](data-sources--udp_loadbalancer--properties--origin_pools_weights--cluster.md#schema-origin_pools_weights--cluster--tenant) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](data-sources--udp_loadbalancer--properties--origin_pools_weights--endpoint_subsets.md#section) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](data-sources--udp_loadbalancer--properties--origin_pools_weights--pool.md#section) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](data-sources--udp_loadbalancer--properties--origin_pools_weights--pool.md#schema-origin_pools_weights--pool--name) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](data-sources--udp_loadbalancer--properties--origin_pools_weights--pool.md#schema-origin_pools_weights--pool--namespace) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](data-sources--udp_loadbalancer--properties--origin_pools_weights--pool.md#schema-origin_pools_weights--pool--tenant) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](data-sources--udp_loadbalancer--properties--origin_pools_weights.md#schema-origin_pools_weights--priority) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](data-sources--udp_loadbalancer--properties--origin_pools_weights.md#schema-origin_pools_weights--weight) |
| `port_ranges` | [port_ranges](data-sources--udp_loadbalancer--reference.md#schema-port_ranges) |
| `service_policies_from_namespace` | [service_policies_from_namespace](data-sources--udp_loadbalancer--properties--service_policies_from_namespace.md#section) |
| `udp` | [udp](data-sources--udp_loadbalancer--properties--udp.md#section) |

## Next pages

- [active_service_policies](data-sources--udp_loadbalancer--properties--active_service_policies.md)
- [advertise_custom](data-sources--udp_loadbalancer--properties--advertise_custom.md)
- [advertise_on_public](data-sources--udp_loadbalancer--properties--advertise_on_public.md)
- [advertise_on_public_default_vip](data-sources--udp_loadbalancer--properties--advertise_on_public_default_vip.md)
- [do_not_advertise](data-sources--udp_loadbalancer--properties--do_not_advertise.md)
- [hash_policy_choice_random](data-sources--udp_loadbalancer--properties--hash_policy_choice_random.md)
- [hash_policy_choice_round_robin](data-sources--udp_loadbalancer--properties--hash_policy_choice_round_robin.md)
- [hash_policy_choice_source_ip_stickiness](data-sources--udp_loadbalancer--properties--hash_policy_choice_source_ip_stickiness.md)
- [no_service_policies](data-sources--udp_loadbalancer--properties--no_service_policies.md)
- [origin_pools_weights](data-sources--udp_loadbalancer--properties--origin_pools_weights.md)
- [service_policies_from_namespace](data-sources--udp_loadbalancer--properties--service_policies_from_namespace.md)
- [udp](data-sources--udp_loadbalancer--properties--udp.md)
- [xcsh_udp_loadbalancer](../data-sources/udp_loadbalancer.md)
