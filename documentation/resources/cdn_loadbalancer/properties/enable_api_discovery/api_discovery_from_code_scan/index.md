---
page_title: "enable_api_discovery.api_discovery_from_code_scan"
subcategory: "Load Balancing"
description: "Select Code Base and Repositories."
xcsh_docs: {"aliases": ["enable api discovery api discovery from code scan"], "body_bytes": 1916, "body_sha256": "sha256:e9fdce184f5fca8e4fe570585894006250521b443f038ca1be4b32396bb3b1c6", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery", "path": "documentation/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-010.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan:RequiredObjectAttributes:code_base_integrations", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan"], "schema_version": 1, "sections": [{"aliases": ["enable api discovery api discovery from code scan code base integrations"], "anchor": "section", "description": "Configuration parameter for code base integrations", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:all_repos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "enable_api_discovery.api_discovery_from_code_scan.code_base_integrations:ConflictingListObjectAttributes:all_repos,selected_repos", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:enable_api_discovery:api_discovery_from_code_scan:code_base_integrations:selected_repos", "type": "conflicts"}], "schema_path": ["enable_api_discovery", "api_discovery_from_code_scan", "code_base_integrations"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Select Code Base and Repositories.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enable_api_discovery.api_discovery_from_code_scan

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/)
- enable_api_discovery.api_discovery_from_code_scan

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Code Base and Repositories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
```

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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

## Direct properties

- [code_base_integrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/): complete subsection reference.

## Next pages

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/api_discovery_from_code_scan/code_base_integrations/)
- [enable_api_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/enable_api_discovery/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
