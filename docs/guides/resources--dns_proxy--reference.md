---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 50249, "body_sha256": "sha256:327d70193327182ec84860484696e1e42cc52bb610354b3e8f40b31e7711daad", "canonical_id": "xcsh-docs:resources:dns_proxy:reference", "child_ids": ["xcsh-docs:resources:dns_proxy:properties:cache_profile", "xcsh-docs:resources:dns_proxy:properties:ddos_profile", "xcsh-docs:resources:dns_proxy:properties:irules", "xcsh-docs:resources:dns_proxy:properties:lb_algorithm", "xcsh-docs:resources:dns_proxy:properties:origin_servers", "xcsh-docs:resources:dns_proxy:properties:protocol_inspection", "xcsh-docs:resources:dns_proxy:properties:proxy_advertisement", "xcsh-docs:resources:dns_proxy:properties:timeouts"], "collection_id": "xcsh-docs:resources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_proxy:reference", "parent_id": "xcsh-docs:resources:dns_proxy:fundamentals", "path": "docs/guides/resources--dns_proxy--reference.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_proxy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md)
- Property reference

## Direct properties

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

- [cache_profile](resources--dns_proxy--properties--cache_profile.md): complete subsection reference.

- [ddos_profile](resources--dns_proxy--properties--ddos_profile.md): complete subsection reference.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](resources--dns_proxy--properties--irules.md): complete subsection reference.

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

