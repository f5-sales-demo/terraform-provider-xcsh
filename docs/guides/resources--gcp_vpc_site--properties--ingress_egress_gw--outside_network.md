---
page_title: "ingress_egress_gw.outside_network"
subcategory: "Infrastructure"
description: "ingress_egress_gw.outside_network for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2394, "body_sha256": "sha256:96e9040ba810ca15092ffa52c7bd4c66d567b1eb3e132c5dd371f03dfc5d81eb", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:existing_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network:new_network_autogenerate"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw:outside_network", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_egress_gw", "path": "docs/guides/resources--gcp_vpc_site--properties--ingress_egress_gw--outside_network.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "outside_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_egress_gw/outside_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.outside_network for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.outside_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- ingress_egress_gw.outside_network

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
outside_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_network](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--existing_network.md): complete subsection reference.

- [new_network](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--new_network.md): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--new_network_autogenerate.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.outside_network.existing_network](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--existing_network.md)
- [ingress_egress_gw.outside_network.new_network](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--new_network.md)
- [ingress_egress_gw.outside_network.new_network_autogenerate](resources--gcp_vpc_site--properties--ingress_egress_gw--outside_network--new_network_autogenerate.md)
- [ingress_egress_gw](resources--gcp_vpc_site--properties--ingress_egress_gw.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
