---
page_title: "disable_pfs"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["disable pfs"], "body_bytes": 1134, "body_sha256": "sha256:c5143e465e2fd54c730129a6a971949cdbc1d98fe083476957e24fe58d4f2192", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:ike_phase2_profile:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike_phase2_profile:properties:disable_pfs", "parent_id": "xcsh-docs:resources:ike_phase2_profile:reference", "path": "documentation/resources/ike_phase2_profile/properties/disable_pfs/index.md", "product": "distributed-cloud", "provider_name": "ike_phase2_profile", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3031211031320220-0013203313003321-1310032113330010-1101300130320003-1211220101113133-1300032010002332-2012121201003130-0332131300131013", "registry_path": "docs/guides/resources--ike_phase2_profile--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["disable_pfs"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike_phase2_profile/properties/disable_pfs/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ike_phase2_profileCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# disable_pfs

Breadcrumbs:

- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- disable_pfs

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable pfs.

Upstream description:

This can be used for messages where no values are needed.

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
disable_pfs = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/properties/)
- [xcsh_ike_phase2_profile](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike_phase2_profile/)
