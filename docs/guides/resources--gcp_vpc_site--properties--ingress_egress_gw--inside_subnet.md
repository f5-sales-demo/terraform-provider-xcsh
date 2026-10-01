---
page_title: "ingress_egress_gw.inside_subnet"
subcategory: "Infrastructure"
description: "ingress_egress_gw.inside_subnet for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1823, "body_sha256": "sha256:0233a2f428812e6f11c7bc2fedff295885da5f2e00f1c81acbbce27197116c92", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:existing_subnet", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet:new_subnet"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:inside_subnet", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/resources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "inside_subnet"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/inside_subnet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.inside_subnet for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.inside_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.inside_subnet

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_subnet](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--existing_subnet.md): complete subsection reference.

- [new_subnet](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--new_subnet.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.inside_subnet.existing_subnet](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--existing_subnet.md)
- [ingress_egress_gw.inside_subnet.new_subnet](resources--gcp_vpc_site--properties--ingress_egress_gw--inside_subnet--new_subnet.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
