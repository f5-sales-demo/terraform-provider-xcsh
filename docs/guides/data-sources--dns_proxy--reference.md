---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_proxy."
xcsh_docs: {"aliases": [], "body_bytes": 48712, "body_sha256": "sha256:4305a2e92363dd77a8c4aa96cbe914c804efea242e96e68185bb3c27fbc50361", "canonical_id": "xcsh-docs:data-sources:dns_proxy:reference", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:cache_profile", "xcsh-docs:data-sources:dns_proxy:properties:ddos_profile", "xcsh-docs:data-sources:dns_proxy:properties:irules", "xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "xcsh-docs:data-sources:dns_proxy:properties:protocol_inspection", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement"], "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:reference", "parent_id": "xcsh-docs:data-sources:dns_proxy:fundamentals", "path": "docs/guides/data-sources--dns_proxy--reference.md", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_dns_proxy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
- Property reference

## Direct properties

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

- [cache_profile](data-sources--dns_proxy--properties--cache_profile.md): complete subsection reference.

- [ddos_profile](data-sources--dns_proxy--properties--ddos_profile.md): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the DNSProxy.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](data-sources--dns_proxy--properties--irules.md): complete subsection reference.

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

- [lb_algorithm](data-sources--dns_proxy--properties--lb_algorithm.md): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the DNSProxy.

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

Type: `"string"`. Optional, Computed.

Namespace where the DNSProxy exists.

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

- [origin_servers](data-sources--dns_proxy--properties--origin_servers.md): complete subsection reference.

- [protocol_inspection](data-sources--dns_proxy--properties--protocol_inspection.md): complete subsection reference.

- [proxy_advertisement](data-sources--dns_proxy--properties--proxy_advertisement.md): complete subsection reference.

<a id="schema-transport_type"></a>

### transport_type property

Type: `"string"`. Computed.

