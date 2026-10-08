---
page_title: "direct_connect_enabled.hosted_vifs"
subcategory: ""
description: "AWS Direct Connect Hosted VIF Configuration."
xcsh_docs: {"aliases": ["direct connect enabled hosted vifs"], "body_bytes": 1987, "body_sha256": "sha256:6b22d9c551da06a68b9526e260f423c423015e9f42b78f3b99b1280bd780589b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "path": "documentation/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs:ConflictingObjectAttributes:site_registration_over_direct_connect,site_registration_over_internet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs:ConflictingObjectAttributes:site_registration_over_direct_connect,site_registration_over_internet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs"], "schema_version": 1, "sections": [{"aliases": ["direct connect enabled hosted vifs site registration over direct connect"], "anchor": "section", "description": "CloudLink ADN Network Config.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect:RequiredObjectAttributes:cloudlink_network_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "type": "requires"}], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_direct_connect"], "syntax": "block", "type": "object"}, {"aliases": ["direct connect enabled hosted vifs site registration over internet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_internet"], "syntax": "attribute", "type": "object"}, {"aliases": ["direct connect enabled hosted vifs vif list"], "anchor": "section", "description": "List of Hosted VIF Config.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--other_region", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:ConflictingListObjectAttributes:other_region,same_as_site_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:ConflictingListObjectAttributes:other_region,same_as_site_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list:same_as_site_region", "type": "conflicts"}, {"anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:RequiredListObjectAttributes:vif_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "type": "requires"}], "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "AWS Direct Connect Hosted VIF Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/)
- direct_connect_enabled.hosted_vifs

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

AWS Direct Connect Hosted VIF Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_registration_over_direct_connect",
    "site_registration_over_internet")}
```

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

Terraform syntax:

```terraform
hosted_vifs {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/): complete subsection reference.

- [site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/): complete subsection reference.

- [vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/): complete subsection reference.
