---
page_title: "gre"
subcategory: ""
description: "External Connector with GRE tunnel."
xcsh_docs: {"aliases": ["gre"], "body_bytes": 1665, "body_sha256": "sha256:33a32a2cef7be47e4201c97fa9c23e01e4fdc8caa8ca9a64576b4cfdf3b5f11f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:external_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:external_connector:properties:gre", "parent_id": "xcsh-docs:data-sources:external_connector:reference", "path": "documentation/data-sources/external_connector/properties/gre/index.md", "product": "distributed-cloud", "provider_name": "external_connector", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021", "registry_path": "docs/guides/data-sources--external_connector--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["gre"], "schema_version": 1, "sections": [{"aliases": ["gre parameters"], "anchor": "section", "description": "GRE configuration parameters required for GRE Connection type.", "document_id": "xcsh-docs:data-sources:external_connector:properties:gre:gre_parameters", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["gre", "gre_parameters"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/external_connector/properties/gre/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "External Connector with GRE tunnel.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["external_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gre

Breadcrumbs:

- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- gre

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: gre, ipsec\] GRE. External Connector with GRE tunnel.

Upstream description:

External Connector with GRE tunnel.

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

OneOf alternatives in this subsection:

- [gre](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/#section)
- [ipsec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/ipsec/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/): complete subsection reference.

## Next pages

- [gre.gre_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/gre/gre_parameters/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/properties/)
- [xcsh_external_connector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/external_connector/)
