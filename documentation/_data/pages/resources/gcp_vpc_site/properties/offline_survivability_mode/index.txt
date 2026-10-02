---
page_title: "offline_survivability_mode"
subcategory: "Infrastructure"
description: "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is"
xcsh_docs: {"aliases": ["offline survivability mode"], "body_bytes": 3119, "body_sha256": "sha256:16e9b77f3f8409c1e8c7ef53424b7fad859a5c751aa5af4da0afc007a90de87e", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "xcsh-docs:resources:gcp_vpc_site:properties:offline_survivability_mode:no_offline_survivability_mode"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:offline_survivability_mode", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "documentation/resources/gcp_vpc_site/properties/offline_survivability_mode/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0302011120033022-2021220221211032-3112222301002021-1021012200222002-0322233013032322-0132210231022013-1300001101112021-3332323033020031", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "offline_survivability_mode:ConflictingObjectAttributes:enable_offline_survivability_mode,no_offline_survivability_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "offline_survivability_mode:ConflictingObjectAttributes:enable_offline_survivability_mode,no_offline_survivability_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:offline_survivability_mode:no_offline_survivability_mode", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["offline_survivability_mode"], "schema_version": 1, "sections": [{"aliases": ["enable offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "enable_offline_survivability_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["no offline survivability mode"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:offline_survivability_mode:no_offline_survivability_mode", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["offline_survivability_mode", "no_offline_survivability_mode"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/offline_survivability_mode/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# offline_survivability_mode

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- offline_survivability_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/offline_survivability_mode/enable_offline_survivability_mode/): complete subsection reference.

- [no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/offline_survivability_mode/no_offline_survivability_mode/): complete subsection reference.

## Next pages

- [offline_survivability_mode.enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/offline_survivability_mode/enable_offline_survivability_mode/)
- [offline_survivability_mode.no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/offline_survivability_mode/no_offline_survivability_mode/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
