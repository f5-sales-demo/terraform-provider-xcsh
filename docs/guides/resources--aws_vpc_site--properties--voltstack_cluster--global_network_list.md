---
page_title: "voltstack_cluster.global_network_list"
subcategory: "Infrastructure"
description: "voltstack_cluster.global_network_list for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1557, "body_sha256": "sha256:4dcea65ba69a53b31f98dfdaacf7e058cbee4b2e6edd77540dd3d95df905a079", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:global_network_list", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:global_network_list:global_network_connections"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster:global_network_list", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:voltstack_cluster", "path": "docs/guides/resources--aws_vpc_site--properties--voltstack_cluster--global_network_list.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["voltstack_cluster", "global_network_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/voltstack_cluster/global_network_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "voltstack_cluster.global_network_list for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.global_network_list

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md)
- voltstack_cluster.global_network_list

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

- [global_network_connections](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections.md): complete subsection reference.

## Next pages

- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--properties--voltstack_cluster--global_network_list--global_network_connections.md)
- [voltstack_cluster](resources--aws_vpc_site--properties--voltstack_cluster.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
