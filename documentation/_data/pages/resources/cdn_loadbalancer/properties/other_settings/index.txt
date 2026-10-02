---
page_title: "other_settings"
subcategory: "Load Balancing"
description: "Other Settings."
xcsh_docs: {"aliases": ["other settings"], "body_bytes": 2278, "body_sha256": "sha256:59ed201b94507e04dd4f2f5275f6dc816d16bb46481ba3de0543d5655d81937f", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:header_options", "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:reference", "path": "documentation/resources/cdn_loadbalancer/properties/other_settings/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["other_settings"], "schema_version": 1, "sections": [{"aliases": ["add location"], "anchor": "schema-other_settings--add_location", "description": "X-example: true Appends header x-F5 Distributed Cloud-location = <RE-site-name> in responses.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["other_settings", "add_location"], "syntax": "attribute", "type": "bool"}, {"aliases": ["header options"], "anchor": "section", "description": "This defines various OPTIONS related to request/response headers.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:header_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "header_options"], "syntax": "block", "type": "object"}, {"aliases": ["logging options"], "anchor": "section", "description": "This defines various OPTIONS related to logging.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:other_settings:logging_options", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["other_settings", "logging_options"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/other_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Other Settings.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# other_settings

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- other_settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for other settings.

Upstream description:

Other Settings.

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
other_settings {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-other_settings--add_location"></a>

### add_location property

Type: `"bool"`. Optional.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.

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

- [header_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/header_options/): complete subsection reference.

- [logging_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/logging_options/): complete subsection reference.

## Next pages

- [other_settings.header_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/header_options/)
- [other_settings.logging_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/other_settings/logging_options/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
