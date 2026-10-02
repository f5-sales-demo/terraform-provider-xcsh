---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_bigip_http_proxy."
xcsh_docs: {"aliases": ["bigip http proxy"], "body_bytes": 105753, "body_sha256": "sha256:1726c6a2a36d2a49ebdaa3f1a4a9540dd19df17d1e87368bf33054f0b00148f0", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:bigip_http_proxy:properties:advanced_profile", "xcsh-docs:data-sources:bigip_http_proxy:properties:ddos_profile", "xcsh-docs:data-sources:bigip_http_proxy:properties:irules", "xcsh-docs:data-sources:bigip_http_proxy:properties:lb_algorithm", "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement", "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_http_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "parent_id": "xcsh-docs:data-sources:bigip_http_proxy:fundamentals", "path": "documentation/data-sources/bigip_http_proxy/properties/index.md", "product": "distributed-cloud", "provider_name": "bigip_http_proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002", "registry_path": "docs/guides/data-sources--bigip_http_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["advanced profile"], "anchor": "section", "description": "This defines various advanced Profile OPTIONS for a Loadbalancer.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:advanced_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["advanced_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["ddos profile"], "anchor": "section", "description": "BIG-IP DDoS Protection Rules.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:ddos_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ddos_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["irules"], "anchor": "section", "description": "IRules Configuration for downstream connections.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:irules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["irules"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["lb algorithm"], "anchor": "section", "description": "Load Balancing Algorithm Type.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:lb_algorithm", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["lb_algorithm"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "origin pools", "origin servers", "upstream servers"], "anchor": "section", "description": "List of Origin Pools.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:origin_pools", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_pools"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement"], "anchor": "section", "description": "Proxy Advertisement Type.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_advertisement", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy config"], "anchor": "section", "description": "HTTP/HTTPS Load balancer.", "document_id": "xcsh-docs:data-sources:bigip_http_proxy:properties:proxy_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_config"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_http_proxy/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_bigip_http_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bigip_http_proxyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
- Property reference

## Direct properties

- [advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/advanced_profile/): complete subsection reference.

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

- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/ddos_profile/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the BigIPHTTPProxy.

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

- [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/irules/): complete subsection reference.

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

- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/lb_algorithm/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the BigIPHTTPProxy.

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

Namespace where the BigIPHTTPProxy exists.

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

- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/): complete subsection reference.

- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/): complete subsection reference.

- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_profile` | [advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/advanced_profile/#section) |
| `advanced_profile.disable_spec` | [advanced_profile.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/advanced_profile/disable_spec/#section) |
| `advanced_profile.enable_default_profile` | [advanced_profile.enable_default_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/advanced_profile/enable_default_profile/#section) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/#schema-annotations) |
| `ddos_profile` | [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/ddos_profile/#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/ddos_profile/disable_ddos_mitigation/#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/ddos_profile/enable_ddos_mitigation/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/#schema-id) |
| `irules` | [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/irules/#section) |
| `irules.irules` | [irules.irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/irules/irules/#section) |
| `irules.irules.name` | [irules.irules.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/irules/irules/#schema-irules--irules--name) |
| `irules.irules.namespace` | [irules.irules.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/irules/irules/#schema-irules--irules--namespace) |
| `irules.irules.tenant` | [irules.irules.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/irules/irules/#schema-irules--irules--tenant) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/#schema-labels) |
| `lb_algorithm` | [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/lb_algorithm/#section) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/lb_algorithm/round_robin/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/#schema-namespace) |
| `origin_pools` | [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/#section) |
| `origin_pools.pools` | [origin_pools.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/#section) |
| `origin_pools.pools.name` | [origin_pools.pools.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/#schema-origin_pools--pools--name) |
| `origin_pools.pools.origin_servers` | [origin_pools.pools.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/#section) |
| `origin_pools.pools.origin_servers.automatic_port` | [origin_pools.pools.origin_servers.automatic_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/automatic_port/#section) |
| `origin_pools.pools.origin_servers.health_checks` | [origin_pools.pools.origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check` | [origin_pools.pools.origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/icmp_health_check/#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/tcp_health_check/#section) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/tcp_health_check/#schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--expected_response) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/health_check/tcp_health_check/#schema-origin_pools--pools--origin_servers--health_checks--health_check--tcp_health_check--send_payload) |
| `origin_pools.pools.origin_servers.health_checks.healthy_threshold` | [origin_pools.pools.origin_servers.health_checks.healthy_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/#schema-origin_pools--pools--origin_servers--health_checks--healthy_threshold) |
| `origin_pools.pools.origin_servers.health_checks.interval` | [origin_pools.pools.origin_servers.health_checks.interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/#schema-origin_pools--pools--origin_servers--health_checks--interval) |
| `origin_pools.pools.origin_servers.health_checks.timeout` | [origin_pools.pools.origin_servers.health_checks.timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/#schema-origin_pools--pools--origin_servers--health_checks--timeout) |
| `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` | [origin_pools.pools.origin_servers.health_checks.unhealthy_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/health_checks/#schema-origin_pools--pools--origin_servers--health_checks--unhealthy_threshold) |
| `origin_pools.pools.origin_servers.lb_port` | [origin_pools.pools.origin_servers.lb_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/lb_port/#section) |
| `origin_pools.pools.origin_servers.origin_servers` | [origin_pools.pools.origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service` | [origin_pools.pools.origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/inside_network/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/outside_network/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--protocol) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--service_name) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/site/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/site/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site--name) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/site/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/site/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site--name) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--site_locator--virtual_site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/no_snat_pool/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/#section) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/#schema-origin_pools--pools--origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/k8s_service/vk8s_networks/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/inside_network/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--ip) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/outside_network/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/segment/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/segment/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--segment--name) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/segment/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--segment--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/segment/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--segment--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/site/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/site/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site--name) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/site/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/site/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/virtual_site/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/virtual_site/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site--name) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/virtual_site/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site--namespace) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/site_locator/virtual_site/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--site_locator--virtual_site--tenant) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/no_snat_pool/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/snat_pool/#section) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/private_ip/snat_pool/snat_pool/#schema-origin_pools--pools--origin_servers--origin_servers--private_ip--snat_pool--snat_pool--prefixes) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/public_ip/#section) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/public_ip/#schema-origin_pools--pools--origin_servers--origin_servers--public_ip--ip) |
| `origin_pools.pools.origin_servers.origin_servers.public_name` | [origin_pools.pools.origin_servers.origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/public_name/#section) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` | [origin_pools.pools.origin_servers.origin_servers.public_name.dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/public_name/#schema-origin_pools--pools--origin_servers--origin_servers--public_name--dns_name) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` | [origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/origin_servers/public_name/#schema-origin_pools--pools--origin_servers--origin_servers--public_name--refresh_interval) |
| `origin_pools.pools.origin_servers.port` | [origin_pools.pools.origin_servers.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/origin_servers/#schema-origin_pools--pools--origin_servers--port) |
| `origin_pools.pools.priority` | [origin_pools.pools.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/#schema-origin_pools--pools--priority) |
| `origin_pools.pools.weight` | [origin_pools.pools.weight](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/pools/#schema-origin_pools--pools--weight) |
| `proxy_advertisement` | [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/#section) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/#section) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/#schema-proxy_advertisement--advertise_custom--advertise_where--port) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/#schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/use_default_port/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/default_v6_vip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/default_vip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_v6_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--tenant) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/do_not_advertise/#section) |
| `proxy_config` | [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/#section) |
| `proxy_config.domains` | [proxy_config.domains](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/#schema-proxy_config--domains) |
| `proxy_config.http` | [proxy_config.http](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/http/#section) |
| `proxy_config.http.dns_volterra_managed` | [proxy_config.http.dns_volterra_managed](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/http/#schema-proxy_config--http--dns_volterra_managed) |
| `proxy_config.http.port` | [proxy_config.http.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/http/#schema-proxy_config--http--port) |
| `proxy_config.http.port_ranges` | [proxy_config.http.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/http/#schema-proxy_config--http--port_ranges) |
| `proxy_config.https` | [proxy_config.https](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#section) |
| `proxy_config.https.add_hsts` | [proxy_config.https.add_hsts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#schema-proxy_config--https--add_hsts) |
| `proxy_config.https.append_server_name` | [proxy_config.https.append_server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#schema-proxy_config--https--append_server_name) |
| `proxy_config.https.coalescing_options` | [proxy_config.https.coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/#section) |
| `proxy_config.https.coalescing_options.default_coalescing` | [proxy_config.https.coalescing_options.default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/default_coalescing/#section) |
| `proxy_config.https.coalescing_options.strict_coalescing` | [proxy_config.https.coalescing_options.strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/coalescing_options/strict_coalescing/#section) |
| `proxy_config.https.connection_idle_timeout` | [proxy_config.https.connection_idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#schema-proxy_config--https--connection_idle_timeout) |
| `proxy_config.https.default_header` | [proxy_config.https.default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/default_header/#section) |
| `proxy_config.https.default_loadbalancer` | [proxy_config.https.default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/default_loadbalancer/#section) |
| `proxy_config.https.disable_path_normalize` | [proxy_config.https.disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/disable_path_normalize/#section) |
| `proxy_config.https.enable_path_normalize` | [proxy_config.https.enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/enable_path_normalize/#section) |
| `proxy_config.https.http_protocol_options` | [proxy_config.https.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v1_v2/#section) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/http_protocol_options/http_protocol_enable_v2_only/#section) |
| `proxy_config.https.http_redirect` | [proxy_config.https.http_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#schema-proxy_config--https--http_redirect) |
| `proxy_config.https.non_default_loadbalancer` | [proxy_config.https.non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/non_default_loadbalancer/#section) |
| `proxy_config.https.pass_through` | [proxy_config.https.pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/pass_through/#section) |
| `proxy_config.https.port` | [proxy_config.https.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#schema-proxy_config--https--port) |
| `proxy_config.https.port_ranges` | [proxy_config.https.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#schema-proxy_config--https--port_ranges) |
| `proxy_config.https.server_name` | [proxy_config.https.server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/#schema-proxy_config--https--server_name) |
| `proxy_config.https.tls_cert_params` | [proxy_config.https.tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/#section) |
| `proxy_config.https.tls_cert_params.certificates` | [proxy_config.https.tls_cert_params.certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/certificates/#section) |
| `proxy_config.https.tls_cert_params.certificates.name` | [proxy_config.https.tls_cert_params.certificates.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/certificates/#schema-proxy_config--https--tls_cert_params--certificates--name) |
| `proxy_config.https.tls_cert_params.certificates.namespace` | [proxy_config.https.tls_cert_params.certificates.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/certificates/#schema-proxy_config--https--tls_cert_params--certificates--namespace) |
| `proxy_config.https.tls_cert_params.certificates.tenant` | [proxy_config.https.tls_cert_params.certificates.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/certificates/#schema-proxy_config--https--tls_cert_params--certificates--tenant) |
| `proxy_config.https.tls_cert_params.no_mtls` | [proxy_config.https.tls_cert_params.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/no_mtls/#section) |
| `proxy_config.https.tls_cert_params.tls_config` | [proxy_config.https.tls_cert_params.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/#section) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security` | [proxy_config.https.tls_cert_params.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/custom_security/#section) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/custom_security/#schema-proxy_config--https--tls_cert_params--tls_config--custom_security--cipher_suites) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.max_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/custom_security/#schema-proxy_config--https--tls_cert_params--tls_config--custom_security--max_version) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.min_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/custom_security/#schema-proxy_config--https--tls_cert_params--tls_config--custom_security--min_version) |
| `proxy_config.https.tls_cert_params.tls_config.default_security` | [proxy_config.https.tls_cert_params.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/default_security/#section) |
| `proxy_config.https.tls_cert_params.tls_config.low_security` | [proxy_config.https.tls_cert_params.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/low_security/#section) |
| `proxy_config.https.tls_cert_params.tls_config.medium_security` | [proxy_config.https.tls_cert_params.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/tls_config/medium_security/#section) |
| `proxy_config.https.tls_cert_params.use_mtls` | [proxy_config.https.tls_cert_params.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/#section) |
| `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` | [proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/#schema-proxy_config--https--tls_cert_params--use_mtls--client_certificate_optional) |
| `proxy_config.https.tls_cert_params.use_mtls.crl` | [proxy_config.https.tls_cert_params.use_mtls.crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/crl/#section) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.name` | [proxy_config.https.tls_cert_params.use_mtls.crl.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/crl/#schema-proxy_config--https--tls_cert_params--use_mtls--crl--name) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` | [proxy_config.https.tls_cert_params.use_mtls.crl.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/crl/#schema-proxy_config--https--tls_cert_params--use_mtls--crl--namespace) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.tenant` | [proxy_config.https.tls_cert_params.use_mtls.crl.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/crl/#schema-proxy_config--https--tls_cert_params--use_mtls--crl--tenant) |
| `proxy_config.https.tls_cert_params.use_mtls.no_crl` | [proxy_config.https.tls_cert_params.use_mtls.no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/no_crl/#section) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/trusted_ca/#section) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/trusted_ca/#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca--name) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/trusted_ca/#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca--namespace) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/trusted_ca/#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca--tenant) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/#schema-proxy_config--https--tls_cert_params--use_mtls--trusted_ca_url) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_disabled/#section) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_options/#section) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_cert_params/use_mtls/xfcc_options/#schema-proxy_config--https--tls_cert_params--use_mtls--xfcc_options--xfcc_header_elements) |
| `proxy_config.https.tls_parameters` | [proxy_config.https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/#section) |
| `proxy_config.https.tls_parameters.no_mtls` | [proxy_config.https.tls_parameters.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/no_mtls/#section) |
| `proxy_config.https.tls_parameters.tls_certificates` | [proxy_config.https.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/#section) |
| `proxy_config.https.tls_parameters.tls_certificates.certificate_url` | [proxy_config.https.tls_parameters.tls_certificates.certificate_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/#schema-proxy_config--https--tls_parameters--tls_certificates--certificate_url) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/custom_hash_algorithms/#section) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/custom_hash_algorithms/#schema-proxy_config--https--tls_parameters--tls_certificates--custom_hash_algorithms--hash_algorithms) |
| `proxy_config.https.tls_parameters.tls_certificates.description_spec` | [proxy_config.https.tls_parameters.tls_certificates.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/#schema-proxy_config--https--tls_parameters--tls_certificates--description_spec) |
| `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` | [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/disable_ocsp_stapling/#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key` | [proxy_config.https.tls_parameters.tls_certificates.private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/blindfold_secret_info/#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/blindfold_secret_info/#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info--decryption_provider) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/blindfold_secret_info/#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info--location) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/blindfold_secret_info/#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--blindfold_secret_info--store_provider) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/clear_secret_info/#section) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/clear_secret_info/#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info--provider_ref) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/private_key/clear_secret_info/#schema-proxy_config--https--tls_parameters--tls_certificates--private_key--clear_secret_info--url) |
| `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` | [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_certificates/use_system_defaults/#section) |
| `proxy_config.https.tls_parameters.tls_config` | [proxy_config.https.tls_parameters.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/#section) |
| `proxy_config.https.tls_parameters.tls_config.custom_security` | [proxy_config.https.tls_parameters.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/custom_security/#section) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/custom_security/#schema-proxy_config--https--tls_parameters--tls_config--custom_security--cipher_suites) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.max_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/custom_security/#schema-proxy_config--https--tls_parameters--tls_config--custom_security--max_version) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.min_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/custom_security/#schema-proxy_config--https--tls_parameters--tls_config--custom_security--min_version) |
| `proxy_config.https.tls_parameters.tls_config.default_security` | [proxy_config.https.tls_parameters.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/default_security/#section) |
| `proxy_config.https.tls_parameters.tls_config.low_security` | [proxy_config.https.tls_parameters.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/low_security/#section) |
| `proxy_config.https.tls_parameters.tls_config.medium_security` | [proxy_config.https.tls_parameters.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/tls_config/medium_security/#section) |
| `proxy_config.https.tls_parameters.use_mtls` | [proxy_config.https.tls_parameters.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/#section) |
| `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` | [proxy_config.https.tls_parameters.use_mtls.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/#schema-proxy_config--https--tls_parameters--use_mtls--client_certificate_optional) |
| `proxy_config.https.tls_parameters.use_mtls.crl` | [proxy_config.https.tls_parameters.use_mtls.crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/crl/#section) |
| `proxy_config.https.tls_parameters.use_mtls.crl.name` | [proxy_config.https.tls_parameters.use_mtls.crl.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/crl/#schema-proxy_config--https--tls_parameters--use_mtls--crl--name) |
| `proxy_config.https.tls_parameters.use_mtls.crl.namespace` | [proxy_config.https.tls_parameters.use_mtls.crl.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/crl/#schema-proxy_config--https--tls_parameters--use_mtls--crl--namespace) |
| `proxy_config.https.tls_parameters.use_mtls.crl.tenant` | [proxy_config.https.tls_parameters.use_mtls.crl.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/crl/#schema-proxy_config--https--tls_parameters--use_mtls--crl--tenant) |
| `proxy_config.https.tls_parameters.use_mtls.no_crl` | [proxy_config.https.tls_parameters.use_mtls.no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/no_crl/#section) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/trusted_ca/#section) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/trusted_ca/#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca--name) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/trusted_ca/#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca--namespace) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/trusted_ca/#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca--tenant) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/#schema-proxy_config--https--tls_parameters--use_mtls--trusted_ca_url) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` | [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/xfcc_disabled/#section) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/xfcc_options/#section) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https/tls_parameters/use_mtls/xfcc_options/#schema-proxy_config--https--tls_parameters--use_mtls--xfcc_options--xfcc_header_elements) |
| `proxy_config.https_auto_cert` | [proxy_config.https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#section) |
| `proxy_config.https_auto_cert.add_hsts` | [proxy_config.https_auto_cert.add_hsts](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#schema-proxy_config--https_auto_cert--add_hsts) |
| `proxy_config.https_auto_cert.append_server_name` | [proxy_config.https_auto_cert.append_server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#schema-proxy_config--https_auto_cert--append_server_name) |
| `proxy_config.https_auto_cert.coalescing_options` | [proxy_config.https_auto_cert.coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/coalescing_options/#section) |
| `proxy_config.https_auto_cert.coalescing_options.default_coalescing` | [proxy_config.https_auto_cert.coalescing_options.default_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/coalescing_options/default_coalescing/#section) |
| `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` | [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/coalescing_options/strict_coalescing/#section) |
| `proxy_config.https_auto_cert.connection_idle_timeout` | [proxy_config.https_auto_cert.connection_idle_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#schema-proxy_config--https_auto_cert--connection_idle_timeout) |
| `proxy_config.https_auto_cert.default_header` | [proxy_config.https_auto_cert.default_header](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/default_header/#section) |
| `proxy_config.https_auto_cert.default_loadbalancer` | [proxy_config.https_auto_cert.default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/default_loadbalancer/#section) |
| `proxy_config.https_auto_cert.disable_path_normalize` | [proxy_config.https_auto_cert.disable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/disable_path_normalize/#section) |
| `proxy_config.https_auto_cert.enable_path_normalize` | [proxy_config.https_auto_cert.enable_path_normalize](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/enable_path_normalize/#section) |
| `proxy_config.https_auto_cert.http_protocol_options` | [proxy_config.https_auto_cert.http_protocol_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/default_header_transformation/#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/preserve_case_header_transformation/#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_only/header_transformation/proper_case_header_transformation/#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v1_v2/#section) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/http_protocol_options/http_protocol_enable_v2_only/#section) |
| `proxy_config.https_auto_cert.http_redirect` | [proxy_config.https_auto_cert.http_redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#schema-proxy_config--https_auto_cert--http_redirect) |
| `proxy_config.https_auto_cert.no_mtls` | [proxy_config.https_auto_cert.no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/no_mtls/#section) |
| `proxy_config.https_auto_cert.non_default_loadbalancer` | [proxy_config.https_auto_cert.non_default_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/non_default_loadbalancer/#section) |
| `proxy_config.https_auto_cert.pass_through` | [proxy_config.https_auto_cert.pass_through](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/pass_through/#section) |
| `proxy_config.https_auto_cert.port` | [proxy_config.https_auto_cert.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#schema-proxy_config--https_auto_cert--port) |
| `proxy_config.https_auto_cert.port_ranges` | [proxy_config.https_auto_cert.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#schema-proxy_config--https_auto_cert--port_ranges) |
| `proxy_config.https_auto_cert.server_name` | [proxy_config.https_auto_cert.server_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/#schema-proxy_config--https_auto_cert--server_name) |
| `proxy_config.https_auto_cert.tls_config` | [proxy_config.https_auto_cert.tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/#section) |
| `proxy_config.https_auto_cert.tls_config.custom_security` | [proxy_config.https_auto_cert.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/custom_security/#section) |
| `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` | [proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/custom_security/#schema-proxy_config--https_auto_cert--tls_config--custom_security--cipher_suites) |
| `proxy_config.https_auto_cert.tls_config.custom_security.max_version` | [proxy_config.https_auto_cert.tls_config.custom_security.max_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/custom_security/#schema-proxy_config--https_auto_cert--tls_config--custom_security--max_version) |
| `proxy_config.https_auto_cert.tls_config.custom_security.min_version` | [proxy_config.https_auto_cert.tls_config.custom_security.min_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/custom_security/#schema-proxy_config--https_auto_cert--tls_config--custom_security--min_version) |
| `proxy_config.https_auto_cert.tls_config.default_security` | [proxy_config.https_auto_cert.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/default_security/#section) |
| `proxy_config.https_auto_cert.tls_config.low_security` | [proxy_config.https_auto_cert.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/low_security/#section) |
| `proxy_config.https_auto_cert.tls_config.medium_security` | [proxy_config.https_auto_cert.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/tls_config/medium_security/#section) |
| `proxy_config.https_auto_cert.use_mtls` | [proxy_config.https_auto_cert.use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/#section) |
| `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` | [proxy_config.https_auto_cert.use_mtls.client_certificate_optional](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/#schema-proxy_config--https_auto_cert--use_mtls--client_certificate_optional) |
| `proxy_config.https_auto_cert.use_mtls.crl` | [proxy_config.https_auto_cert.use_mtls.crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/crl/#section) |
| `proxy_config.https_auto_cert.use_mtls.crl.name` | [proxy_config.https_auto_cert.use_mtls.crl.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/crl/#schema-proxy_config--https_auto_cert--use_mtls--crl--name) |
| `proxy_config.https_auto_cert.use_mtls.crl.namespace` | [proxy_config.https_auto_cert.use_mtls.crl.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/crl/#schema-proxy_config--https_auto_cert--use_mtls--crl--namespace) |
| `proxy_config.https_auto_cert.use_mtls.crl.tenant` | [proxy_config.https_auto_cert.use_mtls.crl.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/crl/#schema-proxy_config--https_auto_cert--use_mtls--crl--tenant) |
| `proxy_config.https_auto_cert.use_mtls.no_crl` | [proxy_config.https_auto_cert.use_mtls.no_crl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/no_crl/#section) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca` | [proxy_config.https_auto_cert.use_mtls.trusted_ca](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/trusted_ca/#section) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/trusted_ca/#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca--name) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/trusted_ca/#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca--namespace) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/trusted_ca/#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca--tenant) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` | [proxy_config.https_auto_cert.use_mtls.trusted_ca_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/#schema-proxy_config--https_auto_cert--use_mtls--trusted_ca_url) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` | [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_disabled/#section) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options` | [proxy_config.https_auto_cert.use_mtls.xfcc_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_options/#section) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/https_auto_cert/use_mtls/xfcc_options/#schema-proxy_config--https_auto_cert--use_mtls--xfcc_options--xfcc_header_elements) |

## Next pages

- [advanced_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/advanced_profile/)
- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/ddos_profile/)
- [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/irules/)
- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/lb_algorithm/)
- [origin_pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/origin_pools/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_advertisement/)
- [proxy_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/properties/proxy_config/)
- [xcsh_bigip_http_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_http_proxy/)
