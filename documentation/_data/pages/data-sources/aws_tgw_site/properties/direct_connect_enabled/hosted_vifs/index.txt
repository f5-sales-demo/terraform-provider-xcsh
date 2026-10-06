---
page_title: "direct_connect_enabled.hosted_vifs"
subcategory: ""
description: "AWS Direct Connect Hosted VIF Configuration."
xcsh_docs: {"aliases": ["direct connect enabled hosted vifs"], "body_bytes": 1637, "body_sha256": "sha256:31f828efe69f1e51e4119b0be6324ef9c08a56cf47a55ab14015f3fcf4c5825c", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled", "path": "documentation/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013", "registry_path": "docs/guides/data-sources--aws_tgw_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs"], "schema_version": 1, "sections": [{"aliases": ["direct connect enabled hosted vifs site registration over direct connect"], "anchor": "section", "description": "CloudLink ADN Network Config.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_direct_connect"], "syntax": "attribute", "type": "object"}, {"aliases": ["direct connect enabled hosted vifs site registration over internet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_internet"], "syntax": "attribute", "type": "object"}, {"aliases": ["direct connect enabled hosted vifs vif list"], "anchor": "section", "description": "List of Hosted VIF Config.", "document_id": "xcsh-docs:data-sources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "AWS Direct Connect Hosted VIF Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/)
- direct_connect_enabled.hosted_vifs

<a id="section"></a>

Type: `"single"`. Computed.

AWS Direct Connect Hosted VIF Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

## Direct properties

- [site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/): complete subsection reference.

- [site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/): complete subsection reference.

- [vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/): complete subsection reference.
