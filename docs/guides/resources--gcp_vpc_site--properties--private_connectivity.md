---
page_title: "private_connectivity"
subcategory: "Infrastructure"
description: "private_connectivity for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1686, "body_sha256": "sha256:886bc94242d49faae90069f423e75fb8736370989221acf2e03a1ffefcb8b17d", "canonical_id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:cloud_link", "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:inside", "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity:outside"], "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:private_connectivity", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "docs/guides/resources--gcp_vpc_site--properties--private_connectivity.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connectivity for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# private_connectivity

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
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

- [cloud_link](resources--gcp_vpc_site--properties--private_connectivity--cloud_link.md): complete subsection reference.

- [inside](resources--gcp_vpc_site--properties--private_connectivity--inside.md): complete subsection reference.

- [outside](resources--gcp_vpc_site--properties--private_connectivity--outside.md): complete subsection reference.

## Next pages

- [private_connectivity.cloud_link](resources--gcp_vpc_site--properties--private_connectivity--cloud_link.md)
- [private_connectivity.inside](resources--gcp_vpc_site--properties--private_connectivity--inside.md)
- [private_connectivity.outside](resources--gcp_vpc_site--properties--private_connectivity--outside.md)
- [Property reference](resources--gcp_vpc_site--reference.md)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md)
