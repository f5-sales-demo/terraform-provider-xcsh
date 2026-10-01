---
page_title: "voltstack_cluster.site_local_network"
subcategory: "Infrastructure"
description: "voltstack_cluster.site_local_network for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 2981, "body_sha256": "sha256:52074bbb69722d131a46066b73c7634989fd9a40f1d211a8a1391c6a1661bfda", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:existing_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network", "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network:new_network_autogenerate"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster:site_local_network", "parent_id": "xcsh-docs:resources:gcp_vpc_site:properties:voltstack_cluster", "path": "documentation/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/index.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["voltstack_cluster", "site_local_network"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.site_local_network for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.site_local_network

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- voltstack_cluster.site_local_network

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
site_local_network {
  # Configure direct properties listed below.
}
```

## Direct properties

- [existing_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/existing_network/): complete subsection reference.

- [new_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network/): complete subsection reference.

- [new_network_autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network_autogenerate/): complete subsection reference.

## Next pages

- [voltstack_cluster.site_local_network.existing_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/existing_network/)
- [voltstack_cluster.site_local_network.new_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network/)
- [voltstack_cluster.site_local_network.new_network_autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/site_local_network/new_network_autogenerate/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/voltstack_cluster/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
