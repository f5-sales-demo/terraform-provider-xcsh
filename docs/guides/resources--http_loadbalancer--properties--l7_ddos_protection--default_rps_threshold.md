---
page_title: "l7_ddos_protection.default_rps_threshold"
subcategory: "Load Balancing"
description: "l7_ddos_protection.default_rps_threshold for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1008, "body_sha256": "sha256:d9a9b51b909e6ff0b7f7363beb564f2e541f7c589d6d60134be914a43b946177", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection:default_rps_threshold", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:l7_ddos_protection", "path": "docs/guides/resources--http_loadbalancer--properties--l7_ddos_protection--default_rps_threshold.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["l7_ddos_protection", "default_rps_threshold"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/l7_ddos_protection/default_rps_threshold/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "l7_ddos_protection.default_rps_threshold for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# l7_ddos_protection.default_rps_threshold

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [l7_ddos_protection](resources--http_loadbalancer--properties--l7_ddos_protection.md)
- l7_ddos_protection.default_rps_threshold

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default rps threshold.

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
default_rps_threshold = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [l7_ddos_protection](resources--http_loadbalancer--properties--l7_ddos_protection.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
