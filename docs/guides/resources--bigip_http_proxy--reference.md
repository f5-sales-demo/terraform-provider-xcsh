---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 91652, "body_sha256": "sha256:fee13fabd9d31a8b42328c3e53fef90ed4a55e00efb083a61940766c44b8e551", "canonical_id": "xcsh-docs:resources:bigip_http_proxy:reference", "child_ids": ["xcsh-docs:resources:bigip_http_proxy:properties:advanced_profile", "xcsh-docs:resources:bigip_http_proxy:properties:ddos_profile", "xcsh-docs:resources:bigip_http_proxy:properties:irules", "xcsh-docs:resources:bigip_http_proxy:properties:lb_algorithm", "xcsh-docs:resources:bigip_http_proxy:properties:origin_pools", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_advertisement", "xcsh-docs:resources:bigip_http_proxy:properties:proxy_config", "xcsh-docs:resources:bigip_http_proxy:properties:timeouts"], "collection_id": "xcsh-docs:resources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:bigip_http_proxy:reference", "parent_id": "xcsh-docs:resources:bigip_http_proxy:fundamentals", "path": "docs/guides/resources--bigip_http_proxy--reference.md", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bigip_http_proxy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bigip_http_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
- Property reference

## Direct properties

- [advanced_profile](resources--bigip_http_proxy--properties--advanced_profile.md): complete subsection reference.

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

- [ddos_profile](resources--bigip_http_proxy--properties--ddos_profile.md): complete subsection reference.

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

- [irules](resources--bigip_http_proxy--properties--irules.md): complete subsection reference.

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

- [lb_algorithm](resources--bigip_http_proxy--properties--lb_algorithm.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the BIG-IP HTTP Proxy. Must be unique within the namespace.

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

Namespace where the BIG-IP HTTP Proxy is created.

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

- [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md): complete subsection reference.

- [proxy_advertisement](resources--bigip_http_proxy--properties--proxy_advertisement.md): complete subsection reference.

- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md): complete subsection reference.

- [timeouts](resources--bigip_http_proxy--properties--timeouts.md): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_profile` | [advanced_profile](resources--bigip_http_proxy--properties--advanced_profile.md#section) |
| `advanced_profile.disable_spec` | [advanced_profile.disable_spec](resources--bigip_http_proxy--properties--advanced_profile--disable_spec.md#section) |
| `advanced_profile.enable_default_profile` | [advanced_profile.enable_default_profile](resources--bigip_http_proxy--properties--advanced_profile--enable_default_profile.md#section) |
| `annotations` | [annotations](resources--bigip_http_proxy--reference.md#schema-annotations) |
| `ddos_profile` | [ddos_profile](resources--bigip_http_proxy--properties--ddos_profile.md#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](resources--bigip_http_proxy--properties--ddos_profile--disable_ddos_mitigation.md#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](resources--bigip_http_proxy--properties--ddos_profile--enable_ddos_mitigation.md#section) |
| `description` | [description](resources--bigip_http_proxy--reference.md#schema-description) |
| `disable` | [disable](resources--bigip_http_proxy--reference.md#schema-disable) |
| `id` | [id](resources--bigip_http_proxy--reference.md#schema-id) |
| `irules` | [irules](resources--bigip_http_proxy--properties--irules.md#section) |
| `irules.irules` | [irules.irules](resources--bigip_http_proxy--properties--irules--irules.md#section) |
| `irules.irules.name` | [irules.irules.name](resources--bigip_http_proxy--properties--irules--irules.md#schema-irules--irules--name) |
| `irules.irules.namespace` | [irules.irules.namespace](resources--bigip_http_proxy--properties--irules--irules.md#schema-irules--irules--namespace) |
| `irules.irules.tenant` | [irules.irules.tenant](resources--bigip_http_proxy--properties--irules--irules.md#schema-irules--irules--tenant) |
| `labels` | [labels](resources--bigip_http_proxy--reference.md#schema-labels) |
| `lb_algorithm` | [lb_algorithm](resources--bigip_http_proxy--properties--lb_algorithm.md#section) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](resources--bigip_http_proxy--properties--lb_algorithm--round_robin.md#section) |
| `name` | [name](resources--bigip_http_proxy--reference.md#schema-name) |
| `namespace` | [namespace](resources--bigip_http_proxy--reference.md#schema-namespace) |
| `origin_pools` | [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md#section) |
| `origin_pools.pools` | [origin_pools.pools](resources--bigip_http_proxy--properties--origin_pools--pools.md#section) |
| `origin_pools.pools.name` | [origin_pools.pools.name](resources--bigip_http_proxy--properties--origin_pools--pools.md#schema-origin_pools--pools--name) |
| `origin_pools.pools.origin_servers` | [origin_pools.pools.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md#section) |
| `origin_pools.pools.origin_servers.automatic_port` | [origin_pools.pools.origin_servers.automatic_port](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--automatic_port.md#section) |
| `origin_pools.pools.origin_servers.health_checks` | [origin_pools.pools.origin_servers.health_checks](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check` | [origin_pools.pools.origin_servers.health_checks.health_check](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check.md#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check--icmp_health_check.md#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check.md#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check.md#schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--expected_response) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check.md#schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--send_payload) |
| `origin_pools.pools.origin_servers.health_checks.healthy_threshold` | [origin_pools.pools.origin_servers.health_checks.healthy_threshold](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md#schema-origin_pools--pools--origin_servers--health_checks--healthy_threshold) |
| `origin_pools.pools.origin_servers.health_checks.interval` | [origin_pools.pools.origin_servers.health_checks.interval](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md#schema-origin_pools--pools--origin_servers--health_checks--interval) |
| `origin_pools.pools.origin_servers.health_checks.timeout` | [origin_pools.pools.origin_servers.health_checks.timeout](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md#schema-origin_pools--pools--origin_servers--health_checks--timeout) |
| `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` | [origin_pools.pools.origin_servers.health_checks.unhealthy_threshold](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--health_checks.md#schema-origin_pools--pools--origin_servers--health_checks--unhealthy_threshold) |
| `origin_pools.pools.origin_servers.lb_port` | [origin_pools.pools.origin_servers.lb_port](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--lb_port.md#section) |
| `origin_pools.pools.origin_servers.origin_servers` | [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service` | [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--inside_network.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--outside_network.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--protocol) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--service_name) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site--name) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site--name) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--no_snat_pool.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--k8s_service--vk8s_networks.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--inside_network.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip.ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--ip) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--outside_network.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--segment.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--segment.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--segment--name) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--segment.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--segment--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--segment.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--segment--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site--name) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site--name) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--no_snat_pool.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool.md#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool--prefixes) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_ip.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip.ip](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_ip.md#schema-origin_pools--pools--origin_servers--origin_servers--public_ip--ip) |
| `origin_pools.pools.origin_servers.origin_servers.public_name` | [origin_pools.pools.origin_servers.origin_servers.public_name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_name.md#section) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` | [origin_pools.pools.origin_servers.origin_servers.public_name.dns_name](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_name.md#schema-origin_pools--pools--origin_servers--origin_servers--public_name--dns_name) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` | [origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers--origin_servers--public_name.md#schema-origin_pools--pools--origin_servers--origin_servers--public_name--refresh_interval) |
| `origin_pools.pools.origin_servers.port` | [origin_pools.pools.origin_servers.port](resources--bigip_http_proxy--properties--origin_pools--pools--origin_servers.md#schema-origin_pools--pools--origin_servers--port) |
| `origin_pools.pools.priority` | [origin_pools.pools.priority](resources--bigip_http_proxy--properties--origin_pools--pools.md#schema-origin_pools--pools--priority) |
| `origin_pools.pools.weight` | [origin_pools.pools.weight](resources--bigip_http_proxy--properties--origin_pools--pools.md#schema-origin_pools--pools--weight) |
| `proxy_advertisement` | [proxy_advertisement](resources--bigip_http_proxy--properties--proxy_advertisement.md#section) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#schema-proxy_advertisement--advertise_custom--advertise_where--port) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--use_default_port.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_v6_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_v6_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](resources--bigip_http_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--tenant) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](resources--bigip_http_proxy--properties--proxy_advertisement--do_not_advertise.md#section) |
| `proxy_config` | [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md#section) |
| `proxy_config.domains` | [proxy_config.domains](resources--bigip_http_proxy--properties--proxy_config.md#schema-proxy_config--domains) |
| `proxy_config.http` | [proxy_config.http](resources--bigip_http_proxy--properties--proxy_config--http.md#section) |
| `proxy_config.http.dns_volterra_managed` | [proxy_config.http.dns_volterra_managed](resources--bigip_http_proxy--properties--proxy_config--http.md#schema-proxy_config--http--dns_volterra_managed) |
| `proxy_config.http.port` | [proxy_config.http.port](resources--bigip_http_proxy--properties--proxy_config--http.md#schema-proxy_config--http--port) |
| `proxy_config.http.port_ranges` | [proxy_config.http.port_ranges](resources--bigip_http_proxy--properties--proxy_config--http.md#schema-proxy_config--http--port_ranges) |
| `proxy_config.https` | [proxy_config.https](resources--bigip_http_proxy--properties--proxy_config--https.md#section) |
| `proxy_config.https.add_hsts` | [proxy_config.https.add_hsts](resources--bigip_http_proxy--properties--proxy_config--https.md#schema-proxy_config--https--add_hsts) |
| `proxy_config.https.append_server_name` | [proxy_config.https.append_server_name](resources--bigip_http_proxy--properties--proxy_config--https.md#schema-proxy_config--https--append_server_name) |
| `proxy_config.https.coalescing_options` | [proxy_config.https.coalescing_options](resources--bigip_http_proxy--properties--proxy_config--https--coalescing_options.md#section) |
| `proxy_config.https.coalescing_options.default_coalescing` | [proxy_config.https.coalescing_options.default_coalescing](resources--bigip_http_proxy--properties--proxy_config--https--coalescing_options--default_coalescing.md#section) |
| `proxy_config.https.coalescing_options.strict_coalescing` | [proxy_config.https.coalescing_options.strict_coalescing](resources--bigip_http_proxy--properties--proxy_config--https--coalescing_options--strict_coalescing.md#section) |
| `proxy_config.https.connection_idle_timeout` | [proxy_config.https.connection_idle_timeout](resources--bigip_http_proxy--properties--proxy_config--https.md#schema-proxy_config--https--connection_idle_timeout) |
| `proxy_config.https.default_header` | [proxy_config.https.default_header](resources--bigip_http_proxy--properties--proxy_config--https--default_header.md#section) |
| `proxy_config.https.default_loadbalancer` | [proxy_config.https.default_loadbalancer](resources--bigip_http_proxy--properties--proxy_config--https--default_loadbalancer.md#section) |
| `proxy_config.https.disable_path_normalize` | [proxy_config.https.disable_path_normalize](resources--bigip_http_proxy--properties--proxy_config--https--disable_path_normalize.md#section) |
| `proxy_config.https.enable_path_normalize` | [proxy_config.https.enable_path_normalize](resources--bigip_http_proxy--properties--proxy_config--https--enable_path_normalize.md#section) |
| `proxy_config.https.http_protocol_options` | [proxy_config.https.http_protocol_options](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options.md#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only.md#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v1_v2.md#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--properties--proxy_config--https--http_protocol_options--http_protocol_enable_v2_only.md#section) |
| `proxy_config.https.http_redirect` | [proxy_config.https.http_redirect](resources--bigip_http_proxy--properties--proxy_config--https.md#schema-proxy_config--https--http_redirect) |
| `proxy_config.https.non_default_loadbalancer` | [proxy_config.https.non_default_loadbalancer](resources--bigip_http_proxy--properties--proxy_config--https--non_default_loadbalancer.md#section) |
| `proxy_config.https.pass_through` | [proxy_config.https.pass_through](resources--bigip_http_proxy--properties--proxy_config--https--pass_through.md#section) |
| `proxy_config.https.port` | [proxy_config.https.port](resources--bigip_http_proxy--properties--proxy_config--https.md#schema-proxy_config--https--port) |
| `proxy_config.https.port_ranges` | [proxy_config.https.port_ranges](resources--bigip_http_proxy--properties--proxy_config--https.md#schema-proxy_config--https--port_ranges) |
| `proxy_config.https.server_name` | [proxy_config.https.server_name](resources--bigip_http_proxy--properties--proxy_config--https.md#schema-proxy_config--https--server_name) |
| `proxy_config.https.tls_cert_params` | [proxy_config.https.tls_cert_params](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params.md#section) |
| `proxy_config.https.tls_cert_params.certificates` | [proxy_config.https.tls_cert_params.certificates](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--certificates.md#section) |
| `proxy_config.https.tls_cert_params.certificates.name` | [proxy_config.https.tls_cert_params.certificates.name](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--certificates.md#schema-proxy_config--https--tls_cert_params--certificates--name) |
| `proxy_config.https.tls_cert_params.certificates.namespace` | [proxy_config.https.tls_cert_params.certificates.namespace](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--certificates.md#schema-proxy_config--https--tls_cert_params--certificates--namespace) |
| `proxy_config.https.tls_cert_params.certificates.tenant` | [proxy_config.https.tls_cert_params.certificates.tenant](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--certificates.md#schema-proxy_config--https--tls_cert_params--certificates--tenant) |
| `proxy_config.https.tls_cert_params.no_mtls` | [proxy_config.https.tls_cert_params.no_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--no_mtls.md#section) |
| `proxy_config.https.tls_cert_params.tls_config` | [proxy_config.https.tls_cert_params.tls_config](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config.md#section) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security` | [proxy_config.https.tls_cert_params.tls_config.custom_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md#section) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md#schema-proxy_config--https--tls_cert_params--tls_config--custom_security--cipher_suites) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.max_version](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md#schema-proxy_config--https--tls_cert_params--tls_config--custom_security--max_version) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.min_version](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--custom_security.md#schema-proxy_config--https--tls_cert_params--tls_config--custom_security--min_version) |
| `proxy_config.https.tls_cert_params.tls_config.default_security` | [proxy_config.https.tls_cert_params.tls_config.default_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--default_security.md#section) |
| `proxy_config.https.tls_cert_params.tls_config.low_security` | [proxy_config.https.tls_cert_params.tls_config.low_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--low_security.md#section) |
| `proxy_config.https.tls_cert_params.tls_config.medium_security` | [proxy_config.https.tls_cert_params.tls_config.medium_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--tls_config--medium_security.md#section) |
| `proxy_config.https.tls_cert_params.use_mtls` | [proxy_config.https.tls_cert_params.use_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls.md#section) |
| `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` | [proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls.md#schema-proxy_config--https--tls_cert_params--use_mtls--client_certificate_optional) |
| `proxy_config.https.tls_cert_params.use_mtls.crl` | [proxy_config.https.tls_cert_params.use_mtls.crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--crl.md#section) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.name` | [proxy_config.https.tls_cert_params.use_mtls.crl.name](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--crl.md#schema-proxy_config--https--tls_cert_params--use_mtls--crl--name) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` | [proxy_config.https.tls_cert_params.use_mtls.crl.namespace](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--crl.md#schema-proxy_config--https--tls_cert_params--use_mtls--crl--namespace) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.tenant` | [proxy_config.https.tls_cert_params.use_mtls.crl.tenant](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--crl.md#schema-proxy_config--https--tls_cert_params--use_mtls--crl--tenant) |
| `proxy_config.https.tls_cert_params.use_mtls.no_crl` | [proxy_config.https.tls_cert_params.use_mtls.no_crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--no_crl.md#section) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--trusted_ca.md#section) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--trusted_ca.md#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca--name) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--trusted_ca.md#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca--namespace) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--trusted_ca.md#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca--tenant) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls.md#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca_url) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--xfcc_disabled.md#section) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--xfcc_options.md#section) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--properties--proxy_config--https--tls_cert_params--use_mtls--xfcc_options.md#schema-proxy_config--https--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements) |
| `proxy_config.https.tls_parameters` | [proxy_config.https.tls_parameters](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters.md#section) |
| `proxy_config.https.tls_parameters.no_mtls` | [proxy_config.https.tls_parameters.no_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--no_mtls.md#section) |
| `proxy_config.https.tls_parameters.tls_certificates` | [proxy_config.https.tls_parameters.tls_certificates](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates.md#section) |
| `proxy_config.https.tls_parameters.tls_certificates.certificate_url` | [proxy_config.https.tls_parameters.tls_certificates.certificate_url](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates.md#schema-proxy_config--https--tls_parameters--tls_certificates--certificate_url) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--custom_hash_algorithms.md#section) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--custom_hash_algorithms.md#schema-proxy_config--https--tls_parameters--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `proxy_config.https.tls_parameters.tls_certificates.description_spec` | [proxy_config.https.tls_parameters.tls_certificates.description_spec](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates.md#schema-proxy_config--https--tls_parameters--tls_certificates--description_spec) |
| `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` | [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--disable_ocsp_stapling.md#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key` | [proxy_config.https.tls_parameters.tls_certificates.private_key](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key.md#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info--location) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info.md#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info.md#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info--url) |
| `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` | [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_certificates--use_system_defaults.md#section) |
| `proxy_config.https.tls_parameters.tls_config` | [proxy_config.https.tls_parameters.tls_config](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config.md#section) |
| `proxy_config.https.tls_parameters.tls_config.custom_security` | [proxy_config.https.tls_parameters.tls_config.custom_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--custom_security.md#section) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--custom_security.md#schema-proxy_config--https--tls_parameters--tls_config--custom_security--cipher_suites) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.max_version](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--custom_security.md#schema-proxy_config--https--tls_parameters--tls_config--custom_security--max_version) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.min_version](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--custom_security.md#schema-proxy_config--https--tls_parameters--tls_config--custom_security--min_version) |
| `proxy_config.https.tls_parameters.tls_config.default_security` | [proxy_config.https.tls_parameters.tls_config.default_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--default_security.md#section) |
| `proxy_config.https.tls_parameters.tls_config.low_security` | [proxy_config.https.tls_parameters.tls_config.low_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--low_security.md#section) |
| `proxy_config.https.tls_parameters.tls_config.medium_security` | [proxy_config.https.tls_parameters.tls_config.medium_security](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--tls_config--medium_security.md#section) |
| `proxy_config.https.tls_parameters.use_mtls` | [proxy_config.https.tls_parameters.use_mtls](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls.md#section) |
| `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` | [proxy_config.https.tls_parameters.use_mtls.client_certificate_optional](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls.md#schema-proxy_config--https--tls_parameters--use_mtls--client_certificate_optional) |
| `proxy_config.https.tls_parameters.use_mtls.crl` | [proxy_config.https.tls_parameters.use_mtls.crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--crl.md#section) |
| `proxy_config.https.tls_parameters.use_mtls.crl.name` | [proxy_config.https.tls_parameters.use_mtls.crl.name](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--crl.md#schema-proxy_config--https--tls_parameters--use_mtls--crl--name) |
| `proxy_config.https.tls_parameters.use_mtls.crl.namespace` | [proxy_config.https.tls_parameters.use_mtls.crl.namespace](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--crl.md#schema-proxy_config--https--tls_parameters--use_mtls--crl--namespace) |
| `proxy_config.https.tls_parameters.use_mtls.crl.tenant` | [proxy_config.https.tls_parameters.use_mtls.crl.tenant](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--crl.md#schema-proxy_config--https--tls_parameters--use_mtls--crl--tenant) |
| `proxy_config.https.tls_parameters.use_mtls.no_crl` | [proxy_config.https.tls_parameters.use_mtls.no_crl](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--no_crl.md#section) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--trusted_ca.md#section) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.name](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--trusted_ca.md#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca--name) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--trusted_ca.md#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca--namespace) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--trusted_ca.md#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca--tenant) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca_url](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls.md#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca_url) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` | [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--xfcc_disabled.md#section) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--xfcc_options.md#section) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--properties--proxy_config--https--tls_parameters--use_mtls--xfcc_options.md#schema-proxy_config--https--tls_parameters--use_mtls--xfcc_options--xfcc_header_elements) |
| `proxy_config.https_auto_cert` | [proxy_config.https_auto_cert](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#section) |
| `proxy_config.https_auto_cert.add_hsts` | [proxy_config.https_auto_cert.add_hsts](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#schema-proxy_config--https_auto_cert--add_hsts) |
| `proxy_config.https_auto_cert.append_server_name` | [proxy_config.https_auto_cert.append_server_name](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#schema-proxy_config--https_auto_cert--append_server_name) |
| `proxy_config.https_auto_cert.coalescing_options` | [proxy_config.https_auto_cert.coalescing_options](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options.md#section) |
| `proxy_config.https_auto_cert.coalescing_options.default_coalescing` | [proxy_config.https_auto_cert.coalescing_options.default_coalescing](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options--default_coalescing.md#section) |
| `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` | [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--coalescing_options--strict_coalescing.md#section) |
| `proxy_config.https_auto_cert.connection_idle_timeout` | [proxy_config.https_auto_cert.connection_idle_timeout](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#schema-proxy_config--https_auto_cert--connection_idle_timeout) |
| `proxy_config.https_auto_cert.default_header` | [proxy_config.https_auto_cert.default_header](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--default_header.md#section) |
| `proxy_config.https_auto_cert.default_loadbalancer` | [proxy_config.https_auto_cert.default_loadbalancer](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--default_loadbalancer.md#section) |
| `proxy_config.https_auto_cert.disable_path_normalize` | [proxy_config.https_auto_cert.disable_path_normalize](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--disable_path_normalize.md#section) |
| `proxy_config.https_auto_cert.enable_path_normalize` | [proxy_config.https_auto_cert.enable_path_normalize](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--enable_path_normalize.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options` | [proxy_config.https_auto_cert.http_protocol_options](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--default_header_transformation.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--preserve_case_header_transformation.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options--http_protocol_enable_v1_only--header_transformation--proper_case_header_transformation.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options--http_protocol_enable_v1_v2.md#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--http_protocol_options--http_protocol_enable_v2_only.md#section) |
| `proxy_config.https_auto_cert.http_redirect` | [proxy_config.https_auto_cert.http_redirect](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#schema-proxy_config--https_auto_cert--http_redirect) |
| `proxy_config.https_auto_cert.no_mtls` | [proxy_config.https_auto_cert.no_mtls](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--no_mtls.md#section) |
| `proxy_config.https_auto_cert.non_default_loadbalancer` | [proxy_config.https_auto_cert.non_default_loadbalancer](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--non_default_loadbalancer.md#section) |
| `proxy_config.https_auto_cert.pass_through` | [proxy_config.https_auto_cert.pass_through](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--pass_through.md#section) |
| `proxy_config.https_auto_cert.port` | [proxy_config.https_auto_cert.port](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#schema-proxy_config--https_auto_cert--port) |
| `proxy_config.https_auto_cert.port_ranges` | [proxy_config.https_auto_cert.port_ranges](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#schema-proxy_config--https_auto_cert--port_ranges) |
| `proxy_config.https_auto_cert.server_name` | [proxy_config.https_auto_cert.server_name](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert.md#schema-proxy_config--https_auto_cert--server_name) |
| `proxy_config.https_auto_cert.tls_config` | [proxy_config.https_auto_cert.tls_config](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config.md#section) |
| `proxy_config.https_auto_cert.tls_config.custom_security` | [proxy_config.https_auto_cert.tls_config.custom_security](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--custom_security.md#section) |
| `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` | [proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--custom_security.md#schema-proxy_config--https_auto_cert--tls_config--custom_security--cipher_suites) |
| `proxy_config.https_auto_cert.tls_config.custom_security.max_version` | [proxy_config.https_auto_cert.tls_config.custom_security.max_version](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--custom_security.md#schema-proxy_config--https_auto_cert--tls_config--custom_security--max_version) |
| `proxy_config.https_auto_cert.tls_config.custom_security.min_version` | [proxy_config.https_auto_cert.tls_config.custom_security.min_version](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--custom_security.md#schema-proxy_config--https_auto_cert--tls_config--custom_security--min_version) |
| `proxy_config.https_auto_cert.tls_config.default_security` | [proxy_config.https_auto_cert.tls_config.default_security](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--default_security.md#section) |
| `proxy_config.https_auto_cert.tls_config.low_security` | [proxy_config.https_auto_cert.tls_config.low_security](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--low_security.md#section) |
| `proxy_config.https_auto_cert.tls_config.medium_security` | [proxy_config.https_auto_cert.tls_config.medium_security](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--tls_config--medium_security.md#section) |
| `proxy_config.https_auto_cert.use_mtls` | [proxy_config.https_auto_cert.use_mtls](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls.md#section) |
| `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` | [proxy_config.https_auto_cert.use_mtls.client_certificate_optional](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls.md#schema-proxy_config--https_auto_cert--use_mtls--client_certificate_optional) |
| `proxy_config.https_auto_cert.use_mtls.crl` | [proxy_config.https_auto_cert.use_mtls.crl](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--crl.md#section) |
| `proxy_config.https_auto_cert.use_mtls.crl.name` | [proxy_config.https_auto_cert.use_mtls.crl.name](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--crl.md#schema-proxy_config--https_auto_cert--use_mtls--crl--name) |
| `proxy_config.https_auto_cert.use_mtls.crl.namespace` | [proxy_config.https_auto_cert.use_mtls.crl.namespace](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--crl.md#schema-proxy_config--https_auto_cert--use_mtls--crl--namespace) |
| `proxy_config.https_auto_cert.use_mtls.crl.tenant` | [proxy_config.https_auto_cert.use_mtls.crl.tenant](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--crl.md#schema-proxy_config--https_auto_cert--use_mtls--crl--tenant) |
| `proxy_config.https_auto_cert.use_mtls.no_crl` | [proxy_config.https_auto_cert.use_mtls.no_crl](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--no_crl.md#section) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca` | [proxy_config.https_auto_cert.use_mtls.trusted_ca](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--trusted_ca.md#section) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.name](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--trusted_ca.md#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca--name) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--trusted_ca.md#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca--namespace) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--trusted_ca.md#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca--tenant) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` | [proxy_config.https_auto_cert.use_mtls.trusted_ca_url](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls.md#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca_url) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` | [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--xfcc_disabled.md#section) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options` | [proxy_config.https_auto_cert.use_mtls.xfcc_options](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--xfcc_options.md#section) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](resources--bigip_http_proxy--properties--proxy_config--https_auto_cert--use_mtls--xfcc_options.md#schema-proxy_config--https_auto_cert--use_mtls--xfcc_options--xfcc_header_elements) |
| `timeouts` | [timeouts](resources--bigip_http_proxy--properties--timeouts.md#section) |
| `timeouts.create` | [timeouts.create](resources--bigip_http_proxy--properties--timeouts.md#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](resources--bigip_http_proxy--properties--timeouts.md#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](resources--bigip_http_proxy--properties--timeouts.md#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](resources--bigip_http_proxy--properties--timeouts.md#schema-timeouts--update) |

## Next pages

- [advanced_profile](resources--bigip_http_proxy--properties--advanced_profile.md)
- [ddos_profile](resources--bigip_http_proxy--properties--ddos_profile.md)
- [irules](resources--bigip_http_proxy--properties--irules.md)
- [lb_algorithm](resources--bigip_http_proxy--properties--lb_algorithm.md)
- [origin_pools](resources--bigip_http_proxy--properties--origin_pools.md)
- [proxy_advertisement](resources--bigip_http_proxy--properties--proxy_advertisement.md)
- [proxy_config](resources--bigip_http_proxy--properties--proxy_config.md)
- [timeouts](resources--bigip_http_proxy--properties--timeouts.md)
- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md)
