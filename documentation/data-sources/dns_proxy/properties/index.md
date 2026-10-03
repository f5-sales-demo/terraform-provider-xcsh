---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_dns_proxy."
xcsh_docs: {"aliases": ["dns proxy"], "body_bytes": 56917, "body_sha256": "sha256:f855c44b2705cc836c1859760b41148e134b6f5b0ef6a05292a0c8e1fda20df3", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_proxy:properties:cache_profile", "xcsh-docs:data-sources:dns_proxy:properties:ddos_profile", "xcsh-docs:data-sources:dns_proxy:properties:irules", "xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm", "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "xcsh-docs:data-sources:dns_proxy:properties:protocol_inspection", "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_proxy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_proxy:reference", "parent_id": "xcsh-docs:data-sources:dns_proxy:fundamentals", "path": "documentation/data-sources/dns_proxy/properties/index.md", "product": "distributed-cloud", "provider_name": "dns_proxy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020", "registry_path": "docs/guides/data-sources--dns_proxy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:dns_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["cache profile"], "anchor": "section", "description": "DNS Cache specifies cache configuration.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:cache_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["cache_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["ddos profile"], "anchor": "section", "description": "DDoS Protection Rule for DNS.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:ddos_profile", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ddos_profile"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:dns_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:dns_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["irules"], "anchor": "section", "description": "OPTIONS for attaching iRules to DNS proxy.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:irules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["irules"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:dns_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["lb algorithm"], "anchor": "section", "description": "Load Balancing Algorithm Type.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:lb_algorithm", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["lb_algorithm"], "syntax": "attribute", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:dns_proxy:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:dns_proxy:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "origin servers", "upstream servers"], "anchor": "section", "description": "List of origin Servers for the DNS proxy.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:origin_servers", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["origin_servers"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol inspection"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:protocol_inspection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["protocol_inspection"], "syntax": "attribute", "type": "object"}, {"aliases": ["proxy advertisement"], "anchor": "section", "description": "Proxy Advertisement Type.", "document_id": "xcsh-docs:data-sources:dns_proxy:properties:proxy_advertisement", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["proxy_advertisement"], "syntax": "attribute", "type": "object"}, {"aliases": ["transport type"], "anchor": "schema-transport_type", "description": "Transport Type - UDP: UDP - TCP: TCP - BothTCPAndUDP: Both TCP and UDP.", "document_id": "xcsh-docs:data-sources:dns_proxy:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["transport_type"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_proxy/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_dns_proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["dns_proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
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

- [cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/cache_profile/): complete subsection reference.

- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/): complete subsection reference.

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

- [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/irules/): complete subsection reference.

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

- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/lb_algorithm/): complete subsection reference.

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

- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/): complete subsection reference.

- [protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/protocol_inspection/): complete subsection reference.

- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/): complete subsection reference.

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
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/#schema-annotations) |
| `cache_profile` | [cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/cache_profile/#section) |
| `cache_profile.cache_size` | [cache_profile.cache_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/cache_profile/#schema-cache_profile--cache_size) |
| `cache_profile.disable_cache_profile` | [cache_profile.disable_cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/cache_profile/disable_cache_profile/#section) |
| `ddos_profile` | [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/#section) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/disable_ddos_mitigation/#section) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/enable_ddos_mitigation/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/#schema-description) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/#schema-id) |
| `irules` | [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/irules/#section) |
| `irules.name` | [irules.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/irules/#schema-irules--name) |
| `irules.namespace` | [irules.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/irules/#schema-irules--namespace) |
| `irules.tenant` | [irules.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/irules/#schema-irules--tenant) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/#schema-labels) |
| `lb_algorithm` | [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/lb_algorithm/#section) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/lb_algorithm/round_robin/#section) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/#schema-namespace) |
| `origin_servers` | [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/#section) |
| `origin_servers.health_checks` | [origin_servers.health_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/#section) |
| `origin_servers.health_checks.health_check` | [origin_servers.health_checks.health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/#section) |
| `origin_servers.health_checks.health_check.dns_health_check` | [origin_servers.health_checks.health_check.dns_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/#section) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_rcode` | [origin_servers.health_checks.health_check.dns_health_check.expected_rcode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/#schema-origin_servers--health_checks--health_check--dns_health_check--expected_rcode) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_record_type` | [origin_servers.health_checks.health_check.dns_health_check.expected_record_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/#schema-origin_servers--health_checks--health_check--dns_health_check--expected_record_type) |
| `origin_servers.health_checks.health_check.dns_health_check.expected_response` | [origin_servers.health_checks.health_check.dns_health_check.expected_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/#schema-origin_servers--health_checks--health_check--dns_health_check--expected_response) |
| `origin_servers.health_checks.health_check.dns_health_check.query_name` | [origin_servers.health_checks.health_check.dns_health_check.query_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/#schema-origin_servers--health_checks--health_check--dns_health_check--query_name) |
| `origin_servers.health_checks.health_check.dns_health_check.query_type` | [origin_servers.health_checks.health_check.dns_health_check.query_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/#schema-origin_servers--health_checks--health_check--dns_health_check--query_type) |
| `origin_servers.health_checks.health_check.dns_health_check.reverse` | [origin_servers.health_checks.health_check.dns_health_check.reverse](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/dns_health_check/#schema-origin_servers--health_checks--health_check--dns_health_check--reverse) |
| `origin_servers.health_checks.health_check.icmp_health_check` | [origin_servers.health_checks.health_check.icmp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/icmp_health_check/#section) |
| `origin_servers.health_checks.health_check.tcp_health_check` | [origin_servers.health_checks.health_check.tcp_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/#section) |
| `origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_servers.health_checks.health_check.tcp_health_check.expected_response](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/#schema-origin_servers--health_checks--health_check--tcp_health_check--expected_response) |
| `origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_servers.health_checks.health_check.tcp_health_check.send_payload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/health_check/tcp_health_check/#schema-origin_servers--health_checks--health_check--tcp_health_check--send_payload) |
| `origin_servers.health_checks.healthy_threshold` | [origin_servers.health_checks.healthy_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/#schema-origin_servers--health_checks--healthy_threshold) |
| `origin_servers.health_checks.interval` | [origin_servers.health_checks.interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/#schema-origin_servers--health_checks--interval) |
| `origin_servers.health_checks.timeout` | [origin_servers.health_checks.timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/#schema-origin_servers--health_checks--timeout) |
| `origin_servers.health_checks.unhealthy_threshold` | [origin_servers.health_checks.unhealthy_threshold](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/health_checks/#schema-origin_servers--health_checks--unhealthy_threshold) |
| `origin_servers.origin_servers` | [origin_servers.origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/#section) |
| `origin_servers.origin_servers.k8s_service` | [origin_servers.origin_servers.k8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/#section) |
| `origin_servers.origin_servers.k8s_service.inside_network` | [origin_servers.origin_servers.k8s_service.inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/inside_network/#section) |
| `origin_servers.origin_servers.k8s_service.outside_network` | [origin_servers.origin_servers.k8s_service.outside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/outside_network/#section) |
| `origin_servers.origin_servers.k8s_service.protocol` | [origin_servers.origin_servers.k8s_service.protocol](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/#schema-origin_servers--origin_servers--k8s_service--protocol) |
| `origin_servers.origin_servers.k8s_service.service_name` | [origin_servers.origin_servers.k8s_service.service_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/#schema-origin_servers--origin_servers--k8s_service--service_name) |
| `origin_servers.origin_servers.k8s_service.site_locator` | [origin_servers.origin_servers.k8s_service.site_locator](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_servers.origin_servers.k8s_service.site_locator.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/site/#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_servers.origin_servers.k8s_service.site_locator.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/site/#schema-origin_servers--origin_servers--k8s_service--site_locator--site--name) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/site/#schema-origin_servers--origin_servers--k8s_service--site_locator--site--namespace) |
| `origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/site/#schema-origin_servers--origin_servers--k8s_service--site_locator--site--tenant) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#section) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--name) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--namespace) |
| `origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/site_locator/virtual_site/#schema-origin_servers--origin_servers--k8s_service--site_locator--virtual_site--tenant) |
| `origin_servers.origin_servers.k8s_service.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/no_snat_pool/#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/#section) |
| `origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/snat_pool/snat_pool/#schema-origin_servers--origin_servers--k8s_service--snat_pool--snat_pool--prefixes) |
| `origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_servers.origin_servers.k8s_service.vk8s_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/k8s_service/vk8s_networks/#section) |
| `origin_servers.origin_servers.no_preference` | [origin_servers.origin_servers.no_preference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/no_preference/#section) |
| `origin_servers.origin_servers.public_ip` | [origin_servers.origin_servers.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_ip/#section) |
| `origin_servers.origin_servers.public_ip.ip` | [origin_servers.origin_servers.public_ip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_ip/#schema-origin_servers--origin_servers--public_ip--ip) |
| `origin_servers.origin_servers.public_name` | [origin_servers.origin_servers.public_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_name/#section) |
| `origin_servers.origin_servers.public_name.dns_name` | [origin_servers.origin_servers.public_name.dns_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_name/#schema-origin_servers--origin_servers--public_name--dns_name) |
| `origin_servers.origin_servers.public_name.refresh_interval` | [origin_servers.origin_servers.public_name.refresh_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/public_name/#schema-origin_servers--origin_servers--public_name--refresh_interval) |
| `origin_servers.origin_servers.site_preferences` | [origin_servers.origin_servers.site_preferences](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/#section) |
| `origin_servers.origin_servers.site_preferences.refs` | [origin_servers.origin_servers.site_preferences.refs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/#section) |
| `origin_servers.origin_servers.site_preferences.refs.name` | [origin_servers.origin_servers.site_preferences.refs.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/#schema-origin_servers--origin_servers--site_preferences--refs--name) |
| `origin_servers.origin_servers.site_preferences.refs.namespace` | [origin_servers.origin_servers.site_preferences.refs.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/#schema-origin_servers--origin_servers--site_preferences--refs--namespace) |
| `origin_servers.origin_servers.site_preferences.refs.tenant` | [origin_servers.origin_servers.site_preferences.refs.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/origin_servers/site_preferences/refs/#schema-origin_servers--origin_servers--site_preferences--refs--tenant) |
| `protocol_inspection` | [protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/protocol_inspection/#section) |
| `protocol_inspection.name` | [protocol_inspection.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/protocol_inspection/#schema-protocol_inspection--name) |
| `protocol_inspection.namespace` | [protocol_inspection.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/protocol_inspection/#schema-protocol_inspection--namespace) |
| `protocol_inspection.tenant` | [protocol_inspection.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/protocol_inspection/#schema-protocol_inspection--tenant) |
| `proxy_advertisement` | [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/#section) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/#section) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_custom--advertise_where--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/#schema-proxy_advertisement--advertise_custom--advertise_where--port) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/#schema-proxy_advertisement--advertise_custom--advertise_where--port_ranges) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/site/site/#schema-proxy_advertisement--advertise_custom--advertise_where--site--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/use_default_port/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/default_v6_vip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/default_vip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_v6_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--specific_vip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_network/virtual_network/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_network--virtual_network--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--ip) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--network) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/virtual_site_with_vip/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--virtual_site_with_vip--virtual_site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--site--tenant) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#section) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--name) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--namespace) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_custom/advertise_where/vk8s_service/virtual_site/#schema-proxy_advertisement--advertise_custom--advertise_where--vk8s_service--virtual_site--tenant) |
| `proxy_advertisement.advertise_dualstack_on_public` | [proxy_advertisement.advertise_dualstack_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/#section) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_dualstack_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--name) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_dualstack_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_dualstack_on_public/public_ip/#schema-proxy_advertisement--advertise_dualstack_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_on_public` | [proxy_advertisement.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public/#section) |
| `proxy_advertisement.advertise_on_public.public_ip` | [proxy_advertisement.advertise_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_on_public--public_ip--name) |
| `proxy_advertisement.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public/public_ip/#schema-proxy_advertisement--advertise_on_public--public_ip--tenant) |
| `proxy_advertisement.advertise_on_public_default_dualstack_vip` | [proxy_advertisement.advertise_on_public_default_dualstack_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_dualstack_vip/#section) |
| `proxy_advertisement.advertise_on_public_default_ipv6_vip` | [proxy_advertisement.advertise_on_public_default_ipv6_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_ipv6_vip/#section) |
| `proxy_advertisement.advertise_on_public_default_vip` | [proxy_advertisement.advertise_on_public_default_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_on_public_default_vip/#section) |
| `proxy_advertisement.advertise_v6_on_public` | [proxy_advertisement.advertise_v6_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/#section) |
| `proxy_advertisement.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_v6_on_public.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/public_ip/#section) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_v6_on_public.public_ip.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_v6_on_public--public_ip--name) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_v6_on_public.public_ip.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_v6_on_public--public_ip--namespace) |
| `proxy_advertisement.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_v6_on_public.public_ip.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/advertise_v6_on_public/public_ip/#schema-proxy_advertisement--advertise_v6_on_public--public_ip--tenant) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/do_not_advertise/#section) |
| `transport_type` | [transport_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/#schema-transport_type) |

## Next pages

- [cache_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/cache_profile/)
- [ddos_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/ddos_profile/)
- [irules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/irules/)
- [lb_algorithm](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/lb_algorithm/)
- [origin_servers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/origin_servers/)
- [protocol_inspection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/protocol_inspection/)
- [proxy_advertisement](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/properties/proxy_advertisement/)
- [xcsh_dns_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_proxy/)
