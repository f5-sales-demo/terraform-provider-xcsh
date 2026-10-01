---
page_title: "api_protection_rules"
subcategory: "Load Balancing"
description: "api_protection_rules for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1442, "body_sha256": "sha256:2d3b58f78274b9effa65cabec59bc76af1cb4b5d4e8d495657e6e85e6a234006", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_endpoint_rules", "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules:api_groups_rules"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_protection_rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--api_protection_rules.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_protection_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_protection_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_protection_rules for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_protection_rules

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- api_protection_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

API Protection Rules. API Protection Rules.

Upstream description:

API Protection Rules.

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
api_protection_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [api_endpoint_rules](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules.md): complete subsection reference.

- [api_groups_rules](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md): complete subsection reference.

## Next pages

- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--properties--api_protection_rules--api_endpoint_rules.md)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--properties--api_protection_rules--api_groups_rules.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
