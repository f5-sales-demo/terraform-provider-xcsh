---
page_title: "private_connectivity"
subcategory: ""
description: "private_connectivity for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 2194, "body_sha256": "sha256:4c899532943ec6e7f22ef9ece6c781fe3e4d22b648ae6be1c0416a5c66697441", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:private_connectivity:cloud_link", "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity:inside", "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity:outside"], "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:private_connectivity", "parent_id": "xcsh-docs:resources:aws_tgw_site:reference", "path": "documentation/resources/aws_tgw_site/properties/private_connectivity/index.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["private_connectivity"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/private_connectivity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "private_connectivity for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# private_connectivity

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
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

- [cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/cloud_link/): complete subsection reference.

- [inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/inside/): complete subsection reference.

- [outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/outside/): complete subsection reference.

## Next pages

- [private_connectivity.cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/cloud_link/)
- [private_connectivity.inside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/inside/)
- [private_connectivity.outside](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/private_connectivity/outside/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
