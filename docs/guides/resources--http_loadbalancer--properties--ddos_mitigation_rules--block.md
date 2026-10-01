---
page_title: "ddos_mitigation_rules.block"
subcategory: "Load Balancing"
description: "ddos_mitigation_rules.block for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1058, "body_sha256": "sha256:e064f4f619324b99fe3861e2450833e0e2080b59c95cd79da9041c369edc6007", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:ddos_mitigation_rules:block", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ddos_mitigation_rules:block", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:ddos_mitigation_rules", "path": "docs/guides/resources--http_loadbalancer--properties--ddos_mitigation_rules--block.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ddos_mitigation_rules", "block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ddos_mitigation_rules/block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ddos_mitigation_rules.block for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ddos_mitigation_rules.block

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [ddos_mitigation_rules](resources--http_loadbalancer--properties--ddos_mitigation_rules.md)
- ddos_mitigation_rules.block

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
block {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [ddos_mitigation_rules](resources--http_loadbalancer--properties--ddos_mitigation_rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
