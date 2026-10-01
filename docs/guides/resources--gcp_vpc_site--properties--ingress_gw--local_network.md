---
page_title: "ingress_gw.local_network"
subcategory: "Infrastructure"
description: "ingress_gw.local_network for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2265, "body_sha256": "sha256:5b7f5c261c65b724474156f23e683d580410620ddd8df3feb9b9be9b983d7941", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:existing_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network", "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network:new_network_autogenerate"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw:local_network", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:ingress_gw", "path": "docs/guides/resources--gcp_vpc_site--properties--ingress_gw--local_network.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw", "local_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/ingress_gw/local_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw.local_network for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.local_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [ingress_gw](resources--gcp_vpc_site--properties--ingress_gw.md)
- ingress_gw.local_network

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
local_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_network](resources--gcp_vpc_site--properties--ingress_gw--local_network--existing_network.md): complete subsection reference.

- [new_network](resources--gcp_vpc_site--properties--ingress_gw--local_network--new_network.md): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--properties--ingress_gw--local_network--new_network_autogenerate.md): complete subsection reference.

## Next pages

- [ingress_gw.local_network.existing_network](resources--gcp_vpc_site--properties--ingress_gw--local_network--existing_network.md)
- [ingress_gw.local_network.new_network](resources--gcp_vpc_site--properties--ingress_gw--local_network--new_network.md)
- [ingress_gw.local_network.new_network_autogenerate](resources--gcp_vpc_site--properties--ingress_gw--local_network--new_network_autogenerate.md)
- [ingress_gw](resources--gcp_vpc_site--properties--ingress_gw.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
