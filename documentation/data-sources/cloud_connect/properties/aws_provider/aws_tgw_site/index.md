---
page_title: "aws_provider.aws_tgw_site"
subcategory: ""
description: "Cloud Connect AWS TGW Site Type."
xcsh_docs: {"aliases": ["aws provider aws tgw site"], "body_bytes": 2168, "body_sha256": "sha256:4ca1cf9c0f291553a98a6a1a686ad34835c2a246f9723a9eb1e5f2abc2c879b0", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:site", "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site", "parent_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider", "path": "documentation/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/index.md", "product": "distributed-cloud", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1210201012310011-1202111032113023-2120321323312011-1101121121031011-0113230133300122-3231021212301121-1230133321330321-3230000102301202", "registry_path": "docs/guides/data-sources--cloud_connect--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_provider", "aws_tgw_site"], "schema_version": 1, "sections": [{"aliases": ["cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "cred"], "syntax": "attribute", "type": "object"}, {"aliases": ["site"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["vpc attachments"], "anchor": "section", "description": "Configuration parameter for vpc attachments.", "document_id": "xcsh-docs:data-sources:cloud_connect:properties:aws_provider:aws_tgw_site:vpc_attachments", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["aws_provider", "aws_tgw_site", "vpc_attachments"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Cloud Connect AWS TGW Site Type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_provider.aws_tgw_site

Breadcrumbs:

- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/)
- aws_provider.aws_tgw_site

<a id="section"></a>

Type: `"single"`. Computed.

AWS TGW Site Type. Cloud Connect AWS TGW Site Type.

Upstream description:

Cloud Connect AWS TGW Site Type.

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

## Direct properties

- [cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/): complete subsection reference.

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/site/): complete subsection reference.

- [vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/): complete subsection reference.

## Next pages

- [aws_provider.aws_tgw_site.cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/cred/)
- [aws_provider.aws_tgw_site.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/site/)
- [aws_provider.aws_tgw_site.vpc_attachments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/aws_tgw_site/vpc_attachments/)
- [aws_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/properties/aws_provider/)
- [xcsh_cloud_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_connect/)
