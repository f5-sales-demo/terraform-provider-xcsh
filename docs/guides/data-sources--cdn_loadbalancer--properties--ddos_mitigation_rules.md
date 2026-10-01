---
page_title: "ddos_mitigation_rules"
subcategory: "Load Balancing"
description: "ddos_mitigation_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3547, "body_sha256": "sha256:1f2227c4779978dac0e61a0d2363bfa325e1b6b6f6d8fb9e2268f8cf3169675b", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules:block", "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules:ddos_client_source", "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules:ip_prefix_list", "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules:metadata"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:ddos_mitigation_rules", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ddos_mitigation_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/ddos_mitigation_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ddos_mitigation_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_mitigation_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- ddos_mitigation_rules

<a id="section"></a>

Type: `"list"`. Computed.

Define manual mitigation rules to block L7 DDoS attacks.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [block](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--block.md): complete subsection reference.

- [ddos_client_source](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--ddos_client_source.md): complete subsection reference.

<a id="schema-ddos_mitigation_rules--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--ip_prefix_list.md): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--metadata.md): complete subsection reference.

## Next pages

- [ddos_mitigation_rules.block](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--block.md)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--ddos_client_source.md)
- [ddos_mitigation_rules.ip_prefix_list](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--ip_prefix_list.md)
- [ddos_mitigation_rules.metadata](data-sources--cdn_loadbalancer--properties--ddos_mitigation_rules--metadata.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
