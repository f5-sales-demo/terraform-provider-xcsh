---
page_title: "vn_config.global_network_list"
subcategory: ""
description: "List of global network connections."
xcsh_docs: {"aliases": ["vn config global network list"], "body_bytes": 1840, "body_sha256": "sha256:4d1081431fc0318626cc58b13653eded06a90646616bfbb02bc7c8f2da21e93d", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config", "path": "documentation/resources/aws_tgw_site/properties/vn_config/global_network_list/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1311000133311331-0133032032130101-2311221330311100-0022033203212131-3312123023132201-3130101000123111-0013130131023303-2323231132300030", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.global_network_list:RequiredObjectAttributes:global_network_connections", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["vn_config", "global_network_list"], "schema_version": 1, "sections": [{"aliases": ["global network connections"], "anchor": "section", "description": "Global network connections.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.global_network_list.global_network_connections:ConflictingListObjectAttributes:sli_to_global_dr,slo_to_global_dr", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:sli_to_global_dr", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "vn_config.global_network_list.global_network_connections:ConflictingListObjectAttributes:sli_to_global_dr,slo_to_global_dr", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:vn_config:global_network_list:global_network_connections:slo_to_global_dr", "type": "conflicts"}], "schema_path": ["vn_config", "global_network_list", "global_network_connections"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/vn_config/global_network_list/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of global network connections.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# vn_config.global_network_list

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- vn_config.global_network_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/): complete subsection reference.

## Next pages

- [vn_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/global_network_list/global_network_connections/)
- [vn_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/vn_config/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