\[Enum: UDP|TCP|BothTCPAndUDP\] Transport Type - UDP: UDP - TCP: TCP - BothTCPAndUDP: Both TCP and
UDP. Possible values are \`UDP\`, \`TCP\`, \`BothTCPAndUDP\`. Defaults to \`UDP\`.

Upstream description:

Transport Type

&#8203;- UDP: UDP

&#8203;- TCP: TCP

&#8203;- BothTCPAndUDP: Both TCP and UDP.

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
| `annotations` | [annotations](data-sources--dns_proxy--reference.md#schema-annotations) |
| `cache_profile` | [cache_profile](data-sources--dns_proxy--properties--cache_profile.md#section) |
| `cache_profile.cache_size` | [cache_profile.cache_size](data-sources--dns_proxy--properties--cache_profile.md#schema-cache_profile--cache_size) |
| `cache_profile.disable_cache_profile` | [cache_profile.disable_cache_profile](data-sources--dns_proxy--properties--cache_profile--disable_cache_profile.md#section) |
| `ddos_profile` | [ddos_profile](data-sources--dns_proxy--properties--ddos_profile.md#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--dns_proxy--properties--ddos_profile--disable_ddos_mitigation.md#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--dns_proxy--properties--ddos_profile--enable_ddos_mitigation.md#section) |
| `description` | [description](data-sources--dns_proxy--reference.md#schema-description) |
| `id` | [id](data-sources--dns_proxy--reference.md#schema-id) |
| `irules` | [irules](data-sources--dns_proxy--properties--irules.md#section) |
| `irules.name` | [irules.name](data-sources--dns_proxy--properties--irules.md#schema-irules--name) |
| `irules.namespace` | [irules.namespace](data-sources--dns_proxy--properties--irules.md#schema-irules--namespace) |
| `irules.tenant` | [irules.tenant](data-sources--dns_proxy--properties--irules.md#schema-irules--tenant) |
| `labels` | [labels](data-sources--dns_proxy--reference.md#schema-labels) |
| `lb_algorithm` | [lb_algorithm](data-sources--dns_proxy--properties--lb_algorithm.md#section) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](data-sources--dns_proxy--properties--lb_algorithm--round_robin.md#section) |
| `name` | [name](data-sources--dns_proxy--reference.md#schema-name) |
| `namespace` | [namespace](data-sources--dns_proxy--reference.md#schema-namespace) |
| `origin_servers` | [origin_servers](data-sources--dns_proxy--properties--origin_servers.md#section) |
| `origin_servers.health_checks` | [origin_servers.health_checks](data-sources--dns_proxy--properties--origin_servers--health_checks.md#section) |
| `origin_servers.health_checks.health_check` | [origin_servers.health_checks.health_check](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check.md#section) |
| `origin_servers.health_checks.health_check.dns_health_check` | [origin_servers.health_checks.health_check.dns_health_check](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#section) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` | [origin_servers.health_checks.health_check.dns_health_check.expected_rcode](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--expected_rcode) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` | [origin_servers.health_checks.health_check.dns_health_check.expected_record_type](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--expected_record_type) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_response` | [origin_servers.health_checks.health_check.dns_health_check.expected_response](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--expected_response) |
| `origin_servers.health_checks.health_check.dns_health_check.query_name` | [origin_servers.health_checks.health_check.dns_health_check.query_name](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--query_name) |
| `origin_servers.health_checks.health_check.dns_health_check.query_type` | [origin_servers.health_checks.health_check.dns_health_check.query_type](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--query_type) |
| `origin_servers.health_checks.health_check.dns_health_check.reverse` | [origin_servers.health_checks.health_check.dns_health_check.reverse](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--dns_health_check.md#schema-origin_servers--health_checks--health_check--dns_health_check--reverse) |
| `origin_servers.health_checks.health_check.icmp_health_check` | [origin_servers.health_checks.health_check.icmp_health_check](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--icmp_health_check.md#section) |
| `origin_servers.health_checks.health_check.tcp_health_check` | [origin_servers.health_checks.health_check.tcp_health_check](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md#section) |
| `origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_servers.health_checks.health_check.tcp_health_check.expected_response](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md#schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response) |
| `origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_servers.health_checks.health_check.tcp_health_check.send_payload](data-sources--dns_proxy--properties--origin_servers--health_checks--health_check--tcp_health_check.md#schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload) |
| `origin_servers.health_checks.healthy_threshold` | [origin_servers.health_checks.healthy_threshold](data-sources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--healthy_threshold) |
| `origin_servers.health_checks.interval` | [origin_servers.health_checks.interval](data-sources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--interval) |
| `origin_servers.health_checks.timeout` | [origin_servers.health_checks.timeout](data-sources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--timeout) |
| `origin_servers.health_checks.unhealthy_threshold` | [origin_servers.health_checks.unhealthy_threshold](data-sources--dns_proxy--properties--origin_servers--health_checks.md#schema-origin_servers--health_checks--unhealthy_threshold) |
| `origin_servers.origin_servers` | [origin_servers.origin_servers](data-sources--dns_proxy--properties--origin_servers--origin_servers.md#section) |
| `origin_servers.origin_servers.k8s_service` | [origin_servers.origin_servers.k8s_service](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md#section) |
| `origin_servers.origin_servers.k8s_service.inside_network` | [origin_servers.origin_servers.k8s_service.inside_network](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--inside_network.md#section) |
| `origin_servers.origin_servers.k8s_service.outside_network` | [origin_servers.origin_servers.k8s_service.outside_network](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--outside_network.md#section) |
| `origin_servers.origin_servers.k8s_service.protocol` | [origin_servers.origin_servers.k8s_service.protocol](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md#schema-origin_servers--origin_servers--k8s_service--protocol) |
| `origin_servers.origin_servers.k8s_service.service_name` | [origin_servers.origin_servers.k8s_service.service_name](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service.md#schema-origin_servers--origin_servers--k8s_service--service_name) |
| `origin_servers.origin_servers.k8s_service.site_locator` | [origin_servers.origin_servers.k8s_service.site_locator](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator.md#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_servers.origin_servers.k8s_service.site_locator.site.name](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--site--name) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.site.namespace](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--site--namespace) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.site.tenant](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--site--tenant) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--name) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--namespace) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--site_locator--virtual_site.md#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--tenant) |
| `origin_servers.origin_servers.k8s_service.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool.md#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--no_snat_pool.md#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool.md#schema-origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes) |
| `origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--dns_proxy--properties--origin_servers--origin_servers--k8s_service--vk8s_networks.md#section) |
| `origin_servers.origin_servers.no_preference` | [origin_servers.origin_servers.no_preference](data-sources--dns_proxy--properties--origin_servers--origin_servers--no_preference.md#section) |
| `origin_servers.origin_servers.public_ip` | [origin_servers.origin_servers.public_ip](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md#section) |
| `origin_servers.origin_servers.public_ip.ip` | [origin_servers.origin_servers.public_ip.ip](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_ip.md#schema-origin_servers--origin_servers--public_ip--ip) |
| `origin_servers.origin_servers.public_name` | [origin_servers.origin_servers.public_name](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_name.md#section) |
| `origin_servers.origin_servers.public_name.dns_name` | [origin_servers.origin_servers.public_name.dns_name](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_name.md#schema-origin_servers--origin_servers--public_name--dns_name) |
| `origin_servers.origin_servers.public_name.refresh_interval` | [origin_servers.origin_servers.public_name.refresh_interval](data-sources--dns_proxy--properties--origin_servers--origin_servers--public_name.md#schema-origin_servers--origin_servers--public_name--refresh_interval) |
| `origin_servers.origin_servers.site_preferences` | [origin_servers.origin_servers.site_preferences](data-sources--dns_proxy--properties--origin_servers--origin_servers--site_preferences.md#section) |
| `origin_servers.origin_servers.site_preferences.refs` | [origin_servers.origin_servers.site_preferences.refs](data-sources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#section) |
| `origin_servers.origin_servers.site_preferences.refs.name` | [origin_servers.origin_servers.site_preferences.refs.name](data-sources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#schema-origin_servers--origin_servers--site_preferences--refs--name) |
| `origin_servers.origin_servers.site_preferences.refs.namespace` | [origin_servers.origin_servers.site_preferences.refs.namespace](data-sources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#schema-origin_servers--origin_servers--site_preferences--refs--namespace) |
| `origin_servers.origin_servers.site_preferences.refs.tenant` | [origin_servers.origin_servers.site_preferences.refs.tenant](data-sources--dns_proxy--properties--origin_servers--origin_servers--site_preferences--refs.md#schema-origin_servers--origin_servers--site_preferences--refs--tenant) |
| `protocol_inspection` | [protocol_inspection](data-sources--dns_proxy--properties--protocol_inspection.md#section) |
| `protocol_inspection.name` | [protocol_inspection.name](data-sources--dns_proxy--properties--protocol_inspection.md#schema-protocol_inspection--name) |
| `protocol_inspection.namespace` | [protocol_inspection.namespace](data-sources--dns_proxy--properties--protocol_inspection.md#schema-protocol_inspection--namespace) |
| `protocol_inspection.tenant` | [protocol_inspection.tenant](data-sources--dns_proxy--properties--protocol_inspection.md#schema-protocol_inspection--tenant) |
| `proxy_advertisement` | [proxy_advertisement](data-sources--dns_proxy--properties--proxy_advertisement.md#section) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#schema-proxy_advertisement--advertise_custom--advertise_where--port) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where.md#schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--site--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--use_default_port.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_v6_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--default_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_v6_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site.md#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--tenant) |
| `proxy_advertisement.advertise_dualstack_on_public` | [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public.md#section) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_dualstack_on_public--public_ip.md#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_on_public` | [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public.md#section) |
| `proxy_advertisement.advertise_on_public.public_ip` | [proxy_advertisement.advertise_on_public.public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_on_public.public_ip.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_on_public.public_ip.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_on_public.public_ip.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public--public_ip.md#schema-proxy_advertisement--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_on_public_default_dualstack_vip` | [proxy_advertisement.advertise_on_public_default_dualstack_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_dualstack_vip.md#section) |
| `proxy_advertisement.advertise_on_public_default_ipv6_vip` | [proxy_advertisement.advertise_on_public_default_ipv6_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_ipv6_vip.md#section) |
| `proxy_advertisement.advertise_on_public_default_vip` | [proxy_advertisement.advertise_on_public_default_vip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_on_public_default_vip.md#section) |
| `proxy_advertisement.advertise_v6_on_public` | [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public.md#section) |
| `proxy_advertisement.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_v6_on_public.public_ip](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#section) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_v6_on_public.public_ip.name](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_v6_on_public.public_ip.namespace](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_v6_on_public.public_ip.tenant](data-sources--dns_proxy--properties--proxy_advertisement--advertise_v6_on_public--public_ip.md#schema-proxy_advertisement--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](data-sources--dns_proxy--properties--proxy_advertisement--do_not_advertise.md#section) |
| `transport_type` | [transport_type](data-sources--dns_proxy--reference.md#schema-transport_type) |

## Next pages

- [cache_profile](data-sources--dns_proxy--properties--cache_profile.md)
- [ddos_profile](data-sources--dns_proxy--properties--ddos_profile.md)
- [irules](data-sources--dns_proxy--properties--irules.md)
- [lb_algorithm](data-sources--dns_proxy--properties--lb_algorithm.md)
- [origin_servers](data-sources--dns_proxy--properties--origin_servers.md)
- [protocol_inspection](data-sources--dns_proxy--properties--protocol_inspection.md)
- [proxy_advertisement](data-sources--dns_proxy--properties--proxy_advertisement.md)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md)
