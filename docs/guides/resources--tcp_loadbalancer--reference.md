---
page_title: "Property reference"
subcategory: "Load Balancing"
description: "Property reference for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 64338, "body_sha256": "sha256:934add54f83049d9b1f2fcf72cf3104676c05d873582499fefa2793de58f758b", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:reference", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:active_service_policies", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_custom", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_on_public", "xcsh-docs:resources:tcp_loadbalancer:properties:advertise_on_public_default_vip", "xcsh-docs:resources:tcp_loadbalancer:properties:default_lb_with_sni", "xcsh-docs:resources:tcp_loadbalancer:properties:do_not_advertise", "xcsh-docs:resources:tcp_loadbalancer:properties:do_not_retract_cluster", "xcsh-docs:resources:tcp_loadbalancer:properties:hash_policy_choice_least_active", "xcsh-docs:resources:tcp_loadbalancer:properties:hash_policy_choice_random", "xcsh-docs:resources:tcp_loadbalancer:properties:hash_policy_choice_round_robin", "xcsh-docs:resources:tcp_loadbalancer:properties:hash_policy_choice_source_ip_stickiness", "xcsh-docs:resources:tcp_loadbalancer:properties:no_service_policies", "xcsh-docs:resources:tcp_loadbalancer:properties:no_sni", "xcsh-docs:resources:tcp_loadbalancer:properties:origin_pools_weights", "xcsh-docs:resources:tcp_loadbalancer:properties:retract_cluster", "xcsh-docs:resources:tcp_loadbalancer:properties:service_policies_from_namespace", "xcsh-docs:resources:tcp_loadbalancer:properties:sni", "xcsh-docs:resources:tcp_loadbalancer:properties:tcp", "xcsh-docs:resources:tcp_loadbalancer:properties:timeouts", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp_auto_cert"], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:reference", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:fundamentals", "path": "docs/guides/resources--tcp_loadbalancer--reference.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- Property reference

## Direct properties

- [active_service_policies](resources--tcp_loadbalancer--properties--active_service_policies.md): complete subsection reference.

- [advertise_custom](resources--tcp_loadbalancer--properties--advertise_custom.md): complete subsection reference.

- [advertise_on_public](resources--tcp_loadbalancer--properties--advertise_on_public.md): complete subsection reference.

- [advertise_on_public_default_vip](resources--tcp_loadbalancer--properties--advertise_on_public_default_vip.md): complete subsection reference.

<a id="schema-annotations"></a>

### annotations property

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

- [default_lb_with_sni](resources--tcp_loadbalancer--properties--default_lb_with_sni.md): complete subsection reference.

<a id="schema-description"></a>

### description property

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

<a id="schema-disable"></a>

### disable property

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

<a id="schema-dns_volterra_managed"></a>

### dns_volterra_managed property

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

- [do_not_advertise](resources--tcp_loadbalancer--properties--do_not_advertise.md): complete subsection reference.

- [do_not_retract_cluster](resources--tcp_loadbalancer--properties--do_not_retract_cluster.md): complete subsection reference.

<a id="schema-domains"></a>

### domains property

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

- [hash_policy_choice_least_active](resources--tcp_loadbalancer--properties--hash_policy_choice_least_active.md): complete subsection reference.

- [hash_policy_choice_random](resources--tcp_loadbalancer--properties--hash_policy_choice_random.md): complete subsection reference.

- [hash_policy_choice_round_robin](resources--tcp_loadbalancer--properties--hash_policy_choice_round_robin.md): complete subsection reference.