- [lb_algorithm](resources--dns_proxy--properties--lb_algorithm.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the DNS Proxy. Must be unique within the namespace.

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

Type: `"string"`. Optional, Computed.

Namespace for the DNS Proxy. The F5 XC API restricts this resource to the system namespace; it
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

- [origin_servers](resources--dns_proxy--properties--origin_servers.md): complete subsection reference.

- [protocol_inspection](resources--dns_proxy--properties--protocol_inspection.md): complete subsection reference.

- [proxy_advertisement](resources--dns_proxy--properties--proxy_advertisement.md): complete subsection reference.

- [timeouts](resources--dns_proxy--properties--timeouts.md): complete subsection reference.

<a id="schema-transport_type"></a>

### transport_type property

Type: `"string"`. Optional, Computed.

\[Enum: UDP|TCP|BothTCPAndUDP\] Transport Type - UDP: UDP - TCP: TCP - BothTCPAndUDP: Both TCP and
UDP. Possible values are \`UDP\`, \`TCP\`, \`BothTCPAndUDP\`. Defaults to \`UDP\`.

Upstream description:

Transport Type

&#8203;- UDP: UDP

&#8203;- TCP: TCP

&#8203;- BothTCPAndUDP: Both TCP and UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UDP",
    "TCP",
    "BothTCPAndUDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "UDP",
  "enum": [
    "UDP",
    "TCP",
    "BothTCPAndUDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--dns_proxy--reference.md#schema-annotations) |
| `cache_profile` | [cache_profile](resources--dns_proxy--properties--cache_profile.md#section) |
| `cache_profile.cache_size` | [cache_profile.cache_size](resources--dns_proxy--properties--cache_profile.md#schema-cache_profile--cache_size) |
| `cache_profile.disable_cache_profile` | [cache_profile.disable_cache_profile](resources--dns_proxy--properties--cache_profile--disable_cache_profile.md#section) |
| `ddos_profile` | [ddos_profile](resources--dns_proxy--properties--ddos_profile.md#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--dns_proxy--properties--ddos_profile--disable_ddos_mitigation.md#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--dns_proxy--properties--ddos_profile--enable_ddos_mitigation.md#section) |
| `description` | [description](resources--dns_proxy--reference.md#schema-description) |
| `disable` | [disable](resources--dns_proxy--reference.md#schema-disable) |
| `id` | [id](resources--dns_proxy--reference.md#schema-id) |
| `irules` | [irules](resources--dns_proxy--properties--irules.md#section) |
| `irules.name` | [irules.name](resources--dns_proxy--properties--irules.md#schema-irules--name) |
| `irules.namespace` | [irules.namespace](resources--dns_proxy--properties--irules.md#schema-irules--namespace) |
| `irules.tenant` | [irules.tenant](resources--dns_proxy--properties--irules.md#schema-irules--tenant) |
| `labels` | [labels](resources--dns_proxy--reference.md#schema-labels) |
| `lb_algorithm` | [lb_algorithm](resources--dns_proxy--properties--lb_algorithm.md#section) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](resources--dns_proxy--properties--lb_algorithm--round_robin.md#section) |
| `name` | [name](resources--dns_proxy--reference.md#schema-name) |
| `namespace` | [namespace](resources--dns_proxy--reference.md#schema-namespace) |
| `origin_servers` | [origin_servers](resources--dns_proxy--properties--origin_servers.md#section) |
| `origin_servers.health_checks` | [origin_servers.health_checks](resources--dns_proxy--properties--origin_servers--health_checks.md#section) |
| `origin_servers.health_checks.health_check` | [origin_servers.health_checks.health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check.md#section) |
| `origin_servers.health_checks.health_check.dns_health_check` | [origin_servers.health_checks.health_check.dns_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#section) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` | [origin_servers.health_checks.health_check.dns_health_check.expected_rcode](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--expected_rcode) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` | [origin_servers.health_checks.health_check.dns_health_check.expected_record_type](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--expected_record_type) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_response` | [origin_servers.health_checks.health_check.dns_health_check.expected_response](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--expected_response) |
| `origin_servers.health_checks.health_check.dns_health_check.query_name` | [origin_servers.health_checks.health_check.dns_health_check.query_name](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--query_name) |
| `origin_servers.health_checks.health_check.dns_health_check.query_type` | [origin_servers.health_checks.health_check.dns_health_check.query_type](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--query_type) |
| `origin_servers.health_checks.health_check.dns_health_check.reverse` | [origin_servers.health_checks.health_check.dns_health_check.reverse](resources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--reverse) |
| `origin_servers.health_checks.health_check.icmp_health_check` | [origin_servers.health_checks.health_check.icmp_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--icmp_health_check.md#section) |
| `origin_servers.health_checks.health_check.tcp_health_check` | [origin_servers.health_checks.health_check.tcp_health_check](resources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md#section) |
| `origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_servers.health_checks.health_check.tcp_health_check.expected_response](resources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md#schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response) |
| `origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_servers.health_checks.health_check.tcp_health_check.send_payload](resources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md#schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload) |
| `origin_servers.health_checks.healthy_threshold` | [origin_servers.health_checks.healthy_threshold](resources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--healthy_threshold) |
| `origin_servers.health_checks.interval` | [origin_servers.health_checks.interval](resources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--interval) |
| `origin_servers.health_checks.timeout` | [origin_servers.health_checks.timeout](resources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--timeout) |
| `origin_servers.health_checks.unhealthy_threshold` | [origin_servers.health_checks.unhealthy_threshold](resources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--unhealthy_threshold) |
| `origin_servers.origin_servers` | [origin_servers.origin_servers](resources--dns_proxy--properties--origin_servers--origin_servers.md#section) |
| `origin_servers.origin_servers.k8s_service` | [origin_servers.origin_servers.k8s_service](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md#section) |
| `origin_servers.origin_servers.k8s_service.inside_network` | [origin_servers.origin_servers.k8s_service.inside_network](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--inside_network.md#section) |
| `origin_servers.origin_servers.k8s_service.outside_network` | [origin_servers.origin_servers.k8s_service.outside_network](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--outside_network.md#section) |
| `origin_servers.origin_servers.k8s_service.protocol` | [origin_servers.origin_servers.k8s_service.protocol](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md#schema-origin_servers--origin_servers--k8s_service--protocol) |
| `origin_servers.origin_servers.k8s_service.service_name` | [origin_servers.origin_servers.k8s_service.service_name](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md#schema-origin_servers--origin_servers--k8s_service--service_name) |
| `origin_servers.origin_servers.k8s_service.site_locator` | [origin_servers.origin_servers.k8s_service.site_locator](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator.md#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_servers.origin_servers.k8s_service.site_locator.site](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_servers.origin_servers.k8s_service.site_locator.site.name](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--site--name) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.site.namespace](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--site--namespace) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.site.tenant](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--site--tenant) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--name) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--namespace) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--tenant) |
| `origin_servers.origin_servers.k8s_service.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool.md#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--no_snat_pool.md#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md#schema-origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes) |
| `origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_servers.origin_servers.k8s_service.vk8s_networks](resources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--vk8s_networks.md#section) |
| `origin_servers.origin_servers.no_preference` | [origin_servers.origin_servers.no_preference](resources--dns_proxy--properties--origin_servers--origin_servers--no_preference.md#section) |
| `origin_servers.origin_servers.public_ip` | [origin_servers.origin_servers.public_ip](resources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md#section) |
| `origin_servers.origin_servers.public_ip.ip` | [origin_servers.origin_servers.public_ip.ip](resources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md#schema-origin_servers--origin_servers--public_ip--ip) |
| `origin_servers.origin_servers.public_name` | [origin_servers.origin_servers.public_name](resources--dns_proxy--properties--origin_servers--origin_servers--public_name.md#section) |
| `origin_servers.origin_servers.public_name.dns_name` | [origin_servers.origin_servers.public_name.dns_name](resources--dns_proxy--properties--origin_servers--origin_servers--public_name.md#schema-origin_servers--origin_servers--public_name--dns_name) |
| `origin_servers.origin_servers.public_name.refresh_interval` | [origin_servers.origin_servers.public_name.refresh_interval](resources--dns_proxy--properties--origin_servers--origin_servers--public_name.md#schema-origin_servers--origin_servers--public_name--refresh_interval) |
| `origin_servers.origin_servers.site_preferences` | [origin_servers.origin_servers.site_preferences](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences.md#section) |
| `origin_servers.origin_servers.site_preferences.refs` | [origin_servers.origin_servers.site_preferences.refs](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#section) |
| `origin_servers.origin_servers.site_preferences.refs.name` | [origin_servers.origin_servers.site_preferences.refs.name](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#schema-origin_servers--origin_servers--site_preferences--refs--name) |
| `origin_servers.origin_servers.site_preferences.refs.namespace` | [origin_servers.origin_servers.site_preferences.refs.namespace](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#schema-origin_servers--origin_servers--site_preferences--refs--namespace) |
| `origin_servers.origin_servers.site_preferences.refs.tenant` | [origin_servers.origin_servers.site_preferences.refs.tenant](resources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#schema-origin_servers--origin_servers--site_preferences--refs--tenant) |
| `protocol_inspection` | [protocol_inspection](resources--dns_proxy--properties--protocol_inspection.md#section) |
| `protocol_inspection.name` | [protocol_inspection.name](resources--dns_proxy--properties--protocol_inspection.md#schema-protocol_inspection--name) |
| `protocol_inspection.namespace` | [protocol_inspection.namespace](resources--dns_proxy--properties--protocol_inspection.md#schema-protocol_inspection--namespace) |
| `protocol_inspection.tenant` | [protocol_inspection.tenant](resources--dns_proxy--properties--protocol_inspection.md#schema-protocol_inspection--tenant) |
| `proxy_advertisement` | [proxy_advertisement](resources--dns_proxy--properties--proxy_advertisement.md#section) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](resources--dns_proxy--properties--proxy_advertisement--advertise_custom.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#schema-proxy_advertisement--advertise_custom--advertise_where--port) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--use_default_port.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_v6_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_v6_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--tenant) |
| `proxy_advertisement.advertise_dualstack_on_public` | [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public.md#section) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_dualstack_on_public.public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.name](resources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_on_public` | [proxy_advertisement.advertise_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public.md#section) |
| `proxy_advertisement.advertise_on_public.public_ip` | [proxy_advertisement.advertise_on_public.public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_on_public.public_ip.name](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_on_public.public_ip.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_on_public.public_ip.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_on_public_default_dualstack_vip` | [proxy_advertisement.advertise_on_public_default_dualstack_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_dualstack_vip.md#section) |
| `proxy_advertisement.advertise_on_public_default_ipv6_vip` | [proxy_advertisement.advertise_on_public_default_ipv6_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_ipv6_vip.md#section) |
| `proxy_advertisement.advertise_on_public_default_vip` | [proxy_advertisement.advertise_on_public_default_vip](resources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_vip.md#section) |
| `proxy_advertisement.advertise_v6_on_public` | [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public.md#section) |
| `proxy_advertisement.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_v6_on_public.public_ip](resources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_v6_on_public.public_ip.name](resources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_v6_on_public.public_ip.namespace](resources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_v6_on_public.public_ip.tenant](resources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](resources--dns_proxy--properties--proxy_advertisement--do_not_advertise.md#section) |
| `timeouts` | [timeouts](resources--dns_proxy--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--dns_proxy--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--dns_proxy--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--dns_proxy--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--dns_proxy--properties--timeouts.md#schema-timeouts--update) |
| `transport_type` | [transport_type](resources--dns_proxy--reference.md#schema-transport_type) |

## Next pages

- [cache_profile](resources--dns_proxy--properties--cache_profile.md)
- [ddos_profile](resources--dns_proxy--properties--ddos_profile.md)
- [irules](resources--dns_proxy--properties--irules.md)
- [lb_algorithm](resources--dns_proxy--properties--lb_algorithm.md)
- [origin_servers](resources--dns_proxy--properties--origin_servers.md)
- [protocol_inspection](resources--dns_proxy--properties--protocol_inspection.md)
- [proxy_advertisement](resources--dns_proxy--properties--proxy_advertisement.md)
- [timeouts](resources--dns_proxy--properties--timeouts.md)
- [xcsh_dns_proxy](../resources/dns_proxy.md)
