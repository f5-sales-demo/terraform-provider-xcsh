---
page_title: "private_connectivity"
subcategory: "Infrastructure"
description: "private_connectivity for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1785, "body_sha256": "sha256:89c1362d2b5fd78ed792b6dc71ca8fed6b13af87c0189349776efcda83cdb80e", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:private_connectivity", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:private_connectivity:cloud_link", "xcsh-docs:resources:aws_vpc_site:properties:private_connectivity:inside", "xcsh-docs:resources:aws_vpc_site:properties:private_connectivity:outside"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:private_connectivity", "parent_id": "xcsh-docs:resources:aws_vpc_site:reference", "path": "docs/guides/resources--aws_vpc_site--properties--private_connectivity.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connectivity for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# private_connectivity

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
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

- [cloud_link](resources--aws_vpc_site--properties--private_connectivity--cloud_link.md): complete subsection reference.

- [inside](resources--aws_vpc_site--properties--private_connectivity--inside.md): complete subsection reference.

- [outside](resources--aws_vpc_site--properties--private_connectivity--outside.md): complete subsection reference.

## Next pages

- [private_connectivity.cloud_link](resources--aws_vpc_site--properties--private_connectivity--cloud_link.md)
- [private_connectivity.inside](resources--aws_vpc_site--properties--private_connectivity--inside.md)
- [private_connectivity.outside](resources--aws_vpc_site--properties--private_connectivity--outside.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
