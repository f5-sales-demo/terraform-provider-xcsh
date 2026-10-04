---
page_title: "tgw_security.active_east_west_service_policies"
subcategory: ""
description: "Active service policies for the east-west proxy."
xcsh_docs: {"aliases": ["tgw security active east west service policies"], "body_bytes": 1667, "body_sha256": "sha256:dc0fa60f2a63586e34afdc42a75613d5b42e3e5346e01c1667dbb8aecd404226", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies:service_policies"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security", "path": "documentation/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tgw_security", "active_east_west_service_policies"], "schema_version": 1, "sections": [{"aliases": ["tgw security active east west service policies service policies"], "anchor": "section", "description": "A list of references to service_policy objects.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies:service_policies", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-tgw_security--active_east_west_service_policies--service_policies--name", "enforcement": "provider-schema", "group": "tgw_security.active_east_west_service_policies.service_policies:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:tgw_security:active_east_west_service_policies:service_policies", "type": "requires"}], "schema_path": ["tgw_security", "active_east_west_service_policies", "service_policies"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Active service policies for the east-west proxy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tgw_security.active_east_west_service_policies

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [tgw_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/)
- tgw_security.active_east_west_service_policies

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Active service policies for the east-west proxy.

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
active_east_west_service_policies {
  # Configure direct properties listed below.
}
```

## Direct properties

- [service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/service_policies/): complete subsection reference.

## Next pages

- [tgw_security.active_east_west_service_policies.service_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/active_east_west_service_policies/service_policies/)
- [tgw_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/tgw_security/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
