---
page_title: "private_connectivity"
subcategory: "Infrastructure"
description: "Private Connect Configuration."
xcsh_docs: {"aliases": ["private connectivity"], "body_bytes": 2293, "body_sha256": "sha256:a9b18a1b1c1c0cca0165f63d89eb4dae144ced77f6b1daa5031b2693769327e6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:cloud_link", "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:inside", "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:outside"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "documentation/resources/gcp_vpc_site/properties/private_connectivity/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0223011112231012-1302132230002202-0010012111321332-1313123023032111-0101211323012012-2102120320023221-2221022121130120-3120021033132133", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "private_connectivity:ConflictingObjectAttributes:inside,outside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:inside", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "private_connectivity:ConflictingObjectAttributes:inside,outside", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:outside", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "sections": [{"aliases": ["cloud link"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:cloud_link", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-private_connectivity--cloud_link--name", "enforcement": "provider-schema", "group": "private_connectivity.cloud_link:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:cloud_link", "type": "requires"}], "schema_path": ["private_connectivity", "cloud_link"], "syntax": "block", "type": "object"}, {"aliases": ["inside"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:inside", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "inside"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:outside", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["private_connectivity", "outside"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Private Connect Configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- private_connectivity

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside",
    "outside")}
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
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

Terraform syntax:

```terraform
private_connectivity {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/private_connectivity/cloud_link/): complete subsection reference.

- [inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/private_connectivity/inside/): complete subsection reference.

- [outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/private_connectivity/outside/): complete subsection reference.

## Next pages

- [private_connectivity.cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/private_connectivity/cloud_link/)
- [private_connectivity.inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/private_connectivity/inside/)
- [private_connectivity.outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/private_connectivity/outside/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
