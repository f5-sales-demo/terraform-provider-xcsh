---
page_title: "other_settings.logging_options"
subcategory: "Load Balancing"
description: "This defines various OPTIONS related to logging."
xcsh_docs: {"aliases": ["other settings logging options"], "body_bytes": 2091, "body_sha256": "sha256:e317c640159fe6ad161f502b9a91e4d518290c145128b6c09f608b9fb800ff83", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings", "path": "documentation/resources/cdn_loadbalancer/properties/other_settings/logging_options/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["other_settings", "logging_options"], "schema_version": 1, "sections": [{"aliases": ["client log options"], "anchor": "section", "description": "List of headers to Log.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:client_log_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options", "client_log_options"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "origin log options", "origin servers", "upstream servers"], "anchor": "section", "description": "List of headers to Log.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options:origin_log_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options", "origin_log_options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/other_settings/logging_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This defines various OPTIONS related to logging.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings.logging_options

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [other_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/)
- other_settings.logging_options

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS related to logging.

Upstream description:

This defines various OPTIONS related to logging.

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
logging_options {
  # Configure direct properties listed below.
}
```

## Direct properties

- [client_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/logging_options/client_log_options/): complete subsection reference.

- [origin_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/logging_options/origin_log_options/): complete subsection reference.

## Next pages

- [other_settings.logging_options.client_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/logging_options/client_log_options/)
- [other_settings.logging_options.origin_log_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/logging_options/origin_log_options/)
- [other_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
