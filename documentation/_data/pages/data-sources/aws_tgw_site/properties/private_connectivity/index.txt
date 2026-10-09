---
page_title: "private_connectivity"
subcategory: ""
description: "Private Connect Configuration."
xcsh_docs: {"aliases": ["private connectivity"], "body_bytes": 1323, "body_sha256": "sha256:880fde05c67f2254df747f180b400917f6c117a738079e2a4b1af27290c7a069", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:cloud_link", "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:inside", "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:outside"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "documentation/data-sources/aws_tgw_site/properties/private_connectivity/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "sections": [{"aliases": ["private connectivity cloud link"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:cloud_link", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["private_connectivity", "cloud_link"], "syntax": "attribute", "type": "object"}, {"aliases": ["private connectivity inside"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:inside", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "inside"], "syntax": "attribute", "type": "object"}, {"aliases": ["private connectivity outside"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:private_connectivity:outside", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "outside"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Private Connect Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- private_connectivity

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for private connectivity.

Additional upstream details:

Private Connect Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

## Direct properties

- [cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/cloud_link/): complete subsection reference.

- [inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/inside/): complete subsection reference.

- [outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/private_connectivity/outside/): complete subsection reference.