- [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--properties--hash_policy_choice_source_ip_stickiness.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="schema-idle_timeout"></a>

### idle_timeout property

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

<a id="schema-labels"></a>

### labels property

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

<a id="schema-listen_port"></a>

### listen_port property

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

- [listen_port](resources--tcp_loadbalancer--reference.md#schema-listen_port)
- [port_ranges](resources--tcp_loadbalancer--reference.md#schema-port_ranges)

Select alternatives according to the provider validators above.

<a id="schema-name"></a>

### name property

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

<a id="schema-namespace"></a>

### namespace property

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

- [no_service_policies](resources--tcp_loadbalancer--properties--no_service_policies.md): complete subsection reference.

- [no_sni](resources--tcp_loadbalancer--properties--no_sni.md): complete subsection reference.

- [origin_pools_weights](resources--tcp_loadbalancer--properties--origin_pools_weights.md): complete subsection reference.

<a id="schema-port_ranges"></a>

### port_ranges property

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

- [retract_cluster](resources--tcp_loadbalancer--properties--retract_cluster.md): complete subsection reference.

- [service_policies_from_namespace](resources--tcp_loadbalancer--properties--service_policies_from_namespace.md): complete subsection reference.

- [sni](resources--tcp_loadbalancer--properties--sni.md): complete subsection reference.

- [tcp](resources--tcp_loadbalancer--properties--tcp.md): complete subsection reference.

- [timeouts](resources--tcp_loadbalancer--properties--timeouts.md): complete subsection reference.

- [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md): complete subsection reference.

- [tls_tcp_auto_cert](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `active_service_policies` | [active_service_policies](resources--tcp_loadbalancer--properties--active_service_policies.md#section) |
| `active_service_policies.policies` | [active_service_policies.policies](resources--tcp_loadbalancer--properties--active_service_policies--policies.md#section) |
| `active_service_policies.policies.name` | [active_service_policies.policies.name](resources--tcp_loadbalancer--properties--active_service_policies--policies.md#schema-active_service_policies--policies--name) |
| `active_service_policies.policies.namespace` | [active_service_policies.policies.namespace](resources--tcp_loadbalancer--properties--active_service_policies--policies.md#schema-active_service_policies--policies--namespace) |
| `active_service_policies.policies.tenant` | [active_service_policies.policies.tenant](resources--tcp_loadbalancer--properties--active_service_policies--policies.md#schema-active_service_policies--policies--tenant) |
| `advertise_custom` | [advertise_custom](resources--tcp_loadbalancer--properties--advertise_custom.md#section) |
| `advertise_custom.advertise_where` | [advertise_custom.advertise_where](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where.md#section) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public` | [advertise_custom.advertise_where.advertise_dualstack_on_public](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public.md#section) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#section) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--namespace) |
| `advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--tenant) |
| `advertise_custom.advertise_where.advertise_on_public` | [advertise_custom.advertise_where.advertise_on_public](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public.md#section) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip` | [advertise_custom.advertise_where.advertise_on_public.public_ip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#section) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_on_public--public_ip--name) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_on_public--public_ip--namespace) |
| `advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_on_public--public_ip--tenant) |
| `advertise_custom.advertise_where.advertise_v6_on_public` | [advertise_custom.advertise_where.advertise_v6_on_public](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public.md#section) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#section) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace) |
| `advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant) |
| `advertise_custom.advertise_where.port` | [advertise_custom.advertise_where.port](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where.md#schema-advertise_custom--advertise_where--port) |
| `advertise_custom.advertise_where.port_ranges` | [advertise_custom.advertise_where.port_ranges](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where.md#schema-advertise_custom--advertise_where--port_ranges) |
| `advertise_custom.advertise_where.site` | [advertise_custom.advertise_where.site](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--site.md#section) |
| `advertise_custom.advertise_where.site.ip` | [advertise_custom.advertise_where.site.ip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--site.md#schema-advertise_custom--advertise_where--site--ip) |
| `advertise_custom.advertise_where.site.network` | [advertise_custom.advertise_where.site.network](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--site.md#schema-advertise_custom--advertise_where--site--network) |
| `advertise_custom.advertise_where.site.site` | [advertise_custom.advertise_where.site.site](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#section) |
| `advertise_custom.advertise_where.site.site.name` | [advertise_custom.advertise_where.site.site.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#schema-advertise_custom--advertise_where--site--site--name) |
| `advertise_custom.advertise_where.site.site.namespace` | [advertise_custom.advertise_where.site.site.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#schema-advertise_custom--advertise_where--site--site--namespace) |
| `advertise_custom.advertise_where.site.site.tenant` | [advertise_custom.advertise_where.site.site.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--site--site.md#schema-advertise_custom--advertise_where--site--site--tenant) |
| `advertise_custom.advertise_where.use_default_port` | [advertise_custom.advertise_where.use_default_port](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--use_default_port.md#section) |
| `advertise_custom.advertise_where.virtual_network` | [advertise_custom.advertise_where.virtual_network](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md#section) |
| `advertise_custom.advertise_where.virtual_network.default_v6_vip` | [advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_v6_vip.md#section) |
| `advertise_custom.advertise_where.virtual_network.default_vip` | [advertise_custom.advertise_where.virtual_network.default_vip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--default_vip.md#section) |
| `advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--specific_v6_vip) |
| `advertise_custom.advertise_where.virtual_network.specific_vip` | [advertise_custom.advertise_where.virtual_network.specific_vip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--specific_vip) |
| `advertise_custom.advertise_where.virtual_network.virtual_network` | [advertise_custom.advertise_where.virtual_network.virtual_network](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#section) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.name` | [advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--virtual_network--name) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--virtual_network--namespace) |
| `advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-advertise_custom--advertise_where--virtual_network--virtual_network--tenant) |
| `advertise_custom.advertise_where.virtual_site` | [advertise_custom.advertise_where.virtual_site](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md#section) |
| `advertise_custom.advertise_where.virtual_site.network` | [advertise_custom.advertise_where.virtual_site.network](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--network) |
| `advertise_custom.advertise_where.virtual_site.virtual_site` | [advertise_custom.advertise_where.virtual_site.virtual_site](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#section) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.name` | [advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--virtual_site--name) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--virtual_site--namespace) |
| `advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site--virtual_site--tenant) |
| `advertise_custom.advertise_where.virtual_site_with_vip` | [advertise_custom.advertise_where.virtual_site_with_vip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md#section) |
| `advertise_custom.advertise_where.virtual_site_with_vip.ip` | [advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--ip) |
| `advertise_custom.advertise_where.virtual_site_with_vip.network` | [advertise_custom.advertise_where.virtual_site_with_vip.network](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--network) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#section) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--namespace) |
| `advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--tenant) |
| `advertise_custom.advertise_where.vk8s_service` | [advertise_custom.advertise_where.vk8s_service](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service.md#section) |
| `advertise_custom.advertise_where.vk8s_service.site` | [advertise_custom.advertise_where.vk8s_service.site](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#section) |
| `advertise_custom.advertise_where.vk8s_service.site.name` | [advertise_custom.advertise_where.vk8s_service.site.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#schema-advertise_custom--advertise_where--vk8s_service--site--name) |
| `advertise_custom.advertise_where.vk8s_service.site.namespace` | [advertise_custom.advertise_where.vk8s_service.site.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#schema-advertise_custom--advertise_where--vk8s_service--site--namespace) |
| `advertise_custom.advertise_where.vk8s_service.site.tenant` | [advertise_custom.advertise_where.vk8s_service.site.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--site.md#schema-advertise_custom--advertise_where--vk8s_service--site--tenant) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site` | [advertise_custom.advertise_where.vk8s_service.virtual_site](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#section) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-advertise_custom--advertise_where--vk8s_service--virtual_site--name) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-advertise_custom--advertise_where--vk8s_service--virtual_site--namespace) |
| `advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--tcp_loadbalancer--properties--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-advertise_custom--advertise_where--vk8s_service--virtual_site--tenant) |
| `advertise_on_public` | [advertise_on_public](resources--tcp_loadbalancer--properties--advertise_on_public.md#section) |
| `advertise_on_public.public_ip` | [advertise_on_public.public_ip](resources--tcp_loadbalancer--properties--advertise_on_public--public_ip.md#section) |
| `advertise_on_public.public_ip.name` | [advertise_on_public.public_ip.name](resources--tcp_loadbalancer--properties--advertise_on_public--public_ip.md#schema-advertise_on_public--public_ip--name) |
| `advertise_on_public.public_ip.namespace` | [advertise_on_public.public_ip.namespace](resources--tcp_loadbalancer--properties--advertise_on_public--public_ip.md#schema-advertise_on_public--public_ip--namespace) |
| `advertise_on_public.public_ip.tenant` | [advertise_on_public.public_ip.tenant](resources--tcp_loadbalancer--properties--advertise_on_public--public_ip.md#schema-advertise_on_public--public_ip--tenant) |
| `advertise_on_public_default_vip` | [advertise_on_public_default_vip](resources--tcp_loadbalancer--properties--advertise_on_public_default_vip.md#section) |
| `annotations` | [annotations](resources--tcp_loadbalancer--reference.md#schema-annotations) |
| `default_lb_with_sni` | [default_lb_with_sni](resources--tcp_loadbalancer--properties--default_lb_with_sni.md#section) |
| `description` | [description](resources--tcp_loadbalancer--reference.md#schema-description) |
| `disable` | [disable](resources--tcp_loadbalancer--reference.md#schema-disable) |
| `dns_volterra_managed` | [dns_volterra_managed](resources--tcp_loadbalancer--reference.md#schema-dns_volterra_managed) |
| `do_not_advertise` | [do_not_advertise](resources--tcp_loadbalancer--properties--do_not_advertise.md#section) |
| `do_not_retract_cluster` | [do_not_retract_cluster](resources--tcp_loadbalancer--properties--do_not_retract_cluster.md#section) |
| `domains` | [domains](resources--tcp_loadbalancer--reference.md#schema-domains) |
| `hash_policy_choice_least_active` | [hash_policy_choice_least_active](resources--tcp_loadbalancer--properties--hash_policy_choice_least_active.md#section) |
| `hash_policy_choice_random` | [hash_policy_choice_random](resources--tcp_loadbalancer--properties--hash_policy_choice_random.md#section) |
| `hash_policy_choice_round_robin` | [hash_policy_choice_round_robin](resources--tcp_loadbalancer--properties--hash_policy_choice_round_robin.md#section) |
| `hash_policy_choice_source_ip_stickiness` | [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--properties--hash_policy_choice_source_ip_stickiness.md#section) |
| `id` | [id](resources--tcp_loadbalancer--reference.md#schema-id) |
| `idle_timeout` | [idle_timeout](resources--tcp_loadbalancer--reference.md#schema-idle_timeout) |
| `labels` | [labels](resources--tcp_loadbalancer--reference.md#schema-labels) |
| `listen_port` | [listen_port](resources--tcp_loadbalancer--reference.md#schema-listen_port) |
| `name` | [name](resources--tcp_loadbalancer--reference.md#schema-name) |
| `namespace` | [namespace](resources--tcp_loadbalancer--reference.md#schema-namespace) |
| `no_service_policies` | [no_service_policies](resources--tcp_loadbalancer--properties--no_service_policies.md#section) |
| `no_sni` | [no_sni](resources--tcp_loadbalancer--properties--no_sni.md#section) |
| `origin_pools_weights` | [origin_pools_weights](resources--tcp_loadbalancer--properties--origin_pools_weights.md#section) |
| `origin_pools_weights.cluster` | [origin_pools_weights.cluster](resources--tcp_loadbalancer--properties--origin_pools_weights--cluster.md#section) |
| `origin_pools_weights.cluster.name` | [origin_pools_weights.cluster.name](resources--tcp_loadbalancer--properties--origin_pools_weights--cluster.md#schema-origin_pools_weights--cluster--name) |
| `origin_pools_weights.cluster.namespace` | [origin_pools_weights.cluster.namespace](resources--tcp_loadbalancer--properties--origin_pools_weights--cluster.md#schema-origin_pools_weights--cluster--namespace) |
| `origin_pools_weights.cluster.tenant` | [origin_pools_weights.cluster.tenant](resources--tcp_loadbalancer--properties--origin_pools_weights--cluster.md#schema-origin_pools_weights--cluster--tenant) |
| `origin_pools_weights.endpoint_subsets` | [origin_pools_weights.endpoint_subsets](resources--tcp_loadbalancer--properties--origin_pools_weights--endpoint_subsets.md#section) |
| `origin_pools_weights.pool` | [origin_pools_weights.pool](resources--tcp_loadbalancer--properties--origin_pools_weights--pool.md#section) |
| `origin_pools_weights.pool.name` | [origin_pools_weights.pool.name](resources--tcp_loadbalancer--properties--origin_pools_weights--pool.md#schema-origin_pools_weights--pool--name) |
| `origin_pools_weights.pool.namespace` | [origin_pools_weights.pool.namespace](resources--tcp_loadbalancer--properties--origin_pools_weights--pool.md#schema-origin_pools_weights--pool--namespace) |
| `origin_pools_weights.pool.tenant` | [origin_pools_weights.pool.tenant](resources--tcp_loadbalancer--properties--origin_pools_weights--pool.md#schema-origin_pools_weights--pool--tenant) |
| `origin_pools_weights.priority` | [origin_pools_weights.priority](resources--tcp_loadbalancer--properties--origin_pools_weights.md#schema-origin_pools_weights--priority) |
| `origin_pools_weights.weight` | [origin_pools_weights.weight](resources--tcp_loadbalancer--properties--origin_pools_weights.md#schema-origin_pools_weights--weight) |
| `port_ranges` | [port_ranges](resources--tcp_loadbalancer--reference.md#schema-port_ranges) |
| `retract_cluster` | [retract_cluster](resources--tcp_loadbalancer--properties--retract_cluster.md#section) |
| `service_policies_from_namespace` | [service_policies_from_namespace](resources--tcp_loadbalancer--properties--service_policies_from_namespace.md#section) |
| `sni` | [sni](resources--tcp_loadbalancer--properties--sni.md#section) |
| `tcp` | [tcp](resources--tcp_loadbalancer--properties--tcp.md#section) |
| `timeouts` | [timeouts](resources--tcp_loadbalancer--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--tcp_loadbalancer--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--tcp_loadbalancer--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--tcp_loadbalancer--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--tcp_loadbalancer--properties--timeouts.md#schema-timeouts--update) |
| `tls_tcp` | [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md#section) |
| `tls_tcp.tls_cert_params` | [tls_tcp.tls_cert_params](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params.md#section) |
| `tls_tcp.tls_cert_params.certificates` | [tls_tcp.tls_cert_params.certificates](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--certificates.md#section) |
| `tls_tcp.tls_cert_params.certificates.name` | [tls_tcp.tls_cert_params.certificates.name](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--certificates.md#schema-tls_tcp--tls_cert_params--certificates--name) |
| `tls_tcp.tls_cert_params.certificates.namespace` | [tls_tcp.tls_cert_params.certificates.namespace](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--certificates.md#schema-tls_tcp--tls_cert_params--certificates--namespace) |
| `tls_tcp.tls_cert_params.certificates.tenant` | [tls_tcp.tls_cert_params.certificates.tenant](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--certificates.md#schema-tls_tcp--tls_cert_params--certificates--tenant) |
| `tls_tcp.tls_cert_params.no_mtls` | [tls_tcp.tls_cert_params.no_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--no_mtls.md#section) |
| `tls_tcp.tls_cert_params.tls_config` | [tls_tcp.tls_cert_params.tls_config](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config.md#section) |
| `tls_tcp.tls_cert_params.tls_config.custom_security` | [tls_tcp.tls_cert_params.tls_config.custom_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--custom_security.md#section) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_cert_params.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--custom_security.md#schema-tls_tcp--tls_cert_params--tls_config--custom_security--cipher_suites) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.max_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.max_version](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--custom_security.md#schema-tls_tcp--tls_cert_params--tls_config--custom_security--max_version) |
| `tls_tcp.tls_cert_params.tls_config.custom_security.min_version` | [tls_tcp.tls_cert_params.tls_config.custom_security.min_version](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--custom_security.md#schema-tls_tcp--tls_cert_params--tls_config--custom_security--min_version) |
| `tls_tcp.tls_cert_params.tls_config.default_security` | [tls_tcp.tls_cert_params.tls_config.default_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--default_security.md#section) |
| `tls_tcp.tls_cert_params.tls_config.low_security` | [tls_tcp.tls_cert_params.tls_config.low_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--low_security.md#section) |
| `tls_tcp.tls_cert_params.tls_config.medium_security` | [tls_tcp.tls_cert_params.tls_config.medium_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--tls_config--medium_security.md#section) |
| `tls_tcp.tls_cert_params.use_mtls` | [tls_tcp.tls_cert_params.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls.md#section) |
| `tls_tcp.tls_cert_params.use_mtls.client_certificate_optional` | [tls_tcp.tls_cert_params.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls.md#schema-tls_tcp--tls_cert_params--use_mtls--client_certificate_optional) |
| `tls_tcp.tls_cert_params.use_mtls.crl` | [tls_tcp.tls_cert_params.use_mtls.crl](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--crl.md#section) |
| `tls_tcp.tls_cert_params.use_mtls.crl.name` | [tls_tcp.tls_cert_params.use_mtls.crl.name](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--crl.md#schema-tls_tcp--tls_cert_params--use_mtls--crl--name) |
| `tls_tcp.tls_cert_params.use_mtls.crl.namespace` | [tls_tcp.tls_cert_params.use_mtls.crl.namespace](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--crl.md#schema-tls_tcp--tls_cert_params--use_mtls--crl--namespace) |
| `tls_tcp.tls_cert_params.use_mtls.crl.tenant` | [tls_tcp.tls_cert_params.use_mtls.crl.tenant](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--crl.md#schema-tls_tcp--tls_cert_params--use_mtls--crl--tenant) |
| `tls_tcp.tls_cert_params.use_mtls.no_crl` | [tls_tcp.tls_cert_params.use_mtls.no_crl](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--no_crl.md#section) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--trusted_ca.md#section) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.name` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--trusted_ca.md#schema-tls_tcp--tls_cert_params--use_mtls--trusted_ca--name) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--trusted_ca.md#schema-tls_tcp--tls_cert_params--use_mtls--trusted_ca--namespace) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--trusted_ca.md#schema-tls_tcp--tls_cert_params--use_mtls--trusted_ca--tenant) |
| `tls_tcp.tls_cert_params.use_mtls.trusted_ca_url` | [tls_tcp.tls_cert_params.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls.md#schema-tls_tcp--tls_cert_params--use_mtls--trusted_ca_url) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_disabled` | [tls_tcp.tls_cert_params.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--xfcc_disabled.md#section) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--xfcc_options.md#section) |
| `tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--properties--tls_tcp--tls_cert_params--use_mtls--xfcc_options.md#schema-tls_tcp--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements) |
| `tls_tcp.tls_parameters` | [tls_tcp.tls_parameters](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters.md#section) |
| `tls_tcp.tls_parameters.no_mtls` | [tls_tcp.tls_parameters.no_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--no_mtls.md#section) |
| `tls_tcp.tls_parameters.tls_certificates` | [tls_tcp.tls_parameters.tls_certificates](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md#section) |
| `tls_tcp.tls_parameters.tls_certificates.certificate_url` | [tls_tcp.tls_parameters.tls_certificates.certificate_url](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md#schema-tls_tcp--tls_parameters--tls_certificates--certificate_url) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--custom_hash_algorithms.md#section) |
| `tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_tcp.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--custom_hash_algorithms.md#schema-tls_tcp--tls_parameters--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `tls_tcp.tls_parameters.tls_certificates.description_spec` | [tls_tcp.tls_parameters.tls_certificates.description_spec](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates.md#schema-tls_tcp--tls_parameters--tls_certificates--description_spec) |
| `tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling` | [tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--disable_ocsp_stapling.md#section) |
| `tls_tcp.tls_parameters.tls_certificates.private_key` | [tls_tcp.tls_parameters.tls_certificates.private_key](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key.md#section) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info--location) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#schema-tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info.md#section) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info.md#schema-tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info.url](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info.md#schema-tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info--url) |
| `tls_tcp.tls_parameters.tls_certificates.use_system_defaults` | [tls_tcp.tls_parameters.tls_certificates.use_system_defaults](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_certificates--use_system_defaults.md#section) |
| `tls_tcp.tls_parameters.tls_config` | [tls_tcp.tls_parameters.tls_config](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config.md#section) |
| `tls_tcp.tls_parameters.tls_config.custom_security` | [tls_tcp.tls_parameters.tls_config.custom_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config--custom_security.md#section) |
| `tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites` | [tls_tcp.tls_parameters.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config--custom_security.md#schema-tls_tcp--tls_parameters--tls_config--custom_security--cipher_suites) |
| `tls_tcp.tls_parameters.tls_config.custom_security.max_version` | [tls_tcp.tls_parameters.tls_config.custom_security.max_version](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config--custom_security.md#schema-tls_tcp--tls_parameters--tls_config--custom_security--max_version) |
| `tls_tcp.tls_parameters.tls_config.custom_security.min_version` | [tls_tcp.tls_parameters.tls_config.custom_security.min_version](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config--custom_security.md#schema-tls_tcp--tls_parameters--tls_config--custom_security--min_version) |
| `tls_tcp.tls_parameters.tls_config.default_security` | [tls_tcp.tls_parameters.tls_config.default_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config--default_security.md#section) |
| `tls_tcp.tls_parameters.tls_config.low_security` | [tls_tcp.tls_parameters.tls_config.low_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config--low_security.md#section) |
| `tls_tcp.tls_parameters.tls_config.medium_security` | [tls_tcp.tls_parameters.tls_config.medium_security](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--tls_config--medium_security.md#section) |
| `tls_tcp.tls_parameters.use_mtls` | [tls_tcp.tls_parameters.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls.md#section) |
| `tls_tcp.tls_parameters.use_mtls.client_certificate_optional` | [tls_tcp.tls_parameters.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls.md#schema-tls_tcp--tls_parameters--use_mtls--client_certificate_optional) |
| `tls_tcp.tls_parameters.use_mtls.crl` | [tls_tcp.tls_parameters.use_mtls.crl](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--crl.md#section) |
| `tls_tcp.tls_parameters.use_mtls.crl.name` | [tls_tcp.tls_parameters.use_mtls.crl.name](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--crl.md#schema-tls_tcp--tls_parameters--use_mtls--crl--name) |
| `tls_tcp.tls_parameters.use_mtls.crl.namespace` | [tls_tcp.tls_parameters.use_mtls.crl.namespace](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--crl.md#schema-tls_tcp--tls_parameters--use_mtls--crl--namespace) |
| `tls_tcp.tls_parameters.use_mtls.crl.tenant` | [tls_tcp.tls_parameters.use_mtls.crl.tenant](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--crl.md#schema-tls_tcp--tls_parameters--use_mtls--crl--tenant) |
| `tls_tcp.tls_parameters.use_mtls.no_crl` | [tls_tcp.tls_parameters.use_mtls.no_crl](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--no_crl.md#section) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca` | [tls_tcp.tls_parameters.use_mtls.trusted_ca](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--trusted_ca.md#section) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.name` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--trusted_ca.md#schema-tls_tcp--tls_parameters--use_mtls--trusted_ca--name) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--trusted_ca.md#schema-tls_tcp--tls_parameters--use_mtls--trusted_ca--namespace) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant` | [tls_tcp.tls_parameters.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--trusted_ca.md#schema-tls_tcp--tls_parameters--use_mtls--trusted_ca--tenant) |
| `tls_tcp.tls_parameters.use_mtls.trusted_ca_url` | [tls_tcp.tls_parameters.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls.md#schema-tls_tcp--tls_parameters--use_mtls--trusted_ca_url) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_disabled` | [tls_tcp.tls_parameters.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--xfcc_disabled.md#section) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options` | [tls_tcp.tls_parameters.use_mtls.xfcc_options](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--xfcc_options.md#section) |
| `tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--properties--tls_tcp--tls_parameters--use_mtls--xfcc_options.md#schema-tls_tcp--tls_parameters--use_mtls--xfcc_options--xfcc_header_elements) |
| `tls_tcp_auto_cert` | [tls_tcp_auto_cert](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md#section) |
| `tls_tcp_auto_cert.no_mtls` | [tls_tcp_auto_cert.no_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--no_mtls.md#section) |
| `tls_tcp_auto_cert.tls_config` | [tls_tcp_auto_cert.tls_config](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config.md#section) |
| `tls_tcp_auto_cert.tls_config.custom_security` | [tls_tcp_auto_cert.tls_config.custom_security](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--custom_security.md#section) |
| `tls_tcp_auto_cert.tls_config.custom_security.cipher_suites` | [tls_tcp_auto_cert.tls_config.custom_security.cipher_suites](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--custom_security.md#schema-tls_tcp_auto_cert--tls_config--custom_security--cipher_suites) |
| `tls_tcp_auto_cert.tls_config.custom_security.max_version` | [tls_tcp_auto_cert.tls_config.custom_security.max_version](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--custom_security.md#schema-tls_tcp_auto_cert--tls_config--custom_security--max_version) |
| `tls_tcp_auto_cert.tls_config.custom_security.min_version` | [tls_tcp_auto_cert.tls_config.custom_security.min_version](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--custom_security.md#schema-tls_tcp_auto_cert--tls_config--custom_security--min_version) |
| `tls_tcp_auto_cert.tls_config.default_security` | [tls_tcp_auto_cert.tls_config.default_security](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--default_security.md#section) |
| `tls_tcp_auto_cert.tls_config.low_security` | [tls_tcp_auto_cert.tls_config.low_security](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--low_security.md#section) |
| `tls_tcp_auto_cert.tls_config.medium_security` | [tls_tcp_auto_cert.tls_config.medium_security](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--tls_config--medium_security.md#section) |
| `tls_tcp_auto_cert.use_mtls` | [tls_tcp_auto_cert.use_mtls](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md#section) |
| `tls_tcp_auto_cert.use_mtls.client_certificate_optional` | [tls_tcp_auto_cert.use_mtls.client_certificate_optional](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md#schema-tls_tcp_auto_cert--use_mtls--client_certificate_optional) |
| `tls_tcp_auto_cert.use_mtls.crl` | [tls_tcp_auto_cert.use_mtls.crl](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--crl.md#section) |
| `tls_tcp_auto_cert.use_mtls.crl.name` | [tls_tcp_auto_cert.use_mtls.crl.name](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--crl.md#schema-tls_tcp_auto_cert--use_mtls--crl--name) |
| `tls_tcp_auto_cert.use_mtls.crl.namespace` | [tls_tcp_auto_cert.use_mtls.crl.namespace](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--crl.md#schema-tls_tcp_auto_cert--use_mtls--crl--namespace) |
| `tls_tcp_auto_cert.use_mtls.crl.tenant` | [tls_tcp_auto_cert.use_mtls.crl.tenant](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--crl.md#schema-tls_tcp_auto_cert--use_mtls--crl--tenant) |
| `tls_tcp_auto_cert.use_mtls.no_crl` | [tls_tcp_auto_cert.use_mtls.no_crl](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--no_crl.md#section) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca` | [tls_tcp_auto_cert.use_mtls.trusted_ca](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--trusted_ca.md#section) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.name` | [tls_tcp_auto_cert.use_mtls.trusted_ca.name](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--trusted_ca.md#schema-tls_tcp_auto_cert--use_mtls--trusted_ca--name) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.namespace` | [tls_tcp_auto_cert.use_mtls.trusted_ca.namespace](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--trusted_ca.md#schema-tls_tcp_auto_cert--use_mtls--trusted_ca--namespace) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca.tenant` | [tls_tcp_auto_cert.use_mtls.trusted_ca.tenant](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--trusted_ca.md#schema-tls_tcp_auto_cert--use_mtls--trusted_ca--tenant) |
| `tls_tcp_auto_cert.use_mtls.trusted_ca_url` | [tls_tcp_auto_cert.use_mtls.trusted_ca_url](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls.md#schema-tls_tcp_auto_cert--use_mtls--trusted_ca_url) |
| `tls_tcp_auto_cert.use_mtls.xfcc_disabled` | [tls_tcp_auto_cert.use_mtls.xfcc_disabled](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--xfcc_disabled.md#section) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options` | [tls_tcp_auto_cert.use_mtls.xfcc_options](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--xfcc_options.md#section) |
| `tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [tls_tcp_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert--use_mtls--xfcc_options.md#schema-tls_tcp_auto_cert--use_mtls--xfcc_options--xfcc_header_elements) |

## Next pages

- [active_service_policies](resources--tcp_loadbalancer--properties--active_service_policies.md)
- [advertise_custom](resources--tcp_loadbalancer--properties--advertise_custom.md)
- [advertise_on_public](resources--tcp_loadbalancer--properties--advertise_on_public.md)
- [advertise_on_public_default_vip](resources--tcp_loadbalancer--properties--advertise_on_public_default_vip.md)
- [default_lb_with_sni](resources--tcp_loadbalancer--properties--default_lb_with_sni.md)
- [do_not_advertise](resources--tcp_loadbalancer--properties--do_not_advertise.md)
- [do_not_retract_cluster](resources--tcp_loadbalancer--properties--do_not_retract_cluster.md)
- [hash_policy_choice_least_active](resources--tcp_loadbalancer--properties--hash_policy_choice_least_active.md)
- [hash_policy_choice_random](resources--tcp_loadbalancer--properties--hash_policy_choice_random.md)
- [hash_policy_choice_round_robin](resources--tcp_loadbalancer--properties--hash_policy_choice_round_robin.md)
- [hash_policy_choice_source_ip_stickiness](resources--tcp_loadbalancer--properties--hash_policy_choice_source_ip_stickiness.md)
- [no_service_policies](resources--tcp_loadbalancer--properties--no_service_policies.md)
- [no_sni](resources--tcp_loadbalancer--properties--no_sni.md)
- [origin_pools_weights](resources--tcp_loadbalancer--properties--origin_pools_weights.md)
- [retract_cluster](resources--tcp_loadbalancer--properties--retract_cluster.md)
- [service_policies_from_namespace](resources--tcp_loadbalancer--properties--service_policies_from_namespace.md)
- [sni](resources--tcp_loadbalancer--properties--sni.md)
- [tcp](resources--tcp_loadbalancer--properties--tcp.md)
- [timeouts](resources--tcp_loadbalancer--properties--timeouts.md)
- [tls_tcp](resources--tcp_loadbalancer--properties--tls_tcp.md)
- [tls_tcp_auto_cert](resources--tcp_loadbalancer--properties--tls_tcp_auto_cert.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
