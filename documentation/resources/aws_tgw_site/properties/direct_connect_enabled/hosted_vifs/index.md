---
page_title: "direct_connect_enabled.hosted_vifs"
subcategory: ""
description: "AWS Direct Connect Hosted VIF Configuration."
xcsh_docs: {"aliases": ["direct connect enabled hosted vifs"], "body_bytes": 2874, "body_sha256": "sha256:4237af26350674b264da76eef2cbf075f78b5b07a63352a32c1cc58202af094f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled", "path": "documentation/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs:ConflictingObjectAttributes:site_registration_over_direct_connect,site_registration_over_internet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs:ConflictingObjectAttributes:site_registration_over_direct_connect,site_registration_over_internet", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs"], "schema_version": 1, "sections": [{"aliases": ["site registration over direct connect"], "anchor": "section", "description": "CloudLink ADN Network Config.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-direct_connect_enabled--hosted_vifs--site_registration_over_direct_connect--cloudlink_network_name", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect:RequiredObjectAttributes:cloudlink_network_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_direct_connect", "type": "requires"}], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_direct_connect"], "syntax": "block", "type": "object"}, {"aliases": ["site registration over internet"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:site_registration_over_internet", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "site_registration_over_internet"], "syntax": "attribute", "type": "object"}, {"aliases": ["vif list"], "anchor": "section", "description": "List of Hosted VIF Config.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--other_region", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:ConflictingListObjectAttributes:other_region,same_as_site_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:ConflictingListObjectAttributes:other_region,same_as_site_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list:same_as_site_region", "type": "conflicts"}, {"anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:RequiredListObjectAttributes:vif_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "type": "requires"}], "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "AWS Direct Connect Hosted VIF Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_direct_connect/)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/site_registration_over_internet/)
- [direct_connect_enabled.hosted_vifs.vif_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
